// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package registrarcontract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RegistrarcontractABI is the input ABI used to generate the binding from.
const RegistrarcontractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"releaseDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getAllowedTime\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unhashedName\",\"type\":\"string\"}],\"name\":\"invalidateName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"shaBid\",\"outputs\":[{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"seal\",\"type\":\"bytes32\"}],\"name\":\"cancelBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"entries\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"unsealBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"transferRegistrars\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sealedBids\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"finalizeAuction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registryStarted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"launchLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"name\":\"newBid\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"labels\",\"type\":\"bytes32[]\"}],\"name\":\"eraseNode\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"startAuctions\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"deed\",\"type\":\"address\"},{\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"acceptRegistrarTransfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashes\",\"type\":\"bytes32[]\"},{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"name\":\"startAuctionsAndBid\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_rootNode\",\"type\":\"bytes32\"},{\"name\":\"_startDate\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"HashRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"HashReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"HashInvalidated\",\"type\":\"event\"}]"

// Registrarcontract is an auto generated Go binding around an Ethereum contract.
type Registrarcontract struct {
	RegistrarcontractCaller     // Read-only binding to the contract
	RegistrarcontractTransactor // Write-only binding to the contract
}

// RegistrarcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarcontractSession struct {
	Contract     *Registrarcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RegistrarcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarcontractCallerSession struct {
	Contract *RegistrarcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RegistrarcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarcontractTransactorSession struct {
	Contract     *RegistrarcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RegistrarcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarcontractRaw struct {
	Contract *Registrarcontract // Generic contract binding to access the raw methods on
}

// RegistrarcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarcontractCallerRaw struct {
	Contract *RegistrarcontractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarcontractTransactorRaw struct {
	Contract *RegistrarcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrarcontract creates a new instance of Registrarcontract, bound to a specific deployed contract.
func NewRegistrarcontract(address common.Address, backend bind.ContractBackend) (*Registrarcontract, error) {
	contract, err := bindRegistrarcontract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrarcontract{RegistrarcontractCaller: RegistrarcontractCaller{contract: contract}, RegistrarcontractTransactor: RegistrarcontractTransactor{contract: contract}}, nil
}

// NewRegistrarcontractCaller creates a new read-only instance of Registrarcontract, bound to a specific deployed contract.
func NewRegistrarcontractCaller(address common.Address, caller bind.ContractCaller) (*RegistrarcontractCaller, error) {
	contract, err := bindRegistrarcontract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarcontractCaller{contract: contract}, nil
}

// NewRegistrarcontractTransactor creates a new write-only instance of Registrarcontract, bound to a specific deployed contract.
func NewRegistrarcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarcontractTransactor, error) {
	contract, err := bindRegistrarcontract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RegistrarcontractTransactor{contract: contract}, nil
}

// bindRegistrarcontract binds a generic wrapper to an already deployed contract.
func bindRegistrarcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrarcontract *RegistrarcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrarcontract.Contract.RegistrarcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrarcontract *RegistrarcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrarcontract.Contract.RegistrarcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrarcontract *RegistrarcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrarcontract.Contract.RegistrarcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrarcontract *RegistrarcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrarcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrarcontract *RegistrarcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrarcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrarcontract *RegistrarcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrarcontract.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_Registrarcontract *RegistrarcontractCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_Registrarcontract *RegistrarcontractSession) Ens() (common.Address, error) {
	return _Registrarcontract.Contract.Ens(&_Registrarcontract.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_Registrarcontract *RegistrarcontractCallerSession) Ens() (common.Address, error) {
	return _Registrarcontract.Contract.Ens(&_Registrarcontract.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(_hash bytes32) constant returns(uint8, address, uint256, uint256, uint256)
func (_Registrarcontract *RegistrarcontractCaller) Entries(opts *bind.CallOpts, _hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Registrarcontract.contract.Call(opts, out, "entries", _hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(_hash bytes32) constant returns(uint8, address, uint256, uint256, uint256)
func (_Registrarcontract *RegistrarcontractSession) Entries(_hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Registrarcontract.Contract.Entries(&_Registrarcontract.CallOpts, _hash)
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(_hash bytes32) constant returns(uint8, address, uint256, uint256, uint256)
func (_Registrarcontract *RegistrarcontractCallerSession) Entries(_hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Registrarcontract.Contract.Entries(&_Registrarcontract.CallOpts, _hash)
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(_hash bytes32) constant returns(timestamp uint256)
func (_Registrarcontract *RegistrarcontractCaller) GetAllowedTime(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "getAllowedTime", _hash)
	return *ret0, err
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(_hash bytes32) constant returns(timestamp uint256)
func (_Registrarcontract *RegistrarcontractSession) GetAllowedTime(_hash [32]byte) (*big.Int, error) {
	return _Registrarcontract.Contract.GetAllowedTime(&_Registrarcontract.CallOpts, _hash)
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(_hash bytes32) constant returns(timestamp uint256)
func (_Registrarcontract *RegistrarcontractCallerSession) GetAllowedTime(_hash [32]byte) (*big.Int, error) {
	return _Registrarcontract.Contract.GetAllowedTime(&_Registrarcontract.CallOpts, _hash)
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(_hash bytes32, _timestamp uint256) constant returns(allowed bool)
func (_Registrarcontract *RegistrarcontractCaller) IsAllowed(opts *bind.CallOpts, _hash [32]byte, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "isAllowed", _hash, _timestamp)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(_hash bytes32, _timestamp uint256) constant returns(allowed bool)
func (_Registrarcontract *RegistrarcontractSession) IsAllowed(_hash [32]byte, _timestamp *big.Int) (bool, error) {
	return _Registrarcontract.Contract.IsAllowed(&_Registrarcontract.CallOpts, _hash, _timestamp)
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(_hash bytes32, _timestamp uint256) constant returns(allowed bool)
func (_Registrarcontract *RegistrarcontractCallerSession) IsAllowed(_hash [32]byte, _timestamp *big.Int) (bool, error) {
	return _Registrarcontract.Contract.IsAllowed(&_Registrarcontract.CallOpts, _hash, _timestamp)
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_Registrarcontract *RegistrarcontractCaller) LaunchLength(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "launchLength")
	return *ret0, err
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_Registrarcontract *RegistrarcontractSession) LaunchLength() (uint32, error) {
	return _Registrarcontract.Contract.LaunchLength(&_Registrarcontract.CallOpts)
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_Registrarcontract *RegistrarcontractCallerSession) LaunchLength() (uint32, error) {
	return _Registrarcontract.Contract.LaunchLength(&_Registrarcontract.CallOpts)
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_Registrarcontract *RegistrarcontractCaller) RegistryStarted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "registryStarted")
	return *ret0, err
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_Registrarcontract *RegistrarcontractSession) RegistryStarted() (*big.Int, error) {
	return _Registrarcontract.Contract.RegistryStarted(&_Registrarcontract.CallOpts)
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_Registrarcontract *RegistrarcontractCallerSession) RegistryStarted() (*big.Int, error) {
	return _Registrarcontract.Contract.RegistryStarted(&_Registrarcontract.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_Registrarcontract *RegistrarcontractCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_Registrarcontract *RegistrarcontractSession) RootNode() ([32]byte, error) {
	return _Registrarcontract.Contract.RootNode(&_Registrarcontract.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_Registrarcontract *RegistrarcontractCallerSession) RootNode() ([32]byte, error) {
	return _Registrarcontract.Contract.RootNode(&_Registrarcontract.CallOpts)
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids( address,  bytes32) constant returns(address)
func (_Registrarcontract *RegistrarcontractCaller) SealedBids(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "sealedBids", arg0, arg1)
	return *ret0, err
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids( address,  bytes32) constant returns(address)
func (_Registrarcontract *RegistrarcontractSession) SealedBids(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _Registrarcontract.Contract.SealedBids(&_Registrarcontract.CallOpts, arg0, arg1)
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids( address,  bytes32) constant returns(address)
func (_Registrarcontract *RegistrarcontractCallerSession) SealedBids(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _Registrarcontract.Contract.SealedBids(&_Registrarcontract.CallOpts, arg0, arg1)
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(hash bytes32, owner address, value uint256, salt bytes32) constant returns(sealedBid bytes32)
func (_Registrarcontract *RegistrarcontractCaller) ShaBid(opts *bind.CallOpts, hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "shaBid", hash, owner, value, salt)
	return *ret0, err
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(hash bytes32, owner address, value uint256, salt bytes32) constant returns(sealedBid bytes32)
func (_Registrarcontract *RegistrarcontractSession) ShaBid(hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	return _Registrarcontract.Contract.ShaBid(&_Registrarcontract.CallOpts, hash, owner, value, salt)
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(hash bytes32, owner address, value uint256, salt bytes32) constant returns(sealedBid bytes32)
func (_Registrarcontract *RegistrarcontractCallerSession) ShaBid(hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	return _Registrarcontract.Contract.ShaBid(&_Registrarcontract.CallOpts, hash, owner, value, salt)
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(_hash bytes32) constant returns(uint8)
func (_Registrarcontract *RegistrarcontractCaller) State(opts *bind.CallOpts, _hash [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Registrarcontract.contract.Call(opts, out, "state", _hash)
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(_hash bytes32) constant returns(uint8)
func (_Registrarcontract *RegistrarcontractSession) State(_hash [32]byte) (uint8, error) {
	return _Registrarcontract.Contract.State(&_Registrarcontract.CallOpts, _hash)
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(_hash bytes32) constant returns(uint8)
func (_Registrarcontract *RegistrarcontractCallerSession) State(_hash [32]byte) (uint8, error) {
	return _Registrarcontract.Contract.State(&_Registrarcontract.CallOpts, _hash)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(hash bytes32, deed address, registrationDate uint256) returns()
func (_Registrarcontract *RegistrarcontractTransactor) AcceptRegistrarTransfer(opts *bind.TransactOpts, hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "acceptRegistrarTransfer", hash, deed, registrationDate)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(hash bytes32, deed address, registrationDate uint256) returns()
func (_Registrarcontract *RegistrarcontractSession) AcceptRegistrarTransfer(hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _Registrarcontract.Contract.AcceptRegistrarTransfer(&_Registrarcontract.TransactOpts, hash, deed, registrationDate)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(hash bytes32, deed address, registrationDate uint256) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) AcceptRegistrarTransfer(hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _Registrarcontract.Contract.AcceptRegistrarTransfer(&_Registrarcontract.TransactOpts, hash, deed, registrationDate)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(bidder address, seal bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) CancelBid(opts *bind.TransactOpts, bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "cancelBid", bidder, seal)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(bidder address, seal bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) CancelBid(bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.CancelBid(&_Registrarcontract.TransactOpts, bidder, seal)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(bidder address, seal bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) CancelBid(bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.CancelBid(&_Registrarcontract.TransactOpts, bidder, seal)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(labels bytes32[]) returns()
func (_Registrarcontract *RegistrarcontractTransactor) EraseNode(opts *bind.TransactOpts, labels [][32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "eraseNode", labels)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(labels bytes32[]) returns()
func (_Registrarcontract *RegistrarcontractSession) EraseNode(labels [][32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.EraseNode(&_Registrarcontract.TransactOpts, labels)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(labels bytes32[]) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) EraseNode(labels [][32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.EraseNode(&_Registrarcontract.TransactOpts, labels)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) FinalizeAuction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "finalizeAuction", _hash)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) FinalizeAuction(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.FinalizeAuction(&_Registrarcontract.TransactOpts, _hash)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) FinalizeAuction(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.FinalizeAuction(&_Registrarcontract.TransactOpts, _hash)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(unhashedName string) returns()
func (_Registrarcontract *RegistrarcontractTransactor) InvalidateName(opts *bind.TransactOpts, unhashedName string) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "invalidateName", unhashedName)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(unhashedName string) returns()
func (_Registrarcontract *RegistrarcontractSession) InvalidateName(unhashedName string) (*types.Transaction, error) {
	return _Registrarcontract.Contract.InvalidateName(&_Registrarcontract.TransactOpts, unhashedName)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(unhashedName string) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) InvalidateName(unhashedName string) (*types.Transaction, error) {
	return _Registrarcontract.Contract.InvalidateName(&_Registrarcontract.TransactOpts, unhashedName)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(sealedBid bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) NewBid(opts *bind.TransactOpts, sealedBid [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "newBid", sealedBid)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(sealedBid bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) NewBid(sealedBid [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.NewBid(&_Registrarcontract.TransactOpts, sealedBid)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(sealedBid bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) NewBid(sealedBid [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.NewBid(&_Registrarcontract.TransactOpts, sealedBid)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) ReleaseDeed(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "releaseDeed", _hash)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) ReleaseDeed(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.ReleaseDeed(&_Registrarcontract.TransactOpts, _hash)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) ReleaseDeed(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.ReleaseDeed(&_Registrarcontract.TransactOpts, _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) StartAuction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "startAuction", _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) StartAuction(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.StartAuction(&_Registrarcontract.TransactOpts, _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) StartAuction(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.StartAuction(&_Registrarcontract.TransactOpts, _hash)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(_hashes bytes32[]) returns()
func (_Registrarcontract *RegistrarcontractTransactor) StartAuctions(opts *bind.TransactOpts, _hashes [][32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "startAuctions", _hashes)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(_hashes bytes32[]) returns()
func (_Registrarcontract *RegistrarcontractSession) StartAuctions(_hashes [][32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.StartAuctions(&_Registrarcontract.TransactOpts, _hashes)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(_hashes bytes32[]) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) StartAuctions(_hashes [][32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.StartAuctions(&_Registrarcontract.TransactOpts, _hashes)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(hashes bytes32[], sealedBid bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) StartAuctionsAndBid(opts *bind.TransactOpts, hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "startAuctionsAndBid", hashes, sealedBid)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(hashes bytes32[], sealedBid bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) StartAuctionsAndBid(hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.StartAuctionsAndBid(&_Registrarcontract.TransactOpts, hashes, sealedBid)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(hashes bytes32[], sealedBid bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) StartAuctionsAndBid(hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.StartAuctionsAndBid(&_Registrarcontract.TransactOpts, hashes, sealedBid)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(_hash bytes32, newOwner address) returns()
func (_Registrarcontract *RegistrarcontractTransactor) Transfer(opts *bind.TransactOpts, _hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "transfer", _hash, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(_hash bytes32, newOwner address) returns()
func (_Registrarcontract *RegistrarcontractSession) Transfer(_hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Registrarcontract.Contract.Transfer(&_Registrarcontract.TransactOpts, _hash, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(_hash bytes32, newOwner address) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) Transfer(_hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Registrarcontract.Contract.Transfer(&_Registrarcontract.TransactOpts, _hash, newOwner)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) TransferRegistrars(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "transferRegistrars", _hash)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) TransferRegistrars(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.TransferRegistrars(&_Registrarcontract.TransactOpts, _hash)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(_hash bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) TransferRegistrars(_hash [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.TransferRegistrars(&_Registrarcontract.TransactOpts, _hash)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(_hash bytes32, _value uint256, _salt bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactor) UnsealBid(opts *bind.TransactOpts, _hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.contract.Transact(opts, "unsealBid", _hash, _value, _salt)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(_hash bytes32, _value uint256, _salt bytes32) returns()
func (_Registrarcontract *RegistrarcontractSession) UnsealBid(_hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.UnsealBid(&_Registrarcontract.TransactOpts, _hash, _value, _salt)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(_hash bytes32, _value uint256, _salt bytes32) returns()
func (_Registrarcontract *RegistrarcontractTransactorSession) UnsealBid(_hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _Registrarcontract.Contract.UnsealBid(&_Registrarcontract.TransactOpts, _hash, _value, _salt)
}
