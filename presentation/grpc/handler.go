package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/shellingford330/auth/pkg/grpc/go/auth"
	"github.com/shellingford330/auth/usecase"
)

type Handler struct {
	pb.UnimplementedAuthServer

	*usecase.UseCase
}

func NewHandler(u *usecase.UseCase) *Handler {
	return &Handler{UseCase: u}
}

func (s *Handler) VerifyAccessToken(ctx context.Context, r *pb.VerifyAccessTokenRequest) (*pb.VerifyAccessTokenResponse, error) {
	log.Printf("accept to verify access token. accessToken=%s\n", r.AccessToken)
	user, err := s.UserUseCase.VerifyAccessToken(ctx, r.AccessToken)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to verify access token: %w", err)
	}
	return &pb.VerifyAccessTokenResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}, nil
}
