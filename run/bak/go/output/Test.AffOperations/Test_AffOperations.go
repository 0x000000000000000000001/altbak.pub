package Test_AffOperations

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Aff Operations (Asynchronous Delays):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.UncurriedApp2(pkg_Effect_Aff.Get__makeFiber(), pkg_Effect_Aff.Get_ffiUtil(), gopurs_runtime.Apply2(pkg_Effect_Aff.Get__bind(), gopurs_runtime.UncurriedApp2(pkg_Effect_Aff.Get__delay(), pkg_Data_Either.Get_Right(), gopurs_runtime.Float(10.0)), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Aff.Get__liftEffect(), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("10")))
})))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fiber_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = fiber_1_1
_dollar__unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(fiber_1_1, "run"), gopurs_runtime.Value{})
_ = _dollar__unused_2_2
return pkg_Data_Unit.Get_unit()
})
}()
	})
	return cache_act
}




