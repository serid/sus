package valexpr

type Plus struct {
	e1 ValExpr
	e2 ValExpr
}

func (Plus) tagValExpr() {}

func NewPlus(e1, e2 ValExpr) Plus {
	return Plus{e1: e1, e2: e2}
}

func (pl Plus) E1() ValExpr {
	return pl.e1
}

func (pl Plus) E2() ValExpr {
	return pl.e2
}
