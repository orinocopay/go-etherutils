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
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/resolvercontract"
)

var zeroHash = make([]byte, 32)
var UnknownAddress = common.HexToAddress("00")

// PublicResolver obtains the public resolver for a chain
func PublicResolver(chainID *big.Int, client *ethclient.Client) (address common.Address, err error) {
	// Instantiate the registry contract
	if chainID.Cmp(params.MainnetChainConfig.ChainId) == 0 {
		address = common.HexToAddress("5FfC014343cd971B7eb70732021E26C35B744cc4")
	} else if chainID.Cmp(params.TestnetChainConfig.ChainId) == 0 {
		address = common.HexToAddress("4c641fb9bad9b60ef180c31f56051ce826d21a9a")
		// TODO does Rinkeby have a public resolver?
		//	} else if chainID.Cmp(params.RinkebyChainConfig.ChainId) == 0 {
		//		address = common.HexToAddress("")
	} else {
		err = errors.New("Unknown network ID")
	}
	return
}

// Resolve resolves an ENS name in to an Etheruem address
// This will return an error if the name is not found or otherwise 0
func Resolve(client *ethclient.Client, input string) (address common.Address, err error) {
	if strings.HasSuffix(input, ".eth") {
		nameHash := NameHash(input)
		if bytes.Compare(nameHash[:], zeroHash) == 0 {
			err = errors.New("Bad name")
		} else {
			address, err = resolveHash(client, input)
		}
	} else {
		address = common.HexToAddress(input)
		if address == UnknownAddress {
			err = errors.New("could not parse address")
		}
	}

	return
}

func resolveHash(client *ethclient.Client, name string) (address common.Address, err error) {
	contract, err := ResolverContract(client, name)
	if err != nil {
		return UnknownAddress, err
	}

	// Resolve the name
	address, err = contract.Addr(nil, NameHash(name))
	if err != nil {
		return UnknownAddress, err
	}
	if bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		return UnknownAddress, errors.New("no address")
	}

	return
}

// CreateResolverSession creates a session suitable for multiple calls
// TODO how to handle changes in gas limit?
func CreateResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *resolvercontract.Resolvercontract, gasLimit *big.Int, gasPrice *big.Int) *resolvercontract.ResolvercontractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &resolvercontract.ResolvercontractSession{
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
func SetResolution(session *resolvercontract.ResolvercontractSession, name string, resolutionAddress *common.Address) (tx *types.Transaction, err error) {
	nameHash := NameHash(name)
	tx, err = session.SetAddr(nameHash, *resolutionAddress)
	return
}

// ResolverContractByAddress instantiates the resolver contract at aspecific address
func ResolverContractByAddress(client *ethclient.Client, resolverAddress common.Address) (resolver *resolvercontract.Resolvercontract, err error) {
	// Instantiate the resolver contract
	resolver, err = resolvercontract.NewResolvercontract(resolverAddress, client)
	if err != nil {
		return nil, err
	}

	return
}

// ResolverContract obtains the resolver contract for a name
func ResolverContract(client *ethclient.Client, name string) (resolver *resolvercontract.Resolvercontract, err error) {
	nameHash := NameHash(name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	registryContract, err := RegistryContract(chainID, client)
	if err != nil {
		return nil, err
	}

	// Check that this name is owned
	ownerAddress, err := registryContract.Owner(nil, nameHash)
	if err != nil {
		return nil, err
	}
	if bytes.Compare(ownerAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		return nil, errors.New("unregistered name")
	}

	// Obtain the resolver for this name
	resolverAddress, err := registryContract.Resolver(nil, nameHash)
	if err != nil {
		return nil, err
	}
	if bytes.Compare(resolverAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		return nil, errors.New("no resolver")
	}

	return ResolverContractByAddress(client, resolverAddress)
}
