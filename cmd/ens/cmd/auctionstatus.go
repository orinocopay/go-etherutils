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
	"os"
	"time"

	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

// auctionStatusCmd represents the auction status command
var auctionStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Obtain the status of an ENS auction",
	Long: `Obtain the status of an auction with the Ethereum Name Service (ENS).  For example:

    ens auction status enstest.eth

In quiet mode this will return 0 if the auction is currently bidding, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		registrarContract, err := ens.RegistrarContract(client, rpcclient)
		registrationDate, err := ens.RegistrationDate(registrarContract, args[0])
		cli.ErrCheck(err, quiet, "Cannot obtain auction status")

		twoDaysAgo := time.Duration(-48) * time.Hour
		fiveDaysAgo := time.Duration(-120) * time.Hour
		if quiet {
			if registrationDate.Add(twoDaysAgo).After(time.Now()) {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		} else {
			fmt.Println("     Auction:", registrationDate.Add(fiveDaysAgo))
			fmt.Println("      Reveal:", registrationDate.Add(twoDaysAgo))
			fmt.Println("Registration:", registrationDate)
		}
	},
}

func init() {
	auctionCmd.AddCommand(auctionStatusCmd)
}
