package mysql

import "database/sql"

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:password@/shellingford")
	if err != nil {
		panic(err)
	}
}
