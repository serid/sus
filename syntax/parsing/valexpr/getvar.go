package valexpr

import (
	"sus/syntax/lexing/lexeme"
)

type GetVar struct {
	varnum lexeme.VarNum
}

func (GetVar) tagValExpr() {}

func NewGetVar(varnum lexeme.VarNum) GetVar {
	return GetVar{varnum: varnum}
}

func (gv GetVar) VarNum() lexeme.VarNum {
	return gv.varnum
}
