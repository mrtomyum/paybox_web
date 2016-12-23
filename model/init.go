package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db         *sqlx.DB
	host Host
	MyHub Hub
	coinHopper CoinHopper
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	// Mock Init Data
	host = Host{
		Id:     "1",
		OnHand: 200,
	}
	MyHub = Hub{
		Clients:      make(map[*Client]bool),
		//Broadcast:    make(chan Msg),
		Send:         make(chan *Client),
		AddClient:    make(chan *Client),
		DelClient: make(chan *Client),
	}
	coinHopper = CoinHopper{
		Status: "ready",
	}
}
