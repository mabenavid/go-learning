package main

import "fmt"

func main() {
	x := foo()
	i, s := bar()
	fmt.Println(x)
	fmt.Println(i, s)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 11, "Paul"
}
