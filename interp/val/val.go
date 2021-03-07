package val

// Sum type for Values
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
