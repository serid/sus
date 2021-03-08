package lexeme

import (
	"fmt"
	"sus/interp/bcinterp/bytecode"
)

type Kind byte

const (
	KindInvalid Kind = iota
	KindPlus
	KindAsterisk
	KindComma // Comma is a binary operator creating a syntactic pair of values (a 2-tuple)
	KindEqual
	KindRuleCall
	KindConj // Conjunction
	KindDisj // Disjunction
	KindParenL
	KindParenR
	KindUnit
	KindInt
	KindAt
	KindIdent
	KindAny
)

type Lexeme struct {
	kind Kind
	data Data
}

func (l Lexeme) Kind() Kind {
	return l.kind
}

func (l Lexeme) Data() Data {
	return l.data
}

func Plus() Lexeme {
	return Lexeme{kind: KindPlus, data: nil}
}

func Asterisk() Lexeme {
	return Lexeme{kind: KindAsterisk, data: nil}
}

func Comma() Lexeme {
	return Lexeme{kind: KindComma, data: nil}
}

func Equal() Lexeme {
	return Lexeme{kind: KindEqual, data: nil}
}

func RuleCall() Lexeme {
	return Lexeme{kind: KindRuleCall, data: nil}
}

func Conj() Lexeme {
	return Lexeme{kind: KindConj, data: nil}
}

func Disj() Lexeme {
	return Lexeme{kind: KindDisj, data: nil}
}

func ParenL() Lexeme {
	return Lexeme{kind: KindParenL, data: nil}
}

func ParenR() Lexeme {
	return Lexeme{kind: KindParenR, data: nil}
}

func Unit() Lexeme {
	return Lexeme{kind: KindUnit, data: nil}
}

func Int(i int) Lexeme {
	return Lexeme{kind: KindInt, data: IntData{data: i}}
}

func At(i bytecode.VarNum) Lexeme {
	return Lexeme{kind: KindAt, data: AtData{data: i}}
}

func Ident(s string) Lexeme {
	return Lexeme{kind: KindIdent, data: IdentData{data: s}}
}

func Any() Lexeme {
	return Lexeme{kind: KindAny, data: nil}
}

func (l Lexeme) IsOperatorLexeme() bool {
	switch l.kind {
	case KindUnit, KindInt, KindAt, KindIdent:
		return false
	case KindPlus, KindAsterisk, KindComma, KindEqual, KindRuleCall, KindConj, KindDisj:
		return true
	case KindParenL, KindParenR:
		//panic("Attempted to query if ParenL or ParenR is an operator.")
		return false
	default:
		panic(fmt.Sprintf("Unhandled Lexeme: %#v.", l))
	}
}

// Checks if a lexeme can participate in formation of borders of a ValExpr
func (l Lexeme) IsValBorderLexeme() bool {
	if l.IsOperatorLexeme() {
		return false
	}

	switch l.kind {
	case KindParenL, KindParenR:
		return false
	case KindInt, KindAt, KindUnit, KindIdent:
		return true
	default:
		panic("Unhandled lexeme")
	}
}

type OperatorType byte

const (
	TypeVal OperatorType = iota
	TypeProp
	TypeRuleCall
)

// Val Operator's arguments are ValExprs
// Prop Operator's arguments are PropExprs
// RuleCall Operator's arguments are Ident and tuple-linked-list
func (l Lexeme) IsValOrPropOperatorLexeme() OperatorType {
	if !l.IsOperatorLexeme() {
		panic("Attempted to query if a non-operator lexeme is a val or prop operator.")
	}

	switch l.kind {
	case KindPlus, KindAsterisk, KindComma, KindEqual:
		return TypeVal
	case KindConj, KindDisj:
		return TypeProp
	case KindRuleCall:
		return TypeRuleCall
	default:
		panic("Unreachable.")
	}
}

func CompareLexemeSlices(a, b []Lexeme) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
