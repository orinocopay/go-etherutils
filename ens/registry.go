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
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/registrycontract"
)

// RegistryContract obtains the registry contract for a chain
func RegistryContract(client *ethclient.Client, rpcclient *rpc.Client) (registry *registrycontract.Registrycontract, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//chainID, err := client.NetworkID(ctx)
	chainID, err := etherutils.NetworkID(ctx, rpcclient)
	if err != nil {
		return nil, err
	}

	// Instantiate the registry contract
	if chainID.Cmp(params.MainnetChainConfig.ChainId) == 0 {
		registry, err = registrycontract.NewRegistrycontract(common.HexToAddress("314159265dd8dbb310642f98f50c066173c1259b"), client)
	} else if chainID.Cmp(params.TestnetChainConfig.ChainId) == 0 {
		registry, err = registrycontract.NewRegistrycontract(common.HexToAddress("112234455c3a32fd11230c42e7bccd4a84e02010"), client)
	} else if chainID.Cmp(params.RinkebyChainConfig.ChainId) == 0 {
		registry, err = registrycontract.NewRegistrycontract(common.HexToAddress("e7410170f87102DF0055eB195163A03B7F2Bff4A"), client)
	} else {
		err = errors.New("Unknown network ID")
	}
	return
}

// Resolver obtains the address of the resolver for a .eth name
func Resolver(contract *registrycontract.Registrycontract, name string) (address common.Address, err error) {
	address, err = contract.Resolver(nil, NameHash(name))
	if err == nil && bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
	}
	return
}

// SetResolver sets the resolver for a name
func SetResolver(session *registrycontract.RegistrycontractSession, name string, resolverAddr *common.Address) (tx *types.Transaction, err error) {
	// Set the resolver for this name
	nameHash := NameHash(name)
	tx, err = session.SetResolver(nameHash, *resolverAddr)
	return
}

// CreateRegistrySession creates a session suitable for multiple calls
// TODO how to handle changes in gas limit?
func CreateRegistrySession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *registrycontract.Registrycontract, gasLimit *big.Int, gasPrice *big.Int) *registrycontract.RegistrycontractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &registrycontract.RegistrycontractSession{
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
