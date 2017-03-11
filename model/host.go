package model

type Host struct {
	Id               string  // รหัสเมนบอร์ดตู้
	IsNetOnline      bool    // สถานะ GSM ปัจจุบัน (Real time)
	IsServerOnline   bool    // สถานะเซิร์ฟเวอร์ครั้งสุดท้ายที่สื่อสาร
	Web              *Socket // Web Client object ที่เปิดคอนเนคชั่นอยู่
	Hw               *Socket // Device Client object ที่เปิดคอนเนคชั่นอยู่
	LastTicketNumber int     // เลขคิวตั๋วล่าสุด ของแต่ละวัน ปิดเครื่องต้องยังอยู่ ขึ้นวันใหม่ต้อง Reset
}
