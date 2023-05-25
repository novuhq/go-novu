package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type IEvent interface {
	Trigger(ctx context.Context, eventId string, data ITriggerPayloadOptions) (EventResponse, error)
	TriggerBulk(ctx context.Context, data []BulkTriggerOptions) ([]EventResponse, error)
	BroadcastToAll(ctx context.Context, data BroadcastEventToAll) (EventResponse, error)
	CancelTrigger(ctx context.Context, transactionId string) (bool, error)
}

type EventService service

func (e *EventService) Trigger(ctx context.Context, eventId string, data ITriggerPayloadOptions) (EventResponse, error) {
	var resp EventResponse
	URL := e.client.config.BackendURL.JoinPath("events/trigger")

	reqBody := EventRequest{
		Name:          eventId,
		To:            data.To,
		Payload:       data.Payload,
		Overrides:     data.Overrides,
		TransactionId: data.TransactionId,
		Actor:         data.Actor,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (e *EventService) TriggerBulk(ctx context.Context, data []BulkTriggerOptions) ([]EventResponse, error) {
	var resp []EventResponse
	URL := e.client.config.BackendURL.JoinPath("events/trigger/bulk")

	reqBody := BulkTriggerEvent{
		Events: data,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (e *EventService) BroadcastToAll(ctx context.Context, data BroadcastEventToAll) (EventResponse, error) {
	var resp EventResponse
	URL := e.client.config.BackendURL.JoinPath("events/trigger/broadcast")

	reqBody := BroadcastEventToAll{
		Name:          data.Name,
		Payload:       data.Payload,
		Overrides:     data.Overrides,
		TransactionId: data.TransactionId,
		Actor:         data.Actor,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (e *EventService) CancelTrigger(ctx context.Context, transactionId string) (bool, error) {
	var resp bool
	URL := e.client.config.BackendURL.JoinPath("events/trigger/" + transactionId)

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

var _ IEvent = &EventService{}
