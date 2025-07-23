package http

import (
	"log"
	"net/http"
)

// NewRouter returns an http.ServeMux instance for routing HTTP requests.
// This is a simple HTTP router from Go's standard library.
func NewRouter() *http.ServeMux {
	// http.NewServeMux is a multiplexer that routes HTTP requests to handlers based on URL paths.
	mux := http.NewServeMux()

	// For example, a general logging middleware could be added for all requests here.
	// However, in Clean Architecture, middlewares are typically defined above the main router,
	// either in main.go or with a more complex wrapper here.
	// For now, let's keep it simple.

	log.Println("HTTP router initialized.")
	return mux
}

// HandlerFuncWrapper creates an adapter suitable for http.HandlerFunc.
// This is a helper function to directly bind Controller methods to http.HandleFunc.
// It can be extended in the future for more advanced middlewares or error handling.
type HandlerFuncWrapper func(http.ResponseWriter, *http.Request)

// Methods is a helper to chain methods, similar to some web frameworks.
// This is a simplified example to make routing definitions cleaner.
// For more robust routing, consider a dedicated router library (Gorilla Mux, Gin, Echo).
func (h HandlerFuncWrapper) Methods(methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := false
		for _, m := range methods {
			if r.Method == m {
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}
