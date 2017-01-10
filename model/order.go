package model

import "time"

type Order struct {
	Id    uint64
	Time  *time.Time
	Total float64
	Items []*OrderSub
}

type OrderSub struct {
	Line     uint64 `json:"line"`
	OrderId  uint64
	ItemId   uint64  `json:"item_id"`
	ItemName string  `json:"item_name"`
	SizeId   int     `json:"size_id"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
}
