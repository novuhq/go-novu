package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type IIntegration interface {
	Create(ctx context.Context, request CreateIntegrationRequest) (*IntegrationResponse, error)
	GetAll(ctx context.Context) (*GetIntegrationsResponse, error)
	GetActive(ctx context.Context) (*GetIntegrationsResponse, error)
	GetWebhookSupportStatus(ctx context.Context, providerId string) (bool, error)
	Update(ctx context.Context, integrationId string, request UpdateIntegrationRequest) (*IntegrationResponse, error)
	Delete(ctx context.Context, integrationId string) (*IntegrationResponse, error)
	SetPrimary(ctx context.Context, integrationId string) (*IntegrationResponse, error)
}

type IntegrationService service

func (i IntegrationService) Create(ctx context.Context, request CreateIntegrationRequest) (*IntegrationResponse, error) {
	var response IntegrationResponse
	URL := i.client.config.BackendURL.JoinPath("integrations")

	requestBody := CreateIntegrationRequest{
		ProviderID:  request.ProviderID,
		Channel:     request.Channel,
		Credentials: request.Credentials,
		Active:      request.Active,
		Check:       request.Check,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (i IntegrationService) GetAll(ctx context.Context) (*GetIntegrationsResponse, error) {
	var response GetIntegrationsResponse
	URL := i.client.config.BackendURL.JoinPath("integrations")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (i IntegrationService) GetActive(ctx context.Context) (*GetIntegrationsResponse, error) {
	var response GetIntegrationsResponse
	URL := i.client.config.BackendURL.JoinPath("integrations", "active")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (i IntegrationService) GetWebhookSupportStatus(ctx context.Context, providerId string) (bool, error) {
	URL := i.client.config.BackendURL.JoinPath("integrations", "webhook", "provider", providerId, "status")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)

	if err != nil {
		return false, err
	}

	var status bool
	_, err = i.client.sendRequest(req, &status)

	if err != nil {
		return false, err
	}

	return status, nil
}

func (i IntegrationService) Update(ctx context.Context, integrationId string, request UpdateIntegrationRequest) (*IntegrationResponse, error) {
	var response IntegrationResponse
	URL := i.client.config.BackendURL.JoinPath("integrations", integrationId)

	requestBody := UpdateIntegrationRequest{
		Credentials: request.Credentials,
		Active:      request.Active,
		Check:       request.Check,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, URL.String(), bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (i IntegrationService) Delete(ctx context.Context, integrationId string) (*IntegrationResponse, error) {
	var response IntegrationResponse
	URL := i.client.config.BackendURL.JoinPath("integrations", integrationId)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (i IntegrationService) SetPrimary(ctx context.Context, integrationId string) (*IntegrationResponse, error) {
	var response IntegrationResponse
	URL := i.client.config.BackendURL.JoinPath("integrations", integrationId, "set-primary")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
