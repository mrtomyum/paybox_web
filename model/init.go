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
	PM *Payment
	CB *CashBox
	AV *AcceptedValue
	AB *AcceptedBill
	S  *Sale
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	H = &Host{
		Id:                      "001",
		IsNetOnline:             true,
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
	PM = &Payment{
		Coin: 0,
		Bill: 0,
		Total:0,
		Remain: 0,
		Received: make(chan *Message),
	}
	CB = &CashBox{
		Hopper:1000, // todo: เพิ่ม API สั่งเพิ่มเหรียญ
		Coin:  0,
		Bill:  0,
		Total: 0,
	}
	AV = &AcceptedValue{
		B20:  0,
		B50:  0,
		B100: 0,
		B500: 300,
		B1000:700,
	}
	AB = &AcceptedBill{
		B20:  true,
		B50:  true,
		B100: true,
		B500: true,
		B1000:true,
	}
	S = &Sale{
		Total:   0,
		Payment: 0,
		Change:  0,
	}
}
