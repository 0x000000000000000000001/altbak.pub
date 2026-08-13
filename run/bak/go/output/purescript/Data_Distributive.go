package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Distributive_unwrap gopurs_runtime.Value
var once_Data_Distributive_unwrap sync.Once
func Get_Data_Distributive_unwrap() gopurs_runtime.Value {
	once_Data_Distributive_unwrap.Do(func() {
		cache_Data_Distributive_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Distributive_unwrap
}

var cache_Data_Distributive_unwrap1 gopurs_runtime.Value
var once_Data_Distributive_unwrap1 sync.Once
func Get_Data_Distributive_unwrap1() gopurs_runtime.Value {
	once_Data_Distributive_unwrap1.Do(func() {
		cache_Data_Distributive_unwrap1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Distributive_unwrap1
}

var cache_Data_Distributive_identity gopurs_runtime.Value
var once_Data_Distributive_identity sync.Once
func Get_Data_Distributive_identity() gopurs_runtime.Value {
	once_Data_Distributive_identity.Do(func() {
		cache_Data_Distributive_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_identity(x_0_box)
})
	})
	return cache_Data_Distributive_identity
}

var cache_Data_Distributive_Distributive_dollarDict gopurs_runtime.Value
var once_Data_Distributive_Distributive_dollarDict sync.Once
func Get_Data_Distributive_Distributive_dollarDict() gopurs_runtime.Value {
	once_Data_Distributive_Distributive_dollarDict.Do(func() {
		cache_Data_Distributive_Distributive_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_Distributive_dollarDict(x_0_box)
})
	})
	return cache_Data_Distributive_Distributive_dollarDict
}

var cache_Data_Distributive_distributiveIdentity gopurs_runtime.Value
var once_Data_Distributive_distributiveIdentity sync.Once
func Get_Data_Distributive_distributiveIdentity() gopurs_runtime.Value {
	once_Data_Distributive_distributiveIdentity.Do(func() {
		cache_Data_Distributive_distributiveIdentity = gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer(&Constructor_Data_Distributive_Distributive{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_2)
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3)
})
})
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), Get_Unsafe_Coerce_unsafeCoerce())
_ = __local_var_1_1
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
})
})})}
	})
	return cache_Data_Distributive_distributiveIdentity
}

var cache_Data_Distributive_distribute gopurs_runtime.Value
var once_Data_Distributive_distribute sync.Once
func Get_Data_Distributive_distribute() gopurs_runtime.Value {
	once_Data_Distributive_distribute.Do(func() {
		cache_Data_Distributive_distribute = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_distribute(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dict_0_box))
})
	})
	return cache_Data_Distributive_distribute
}

var cache_Data_Distributive_distributiveFunction gopurs_runtime.Value
var once_Data_Distributive_distributiveFunction sync.Once
func Get_Data_Distributive_distributiveFunction() gopurs_runtime.Value {
	once_Data_Distributive_distributiveFunction.Do(func() {
		cache_Data_Distributive_distributiveFunction = gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer(&Constructor_Data_Distributive_Distributive{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_functorFn()))}
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](Get_Data_Distributive_distributiveFunction()).V2), dictFunctor_0)
_ = __local_var_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, e_2)
}), a_1)
})
})
})})}
	})
	return cache_Data_Distributive_distributiveFunction
}

var cache_Data_Distributive_cotraverse gopurs_runtime.Value
var once_Data_Distributive_cotraverse sync.Once
func Get_Data_Distributive_cotraverse() gopurs_runtime.Value {
	once_Data_Distributive_cotraverse.Do(func() {
		cache_Data_Distributive_cotraverse = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_cotraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dictDistributive_0_box))
})
	})
	return cache_Data_Distributive_cotraverse
}

var cache_Data_Distributive_collectDefault gopurs_runtime.Value
var once_Data_Distributive_collectDefault sync.Once
func Get_Data_Distributive_collectDefault() gopurs_runtime.Value {
	once_Data_Distributive_collectDefault.Do(func() {
		cache_Data_Distributive_collectDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_collectDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dictDistributive_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_1_box))
})
	})
	return cache_Data_Distributive_collectDefault
}

var cache_Data_Distributive_distributiveTuple gopurs_runtime.Value
var once_Data_Distributive_distributiveTuple sync.Once
func Get_Data_Distributive_distributiveTuple() gopurs_runtime.Value {
	once_Data_Distributive_distributiveTuple.Do(func() {
		cache_Data_Distributive_distributiveTuple = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_distributiveTuple(dictTypeEquals_0_box)
})
	})
	return cache_Data_Distributive_distributiveTuple
}

