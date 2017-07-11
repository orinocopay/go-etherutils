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

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
)

func main() {
	ownerAddressPtr := flag.String("o", "", "The address that owns the name")
	resolutionAddressPtr := flag.String("r", "", "The address to which the name should resolve (when setting)")
	connectionPtr := flag.String("c", "https://api.orinocopay.com:8546/", "Ethereum connection")
	passphrasePtr := flag.String("pp", "", "The passphrase to unlock the owning account (when setting)")
	quietPtr := flag.Bool("q", false, "Quiet mode")
	flag.Parse()
	quiet := *quietPtr
	passphrase := *passphrasePtr

	// Obtain the name we are working with
	cli.Assert(len(flag.Args()) == 1, quiet, "Usage: ens [-r resolution address] name")
	name := flag.Arg(0)

	// Create a connection to an Ethereum node
	client, err := ethclient.Dial(*connectionPtr)
	cli.ErrCheck(err, quiet, "Failed to connect to Ethereum")

	if *resolutionAddressPtr == "" {
		query(quiet, client, name)
	} else {
		resolutionAddress, err := ens.Resolve(client, *resolutionAddressPtr)
		cli.ErrCheck(err, quiet, "Unknown resolution address")
		ownerAddress, err := ens.Resolve(client, *ownerAddressPtr)
		cli.ErrCheck(err, quiet, "Unknown owner address")
		err = set(quiet, client, ownerAddress, name, resolutionAddress, passphrase)
		cli.ErrCheck(err, quiet, "Failed to set resolution")
	}
}

func query(quiet bool, client *ethclient.Client, name string) {
	result, err := ens.Resolve(client, name)
	cli.ErrCheck(err, quiet, "Failed to obtain address")
	if !quiet {
		fmt.Println(result.Hex())
	}
}

func set(quiet bool, client *ethclient.Client, ownerAddress common.Address, name string, resolutionAddress common.Address, passphrase string) (err error) {
	// Obtain the wallet and account
	wallet, err := cli.ObtainWallet(ownerAddress)
	cli.ErrCheck(err, quiet, "Failed to obtain an account for the owner address")
	account, err := cli.ObtainAccount(wallet, ownerAddress, passphrase)
	cli.ErrCheck(err, quiet, "Failed to obtain an account for the owner address")

	// Obtain the network ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return
	}

	// TODO where to fetch these?
	gasLimit := big.NewInt(500000)
	gasPrice := big.NewInt(21000000000)

	// Fetch the registry contract
	registryContract, err := ens.RegistryContract(chainID, client)
	// Set up our session
	registrySession := ens.CreateRegistrySession(chainID, &wallet, account, passphrase, registryContract, gasLimit, gasPrice)
	if err != nil {
		// No registry
		return
	}

	// Fetch the resolver
	resolverAddress, err := ens.Resolver(registryContract, name)
	if err != nil {
		// No resolver; create one
		resolverAddress, err = ens.PublicResolver(chainID, client)
		if err != nil {
			// No public resolver
			return errors.New("No public resolver known for that network")
		}
		tx, err := ens.SetResolver(registrySession, name, &resolverAddress)
		if err != nil {
			return err
		}
		if !quiet {
			fmt.Println("SetResolver transaction is", tx.Hash().Hex())
		}
	}

	// Set the address to which we resolve
	resolverContract, err := ens.ResolverContractByAddress(client, resolverAddress)
	if err != nil {
		return
	}
	resolverSession := ens.CreateResolverSession(chainID, &wallet, account, passphrase, resolverContract, gasLimit, gasPrice)
	tx, err := ens.SetResolution(resolverSession, name, &resolutionAddress)
	if err != nil {
		return
	}
	if !quiet {
		fmt.Println("SetResolution transaction is", tx.Hash().Hex())
	}

	return
}
