package db

import (
	"database/sql"
	"fmt"

	"github.com/gheysiell/todo_backend/configs"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	conn, err := sql.Open("mysql", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
