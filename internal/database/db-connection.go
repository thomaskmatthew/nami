package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

func Connection() *DataBase {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	return &DataBase{DB: db}

}

func (db *DataBase) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DataBase) CreateTable(statement string) bool {
	_, err := db.Exec(statement)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (db *DataBase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}

func (db *DataBase) QueryRow(query string, args ...interface{}) *sql.Row {
	row := db.DB.QueryRow(query, args...)
	return row
}

func (db *DataBase) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.DB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
