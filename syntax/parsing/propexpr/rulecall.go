package propexpr

import (
	"sus/interp/bcinterp/bytecode"
	"sus/syntax/parsing/valexpr"
)

type RuleCall struct {
	Rid  bytecode.RuleId
	Args []valexpr.ValExpr
}

func (RuleCall) tagPropExpr() {}

func NewRuleCall(rid bytecode.RuleId, args []valexpr.ValExpr) RuleCall {
	return RuleCall{Rid: rid, Args: args}
}
