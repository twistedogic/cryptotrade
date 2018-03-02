package bitmex

import (
	"strings"
	"time"

	"github.com/twistedogic/cryptotrade/common/pair"
	"github.com/twistedogic/cryptotrade/exchange"
)

const (
	exchangeName = "bitmex"
)

var actionMapping = map[string]exchange.ActionType{
	"insert":  exchange.INSERT,
	"delete":  exchange.DELETE,
	"partial": exchange.INSERT,
	"update":  exchange.UPDATE,
}

var sideMapping = map[string]exchange.OrderType{
	"buy":  exchange.BID,
	"sell": exchange.ASK,
}

type order struct {
	Symbol string  `json:"symbol"`
	ID     int     `json:"id"`
	Side   string  `json:"side,omitempty"`
	Size   float64 `json:"size,omitempty"`
	Price  float64 `json:"price,omitempty"`
}

func (o order) ToExchangeOrder(action exchange.ActionType) exchange.Order {
	var side exchange.OrderType
	if val, ok := sideMapping[strings.ToLower(o.Side)]; ok {
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

type event struct {
	Action  string  `json:"action"`
	Success bool    `json:"success,omitempty"`
	Order   []order `json:"data"`
}

func (e event) ToExchangeOrders() []exchange.Order {
	orders := make([]exchange.Order, 0, len(e.Order))
	if action, ok := actionMapping[e.Action]; ok {
		for _, v := range e.Order {
			orders = append(orders, v.ToExchangeOrder(action))
		}
	}
	return orders
}
