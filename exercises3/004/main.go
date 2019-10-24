package main

import "fmt"

func main() {

	born := 1981
	for {
		if born == 2020 {
			break
		}
		fmt.Println(born)
		born++
	}
}
