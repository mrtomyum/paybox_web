package model

import "time"

type Order struct {
	Id   uint64
	Time *time.Time
}

type OrderItem struct {
	OrderId uint64
	ItemId uint64
	Unit string
	Size ItemSize
	Price float32
}
