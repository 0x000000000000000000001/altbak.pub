package Data_Function_Uncurried

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_runFn1 gopurs_runtime.Value
var once_runFn1 sync.Once
func Get_runFn1() gopurs_runtime.Value {
	once_runFn1.Do(func() {
		cache_runFn1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runFn1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_runFn1
}

var cache_mkFn1 gopurs_runtime.Value
var once_mkFn1 sync.Once
func Get_mkFn1() gopurs_runtime.Value {
	once_mkFn1.Do(func() {
		cache_mkFn1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mkFn1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_mkFn1
}

var cache_mkFn0 gopurs_runtime.Value
var once_mkFn0 sync.Once
func Get_mkFn0() gopurs_runtime.Value {
	once_mkFn0.Do(func() {
		cache_mkFn0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn0(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
})
})
	})
	return cache_mkFn0
}

var cache_mkFn10 gopurs_runtime.Value
var once_mkFn10 sync.Once
func Get_mkFn10() gopurs_runtime.Value {
	once_mkFn10.Do(func() {
		cache_mkFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn10(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}, inner_arg8 interface{}, inner_arg9 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply10(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7), gopurs_runtime.Any(inner_arg8), gopurs_runtime.Any(inner_arg9)))
})
})
	})
	return cache_mkFn10
}

var cache_mkFn2 gopurs_runtime.Value
var once_mkFn2 sync.Once
func Get_mkFn2() gopurs_runtime.Value {
	once_mkFn2.Do(func() {
		cache_mkFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn2(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
})
})
	})
	return cache_mkFn2
}

var cache_mkFn3 gopurs_runtime.Value
var once_mkFn3 sync.Once
func Get_mkFn3() gopurs_runtime.Value {
	once_mkFn3.Do(func() {
		cache_mkFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn3(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2)))
})
})
	})
	return cache_mkFn3
}

var cache_mkFn4 gopurs_runtime.Value
var once_mkFn4 sync.Once
func Get_mkFn4() gopurs_runtime.Value {
	once_mkFn4.Do(func() {
		cache_mkFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn4(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply4(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3)))
})
})
	})
	return cache_mkFn4
}

var cache_mkFn5 gopurs_runtime.Value
var once_mkFn5 sync.Once
func Get_mkFn5() gopurs_runtime.Value {
	once_mkFn5.Do(func() {
		cache_mkFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn5(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply5(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4)))
})
})
	})
	return cache_mkFn5
}

var cache_mkFn6 gopurs_runtime.Value
var once_mkFn6 sync.Once
func Get_mkFn6() gopurs_runtime.Value {
	once_mkFn6.Do(func() {
		cache_mkFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn6(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply6(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5)))
})
})
	})
	return cache_mkFn6
}

var cache_mkFn7 gopurs_runtime.Value
var once_mkFn7 sync.Once
func Get_mkFn7() gopurs_runtime.Value {
	once_mkFn7.Do(func() {
		cache_mkFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn7(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply7(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6)))
})
})
	})
	return cache_mkFn7
}

var cache_mkFn8 gopurs_runtime.Value
var once_mkFn8 sync.Once
func Get_mkFn8() gopurs_runtime.Value {
	once_mkFn8.Do(func() {
		cache_mkFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn8(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply8(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7)))
})
})
	})
	return cache_mkFn8
}

var cache_mkFn9 gopurs_runtime.Value
var once_mkFn9 sync.Once
func Get_mkFn9() gopurs_runtime.Value {
	once_mkFn9.Do(func() {
		cache_mkFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkFn9(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}, inner_arg8 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply9(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7), gopurs_runtime.Any(inner_arg8)))
})
})
	})
	return cache_mkFn9
}

var cache_runFn0 gopurs_runtime.Value
var once_runFn0 sync.Once
func Get_runFn0() gopurs_runtime.Value {
	once_runFn0.Do(func() {
		cache_runFn0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn0(arg0))
})
	})
	return cache_runFn0
}

var cache_runFn10 gopurs_runtime.Value
var once_runFn10 sync.Once
func Get_runFn10() gopurs_runtime.Value {
	once_runFn10.Do(func() {
		cache_runFn10 = gopurs_runtime.Func11(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value, arg10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn10(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7), gopurs_runtime.UnboxAny(arg8), gopurs_runtime.UnboxAny(arg9), gopurs_runtime.UnboxAny(arg10)))
})
	})
	return cache_runFn10
}

var cache_runFn2 gopurs_runtime.Value
var once_runFn2 sync.Once
func Get_runFn2() gopurs_runtime.Value {
	once_runFn2.Do(func() {
		cache_runFn2 = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn2(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2)))
})
	})
	return cache_runFn2
}

var cache_runFn3 gopurs_runtime.Value
var once_runFn3 sync.Once
func Get_runFn3() gopurs_runtime.Value {
	once_runFn3.Do(func() {
		cache_runFn3 = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn3(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3)))
})
	})
	return cache_runFn3
}

var cache_runFn4 gopurs_runtime.Value
var once_runFn4 sync.Once
func Get_runFn4() gopurs_runtime.Value {
	once_runFn4.Do(func() {
		cache_runFn4 = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn4(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4)))
})
	})
	return cache_runFn4
}

var cache_runFn5 gopurs_runtime.Value
var once_runFn5 sync.Once
func Get_runFn5() gopurs_runtime.Value {
	once_runFn5.Do(func() {
		cache_runFn5 = gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn5(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5)))
})
	})
	return cache_runFn5
}

var cache_runFn6 gopurs_runtime.Value
var once_runFn6 sync.Once
func Get_runFn6() gopurs_runtime.Value {
	once_runFn6.Do(func() {
		cache_runFn6 = gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn6(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6)))
})
	})
	return cache_runFn6
}

var cache_runFn7 gopurs_runtime.Value
var once_runFn7 sync.Once
func Get_runFn7() gopurs_runtime.Value {
	once_runFn7.Do(func() {
		cache_runFn7 = gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn7(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7)))
})
	})
	return cache_runFn7
}

var cache_runFn8 gopurs_runtime.Value
var once_runFn8 sync.Once
func Get_runFn8() gopurs_runtime.Value {
	once_runFn8.Do(func() {
		cache_runFn8 = gopurs_runtime.Func9(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn8(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7), gopurs_runtime.UnboxAny(arg8)))
})
	})
	return cache_runFn8
}

var cache_runFn9 gopurs_runtime.Value
var once_runFn9 sync.Once
func Get_runFn9() gopurs_runtime.Value {
	once_runFn9.Do(func() {
		cache_runFn9 = gopurs_runtime.Func10(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(RunFn9(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7), gopurs_runtime.UnboxAny(arg8), gopurs_runtime.UnboxAny(arg9)))
})
	})
	return cache_runFn9
}

func Call_runFn1(f_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_mkFn1(f_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(arg0)))
})
}
