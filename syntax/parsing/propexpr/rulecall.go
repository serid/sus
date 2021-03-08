package propexpr

import (
	"sus/interp/bcinterp/bytecode"
	"sus/syntax/parsing/valexpr"
)

type RuleCall struct {
	rid  bytecode.RuleId
	args []valexpr.ValExpr
}

func (RuleCall) tagPropExpr() {}

func NewRuleCall(rid bytecode.RuleId, args []valexpr.ValExpr) RuleCall {
	return RuleCall{rid: rid, args: args}
}

func (rc RuleCall) Rid() bytecode.RuleId {
	return rc.rid
}

func (rc RuleCall) Args() []valexpr.ValExpr {
	return rc.args
}
