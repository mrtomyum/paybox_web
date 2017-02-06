package test

import (
	"testing"
	"github.com/mrtomyum/paybox_terminal/model"
)

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

	err := model.P.PrintTest(data)
	if err != nil {
		t.Log(err.Error())
	}

}
