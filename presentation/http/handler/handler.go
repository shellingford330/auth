package handler

import (
	"net/http"

	"github.com/shellingford330/auth/presentation/http/middleware"
	"github.com/shellingford330/auth/usecase"
)

func New(u usecase.UserUseCase, a usecase.AccountUseCase, s usecase.SessionUseCase) http.Handler {
	userHandler := NewUserHandler(u)
	accountHandler := NewAccountHandler(a)
	sessionHandler := NewSessionHandler(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/user/create", middleware.Use(userHandler.HandleCreate))
	mux.HandleFunc("/user", middleware.Use(userHandler.HandleGet))
	mux.HandleFunc("/user/account", middleware.Use(userHandler.HandleGetByProviderAccountID))
	mux.HandleFunc("/user/update", middleware.Use(userHandler.HandleUpdate))
	mux.HandleFunc("/account/link", middleware.Use(accountHandler.HandleLinkAccount))
	mux.HandleFunc("/session/create", middleware.Use(sessionHandler.HandleCreate))
	mux.HandleFunc("/session", middleware.Use(sessionHandler.HandleGet))
	mux.HandleFunc("/session/update", middleware.Use(sessionHandler.HandleUpdate))
	mux.HandleFunc("/session/delete", middleware.Use(sessionHandler.HandleDelete))
	return mux
}
