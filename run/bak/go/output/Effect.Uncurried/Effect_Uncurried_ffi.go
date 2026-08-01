package Effect_Uncurried

import "gopurs/output/gopurs_runtime"


func MkEffectFn1(f interface{}) interface{} {
	return func(a interface{}) interface{} {
		return f.(func(interface{}) interface{})(a)
	}
}

func MkEffectFn2(f interface{}) interface{} {
	return func(a, b interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b)
	}
}

func MkEffectFn3(f interface{}) interface{} {
	return func(a, b, c interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c)
	}
}

func MkEffectFn4(f interface{}) interface{} {
	return func(a, b, c, d interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d)
	}
}

func MkEffectFn5(f interface{}) interface{} {
	return func(a, b, c, d, e interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e)
	}
}

func MkEffectFn6(f interface{}) interface{} {
	return func(a, b, c, d, e, g interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g)
	}
}

func MkEffectFn7(f interface{}) interface{} {
	return func(a, b, c, d, e, g, h interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h)
	}
}

func MkEffectFn8(f interface{}) interface{} {
	return func(a, b, c, d, e, g, h, i interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i)
	}
}

func MkEffectFn9(f interface{}) interface{} {
	return func(a, b, c, d, e, g, h, i, j interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j)
	}
}

func MkEffectFn10(f interface{}) interface{} {
	return func(a, b, c, d, e, g, h, i, j, k interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j).(func(interface{}) interface{})(k)
	}
}

func RunEffectFn1(f interface{}, a interface{}) interface{} {
	return f.(func(interface{}) interface{})(a)
}

func RunEffectFn2(f interface{}, a interface{}, b interface{}) interface{} {
	return f.(func(interface{}, interface{}) interface{})(a, b)
}

func RunEffectFn3(f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
	return f.(func(interface{}, interface{}, interface{}) interface{})(a, b, c)
}

func RunEffectFn4(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	return f.(func(interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d)
}

func RunEffectFn5(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e)
}

func RunEffectFn6(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, g interface{}) interface{} {
	var args []interface{}
	args = append(args, a, b, c, d, e, g)
	return f.(func(interface{}) interface{})(args)
}

func RunEffectFn7(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, g interface{}, h interface{}) interface{} {
	var args []interface{}
	args = append(args, a, b, c, d, e, g, h)
	return f.(func(interface{}) interface{})(args)
}

func RunEffectFn8(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	var args []interface{}
	args = append(args, a, b, c, d, e, g, h, i)
	return f.(func(interface{}) interface{})(args)
}

func RunEffectFn9(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	var args []interface{}
	args = append(args, a, b, c, d, e, g, h, i, j)
	return f.(func(interface{}) interface{})(args)
}

func RunEffectFn10(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	var args []interface{}
	args = append(args, a, b, c, d, e, g, h, i, j, k)
	return f.(func(interface{}) interface{})(args)
}


// --- Auto-generated FFI wrappers ---
func Call_mkEffectFn1(arg0 interface{}) interface{} {
	return MkEffectFn1(arg0)
}
var _Gopurs_MkEffectFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn2(arg0 interface{}) interface{} {
	return MkEffectFn2(arg0)
}
var _Gopurs_MkEffectFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn3(arg0 interface{}) interface{} {
	return MkEffectFn3(arg0)
}
var _Gopurs_MkEffectFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn4(arg0 interface{}) interface{} {
	return MkEffectFn4(arg0)
}
var _Gopurs_MkEffectFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn5(arg0 interface{}) interface{} {
	return MkEffectFn5(arg0)
}
var _Gopurs_MkEffectFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn6(arg0 interface{}) interface{} {
	return MkEffectFn6(arg0)
}
var _Gopurs_MkEffectFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn7(arg0 interface{}) interface{} {
	return MkEffectFn7(arg0)
}
var _Gopurs_MkEffectFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn8(arg0 interface{}) interface{} {
	return MkEffectFn8(arg0)
}
var _Gopurs_MkEffectFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn9(arg0 interface{}) interface{} {
	return MkEffectFn9(arg0)
}
var _Gopurs_MkEffectFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkEffectFn10(arg0 interface{}) interface{} {
	return MkEffectFn10(arg0)
}
var _Gopurs_MkEffectFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkEffectFn10(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn1(arg0 interface{}, arg1 interface{}) interface{} {
	return RunEffectFn1(arg0, arg1)
}
var _Gopurs_RunEffectFn1 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := RunEffectFn1(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn2(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return RunEffectFn2(arg0, arg1, arg2)
}
var _Gopurs_RunEffectFn2 = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := RunEffectFn2(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn3(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} {
	return RunEffectFn3(arg0, arg1, arg2, arg3)
}
var _Gopurs_RunEffectFn3 = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := RunEffectFn3(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn4(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}) interface{} {
	return RunEffectFn4(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_RunEffectFn4 = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := RunEffectFn4(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn5(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}) interface{} {
	return RunEffectFn5(arg0, arg1, arg2, arg3, arg4, arg5)
}
var _Gopurs_RunEffectFn5 = gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_res := RunEffectFn5(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn6(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}, arg6 interface{}) interface{} {
	return RunEffectFn6(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
var _Gopurs_RunEffectFn6 = gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_res := RunEffectFn6(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn7(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}, arg6 interface{}, arg7 interface{}) interface{} {
	return RunEffectFn7(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
var _Gopurs_RunEffectFn7 = gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_res := RunEffectFn7(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn8(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}, arg6 interface{}, arg7 interface{}, arg8 interface{}) interface{} {
	return RunEffectFn8(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}
var _Gopurs_RunEffectFn8 = gopurs_runtime.Func9(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_res := RunEffectFn8(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn9(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}, arg6 interface{}, arg7 interface{}, arg8 interface{}, arg9 interface{}) interface{} {
	return RunEffectFn9(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}
var _Gopurs_RunEffectFn9 = gopurs_runtime.Func10(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_arg9 := arg9
	go_res := RunEffectFn9(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9)
	return gopurs_runtime.Box(go_res)
})
func Call_runEffectFn10(arg0 interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}, arg4 interface{}, arg5 interface{}, arg6 interface{}, arg7 interface{}, arg8 interface{}, arg9 interface{}, arg10 interface{}) interface{} {
	return RunEffectFn10(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}
var _Gopurs_RunEffectFn10 = gopurs_runtime.Func11(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value, arg10 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_arg9 := arg9
	go_arg10 := arg10
	go_res := RunEffectFn10(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9, go_arg10)
	return gopurs_runtime.Box(go_res)
})
