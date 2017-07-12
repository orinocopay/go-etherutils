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
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/orinocopay/go-etherutils/ens/registrarcontract"
)

// RegistrarContract obtains the registrar contract for a chain
func RegistrarContract(chainID *big.Int, client *ethclient.Client) (registrar *registrarcontract.Registrarcontract, err error) {
	// Instantiate the registrar contract
	if chainID.Cmp(params.MainnetChainConfig.ChainId) == 0 {
		registrar, err = registrarcontract.NewRegistrarcontract(common.HexToAddress("6090A6e47849629b7245Dfa1Ca21D94cd15878Ef"), client)
	} else if chainID.Cmp(params.TestnetChainConfig.ChainId) == 0 {
		registrar, err = registrarcontract.NewRegistrarcontract(common.HexToAddress("c19fd9004b5c9789391679de6d766b981db94610"), client)
	} else if chainID.Cmp(params.RinkebyChainConfig.ChainId) == 0 {
		registrar, err = registrarcontract.NewRegistrarcontract(common.HexToAddress("21397c1a1f4acd9132fe36df011610564b87e24b"), client)
	} else {
		err = errors.New("Unknown network ID")
	}
	return
}

// NameInState checks if a name is in agiven state, and errors if not.
func NameInState(contract *registrarcontract.Registrarcontract, name string, desiredState uint8) (inState bool, err error) {
	// Need the second-level domain name (e.g. bar in foo.bar.eth)
	nameBits := strings.Split(name, ".")
	if len(nameBits) < 2 {
		err = errors.New("invalid name")
		return
	}

	state, err := contract.State(nil, LabelHash(nameBits[len(nameBits)-2]))
	if err == nil {
		if state == desiredState {
			inState = true
		} else {
			switch state {
			case 0:
				err = errors.New("this name has not been auctioned")
			case 1:
				err = errors.New("this name is being auctioned")
			case 2:
				err = errors.New("this name is owned")
			case 3:
				err = errors.New("this name is unavailable")
			case 4:
				err = errors.New("this name is being revealed")
			case 5:
				err = errors.New("this name is not yet available")
			default:
				err = errors.New("this name is in an unknown state")
			}
		}
	}
	return
}
