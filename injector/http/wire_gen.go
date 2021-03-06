// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http

import (
	"github.com/shellingford330/auth/domain/service"
	"github.com/shellingford330/auth/infra/rdb"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation/http"
	"github.com/shellingford330/auth/presentation/http/handler"
	"github.com/shellingford330/auth/usecase"
)

// Injectors from wire.go:

func InitServer() (*http.Server, error) {
	db := _wireDBValue
	userRepository := rdb.NewUserRepository(db)
	userQueryService := rdb.NewUserQueryService(db)
	sessionRepository := rdb.NewSessionRepository(db)
	sessionService := service.NewSessionService(sessionRepository)
	userUseCase := usecase.NewUserUseCase(userRepository, userQueryService, sessionService)
	accountRepository := rdb.NewAccountRepository(db)
	accountUseCase := usecase.NewAccountUseCase(accountRepository)
	sessionQueryService := rdb.NewSessionQueryService(db)
	sessionUseCase := usecase.NewSessionUseCase(sessionRepository, sessionQueryService)
	httpHandler := handler.New(userUseCase, accountUseCase, sessionUseCase)
	server := http.NewServer(httpHandler)
	return server, nil
}

var (
	_wireDBValue = mysql.DB
)
