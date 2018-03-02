package util

import (
	"strconv"
)

func StringSliceToFloatSlice(s []string) ([]float64, error) {
	output := make([]float64, len(s))
	for i, v := range s {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return output, err
		}
		output[i] = f
	}
	return output, nil
}
