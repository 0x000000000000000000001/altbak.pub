package Data_Array_ST_Partial

import "gopurs/output/gopurs_runtime"

func PeekImpl(i int64, xs *[]interface{}, _ interface{}) interface{} {
	return (*xs)[i]
}
func PokeImpl(i int64, a interface{}, xs *[]interface{}, _ interface{}) bool {
	(*xs)[i] = a
	return true
}


// --- Auto-generated FFI wrappers ---
func Call_peekImpl(arg0 int64, arg1 *[]interface{}, arg2 interface{}) interface{} {
	return PeekImpl(arg0, arg1, arg2)
}
var _Gopurs_PeekImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[*[]interface{}](arg1)
	go_arg2 := arg2
	go_res := PeekImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_pokeImpl(arg0 int64, arg1 interface{}, arg2 *[]interface{}, arg3 interface{}) bool {
	return PokeImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_PokeImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*[]interface{}](arg2)
	go_arg3 := arg3
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
