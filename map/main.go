package main

import "fmt"

func main() {
	// make map with composite literal
	m := map[string]int{
		"Todd": 45,
		"Matt": 40,
	}
	fmt.Println(m["Matt"])
	fmt.Println(m)

	// make map with `make` keyword
	x := make(map[string]int)
	x["Todd"] = 45
	x["Matt"] = 40
	fmt.Println(x["Todd"])
	fmt.Println(x)

}
