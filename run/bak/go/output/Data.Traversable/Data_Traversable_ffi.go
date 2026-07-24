package Data_Traversable

import "gopurs/output/gopurs_runtime"

func TraverseArrayImpl(apply func(interface{}) func(interface{}) interface{}, mapFn func(interface{}) func(interface{}) interface{}, pure func(interface{}) interface{}, f func(interface{}) interface{}, arrayVal []interface{}) interface{} {
	array1 := func(a interface{}) interface{} {
		return []interface{}{a}
	}
	
	array2 := func(a interface{}) func(interface{}) interface{} {
		return func(b interface{}) interface{} {
			return []interface{}{a, b}
		}
	}
	
	array3 := func(a interface{}) func(interface{}) interface{} {
		return func(b interface{}) interface{} {
			return func(c interface{}) interface{} {
				return []interface{}{a, b, c}
			}
		}
	}
	
	concat2 := func(xsVal interface{}) func(interface{}) interface{} {
		return func(ysVal interface{}) interface{} {
			var xs, ys []interface{}
			if vx, ok := xsVal.(gopurs_runtime.Value); ok {
				if arr, ok := vx.PtrVal.([]gopurs_runtime.Value); ok {
					xs = make([]interface{}, len(arr))
					for i, x := range arr { xs[i] = x }
				}
			} else if arr, ok := xsVal.([]interface{}); ok {
				xs = arr
			} else {
				xs = xsVal.([]interface{})
			}

			if vy, ok := ysVal.(gopurs_runtime.Value); ok {
				if arr, ok := vy.PtrVal.([]gopurs_runtime.Value); ok {
					ys = make([]interface{}, len(arr))
					for i, y := range arr { ys[i] = y }
				}
			} else if arr, ok := ysVal.([]interface{}); ok {
				ys = arr
			} else {
				ys = ysVal.([]interface{})
			}

			res := make([]interface{}, 0, len(xs)+len(ys))
			res = append(res, xs...)
			res = append(res, ys...)
			return res
		}
	}
	
	var goFn func(int, int) interface{}
	goFn = func(bot, top int) interface{} {
		switch top - bot {
		case 0:
			return pure([]interface{}{})
		case 1:
			return mapFn(array1)(f(arrayVal[bot]))
		case 2:
			return apply(mapFn(array2)(f(arrayVal[bot])))(f(arrayVal[bot+1]))
		case 3:
			return apply(apply(mapFn(array3)(f(arrayVal[bot])))(f(arrayVal[bot+1])))(f(arrayVal[bot+2]))
		default:
			pivot := bot + ((top - bot) / 4) * 2
			return apply(mapFn(concat2)(goFn(bot, pivot)))(goFn(pivot, top))
		}
	}
	
	return goFn(0, len(arrayVal))
}


// --- Auto-generated FFI wrappers ---
func Call_traverseArrayImpl(arg0 func(interface{}) func(interface{}) interface{}, arg1 func(interface{}) func(interface{}) interface{}, arg2 func(interface{}) interface{}, arg3 func(interface{}) interface{}, arg4 []interface{}) interface{} {
	return TraverseArrayImpl(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_TraverseArrayImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	go_arg3 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
		}
	arg4_arr := arg4.PtrVal().([]gopurs_runtime.Value)
	go_arg4 := make([]interface{}, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = v }
	go_res := TraverseArrayImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
