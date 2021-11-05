// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/shellingford330/auth/infra/rdb"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation"
	"github.com/shellingford330/auth/presentation/http"
	"github.com/shellingford330/auth/presentation/http/handler"
	"github.com/shellingford330/auth/usecase"
)

// Injectors from wire.go:

func initializeServer() presentation.Server {
	db := _wireDBValue
	userRepository := rdb.NewUserRepository(db)
	userQueryService := rdb.NewUserQueryService(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userQueryService)
	accountRepository := rdb.NewAccountRepository(db)
	accountUseCase := usecase.NewAccountUseCase(accountRepository)
	sessionRepository := rdb.NewSessionRepository(db)
	sessionQueryService := rdb.NewSessionQueryService(db)
	sessionUseCase := usecase.NewSessionUseCase(sessionRepository, sessionQueryService)
	handlerHandler := handler.NewHandler(userUseCase, accountUseCase, sessionUseCase)
	server := http.NewServer(handlerHandler)
	return server
}

var (
	_wireDBValue = mysql.DB
)
