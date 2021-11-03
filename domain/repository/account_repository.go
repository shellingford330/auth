package repository

import (
	"context"

	"github.com/shellingford330/auth/domain/model"
)

type AccountRepository interface {
	InsertAccount(ctx context.Context, account *model.Account) (*model.Account, error)
}
