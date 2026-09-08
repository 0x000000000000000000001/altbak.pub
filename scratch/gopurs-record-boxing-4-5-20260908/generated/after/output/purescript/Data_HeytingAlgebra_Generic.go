package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict sync.Once
func Get_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict.Do(func() {
		cache_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(Call_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict(func() struct{
	genericConj_prime_ gopurs_runtime.Value
	genericDisj_prime_ gopurs_runtime.Value
	genericFF_prime_ gopurs_runtime.Value
	genericImplies_prime_ gopurs_runtime.Value
	genericNot_prime_ gopurs_runtime.Value
	genericTT_prime_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericConj_prime_ gopurs_runtime.Value
	genericDisj_prime_ gopurs_runtime.Value
	genericFF_prime_ gopurs_runtime.Value
	genericImplies_prime_ gopurs_runtime.Value
	genericNot_prime_ gopurs_runtime.Value
	genericTT_prime_ gopurs_runtime.Value
}{}
					clone.genericConj_prime_ = gopurs_runtime.RecordGet(orig, "genericConj'")
					clone.genericDisj_prime_ = gopurs_runtime.RecordGet(orig, "genericDisj'")
					clone.genericFF_prime_ = gopurs_runtime.RecordGet(orig, "genericFF'")
					clone.genericImplies_prime_ = gopurs_runtime.RecordGet(orig, "genericImplies'")
					clone.genericNot_prime_ = gopurs_runtime.RecordGet(orig, "genericNot'")
					clone.genericTT_prime_ = gopurs_runtime.RecordGet(orig, "genericTT'")
					return clone
				}()))}
})
	})
	return cache_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict
}

var cache_Data_HeytingAlgebra_Generic_genericTT_prime_ gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericTT_prime_ sync.Once
func Get_Data_HeytingAlgebra_Generic_genericTT_prime_() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericTT_prime_.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericTT_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericTT_prime_(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericTT_prime_
}

var cache_Data_HeytingAlgebra_Generic_genericTT gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericTT sync.Once
func Get_Data_HeytingAlgebra_Generic_genericTT() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericTT.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericTT = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericTT(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericTT
}

var cache_Data_HeytingAlgebra_Generic_genericNot_prime_ gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericNot_prime_ sync.Once
func Get_Data_HeytingAlgebra_Generic_genericNot_prime_() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericNot_prime_.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericNot_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericNot_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericNot_prime_
}

var cache_Data_HeytingAlgebra_Generic_genericNot gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericNot sync.Once
func Get_Data_HeytingAlgebra_Generic_genericNot() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericNot.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericNot = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericNot(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericNot
}

var cache_Data_HeytingAlgebra_Generic_genericImplies_prime_ gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericImplies_prime_ sync.Once
func Get_Data_HeytingAlgebra_Generic_genericImplies_prime_() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericImplies_prime_.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericImplies_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericImplies_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericImplies_prime_
}

var cache_Data_HeytingAlgebra_Generic_genericImplies gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericImplies sync.Once
func Get_Data_HeytingAlgebra_Generic_genericImplies() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericImplies.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericImplies = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericImplies(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericImplies
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(Rebox_Data_HeytingAlgebra_Generic_1092464188_1111920849((&Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), 1454898258, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), 1454898258})))}
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument
}

var cache_Data_HeytingAlgebra_Generic_genericFF_prime_ gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericFF_prime_ sync.Once
func Get_Data_HeytingAlgebra_Generic_genericFF_prime_() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericFF_prime_.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericFF_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericFF_prime_(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericFF_prime_
}

var cache_Data_HeytingAlgebra_Generic_genericFF gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericFF sync.Once
func Get_Data_HeytingAlgebra_Generic_genericFF() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericFF.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericFF = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericFF(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericFF
}

var cache_Data_HeytingAlgebra_Generic_genericDisj_prime_ gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericDisj_prime_ sync.Once
func Get_Data_HeytingAlgebra_Generic_genericDisj_prime_() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericDisj_prime_.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericDisj_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericDisj_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericDisj_prime_
}

var cache_Data_HeytingAlgebra_Generic_genericDisj gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericDisj sync.Once
func Get_Data_HeytingAlgebra_Generic_genericDisj() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericDisj.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericDisj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericDisj(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericDisj
}

