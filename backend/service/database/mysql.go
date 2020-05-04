package database

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

func InitDb(user, pass, host, db string) (*sql.DB, error) {
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		user,
		pass,
		host,
		db,
		param,
	)
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("fail to open mysql")
	}
	err = sqldb.Ping()
	if err != nil {
		return nil, fmt.Errorf("fail to open mysql")
	}
	return sqldb, nil
}
