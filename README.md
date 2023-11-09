# Novu's API v1 Go Library

Novu's API exposes the entire Novu features via a standardized programmatic interface. Please refer to the full [documentation](https://docs.novu.co/docs/overview/introduction) to learn more.

## Installation & Usage

Install the package to your GoLang project.

```golang
go get github.com/novuhq/go-novu
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```golang
package main

import (
	"context"
	"fmt"
	novu "github.com/novuhq/go-novu/lib"
	"log"
)

func main() {
	subscriberID := "<<REPLACE_WITH_YOUR_SUBSCRIBER>"
	apiKey := "<REPLACE_WITH_YOUR_API_KEY>"
	eventId := "<REPLACE_WITH_YOUR_EVENT_ID>"

	ctx := context.Background()
	to := map[string]interface{}{
		"lastName":     "Doe",
		"firstName":    "John",
		"subscriberId": subscriberID,
		"email":        "john@doemail.com",
	}

	payload := map[string]interface{}{
		"name": "Hello World",
		"organization": map[string]interface{}{
			"logo": "https://happycorp.com/logo.png",
		},
	}

	data := novu.ITriggerPayloadOptions{To: to, Payload: payload}
	novuClient := novu.NewAPIClient(apiKey, &novu.Config{})

	resp, err := novuClient.EventApi.Trigger(ctx, eventId, data)
	if err != nil {
		log.Fatal("novu error", err.Error())
		return
	}

	fmt.Println(resp)

	// get integrations
	integrations, err := novuClient.IntegrationsApi.GetAll(ctx)
	if err != nil {
		log.Fatal("Get all integrations error: ", err.Error())
	}
	fmt.Println(integrations)
}
```

**NOTE**
Check the `cmd` directory to see a sample implementation and test files to see sample tests

## Documentation for API Endpoints

| Class             | Method                                                                                     | HTTP request                                                 | Description                                            |
| ----------------- | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------ |
| _EventApi_        | [**Trigger**](https://docs.novu.co/platform/subscribers#removing-a-subscriber)             | **Post** /events/trigger                                     | Trigger                                                |
| _EventApi_        | [**TriggerBulk**](https://docs.novu.co/api/trigger-event/)                                 | **Post** /v1/events/trigger/bulk                             | Bulk trigger event                                     |
| _EventApi_        | [**BroadcastToAll**](https://docs.novu.co/api/broadcast-event-to-all/)                     | **Post** /v1/events/trigger/broadcast                        | Broadcast event to all                                 |
| _EventApi_        | [**CancelTrigger**](https://docs.novu.co/api/cancel-triggered-event/)                      | **Delete** /v1/events/trigger/:transactionId                 | Cancel triggered event                                 |
| _SubscriberApi_   | [**Get**](https://docs.novu.co/api/get-subscriber/)                                        | **Get** /subscribers/:subscriberId                           | Get a subscriber                                       |
| _SubscriberApi_   | [**Identify**](https://docs.novu.co/platform/subscribers#creating-a-subscriber)            | **Post** /subscribers                                        | Create a subscriber                                    |
| _SubscriberApi_   | [**Update**](https://docs.novu.co/platform/subscribers#updating-subscriber-data)           | **Put** /subscribers/:subscriberID                           | Update subscriber data                                 |
| _SubscriberApi_   | [**Delete**](https://docs.novu.co/platform/subscribers#removing-a-subscriber)              | **Delete** /subscribers/:subscriberID                        | Removing a subscriber                                  |
| _SubscriberApi_   | [**Get**](https://docs.novu.co/api/get-a-notification-feed-for-a-particular-subscriber)    | **Get** /subscribers/:subscriberId/notifications/feed        | Get a notification feed for a particular subscriber    |
| _SubscriberApi_   | [**Get**](https://docs.novu.co/api/get-the-unseen-notification-count-for-subscribers-feed) | **Get** /subscribers/:subscriberId/notifications/feed        | Get the unseen notification count for subscribers feed |
| _SubscriberApi_   | [**Post**](https://docs.novu.co/api/mark-a-subscriber-feed-message-as-seen)                | **Post** /v1/subscribers/:subscriberId/messages/markAs       | Mark a subscriber feed message as seen                 |
| _SubscriberApi_   | [**Get**](https://docs.novu.co/api/get-subscriber-preferences/)                            | **Get** /subscribers/:subscriberId/preferences               | Get subscriber preferences                             |
| _SubscriberApi_   | [**Patch**](https://docs.novu.co/api/update-subscriber-preference/)                        | **Patch** /subscribers/:subscriberId/preferences/:templateId | Update subscriber preference                           |
| _IntegrationsApi_ | [**Create**](https://docs.novu.co/platform/integrations)                                   | **Post** /integrations                                       | Create an integration                                  |
| _IntegrationsApi_ | [**Update**](https://docs.novu.co/platform/integrations)                                   | **Put** /integrations/:integrationId                         | Update an integration                                  |
| _IntegrationsApi_ | [**Delete**](https://docs.novu.co/platform/integrations)                                   | **Delete** /integrations/:integrationId                      | Delete an integration                                  |
| _IntegrationsApi_ | [**Get**](https://docs.novu.co/platform/integrations)                                      | **Get** /integrations                                        | Get all integrations                                   |
| _IntegrationsApi_ | [**GetActive**](https://docs.novu.co/platform/intergations)                                | **Get** /integrations/active                                 | Get all active integrations                            |
| _IntegrationsApi_ | [**SetIntegrationAsPrimary**](https://docs.novu.co/platform/intergations)                  | **Post** /integrations/{integrationId}/set-primary           | Set the integration as primary                         |
| _IntegrationsApi_ | [**GetChannelLimit**](https://docs.novu.co/platform/intergations)                          | **Get** /integrations/{channelType}/limit                    | Get the limits of the channel                          |

_InboundParserApi_ | [**Get**](https://docs.novu.co/platform/inbound-parse-webhook/) | **Get** /inbound-parse/mx/status | Validate the mx record setup for the inbound parse functionality

## Authorization (api-key)

- **Type**: API key
- **API key parameter name**: ApiKey
- **Location**: HTTP header

## Support and Feedback

Be sure to visit the Novu official [documentation website](https://docs.novu.co/docs) for additional information about our API.

If you find a bug, please post the issue on [Github](https://github.com/novuhq/go-novu/issues).

As always, if you need additional assistance, join our Discord us a note [here](https://discord.gg/novu).

## Contributors

| Name                                           |
| ---------------------------------------------- |
| [Oyewole Samuel](https://github.com/samsoft00) |
| [Dima Grossman](https://github.com/scopsy)     |
