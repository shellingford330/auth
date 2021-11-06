package query

import (
	"context"

	"github.com/shellingford330/auth/domain/model"
)

type SessionQueryService interface {
	GetSessionBySessionToken(ctx context.Context, sessionToken string) (*model.Session, error)
}
