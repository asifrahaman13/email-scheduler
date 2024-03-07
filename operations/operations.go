package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asifrahaman13/event_management/connection"
	"github.com/asifrahaman13/event_management/models"
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love a sample response %s!", r.URL.Path[1:])
}

func Helloworld() {
	fmt.Println("Hello World")
}

// InsertEmail function to insert email into the database.
func InsertEmail(w http.ResponseWriter, r *http.Request) {

	// Connect with the mongodb database.
	client, err := connection.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}

	// Disconnect the mongodb database after the function ends.
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %s", err)
		}
	}()

	// Extract the collection. 
	coll := client.Database("email_scheduling").Collection("emails")

	var email models.EmailStruct

	// Decode the request body into the email struct.
	err = json.NewDecoder(r.Body).Decode(&email)

	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}
    
	// Insert into the database collection. 
	result, err := coll.InsertOne(context.TODO(), &email)

	if err != nil {
		log.Fatalf("Error inserting email: %s", err)
	}

	log.Printf("Inserted a single document with ID %s", result.InsertedID)

	fmt.Fprintf(w, "Email Inserted: %+v\n", email)
}
