package main

import (
	"fmt"
	"sort"
)

func kuji(m int, k []int) bool {
	l := len(k)
	kk := []int{}
	for c := 0; c < l; c++ {
		for d := 0; d < l; d++ {
			kk = append(kk, k[c]+k[d])
		}
	}

	sort.Ints(kk)
	for a := 0; a < l; a++ {
		for b := 0; b < l; b++ {
			target := m - k[a] - k[b]
			ans := sort.SearchInts(kk, target)
			if ans < len(kk) && kk[ans] == target {
				return true
			}
		}
	}

	return false
}

func main() {

	fmt.Println(kuji(10, []int{1, 3, 5}))
	fmt.Println(kuji(9, []int{1, 3, 5}))

}
