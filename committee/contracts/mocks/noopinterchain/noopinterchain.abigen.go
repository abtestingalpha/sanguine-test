// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package noopinterchain

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"getWriterNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b8a740e0": "getEntry(address,uint256)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"4a30a686": "getWriterNonce(address)",
		"d48588e0": "readEntry(address,(uint256,bytes32,uint256,bytes32))",
		"b4f16bae": "requestVerification(uint256,address,uint256,address[])",
		"9cbc6dd5": "verifyEntry((uint256,bytes32,uint256,bytes32))",
		"2ad8c706": "writeEntry(bytes32)",
		"67c769af": "writeEntryWithVerification(uint256,bytes32,address[])",
	},
}

// IInterchainDBABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainDBMetaData.ABI instead.
var IInterchainDBABI = IInterchainDBMetaData.ABI

// Deprecated: Use IInterchainDBMetaData.Sigs instead.
// IInterchainDBFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainDBFuncSigs = IInterchainDBMetaData.Sigs

// IInterchainDB is an auto generated Go binding around an Ethereum contract.
type IInterchainDB struct {
	IInterchainDBCaller     // Read-only binding to the contract
	IInterchainDBTransactor // Write-only binding to the contract
	IInterchainDBFilterer   // Log filterer for contract events
}

// IInterchainDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainDBSession struct {
	Contract     *IInterchainDB    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainDBCallerSession struct {
	Contract *IInterchainDBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IInterchainDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainDBTransactorSession struct {
	Contract     *IInterchainDBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IInterchainDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainDBRaw struct {
	Contract *IInterchainDB // Generic contract binding to access the raw methods on
}

// IInterchainDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainDBCallerRaw struct {
	Contract *IInterchainDBCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainDBTransactorRaw struct {
	Contract *IInterchainDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainDB creates a new instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDB(address common.Address, backend bind.ContractBackend) (*IInterchainDB, error) {
	contract, err := bindIInterchainDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainDB{IInterchainDBCaller: IInterchainDBCaller{contract: contract}, IInterchainDBTransactor: IInterchainDBTransactor{contract: contract}, IInterchainDBFilterer: IInterchainDBFilterer{contract: contract}}, nil
}

// NewIInterchainDBCaller creates a new read-only instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBCaller(address common.Address, caller bind.ContractCaller) (*IInterchainDBCaller, error) {
	contract, err := bindIInterchainDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBCaller{contract: contract}, nil
}

// NewIInterchainDBTransactor creates a new write-only instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainDBTransactor, error) {
	contract, err := bindIInterchainDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBTransactor{contract: contract}, nil
}

// NewIInterchainDBFilterer creates a new log filterer instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainDBFilterer, error) {
	contract, err := bindIInterchainDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBFilterer{contract: contract}, nil
}

// bindIInterchainDB binds a generic wrapper to an already deployed contract.
func bindIInterchainDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainDB *IInterchainDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainDB.Contract.IInterchainDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainDB *IInterchainDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainDB.Contract.IInterchainDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainDB *IInterchainDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainDB.Contract.IInterchainDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainDB *IInterchainDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainDB *IInterchainDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainDB *IInterchainDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainDB.Contract.contract.Transact(opts, method, params...)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetEntry(opts *bind.CallOpts, writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntry", writer, writerNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, writer, writerNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, writer, writerNonce)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", destChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, destChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, destChainId, srcModules)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetWriterNonce(opts *bind.CallOpts, writer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getWriterNonce", writer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetWriterNonce(&_IInterchainDB.CallOpts, writer)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetWriterNonce(&_IInterchainDB.CallOpts, writer)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestVerification", destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208d1ebcab36f6edf8815af2f3b9ccb1277247e1c6668f8bdd0dc39f0eef9721c964736f6c63430008140033",
}

// InterchainEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainEntryLibMetaData.ABI instead.
var InterchainEntryLibABI = InterchainEntryLibMetaData.ABI

// InterchainEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainEntryLibMetaData.Bin instead.
var InterchainEntryLibBin = InterchainEntryLibMetaData.Bin

