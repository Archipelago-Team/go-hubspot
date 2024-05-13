package hubspot

import "fmt"

const (
	crmBasePath = "crm"

	objectsBasePath = "objects"
)

type CRM struct {
	Contact    ContactService
	Company    CompanyService
	Deal       DealService
	Imports    CrmImportsService
	Exports    CrmExportsService
	Schemas    CrmSchemasService
	Properties CrmPropertiesService
	Tickets    CrmTicketsService
	List       CrmListService
}

func newCRM(c *Client) *CRM {
	crmPath := fmt.Sprintf("%s/%s", crmBasePath, c.apiVersion)
	return &CRM{
		Contact: &ContactServiceOp{
			contactPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, contactBasePath),
			client:      c,
		},
		Company: &CompanyServiceOp{
			companyPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, companyBasePath),
			client:      c,
		},
		Deal: &DealServiceOp{
			dealPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, dealBasePath),
			client:   c,
		},
		Imports: &CrmImportsServiceOp{
			crmImportsPath: fmt.Sprintf("%s/%s", crmPath, crmImportsBasePath),
			client:         c,
		},
		Exports: &CrmExportsServiceOp{
			crmExportsPath: fmt.Sprintf("%s/%s", crmPath, crmExportsBasePath),
			client:         c,
		},
		Schemas: &CrmSchemasServiceOp{
			crmSchemasPath: fmt.Sprintf("%s/%s", crmPath, crmSchemasPath),
			client:         c,
		},
		Properties: &CrmPropertiesServiceOp{
			crmPropertiesPath: fmt.Sprintf("%s/%s", crmPath, crmPropertiesPath),
			client:            c,
		},
		Tickets: &CrmTicketsServiceOp{
			crmTicketsPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, crmTicketsBasePath),
			client:         c,
		},
		List: &CrmListServiceOp{
			crmListPath: fmt.Sprintf("%s/%s", crmPath, crmListBasePath),
			client:      c,
		},
	}
}
