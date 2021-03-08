package bytecode

import (
	"sus/types"
)

// Data for Operations operating on solutions

type SolCloneSolutionData struct {
	Input SlotNum
	Output SlotNum
}

func (SolCloneSolutionData) tagData() {}

type SolUnifyData struct {
	Context SlotNum
	Input1 VarNum
	Input2 VarNum
	Output SlotNum
}

func (SolUnifyData) tagData() {}

// The `OpConjunctionPart1` and `OpDisjunctionPart1` are opcodes inserted after calculation of first operand and
// before calculation of second operand. The may decide to skip calculation of second operand if information from
// first operand is sufficient.
// `OpConjunctionPart1` will skip calculation of second operand if first operand turned out to be nil.
// `OpDisjunctionPart1` will skip calculation of second operand if first operand turned out to be not nil. (until
// multisolution return is implemented.)

// `SkipAddress` is index in array of op-s that interpreter should jump to if it wants to skip calculating second operand
// it is typically the adress of `Op(Con|Dis)junctionPart2`.
//
// Example illustraion of opcodes in array
// 1) *calulation of first operand*
// 10) `OpDisjunctionPart1` SkipAddress = 30
// 11) *calulation of first operand*
// 30) `OpDisjunctionPart2`
type SolMaybeSkipData struct {
	Input1      SlotNum
	SkipAddress BodyAddress
}

func (SolMaybeSkipData) tagData() {}

type BodyAddress int

// SolData for binary proposition (solution) operations that take input from two slots and put output into a third slot
type SolBinOpData struct {
	Input1 SlotNum
	Input2 SlotNum
	Output SlotNum
}

func (SolBinOpData) tagData() {}

type SolRuleCallData struct {
	Context SlotNum
	Rid    types.RuleId
	Input  []VarNum
	Output SlotNum
}

func (SolRuleCallData) tagData() {}
