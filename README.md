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
	apiKey := "c654b3ac3b5cee35a741f649daa3192b"
	eventId := "gs-cooperative"

	ctx := context.Background()
	to := map[string]interface{}{
		"lastName":     "Doe",
		"firstName":    "John",
		"subscriberId": "john@doemail.com",
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

## Documentation for API Endpoints

Class | Method                                         | HTTP request                          | Description
------------ |------------------------------------------------|---------------------------------------| -------------
*EventApi* | [**Trigger**](https://docs.novu.co/platform/subscribers#removing-a-subscriber) | **Post** /events/trigger              | Trigger
*SubscriberApi* | [**Identify**](https://docs.novu.co/platform/subscribers#creating-a-subscriber) | **Post** /subscribers                 | Create a subscriber
*SubscriberApi* | [**Update**](https://docs.novu.co/platform/subscribers#updating-subscriber-data)     | **Put** /subscribers/:subscriberID    | Update subscriber data
*SubscriberApi* | [**Delete**](https://docs.novu.co/platform/subscribers#removing-a-subscriber)     | **Delete** /subscribers/:subscriberID | Removing a subscriber

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