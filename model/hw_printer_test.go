package model_test

import (
	"testing"
	"github.com/mrtomyum/paybox_web/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	H  *model.Host
	BA *model.BillAcceptor
	CA *model.CoinAcceptor
	CH *model.CoinHopper
	P  *model.Printer
	M  *model.MainBoard
	PM *model.Payment
	CB *model.CashBox
	AV *model.AcceptedValue
	AB *model.AcceptedBill
	S  *model.Sale
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
	P = &model.Printer{
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
	S = &model.Sale{
		HostId: "001",
		Total:  0,
		Pay:    0,
		Change: 0,
	}

	//r := gin.Default()
	//app := ctrl.Router(r)
	//go ctrl.CallDev()
	//app.Run(":8888")

}

func TestPrint(t *testing.T) {
	data := `[
        {
            "action": "print",
            "action_data": "Hello World"
        },
        {
        "action": "print",
        "action_data": "Hello World"
        }
    ]`

	err := P.PrintTest(data)
	if err != nil {
		t.Log(err.Error())
	}

}
