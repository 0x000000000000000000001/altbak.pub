package Data_Array

import "gopurs/output/gopurs_runtime"



func RangeImpl(start int, end int) []int {
	step := 1
	if start > end {
		step = -1
	}
	size := (end - start) * step + 1
	result := make([]int, size)
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

func ReplicateImpl(count int, value interface{}) []interface{} {
	if count < 1 {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, count)
	for i := 0; i < count; i++ {
		result[i] = value
	}
	return result
}

func Length(xs []interface{}) int {
	return len(xs)
}

func UnconsImpl(empty func(interface{}) interface{}, next func(interface{}) func([]interface{}) interface{}, xs []interface{}) interface{} {
	if len(xs) == 0 {
		return empty(nil)
	}
	head := xs[0]
	tail := make([]interface{}, len(xs)-1)
	copy(tail, xs[1:])
	return next(head)(tail)
}

func IndexImpl(just func(interface{}) interface{}, nothing interface{}, xs []interface{}, i int) interface{} {
	if i < 0 || i >= len(xs) {
		return nothing
	}
	return just(xs[i])
}

func _UpdateAt(just func([]interface{}) interface{}, nothing interface{}, i int, a interface{}, xs []interface{}) interface{} {
	if i < 0 || i >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, len(xs))
	copy(l1, xs)
	l1[i] = a
	return just(l1)
}

func _InsertAt(just func([]interface{}) interface{}, nothing interface{}, i int, a interface{}, xs []interface{}) interface{} {
	if i < 0 || i > len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)+1)
	l1 = append(l1, xs[:i]...)
	l1 = append(l1, a)
	l1 = append(l1, xs[i:]...)
	return just(l1)
}

func _DeleteAt(just func([]interface{}) interface{}, nothing interface{}, i int, xs []interface{}) interface{} {
	if i < 0 || i >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)-1)
	l1 = append(l1, xs[:i]...)
	l1 = append(l1, xs[i+1:]...)
	return just(l1)
}

func Reverse(xs []interface{}) []interface{} {
	l := len(xs)
	l1 := make([]interface{}, l)
	for i := 0; i < l; i++ {
		l1[i] = xs[l-1-i]
	}
	return l1
}

func Concat(xss [][]interface{}) []interface{} {
	var result []interface{}
	for _, xs := range xss {
		result = append(result, xs...)
	}
	return result
}

func FilterImpl(f func(interface{}) bool, xs []interface{}) []interface{} {
	var result []interface{}
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func SliceImpl(s int, e int, l []interface{}) []interface{} {
	if s < 0 {
		s = len(l) + s
	}
	if e < 0 {
		e = len(l) + e
	}
	if s < 0 { s = 0 }
	if e > len(l) { e = len(l) }
	if s > e { s = e }
	
	res := make([]interface{}, e-s)
	copy(res, l[s:e])
	return res
}

func ZipWithImpl(f func(interface{}) func(interface{}) interface{}, xs []interface{}, ys []interface{}) []interface{} {
	length := len(xs)
	if len(ys) < length {
		length = len(ys)
	}
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		result[i] = f(xs[i])(ys[i])
	}
	return result
}

func UnsafeIndexImpl(xs []interface{}, n int) interface{} {
	return xs[n]
}

func SortByImpl(compare func(interface{}) func(interface{}) interface{}, fromOrdering func(interface{}) int, xs []interface{}) []interface{} {
	if len(xs) < 2 {
		return xs
	}
	out := make([]interface{}, len(xs))
	copy(out, xs)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			c := fromOrdering(compare(out[i])(out[j]))
			if c > 0 { // GT
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func ScanrImpl(f func(interface{}) func(interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i])(acc)
		out[i] = acc
	}
	return out
}

func ScanlImpl(f func(interface{}) func(interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := 0; i < len(xs); i++ {
		acc = f(acc)(xs[i])
		out[i] = acc
	}
	return out
}

func PartitionImpl(f func(interface{}) bool, xs []interface{}) map[string]interface{} {
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

func FromFoldableImpl(foldr interface{}, xsVal interface{}) []interface{} {
	panic("Not implemented: FromFoldableImpl (complex callback)")
}

func FindMapImpl(nothing interface{}, isJust func(interface{}) bool, f func(interface{}) interface{}, xs []interface{}) interface{} {
	for _, x := range xs {
		res := f(x)
		if isJust(res) {
			return res
		}
	}
	return nothing
}

func FindLastIndexImpl(just func(int) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return just(i)
		}
	}
	return nothing
}

func FindIndexImpl(just func(int) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := 0; i < len(xs); i++ {
		if f(xs[i]) {
			return just(i)
		}
	}
	return nothing
}

