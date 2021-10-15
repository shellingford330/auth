package main

import (
	"net/http"

	"github.com/shellingford330/auth/handler"
)

func main() {
	http.HandleFunc("/user", handler.HandleUserCreate)
	http.ListenAndServe(":8080", nil)
}
