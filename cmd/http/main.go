package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shellingford330/auth/handler"
	"github.com/shellingford330/auth/infra/mysql"
	"github.com/shellingford330/auth/usecase"
)

func main() {
	userHandler := handler.UserHandler{
		UserUseCase: usecase.NewUserUseCase(mysql.NewUserRepository(mysql.DB)),
	}
	http.HandleFunc("/user/create", userHandler.HandleCreate)
	http.HandleFunc("/user", userHandler.HandleGet)
	http.ListenAndServe(":8080", nil)
}
