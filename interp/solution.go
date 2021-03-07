package interp

import (
	"fmt"
	"sus/interp/val"
)

type Solution []val.Val

func (sol Solution) Clone() Solution {
	newSolution := make(Solution, len(sol))

	for i, value := range sol {
		newSolution[i] = val.CloneNillable(value)
	}

	return newSolution
}

func (sol Solution) String() string {
	return fmt.Sprintf("Solution%v", sol)
}

func ArrayCmp(a, b interface{}) bool {
	ra := a.(Solution)
	rb := b.(Solution)

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
