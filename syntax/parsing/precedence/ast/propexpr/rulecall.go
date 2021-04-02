package propexpr

import (
	"sus/syntax/parsing/precedence/ast/valexpr"
)

type RuleCall struct {
	Name string
	Args []valexpr.ValExpr
}

func (RuleCall) tagPropExpr() {}

func NewRuleCall(name string, args []valexpr.ValExpr) RuleCall {
	return RuleCall{Name: name, Args: args}
}
