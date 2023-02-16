package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Config struct {
	BackendURL string
	HttpClient *http.Client
}

type APIClient struct {
	apiKey string
	config *Config
	common service

	// Api Service
	SubscriberApi *SubscriberService
	EventApi      *EventService
	TopicsApi     *TopicService
}

type service struct {
	client *APIClient
}

func NewAPIClient(apiKey string, cfg *Config) *APIClient {
	cfg.BackendURL = buildBackendURL(cfg)

	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{Timeout: 20 * time.Second}
	}

	c := &APIClient{apiKey: apiKey}
	c.config = cfg
	c.common.client = c

	// API Services
	c.EventApi = (*EventService)(&c.common)
	c.SubscriberApi = (*SubscriberService)(&c.common)
	c.TopicsApi = (*TopicService)(&c.common)

	return c
}

func (c APIClient) sendRequest(req *http.Request, resp interface{}) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", c.apiKey))

	res, err := c.config.HttpClient.Do(req)
	if err != nil {
		return res, errors.Wrap(err, "failed to execute request")
	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode >= http.StatusMultipleChoices {
		return res, errors.Errorf(
			`request was not successful, status code %d, %s`, res.StatusCode,
			string(body),
		)
	}

	if string(body) == "" {
		resp = map[string]string{}
		return res, nil
	}

	err = c.decode(&resp, body)
	if err != nil {
		return res, errors.Wrap(err, "unable to unmarshal response body")
	}

	return res, nil
}

func (c APIClient) mergeStruct(target, patch interface{}) (interface{}, error) {
	var m map[string]interface{}

	targetPayload, _ := json.Marshal(target)
	patchPayload, _ := json.Marshal(patch)

	_ = json.Unmarshal(targetPayload, &m)
	_ = json.Unmarshal(patchPayload, &m)

	return m, nil
}

func (c APIClient) decode(v interface{}, b []byte) (err error) {
	if err = json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}

func buildBackendURL(cfg *Config) string {
	novuVersion := "v1"

	if cfg.BackendURL == "" {
		return fmt.Sprintf("https://api.novu.co/%s", novuVersion)
	}

	if strings.Contains(cfg.BackendURL, "novu.co/v") {
		return cfg.BackendURL
	}
	return fmt.Sprintf(cfg.BackendURL+"/%s", novuVersion)
}
