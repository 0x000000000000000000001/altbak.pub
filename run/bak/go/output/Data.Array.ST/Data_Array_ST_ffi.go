package Data_Array_ST

import "gopurs/output/gopurs_runtime"


func NewImpl(s interface{}) *[]interface{} {
	arr := make([]interface{}, 0)
	return &arr
}

func PeekImpl(just func(interface{}) interface{}, nothing interface{}, i int64, arr *[]interface{}) interface{} {
	if i >= 0 && i < int64(len(*arr)) {
		return just((*arr)[i])
	}
	return nothing
}

func PokeImpl(i int64, a interface{}, arr *[]interface{}) interface{} {
	if i >= 0 && i < int64(len(*arr)) {
		(*arr)[i] = a
		return true
	}
	return false
}

func LengthImpl(arr *[]interface{}) interface{} {
	return int64(len(*arr))
}

func PopImpl(just func(interface{}) interface{}, nothing interface{}, arr *[]interface{}) interface{} {
	if len(*arr) > 0 {
		last := (*arr)[len(*arr)-1]
		*arr = (*arr)[:len(*arr)-1]
		return just(last)
	}
	return nothing
}

func PushAllImpl(xs []interface{}, arr *[]interface{}) interface{} {
	*arr = append(*arr, xs...)
	return int64(len(*arr))
}

func PushImpl(x interface{}, arr *[]interface{}) interface{} {
	*arr = append(*arr, x)
	return int64(len(*arr))
}

func ShiftImpl(just func(interface{}) interface{}, nothing interface{}, arr *[]interface{}) interface{} {
	if len(*arr) > 0 {
		first := (*arr)[0]
		*arr = (*arr)[1:]
		return just(first)
	}
	return nothing
}

func UnshiftAllImpl(xs []interface{}, arr *[]interface{}) interface{} {
	*arr = append(xs, *arr...)
	return int64(len(*arr))
}

func UnshiftImpl(x interface{}, arr *[]interface{}) interface{} {
	*arr = append([]interface{}{x}, *arr...)
	return int64(len(*arr))
}

func SpliceImpl(start int64, count int64, xs []interface{}, arr *[]interface{}) interface{} {
	removed := make([]interface{}, count)
	copy(removed, (*arr)[start:start+count])
	
	newArr := make([]interface{}, 0, len(*arr) - int(count) + len(xs))
	newArr = append(newArr, (*arr)[:start]...)
	newArr = append(newArr, xs...)
	newArr = append(newArr, (*arr)[start+count:]...)
	*arr = newArr
	return removed
}

func UnsafeFreezeImpl(arr *[]interface{}) []interface{} {
	return *arr
}

func UnsafeThawImpl(xs []interface{}) *[]interface{} {
	return &xs
}

func FreezeImpl(arr *[]interface{}) []interface{} {
	return *arr
}

func ThawImpl(xs []interface{}) *[]interface{} {
	return &xs
}

func CloneImpl(arr *[]interface{}) *[]interface{} {
	res := make([]interface{}, len(*arr))
	copy(res, *arr)
	return &res
}

func SortByImpl(f func(interface{}, interface{}) interface{}, toInt func(interface{}) int64, arr *[]interface{}) interface{} {
	panic("Not implemented: sortByImpl")
}

func ToAssocArrayImpl(arr *[]interface{}) interface{} {
	panic("Not implemented: toAssocArrayImpl")
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_CloneImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*[]interface{}](arg0)
	go_res := CloneImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FreezeImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (TypeVar a))])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*[]interface{}](arg0)
	go_res := FreezeImpl(go_arg0)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_LengthImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*[]interface{}](arg0)
	go_res := LengthImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NewImpl = // TAST: (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := NewImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PeekImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn4"] [(Func [(TypeVar a)] (TypeVar r)), (TypeVar r), Int, (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (TypeVar r)])
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	go_arg3 := gopurs_runtime.Unbox[*[]interface{}](arg3)
	go_res := PeekImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PokeImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn3"] [Int, (TypeVar a), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Boolean])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*[]interface{}](arg2)
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PopImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*[]interface{}](arg2)
	go_res := PopImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PushAllImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(Array (TypeVar a)), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := gopurs_runtime.Unbox[*[]interface{}](arg1)
	go_res := PushAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_PushImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(TypeVar a), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[*[]interface{}](arg1)
	go_res := PushImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShiftImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)]), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*[]interface{}](arg2)
	go_res := ShiftImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_SortByImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] [])), (Func [(ADT ["Data","Ordering","Ordering"] [])] Int), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := func(p0_0 any) int64 {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[int64](inner_res0)
		}
	go_arg2 := gopurs_runtime.Unbox[*[]interface{}](arg2)
	go_res := SortByImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_SpliceImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn4"] [Int, Int, (Array (TypeVar a)), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (TypeVar a))])
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_arg3 := gopurs_runtime.Unbox[*[]interface{}](arg3)
	go_res := SpliceImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ThawImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(Array (TypeVar a)), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := ThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToAssocArrayImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array Any)])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*[]interface{}](arg0)
	go_res := ToAssocArrayImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_UnsafeFreezeImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), (Array (TypeVar a))])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*[]interface{}](arg0)
	go_res := UnsafeFreezeImpl(go_arg0)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_UnsafeThawImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(Array (TypeVar a)), (TypeVar h), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := UnsafeThawImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_UnshiftAllImpl = // TAST: (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(Array (TypeVar a)), (ADT ["Data","Array","ST","STArray"] [(TypeVar h), (TypeVar a)]), (TypeVar h), Int])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := gopurs_runtime.Unbox[*[]interface{}](arg1)
	go_res := UnshiftAllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})