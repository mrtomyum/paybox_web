package model

type IndexPage struct {
	LangId   int `json:"lang_id"`
	LangName string `json:"lang_name"`
	Menu     []*Menu `json:"menu"`
}

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

func (m *Menu) Index() (data []*IndexPage, err error) {
	//data = IndexPage{
	//	LangId: 1,
	//	LangName: "UK English Female",
	//}
	//page := IndexPage{}
	//menus := []*Menu{}
	//for i, v := range data {
	//	sql := `SELECT * FROM menu`
	//	err = db.Select(&menus, sql)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//}
	//
	//
	//data = append(data, page)
	return data, nil
}

