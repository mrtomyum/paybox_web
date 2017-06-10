package model_test

import (
	"testing"
	"github.com/mrtomyum/paybox_web/model"
)

func TestCoinHopper_CoinCount(t *testing.T) {
	// Arrange
	var c = []struct {
		in     int
		expect int
	}{
		{10, 10},
		{10, 10},
		{10, 10},
		{10, 10},
	}
	model.CH.Reset()

	model.CH.SetCoinCount(c[0].in, c[1].in, c[2].in, c[3].in)
	// Act
	err := model.CH.CoinCount()
	if err != nil {
		t.Error(err)
	}

	// Assert
	if model.CH.c1 != c[1].expect ||
		model.CH.c2 != c[2].expect ||
		model.CH.c5 != c[3].expect ||
		model.CH.c5 != c[4].expect {
		t.Fail()
		t.Logf("Expected C1 = %v , but got: %v\n", c[1].expect, model.CH.c1)
	}
}
