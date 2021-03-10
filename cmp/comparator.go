package cmp

import (
	"fmt"
	"strings"
	"sus/interp/bcinterp/bytecode"
	"sus/stuff"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// Tries to compare `a` and `b` using simple comparison operator and,
// if they are not of comparable type, uses `fallbackCmp` function.
//
// NOTE: if a and b have differing types, the comparison fails as expected from (==) operator
// NOTE: "comparable type" is used as defined in go spec
// https://golang.org/ref/spec#Comparison_operators
func SmartCmp(a, b interface{}, fallbackCmp func(a, b interface{}) bool) bool {
	p := stuff.Catch(func() interface{} {
		return a == b
	})

	// If simple (==) successfully returned a boolean, pass the return value
	if b, ok := p.(bool); ok {
		return b
	}

	// Otherwise, a panic occurred. Check if it is the "comparing uncomparable type" panic
	err := fmt.Sprint(p)

	if !strings.HasPrefix(err, "runtime error: comparing uncomparable type ") {
		panic("unreachable")
	}

	// Call custom comparison routine
	return fallbackCmp(a, b)
}

// This function tries to compare any types in this package
func Cmp(a, b interface{}) bool {
	isEqual := SmartCmp(a, b, func(a, b interface{}) bool {
		switch a1 := a.(type) {
		case propexpr.PropExpr:
			b1, ok := b.(propexpr.PropExpr)
			if !ok {
				return false
			}
			return PropExpr(a1, b1)
		case valexpr.ValExpr:
			b1, ok := b.(valexpr.ValExpr)
			if !ok {
				return false
			}
			return ValExpr(a1, b1)
		case bytecode.Op:
			b1, ok := b.(bytecode.Op)
			if !ok {
				return false
			}
			return a1.OpCode == b1.OpCode && Cmp(a1.Data, b1.Data)
		case bytecode.SolRuleCallData:
			b1, ok := b.(bytecode.SolRuleCallData)
			if !ok {
				return false
			}
			return a1.Rid == b1.Rid && a1.Output == b1.Output && Cmp(a1.Output, b1.Output)
		default:
			panic(fmt.Sprintf("unhandled arg in `Cmp`: %#v", a))
		}
	})

	return isEqual
}

func PropExpr(a, b propexpr.PropExpr) bool {
	switch a1 := a.(type) {
	case propexpr.Conjunction:
		b1, ok := b.(propexpr.Conjunction)
		if !ok {
			return false
		}
		return PropExpr(a1.E1(), b1.E1()) && PropExpr(a1.E2(), b1.E2())
	case propexpr.Disjunction:
		b1, ok := b.(propexpr.Disjunction)
		if !ok {
			return false
		}
		return PropExpr(a1.E1(), b1.E1()) && PropExpr(a1.E2(), b1.E2())
	case propexpr.Unification:
		b1, ok := b.(propexpr.Unification)
		if !ok {
			return false
		}
		return ValExpr(a1.E1, b1.E1) && ValExpr(a1.E2, b1.E2)
	case propexpr.RuleCall:
		b1, ok := b.(propexpr.RuleCall)
		if !ok {
			return false
		}
		return a1.Rid() == b1.Rid() && ValExprSliceEquivalent(a1.Args(), b1.Args())
	case propexpr.True:
		return a == b
	default:
		panic("Unhandled arg in PropExpr.")
	}
}

func ValExpr(a, b valexpr.ValExpr) bool {
	switch a1 := a.(type) {
	case valexpr.Plus:
		b1, ok := b.(valexpr.Plus)
		if !ok {
			return false
		}
		return ValExpr(a1.E1(), b1.E1()) && ValExpr(a1.E2(), b1.E2())
	case valexpr.Mul:
		b1, ok := b.(valexpr.Mul)
		if !ok {
			return false
		}
		return ValExpr(a1.E1(), b1.E1()) && ValExpr(a1.E2(), b1.E2())
	case valexpr.CommaListPair:
		b1, ok := b.(valexpr.CommaListPair)
		if !ok {
			return false
		}
		return ValExpr(a1.V(), b1.V()) && ValExpr(a1.Tail(), b1.Tail())
	case valexpr.IntLit, valexpr.GetVar, valexpr.Unit:
		return a == b
	default:
		panic("Unhandled arg in ValExpr.")
	}
}

func ValExprSliceEquivalent(a, b []valexpr.ValExpr) bool {
	if (a == nil) || (b == nil) {
		panic("Attempted to comapre nil slices for equivalence.")
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !ValExpr(a[i], b[i]) {
			return false
		}
	}
	return true
}
