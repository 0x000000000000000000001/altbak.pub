package Effect

import "gopurs/output/gopurs_runtime"


func PureE(a any, _ interface{}) any {
	return a
}
func BindE(a func(interface{}) any, f func(any) func(interface{}) any, _ interface{}) any {
	resA := a(nil)
	return f(resA)(nil)
}
func getBool(v any) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	return v.(gopurs_runtime.Value).BoolVal()
}

func UntilE(f func(interface{}) any, _ interface{}) any {
	for {
		if getBool(f(nil)) {
			break
		}
	}
	return nil
}
func WhileE(f func(interface{}) any, a func(interface{}) any, _ interface{}) any {
	for {
		if !getBool(f(nil)) {
			break
		}
		a(nil)
	}
	return nil
}
func ForE(lo int64, hi int64, f func(any) func(interface{}) any, _ interface{}) any {
	for i := lo; i < hi; i++ {
		f(i)(nil)
	}
	return nil
}
func ForeachE(as []any, f func(any) func(interface{}) any, _ interface{}) any {
	for _, a := range as {
		f(a)(nil)
	}
	return nil
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_BindE = // TAST: (ForAll [a, b] (Func [(ADT ["Effect","Effect"] [(TypeVar a)]), (Func [(TypeVar a)] (ADT ["Effect","Effect"] [(TypeVar b)]))] (ADT ["Effect","Effect"] [(TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := BindE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ForE = // TAST: (Func [Int, Int, (Func [Int] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg3 := arg3
	go_res := ForE(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ForeachE = // TAST: (ForAll [a] (Func [(Array (TypeVar a)), (Func [(TypeVar a)] (ADT ["Effect","Effect"] [Unit]))] (ADT ["Effect","Effect"] [Unit])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := arg2
	go_res := ForeachE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PureE = // TAST: (ForAll [a] (Func [(TypeVar a)] (ADT ["Effect","Effect"] [(TypeVar a)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := PureE(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_UntilE = // TAST: (Func [(ADT ["Effect","Effect"] [Boolean])] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_res := UntilE(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_WhileE = // TAST: (ForAll [a] (Func [(ADT ["Effect","Effect"] [Boolean]), (ADT ["Effect","Effect"] [(TypeVar a)])] (ADT ["Effect","Effect"] [Unit])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_res := WhileE(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})