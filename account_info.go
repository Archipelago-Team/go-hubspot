package hubspot

import "fmt"

const (
	accountInfoBasePath = "account-info"
)

type AccountInfoService interface {
	Details() (*AccountInfo, error)
}

type AccountInfoServiceOp struct {
	client          *Client
	accountInfoPath string
}

type AccountInfo struct {
	PortalId HsInt `json:"portalId"`
	UiDomain HsStr `json:"uiDomain"`
}

var _ AccountInfoService = (*AccountInfoServiceOp)(nil)

func (op *AccountInfoServiceOp) Details() (*AccountInfo, error) {
	var respone AccountInfo
	path := fmt.Sprintf("%s/%s", op.accountInfoPath, "details")
	err := op.client.Get(path, &respone, nil)
	if err != nil {
		return nil, err
	}
	return &respone, nil
}
