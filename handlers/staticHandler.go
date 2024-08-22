package handlers

import (
	"net/http"
	"os"
)

func StaticHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := os.Open("./static/styles/master")

		if os.IsNotExist(err) {
			ErrorHandler(w, "Page Not Found", http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
