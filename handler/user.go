package handler

import (
	"github.com/shellingford330/auth/domain/repository"
)

type UserHandler struct {
	repository.UserRepository
}
