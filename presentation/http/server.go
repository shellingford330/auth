package http

import (
	"log"
	"net/http"

	"github.com/google/wire"
	"github.com/shellingford330/auth/presentation"
	"github.com/shellingford330/auth/presentation/http/handler"
)

type Server struct {
	*handler.Handler
}

var Set = wire.NewSet(NewServer, handler.NewHandler)

func NewServer(h *handler.Handler) presentation.Server {
	return &Server{h}
}

func (s *Server) Start() error {
	http.HandleFunc("/user/create", s.UserHandler.HandleCreate)
	http.HandleFunc("/user", s.UserHandler.HandleGet)
	http.HandleFunc("/user/account", s.UserHandler.HandleGetByProviderAccountID)
	http.HandleFunc("/user/update", s.UserHandler.HandleUpdate)
	http.HandleFunc("/account/link", s.AccountHandler.HandleLinkAccount)
	http.HandleFunc("/session/create", s.SessionHandler.HandleCreate)
	http.HandleFunc("/session", s.SessionHandler.HandleGet)
	http.HandleFunc("/session/update", s.SessionHandler.HandleUpdate)
	// TODO: ENVから取得する
	log.Println("Server running ...")
	return http.ListenAndServe(":8080", nil)
}
