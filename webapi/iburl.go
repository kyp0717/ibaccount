package cmd

const (
	BaseURL   = "https://localhost:5000/v1/portal"
	AcctEndPt = "/iserver/accounts"
)

type IbAccunts struct {
	Accounts []string `json:"accounts"`
	Aliases  struct {
	} `json:"aliases"`
	SelectedAccount string `json:"selectedAccount"`
}
