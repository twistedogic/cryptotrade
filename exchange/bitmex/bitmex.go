package bitmex

import (
	"github.com/gorilla/websocket"
	"github.com/twistedogic/cryptotrade/exchange"
)

const (
	base                 = "wss://www.bitmex.com/realtime"
	subscribe            = "subscribe"
	operationOrderBookL2 = "orderBookL2"
	operationQuote       = "quote"
)

type Bitmex struct {
	Client *websocket.Conn
}

func New() (*Bitmex, error) {
	c, _, err := websocket.DefaultDialer.Dial(base, nil)
	if err != nil {
		return nil, err
	}
	return &Bitmex{Client: c}, nil
}

type operation struct {
	Operation string `json:"op"`
	Argument  string `json:"args"`
}

func (b *Bitmex) StreamOrder(ch chan exchange.Order) error {
	err := b.Client.WriteJSON(operation{
		Operation: subscribe,
		Argument:  operationOrderBookL2,
	})
	if err != nil {
		return err
	}
	go func() {
		for {
			var res event
			_ = b.Client.ReadJSON(&res)
			for _, data := range res.ToExchangeOrders() {
				ch <- data
			}
		}
	}()
	return nil
}
