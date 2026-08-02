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
var _Gopurs_MkFn0 = // TAST: (Func [(Func [(ADT ["Data","Unit","Unit"] [])] (TypeVar a))] (ADT ["Data","Function","Uncurried","Fn0"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MkFn0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_MkFn10 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)] (TypeVar k))] (ADT ["Data","Function","Uncurried","Fn10"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j), (TypeVar k)]))
gopurs_runtime.Box(MkFn10)
var _Gopurs_MkFn2 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar c))] (ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar a), (TypeVar b), (TypeVar c)]))
gopurs_runtime.Box(MkFn2)
var _Gopurs_MkFn3 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c)] (TypeVar d))] (ADT ["Data","Function","Uncurried","Fn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)]))
gopurs_runtime.Box(MkFn3)
var _Gopurs_MkFn4 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (TypeVar e))] (ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)]))
gopurs_runtime.Box(MkFn4)
var _Gopurs_MkFn5 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (TypeVar f))] (ADT ["Data","Function","Uncurried","Fn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)]))
gopurs_runtime.Box(MkFn5)
var _Gopurs_MkFn6 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (TypeVar g))] (ADT ["Data","Function","Uncurried","Fn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)]))
gopurs_runtime.Box(MkFn6)
var _Gopurs_MkFn7 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (TypeVar h))] (ADT ["Data","Function","Uncurried","Fn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)]))
gopurs_runtime.Box(MkFn7)
var _Gopurs_MkFn8 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (TypeVar i))] (ADT ["Data","Function","Uncurried","Fn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)]))
gopurs_runtime.Box(MkFn8)
var _Gopurs_MkFn9 = // TAST: (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (TypeVar j))] (ADT ["Data","Function","Uncurried","Fn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)]))
gopurs_runtime.Box(MkFn9)
var _Gopurs_RunFn0 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn0"] [(TypeVar a)])] (TypeVar a))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := RunFn0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_RunFn10 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn10"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j), (TypeVar k)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)] (TypeVar k))
gopurs_runtime.Box(RunFn10)
var _Gopurs_RunFn2 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar a), (TypeVar b), (TypeVar c)]), (TypeVar a), (TypeVar b)] (TypeVar c))
gopurs_runtime.Box(RunFn2)
var _Gopurs_RunFn3 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)]), (TypeVar a), (TypeVar b), (TypeVar c)] (TypeVar d))
gopurs_runtime.Box(RunFn3)
var _Gopurs_RunFn4 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (TypeVar e))
gopurs_runtime.Box(RunFn4)
var _Gopurs_RunFn5 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (TypeVar f))
gopurs_runtime.Box(RunFn5)
var _Gopurs_RunFn6 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (TypeVar g))
gopurs_runtime.Box(RunFn6)
var _Gopurs_RunFn7 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (TypeVar h))
gopurs_runtime.Box(RunFn7)
var _Gopurs_RunFn8 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (TypeVar i))
gopurs_runtime.Box(RunFn8)
var _Gopurs_RunFn9 = // TAST: (Func [(ADT ["Data","Function","Uncurried","Fn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (TypeVar j))
gopurs_runtime.Box(RunFn9)