package middleware

import (
	"log"
	"net/http"
)

func AccessLog(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host=%s, method=%s, path=%s, query=%s\n", r.RemoteAddr, r.Method, r.URL.Path, r.URL.Query())

		nextFunc(w, r)
	}
}
