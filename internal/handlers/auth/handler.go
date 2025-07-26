package auth

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle login logic here
		w.Write([]byte("Login successful"))
		return
	}
	http.ServeFile(w, r, "templates/login.html")
}
