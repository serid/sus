package bytecode

func Assign(context SlotNum, input, output VarNum) Op {
	return Op{OpCode: OpAssign, Data: AssignData{Context: context, Input: input, Output: output}}
}

func PutInt(context SlotNum, data int, output VarNum) Op {
	return Op{OpCode: OpPutInt, Data: PutIntData{Context: context, Data: data, Output: output}}
}

func PutVarNum(context SlotNum, data VarNum, output VarNum) Op {
	return Op{OpCode: OpPutVarNum, Data: PutVarNumData{Context: context, Data: data, Output: output}}
}

func Add(context SlotNum, input1, input2, output VarNum) Op {
	return Op{OpCode: OpAdd, Data: ValBinOpData{
		Context: context,
		Input1:  input1,
		Input2:  input2,
		Output:  output,
	}}
}
