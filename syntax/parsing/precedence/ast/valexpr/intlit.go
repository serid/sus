package valexpr

type IntLit struct {
	Data int
}

func (IntLit) tagValExpr() {}

func NewIntLit(data int) IntLit {
	return IntLit{Data: data}
}
