package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	// "log"
	"crypto/tls"
	"encoding/json"
	"github.com/fatih/color"
	"net/http"
)

// getstatusCmd represents the getstatus command
var getAcctCmd = &cobra.Command{
	Use:   "accounts",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		getAccts()
	},
}

func init() {
	rootCmd.AddCommand(getAcctCmd)
}

func getAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbGet(url)
	return data, err
}

func (a IbAccounts) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}
