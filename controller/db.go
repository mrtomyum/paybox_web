package controller

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func init() {
	DB = sqlx.MustConnect("sqlite3", "./paybox.db")
	// Load Dummy Data.
}

