package Data_Semiring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var cache_genericZero_prime gopurs_runtime.Value
var once_genericZero_prime sync.Once
func Get_genericZero_prime() gopurs_runtime.Value {
	once_genericZero_prime.Do(func() {
		cache_genericZero_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericZero_prime(dict_0_box)
})
	})
	return cache_genericZero_prime
}

var cache_genericZero_prime__gopurs_runtime_Value_2711230462 gopurs_runtime.Value
var once_genericZero_prime__gopurs_runtime_Value_2711230462 sync.Once
func Get_genericZero_prime__gopurs_runtime_Value_2711230462() gopurs_runtime.Value {
	once_genericZero_prime__gopurs_runtime_Value_2711230462.Do(func() {
		cache_genericZero_prime__gopurs_runtime_Value_2711230462 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericZero_prime__gopurs_runtime_Value_2711230462(dict_0_box)
})
	})
	return cache_genericZero_prime__gopurs_runtime_Value_2711230462
}

var cache_genericZero gopurs_runtime.Value
var once_genericZero sync.Once
func Get_genericZero() gopurs_runtime.Value {
	once_genericZero.Do(func() {
		cache_genericZero = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericZero(dictGeneric_0_box, dictGenericSemiring_1_box)
})
	})
	return cache_genericZero
}

var cache_genericSemiringNoArguments gopurs_runtime.Value
var once_genericSemiringNoArguments sync.Once
func Get_genericSemiringNoArguments() gopurs_runtime.Value {
	once_genericSemiringNoArguments.Do(func() {
		cache_genericSemiringNoArguments = gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil})
	})
	return cache_genericSemiringNoArguments
}

var cache_genericSemiringArgument gopurs_runtime.Value
var once_genericSemiringArgument sync.Once
func Get_genericSemiringArgument() gopurs_runtime.Value {
	once_genericSemiringArgument.Do(func() {
		cache_genericSemiringArgument = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemiringArgument(dictSemiring_0_box)
})
	})
	return cache_genericSemiringArgument
}

var cache_genericOne_prime gopurs_runtime.Value
var once_genericOne_prime sync.Once
func Get_genericOne_prime() gopurs_runtime.Value {
	once_genericOne_prime.Do(func() {
		cache_genericOne_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOne_prime(dict_0_box)
})
	})
	return cache_genericOne_prime
}

var cache_genericOne_prime__gopurs_runtime_Value_2711230462 gopurs_runtime.Value
var once_genericOne_prime__gopurs_runtime_Value_2711230462 sync.Once
func Get_genericOne_prime__gopurs_runtime_Value_2711230462() gopurs_runtime.Value {
	once_genericOne_prime__gopurs_runtime_Value_2711230462.Do(func() {
		cache_genericOne_prime__gopurs_runtime_Value_2711230462 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOne_prime__gopurs_runtime_Value_2711230462(dict_0_box)
})
	})
	return cache_genericOne_prime__gopurs_runtime_Value_2711230462
}

var cache_genericOne gopurs_runtime.Value
var once_genericOne sync.Once
func Get_genericOne() gopurs_runtime.Value {
	once_genericOne.Do(func() {
		cache_genericOne = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOne(dictGeneric_0_box, dictGenericSemiring_1_box)
})
	})
	return cache_genericOne
}

var cache_genericMul_prime gopurs_runtime.Value
var once_genericMul_prime sync.Once
func Get_genericMul_prime() gopurs_runtime.Value {
	once_genericMul_prime.Do(func() {
		cache_genericMul_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMul_prime(dict_0_box)
})
	})
	return cache_genericMul_prime
}

var cache_genericMul_prime__gopurs_runtime_Value_1686592254 gopurs_runtime.Value
var once_genericMul_prime__gopurs_runtime_Value_1686592254 sync.Once
func Get_genericMul_prime__gopurs_runtime_Value_1686592254() gopurs_runtime.Value {
	once_genericMul_prime__gopurs_runtime_Value_1686592254.Do(func() {
		cache_genericMul_prime__gopurs_runtime_Value_1686592254 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMul_prime__gopurs_runtime_Value_1686592254(dict_0_box)
})
	})
	return cache_genericMul_prime__gopurs_runtime_Value_1686592254
}

var cache_genericMul gopurs_runtime.Value
var once_genericMul sync.Once
func Get_genericMul() gopurs_runtime.Value {
	once_genericMul.Do(func() {
		cache_genericMul = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMul(dictGeneric_0_box, dictGenericSemiring_1_box, x_2_box, y_3_box)
})
	})
	return cache_genericMul
}

