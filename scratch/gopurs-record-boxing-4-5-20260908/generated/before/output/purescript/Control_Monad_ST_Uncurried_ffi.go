package purescript

import "gopurs/output/gopurs_runtime"


func Control_Monad_ST_Uncurried_MkSTFn1(fn func(any) any) any {
	return func(a any) any { return fn(a) }
}
func Control_Monad_ST_Uncurried_MkSTFn2(fn func(any) func(any) any) any {
	return func(a any, b any) any { return fn(a)(b) }
}
func Control_Monad_ST_Uncurried_MkSTFn3(fn func(any) func(any) func(any) any) any {
	return func(a any, b any, c any) any { return fn(a)(b)(c) }
}
func Control_Monad_ST_Uncurried_MkSTFn4(fn func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any) any { return fn(a)(b)(c)(d) }
}
func Control_Monad_ST_Uncurried_MkSTFn5(fn func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any) any { return fn(a)(b)(c)(d)(e) }
}
func Control_Monad_ST_Uncurried_MkSTFn6(fn func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any) any { return fn(a)(b)(c)(d)(e)(f) }
}
func Control_Monad_ST_Uncurried_MkSTFn7(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any) any { return fn(a)(b)(c)(d)(e)(f)(g) }
}
func Control_Monad_ST_Uncurried_MkSTFn8(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any, h any) any { return fn(a)(b)(c)(d)(e)(f)(g)(h) }
}
func Control_Monad_ST_Uncurried_MkSTFn9(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	return func(a any, b any, c any, d any, e any, f any, g any, h any, i any) any { return fn(a)(b)(c)(d)(e)(f)(g)(h)(i) }
}
// func Control_Monad_ST_Uncurried_MkSTFn10(fn func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) func(any) any) any {
	// return func any { return fn(a)(b)(c)(d)(e)(f)(g)(h)(i)(j) }
