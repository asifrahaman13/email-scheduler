package main

import (
	"fmt"
	"github.com/asifrahaman13/event_management/operations"
  
	"log"
	"net/http"
)

func main() {

    // Initialize the connection function to connect to teh mongodb database. 
  

	http.HandleFunc("/", operations.HandleRequest)

    http.HandleFunc("/insert-email", operations.InsertEmail)
	fmt.Println("Server started ...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
