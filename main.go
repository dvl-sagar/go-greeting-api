package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GreetingRequest struct {
	Name string `json:"name"`
}

type GreetingResponse struct {
	Message string `json:"message"`
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var req GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create the greeting message
	message := fmt.Sprintf("Hello, %s!", req.Name)

	// Send the response
	resp := GreetingResponse{Message: message}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Serve the frontend files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Register the API endpoint
	http.HandleFunc("/api/greet", greetingHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
