package purescript

import "gopurs/output/gopurs_runtime"



func Data_Array_RangeImpl(start int64, end int64) []int64 {
	step := int64(1)
	if start > end {
		step = -1
	}
	size := (end - start) * step + 1
	result := make([]int64, size)
	i := start
	n := 0
	for i != end {
		result[n] = i
		n++
		i += step
	}
	result[n] = i
	return result
}

func Data_Array_ReplicateImpl(count int64, value interface{}) []interface{} {
	if count < 1 {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, count)
	for i := 0; i < int(count); i++ {
		result[i] = value
	}
	return result
}

func Data_Array_Length(xs []interface{}) int64 {
	return int64(len(xs))
}

func Data_Array_UnconsImpl(empty func(interface{}) interface{}, next func(interface{}, []interface{}) interface{}, xs []interface{}) interface{} {
	if len(xs) == 0 {
		return empty(nil)
	}
	head := xs[0]
	tail := make([]interface{}, len(xs)-1)
	copy(tail, xs[1:])
	return next(head, tail)
}

func Data_Array_IndexImpl(just func(interface{}) interface{}, nothing interface{}, xs []interface{}, i int64) interface{} {
	if i < 0 || int(i) >= len(xs) {
		return nothing
	}
	return just(xs[int(i)])
}

func Data_Array__UpdateAt(just func([]interface{}) interface{}, nothing interface{}, i int64, a interface{}, xs []interface{}) interface{} {
	if i < 0 || int(i) >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, len(xs))
	copy(l1, xs)
	l1[int(i)] = a
	return just(l1)
}

func Data_Array__InsertAt(just func([]interface{}) interface{}, nothing interface{}, i int64, a interface{}, xs []interface{}) interface{} {
	if i < 0 || int(i) > len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)+1)
	l1 = append(l1, xs[:int(i)]...)
	l1 = append(l1, a)
	l1 = append(l1, xs[int(i):]...)
	return just(l1)
}

func Data_Array__DeleteAt(just func([]interface{}) interface{}, nothing interface{}, i int64, xs []interface{}) interface{} {
	if i < 0 || int(i) >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)-1)
	l1 = append(l1, xs[:int(i)]...)
	l1 = append(l1, xs[int(i)+1:]...)
	return just(l1)
}

func Data_Array_Reverse(xs []interface{}) []interface{} {
	l := len(xs)
	l1 := make([]interface{}, l)
	for i := 0; i < l; i++ {
		l1[i] = xs[l-1-i]
	}
	return l1
}

func Data_Array_Concat(xss [][]interface{}) []interface{} {
	var result []interface{}
	for _, xs := range xss {
		result = append(result, xs...)
	}
	return result
}

func Data_Array_FilterImpl(f func(interface{}) bool, xs []interface{}) []interface{} {
	var result []interface{}
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func Data_Array_SliceImpl(s int64, e int64, l []interface{}) []interface{} {
	sInt := int(s)
	eInt := int(e)
	if sInt < 0 {
		sInt = len(l) + sInt
	}
	if eInt < 0 {
		eInt = len(l) + eInt
	}
	if sInt < 0 { sInt = 0 }
	if eInt > len(l) { eInt = len(l) }
	if sInt > eInt { sInt = eInt }
	
	res := make([]interface{}, eInt-sInt)
	copy(res, l[sInt:eInt])
	return res
}

func Data_Array_ZipWithImpl(f func(interface{}, interface{}) interface{}, xs []interface{}, ys []interface{}) []interface{} {
	length := len(xs)
	if len(ys) < length {
		length = len(ys)
	}
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		result[i] = f(xs[i], ys[i])
	}
	return result
}

func Data_Array_UnsafeIndexImpl(xs []interface{}, n int64) interface{} {
	return xs[int(n)]
}

