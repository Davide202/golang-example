package function

import (
	"context"
	"encoding/json"
	
	"fmt"
	"log"
	"os"
	

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Info    string `json:"info"`
}

func (u User) String() string {
	return fmt.Sprintf("\nName: %v, \nSurname: %v", u.Name, u.Surname)
}

func init() {
	functions.CloudEvent("HelloPubSub", helloPubSub)
}

// helloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func helloPubSub(ctx context.Context, e event.Event) error {
	var err error
	var user User
	log.Printf("Received: %v", e.String())

	//VERSION 1
	/*	
	log.Printf("ID: %v", e.ID())
	log.Printf("Extensions: %v", e.Extensions())

	
	if err = e.DataAs(&user); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}
	log.Printf("User: %v", user.String())

	err = json.Unmarshal(e.Data(), &user)
	if err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}
	log.Printf("User: %v", user.String())
*/
	// VERSION 2
	//topicID := os.Getenv("TOPIC")
	projectID := os.Getenv("PROJECT_ID")
	subscriptionID := os.Getenv("SUBSCRIPTION")

	//ctx = context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	sub := client.Subscription(subscriptionID)
	if err != nil {
		log.Fatalf("Start: %v\n", err)
	}
	sub.Receive(ctx,func(ctx context.Context, m *pubsub.Message) {
		err = json.Unmarshal(m.Data, &user)
		log.Printf("User: %v", user.String())
		m.Ack()
	})




	return nil
}
