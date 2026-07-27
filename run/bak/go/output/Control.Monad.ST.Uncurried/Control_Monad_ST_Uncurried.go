package Control_Monad_ST_Uncurried

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_mkSTFn1 gopurs_runtime.Value
var once_mkSTFn1 sync.Once
func Get_mkSTFn1() gopurs_runtime.Value {
	once_mkSTFn1.Do(func() {
		cache_mkSTFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn1(func(inner_arg0 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)), nil))
}
})
})
	})
	return cache_mkSTFn1
}

var cache_mkSTFn10 gopurs_runtime.Value
var once_mkSTFn10 sync.Once
func Get_mkSTFn10() gopurs_runtime.Value {
	once_mkSTFn10.Do(func() {
		cache_mkSTFn10 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn10(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}, inner_arg8 interface{}, inner_arg9 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply10(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7), gopurs_runtime.Any(inner_arg8), gopurs_runtime.Any(inner_arg9)), nil))
}
})
})
	})
	return cache_mkSTFn10
}

var cache_mkSTFn2 gopurs_runtime.Value
var once_mkSTFn2 sync.Once
func Get_mkSTFn2() gopurs_runtime.Value {
	once_mkSTFn2.Do(func() {
		cache_mkSTFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn2(func(inner_arg0 interface{}, inner_arg1 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)), nil))
}
})
})
	})
	return cache_mkSTFn2
}

var cache_mkSTFn3 gopurs_runtime.Value
var once_mkSTFn3 sync.Once
func Get_mkSTFn3() gopurs_runtime.Value {
	once_mkSTFn3.Do(func() {
		cache_mkSTFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn3(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply3(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2)), nil))
}
})
})
	})
	return cache_mkSTFn3
}

var cache_mkSTFn4 gopurs_runtime.Value
var once_mkSTFn4 sync.Once
func Get_mkSTFn4() gopurs_runtime.Value {
	once_mkSTFn4.Do(func() {
		cache_mkSTFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn4(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply4(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3)), nil))
}
})
})
	})
	return cache_mkSTFn4
}

var cache_mkSTFn5 gopurs_runtime.Value
var once_mkSTFn5 sync.Once
func Get_mkSTFn5() gopurs_runtime.Value {
	once_mkSTFn5.Do(func() {
		cache_mkSTFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn5(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply5(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4)), nil))
}
})
})
	})
	return cache_mkSTFn5
}

var cache_mkSTFn6 gopurs_runtime.Value
var once_mkSTFn6 sync.Once
func Get_mkSTFn6() gopurs_runtime.Value {
	once_mkSTFn6.Do(func() {
		cache_mkSTFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn6(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply6(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5)), nil))
}
})
})
	})
	return cache_mkSTFn6
}

var cache_mkSTFn7 gopurs_runtime.Value
var once_mkSTFn7 sync.Once
func Get_mkSTFn7() gopurs_runtime.Value {
	once_mkSTFn7.Do(func() {
		cache_mkSTFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn7(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply7(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6)), nil))
}
})
})
	})
	return cache_mkSTFn7
}

var cache_mkSTFn8 gopurs_runtime.Value
var once_mkSTFn8 sync.Once
func Get_mkSTFn8() gopurs_runtime.Value {
	once_mkSTFn8.Do(func() {
		cache_mkSTFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn8(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply8(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7)), nil))
}
})
})
	})
	return cache_mkSTFn8
}

var cache_mkSTFn9 gopurs_runtime.Value
var once_mkSTFn9 sync.Once
func Get_mkSTFn9() gopurs_runtime.Value {
	once_mkSTFn9.Do(func() {
		cache_mkSTFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return MkSTFn9(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}, inner_arg8 interface{}) func() interface{} {
return func() interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply9(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7), gopurs_runtime.Any(inner_arg8)), nil))
}
})
})
	})
	return cache_mkSTFn9
}

var cache_runSTFn1 gopurs_runtime.Value
var once_runSTFn1 sync.Once
func Get_runSTFn1() gopurs_runtime.Value {
	once_runSTFn1.Do(func() {
		cache_runSTFn1 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn1(arg0, gopurs_runtime.UnboxAny(arg1))())
})
})
	})
	return cache_runSTFn1
}

var cache_runSTFn10 gopurs_runtime.Value
var once_runSTFn10 sync.Once
func Get_runSTFn10() gopurs_runtime.Value {
	once_runSTFn10.Do(func() {
		cache_runSTFn10 = gopurs_runtime.Func11(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value, arg10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn10(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7), gopurs_runtime.UnboxAny(arg8), gopurs_runtime.UnboxAny(arg9), gopurs_runtime.UnboxAny(arg10))())
})
})
	})
	return cache_runSTFn10
}

var cache_runSTFn2 gopurs_runtime.Value
var once_runSTFn2 sync.Once
func Get_runSTFn2() gopurs_runtime.Value {
	once_runSTFn2.Do(func() {
		cache_runSTFn2 = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn2(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2))())
})
})
	})
	return cache_runSTFn2
}

var cache_runSTFn3 gopurs_runtime.Value
var once_runSTFn3 sync.Once
func Get_runSTFn3() gopurs_runtime.Value {
	once_runSTFn3.Do(func() {
		cache_runSTFn3 = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn3(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3))())
})
})
	})
	return cache_runSTFn3
}

var cache_runSTFn4 gopurs_runtime.Value
var once_runSTFn4 sync.Once
func Get_runSTFn4() gopurs_runtime.Value {
	once_runSTFn4.Do(func() {
		cache_runSTFn4 = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn4(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4))())
})
})
	})
	return cache_runSTFn4
}

var cache_runSTFn5 gopurs_runtime.Value
var once_runSTFn5 sync.Once
func Get_runSTFn5() gopurs_runtime.Value {
	once_runSTFn5.Do(func() {
		cache_runSTFn5 = gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn5(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5))())
})
})
	})
	return cache_runSTFn5
}

var cache_runSTFn6 gopurs_runtime.Value
var once_runSTFn6 sync.Once
func Get_runSTFn6() gopurs_runtime.Value {
	once_runSTFn6.Do(func() {
		cache_runSTFn6 = gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn6(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6))())
})
})
	})
	return cache_runSTFn6
}

var cache_runSTFn7 gopurs_runtime.Value
var once_runSTFn7 sync.Once
func Get_runSTFn7() gopurs_runtime.Value {
	once_runSTFn7.Do(func() {
		cache_runSTFn7 = gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn7(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7))())
})
})
	})
	return cache_runSTFn7
}

var cache_runSTFn8 gopurs_runtime.Value
var once_runSTFn8 sync.Once
func Get_runSTFn8() gopurs_runtime.Value {
	once_runSTFn8.Do(func() {
		cache_runSTFn8 = gopurs_runtime.Func9(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn8(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7), gopurs_runtime.UnboxAny(arg8))())
})
})
	})
	return cache_runSTFn8
}

var cache_runSTFn9 gopurs_runtime.Value
var once_runSTFn9 sync.Once
func Get_runSTFn9() gopurs_runtime.Value {
	once_runSTFn9.Do(func() {
		cache_runSTFn9 = gopurs_runtime.Func10(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(RunSTFn9(arg0, gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2), gopurs_runtime.UnboxAny(arg3), gopurs_runtime.UnboxAny(arg4), gopurs_runtime.UnboxAny(arg5), gopurs_runtime.UnboxAny(arg6), gopurs_runtime.UnboxAny(arg7), gopurs_runtime.UnboxAny(arg8), gopurs_runtime.UnboxAny(arg9))())
})
})
	})
	return cache_runSTFn9
}


