package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func ConfigDb() *sql.DB {

	cfg := mysql.Config{
		User:   "",
		Passwd: "",
		Net:    "",
		Addr:   "",
		DBName: "",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	// defer db.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("Rodando db")

	return db
}
