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

func Ident(identData lexeme.IdentData) OutStackItem {
	return OutStackItem{
		Kind: KindIdent,
		Data: identData,
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
	KindIdent
	KindCommaListNode
)
