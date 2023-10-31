package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConfigDb() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	cfg := mysql.Config{
		User:   os.Getenv("USER"),
		Passwd: os.Getenv("PASSWD"),
		Net:    os.Getenv("NET"),
		Addr:   os.Getenv("ADDR"),
		DBName: os.Getenv("DBNAME"),
	}

	db, err := sql.Open(os.Getenv("DRIVERNAME"), cfg.FormatDSN())

	if err != nil {
		panic(err)
	}
	fmt.Println("Rodando db")

	return db
}
