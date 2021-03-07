package val

// Sum type for Values
// nil value means that a variable is not solved yet and can be filled in with any value later.
type Val interface {
	tagVal()

	// Should return value of type Self
	Clone() Val
}

func CloneNillable(value Val) Val {
	if value == nil {
		return nil
	} else {
		return value.Clone()
	}
}
