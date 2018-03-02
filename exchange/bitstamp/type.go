package bitstamp

import (
	"time"

	"github.com/twistedogic/cryptotrade/common/pair"
	"github.com/twistedogic/cryptotrade/exchange"
)

var actionMapping = map[string]exchange.ActionType{
	"order_created": exchange.INSERT,
	"order_changed": exchange.UPDATE,
	"order_deleted": exchange.DELETE,
}

var sideMapping = map[int]exchange.OrderType{
	0: exchange.BID,
	1: exchange.ASK,
}

type order struct {
	Symbol string
	ID     int     `json:"id"`
	Size   float64 `json:"amount"`
	Price  float64 `json:"price"`
	Side   int     `json:"order_type"`
}

func (o order) ToExchangeOrder(action exchange.ActionType) exchange.Order {
	var side exchange.OrderType
	if val, ok := sideMapping[o.Side]; ok {
		side = val
	}
	return exchange.Order{
		Exchange:  exchangeName,
		Symbol:    pair.NewCurrencyPairFromString(o.Symbol),
		Side:      side,
		Size:      o.Size,
		Price:     o.Price,
		ID:        o.ID,
		Action:    action,
		Timestamp: time.Now(),
	}
}
