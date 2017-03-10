/*
สรุป Message Command ของ Coin Hopper
	case "machine_id":     // ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins Hopper
	case "status":         // ร้องขอสถานะต่างๆของอุปกรณ์
	case "cash_amount":    // ร้องขอจานวนเงินคงเหลือใน Coins Hopper
	case "coin_count":     // ร้องขอจานวนเงินเหรียญคงเหลือใน Coins Hopper
	case "set_coin_count": // ตั้งค่าจำนวนเงินคงเหลือใน Coins Hopper
	case "payout_by_cash": // ร้องขอการจ่ายเหรียญออกทางด้านหน้าเครื่องโดยระบุจานวนเป็นยอดเงิน
	case "payout_by_coin": // ร้องขอการจ่ายเหรียญออกทางด้านหน้าเครื่องโดยระบุจานวนเป็นจานวนเหรียญ
	case "empty":          // ร้องขอการปล่อยเหรียญทั้งหมดออกทางด้านล่าง
	case "reset":          // ร้องขอการ Reset ตัวเครื่อง เพ่ือเคลียร์ค่า Error ต่างๆ
	case "status_change":  // Event น้ีจะเกิดข้ึนเม่ือสถานะใดๆของ Coins Hopper มีการเปลี่ยนแปลง
*/
package model

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"errors"
)

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
	machineId string `json:"machine_id"`
	Status    string
	Response  chan *Message
	C1        int
	C2        int
	C5        int
	C10       int
}

// Todo: Try to Construct CoinHopper Object after Hw Client opened connection????
func (ch *CoinHopper) Setup() {
	CH.GetId()
}

//ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins hopper
func (ch *CoinHopper) GetId() {
	fmt.Println("CoinHopper.GetId() start")
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "machine_id",
	}
	H.Hw.Send <- m

	// เปิด Goroutine เพื่อรอรับ MessagMessagee กลับมาจาก Channel ch.Response
	go func() {
		m := <-ch.Response
		ch.machineId = m.Data.(string)
		fmt.Println("Got Response from CoinHopper ID:", ch.machineId, "Status:", ch.Status)
	}()
}

func (ch *CoinHopper) Event(c *Client) {
	switch c.Msg.Command {
	//case "status":         // ร้องขอสถานะต่างๆของอุปกรณ์
	//case "cash_amount":    // ร้องขอจานวนเงินคงเหลือใน Coins hopper
	//case "coin_count":     // ร้องขอจานวนเงินเหรียญคงเหลือใน Coins hopper
	//case "set_coin_count": // ตั้งค่าจำนวนเงินคงเหลือใน Coins hopper
	//case "payout_by_cash": // ร้องขอการจ่ายเหรียญออกทางด้านหน้าเครื่องโดยระบุจานวนเป็นยอดเงิน
	//case "payout_by_coin": // ร้องขอการจ่ายเหรียญออกทางด้านหน้าเครื่องโดยระบุจานวนเป็นจานวนเหรียญ
	//case "empty":          // ร้องขอการปล่อยเหรียญทั้งหมดออกทางด้านล่าง
	//case "reset":          // ร้องขอการ Reset ตัวเครื่อง เพ่ือเคลียร์ค่า Error ต่างๆ
	case "status", "cash_amount", "coin_count", "set_coin_count", "paybout_by_cash", "payout_by_coin", "empty", "reset":
		ch.Response <- c.Msg
		log.Println("ch.Response <-c.Msg", c.Msg)
	case "status_change": // Event น้ีจะเกิดข้ึนเม่ือสถานะใดๆของ Coins hopper มีการเปลี่ยนแปลง
		ch.StatusChange(c)
	}
}

func (ch *CoinHopper) PayoutByCash(v float64) error {
	// command to send to devClient for "payout" value = v
	fmt.Println("====Send Command to CoinHopper payout_by_cash, Value:", v, "====")
	defer fmt.Println("============================================================")
	waitChannel := make(chan *Message)
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "payout_by_cash",
		Data:    v,
	}
	H.Hw.Send <- m
	fmt.Println("Waiting response from coin hopper.")
	m = <-ch.Response
	log.Println("Got response from coin hopper:", m)
	close(waitChannel)
	// todo: ให้ตรวจ result == false  และ return error ด้วย เช่นกรณีเหรียญหมด
	return nil
}

// StatusChange ตอบสนองต่อ Event ที่ถูกส่งมาจาก CoinHopper
// Message เฉพาะบางรายการที่จำเป็นจะส่งแจ้งเตือนให้ Web นำไปแสดงผลบอก User
// แต่โดยทั่วไปจะต้องส่งขึ้น Cloud ทันทีท้งนี้หากติดต่อไม่ได้ ต้องลง Errorlog เก็บไว้
func (ch *CoinHopper) StatusChange(c *Client) {
	fmt.Println("CoinHopper.StatusChange() start")
	switch c.Msg.Data.(string) {
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
	default:
		log.Println("Error CoinHopper.StatusChange: Unknown Msg.Data=>", c.Msg.Data.(string))
	}

	// Todo: Post to Cloud to Log Status
	// If Net == online Send msg to Cloud
	// If Net != online SQL ErrorLog
	H.Web.Send <- c.Msg // ตอนนี้กำหนดให้ทุกสถานะจะส่งไปให้ Web ด้วย
}

func (ch *CoinHopper) Reset() error {
	return nil
}

// CoinCount() ร้องขอจำนวนเหรียญแต่ละขนาดที่เหลือใน Hopper
func (ch *CoinHopper) CoinCount() error {
	return nil
}

// SetCoinCount() คำสั่ง เพิ่ม/ลด จำนวนเหรียญคงเหลือใน Coins Hopper
// ระวัง เมธอดนี้จะเพิ่ม หรือลด จากค่าเดิมเท่านั้น
func (ch *CoinHopper) SetCoinCount(c1, c2, c5, c10 int) error {
	fmt.Println("===*CoinHopper.SetCoinCount()===START")
	defer fmt.Println("===*CoinHopper.SetCoinCount()===END")

	data := gin.H{
		"coin_1":  c1,
		"coin_2":  c2,
		"coin_5":  c5,
		"coin_10": c10,
	}
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "set_coin_count",
		Data:    data,
	}
	H.Hw.Send <- m
	m = <-ch.Response
	fmt.Println("Response from CoinHopper:", m)
	errSetCoinCount := errors.New("Error set_coin_count to CoinHopper.")
	if !m.Result {
		return errSetCoinCount
	}
	return nil
}

// Empty() ปล่อยเหรียญลง CoinBox ทั้งหมด และน่าจะรีเซ็ท coin_count ด้วย
func (ch *CoinHopper) Empty() error {
	return nil
}

