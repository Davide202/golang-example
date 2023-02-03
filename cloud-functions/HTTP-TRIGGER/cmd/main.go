package main

import (
	"context"
	"log"
	"os"
	_ "my-cloud-function"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

const (
	projectID = "my-project-id"
	topicID   = "my-topic"
)

func main() {

	//sempre necessarie
	err := os.Setenv("PROJECT_ID", projectID)
	err = os.Setenv("TOPIC", topicID)

	//necessarie solo in locale
	err = os.Setenv("FUNCTION_TARGET", "publisher")
	err = os.Setenv("PUBSUB_EMULATOR_HOST", "0.0.0.0:8085")

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)

	topic := client.Topic(topicID)

	ok, err := topic.Exists(ctx)

	if !ok {
		_, err = client.CreateTopic(ctx, topicID)
	}

	if err != nil {
		log.Fatalf("Start: %v\n", err)
	}

	port := "8081"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("Start: %v\n", err)
	}
}
