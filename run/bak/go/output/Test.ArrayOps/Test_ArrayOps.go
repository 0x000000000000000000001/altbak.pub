package Test_ArrayOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_range_(start_0_box, end_1_box)
})
	})
	return cache_range_
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(arr_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var arr_0 gopurs_runtime.Value = arr_0_loop
_ = arr_0
return gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), x_1, gopurs_runtime.Int(2)).IntVal) == (0))
}), arr_0)
}()
})
	})
	return cache_filterEvens
}

var cache_sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		cache_sumEvens = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
return func() gopurs_runtime.Value {
arr_val_foldlArray0 := gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), x_1, gopurs_runtime.Int(2)).IntVal) == (0))
}), gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), gopurs_runtime.Int(1), n_0))
_ = arr_val_foldlArray0
res_go_foldlArray0 := gopurs_runtime.Int(0)
_ = res_go_foldlArray0
arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
_ = arr_go_foldlArray0
for _, v_foldlArray0 := range *arr_go_foldlArray0 {
res_go_foldlArray0 = gopurs_runtime.Apply2(pkg_Data_Semiring.Get_intAdd(), res_go_foldlArray0, v_foldlArray0)
}
return res_go_foldlArray0
}()
}()
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
		cache_act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_sumEvens(), dummy_1_1))), gopurs_runtime.Value{})
})
}()
	})
	return cache_act
}

func Call_range_(start_0_loop gopurs_runtime.Value, end_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 gopurs_runtime.Value = start_0_loop
_ = start_0
var end_1 gopurs_runtime.Value = end_1_loop
_ = end_1
return gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), start_0, end_1)
}


