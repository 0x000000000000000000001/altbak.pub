package Control_Monad_ST_Uncurried

import "gopurs/output/gopurs_runtime"

func MkSTFn1(fn func(any) any) any {
	return func(a any) any { return fn(a) }
}
func MkSTFn2(fn func(any) func(any) any) any {
	return func(a any, b any) any { return fn(a)(b) }
}
func MkSTFn3(fn func(any) func(any) func(any) any) any {
	return func(a any, b any, c any) any { return fn(a)(b)(c) }
}
func MkSTFn4(fn func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any) any { return fn(a)(b)(c)(d) }
}
func MkSTFn5(fn func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any) any { return fn(a)(b)(c)(d)(e) }
}
func MkSTFn6(fn func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any) any { return fn(a)(b)(c)(d)(e)(f) }
}
func MkSTFn7(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any) any { return fn(a)(b)(c)(d)(e)(f)(g) }
}
func MkSTFn8(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any, h any) any { return fn(a)(b)(c)(d)(e)(f)(g)(h) }
}
func MkSTFn9(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any, h any, i any) any { return fn(a)(b)(c)(d)(e)(f)(g)(h)(i) }
}
// func MkSTFn10(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	// return func any { return fn(a)(b)(c)(d)(e)(f)(g)(h)(i)(j) }
// }
func RunSTFn1(fn any, a any, _ interface{}) any { return fn.(func(any) any)(a) }
func RunSTFn2(fn any, a any, b any, _ interface{}) any { return fn.(func(any, any) any)(a, b) }
func RunSTFn3(fn any, a any, b any, c any, _ interface{}) any { return fn.(func(any, any, any) any)(a, b, c) }
func RunSTFn4(fn any, a any, b any, c any, d any, _ interface{}) any { return fn.(func(any, any, any, any) any)(a, b, c, d) }
func RunSTFn5(fn any, a any, b any, c any, d any, e any, _ interface{}) any { return fn.(func(any, any, any, any, any) any)(a, b, c, d, e) }
func RunSTFn6(fn any, a any, b any, c any, d any, e any, f any, _ interface{}) any { return fn.(func(any, any, any, any, any, any) any)(a, b, c, d, e, f) }
func RunSTFn7(fn any, a any, b any, c any, d any, e any, f any, g any, _ interface{}) any { return fn.(func(any, any, any, any, any, any, any) any)(a, b, c, d, e, f, g) }
func RunSTFn8(fn any, a any, b any, c any, d any, e any, f any, g any, h any, _ interface{}) any { return fn.(func(any, any, any, any, any, any, any, any) any)(a, b, c, d, e, f, g, h) }
func RunSTFn9(fn any, a any, b any, c any, d any, e any, f any, g any, h any, i any, _ interface{}) any { return fn.(func(any, any, any, any, any, any, any, any, any) any)(a, b, c, d, e, f, g, h, i) }
// func RunSTFn10(fn any, a any, b any, c any, d any, e any, f any, g any, h any, i any, j any, _ interface{}) any { return fn.(func(any, any, any, any, any, any, any, any, any, any) any)(a, b, c, d, e, f, g, h, i, j) }


