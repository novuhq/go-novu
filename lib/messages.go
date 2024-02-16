package lib

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

type MessagesService service

func (e *MessagesService) GetMessages(ctx context.Context, q QueryBuilder) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("messages")
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

func (e *MessagesService) DeleteMessage(ctx context.Context, messageId string) (JsonResponse, error) {
	var resp JsonResponse
	URL := e.client.config.BackendURL.JoinPath("messages", messageId)
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

func (q MessagesQueryParams) BuildQuery() string {
	params := url.Values{}
	if q.Channel != "" {
		params.Add("channel", q.Channel)
	}
	if q.SubscriberId != "" {
		params.Add("subscriberId", q.SubscriberId)
	}
	if q.TransactionId != nil && len(q.TransactionId) > 0 {
		for _, s := range q.TransactionId {
			params.Add("transactionId", s)
		}
	}
	if q.Page != 0 {
		i := strconv.Itoa(q.Page)
		params.Add("page", i)
	}
	if q.Limit != 0 {
		i := strconv.Itoa(q.Limit)
		params.Add("limit", i)
	}
	return params.Encode()
}
