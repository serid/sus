package odesc

type OperatorDescription struct {
	PrecedenceLevel int
	AssociativityLR Associativity
}

func NewOperatorDescription(precedenceLevel int, associativityLR Associativity) OperatorDescription {
	return OperatorDescription{PrecedenceLevel: precedenceLevel, AssociativityLR: associativityLR}
}

type Associativity byte

const (
	AssociativityInvalid Associativity = iota
	AssociativityLeft
	AssociativityRight
	// AssociativityAssociative -- unused
)
