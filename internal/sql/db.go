package sql

import "database/sql"

type DB struct {
	Conn *sql.Db
}

func New(db *sql.DB) *DB {
	return &DB{Conn: db}
}
