package Test_STArray

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var cache_sumArray gopurs_runtime.Value
var once_sumArray sync.Once
func Get_sumArray() gopurs_runtime.Value {
	once_sumArray.Do(func() {
		cache_sumArray = gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(arr_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushAllImpl(), func() gopurs_runtime.Value {
					arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}(), arr_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(pkg_Data_Array_ST.Get_popImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, arr_0)
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Int(0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}))
}))).IntVal)
	})
	return cache_sumArray
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("STArray Operations:"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), Get_sumArray()))
	})
	return cache_act
}




