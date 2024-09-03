/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github/renlulu/etherscansdk-go/internal"
)

// verifyproxycontractCmd represents the verifyproxycontract command
var verifyproxycontractCmd = &cobra.Command{
	Use:   "verifyproxycontract",
	Short: "Submits a proxy contract source code to Etherscan for verification",
	Long:  `Submits a proxy contract source code to Etherscan for verification.`,
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
		c := internal.NewContract(api, apiKey)
		resp, err := c.VerifyProxy(context.Background(), address)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	contractsCmd.AddCommand(verifyproxycontractCmd)
	verifyproxycontractCmd.Flags().StringP("address", "a", "", "Address of the contract")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyproxycontractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyproxycontractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
