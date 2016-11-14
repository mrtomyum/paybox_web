package model

type BaseDevice struct {
	Id     int
	Type   int
	Name   string
	Status string
}

type CoinAcceptor struct {
	BaseDevice
}

type BillAcceptor struct {
	BaseDevice
}

func (b BillAcceptor) Status() string {
	var status string
	return status
}

type Devicer interface {
	Status() string
}

func CheckStatus(d Devicer) string {
	status := d.Status()
	return status
}