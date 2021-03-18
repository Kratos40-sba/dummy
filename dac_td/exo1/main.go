package main

import "fmt"

func powerToNRecursive(a, n int) int {
	if n == 0 {
		return 1
	} else if a == 0 {
		return 0
	} else {
		return a * powerToNRecursive(a, n-1)
	}
}
func powerToNDac(a, n int) int {
	if n == 0 {
		return 1
	} else if a == 0 {
		return 0
	} else {
		if n%2 == 0 {
			return powerToNRecursive(a*a, n/2)
		} else {
			return a * powerToNRecursive(a*a, n/2)
		}
	}
}
func iterativeApproach(a, n int) int64 {
	if n == 0 {
		return 1
	} else if a == 0 {
		return 0
	}
	var t int64 = 1
	for i := 0; i < n; i++ {
		t = int64(a) * t
	}
	return t
}

func main() {
	fmt.Println(powerToNRecursive(2, 30))
	fmt.Println(powerToNDac(2, 30))
	fmt.Println(iterativeApproach(2, 30))
}
