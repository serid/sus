package interp

import (
	"fmt"
	"sus/interp/val"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

func Query(propExpr propexpr.PropExpr, opts val.OptArray) bool {
	switch l := propExpr.(type) {
	case propexpr.True:
		return true
	case propexpr.Unification:
		var val1 = evalValExpr(l.E1(), opts)
		var val2 = evalValExpr(l.E2(), opts)

		// If one of args to Unification is nil, copy value overwriting None

		if val1.IsNone() {
			val1.SetData(val2.Unwrap())
			return true
		}
		if val2.IsNone() {
			val2.SetData(val1.Unwrap())
			return true
		}

		// If Values in expressions are equal, return true, otherwise false
		if val1.Unwrap() == val2.Unwrap() {
			return true
		} else {
			return false
		}
	case propexpr.Conjunction:
		var success1 = Query(l.E1(), opts)

		if success1 {
			var success2 = Query(l.E2(), opts)

			if success2 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	case propexpr.Disjunction:
		var success1 = Query(l.E1(), opts.Clone())

		if success1 {
			return true
		} else {
			var success2 = Query(l.E2(), opts)

			if success2 {
				return true
			} else {
				return false
			}
		}
	default:
		panic(fmt.Sprintf("Unhandled PropExpr: %#v.", l))
	}
}

func evalValExpr(expr valexpr.ValExpr, opts val.OptArray) *val.Option {
	switch expr := expr.(type) {
	case valexpr.GetVar:
		return &opts.Options()[expr.VarNum()]
	case valexpr.IntLit:
		v := val.NewOption(val.NewInt(expr.Data()))
		return &v
	default:
		panic(fmt.Sprintf("Unhandled ValExpr: %#v.", expr))
	}
}
