package model

import (
	"log"
	"fmt"
)

type Item struct {
	Id      int
	Name    string  `json:"name" db:"name"`
	NameEn  string  `json:"name_en,omitempty" db:"name_en"`
	NameCn  string  `json:"name_cn,omitempty" db:"name_cn"`
	Unit    string  `json:"unit"`
	UnitEn  string  `json:"unit_en,omitempty" db:"unit_en"`
	UnitCn  string  `json:"unit_cn,omitempty" db:"unit_cn"`
	MenuId  uint64  `json:"menu_id,omitempty" db:"menu_id"`
	MenuSeq int     `json:"menu_seq,omitempty" db:"menu_seq"`
	Image   string  `json:"image" db:"image"`
	Prices  []*Price `json:"prices"`
}

type Price struct {
	Id     int     `json:"id"`
	ItemId int64   `json:"-" db:"item_id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
}

func (i *Item) Get(id int64) (err error) {
	sql := `SELECT * FROM item WHERE id = ?`
	err = db.Get(i, sql, id)
	//err = db.QueryRowx(sql, id).StructScan(i)
	if err != nil {
		return err
	}
	// ดึงข้อมูลราคาทั้งหมดของสินค้ารายการนี้
	sizes := []*Price{}
	sql = `SELECT * FROM price WHERE item_id = ?`
	err = db.Select(&sizes, sql, id)
	if err != nil {
		return err
	}
	i.Prices = sizes
	return nil
}

func (i *Item) ByMenuId(id int64) ([]*Lang, error) {
	fmt.Println("call method: Item.ByMenuId::lang:", langs)
	var sql string
	langInit()
	for _, l := range langs {
		items := []*Item{}
		switch l.Id {
		case 1:
			sql = `SELECT id, name, unit, menu_seq, image FROM item WHERE menu_id = ?`
		case 2:
			sql = `SELECT id, name_en as name, unit_en as unit, menu_seq, image FROM item WHERE menu_id = ?`
		case 3:
			sql = `SELECT id, name_cn as name, unit_cn as unit, menu_seq, image FROM item WHERE menu_id = ?`
		}
		fmt.Println("case:", l.Id, l.Name)
		err := db.Select(&items, sql, id)
		if err != nil {
			log.Println("Error select items")
			return nil, err
		}
		fmt.Println("items:", items)
		// query Size{}
		for _, i := range items {
			prices := []*Price{}
			sql = `SELECT * FROM price WHERE item_id = ?`
			item_id := int(i.Id)
			err = db.Select(&prices, sql, item_id)
			if err != nil {
				return nil, err
			}
			i.Prices = prices
		}
		l.MenuId = int(id)
		l.Items = items
	}
	return langs, nil
}
