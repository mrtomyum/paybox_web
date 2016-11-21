package model

import "log"

type Lang struct {
	Id   int     `json:"lang_id"`
	Name string  `json:"lang_name"`
	Menu []*Menu `json:"menu"`
}

type Menu struct {
	Id      int
	Name    string `json:"name" db:"name"`
	NameEn  string `json:"name_en,omitempty" db:"name_en"`
	NameCn  string `json:"name_cn,omitempty" db:"name_cn"`
	Short   string `json:"short,omitempty" db:"short"`
	ShortEn string `json:"short_en,omitempty" db:"short_en"`
	ShortCn string `json:"short_cn,omitempty" db:"short_cn"`
	Image   string `json:"image" db:"image"`
	Link    string `json:"link" db:"link"`
}

func (m *Menu) Index() ([]*Lang, error) {
	langs := make([]*Lang, 3)
	langs[0] = &Lang{Id: 1, Name: "Thai Female"}
	langs[1] = &Lang{Id: 2, Name: "UK English Female"}
	langs[2] = &Lang{Id: 3, Name: "Chinese Female"}

	var sql string

	for _, l := range langs {
		menus := []*Menu{}
		switch l.Id {
		case 1:
			sql = `SELECT id, name, image, link FROM menu`
		case 2:
			sql = `SELECT id,  name_en as name, image, link FROM menu`
		case 3:
			sql = `SELECT id, name_cn as name, image, link FROM menu`
		}
		log.Println("case:", l.Id)
		err := db.Select(&menus, sql)
		if err != nil {
			return nil, err
		}
		l.Menu = menus
		log.Println(l)
	}
	log.Println(langs)
	return langs, nil
}
