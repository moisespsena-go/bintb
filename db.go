package bintb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDBv(name string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", name+"?_vacuum=1")
	if err != nil {
		return
	}
	return
}
func OpenDB(name string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", name)
	if err != nil {
		return
	}
	return
}
