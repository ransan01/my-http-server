package netHttp

import (
	"net/http"

	"github.com/ransan01/my-http-server/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("HEAD /users", handlers.GetUsers)

	// Read Operations
	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("GET /users/{id}", handlers.GetUser)

	// Create Operations
	// usually handles a single user object in the request body,
	// can be designed to accept an array of objects to perform a bulk create
	mux.HandleFunc("POST /users", handlers.CreateUser)

	// Update or Replace Operations
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser)

	// Partial Update Operations
	mux.HandleFunc("PATCH /users/{id}", handlers.PatchUser)

	// Delete Operations
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	// Bulk Delete Operations
	mux.HandleFunc("DELETE /users", handlers.DeleteUsers)

	return mux
}