// DeployInterchainEntryLib deploys a new Ethereum contract, binding an instance of InterchainEntryLib to it.
func DeployInterchainEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainEntryLib, error) {
	parsed, err := InterchainEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainEntryLib{InterchainEntryLibCaller: InterchainEntryLibCaller{contract: contract}, InterchainEntryLibTransactor: InterchainEntryLibTransactor{contract: contract}, InterchainEntryLibFilterer: InterchainEntryLibFilterer{contract: contract}}, nil
}

// InterchainEntryLib is an auto generated Go binding around an Ethereum contract.
type InterchainEntryLib struct {
	InterchainEntryLibCaller     // Read-only binding to the contract
	InterchainEntryLibTransactor // Write-only binding to the contract
	InterchainEntryLibFilterer   // Log filterer for contract events
}

// InterchainEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainEntryLibSession struct {
	Contract     *InterchainEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainEntryLibCallerSession struct {
	Contract *InterchainEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainEntryLibTransactorSession struct {
	Contract     *InterchainEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainEntryLibRaw struct {
	Contract *InterchainEntryLib // Generic contract binding to access the raw methods on
}

// InterchainEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainEntryLibCallerRaw struct {
	Contract *InterchainEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainEntryLibTransactorRaw struct {
	Contract *InterchainEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainEntryLib creates a new instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLib(address common.Address, backend bind.ContractBackend) (*InterchainEntryLib, error) {
	contract, err := bindInterchainEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLib{InterchainEntryLibCaller: InterchainEntryLibCaller{contract: contract}, InterchainEntryLibTransactor: InterchainEntryLibTransactor{contract: contract}, InterchainEntryLibFilterer: InterchainEntryLibFilterer{contract: contract}}, nil
}

// NewInterchainEntryLibCaller creates a new read-only instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainEntryLibCaller, error) {
	contract, err := bindInterchainEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibCaller{contract: contract}, nil
}

// NewInterchainEntryLibTransactor creates a new write-only instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainEntryLibTransactor, error) {
	contract, err := bindInterchainEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibTransactor{contract: contract}, nil
}

// NewInterchainEntryLibFilterer creates a new log filterer instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainEntryLibFilterer, error) {
	contract, err := bindInterchainEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibFilterer{contract: contract}, nil
}

// bindInterchainEntryLib binds a generic wrapper to an already deployed contract.
func bindInterchainEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainEntryLib *InterchainEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainEntryLib.Contract.InterchainEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainEntryLib *InterchainEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.InterchainEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainEntryLib *InterchainEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.InterchainEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainEntryLib *InterchainEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainEntryLib *InterchainEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainEntryLib *InterchainEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.contract.Transact(opts, method, params...)
}

