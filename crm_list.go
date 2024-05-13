package hubspot

import "fmt"

const (
	crmListBasePath = "lists"
)

type CrmListProcessingType string

const (
	CrmListProcessingTypeDynamic  CrmListProcessingType = "DYNAMIC"
	CrmListProcessingTypeManual   CrmListProcessingType = "MANUAL"
	CrmListProcessingTypeSnapshot CrmListProcessingType = "SNAPSHOT"
)

type CrmListService interface {
	Search(req *CrmListSearchRequest) (*CrmListSearchResponse, error)
	Get(listId string, req *CrmListGetRequest) (*CrmList, error)
}

type CrmListServiceOp struct {
	client      *Client
	crmListPath string
}

var _ CrmListService = (*CrmListServiceOp)(nil)

type CrmListSearchRequest struct {
	ListIds         []string                `json:"listIds,omitempty"`
	Offset          int                     `json:"offset,omitempty"`
	Query           string                  `json:"query,omitempty"`
	Count           int                     `json:"count,omitempty"`
	ProcessingTypes []CrmListProcessingType `json:"processingTypes,omitempty"`
}

type CrmListSearchResponse struct {
	Offset  HsInt      `json:"offset"`
	Total   HsInt      `json:"total"`
	HasMore HsBool     `json:"hasMore"`
	Lists   []*CrmList `json:"lists"`
}

type CrmListGetRequest struct {
	IncludeFilters bool `json:"includeFilters,omitempty"`
}

type CrmListGetResponse struct {
	List CrmList `json:"list"`
}

type CrmList struct {
	ListID               HsStr                 `json:"listId"`
	ListVersion          HsInt                 `json:"listVersion"`
	CreatedAt            HsTime                `json:"createdAt"`
	UpdatedAt            HsTime                `json:"updatedAt"`
	DeletedAt            *HsTime               `json:"deletedAt"`
	FiltersUpdatedAt     *HsTime               `json:"filtersUpdatedAt"`
	ProcessingStatus     HsStr                 `json:"processingStatus"`
	CreatedByID          HsStr                 `json:"createdById"`
	UpdatedByID          HsStr                 `json:"updatedById"`
	ProcessingType       CrmListProcessingType `json:"processingType"`
	ObjectTypeID         HsStr                 `json:"objectTypeId"`
	Name                 HsStr                 `json:"name"`
	Size                 *HsInt                `json:"size"`
	AdditionalProperties map[HsStr]any         `json:"additionalProperties"`
}

func (op *CrmListServiceOp) Search(req *CrmListSearchRequest) (*CrmListSearchResponse, error) {
	var response CrmListSearchResponse
	path := fmt.Sprintf("%s/%s", op.crmListPath, "search")
	err := op.client.Post(path, req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (op *CrmListServiceOp) Get(listId string, req *CrmListGetRequest) (*CrmList, error) {
	var response CrmListGetResponse
	path := fmt.Sprintf("%s/%s", op.crmListPath, listId)
	err := op.client.Get(path, &response, req)
	if err != nil {
		return nil, err
	}
	return &response.List, nil
}
