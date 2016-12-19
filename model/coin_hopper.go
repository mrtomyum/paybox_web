package model

import "log"

type CoinHopper struct {
	Msg
}

func (h *CoinHopper) Action(Msg) {
	switch h.Payload.Type {
	case "request": // Msg from web client.
		h.OnRequest()
	case "response": // Response from Device
		h.OnResponse()
	case "event":
		h.OnEvent()
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
	case "hopper_empty": // Legacy
	case "memory_error":
	case "sensors_not_initialised":
	case "lid_remove": // Legacy
	}
}

func (ch *CoinHopper) Serial() (serial string) {

	return serial
}

func (ch *CoinHopper) CashAmount() (amount int64) {

	return amount
}

