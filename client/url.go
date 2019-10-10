package client

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:5000"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath    string = "/v1/portal"
	AcctsPt                   = "/iserver/accounts"
	LedgerPt                  = "/portfolio/{accountId}/ledger"
	MetaPt                    = "/portfolio/{accountId}/meta"
	SummaryPt                 = "/portfolio/{accountId}/summary"
	PortfolioAcctPt           = "/portfolio/accounts"
	PortfolioSubAcctPt        = "/portfolio/subaccounts"
	AcctPt                    = "/iserver/account"

	// DefaultSchemes are the default schemes found in Meta (info) section of spec file
	DefaultScheme = "https"
)

const BaseURL = "https://localhost:5000/v1/portal"

// Account API - 8 endpoints total
// Get Brokerage Acct
const AcctURL = "/iserver/accounts"
