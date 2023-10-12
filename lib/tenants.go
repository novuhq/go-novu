package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type TenantService service

func (e *TenantService) CreateTenant(ctx context.Context, name string,identifier string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("tenants")
	n := map[string]string{"name": name,"identifier":identifier}
	jsonBody, _ := json.Marshal(n)
	b := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), b)
	if err != nil {
		return resp, err
	}
	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (e *TenantService) GetTenants(ctx context.Context,page string,limit string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("tenants")
	v := URL.Query();
	v.Set("page",page)
	v.Set("limit",limit)
	URL.RawQuery = v.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (e *TenantService) GetTenant(ctx context.Context,identifier string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("tenants",identifier)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (e *TenantService) DeleteTenant(ctx context.Context, identifier string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("tenants", identifier)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}


func (e *TenantService) UpdateTenant(ctx context.Context, identifier string,updateTenantObject *UpdateTenantRequest) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("tenants", identifier)
	jsonBody, _ := json.Marshal(updateTenantObject)
	b := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, URL.String(), b)
	if err != nil {
		return resp, err
	}
	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
