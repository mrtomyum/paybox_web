package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	host Host
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")

	// Mock Init Data
	host = Host{
		Id:     "1",
		OnHand: 0,
	}

}

