package odesc

type OperatorDescription struct {
	precedenceLevel int
	associativityLR Associativity
}

func NewOperatorDescription(precedenceLevel int, associativityLR Associativity) OperatorDescription {
	return OperatorDescription{precedenceLevel: precedenceLevel, associativityLR: associativityLR}
}

func (od OperatorDescription) PrecedenceLevel() int {
	return od.precedenceLevel
}

func (od OperatorDescription) AssociativityLR() Associativity {
	return od.associativityLR
}

type Associativity byte

const (
	AssociativityInvalid Associativity = iota
	AssociativityLeft
	AssociativityRight
	// AssociativityAssociative -- unused
)
