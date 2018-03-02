package util

import (
	"reflect"
	"testing"
)

func TestStringSliceToFloatSlice(t *testing.T) {
	var testCases = []struct {
		input  []string
		expect []float64
	}{
		{
			[]string{"1.1000"},
			[]float64{1.1000},
		},
	}
	for _, test := range testCases {
		output, err := StringSliceToFloatSlice(test.input)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(output, test.expect) {
			t.Fail()
		}
	}
}
