package val

import "fmt"

type Array struct {
	vals []Val
}

func NewArray(vals []Val) Array {
	return Array{vals: vals}
}

func (vala Array) Vals() []Val {
	return vala.vals
}

func (vala Array) IsNone() bool {
	return vala.vals == nil
}

func (vala Array) IsSome() bool {
	return vala.vals != nil
}

func (vala Array) Clone() Array {
	newArray := Array{vals: make([]Val, len(vala.vals))}

	for i, value := range vala.vals {
		newArray.vals[i] = CloneNillable(value)
	}

	return newArray
}

func (vala Array) String() string {
	return fmt.Sprintf("Array%v", vala.vals)
}

func ArrayCmp(a, b interface{}) bool {
	raa := a.(Array)
	rab := b.(Array)

	ra := raa.vals
	rb := rab.vals

	if (ra == nil) || (rb == nil) {
		panic("Attempted to comapre nil slices for equivalence.")
	}

	if len(ra) != len(rb) {
		return false
	}

	for i := range ra {
		if ra[i] != rb[i] {
			return false
		}
	}
	return true
}
