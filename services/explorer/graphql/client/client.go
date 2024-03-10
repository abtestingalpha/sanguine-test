// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package client

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) *Client {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	BridgeTransactions     []*model.BridgeTransaction      "json:\"bridgeTransactions\" graphql:\"bridgeTransactions\""
	MessageBusTransactions []*model.MessageBusTransaction  "json:\"messageBusTransactions\" graphql:\"messageBusTransactions\""
	CountByChainID         []*model.TransactionCountResult "json:\"countByChainId\" graphql:\"countByChainId\""
	CountByTokenAddress    []*model.TokenCountResult       "json:\"countByTokenAddress\" graphql:\"countByTokenAddress\""
	AddressRanking         []*model.AddressRanking         "json:\"addressRanking\" graphql:\"addressRanking\""
	AmountStatistic        *model.ValueResult              "json:\"amountStatistic\" graphql:\"amountStatistic\""
	DailyStatisticsByChain []*model.DateResultByChain      "json:\"dailyStatisticsByChain\" graphql:\"dailyStatisticsByChain\""
	RankedChainIDsByVolume []*model.VolumeByChainID        "json:\"rankedChainIDsByVolume\" graphql:\"rankedChainIDsByVolume\""
	AddressData            *model.AddressData              "json:\"addressData\" graphql:\"addressData\""
	Leaderboard            []*model.Leaderboard            "json:\"leaderboard\" graphql:\"leaderboard\""
	GetOriginBridgeTx      *model.BridgeWatcherTx          "json:\"getOriginBridgeTx\" graphql:\"getOriginBridgeTx\""
	GetDestinationBridgeTx *model.BridgeWatcherTx          "json:\"getDestinationBridgeTx\" graphql:\"getDestinationBridgeTx\""
	GetBlockHeight         []*model.BlockHeight            "json:\"getBlockHeight\" graphql:\"getBlockHeight\""
}
type GetBridgeTransactions struct {
	Response []*struct {
		FromInfo *struct {
			ChainID            *int     "json:\"chainID\" graphql:\"chainID\""
			DestinationChainID *int     "json:\"destinationChainID\" graphql:\"destinationChainID\""
			Address            *string  "json:\"address\" graphql:\"address\""
			TxnHash            *string  "json:\"txnHash\" graphql:\"txnHash\""
			Value              *string  "json:\"value\" graphql:\"value\""
			FormattedValue     *float64 "json:\"formattedValue\" graphql:\"formattedValue\""
			USDValue           *float64 "json:\"USDValue\" graphql:\"USDValue\""
			TokenAddress       *string  "json:\"tokenAddress\" graphql:\"tokenAddress\""
			TokenSymbol        *string  "json:\"tokenSymbol\" graphql:\"tokenSymbol\""
			BlockNumber        *int     "json:\"blockNumber\" graphql:\"blockNumber\""
			Time               *int     "json:\"time\" graphql:\"time\""
			FormattedTime      *string  "json:\"formattedTime\" graphql:\"formattedTime\""
		} "json:\"fromInfo\" graphql:\"fromInfo\""
		ToInfo *struct {
			ChainID        *int     "json:\"chainID\" graphql:\"chainID\""
			Address        *string  "json:\"address\" graphql:\"address\""
			TxnHash        *string  "json:\"txnHash\" graphql:\"txnHash\""
			Value          *string  "json:\"value\" graphql:\"value\""
			FormattedValue *float64 "json:\"formattedValue\" graphql:\"formattedValue\""
			USDValue       *float64 "json:\"USDValue\" graphql:\"USDValue\""
			TokenAddress   *string  "json:\"tokenAddress\" graphql:\"tokenAddress\""
			TokenSymbol    *string  "json:\"tokenSymbol\" graphql:\"tokenSymbol\""
			BlockNumber    *int     "json:\"blockNumber\" graphql:\"blockNumber\""
			Time           *int     "json:\"time\" graphql:\"time\""
			FormattedTime  *string  "json:\"formattedTime\" graphql:\"formattedTime\""
		} "json:\"toInfo\" graphql:\"toInfo\""
		Kappa       *string "json:\"kappa\" graphql:\"kappa\""
		Pending     *bool   "json:\"pending\" graphql:\"pending\""
		SwapSuccess *bool   "json:\"swapSuccess\" graphql:\"swapSuccess\""
	} "json:\"response\" graphql:\"response\""
}
type GetCountByChainID struct {
	Response []*struct {
		Count   *int "json:\"count\" graphql:\"count\""
		ChainID *int "json:\"chainID\" graphql:\"chainID\""
	} "json:\"response\" graphql:\"response\""
}
type GetCountByTokenAddress struct {
	Response []*struct {
		ChainID      *int    "json:\"chainID\" graphql:\"chainID\""
		TokenAddress *string "json:\"tokenAddress\" graphql:\"tokenAddress\""
		Count        *int    "json:\"count\" graphql:\"count\""
	} "json:\"response\" graphql:\"response\""
}
type GetAddressRanking struct {
	Response []*struct {
		Address *string "json:\"address\" graphql:\"address\""
		Count   *int    "json:\"count\" graphql:\"count\""
	} "json:\"response\" graphql:\"response\""
}
type GetRankedChainIDsByVolume struct {
	Response []*struct {
		ChainID *int     "json:\"chainID\" graphql:\"chainID\""
		Total   *float64 "json:\"total\" graphql:\"total\""
	} "json:\"response\" graphql:\"response\""
}
type GetBlockHeight struct {
	Response []*struct {
		ChainID     *int                "json:\"chainID\" graphql:\"chainID\""
		Type        *model.ContractType "json:\"type\" graphql:\"type\""
		BlockNumber *int                "json:\"blockNumber\" graphql:\"blockNumber\""
	} "json:\"response\" graphql:\"response\""
}
type GetAmountStatistic struct {
	Response *struct {
		Value *string "json:\"value\" graphql:\"value\""
	} "json:\"response\" graphql:\"response\""
}
type GetDailyStatisticsByChain struct {
	Response []*struct {
		Date      *string  "json:\"date\" graphql:\"date\""
		Ethereum  *float64 "json:\"ethereum\" graphql:\"ethereum\""
		Optimism  *float64 "json:\"optimism\" graphql:\"optimism\""
		Cronos    *float64 "json:\"cronos\" graphql:\"cronos\""
		Bsc       *float64 "json:\"bsc\" graphql:\"bsc\""
		Polygon   *float64 "json:\"polygon\" graphql:\"polygon\""
		Fantom    *float64 "json:\"fantom\" graphql:\"fantom\""
		Boba      *float64 "json:\"boba\" graphql:\"boba\""
		Metis     *float64 "json:\"metis\" graphql:\"metis\""
		Moonbeam  *float64 "json:\"moonbeam\" graphql:\"moonbeam\""
		Moonriver *float64 "json:\"moonriver\" graphql:\"moonriver\""
		Klaytn    *float64 "json:\"klaytn\" graphql:\"klaytn\""
		Arbitrum  *float64 "json:\"arbitrum\" graphql:\"arbitrum\""
		Avalanche *float64 "json:\"avalanche\" graphql:\"avalanche\""
		Dfk       *float64 "json:\"dfk\" graphql:\"dfk\""
		Aurora    *float64 "json:\"aurora\" graphql:\"aurora\""
		Harmony   *float64 "json:\"harmony\" graphql:\"harmony\""
		Canto     *float64 "json:\"canto\" graphql:\"canto\""
		Dogechain *float64 "json:\"dogechain\" graphql:\"dogechain\""
		Base      *float64 "json:\"base\" graphql:\"base\""
		Blast     *float64 "json:\"blast\" graphql:\"blast\""
		Total     *float64 "json:\"total\" graphql:\"total\""
	} "json:\"response\" graphql:\"response\""
}
type GetMessageBusTransactions struct {
	Response []*struct {
		FromInfo *struct {
			ChainID              *int    "json:\"chainID\" graphql:\"chainID\""
			ChainName            *string "json:\"chainName\" graphql:\"chainName\""
			DestinationChainID   *int    "json:\"destinationChainID\" graphql:\"destinationChainID\""
			DestinationChainName *string "json:\"destinationChainName\" graphql:\"destinationChainName\""
			ContractAddress      *string "json:\"contractAddress\" graphql:\"contractAddress\""
			TxnHash              *string "json:\"txnHash\" graphql:\"txnHash\""
			Message              *string "json:\"message\" graphql:\"message\""
			MessageType          *struct {
				TearType struct {
					Recipient string "json:\"recipient\" graphql:\"recipient\""
					Amount    string "json:\"amount\" graphql:\"amount\""
				} "graphql:\"... on TearType\""
				HeroType struct {
					Recipient string "json:\"recipient\" graphql:\"recipient\""
					HeroID    string "json:\"heroID\" graphql:\"heroID\""
				} "graphql:\"... on HeroType\""
				PetType struct {
					Recipient string "json:\"recipient\" graphql:\"recipient\""
					PetID     string "json:\"petID\" graphql:\"petID\""
					Name      string "json:\"name\" graphql:\"name\""
				} "graphql:\"... on PetType\""
				UnknownType struct {
					Known bool "json:\"known\" graphql:\"known\""
				} "graphql:\"... on UnknownType\""
			} "json:\"messageType\" graphql:\"messageType\""
			BlockNumber   *int    "json:\"blockNumber\" graphql:\"blockNumber\""
			Time          *int    "json:\"time\" graphql:\"time\""
			FormattedTime *string "json:\"formattedTime\" graphql:\"formattedTime\""
		} "json:\"fromInfo\" graphql:\"fromInfo\""
		ToInfo *struct {
			ChainID         *int    "json:\"chainID\" graphql:\"chainID\""
			ChainName       *string "json:\"chainName\" graphql:\"chainName\""
			ContractAddress *string "json:\"contractAddress\" graphql:\"contractAddress\""
			TxnHash         *string "json:\"txnHash\" graphql:\"txnHash\""
			Message         *string "json:\"message\" graphql:\"message\""
			MessageType     *struct {
				TearType struct {
					Recipient string "json:\"recipient\" graphql:\"recipient\""
					Amount    string "json:\"amount\" graphql:\"amount\""
				} "graphql:\"... on TearType\""
				HeroType struct {
					Recipient string "json:\"recipient\" graphql:\"recipient\""
					HeroID    string "json:\"heroID\" graphql:\"heroID\""
				} "graphql:\"... on HeroType\""
				PetType struct {
					Recipient string "json:\"recipient\" graphql:\"recipient\""
					PetID     string "json:\"petID\" graphql:\"petID\""
					Name      string "json:\"name\" graphql:\"name\""
				} "graphql:\"... on PetType\""
				UnknownType struct {
					Known bool "json:\"known\" graphql:\"known\""
				} "graphql:\"... on UnknownType\""
			} "json:\"messageType\" graphql:\"messageType\""
			BlockNumber    *int    "json:\"blockNumber\" graphql:\"blockNumber\""
			Time           *int    "json:\"time\" graphql:\"time\""
			FormattedTime  *string "json:\"formattedTime\" graphql:\"formattedTime\""
			RevertedReason *string "json:\"revertedReason\" graphql:\"revertedReason\""
		} "json:\"toInfo\" graphql:\"toInfo\""
		MessageID *string "json:\"messageID\" graphql:\"messageID\""
		Pending   *bool   "json:\"pending\" graphql:\"pending\""
	} "json:\"response\" graphql:\"response\""
}
type GetAddressData struct {
	Response *struct {
		BridgeVolume *float64 "json:\"bridgeVolume\" graphql:\"bridgeVolume\""
		BridgeFees   *float64 "json:\"bridgeFees\" graphql:\"bridgeFees\""
		BridgeTxs    *int     "json:\"bridgeTxs\" graphql:\"bridgeTxs\""
		SwapVolume   *float64 "json:\"swapVolume\" graphql:\"swapVolume\""
		SwapFees     *float64 "json:\"swapFees\" graphql:\"swapFees\""
		SwapTxs      *int     "json:\"swapTxs\" graphql:\"swapTxs\""
		Rank         *int     "json:\"rank\" graphql:\"rank\""
		EarliestTx   *int     "json:\"earliestTx\" graphql:\"earliestTx\""
		ChainRanking []*struct {
			ChainID   *int     "json:\"chainID\" graphql:\"chainID\""
			VolumeUsd *float64 "json:\"volumeUsd\" graphql:\"volumeUsd\""
			Rank      *int     "json:\"rank\" graphql:\"rank\""
		} "json:\"chainRanking\" graphql:\"chainRanking\""
		DailyData []*struct {
			Date  *string "json:\"date\" graphql:\"date\""
			Count *int    "json:\"count\" graphql:\"count\""
		} "json:\"dailyData\" graphql:\"dailyData\""
	} "json:\"response\" graphql:\"response\""
}
type GetLeaderboard struct {
	Response []*struct {
		Address      *string  "json:\"address\" graphql:\"address\""
		VolumeUsd    *float64 "json:\"volumeUSD\" graphql:\"volumeUSD\""
		Fees         *float64 "json:\"fees\" graphql:\"fees\""
		Txs          *int     "json:\"txs\" graphql:\"txs\""
		Rank         *int     "json:\"rank\" graphql:\"rank\""
		AvgVolumeUsd *float64 "json:\"avgVolumeUSD\" graphql:\"avgVolumeUSD\""
	} "json:\"response\" graphql:\"response\""
}
type GetOriginBridgeTx struct {
	Response *struct {
		BridgeTx *struct {
			ChainID            *int     "json:\"chainID\" graphql:\"chainID\""
			DestinationChainID *int     "json:\"destinationChainID\" graphql:\"destinationChainID\""
			Address            *string  "json:\"address\" graphql:\"address\""
			TxnHash            *string  "json:\"txnHash\" graphql:\"txnHash\""
			Value              *string  "json:\"value\" graphql:\"value\""
			FormattedValue     *float64 "json:\"formattedValue\" graphql:\"formattedValue\""
			USDValue           *float64 "json:\"USDValue\" graphql:\"USDValue\""
			TokenAddress       *string  "json:\"tokenAddress\" graphql:\"tokenAddress\""
			TokenSymbol        *string  "json:\"tokenSymbol\" graphql:\"tokenSymbol\""
			BlockNumber        *int     "json:\"blockNumber\" graphql:\"blockNumber\""
			Time               *int     "json:\"time\" graphql:\"time\""
			FormattedTime      *string  "json:\"formattedTime\" graphql:\"formattedTime\""
		} "json:\"bridgeTx\" graphql:\"bridgeTx\""
		Pending     *bool               "json:\"pending\" graphql:\"pending\""
		Type        *model.BridgeTxType "json:\"type\" graphql:\"type\""
		Kappa       *string             "json:\"kappa\" graphql:\"kappa\""
		KappaStatus *model.KappaStatus  "json:\"kappaStatus\" graphql:\"kappaStatus\""
	} "json:\"response\" graphql:\"response\""
}
type GetDestinationBridgeTx struct {
	Response *struct {
		BridgeTx *struct {
			ChainID            *int     "json:\"chainID\" graphql:\"chainID\""
			DestinationChainID *int     "json:\"destinationChainID\" graphql:\"destinationChainID\""
			Address            *string  "json:\"address\" graphql:\"address\""
			TxnHash            *string  "json:\"txnHash\" graphql:\"txnHash\""
			Value              *string  "json:\"value\" graphql:\"value\""
			FormattedValue     *float64 "json:\"formattedValue\" graphql:\"formattedValue\""
			USDValue           *float64 "json:\"USDValue\" graphql:\"USDValue\""
			TokenAddress       *string  "json:\"tokenAddress\" graphql:\"tokenAddress\""
			TokenSymbol        *string  "json:\"tokenSymbol\" graphql:\"tokenSymbol\""
			BlockNumber        *int     "json:\"blockNumber\" graphql:\"blockNumber\""
			Time               *int     "json:\"time\" graphql:\"time\""
			FormattedTime      *string  "json:\"formattedTime\" graphql:\"formattedTime\""
		} "json:\"bridgeTx\" graphql:\"bridgeTx\""
		Pending     *bool               "json:\"pending\" graphql:\"pending\""
		Type        *model.BridgeTxType "json:\"type\" graphql:\"type\""
		Kappa       *string             "json:\"kappa\" graphql:\"kappa\""
		KappaStatus *model.KappaStatus  "json:\"kappaStatus\" graphql:\"kappaStatus\""
	} "json:\"response\" graphql:\"response\""
}

