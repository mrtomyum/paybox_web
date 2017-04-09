package model

type Host struct {
	Id               string  // รหัสเมนบอร์ดตู้
	IsNetOnline      bool    // สถานะ GSM ปัจจุบัน (Real time)
	IsServerOnline   bool    // สถานะเซิร์ฟเวอร์ครั้งสุดท้ายที่สื่อสาร
	Web              *Socket // Web Client object ที่เปิดคอนเนคชั่นอยู่
	Hw               *Socket // Device Client object ที่เปิดคอนเนคชั่นอยู่
	LastTicketNumber int     // เลขคิวตั๋วล่าสุด ของแต่ละวัน ปิดเครื่องต้องยังอยู่ ขึ้นวันใหม่ต้อง Reset
}

// Init() เติมเงินเริ่มงานวันใหม่ ระบบจะถือว่าเงินใน Hopperเป็น 0 ในแต่ละครั้งที่เข้าโหมด Init คือต้องสั่ง Empty CHไปแล้ว
func (h *Host) Init() {
	// User input Coin start.
	//
}

// Shutdown() สั่งปิดตู้ ระบบจะทำการ Empty เหรียญ ในตู้ลง CashBox และ ส่งยอดและสถานะต่างๆ ไปยัง Cloud API
func (h *Host) Shutdown() {
	// Empty Coin_Hopper
	// Post EndDay Sale
	// Post Status
	// Post CashBox
	// if no network Save()
}
