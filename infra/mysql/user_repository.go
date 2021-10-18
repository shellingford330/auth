package mysql

import (
	"context"
	"database/sql"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	stmt, err := u.DB.Prepare("INSERT INTO users (name, email, image) VALUES (?, ?, ?) RETURNING id")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(user.Name, user.Email, user.Image).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImpl) GetUser(ctx context.Context, id int) (*model.User, error) {
	user := model.User{ID: id}
	err := u.DB.QueryRow("SELECT name, email, image FROM users WHERE id = ?", id).Scan(&user.Name, &user.Email, &user.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
