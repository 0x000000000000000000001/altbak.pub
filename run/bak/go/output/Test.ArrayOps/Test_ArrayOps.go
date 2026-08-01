package Test_ArrayOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
)

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_go__range(start_0_box.IntVal, end_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_go__range
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(arr_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_filterEvens(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(arr_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_filterEvens
}

var cache_sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		cache_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sumEvens(n_0_box.IntVal))
})
	})
	return cache_sumEvens
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Array Processing (900 elements):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_sumEvens(dummy_0.IntVal))))
}))
	})
	return cache_act
}

func Call_go__range(start_0_loop int64, end_1_loop int64) []int64 {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_filterEvens(arr_0_loop []int64) []int64 {
var arr_0 []int64 = arr_0_loop
_ = arr_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), x_1, gopurs_runtime.Int(2)).IntVal) == (0))
}), func() gopurs_runtime.Value {
					arr := arr_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), x_1, gopurs_runtime.Int(2)).IntVal) == (0))
}), gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(n_0)))).IntVal
}


