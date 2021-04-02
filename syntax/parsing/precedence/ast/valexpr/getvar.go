package valexpr

type GetVar struct {
	Name string
}

func (GetVar) tagValExpr() {}

func NewGetVar(varnum string) GetVar {
	return GetVar{Name: varnum}
}
