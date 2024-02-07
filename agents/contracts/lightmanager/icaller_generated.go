// autogenerated file

package lightmanager

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ILightManagerCaller ...
type ILightManagerCaller interface {
	// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
	//
	// Solidity: function agentRoot() view returns(bytes32)
	AgentRoot(opts *bind.CallOpts) ([32]byte, error)
	// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
	//
	// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
	AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error)
	// Destination is a free data retrieval call binding the contract method 0xb269681d.
	//
	// Solidity: function destination() view returns(address)
	Destination(opts *bind.CallOpts) (common.Address, error)
	// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
	//
	// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
	DisputeStatus(opts *bind.CallOpts, agent common.Address) (struct {
		Flag        uint8
		Rival       common.Address
		FraudProver common.Address
		DisputePtr  *big.Int
	}, error)
	// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
	//
	// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
	GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
		Agent  common.Address
		Status AgentStatus
	}, error)
	// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
	//
	// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
	GetDispute(opts *bind.CallOpts, index *big.Int) (struct {
		Guard           common.Address
		Notary          common.Address
		SlashedAgent    common.Address
		FraudProver     common.Address
		ReportPayload   []byte
		ReportSignature []byte
	}, error)
	// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
	//
	// Solidity: function getDisputesAmount() view returns(uint256)
	GetDisputesAmount(opts *bind.CallOpts) (*big.Int, error)
	// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
	//
	// Solidity: function inbox() view returns(address)
	Inbox(opts *bind.CallOpts) (common.Address, error)
	// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
	//
	// Solidity: function localDomain() view returns(uint32)
	LocalDomain(opts *bind.CallOpts) (uint32, error)
	// Origin is a free data retrieval call binding the contract method 0x938b5f32.
	//
	// Solidity: function origin() view returns(address)
	Origin(opts *bind.CallOpts) (common.Address, error)
	// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
	//
	// Solidity: function owner() view returns(address)
	Owner(opts *bind.CallOpts) (common.Address, error)
	// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
	//
	// Solidity: function pendingOwner() view returns(address)
	PendingOwner(opts *bind.CallOpts) (common.Address, error)
	// ProposedAgentRootData is a free data retrieval call binding the contract method 0x5396feef.
	//
	// Solidity: function proposedAgentRootData() view returns(bytes32 agentRoot_, uint256 proposedAt_)
	ProposedAgentRootData(opts *bind.CallOpts) (struct {
		AgentRoot  [32]byte
		ProposedAt *big.Int
	}, error)
	// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
	//
	// Solidity: function synapseDomain() view returns(uint32)
	SynapseDomain(opts *bind.CallOpts) (uint32, error)
	// Version is a free data retrieval call binding the contract method 0x54fd4d50.
	//
	// Solidity: function version() view returns(string versionString)
	Version(opts *bind.CallOpts) (string, error)
}
