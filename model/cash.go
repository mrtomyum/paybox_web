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

type AcceptedBill struct {
	THB20   bool
	THB50   bool
	THB100  bool
	THB500  bool
	THB1000 bool
}
