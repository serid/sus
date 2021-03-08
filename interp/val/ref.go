package val

import "sus/interp/bcinterp/bytecode"

type Ref struct {
	Value bytecode.VarNum
}

func (Ref) tagVal() {}

func (i Ref) Clone() Val {
	// Copies struct byte-wise
	return i
}
