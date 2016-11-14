package model

import "time"

type Order struct {
	Id    uint64
	Time  *time.Time
	Total float64
	Items []*OrderItem
}

type OrderItem struct {
	Line    uint64
	OrderId uint64
	ItemId  uint64
	Size    ItemSize
	Price   float32
	Qty     int
}
