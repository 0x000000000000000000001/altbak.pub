package Data_Distributive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var distributiveIdentity gopurs_runtime.Value
var once_distributiveIdentity sync.Once
func Get_distributiveIdentity() gopurs_runtime.Value {
	once_distributiveIdentity.Do(func() {
		distributiveIdentity = gopurs_runtime.RecordDict3("distribute", "collect", "Functor0", gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Unsafe_Coerce.Get_unsafeCoerce())
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_2)
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}))
	})
	return distributiveIdentity
}

var distribute gopurs_runtime.Value
var once_distribute sync.Once
func Get_distribute() gopurs_runtime.Value {
	once_distribute.Do(func() {
		distribute = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "distribute")
})
	})
	return distribute
}

var distributiveFunction gopurs_runtime.Value
var once_distributiveFunction sync.Once
func Get_distributiveFunction() gopurs_runtime.Value {
	once_distributiveFunction.Do(func() {
		distributiveFunction = gopurs_runtime.RecordDict3("distribute", "collect", "Functor0", gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, e_2)
}), a_1)
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_distributiveFunction(), "distribute"), dictFunctor_0)
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}))
	})
	return distributiveFunction
}

var cotraverse gopurs_runtime.Value
var once_cotraverse sync.Once
func Get_cotraverse() gopurs_runtime.Value {
	once_cotraverse.Do(func() {
		cotraverse = gopurs_runtime.Func2(Call_cotraverse)
	})
	return cotraverse
}

var collectDefault gopurs_runtime.Value
var once_collectDefault sync.Once
func Get_collectDefault() gopurs_runtime.Value {
	once_collectDefault.Do(func() {
		collectDefault = gopurs_runtime.Func2(Call_collectDefault)
	})
	return collectDefault
}

var distributiveTuple gopurs_runtime.Value
var once_distributiveTuple sync.Once
func Get_distributiveTuple() gopurs_runtime.Value {
	once_distributiveTuple.Do(func() {
		distributiveTuple = gopurs_runtime.Func(func(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
distributiveTuple:
for {
if false { continue distributiveTuple }
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
from_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTypeEquals_0_loop, "proof"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
_ = from_1_0
return gopurs_runtime.RecordDict3("collect", "distribute", "Functor0", gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_distributiveTuple(), dictTypeEquals_0_loop), "distribute"), dictFunctor_2)
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
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}))
}
}()
})
	})
	return distributiveTuple
}

var collect gopurs_runtime.Value
var once_collect sync.Once
func Get_collect() gopurs_runtime.Value {
	once_collect.Do(func() {
		collect = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "collect")
}()
})
	})
	return collect
}

var distributeDefault gopurs_runtime.Value
var once_distributeDefault sync.Once
func Get_distributeDefault() gopurs_runtime.Value {
	once_distributeDefault.Do(func() {
		distributeDefault = gopurs_runtime.Func2(Call_distributeDefault)
	})
	return distributeDefault
}

func Call_cotraverse(dictDistributive_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0_loop, "distribute"), dictFunctor_1_loop)
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0_loop, "Functor0"), gopurs_runtime.Value{}), "map"), f_3)
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
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0_loop, "distribute"), dictFunctor_1_loop)
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1_loop, "map"), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
}

func Call_distributeDefault(dictDistributive_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictDistributive_0_loop, "collect"), dictFunctor_1_loop, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}


