//go:build wireinject
// +build wireinject

package http

import (
	"github.com/google/wire"
	"github.com/shellingford330/auth/domain/service"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation/http"
	"github.com/shellingford330/auth/usecase"
)

func InitServer() (*http.Server, error) {
	wire.Build(
		http.Set,
		service.Set,
		usecase.Set,
		mysql.Set,
	)
	return nil, nil
}
