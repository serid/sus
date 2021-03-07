package val

import "fmt"

type Array []Val

func (vala Array) Clone() Array {
	newArray := make([]Val, len(vala))

	for i, value := range vala {
		newArray[i] = CloneNillable(value)
	}

	return newArray
}

func (vala Array) String() string {
	return fmt.Sprintf("Array%v", vala)
}

func ArrayCmp(a, b interface{}) bool {
	ra := a.(Array)
	rb := b.(Array)

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
