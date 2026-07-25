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
		cache_sumArray = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
arr_0_0 := gopurs_runtime.Apply(pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Value{})
_ = arr_0_0
_dollar__unused_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_pushAll(), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(1), gopurs_runtime.Int(2), gopurs_runtime.Int(3), gopurs_runtime.Int(4), gopurs_runtime.Int(5), gopurs_runtime.Int(6), gopurs_runtime.Int(7), gopurs_runtime.Int(8), gopurs_runtime.Int(9), gopurs_runtime.Int(10)}), arr_0_0), gopurs_runtime.Value{})
_ = _dollar__unused_1_1
x_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_ST.Get_pop(), arr_0_0), gopurs_runtime.Value{})
_ = x_2_2
var __t3 gopurs_runtime.Value
{
if (x_2_2.Type == 9 && x_2_2.IntVal == 930809136) {
__t3 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Maybe.Data_Data_Maybe_Just)(x_2_2.UnsafePtr).V0
})
goto end_branch_3
} else {

}
}
{
if (x_2_2.Type == 9 && x_2_2.IntVal == 3589588149) {
__t3 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Apply(__t3, gopurs_runtime.Value{})
}))
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
		cache_act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), Get_sumArray()))
	})
	return cache_act
}




