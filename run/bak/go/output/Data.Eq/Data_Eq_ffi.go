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
var _Gopurs_EqArrayImpl = // TAST: (ForAll [a] (Func [(Func [(TypeVar a), (TypeVar a)] Boolean), (Array (TypeVar a)), (Array (TypeVar a))] Boolean))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) bool {
			inner_res0 := gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := EqArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_EqBooleanImpl = // TAST: (Func [Boolean, Boolean] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[bool](arg0)
	go_arg1 := gopurs_runtime.Unbox[bool](arg1)
	go_res := EqBooleanImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_EqCharImpl = // TAST: (Func [Char, Char] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := EqCharImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_EqIntImpl = // TAST: (Func [Int, Int] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := EqIntImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_EqNumberImpl = // TAST: (Func [Number, Number] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := EqNumberImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_EqStringImpl = // TAST: (Func [String, String] Boolean)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := EqStringImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})