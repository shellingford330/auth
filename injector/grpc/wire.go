//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/google/wire"
	"github.com/shellingford330/auth/domain/service"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation/grpc"
	"github.com/shellingford330/auth/usecase"
)

func InitServer() (*grpc.Server, error) {
	wire.Build(
		grpc.Set,
		service.Set,
		usecase.Set,
		mysql.Set,
	)
	return nil, nil
}
