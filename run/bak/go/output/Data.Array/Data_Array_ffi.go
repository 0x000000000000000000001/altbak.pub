package Data_Array

import "gopurs/output/gopurs_runtime"

func RangeImpl(start int64, end int64) []int64 {
	var step int64 = 1
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

func ReplicateImpl[T any](count int64, value T) []T {
	if count < 1 {
		return make([]T, 0)
	}
	result := make([]T, count)
	for i := int64(0); i < count; i++ {
		result[i] = value
	}
	return result
}

func Length[T any](xs []T) int64 {
	return int64(len(xs))
}

func UnconsImpl[T any](empty func(interface{}) interface{}, next func(T) func([]T) interface{}, xs []T) interface{} {
	if len(xs) == 0 {
		return empty(nil)
	}
	head := xs[0]
	tail := make([]T, len(xs)-1)
	copy(tail, xs[1:])
	return next(head)(tail)
}

func IndexImpl[T any](just func(T) interface{}, nothing interface{}, xs []T, i int64) interface{} {
	if i < 0 || i >= int64(len(xs)) {
		return nothing
	}
	return just(xs[i])
}

func _UpdateAt[T any](just func([]T) interface{}, nothing interface{}, i int64, a T, xs []T) interface{} {
	if i < 0 || i >= int64(len(xs)) {
		return nothing
	}
	l1 := make([]T, len(xs))
	copy(l1, xs)
	l1[i] = a
	return just(l1)
}

func _InsertAt[T any](just func([]T) interface{}, nothing interface{}, i int64, a T, xs []T) interface{} {
	if i < 0 || i > int64(len(xs)) {
		return nothing
	}
	l1 := make([]T, 0, len(xs)+1)
	l1 = append(l1, xs[:i]...)
	l1 = append(l1, a)
	l1 = append(l1, xs[i:]...)
	return just(l1)
}

func _DeleteAt[T any](just func([]T) interface{}, nothing interface{}, i int64, xs []T) interface{} {
	if i < 0 || i >= int64(len(xs)) {
		return nothing
	}
	l1 := make([]T, 0, len(xs)-1)
	l1 = append(l1, xs[:i]...)
	l1 = append(l1, xs[i+1:]...)
	return just(l1)
}

func Reverse[T any](xs []T) []T {
	l := len(xs)
	l1 := make([]T, l)
	for i := 0; i < l; i++ {
		l1[i] = xs[l-1-i]
	}
	return l1
}

func Concat[T any](xss [][]T) []T {
	var result []T
	for _, xs := range xss {
		result = append(result, xs...)
	}
	return result
}

func FilterImpl[T any](f func(T) bool, xs []T) []T {
	var result []T
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func SliceImpl[T any](s int64, e int64, l []T) []T {
	if s < 0 {
		s = int64(len(l)) + s
	}
	if e < 0 {
		e = int64(len(l)) + e
	}
	if s < 0 { s = 0 }
	if e > int64(len(l)) { e = int64(len(l)) }
	if s > e { s = e }
	
	res := make([]T, e-s)
	copy(res, l[s:e])
	return res
}

func ZipWithImpl[T1 any, T2 any, T3 any](f func(T1) func(T2) T3, xs []T1, ys []T2) []T3 {
	length := len(xs)
	if len(ys) < length {
		length = len(ys)
	}
	result := make([]T3, length)
	for i := 0; i < length; i++ {
		result[i] = f(xs[i])(ys[i])
	}
	return result
}

func UnsafeIndexImpl[T any](xs []T, n int64) T {
	return xs[n]
}

func SortByImpl[T any](compare func(T) func(T) interface{}, fromOrdering func(interface{}) int64, xs []T) []T {
	if len(xs) < 2 {
		return xs
	}
	out := make([]T, len(xs))
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

func ScanrImpl[T1 any, T2 any](f func(T1) func(T2) T2, b T2, xs []T1) []T2 {
	out := make([]T2, len(xs))
	acc := b
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i])(acc)
		out[i] = acc
	}
	return out
}

func ScanlImpl[T1 any, T2 any](f func(T2) func(T1) T2, b T2, xs []T1) []T2 {
	out := make([]T2, len(xs))
	acc := b
	for i := 0; i < len(xs); i++ {
		acc = f(acc)(xs[i])
		out[i] = acc
	}
	return out
}

