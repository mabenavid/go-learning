package main

import "fmt"

func main() {
	x := foo([]int{1, 2, 3, 4, 5}...)
	fmt.Println(x)
	y := bar([]int{1, 2, 3, 4, 5})
	fmt.Println(y)
}

func foo(num ...int) int {
	total := 0
	for _, v := range num {
		total += v
	}
	return total
}

func bar(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
