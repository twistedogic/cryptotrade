package bitmex

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEventUnmarshal(t *testing.T) {
	input := []byte(`{
    "table":"orderBookL2",
    "keys":["symbol","id","side"],
    "types":{"id":"long","price":"float","side":"symbol","size":"long","symbol":"symbol"},
    "foreignKeys":{"side":"side","symbol":"instrument"},
    "attributes":{"id":"sorted","symbol":"grouped"},
    "action":"partial",
    "data":[
      {"symbol":"XBTUSD","id":17999992000,"side":"Sell","size":100,"price":80},
      {"symbol":"XBTUSD","id":17999995000,"side":"Buy","size":10,"price":50}
    ]
  }`)
	output := event{}
	expect := event{
		Action: "partial",
		Order: []order{
			order{Symbol: "XBTUSD", ID: 17999992000, Side: "Sell", Size: float64(100), Price: float64(80)},
			order{Symbol: "XBTUSD", ID: 17999995000, Side: "Buy", Size: float64(10), Price: float64(50)},
		},
	}
	err := json.Unmarshal(input, &output)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(output, expect) {
		t.Fail()
	}
}
