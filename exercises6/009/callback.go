package main

import (
	"fmt"
)

func callback() {
	fmt.Println("callback executed")
}

func foo(f func()) {

	fmt.Println("foo executed")

	f()

}

func main() {

	f := callback
	foo(f)

}
