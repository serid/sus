package parsing

import (
	"fmt"
	"sus/syntax/lexing/lexeme"
)

type EOFError struct {
	lex lexeme.Lexeme
}

func NewEOFError(lex lexeme.Lexeme) EOFError {
	return EOFError{lex: lex}
}

func (ee EOFError) Error() string {
	return fmt.Sprintf("expected lexeme (%v), found EOF", ee.lex)
}
