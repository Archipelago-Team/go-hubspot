package hubspot

type CrmExportConfig struct {
	ExportType             string                 `json:"exportType"`
	Format                 string                 `json:"format"`
	ExportName             string                 `json:"exportName"`
	Language               string                 `json:"language"`
	ObjectType             string                 `json:"objectType"`
	AssociatedObjectType   string                 `json:"associatedObjectType"`
	ObjectProperties       []string               `json:"objectProperties"`
	PublicCrmSearchRequest CrmExportSearchRequest `json:"publicCrmSearchRequest"`
}

type CrmExportFilter struct {
	PropertyName string `json:"propertyName"`
	Operator     string `json:"operator"`
	Value        string `json:"value"`
}
type CrmExportSort struct {
	PropertyName string `json:"propertyName"`
	Order        string `json:"order"`
}
type CrmExportSearchRequest struct {
	Filters []CrmExportFilter
	Query   string `json:"query"`
	Sorts   []CrmExportSort
}

func (s *CrmExportsServiceOp) Start(exportRequest *CrmExportConfig) (interface{}, error) {
	resource := make(map[string]interface{})
	err := s.client.Post(s.crmExportsPath, exportRequest, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
