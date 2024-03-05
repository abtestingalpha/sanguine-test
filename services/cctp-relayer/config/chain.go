package config

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/richardwilkes/toolbox/collection"
)

// ChainConfig defines the config for a specific chain.
type ChainConfig struct { // ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// SynapseCCTPAddress is the address of the SynapseCCTP contract.
	SynapseCCTPAddress string `yaml:"synapse_cctp_address"`
	// CircleCCTPAddress is the address of the CircleCCTP contract.
	CircleCCTPAddress string `yaml:"circle_cctp_address"`
}

// GetSynapseCCTPAddress returns the SynapseCCTP address.
func (c ChainConfig) GetSynapseCCTPAddress() common.Address {
	return common.HexToAddress(c.SynapseCCTPAddress)
}

// GetCircleCCTPAddress returns the CircleCCTP address.
func (c ChainConfig) GetCircleCCTPAddress() common.Address {
	return common.HexToAddress(c.CircleCCTPAddress)
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
func (c ChainConfigs) IsValid(_ context.Context) (ok bool, err error) {
	intSet := collection.Set[uint32]{}

	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice: %s", cfg.ChainID, "duplicate chain id")
		}
		intSet.Add(cfg.ChainID)

		if len(cfg.SynapseCCTPAddress) > 0 && !common.IsHexAddress(cfg.SynapseCCTPAddress) {
			return false, fmt.Errorf("invalid address %s: %s", cfg.SynapseCCTPAddress, "invalid address")
		}

		if len(cfg.CircleCCTPAddress) > 0 && !common.IsHexAddress(cfg.CircleCCTPAddress) {
			return false, fmt.Errorf("invalid address %s: %s", cfg.CircleCCTPAddress, "invalid address")
		}
	}

	return true, nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("%s: chain ID cannot be 0", "invalid chain id")
	}

	if c.SynapseCCTPAddress == "" && c.CircleCCTPAddress == "" {
		return false, fmt.Errorf("at least one of SynapseCCTPAddress and CircleCCTPAddress is required")
	}

	return true, nil
}
