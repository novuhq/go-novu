package lib

import "io"

type ChannelType string
type GeneralError error

const Version = "v1"

const (
	HTTPStatusOk       = 200
	HTTPStatusCreated  = 201
	HTTPStatusNoChange = 204
	HTTPRedirectOk     = 300
)

const (
	EMAIL  ChannelType = "EMAIL"
	SMS                = "SMS"
	DIRECT             = "DIRECT"
)

type Data struct {
	Acknowledged bool   `json:"acknowledged"`
	Status       string `json:"status"`
}

type Response struct {
	Data Data `json:"data"`
}

type ITriggerPayloadOptions struct {
	To            interface{} `json:"to,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type TriggerRecipientsTypeArray interface {
	[]string | []SubscriberPayload
}
type TriggerRecipientsTypeSingle interface {
	string | SubscriberPayload
}

type TriggerTopicRecipientsTypeSingle struct {
	TopicKey string `json:"topicKey,omitempty"`
	Type     string `json:"type,omitempty"`
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
	TriggerRecipientsTypeSingle | TriggerRecipientsTypeArray | TriggerTopicRecipientsTypeSingle | []TriggerTopicRecipientsTypeSingle
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
	Name          string      `json:"name"`
	To            interface{} `json:"to"`
	Payload       interface{} `json:"payload"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type SubscriberResponse struct {
	Data interface{} `json:"data"`
}

type ListTopicsResponse struct {
	Page       int                `json:"name"`
	PageSize   int                `json:"pageSize"`
	TotalCount int                `json:"totalCount"`
	Data       []GetTopicResponse `json:"data"`
}

type GetTopicResponse struct {
	Id             string   `json:"_id"`
	OrganizationId string   `json:"_organizationId"`
	EnvironmentId  string   `json:"_environmentId"`
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Subscribers    []string `json:"subscribers"`
}

type ListTopicsOptions struct {
	Page     *int    `json:"page,omitempty"`
	PageSize *int    `json:"pageSize,omitempty"`
	Key      *string `json:"key,omitempty"`
}

type CreateTopicRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type RenameTopicRequest struct {
	Name string `json:"name"`
}

type SubscribersTopicRequest struct {
	Subscribers []string `json:"subscribers"`
}

type IntegrationCredentials struct {
	ApiKey           string                 `json:"apiKey,omitempty"`
	User             string                 `json:"user,omitempty"`
	SecretKey        string                 `json:"secretKey,omitempty"`
	Domain           string                 `json:"domain,omitempty"`
	Password         string                 `json:"password,omitempty"`
	Host             string                 `json:"host,omitempty"`
	Port             string                 `json:"port,omitempty"`
	Secure           bool                   `json:"secure,omitempty"`
	Region           string                 `json:"region,omitempty"`
	AccountSid       string                 `json:"accountSid,omitempty"`
	MessageProfileId string                 `json:"messageProfileId,omitempty"`
	Token            string                 `json:"token,omitempty"`
	From             string                 `json:"from,omitempty"`
	SenderName       string                 `json:"senderName,omitempty"`
	ProjectName      string                 `json:"projectName,omitempty"`
	ApplicationId    string                 `json:"applicationId,omitempty"`
	ClientId         string                 `json:"clientId,omitempty"`
	RequireTls       bool                   `json:"requireTls,omitempty"`
	IgnoreTls        bool                   `json:"ignoreTls,omitempty"`
	TlsOptions       map[string]interface{} `json:"tlsOptions,omitempty"`
}

type CreateIntegrationRequest struct {
	ProviderId  string                 `json:"providerId"`
	Channel     ChannelType            `json:"channel"`
	Credentials IntegrationCredentials `json:"credentials,omitempty"`
	Active      bool                   `json:"active"`
	Check       bool                   `json:"check"`
}

type UpdateIntegrationRequest struct {
	Credentials IntegrationCredentials `json:"credentials"`
	Active      bool                   `json:"active"`
	Check       bool                   `json:"check"`
}

type Integration struct {
	Id             string                 `json:"_id"`
	EnvironmentId  string                 `json:"_environmentId"`
	OrganizationId string                 `json:"_organizationId"`
	ProviderId     string                 `json:"providerId"`
	Channel        ChannelType            `json:"channel"`
	Credentials    IntegrationCredentials `json:"credentials"`
	Active         bool                   `json:"active"`
	Deleted        bool                   `json:"deleted"`
	UpdatedAt      string                 `json:"updatedAt"`
	DeletedAt      string                 `json:"deletedAt"`
	DeletedBy      string                 `json:"deletedBy"`
}

type IntegrationResponse struct {
	Data Integration `json:"data"`
}

type GetIntegrationsResponse struct {
	Data []Integration `json:"data"`
}
