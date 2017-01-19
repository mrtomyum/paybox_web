package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	H  *Host
	B  *BillAcceptor
	CA *CoinAcceptor
	CH *CoinHopper
	P  *Printer
	M  *MainBoard
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	H = &Host{
		Id:                 "001",
		IsNetOnline:             true,
		TotalEscrow:        0,
		BillEscrow:         100,
		TotalBill:          0,
		TotalCoinHopper:    0,
		TotalCainBox:       0,
		TotalCash:          0,

	}
	B = &BillAcceptor{
		Status: "ok",
		Send:   make(chan *Message),
	}
	CA = &CoinAcceptor{
		Status: "ok",
		Send:   make(chan *Message),
	}

	CH = &CoinHopper{
		Status: "ok",
		Send:   make(chan *Message),
	}
	P = &Printer{

		Status: "ok",
		Send:   make(chan *Message),
	}
	M = &MainBoard{
		Status: "ok",
		Send:   make(chan *Message),
	}
}
