package hubspot_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
)

func TestCrmListSearch(t *testing.T) {
	tests := []struct {
		name    string
		client  *hubspot.Client
		request *hubspot.CrmListSearchRequest
		want    *hubspot.CrmListSearchResponse
	}{
		{
			name: "search",
			client: hubspot.NewMockClient(&hubspot.MockConfig{
				Status: http.StatusOK,
				Header: http.Header{},
				Body: []byte(`
{
  "offset": 1,
  "hasMore": false,
  "lists": [
    {
      "listId": "8",
      "listVersion": 1,
      "createdAt": "2024-05-13T19:44:45.137Z",
      "updatedAt": "2024-05-13T19:44:45.137Z",
      "filtersUpdatedAt": "2024-05-13T19:44:45.137Z",
      "processingStatus": "COMPLETE",
      "createdById": "60342733",
      "updatedById": "60342733",
      "processingType": "DYNAMIC",
      "objectTypeId": "0-1",
      "name": "localdev active list",
      "additionalProperties": {
        "hs_list_reference_count": "0",
        "hs_list_size": "5",
        "hs_last_record_added_at": "1715629490454"
      }
    }
  ],
  "total": 1
}`,
				),
			}),
			request: &hubspot.CrmListSearchRequest{},
			want: &hubspot.CrmListSearchResponse{
				Offset:  hubspot.HsInt(1),
				Total:   hubspot.HsInt(1),
				HasMore: hubspot.HsBool(false),
				Lists: []*hubspot.CrmList{
					{
						ListID:           hubspot.HsStr("8"),
						ListVersion:      hubspot.HsInt(1),
						CreatedAt:        hubspot.HsTime(time.Date(2024, 5, 13, 19, 44, 45, 137000000, time.UTC)),
						UpdatedAt:        hubspot.HsTime(time.Date(2024, 5, 13, 19, 44, 45, 137000000, time.UTC)),
						FiltersUpdatedAt: hubspot.NewTime(time.Date(2024, 5, 13, 19, 44, 45, 137000000, time.UTC)),
						ProcessingStatus: hubspot.HsStr("COMPLETE"),
						CreatedByID:      hubspot.HsStr("60342733"),
						UpdatedByID:      hubspot.HsStr("60342733"),
						ProcessingType:   hubspot.CrmListProcessingTypeDynamic,
						ObjectTypeID:     hubspot.HsStr("0-1"),
						Name:             hubspot.HsStr("localdev active list"),
						AdditionalProperties: map[hubspot.HsStr]interface{}{
							"hs_list_reference_count": "0",
							"hs_list_size":            "5",
							"hs_last_record_added_at": "1715629490454",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.CRM.List.Search(tt.request)
			if err != nil {
				t.Fatalf("error = %v", err)
			} else if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("response mismatch (-want +got): %s", diff)
			}
		})
	}
}

func TestCrmListGet(t *testing.T) {
	tests := []struct {
		name    string
		client  *hubspot.Client
		listId  string
		request *hubspot.CrmListGetRequest
		want    *hubspot.CrmList
	}{
		{
			name: "get",
			client: hubspot.NewMockClient(&hubspot.MockConfig{
				Status: http.StatusCreated,
				Header: http.Header{},
				Body: []byte(`
{
  "list": {
    "listId": "8",
    "listVersion": 1,
    "createdAt": "2024-05-13T19:44:45.137Z",
    "updatedAt": "2024-05-13T19:44:45.137Z",
    "filtersUpdatedAt": "2024-05-13T19:44:45.137Z",
    "processingStatus": "COMPLETE",
    "createdById": "60342733",
    "updatedById": "60342733",
    "processingType": "DYNAMIC",
    "objectTypeId": "0-1",
    "name": "localdev active list",
    "size": 5
  }
}
`,
				),
			}),
			listId:  "8",
			request: &hubspot.CrmListGetRequest{},
			want: &hubspot.CrmList{
				ListID:               hubspot.HsStr("8"),
				ListVersion:          hubspot.HsInt(1),
				CreatedAt:            hubspot.HsTime(time.Date(2024, 5, 13, 19, 44, 45, 137000000, time.UTC)),
				UpdatedAt:            hubspot.HsTime(time.Date(2024, 5, 13, 19, 44, 45, 137000000, time.UTC)),
				FiltersUpdatedAt:     hubspot.NewTime(time.Date(2024, 5, 13, 19, 44, 45, 137000000, time.UTC)),
				ProcessingStatus:     hubspot.HsStr("COMPLETE"),
				CreatedByID:          hubspot.HsStr("60342733"),
				UpdatedByID:          hubspot.HsStr("60342733"),
				ProcessingType:       hubspot.CrmListProcessingTypeDynamic,
				ObjectTypeID:         hubspot.HsStr("0-1"),
				Name:                 hubspot.HsStr("localdev active list"),
				Size:                 hubspot.NewInt(5),
				AdditionalProperties: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.CRM.List.Get(tt.listId, tt.request)
			if err != nil {
				t.Fatalf("error = %v", err)
			} else if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("response mismatch (-want +got): %s", diff)
			}
		})
	}
}
