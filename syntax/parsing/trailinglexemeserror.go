package parsing

import "fmt"

type TrailingLexemesError struct{}

func (TrailingLexemesError) Error() string {
	return fmt.Sprintf("expected EOF, found lexemes")
}
