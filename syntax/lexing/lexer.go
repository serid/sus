package lexing

import (
	"sus/interp/bcinterp/bytecode"
	"sus/stuff"
	"sus/syntax/lexing/lexeme"
	"unicode"
)

func LexateE(s1 string) ([]lexeme.Lexeme, error) {
	var result []lexeme.Lexeme = nil

	// Convert string to array of runes
	s := []rune(s1)

	i := 0

	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else if s[i] == '@' {
			i++
			if unicode.IsDigit(s[i]) {
				// Parse an integer
				n := 0
				for i < len(s) && unicode.IsDigit(s[i]) {
					n *= 10
					n += int(s[i]) - '0'
					i++
				}
				result = append(result, lexeme.At(bytecode.VarNum(n)))
			} else {
				panic("'@' should be followed by a number")
			}
		} else if s[i] == '+' {
			result = append(result, lexeme.Plus())
			i++
		} else if s[i] == '*' {
			result = append(result, lexeme.Asterisk())
			i++
		} else if s[i] == ',' {
			result = append(result, lexeme.Comma())
			i++
		} else if PrefixRunes([]rune("()"), s[i:]) {
			result = append(result, lexeme.Unit())
			i += len("()")
		} else if PrefixRunes([]rune("/\\"), s[i:]) {
			result = append(result, lexeme.Conj())
			i += len("/\\")
		} else if PrefixRunes([]rune("\\/"), s[i:]) {
			result = append(result, lexeme.Disj())
			i += len("\\/")
		} else if s[i] == '=' {
			result = append(result, lexeme.Equal())
			i++
		} else if s[i] == '(' {
			result = append(result, lexeme.ParenL())
			i++
		} else if s[i] == ')' {
			result = append(result, lexeme.ParenR())
			i++
		} else if unicode.IsDigit(s[i]) {
			// Parse an integer
			n := 0
			for i < len(s) && unicode.IsDigit(s[i]) {
				n *= 10
				n += int(s[i]) - '0'
				i++
			}
			result = append(result, lexeme.Int(n))
		} else if unicode.IsLetter(s[i]) {
			// Parse an ident
			ident := make([]rune, 0)
			for i < len(s) && unicode.IsLetter(s[i]) {
				ident = append(ident, s[i])
				i++
			}
			result = append(result, lexeme.Ident(string(ident)))
		} else {
			return nil, NewUnrecognizedCharacterError(s[i])
		}

		// After pushing a lexeme, insert a `RuleCall` between last two lexemes if they are value lexemes
		if len(result)-2 >= 0 {
			if result[len(result)-2].IsValBorderLexeme() || result[len(result)-2].Kind() == lexeme.KindParenR {
				if result[len(result)-1].IsValBorderLexeme() || result[len(result)-1].Kind() == lexeme.KindParenL {
					// Insert an element at `index`
					index := len(result) - 1
					result = append(result[:index+1], result[index:]...)
					result[index] = lexeme.RuleCall()
				}
			}
		}
	}

	return result, nil
}

func Lexate(s string) []lexeme.Lexeme {
	r, err := LexateE(s)
	stuff.Unwrap(err)
	return r
}

// Checks if array `prefix` is a prefix of array `s`
func PrefixRunes(prefix []rune, s []rune) bool {
	if len(prefix) > len(s) {
		return false
	}

	for i := range prefix {
		if prefix[i] != s[i] {
			return false
		}
	}
	return true
}
