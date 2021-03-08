package propexpr

import (
	"sus/syntax/parsing/valexpr"
	"sus/types"
)

type RuleCall struct {
	rid  types.RuleId
	args []valexpr.ValExpr
}

func (RuleCall) tagPropExpr() {}

func NewRuleCall(rid types.RuleId, args []valexpr.ValExpr) RuleCall {
	return RuleCall{rid: rid, args: args}
}

func (rc RuleCall) Rid() types.RuleId {
	return rc.rid
}

func (rc RuleCall) Args() []valexpr.ValExpr {
	return rc.args
}
