package client

import (
	"github.com/twistedogic/cryptotrade/common/pair"
)

type Client interface {
	StreamData(pair.CurrencyPair) (<-chan []byte, error)
}