// NoOpInterchainMetaData contains all meta data concerning the NoOpInterchain contract.
var NoOpInterchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"getWriterNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b8a740e0": "getEntry(address,uint256)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"4a30a686": "getWriterNonce(address)",
		"d48588e0": "readEntry(address,(uint256,bytes32,uint256,bytes32))",
		"b4f16bae": "requestVerification(uint256,address,uint256,address[])",
		"9cbc6dd5": "verifyEntry((uint256,bytes32,uint256,bytes32))",
		"2ad8c706": "writeEntry(bytes32)",
		"67c769af": "writeEntryWithVerification(uint256,bytes32,address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b5061054e806100206000396000f3fe60806040526004361061007b5760003560e01c8063b4f16bae1161004e578063b4f16bae14610106578063b8a740e01461011a578063d48588e0146101af578063fc7686ec146101d257600080fd5b80632ad8c706146100805780634a30a686146100b457806367c769af146100cf5780639cbc6dd5146100e6575b600080fd5b34801561008c57600080fd5b506100a161009b3660046101ed565b50600090565b6040519081526020015b60405180910390f35b3480156100c057600080fd5b506100a161009b36600461022f565b6100a16100dd366004610340565b60009392505050565b3480156100f257600080fd5b506101046101013660046103f6565b50565b005b610104610114366004610412565b50505050565b34801561012657600080fd5b5061017c610135366004610473565b505060408051608080820183526000808352602080840182905283850182905260609384018290528451928301855281835282018190529281018390529081019190915290565b6040516100ab91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b3480156101bb57600080fd5b506100a16101ca36600461049d565b600092915050565b3480156101de57600080fd5b506100a16101ca3660046104d1565b6000602082840312156101ff57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461022a57600080fd5b919050565b60006020828403121561024157600080fd5b61024a82610206565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261029157600080fd5b8135602067ffffffffffffffff808311156102ae576102ae610251565b8260051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811084821117156102f1576102f1610251565b60405293845285810183019383810192508785111561030f57600080fd5b83870191505b848210156103355761032682610206565b83529183019190830190610315565b979650505050505050565b60008060006060848603121561035557600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561037a57600080fd5b61038686828701610280565b9150509250925092565b6000608082840312156103a257600080fd5b6040516080810181811067ffffffffffffffff821117156103c5576103c5610251565b8060405250809150823581526020830135602082015260408301356040820152606083013560608201525092915050565b60006080828403121561040857600080fd5b61024a8383610390565b6000806000806080858703121561042857600080fd5b8435935061043860208601610206565b925060408501359150606085013567ffffffffffffffff81111561045b57600080fd5b61046787828801610280565b91505092959194509250565b6000806040838503121561048657600080fd5b61048f83610206565b946020939093013593505050565b60008060a083850312156104b057600080fd5b6104b983610206565b91506104c88460208501610390565b90509250929050565b600080604083850312156104e457600080fd5b82359150602083013567ffffffffffffffff81111561050257600080fd5b61050e85828601610280565b915050925092905056fea264697066735822122032b617182fcd414981a26963db8cde3c706e66143b001d9991e6e56b88f5cf5564736f6c63430008140033",
}

// NoOpInterchainABI is the input ABI used to generate the binding from.
// Deprecated: Use NoOpInterchainMetaData.ABI instead.
var NoOpInterchainABI = NoOpInterchainMetaData.ABI

// Deprecated: Use NoOpInterchainMetaData.Sigs instead.
// NoOpInterchainFuncSigs maps the 4-byte function signature to its string representation.
var NoOpInterchainFuncSigs = NoOpInterchainMetaData.Sigs

// NoOpInterchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NoOpInterchainMetaData.Bin instead.
var NoOpInterchainBin = NoOpInterchainMetaData.Bin

// DeployNoOpInterchain deploys a new Ethereum contract, binding an instance of NoOpInterchain to it.
func DeployNoOpInterchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOpInterchain, error) {
	parsed, err := NoOpInterchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoOpInterchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOpInterchain{NoOpInterchainCaller: NoOpInterchainCaller{contract: contract}, NoOpInterchainTransactor: NoOpInterchainTransactor{contract: contract}, NoOpInterchainFilterer: NoOpInterchainFilterer{contract: contract}}, nil
}

// NoOpInterchain is an auto generated Go binding around an Ethereum contract.
type NoOpInterchain struct {
	NoOpInterchainCaller     // Read-only binding to the contract
	NoOpInterchainTransactor // Write-only binding to the contract
	NoOpInterchainFilterer   // Log filterer for contract events
}

// NoOpInterchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoOpInterchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpInterchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoOpInterchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpInterchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoOpInterchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpInterchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoOpInterchainSession struct {
	Contract     *NoOpInterchain   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoOpInterchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoOpInterchainCallerSession struct {
	Contract *NoOpInterchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NoOpInterchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoOpInterchainTransactorSession struct {
	Contract     *NoOpInterchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NoOpInterchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoOpInterchainRaw struct {
	Contract *NoOpInterchain // Generic contract binding to access the raw methods on
}

// NoOpInterchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoOpInterchainCallerRaw struct {
	Contract *NoOpInterchainCaller // Generic read-only contract binding to access the raw methods on
}

// NoOpInterchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoOpInterchainTransactorRaw struct {
	Contract *NoOpInterchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoOpInterchain creates a new instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchain(address common.Address, backend bind.ContractBackend) (*NoOpInterchain, error) {
	contract, err := bindNoOpInterchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchain{NoOpInterchainCaller: NoOpInterchainCaller{contract: contract}, NoOpInterchainTransactor: NoOpInterchainTransactor{contract: contract}, NoOpInterchainFilterer: NoOpInterchainFilterer{contract: contract}}, nil
}

// NewNoOpInterchainCaller creates a new read-only instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchainCaller(address common.Address, caller bind.ContractCaller) (*NoOpInterchainCaller, error) {
	contract, err := bindNoOpInterchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainCaller{contract: contract}, nil
}

