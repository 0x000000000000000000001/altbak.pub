package Test_Native

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
)

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
})
}()
	})
	return cache_greaterThan
}

var cache_sumNative gopurs_runtime.Value
var once_sumNative sync.Once
func Get_sumNative() gopurs_runtime.Value {
	once_sumNative.Do(func() {
		cache_sumNative = gopurs_runtime.Func(func(arr_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sumNative(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(arr_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))
})
	})
	return cache_sumNative
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Native Intrinsics Test (900 elements):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), func() gopurs_runtime.Value {
arr_val_foldlArray4 := func() gopurs_runtime.Value {
arr_val_filter5 := func() gopurs_runtime.Value {
arr_val_arrayMap6 := func() gopurs_runtime.Value {
arr_val_arrayMap7 := []int64{1, 2, 3}
_ = arr_val_arrayMap7
res_go_arrayMap7 := make([]gopurs_runtime.Value, len(arr_val_arrayMap7))
_ = res_go_arrayMap7
for i_arrayMap7, v_arrayMap7 := range arr_val_arrayMap7 {
res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Int(v_arrayMap7))
}
return gopurs_runtime.Array(res_go_arrayMap7)
}()
_ = arr_val_arrayMap6
arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
_ = arr_go_arrayMap6
res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
_ = res_go_arrayMap6
for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_1.IntVal) * (2))
}), v_arrayMap6)
}
return gopurs_runtime.Array(res_go_arrayMap6)
}()
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if (gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_greaterThan(), x_1, gopurs_runtime.Int(0))
}), v_filter5).IntVal) == (1) {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}()
_ = arr_val_foldlArray4
arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
_ = arr_go_foldlArray4
res_go_foldlArray4 := gopurs_runtime.Int(0)
_ = res_go_foldlArray4
for _, v_foldlArray4 := range *arr_go_foldlArray4 {
res_go_foldlArray4 = gopurs_runtime.Apply2(pkg_Data_Semiring.Get_intAdd(), res_go_foldlArray4, v_foldlArray4)
}
return res_go_foldlArray4
}()))
}))
	})
	return cache_act
}

func Call_sumNative(arr_0_loop []int64) int64 {
var arr_0 []int64 = arr_0_loop
_ = arr_0
return func() gopurs_runtime.Value {
arr_val_foldlArray0 := arr_0
_ = arr_val_foldlArray0
res_go_foldlArray0 := gopurs_runtime.Int(0)
_ = res_go_foldlArray0
for _, v_foldlArray0 := range arr_val_foldlArray0 {
res_go_foldlArray0 = gopurs_runtime.Apply2(pkg_Data_Semiring.Get_intAdd(), res_go_foldlArray0, gopurs_runtime.Int(v_foldlArray0))
}
return res_go_foldlArray0
}().IntVal
}

func Get_arrayMap() gopurs_runtime.Value {
	return _Gopurs_ArrayMap
}

func Get_filter() gopurs_runtime.Value {
	return _Gopurs_Filter
}

func Get_foldlArray() gopurs_runtime.Value {
	return _Gopurs_FoldlArray
}
