package Test_AffOperations

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Aff Operations (Asynchronous Delays)"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.UncurriedApp2(pkg_Effect_Aff.Get__delay(), pkg_Data_Either.Get_Right(), gopurs_runtime.Float(10.0)), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("10")))
}))
	})
	return cache_act
}




