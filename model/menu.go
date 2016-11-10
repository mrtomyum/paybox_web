package model

import "github.com/jmoiron/sqlx"

type Menu struct {
	Id      int
	Name    string `json:"name" db:"name"`
	NameEn  string `json:"name_en" db:"name_en"`
	NameCn  string `json:"name_cn" db:"name_cn"`
	Short   string `json:"short" db:"short"`
	ShortEn string `json:"short_en" db:"short_en"`
	ShortCn string `json:"short_cn" db:"short_cn"`
	Image   string `json:"image" db:"image"`
	Link    string `json:"link" db:"link"`
}

func (m *Menu) Index(db *sqlx.DB) (menus []*Menu, err error) {
	sql := `SELECT * FROM menu`
	err = db.Select(&menus, sql)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
