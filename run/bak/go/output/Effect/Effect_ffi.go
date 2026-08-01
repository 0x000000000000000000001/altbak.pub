package Effect

import "gopurs/output/gopurs_runtime"



func PureE(a any) func() any {
	return func() any {
		return a
	}
}
func BindE(a func() any, f func(any) func() any) func() any {
	return func() any {
		resA := a()
		return f(resA)()
	}
}
func UntilE(f func() any) func() any {
	return func() any {
		for {
			if f().(bool) {
				break
			}
		}
		return nil
	}
}
func WhileE(f func() any, a func() any) func() any {
	return func() any {
		for {
			if !f().(bool) {
				break
			}
			a()
		}
		return nil
	}
}
func ForE(lo int64, hi int64, f func(any) func() any) func() any {
	return func() any {
		for i := lo; i < hi; i++ {
			f(i)()
		}
		return nil
	}
}
func ForeachE(as []any, f func(any) func() any) func() any {
	return func() any {
		for _, a := range as {
			f(a)()
		}
		return nil
	}
}


// --- Auto-generated FFI wrappers ---
func Call_pureE(arg0 any) func() any {
	return PureE(arg0)
}
var _Gopurs_PureE = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := PureE(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_bindE(arg0 func() any, arg1 func(any) func() any) func() any {
	return BindE(arg0, arg1)
}
var _Gopurs_BindE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_arg1 := func(p0_0 any) func() any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func() any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := BindE(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_untilE(arg0 func() any) func() any {
	return UntilE(arg0)
}
var _Gopurs_UntilE = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_res := UntilE(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_whileE(arg0 func() any, arg1 func() any) func() any {
	return WhileE(arg0, arg1)
}
var _Gopurs_WhileE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_arg1 := func() any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Value{})
		}
	go_res := WhileE(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_forE(arg0 int64, arg1 int64, arg2 func(any) func() any) func() any {
	return ForE(arg0, arg1, arg2)
}
var _Gopurs_ForE = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func(p0_0 any) func() any {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func() any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := ForE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_foreachE(arg0 []any, arg1 func(any) func() any) func() any {
	return ForeachE(arg0, arg1)
}
var _Gopurs_ForeachE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 any) func() any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func() any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := ForeachE(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
