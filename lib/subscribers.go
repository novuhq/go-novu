package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type ISubscribers interface {
	Identify(ctx context.Context, subscriberID string, data interface{}) (SubscriberResponse, error)
	Update(ctx context.Context, subscriberID string, data interface{}) (SubscriberResponse, error)
	Delete(ctx context.Context, subscriberID string) (SubscriberResponse, error)
}

type SubscriberService service

func (s *SubscriberService) Identify(ctx context.Context, subscriberID string, data interface{}) (SubscriberResponse, error) {
	var resp SubscriberResponse
	URL := s.client.config.BackendURL.JoinPath("subscribers")

	reqBody, err := s.client.mergeStruct(data, map[string]interface{}{"subscriberId": subscriberID})
	if err != nil {
		return resp, errors.Wrap(err, "unable to merge struct")
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	_, err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *SubscriberService) Update(ctx context.Context, subscriberID string, data interface{}) (SubscriberResponse, error) {
	var resp SubscriberResponse
	URL := s.client.config.BackendURL.JoinPath("subscribers", subscriberID)

	jsonBody, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	_, err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *SubscriberService) Delete(ctx context.Context, subscriberID string) (SubscriberResponse, error) {
	var resp SubscriberResponse
	URL := s.client.config.BackendURL.JoinPath("subscribers", subscriberID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var _ ISubscribers = &SubscriberService{}
