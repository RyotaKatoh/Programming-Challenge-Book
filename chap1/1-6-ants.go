package main

import (
	"fmt"
	"math"
)

func ants(L int, x []int) (int, int) {
	min := 0.0
	max := 0.0
	for _, v := range x {
		min = math.Max(min, math.Min(float64(v), float64(L-v)))
		max = math.Max(max, math.Max(float64(v), float64(L-v)))
	}

	return int(min), int(max)
}

func main() {
	min, max := ants(10, []int{2, 6, 7})
	fmt.Println(min, max)
}
