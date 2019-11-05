package main

import (
	"fmt"
)

func returnFunc() func() {
	return func() {
		fmt.Println("func returned")
	}
}

func main() {

	f := returnFunc()

	f()

}
