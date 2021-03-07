package valexpr

type Mul struct {
	e1 ValExpr
	e2 ValExpr
}

func (Mul) tagValExpr() {}

func NewMul(e1, e2 ValExpr) Mul {
	return Mul{e1: e1, e2: e2}
}

func (mul Mul) E1() ValExpr {
	return mul.e1
}

func (mul Mul) E2() ValExpr {
	return mul.e2
}
