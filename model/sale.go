package model

import (
	"fmt"
	"time"
)

type Sale struct {
	Id       int64
	Created  *time.Time
	HostId   string `json:"host_id" db:"host_id"`
	Total    float64 `json:"total"`
	Payment  float64 `json:"payment"`
	Change   float64 `json:"change"`
	Type     string `json:"type" db:"type"`
	IsPosted bool `json:"is_posted" db:"is_posted"`
	SaleSubs []*SaleSub `json:"sale_subs" `
}

type SaleSub struct {
	Line     uint64 `json:"line"`
	SaleId   uint64 `json:"sale_id" db:"sale_id"`
	ItemId   uint64  `json:"item_id" db:"item_id"`
	ItemName string  `json:"item_name" db:"item_name"`
	PriceId  int     `json:"price_id" db:"price_id"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
	Unit     string `json:"unit"`
}

func (s *Sale) Post() error {
	// Ping Server api.paybox.work:8080/ping
	// if post Error s.IsPosted = false
	// IsNetOnline => Post Order ขึ้น Cloud
	s.IsPosted = true
	return nil
}

func (s *Sale) Save() error {
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

	// Todo: Add time to "created" field
	//created := time.Now()
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
	fmt.Println("s.Id =", s.Id)
	ss := SaleSub{}
	sql2 := `INSERT INTO sale_sub(
		sale_id,
		item_id,
		qty,
		price_id,
		price
		)
	VALUES(?,?,?,?,?)`
	// Todo: Loop til end SaleSub
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
	// Check result
	sales := []*Sale{}
	err = db.Select(&sales, "SELECT * FROM sale WHERE id = ?", s.Id)
	if err != nil {
		fmt.Printf("Error when db.Get(&s) %v", err.Error())
		return err
	}
	for _, v := range sales {
		fmt.Println("Read database row->", v)
	}
	fmt.Println("*Sale.Save() completed, data->", sales)
	return nil
}
