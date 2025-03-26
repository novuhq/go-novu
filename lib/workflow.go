package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type WorkflowService service

func (w *WorkflowService) List(ctx context.Context, options *WorkflowGetRequestOptions) (*WorkflowGetResponse, error) {
	var resp WorkflowGetResponse
	URL := w.client.config.BackendURL.JoinPath("workflows")
	if options == nil {
		options = &WorkflowGetRequestOptions{}
	}
	queryParams, _ := json.Marshal(options)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), bytes.NewBuffer(queryParams))
	if err != nil {
		return nil, err
	}

	_, err = w.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WorkflowService) Create(ctx context.Context, request WorkflowData) (*WorkflowData, error) {
	var resp WorkflowData
	URL := w.client.config.BackendURL.JoinPath("workflows")

	requestBody := WorkflowData{
		Name:                request.Name,
		NotificationGroupId: request.NotificationGroupId,
		Tags:                request.Tags,
		Description:         request.Description,
		Steps:               request.Steps,
		Active:              request.Active,
		Critical:            request.Critical,
		PreferenceSettings:  request.PreferenceSettings,
		BlueprintId:         request.BlueprintId,
		Data:                request.Data,
	}

	jsonBody, _ := json.Marshal(requestBody)
	b := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), b)
	if err != nil {
		return nil, err
	}
	_, err = w.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WorkflowService) Update(ctx context.Context, key string, request WorkflowData) (*WorkflowData, error) {

	var resp WorkflowData
	URL := w.client.config.BackendURL.JoinPath("workflows", key)

	requestBody := WorkflowData{
		Name:                request.Name,
		Tags:                request.Tags,
		Description:         request.Description,
		Steps:               request.Steps,
		NotificationGroupId: request.NotificationGroupId,
		Critical:            request.Critical,
		PreferenceSettings:  request.PreferenceSettings,
		Data:                request.Data,
	}

	jsonBody, _ := json.Marshal(requestBody)
	b := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, URL.String(), b)
	if err != nil {
		return nil, err
	}
	_, err = w.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WorkflowService) Delete(ctx context.Context, key string) (*WorkflowDeleteResponse, error) {
	var resp WorkflowDeleteResponse
	URL := w.client.config.BackendURL.JoinPath("workflows", key)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}
	_, err = w.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WorkflowService) Get(ctx context.Context, key string) (*WorkflowData, error) {
	var resp WorkflowData
	URL := w.client.config.BackendURL.JoinPath("workflows", key)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	_, err = w.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WorkflowService) UpdateStatus(ctx context.Context, key string) (*WorkflowData, error) {
	var resp WorkflowData
	URL := w.client.config.BackendURL.JoinPath("workflows", key, "status")

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, URL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}
	_, err = w.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil

}
