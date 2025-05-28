package main

import (
	"fmt"
	"net/http"

	"github.com/dev-crusader/http-server/handler"
	"github.com/dev-crusader/http-server/startup"
	mdware "github.com/dev-crusader/http-server/startup/middleware"
)

func main() {
	startup.Load("")
	// Chain the valid methodType, logger middleware, authentication middleware, and HandleMessage function for the '/message' endpoint
	http.HandleFunc("/message", mdware.MethodType(mdware.Logger(mdware.AuthMiddleware(handler.HandleMessage)), http.MethodPost))

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
