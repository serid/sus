package bytecode

type AssignData struct {
	Context SlotNum
	Input   VarNum
	Output  VarNum
}

func (AssignData) tagData() {}

type PutIntData struct {
	Context SlotNum
	Data    int
	Output  VarNum
}

func (PutIntData) tagData() {}

type PutVarNumData struct {
	Context SlotNum
	Data    VarNum
	Output  VarNum
}

func (PutVarNumData) tagData() {}

// ValData for simple binary operations that take input from two variables and put output into a third variable
type ValBinOpData struct {
	Context SlotNum
	Input1  VarNum
	Input2  VarNum
	Output  VarNum
}

func (ValBinOpData) tagData() {}
