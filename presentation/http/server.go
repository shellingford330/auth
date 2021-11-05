package http

import (
	"log"
	"net/http"

	"github.com/google/wire"
	"github.com/shellingford330/auth/presentation"
	"github.com/shellingford330/auth/presentation/http/handler"
	"github.com/shellingford330/auth/presentation/http/middleware"
)

type Server struct {
	*handler.Handler
}

var Set = wire.NewSet(NewServer, handler.NewHandler)

func NewServer(h *handler.Handler) presentation.Server {
	return &Server{h}
}

func (s *Server) Start() error {
	http.HandleFunc("/user/create", middleware.Use(s.UserHandler.HandleCreate))
	http.HandleFunc("/user", middleware.Use(s.UserHandler.HandleGet))
	http.HandleFunc("/user/account", middleware.Use(s.UserHandler.HandleGetByProviderAccountID))
	http.HandleFunc("/user/update", middleware.Use(s.UserHandler.HandleUpdate))
	http.HandleFunc("/account/link", middleware.Use(s.AccountHandler.HandleLinkAccount))
	http.HandleFunc("/session/create", middleware.Use(s.SessionHandler.HandleCreate))
	http.HandleFunc("/session", middleware.Use(s.SessionHandler.HandleGet))
	http.HandleFunc("/session/update", middleware.Use(s.SessionHandler.HandleUpdate))
	http.HandleFunc("/session/delete", middleware.Use(s.SessionHandler.HandleDelete))
	// TODO: ENVから取得する
	log.Println("Server running ...")
	return http.ListenAndServe(":8080", nil)
}
