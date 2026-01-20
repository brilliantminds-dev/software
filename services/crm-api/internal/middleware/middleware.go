package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	//logs the requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s | %s | %s | %s", r.Method, r.URL, r.Host, time.Now().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}
