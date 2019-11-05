package main

import (
	"fmt"
)

func main() {
	n := fibonacci(10)
	fmt.Println(n)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
