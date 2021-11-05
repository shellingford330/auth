package usecase

import "github.com/google/wire"

type UseCase struct {
	UserUseCase
	AccountUseCase
	SessionUseCase
}

var Set = wire.NewSet(NewUseCase, NewUserUseCase, NewAccountUseCase, NewSessionUseCase)

func NewUseCase(u UserUseCase, a AccountUseCase, s SessionUseCase) *UseCase {
	return &UseCase{u, a, s}
}