var cache_Data_HeytingAlgebra_Generic_genericConj_prime_ gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericConj_prime_ sync.Once
func Get_Data_HeytingAlgebra_Generic_genericConj_prime_() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericConj_prime_.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericConj_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericConj_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericConj_prime_
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct = gopurs_runtime.Func2(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_box, dictGenericHeytingAlgebra1_1_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct
}

var cache_Data_HeytingAlgebra_Generic_genericConj gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericConj sync.Once
func Get_Data_HeytingAlgebra_Generic_genericConj() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericConj.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericConj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericConj(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericConj
}

type Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 T_a
}


func init() {
	gopurs_runtime.StructGetters[2831861733] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[any])(ptr)
		_ = c
		switch key {
		case "genericConj'": return gopurs_runtime.Box(c.V0)
		case "genericDisj'": return gopurs_runtime.Box(c.V1)
		case "genericFF'": return gopurs_runtime.Box(c.V2)
		case "genericImplies'": return gopurs_runtime.Box(c.V3)
		case "genericNot'": return gopurs_runtime.Box(c.V4)
		case "genericTT'": return gopurs_runtime.Box(c.V5)
		default: panic("Key not found in dictionary Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra: " + key)
		}
	}
}


func Call_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollar_Dict(x_0_loop struct{
	genericConj_prime_ gopurs_runtime.Value
	genericDisj_prime_ gopurs_runtime.Value
	genericFF_prime_ gopurs_runtime.Value
	genericImplies_prime_ gopurs_runtime.Value
	genericNot_prime_ gopurs_runtime.Value
	genericTT_prime_ gopurs_runtime.Value
}) *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] {
var x_0 struct{
	genericConj_prime_ gopurs_runtime.Value
	genericDisj_prime_ gopurs_runtime.Value
	genericFF_prime_ gopurs_runtime.Value
	genericImplies_prime_ gopurs_runtime.Value
	genericNot_prime_ gopurs_runtime.Value
	genericTT_prime_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{orig.genericConj_prime_, orig.genericDisj_prime_, orig.genericFF_prime_, orig.genericImplies_prime_, orig.genericNot_prime_, orig.genericTT_prime_})
				}())
}

func Call_Data_HeytingAlgebra_Generic_genericTT_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
}

func Call_Data_HeytingAlgebra_Generic_genericTT(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericTT'"))
}

func Call_Data_HeytingAlgebra_Generic_genericNot_prime_(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_Generic_genericNot(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericHeytingAlgebra_1.V4), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2)))
}

func Call_Data_HeytingAlgebra_Generic_genericImplies_prime_(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_Generic_genericImplies(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictGenericHeytingAlgebra_1.V3), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer((&Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), v_1, v1_2)
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), v_1)
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")}))}
}

func Call_Data_HeytingAlgebra_Generic_genericFF_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericFF'")
}

func Call_Data_HeytingAlgebra_Generic_genericFF(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericFF'"))
}

func Call_Data_HeytingAlgebra_Generic_genericDisj_prime_(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_Generic_genericDisj(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictGenericHeytingAlgebra_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_HeytingAlgebra_Generic_genericConj_prime_(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 gopurs_runtime.Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
return gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer((&Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), v_1, v1_2)
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), v_1)
}), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'")}))}
}

func Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 gopurs_runtime.Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
var dictGenericHeytingAlgebra1_1 gopurs_runtime.Value = dictGenericHeytingAlgebra1_1_loop
_ = dictGenericHeytingAlgebra1_1
return gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(Rebox_Data_HeytingAlgebra_Generic_3342988926_1111920849((&Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericConj'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericDisj'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
}), (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericFF'")}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericImplies'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericNot'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)}))}
}), (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericTT'")})})))}
}

func Call_Data_HeytingAlgebra_Generic_genericConj(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictGenericHeytingAlgebra_1.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Rebox_Data_HeytingAlgebra_Generic_1092464188_1111920849(in *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[uint32]) *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = in.V3
		out.V4 = in.V4
		out.V5 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V5), UnsafePtr: nil}
	return out
}

func Rebox_Data_HeytingAlgebra_Generic_3342988926_1111920849(in *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V2)}
		out.V3 = in.V3
		out.V4 = in.V4
		out.V5 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V5)}
	return out
}


