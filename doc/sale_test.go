package test

import (
	"testing"
	"github.com/mrtomyum/paybox_terminal/model"
	"github.com/matryer/silk/runner"
)

func TestPrinter_makeSaleSlip(t *testing.T) {
	ss1 := model.SaleSub{
		SaleId:  1,
		Line:    1,
		ItemId:  3,
		ItemName:"Cappuchino Ice Latte",
		PriceId: 1,
		Price:   50,
		Qty:     1,
		Unit:    "แก้ว",
	}
	ss2 := model.SaleSub{
		SaleId:  1,
		Line:    2,
		ItemId:  4,
		ItemName:"Cappuchino Ice Freppe",
		PriceId: 1,
		Price:   70,
		Qty:     1,
		Unit:    "แก้ว",
	}
	s := model.Sale{
		Id:       1,
		Total:    120,
		Pay:      150,
		Change:   30,
		Type:     "take_home",
		SaleSubs: [{&ss1}, {&ss2}],
	}
	s.SaleSubs = append(s.SaleSubs, &ss1)
	s.SaleSubs = append(s.SaleSubs, &ss2)

	model.P.Print(&s)

}

func TestSale_save(t *testing.T) {
	ss := model.SaleSub{}
	sp := model.SalePay{}
	s := model.Sale{}

}