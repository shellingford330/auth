package service

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type SessionService interface {
	CreateSession(ctx context.Context, newSession *model.Session) (*model.Session, error)
}

type sessionServiceImpl struct {
	repository.SessionRepository
}

func NewSessionService(r repository.SessionRepository) SessionService {
	return &sessionServiceImpl{r}
}

func (s *sessionServiceImpl) CreateSession(ctx context.Context, newSession *model.Session) (*model.Session, error) {
	session, err := s.SessionRepository.GetSessionByUserID(ctx, newSession.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session. userID=%s: %w", newSession.UserID, err)
	}
	if session != nil {
		// 指定されたユーザのセッションが存在する場合、既存のセッションを削除して新たなセッションを作成する
		if err = s.SessionRepository.DeleteSession(ctx, session.SessionToken); err != nil {
			return nil, fmt.Errorf("failed to delete session: %w", err)
		}
	}

	newSession, err = s.SessionRepository.InsertSession(ctx, newSession)
	if err != nil {
		return nil, fmt.Errorf("failed to insert session: %w", err)
	}
	return newSession, nil
}
