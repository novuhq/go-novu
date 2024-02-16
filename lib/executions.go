package lib

import (
	"context"
	"net/http"
	"net/url"
)

type ExecutionsService service

func (e *ExecutionsService) GetExecutions(ctx context.Context, q QueryBuilder) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("execution-details")
	URL.RawQuery = q.BuildQuery()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}
	_, err = e.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (q ExecutionsQueryParams) BuildQuery() string {
	params := url.Values{}
	if q.NotificationId != "" {
		params.Add("notificationId", q.NotificationId)
	}
	if q.SubscriberId != "" {
		params.Add("subscriberId", q.SubscriberId)
	}
	return params.Encode()
}
