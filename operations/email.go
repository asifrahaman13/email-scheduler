package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asifrahaman13/event_management/connection"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"sync"
)

// Function to send email to the receivers.
func Email(Receiver string, EmailBody string) {

	// Load the .env file.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file is specified.")
	}


	// Extract the email address from the .env file. 
	emailAddress := os.Getenv("EMAIL_ADDRESS")

	if emailAddress == "" {
		log.Println("No email address is specified.")
	}

    // Extract the email password from the .env file.
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	if emailPassword == "" {
		log.Println("No email password is specified.")
	}

	// Sender data.
	from := emailAddress
	password := emailPassword

	// Receiver email address.
	to := []string{
		Receiver,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(EmailBody)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Send a success message. 
	fmt.Println("Email Sent Successfully!")
}

func AllEmails(w http.ResponseWriter, r *http.Request) {

	// Initialize the wait group. 
	var wg sync.WaitGroup

	// Set the content type of the response. 
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "All Emails")
    
	// Connect to mongoose.
	client, err := connection.ConnectToMongoDB()

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}

	// Disconnect in case of any error. 
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %s", err)
		}
	}()

	// Connect to the collection.
	coll := client.Database("email_scheduling").Collection("emails")

	// Create a slice to store the results
	var results []bson.M

	// Pass an empty filter to find all documents
	cursor, err := coll.Find(context.TODO(), bson.D{})
	
	if err != nil {
		log.Fatalf("Error finding documents: %s", err)
	}
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and decode each document into the results slice
	// Iterate over the cursor and decode each document into the results slice
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatalf("Error decoding document: %s", err)
		}
		fmt.Printf("result: %v\n", result)

		// Extract the email field as a string
		email, ok := result["receiveremail"].(string)
		if !ok {
			fmt.Println("Email not found or not a string")
			continue
		}
		fmt.Printf("Email: %s\n", email)

		message, ok := result["emailbody"].(string)
		if !ok {
			fmt.Println("Email body not found or not a string")
			continue
		}
		wg.Add(1)

		// Start a goroutine to send the email
		go func(email, message string) {
			defer wg.Done()
			Email(email, message)
		}(email, message)

		// cursor.Close(context.TODO())
		wg.Wait()

		// Email(email, message)

		results = append(results, result)

	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %s", err)
	}

	// Marshal the results slice into JSON and write it to the response
	jsonResults, err := json.Marshal(results)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %s", err)
	}

	w.Write(jsonResults)
}
