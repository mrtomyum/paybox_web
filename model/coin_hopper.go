package model

import "log"

type CoinHopper struct {
	Msg
	Status string
}

func (h *CoinHopper) Action(d Device, m Msg) {
	switch h.Payload.Type {
	case "request": // Send from web client.
		h.OnRequest(d, m)
	case "response": // Response from Device
		h.OnResponse(d, m)
	case "event":
		h.OnEvent(d, m)
	}
}

func (h *CoinHopper) OnRequest(d Device, m Msg) {
	switch m.Payload.Command {
	case "status":
		m.Payload.Data = h.Status
		d.Send <- m
	}
}

func (ch *CoinHopper) OnResponse(d Device, m Msg) {

}

func (ch *CoinHopper) OnEvent(d Device, m Msg) {
	// Sent data string to web socket client
	status := ch.Payload.Command
	data := ch.Payload.Data
	if status != "status_changed" {
		log.Println("Coin Hopper send unknown status:", status)
	}

	switch data {
	case "ready": // do nothing
	case "disable":
	case "calibration_fault":
	case "no_key_set":
	case "coin_jammed":
	// Send msg to web client...How?
	//c.Send <- m
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

