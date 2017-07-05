// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package enscontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EnscontractABI is the input ABI used to generate the binding from.
const EnscontractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"}]"

// Enscontract is an auto generated Go binding around an Ethereum contract.
type Enscontract struct {
	EnscontractCaller     // Read-only binding to the contract
	EnscontractTransactor // Write-only binding to the contract
}

// EnscontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnscontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnscontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnscontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnscontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnscontractSession struct {
	Contract     *Enscontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnscontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnscontractCallerSession struct {
	Contract *EnscontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EnscontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnscontractTransactorSession struct {
	Contract     *EnscontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EnscontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnscontractRaw struct {
	Contract *Enscontract // Generic contract binding to access the raw methods on
}

// EnscontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnscontractCallerRaw struct {
	Contract *EnscontractCaller // Generic read-only contract binding to access the raw methods on
}

// EnscontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnscontractTransactorRaw struct {
	Contract *EnscontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnscontract creates a new instance of Enscontract, bound to a specific deployed contract.
func NewEnscontract(address common.Address, backend bind.ContractBackend) (*Enscontract, error) {
	contract, err := bindEnscontract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Enscontract{EnscontractCaller: EnscontractCaller{contract: contract}, EnscontractTransactor: EnscontractTransactor{contract: contract}}, nil
}

// NewEnscontractCaller creates a new read-only instance of Enscontract, bound to a specific deployed contract.
func NewEnscontractCaller(address common.Address, caller bind.ContractCaller) (*EnscontractCaller, error) {
	contract, err := bindEnscontract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EnscontractCaller{contract: contract}, nil
}

// NewEnscontractTransactor creates a new write-only instance of Enscontract, bound to a specific deployed contract.
func NewEnscontractTransactor(address common.Address, transactor bind.ContractTransactor) (*EnscontractTransactor, error) {
	contract, err := bindEnscontract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EnscontractTransactor{contract: contract}, nil
}

// bindEnscontract binds a generic wrapper to an already deployed contract.
func bindEnscontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnscontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enscontract *EnscontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Enscontract.Contract.EnscontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enscontract *EnscontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enscontract.Contract.EnscontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enscontract *EnscontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enscontract.Contract.EnscontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enscontract *EnscontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Enscontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enscontract *EnscontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enscontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enscontract *EnscontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enscontract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_Enscontract *EnscontractCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enscontract.contract.Call(opts, out, "owner", node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_Enscontract *EnscontractSession) Owner(node [32]byte) (common.Address, error) {
	return _Enscontract.Contract.Owner(&_Enscontract.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_Enscontract *EnscontractCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _Enscontract.Contract.Owner(&_Enscontract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_Enscontract *EnscontractCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enscontract.contract.Call(opts, out, "resolver", node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_Enscontract *EnscontractSession) Resolver(node [32]byte) (common.Address, error) {
	return _Enscontract.Contract.Resolver(&_Enscontract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_Enscontract *EnscontractCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _Enscontract.Contract.Resolver(&_Enscontract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_Enscontract *EnscontractCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Enscontract.contract.Call(opts, out, "ttl", node)
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_Enscontract *EnscontractSession) Ttl(node [32]byte) (uint64, error) {
	return _Enscontract.Contract.Ttl(&_Enscontract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_Enscontract *EnscontractCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _Enscontract.Contract.Ttl(&_Enscontract.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_Enscontract *EnscontractTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Enscontract.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_Enscontract *EnscontractSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Enscontract.Contract.SetOwner(&_Enscontract.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_Enscontract *EnscontractTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Enscontract.Contract.SetOwner(&_Enscontract.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Enscontract *EnscontractTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Enscontract.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Enscontract *EnscontractSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Enscontract.Contract.SetResolver(&_Enscontract.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Enscontract *EnscontractTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Enscontract.Contract.SetResolver(&_Enscontract.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_Enscontract *EnscontractTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Enscontract.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_Enscontract *EnscontractSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Enscontract.Contract.SetSubnodeOwner(&_Enscontract.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_Enscontract *EnscontractTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Enscontract.Contract.SetSubnodeOwner(&_Enscontract.TransactOpts, node, label, owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Enscontract *EnscontractTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Enscontract.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Enscontract *EnscontractSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Enscontract.Contract.SetTTL(&_Enscontract.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Enscontract *EnscontractTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Enscontract.Contract.SetTTL(&_Enscontract.TransactOpts, node, ttl)
}
