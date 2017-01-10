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
	Id     string
	Status string
	Send   chan *Message
}

// Todo: Try to Construct CoinHopper Object after DevClient opened connection????
func (ch *CoinHopper) Setup() {
	CH.GetId()
}

//ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins Hopper
func (ch *CoinHopper) GetId() {
	fmt.Println("CoinHopper.GetId() start")
	m := &Message{
		Device:  "coin_hopper",
		Type:    "request",
		Command: "machine_id",
	}
	H.Dev.Send <- m

	// Todo: Ws.ReadJSON() ไม่สามารถดัก Socket "response" ได้

	go func() {
		for {
			select {
			case m := <-ch.Send:
				fmt.Println("Get Response from CoinHopper:", m)
				ch.Id = m.Data.(string)
				break
			}
		}
	}()
	fmt.Println("CoinHopper ID:", ch.Id, "Status:", ch.Status)
}

func (ch *CoinHopper) Event(c *Client) {
	switch c.Msg.Command {
	case "status":         // ร้องขอสถานะต่างๆของอุปกรณ์
	case "cash_amount":    // ร้องขอจานวนเงินคงเหลือใน Coins Hopper
	case "coin_count":     // ร้องขอจานวนเงินเหรียญคงเหลือใน Coins Hopper
	case "set_coin_count": // ตั้งค่าจำนวนเงินคงเหลือใน Coins Hopper
	case "payout_by_cash": // ร้องขอการจ่ายเหรียญออกทางด้านหน้าเครื่องโดยระบุจานวนเป็นยอดเงิน
	case "payout_by_coin": // ร้องขอการจ่ายเหรียญออกทางด้านหน้าเครื่องโดยระบุจานวนเป็นจานวนเหรียญ
	case "empty":          // ร้องขอการปล่อยเหรียญทั้งหมดออกทางด้านล่าง
	case "reset":          // ร้องขอการ Reset ตัวเครื่อง เพ่ือเคลียร์ค่า Error ต่างๆ
	case "status_change": // Event น้ีจะเกิดข้ึนเม่ือสถานะใดๆของ Coins Hopper มีการเปลี่ยนแปลง
		ch.StatusChange(c)
	}
}

func (ch *CoinHopper) PayoutByCash(v int) error {
	// command to send to devClient for "payout" value = v
	fmt.Println("CoinHopper Command=>Payout, Value:", v)
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
