package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_AffOperations_liftEffect gopurs_runtime.Value
var once_Test_AffOperations_liftEffect sync.Once

func Get_Test_AffOperations_liftEffect() gopurs_runtime.Value {
	once_Test_AffOperations_liftEffect.Do(func() {
		cache_Test_AffOperations_liftEffect = Get_Effect_Aff__liftEffect()
	})
	return cache_Test_AffOperations_liftEffect
}

var cache_Test_AffOperations_describe gopurs_runtime.Value
var once_Test_AffOperations_describe sync.Once

func Get_Test_AffOperations_describe() gopurs_runtime.Value {
	once_Test_AffOperations_describe.Do(func() {
		cache_Test_AffOperations_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Aff Operations (Asynchronous Delays)"))
	})
	return cache_Test_AffOperations_describe
}

var cache_Test_AffOperations_act gopurs_runtime.Value
var once_Test_AffOperations_act sync.Once

func Get_Test_AffOperations_act() gopurs_runtime.Value {
	once_Test_AffOperations_act.Do(func() {
		cache_Test_AffOperations_act = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.UncurriedApp2(Get_Effect_Aff__delay(), Get_Data_Either_Right(), gopurs_runtime.Float(10.0)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("10")))
		}))
	})
	return cache_Test_AffOperations_act
}
