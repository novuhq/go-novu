package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type IEvent interface {
	Trigger(ctx context.Context, eventId string, data ITriggerPayloadOptions) (EventResponse, error)
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

var _ IEvent = &EventService{}
