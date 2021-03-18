package main

import (
	"fmt"
	"sort"
)

func returnCoins(n int, l []int) int {
	sort.Ints(l)
	r := 0
	for n > 0 {
		for i := len(l) - 1; i >= 0; i-- {
			if n-l[i] >= 0 {
				n = n - l[i]
				r++
			}
		}
	}
	return r
}
func main() {
	fmt.Println(returnCoins(263, []int{1, 2, 5, 10, 50, 100, 200}))
}
