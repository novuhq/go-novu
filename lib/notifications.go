package lib

import (
	"context"
	"github.com/novuhq/go-novu/utils"
	"net/http"
	"net/url"
	"time"
)

type INotifications interface {
	GetNotifications(ctx context.Context, params ...utils.QueryParam) (NotificationResponse, error)
	GetNotification(ctx context.Context, notificationID string) (Notification, error)
	NotificationGraphStats(ctx context.Context, params ...utils.QueryParam) (NotificationGraphStats, error)
	NotificationStats(ctx context.Context) (NotificationStats, error)
}

type NotificationService service

type NotificationSubscriber struct{}
type NotificationTemplate struct{}

type Notification struct {
	ID             string                 `json:"_id"`
	EnvironmentID  string                 `json:"_environmentId"`
	OrganizationID string                 `json:"_organizationId"`
	TransactionID  string                 `json:"transactionId"`
	CreatedAt      time.Time              `json:"createdAt"`
	Channels       string                 `json:"channels"`
	Subscriber     NotificationSubscriber `json:"subscriber"`
	Template       NotificationTemplate   `json:"template"`
	Jobs           []string               `json:"jobs"`
}

type NotificationResponse struct {
	TotalCount int            `json:"totalCount"`
	Data       []Notification `json:"data"`
	PageSize   int            `json:"pageSize"`
	Page       int            `json:"page"`
}

type NotificationStats struct {
	WeeklySent  int `json:"weeklySent"`
	MonthlySent int `json:"monthlySent"`
	YearlySent  int `json:"yearlySent"`
}

type NotificationGraphStats struct {
	ID        string   `json:"_id"`
	Count     int      `json:"count"`
	Templates []string `json:"templates"`
	Channels  []string `json:"channels"`
}

func (n NotificationService) GetNotifications(ctx context.Context, params ...utils.QueryParam) (NotificationResponse, error) {
	var resp NotificationResponse
	URL := n.client.config.BackendURL.JoinPath("notifications")

	query := url.Values{}

	for _, opt := range params {
		query.Add(opt.Key, opt.Value)
	}

	URL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = n.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (n NotificationService) GetNotification(ctx context.Context, notificationID string) (Notification, error) {
	var resp Notification
	URL := n.client.config.BackendURL.JoinPath("notifications", notificationID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = n.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (n NotificationService) NotificationGraphStats(ctx context.Context, params ...utils.QueryParam) (NotificationGraphStats, error) {
	var resp NotificationGraphStats
	URL := n.client.config.BackendURL.JoinPath("notifications/graph/stats")

	query := url.Values{}

	for _, opt := range params {
		query.Add(opt.Key, opt.Value)
	}

	URL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = n.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (n NotificationService) NotificationStats(ctx context.Context) (NotificationStats, error) {
	var resp NotificationStats
	URL := n.client.config.BackendURL.JoinPath("notifications/stats")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return resp, err
	}

	_, err = n.client.sendRequest(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var _ INotifications = &NotificationService{}
