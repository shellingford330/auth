package handler

import "github.com/shellingford330/auth/usecase"

type Handler struct {
	*UserHandler
	*AccountHandler
	*SessionHandler
}

func NewHandler(u usecase.UserUseCase, a usecase.AccountUseCase, s usecase.SessionUseCase) *Handler {
	return &Handler{
		UserHandler:    NewUserHandler(u),
		AccountHandler: NewAccountHandler(a),
		SessionHandler: NewSessionHandler(s),
	}
}
