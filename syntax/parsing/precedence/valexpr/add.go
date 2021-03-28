package valexpr

type Plus struct {
	E1 ValExpr
	E2 ValExpr
}

func (Plus) tagValExpr() {}

func NewPlus(e1, e2 ValExpr) Plus {
	return Plus{E1: e1, E2: e2}
}
