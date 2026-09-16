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
		cache_Data_Distributive_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Distributive_unwrap
}

var cache_Data_Distributive_unwrap1 gopurs_runtime.Value
var once_Data_Distributive_unwrap1 sync.Once
func Get_Data_Distributive_unwrap1() gopurs_runtime.Value {
	once_Data_Distributive_unwrap1.Do(func() {
		cache_Data_Distributive_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Distributive_unwrap1
}

var cache_Data_Distributive_identity gopurs_runtime.Value
var once_Data_Distributive_identity sync.Once
func Get_Data_Distributive_identity() gopurs_runtime.Value {
	once_Data_Distributive_identity.Do(func() {
		cache_Data_Distributive_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Distributive_identity
}

var cache_Data_Distributive_Distributive_dollar_Dict gopurs_runtime.Value
var once_Data_Distributive_Distributive_dollar_Dict sync.Once
func Get_Data_Distributive_Distributive_dollar_Dict() gopurs_runtime.Value {
	once_Data_Distributive_Distributive_dollar_Dict.Do(func() {
		cache_Data_Distributive_Distributive_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer(Call_Data_Distributive_Distributive_dollar_Dict(func() struct{
	Functor0 gopurs_runtime.Value
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Functor0 gopurs_runtime.Value
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}{}
					clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
					clone.collect = gopurs_runtime.RecordGet(orig, "collect")
					clone.distribute = gopurs_runtime.RecordGet(orig, "distribute")
					return clone
				}()))}
})
	})
	return cache_Data_Distributive_Distributive_dollar_Dict
}

var cache_Data_Distributive_distributiveIdentity gopurs_runtime.Value
var once_Data_Distributive_distributiveIdentity sync.Once
func Get_Data_Distributive_distributiveIdentity() gopurs_runtime.Value {
	once_Data_Distributive_distributiveIdentity.Do(func() {
		cache_Data_Distributive_distributiveIdentity = gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer((&Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), f_1)))
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
})}))}
	})
	return cache_Data_Distributive_distributiveIdentity
}

var cache_Data_Distributive_distribute gopurs_runtime.Value
var once_Data_Distributive_distribute sync.Once
func Get_Data_Distributive_distribute() gopurs_runtime.Value {
	once_Data_Distributive_distribute.Do(func() {
		cache_Data_Distributive_distribute = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_distribute(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Distributive_distribute
}

var cache_Data_Distributive_distributiveFunction gopurs_runtime.Value
var once_Data_Distributive_distributiveFunction sync.Once
func Get_Data_Distributive_distributiveFunction() gopurs_runtime.Value {
	once_Data_Distributive_distributiveFunction.Do(func() {
		cache_Data_Distributive_distributiveFunction = gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer((&Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()))}
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Data_Distributive_distribute(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](Get_Data_Distributive_distributiveFunction())), dictFunctor_0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1))
}), gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, e_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, e_2)
}), a_1)
})}))}
	})
	return cache_Data_Distributive_distributiveFunction
}

var cache_Data_Distributive_cotraverse gopurs_runtime.Value
var once_Data_Distributive_cotraverse sync.Once
func Get_Data_Distributive_cotraverse() gopurs_runtime.Value {
	once_Data_Distributive_cotraverse.Do(func() {
		cache_Data_Distributive_cotraverse = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_cotraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](dictDistributive_0_box))
})
	})
	return cache_Data_Distributive_cotraverse
}

var cache_Data_Distributive_collectDefault gopurs_runtime.Value
var once_Data_Distributive_collectDefault sync.Once
func Get_Data_Distributive_collectDefault() gopurs_runtime.Value {
	once_Data_Distributive_collectDefault.Do(func() {
		cache_Data_Distributive_collectDefault = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_collectDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](dictDistributive_0_box))
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
return Call_Data_Distributive_collect(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Distributive_collect
}

var cache_Data_Distributive_distributeDefault gopurs_runtime.Value
var once_Data_Distributive_distributeDefault sync.Once
func Get_Data_Distributive_distributeDefault() gopurs_runtime.Value {
	once_Data_Distributive_distributeDefault.Do(func() {
		cache_Data_Distributive_distributeDefault = gopurs_runtime.Func2(func(dictDistributive_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Distributive_distributeDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](dictDistributive_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box))
})
	})
	return cache_Data_Distributive_distributeDefault
}

type Constructor_Data_Distributive_Distributive[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[457335066] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Distributive_Distributive[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0": return gopurs_runtime.Box(c.V0)
		case "collect": return gopurs_runtime.Box(c.V1)
		case "distribute": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Distributive_Distributive: " + key)
		}
	}
}


func Call_Data_Distributive_Distributive_dollar_Dict(x_0_loop struct{
	Functor0 gopurs_runtime.Value
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}) *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] {
var x_0 struct{
	Functor0 gopurs_runtime.Value
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", orig.Functor0, orig.collect, orig.distribute)
				}())
}

func Call_Data_Distributive_distribute(dict_0_loop *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_Data_Distributive_cotraverse(dictDistributive_0_loop *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope39)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictDistributive_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): distribute1_2_1 shape=App(Var) bindingType=Any
distribute1_2_1 := Call_Data_Distributive_distribute(dictDistributive_0)
_ = distribute1_2_1
return gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): distribute2_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope40) [(TypeApp (TypeVar f$scope39) [(TypeVar a$scope37)])])] (TypeApp (TypeVar f$scope39) [(TypeApp (TypeVar g$scope40) [(TypeVar a$scope37)])]))
distribute2_4_2 := gopurs_runtime.Apply(distribute1_2_1, dictFunctor_3)
_ = distribute2_4_2
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_1_0.V0, f_5), distribute2_4_2)
})
})
}

func Call_Data_Distributive_collectDefault(dictDistributive_0_loop *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): distribute1_1_0 shape=App(Var) bindingType=Any
distribute1_1_0 := Call_Data_Distributive_distribute(dictDistributive_0)
_ = distribute1_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): distribute2_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope49) [(TypeApp (TypeVar f$scope48) [(TypeVar b$scope47)])])] (TypeApp (TypeVar f$scope48) [(TypeApp (TypeVar g$scope49) [(TypeVar b$scope47)])]))
distribute2_3_1 := gopurs_runtime.Apply(distribute1_1_0, dictFunctor_2)
_ = distribute2_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), distribute2_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_4))
})
})
}

func Call_Data_Distributive_distributiveTuple(dictTypeEquals_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveTuple:
for {
if false { continue distributiveTuple }
var dictTypeEquals_0 gopurs_runtime.Value = dictTypeEquals_0_loop
_ = dictTypeEquals_0
return gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer(Rebox_Data_Distributive_1559819651_3927024206((&Constructor_Data_Distributive_Distributive[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Distributive_2363162019_2812149806(Rebox_Data_Distributive_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Tuple_functorTuple()))))}
}), gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Distributive_collectDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]](Call_Data_Distributive_distributiveTuple(dictTypeEquals_0))), dictFunctor_1)
}), gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Apply2(Call_Type_Equality_proof(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0)), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return a_2
}), Get_Data_Unit_unit())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_1, "map"), Get_Data_Tuple_snd()))
})})))}
}
}

func Call_Data_Distributive_collect(dict_0_loop *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Data_Distributive_distributeDefault(dictDistributive_0_loop *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDistributive_0 *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] = dictDistributive_0_loop
_ = dictDistributive_0
var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
return gopurs_runtime.Apply2(dictDistributive_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_1)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Rebox_Data_Distributive_1559819651_3927024206(in *Constructor_Data_Distributive_Distributive[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Distributive_Distributive[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Distributive_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Distributive_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}


