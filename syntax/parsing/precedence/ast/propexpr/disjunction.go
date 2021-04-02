package propexpr

type Disjunction struct {
	E1 PropExpr
	E2 PropExpr
}

func (Disjunction) tagPropExpr() {}

func NewDisjunction(e1, e2 PropExpr) Disjunction {
	return Disjunction{E1: e1, E2: e2}
}
