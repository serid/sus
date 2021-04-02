package osi

import (
	"sus/syntax/parsing/precedence/ast/propexpr"
	"sus/syntax/parsing/precedence/ast/valexpr"
)

// `data` may be one of
// `valexpr.ValExpr` for a value expression
// `propexpr.PropExpr` for a prop expression
// `valexpr.CommaListNode` for a CommaListNode // It has a separate kind to differentiate it from other valexpr-s
type OutStackItem struct {
	Kind Kind
	Data interface{}
}

func ValExpr(expr valexpr.ValExpr) OutStackItem {
	return OutStackItem{
		Kind: KindValExpr,
		Data: expr,
	}
}

func PropExpr(expr propexpr.PropExpr) OutStackItem {
	return OutStackItem{
		Kind: KindPropExpr,
		Data: expr,
	}
}

func CommaListNode(node valexpr.CommaListNode) OutStackItem {
	return OutStackItem{
		Kind: KindCommaListNode,
		Data: node,
	}
}

type Kind byte

const (
	KindValExpr Kind = iota
	KindPropExpr
	KindCommaListNode
)
