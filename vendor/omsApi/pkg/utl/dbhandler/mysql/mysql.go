package dbhandler

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func New(driverName string, dataSourceName string, maxPool int, maxIdle int) (*sql.DB, error) {

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxPool)
	db.SetMaxIdleConns(maxIdle)

	err = db.Ping()
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
		return nil, err
	}

	return db, nil
}
