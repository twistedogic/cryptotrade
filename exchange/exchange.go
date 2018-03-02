package exchange

import (
	"time"

	"github.com/twistedogic/cryptotrade/common/pair"
)

//OrderType enum type
type OrderType uint

//ActionType enum type
type ActionType uint

const (
	//ASK order enum
	ASK OrderType = iota
	//BID order enum
	BID
)

const (
	//INSERT event action enum
	INSERT ActionType = iota
	//UPDATE event action enum
	UPDATE
	//DELETE event action enum
	DELETE
)

//Order exchange order datatype
type Order struct {
	ID        int
	Exchange  string
	Symbol    pair.CurrencyPair
	Price     float64
	Size      float64
	Side      OrderType
	Timestamp time.Time
	Action    ActionType
}

type Exchange interface {
	PlaceOrder(pair.CurrencyPair, float64) error
}

type DataSource interface {
	StreamOrder(chan Order) error
}
