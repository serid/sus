package propexpr

import (
	"sus/syntax/parsing/valexpr"
)

// id of a rule, alternative to name
type RuleId int

type RuleCall struct {
	rid  RuleId
	args []valexpr.ValExpr
}

func (RuleCall) tagPropExpr() {}

func NewRuleCall(rid RuleId, args []valexpr.ValExpr) RuleCall {
	return RuleCall{rid: rid, args: args}
}

func (rc RuleCall) Rid() RuleId {
	return rc.rid
}

func (rc RuleCall) Args() []valexpr.ValExpr {
	return rc.args
}
