package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shellingford330/auth/handler"
	"github.com/shellingford330/auth/infra/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:password@/shellingford")
	if err != nil {
		panic(err)
	}
}

func main() {
	userHandler := handler.UserHandler{
		UserRepository: mysql.NewUserRepository(DB),
	}
	http.HandleFunc("/user/create", userHandler.HandleCreate)
	http.HandleFunc("/user", userHandler.HandleGet)
	http.ListenAndServe(":8080", nil)
}
