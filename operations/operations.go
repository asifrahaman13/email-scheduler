package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/asifrahaman13/event_management/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love a sample response %s!", r.URL.Path[1:])
}

func Helloworld() {
	fmt.Println("Hello World")
}
var client *mongo.Client



func InsertEmail(w http.ResponseWriter, r *http.Request) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file is specified.")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
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

