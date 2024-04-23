package main

import (
	"fmt"
	"net/http"

	"github.com/dev-crusader404/http-server/handler"
	"github.com/dev-crusader404/http-server/startup"
	mdware "github.com/dev-crusader404/http-server/startup/middleware"
)

func main() {
	startup.Load("")
	// Chain the logger middleware, authentication middleware, and handleHello function for the '/hello' endpoint
	http.HandleFunc("/message", mdware.MethodType(mdware.Logger(mdware.AuthMiddleware(handler.HandleMessage)), http.MethodPost))

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
