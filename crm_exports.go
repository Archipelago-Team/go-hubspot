package hubspot

import (
	"fmt"
)

const (
	crmExportsBasePath = "exports/export/async"
)

// CrmExportsService is an interface of CRM bulk export endpoints of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/crm/exports
type CrmExportsService interface {
	Get(int64) (interface{}, error)
	Start(*CrmExportConfig) (interface{}, error)
}

// CrmExportsServiceOp handles communication with the bulk CRM export endpoints of the HubSpot API.
type CrmExportsServiceOp struct {
	client         *Client
	crmExportsPath string
}

var _ CrmExportsService = (*CrmExportsServiceOp)(nil)

func (s *CrmExportsServiceOp) Get(exportId int64) (interface{}, error) {
	resource := make(map[string]interface{})
	path := fmt.Sprintf("%s/tasts/%d/status", s.crmExportsPath, exportId)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return resource, nil
}
