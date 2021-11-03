package handler

import "github.com/shellingford330/auth/usecase"

type Handler struct {
	*UserHandler
	*AccountHandler
}

func NewHandler(u usecase.UserUseCase, a usecase.AccountUseCase) *Handler {
	return &Handler{
		UserHandler:    NewUserHandler(u),
		AccountHandler: NewAccountHandler(a),
	}
}