var cache_Data_Distributive_collect gopurs_runtime.Value
var once_Data_Distributive_collect sync.Once
func Get_Data_Distributive_collect() gopurs_runtime.Value {
	once_Data_Distributive_collect.Do(func() {
		cache_Data_Distributive_collect = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_collect(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dict_0_box))
})
	})
	return cache_Data_Distributive_collect
}

var cache_Data_Distributive_distributeDefault gopurs_runtime.Value
var once_Data_Distributive_distributeDefault sync.Once
func Get_Data_Distributive_distributeDefault() gopurs_runtime.Value {
	once_Data_Distributive_distributeDefault.Do(func() {
		cache_Data_Distributive_distributeDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_distributeDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dictDistributive_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_1_box))
})
	})
	return cache_Data_Distributive_distributeDefault
}

var cache_Data_Distributive_collect__1176340970 gopurs_runtime.Value
var once_Data_Distributive_collect__1176340970 sync.Once
func Get_Data_Distributive_collect__1176340970() gopurs_runtime.Value {
	once_Data_Distributive_collect__1176340970.Do(func() {
		cache_Data_Distributive_collect__1176340970 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_collect__1176340970(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dict_0_box))
})
	})
	return cache_Data_Distributive_collect__1176340970
}

var cache_Data_Distributive_collect__2649935082 gopurs_runtime.Value
var once_Data_Distributive_collect__2649935082 sync.Once
func Get_Data_Distributive_collect__2649935082() gopurs_runtime.Value {
	once_Data_Distributive_collect__2649935082.Do(func() {
		cache_Data_Distributive_collect__2649935082 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_collect__2649935082(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dict_0_box))
})
	})
	return cache_Data_Distributive_collect__2649935082
}

var cache_Data_Distributive_collect__4289358698 gopurs_runtime.Value
var once_Data_Distributive_collect__4289358698 sync.Once
func Get_Data_Distributive_collect__4289358698() gopurs_runtime.Value {
	once_Data_Distributive_collect__4289358698.Do(func() {
		cache_Data_Distributive_collect__4289358698 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_collect__4289358698(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dict_0_box))
})
	})
	return cache_Data_Distributive_collect__4289358698
}

var cache_Data_Distributive_distribute__2045770422 gopurs_runtime.Value
var once_Data_Distributive_distribute__2045770422 sync.Once
func Get_Data_Distributive_distribute__2045770422() gopurs_runtime.Value {
	once_Data_Distributive_distribute__2045770422.Do(func() {
		cache_Data_Distributive_distribute__2045770422 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_distribute__2045770422(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](dict_0_box))
})
	})
	return cache_Data_Distributive_distribute__2045770422
}

type Constructor_Data_Distributive_Distributive struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[457335066] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Distributive_Distributive)(ptr)
		_ = c
		switch key {
		case "Functor0": return gopurs_runtime.Box(c.V0)
		case "collect": return gopurs_runtime.Box(c.V1)
		case "distribute": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Distributive_Distributive: " + key)
		}
	}
}


func Call_Data_Distributive_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Distributive_Distributive_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Distributive_distribute(dict_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Distributive_cotraverse(dictDistributive_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Data_Distributive_Distributive = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictDistributive_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): distribute2_3_1 -> gopurs_runtime.Value
distribute2_3_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictDistributive_0.V2), dictFunctor_2)
_ = distribute2_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(distribute2_3_1, x_6))
})
})
})
}

func Call_Data_Distributive_collectDefault(dictDistributive_0_loop *Constructor_Data_Distributive_Distributive, dictFunctor_1_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Data_Distributive_Distributive = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 *Constructor_Data_Functor_Functor = dictFunctor_1_loop
_ = dictFunctor_1
// TAST (Let): distribute2_2_0 -> gopurs_runtime.Value
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictDistributive_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_1)})
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_1.V0), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
}

func Call_Data_Distributive_distributiveTuple(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveTuple:
for {
if false { continue distributiveTuple }
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer(&Constructor_Data_Distributive_Distributive{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Tuple_functorTuple()))}
}), gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): distribute2_2_0 -> gopurs_runtime.Value
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Distributive_distributiveTuple(dictTypeEquals_0), "distribute"), dictFunctor_1)
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
}), gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a_2
}), Get_Data_Unit_unit()))
_ = __local_var_2_2
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), Get_Data_Tuple_snd())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
})})}
}
}

func Call_Data_Distributive_collect(dict_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Distributive_distributeDefault(dictDistributive_0_loop *Constructor_Data_Distributive_Distributive, dictFunctor_1_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Data_Distributive_Distributive = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 *Constructor_Data_Functor_Functor = dictFunctor_1_loop
_ = dictFunctor_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictDistributive_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Distributive_collect__1176340970(dict_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Distributive_collect__2649935082(dict_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Distributive_collect__4289358698(dict_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Distributive_distribute__2045770422(dict_0_loop *Constructor_Data_Distributive_Distributive) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}


