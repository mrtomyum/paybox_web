package main

import (
	"testing"
	"github.com/mrtomyum/paybox_terminal/model"
	"github.com/mrtomyum/paybox_terminal/controller"
)
//var db *sqlx.DB
//
//func init() {
//	db = sqlx.MustConnect("sqlite3", "./paybox.db")
//}

func TestGetItemIndex(t *testing.T) {
	var i model.Item
	var items []*model.Item
	items, err := i.GetIndex(controller.DB)
	if err != nil {
		t.Fail()
	}
	t.Logf("Get Item success: %v", items)
}

//var i1 = model.Item{
//	Name:"คาปูชิโน่ร้อน,
//	NameEn:"Hot Cappuchino"
//	NameCn:"xxxxx"
//	Unit:"แก้ว",
//	Price:45.00,
//	PriceS:50.00,
//	PriceM:60.00,
//	PriceL:70.00,
//	MenuId:1,
//	MenuSeq:1,
//}