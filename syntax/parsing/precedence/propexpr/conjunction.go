package propexpr

type Conjunction struct {
	E1 PropExpr
	E2 PropExpr
}

func (Conjunction) tagPropExpr() {}

func NewConjunction(e1, e2 PropExpr) Conjunction {
	return Conjunction{E1: e1, E2: e2}
}
