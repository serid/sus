package valexpr

import (
	"sus/interp/bcinterp/bytecode"
)

type GetVar struct {
	VarNum bytecode.VarNum
}

func (GetVar) tagValExpr() {}

func NewGetVar(varnum bytecode.VarNum) GetVar {
	return GetVar{VarNum: varnum}
}
