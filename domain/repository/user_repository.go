package repository

import (
	"context"

	"github.com/shellingford330/auth/domain/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, id int, email string) (*model.User, error)
}
