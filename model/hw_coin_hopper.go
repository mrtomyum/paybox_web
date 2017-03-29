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

//type CoinHopperStatus int
//
//const (
//	DISABLE                 CoinHopperStatus = iota
//	CALIBRATION_FAULT
//	NO_KEY_SET
//	COIN_JAMMED
//	FRAUD
//	HOPPER_EMPTY
//	MEMORY_ERROR
//	SENSORS_NOT_INITIALISED
//	LID_REMOVED
//)

type CoinHopper struct {
	machineId string `json:"machine_id"`
	Status    string
	Send      chan *Message
	C1        int
	C2        int
	C5        int
	C10       int
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
	m = <-ch.Send
	ch.machineId = m.Data.(string)
	fmt.Println("Got Response from CoinHopper ID:", ch.machineId, "Status:", ch.Status)
}

func (ch *CoinHopper) Event(c *Socket) {
	switch c.Msg.Command {
	case "status", "cash_amount", "coin_count", "set_coin_count", "paybout_by_cash", "payout_by_coin", "empty", "reset":
		ch.Send <- c.Msg
		log.Println("ch.Send <-c.Msg", c.Msg)
	case "status_change": // Event น้ีจะเกิดข้ึนเม่ือสถานะใดๆของ Coins hopper มีการเปลี่ยนแปลง
		ch.StatusChange(c)
	}
}

func (ch *CoinHopper) PayoutByCoin(c1, c2, c5, c10 int) error {
	fmt.Println("==== Send Command to CoinHopper payout_by_coin, Value: c1=%v c2=%v c5=%v c10=%v ", c1, c2, c5, c10)
	defer fmt.Println("============================================================")

	data := gin.H{
		"coin_1":  c1,
		"coin_2":  c2,
		"coin_5":  c5,
		"coin_10": c10,
	}
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "payout_by_coin",
		Data:    data,
	}
	H.Hw.Send <- m
	fmt.Println("Waiting response from coin hopper.")
	m = <-ch.Send
	log.Println("Got response from coin hopper:", m)
	// ตรวจ result == false  และ return error ด้วย เช่นกรณีเหรียญหมด
	if !m.Result {
		return errors.New("Error payout from Hopper.")
	}
	//นับมูลค่าเหรียญที่จ่าย
	v2 := 2 * c2
	v5 := 5 * c5
	v10 := 10 * c10
	value := c1 + v2 + v5 + v10
	CB.hopper -= float64(value)
	// todo: อ่านยอดคงเหลือของเหรียญแต่ละขนาด
	return nil
}

// PayoutByCash จ่ายเหรียญออกจาก Hopper โดยระบุยอดเงิน
func (ch *CoinHopper) PayoutByCash(v float64) error {
	// command to send to devClient for "payout" value = v
	fmt.Println("====Send Command to CoinHopper payout_by_cash, Value:", v, "====")
	defer fmt.Println("============================================================")
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "payout_by_cash",
		Data:    v,
	}
	H.Hw.Send <- m
	fmt.Println("Waiting response from coin hopper.")
	m = <-ch.Send
	log.Println("Got response from coin hopper:", m)
	// ตรวจ result == false  และ return error ด้วย เช่นกรณีเหรียญหมด
	if !m.Result {
		return errors.New("Error payout from Hopper.")
	}
	CB.hopper -= m.Data.(float64) //ลดยอดเหรียญใน hopper
	return nil
}

// StatusChange ตอบสนองต่อ Event ที่ถูกส่งมาจาก CoinHopper
// Message เฉพาะบางรายการที่จำเป็นจะส่งแจ้งเตือนให้ Web นำไปแสดงผลบอก User
// แต่โดยทั่วไปจะต้องส่งขึ้น Cloud ทันทีท้งนี้หากติดต่อไม่ได้ ต้องลง Errorlog เก็บไว้
func (ch *CoinHopper) StatusChange(c *Socket) {
	fmt.Println("CoinHopper.StatusChange() start")
	switch c.Msg.Data.(string) {
	case "ready", "disable", "calibration_fault", "no_key_set", "coin_jammed", "fraud", "hopper_empty", "memory_error", "sensors_not_initialised", "lid_remove": // Legacy
		H.Web.Send <- c.Msg                                                                                                                                      // ตอนนี้กำหนดให้ทุกสถานะจะส่งไปให้ Web ด้วย
		ch.Status = c.Msg.Data.(string)
	default:
		log.Println("Error CoinHopper.StatusChange: Unknown Msg.Data=>", c.Msg.Data.(string))
	}
}

// CoinCount() ร้องขอจำนวนเหรียญแต่ละขนาดที่เหลือใน Hopper
func (ch *CoinHopper) CoinCount() error {
	// กำหนดให้ตรวจสอบทุกรอบที่จบการขาย/ทอนเงิน แล้วส่งเป็น WebSocket Event ให้กับ WebUI
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "coin_count",
	}
	H.Hw.Send <- m
	m = <-ch.Send
	if !m.Result {
		return errors.New("Error response from command coin_hopper.coin_count")
	}
	data := m.Data.(map[string]interface{})
	ch.C1 = data["coin_1"].(int)
	ch.C2 = data["coin_2"].(int)
	ch.C5 = data["coin_5"].(int)
	ch.C10 = data["coin_10"].(int)
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
	m = <-ch.Send
	fmt.Println("Response from CoinHopper:", m)
	errSetCoinCount := errors.New("Error set_coin_count to CoinHopper.")
	if !m.Result {
		return errSetCoinCount
	}
	return nil
}

// Empty() ปล่อยเหรียญลง CoinBox ทั้งหมด และน่าจะรีเซ็ท coin_count ด้วย
func (ch *CoinHopper) Empty() error {
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "empty",
	}
	H.Hw.Send <- m
	m = <-ch.Send
	if !m.Result {
		return errors.New("Error from EMPTY coin hopper.")
	}
	fmt.Println("SUCCESS EMPTY coin hopper.")
	return nil
}

// Reset() จะปล่อยเหรียญออก และปรับค่ายอดเหรียญคงเหลือในตัว Hopper = 0
func (ch *CoinHopper) Reset() error {
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "reset",
	}
	H.Hw.Send <- m
	m = <-ch.Send
	if !m.Result {
		return errors.New("Error from RESET coin hopper.")
	}
	fmt.Println("SUCCESS RESET coin hopper.")
	return nil
}
