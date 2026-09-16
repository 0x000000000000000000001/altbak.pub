package Data_Lazy

import (
	"sync"
	"gopurs/output/gopurs_runtime"
)

func Defer(thunk gopurs_runtime.Value) gopurs_runtime.Value {
	var once sync.Once
	var result gopurs_runtime.Value

	return gopurs_runtime.Func(func(_dollar__unused gopurs_runtime.Value) gopurs_runtime.Value {
		once.Do(func() {
			result = gopurs_runtime.Apply(thunk, gopurs_runtime.Value{})
		})
		return result
	})
}

func Force(l gopurs_runtime.Value) gopurs_runtime.Value {
	res := gopurs_runtime.Apply(l, gopurs_runtime.Value{})
	return res
}
