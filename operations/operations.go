package operations

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love a sample response %s!", r.URL.Path[1:])
}

func Helloworld() {
	fmt.Println("Hello World")
}

func InsertEmail(w http.ResponseWriter, r *http.Request) {

	// Print the body of the post request endpoint.

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
    
	// json.Marshal will convert the golang data type into the JSON representation. 
	jsonBody, err:=json.Marshal(string(body))

	if err != nil {
		panic(err)
	}

	fmt.Printf("The json is: %s\n", jsonBody)

	fmt.Fprintf(w, "Email Inserted and the body is: %s\n", string(body))
}
