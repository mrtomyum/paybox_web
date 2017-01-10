package model

import (
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	H  *Host
	B  *BillAcceptor
	C  *CoinAcceptor
	CH *CoinHopper
	P  *Printer
	M  *MainBoard
)

func init() {
	H = &Host{
		Id:            "001",
		Online:        true,
		TotalEscrow:   0,
		BillEscrow:    0,
		TotalBill:       0,
		TotalCoinHopper: 0,
		TotalCainBox:       0,
		TotalCash:     0,

	}
	B = &BillAcceptor{
		Status: "ok",
	}
	C = &CoinAcceptor{
		Status: "ok",
	}

	CH = &CoinHopper{
		Status: "ok",
		Send:   make(chan *Message),
	}
	P = &Printer{
		Status: "ok",
	}
	M = &MainBoard{
		Status: "ok",
	}
}
