// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package registrycontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RegistrycontractABI is the input ABI used to generate the binding from.
const RegistrycontractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"}]"

// Registrycontract is an auto generated Go binding around an Ethereum contract.
type Registrycontract struct {
	RegistrycontractCaller     // Read-only binding to the contract
	RegistrycontractTransactor // Write-only binding to the contract
}

// RegistrycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrycontractSession struct {
	Contract     *Registrycontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrycontractCallerSession struct {
	Contract *RegistrycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RegistrycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrycontractTransactorSession struct {
	Contract     *RegistrycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RegistrycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrycontractRaw struct {
	Contract *Registrycontract // Generic contract binding to access the raw methods on
}

// RegistrycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrycontractCallerRaw struct {
	Contract *RegistrycontractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrycontractTransactorRaw struct {
	Contract *RegistrycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrycontract creates a new instance of Registrycontract, bound to a specific deployed contract.
func NewRegistrycontract(address common.Address, backend bind.ContractBackend) (*Registrycontract, error) {
	contract, err := bindRegistrycontract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrycontract{RegistrycontractCaller: RegistrycontractCaller{contract: contract}, RegistrycontractTransactor: RegistrycontractTransactor{contract: contract}}, nil
}

// NewRegistrycontractCaller creates a new read-only instance of Registrycontract, bound to a specific deployed contract.
func NewRegistrycontractCaller(address common.Address, caller bind.ContractCaller) (*RegistrycontractCaller, error) {
	contract, err := bindRegistrycontract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrycontractCaller{contract: contract}, nil
}

// NewRegistrycontractTransactor creates a new write-only instance of Registrycontract, bound to a specific deployed contract.
func NewRegistrycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrycontractTransactor, error) {
	contract, err := bindRegistrycontract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RegistrycontractTransactor{contract: contract}, nil
}

// bindRegistrycontract binds a generic wrapper to an already deployed contract.
func bindRegistrycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrycontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrycontract *RegistrycontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrycontract.Contract.RegistrycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrycontract *RegistrycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrycontract.Contract.RegistrycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrycontract *RegistrycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrycontract.Contract.RegistrycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrycontract *RegistrycontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrycontract *RegistrycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrycontract *RegistrycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrycontract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_Registrycontract *RegistrycontractCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registrycontract.contract.Call(opts, out, "owner", node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_Registrycontract *RegistrycontractSession) Owner(node [32]byte) (common.Address, error) {
	return _Registrycontract.Contract.Owner(&_Registrycontract.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_Registrycontract *RegistrycontractCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _Registrycontract.Contract.Owner(&_Registrycontract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_Registrycontract *RegistrycontractCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registrycontract.contract.Call(opts, out, "resolver", node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_Registrycontract *RegistrycontractSession) Resolver(node [32]byte) (common.Address, error) {
	return _Registrycontract.Contract.Resolver(&_Registrycontract.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_Registrycontract *RegistrycontractCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _Registrycontract.Contract.Resolver(&_Registrycontract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_Registrycontract *RegistrycontractCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Registrycontract.contract.Call(opts, out, "ttl", node)
	return *ret0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_Registrycontract *RegistrycontractSession) Ttl(node [32]byte) (uint64, error) {
	return _Registrycontract.Contract.Ttl(&_Registrycontract.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(node bytes32) constant returns(uint64)
func (_Registrycontract *RegistrycontractCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _Registrycontract.Contract.Ttl(&_Registrycontract.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_Registrycontract *RegistrycontractTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Registrycontract.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_Registrycontract *RegistrycontractSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetOwner(&_Registrycontract.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_Registrycontract *RegistrycontractTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetOwner(&_Registrycontract.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Registrycontract *RegistrycontractTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Registrycontract.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Registrycontract *RegistrycontractSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetResolver(&_Registrycontract.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_Registrycontract *RegistrycontractTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetResolver(&_Registrycontract.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_Registrycontract *RegistrycontractTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Registrycontract.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_Registrycontract *RegistrycontractSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetSubnodeOwner(&_Registrycontract.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_Registrycontract *RegistrycontractTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetSubnodeOwner(&_Registrycontract.TransactOpts, node, label, owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Registrycontract *RegistrycontractTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Registrycontract.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Registrycontract *RegistrycontractSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetTTL(&_Registrycontract.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(node bytes32, ttl uint64) returns()
func (_Registrycontract *RegistrycontractTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Registrycontract.Contract.SetTTL(&_Registrycontract.TransactOpts, node, ttl)
}
