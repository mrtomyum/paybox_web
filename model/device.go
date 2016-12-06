package model

import "log"

type Money struct {
	Job    string
	Amount int
}

type OnHand struct {
	OnhandAmount int
	Job          string
}

type Payload struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Result  bool   `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Msg struct {
	Device  string  `json:"device"`
	Payload Payload `json:"payload"`
}

type Devicer interface {
	Serial() string
	Status() string
	CashAmount() float64
}

type CoinHopper struct {
	Msg
}

func (ch *CoinHopper) CheckMsg() {
	switch ch.Payload.Type {
	case "request": // Msg from web client.
		ch.OnRequest()
	case "response": // Response from Device
		ch.OnResponse()
	case "event":
		ch.OnEvent()
	}
}

func (ch *CoinHopper) OnRequest() {

}

func (ch *CoinHopper) OnResponse() {

}

func (ch *CoinHopper) OnEvent() {
	// Sent data string to web socket client
	status := ch.Payload.Command
	data := ch.Payload.Data
	if status != "status_changed" {
		log.Println("Coin Hopper send unknown status:", status)
	}
	switch data {
	case "ready":
	case "disable":
	case "calibration_fault":
	case "no_key_set":
	case "coin_jammed":
	case "fraud":
	case "hopper_empty":// Legacy
	case "memory_error":
	case "sensors_not_initialised":
	case "lid_remove": // Legacy
	}
}

func (ch *CoinHopper) Serial() (serial string) {

	return serial
}

func (ch *CoinHopper) CashAmount() (amount float64) {

	return amount
}

type CoinAcceptor struct {
	Msg
}
type BillAcceptor struct {
	Msg
}

func (b BillAcceptor) Status() string {
	var status string
	return status
}
