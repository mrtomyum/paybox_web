package model

// Onhand คือยอดเงินพัก ยังไม่ได้รับชำระ
type Onhand struct {
	Coin  float64 // มูลค่าเหรียญพัก ที่ยังไม่ได้รับชำระ
	Bill  float64 // มูลค่าธนบัตรที่พักอยู่ในเครื่องรับธนบัตร
	Total float64 // มูลค่าเงินพักทั้งหมด
}

// CashBox คือถังเก็บเงิน แยกเป็น 3 จุด
// คือ Hopper ถังเก็บเหรียญ CoinBox และถังเก็บธนบัตร BillBox
type CashBox struct {
	Hopper float64 // มูลค่าเหรียญใน Coin Hopper
	Coin   float64 // มูลค่าเหรียญใน CainBox
	Bill   float64 // มูลค่าธนบัตรในกล่องเก็บธนบัตร
	Total  float64 // รวมมูลค่าเงินในตู้นี้
}

// AcceptedBill ระบุค่ายอดขายขั้นต่ำที่ยอมรับธนบัตรแต่ละขนาด 0 = ไม่จำกัด
type AcceptedBill struct {
	THB20   int `json:"thb_20"`
	THB50   int `json:"thb_50"`
	THB100  int `json:"thb_100"`
	THB500  int `json:"thb_500"`
	THB1000 int `json:"thb_1000"`
}

