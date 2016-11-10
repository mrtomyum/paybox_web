package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	// Load Dummy Data.
}

