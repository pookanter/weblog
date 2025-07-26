package admin

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	adminMux := http.NewServeMux()
	adminMux.Handle("/admin/", http.StripPrefix("/admin", mux))
	adminMux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login Page"))
	})
}
