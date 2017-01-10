package model

type Printer struct {
	Id     string
	Status string
	Send   chan *Message
}

func (p *Printer) Print(o *Order) error {

	return nil
}