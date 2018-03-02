package bitstamp

import (
	"encoding/json"
	"strings"

	"github.com/toorop/go-pusher"
	"github.com/twistedogic/cryptotrade/exchange"
)

const (
	exchangeName = "bitstamp"
	pusherKey    = "de504dc5763aeef9ff52"
)

var currency = map[string]string{
	"BTCUSD": "live_orders",
	"BTCEUR": "live_orders_btceur",
	"EURUSD": "live_orders_eurusd",
	"XRPUSD": "live_orders_xrpusd",
	"XRPEUR": "live_orders_xrpeur",
	"XRPBTC": "live_orders_xrpbtc",
	"LTCUSD": "live_orders_ltcusd",
	"LTCEUR": "live_orders_ltceur",
	"LTCBTC": "live_orders_ltcbtc",
	"ETHUSD": "live_orders_ethusd",
	"ETHEUR": "live_orders_etheur",
	"ETHBTC": "live_orders_ethbtc",
	"BCHUSD": "live_orders_bchusd",
	"BCHEUR": "live_orders_bcheur",
	"BCHBTC": "live_orders_bchbtc",
}

// Bitstamp exchange struct
type Bitstamp struct {
	Client *pusher.Client
}

// New returns pointer to Bitstamp
func New() (*Bitstamp, error) {
	client, err := pusher.NewClient(pusherKey)
	if err != nil {
		return nil, err
	}
	return &Bitstamp{Client: client}, nil
}

func eventToExchangeOrder(event *pusher.Event) (exchange.Order, error) {
	var res order
	var exchangeOrder exchange.Order
	symbol := "BTCUSD"
	channelName := strings.Split(event.Channel, "_")
	if len(channelName) > 2 {
		symbol = strings.ToUpper(channelName[len(channelName)-1])
	}
	err := json.Unmarshal([]byte(event.Data), &res)
	if err != nil {
		return exchangeOrder, err
	}
	res.Symbol = symbol
	if action, ok := actionMapping[event.Event]; ok {
		exchangeOrder = res.ToExchangeOrder(action)
	}
	return exchangeOrder, err
}

//StreamOrder write event stream to channel with given currency pair
func (b *Bitstamp) StreamOrder(ch chan exchange.Order) error {
	channel := currency["BTCUSD"]
	err := b.Client.Subscribe(channel)
	if err != nil {
		return err
	}
	streams := make([]chan *pusher.Event, 0, len(actionMapping))
	for action := range actionMapping {
		stream, err := b.Client.Bind(action)
		if err != nil {
			return err
		}
		streams = append(streams, stream)
	}
	for _, stream := range streams {
		go func(eventStream chan *pusher.Event) {
			for event := range eventStream {
				exchangeOrder, err := eventToExchangeOrder(event)
				if err == nil {
					ch <- exchangeOrder
				}
			}
		}(stream)
	}
	return err
}
