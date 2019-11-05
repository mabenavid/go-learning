package main

import (
	"fmt"
)

func main() {

	go func() {
		fmt.Println("anonymus func")
	}()

	f := func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}

	f()

}
