// Copyright © 2017 Orinoco Payments
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

package cmd

import (
	"fmt"
	"math/big"

	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

var auctionFinishPassphrase string
var auctionFinishAddressStr string
var auctionFinishGasPriceStr string

// auctionFinishCmd represents the auction reveal command
var auctionFinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish an auction for an ENS name",
	Long: `Finish an auction for a name with the Ethereum Name Service (ENS).  For example:

    ens auction finish --address=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase="my secret passphrase" enstest.eth

The keystore for the address must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to finish the auction is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Ensure that the name is in a suitable state
		registrarContract, err := ens.RegistrarContract(client, rpcclient)
		inState, err := ens.NameInState(registrarContract, args[0], "Revealing")
		cli.ErrAssert(inState, err, quiet, "Name not in a suitable state for bid to be revealed")

		// Fetch the wallet and account for the address
		auctionFinishAddress, err := ens.Resolve(client, auctionFinishAddressStr, rpcclient)
		cli.ErrCheck(err, quiet, "Failed to obtain auction address")
		wallet, err := cli.ObtainWallet(chainID, auctionFinishAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain a wallet for the address")
		account, err := cli.ObtainAccount(wallet, auctionFinishAddress, auctionFinishPassphrase)
		cli.ErrCheck(err, quiet, "Failed to obtain an account for the address")

		gasLimit := big.NewInt(500000)
		gasPrice, err := etherutils.StringToWei(auctionFinishGasPriceStr)
		cli.ErrCheck(err, quiet, "Invalid gas price")

		// Set up our session
		session := ens.CreateRegistrarSession(chainID, &wallet, account, auctionFinishPassphrase, registrarContract, gasLimit, gasPrice)

		// Finish the bid
		tx, err := ens.FinishAuction(session, args[0])
		cli.ErrCheck(err, quiet, "Failed to send transaction")
		if !quiet {
			fmt.Println("Transaction ID is", tx.Hash().Hex())
		}
	},
}

func init() {
	auctionCmd.AddCommand(auctionFinishCmd)

	auctionFinishCmd.Flags().StringVarP(&auctionFinishPassphrase, "passphrase", "p", "", "Passphrase for the account that owns the bidding address")
	auctionFinishCmd.Flags().StringVarP(&auctionFinishAddressStr, "address", "a", "", "Address doing the bidding")
	auctionFinishCmd.Flags().StringVarP(&auctionFinishGasPriceStr, "gasprice", "g", "20 GWei", "Gas price for the transaction")
}
