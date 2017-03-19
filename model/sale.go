package model

import (
	"fmt"
	"time"
	"net/http"
	"errors"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"github.com/jmoiron/sqlx"
)

// Sale เป็นหัวเอกสารขายแต่ละครั้ง
type Sale struct {
	Id       uint64
	Created  *time.Time
	HostId   string `json:"host_id" db:"host_id"`
	Total    float64 `json:"total"`
	Pay      float64 `json:"payment"`
	Change   float64 `json:"change"`
	Type     string `json:"type" db:"type"`
	IsPosted bool `json:"is_posted" db:"is_posted"`
	SaleSubs []*SaleSub `json:"sale_subs"`
	SalePay  SalePay
}

// SaleSub เป็นรายการสินค้าที่ขายใน Sale
type SaleSub struct {
	SaleId    uint64 `json:"sale_id" db:"sale_id"`
	Line      uint64 `json:"line"`
	ItemId    uint64  `json:"item_id" db:"item_id"`
	ItemName  string  `json:"item_name" db:"item_name"`
	PriceId   int `json:"price_id" db:"price_id"`
	PriceName string `json:"price_name" db:"price_name"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	Unit      string `json:"unit"`
}

// Payment เก็บรายละเอียดการชำระเงิน เหรียญ ธนบัตร หรือในอนาคตจะเพิ่มบัตรเครดิต และ Cashless Payment ได้ด้วย
type SalePay struct {
	SaleId  int64
	TH025C  int `json:"th025c,omitempty"`  // จำนวนเหรียญ 25 สตางค์
	TH050C  int `json:"th050c,omitempty"`  // จำนวนเหรียญ 50 สตางค์
	TH1C    int `json:"th1c,omitempty"`    // จำนวนเหรียญ 1 บาท
	TH2C    int `json:"th2c,omitempty"`    // จำนวนเหรียญ 2 บาท
	TH5C    int `json:"th5c,omitempty"`    // จำนวนเหรียญ 5 บาท
	TH10C   int `json:"th10c,omitempty"`   // จำนวนเหรียญ 10 บาท
	TH20B   int `json:"th20b,omitempty"`   // จำนวนธนบัตรใบละ 20 บาท
	TH50B   int `json:"th50b,omitempty"`   // จำนวนธนบัตรใบละ 50 บาท
	TH100B  int `json:"th100b,omitempty"`  // จำนวนธนบัตรใบละ 100 บาท
	TH500B  int `json:"th500b,omitempty"`  // จำนวนธนบัตรใบละ 500 บาท
	TH1000B int `json:"th1000b,omitempty"` // จำนวนธนบัตรใบละ 1000 บาท
}

// *ไม่น่าได้ใช้* GetSaleSubFK เพราะ WebUI อาจส่งชื่อสินค้า/ราคามาผิด เลยตรวจใหม่ซ้ำ
func (s *Sale) GetSaleSubFK(db *sqlx.DB) error {
	sql1 := `SELECT name FROM item WHERE id = ? LIMIT 1`
	sql2 := `SELECT name FROM price WHERE id = ? LIMIT 1`
	var name string
	ss := s.SaleSubs
	for _, sub := range ss {
		err := db.Get(&name, sql1, sub.ItemId)
		if err != nil {
			return err
		}
		sub.ItemName = name
		err = db.Get(&name, sql2, sub.PriceId)
		if err != nil {
			return err
		}
		sub.PriceName = name
	}
	return nil
}

func (s *Sale) Post() error {
	fmt.Println("method *Sale.Post()")
	// เช็คสถานะ Network และ Server ว่า IsNetOnline อยู่หรือไม่?
	isOnline, err := MB.IsOnline()
	if err != nil {
		return err
	}
	if isOnline {
		fmt.Println("Offline => Save sale to disk")
		return errors.New("Offline => Save sale to disk and try again.")
	}

	// Ping Server api.paybox.work:8080/ping
	url := "http://paybox.work/api/v1/vending/sell"
	fmt.Println("URL:>", url)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(s)
	//req, err := http.NewRequest("POST", url, b)
	//req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")

	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	return err
	//}
	resp, _ := http.Post(url, "application/json; charset=utf-8", b)
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if string(body) == "post error" { // todo: น่าจะใช้ JSON response ไหม?
		s.IsPosted = false
	}
	// IsNetOnline => Post Order ขึ้น Cloud
	s.IsPosted = true
	fmt.Println("Post ยอดขายขึ้น Cloud -> sale.Post()")
	return nil
}

func (s *Sale) Save() error {
	s.Pay = PM.total
	s.Change = s.Pay - s.Total
	fmt.Println("*Sale.Save() start")
	sql1 := `INSERT INTO sale(
		host_id,
		total,
		payment,
		change,
		type,
		is_posted
		)
	VALUES (?,?,?,?,?,?)`
	rs, err := db.Exec(sql1,
		s.HostId,
		s.Total,
		s.Pay,
		s.Change,
		s.Type,
		s.IsPosted,
	)
	if err != nil {
		fmt.Printf("Error when db.Exec(sql1) %v", err.Error())
		return err
	}
	id, _ := rs.LastInsertId()
	s.Id = uint64(id)
	fmt.Println("s.machineId =", s.Id)

	sql2 := `INSERT INTO sale_sub(
		sale_id,
		line,
		item_id,
		item_name,
		price_id,
		price,
		qty,
		unit
		)
	VALUES(?,?,?,?,?,?,?,?)`
	for _, ss := range s.SaleSubs {
		fmt.Println("start for range s.SaleSubs")
		rs, err = db.Exec(sql2,
			s.Id,
			ss.Line,
			ss.ItemId,
			ss.ItemName,
			ss.PriceId,
			ss.Price,
			ss.Qty,
			ss.Unit,
		)
		if err != nil {
			fmt.Printf("Error when db.Exec(sql2) %v\n", err.Error())
			return err
		}
		fmt.Println("Insert sale_sub line ", ss)
	}
	fmt.Println("Save data sucess: sale =", s)

	// Check result
	//sales := []*Sale{}
	//err = db.Select(&sales, "SELECT * FROM sale WHERE id = ?", s.Id)
	//if err != nil {
	//	fmt.Printf("Error when db.Get(&s) %v", err.Error())
	//	return err
	//}
	//for _, v := range sales {
	//	fmt.Println("Read database row->", v)
	//}
	//fmt.Println("*Sale.Save() completed, data->", sales)
	return nil
}
