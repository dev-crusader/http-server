# Basic Client-Server Application in Go

This is a simple client-server application implemented in Go using the `net/http` package. The server listens for incoming HTTP requests from clients on port 8080 and responds with a simple message. Application properties are loaded from the app.properties file at the startup. 

The server also includes a middleware for authentication and generates a unique request ID for each REST call.
The client makes HTTP calls with basic authentication in the header and receives the message response in JSON format.

## How to run the application
1. Clone the repository:
   ```https://github.com/dev-crusader404/http-server.git```
2. Update the `app.properties` file with the necessary user names and passwords.
3. Run the server:
   `go run main.go`
4. Client code is within the package `https://github.com/dev-crusader404/http-server/client`.
   Run another instance of client:
   `go run client\client.go`
