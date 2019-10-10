package cmd

import (

	//"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	ib "ibaccount/client"
)

var AcctSetCmd = &cobra.Command{
	Use:   "set",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		ib.SetAcct(ib.UrlAcctList)

	},
}

func init() {
	rootCmd.AddCommand(AccountCmd)
	rootCmd.AddCommand(setAcctCmd)
	AccountCmd.AddCommand(getAcctListCmd)
}

func setAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbPost(url)
	return data, err
}
