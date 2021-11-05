//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation"
	"github.com/shellingford330/auth/presentation/http"
	"github.com/shellingford330/auth/usecase"
)

func initializeServer() presentation.Server {
	wire.Build(
		http.Set,
		usecase.Set,
		mysql.Set,
	)
	return nil
}
