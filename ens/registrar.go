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
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/registrarcontract"
)

// RegistrarContract obtains the registrar contract for a chain
func RegistrarContract(client *ethclient.Client, rpcclient *rpc.Client) (registrar *registrarcontract.Registrarcontract, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//chainID, err := client.NetworkID(ctx)
	chainID, err := etherutils.NetworkID(ctx, rpcclient)
	if err != nil {
		return nil, err
	}

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

// CreateRegistrarSession creates a session suitable for multiple calls
func CreateRegistrarSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *registrarcontract.Registrarcontract, gasLimit *big.Int, gasPrice *big.Int) *registrarcontract.RegistrarcontractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &registrarcontract.RegistrarcontractSession{
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

// SealBid seals the elements of a bid in to a single hash
func SealBid(name string, owner *common.Address, amount big.Int, salt string) (hash common.Hash, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}
	domainHash := LabelHash(domain)

	sha := sha3.NewKeccak256()
	sha.Write(domainHash[:])
	sha.Write(owner.Bytes())
	// Amount needs to be exactly 32 bytes
	var amountBytes [32]byte
	copy(amountBytes[len(amountBytes)-len(amount.Bytes()):], amount.Bytes()[:])
	sha.Write(amountBytes[:])
	saltHash := LabelHash(salt)
	sha.Write(saltHash[:])
	sha.Sum(hash[:0])
	return
}

// StartAuctionAndBid starts an auction and bids in the same transaction
func StartAuctionAndBid(session *registrarcontract.RegistrarcontractSession, name string, owner *common.Address, amount big.Int, salt string) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	sealedBid, err := SealBid(name, owner, amount, salt)
	if err != nil {
		return
	}

	var domainHashes [][32]byte
	domainHashes = make([][32]byte, 0, 1)
	domainHashes = append(domainHashes, LabelHash(domain))
	tx, err = session.StartAuctionsAndBid(domainHashes, sealedBid)
	return
}

// NewBid bids on an existing auction
func NewBid(session *registrarcontract.RegistrarcontractSession, name string, owner *common.Address, amount big.Int, salt string) (tx *types.Transaction, err error) {
	sealedBid, err := SealBid(name, owner, amount, salt)
	if err != nil {
		return
	}

	tx, err = session.NewBid(sealedBid)
	return
}

// RevealBid reveals an existing bid on an existing auction
func RevealBid(session *registrarcontract.RegistrarcontractSession, name string, owner *common.Address, amount big.Int, salt string) (tx *types.Transaction, err error) {
	tx, err = session.UnsealBid(LabelHash(name), &amount, LabelHash(salt))
	return
}

// FinishAuction reveals an existing bid on an existing auction
func FinishAuction(session *registrarcontract.RegistrarcontractSession, name string) (tx *types.Transaction, err error) {
	tx, err = session.FinalizeAuction(LabelHash(name))
	return
}

// State obains the current state of a name
func State(contract *registrarcontract.Registrarcontract, name string) (state string, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	status, err := contract.State(nil, LabelHash(domain))
	if err != nil {
		return
	}
	switch status {
	case 0:
		state = "Available"
	case 1:
		state = "Bidding"
	case 2:
		state = "Owned"
	case 3:
		state = "Forbidden"
	case 4:
		state = "Revealing"
	case 5:
		state = "Unavailable"
	default:
		state = "Unknown"
	}

	return
}

// NameInState checks if a name is in a given state, and errors if not.
func NameInState(contract *registrarcontract.Registrarcontract, name string, desiredState string) (inState bool, err error) {
	state, err := State(contract, name)
	if err == nil {
		if state == desiredState {
			inState = true
		} else {
			switch state {
			case "Available":
				err = errors.New("this name has not been auctioned")
			case "Bidding":
				err = errors.New("this name is being auctioned")
			case "Owned":
				err = errors.New("this name is owned")
			case "Forbidden":
				err = errors.New("this name is unavailable")
			case "Revealing":
				err = errors.New("this name is being revealed")
			case "Unavailable":
				err = errors.New("this name is not yet available")
			default:
				err = errors.New("this name is in an unknown state")
			}
		}
	}
	return
}
