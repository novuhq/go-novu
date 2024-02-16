package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type ChangesService service

func (c *ChangesService) GetChangesCount(ctx context.Context) (ChangesCountResponse, error) {
	var resp ChangesCountResponse
	URL := c.client.config.BackendURL.JoinPath("changes", "count")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = c.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *ChangesService) GetChanges(ctx context.Context, q ChangesGetQuery) (ChangesGetResponse, error) {
	var resp ChangesGetResponse
	URL := c.client.config.BackendURL.JoinPath("changes")
	URL.RawQuery = q.BuildQuery()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = c.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *ChangesService) ApplyChange(ctx context.Context, changeId string) (ChangesApplyResponse, error) {
	var resp ChangesApplyResponse
	URL := c.client.config.BackendURL.JoinPath("changes", changeId, "apply")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = c.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *ChangesService) ApplyBulkChanges(ctx context.Context, payload ChangesBulkApplyPayload) (ChangesApplyResponse, error) {
	var resp ChangesApplyResponse
	URL := c.client.config.BackendURL.JoinPath("changes", "bulk", "apply")
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return resp, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	_, err = c.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *ChangesGetQuery) BuildQuery() string {
	params := url.Values{}

	if c.Page == 0 {
		c.Page = 1
	}

	if c.Limit == 0 {
		c.Limit = 10
	}

	if c.Promoted == "" {
		c.Promoted = "false"
	}
	params.Add("page", strconv.Itoa(c.Page))
	params.Add("limit", strconv.Itoa(c.Limit))
	params.Add("promoted", c.Promoted)
	return params.Encode()
}
