package main

import (
	"context"
	"fmt"
	novu "github.com/novuhq/go-novu/lib"
	"github.com/novuhq/go-novu/utils"
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

	novuClient := novu.NewAPIClient(apiKey, &novu.Config{})

	// Trigger
	triggerResp, err := novuClient.EventApi.Trigger(ctx, eventId, novu.ITriggerPayloadOptions{
		To:      to,
		Payload: payload,
	})
	if err != nil {
		log.Fatal("Novu error", err.Error())
		return
	}

	fmt.Println(triggerResp)

	// Subscriber
	subscriber := novu.SubscriberPayload{
		LastName: "Skjæveland",
		Email:    "benedicte.skjaeveland@example.com",
		Avatar:   "https://randomuser.me/api/portraits/thumb/women/79.jpg",
		Data: map[string]interface{}{
			"location": map[string]interface{}{
				"city":     "Ballangen",
				"state":    "Aust-Agder",
				"country":  "Norway",
				"postcode": "7481",
			},
		},
	}

	resp, err := novuClient.SubscriberApi.Identify(ctx, subscriberID, subscriber)
	if err != nil {
		log.Fatal("Subscriber error: ", err.Error())
		return
	}

	fmt.Println(resp)

	// update subscriber
	updateSubscriber := novu.SubscriberPayload{FirstName: "Susan"}

	updateResp, err := novuClient.SubscriberApi.Update(ctx, subscriberID, updateSubscriber)
	if err != nil {
		log.Fatal("Update subscriber error: ", err.Error())
		return
	}

	fmt.Println(updateResp)

	// delete subscriber
	deleteResp, err := novuClient.SubscriberApi.Delete(ctx, subscriberID)
	if err != nil {
		log.Fatal("Update subscriber error: ", err.Error())
		return
	}
	fmt.Println(deleteResp)

	// Get Notification
	_, err = novuClient.NotificationApi.GetNotifications(ctx,
		utils.NewQueryParam("channels", []string{"z", "q"}),
		utils.NewQueryParam("templates", []string{}),
		utils.NewQueryParam("emails", ""),
		utils.NewQueryParam("search", ""))

	if err != nil {
		panic(err)
	}
}
