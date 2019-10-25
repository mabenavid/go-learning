package main

import "fmt"

func main() {

	s := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v := range s {
		fmt.Println(i)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

}
