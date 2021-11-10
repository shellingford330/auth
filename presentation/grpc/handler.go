package grpc

import (
	"context"

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
	user, err := s.UserUseCase.VerifyAccessToken(ctx, r.AccessToken, r.UserId)
	return &pb.VerifyAccessTokenResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}, err
}
