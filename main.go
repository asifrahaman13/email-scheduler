package main

import (
	"fmt"
	"github.com/asifrahaman13/event_management/operations"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	// Initialize the connection function to connect to teh mongodb database.

	router.HandleFunc("/", operations.HandleRequest)

	router.HandleFunc("/insert-email", operations.InsertEmail).Methods("POST")
	router.HandleFunc("/all-email", operations.AllEmails).Methods("GET")

	// router.HandleFunc("/send-email", operations.Email).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Server started ...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