var cache_genericAdd_prime gopurs_runtime.Value
var once_genericAdd_prime sync.Once
func Get_genericAdd_prime() gopurs_runtime.Value {
	once_genericAdd_prime.Do(func() {
		cache_genericAdd_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAdd_prime(dict_0_box)
})
	})
	return cache_genericAdd_prime
}

var cache_genericAdd_prime__gopurs_runtime_Value_1686592254 gopurs_runtime.Value
var once_genericAdd_prime__gopurs_runtime_Value_1686592254 sync.Once
func Get_genericAdd_prime__gopurs_runtime_Value_1686592254() gopurs_runtime.Value {
	once_genericAdd_prime__gopurs_runtime_Value_1686592254.Do(func() {
		cache_genericAdd_prime__gopurs_runtime_Value_1686592254 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAdd_prime__gopurs_runtime_Value_1686592254(dict_0_box)
})
	})
	return cache_genericAdd_prime__gopurs_runtime_Value_1686592254
}

var cache_genericSemiringConstructor gopurs_runtime.Value
var once_genericSemiringConstructor sync.Once
func Get_genericSemiringConstructor() gopurs_runtime.Value {
	once_genericSemiringConstructor.Do(func() {
		cache_genericSemiringConstructor = gopurs_runtime.Func(func(dictGenericSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemiringConstructor(dictGenericSemiring_0_box)
})
	})
	return cache_genericSemiringConstructor
}

var cache_genericSemiringProduct gopurs_runtime.Value
var once_genericSemiringProduct sync.Once
func Get_genericSemiringProduct() gopurs_runtime.Value {
	once_genericSemiringProduct.Do(func() {
		cache_genericSemiringProduct = gopurs_runtime.Func(func(dictGenericSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemiringProduct(dictGenericSemiring_0_box)
})
	})
	return cache_genericSemiringProduct
}

var cache_genericAdd gopurs_runtime.Value
var once_genericAdd sync.Once
func Get_genericAdd() gopurs_runtime.Value {
	once_genericAdd.Do(func() {
		cache_genericAdd = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAdd(dictGeneric_0_box, dictGenericSemiring_1_box, x_2_box, y_3_box)
})
	})
	return cache_genericAdd
}

func Call_genericZero_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericZero'")
}

func Call_genericZero_prime__gopurs_runtime_Value_2711230462(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericZero'")
}

func Call_genericZero(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericZero'"))
}

func Call_genericSemiringArgument(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring_0, "zero"))
}

func Call_genericOne_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericOne'")
}

func Call_genericOne_prime__gopurs_runtime_Value_2711230462(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericOne'")
}

func Call_genericOne(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericOne'"))
}

func Call_genericMul_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMul'")
}

func Call_genericMul_prime__gopurs_runtime_Value_1686592254(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMul'")
}

func Call_genericMul(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericMul'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
}

func Call_genericAdd_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericAdd'")
}

func Call_genericAdd_prime__gopurs_runtime_Value_1686592254(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericAdd'")
}

func Call_genericSemiringConstructor(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'"))
}

func Call_genericSemiringProduct(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
genericZero_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'")
_ = genericZero_prime1_1_0
genericOne_prime1_2_1 := gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'")
_ = genericOne_prime1_2_1
return gopurs_runtime.Func(func(dictGenericSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((genericOne_prime1_2_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(genericOne_prime1_2_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericAdd'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)
return genericOne_prime1_2_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericAdd'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
}
}()
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((genericOne_prime1_2_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(genericOne_prime1_2_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericMul'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)
return genericOne_prime1_2_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericMul'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
}
}()
}), func() gopurs_runtime.Value {
if ((genericOne_prime1_2_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(genericOne_prime1_2_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V0 = genericOne_prime1_2_1
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V1 = gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericOne'")
return genericOne_prime1_2_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericOne_prime1_2_1, gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericOne'")})}
}
}(), func() gopurs_runtime.Value {
if ((genericOne_prime1_2_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(genericOne_prime1_2_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V0 = genericZero_prime1_1_0
(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(genericOne_prime1_2_1.UnsafePtr).V1 = gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericZero'")
return genericOne_prime1_2_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericZero_prime1_1_0, gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericZero'")})}
}
}())
})
}

func Call_genericAdd(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericAdd'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
}


