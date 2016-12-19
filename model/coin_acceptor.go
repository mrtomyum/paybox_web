package model

type CoinAcceptor struct {
	Msg
}

func (c *CoinAcceptor) Action(Msg) {
	switch c.Payload.Type {
	case "request": // Msg from web client.
		c.OnRequest()
	case "response": // Response from Device
		c.OnResponse()
	case "event":
		c.OnEvent()
	}
}

func (c *CoinAcceptor) OnRequest() {

}

func (c *CoinAcceptor) OnResponse() {

}

func (c *CoinAcceptor) OnEvent() {

}