package bcinterp

import "sus/interp/bcinterp/bytecode"

// Each compilation function generates instructions that put values/solutions in variables/slots
// The register allocation algorithm keeps track of occupied registers by allcating in order 0,1,2... and saving
// number of next free register in compilerState
type compilerState struct {
	NextSlotNum bytecode.SlotNum
	NextVarNum  bytecode.VarNum
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
	return compilerState{
		NextSlotNum: cs.NextSlotNum + 1,
		NextVarNum:  cs.NextVarNum,
	}
}

func (cs compilerState) SkipVar() compilerState {
	return compilerState{
		NextSlotNum: cs.NextSlotNum,
		NextVarNum:  cs.NextVarNum + 1,
	}
}


