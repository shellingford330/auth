package http

import (
	"net/http"

	"github.com/shellingford330/auth/presentation/http/handler"
)

func Serve(handler handler.Handler) {
	http.HandleFunc("/user/create", handler.UserHandler.HandleCreate)
	http.HandleFunc("/user", handler.UserHandler.HandleGet)
	// TODO: ENVから取得する
	http.ListenAndServe(":8080", nil)
}
