package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type ISubscribers interface {
	Identify(ctx context.Context, subscriberID string, data SubscriberIdentify) (Subscriber, error)
	Subscriber(ctx context.Context, subscriberID string) (Subscriber, error)
	Update(ctx context.Context, subscriberID string, data SubscriberIdentify) (Subscriber, error)
	UpdateCredentials(ctx context.Context, subscriberID string, data SubscriberCredentials) (Subscriber, error)
	Delete(ctx context.Context, subscriberID string) (SubscriberResponse, error)
	Notifications(ctx context.Context, subscriberID string, opts *NotificationsOptions) (NotificationFeedResponse, error)
	UnseenNotificationsCount(ctx context.Context, subscribierID string, seen bool) (UnseenNotificationsCountResponse, error)
	MessagesMarkAs(ctx context.Context, subscriberID string, data MarkRequest) (MarkResponse, error)
}

type SubscriberService service

func (s *SubscriberService) Identify(ctx context.Context, subscriberID string, data SubscriberIdentify) (Subscriber, error) {
	var resp SubscriberResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/%s", "subscribers")

	reqBody, err := s.client.mergeStruct(data, map[string]interface{}{"subscriberId": subscriberID})
	if err != nil {
		return resp.Data, errors.Wrap(err, "unable to merge struct")
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp.Data, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp.Data, err
	}

	return resp.Data, nil
}

func (s *SubscriberService) Subscriber(ctx context.Context, subscriberID string) (Subscriber, error) {
	var resp SubscriberResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s", subscriberID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, http.NoBody)
	if err != nil {
		return resp.Data, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp.Data, err
	}

	return resp.Data, nil
}

func (s *SubscriberService) Update(ctx context.Context, subscriberID string, data SubscriberIdentify) (Subscriber, error) {
	var resp SubscriberResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s", subscriberID)

	jsonBody, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp.Data, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp.Data, err
	}

	return resp.Data, nil
}

func (s *SubscriberService) UpdateCredentials(ctx context.Context, subscriberID string, data SubscriberCredentials) (Subscriber, error) {
	var resp SubscriberResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s/credentials", subscriberID)

	jsonBody, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp.Data, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp.Data, err
	}

	return resp.Data, nil
}

func (s *SubscriberService) Delete(ctx context.Context, subscriberID string) (SubscriberResponse, error) {
	var resp SubscriberResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s", subscriberID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL, http.NoBody)
	if err != nil {
		return resp, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *SubscriberService) Notifications(ctx context.Context, subscriberID string, opts *NotificationsOptions) (NotificationFeedResponse, error) {
	var resp NotificationFeedResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s/notifications/feed", subscriberID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, http.NoBody)
	if err != nil {
		return resp, err
	}

	if opts != nil {
		q := req.URL.Query()

		q.Add("page", strconv.Itoa(opts.Page))

		if opts.Seen != nil {
			q.Add("seen", fmt.Sprintf("%t", *opts.Seen))
		}

		if opts.FeedIdentifier != "" {
			q.Add("feedIdentifier", opts.FeedIdentifier)
		}

		req.URL.RawQuery = q.Encode()
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *SubscriberService) UnseenNotificationsCount(ctx context.Context, subscriberID string, seen bool) (UnseenNotificationsCountResponse, error) {
	var resp UnseenNotificationsCountResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s/notifications/unseen?seen=%b", subscriberID, seen)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, http.NoBody)
	if err != nil {
		return resp, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *SubscriberService) MessagesMarkAs(ctx context.Context, subscriberID string, data MarkRequest) (MarkResponse, error) {
	var resp MarkResponse
	URL := fmt.Sprintf(s.client.config.BackendURL+"/subscribers/%s/messages/markAs", subscriberID)

	jsonBody, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return resp, err
	}

	err = s.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var _ ISubscribers = &SubscriberService{}
