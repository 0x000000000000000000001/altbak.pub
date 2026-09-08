package main

import (
 gopurs_runtime "gopurs/output/gopurs_runtime"
 ps "gopurs/output/purescript"
)

// Source copied verbatim except name and package qualification.
func copiedOriginalSumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return func() gopurs_runtime.Value {
arr_val_foldlArray0 := func() gopurs_runtime.Value {
					arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
arr_val_filterImpl1 := gopurs_runtime.UncurriedApp2(ps.Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(n_0))
_ = arr_val_filterImpl1
_ = arr_val_filterImpl1
arr_go_filterImpl1 := (*[]gopurs_runtime.Value)(arr_val_filterImpl1.UnsafePtr)
_ = arr_go_filterImpl1
res_go_filterImpl1 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl1
for _, v_filterImpl1 := range *arr_go_filterImpl1 {
if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((x_1.IntVal) % (int64(2))) == (int64(0)))
}), v_filterImpl1).BoolVal() {
res_go_filterImpl1 = append(res_go_filterImpl1, v_filterImpl1)
} else {

}
}
return res_go_filterImpl1
}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
_ = arr_val_foldlArray0
res_go_foldlArray0 := gopurs_runtime.Int(int64(0))
_ = res_go_foldlArray0
arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
_ = arr_go_foldlArray0
for _, v_foldlArray0 := range *arr_go_foldlArray0 {
res_go_foldlArray0 = gopurs_runtime.Apply2(ps.Get_Data_Semiring_intAdd(), res_go_foldlArray0, v_foldlArray0)
}
return res_go_foldlArray0
}().IntVal
}

// Audit-only counterfactual: remove []Value -> []int64 -> []Value.
// Keep the original range, filter, append policy and Apply/Apply2 dispatch.
func withoutArrayRoundtrip(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return func() gopurs_runtime.Value {
arr_val_foldlArray0 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
arr_val_filterImpl1 := gopurs_runtime.UncurriedApp2(ps.Get_Data_Array_rangeImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(n_0))
_ = arr_val_filterImpl1
_ = arr_val_filterImpl1
arr_go_filterImpl1 := (*[]gopurs_runtime.Value)(arr_val_filterImpl1.UnsafePtr)
_ = arr_go_filterImpl1
res_go_filterImpl1 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl1
for _, v_filterImpl1 := range *arr_go_filterImpl1 {
if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((x_1.IntVal) % (int64(2))) == (int64(0)))
}), v_filterImpl1).BoolVal() {
res_go_filterImpl1 = append(res_go_filterImpl1, v_filterImpl1)
} else {

}
}
return res_go_filterImpl1
}())
_ = arr_val_foldlArray0
res_go_foldlArray0 := gopurs_runtime.Int(int64(0))
_ = res_go_foldlArray0
arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
_ = arr_go_foldlArray0
for _, v_foldlArray0 := range *arr_go_foldlArray0 {
res_go_foldlArray0 = gopurs_runtime.Apply2(ps.Get_Data_Semiring_intAdd(), res_go_foldlArray0, v_foldlArray0)
}
return res_go_foldlArray0
}().IntVal
}
