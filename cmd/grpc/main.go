package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shellingford330/auth/infra/rdb"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation/grpc"
	"github.com/shellingford330/auth/usecase"
)

func main() {
	server := &grpc.Server{
		UseCase: usecase.NewUseCase(usecase.NewUserUseCase(
			rdb.NewUserRepository(mysql.DB),
			rdb.NewUserQueryService(mysql.DB))),
	}
	grpc.Serve(server)
}
