package propexpr

type Conjunction struct {
	e1 PropExpr
	e2 PropExpr
}

func (Conjunction) tagPropExpr() {}

func NewConjunction(e1, e2 PropExpr) Conjunction {
	return Conjunction{e1: e1, e2: e2}
}

func (cj Conjunction) E1() PropExpr {
	return cj.e1
}

func (cj Conjunction) E2() PropExpr {
	return cj.e2
}
