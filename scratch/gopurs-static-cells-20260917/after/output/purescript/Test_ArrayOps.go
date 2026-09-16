package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_ArrayOps_add gopurs_runtime.Value
var once_Test_ArrayOps_add sync.Once
func Get_Test_ArrayOps_add() gopurs_runtime.Value {
	once_Test_ArrayOps_add.Do(func() {
		cache_Test_ArrayOps_add = Call_Data_Semiring_add(Rebox_Test_ArrayOps_348932501_2826095630(Rebox_Test_ArrayOps_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))
	})
	return cache_Test_ArrayOps_add
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
		cache_Test_ArrayOps_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(int64(900)))
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_ArrayOps_sumEvens(__local_var_1_1.IntVal))).StrVal())
})
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
arr_val_filterImpl0 := arr_0
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]int64, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.IntMod(x_1.IntVal, int64(2))) == (int64(0)))
}), gopurs_runtime.Int(v_filterImpl0)).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()
}

func Call_Test_ArrayOps_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply3(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray())), Call_Data_Semiring_add(Rebox_Test_ArrayOps_348932501_2826095630(Rebox_Test_ArrayOps_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt())))), gopurs_runtime.Int(int64(0)), func() gopurs_runtime.Value {
					arr := func() []int64 {
arr_val_filterImpl1 := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(n_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
_ = arr_val_filterImpl1
_ = arr_val_filterImpl1
arr_go_filterImpl1 := arr_val_filterImpl1
_ = arr_go_filterImpl1
res_go_filterImpl1 := make([]int64, 0)
_ = res_go_filterImpl1
for _, v_filterImpl1 := range arr_go_filterImpl1 {
if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.IntMod(x_1.IntVal, int64(2))) == (int64(0)))
}), gopurs_runtime.Int(v_filterImpl1)).BoolVal() {
res_go_filterImpl1 = append(res_go_filterImpl1, v_filterImpl1)
} else {

}
}
return res_go_filterImpl1
}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal
}

func Rebox_Test_ArrayOps_2826095630_348932501(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2.IntVal
		out.V3 = in.V3.IntVal
	return out
}

func Rebox_Test_ArrayOps_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Int(in.V2)
		out.V3 = gopurs_runtime.Int(in.V3)
	return out
}


