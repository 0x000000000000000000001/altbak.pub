package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ring_Generic_GenericRing_dollar_Dict gopurs_runtime.Value
var once_Data_Ring_Generic_GenericRing_dollar_Dict sync.Once
func Get_Data_Ring_Generic_GenericRing_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ring_Generic_GenericRing_dollar_Dict.Do(func() {
		cache_Data_Ring_Generic_GenericRing_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Call_Data_Ring_Generic_GenericRing_dollar_Dict(func() struct{
	genericSub_prime_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericSub_prime_ gopurs_runtime.Value
}{}
					clone.genericSub_prime_ = gopurs_runtime.RecordGet(orig, "genericSub'")
					return clone
				}()))}
})
	})
	return cache_Data_Ring_Generic_GenericRing_dollar_Dict
}

var cache_Data_Ring_Generic_genericSub_prime_ gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub_prime_ sync.Once
func Get_Data_Ring_Generic_genericSub_prime_() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub_prime_.Do(func() {
		cache_Data_Ring_Generic_genericSub_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericSub_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ring_Generic_genericSub_prime_
}

var cache_Data_Ring_Generic_genericSub gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub sync.Once
func Get_Data_Ring_Generic_genericSub() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub.Do(func() {
		cache_Data_Ring_Generic_genericSub = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericRing_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericSub(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]](dictGenericRing_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_Ring_Generic_genericSub
}

var cache_Data_Ring_Generic_genericRingProduct gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingProduct sync.Once
func Get_Data_Ring_Generic_genericRingProduct() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingProduct.Do(func() {
		cache_Data_Ring_Generic_genericRingProduct = gopurs_runtime.Func2(func(dictGenericRing_0_box gopurs_runtime.Value, dictGenericRing1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericRingProduct(dictGenericRing_0_box, dictGenericRing1_1_box)
})
	})
	return cache_Data_Ring_Generic_genericRingProduct
}

var cache_Data_Ring_Generic_genericRingNoArguments gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingNoArguments sync.Once
func Get_Data_Ring_Generic_genericRingNoArguments() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingNoArguments.Do(func() {
		cache_Data_Ring_Generic_genericRingNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_Generic_1088399036_3123454929((&Constructor_Data_Ring_Generic_GenericRing[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Ring_Generic_genericRingNoArguments
}

var cache_Data_Ring_Generic_genericRingConstructor gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingConstructor sync.Once
func Get_Data_Ring_Generic_genericRingConstructor() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingConstructor.Do(func() {
		cache_Data_Ring_Generic_genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericRingConstructor(dictGenericRing_0_box)
})
	})
	return cache_Data_Ring_Generic_genericRingConstructor
}

var cache_Data_Ring_Generic_genericRingArgument gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingArgument sync.Once
func Get_Data_Ring_Generic_genericRingArgument() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingArgument.Do(func() {
		cache_Data_Ring_Generic_genericRingArgument = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericRingArgument(dictRing_0_box)
})
	})
	return cache_Data_Ring_Generic_genericRingArgument
}

type Constructor_Data_Ring_Generic_GenericRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3896698597] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_Generic_GenericRing[any])(ptr)
		_ = c
		switch key {
		case "genericSub'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Ring_Generic_GenericRing: " + key)
		}
	}
}


func Call_Data_Ring_Generic_GenericRing_dollar_Dict(x_0_loop struct{
	genericSub_prime_ gopurs_runtime.Value
}) *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] {
var x_0 struct{
	genericSub_prime_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("genericSub'", orig.genericSub_prime_)
				}())
}

func Call_Data_Ring_Generic_genericSub_prime_(dict_0_loop *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Ring_Generic_genericSub(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericRing_1_loop *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericRing_1 *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] = dictGenericRing_1_loop
_ = dictGenericRing_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictGenericRing_1.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_Ring_Generic_genericRingProduct(dictGenericRing_0_loop gopurs_runtime.Value, dictGenericRing1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
var dictGenericRing1_1 gopurs_runtime.Value = dictGenericRing1_1_loop
_ = dictGenericRing1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_Generic_2426474238_3123454929((&Constructor_Data_Ring_Generic_GenericRing[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
})})))}
}

func Call_Data_Ring_Generic_genericRingConstructor(dictGenericRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
})}))}
}

func Call_Data_Ring_Generic_genericRingArgument(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
})}))}
}

func Rebox_Data_Ring_Generic_1088399036_3123454929(in *Constructor_Data_Ring_Generic_GenericRing[uint32]) *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ring_Generic_2426474238_3123454929(in *Constructor_Data_Ring_Generic_GenericRing[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}


