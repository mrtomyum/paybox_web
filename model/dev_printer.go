package model

import "fmt"

type Printer struct {
	Id     string
	Status string
	Send   chan *Message
}

func (p *Printer) Print(o *Order) error {
	fmt.Println("p.Print() run")
	return nil
}