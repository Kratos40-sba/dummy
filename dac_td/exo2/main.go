package main

import "fmt"

func dac(t []int) int {
	if len(t) == 0 {
		return 0
	} else if len(t) == 1 {
		return t[0]
	} else if len(t) == 2 {
		return max(t[0], t[1])
	} else {
		m := len(t) / 2
		return max(dac(t[:m]), dac(t[m+1:]))
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	t := []int{12, 5, 23, 44, 70, 1, 2, 3, 5, 88}
	fmt.Println(dac(t))
}
