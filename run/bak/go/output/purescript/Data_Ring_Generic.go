package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ring_Generic_GenericRing_dollarDict gopurs_runtime.Value
var once_Data_Ring_Generic_GenericRing_dollarDict sync.Once
func Get_Data_Ring_Generic_GenericRing_dollarDict() gopurs_runtime.Value {
	once_Data_Ring_Generic_GenericRing_dollarDict.Do(func() {
		cache_Data_Ring_Generic_GenericRing_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_GenericRing_dollarDict(x_0_box)
})
	})
	return cache_Data_Ring_Generic_GenericRing_dollarDict
}

var cache_Data_Ring_Generic_genericSub_prime gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub_prime sync.Once
func Get_Data_Ring_Generic_genericSub_prime() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub_prime.Do(func() {
		cache_Data_Ring_Generic_genericSub_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericSub_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing](dict_0_box))
})
	})
	return cache_Data_Ring_Generic_genericSub_prime
}

var cache_Data_Ring_Generic_genericSub gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub sync.Once
func Get_Data_Ring_Generic_genericSub() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub.Do(func() {
		cache_Data_Ring_Generic_genericSub = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericRing_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericSub(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing](dictGenericRing_1_box), x_2_box, y_3_box)
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
		cache_Data_Ring_Generic_genericRingNoArguments = gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}))
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

var cache_Data_Ring_Generic_genericSub_prime__469823367 gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub_prime__469823367 sync.Once
func Get_Data_Ring_Generic_genericSub_prime__469823367() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub_prime__469823367.Do(func() {
		cache_Data_Ring_Generic_genericSub_prime__469823367 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Generic_genericSub_prime__469823367(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing](dict_0_box))
})
	})
	return cache_Data_Ring_Generic_genericSub_prime__469823367
}

type Constructor_Data_Ring_Generic_GenericRing struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3896698597] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_Generic_GenericRing)(ptr)
		_ = c
		switch key {
		case "genericSub'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Ring_Generic_GenericRing: " + key)
		}
	}
}


func Call_Data_Ring_Generic_GenericRing_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ring_Generic_genericSub_prime(dict_0_loop *Constructor_Data_Ring_Generic_GenericRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Generic_GenericRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Ring_Generic_genericSub(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericRing_1_loop *Constructor_Data_Ring_Generic_GenericRing, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericRing_1 *Constructor_Data_Ring_Generic_GenericRing = dictGenericRing_1_loop
_ = dictGenericRing_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_Ring_Generic_genericSub_prime__469823367(), gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(dictGenericRing_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_Ring_Generic_genericRingProduct(dictGenericRing_0_loop gopurs_runtime.Value, dictGenericRing1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
var dictGenericRing1_1 gopurs_runtime.Value = dictGenericRing1_1_loop
_ = dictGenericRing1_1
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
}))
}

func Call_Data_Ring_Generic_genericRingConstructor(dictGenericRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
})
}))
}

func Call_Data_Ring_Generic_genericRingArgument(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
})
}))
}

func Call_Data_Ring_Generic_genericSub_prime__469823367(dict_0_loop *Constructor_Data_Ring_Generic_GenericRing) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Generic_GenericRing = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}


