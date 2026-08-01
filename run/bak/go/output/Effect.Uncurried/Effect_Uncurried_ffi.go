package Effect_Uncurried

import "gopurs/output/gopurs_runtime"

func box(v interface{}) gopurs_runtime.Value {
    if val, ok := v.(gopurs_runtime.Value); ok {
        return val
    }
    return gopurs_runtime.Box(v)
}

func MkEffectFn1(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply(fn, a)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn2(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func2(func(a gopurs_runtime.Value, b gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply2(fn, a, b)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn3(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func3(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply3(fn, a, b, c)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn4(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func4(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply4(fn, a, b, c, d)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn5(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func5(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply5(fn, a, b, c, d, e)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn6(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func6(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply6(fn, a, b, c, d, e, argf)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn7(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func7(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply7(fn, a, b, c, d, e, argf, g)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn8(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func8(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply8(fn, a, b, c, d, e, argf, g, h)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn9(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func9(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value, i gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply9(fn, a, b, c, d, e, argf, g, h, i)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn10(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func10(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value, i gopurs_runtime.Value, j gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply10(fn, a, b, c, d, e, argf, g, h, i, j)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func RunEffectFn1(f interface{}, a interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.Apply(box(f), box(a)).UnsafePtr
    }
}

func RunEffectFn2(f interface{}, a interface{}, b interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp2(box(f), box(a), box(b)).UnsafePtr
    }
}

func RunEffectFn3(f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp3(box(f), box(a), box(b), box(c)).UnsafePtr
    }
}

func RunEffectFn4(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp4(box(f), box(a), box(b), box(c), box(d)).UnsafePtr
    }
}

func RunEffectFn5(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp5(box(f), box(a), box(b), box(c), box(d), box(e)).UnsafePtr
    }
}

func RunEffectFn6(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp6(box(f), box(a), box(b), box(c), box(d), box(e), box(argf)).UnsafePtr
    }
}

func RunEffectFn7(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp7(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g)).UnsafePtr
    }
}

func RunEffectFn8(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}, h interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp8(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g), box(h)).UnsafePtr
    }
}

func RunEffectFn9(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}, h interface{}, i interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp9(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g), box(h), box(i)).UnsafePtr
    }
}

func RunEffectFn10(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp10(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g), box(h), box(i), box(j)).UnsafePtr
    }
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
