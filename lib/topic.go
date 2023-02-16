package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ITopic interface {
	Create(ctx context.Context, key string, name string) error
	List(ctx context.Context, options *ListTopicsOptions) (*ListTopicsResponse, error)
	AddSubscribers(ctx context.Context, key string, subscribers []string) error
	RemoveSubscribers(ctx context.Context, key string, subscribers []string) error
	Get(ctx context.Context, key string) (*GetTopicResponse, error)
	Rename(ctx context.Context, key string, name string) (*GetTopicResponse, error)
}

type TopicService service

func (t *TopicService) Create(ctx context.Context, key string, name string) error {
	var resp interface{}
	URL := fmt.Sprintf(t.client.config.BackendURL+"/%s", "topics")

	reqBody := CreateTopicRequest{
		Name: name,
		Key:  key,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	httpResponse, err := t.client.sendRequest(req, &resp)
	if err != nil {
		return err
	}

	if httpResponse.StatusCode != HTTPStatusCreated {
		return errors.Wrap(err, "unable to create topic")
	}

	return nil
}

func (t *TopicService) List(ctx context.Context, options *ListTopicsOptions) (*ListTopicsResponse, error) {
	var resp ListTopicsResponse
	URL := fmt.Sprintf(t.client.config.BackendURL+"/%s", "topics")

	if options == nil {
		options = &ListTopicsOptions{}
	}
	queryParams, _ := json.Marshal(options)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, bytes.NewBuffer(queryParams))
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (t *TopicService) AddSubscribers(ctx context.Context, key string, subscribers []string) error {
	URL := fmt.Sprintf(t.client.config.BackendURL+"/%s/%s/subscribers", "topics", key)

	queryParams, _ := json.Marshal(SubscribersTopicRequest{
		Subscribers: subscribers,
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewBuffer(queryParams))
	if err != nil {
		return err
	}

	_, err = t.client.sendRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}

func (t *TopicService) RemoveSubscribers(ctx context.Context, key string, subscribers []string) error {
	URL := fmt.Sprintf(t.client.config.BackendURL+"/%s/%s/subscribers/removal", "topics", key)

	queryParams, _ := json.Marshal(SubscribersTopicRequest{
		Subscribers: subscribers,
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewBuffer(queryParams))
	if err != nil {
		return err
	}

	_, err = t.client.sendRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (t *TopicService) Get(ctx context.Context, key string) (*GetTopicResponse, error) {
	var resp GetTopicResponse
	URL := fmt.Sprintf(t.client.config.BackendURL+"/%s/%s", "topics", key)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (t *TopicService) Rename(ctx context.Context, key string, name string) (*GetTopicResponse, error) {
	var resp GetTopicResponse
	URL := fmt.Sprintf(t.client.config.BackendURL+"/%s/%s", "topics", key)

	reqBody := RenameTopicRequest{
		Name: name,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	_, err = t.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
