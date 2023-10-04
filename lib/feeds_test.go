package lib_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/novuhq/go-novu/lib"
)

var feedsApiResponse = `{
    "data": {
        "_id": "string",
        "name": "string",
        "identifier": "string",
        "_environmentId": "string",
        "_organizationId": "string"
    }
}
`

func TestCreateFeed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Want POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/feeds" {
			t.Errorf("Want /v1/feeds, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(feedsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.FeedsApi.CreateFeed(context.Background(), "FeedyMcFeederson")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestGetFeeds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Want GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/feeds" {
			t.Errorf("Want /v1/feeds, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(feedsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.FeedsApi.GetFeeds(context.Background())
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestDeleteFeed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Want DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/feeds/FeedId" {
			t.Errorf("Want /v1/feeds/FeedId, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(feedsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.FeedsApi.DeleteFeed(context.Background(), "FeedId")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}
