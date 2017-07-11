package etherutils

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// KeySigner generates a signer using a private key
//func KeySigner() (signerfn bind.SignerFn) {
//	signerfn = func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
//		if address != keyAddr {
//			return nil, errors.New("not authorized to sign this account")
//		}
//		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
//		if err != nil {
//			return nil, err
//		}
//		return tx.WithSignature(signer, signature)
//	}
//
//	return
//}

// AccountSigner generates a signer using an account
func AccountSigner(chainId *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string) (signerfn bind.SignerFn) {
	signerfn = func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != account.Address {
			return nil, errors.New("not authorized to sign this account")
		}
		return (*wallet).SignTxWithPassphrase(*account, passphrase, tx, chainId)
	}
	return
}

//func KeySigner() (signerfn bind.SignerFn) {
//	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
//
//	signerfn = func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
//		if address != keyAddr {
//			return nil, errors.New("not authorized to sign this account")
//		}
//		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
//		if err != nil {
//			return nil, err
//		}
//		return tx.WithSignature(signer, signature)
//	}
//
//	return
//}
