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

	opts := val.NewOptArray([]val.Option{val.NewOption(val.NewInt(100)), val.NewOption(nil)})
	expr := parsing.DefaultParser().Parse("@0 = @1").(propexpr.PropExpr)
	b := interp.Query(expr, opts)

	fmt.Println(opts)

	fmt.Println(b)
}
