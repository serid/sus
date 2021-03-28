package bcinterp

import "sus/interp/bcinterp/bytecode"

// Each compilation function generates instructions that put values/solutions in variables/slots
// The register allocation algorithm keeps track of occupied registers by allcating in order 0,1,2... and saving
// number of next free register in compilerState
type compilerState struct {
	NextSlotNum bytecode.SlotNum
	NextVarNum  bytecode.VarNum

	// VarTable keeps track of where a named variable is stored
	VarTable map[string]bytecode.VarNum
}

// By convention operation result is stored in last occupied register
func (cs compilerState) SlotResult() bytecode.SlotNum {
	return cs.NextSlotNum - 1
}

// By convention operation result is stored in last occupied register
func (cs compilerState) VarResult() bytecode.VarNum {
	return cs.NextVarNum - 1
}

func (cs compilerState) SkipSol() compilerState {
	cs.NextSlotNum += 1
	return cs
}

func (cs compilerState) SkipVar() compilerState {
	cs.NextVarNum += 1
	return cs
}

//goland:noinspection GoAssignmentToReceiver
func (cs compilerState) GetNamedVar(Name string) (bytecode.VarNum, compilerState) {
	if vn, ok := cs.VarTable[Name]; ok {
		return vn, cs
	} else {
		// Allocate a VarNum for Name
		newVarNum := cs.NextVarNum
		cs = cs.SkipVar()

		newVarTable := make(map[string]bytecode.VarNum, len(cs.VarTable)+1)
		for k, v := range cs.VarTable {
			newVarTable[k] = v
		}

		newVarTable[Name] = newVarNum

		cs.VarTable = newVarTable

		return newVarNum, cs
	}
}
