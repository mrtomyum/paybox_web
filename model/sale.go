package model

import (
	"fmt"
	"time"
)

// Sale เป็นหัวเอกสารขายแต่ละครั้ง
type Sale struct {
	Id       int64
	Created  *time.Time
	HostId   string `json:"host_id" db:"host_id"`
	Total    float64 `json:"total"`
	Payment  float64 `json:"payment"`
	Change   float64 `json:"change"`
	Type     string `json:"type" db:"type"`
	IsPosted bool `json:"is_posted" db:"is_posted"`
	SaleSubs []*SaleSub `json:"sale_subs"`
}

// SaleSub เป็นรายการสินค้าที่ขายใน Sale
type SaleSub struct {
	SaleId   uint64 `json:"sale_id" db:"sale_id"`
	Line     uint64 `json:"line"`
	ItemId   uint64  `json:"item_id" db:"item_id"`
	ItemName string  `json:"item_name" db:"item_name"`
	PriceId  int     `json:"price_id" db:"price_id"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
	Unit     string `json:"unit"`
}

// Payment เก็บรายละเอียดการชำระเงิน เหรียญ ธนบัตร หรือในอนาคตจะเพิ่มบัตรเครดิต และ Cashless Payment ได้ด้วย
type SalePay struct {
	SaleId  int64
	THB20   int // จำนวนธนบัตรใบละ 20 บาท
	THB50   int // จำนวนธนบัตรใบละ 50 บาท
	THB100  int // จำนวนธนบัตรใบละ 100 บาท
	THB500  int // จำนวนธนบัตรใบละ 500 บาท
	THB1000 int // จำนวนธนบัตรใบละ 1000 บาท
	THB1C   int // จำนวนเหรียญ 1 บาท
	THB2C   int // จำนวนเหรียญ 2 บาท
	THB5C   int // จำนวนเหรียญ 5 บาท
	THB10C  int // จำนวนเหรียญ 10 บาท
}

func (s *Sale) Post() error {
	fmt.Println("method *Sale.Post()")
	// Ping Server api.paybox.work:8080/ping
	url := "http://paybox.work/api/v1/vending/sell"
	fmt.Println("URL:>", url)

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	// if post Error s.IsPosted = false
	// IsNetOnline => Post Order ขึ้น Cloud
	s.IsPosted = true
	return nil
}

func (s *Sale) Save() error {
	s.Payment = PM.Total
	s.Change = s.Payment - s.Total
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
		s.Payment,
		s.Change,
		s.Type,
		s.IsPosted,
	)
	if err != nil {
		fmt.Printf("Error when db.Exec(sql1) %v", err.Error())
		return err
	}
	s.Id, _ = rs.LastInsertId()
	fmt.Println("s.machineId =", s.Id)

	sql2 := `INSERT INTO sale_sub(
		sale_id,
		item_id,
		item_name,
		price_id,
		price
		qty,
		unit
		)
	VALUES(?,?,?,?,?)`
	for _, ss := range s.SaleSubs {
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
			fmt.Printf("Error when db.Exec(sql2) %v", err.Error())
			return err
		}

	}

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
