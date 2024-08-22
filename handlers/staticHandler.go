package handlers

import (
	"net/http"
	"strings"
)

func StaticHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			ErrorHandler(w, "Page Not Found", http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
