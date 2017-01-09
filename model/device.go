package model

import (

	"fmt"
)

type Devicer interface {
	Status() string
}

type CoinHopperStatus int

const (
	DISABLE                 CoinHopperStatus = iota
	CALIBRATION_FAULT
	NO_KEY_SET
	COIN_JAMMED
	FRAUD
	HOPPER_EMPTY
	MEMORY_ERROR
	SENSORS_NOT_INITIALISED
	LID_REMOVED
)

type CoinHopper struct {
	status string
}

func (ch *CoinHopper) Status() string {
	return ch.status
}

func (ch *CoinHopper) Payout(v int) error {
	// command to send to devClient for "payout" value = v
	fmt.Println("CoinHopper Command=>Payout, Value:", v)
	return nil
}

type BillAcceptor struct {
	status string
}

type CoinAcceptor struct {
	status string
}

type Acceptor interface {
	Serial() string
	Status() string
	CashReceive() int64
}
