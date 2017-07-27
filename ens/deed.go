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
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/deedcontract"
)

// DeedContract obtains the deed contract at a particular address
func DeedContract(client *ethclient.Client, address *common.Address) (deed *deedcontract.DeedContract, err error) {
	deed, err = deedcontract.NewDeedContract(*address, client)
	return
}

// CreateDeedSession creates a session suitable for multiple calls
func CreateDeedSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *deedcontract.DeedContract, gasLimit *big.Int, gasPrice *big.Int) *deedcontract.DeedContractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &deedcontract.DeedContractSession{
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
