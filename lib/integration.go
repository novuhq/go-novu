package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type IIntegration interface {
	Create(ctx context.Context, request CreateIntegrationRequest) (*IntegrationResponse, error)
	GetAll(ctx context.Context) (*GetIntegrationsResponse, error)
	GetActive(ctx context.Context) (*GetIntegrationsResponse, error)
	Update(ctx context.Context, integrationId string, request UpdateIntegrationRequest) (*IntegrationResponse, error)
	Delete(ctx context.Context, integrationId string) (*IntegrationResponse, error)
}

type IntegrationService service

func (integration *IntegrationService) Create(ctx context.Context, request CreateIntegrationRequest) (*IntegrationResponse, error) {
	var response IntegrationResponse
	Url := integration.client.config.BackendURL.JoinPath("integrations")

	requestBody := CreateIntegrationRequest{
		ProviderId:  request.ProviderId,
		Channel:     request.Channel,
		Credentials: request.Credentials,
		Active:      request.Active,
		Check:       request.Check,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, Url.String(), bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	httpResponse, err := integration.client.sendRequest(req, &response)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != HTTPStatusCreated {
		return nil, errors.Wrap(err, "Unable to create integration")
	}

	return &response, nil
}

func (integration *IntegrationService) GetAll(ctx context.Context) (*GetIntegrationsResponse, error) {
	var response GetIntegrationsResponse
	Url := integration.client.config.BackendURL.JoinPath("integrations")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, Url.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = integration.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (integration *IntegrationService) GetActive(ctx context.Context) (*GetIntegrationsResponse, error) {
	var response GetIntegrationsResponse
	Url := integration.client.config.BackendURL.JoinPath("integrations", "active")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, Url.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = integration.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (integration *IntegrationService) Update(ctx context.Context, integrationId string, request UpdateIntegrationRequest) (*IntegrationResponse, error) {
	var response IntegrationResponse
	Url := integration.client.config.BackendURL.JoinPath("integrations", integrationId)

	requestBody := UpdateIntegrationRequest{
		Credentials: request.Credentials,
		Active:      request.Active,
		Check:       request.Check,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, Url.String(), bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	_, err = integration.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (integration *IntegrationService) Delete(ctx context.Context, integrationId string) (*IntegrationResponse, error) {
	var response IntegrationResponse
	Url := integration.client.config.BackendURL.JoinPath("integrations", integrationId)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, Url.String(), http.NoBody)

	if err != nil {
		return nil, err
	}

	_, err = integration.client.sendRequest(req, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
