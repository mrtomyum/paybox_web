package model

import "github.com/jmoiron/sqlx"

var (
	db *sqlx.DB
	H  Host
	B  BillAcceptor
	C  CoinAcceptor
	CH CoinHopper
)

func init() {
	H = Host{
		Id:            "001",
		Online:        true,
		TotalEscrow:   0,
		BillEscrow:    0,
		BillBox:       0,
		CoinHopperBox: 0,
		CoinBox:       0,
		TotalCash:     0,
		SetWebClient:  make(chan *Client),
		SetDevClient:  make(chan *Client),
	}
	B = BillAcceptor{
		status: "ok",
	}
	C = CoinAcceptor{
		status: "ok",
	}
	CH = CoinHopper{
		status: "ok",
	}
}
