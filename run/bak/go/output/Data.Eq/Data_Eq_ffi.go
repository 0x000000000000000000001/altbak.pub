package Data_Eq

import "gopurs/output/gopurs_runtime"

func refEq(r1 any, r2 any) bool {
	return r1 == r2
}
func EqBooleanImpl(r1 bool, r2 bool) bool { return r1 == r2 }
func EqIntImpl(r1 int, r2 int) bool { return r1 == r2 }
func EqNumberImpl(r1 float64, r2 float64) bool { return r1 == r2 }
func EqCharImpl(r1 string, r2 string) bool { return r1 == r2 }
func EqStringImpl(r1 string, r2 string) bool { return r1 == r2 }
func EqArrayImpl(f func(any) func(any) bool, xs []any, ys []any) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !f(xs[i])(ys[i]) {
			return false
		}
	}
	return true
}


// --- Auto-generated FFI wrappers ---
func Call_eqBooleanImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := EqBooleanImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_EqBooleanImpl = gopurs_runtime.Func2(Call_eqBooleanImpl)
func Call_eqIntImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := EqIntImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_EqIntImpl = gopurs_runtime.Func2(Call_eqIntImpl)
func Call_eqNumberImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := EqNumberImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_EqNumberImpl = gopurs_runtime.Func2(Call_eqNumberImpl)
func Call_eqCharImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := EqCharImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_EqCharImpl = gopurs_runtime.Func2(Call_eqCharImpl)
func Call_eqStringImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := EqStringImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_EqStringImpl = gopurs_runtime.Func2(Call_eqStringImpl)
func Call_eqArrayImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) bool {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[bool](inner_res1)
		}
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := EqArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_EqArrayImpl = gopurs_runtime.Func3(Call_eqArrayImpl)
