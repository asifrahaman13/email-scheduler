package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/asifrahaman13/event_management/connection"
	"github.com/asifrahaman13/event_management/models"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love a sample response %s!", r.URL.Path[1:])
}

func Helloworld() {
	fmt.Println("Hello World")
}

func InsertEmail(w http.ResponseWriter, r *http.Request) {
	client, err := connection.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %s", err)
		}
	}()

	coll := client.Database("email_scheduling").Collection("emails")

	var email models.EmailStruct
	err = json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}

	result, err := coll.InsertOne(context.TODO(), &email)
	if err != nil {
		log.Fatalf("Error inserting email: %s", err)
	}

	log.Printf("Inserted a single document with ID %s", result.InsertedID)

	fmt.Fprintf(w, "Email Inserted: %+v\n", email)
}
