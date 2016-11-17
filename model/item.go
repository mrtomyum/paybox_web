package model

import (
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

func (i *Item) GetIndex(menuId int64) ([]*Item, error) {
	var items []*Item
	sql := `SELECT * FROM item WHERE menu_id = ?`
	err := db.Select(&items, sql, menuId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *Item) FindById(id int64) (items []*Item, err error) {
	sql := `SELECT * FROM item WHERE menu_id = ?`
	err = db.Select(&items, sql, id)
	if err != nil {
		return nil, err
	}
	return items, nil
}