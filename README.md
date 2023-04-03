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
}
```
**NOTE**
Check the `cmd` directory to see a sample implementation and test files to see sample tests

## Documentation for API Endpoints

Class | Method                                         | HTTP request                          | Description
------------ |------------------------------------------------|---------------------------------------| -------------
*EventApi* | [**Trigger**](https://docs.novu.co/platform/subscribers#removing-a-subscriber) | **Post** /events/trigger              | Trigger
*SubscriberApi* | [**Identify**](https://docs.novu.co/platform/subscribers#creating-a-subscriber) | **Post** /subscribers                 | Create a subscriber
*SubscriberApi* | [**Update**](https://docs.novu.co/platform/subscribers#updating-subscriber-data)     | **Put** /subscribers/:subscriberID    | Update subscriber data
*SubscriberApi* | [**Delete**](https://docs.novu.co/platform/subscribers#removing-a-subscriber)     | **Delete** /subscribers/:subscriberID | Removing a subscriber
*SubscriberApi* | [**Get**](https://docs.novu.co/api/get-a-notification-feed-for-a-particular-subscriber)     | **Get** /subscribers/:subscriberId/notifications/feed | Get a notification feed for a particular subscriber
*SubscriberApi* | [**Get**](https://docs.novu.co/api/get-the-unseen-notification-count-for-subscribers-feed)     | **Get** /subscribers/:subscriberId/notifications/feed | Get the unseen notification count for subscribers feed
*SubscriberApi* | [**Post**](https://docs.novu.co/api/mark-a-subscriber-feed-message-as-seen)     | **Post** /v1/subscribers/:subscriberId/messages/markAs | Mark a subscriber feed message as seen

## Authorization (api-key)

- **Type**: API key
- **API key parameter name**: ApiKey
- **Location**: HTTP header

## Support and Feedback

Be sure to visit the Novu official [documentation website](https://docs.novu.co/docs) for additional information about our API.

If you find a bug, please post the issue on [Github](https://github.com/novuhq/go-novu/issues).

As always, if you need additional assistance, join our Discord us a note [here](https://discord.gg/TT6TttXjRe).

## Contributors

Name |   
------------ |
[Oyewole Samuel](https://github.com/samsoft00) |
[Dima Grossman](https://github.com/scopsy) |