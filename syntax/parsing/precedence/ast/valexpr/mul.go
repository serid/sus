package valexpr

type Mul struct {
	E1 ValExpr
	E2 ValExpr
}

func (Mul) tagValExpr() {}

func NewMul(e1, e2 ValExpr) Mul {
	return Mul{E1: e1, E2: e2}
}
