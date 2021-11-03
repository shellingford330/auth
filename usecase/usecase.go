package usecase

import "github.com/google/wire"

type UseCase struct {
	UserUseCase
	AccountUseCase
}

var Set = wire.NewSet(NewUseCase, NewUserUseCase, NewAccountUseCase)

func NewUseCase(u UserUseCase, a AccountUseCase) *UseCase {
	return &UseCase{u, a}
}
