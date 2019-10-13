package cmd

import (
	"fmt"
	"ibaccount/models"
	"github.com/spf13/cobra"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
)


type AccountList struct {
	// Response Struct for endpt ....
	Accounts []string `json:"accounts"`
	Aliases  struct {
	} `json:"aliases"`
	SelectedAccount string `json:"selectedAccount"`
}

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

// return a list of accounts based on the login session
func GetAccts() (models.AccountList, error) {
	url := BaseURL + AcctURL
	data, _ := IbGet(url)
	var accts models.AccountList
	json.Unmarshal([]byte(data), &accts)
	return accts, nil
}

func (a models.AccountList) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}

func (a IbAccounts) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}
