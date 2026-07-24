package Data_HeytingAlgebra

import "gopurs/output/gopurs_runtime"

func BoolConj(b1 bool, b2 bool) bool {
	return b1 && b2
}
func BoolDisj(b1 bool, b2 bool) bool {
	return b1 || b2
}
func BoolNot(b bool) bool {
	return !b
}


// --- Auto-generated FFI wrappers ---
func Call_boolConj(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := BoolConj(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_BoolConj = gopurs_runtime.Func2(Call_boolConj)
func Call_boolDisj(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := BoolDisj(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_BoolDisj = gopurs_runtime.Func2(Call_boolDisj)
func Call_boolNot(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_res := BoolNot(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_BoolNot = gopurs_runtime.Func(Call_boolNot)
