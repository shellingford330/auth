package rdb

import (
	"context"
	"database/sql"
	"strings"
	"time"

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
	stmt, err := u.DB.Prepare("INSERT INTO users (name, email, image) VALUES (?, ?, ?) RETURNING id, created_at, updated_at")
	if err != nil {
		return nil, err
	}
	var id int
	var createdAt, updatedAt time.Time
	if err = stmt.QueryRow(user.Name, user.Email, user.Image).Scan(&id, &createdAt, &updatedAt); err != nil {
		return nil, err
	}
	if err := user.SetID(id); err != nil {
		return nil, err
	}
	if err := user.SetCreatedAt(createdAt); err != nil {
		return nil, err
	}
	if err := user.SetUpdatedAt(updatedAt); err != nil {
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

func (u *userRepositoryImpl) GetUser(ctx context.Context, id int, email string) (*model.User, error) {
	conds, params := []string{}, []interface{}{}
	if id != 0 {
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
