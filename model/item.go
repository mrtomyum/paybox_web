package model

import (
	sys "github.com/mrtomyum/sys/model"
	"github.com/jmoiron/sqlx"
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
	NameTh string `json:"name_th" db:"name_th"`
	NameEn string `json:"name_en" db:"name_en"`
	NameCn string `json:"name_cn" db:"name_cn"`
	Unit   string
	Price  float32
	PriceS float32
	PriceM float32
	PriceL float32
	MenuId uint64
	MenuSeq int
}

func (i *Item) GetIndex(db *sqlx.DB) ([]*Item, error) {
	var items []*Item
	sql := `SELECT * FROM item WHERE deleted IS NOT null`
	err := db.Select(&items, sql)
	if err != nil {
		return nil, err
	}
	return items, nil
}