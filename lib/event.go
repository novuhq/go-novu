package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type IEvent interface {
	Trigger(ctx context.Context, eventId string, data ITriggerPayloadOptions) (EventResponse, error)
}

type EventService service

func (e *EventService) Trigger(ctx context.Context, eventId string, data ITriggerPayloadOptions) (EventResponse, error) {
	var resp EventResponse
	URL := fmt.Sprintf(e.client.config.BackendURL+"/%s", "events/trigger")

	reqBody := EventRequest{
		Name:      eventId,
		To:        data.To,
		Payload:   data.Payload,
		Overrides: data.Overrides,
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var _ IEvent = &EventService{}
