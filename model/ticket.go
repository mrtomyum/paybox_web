package model

// Ticket เก็บฟอร์มตั๋วรูปแบบต่างๆ โดยรวม FromRow แต่ละแถวมาเข้าชุดกัน
type Ticket struct {
	Id   int
	Name string
	Data map[string]interface{}
}

// TicketForm  เป็น Junction Table ผูกฟอร์มและ Tiketไว้ด้วยกัน
type TicketAction struct {
	TicketId int
	Seq      int
	ActionId int
}

// Form เก็บอักขระฟอร์มต้นแบบ เป็นแถว
type Action struct {
	Print        string `json:"print"`
	Printline    string `json:"printline"`
	SetTextSize  int `json:"set_text_size"`
	PrintBarcode PrintBarcode `json:"print_barcode"`
	PrintQr      PrintQr `json:"print_qr"`
}

type PrintBarcode struct {
	Type string // type คือชนิดของ Barcode “UPC-A”, “UPC-E”, “JAN13”, “JAN8”, “CODE39”, “ITF”, “CODABAR”, “CODE128”
	Data string // data คือ ข้อมูลที่จะบรรจุลงไปใน Barcode
}

type PrintQr struct {
	Mag      int `json:"mag"`             // mag คือ ขนาดของ QR Code โดย 0=ขนาดปกติ, 1=ขนาด2เท่า, 2=ขนาด3เท่า, 3=ขนาด4เท่า
	Ecl      int `json:"ecl"`             // ecl คือค่า Error Correction Level โดย 0=30%, 1=25%, 2=15%, 3=7%
	DataType string  `json:"data_type"`   // data_type คือ ค่าชนิดของข้อมูล โดย “number” หมายถึงข้อมูลที่มีแต่ตัวเลข “0-9”“alpha” หมายถึงข้อมูลที่เป็นตัวหนังสือทั่วไป “8bit” หมายถึงข้อมูลประเภท Byte
	Data     interface{}`json:"data"`     // data คือ ข้อมูลที่จะบรรจุลงไปใน QR Code, ในกรณีที่ข้อมูลเป็น 8bit ให้ส่งข้อมูลมาในรูปแบบ base64
	PaperCut PaperCut  `json:"paper_cut"` // คำสั่งสำหรับสั่งตัดกระดาษ
}

type PaperCut struct {
	Type string //โดย type คือรูปแบบของการตัดกระดาษ มี 2 ชนิด ได้แก่ “full_cut” และ “partial_cut”
	Feed int    // feed คือความยาวของกระดาษที่จะให้ทำการ Feed ออกมาก่อนที่จะติด
}

func (t *Ticket) PrintHeader() {

}

func (t *Ticket) PrintBody() {

}

func (t *Ticket) PrintBarcode() {

}
