package model

// Ticket เก็บฟอร์มตั๋วรูปแบบต่างๆ โดยรวม FromRow แต่ละแถวมาเข้าชุดกัน
type Ticket struct {
	Id    int
	Name  string
	Forms []*TicketForm
}

// TicketForm  เป็น Junction Table ผูกฟอร์มและ Tiketไว้ด้วยกัน
type TicketForm struct {
	TicketId int
	Seq      int
	FormId   int
}

// Form เก็บอักขระฟอร์มต้นแบบ เป็นแถว
type Form struct {
	Id   int
	Name string
	Row  string
}