func Data_Array_SortByImpl(compare func(interface{}, interface{}) interface{}, fromOrdering func(interface{}) int64, xs []interface{}) []interface{} {
	if len(xs) < 2 {
		return xs
	}
	out := make([]interface{}, len(xs))
	copy(out, xs)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			c := fromOrdering(compare(out[i], out[j]))
			if c > 0 { // GT
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func Data_Array_ScanrImpl(f func(interface{}, interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
		out[i] = acc
	}
	return out
}

func Data_Array_ScanlImpl(f func(interface{}, interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := 0; i < len(xs); i++ {
		acc = f(acc, xs[i])
		out[i] = acc
	}
	return out
}

func Data_Array_PartitionImpl(f func(interface{}) bool, xs []interface{}) map[string]interface{} {
	var yes []interface{}
	var no []interface{}
	for _, x := range xs {
		if f(x) {
			yes = append(yes, x)
		} else {
			no = append(no, x)
		}
	}
	return map[string]interface{}{
		"yes": yes,
		"no":  no,
	}
}

type consList struct {
	head interface{}
	tail interface{}
}

func Data_Array_FromFoldableImpl(foldr func(func(interface{}) func(interface{}) interface{}, interface{}, interface{}) interface{}, xsVal interface{}) []interface{} {
	var emptyList interface{} = nil

	curryCons := func(head interface{}) func(interface{}) interface{} {
		return func(tail interface{}) interface{} {
			return &consList{head: head, tail: tail}
		}
	}

	list := foldr(curryCons, emptyList, xsVal)

	var unboxAny func(interface{}) interface{}
	unboxAny = func(v interface{}) interface{} {
		if val, ok := v.(gopurs_runtime.Value); ok && val.Type == gopurs_runtime.TypeAny {
			if val.UnsafePtr != nil {
				return unboxAny(*(*any)(val.UnsafePtr))
			}
			return nil
		}
		return v
	}

	list = unboxAny(list)

	var result []interface{}
	curr, ok := list.(*consList)
	for ok && curr != nil {
		result = append(result, curr.head)
		curr, ok = unboxAny(curr.tail).(*consList)
	}
	
	if result == nil {
		return make([]interface{}, 0)
	}
	return result
}

func Data_Array_FindMapImpl(nothing interface{}, isJust func(interface{}) bool, f func(interface{}) interface{}, xs []interface{}) interface{} {
	for _, x := range xs {
		res := f(x)
		if isJust(res) {
			return res
		}
	}
	return nothing
}

func Data_Array_FindLastIndexImpl(just func(int64) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return just(int64(i))
		}
	}
	return nothing
}

func Data_Array_FindIndexImpl(just func(int64) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := 0; i < len(xs); i++ {
		if f(xs[i]) {
			return just(int64(i))
		}
	}
	return nothing
}

func Data_Array_AnyImpl(p func(interface{}) bool, xs []interface{}) bool {
	for _, x := range xs {
		if p(x) {
			return true
		}
	}
	return false
}

func Data_Array_AllImpl(p func(interface{}) bool, xs []interface{}) bool {
	for _, x := range xs {
		if !p(x) {
			return false
		}
	}
	return true
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Array__DeleteAt = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn4"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), Int, (Array (TypeVar a)), (ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a))])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]any, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := Data_Array__DeleteAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array__InsertAt = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn5"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), Int, (TypeVar a), (Array (TypeVar a)), (ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a))])]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	go_arg3 := arg3
	arg4_arr := *(*[]gopurs_runtime.Value)(arg4.UnsafePtr)
	go_arg4 := make([]any, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = v }
	go_res := Data_Array__InsertAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array__UpdateAt = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn5"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), Int, (TypeVar a), (Array (TypeVar a)), (ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a))])]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	go_arg3 := arg3
	arg4_arr := *(*[]gopurs_runtime.Value)(arg4.UnsafePtr)
	go_arg4 := make([]any, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = v }
	go_res := Data_Array__UpdateAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_AllImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a)] Boolean), (Array (TypeVar a)), Boolean]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Data_Array_AllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_AnyImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a)] Boolean), (Array (TypeVar a)), Boolean]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Data_Array_AnyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_Concat = // TAST: (ForAll [a] (Func [(Array (Array (TypeVar a)))] (Array (TypeVar a))))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([][]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[[]any](v) }
	go_res := Data_Array_Concat(go_arg0)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_FilterImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a)] Boolean), (Array (TypeVar a)), (Array (TypeVar a))]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Data_Array_FilterImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_FindIndexImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn4"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (Func [(TypeVar a)] Boolean), (Array (TypeVar a)), (ADT ["Data","Maybe","Maybe"] [Int])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int64) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]any, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := Data_Array_FindIndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_FindLastIndexImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn4"] [(ForAll [b] (Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]))), (ForAll [b] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (Func [(TypeVar a)] Boolean), (Array (TypeVar a)), (ADT ["Data","Maybe","Maybe"] [Int])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int64) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]any, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := Data_Array_FindLastIndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_FindMapImpl = // TAST: (ForAll [a, b] (ADT ["Data","Function","Uncurried","Fn4"] [(ForAll [c] (ADT ["Data","Maybe","Maybe"] [(TypeVar c)])), (ForAll [c] (Func [(ADT ["Data","Maybe","Maybe"] [(TypeVar c)])] Boolean)), (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (Array (TypeVar a)), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg2 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]any, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := Data_Array_FindMapImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_FromFoldableImpl = // TAST: (ForAll [f, a] (ADT ["Data","Function","Uncurried","Fn2"] [(ForAll [b] (Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))), (TypeApp (TypeVar f) [(TypeVar a)]), (Array (TypeVar a))]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 func(any) func(any) any, p0_1 any, p0_2 any) any {
			return gopurs_runtime.Apply3(arg0, gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := p0_0(arg)
						return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := inner_res(arg)
						return gopurs_runtime.Box(inner_res)
					})
					}), gopurs_runtime.Box(p0_1), gopurs_runtime.Box(p0_2))
		}
	go_arg1 := arg1
	go_res := Data_Array_FromFoldableImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_IndexImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn4"] [(ForAll [r] (Func [(TypeVar r)] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)]))), (ForAll [r] (ADT ["Data","Maybe","Maybe"] [(TypeVar r)])), (Array (TypeVar a)), Int, (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_arg3 := gopurs_runtime.Unbox[int64](arg3)
	go_res := Data_Array_IndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_Length = // TAST: (ForAll [a] (Func [(Array (TypeVar a))] Int))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := Data_Array_Length(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_PartitionImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(TypeVar a)] Boolean), (Array (TypeVar a)), (Record (Row [yes: (Array (TypeVar a)), no: (Array (TypeVar a))] Empty))]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Data_Array_PartitionImpl(go_arg0, go_arg1)
	return gopurs_runtime.Any(go_res)
})
var _Gopurs_Data_Array_RangeImpl = // TAST: (ADT ["Data","Function","Uncurried","Fn2"] [Int, Int, (Array Int)])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_Array_RangeImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_ReplicateImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [Int, (TypeVar a), (Array (TypeVar a))]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	go_res := Data_Array_ReplicateImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_Reverse = // TAST: (ForAll [a] (Func [(Array (TypeVar a))] (Array (TypeVar a))))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := Data_Array_Reverse(go_arg0)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_ScanlImpl = // TAST: (ForAll [a, b] (ADT ["Data","Function","Uncurried","Fn3"] [(Func [(TypeVar b), (TypeVar a)] (TypeVar b)), (TypeVar b), (Array (TypeVar a)), (Array (TypeVar b))]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := Data_Array_ScanlImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_ScanrImpl = // TAST: (ForAll [a, b] (ADT ["Data","Function","Uncurried","Fn3"] [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (Array (TypeVar a)), (Array (TypeVar b))]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := Data_Array_ScanrImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_SliceImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn3"] [Int, Int, (Array (TypeVar a)), (Array (TypeVar a))]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := Data_Array_SliceImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_SortByImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn3"] [(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] [])), (Func [(ADT ["Data","Ordering","Ordering"] [])] Int), (Array (TypeVar a)), (Array (TypeVar a))]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := func(p0_0 any) int64 {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[int64](inner_res0)
		}
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := Data_Array_SortByImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_Array_UnconsImpl = // TAST: (ForAll [a, b] (ADT ["Data","Function","Uncurried","Fn3"] [(Func [Unit] (TypeVar b)), (Func [(TypeVar a), (Array (TypeVar a))] (TypeVar b)), (Array (TypeVar a)), (TypeVar b)]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 any, p0_1 []any) any {
			return gopurs_runtime.Apply2(arg1, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := Data_Array_UnconsImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_UnsafeIndexImpl = // TAST: (ForAll [a] (ADT ["Data","Function","Uncurried","Fn2"] [(Array (TypeVar a)), Int, (TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_Array_UnsafeIndexImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Array_ZipWithImpl = // TAST: (ForAll [a, b, c] (ADT ["Data","Function","Uncurried","Fn3"] [(Func [(TypeVar a), (TypeVar b)] (TypeVar c)), (Array (TypeVar a)), (Array (TypeVar b)), (Array (TypeVar c))]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := Data_Array_ZipWithImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})