const GetBridgeTransactionsDocument = `query GetBridgeTransactions ($chainIDTo: [Int], $chainIDFrom: [Int], $addressTo: String, $addressFrom: String, $maxAmount: Int, $minAmount: Int, $maxAmountUSD: Int, $minAmountUSD: Int, $startTime: Int, $endTime: Int, $txHash: String, $kappa: String, $pending: Boolean, $page: Int, $tokenAddressFrom: [String], $tokenAddressTo: [String], $useMv: Boolean) {
	response: bridgeTransactions(chainIDTo: $chainIDTo, chainIDFrom: $chainIDFrom, addressTo: $addressTo, addressFrom: $addressFrom, maxAmount: $maxAmount, minAmount: $minAmount, maxAmountUsd: $maxAmountUSD, minAmountUsd: $minAmountUSD, startTime: $startTime, endTime: $endTime, txnHash: $txHash, kappa: $kappa, pending: $pending, page: $page, tokenAddressTo: $tokenAddressTo, tokenAddressFrom: $tokenAddressFrom, useMv: $useMv) {
		fromInfo {
			chainID
			destinationChainID
			address
			txnHash
			value
			formattedValue
			USDValue
			tokenAddress
			tokenSymbol
			blockNumber
			time
			formattedTime
		}
		toInfo {
			chainID
			address
			txnHash
			value
			formattedValue
			USDValue
			tokenAddress
			tokenSymbol
			blockNumber
			time
			formattedTime
		}
		kappa
		pending
		swapSuccess
	}
}
`

