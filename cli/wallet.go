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

package cli

import (
	"errors"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

// ObtainWallet fetches the wallet for a given address
func ObtainWallet(chainID *big.Int, address common.Address) (accounts.Wallet, error) {
	keydir := node.DefaultDataDir()
	if chainID.Cmp(params.MainnetChainConfig.ChainId) == 0 {
		// Nothing to add for mainnet
	} else if chainID.Cmp(params.TestnetChainConfig.ChainId) == 0 {
		keydir = filepath.Join(keydir, "testnet")
	} else if chainID.Cmp(params.RinkebyChainConfig.ChainId) == 0 {
		keydir = filepath.Join(keydir, "rinkeby")
	}
	keydir = filepath.Join(keydir, "keystore")
	backends := []accounts.Backend{keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)}
	accountManager := accounts.NewManager(backends...)
	defer accountManager.Close()
	account := accounts.Account{Address: address}
	wallet, err := accountManager.Find(account)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// ObtainAccount fetches the account for a given address
func ObtainAccount(wallet accounts.Wallet, address common.Address, passphrase string) (*accounts.Account, error) {
	for _, account := range wallet.Accounts() {
		if address == account.Address {
			if !VerifyPassphrase(wallet, account, passphrase) {
				return nil, errors.New("invalid passphrase")
			}
			return &account, nil
		}
	}
	return nil, errors.New("account not found")
}

// VerifyPassphrase confirms that a passphrase is correct for an account
func VerifyPassphrase(wallet accounts.Wallet, account accounts.Account, passphrase string) bool {

	_, err := wallet.SignHashWithPassphrase(account, passphrase, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return err == nil
}