func AnyImpl(p func(interface{}) bool, xs []interface{}) bool {
	for _, x := range xs {
		if p(x) {
			return true
		}
	}
	return false
}

func AllImpl(p func(interface{}) bool, xs []interface{}) bool {
	for _, x := range xs {
		if !p(x) {
			return false
		}
	}
	return true
}


// --- Auto-generated FFI wrappers ---
func Call_rangeImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := RangeImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_RangeImpl = gopurs_runtime.Func2(Call_rangeImpl)
func Call_replicateImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := arg1
	go_res := ReplicateImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_ReplicateImpl = gopurs_runtime.Func2(Call_replicateImpl)
func Call_length(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := Length(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_Length = gopurs_runtime.Func(Call_length)
func Call_unconsImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 interface{}) func([]interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 []interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := UnconsImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_UnconsImpl = gopurs_runtime.Func3(Call_unconsImpl)
func Call_indexImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_res := IndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_IndexImpl = gopurs_runtime.Func4(Call_indexImpl)
func Call__UpdateAt(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := arg3
	arg4_arr := arg4.PtrVal.([]gopurs_runtime.Value)
	go_arg4 := make([]interface{}, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = v }
	go_res := _UpdateAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__UpdateAt = gopurs_runtime.Func5(Call__UpdateAt)
func Call__InsertAt(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := arg3
	arg4_arr := arg4.PtrVal.([]gopurs_runtime.Value)
	go_arg4 := make([]interface{}, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = v }
	go_res := _InsertAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__InsertAt = gopurs_runtime.Func5(Call__InsertAt)
func Call__DeleteAt(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	arg3_arr := arg3.PtrVal.([]gopurs_runtime.Value)
	go_arg3 := make([]interface{}, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := _DeleteAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__DeleteAt = gopurs_runtime.Func4(Call__DeleteAt)
func Call_reverse(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_res := Reverse(go_arg0)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_Reverse = gopurs_runtime.Func(Call_reverse)
func Call_concat(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([][]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[[]interface{}](v) }
	go_res := Concat(go_arg0)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_Concat = gopurs_runtime.Func(Call_concat)
func Call_filterImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := FilterImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_FilterImpl = gopurs_runtime.Func2(Call_filterImpl)
func Call_sliceImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := SliceImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_SliceImpl = gopurs_runtime.Func3(Call_sliceImpl)
func Call_zipWithImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := ZipWithImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_ZipWithImpl = gopurs_runtime.Func3(Call_zipWithImpl)
func Call_unsafeIndexImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([]interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := UnsafeIndexImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_UnsafeIndexImpl = gopurs_runtime.Func2(Call_unsafeIndexImpl)
func Call_sortByImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := func(p0_0 interface{}) int {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[int](inner_res0)
		}
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := SortByImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_SortByImpl = gopurs_runtime.Func3(Call_sortByImpl)
func Call_scanrImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := arg1
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := ScanrImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_ScanrImpl = gopurs_runtime.Func3(Call_scanrImpl)
func Call_scanlImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := arg1
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := ScanlImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_ScanlImpl = gopurs_runtime.Func3(Call_scanlImpl)
func Call_partitionImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := PartitionImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_map := make(map[string]gopurs_runtime.Value)
			for k, v := range go_res { res_map[k] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Record(res_map)
		}()
}
var _Gopurs_PartitionImpl = gopurs_runtime.Func2(Call_partitionImpl)
func Call_fromFoldableImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := FromFoldableImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_FromFoldableImpl = gopurs_runtime.Func2(Call_fromFoldableImpl)
func Call_findMapImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg2 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	arg3_arr := arg3.PtrVal.([]gopurs_runtime.Value)
	go_arg3 := make([]interface{}, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := FindMapImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_FindMapImpl = gopurs_runtime.Func4(Call_findMapImpl)
func Call_findLastIndexImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg3_arr := arg3.PtrVal.([]gopurs_runtime.Value)
	go_arg3 := make([]interface{}, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := FindLastIndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_FindLastIndexImpl = gopurs_runtime.Func4(Call_findLastIndexImpl)
func Call_findIndexImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg3_arr := arg3.PtrVal.([]gopurs_runtime.Value)
	go_arg3 := make([]interface{}, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = v }
	go_res := FindIndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_FindIndexImpl = gopurs_runtime.Func4(Call_findIndexImpl)
func Call_anyImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := AnyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_AnyImpl = gopurs_runtime.Func2(Call_anyImpl)
func Call_allImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := AllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_AllImpl = gopurs_runtime.Func2(Call_allImpl)
