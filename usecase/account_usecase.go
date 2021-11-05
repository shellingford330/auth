package usecase

import (
	"context"
	"fmt"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type AccountUseCase interface {
	LinkAccount(ctx context.Context, params *LinkAccountParams) (*model.Account, error)
}

type accountUseCaseImpl struct {
	repository.AccountRepository
}

func NewAccountUseCase(a repository.AccountRepository) AccountUseCase {
	return &accountUseCaseImpl{a}
}

func (a *accountUseCaseImpl) LinkAccount(ctx context.Context, params *LinkAccountParams) (*model.Account, error) {
	account, err := model.NewAccount(params.ProviderID, params.ProviderType, params.ProviderAccountID, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize account: %w", err)
	}
	account, err = a.AccountRepository.InsertAccount(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("failed to insert account: %w", err)
	}
	return account, nil
}

type LinkAccountParams struct {
	ProviderID        string
	ProviderType      string
	ProviderAccountID string
	UserID            string
}
