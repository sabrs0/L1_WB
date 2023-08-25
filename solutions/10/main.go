package main

import (
	"math"

	"fmt"
)

var (
	temps = []float64{-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}
)

func main() {
	ans := make(map[int][]float64)

	for _, val := range temps {
		key := int(math.Trunc(val/10) * 10)
		ans[key] = append(ans[key], val)
	}
	fmt.Println(ans)
}
