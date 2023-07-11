package lib

import (
	"context"
	"net/http"
)

type IInboundParser interface {
	Get(ctx context.Context) (*InboundParserResponse, error)
}

type InboundParserService service

func (i *InboundParserService) Get(ctx context.Context) (*InboundParserResponse, error) {
	var resp InboundParserResponse

	URL := i.client.config.BackendURL.JoinPath("inbound-parse", "mx", "status")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	_, err = i.client.sendRequest(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
