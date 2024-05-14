package hubspot_test

import (
	"net/http"
	"testing"

	"github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
)

func TestAccountInfoDetails(t *testing.T) {
	tests := []struct {
		name   string
		client *hubspot.Client
		want   *hubspot.AccountInfo
	}{
		{
			name: "search",
			client: hubspot.NewMockClient(&hubspot.MockConfig{
				Status: http.StatusOK,
				Header: http.Header{},
				Body:   []byte(`{"portalId":143328477,"accountType":"DEVELOPER_TEST","timeZone":"US/Eastern","companyCurrency":"USD","additionalCurrencies":[],"utcOffset":"-04:00","utcOffsetMilliseconds":-14400000,"uiDomain":"app-eu1.hubspot.com","dataHostingLocation":"eu1"}`),
			}),
			want: &hubspot.AccountInfo{
				PortalId: 143328477,
				UiDomain: "app-eu1.hubspot.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.Account.Info.Details()
			if err != nil {
				t.Fatalf("error = %v", err)
			} else if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("response mismatch (-want +got): %s", diff)
			}
		})
	}
}
