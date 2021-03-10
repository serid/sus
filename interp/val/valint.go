package val

import "strconv"

type Int struct {
	Value int
}

func (Int) tagVal() {}

func NewInt(value int) Int {
	return Int{Value: value}
}

func (i Int) Clone() Val {
	// Copies struct byte-wise
	return i
}

func (i Int) String() string {
	return strconv.Itoa(i.Value)
}
