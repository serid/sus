package osi

import (
	"sus/syntax/lexing/lexeme"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// `data` may be one of
// `valexpr.ValExpr` for a value expression
// `propexpr.PropExpr` for a prop expression
// `lexeme.IdentData` for an ident
// `valexpr.CommaListNode` for a CommaListNode // It has a separate kind to differentiate it from other valexpr-s
type OutStackItem struct {
	kind Kind
	data interface{}
}

func (osi OutStackItem) Kind() Kind {
	return osi.kind
}

func (osi OutStackItem) Data() interface{} {
	return osi.data
}

func ValExpr(expr valexpr.ValExpr) OutStackItem {
	return OutStackItem{
		kind: KindValExpr,
		data: expr,
	}
}

func PropExpr(expr propexpr.PropExpr) OutStackItem {
	return OutStackItem{
		kind: KindPropExpr,
		data: expr,
	}
}

func Ident(identData lexeme.IdentData) OutStackItem {
	return OutStackItem{
		kind: KindIdent,
		data: identData,
	}
}

func CommaListNode(node valexpr.CommaListNode) OutStackItem {
	return OutStackItem{
		kind: KindCommaListNode,
		data: node,
	}
}

type Kind byte

const (
	KindValExpr Kind = iota
	KindPropExpr
	KindIdent
	KindCommaListNode
)
