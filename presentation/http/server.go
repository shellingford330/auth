package http

import (
	"net/http"

	"github.com/shellingford330/auth/presentation/http/handler"
)

func Serve(h handler.Handler) {
	http.HandleFunc("/user/create", h.UserHandler.HandleCreate)
	http.HandleFunc("/user", h.UserHandler.HandleGet)
	http.HandleFunc("/user/account", h.UserHandler.HandleGetByProviderAccountID)
	http.HandleFunc("/user/update", h.UserHandler.HandleUpdate)
	// TODO: ENVから取得する
	http.ListenAndServe(":8080", nil)
}
