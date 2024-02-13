// Package pricer contains pricing logic for RFQ relayer quotes.
package pricer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// FeePricer is the interface for the fee pricer.
type FeePricer interface {
	// Start starts the fee pricer.
	Start(ctx context.Context)
	// GetOriginFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error)
	// GetDestinationFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error)
	// GetGasPrice returns the gas price for a given chainID in native units.
	GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error)
}

type feePricer struct {
	// config is the relayer config.
	config relconfig.Config
	// gasPriceCache maps chainID -> gas price
	gasPriceCache *ttlcache.Cache[uint32, *big.Int]
	// originGasEstimateCache maps chainID -> gas estimate
	originGasEstimateCache *ttlcache.Cache[uint32, uint64]
	// destGasEstimateCache maps chainID -> gas estimate
	destGasEstimateCache *ttlcache.Cache[uint32, uint64]
	// tokenPriceCache maps token name -> token price
	tokenPriceCache *ttlcache.Cache[string, float64]
	// clientFetcher is used to fetch clients.
	clientFetcher submitter.ClientFetcher
	// handler is the metrics handler.
	handler metrics.Handler
	// priceFetcher is used to fetch prices from coingecko.
	priceFetcher CoingeckoPriceFetcher
	// arbitrumSDK is the SDK for interacting with Arbitrum.
	arbitrumSDK arbitrum.SDK
	// feeSigner is a signer used for gas estimation.
	feeSigner signer.Signer
	// bridges is a map of chain ID -> bridge contract (used for gas estimation).
	bridges map[uint32]*fastbridge.FastBridgeRef
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(ctx context.Context, config relconfig.Config, clientFetcher submitter.ClientFetcher, priceFetcher CoingeckoPriceFetcher, feeSigner signer.Signer, handler metrics.Handler) (FeePricer, error) {
	// setup caches
	gasPriceCache := ttlcache.New[uint32, *big.Int](
		ttlcache.WithTTL[uint32, *big.Int](time.Second*time.Duration(config.GetFeePricer().GasPriceCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[uint32, *big.Int](),
	)
	tokenPriceCache := ttlcache.New[string, float64](
		ttlcache.WithTTL[string, float64](time.Second*time.Duration(config.GetFeePricer().TokenPriceCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[string, float64](),
	)
	originGasEstimateCache := ttlcache.New[uint32, uint64](
		ttlcache.WithTTL[uint32, uint64](time.Second*time.Duration(config.GetFeePricer().GasEstimateCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[uint32, uint64](),
	)
	destGasEstimateCache := ttlcache.New[uint32, uint64](
		ttlcache.WithTTL[uint32, uint64](time.Second*time.Duration(config.GetFeePricer().GasEstimateCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[uint32, uint64](),
	)

	// setup contracts for dynamic gas estimates
	bridges := map[uint32]*fastbridge.FastBridgeRef{}
	for chainID, chainConfig := range config.Chains {
		client, err := clientFetcher.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		bridges[uint32(chainID)], err = fastbridge.NewFastBridgeRef(common.HexToAddress(chainConfig.Bridge), client)
		if err != nil {
			return nil, fmt.Errorf("could not get bridge contract: %w", err)
		}
	}

	return &feePricer{
		config:                 config,
		gasPriceCache:          gasPriceCache,
		originGasEstimateCache: originGasEstimateCache,
		destGasEstimateCache:   destGasEstimateCache,
		tokenPriceCache:        tokenPriceCache,
		clientFetcher:          clientFetcher,
		handler:                handler,
		priceFetcher:           priceFetcher,
		feeSigner:              feeSigner,
		bridges:                bridges,
	}, nil
}

func (f *feePricer) Start(ctx context.Context) {
	g, _ := errgroup.WithContext(ctx)

	// start the TTL caches
	g.Go(func() error {
		f.gasPriceCache.Start()
		return nil
	})
	g.Go(func() error {
		f.originGasEstimateCache.Start()
		return nil
	})
	g.Go(func() error {
		f.destGasEstimateCache.Start()
		return nil
	})
	g.Go(func() error {
		f.tokenPriceCache.Start()
		return nil
	})
}

var nativeDecimalsFactor = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18)), nil)

func (f *feePricer) GetOriginFee(parentCtx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error) {
	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getOriginFee", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(origin)),
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("use_multiplier", useMultiplier),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Calculate the origin fee
	gasEstimate, err := f.getGasEstimate(ctx, origin, true)
	if err != nil {
		return nil, fmt.Errorf("could not get origin gas estimate: %w", err)
	}
	fee, err := f.getFee(ctx, origin, destination, gasEstimate, denomToken, useMultiplier)
	if err != nil {
		return nil, err
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(origin, true)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, uint64(l1GasEstimate), denomToken, useMultiplier)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("origin_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetDestinationFee(parentCtx context.Context, _, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error) {
	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getDestinationFee", trace.WithAttributes(
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("use_multiplier", useMultiplier),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Calculate the destination fee
	gasEstimate, err := f.getGasEstimate(ctx, destination, false)
	if err != nil {
		return nil, fmt.Errorf("could not get dest gas estimate: %w", err)
	}
	fee, err := f.getFee(ctx, destination, destination, gasEstimate, denomToken, useMultiplier)
	if err != nil {
		return nil, err
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(destination, false)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, uint64(l1GasEstimate), denomToken, useMultiplier)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("destination_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetTotalFee(parentCtx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (_ *big.Int, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getTotalFee", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(origin)),
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("use_multiplier", useMultiplier),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	originFee, err := f.GetOriginFee(ctx, origin, destination, denomToken, useMultiplier)
	if err != nil {
		span.AddEvent("could not get origin fee", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
		return nil, err
	}
	destFee, err := f.GetDestinationFee(ctx, origin, destination, denomToken, useMultiplier)
	if err != nil {
		span.AddEvent("could not get destination fee", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
		return nil, err
	}
	totalFee := new(big.Int).Add(originFee, destFee)
	span.SetAttributes(
		attribute.String("origin_fee", originFee.String()),
		attribute.String("dest_fee", destFee.String()),
		attribute.String("total_fee", totalFee.String()),
	)
	return totalFee, nil
}

func (f *feePricer) getFee(parentCtx context.Context, gasChain, denomChain uint32, gasEstimate uint64, denomToken string, useMultiplier bool) (_ *big.Int, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getFee", trace.WithAttributes(
		attribute.Int("gas_chain", int(gasChain)),
		attribute.Int("denom_chain", int(denomChain)),
		attribute.Int("gas_estimate", int(gasEstimate)),
		attribute.String("denom_token", denomToken),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	gasPrice, err := f.GetGasPrice(ctx, gasChain)
	if err != nil {
		return nil, err
	}
	nativeToken, err := f.config.GetNativeToken(int(gasChain))
	if err != nil {
		return nil, err
	}
	nativeTokenPrice, err := f.getTokenPrice(ctx, nativeToken)
	if err != nil {
		return nil, err
	}
	denomTokenPrice, err := f.getTokenPrice(ctx, denomToken)
	if err != nil {
		return nil, err
	}
	denomTokenDecimals, err := f.config.GetTokenDecimals(denomChain, denomToken)
	if err != nil {
		return nil, err
	}
	denomDecimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(denomTokenDecimals)), nil)

	// Compute the fee.
	var feeDenom *big.Float
	feeWei := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), new(big.Float).SetFloat64(float64(gasEstimate)))
	if denomToken == nativeToken {
		// Denomination token is native token, so no need for unit conversion.
		feeDenom = feeWei
	} else {
		// Convert the fee from ETH to denomToken terms.
		feeEth := new(big.Float).Quo(feeWei, new(big.Float).SetInt(nativeDecimalsFactor))
		feeUSD := new(big.Float).Mul(feeEth, new(big.Float).SetFloat64(nativeTokenPrice))
		feeUSDC := new(big.Float).Mul(feeUSD, new(big.Float).SetFloat64(denomTokenPrice))
		feeDenom = new(big.Float).Mul(feeUSDC, new(big.Float).SetInt(denomDecimalsFactor))
		span.SetAttributes(
			attribute.String("fee_wei", feeWei.String()),
			attribute.String("fee_eth", feeEth.String()),
			attribute.String("fee_usd", feeUSD.String()),
			attribute.String("fee_usdc", feeUSDC.String()),
		)
	}

	multiplier := 1.
	if useMultiplier {
		multiplier, err = f.config.GetFixedFeeMultiplier(int(gasChain))
		if err != nil {
			return nil, fmt.Errorf("could not get fixed fee multiplier: %w", err)
		}
	}

	// Apply the fixed fee multiplier.
	// Note that this step rounds towards zero- we may need to apply rounding here if
	// we want to be conservative and lean towards overestimating fees.
	feeUSDCDecimalsScaled, _ := new(big.Float).Mul(feeDenom, new(big.Float).SetFloat64(multiplier)).Int(nil)
	span.SetAttributes(
		attribute.String("gas_price", gasPrice.String()),
		attribute.Float64("native_token_price", nativeTokenPrice),
		attribute.Float64("denom_token_price", denomTokenPrice),
		attribute.Int("denom_token_decimals", int(denomTokenDecimals)),
		attribute.String("fee_wei", feeWei.String()),
		attribute.String("fee_denom", feeDenom.String()),
		attribute.String("fee_usdc_decimals_scaled", feeUSDCDecimalsScaled.String()),
	)
	return feeUSDCDecimalsScaled, nil
}

// getGasPrice returns the gas price for a given chainID in native units.
func (f *feePricer) GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error) {
	// Attempt to fetch gas price from cache.
	gasPriceItem := f.gasPriceCache.Get(chainID)
	var gasPrice *big.Int
	if gasPriceItem == nil {
		// Fetch gas price from omnirpc.
		client, err := f.clientFetcher.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return nil, err
		}
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}
		gasPrice = header.BaseFee
		f.gasPriceCache.Set(chainID, gasPrice, 0)
	} else {
		gasPrice = gasPriceItem.Value()
	}
	return gasPrice, nil
}

// getTokenPrice returns the price of a token in USD.
func (f *feePricer) getTokenPrice(ctx context.Context, token string) (price float64, err error) {
	// Attempt to fetch gas price from cache.
	tokenPriceItem := f.tokenPriceCache.Get(token)
	//nolint:nestif
	if tokenPriceItem == nil {
		// Try to get price from coingecko.
		price, err = f.priceFetcher.GetPrice(ctx, token)
		if err == nil {
			f.tokenPriceCache.Set(token, price, 0)
		} else {
			// Fallback to configured token price.
			price, err = f.getTokenPriceFromConfig(token)
			if err != nil {
				return 0, err
			}
		}
	} else {
		price = tokenPriceItem.Value()
	}
	return price, nil
}

func (f *feePricer) getTokenPriceFromConfig(token string) (float64, error) {
	for _, chainConfig := range f.config.GetChains() {
		for tokenName, tokenConfig := range chainConfig.Tokens {
			if token == tokenName {
				return tokenConfig.PriceUSD, nil
			}
		}
	}
	return 0, fmt.Errorf("could not get price for token: %s", token)
}

func (f *feePricer) getGasEstimate(parentCtx context.Context, chainID uint32, isOrigin bool) (gasEstimate uint64, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getGasEstimate", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(chainID)),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// attempt to load the gas estimate from the cache
	var cache *ttlcache.Cache[uint32, uint64]
	if isOrigin {
		cache = f.originGasEstimateCache
	} else {
		cache = f.destGasEstimateCache
	}
	gasEstimateItem := cache.Get(chainID)
	if gasEstimateItem != nil {
		gasEstimate = gasEstimateItem.Value()
		return gasEstimate, nil
	}

	// if dynamic gas estimation is enabled, attempt to get the gas estimate from the client
	dynamic, err := f.config.GetDynamicGasEstimate(int(chainID))
	if err != nil {
		return 0, fmt.Errorf("could not get dynamic gas estimate from config: %w", err)
	}
	if dynamic {
		gasEstimate, err = f.getGasEstimateFromClient(ctx, chainID, isOrigin)
		if err == nil {
			// cache the dynamic gas estimate
			cache.Set(chainID, gasEstimate, 0)
			return gasEstimate, nil
		}
		span.AddEvent("could not get gas estimate from client", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
	}

	// fall back to gas estimate from the config
	var gasEstimateRaw int
	if isOrigin {
		gasEstimateRaw, err = f.config.GetOriginGasEstimate(int(chainID))
	} else {
		gasEstimateRaw, err = f.config.GetDestGasEstimate(int(chainID))
	}
	if err != nil {
		return 0, fmt.Errorf("could not get gas estimate from config: %w", err)
	}
	gasEstimate = uint64(gasEstimateRaw)

	return gasEstimate, nil
}

func (f *feePricer) getGasEstimateFromClient(parentCtx context.Context, chainID uint32, isOrigin bool) (gasEstimate uint64, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getGasEstimateFromClient", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(chainID)),
	))

	defer func() {
		span.AddEvent("estimated_gas", trace.WithAttributes(attribute.Int64("gas", int64(gasEstimate))))
		metrics.EndSpanWithErr(span, err)
	}()

	chainClient, err := f.clientFetcher.GetClient(ctx, big.NewInt(int64(chainID)))
	if err != nil {
		return 0, fmt.Errorf("could not get client: %w", err)
	}

	// build the mock contract calls
	calls, err := f.getCalls(ctx, chainID, isOrigin)
	if err != nil {
		return 0, err
	}

	// load the gas estimator
	var gasEstimator submitter.GasEstimator
	switch submitter.GetGasEstimationMethod(&f.config.SubmitterConfig, int(chainID)) {
	case submitter.ArbitrumGasEstimation:
		if f.arbitrumSDK == nil {
			f.arbitrumSDK, err = arbitrum.NewArbitrumSDK(chainClient, arbitrum.WithMetrics(f.handler))
			if err != nil {
				return 0, fmt.Errorf("could not get arbitrum SDK: %w", err)
			}
		}
		gasEstimator = f.arbitrumSDK
	case submitter.GethGasEstimation:
		gasEstimator = chainClient
	default:
		return 0, fmt.Errorf("unknown gas estimation method")
	}

	// fetch the gas estimates
	for _, call := range calls {
		var estimate uint64
		estimate, err = gasEstimator.EstimateGas(ctx, *call)
		if err != nil {
			return 0, fmt.Errorf("could not estimate gas: %w", err)
		}
		gasEstimate += estimate
	}
	return gasEstimate, nil
}

func (f *feePricer) getCalls(parentCtx context.Context, chainID uint32, isOrigin bool) (calls []*ethereum.CallMsg, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getCalls")
	defer func() {
		span.SetAttributes(attribute.String("error", err.Error()))
		metrics.EndSpanWithErr(span, err)
	}()

	bridge, ok := f.bridges[chainID]
	if !ok {
		return calls, fmt.Errorf("could not get bridge for chain: %d", chainID)
	}
	transactor, err := f.feeSigner.GetTransactor(ctx, big.NewInt(int64(chainID)))
	if err != nil {
		return calls, fmt.Errorf("could not get transactor: %w", err)
	}

	calls = []*ethereum.CallMsg{}
	var call *ethereum.CallMsg
	if isOrigin {
		// get claim call
		call, err = getCall(transactor, bridge, claimCallType)
		if err != nil {
			return calls, fmt.Errorf("could not get claim call: %w", err)
		}
		calls = append(calls, call)

		// get prove call
		call, err = getCall(transactor, bridge, proveCallType)
		if err != nil {
			return calls, fmt.Errorf("could not get claim call: %w", err)
		}
		calls = append(calls, call)
	} else {
		// get relay call
		call, err = getCall(transactor, bridge, proveCallType)
		if err != nil {
			return calls, fmt.Errorf("could not get claim call: %w", err)
		}
		calls = append(calls, call)
	}
	return calls, nil
}