// }
func Control_Monad_ST_Uncurried_RunSTFn1(fn func(any) any, a any, _ interface{}) any { return fn(a) }
func Control_Monad_ST_Uncurried_RunSTFn2(fn func(any, any) any, a any, b any, _ interface{}) any { return fn(a, b) }
func Control_Monad_ST_Uncurried_RunSTFn3(fn func(any, any, any) any, a any, b any, c any, _ interface{}) any { return fn(a, b, c) }
func Control_Monad_ST_Uncurried_RunSTFn4(fn func(any, any, any, any) any, a any, b any, c any, d any, _ interface{}) any { return fn(a, b, c, d) }
func Control_Monad_ST_Uncurried_RunSTFn5(fn func(any, any, any, any, any) any, a any, b any, c any, d any, e any, _ interface{}) any { return fn(a, b, c, d, e) }
func Control_Monad_ST_Uncurried_RunSTFn6(fn func(any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, _ interface{}) any { return fn(a, b, c, d, e, f) }
func Control_Monad_ST_Uncurried_RunSTFn7(fn func(any, any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, g any, _ interface{}) any { return fn(a, b, c, d, e, f, g) }
func Control_Monad_ST_Uncurried_RunSTFn8(fn func(any, any, any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, g any, h any, _ interface{}) any { return fn(a, b, c, d, e, f, g, h) }
func Control_Monad_ST_Uncurried_RunSTFn9(fn func(any, any, any, any, any, any, any, any, any) any, a any, b any, c any, d any, e any, f any, g any, h any, i any, _ interface{}) any { return fn(a, b, c, d, e, f, g, h, i) }
// func Control_Monad_ST_Uncurried_RunSTFn10(fn any, a any, b any, c any, d any, e any, f any, g any, h any, i any, j any, _ interface{}) any { return fn.(func(any, any, any, any, any, any, any, any, any, any) any)(a, b, c, d, e, f, g, h, i, j) }


// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn1 = // TAST: (ForAll [a, t, r] (Func [(Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(TypeVar a), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := Control_Monad_ST_Uncurried_MkSTFn1(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { panic("FFI not implemented: mkSTFn10"); return gopurs_runtime.Value{} })
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn2 = // TAST: (ForAll [a, b, t, r] (Func [(Func [(TypeVar a), (TypeVar b)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(TypeVar a), (TypeVar b), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_res := Control_Monad_ST_Uncurried_MkSTFn2(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn3 = // TAST: (ForAll [a, b, c, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) any {
			return gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
		}
		}
		}
	go_res := Control_Monad_ST_Uncurried_MkSTFn3(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn4 = // TAST: (ForAll [a, b, c, d, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := Control_Monad_ST_Uncurried_MkSTFn4(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn5 = // TAST: (ForAll [a, b, c, d, e, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := Control_Monad_ST_Uncurried_MkSTFn5(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn6 = // TAST: (ForAll [a, b, c, d, e, f, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := Control_Monad_ST_Uncurried_MkSTFn6(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := Control_Monad_ST_Uncurried_MkSTFn7(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := Control_Monad_ST_Uncurried_MkSTFn8(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	go_res := Control_Monad_ST_Uncurried_MkSTFn9(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn1 = // TAST: (ForAll [a, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn1"] [(TypeVar a), (TypeVar t), (TypeVar r)]), (TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Uncurried_RunSTFn1(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { panic("FFI not implemented: runSTFn10"); return gopurs_runtime.Value{} })
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn2 = // TAST: (ForAll [a, b, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn2"] [(TypeVar a), (TypeVar b), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Control_Monad_ST_Uncurried_RunSTFn2(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn3 = // TAST: (ForAll [a, b, c, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any) any {
			return gopurs_runtime.Apply3(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Control_Monad_ST_Uncurried_RunSTFn3(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn4 = // TAST: (ForAll [a, b, c, d, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any, p0_3 any) any {
			return gopurs_runtime.Apply4(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2), gopurs_runtime.Box(p0_3))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_res := Control_Monad_ST_Uncurried_RunSTFn4(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn5 = // TAST: (ForAll [a, b, c, d, e, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any, p0_3 any, p0_4 any) any {
			return gopurs_runtime.Apply5(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2), gopurs_runtime.Box(p0_3), gopurs_runtime.Box(p0_4))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_res := Control_Monad_ST_Uncurried_RunSTFn5(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn6 = // TAST: (ForAll [a, b, c, d, e, f, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any, p0_3 any, p0_4 any, p0_5 any) any {
			return gopurs_runtime.Apply6(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2), gopurs_runtime.Box(p0_3), gopurs_runtime.Box(p0_4), gopurs_runtime.Box(p0_5))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_res := Control_Monad_ST_Uncurried_RunSTFn6(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func9(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any, p0_3 any, p0_4 any, p0_5 any, p0_6 any) any {
			return gopurs_runtime.Apply7(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2), gopurs_runtime.Box(p0_3), gopurs_runtime.Box(p0_4), gopurs_runtime.Box(p0_5), gopurs_runtime.Box(p0_6))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_res := Control_Monad_ST_Uncurried_RunSTFn7(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func10(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any, p0_3 any, p0_4 any, p0_5 any, p0_6 any, p0_7 any) any {
			return gopurs_runtime.Apply8(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2), gopurs_runtime.Box(p0_3), gopurs_runtime.Box(p0_4), gopurs_runtime.Box(p0_5), gopurs_runtime.Box(p0_6), gopurs_runtime.Box(p0_7))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_arg9 := arg9
	go_res := Control_Monad_ST_Uncurried_RunSTFn8(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func11(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value, arg10 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any, p0_2 any, p0_3 any, p0_4 any, p0_5 any, p0_6 any, p0_7 any, p0_8 any) any {
			return gopurs_runtime.Apply9(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2), gopurs_runtime.Box(p0_3), gopurs_runtime.Box(p0_4), gopurs_runtime.Box(p0_5), gopurs_runtime.Box(p0_6), gopurs_runtime.Box(p0_7), gopurs_runtime.Box(p0_8))
		}
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
	go_res := Control_Monad_ST_Uncurried_RunSTFn9(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9, go_arg10)
	return gopurs_runtime.Box(go_res)
})