package lib

import "io"

type ChannelType string
type GeneralError error

const Version = "v1"

const (
	HTTPStatusOk      = 200
	HTTPStatusCreated = 201
	HTTPRedirectOk    = 300
)

const (
	EMAIL  ChannelType = "EMAIL"
	SMS    ChannelType = "SMS"
	DIRECT ChannelType = "DIRECT"
)

type Data struct {
	Acknowledged bool   `json:"acknowledged"`
	Status       string `json:"status"`
}

type Response struct {
	Data Data `json:"data"`
}

type ITriggerPayloadOptions struct {
	To        interface{} `json:"to,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
	Overrides interface{} `json:"overrides,omitempty"`
}

type TriggerRecipientsTypeArray interface {
	[]string | []SubscriberPayload
}
type TriggerRecipientsTypeSingle interface {
	string | SubscriberPayload
}

type SubscriberPayload struct {
	FirstName string                 `json:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Phone     string                 `json:"phone,omitempty"`
	Avatar    string                 `json:"avatar,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type TriggerRecipientsType interface {
	TriggerRecipientsTypeSingle | TriggerRecipientsTypeArray
}

type ITriggerPayload interface {
	string | []string | bool | int64 | IAttachmentOptions | []IAttachmentOptions
}

type IAttachmentOptions struct {
	Mime     string        `json:"mime,omitempty"`
	File     io.Reader     `json:"file,omitempty"`
	Name     string        `json:"name,omitempty"`
	Channels []ChannelType `json:"channels,omitempty"`
}

type EventResponse struct {
	Data interface{} `json:"data"`
}

type EventRequest struct {
	Name      string      `json:"name"`
	To        interface{} `json:"to"`
	Payload   interface{} `json:"payload"`
	Overrides interface{} `json:"overrides,omitempty"`
}

type Subscriber struct {
	ID             string                  `json:"_id,omitempty"`
	FirstName      string                  `json:"firstName,omitempty"`
	LastName       string                  `json:"lastName,omitempty"`
	Email          string                  `json:"email,omitempty"`
	Phone          string                  `json:"phone,omitempty"`
	Avatar         string                  `json:"avatar,omitempty"`
	Locale         string                  `json:"locale,omitempty"`
	SubscriberID   string                  `json:"subscriberId,omitempty"`
	Channels       []SubscriberCredentials `json:"channels,omitempty"`
	IsOnline       bool                    `json:"isOnline,omitempty"`
	LastOnlineAt   string                  `json:"lastOnlineAt,omitempty"`
	OrganizationID string                  `json:"_organizationId,omitempty"`
	EnvironmentID  string                  `json:"_environmentId,omitempty"`
	Deleted        bool                    `json:"deleted,omitempty"`
	CreatedAt      string                  `json:"createdAt,omitempty"`
	UpdatedAt      string                  `json:"updatedAt,omitempty"`
}

type SubscriberIdentify struct {
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Locale       string `json:"locale,omitempty"`
	SubscriberID string `json:"subscriberId,omitempty"`
}

type SubscriberResponse struct {
	Data Subscriber
}

type Credentials struct {
	WebhookURL   string   `json:"webhookUrl,omitempty"`
	DeviceTokens []string `json:"deviceTokens,omitempty"`
}

type SubscriberCredentials struct {
	ProviderID  string      `json:"providerId,omitempty"`
	Credentials Credentials `json:"credentials,omitempty"`
}

type NotificationsOptions struct {
	Page           int    `json:"page,omitempty"`
	FeedIdentifier string `json:"feedIdentifier,omitempty"`
	Seen           *bool  `json:"seen,omitempty"`
}

type NotificationFeedResponse struct {
	TotalCount int                      `json:"totalCount,omitempty"`
	Data       []map[string]interface{} `json:"data,omitempty"`
	PageSize   int                      `json:"pageSize,omitempty"`
	Page       int                      `json:"page,omitempty"`
}

type Mark struct {
	Seen bool `json:"seen,omitempty"`
	Read bool `json:"read,omitempty"`
}

type MarkRequest struct {
	MessageID string `json:"messageId,omitempty"`
	Mark      Mark   `json:"mark"`
}

type MarkResponse struct {
	Data []NotificationResponse `json:"data"`
}

type NotificationResponse map[string]interface{}

type UnseenNotificationsCount struct {
	Count int `json:"count"`
}

type UnseenNotificationsCountResponse struct {
	Data UnseenNotificationsCount `json:"data"`
}

const (
	ChatProviderSlack   string = "slack"
	ChatProviderDiscord string = "discord"
	ChatProviderMSTeams string = "msteams"
	PushProviderFCM     string = "fcm"
	PushProviderAPNS    string = "apns"
	PushProviderEXPO    string = "expo"
)
