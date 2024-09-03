/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github/renlulu/etherscansdk-go/internal"

	"github.com/spf13/cobra"
)

var address string

// getabiCmd represents the getabi command
var getabiCmd = &cobra.Command{
	Use:   "getabi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		abi, err := c.GetContractABI(address)
		if err != nil {
			panic(err)
		}
		fmt.Println(abi)
	},
}

func init() {
	contractsCmd.AddCommand(getabiCmd)
	getabiCmd.Flags().StringP("address", "a", "", "Address of the contract")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getabiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getabiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
