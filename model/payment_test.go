package model_test

import (
	"testing"
	"github.com/mrtomyum/paybox_web/model"
)

func TestPayment_S90PayB100ChangeMustBe10(t *testing.T) {
	s := &model.Sale{
		Total: 90,
	}
	model.PM.New(s)
	mockHW.BillPay = "B100"
	// Expected change = 10
	if model.PM.Change() != 10.0 {
		t.Errorf("Fail: Sale %v Pay 100 expected change = 10 but got = %v", s.Total, model.PM.Change())
	}
	t.Logf("PASS: Sale %v Pay 100 got change = %v", s.Total, model.PM.Change())
}


