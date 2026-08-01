package Effect

import "gopurs/output/gopurs_runtime"

func PureE(a any, _ interface{}) any {
	return a
}
func BindE(a func(interface{}) any, f func(any) func(interface{}) any, _ interface{}) any {
	resA := a(nil)
	return f(resA)(nil)
}
func UntilE(f func(interface{}) any, _ interface{}) any {
	for {
		if f(nil).(bool) {
			break
		}
	}
	return nil
}
func WhileE(f func(interface{}) any, a func(interface{}) any, _ interface{}) any {
	for {
		if !f(nil).(bool) {
			break
		}
		a(nil)
	}
	return nil
}
func ForE(lo int64, hi int64, f func(any) func(interface{}) any, _ interface{}) any {
	for i := lo; i < hi; i++ {
		f(i)(nil)
	}
	return nil
}
func ForeachE(as []any, f func(any) func(interface{}) any, _ interface{}) any {
	for _, a := range as {
		f(a)(nil)
	}
	return nil
}


// --- Auto-generated FFI wrappers ---
func Call_pureE(arg0 any, arg1 interface{}) any {
	return PureE(arg0, arg1)
}
var _Gopurs_PureE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := PureE(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_bindE(arg0 func(interface{}) any, arg1 func(any) func(interface{}) any, arg2 interface{}) any {
	return BindE(arg0, arg1, arg2)
}
var _Gopurs_BindE = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 any) func(interface{}) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := BindE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_untilE(arg0 func(interface{}) any, arg1 interface{}) any {
	return UntilE(arg0, arg1)
}
var _Gopurs_UntilE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_res := UntilE(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_whileE(arg0 func(interface{}) any, arg1 func(interface{}) any, arg2 interface{}) any {
	return WhileE(arg0, arg1, arg2)
}
var _Gopurs_WhileE = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 interface{}) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := WhileE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_forE(arg0 int64, arg1 int64, arg2 func(any) func(interface{}) any, arg3 interface{}) any {
	return ForE(arg0, arg1, arg2, arg3)
}
var _Gopurs_ForE = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func(p0_0 any) func(interface{}) any {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg3 := arg3
	go_res := ForE(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_foreachE(arg0 []any, arg1 func(any) func(interface{}) any, arg2 interface{}) any {
	return ForeachE(arg0, arg1, arg2)
}
var _Gopurs_ForeachE = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 any) func(interface{}) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := ForeachE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
