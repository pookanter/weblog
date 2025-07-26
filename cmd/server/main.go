package main

import (
	"fmt"
	"net/http"
	"weblog/internal/routers"
)

func main() {
	mux := http.NewServeMux()
	routers.RegisterRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