// NewNoOpInterchainTransactor creates a new write-only instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchainTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOpInterchainTransactor, error) {
	contract, err := bindNoOpInterchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainTransactor{contract: contract}, nil
}

// NewNoOpInterchainFilterer creates a new log filterer instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchainFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOpInterchainFilterer, error) {
	contract, err := bindNoOpInterchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainFilterer{contract: contract}, nil
}

// bindNoOpInterchain binds a generic wrapper to an already deployed contract.
func bindNoOpInterchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoOpInterchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOpInterchain *NoOpInterchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpInterchain.Contract.NoOpInterchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOpInterchain *NoOpInterchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.NoOpInterchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOpInterchain *NoOpInterchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.NoOpInterchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOpInterchain *NoOpInterchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpInterchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOpInterchain *NoOpInterchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOpInterchain *NoOpInterchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.contract.Transact(opts, method, params...)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_NoOpInterchain *NoOpInterchainCaller) GetEntry(opts *bind.CallOpts, writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _NoOpInterchain.contract.Call(opts, &out, "getEntry", writer, writerNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_NoOpInterchain *NoOpInterchainSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _NoOpInterchain.Contract.GetEntry(&_NoOpInterchain.CallOpts, writer, writerNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_NoOpInterchain *NoOpInterchainCallerSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _NoOpInterchain.Contract.GetEntry(&_NoOpInterchain.CallOpts, writer, writerNonce)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_NoOpInterchain *NoOpInterchainCaller) GetInterchainFee(opts *bind.CallOpts, destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NoOpInterchain.contract.Call(opts, &out, "getInterchainFee", destChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_NoOpInterchain *NoOpInterchainSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _NoOpInterchain.Contract.GetInterchainFee(&_NoOpInterchain.CallOpts, destChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_NoOpInterchain *NoOpInterchainCallerSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _NoOpInterchain.Contract.GetInterchainFee(&_NoOpInterchain.CallOpts, destChainId, srcModules)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_NoOpInterchain *NoOpInterchainCaller) GetWriterNonce(opts *bind.CallOpts, writer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NoOpInterchain.contract.Call(opts, &out, "getWriterNonce", writer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_NoOpInterchain *NoOpInterchainSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _NoOpInterchain.Contract.GetWriterNonce(&_NoOpInterchain.CallOpts, writer)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_NoOpInterchain *NoOpInterchainCallerSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _NoOpInterchain.Contract.GetWriterNonce(&_NoOpInterchain.CallOpts, writer)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_NoOpInterchain *NoOpInterchainCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _NoOpInterchain.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_NoOpInterchain *NoOpInterchainSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _NoOpInterchain.Contract.ReadEntry(&_NoOpInterchain.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_NoOpInterchain *NoOpInterchainCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _NoOpInterchain.Contract.ReadEntry(&_NoOpInterchain.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_NoOpInterchain *NoOpInterchainTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.contract.Transact(opts, "requestVerification", destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_NoOpInterchain *NoOpInterchainSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.RequestVerification(&_NoOpInterchain.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_NoOpInterchain *NoOpInterchainTransactorSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.RequestVerification(&_NoOpInterchain.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_NoOpInterchain *NoOpInterchainTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _NoOpInterchain.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_NoOpInterchain *NoOpInterchainSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.VerifyEntry(&_NoOpInterchain.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_NoOpInterchain *NoOpInterchainTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.VerifyEntry(&_NoOpInterchain.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_NoOpInterchain *NoOpInterchainTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _NoOpInterchain.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_NoOpInterchain *NoOpInterchainSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.WriteEntry(&_NoOpInterchain.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_NoOpInterchain *NoOpInterchainTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.WriteEntry(&_NoOpInterchain.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_NoOpInterchain *NoOpInterchainTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_NoOpInterchain *NoOpInterchainSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.WriteEntryWithVerification(&_NoOpInterchain.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_NoOpInterchain *NoOpInterchainTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.WriteEntryWithVerification(&_NoOpInterchain.TransactOpts, destChainId, dataHash, srcModules)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e71082f6d284e4af03ba7600edb7b3c5fad8044a743f372429d43fbf488c8c1a64736f6c63430008140033",
}

// TypeCastsABI is the input ABI used to generate the binding from.
// Deprecated: Use TypeCastsMetaData.ABI instead.
var TypeCastsABI = TypeCastsMetaData.ABI

// TypeCastsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TypeCastsMetaData.Bin instead.
var TypeCastsBin = TypeCastsMetaData.Bin

// DeployTypeCasts deploys a new Ethereum contract, binding an instance of TypeCasts to it.
func DeployTypeCasts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TypeCasts, error) {
	parsed, err := TypeCastsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TypeCastsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TypeCasts{TypeCastsCaller: TypeCastsCaller{contract: contract}, TypeCastsTransactor: TypeCastsTransactor{contract: contract}, TypeCastsFilterer: TypeCastsFilterer{contract: contract}}, nil
}

// TypeCasts is an auto generated Go binding around an Ethereum contract.
type TypeCasts struct {
	TypeCastsCaller     // Read-only binding to the contract
	TypeCastsTransactor // Write-only binding to the contract
	TypeCastsFilterer   // Log filterer for contract events
}

// TypeCastsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypeCastsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeCastsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypeCastsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeCastsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypeCastsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeCastsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypeCastsSession struct {
	Contract     *TypeCasts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypeCastsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypeCastsCallerSession struct {
	Contract *TypeCastsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TypeCastsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypeCastsTransactorSession struct {
	Contract     *TypeCastsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TypeCastsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypeCastsRaw struct {
	Contract *TypeCasts // Generic contract binding to access the raw methods on
}

// TypeCastsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypeCastsCallerRaw struct {
	Contract *TypeCastsCaller // Generic read-only contract binding to access the raw methods on
}

// TypeCastsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypeCastsTransactorRaw struct {
	Contract *TypeCastsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypeCasts creates a new instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCasts(address common.Address, backend bind.ContractBackend) (*TypeCasts, error) {
	contract, err := bindTypeCasts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TypeCasts{TypeCastsCaller: TypeCastsCaller{contract: contract}, TypeCastsTransactor: TypeCastsTransactor{contract: contract}, TypeCastsFilterer: TypeCastsFilterer{contract: contract}}, nil
}

// NewTypeCastsCaller creates a new read-only instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCastsCaller(address common.Address, caller bind.ContractCaller) (*TypeCastsCaller, error) {
	contract, err := bindTypeCasts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypeCastsCaller{contract: contract}, nil
}

// NewTypeCastsTransactor creates a new write-only instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCastsTransactor(address common.Address, transactor bind.ContractTransactor) (*TypeCastsTransactor, error) {
	contract, err := bindTypeCasts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypeCastsTransactor{contract: contract}, nil
}

// NewTypeCastsFilterer creates a new log filterer instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCastsFilterer(address common.Address, filterer bind.ContractFilterer) (*TypeCastsFilterer, error) {
	contract, err := bindTypeCasts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypeCastsFilterer{contract: contract}, nil
}

// bindTypeCasts binds a generic wrapper to an already deployed contract.
func bindTypeCasts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TypeCastsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypeCasts *TypeCastsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypeCasts.Contract.TypeCastsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypeCasts *TypeCastsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypeCasts.Contract.TypeCastsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypeCasts *TypeCastsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypeCasts.Contract.TypeCastsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypeCasts *TypeCastsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypeCasts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypeCasts *TypeCastsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypeCasts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypeCasts *TypeCastsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypeCasts.Contract.contract.Transact(opts, method, params...)
}
