package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type LayoutService service

func (l *LayoutService) Create(ctx context.Context, request CreateLayoutRequest) (*CreateLayoutResponse, error) {
	var resp CreateLayoutResponse
	URL := l.client.config.BackendURL.JoinPath("layouts")

	requestBody := CreateLayoutRequest{
		Name:        request.Name,
		Identifier:  request.Identifier,
		Description: request.Description,
		Content:     request.Content,
		Variables:   request.Variables,
		IsDefault:   request.IsDefault,
	}

	jsonBody, _ := json.Marshal(requestBody)
	b := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), b)
	if err != nil {
		return nil, err
	}
	_, err = l.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (l *LayoutService) List(ctx context.Context, options *LayoutRequestOptions) (*LayoutsResponse, error) {
	var resp LayoutsResponse
	URL := l.client.config.BackendURL.JoinPath("layouts")
	if options == nil {
		options = &LayoutRequestOptions{}
	}
	queryParams, _ := json.Marshal(options)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), bytes.NewBuffer(queryParams))
	if err != nil {
		return nil, err
	}

	_, err = l.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (l *LayoutService) Get(ctx context.Context, key string) (*LayoutResponse, error) {
	var resp LayoutResponse
	URL := l.client.config.BackendURL.JoinPath("layouts", key)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	_, err = l.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (l *LayoutService) Delete(ctx context.Context, key string) error {
	var resp interface{}
	URL := l.client.config.BackendURL.JoinPath("layouts", key)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), http.NoBody)
	if err != nil {
		return err
	}
	_, err = l.client.sendRequest(req, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (l *LayoutService) Update(ctx context.Context, key string, request CreateLayoutRequest) (*LayoutResponse, error) {
	var resp LayoutResponse
	URL := l.client.config.BackendURL.JoinPath("layouts", key)

	requestBody := CreateLayoutRequest{
		Name:        request.Name,
		Identifier:  request.Identifier,
		Description: request.Description,
		Content:     request.Content,
		Variables:   request.Variables,
		IsDefault:   request.IsDefault,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	_, err = l.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (l *LayoutService) SetDefault(ctx context.Context, key string) error {
	var resp interface{}
	URL := l.client.config.BackendURL.JoinPath("layouts", key, "default")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), http.NoBody)
	if err != nil {
		return err
	}

	_, err = l.client.sendRequest(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
