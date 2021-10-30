package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shellingford330/auth/infra/rdb"
	"github.com/shellingford330/auth/infra/rdb/mysql"
	"github.com/shellingford330/auth/presentation/http"
	"github.com/shellingford330/auth/presentation/http/handler"
	"github.com/shellingford330/auth/usecase"
)

func main() {
	handler := handler.Handler{
		handler.UserHandler{
			usecase.NewUserUseCase(
				rdb.NewUserRepository(mysql.DB),
				rdb.NewUserQueryService(mysql.DB),
			),
		},
	}
	http.Serve(handler)
}
