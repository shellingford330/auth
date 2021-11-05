package rdb

import (
	"context"
	"database/sql"
	"strings"

	"github.com/shellingford330/auth/domain/model"
	"github.com/shellingford330/auth/domain/repository"
	"github.com/shellingford330/auth/pkg/ulid"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	id := ulid.Generate()
	stmt, err := u.DB.Prepare("INSERT INTO users (id, name, email, image) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	if _, err = stmt.Exec(id, user.Name, user.Email, user.Image); err != nil {
		return nil, err
	}
	if err := user.SetID(id); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImpl) UpdateUser(ctx context.Context, user *model.User) error {
	stmt, err := u.DB.Prepare("UPDATE users SET name = ?, email = ?, image = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Name, user.Email, user.Image, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImpl) GetUser(ctx context.Context, id, email string) (*model.User, error) {
	conds, params := []string{}, []interface{}{}
	if id != "" {
		conds = append(conds, "id = ?")
		params = append(params, id)
	}
	if email != "" {
		conds = append(conds, "email = ?")
		params = append(params, email)
	}
	user := model.User{}
	err := u.DB.QueryRow("SELECT id, name, email, image FROM users WHERE "+strings.Join(conds, " AND "), params...).Scan(&user.ID, &user.Name, &user.Email, &user.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
