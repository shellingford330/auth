package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shellingford330/auth/infra/mysql"
	"github.com/shellingford330/auth/presentation/grpc"
	"github.com/shellingford330/auth/usecase"
)

func main() {
	server := &grpc.Server{
		UseCase: usecase.NewUseCase(usecase.NewUserUseCase(mysql.NewUserRepository(mysql.DB))),
	}
	grpc.Serve(server)
}
