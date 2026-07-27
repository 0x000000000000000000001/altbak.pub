package Data_Ring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var cache_genericSub_prime gopurs_runtime.Value
var once_genericSub_prime sync.Once
func Get_genericSub_prime() gopurs_runtime.Value {
	once_genericSub_prime.Do(func() {
		cache_genericSub_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSub_prime(dict_0_box)
})
	})
	return cache_genericSub_prime
}

var cache_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253 gopurs_runtime.Value
var once_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253 sync.Once
func Get_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253() gopurs_runtime.Value {
	once_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253.Do(func() {
		cache_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253(dict_0_box)
})
	})
	return cache_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253
}

var cache_genericSub gopurs_runtime.Value
var once_genericSub sync.Once
func Get_genericSub() gopurs_runtime.Value {
	once_genericSub.Do(func() {
		cache_genericSub = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericRing_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericSub(dictGeneric_0_box, dictGenericRing_1_box, gopurs_runtime.UnboxAny(x_2_box), gopurs_runtime.UnboxAny(y_3_box)))
})
	})
	return cache_genericSub
}

var cache_genericRingProduct gopurs_runtime.Value
var once_genericRingProduct sync.Once
func Get_genericRingProduct() gopurs_runtime.Value {
	once_genericRingProduct.Do(func() {
		cache_genericRingProduct = gopurs_runtime.Func2(func(dictGenericRing_0_box gopurs_runtime.Value, dictGenericRing1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericRingProduct(dictGenericRing_0_box, dictGenericRing1_1_box))
})
	})
	return cache_genericRingProduct
}

var cache_genericRingNoArguments gopurs_runtime.Value
var once_genericRingNoArguments sync.Once
func Get_genericRingNoArguments() gopurs_runtime.Value {
	once_genericRingNoArguments.Do(func() {
		cache_genericRingNoArguments = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil})
}))))
	})
	return cache_genericRingNoArguments
}

var cache_genericRingConstructor gopurs_runtime.Value
var once_genericRingConstructor sync.Once
func Get_genericRingConstructor() gopurs_runtime.Value {
	once_genericRingConstructor.Do(func() {
		cache_genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericRingConstructor(dictGenericRing_0_box))
})
	})
	return cache_genericRingConstructor
}

var cache_genericRingArgument gopurs_runtime.Value
var once_genericRingArgument sync.Once
func Get_genericRingArgument() gopurs_runtime.Value {
	once_genericRingArgument.Do(func() {
		cache_genericRingArgument = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericRingArgument(dictRing_0_box))
})
	})
	return cache_genericRingArgument
}

func Call_genericSub_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericSub'")
}

func Call_genericSub_prime__func_gopurs_runtime_Value__interface____interface____interface___1790541253(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericSub'")
}

func Call_genericSub(dictGeneric_0_loop gopurs_runtime.Value, dictGenericRing_1_loop gopurs_runtime.Value, x_2_loop interface{}, y_3_loop interface{}) interface{} {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericRing_1 gopurs_runtime.Value = dictGenericRing_1_loop
_ = dictGenericRing_1
var x_2 interface{} = x_2_loop
_ = x_2
var y_3 interface{} = y_3_loop
_ = y_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_1, "genericSub'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), gopurs_runtime.Any(x_2)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), gopurs_runtime.Any(y_3)))))
}

func Call_genericRingProduct(dictGenericRing_0_loop gopurs_runtime.Value, dictGenericRing1_1_loop gopurs_runtime.Value) interface{} {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
var dictGenericRing1_1 gopurs_runtime.Value = dictGenericRing1_1_loop
_ = dictGenericRing1_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))), gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)))})})
})))
}

func Call_genericRingConstructor(dictGenericRing_0_loop gopurs_runtime.Value) interface{} {
var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
_ = dictGenericRing_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
})))
}

func Call_genericRingArgument(dictRing_0_loop gopurs_runtime.Value) interface{} {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
})))
}
