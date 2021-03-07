package val

import "fmt"

type OptArray struct {
	options []Option
}

func NewOptArray(options []Option) OptArray {
	return OptArray{options: options}
}

func (opt OptArray) Options() []Option {
	return opt.options
}

func (opt OptArray) IsNone() bool {
	return opt.options == nil
}

func (opt OptArray) IsSome() bool {
	return opt.options != nil
}

func (opt OptArray) Clone() OptArray {
	newOptArray := OptArray{options: make([]Option, len(opt.options))}

	for i, option := range opt.options {
		newOptArray.options[i] = option.Clone()
	}

	return newOptArray
}

func (opt OptArray) String() string {
	return fmt.Sprintf("OptArray%v", opt.options)
}

func OptArrayCmp(a, b interface{}) bool {
	raa := a.(OptArray)
	rab := b.(OptArray)

	ra := raa.options
	rb := rab.options

	if (ra == nil) || (rb == nil) {
		panic("Attempted to comapre nil slices for equivalence.")
	}

	if len(ra) != len(rb) {
		return false
	}

	for i := range ra {
		if ra[i] != rb[i] {
			return false
		}
	}
	return true
}
