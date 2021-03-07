package cmp

import (
	"fmt"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// This function tries to compare any types in this package
func Cmp(a, b interface{}) bool {
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
	default:
		panic(fmt.Sprintf("unhandled arg in `Cmp`: %#v", a))
	}
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
		return ValExpr(a1.E1(), b1.E1()) && ValExpr(a1.E2(), b1.E2())
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
