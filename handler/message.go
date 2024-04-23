package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dev-crusader404/http-server/models"
	md "github.com/dev-crusader404/http-server/startup/middleware"
)

// Custom handler function with additional parameters
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	requestID := r.Context().Value(md.RequestIDKey).(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	request := models.MessageBody{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new message struct
	message := models.Message{
		User:      request.User,
		Text:      fmt.Sprintf("Hi, %s! Received your message: '%s'", request.User, request.Msg),
		RequestID: requestID}

	resp := models.HTTPResponse{
		Status:      "ok",
		Message:     message,
		TimeElapsed: time.Since(startTime).Milliseconds(),
	}

	// Encode the message struct into JSON
	response, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content-type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
