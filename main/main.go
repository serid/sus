package main

import (
	"fmt"
	"sus/syntax/parsing"
)

func main() {
	fmt.Printf("%#v", parsing.DefaultParser().Parse("Test(0, 1, 2)"))
}
