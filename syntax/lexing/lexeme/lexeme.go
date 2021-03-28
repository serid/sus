package lexeme

import (
	"fmt"
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
	KindIdent
	KindAny
)

type Lexeme struct {
	Kind Kind
	Data Data
}

func Plus() Lexeme {
	return Lexeme{Kind: KindPlus, Data: nil}
}

func Asterisk() Lexeme {
	return Lexeme{Kind: KindAsterisk, Data: nil}
}

func Comma() Lexeme {
	return Lexeme{Kind: KindComma, Data: nil}
}

func Equal() Lexeme {
	return Lexeme{Kind: KindEqual, Data: nil}
}

func RuleCall() Lexeme {
	return Lexeme{Kind: KindRuleCall, Data: nil}
}

func Conj() Lexeme {
	return Lexeme{Kind: KindConj, Data: nil}
}

func Disj() Lexeme {
	return Lexeme{Kind: KindDisj, Data: nil}
}

func ParenL() Lexeme {
	return Lexeme{Kind: KindParenL, Data: nil}
}

func ParenR() Lexeme {
	return Lexeme{Kind: KindParenR, Data: nil}
}

func Unit() Lexeme {
	return Lexeme{Kind: KindUnit, Data: nil}
}

func Int(i int) Lexeme {
	return Lexeme{Kind: KindInt, Data: IntData{Data: i}}
}

func Ident(s string) Lexeme {
	return Lexeme{Kind: KindIdent, Data: IdentData{Data: s}}
}

func Any() Lexeme {
	return Lexeme{Kind: KindAny, Data: nil}
}

func (l Lexeme) IsOperatorLexeme() bool {
	switch l.Kind {
	case KindUnit, KindInt, KindIdent:
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

	switch l.Kind {
	case KindParenL, KindParenR:
		return false
	case KindInt, KindUnit, KindIdent:
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

	switch l.Kind {
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
