package query

import (
	"context"

	"github.com/shellingford330/auth/domain/model"
)

type UserQueryService interface {
	FetchUserByProviderAccountID(ctx context.Context, providerID, providerAccountID string) (*model.User, error)
}
