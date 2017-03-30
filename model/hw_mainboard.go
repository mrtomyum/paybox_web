package model

import (
	"errors"
)

type MainBoard struct {
	machineId string `json:"machine_id"`
	Status    string
	Send      chan *Message
	PinOpen   int
}

func (mb *MainBoard) event(c *Socket) {
	switch c.Msg.Command {
	case "machine_id":    // ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Main Board
		mb.Send <- c.Msg
	case "set_ex_output": // สั่งงาน External Output ของ Main board
		mb.Send <- c.Msg
	case "get_ex_output": // ใช้สาหรับอ่านค่า External Input ของ Main board
		mb.Send <- c.Msg
	case "get_3g_status": // Event แจ้งสถานะของ Network
		mb.Send <- c.Msg
	}
}

// IsOpen() เช็คสถานะ Magnetic Sensor ว่าฝาตู้เปิดหรือไม่ ให้วน loop ตรวจสอบเซนเซอร์ไว้ทุก 1 วินาที
func (mb *MainBoard) IsOpen() (bool, error) {
	m := &Message{
		Device:  "mainboard",
		Type:    "request",
		Command: "get_ex_output",
		Data:    mb.PinOpen,
	}
	H.Hw.Send <- m
	m = <-mb.Send
	if !m.Result {
		return false, errors.New("Error get 3g status")
	}
	switch m.Data {
	case false:
		return false, nil
	case true:
		return true, nil
	default:
		return false, errors.New("Abnormal Message.Data")
	}
}

// IsOnline() ตรวจเช็คสถานะ Internet และ Server Endpoint ผ่าน Hardware 3G Module
func (mb *MainBoard) IsOnline() bool {
	m := &Message{
		Device:  "mainboard",
		Type:    "request",
		Command: "get_3g_status",
	}
	H.Hw.Send <- m
	m = <-mb.Send
	if !m.Result {
		return false
	}

	// Todo: Check Server Endpoint response this request.
	var response bool
	// Try send REST req.
	// if no response for within timeout sec. then return false.

	switch {
	case m.Data == "offline":
		return false
	case m.Data == "online" && !response:
		return false
	}
	return true
}

func (mb *MainBoard) GetMachineId() string {
	mb.machineId = "xxx" // todo: implement this.
	return ""
}