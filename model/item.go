package model

import (
	sys "github.com/mrtomyum/sys/model"
	"log"
)

type Item struct {
	sys.Base
	Name    string  `json:"name" db:"name"`
	NameEn  string  `json:"name_en,omitempty" db:"name_en"`
	NameCn  string  `json:"name_cn,omitempty" db:"name_cn"`
	Unit    string  `json:"unit"`
	UnitEn  string  `json:"unit_en,omitempty"`
	UnitCn  string  `json:"unit_cn,omitempty"`
	//Price   float32 `json:"price"`
	//PriceS  float32 `json:"price_s" db:"price_s"`
	//PriceM  float32 `json:"price_m" db:"price_m"`
	//PriceL  float32 `json:"price_l" db:"price_l"`
	MenuId  uint64  `json:"menu_id,omitempty" db:"menu_id"`
	MenuSeq int     `json:"menu_seq,omitempty" db:"menu_seq"`
	Image   string  `json:"image" db:"image"`
	Sizes   []*Size `json:"sizes"`
}

type Size struct {
	ItemId int64   `json:"-" db:"item_id"`
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
}

func (i *Item) Get(id int64) (err error) {
	sql := `SELECT * FROM item WHERE id = ?`
	err = db.Get(&i, sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) ByMenuId(id int64) ([]*Lang, error) {
	log.Println("call method: Item.ByMenuId::lang:", langs)
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
		log.Println("case:", l.Id, l.Name)
		err := db.Select(&items, sql, id)
		if err != nil {
			log.Println("Error select items")
			return nil, err
		}
		log.Println("items:", items)
		// query Size{}
		for _, i := range items {
			sizes := []*Size{}
			sql = `SELECT * FROM size WHERE item_id = ?`
			item_id := int(i.Id)
			err = db.Select(&sizes, sql, item_id)
			if err != nil {
				return nil, err
			}
			i.Sizes = sizes
		}
		l.MenuId = int(id)
		l.Items = items
	}
	return langs, nil
}
