package Effect

import "gopurs/output/gopurs_runtime"



func PureE(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}

func BindE(a func() interface{}, f func(interface{}) func() interface{}) func() interface{} {
	return func() interface{} {
		resA := a()
		return f(resA)()
	}
}

func UntilE(f func() bool) func() interface{} {
	return func() interface{} {
		for {
			if f() {
				break
			}
		}
		return nil
	}
}

func WhileE(f func() bool, a func() interface{}) func() interface{} {
	return func() interface{} {
		for {
			if !f() {
				break
			}
			a()
		}
		return nil
	}
}

func ForE(lo int, hi int, f func(int) func() interface{}) func() interface{} {
	return func() interface{} {
		for i := lo; i < hi; i++ {
			f(i)()
		}
		return nil
	}
}

func ForeachE(as []interface{}, f func(interface{}) func() interface{}) func() interface{} {
	return func() interface{} {
		for _, v := range as {
			f(v)()
		}
		return nil
	}
}


// --- Auto-generated FFI wrappers ---
func Call_pureE(arg0 interface{}) func() interface{} {
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
func Call_bindE(arg0 func() interface{}, arg1 func(interface{}) func() interface{}) func() interface{} {
	return BindE(arg0, arg1)
}
var _Gopurs_BindE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_arg1 := func(p0_0 interface{}) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := BindE(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_untilE(arg0 func() bool) func() interface{} {
	return UntilE(arg0)
}
var _Gopurs_UntilE = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_res := UntilE(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_whileE(arg0 func() bool, arg1 func() interface{}) func() interface{} {
	return WhileE(arg0, arg1)
}
var _Gopurs_WhileE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg1 := func() interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Value{})
		}
	go_res := WhileE(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_forE(arg0 int, arg1 int, arg2 func(int) func() interface{}) func() interface{} {
	return ForE(arg0, arg1, arg2)
}
var _Gopurs_ForE = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := func(p0_0 int) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := ForE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_foreachE(arg0 []interface{}, arg1 func(interface{}) func() interface{}) func() interface{} {
	return ForeachE(arg0, arg1)
}
var _Gopurs_ForeachE = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 interface{}) func() interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func() interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Value{})
		}
		}
	go_res := ForeachE(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
