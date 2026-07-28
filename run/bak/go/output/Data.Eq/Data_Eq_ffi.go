package Data_Eq

import "gopurs/output/gopurs_runtime"

func refEq(r1 interface{}, r2 interface{}) bool {
	return r1 == r2
}
func EqBooleanImpl(r1 bool, r2 bool) bool {
	return r1 == r2
}
func EqIntImpl(r1 int64, r2 int64) bool {
	return r1 == r2
}
func EqNumberImpl(r1 float64, r2 float64) bool {
	return r1 == r2
}
func EqCharImpl(r1 string, r2 string) bool {
	return r1 == r2
}
func EqStringImpl(r1 string, r2 string) bool {
	return r1 == r2
}
func EqArrayImpl(f func(interface{}, interface{}) bool, xs []interface{}, ys []interface{}) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !f(xs[i], ys[i]) {
			return false
		}
	}
	return true
}


// --- Auto-generated FFI wrappers ---
func Call_eqBooleanImpl(arg0 bool, arg1 bool) bool {
	return EqBooleanImpl(arg0, arg1)
}
var _Gopurs_EqBooleanImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := EqBooleanImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_eqIntImpl(arg0 int64, arg1 int64) bool {
	return EqIntImpl(arg0, arg1)
}
var _Gopurs_EqIntImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := EqIntImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_eqNumberImpl(arg0 float64, arg1 float64) bool {
	return EqNumberImpl(arg0, arg1)
}
var _Gopurs_EqNumberImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := EqNumberImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_eqCharImpl(arg0 string, arg1 string) bool {
	return EqCharImpl(arg0, arg1)
}
var _Gopurs_EqCharImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := EqCharImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_eqStringImpl(arg0 string, arg1 string) bool {
	return EqStringImpl(arg0, arg1)
}
var _Gopurs_EqStringImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := EqStringImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_eqArrayImpl(arg0 func(interface{}, interface{}) bool, arg1 []interface{}, arg2 []interface{}) bool {
	return EqArrayImpl(arg0, arg1, arg2)
}
var _Gopurs_EqArrayImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}, p0_1 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := EqArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
