package hubspot

import "fmt"

const (
	accountBasePath = "account-info"
)

type Account struct {
	Info AccountInfoService
}

func newAccount(c *Client) *Account {
	accountPath := fmt.Sprintf("%s/%s", accountBasePath, c.apiVersion)
	return &Account{
		Info: &AccountInfoServiceOp{
			accountInfoPath: fmt.Sprintf("%s", accountPath),
			client:          c,
		},
	}
}
