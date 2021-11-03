package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/wire"
	"github.com/shellingford330/auth/infra/rdb"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:password@/shellingford")
	if err != nil {
		panic(err)
	}
}

var Set = wire.NewSet(
	rdb.NewUserRepository,
	rdb.NewAccountRepository,
	rdb.NewUserQueryService,
	wire.Value(DB),
)
