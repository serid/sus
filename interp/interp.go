package interp

import (
	"fmt"
	"sus/interp/val"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// Attempts to find a solution to the `propExpr` using values and holes provided in `opts`.
// Value passed to `opts` is considered to be "moved" and should not be used after calling this function.
// The function might mutate `opts` value and/or return it as a query result.
func Query(propExpr propexpr.PropExpr, opts val.OptArray) val.OptArray {
	switch l := propExpr.(type) {
	case propexpr.True:
		return opts
	case propexpr.Unification:
		var val1 = evalValExpr(l.E1(), opts)
		var val2 = evalValExpr(l.E2(), opts)

		// If one of args to Unification is nil, copy value overwriting None

		if val1.IsNone() {
			val1.SetData(val2.Unwrap())
			return opts
		}
		if val2.IsNone() {
			val2.SetData(val1.Unwrap())
			return opts
		}

		// If Values in expressions are equal, return true, otherwise false
		if val1.Unwrap() == val2.Unwrap() {
			return opts
		} else {
			return val.NewOptArray(nil)
		}
	case propexpr.Conjunction:
		var success1 = Query(l.E1(), opts)

		if success1.IsSome() {
			var success2 = Query(l.E2(), success1)

			if success2.IsSome() {
				return success2
			} else {
				return val.NewOptArray(nil)
			}
		} else {
			return val.NewOptArray(nil)
		}
	case propexpr.Disjunction:
		var success1 = Query(l.E1(), opts.Clone())

		if success1.IsSome() {
			return success1
		} else {
			var success2 = Query(l.E2(), opts)

			if success2.IsSome() {
				return success2
			} else {
				return val.NewOptArray(nil)
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
