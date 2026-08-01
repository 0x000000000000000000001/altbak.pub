package Data_Distributive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_distributiveIdentity gopurs_runtime.Value
var once_distributiveIdentity sync.Once
func Get_distributiveIdentity() gopurs_runtime.Value {
	once_distributiveIdentity.Do(func() {
		cache_distributiveIdentity = gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_2)
}))
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Unsafe_Coerce.Get_unsafeCoerce())
}))
	})
	return cache_distributiveIdentity
}

var cache_distribute gopurs_runtime.Value
var once_distribute sync.Once
func Get_distribute() gopurs_runtime.Value {
	once_distribute.Do(func() {
		cache_distribute = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distribute(dict_0_box)
})
	})
	return cache_distribute
}

var cache_distribute__gopurs_runtime_Value_4119050416 gopurs_runtime.Value
var once_distribute__gopurs_runtime_Value_4119050416 sync.Once
func Get_distribute__gopurs_runtime_Value_4119050416() gopurs_runtime.Value {
	once_distribute__gopurs_runtime_Value_4119050416.Do(func() {
		cache_distribute__gopurs_runtime_Value_4119050416 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distribute__gopurs_runtime_Value_4119050416(dict_0_box)
})
	})
	return cache_distribute__gopurs_runtime_Value_4119050416
}

var cache_distributiveFunction gopurs_runtime.Value
var once_distributiveFunction sync.Once
func Get_distributiveFunction() gopurs_runtime.Value {
	once_distributiveFunction.Do(func() {
		cache_distributiveFunction = gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_distributiveFunction(), "distribute"), dictFunctor_0)
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}), gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, e_2)
}), a_1)
}))
	})
	return cache_distributiveFunction
}

var cache_cotraverse gopurs_runtime.Value
var once_cotraverse sync.Once
func Get_cotraverse() gopurs_runtime.Value {
	once_cotraverse.Do(func() {
		cache_cotraverse = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cotraverse(dictDistributive_0_box, dictFunctor_1_box)
})
	})
	return cache_cotraverse
}

var cache_collectDefault gopurs_runtime.Value
var once_collectDefault sync.Once
func Get_collectDefault() gopurs_runtime.Value {
	once_collectDefault.Do(func() {
		cache_collectDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collectDefault(dictDistributive_0_box, dictFunctor_1_box)
})
	})
	return cache_collectDefault
}

var cache_distributiveTuple gopurs_runtime.Value
var once_distributiveTuple sync.Once
func Get_distributiveTuple() gopurs_runtime.Value {
	once_distributiveTuple.Do(func() {
		cache_distributiveTuple = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distributiveTuple(dictTypeEquals_0_box)
})
	})
	return cache_distributiveTuple
}

var cache_collect gopurs_runtime.Value
var once_collect sync.Once
func Get_collect() gopurs_runtime.Value {
	once_collect.Do(func() {
		cache_collect = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collect(dict_0_box)
})
	})
	return cache_collect
}

var cache_collect__gopurs_runtime_Value_1217025457 gopurs_runtime.Value
var once_collect__gopurs_runtime_Value_1217025457 sync.Once
func Get_collect__gopurs_runtime_Value_1217025457() gopurs_runtime.Value {
	once_collect__gopurs_runtime_Value_1217025457.Do(func() {
		cache_collect__gopurs_runtime_Value_1217025457 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collect__gopurs_runtime_Value_1217025457(dict_0_box)
})
	})
	return cache_collect__gopurs_runtime_Value_1217025457
}

var cache_distributeDefault gopurs_runtime.Value
var once_distributeDefault sync.Once
func Get_distributeDefault() gopurs_runtime.Value {
	once_distributeDefault.Do(func() {
		cache_distributeDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distributeDefault(dictDistributive_0_box, dictFunctor_1_box)
})
	})
	return cache_distributeDefault
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_distribute(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "distribute")
}

func Call_distribute__gopurs_runtime_Value_4119050416(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "distribute")
}

func Call_cotraverse(dictDistributive_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "distribute"), dictFunctor_1)
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{}), "map"), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(distribute2_2_0, x_5))
})
})
}

func Call_collectDefault(dictDistributive_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "distribute"), dictFunctor_1)
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
}

func Call_distributiveTuple(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveTuple:
for {
if false { continue distributiveTuple }
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
from_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
_ = from_1_0
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_distributiveTuple(dictTypeEquals_0), "distribute"), dictFunctor_2)
_ = distribute2_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_3_1, gopurs_runtime.Apply(__local_var_5_2, x_6))
})
})
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply(from_1_0, pkg_Data_Unit.Get_unit()))
_ = __local_var_3_3
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), pkg_Data_Tuple.Get_snd())
_ = __local_var_4_4
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Apply(__local_var_4_4, x_5))
})
}))
}
}

func Call_collect(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "collect")
}

func Call_collect__gopurs_runtime_Value_1217025457(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "collect")
}

func Call_distributeDefault(dictDistributive_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), dictFunctor_1, Get_identity())
}


