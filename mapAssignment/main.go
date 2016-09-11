package main

import "fmt"

func main() {
	at := map[string]string{
		"Patterson": "Along Came a Spider",
		"Brown":     "The Da Vinci Code",
		"Christ":    "Bible",
	}
	fmt.Println(at["Brown"])
	fmt.Println(at)

	ab := make(map[string]string)
	ab["Patterson"] = "Along Came a Spider"
	ab["Brown"] = "The Da Vinci Code"
	ab["Christ"] = "The Bible"
	fmt.Println(ab["Patterson"])
	fmt.Println(ab)
}
