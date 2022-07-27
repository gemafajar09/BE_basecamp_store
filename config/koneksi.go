package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Koneksi() *sql.DB {
	cek := godotenv.Load(".env")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DATABASE := os.Getenv("DB_DATABASE")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DATABASE)
	db, cek := sql.Open("mysql", URL)
	if cek != nil {
		panic(cek.Error)
	}
	return db
}
