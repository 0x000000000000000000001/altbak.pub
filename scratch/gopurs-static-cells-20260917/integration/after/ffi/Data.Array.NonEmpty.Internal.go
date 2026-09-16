package Internal

import "gopurs/output/gopurs_runtime"
func Foldr1Impl(f func(interface{}) func(interface{}) interface{}, xs []interface{}) interface{} {
	acc := xs[len(xs)-1]
	for i := len(xs) - 2; i >= 0; i-- {
		acc = f(xs[i])(acc)
	}
	return acc
}

func Foldl1Impl(f func(interface{}) func(interface{}) interface{}, xs []interface{}) interface{} {
	acc := xs[0]
	length := len(xs)
	for i := 1; i < length; i++ {
		acc = f(acc)(xs[i])
	}
	return acc
}

type listNode struct {
	head interface{}
	tail interface{}
}

func Traverse1Impl(apply func(interface{}) func(interface{}) interface{}, mapFn func(interface{}) func(interface{}) interface{}, f func(interface{}) interface{}, array []interface{}) interface{} {

	emptyList := &listNode{}

	consList := func(x interface{}) func(interface{}) interface{} {
		return func(xs interface{}) interface{} {
			xsNode := gopurs_runtime.Unbox[*listNode](xs.(gopurs_runtime.Value))
			return &listNode{head: x, tail: xsNode}
		}
	}

	finalCell := func(head interface{}) interface{} {
		return &listNode{head: head, tail: emptyList}
	}

	listToArray := func(list interface{}) interface{} {
		var arr []interface{}
		xs := gopurs_runtime.Unbox[*listNode](list.(gopurs_runtime.Value))
		for xs != emptyList {
			arr = append(arr, xs.head)
			xs = xs.tail.(*listNode)
		}
		if arr == nil {
			return []interface{}{}
		}
		return arr
	}

	buildFrom := func(x interface{}) func(interface{}) interface{} {
		return func(ys interface{}) interface{} {
			return apply(mapFn(consList)(f(x)))(ys)
		}
	}

	var goFn func(interface{}, int, []interface{}) interface{}
	goFn = func(acc interface{}, currentLen int, xs []interface{}) interface{} {
		if currentLen == 0 {
			return acc
		}
		last := xs[currentLen-1]
		return func() interface{} {
			return goFn(buildFrom(last)(acc), currentLen-1, xs)
		}
	}

	acc := mapFn(finalCell)(f(array[len(array)-1]))
	result := goFn(acc, len(array)-1, array)

	for {
		fn, isFunc := result.(func() interface{})
		if !isFunc {
			break
		}
		result = fn()
	}

	return mapFn(listToArray)(result)
}
