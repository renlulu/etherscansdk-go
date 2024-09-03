/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github/renlulu/etherscansdk-go/internal"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// verifysourcecodeCmd represents the verifysourcecode command
var verifysourcecodeCmd = &cobra.Command{
	Use:   "verifysourcecode",
	Short: "Submits a contract source code to an Etherscan-like explorer for verification",
	Long:  `Submits a contract source code to an Etherscan-like explorer for verification.`,
	Run: func(cmd *cobra.Command, args []string) {
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			panic(err)
		}
		api, _ := cmd.Flags().GetString("api")
		apiKey, err := cmd.Flags().GetString("apikey")
		if err != nil {
			panic(err)
		}
		sampleSourceCode, err := ioutil.ReadFile("testsuite/sample.json")
		if err != nil {
			panic(err)
		}
		c := internal.NewContract(api, apiKey)
		resp, err := c.VerifySourceCode(
			context.Background(),
			"17000",
			address,
			string(sampleSourceCode),
			"000000000000000000000000055733000064333CaDDbC92763c58BF0192fFeBf0000000000000000000000002dA62Eecc308Cfad483096fb62302Fb03f7730b8000000000000000000000000A44151489861Fe9e3055d95adC98FbD462B948e7",
			"HelloWorldTemplate.sol:HelloWorldTemplate",
			"v0.8.26+commit.8a97fa7a",
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	contractsCmd.AddCommand(verifysourcecodeCmd)
	verifysourcecodeCmd.Flags().StringP("address", "a", "", "Address of the contract")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifysourcecodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifysourcecodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
