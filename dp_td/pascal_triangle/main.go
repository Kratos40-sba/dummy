package main

import "fmt"

func recursivePascalTriangle(n, k int) int {
	if n == k || k == 0 {
		return 1
	} else {
		return recursivePascalTriangle(n-1, k-1) + recursivePascalTriangle(n-1, k)
	}
}
func dp(n, k int) int {
	t := make([][]int, n+1)
	for i := range t {
		t[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		t[i][i] = 1
		t[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			t[i][j] = t[i-1][j-1] + t[i-1][j]
		}
	}

	return t[n][k]
}
func main() {
	fmt.Println(recursivePascalTriangle(100, 20))
	fmt.Println(dp(100, 20))
}