func PartitionImpl[T any](f func(T) bool, xs []T) map[string]interface{} {
	var yes []T
	var no []T
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

func FindMapImpl[T1 any, T2 any](nothing interface{}, isJust func(interface{}) bool, f func(T1) interface{}, xs []T1) interface{} {
	for _, x := range xs {
		res := f(x)
		if isJust(res) {
			return res
		}
	}
	return nothing
}

func FindLastIndexImpl[T any](just func(int64) interface{}, nothing interface{}, f func(T) bool, xs []T) interface{} {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return just(int64(i))
		}
	}
	return nothing
}

func FindIndexImpl[T any](just func(int64) interface{}, nothing interface{}, f func(T) bool, xs []T) interface{} {
	for i := 0; i < len(xs); i++ {
		if f(xs[i]) {
			return just(int64(i))
		}
	}
	return nothing
}

func AnyImpl[T any](p func(T) bool, xs []T) bool {
	for _, x := range xs {
		if p(x) {
			return true
		}
	}
	return false
}

func AllImpl[T any](p func(T) bool, xs []T) bool {
	for _, x := range xs {
		if !p(x) {
			return false
		}
	}
	return true
}


// --- Auto-generated FFI wrappers ---
func Call_rangeImpl(arg0 int64, arg1 int64) []int64 {
	return RangeImpl(arg0, arg1)
}
var _Gopurs_RangeImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := RangeImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_replicateImpl[T any](arg0 int64, arg1 T) []T {
	return ReplicateImpl(arg0, arg1)
}
var _Gopurs_ReplicateImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[T](arg1)
	go_res := ReplicateImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_length[T any](arg0 []T) int64 {
	return Length(arg0)
}
var _Gopurs_Length = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]T, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[T](v) }
	go_res := Length(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_unconsImpl[T any](arg0 func(interface{}) interface{}, arg1 func(T) func([]T) interface{}, arg2 []T) interface{} {
	return UnconsImpl(arg0, arg1, arg2)
}
var _Gopurs_UnconsImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := func(p0_0 T) func([]T) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 []T) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T](v) }
	go_res := UnconsImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_indexImpl[T any](arg0 func(T) interface{}, arg1 interface{}, arg2 []T, arg3 int64) interface{} {
	return IndexImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_IndexImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T](v) }
	go_arg3 := gopurs_runtime.Unbox[int64](arg3)
	go_res := IndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call__UpdateAt[T any](arg0 func([]T) interface{}, arg1 interface{}, arg2 int64, arg3 T, arg4 []T) interface{} {
	return _UpdateAt(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs__UpdateAt = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []T) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	go_arg3 := gopurs_runtime.Unbox[T](arg3)
	arg4_arr := *(*[]gopurs_runtime.Value)(arg4.UnsafePtr)
	go_arg4 := make([]T, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = gopurs_runtime.Unbox[T](v) }
	go_res := _UpdateAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call__InsertAt[T any](arg0 func([]T) interface{}, arg1 interface{}, arg2 int64, arg3 T, arg4 []T) interface{} {
	return _InsertAt(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs__InsertAt = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []T) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	go_arg3 := gopurs_runtime.Unbox[T](arg3)
	arg4_arr := *(*[]gopurs_runtime.Value)(arg4.UnsafePtr)
	go_arg4 := make([]T, len(arg4_arr))
	for i, v := range arg4_arr { go_arg4[i] = gopurs_runtime.Unbox[T](v) }
	go_res := _InsertAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call__DeleteAt[T any](arg0 func([]T) interface{}, arg1 interface{}, arg2 int64, arg3 []T) interface{} {
	return _DeleteAt(arg0, arg1, arg2, arg3)
}
var _Gopurs__DeleteAt = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 []T) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int64](arg2)
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]T, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = gopurs_runtime.Unbox[T](v) }
	go_res := _DeleteAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_reverse[T any](arg0 []T) []T {
	return Reverse(arg0)
}
var _Gopurs_Reverse = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]T, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[T](v) }
	go_res := Reverse(go_arg0)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_concat[T any](arg0 [][]T) []T {
	return Concat(arg0)
}
var _Gopurs_Concat = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([][]T, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[[]T](v) }
	go_res := Concat(go_arg0)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_filterImpl[T any](arg0 func(T) bool, arg1 []T) []T {
	return FilterImpl(arg0, arg1)
}
var _Gopurs_FilterImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]T, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[T](v) }
	go_res := FilterImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_sliceImpl[T any](arg0 int64, arg1 int64, arg2 []T) []T {
	return SliceImpl(arg0, arg1, arg2)
}
var _Gopurs_SliceImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T](v) }
	go_res := SliceImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_zipWithImpl[T1 any, T2 any, T3 any](arg0 func(T1) func(T2) T3, arg1 []T1, arg2 []T2) []T3 {
	return ZipWithImpl(arg0, arg1, arg2)
}
var _Gopurs_ZipWithImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T1) func(T2) T3 {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 T2) T3 {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[T3](inner_res1)
		}
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]T1, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[T1](v) }
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T2, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T2](v) }
	go_res := ZipWithImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_unsafeIndexImpl[T any](arg0 []T, arg1 int64) T {
	return UnsafeIndexImpl(arg0, arg1)
}
var _Gopurs_UnsafeIndexImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]T, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[T](v) }
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := UnsafeIndexImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_sortByImpl[T any](arg0 func(T) func(T) interface{}, arg1 func(interface{}) int64, arg2 []T) []T {
	return SortByImpl(arg0, arg1, arg2)
}
var _Gopurs_SortByImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T) func(T) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 T) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := func(p0_0 interface{}) int64 {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[int64](inner_res0)
		}
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T](v) }
	go_res := SortByImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_scanrImpl[T1 any, T2 any](arg0 func(T1) func(T2) T2, arg1 T2, arg2 []T1) []T2 {
	return ScanrImpl(arg0, arg1, arg2)
}
var _Gopurs_ScanrImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T1) func(T2) T2 {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 T2) T2 {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[T2](inner_res1)
		}
		}
	go_arg1 := gopurs_runtime.Unbox[T2](arg1)
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T1, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T1](v) }
	go_res := ScanrImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_scanlImpl[T1 any, T2 any](arg0 func(T2) func(T1) T2, arg1 T2, arg2 []T1) []T2 {
	return ScanlImpl(arg0, arg1, arg2)
}
var _Gopurs_ScanlImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T2) func(T1) T2 {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 T1) T2 {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[T2](inner_res1)
		}
		}
	go_arg1 := gopurs_runtime.Unbox[T2](arg1)
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]T1, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = gopurs_runtime.Unbox[T1](v) }
	go_res := ScanlImpl(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_partitionImpl[T any](arg0 func(T) bool, arg1 []T) map[string]interface{} {
	return PartitionImpl(arg0, arg1)
}
var _Gopurs_PartitionImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]T, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[T](v) }
	go_res := PartitionImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_map := make(map[string]gopurs_runtime.Value)
			for k, v := range go_res { res_map[k] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Record(res_map)
		}()
})
func Call_fromFoldableImpl(arg0 interface{}, arg1 interface{}) []interface{} {
	return FromFoldableImpl(arg0, arg1)
}
var _Gopurs_FromFoldableImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := FromFoldableImpl(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_findMapImpl[T1 any, T2 any](arg0 interface{}, arg1 func(interface{}) bool, arg2 func(T1) interface{}, arg3 []T1) interface{} {
	return FindMapImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_FindMapImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 interface{}) bool {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg2 := func(p0_0 T1) interface{} {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]T1, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = gopurs_runtime.Unbox[T1](v) }
	go_res := FindMapImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_findLastIndexImpl[T any](arg0 func(int64) interface{}, arg1 interface{}, arg2 func(T) bool, arg3 []T) interface{} {
	return FindLastIndexImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_FindLastIndexImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int64) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := func(p0_0 T) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]T, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = gopurs_runtime.Unbox[T](v) }
	go_res := FindLastIndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_findIndexImpl[T any](arg0 func(int64) interface{}, arg1 interface{}, arg2 func(T) bool, arg3 []T) interface{} {
	return FindIndexImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_FindIndexImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int64) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := func(p0_0 T) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg3_arr := *(*[]gopurs_runtime.Value)(arg3.UnsafePtr)
	go_arg3 := make([]T, len(arg3_arr))
	for i, v := range arg3_arr { go_arg3[i] = gopurs_runtime.Unbox[T](v) }
	go_res := FindIndexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_anyImpl[T any](arg0 func(T) bool, arg1 []T) bool {
	return AnyImpl(arg0, arg1)
}
var _Gopurs_AnyImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]T, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[T](v) }
	go_res := AnyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_allImpl[T any](arg0 func(T) bool, arg1 []T) bool {
	return AllImpl(arg0, arg1)
}
var _Gopurs_AllImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 T) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]T, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[T](v) }
	go_res := AllImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
