package middlewares

import (
	"log"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request details (method, URL, etc.)
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
