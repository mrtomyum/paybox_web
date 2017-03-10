package model_test

import (
	"testing"
	"github.com/mrtomyum/paybox_web/model"
)

func TestCoinHopper_CoinCount(t *testing.T) {
	// Arrange
	model.CH.Reset()
	inC1 := 10
	inC2 := 10
	inC5 := 10
	inC10 := 10
	model.CH.SetCoinCount(inC1, inC2, inC5, inC10)
	// Act
	err := model.CH.CoinCount()
	if err != nil {
		t.Error(err)
	}

	// Assert
	outC1 := model.CH.QtyC1()
	outC2 := model.CH.QtyC2()
	outC5 := model.CH.QtyC5()
	outC10 := model.CH.QtyC10()
	if outC1 != 10 || outC2 != 10 || outC5 != 10 || outC10 != 10 {
		t.Fail()
		t.Logf("Expected C1 = 10 Got model.Ch.QtyC1() = ", 10, "But got":, outC1)
	}
}
