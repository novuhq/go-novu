package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type FeedsService service

func (e *FeedsService) CreateFeed(ctx context.Context, name string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("feeds")
	n := map[string]string{"name": name}
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

func (e *FeedsService) GetFeeds(ctx context.Context) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("feeds")
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

func (e *FeedsService) DeleteFeed(ctx context.Context, feedId string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("feeds", feedId)
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
