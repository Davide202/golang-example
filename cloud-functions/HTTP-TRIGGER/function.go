package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/google/uuid"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Info    string `json:"info"`
}

func (u User) String() string {
	return fmt.Sprintf("\nName: %v, \nSurname %v", u.Name, u.Surname)
}

func init() {
	functions.HTTP("publisher", welcome)
}

func welcome(w http.ResponseWriter, r *http.Request) {

	nameQ := r.FormValue("name")
	surnameQ := r.FormValue("surname")

	var user = User{
		Name:    nameQ,
		Surname: surnameQ,
	}
	log.Printf("User: %v", user.String())
	var err error

	user.Info, err = publishUserToPubSub(w, &user)
	if err != nil {
		fmt.Fprintln(w, "Error")
	}
	json_data, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintln(w, "Error")
	} else {
		//publishUserToPubSub(w, &user)
		//io.WriteString([]string{json_data})
		w.Header().Add("Content-Type", "application/json")
		w.Write(json_data)
		//res := bytes.NewBuffer(json_data)
		//fmt.Fprintln(w, res) //http://localhost:8080?name=Davide&surname=D'Innocente
	}
}

func publishUserToPubSub(w io.Writer, user *User) (string, error) {
	//w := new(bytes.Buffer)
	//projectID := "my-project-id"
	//topicID := "my-topic"
	projectID := os.Getenv("PROJECT_ID")
	topicID := os.Getenv("TOPIC")
	var info string

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return info, fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	json_data, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintln(w, "Error")
	}

	var msg pubsub.Message = pubsub.Message{
		ID:   uuid.New().String(),
		Data: json_data,
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &msg)

	id, err := result.Get(ctx)
	if err != nil {
		return info, fmt.Errorf("get: %v", err)
	}
	info = fmt.Sprintf("Published message with custom attributes; msg ID: %v\n", id)
	//fmt.Fprintf(w, "Published message with custom attributes; msg ID: %v\n", id)
	return info, nil
}
