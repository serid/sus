package lexeme

import (
	"sus/interp/bcinterp/bytecode"
)

// Data structures for lexemes like Int or At

type Data interface {
	tagData()
}

type IntData struct {
	Data int
}

func (IntData) tagData() {}

type AtData struct {
	Data bytecode.VarNum
}

func (AtData) tagData() {}

type IdentData struct {
	Data string
}

func (IdentData) tagData() {}
