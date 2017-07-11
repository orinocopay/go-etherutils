package cli

import (
	"errors"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/node"
)

// ObtainWallet fetches the wallet for a given address
func ObtainWallet(address common.Address) (accounts.Wallet, error) {
	// FIXME hard-coded name
	keydir := filepath.Join(node.DefaultDataDir(), "testnet", "keystore")
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
