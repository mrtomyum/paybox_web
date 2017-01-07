package model

// ====================
// Send
// ====================
type Msg struct {
	Device  string  `json:"device"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Type    string      `json:"type"`
	Command string      `json:"command"`
	Result  bool        `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
