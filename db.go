package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/hello_world_test")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
