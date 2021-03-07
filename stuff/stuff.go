package stuff

import "reflect"

func Unwrap(err error) {
	if err != nil {
		panic(err)
	}
}

func IsNilExt(i interface{}) bool {
	if i == nil {
		return true
	}
	vo := reflect.ValueOf(i)
	return CanBeNil(vo) && vo.IsNil()
}

func CanBeNil(value reflect.Value) bool {
	k := value.Kind()
	return k == reflect.Chan ||
		k == reflect.Func ||
		k == reflect.Interface ||
		k == reflect.Map ||
		k == reflect.Ptr ||
		k == reflect.Slice
}

func Catch(f func() interface{}) (p interface{}) {
	defer func() {
		p = recover()
	}()

	ok := f()

	return ok
}
