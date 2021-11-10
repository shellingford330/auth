package usecase

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
	"github.com/shellingford330/auth/domain/service"
	"github.com/shellingford330/auth/usecase/query"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id, email string) (*model.User, error)
	GetUserByProviderAccountID(ctx context.Context, providerID, providerAccountID string) (*model.User, error)
	CreateUser(ctx context.Context, name, email, image string) (*model.User, error)
	UpdateUser(ctx context.Context, id, name, email, image string) (*model.User, error)
	VerifyAccessToken(ctx context.Context, accessToken, userID string) (*model.User, error)
}

type userUseCaseImpl struct {
	repository.UserRepository
	query.UserQueryService
	service.SessionService
}

func NewUserUseCase(r repository.UserRepository, q query.UserQueryService, s service.SessionService) UserUseCase {
	return &userUseCaseImpl{r, q, s}
}

func (u *userUseCaseImpl) GetUser(ctx context.Context, id, email string) (*model.User, error) {
	user, err := u.UserRepository.GetUser(ctx, id, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCaseImpl) GetUserByProviderAccountID(
	ctx context.Context,
	providerID, providerAccountID string,
) (*model.User, error) {
	user, err := u.UserQueryService.FetchUserByProviderAccountID(ctx, providerID, providerAccountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by provider account id. providerID=%s, providerAccountID=%s: %w",
			providerID, providerAccountID, err)
	}
	return user, nil
}

func (u *userUseCaseImpl) CreateUser(ctx context.Context, name, email, image string) (*model.User, error) {
	user, err := model.NewUser(name, email, image)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize user. %w", err)
	}
	user, err = u.UserRepository.InsertUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user. %w", err)
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

func (u *userUseCaseImpl) VerifyAccessToken(ctx context.Context, accessToken, userID string) (*model.User, error) {
	if err := u.SessionService.VerifyAccessToken(ctx, accessToken, userID); err != nil {
		return nil, fmt.Errorf("failed to verify accessToken: %w", err)
	}

	user, err := u.GetUser(ctx, userID, "")
	if err != nil {
		return nil, fmt.Errorf("failed to get user. userID=%s: %w", userID, err)
	}
	return user, nil
}
