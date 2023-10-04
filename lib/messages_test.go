package lib_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/novuhq/go-novu/lib"
)

var messagesGetResponse = `{
	"hasMore": true,
	"data": [
	  {
		"_id": "string",
		"_environmentId": "string",
		"_organizationId": "string",
		"transactionId": "string",
		"createdAt": "string",
		"channels": "in_app",
		"subscriber": {
		  "firstName": "string",
		  "_id": "string",
		  "lastName": "string",
		  "email": "string",
		  "phone": "string"
		},
		"template": {
		  "_id": "string",
		  "name": "string",
		  "triggers": [
			{
			  "type": "string",
			  "identifier": "string",
			  "variables": [
				{
				  "name": "string"
				}
			  ],
			  "subscriberVariables": [
				{
				  "name": "string"
				}
			  ]
			}
		  ]
		},
		"jobs": [
		  {
			"_id": "string",
			"type": "string",
			"digest": {},
			"executionDetails": [
			  {
				"_id": "string",
				"_jobId": "string",
				"status": "Success",
				"detail": "string",
				"isRetry": true,
				"isTest": true,
				"providerId": {},
				"raw": "string",
				"source": "Credentials"
			  }
			],
			"step": {
			  "_id": "string",
			  "active": true,
			  "filters": {
				"isNegated": true,
				"type": "BOOLEAN",
				"value": "AND",
				"children": [
				  {
					"field": "string",
					"value": "string",
					"operator": "LARGER",
					"on": "subscriber"
				  }
				]
			  },
			  "template": {}
			},
			"payload": {},
			"providerId": {},
			"status": "string"
		  }
		]
	  }
	],
	"pageSize": 0,
	"page": 0
  }
`

var messagesDeleteResponse = `{
	"data": {
	  "acknowledged": true,
	  "status": "deleted"
	}
  }`

func TestGetMessages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Want GET, got %s", r.Method)
		}
		expected := "/v1/messages?channel=email&transactionId=12&transactionId=156"
		if r.URL.String() != expected {
			t.Errorf("Want %s, got %s", expected, r.URL.String())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(messagesGetResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})

	q := lib.MessagesQueryParams{TransactionId: []string{"12", "156"}, Channel: "email"}
	resp, err := c.MessagesApi.GetMessages(context.Background(), q)
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestMessagesBuildQuery(t *testing.T) {
	tests := map[string]struct {
		q    lib.MessagesQueryParams
		want string
	}{
		"none": {
			q:    lib.MessagesQueryParams{},
			want: "",
		},
		"transactionIds only": {
			q:    lib.MessagesQueryParams{TransactionId: []string{"1", "2", "2938a0AAcx42"}},
			want: "transactionId=1&transactionId=2&transactionId=2938a0AAcx42",
		},
		"channel only": {
			q:    lib.MessagesQueryParams{Channel: "email"},
			want: "channel=email",
		},
		"all params": {
			q: lib.MessagesQueryParams{
				SubscriberId:  "subId",
				Channel:       "email",
				TransactionId: []string{"1", "2", "2938a0AAcx42"},
				Page:          2,
				Limit:         25,
			},
			want: "channel=email&limit=25&page=2&subscriberId=subId&transactionId=1&transactionId=2&transactionId=2938a0AAcx42",
		},
	}
	for name, tc := range tests {
		got := tc.q.BuildQuery()
		if got != tc.want {
			t.Errorf("%s-- want: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestDeleteMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Want DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/messages/MessageId" {
			t.Errorf("Want /v1/messages/MessageId, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(messagesDeleteResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.MessagesApi.DeleteMessage(context.Background(), "MessageId")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}
