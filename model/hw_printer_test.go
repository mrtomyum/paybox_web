package model_test

import (
	"testing"
	"github.com/mrtomyum/paybox_web/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

var (
	db *sqlx.DB
	H  *model.Host
	BA *model.BillAcceptor
	CA *model.CoinAcceptor
	CH *model.CoinHopper
	p  *model.Printer
	M  *model.MainBoard
	PM *model.Payment
	CB *model.CashBox
	AV *model.AcceptedValue
	AB *model.AcceptedBill
	s  *model.Sale
)

func init() {
	db = sqlx.MustConnect("sqlite3", "./paybox.db")
	H = &model.Host{
		Id:          "001",
		IsNetOnline: true,
	}
	BA = &model.BillAcceptor{
		Status: "ok",
		Send:   make(chan *model.Message),
	}

	CA = &model.CoinAcceptor{
		Status: "ok",
		Send:   make(chan *model.Message),
	}

	CH = &model.CoinHopper{
		Status: "ok",
		Send:   make(chan *model.Message),
	}
	p = &model.Printer{
		Status: "ok",
		Send:   make(chan *model.Message),
	}
	M = &model.MainBoard{
		Status: "ok",
		Send:   make(chan *model.Message),
	}
	PM = &model.Payment{

	}
	CB = &model.CashBox{

	}
	AV = &model.AcceptedValue{
		B20:   0,
		B50:   0,
		B100:  0,
		B500:  300,
		B1000: 700,
	}
	AB = &model.AcceptedBill{
		B20:   true,
		B50:   true,
		B100:  true,
		B500:  true,
		B1000: true,
	}
	s = &model.Sale{
		HostId:   "001",
		Total:    90,
		Pay:      100,
		Change:   10,
		Type:     "TICKET",
		IsPosted: false,
	}
	ss1 := &model.SaleSub{
		Line:     1,
		ItemId:   12345,
		ItemName: "คาปูชิโน่ร้อน",
		PriceId:  43,
		Price:    35.00,
		Qty:      2,
		Unit:     "แก้ว",
	}
	ss2 := &model.SaleSub{
		Line:     1,
		ItemId:   12345,
		ItemName: "คาปูชิโน่ร้อน",
		PriceId:  43,
		Price:    35.00,
		Qty:      2,
		Unit:     "แก้ว",
	}
	ss := make([]*model.SaleSub, 10)
	ss = append(ss, ss1, ss2)
	s.SaleSubs = ss

}

func TestPrinter_doTicket(t *testing.T) {
	x := p.doTicket(s)
	fmt.Println(x)
}