
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
