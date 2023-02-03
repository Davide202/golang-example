package main

import (
	"context"
	"log"
	"os"

	_ "reader"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

const (
	projectID      = "my-project-id"
	topicID        = "my-topic"
	subscriptionID = "my-subscription"
)

func main() {

	//sempre necessarie
	_ = os.Setenv("PROJECT_ID", projectID)
	_ = os.Setenv("TOPIC", topicID)
	_ = os.Setenv("SUBSCRIPTION", subscriptionID)

	//necessarie solo in locale
	_ = os.Setenv("FUNCTION_TARGET", "HelloPubSub")
	_ = os.Setenv("PUBSUB_EMULATOR_HOST", "0.0.0.0:8085")
	/**/
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Start: %v\n", err)
	}

	//creo il topic se non esiste
	topic := client.Topic(topicID)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Start: %v\n", err)
	}
	if !ok {
		_, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatalf("Start: %v\n", err)
		}
	}

	//creo la subscription se non esiste
	sub := client.Subscription(subscriptionID)
	ok, err = sub.Exists(ctx)
	if !ok {
		_, err = client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
			Topic:            topic,
			AckDeadline:      10 * time.Second,
			ExpirationPolicy: 25 * time.Hour,
		})
		if err != nil {
			return
		}
	}

	if err != nil {
		log.Fatalf("Start: %v\n", err)
	}

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("Start: %v\n", err)
	}
}
