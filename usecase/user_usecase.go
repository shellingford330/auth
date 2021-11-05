package usecase

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
	"github.com/shellingford330/auth/usecase/query"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id, email string) (*model.User, error)
	GetUserByProviderAccountID(ctx context.Context, providerID, providerAccountID string) (*model.User, error)
	CreateUser(ctx context.Context, name, email, image string) (*model.User, error)
	UpdateUser(ctx context.Context, id, name, email, image string) (*model.User, error)
}

type userUseCaseImpl struct {
	repository.UserRepository
	query.UserQueryService
}

func NewUserUseCase(r repository.UserRepository, q query.UserQueryService) UserUseCase {
	return &userUseCaseImpl{r, q}
}

func (u *userUseCaseImpl) GetUser(ctx context.Context, id, email string) (*model.User, error) {
	user, err := u.UserRepository.GetUser(ctx, id, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("cannot get user(id=%s))", id)
	}
	return user, nil
}

func (u *userUseCaseImpl) GetUserByProviderAccountID(
	ctx context.Context,
	providerID, providerAccountID string,
) (*model.User, error) {
	user, err := u.UserQueryService.FetchUserByProviderAccountID(ctx, providerID, providerAccountID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCaseImpl) CreateUser(ctx context.Context, name, email, image string) (*model.User, error) {
	user, err := model.NewUser(name, email, image)
	if err != nil {
		return nil, err
	}
	user, err = u.UserRepository.InsertUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCaseImpl) UpdateUser(ctx context.Context, id, name, email, image string) (*model.User, error) {
	user, err := model.NewUser(name, email, image)
	if err != nil {
		return nil, err
	}
	if err := u.UserRepository.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
