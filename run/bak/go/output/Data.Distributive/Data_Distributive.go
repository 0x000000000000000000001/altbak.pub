package Data_Distributive

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Type_Equality "gopurs/output/Type.Equality"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_2)
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3)
})
})
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Unsafe_Coerce.Get_unsafeCoerce())
_ = __local_var_1_1
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
})
}))
	})
	return cache_distributiveIdentity
}

var cache_distribute gopurs_runtime.Value
var once_distribute sync.Once
func Get_distribute() gopurs_runtime.Value {
	once_distribute.Do(func() {
		cache_distribute = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distribute(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_distribute
}

var cache_distributiveFunction gopurs_runtime.Value
var once_distributiveFunction sync.Once
func Get_distributiveFunction() gopurs_runtime.Value {
	once_distributiveFunction.Do(func() {
		cache_distributiveFunction = gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_distributiveFunction(), "distribute"), dictFunctor_0)
_ = __local_var_2_0
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
}))
	})
	return cache_distributiveFunction
}

var cache_cotraverse gopurs_runtime.Value
var once_cotraverse sync.Once
func Get_cotraverse() gopurs_runtime.Value {
	once_cotraverse.Do(func() {
		cache_cotraverse = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cotraverse(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dictDistributive_0_box))
})
	})
	return cache_cotraverse
}

var cache_collectDefault gopurs_runtime.Value
var once_collectDefault sync.Once
func Get_collectDefault() gopurs_runtime.Value {
	once_collectDefault.Do(func() {
		cache_collectDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collectDefault(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dictDistributive_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box))
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
return Call_collect(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_collect
}

var cache_distributeDefault gopurs_runtime.Value
var once_distributeDefault sync.Once
func Get_distributeDefault() gopurs_runtime.Value {
	once_distributeDefault.Do(func() {
		cache_distributeDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distributeDefault(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dictDistributive_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box))
})
	})
	return cache_distributeDefault
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_collect__1176340970 gopurs_runtime.Value
var once_collect__1176340970 sync.Once
func Get_collect__1176340970() gopurs_runtime.Value {
	once_collect__1176340970.Do(func() {
		cache_collect__1176340970 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collect__1176340970(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_collect__1176340970
}

var cache_collect__2649935082 gopurs_runtime.Value
var once_collect__2649935082 sync.Once
func Get_collect__2649935082() gopurs_runtime.Value {
	once_collect__2649935082.Do(func() {
		cache_collect__2649935082 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collect__2649935082(gopurs_runtime.CoerceToStruct[Constructor_Distributive[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_collect__2649935082
}

var cache_functorFn__20325936 gopurs_runtime.Value
var once_functorFn__20325936 sync.Once
func Get_functorFn__20325936() gopurs_runtime.Value {
	once_functorFn__20325936.Do(func() {
		cache_functorFn__20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__20325936
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__2562444020 gopurs_runtime.Value
var once_map__2562444020 sync.Once
func Get_map__2562444020() gopurs_runtime.Value {
	once_map__2562444020.Do(func() {
		cache_map__2562444020 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2562444020(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2562444020
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__2345808404 gopurs_runtime.Value
var once_map__2345808404 sync.Once
func Get_map__2345808404() gopurs_runtime.Value {
	once_map__2345808404.Do(func() {
		cache_map__2345808404 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2345808404(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2345808404
}

var cache_map__87655540 gopurs_runtime.Value
var once_map__87655540 sync.Once
func Get_map__87655540() gopurs_runtime.Value {
	once_map__87655540.Do(func() {
		cache_map__87655540 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__87655540(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__87655540
}

var cache_map__1938733460 gopurs_runtime.Value
var once_map__1938733460 sync.Once
func Get_map__1938733460() gopurs_runtime.Value {
	once_map__1938733460.Do(func() {
		cache_map__1938733460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1938733460(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1938733460
}

var cache_functorIdentity__943655089 gopurs_runtime.Value
var once_functorIdentity__943655089 sync.Once
func Get_functorIdentity__943655089() gopurs_runtime.Value {
	once_functorIdentity__943655089.Do(func() {
		cache_functorIdentity__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorIdentity__943655089
}

var cache_functorTuple__2249620049 gopurs_runtime.Value
var once_functorTuple__2249620049 sync.Once
func Get_functorTuple__2249620049() gopurs_runtime.Value {
	once_functorTuple__2249620049.Do(func() {
		cache_functorTuple__2249620049 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple__2249620049
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_from__4089948322 gopurs_runtime.Value
var once_from__4089948322 sync.Once
func Get_from__4089948322() gopurs_runtime.Value {
	once_from__4089948322.Do(func() {
		cache_from__4089948322 = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from__4089948322(gopurs_runtime.CoerceToStruct[pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_from__4089948322
}

var cache_from__2366809570 gopurs_runtime.Value
var once_from__2366809570 sync.Once
func Get_from__2366809570() gopurs_runtime.Value {
	once_from__2366809570.Do(func() {
		cache_from__2366809570 = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from__2366809570(gopurs_runtime.CoerceToStruct[pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
})
	})
	return cache_from__2366809570
}

var cache_proof__3363032129 gopurs_runtime.Value
var once_proof__3363032129 sync.Once
func Get_proof__3363032129() gopurs_runtime.Value {
	once_proof__3363032129.Do(func() {
		cache_proof__3363032129 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_proof__3363032129(gopurs_runtime.CoerceToStruct[pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_proof__3363032129
}

type Constructor_Distributive[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[457335066] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Distributive[gopurs_runtime.Value])(ptr)
		switch key {
		case "Functor0": return c.V0
		case "collect": return c.V1
		case "distribute": return c.V2
		default: panic("Key not found in dictionary Constructor_Distributive: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_distribute(dict_0_loop *Constructor_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Distributive[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_cotraverse(dictDistributive_0_loop *Constructor_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Distributive[gopurs_runtime.Value] = dictDistributive_0_loop
_ = dictDistributive_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictDistributive_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_3_1 := gopurs_runtime.Apply(dictDistributive_0.V2, dictFunctor_2)
_ = distribute2_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Functor0_1_0.V0, f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(distribute2_3_1, x_6))
})
})
})
}

func Call_collectDefault(dictDistributive_0_loop *Constructor_Distributive[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Distributive[gopurs_runtime.Value] = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
distribute2_2_0 := gopurs_runtime.Apply(dictDistributive_0.V2, gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_1)})
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(dictFunctor_1.V0, f_3)
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
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
distribute2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_distributiveTuple(dictTypeEquals_0), "distribute"), dictFunctor_1)
_ = distribute2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute2_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
}), gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTypeEquals_0, "proof"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a_2
}), pkg_Data_Unit.Get_unit()))
_ = __local_var_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), pkg_Data_Tuple.Get_snd())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
}))
}
}

func Call_collect(dict_0_loop *Constructor_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Distributive[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_distributeDefault(dictDistributive_0_loop *Constructor_Distributive[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Distributive[gopurs_runtime.Value] = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
return gopurs_runtime.Apply2(dictDistributive_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_1)}, Get_identity())
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_collect__1176340970(dict_0_loop *Constructor_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Distributive[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_collect__2649935082(dict_0_loop *Constructor_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Distributive[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2562444020(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2345808404(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__87655540(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1938733460(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_from__4089948322(dictTypeEquals_0_loop *pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(dictTypeEquals_0.V1, gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_from__2366809570(dictTypeEquals_0_loop *pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTypeEquals_0 *pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Apply(dictTypeEquals_0.V1, gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_1
}))
}

func Call_proof__3363032129(dict_0_loop *pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Type_Equality.Constructor_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


