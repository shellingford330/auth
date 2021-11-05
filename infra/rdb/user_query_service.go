package rdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/usecase/query"
)

type userQueryServiceImpl struct {
	DB *sql.DB
}

func NewUserQueryService(db *sql.DB) query.UserQueryService {
	return &userQueryServiceImpl{db}
}

func (u *userQueryServiceImpl) FetchUserByProviderAccountID(
	ctx context.Context,
	providerID, providerAccountID string,
) (*model.User, error) {
	user := model.User{}
	err := u.DB.QueryRow(
		"SELECT `users`.`id`, `users`.`name`, `users`.`email`, `users`.`image` FROM `users` INNER JOIN `accounts` ON `accounts`.`user_id` = `users`.`id` WHERE `accounts`.`provider_id` = ? AND `accounts`.`provider_account_id = ?`",
		providerID, providerAccountID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
