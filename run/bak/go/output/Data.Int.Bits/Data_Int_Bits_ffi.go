package Data_Int_Bits

import "gopurs/output/gopurs_runtime"

func And(n1 int, n2 int) int { return n1 & n2 }
func Or(n1 int, n2 int) int { return n1 | n2 }
func Xor(n1 int, n2 int) int { return n1 ^ n2 }
func Shl(n1 int, n2 int) int { return n1 << n2 }
func Shr(n1 int, n2 int) int { return n1 >> n2 }
func Zshr(n1 int, n2 int) int { return int(uint(n1) >> uint(n2)) }
func Complement(n int) int { return ^n }


// --- Auto-generated FFI wrappers ---
var _Gopurs_And = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := And(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Or = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Or(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Xor = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Xor(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Shl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Shl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Shr = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Shr(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Zshr = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Zshr(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Complement = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Complement(go_arg0)
	return gopurs_runtime.Box(go_res)
})
