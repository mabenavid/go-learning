package main

import "fmt"

func main() {

	m := make(map[string][]string)

	m["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	m["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	m2 := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)
	fmt.Println(m2)

}
