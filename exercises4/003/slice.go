package main

import "fmt"

func main() {

	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//[42 43 44 45 46]
	fmt.Println(s[:5])
	//[47 48 49 50 51]
	fmt.Println(s[5:])
	//[44 45 46 47 48]
	fmt.Println(s[2:7])
	//[43 44 45 46 47]
	fmt.Println(s[1:6])

}
