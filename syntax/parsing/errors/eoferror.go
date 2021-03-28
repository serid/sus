package errors

import (
	"fmt"
	"sus/syntax/lexing/lexeme"
)

type Eof struct {
	lex lexeme.Lexeme
}

func NewEof(lex lexeme.Lexeme) Eof {
	return Eof{lex: lex}
}

func (ee Eof) Error() string {
	return fmt.Sprintf("expected lexeme (%v), found EOF", ee.lex)
}
