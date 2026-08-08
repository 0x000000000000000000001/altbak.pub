package Data_Array_ST_Partial

import "gopurs/output/gopurs_runtime"


func PeekImpl(i int64, xs *[]interface{}) interface{} {
	return (*xs)[i]
}
func PokeImpl(i int64, a interface{}, xs *[]interface{}) bool {
	(*xs)[i] = a
	return true
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_PeekImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn2"] [Int, (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (TypeVar a)])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[*[]interface{}](arg1)
	go_res := PeekImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PokeImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn3"] [Int, (TypeVar a), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Unit","Unit"] [])])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*[]interface{}](arg2)
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})