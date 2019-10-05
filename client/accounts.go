/******

*****/
package accounts

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
	"ibaccount/models"
)

func GetAccts() (models.IbAccounts, error) {
	url := BaseURL + AcctURL
	data, _ := IbGet(url)
	var accts models.IbAccounts
	json.Unmarshal([]byte(data), &accts)
	return accts, nil
}

func (a IbAccounts) Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(" Accounts: %s \n", cyan(a.Accounts))
}

func SetAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbPost(url)
	return data, err
}
