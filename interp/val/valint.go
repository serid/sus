package val

import "strconv"

type Int struct {
	value int
}

func (Int) tagVal() {}

func NewInt(value int) Int {
	return Int{value: value}
}

func (i Int) Value() int {
	return i.value
}

func (i Int) Clone() Val {
	// Copies struct byte-wise
	return i
}

func (i Int) String() string {
	return strconv.Itoa(i.value)
}
