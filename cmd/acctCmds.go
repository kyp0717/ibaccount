package cmd

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	ib "ibx/webapi"

	"github.com/spf13/cobra"
)

// Get list of accounts

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var AcctListCmd = &cobra.Command{
	Use:   "list",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		ib.GetAcctList(ib.UrlAcctList)

	},
}

var AcctSetCmd = &cobra.Command{
	Use:   "set",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		ib.SetAcct(ib.UrlAcctList)

	},
}

// Get an object containing PnL for selected account
var getPnLCmd = &cobra.Command{
	Use:   "pnl",
	Short: "retrieve pnl",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call pnl cmd:  ibx accounts pnl")

	},
}

func init() {
	rootCmd.AddCommand(AccountCmd)
	rootCmd.AddCommand(setAcctCmd)
	AccountCmd.AddCommand(getAcctListCmd)
}

func getAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbGet(url)
	return data, err
}

func setAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbPost(url)
	return data, err
}
