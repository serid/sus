package val

import "fmt"

// An optional value
type Option struct {
	data Val
}

func NewOption(data Val) Option {
	return Option{data: data}
}

func (op Option) IsNone() bool {
	return op.data == nil
}

func (op Option) IsSome() bool {
	return op.data != nil
}

func (op Option) Unwrap() Val {
	if op.IsNone() {
		panic("called `Unwrap` on a None variant")
	}
	return op.data
}

func (op *Option) SetData(newData Val) {
	if op.IsSome() {
		panic("attempted to overwrite a value inside Option")
	}
	op.data = newData
}

func (op Option) Clone() Option {
	if op.IsNone() {
		return Option{data: nil}
	}
	return Option{data: op.data.Clone()}
}

func (op Option) String() string {
	return fmt.Sprintf("Option(%v)", op.data)
}
