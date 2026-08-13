package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_ArrayOps_add gopurs_runtime.Value
var once_Test_ArrayOps_add sync.Once
func Get_Test_ArrayOps_add() gopurs_runtime.Value {
	once_Test_ArrayOps_add.Do(func() {
		cache_Test_ArrayOps_add = Get_Data_Semiring_intAdd()
	})
	return cache_Test_ArrayOps_add
}

var cache_Test_ArrayOps_logShow gopurs_runtime.Value
var once_Test_ArrayOps_logShow sync.Once
func Get_Test_ArrayOps_logShow() gopurs_runtime.Value {
	once_Test_ArrayOps_logShow.Do(func() {
		cache_Test_ArrayOps_logShow = gopurs_runtime.Apply(Get_Effect_Console_logShow(), Get_Data_Show_showInt())
	})
	return cache_Test_ArrayOps_logShow
}

var cache_Test_ArrayOps_go__range gopurs_runtime.Value
var once_Test_ArrayOps_go__range sync.Once
func Get_Test_ArrayOps_go__range() gopurs_runtime.Value {
	once_Test_ArrayOps_go__range.Do(func() {
		cache_Test_ArrayOps_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Test_ArrayOps_go__range(start_0_box.IntVal, end_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Test_ArrayOps_go__range
}

var cache_Test_ArrayOps_filterEvens gopurs_runtime.Value
var once_Test_ArrayOps_filterEvens sync.Once
func Get_Test_ArrayOps_filterEvens() gopurs_runtime.Value {
	once_Test_ArrayOps_filterEvens.Do(func() {
		cache_Test_ArrayOps_filterEvens = gopurs_runtime.Func(func(arr_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Test_ArrayOps_filterEvens(func() []int64 {
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
	return cache_Test_ArrayOps_filterEvens
}

var cache_Test_ArrayOps_sumEvens gopurs_runtime.Value
var once_Test_ArrayOps_sumEvens sync.Once
func Get_Test_ArrayOps_sumEvens() gopurs_runtime.Value {
	once_Test_ArrayOps_sumEvens.Do(func() {
		cache_Test_ArrayOps_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_ArrayOps_sumEvens(n_0_box.IntVal))
})
	})
	return cache_Test_ArrayOps_sumEvens
}

var cache_Test_ArrayOps_describe gopurs_runtime.Value
var once_Test_ArrayOps_describe sync.Once
func Get_Test_ArrayOps_describe() gopurs_runtime.Value {
	once_Test_ArrayOps_describe.Do(func() {
		cache_Test_ArrayOps_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Array Processing (900 elements):"))
	})
	return cache_Test_ArrayOps_describe
}

var cache_Test_ArrayOps_act gopurs_runtime.Value
var once_Test_ArrayOps_act sync.Once
func Get_Test_ArrayOps_act() gopurs_runtime.Value {
	once_Test_ArrayOps_act.Do(func() {
		cache_Test_ArrayOps_act = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(900))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Show_showInt(), "show"), gopurs_runtime.Int(gopurs_runtime.Int(Call_Test_ArrayOps_sumEvens(__local_var_1_1.IntVal)).IntVal)).StrVal())), gopurs_runtime.Value{})
})
}()
	})
	return cache_Test_ArrayOps_act
}

func Call_Test_ArrayOps_go__range(start_0_loop int64, end_1_loop int64) []int64 {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Test_ArrayOps_filterEvens(arr_0_loop []int64) []int64 {
var arr_0 []int64 = arr_0_loop
_ = arr_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_filterImpl(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(x_1.IntVal), gopurs_runtime.Int(2)).IntVal) == (0))
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

func Call_Test_ArrayOps_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), Get_Data_Semiring_intAdd(), gopurs_runtime.Int(0), func() gopurs_runtime.Value {
					arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_filterImpl(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(x_1.IntVal), gopurs_runtime.Int(2)).IntVal) == (0))
}), func() gopurs_runtime.Value {
					arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(n_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal
}


