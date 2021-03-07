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
func Query(propExpr propexpr.PropExpr, vals Solution) Solution {
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
			return nil
		}
	case propexpr.Conjunction:
		var success1 = Query(l.E1(), vals)

		if success1 != nil {
			var success2 = Query(l.E2(), success1)

			if success2 != nil {
				return success2
			} else {
				return nil
			}
		} else {
			return nil
		}
	case propexpr.Disjunction:
		var success1 = Query(l.E1(), vals.Clone())

		if success1 != nil {
			return success1
		} else {
			var success2 = Query(l.E2(), vals)

			if success2 != nil {
				return success2
			} else {
				return nil
			}
		}
	default:
		panic(fmt.Sprintf("Unhandled PropExpr: %#v.", l))
	}
}

// If expression is GetVar (@0), returns an lvalue (pointer to an nil value) that can be assigned
func evalValExpr(expr valexpr.ValExpr, vals Solution) *val.Val {
	switch expr := expr.(type) {
	case valexpr.GetVar:
		return &vals[expr.VarNum()]
	case valexpr.IntLit:
		var v val.Val = val.NewInt(expr.Data())
		return &v
	case valexpr.Plus:
		var v1 = *evalValExpr(expr.E1(), vals)
		if v1 == nil {
			panic("using unset variables in arithmetic expressions is unsupported")
		}
		var vi1 = v1.(val.Int)

		var v2 = *evalValExpr(expr.E2(), vals)
		if v2 == nil {
			panic("using unset variables in arithmetic expressions is unsupported")
		}
		var vi2 = v2.(val.Int)

		var v val.Val = val.NewInt(vi1.Value() + vi2.Value())
		return &v
	case valexpr.Mul:
		var v1 = *evalValExpr(expr.E1(), vals)
		if v1 == nil {
			panic("using unset variables in arithmetic expressions is unsupported")
		}
		var vi1 = v1.(val.Int)

		var v2 = *evalValExpr(expr.E2(), vals)
		if v2 == nil {
			panic("using unset variables in arithmetic expressions is unsupported")
		}
		var vi2 = v2.(val.Int)

		var v val.Val = val.NewInt(vi1.Value() * vi2.Value())
		return &v
	default:
		panic(fmt.Sprintf("Unhandled ValExpr: %#v.", expr))
	}
}
