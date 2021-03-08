package bytecode

type Op struct {
	OpCode OpCode
	Data   Data
}

type OpCode int

const (
	OpInvalid OpCode = iota
	OpCloneSolution
	OpUnify
	OpConjunctionPart1
	OpConjunctionPart2
	OpDisjunctionPart1
	OpDisjunctionPart2
	OpRuleCall
	OpAssign
	OpPutInt
	OpPutVarNum
	OpAdd
)

// Slot is a register containing a solution
// Slot Number
type SlotNum int

// Variable is a register containing a value
// Variable Number
type VarNum int

// id of a rule, mapped from a name
type RuleId int
