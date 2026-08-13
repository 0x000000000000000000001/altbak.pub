package purescript



import (
	"sync"
	"gopurs/output/gopurs_runtime"
)

func Data_Lazy_Defer(thunk gopurs_runtime.Value) gopurs_runtime.Value {
	var once sync.Once
	var result gopurs_runtime.Value

	return gopurs_runtime.Func(func(_dollar__unused gopurs_runtime.Value) gopurs_runtime.Value {
		once.Do(func() {
			result = gopurs_runtime.Apply(thunk, gopurs_runtime.Value{})
		})
		return result
	})
}

func Data_Lazy_Force(l gopurs_runtime.Value) gopurs_runtime.Value {
	res := gopurs_runtime.Apply(l, gopurs_runtime.Value{})
	return res
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Lazy_Go__defer = // TAST: (ForAll [a] (Func [(Func [Unit] (TypeVar a))] (ADT ["Data","Lazy","Lazy"] [(TypeVar a)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Lazy_Defer(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Lazy_Force = // TAST: (ForAll [a] (Func [(ADT ["Data","Lazy","Lazy"] [(TypeVar a)])] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Lazy_Force(go_arg0)
	return gopurs_runtime.Box(go_res)
})