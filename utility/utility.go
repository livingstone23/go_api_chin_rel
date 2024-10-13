package utility

import (
	"encoding/json"
	"net/http"
)

// Function to respond in JSON format
// help to parse the response in JSON format
func ResponderJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Livingstone Cano", "www.livingstonecano.com")
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(code)
	// Write the response
	w.Write(response)

}