// --- Auto-generated FFI wrappers ---
func Call_mkSTFn1(arg0 func(any) any) any {
	return MkSTFn1(arg0)
}
var _Gopurs_MkSTFn1 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := MkSTFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn2(arg0 func(any) func(any) any) any {
	return MkSTFn2(arg0)
}
var _Gopurs_MkSTFn2 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_res := MkSTFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn3(arg0 func(any) func(any) func(any) any) any {
	return MkSTFn3(arg0)
}
var _Gopurs_MkSTFn3 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) any {
			return gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
		}
		}
		}
	go_res := MkSTFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn4(arg0 func(any) func(any) func(any) func(any) any) any {
	return MkSTFn4(arg0)
}
var _Gopurs_MkSTFn4 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) func(any) any {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) any {
			return gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
		}
		}
		}
		}
	go_res := MkSTFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn5(arg0 func(any) func(any) func(any) func(any) func(any) any) any {
	return MkSTFn5(arg0)
}
var _Gopurs_MkSTFn5 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) func(any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) func(any) func(any) any {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) func(any) any {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return func(p4_0 any) any {
			return gopurs_runtime.Apply(inner_res3, gopurs_runtime.Box(p4_0))
		}
		}
		}
		}
		}
	go_res := MkSTFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn6(arg0 func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return MkSTFn6(arg0)
}
var _Gopurs_MkSTFn6 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) func(any) func(any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) func(any) func(any) func(any) any {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) func(any) func(any) any {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return func(p4_0 any) func(any) any {
			inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Box(p4_0))
			return func(p5_0 any) any {
			return gopurs_runtime.Apply(inner_res4, gopurs_runtime.Box(p5_0))
		}
		}
		}
		}
		}
		}
	go_res := MkSTFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn7(arg0 func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return MkSTFn7(arg0)
}
var _Gopurs_MkSTFn7 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) func(any) func(any) func(any) func(any) any {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) func(any) func(any) func(any) any {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return func(p4_0 any) func(any) func(any) any {
			inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Box(p4_0))
			return func(p5_0 any) func(any) any {
			inner_res5 := gopurs_runtime.Apply(inner_res4, gopurs_runtime.Box(p5_0))
			return func(p6_0 any) any {
			return gopurs_runtime.Apply(inner_res5, gopurs_runtime.Box(p6_0))
		}
		}
		}
		}
		}
		}
		}
	go_res := MkSTFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn8(arg0 func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return MkSTFn8(arg0)
}
var _Gopurs_MkSTFn8 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) func(any) func(any) func(any) func(any) any {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return func(p4_0 any) func(any) func(any) func(any) any {
			inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Box(p4_0))
			return func(p5_0 any) func(any) func(any) any {
			inner_res5 := gopurs_runtime.Apply(inner_res4, gopurs_runtime.Box(p5_0))
			return func(p6_0 any) func(any) any {
			inner_res6 := gopurs_runtime.Apply(inner_res5, gopurs_runtime.Box(p6_0))
			return func(p7_0 any) any {
			return gopurs_runtime.Apply(inner_res6, gopurs_runtime.Box(p7_0))
		}
		}
		}
		}
		}
		}
		}
		}
	go_res := MkSTFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_mkSTFn9(arg0 func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return MkSTFn9(arg0)
}
var _Gopurs_MkSTFn9 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) func(any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) func(any) func(any) func(any) func(any) func(any) any {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return func(p4_0 any) func(any) func(any) func(any) func(any) any {
			inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Box(p4_0))
			return func(p5_0 any) func(any) func(any) func(any) any {
			inner_res5 := gopurs_runtime.Apply(inner_res4, gopurs_runtime.Box(p5_0))
			return func(p6_0 any) func(any) func(any) any {
			inner_res6 := gopurs_runtime.Apply(inner_res5, gopurs_runtime.Box(p6_0))
			return func(p7_0 any) func(any) any {
			inner_res7 := gopurs_runtime.Apply(inner_res6, gopurs_runtime.Box(p7_0))
			return func(p8_0 any) any {
			return gopurs_runtime.Apply(inner_res7, gopurs_runtime.Box(p8_0))
		}
		}
		}
		}
		}
		}
		}
		}
		}
	go_res := MkSTFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn1(arg0 any, arg1 any, arg2 interface{}) any {
	return RunSTFn1(arg0, arg1, arg2)
}
var _Gopurs_RunSTFn1 = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := RunSTFn1(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn2(arg0 any, arg1 any, arg2 any, arg3 interface{}) any {
	return RunSTFn2(arg0, arg1, arg2, arg3)
}
var _Gopurs_RunSTFn2 = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := RunSTFn2(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn3(arg0 any, arg1 any, arg2 any, arg3 any, arg4 interface{}) any {
	return RunSTFn3(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_RunSTFn3 = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := RunSTFn3(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn4(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any, arg5 interface{}) any {
	return RunSTFn4(arg0, arg1, arg2, arg3, arg4, arg5)
}
var _Gopurs_RunSTFn4 = gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_res := RunSTFn4(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn5(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any, arg5 any, arg6 interface{}) any {
	return RunSTFn5(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
var _Gopurs_RunSTFn5 = gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_res := RunSTFn5(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn6(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any, arg5 any, arg6 any, arg7 interface{}) any {
	return RunSTFn6(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
var _Gopurs_RunSTFn6 = gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_res := RunSTFn6(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn7(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any, arg5 any, arg6 any, arg7 any, arg8 interface{}) any {
	return RunSTFn7(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}
var _Gopurs_RunSTFn7 = gopurs_runtime.Func9(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_res := RunSTFn7(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn8(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any, arg5 any, arg6 any, arg7 any, arg8 any, arg9 interface{}) any {
	return RunSTFn8(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}
var _Gopurs_RunSTFn8 = gopurs_runtime.Func10(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := RunSTFn8(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9)
	return gopurs_runtime.Box(go_res)
})
func Call_runSTFn9(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any, arg5 any, arg6 any, arg7 any, arg8 any, arg9 any, arg10 interface{}) any {
	return RunSTFn9(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}
var _Gopurs_RunSTFn9 = gopurs_runtime.Func11(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value, arg10 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := RunSTFn9(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9, go_arg10)
	return gopurs_runtime.Box(go_res)
})
