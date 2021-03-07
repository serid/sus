package valexpr

type CommaListNode interface {
	ValExpr
	tagCommaListNode()
}

// `e2` will always be a `CommaListPair` value or a Unit
// This way CommaPairs form a linked list where an element is stored in e1 and a link to the next node is stored in e2
type CommaListPair struct {
	v    ValExpr
	tail CommaListNode
}

func (CommaListPair) tagValExpr() {}

func (CommaListPair) tagCommaListNode() {}

func NewCommaPair(v ValExpr, tail CommaListNode) CommaListPair {
	return CommaListPair{v: v, tail: tail}
}

func (clp CommaListPair) V() ValExpr {
	return clp.v
}

func (clp CommaListPair) Tail() CommaListNode {
	return clp.tail
}

type Unit struct{}

func (Unit) tagValExpr() {}

func (Unit) tagCommaListNode() {}

// Converts
// a to [a]
// (a, b) to [a, b]
// and
// (a, (b, (c, d))) to [a, b, c, d]
func NestedPairsToSliceOfValExpr(expr CommaListNode) []ValExpr {
	var flattenedSlice = make([]ValExpr, 0)

loop:
	for {
		switch v := expr.(type) {
		case Unit:
			break loop
		case CommaListPair:
			flattenedSlice = append(flattenedSlice, v.V())
			expr = v.Tail()
		default:
			panic("Unreachable.")
		}
	}

	return flattenedSlice
}
