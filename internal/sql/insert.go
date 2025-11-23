package sql

import (
	"log"
	"strings"
)

func (this *DB) insert[T any](fields T) *DB {
	if this == nil {
		log.Fatal("no database is connected")
	}

	cols := getSelectFields(fields)
	statement := "INSERT INTO TABLENAME ("
	statement += cols
	statement += ") "

	statement += "VALUES ("

}
