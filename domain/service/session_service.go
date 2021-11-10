package service

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/repository"
)

type SessionService interface {
	VerifyAccessToken(ctx context.Context, accessToken, userID string) error
}

type sessionServiceImpl struct {
	repository.SessionRepository
}

func NewSessionService(r repository.SessionRepository) SessionService {
	return &sessionServiceImpl{r}
}

// verify that session of the access token belongs to the user
func (s *sessionServiceImpl) VerifyAccessToken(ctx context.Context, accessToken, userID string) error {
	session, err := s.GetSessionByAccessToken(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("failed to get session by accessToken: %w", err)
	}
	if session == nil {
		return fmt.Errorf("session not found. accessToken=%s: %w", accessToken, err)
	}

	if session.UserID != userID {
		return fmt.Errorf("the session of the accessToken doesn't belong to the user. userID=%s: %w", userID, err)
	}
	return nil
}
