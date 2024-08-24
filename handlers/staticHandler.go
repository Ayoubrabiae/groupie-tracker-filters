package handlers

import (
	"net/http"
	"os"
	"strings"
)

func StaticHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := os.Open("./static" + r.URL.Path)

		if err != nil || strings.HasSuffix(r.URL.Path, "/") {
			ErrorHandler(w, "Page Not Found", http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
