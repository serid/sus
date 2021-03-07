package val

// Sum type for Values
type Val interface {
	tagVal()

	// Should return value of type Self
	Clone() Val
}
