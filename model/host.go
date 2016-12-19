package model

import (
	"log"
	"fmt"
)

// ====================
// Host Machine
// ====================
// Host is a Machine Property
type Host struct {
	Id            string
	OnHand        int  // Money onhand
	Online        bool // Network GSM status
	Devices       map[string]bool
	DeviceOnline  chan *Actioner
	DeviceOffline chan *Actioner
}

func (h *Host) GetOnHand(c *Client, msg Msg) {
	fmt.Println("onhand_request_starting....")
	log.Println("hub.Clients:", MyHub.Clients)
	msg.Payload.Type = "response"
	msg.Payload.Result = true
	//msg.Payload.Data = 100 // test dummy data
	msg.Payload.Data = h.OnHand
	c.Send <- msg
}

func (h *Host) GetDevices() {
	log.Println(h.Devices)
}

func (h *Host) Cancel(c *Client, msg Msg) error {
	// TODO: Check Money.Onhand
	// TODO: Calculate refund
	// if error set Payload.Result = fault
	// Reset Onhand Amount เป็น 0 และคืนเงินลูกค้า
	fmt.Println("cancel_request_starting....")
	h.OnHand = 0
	msg.Payload.Type = "response"
	msg.Payload.Result = true
	msg.Payload.Data = "Cancel - Successful"
	c.Send <- msg
	return nil
}

func (h *Host) Billing(c *Client, msg Msg) error {
	// Get Escrow status from  Bill acceptor if error send message to client
	// Check Cash onhand is enough?
	// Print Order/Queue/Invoice and return if error
	fmt.Println("billing_request_starting....")
	msg.Payload.Command = "billing"
	msg.Payload.Data = "Docno : xxxxxx sucessful"
	msg.Payload.Result = true
	msg.Device = "Host"
	msg.Payload.Type = "response"
	c.Send <- msg
	// todo: save into database sqlite
	// todo: reset Onhand
	h.OnHand = 0
	h.GetOnHand(c, msg)
	return nil
}