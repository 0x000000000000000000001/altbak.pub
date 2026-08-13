package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_Generic_GenericSemigroup_dollarDict gopurs_runtime.Value
var once_Data_Semigroup_Generic_GenericSemigroup_dollarDict sync.Once
func Get_Data_Semigroup_Generic_GenericSemigroup_dollarDict() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_GenericSemigroup_dollarDict.Do(func() {
		cache_Data_Semigroup_Generic_GenericSemigroup_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_GenericSemigroup_dollarDict(x_0_box)
})
	})
	return cache_Data_Semigroup_Generic_GenericSemigroup_dollarDict
}

var cache_Data_Semigroup_Generic_genericSemigroupNoConstructors gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericSemigroupNoConstructors sync.Once
func Get_Data_Semigroup_Generic_genericSemigroupNoConstructors() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericSemigroupNoConstructors.Do(func() {
		cache_Data_Semigroup_Generic_genericSemigroupNoConstructors = gopurs_runtime.Value{Type: 9, IntVal: 1022671493, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Generic_GenericSemigroup{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})})}
	})
	return cache_Data_Semigroup_Generic_genericSemigroupNoConstructors
}

var cache_Data_Semigroup_Generic_genericSemigroupNoArguments gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericSemigroupNoArguments sync.Once
func Get_Data_Semigroup_Generic_genericSemigroupNoArguments() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericSemigroupNoArguments.Do(func() {
		cache_Data_Semigroup_Generic_genericSemigroupNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 1022671493, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Generic_GenericSemigroup{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(a_0.IntVal)), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Semigroup_Generic_genericSemigroupNoArguments
}

var cache_Data_Semigroup_Generic_genericSemigroupArgument gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericSemigroupArgument sync.Once
func Get_Data_Semigroup_Generic_genericSemigroupArgument() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericSemigroupArgument.Do(func() {
		cache_Data_Semigroup_Generic_genericSemigroupArgument = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_genericSemigroupArgument(dictSemigroup_0_box)
})
	})
	return cache_Data_Semigroup_Generic_genericSemigroupArgument
}

var cache_Data_Semigroup_Generic_genericAppend_prime gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericAppend_prime sync.Once
func Get_Data_Semigroup_Generic_genericAppend_prime() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericAppend_prime.Do(func() {
		cache_Data_Semigroup_Generic_genericAppend_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_genericAppend_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Generic_GenericSemigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_Generic_genericAppend_prime
}

var cache_Data_Semigroup_Generic_genericSemigroupConstructor gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericSemigroupConstructor sync.Once
func Get_Data_Semigroup_Generic_genericSemigroupConstructor() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericSemigroupConstructor.Do(func() {
		cache_Data_Semigroup_Generic_genericSemigroupConstructor = gopurs_runtime.Func(func(dictGenericSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_genericSemigroupConstructor(dictGenericSemigroup_0_box)
})
	})
	return cache_Data_Semigroup_Generic_genericSemigroupConstructor
}

var cache_Data_Semigroup_Generic_genericSemigroupProduct gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericSemigroupProduct sync.Once
func Get_Data_Semigroup_Generic_genericSemigroupProduct() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericSemigroupProduct.Do(func() {
		cache_Data_Semigroup_Generic_genericSemigroupProduct = gopurs_runtime.Func2(func(dictGenericSemigroup_0_box gopurs_runtime.Value, dictGenericSemigroup1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_genericSemigroupProduct(dictGenericSemigroup_0_box, dictGenericSemigroup1_1_box)
})
	})
	return cache_Data_Semigroup_Generic_genericSemigroupProduct
}

var cache_Data_Semigroup_Generic_genericAppend gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericAppend sync.Once
func Get_Data_Semigroup_Generic_genericAppend() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericAppend.Do(func() {
		cache_Data_Semigroup_Generic_genericAppend = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemigroup_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_genericAppend(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Generic_GenericSemigroup](dictGenericSemigroup_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_Semigroup_Generic_genericAppend
}

var cache_Data_Semigroup_Generic_genericAppend_prime__1736829671 gopurs_runtime.Value
var once_Data_Semigroup_Generic_genericAppend_prime__1736829671 sync.Once
func Get_Data_Semigroup_Generic_genericAppend_prime__1736829671() gopurs_runtime.Value {
	once_Data_Semigroup_Generic_genericAppend_prime__1736829671.Do(func() {
		cache_Data_Semigroup_Generic_genericAppend_prime__1736829671 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Generic_genericAppend_prime__1736829671(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Generic_GenericSemigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_Generic_genericAppend_prime__1736829671
}

type Constructor_Data_Semigroup_Generic_GenericSemigroup struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1022671493] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Generic_GenericSemigroup)(ptr)
		_ = c
		switch key {
		case "genericAppend'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_Generic_GenericSemigroup: " + key)
		}
	}
}


func Call_Data_Semigroup_Generic_GenericSemigroup_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Generic_genericSemigroupArgument(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 1022671493, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Generic_GenericSemigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
})
})})}
}

func Call_Data_Semigroup_Generic_genericAppend_prime(dict_0_loop *Constructor_Data_Semigroup_Generic_GenericSemigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Generic_GenericSemigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_Generic_genericSemigroupConstructor(dictGenericSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemigroup_0 gopurs_runtime.Value = dictGenericSemigroup_0_loop
_ = dictGenericSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 1022671493, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Generic_GenericSemigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_0, "genericAppend'"), v_1, v1_2)
})
})})}
}

func Call_Data_Semigroup_Generic_genericSemigroupProduct(dictGenericSemigroup_0_loop gopurs_runtime.Value, dictGenericSemigroup1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemigroup_0 gopurs_runtime.Value = dictGenericSemigroup_0_loop
_ = dictGenericSemigroup_0
var dictGenericSemigroup1_1 gopurs_runtime.Value = dictGenericSemigroup1_1_loop
_ = dictGenericSemigroup1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1022671493, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Generic_GenericSemigroup{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_0, "genericAppend'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup1_1, "genericAppend'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
})})}
}

func Call_Data_Semigroup_Generic_genericAppend(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericSemigroup_1_loop *Constructor_Data_Semigroup_Generic_GenericSemigroup, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemigroup_1 *Constructor_Data_Semigroup_Generic_GenericSemigroup = dictGenericSemigroup_1_loop
_ = dictGenericSemigroup_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_Semigroup_Generic_genericAppend_prime__1736829671(), gopurs_runtime.Value{Type: 9, IntVal: 1022671493, UnsafePtr: unsafe.Pointer(dictGenericSemigroup_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_Semigroup_Generic_genericAppend_prime__1736829671(dict_0_loop *Constructor_Data_Semigroup_Generic_GenericSemigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Generic_GenericSemigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}


