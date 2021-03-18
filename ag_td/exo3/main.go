package main

import "fmt"

func ag(n int, coins []int) int {
	ctr := 0
	for n > 0 {
		for i := 0; i < len(coins); i++ {
			if n >= coins[i] {
				n = n - coins[i]
				ctr++
			}
		}
	}
	return ctr
}
func main() {
	fmt.Println(ag(30, []int{25, 10, 1}))
}
