package checker

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CheckURL(t *testing.T) {
	upStatus := "UP"
	downStatus := "DOWN"

	tests := []struct {
		name           string
		responseStatus int
		wantStatus     string
		wantStatusCode int
	}{
		{
			name:           "status should be 200",
			responseStatus: http.StatusOK,
			wantStatus:     upStatus,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "status should be 404",
			responseStatus: http.StatusNotFound,
			wantStatus:     downStatus,
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "status should be 500",
			responseStatus: http.StatusInternalServerError,
			wantStatus:     downStatus,
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.responseStatus)
			}))
			defer server.Close()

			urlChecker := NewChecker()

			got := urlChecker.CheckURL(context.Background(), server.URL)

			if got.Status != tt.wantStatus {
				t.Errorf("Status = %s; want %s", got.Status, tt.wantStatus)
			}

			if got.StatusCode != tt.wantStatusCode {
				t.Errorf("StatusCode = %d; want %d", got.StatusCode, tt.wantStatusCode)
			}

			if got.URL != server.URL {
				t.Errorf("URL = %s; want %s", got.URL, server.URL)
			}

			if got.CheckedAt == "" {
				t.Error("CheckedAt should not be empty")
			}
		})
	}
}
