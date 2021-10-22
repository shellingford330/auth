package usecase

import "github.com/google/wire"

type UseCase struct {
	UserUseCase
}

var Set = wire.NewSet(NewUserUseCase)

func NewUseCase(u UserUseCase) *UseCase {
	return &UseCase{u}
}
