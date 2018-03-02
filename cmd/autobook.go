package main

import (
	"fmt"

	"github.com/twistedogic/cryptotrade/exchange"
	"github.com/twistedogic/cryptotrade/exchange/bitstamp"
)

func main() {
	ex, err := bitstamp.New()
	if err != nil {
		panic(err)
	}
	ch := make(chan exchange.Order)
	err = ex.StreamOrder(ch)
	if err != nil {
		panic(err)
	}
	for data := range ch {
		fmt.Printf("%+v\n", data)
	}
}
