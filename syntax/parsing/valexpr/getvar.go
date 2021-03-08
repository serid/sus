package valexpr

import (
	"sus/types"
)

type GetVar struct {
	varnum types.VarNum
}

func (GetVar) tagValExpr() {}

func NewGetVar(varnum types.VarNum) GetVar {
	return GetVar{varnum: varnum}
}

func (gv GetVar) VarNum() types.VarNum {
	return gv.varnum
}
