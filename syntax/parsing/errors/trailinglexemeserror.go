package errors

import "fmt"

type TrailingLexemes struct{}

func (TrailingLexemes) Error() string {
	return fmt.Sprintf("expected EOF, found lexemes")
}
