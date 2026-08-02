package Data_Traversable

import "gopurs/output/gopurs_runtime"


func TraverseArrayImpl(
	apply func(interface{}, interface{}) interface{},
	mapFn func(func(interface{}) interface{}, interface{}) interface{},
	pure func(interface{}) interface{},
	concat2 func(interface{}) func(interface{}) interface{},
	f func(interface{}) interface{},
	arrayVal []interface{},
) interface{} {
	array1 := func(a interface{}) interface{} {
		return []interface{}{a}
	}
	
	array2 := func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return []interface{}{a, b}
		}
	}
	
	array3 := func(a interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return []interface{}{a, b, c}
			}
		}
	}
	
	var goFn func(int, int) interface{}
	goFn = func(bot, top int) interface{} {
		switch top - bot {
		case 0:
			return pure([]interface{}{})
		case 1:
			return mapFn(array1, f(arrayVal[bot]))
		case 2:
			return apply(mapFn(array2, f(arrayVal[bot])), f(arrayVal[bot+1]))
		case 3:
			return apply(apply(mapFn(array3, f(arrayVal[bot])), f(arrayVal[bot+1])), f(arrayVal[bot+2]))
		default:
			pivot := bot + ((top - bot) / 4) * 2
			return apply(mapFn(func(x interface{}) interface{} {
				return concat2(x)
			}, goFn(bot, pivot)), goFn(pivot, top))
		}
	}
	
	return goFn(0, len(arrayVal))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_TraverseArrayImpl = // TAST: (Func [(Func [(TypeVar m), (TypeVar m)] (TypeVar m)), (Func [(Func [(TypeVar x)] (TypeVar y)), (TypeVar m)] (TypeVar m)), (Func [(TypeVar x)] (TypeVar m)), (Func [(Array (TypeVar b)), (Array (TypeVar b))] (Array (TypeVar b))), (Func [(TypeVar a)] (TypeVar m)), (Array (TypeVar a))] (TypeVar m))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := func(p0_0 func(any) any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg1, gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := p0_0(arg)
						return gopurs_runtime.Box(inner_res)
					}), gopurs_runtime.Box(p0_1))
		}
	go_arg2 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	go_arg3 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg4 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg4, gopurs_runtime.Box(p0_0))
		}
	arg5_arr := *(*[]gopurs_runtime.Value)(arg5.UnsafePtr)
	go_arg5 := make([]any, len(arg5_arr))
	for i, v := range arg5_arr { go_arg5[i] = v }
	go_res := TraverseArrayImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})