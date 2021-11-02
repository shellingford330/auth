package model

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email, image string) (*User, error) {
	user := User{}
	if err := user.SetName(name); err != nil {
		return nil, err
	}
	if err := user.SetEmail(email); err != nil {
		return nil, err
	}
	if err := user.SetImage(image); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) SetID(id int) error {
	if id > 0 {
		return errors.New("id is less than or equal 0")
	}
	u.ID = id
	return nil
}

func (u *User) SetName(name string) error {
	if name == "" {
		return errors.New("name is blank")
	}
	u.Name = name
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New("email is blank")
	}
	u.Email = email
	return nil
}

func (u *User) SetImage(image string) error {
	u.Image = image
	return nil
}

func (u *User) SetCreatedAt(createdAt time.Time) error {
	if createdAt.IsZero() {
		return errors.New("created_at is blank")
	}
	u.CreatedAt = createdAt
	return nil
}

func (u *User) SetUpdatedAt(updatedAt time.Time) error {
	if updatedAt.IsZero() {
		return errors.New("updated_at is blank")
	}
	u.UpdatedAt = updatedAt
	return nil
}
