// Package relconfig contains the config yaml object for the relayer.
package relconfig

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"

	"path/filepath"
)

// Config represents the configuration for the relayer.
// TODO: validation function.
//
//go:generate go run github.com/vburenin/ifacemaker -f config.go -s Config -i IConfig -p relconfig -o iconfig_generated.go -c "autogenerated file"
type Config struct {
	// Chains is a map of chainID -> chain config.
	Chains map[int]ChainConfig `yaml:"bridges"`
	// OmniRPCURL is the URL of the OmniRPC server.
	OmniRPCURL string `yaml:"omnirpc_url"`
	// RfqAPIURL is the URL of the RFQ API.
	RfqAPIURL string `yaml:"rfq_url"`
	// RelayerAPIPort is the port of the relayer API.
	RelayerAPIPort string `yaml:"relayer_api_port"`
	// Database is the database config.
	Database DatabaseConfig `yaml:"database"`
	// QuotableTokens is a map of token -> list of quotable tokens.
	QuotableTokens map[string][]string `yaml:"quotable_tokens"`
	// Signer is the signer config.
	Signer config.SignerConfig `yaml:"signer"`
	// Submitter is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	// FeePricer is the fee pricer config.
	FeePricer FeePricerConfig `yaml:"fee_pricer"`
	// MinGasToken is minimum amount of gas that should be leftover after bridging a gas token.
	MinGasToken string `yaml:"min_gas_token"`
}

// ChainConfig represents the configuration for a chain.
type ChainConfig struct {
	// Bridge is the bridge confirmation count.
	Bridge string `yaml:"address"`
	// Confirmations is the number of required confirmations
	Confirmations uint64 `yaml:"confirmations"`
	// Tokens is a map of token ID -> token config.
	Tokens map[string]TokenConfig `yaml:"tokens"`
	// NativeToken is the native token of the chain (pays gas).
	NativeToken string `yaml:"native_token"`
}

// TokenConfig represents the configuration for a token.
type TokenConfig struct {
	// Address is the token address.
	Address string `yaml:"address"`
	// Decimals is the token decimals.
	Decimals uint8 `yaml:"decimals"`
	// For now, specify the USD price of the token in the config.
	PriceUSD float64 `yaml:"price_usd"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}

// FeePricerConfig represents the configuration for the fee pricer.
type FeePricerConfig struct {
	// OriginGasEstimate is the gas required to execute prove + claim transactions on origin chain.
	OriginGasEstimate int `yaml:"origin_gas_estimate"`
	// DestinationGasEstimate is the gas required to execute relay transaction on destination chain.
	DestinationGasEstimate int `yaml:"destination_gas_estimate"`
	// FixedFeeMultiplier is the multiplier for the fixed fee.
	FixedFeeMultiplier float64 `yaml:"fixed_fee_multiplier"`
	// GasPriceCacheTTLSeconds is the TTL for the gas price cache.
	GasPriceCacheTTLSeconds int `yaml:"gas_price_cache_ttl"`
	// TokenPriceCacheTTLSeconds is the TTL for the token price cache.
	TokenPriceCacheTTLSeconds int `yaml:"token_price_cache_ttl"`
}

// LoadConfig loads the config from the given path.
func LoadConfig(path string) (config Config, err error) {
	input, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &config)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return config, nil
}

// GetChains returns the chains config.
func (c Config) GetChains() map[int]ChainConfig {
	return c.Chains
}

// GetOmniRPCURL returns the OmniRPCURL.
func (c Config) GetOmniRPCURL() string {
	return c.OmniRPCURL
}

// GetRfqAPIURL returns the RFQ API URL.
func (c Config) GetRfqAPIURL() string {
	return c.RfqAPIURL
}

// GetDatabase returns the database config.
func (c Config) GetDatabase() DatabaseConfig {
	return c.Database
}

// GetSigner returns the signer config.
func (c Config) GetSigner() config.SignerConfig {
	return c.Signer
}

// GetFeePricer returns the fee pricer config.
func (c Config) GetFeePricer() FeePricerConfig {
	return c.FeePricer
}

// GetTokenID returns the tokenID for the given chain and address.
func (c Config) GetTokenID(chain int, addr string) (string, error) {
	chainConfig, ok := c.Chains[int(chain)]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenID, tokenConfig := range chainConfig.Tokens {
		if tokenConfig.Address == addr {
			return tokenID, nil
		}
	}
	return "", fmt.Errorf("no tokenID found for chain %d and address %s", chain, addr)
}

// GetQuotableTokens returns the quotable tokens for the given token.
func (c Config) GetQuotableTokens(token string) ([]string, error) {
	tokens, ok := c.QuotableTokens[token]
	if !ok {
		return nil, fmt.Errorf("no quotable tokens for token %s", token)
	}
	return tokens, nil
}

// GetNativeToken returns the native token for the given chain.
func (c Config) GetNativeToken(chainID uint32) (string, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return "", fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	if len(chainConfig.NativeToken) == 0 {
		return "", fmt.Errorf("chain config for chainID %d does not have a native token", chainID)
	}
	return chainConfig.NativeToken, nil
}

// GetTokenDecimals returns the token decimals for the given chain and token.
func (c Config) GetTokenDecimals(chainID uint32, token string) (uint8, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return 0, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		if token == tokenName {
			return tokenConfig.Decimals, nil
		}
	}
	return 0, fmt.Errorf("could not get token decimals for chain %d and token %s", chainID, token)
}

// GetTokens returns the tokens for the given chain.
func (c Config) GetTokens(chainID uint32) (map[string]TokenConfig, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return nil, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	return chainConfig.Tokens, nil
}

// GetTokenName returns the token name for the given chain and address.
func (c Config) GetTokenName(chain uint32, addr string) (string, error) {
	chainConfig, ok := c.Chains[int(chain)]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		// TODO: probably a better way to do this.
		if strings.ToLower(tokenConfig.Address) == strings.ToLower(addr) {
			return tokenName, nil
		}
	}
	return "", fmt.Errorf("no tokenName found for chain %d and address %s", chain, addr)
}

const defaultFixedFeeMultiplier = 1

// GetFixedFeeMultiplier returns the fixed fee multiplier.
func (c Config) GetFixedFeeMultiplier() float64 {
	if c.FeePricer.FixedFeeMultiplier <= 0 {
		return defaultFixedFeeMultiplier
	}
	return c.FeePricer.FixedFeeMultiplier
}

// GetMinGasToken returns the min gas token.
func (c Config) GetMinGasToken() (*big.Int, error) {
	minGasToken, ok := new(big.Int).SetString(c.MinGasToken, 10)
	if !ok {
		return nil, fmt.Errorf("could not parse min gas token %s", c.MinGasToken)
	}
	return minGasToken, nil
}

var _ IConfig = &Config{}
