package valexpr

type IntLit struct {
	data int
}

func (IntLit) tagValExpr() {}

func NewIntLit(data int) IntLit {
	return IntLit{data: data}
}

func (il IntLit) Data() int {
	return il.data
}
