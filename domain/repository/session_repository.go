package repository

import (
	"context"

	"github.com/shellingford330/auth/domain/model"
)

type SessionRepository interface {
	GetSessionByUserID(ctx context.Context, userID string) (*model.Session, error)
	InsertSession(ctx context.Context, session *model.Session) (*model.Session, error)
	UpdateSession(ctx context.Context, session *model.Session) (*model.Session, error)
	DeleteSession(ctx context.Context, sessionToken string) error
}
