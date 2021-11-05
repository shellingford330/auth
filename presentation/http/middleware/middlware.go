package middleware

import (
	"net/http"
)

func Use(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		AccessLog(CommonHeader(nextFunc))(w, r)
	}
}
