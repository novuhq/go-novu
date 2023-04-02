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

Class | Method                                                                           | HTTP request                            | Description
------------ |----------------------------------------------------------------------------------|-----------------------------------------| -------------
*EventApi* | [**Trigger**](https://docs.novu.co/platform/subscribers#removing-a-subscriber)   | **Post** /events/trigger                | Trigger
*SubscriberApi* | [**Identify**](https://docs.novu.co/platform/subscribers#creating-a-subscriber)  | **Post** /subscribers                   | Create a subscriber
*SubscriberApi* | [**Update**](https://docs.novu.co/platform/subscribers#updating-subscriber-data) | **Put** /subscribers/:subscriberID      | Update subscriber data
*SubscriberApi* | [**Delete**](https://docs.novu.co/platform/subscribers#removing-a-subscriber)    | **Delete** /subscribers/:subscriberID   | Removing a subscriber
*IntegrationsApi* | [**Create**](https://docs.novu.co/platform/integrations)                         | **Post** /integrations                  | Create an integration
*IntegrationsApi* | [**Update**](https://docs.novu.co/platform/integrations)                         | **Put** /integrations/:integrationId    | Update an integration
*IntegrationsApi* | [**Delete**](https://docs.novu.co/platform/integrations)                         | **Delete** /integrations/:integrationId | Delete an integration
*IntegrationsApi* | [**Get**](https://docs.novu.co/platform/integrations)                            | **Get** /integrations                   | Get all integrations
*IntegrationsApi* | [**GetActive**](https://docs.novu.co/platform/intergations)                      | **Get** /integrations/active            | Get all active integrations


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