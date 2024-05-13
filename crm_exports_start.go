package hubspot

type CrmExportConfig struct {
	ExportType             string                 `json:"exportType,omitempty"`
	Format                 string                 `json:"format,omitempty"`
	ExportName             string                 `json:"exportName,omitempty"`
	Language               string                 `json:"language,omitempty"`
	ObjectType             string                 `json:"objectType,omitempty"`
	AssociatedObjectType   string                 `json:"associatedObjectType,omitempty"`
	ObjectProperties       []string               `json:"objectProperties,omitempty"`
	PublicCrmSearchRequest CrmExportSearchRequest `json:"publicCrmSearchRequest,omitempty"`
	ListId                 string                 `json:"listId,omitempty"`
}

type CrmExportFilter struct {
	PropertyName string `json:"propertyName,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Value        string `json:"value,omitempty"`
}
type CrmExportSort struct {
	PropertyName string `json:"propertyName,omitempty"`
	Order        string `json:"order,omitempty"`
}
type CrmExportSearchRequest struct {
	Filters []CrmExportFilter `json:"filters,omitempty"`
	Query   string            `json:"query,omitempty"`
	Sorts   []CrmExportSort   `json:"sorts,omitempty"`
}

func (s *CrmExportsServiceOp) Start(exportRequest *CrmExportConfig) (interface{}, error) {
	resource := make(map[string]interface{})
	err := s.client.Post(s.crmExportsPath, exportRequest, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
