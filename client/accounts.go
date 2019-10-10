/******

*****/
package client

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/fatih/color"
	"ibaccount/models"
)

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

func SetAccts() ([]byte, error) {
	url := BaseURL + AcctURL
	data, err := IbPost(url)
	return data, err
}
