package usecase

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int, email string) (*model.User, error)
	CreateUser(ctx context.Context, name, email, image string) (*model.User, error)
}

type userUseCaseImpl struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{r}
}

func (u *userUseCaseImpl) GetUser(ctx context.Context, id int, email string) (*model.User, error) {
	user, err := u.UserRepository.GetUser(context.Background(), id, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("cannot get user(id=%d))", id)
	}
	return user, nil
}

func (u *userUseCaseImpl) CreateUser(ctx context.Context, name, email, image string) (*model.User, error) {
	// TODO: コンストラクタ
	user := &model.User{
		Name:  name,
		Email: email,
		Image: image,
	}
	user, err := u.UserRepository.InsertUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
