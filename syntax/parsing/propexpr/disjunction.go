package propexpr

type Disjunction struct {
	e1 PropExpr
	e2 PropExpr
}

func (Disjunction) tagPropExpr() {}

func NewDisjunction(e1, e2 PropExpr) Disjunction {
	return Disjunction{e1: e1, e2: e2}
}

func (dj Disjunction) E1() PropExpr {
	return dj.e1
}

func (dj Disjunction) E2() PropExpr {
	return dj.e2
}
