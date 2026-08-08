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


// --- Auto-generated FFI wrappers ---
var _Gopurs_Go__defer = // TAST: (Func [(Func [(ADT ["Data","Unit","Unit"] [])] (TypeVar a))] (ADT ["Data","Lazy","Lazy"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Defer(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Force = // TAST: (Func [(ADT ["Data","Lazy","Lazy"] [(TypeVar a)])] (TypeVar a))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Force(go_arg0)
	return gopurs_runtime.Box(go_res)
})