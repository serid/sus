package bytecode

func CloneSolution(input, output SlotNum) Op {
	return Op{OpCode: OpCloneSolution, Data: SolCloneSolutionData{
		Input:  input,
		Output: output,
	}}
}

func Unify(context SlotNum, input1, input2 VarNum, output SlotNum) Op {
	return Op{OpCode: OpUnify, Data: SolUnifyData{
		Context: context,
		Input1:  input1,
		Input2:  input2,
		Output:  output,
	}}
}

func ConjunctionPart1(input1 SlotNum, skipAddress BodyAddress) Op {
	return Op{OpCode: OpConjunctionPart1, Data: SolMaybeSkipData{
		Input1:      input1,
		SkipAddress: skipAddress,
	}}
}

func ConjunctionPart2(input1, input2, output SlotNum) Op {
	return Op{OpCode: OpConjunctionPart2, Data: SolBinOpData{
		Input1: input1,
		Input2: input2,
		Output: output,
	}}
}

func DisjunctionPart1(input1 SlotNum, skipAddress BodyAddress) Op {
	return Op{OpCode: OpDisjunctionPart1, Data: SolMaybeSkipData{
		Input1:      input1,
		SkipAddress: skipAddress,
	}}
}

func DisjunctionPart2(input1, input2, output SlotNum) Op {
	return Op{OpCode: OpDisjunctionPart2, Data: SolBinOpData{
		Input1: input1,
		Input2: input2,
		Output: output,
	}}
}

func RuleCall(context SlotNum, rid RuleId, input []VarNum, output SlotNum) Op {
	return Op{OpCode: OpRuleCall, Data: SolRuleCallData{
		Context: context,
		Rid:     rid,
		Input:   input,
		Output:  output,
	}}
}
