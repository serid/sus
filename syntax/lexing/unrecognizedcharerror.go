package lexing

import "fmt"

type UnrecognizedCharacterError struct {
	c rune
}

func NewUnrecognizedCharacterError(c rune) UnrecognizedCharacterError {
	return UnrecognizedCharacterError{c: c}
}

func (uce UnrecognizedCharacterError) Error() string {
	return fmt.Sprintf("unrecognized character: %v", uce.c)
}
