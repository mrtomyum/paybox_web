package model

type Message struct {
	Device  string
	Type    string
	Command string
	Result  bool
	Data    interface{}
}
