package main

import "fmt"

func main() {
	stack := []string{"a", "b"}

	for i := 0; i >= len(stack); i++ {
		fmt.Println(stack[len(stack)-1])
	}
}
