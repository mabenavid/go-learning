package main

import "fmt"

func main() {
	defer foo()
	i, s := bar()
	fmt.Println(i, s)
}

func foo() {
	fmt.Println("last")
}

func bar() (int, string) {
	return 11, "Paul"
}
