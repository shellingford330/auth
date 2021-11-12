package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/shellingford330/auth/internal/config"
	"github.com/shellingford330/auth/presentation"
	"github.com/shellingford330/auth/presentation/http/handler"
)

type Server struct {
	*http.Server
}

var _ presentation.Server = (*Server)(nil)

var Set = wire.NewSet(NewServer, handler.New)

func NewServer(h http.Handler) *Server {
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Config.HTTPServerHost(), config.Config.HTTPServerPort()),
		Handler: h,
	}
	return &Server{s}
}

func (s *Server) Start() error {
	log.Printf("HTTP server running and listen %s ...\n", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) GracefulShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
