package Data_Traversable


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
			if vx, ok := xsVal.(interface{}); ok {
				if arr, ok := vx.([]interface{}); ok {
					xs = make([]interface{}, len(arr))
					for i, x := range arr { xs[i] = x }
				}
			} else if arr, ok := xsVal.([]interface{}); ok {
				xs = arr
			} else {
				xs = xsVal.([]interface{})
			}

			if vy, ok := ysVal.(interface{}); ok {
				if arr, ok := vy.([]interface{}); ok {
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
