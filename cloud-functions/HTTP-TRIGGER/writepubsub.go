// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package function

// [START pubsub_publish_custom_attributes]
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
)

func publishUserToPubSub2(user *User) error {

	projectID := "my-project-id"
	topicID := "my-topic"
	w := new(bytes.Buffer)
	//projectID := "my-project-id"
	//topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	json_data, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintln(w, "Error")
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		ID:   uuid.New().String(),
		Data: json_data,
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published message with custom attributes; msg ID: %v\n", id)
	return nil
}

func publishCustomAttributes(w io.Writer, projectID, topicID string) error {
	// w := buf := new(bytes.Buffer)
	// projectID := "my-project-id"
	// topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello world!"),
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published message with custom attributes; msg ID: %v\n", id)
	return nil
}

// [END pubsub_publish_custom_attributes]
