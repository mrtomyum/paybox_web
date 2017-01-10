package model

type CoinAcceptor struct {
	Id     string
	Status string
	Send   chan *Message
}

