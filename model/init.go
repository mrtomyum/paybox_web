package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"log"
)

var (
	db    *sqlx.DB
	H     *Host
	BA    *BillAcceptor
	CA    *CoinAcceptor
	CH    *CoinHopper
	P     *Printer
	MB    *MainBoard
	PM    *Payment
	CB    *CashBox
	AV    *AcceptedValue
	AB    *AcceptedBill
	Items []*Item
)

func init() {
	pwd, _ := os.Getwd()
	db = sqlx.MustConnect("sqlite3", pwd+"/paybox.db")

	PM = &Payment{
		coin:       0,
		bill:       0,
		billEscrow: 0,
		total:      0,
		remain:     0,
		//receivedCh: make(chan *Message),
	}
	H = &Host{
		Id:          "001",
		IsNetOnline: true,
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
	//Todo: เพิ่ม Struct เก็บจำนวนเหรียญแต่ละขนาดไว้ด้วย แล้วกำหนดให้ตรวจสอบยอดคงเหลือจาก coin_hopper ตอนเริ่มต้น และทุกขั้นตอนที่มีการรับ coin_acc จ่าย coin_hopper
	P = &Printer{
		Status: "ok",
		Send:   make(chan *Message),
	}
	MB = &MainBoard{
		Status:  "ok",
		Send:    make(chan *Message),
		PinOpen: 15, // <- เปลี่ยนหมายเลขพินต่อ Magnetic Sensor ทีนี่
	}

	CB = &CashBox{
		hopper: 1000, // todo: เพิ่ม API สั่งเพิ่มเหรียญ
		coin:   0,
		bill:   0,
		total:  0,
	}
	AV = &AcceptedValue{
		B20:   0,
		B50:   0,
		B100:  0,
		B500:  300,
		B1000: 700,
	}

	i := new(Item)
	Items = make([]*Item, 100)
	var err error
	Items, err = i.GetAll()
	if err != nil {
		log.Println(err.Error())
	}

}
