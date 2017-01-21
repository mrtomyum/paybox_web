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
type Action struct {
	Print        string
	Printline    string
	SetTextSize  int
	PrintBarcode PrintBarcode
	PrintQr      PrintQr
}

type BarcodeType string

const (
//UPC-A BarcodeType  = "UPC-A"
//UPC-E BarcodeType = "UPC-E"
//JAN13 BarcodeType = "JAN13"
//CODE39 BarcodeType = "CODE39"
//ITF BarcodeType = "ITF"
//CODABAR BarcodeType = "CODABAR"
//CODE128 BarcodeType = "CODE128"
)

type PrintBarcode struct {
	Type BarcodeType
	Data string
}

type PrintQr struct {
	Mag      int         // mag คือ ขนาดของ QR Code โดย 0=ขนาดปกติ, 1=ขนาด2เท่า, 2=ขนาด3เท่า, 3=ขนาด4เท่า
	Ecl      int         // ecl คือค่า Error Correction Level โดย 0=30%, 1=25%, 2=15%, 3=7%
	DataType string      // data_type คือ ค่าชนิดของข้อมูล โดย “number” หมายถึงข้อมูลที่มีแต่ตัวเลข “0-9”“alpha” หมายถึงข้อมูลที่เป็นตัวหนังสือทั่วไป “8bit” หมายถึงข้อมูลประเภท Byte
	Data     interface{} // data คือ ข้อมูลที่จะบรรจุลงไปใน QR Code, ในกรณีที่ข้อมูลเป็น 8bit ให้ส่งข้อมูลมาในรูปแบบ base64
	PaperCut PaperCut    // คำสั่งสำหรับสั่งตัดกระดาษ
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
