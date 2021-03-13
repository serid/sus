package parsing

import (
	"fmt"
	"sus/stuff"
	"sus/syntax/lexing"
	"sus/syntax/lexing/lexeme"
	"sus/syntax/parsing/odesc"
	"sus/syntax/parsing/osi"
	"sus/syntax/parsing/propexpr"
	"sus/syntax/parsing/valexpr"
)

// A shunting-yard operator-precedence parser

type OutStack []osi.OutStackItem

type OperatorStackItem lexeme.Lexeme

type OperatorStack []OperatorStackItem

// Parse returns interface{} and not concrete AST node type for testing purposes
func (parser Parser) ParseE(s string) (interface{}, error) {
	lexemes := lexing.Lexate(s)

	var outStack OutStack
	var operatorStack OperatorStack

	for _, l := range lexemes {
		if l.Kind == lexeme.KindParenL {
			// If Lexeme is ParenL, then Push it to stack
			operatorStack = append(operatorStack, OperatorStackItem(l))
		} else if l.Kind == lexeme.KindParenR {
			// If Lexeme is ParenR, then Reduce the stack until a ParenL is found
			err := reduceUntil(&operatorStack, &outStack, func(tos OperatorStackItem) bool {
				return lexeme.Lexeme(tos).Kind == lexeme.KindParenL
			}, func() {
				panic("No matching ParenL found for ParenR.")
			})

			if err != nil {
				return nil, err
			}

			// Pop off the ParenL
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else if l.IsOperatorLexeme() {
			// Reduce until `operatorStack`.TOS is a ParenL
			// or precedence level of `operatorStack`.TOS is lower than level of `l`

			err := reduceUntil(&operatorStack, &outStack, func(tos OperatorStackItem) bool {
				if lexeme.Lexeme(tos).Kind == lexeme.KindParenL {
					return true
				}

				lPrecedence := parser.OperatorDesciption(l).PrecedenceLevel
				lAssociativity := parser.OperatorDesciption(l).AssociativityLR
				tosPrecedence := parser.OperatorDesciption(lexeme.Lexeme(tos)).PrecedenceLevel

				continueEh := tosPrecedence > lPrecedence || (tosPrecedence == lPrecedence && lAssociativity == odesc.AssociativityLeft)

				return !continueEh
			}, func() {
			})

			if err != nil {
				return nil, err
			}

			// Push new operator on stack
			operatorStack = append(operatorStack, OperatorStackItem(l))
		} else {
			shift(l, &outStack)
		}
	}

	// Reduce remanining operators on `operatorStack`
	err := reduceUntil(&operatorStack, &outStack, func(tos OperatorStackItem) bool {
		return false
	}, func() {
	})

	if err != nil {
		return nil, err
	}

	if len(outStack) > 1 {
		return nil, TrailingLexemesError{}
		//panic(fmt.Sprintf("Leftover stack items after parsing: %#v.", outStack))
	}

	//result := outStack[0].Data.(propexpr.PropExpr)
	result := outStack[0].Data

	return result, nil
}

func (parser Parser) Parse(s string) interface{} {
	r, err := parser.ParseE(s)
	stuff.Unwrap(err)
	return r
}

func reduceUntil(operatorStack *OperatorStack, outStack *OutStack, f func(tos OperatorStackItem) bool, onEmptyStack func()) error {
	for {
		if len(*operatorStack) == 0 {
			onEmptyStack()
			break
		}
		if f((*operatorStack)[len(*operatorStack)-1]) {
			break
		}

		poppedOperator := (*operatorStack)[len(*operatorStack)-1]
		*operatorStack = (*operatorStack)[:len(*operatorStack)-1]

		err := reduce(lexeme.Lexeme(poppedOperator), outStack)
		if err != nil {
			return err
		}
	}

	return nil
}

func shift(lex lexeme.Lexeme, outStack *OutStack) {
	switch lex.Kind {
	case lexeme.KindUnit:
		*outStack = append(*outStack, osi.ValExpr(valexpr.Unit{}))
	case lexeme.KindInt:
		n := lex.Data.(lexeme.IntData)
		*outStack = append(*outStack, osi.ValExpr(valexpr.NewIntLit(n.Data)))
	case lexeme.KindAt:
		n := lex.Data.(lexeme.AtData)
		*outStack = append(*outStack, osi.ValExpr(valexpr.NewGetVar(n.Data)))
	case lexeme.KindIdent:
		identData := lex.Data.(lexeme.IdentData)
		*outStack = append(*outStack, osi.Ident(identData))
	default:
		panic(fmt.Sprintf("Unhandled Lexeme: %#v.", lex))
	}
}

func reduce(oper lexeme.Lexeme, outStack *OutStack) error {
	if len(*outStack) >= 2 {
		tos1 := (*outStack)[len(*outStack)-1]
		*outStack = (*outStack)[:len(*outStack)-1]
		tos2 := (*outStack)[len(*outStack)-1]
		*outStack = (*outStack)[:len(*outStack)-1]

		switch oper.IsValOrPropOperatorLexeme() {
		case lexeme.TypeVal:
			e1 := tos1.Data.(valexpr.ValExpr)
			e2 := tos2.Data.(valexpr.ValExpr)

			switch oper.Kind {
			case lexeme.KindPlus:
				*outStack = append(*outStack, osi.ValExpr(valexpr.NewPlus(e2, e1)))
			case lexeme.KindAsterisk:
				*outStack = append(*outStack, osi.ValExpr(valexpr.NewMul(e2, e1)))
			case lexeme.KindComma:
				// Combine 2 valexprs into a linked-list node

				var nodeToPush valexpr.CommaListPair

				v, ok := e1.(valexpr.CommaListNode)

				if ok {
					// If right expr is a node (unit or pair), make a simple node
					nodeToPush = valexpr.NewCommaPair(e2, v)
				} else {
					// Otherwise make a new node
					newNode := valexpr.NewCommaPair(e1, valexpr.Unit{})
					nodeToPush = valexpr.NewCommaPair(e2, newNode)
				}
				*outStack = append(*outStack, osi.CommaListNode(nodeToPush))
			case lexeme.KindEqual:
				*outStack = append(*outStack, osi.PropExpr(propexpr.NewUnification(e2, e1)))
			default:
				panic(fmt.Sprintf("Unsupported binary operator: %#v.", oper))
			}
		case lexeme.TypeProp:
			e1 := tos1.Data.(propexpr.PropExpr)
			e2 := tos2.Data.(propexpr.PropExpr)

			switch oper.Kind {
			case lexeme.KindConj:
				*outStack = append(*outStack, osi.PropExpr(propexpr.NewConjunction(e2, e1)))
			case lexeme.KindDisj:
				*outStack = append(*outStack, osi.PropExpr(propexpr.NewDisjunction(e2, e1)))
			default:
				panic(fmt.Sprintf("Unsupported binary operator: %#v.", oper))
			}
		case lexeme.TypeRuleCall:
			var args valexpr.CommaListNode

			// If call argument is a CommaListNode, just use it
			// If call argument is a single valexpr, convert it to a CommaListNode with one element
			if e, ok := tos1.Data.(valexpr.CommaListNode); ok {
				args = e
			} else if e, ok := tos1.Data.(valexpr.ValExpr); ok {
				args = valexpr.NewCommaPair(e, valexpr.Unit{})
			} else {
				panic("Type cast error")
			}

			ident := tos2.Data.(lexeme.IdentData)
			_ = ident // TODO: convert function name (symbol) to an RuleId
			*outStack = append(*outStack, osi.PropExpr(propexpr.NewRuleCall(1, valexpr.NestedPairsToSliceOfValExpr(args))))
		default:
			panic("Unreachable.")
		}
		return nil
	} else if len(*outStack) == 1 {
		// Reached EOF
		// TODO: provide more useful lexeme for an error message
		return NewEOFError(lexeme.Any())
	} else {
		// len(*outStack) == 0
		// Reached EOF
		return NewEOFError(lexeme.Any())
	}
	//panic(fmt.Sprintf("Could not reduce stack. Stack length: %d.", len(*outStack)))
}

type Parser struct {
	operatorDescriptionTable map[lexeme.Kind]odesc.OperatorDescription
}

func DefaultParser() Parser {
	operatorDescriptionTable := map[lexeme.Kind]odesc.OperatorDescription{
		lexeme.KindDisj:     odesc.NewOperatorDescription(1, odesc.AssociativityLeft),
		lexeme.KindConj:     odesc.NewOperatorDescription(2, odesc.AssociativityLeft),
		lexeme.KindEqual:    odesc.NewOperatorDescription(3, odesc.AssociativityLeft),
		lexeme.KindComma:    odesc.NewOperatorDescription(4, odesc.AssociativityRight),
		lexeme.KindPlus:     odesc.NewOperatorDescription(5, odesc.AssociativityLeft),
		lexeme.KindAsterisk: odesc.NewOperatorDescription(6, odesc.AssociativityLeft),
		lexeme.KindRuleCall: odesc.NewOperatorDescription(7, odesc.AssociativityLeft),
	}
	return Parser{operatorDescriptionTable: operatorDescriptionTable}
}

func (parser Parser) OperatorDesciption(lex lexeme.Lexeme) odesc.OperatorDescription {
	v, ok := parser.operatorDescriptionTable[lex.Kind]
	if !ok {
		panic(fmt.Sprintf("Operator description for a lexeme (%#v) is unknown.", lex))
	}
	return v
}
