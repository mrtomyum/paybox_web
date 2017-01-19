package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	H  *Host
	BA *BillAcceptor
	CA *CoinAcceptor
	CH *CoinHopper
	P  *Printer
	M  *MainBoard
	OH *Onhand
	CB *CashBox
	AB *AcceptedBill
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	H = &Host{
		Id:                      "001",
		IsNetOnline:             true,
		TotalEscrow:             100,
		BillEscrow:              100,
		TotalBill:               0,
		TotalCoinHopper:         0,
		TotalCainBox:            0,
		TotalCash:               0,
	}
	BA = &BillAcceptor{
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
	OH = &Onhand{
		Coin: 0,
		Bill: 0,
		Total:0,
	}
	CB = &CashBox{
		Hopper:0,
		Coin:  0,
		Bill:  0,
		Total: 0,
	}
	AB = &AcceptedBill{
		THB20:  true,
		THB50:  true,
		THB100: true,
		THB500: false,
		THB1000:false,
	}
}
