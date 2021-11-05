package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/shellingford330/auth/pkg/grpc/go/auth"
	"github.com/shellingford330/auth/usecase"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAuthServer

	*usecase.UseCase
}

func Serve(server *Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "localhost", "8081"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuthServer(grpcServer, server)
	grpcServer.Serve(lis)
}

func (s *Server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.UserUseCase.CreateUser(ctx, r.Name, r.Email, r.Image)
	return &pb.CreateUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}, err
}
