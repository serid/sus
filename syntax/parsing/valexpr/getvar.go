package valexpr

import (
	"sus/interp/bcinterp/bytecode"
)

type GetVar struct {
	varnum bytecode.VarNum
}

func (GetVar) tagValExpr() {}

func NewGetVar(varnum bytecode.VarNum) GetVar {
	return GetVar{varnum: varnum}
}

func (gv GetVar) VarNum() bytecode.VarNum {
	return gv.varnum
}
