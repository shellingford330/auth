package rdb

import (
	"context"
	"database/sql"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type accountRepositoryImpl struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) repository.AccountRepository {
	return &accountRepositoryImpl{db}
}

func (a *accountRepositoryImpl) InsertAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (provider_id, provider_type, provider_account_id, user_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	if _, err = stmt.Exec(account.ProviderID, account.ProviderType, account.ProviderAccountID, account.UserID); err != nil {
		return nil, err
	}
	return account, nil
}
