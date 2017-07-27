// Copyright 2017 Orinoco Payments
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ens

import (
	"bytes"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/resolvercontract"
)

var zeroHash = make([]byte, 32)

// UnknownAddress is the address to which unknown entries resolve
var UnknownAddress = common.HexToAddress("00")

// PublicResolver obtains the public resolver for a chain
func PublicResolver(client *ethclient.Client, rpcclient *rpc.Client) (address common.Address, err error) {
	address, err = resolverAddress(client, "resolver.eth", rpcclient)

	return
}

func resolverAddress(client *ethclient.Client, name string, rpcclient *rpc.Client) (address common.Address, err error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return
	}

	registryContract, err := RegistryContract(client, rpcclient)
	if err != nil {
		return
	}

	// Check that this name is owned
	ownerAddress, err := registryContract.Owner(nil, nameHash)
	if err != nil {
		return
	}
	if bytes.Compare(ownerAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("unregistered name")
		return
	}

	// Obtain the resolver address for this name
	address, err = registryContract.Resolver(nil, nameHash)
	if err != nil {
		return
	}
	if bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
		return
	}

	return
}

// Resolve resolves an ENS name in to an Etheruem address
// This will return an error if the name is not found or otherwise 0
func Resolve(client *ethclient.Client, input string, rpcclient *rpc.Client) (address common.Address, err error) {
	if strings.HasSuffix(input, ".eth") {
		return resolveName(client, input, rpcclient)
	} else {
		address = common.HexToAddress(input)
		if address == UnknownAddress {
			err = errors.New("could not parse address")
		}
	}

	return
}

func resolveName(client *ethclient.Client, input string, rpcclient *rpc.Client) (address common.Address, err error) {
	var nameHash [32]byte
	nameHash, err = NameHash(input)
	if err != nil {
		return
	}
	if bytes.Compare(nameHash[:], zeroHash) == 0 {
		err = errors.New("Bad name")
	} else {
		address, err = resolveHash(client, input, rpcclient)
	}
	return
}

func resolveHash(client *ethclient.Client, name string, rpcclient *rpc.Client) (address common.Address, err error) {
	contract, err := ResolverContract(client, name, rpcclient)
	if err != nil {
		return UnknownAddress, err
	}

	// Resolve the name
	nameHash, err := NameHash(name)
	if err != nil {
		return
	}
	address, err = contract.Addr(nil, nameHash)
	if err != nil {
		return UnknownAddress, err
	}
	if bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		return UnknownAddress, errors.New("no address")
	}

	return
}

// CreateResolverSession creates a session suitable for multiple calls
func CreateResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *resolvercontract.ResolverContract, gasLimit *big.Int, gasPrice *big.Int) *resolvercontract.ResolverContractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &resolvercontract.ResolverContractSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     account.Address,
			Signer:   signer,
			GasPrice: gasPrice,
			GasLimit: gasLimit,
		},
	}

	return session
}

// SetResolution sets the address to which a name resolves
func SetResolution(session *resolvercontract.ResolverContractSession, name string, resolutionAddress *common.Address) (tx *types.Transaction, err error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return
	}
	tx, err = session.SetAddr(nameHash, *resolutionAddress)
	return
}

// ResolverContractByAddress instantiates the resolver contract at aspecific address
func ResolverContractByAddress(client *ethclient.Client, resolverAddress common.Address) (resolver *resolvercontract.ResolverContract, err error) {
	// Instantiate the resolver contract
	resolver, err = resolvercontract.NewResolverContract(resolverAddress, client)

	return
}

// ResolverContract obtains the resolver contract for a name
func ResolverContract(client *ethclient.Client, name string, rpcclient *rpc.Client) (resolver *resolvercontract.ResolverContract, err error) {
	resolverAddress, err := resolverAddress(client, name, rpcclient)
	if err != nil {
		return
	}
	if bytes.Compare(resolverAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
		return
	}

	resolver, err = ResolverContractByAddress(client, resolverAddress)
	return
}
