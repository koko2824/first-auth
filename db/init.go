package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
	"log"
	"os"
)

var DB *sql.DB

func Init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	tmp, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatalln(err)
	}

	DB = tmp
}