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

// checkproxyverificationCmd represents the checkproxyverification command
var checkproxyverificationCmd = &cobra.Command{
	Use:   "checkproxyverification",
	Short: "Checking Proxy Contract Verification Submission Status",
	Long:  `Checking Proxy Contract Verification Submission Status.`,
	Run: func(cmd *cobra.Command, args []string) {
		api, _ := cmd.Flags().GetString("api")
		apiKey, err := cmd.Flags().GetString("apikey")
		if err != nil {
			panic(err)
		}
		guid, err := cmd.Flags().GetString("guid")
		if err != nil {
			panic(err)
		}
		c := internal.NewContract(api, apiKey)
		resp, err := c.CheckProxyVerification(context.Background(), guid)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	},
}

func init() {
	contractsCmd.AddCommand(checkproxyverificationCmd)
	checkproxyverificationCmd.Flags().StringP("guid", "g", "", "Global uid for Etherscan")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkproxyverificationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkproxyverificationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
