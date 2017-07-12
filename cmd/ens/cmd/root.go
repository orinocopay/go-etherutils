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
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var quiet bool
var connection string

var client *ethclient.Client
var chainID *big.Int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: persistentPreRun,
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	if cmd.Name() == "help" {
		// User just wants help
		return
	}

	// Ensure that the first argument is present
	if len(args) == 0 {
		cli.Err(quiet, "Thie command requires a name")
	}
	if args[0] == "" {
		cli.Err(quiet, "Thie command requires a name")
	}

	// Add '.eth' to the end of the name if not present
	if !strings.HasSuffix(args[0], ".eth") {
		args[0] += ".eth"
	}

	// Create a connection to an Ethereum node
	var err error
	client, err = ethclient.Dial(connection)
	cli.ErrCheck(err, quiet, "Failed to connect to Ethereum")
	// Fetch the chain ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainID, err = client.NetworkID(ctx)
	cli.ErrCheck(err, quiet, "Failed to obtain chain ID")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "no output")
	RootCmd.PersistentFlags().StringVarP(&connection, "connection", "c", "https://api.orinocopay.com:8546/", "path to the Ethereum connection")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
