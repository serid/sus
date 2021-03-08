package lexeme

import (
	"sus/interp/bcinterp/bytecode"
)

// Data structures for lexemes like Int or At

type Data interface {
	tagData()
}

type IntData struct {
	data int
}

func (IntData) tagData() {}

func (i IntData) Data() int {
	return i.data
}

type AtData struct {
	data bytecode.VarNum
}

func (AtData) tagData() {}

func (at AtData) Data() bytecode.VarNum {
	return at.data
}

type IdentData struct {
	data string
}

func (IdentData) tagData() {}

func (ident IdentData) Data() string {
	return ident.data
}
