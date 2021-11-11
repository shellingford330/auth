package service

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type SessionService interface {
	VerifyAccessToken(ctx context.Context, accessToken string) (*model.Session, error)
}

type sessionServiceImpl struct {
	repository.SessionRepository
}

func NewSessionService(r repository.SessionRepository) SessionService {
	return &sessionServiceImpl{r}
}

// verify that session of the access token belongs to the user
func (s *sessionServiceImpl) VerifyAccessToken(ctx context.Context, accessToken string) (*model.Session, error) {
	session, err := s.GetSessionByAccessToken(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get session by accessToken: %w", err)
	}
	if session == nil {
		return nil, fmt.Errorf("session not found. accessToken=%s: %w", accessToken, err)
	}
	return session, nil
}
