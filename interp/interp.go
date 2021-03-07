package interp

import (
	"fmt"
	"sus/interp/val"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// Attempts to find a solution to the `propExpr` using values and holes provided in `vals`.
// Value passed to `vals` is considered to be "moved" and should not be used after calling this function.
// The function might mutate `vals` value and/or return it as a query result.
func Query(propExpr propexpr.PropExpr, vals val.Array) val.Array {
	switch l := propExpr.(type) {
	case propexpr.True:
		return vals
	case propexpr.Unification:
		var val1 = evalValExpr(l.E1(), vals)
		var val2 = evalValExpr(l.E2(), vals)

		// If one of args to Unification is nil, copy value overwriting None

		if *val1 == nil {
			*val1 = *val2
			return vals
		}
		if *val2 == nil {
			*val2 = *val1
			return vals
		}

		// If Values in expressions are equal, return true, otherwise false
		if *val1 == *val2 {
			return vals
		} else {
			return val.NewArray(nil)
		}
	case propexpr.Conjunction:
		var success1 = Query(l.E1(), vals)

		if success1.IsSome() {
			var success2 = Query(l.E2(), success1)

			if success2.IsSome() {
				return success2
			} else {
				return val.NewArray(nil)
			}
		} else {
			return val.NewArray(nil)
		}
	case propexpr.Disjunction:
		var success1 = Query(l.E1(), vals.Clone())

		if success1.IsSome() {
			return success1
		} else {
			var success2 = Query(l.E2(), vals)

			if success2.IsSome() {
				return success2
			} else {
				return val.NewArray(nil)
			}
		}
	default:
		panic(fmt.Sprintf("Unhandled PropExpr: %#v.", l))
	}
}

func evalValExpr(expr valexpr.ValExpr, vals val.Array) *val.Val {
	switch expr := expr.(type) {
	case valexpr.GetVar:
		return &vals.Vals()[expr.VarNum()]
	case valexpr.IntLit:
		var v val.Val = val.NewInt(expr.Data())
		return &v
	default:
		panic(fmt.Sprintf("Unhandled ValExpr: %#v.", expr))
	}
}
