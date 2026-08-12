package Data_Semigroup_Generic

import (
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genericSemigroupNoConstructors gopurs_runtime.Value
var once_genericSemigroupNoConstructors sync.Once
func Get_genericSemigroupNoConstructors() gopurs_runtime.Value {
	once_genericSemigroupNoConstructors.Do(func() {
		cache_genericSemigroupNoConstructors = gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
}))
	})
	return cache_genericSemigroupNoConstructors
}

var cache_genericSemigroupNoArguments gopurs_runtime.Value
var once_genericSemigroupNoArguments sync.Once
func Get_genericSemigroupNoArguments() gopurs_runtime.Value {
	once_genericSemigroupNoArguments.Do(func() {
		cache_genericSemigroupNoArguments = gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
}))
	})
	return cache_genericSemigroupNoArguments
}

var cache_genericSemigroupArgument gopurs_runtime.Value
var once_genericSemigroupArgument sync.Once
func Get_genericSemigroupArgument() gopurs_runtime.Value {
	once_genericSemigroupArgument.Do(func() {
		cache_genericSemigroupArgument = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemigroupArgument(dictSemigroup_0_box)
})
	})
	return cache_genericSemigroupArgument
}

var cache_genericAppend_prime gopurs_runtime.Value
var once_genericAppend_prime sync.Once
func Get_genericAppend_prime() gopurs_runtime.Value {
	once_genericAppend_prime.Do(func() {
		cache_genericAppend_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAppend_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericSemigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericAppend_prime
}

var cache_genericAppend_prime__gopurs_runtime_Value_1736829671 gopurs_runtime.Value
var once_genericAppend_prime__gopurs_runtime_Value_1736829671 sync.Once
func Get_genericAppend_prime__gopurs_runtime_Value_1736829671() gopurs_runtime.Value {
	once_genericAppend_prime__gopurs_runtime_Value_1736829671.Do(func() {
		cache_genericAppend_prime__gopurs_runtime_Value_1736829671 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAppend_prime__gopurs_runtime_Value_1736829671(gopurs_runtime.CoerceToStruct[Constructor_GenericSemigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericAppend_prime__gopurs_runtime_Value_1736829671
}

var cache_genericSemigroupConstructor gopurs_runtime.Value
var once_genericSemigroupConstructor sync.Once
func Get_genericSemigroupConstructor() gopurs_runtime.Value {
	once_genericSemigroupConstructor.Do(func() {
		cache_genericSemigroupConstructor = gopurs_runtime.Func(func(dictGenericSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemigroupConstructor(dictGenericSemigroup_0_box)
})
	})
	return cache_genericSemigroupConstructor
}

var cache_genericSemigroupProduct gopurs_runtime.Value
var once_genericSemigroupProduct sync.Once
func Get_genericSemigroupProduct() gopurs_runtime.Value {
	once_genericSemigroupProduct.Do(func() {
		cache_genericSemigroupProduct = gopurs_runtime.Func2(func(dictGenericSemigroup_0_box gopurs_runtime.Value, dictGenericSemigroup1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemigroupProduct(dictGenericSemigroup_0_box, dictGenericSemigroup1_1_box)
})
	})
	return cache_genericSemigroupProduct
}

var cache_genericAppend gopurs_runtime.Value
var once_genericAppend sync.Once
func Get_genericAppend() gopurs_runtime.Value {
	once_genericAppend.Do(func() {
		cache_genericAppend = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemigroup_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAppend(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericSemigroup[gopurs_runtime.Value]](dictGenericSemigroup_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericAppend
}

type Constructor_GenericSemigroup[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1022671493] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericSemigroup[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericAppend'": return c.V0
		default: panic("Key not found in dictionary Constructor_GenericSemigroup: " + key)
		}
	}
}


func Call_genericSemigroupArgument(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
})
}))
}

func Call_genericAppend_prime(dict_0_loop *Constructor_GenericSemigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericSemigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericAppend_prime__gopurs_runtime_Value_1736829671(dict_0_loop *Constructor_GenericSemigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericSemigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericSemigroupConstructor(dictGenericSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemigroup_0 gopurs_runtime.Value = dictGenericSemigroup_0_loop
_ = dictGenericSemigroup_0
return gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_0, "genericAppend'"), v_1, v1_2)
})
}))
}

func Call_genericSemigroupProduct(dictGenericSemigroup_0_loop gopurs_runtime.Value, dictGenericSemigroup1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemigroup_0 gopurs_runtime.Value = dictGenericSemigroup_0_loop
_ = dictGenericSemigroup_0
var dictGenericSemigroup1_1 gopurs_runtime.Value = dictGenericSemigroup1_1_loop
_ = dictGenericSemigroup1_1
return gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_0, "genericAppend'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup1_1, "genericAppend'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}))}
})
}))
}

func Call_genericAppend(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericSemigroup_1_loop *Constructor_GenericSemigroup[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemigroup_1 *Constructor_GenericSemigroup[gopurs_runtime.Value] = dictGenericSemigroup_1_loop
_ = dictGenericSemigroup_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply2(dictGenericSemigroup_1.V0, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}


