package repository

import (
	"context"

	"github.com/shellingford330/auth/domain/model"
)

type SessionRepository interface {
	InsertSession(ctx context.Context, session *model.Session) (*model.Session, error)
}