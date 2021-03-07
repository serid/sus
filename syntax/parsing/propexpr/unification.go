package propexpr

import "sus/syntax/parsing/valexpr"

type Unification struct {
	e1 valexpr.ValExpr
	e2 valexpr.ValExpr
}

func (Unification) tagPropExpr() {}

func NewUnification(e1, e2 valexpr.ValExpr) Unification {
	return Unification{e1: e1, e2: e2}
}

func (u Unification) E1() valexpr.ValExpr {
	return u.e1
}

func (u Unification) E2() valexpr.ValExpr {
	return u.e2
}
