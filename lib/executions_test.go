package lib_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/novuhq/go-novu/lib"
)

var executionsGetResponse = `{
	"data": [
	  {
		"_id": "string",
		"_organizationId": "string",
		"_jobId": "string",
		"_environmentId": "string",
		"_notificationId": "string",
		"_notificationTemplateId": "string",
		"_subscriberId": "string",
		"_messageId": "string",
		"providerId": "string",
		"transactionId": "string",
		"channel": "in_app",
		"detail": "string",
		"source": "Credentials",
		"status": "Success",
		"isTest": true,
		"isRetry": true,
		"createdAt": "string"
	  }
	]
  }
`

func TestGetExecutions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Want GET, got %s", r.Method)
		}
		expected := "/v1/execution-details?notificationId=12345&subscriberId=XYZ"
		if r.URL.String() != expected {
			t.Errorf("Want %s, got %s", expected, r.URL.String())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(executionsGetResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})

	q := lib.ExecutionsQueryParams{NotificationId: "12345", SubscriberId: "XYZ"}
	resp, err := c.ExecutionsApi.GetExecutions(context.Background(), q)
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}
