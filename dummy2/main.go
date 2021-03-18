package main

import "fmt"

func dac(a, b int, A []int) int {
	h, l := 0, len(A)
	m := (l + h) / 2
	if A[m] <= b && A[m] >= a {
		return A[m]
	} else if A[m] >= b {
		return dac(a, b, A[:m])
	} else if A[m] <= a {
		return dac(a, b, A[m+1:])
	} else {
		return -1
	}

}
func main() {
	fmt.Println(dac(40, 50, []int{3, 7, 8, 43, 556}))
}
