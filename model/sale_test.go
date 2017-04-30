package model_test

import (
	"github.com/mrtomyum/paybox_web/model"
	"log"
	"testing"
)

var sale model.Sale = model.Sale{
	HostId:   "123",
	Total:    40,
	Pay:      50,
	Change:   10,
	Type:     "test",
	SalePay:  &model.SalePay{C10: 1, B20: 2},
	SaleSubs: saleSubs,
}

var (
	ss1      model.SaleSub = model.SaleSub{SaleId: 1, Line: 1, ItemId: 1, ItemName: "Hello", PriceId: 1, PriceName: "small", Price: 50, Qty: 1, Unit: "cup"}
	ss2      model.SaleSub = model.SaleSub{SaleId: 1, Line: 2}
	saleSubs               = []*model.SaleSub{&ss1, &ss2}
)

func TestSale_Post(t *testing.T) {
	err := sale.Post()
	if err != nil {
		log.Println("Error sale.Post():", err)
	}

}
