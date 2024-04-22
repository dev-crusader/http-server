package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type Message struct {
	Text      string `json:"text"`
	RequestID string `json:"request_id"`
}

type RequestID string

var (
	requestIDKey = RequestID("requestID")
	mutex        sync.Mutex
)

func generateRequestID() string {
	mutex.Lock()
	id := uuid.New()
	mutex.Unlock()
	return id.String()
}

func main() {

	// Chain the logger middleware, authentication middleware, and handleHello function for the '/hello' endpoint
	http.HandleFunc("/message", logger(authenticate(handleMessage)))

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

// Custom handler function with additional parameters
func handleMessage(w http.ResponseWriter, r *http.Request, auth bool) {
	requestID := r.Context().Value(requestIDKey).(string)

	// Create a new message struct
	message := Message{Text: "Hello, World!", RequestID: requestID}

	// Encode the message struct into JSON
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content-type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if username != "abc" || password != "123" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add the user ID to the request context
		ctx := context.WithValue(r.Context(), "user_id", "123")
		next(w, r.WithContext(ctx))
	}
}

// Middleware function for request logging
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		// Log information about the incoming request
		fmt.Printf("[%s] Request: %s %s\n", requestID, r.Method, r.URL.Path)

		// Attach requestID to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, requestIDKey, requestID)

		// Call the next handler with the updated request context
		r = r.WithContext(ctx)
		next(w, r)
	}
}
