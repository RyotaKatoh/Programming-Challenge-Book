package main

import (
	"fmt"
	"math"
)

func triangle(n int, a []int) int {
	l := len(a)
	maxLen := 0.0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				length := a[i] + a[j] + a[k]
				m := math.Max(float64(a[i]), math.Max(float64(a[j]), float64(a[k])))
				if length > 2*int(m) {
					maxLen = math.Max(float64(length), maxLen)
				}
			}
		}
	}
	return int(maxLen)
}

func main() {
	fmt.Println(triangle(5, []int{2, 3, 4, 5, 10}))
	fmt.Println(triangle(4, []int{4, 5, 10, 20}))
}
