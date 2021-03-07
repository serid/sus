package main

import (
	"fmt"
	"sus/interp"
	"sus/interp/val"
	"sus/syntax/parsing"
	"sus/syntax/parsing/propexpr"
)

func main() {
	fmt.Printf("%#v", parsing.DefaultParser().Parse("Test(0, 1, 2)"))
}

func example1() {
	fmt.Printf("%#v", parsing.DefaultParser().Parse("@0 = @1 /\\ @0 = @1"))
}

func example2() {
	fmt.Println("Hello world!")

	vals := []val.Val{val.NewInt(100), nil}
	expr := parsing.DefaultParser().Parse("@0 = @1").(propexpr.PropExpr)
	solution := interp.Query(expr, vals)

	fmt.Println(solution)
}
