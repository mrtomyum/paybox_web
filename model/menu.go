package model

import "github.com/jmoiron/sqlx"

type Menu struct {
	Id int
	NameTh  string `json:"name_th" db:"name_th"`
	NameEn  string `json:"name_en" db:"name_en"`
	NameCN  string `json:"name_cn" db:"name_cn"`
	ShortTh string `json:"short_th" db:"short_th"`
	ShortEn string `json:"short_en" db:"short_en"`
	ShortCn string `json:"short_cn" db:"short_cn"`
}

func (m *Menu) Index(db *sqlx.DB) (menus []*Menu, err error) {
	sql := `SELECT * FROM menu`
	err = db.Select(&menus, sql)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
