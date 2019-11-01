package main

import "fmt"

type person struct {
	firstName, lastName string
	favoriteFlavors     []string
}

func main() {

	p1 := struct {
		firstName, lastName string
		favoriteFlavors     []string
	}{
		firstName:       "Manu",
		lastName:        "Bena",
		favoriteFlavors: []string{"vanilla", "chocolate"},
	}

	fmt.Println(p1)

}
