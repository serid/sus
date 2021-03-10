package propexpr

import "sus/syntax/parsing/valexpr"

type Unification struct {
	E1 valexpr.ValExpr
	E2 valexpr.ValExpr
}

func (Unification) tagPropExpr() {}

func NewUnification(e1, e2 valexpr.ValExpr) Unification {
	return Unification{E1: e1, E2: e2}
}
