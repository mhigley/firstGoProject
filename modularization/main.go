package main

import (
	"fmt"
	// importing default `fmt` package
	someAlias "github.com/GoesToEleven/GolangTraining/02_package/stringutil"
	// applying custom alias to package. referred to now by alias name
)

func main() {
	fmt.Println("test")
	fmt.Println(someAlias.Reverse("!oG olleH"))
}
