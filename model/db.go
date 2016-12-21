package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	host Host
	MyHub Hub
	coinHopper CoinHopper
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	// Mock Init Data
	host = Host{
		Id:     "1",
		OnHand: 0,
	}
	MyHub = Hub{
		Broadcast:    make(chan Msg),
		AddClient:    make(chan *Client),
		RemoveClient: make(chan *Client),
		Clients:      make(map[*Client]bool),
	}
	coinHopper = CoinHopper{
		Status: "ready",
	}
}

