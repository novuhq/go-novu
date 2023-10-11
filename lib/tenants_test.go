package lib_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/novuhq/go-novu/lib"
)

var tenantsApiResponse = `{
	"data": {
	  "_environmentId": "string",
	  "_id": "string",
	  "createdAt": "string",
	  "data": "object",
	  "identifier": "string",
	  "name": "string",
	  "updatedAt": "string"
	}
  }
  
`

func TestCreateTenant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Want POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/tenants" {
			t.Errorf("Want /v1/tenants, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(tenantsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.TenantApi.CreateTenant(context.Background(), "Tenant", "TenantId")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestGetTenants(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Want GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/tenants" {
			t.Errorf("Want /v1/tenants, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(tenantsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.TenantApi.GetTenants(context.Background(), "1", "10")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestGetTenant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Want GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/tenants/TenantId" {
			t.Errorf("Want /v1/feeds, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(tenantsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.TenantApi.GetTenant(context.Background(), "TenantId")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestDeleteTenant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Want DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/tenants/TenantId" {
			t.Errorf("Want /v1/tenants/TenantId, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(tenantsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.TenantApi.DeleteTenant(context.Background(), "TenantId")
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}

func TestUpdateTenant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Want PATCH, got %s", r.Method)
		}
		if r.URL.Path != "/v1/tenants/TenantId" {
			t.Errorf("Want /v1/tenants/TenantId, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(tenantsApiResponse))
	}))
	defer server.Close()
	c := lib.NewAPIClient(novuApiKey, &lib.Config{BackendURL: lib.MustParseURL(server.URL)})
	resp, err := c.TenantApi.UpdateTenant(context.Background(), "TenantId", &lib.UpdateTenantRequest{
		Name: "Tenant2",
	})
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if resp.Data == nil || resp.Data == "" {
		t.Error("Expected response, got none")
	}
}
