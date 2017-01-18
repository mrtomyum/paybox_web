package model

type Message struct {
	Device  string `json:"device"`
	Type    string `json:"type"`
	Command string `json:"command"`
	Result  bool `json:"result"`
	Data    interface{} `json:"data"`
}