func (c *Client) GetBridgeTransactions(ctx context.Context, chainIDTo []*int, chainIDFrom []*int, addressTo *string, addressFrom *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txHash *string, kappa *string, pending *bool, page *int, tokenAddressFrom []*string, tokenAddressTo []*string, useMv *bool, httpRequestOptions ...client.HTTPRequestOption) (*GetBridgeTransactions, error) {
	vars := map[string]interface{}{
		"chainIDTo":        chainIDTo,
		"chainIDFrom":      chainIDFrom,
		"addressTo":        addressTo,
		"addressFrom":      addressFrom,
		"maxAmount":        maxAmount,
		"minAmount":        minAmount,
		"maxAmountUSD":     maxAmountUsd,
		"minAmountUSD":     minAmountUsd,
		"startTime":        startTime,
		"endTime":          endTime,
		"txHash":           txHash,
		"kappa":            kappa,
		"pending":          pending,
		"page":             page,
		"tokenAddressFrom": tokenAddressFrom,
		"tokenAddressTo":   tokenAddressTo,
		"useMv":            useMv,
	}

	var res GetBridgeTransactions
	if err := c.Client.Post(ctx, "GetBridgeTransactions", GetBridgeTransactionsDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetCountByChainIDDocument = `query GetCountByChainId ($chainID: Int, $address: String, $direction: Direction, $hours: Int) {
	response: countByChainId(chainID: $chainID, address: $address, direction: $direction, hours: $hours) {
		count
		chainID
	}
}
`

func (c *Client) GetCountByChainID(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int, httpRequestOptions ...client.HTTPRequestOption) (*GetCountByChainID, error) {
	vars := map[string]interface{}{
		"chainID":   chainID,
		"address":   address,
		"direction": direction,
		"hours":     hours,
	}

	var res GetCountByChainID
	if err := c.Client.Post(ctx, "GetCountByChainId", GetCountByChainIDDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetCountByTokenAddressDocument = `query GetCountByTokenAddress ($chainID: Int, $address: String, $direction: Direction, $hours: Int) {
	response: countByTokenAddress(chainID: $chainID, address: $address, direction: $direction, hours: $hours) {
		chainID
		tokenAddress
		count
	}
}
`

func (c *Client) GetCountByTokenAddress(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int, httpRequestOptions ...client.HTTPRequestOption) (*GetCountByTokenAddress, error) {
	vars := map[string]interface{}{
		"chainID":   chainID,
		"address":   address,
		"direction": direction,
		"hours":     hours,
	}

	var res GetCountByTokenAddress
	if err := c.Client.Post(ctx, "GetCountByTokenAddress", GetCountByTokenAddressDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetAddressRankingDocument = `query GetAddressRanking ($hours: Int) {
	response: addressRanking(hours: $hours) {
		address
		count
	}
}
`

func (c *Client) GetAddressRanking(ctx context.Context, hours *int, httpRequestOptions ...client.HTTPRequestOption) (*GetAddressRanking, error) {
	vars := map[string]interface{}{
		"hours": hours,
	}

	var res GetAddressRanking
	if err := c.Client.Post(ctx, "GetAddressRanking", GetAddressRankingDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetRankedChainIDsByVolumeDocument = `query GetRankedChainIDsByVolume ($duration: Duration) {
	response: rankedChainIDsByVolume(duration: $duration) {
		chainID
		total
	}
}
`

func (c *Client) GetRankedChainIDsByVolume(ctx context.Context, duration *model.Duration, httpRequestOptions ...client.HTTPRequestOption) (*GetRankedChainIDsByVolume, error) {
	vars := map[string]interface{}{
		"duration": duration,
	}

	var res GetRankedChainIDsByVolume
	if err := c.Client.Post(ctx, "GetRankedChainIDsByVolume", GetRankedChainIDsByVolumeDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetBlockHeightDocument = `query GetBlockHeight ($contracts: [ContractQuery]) {
	response: getBlockHeight(contracts: $contracts) {
		chainID
		type
		blockNumber
	}
}
`

func (c *Client) GetBlockHeight(ctx context.Context, contracts []*model.ContractQuery, httpRequestOptions ...client.HTTPRequestOption) (*GetBlockHeight, error) {
	vars := map[string]interface{}{
		"contracts": contracts,
	}

	var res GetBlockHeight
	if err := c.Client.Post(ctx, "GetBlockHeight", GetBlockHeightDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetAmountStatisticDocument = `query GetAmountStatistic ($type: StatisticType!, $platform: Platform, $duration: Duration, $chainID: Int, $address: String, $tokenAddress: String, $useMv: Boolean) {
	response: amountStatistic(type: $type, duration: $duration, platform: $platform, chainID: $chainID, address: $address, tokenAddress: $tokenAddress, useMv: $useMv) {
		value
	}
}
`

func (c *Client) GetAmountStatistic(ctx context.Context, typeArg model.StatisticType, platform *model.Platform, duration *model.Duration, chainID *int, address *string, tokenAddress *string, useMv *bool, httpRequestOptions ...client.HTTPRequestOption) (*GetAmountStatistic, error) {
	vars := map[string]interface{}{
		"type":         typeArg,
		"platform":     platform,
		"duration":     duration,
		"chainID":      chainID,
		"address":      address,
		"tokenAddress": tokenAddress,
		"useMv":        useMv,
	}

	var res GetAmountStatistic
	if err := c.Client.Post(ctx, "GetAmountStatistic", GetAmountStatisticDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetDailyStatisticsByChainDocument = `query GetDailyStatisticsByChain ($chainID: Int, $type: DailyStatisticType, $duration: Duration, $platform: Platform, $useMv: Boolean) {
	response: dailyStatisticsByChain(chainID: $chainID, type: $type, duration: $duration, platform: $platform, useMv: $useMv) {
		date
		ethereum
		optimism
		cronos
		bsc
		polygon
		fantom
		boba
		metis
		moonbeam
		moonriver
		klaytn
		arbitrum
		avalanche
		dfk
		aurora
		harmony
		canto
		dogechain
		base
		blast
		total
	}
}
`

func (c *Client) GetDailyStatisticsByChain(ctx context.Context, chainID *int, typeArg *model.DailyStatisticType, duration *model.Duration, platform *model.Platform, useMv *bool, httpRequestOptions ...client.HTTPRequestOption) (*GetDailyStatisticsByChain, error) {
	vars := map[string]interface{}{
		"chainID":  chainID,
		"type":     typeArg,
		"duration": duration,
		"platform": platform,
		"useMv":    useMv,
	}

	var res GetDailyStatisticsByChain
	if err := c.Client.Post(ctx, "GetDailyStatisticsByChain", GetDailyStatisticsByChainDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetMessageBusTransactionsDocument = `query GetMessageBusTransactions ($chainID: [Int], $contractAddress: String, $startTime: Int, $endTime: Int, $txHash: String, $messageID: String, $pending: Boolean, $reverted: Boolean, $page: Int) {
	response: messageBusTransactions(chainID: $chainID, contractAddress: $contractAddress, startTime: $startTime, endTime: $endTime, txnHash: $txHash, messageID: $messageID, pending: $pending, reverted: $reverted, page: $page) {
		fromInfo {
			chainID
			chainName
			destinationChainID
			destinationChainName
			contractAddress
			txnHash
			message
			messageType {
				... on TearType {
					recipient
					amount
				}
				... on HeroType {
					recipient
					heroID
				}
				... on PetType {
					recipient
					petID
					name
				}
				... on UnknownType {
					known
				}
			}
			blockNumber
			time
			formattedTime
		}
		toInfo {
			chainID
			chainName
			contractAddress
			txnHash
			message
			messageType {
				... on TearType {
					recipient
					amount
				}
				... on HeroType {
					recipient
					heroID
				}
				... on PetType {
					recipient
					petID
					name
				}
				... on UnknownType {
					known
				}
			}
			blockNumber
			time
			formattedTime
			revertedReason
		}
		messageID
		pending
	}
}
`

func (c *Client) GetMessageBusTransactions(ctx context.Context, chainID []*int, contractAddress *string, startTime *int, endTime *int, txHash *string, messageID *string, pending *bool, reverted *bool, page *int, httpRequestOptions ...client.HTTPRequestOption) (*GetMessageBusTransactions, error) {
	vars := map[string]interface{}{
		"chainID":         chainID,
		"contractAddress": contractAddress,
		"startTime":       startTime,
		"endTime":         endTime,
		"txHash":          txHash,
		"messageID":       messageID,
		"pending":         pending,
		"reverted":        reverted,
		"page":            page,
	}

	var res GetMessageBusTransactions
	if err := c.Client.Post(ctx, "GetMessageBusTransactions", GetMessageBusTransactionsDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetAddressDataDocument = `query GetAddressData ($address: String!) {
	response: addressData(address: $address) {
		bridgeVolume
		bridgeFees
		bridgeTxs
		swapVolume
		swapFees
		swapTxs
		rank
		earliestTx
		chainRanking {
			chainID
			volumeUsd
			rank
		}
		dailyData {
			date
			count
		}
	}
}
`

func (c *Client) GetAddressData(ctx context.Context, address string, httpRequestOptions ...client.HTTPRequestOption) (*GetAddressData, error) {
	vars := map[string]interface{}{
		"address": address,
	}

	var res GetAddressData
	if err := c.Client.Post(ctx, "GetAddressData", GetAddressDataDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetLeaderboardDocument = `query GetLeaderboard ($duration: Duration, $chainID: Int, $useMv: Boolean, $page: Int) {
	response: leaderboard(duration: $duration, chainID: $chainID, useMv: $useMv, page: $page) {
		address
		volumeUSD
		fees
		txs
		rank
		avgVolumeUSD
	}
}
`

func (c *Client) GetLeaderboard(ctx context.Context, duration *model.Duration, chainID *int, useMv *bool, page *int, httpRequestOptions ...client.HTTPRequestOption) (*GetLeaderboard, error) {
	vars := map[string]interface{}{
		"duration": duration,
		"chainID":  chainID,
		"useMv":    useMv,
		"page":     page,
	}

	var res GetLeaderboard
	if err := c.Client.Post(ctx, "GetLeaderboard", GetLeaderboardDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetOriginBridgeTxDocument = `query GetOriginBridgeTx ($chainID: Int!, $txnHash: String!, $bridgeType: BridgeType!) {
	response: getOriginBridgeTx(chainID: $chainID, txnHash: $txnHash, bridgeType: $bridgeType) {
		bridgeTx {
			chainID
			destinationChainID
			address
			txnHash
			value
			formattedValue
			USDValue
			tokenAddress
			tokenSymbol
			blockNumber
			time
			formattedTime
		}
		pending
		type
		kappa
		kappaStatus
	}
}
`

func (c *Client) GetOriginBridgeTx(ctx context.Context, chainID int, txnHash string, bridgeType model.BridgeType, httpRequestOptions ...client.HTTPRequestOption) (*GetOriginBridgeTx, error) {
	vars := map[string]interface{}{
		"chainID":    chainID,
		"txnHash":    txnHash,
		"bridgeType": bridgeType,
	}

	var res GetOriginBridgeTx
	if err := c.Client.Post(ctx, "GetOriginBridgeTx", GetOriginBridgeTxDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetDestinationBridgeTxDocument = `query GetDestinationBridgeTx ($chainID: Int!, $kappa: String!, $address: String!, $timestamp: Int!, $bridgeType: BridgeType!, $historical: Boolean) {
	response: getDestinationBridgeTx(chainID: $chainID, address: $address, kappa: $kappa, timestamp: $timestamp, bridgeType: $bridgeType, historical: $historical) {
		bridgeTx {
			chainID
			destinationChainID
			address
			txnHash
			value
			formattedValue
			USDValue
			tokenAddress
			tokenSymbol
			blockNumber
			time
			formattedTime
		}
		pending
		type
		kappa
		kappaStatus
	}
}
`

func (c *Client) GetDestinationBridgeTx(ctx context.Context, chainID int, kappa string, address string, timestamp int, bridgeType model.BridgeType, historical *bool, httpRequestOptions ...client.HTTPRequestOption) (*GetDestinationBridgeTx, error) {
	vars := map[string]interface{}{
		"chainID":    chainID,
		"kappa":      kappa,
		"address":    address,
		"timestamp":  timestamp,
		"bridgeType": bridgeType,
		"historical": historical,
	}

	var res GetDestinationBridgeTx
	if err := c.Client.Post(ctx, "GetDestinationBridgeTx", GetDestinationBridgeTxDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}
