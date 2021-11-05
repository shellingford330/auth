package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
	"github.com/shellingford330/auth/usecase/query"
)

type SessionUseCase interface {
	CreateSession(ctx context.Context, params *CreateSessionParams) (*model.Session, error)
	GetSession(ctx context.Context, sessionToken string) (*model.Session, error)
	UpdateSessionExpires(ctx context.Context, id string, expires time.Time) (*model.Session, error)
}

type sessionUseCaseImpl struct {
	repository.SessionRepository
	query.SessionQueryService
}

func NewSessionUseCase(r repository.SessionRepository, q query.SessionQueryService) SessionUseCase {
	return &sessionUseCaseImpl{r, q}
}

func (s *sessionUseCaseImpl) CreateSession(ctx context.Context, params *CreateSessionParams) (*model.Session, error) {
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

func (s *sessionUseCaseImpl) GetSession(ctx context.Context, sessionToken string) (*model.Session, error) {
	session, err := s.SessionQueryService.GetSessionBySessionToken(ctx, sessionToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get session. sessionToken=%s: %w", sessionToken, err)
	}
	return session, nil
}

func (s *sessionUseCaseImpl) UpdateSessionExpires(ctx context.Context, id string, expires time.Time) (*model.Session, error) {
	session, err := s.SessionQueryService.GetSessionByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("session not found. id=%s: %w", id, err)
	}
	if err := s.SessionRepository.UpdateSession(ctx, session); err != nil {
		return nil, fmt.Errorf("failed to update session expires.: %w", err)
	}
	return session, nil
}
