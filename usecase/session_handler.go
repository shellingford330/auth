package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type SessionUseCase interface {
	CreateSession(ctx context.Context, params *CreateSessionParams) (*model.Session, error)
}

type sessionUseCaseImpl struct {
	repository.SessionRepository
}

func (s *sessionUseCaseImpl) CreateSession(ctx context.Context, params CreateSessionParams) (*model.Session, error) {
	session, err := model.NewSession(params.Expires, params.SessionToken, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize account: %w", err)
	}
	session, err = s.SessionRepository.InsertSession(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("failed to insert session: %w", err)
	}
	return session, nil
}

type CreateSessionParams struct {
	Expires      time.Time
	SessionToken string
	UserID       string
}
