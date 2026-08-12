package Data_Ring_Generic

import (
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genericSub_prime gopurs_runtime.Value
var once_genericSub_prime sync.Once
func Get_genericSub_prime() gopurs_runtime.Value {
	once_genericSub_prime.Do(func() {
		cache_genericSub_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSub_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericSub_prime
}

var cache_genericSub_prime__gopurs_runtime_Value_469823367 gopurs_runtime.Value
var once_genericSub_prime__gopurs_runtime_Value_469823367 sync.Once
func Get_genericSub_prime__gopurs_runtime_Value_469823367() gopurs_runtime.Value {
	once_genericSub_prime__gopurs_runtime_Value_469823367.Do(func() {
		cache_genericSub_prime__gopurs_runtime_Value_469823367 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSub_prime__gopurs_runtime_Value_469823367(gopurs_runtime.CoerceToStruct[Constructor_GenericRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericSub_prime__gopurs_runtime_Value_469823367
}

var cache_genericSub gopurs_runtime.Value
var once_genericSub sync.Once
func Get_genericSub() gopurs_runtime.Value {
	once_genericSub.Do(func() {
		cache_genericSub = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericRing_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSub(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericRing[gopurs_runtime.Value]](dictGenericRing_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericSub
}

var cache_genericRingProduct gopurs_runtime.Value
var once_genericRingProduct sync.Once
func Get_genericRingProduct() gopurs_runtime.Value {
	once_genericRingProduct.Do(func() {
		cache_genericRingProduct = gopurs_runtime.Func2(func(dictGenericRing_0_box gopurs_runtime.Value, dictGenericRing1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericRingProduct(dictGenericRing_0_box, dictGenericRing1_1_box)
})
	})
	return cache_genericRingProduct
}

var cache_genericRingNoArguments gopurs_runtime.Value
var once_genericRingNoArguments sync.Once
func Get_genericRingNoArguments() gopurs_runtime.Value {
	once_genericRingNoArguments.Do(func() {
		cache_genericRingNoArguments = gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_genericRingNoArguments
}

var cache_genericRingConstructor gopurs_runtime.Value
var once_genericRingConstructor sync.Once
func Get_genericRingConstructor() gopurs_runtime.Value {
	once_genericRingConstructor.Do(func() {
		cache_genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericRingConstructor(dictGenericRing_0_box)
})
	})
	return cache_genericRingConstructor
}

var cache_genericRingArgument gopurs_runtime.Value
var once_genericRingArgument sync.Once
func Get_genericRingArgument() gopurs_runtime.Value {
	once_genericRingArgument.Do(func() {
		cache_genericRingArgument = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericRingArgument(dictRing_0_box)
})
	})
	return cache_genericRingArgument
}

type Constructor_GenericRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3896698597] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericRing[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericSub'": return c.V0
		default: panic("Key not found in dictionary Constructor_GenericRing: " + key)
		}
	}
}


func Call_genericSub_prime(dict_0_loop *Constructor_GenericRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericSub_prime__gopurs_runtime_Value_469823367(dict_0_loop *Constructor_GenericRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericSub(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericRing_1_loop *Constructor_GenericRing[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericRing_1 *Constructor_GenericRing[gopurs_runtime.Value] = dictGenericRing_1_loop
_ = dictGenericRing_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply2(dictGenericRing_1.V0, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}

func Call_genericRingProduct(dictGenericRing_0_loop gopurs_runtime.Value, dictGenericRing1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
var dictGenericRing1_1 gopurs_runtime.Value = dictGenericRing1_1_loop
_ = dictGenericRing1_1
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}))}
})
}))
}

func Call_genericRingConstructor(dictGenericRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
})
}))
}

func Call_genericRingArgument(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
})
}))
}


