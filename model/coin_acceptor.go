package model

type CoinAcceptor struct {
	Msg
}

func (ca *CoinAcceptor) Action(d Device, m Msg) {
	switch ca.Payload.Type {
	case "request": // Msg from web client.
		ca.OnRequest(d, m)
	case "response": // Response from Device
		ca.OnResponse(d, m)
	case "event":
		ca.OnEvent(d, m)
	}
}

func (ca *CoinAcceptor) OnRequest(d Device, m Msg) {

}

func (ca *CoinAcceptor) OnResponse(d Device, m Msg) {

}

func (ca *CoinAcceptor) OnEvent(d Device, m Msg) {

}