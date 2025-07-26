package routers

import (
	"net/http"
	"weblog/internal/routers/admin"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Example route registration
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the weblog!"))
	})

	// Add more routes as needed
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	})

	admin.RegisterRoutes(mux)
}
