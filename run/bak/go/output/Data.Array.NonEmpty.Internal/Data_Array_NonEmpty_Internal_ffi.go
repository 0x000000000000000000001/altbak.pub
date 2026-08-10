package Data_Array_NonEmpty_Internal



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


// --- Auto-generated FFI wrappers ---
var _Gopurs_Foldl1Impl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a), (TypeVar a)] (TypeVar a)), (Array (TypeVar a)), (TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Foldl1Impl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foldr1Impl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a), (TypeVar a)] (TypeVar a)), (Array (TypeVar a)), (TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Foldr1Impl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Traverse1Impl = // TAST: (ForAll [m, a, b] (ADT ["Data","Function","Uncurried","Fn3"] [(ForAll [a', b'] (Func [(TypeApp (TypeVar m) [(Func [(TypeVar a')] (TypeVar b'))]), (TypeApp (TypeVar m) [(TypeVar a')])] (TypeApp (TypeVar m) [(TypeVar b')]))), (ForAll [a', b'] (Func [(Func [(TypeVar a')] (TypeVar b')), (TypeApp (TypeVar m) [(TypeVar a')])] (TypeApp (TypeVar m) [(TypeVar b')]))), (Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (Func [(Array (TypeVar a))] (TypeApp (TypeVar m) [(Array (TypeVar b))]))]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]any, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := Traverse1Impl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})