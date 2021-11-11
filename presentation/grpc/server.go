package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/google/wire"
	pb "github.com/shellingford330/auth/pkg/grpc/go/auth"
	"github.com/shellingford330/auth/presentation"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
}

var _ presentation.Server = (*Server)(nil)

var Set = wire.NewSet(NewServer, NewHandler)

func NewServer(h *Handler) *Server {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, h)
	return &Server{grpcServer}
}

func (s *Server) Start() error {
	// TODO: load config
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", "50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.Serve(lis)
}

func (s *Server) GracefulShutdown() error {
	s.GracefulStop()
	return nil
}
