package parsing

import (
	"sus/stuff"
	"sus/syntax/lexing"
	"sus/syntax/parsing/precedence"
)

type Parser struct {
	precedenceParser precedence.Parser
}

func DefaultParser() Parser {
	return Parser{precedenceParser: precedence.DefaultParser()}
}

func (parser Parser) ParseE(s string) (interface{}, error) {
	lexemes := lexing.Lexate(s)

	return parser.precedenceParser.ParseE(lexemes)
}

func (parser Parser) Parse(s string) interface{} {
	r, err := parser.ParseE(s)
	stuff.Unwrap(err)
	return r
}
