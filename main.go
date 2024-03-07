package main

import (
	"fmt"
	"github.com/asifrahaman13/event_management/operations"
    "github.com/asifrahaman13/event_management/connection"
	"log"
	"net/http"
)

func main() {

    // Initialize the connection function to connect to teh mongodb database. 
    connection.Connection()

	http.HandleFunc("/", operations.HandleRequest)

    http.HandleFunc("/insert-email", operations.InsertEmail)
	fmt.Println("Server started ...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
