package Data_Function_Uncurried

import "gopurs/output/gopurs_runtime"




func identity(fn any) any {
	return fn
}

var RunFn2 = identity
var RunFn3 = identity
var RunFn4 = identity
var RunFn5 = identity
var RunFn6 = identity
var RunFn7 = identity
var RunFn8 = identity
var RunFn9 = identity
var RunFn10 = identity

var MkFn2 = identity
var MkFn3 = identity
var MkFn4 = identity
var MkFn5 = identity
var MkFn6 = identity
var MkFn7 = identity
var MkFn8 = identity
var MkFn9 = identity
var MkFn10 = identity

func MkFn0(f any) any {
	return f
}

func RunFn0(f func(any) any) any {
	return f(nil)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_RunFn2 = gopurs_runtime.Box(RunFn2)
var _Gopurs_RunFn3 = gopurs_runtime.Box(RunFn3)
var _Gopurs_RunFn4 = gopurs_runtime.Box(RunFn4)
var _Gopurs_RunFn5 = gopurs_runtime.Box(RunFn5)
var _Gopurs_RunFn6 = gopurs_runtime.Box(RunFn6)
var _Gopurs_RunFn7 = gopurs_runtime.Box(RunFn7)
var _Gopurs_RunFn8 = gopurs_runtime.Box(RunFn8)
var _Gopurs_RunFn9 = gopurs_runtime.Box(RunFn9)
var _Gopurs_RunFn10 = gopurs_runtime.Box(RunFn10)
var _Gopurs_MkFn2 = gopurs_runtime.Box(MkFn2)
var _Gopurs_MkFn3 = gopurs_runtime.Box(MkFn3)
var _Gopurs_MkFn4 = gopurs_runtime.Box(MkFn4)
var _Gopurs_MkFn5 = gopurs_runtime.Box(MkFn5)
var _Gopurs_MkFn6 = gopurs_runtime.Box(MkFn6)
var _Gopurs_MkFn7 = gopurs_runtime.Box(MkFn7)
var _Gopurs_MkFn8 = gopurs_runtime.Box(MkFn8)
var _Gopurs_MkFn9 = gopurs_runtime.Box(MkFn9)
var _Gopurs_MkFn10 = gopurs_runtime.Box(MkFn10)
func Call_mkFn0(arg0 any) any {
	return MkFn0(arg0)
}
var _Gopurs_MkFn0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkFn0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runFn0(arg0 func(any) any) any {
	return RunFn0(arg0)
}
var _Gopurs_RunFn0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := RunFn0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
