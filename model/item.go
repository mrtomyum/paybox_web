package model

import (
	"github.com/jmoiron/sqlx"
	sys "github.com/mrtomyum/sys/model"
)

type ItemSize int

const (
	REGULAR ItemSize = iota
	SMALL
	MEDIUM
	LARGE
)

type Item struct {
	sys.Base
	Name    string `json:"name" db:"name"`
	NameEn  string `json:"name_en" db:"name_en"`
	NameCn  string `json:"name_cn" db:"name_cn"`
	Unit    string
	Price   float32
	PriceS  float32 `json:"price_s" db:"price_s"`
	PriceM  float32 `json:"price_m" db:"price_m"`
	PriceL  float32 `json:"price_l" db:"price_l"`
	MenuId  uint64  `json:"menu_id" db:"menu_id"`
	MenuSeq int     `json:"menu_seq" db:"menu_seq"`
	Image   string  `json:"image" db:"image"`
}

func (i *Item) GetIndex(db *sqlx.DB) ([]*Item, error) {
	var items []*Item
	sql := `SELECT * FROM item`
	err := db.Select(&items, sql)
	if err != nil {
		return nil, err
	}
	return items, nil
}
