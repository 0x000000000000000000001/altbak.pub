package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semiring_Generic_GenericSemiring_dollarDict gopurs_runtime.Value
var once_Data_Semiring_Generic_GenericSemiring_dollarDict sync.Once
func Get_Data_Semiring_Generic_GenericSemiring_dollarDict() gopurs_runtime.Value {
	once_Data_Semiring_Generic_GenericSemiring_dollarDict.Do(func() {
		cache_Data_Semiring_Generic_GenericSemiring_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_GenericSemiring_dollarDict(x_0_box)
})
	})
	return cache_Data_Semiring_Generic_GenericSemiring_dollarDict
}

var cache_Data_Semiring_Generic_genericZero_prime gopurs_runtime.Value
var once_Data_Semiring_Generic_genericZero_prime sync.Once
func Get_Data_Semiring_Generic_genericZero_prime() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericZero_prime.Do(func() {
		cache_Data_Semiring_Generic_genericZero_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericZero_prime(dict_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericZero_prime
}

var cache_Data_Semiring_Generic_genericZero gopurs_runtime.Value
var once_Data_Semiring_Generic_genericZero sync.Once
func Get_Data_Semiring_Generic_genericZero() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericZero.Do(func() {
		cache_Data_Semiring_Generic_genericZero = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericZero(dictGeneric_0_box, dictGenericSemiring_1_box)
})
	})
	return cache_Data_Semiring_Generic_genericZero
}

var cache_Data_Semiring_Generic_genericSemiringNoArguments gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringNoArguments sync.Once
func Get_Data_Semiring_Generic_genericSemiringNoArguments() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringNoArguments.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semiring_Generic_GenericSemiring{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}})}
	})
	return cache_Data_Semiring_Generic_genericSemiringNoArguments
}

var cache_Data_Semiring_Generic_genericSemiringArgument gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringArgument sync.Once
func Get_Data_Semiring_Generic_genericSemiringArgument() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringArgument.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringArgument = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericSemiringArgument(dictSemiring_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericSemiringArgument
}

var cache_Data_Semiring_Generic_genericOne_prime gopurs_runtime.Value
var once_Data_Semiring_Generic_genericOne_prime sync.Once
func Get_Data_Semiring_Generic_genericOne_prime() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericOne_prime.Do(func() {
		cache_Data_Semiring_Generic_genericOne_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericOne_prime(dict_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericOne_prime
}

var cache_Data_Semiring_Generic_genericOne gopurs_runtime.Value
var once_Data_Semiring_Generic_genericOne sync.Once
func Get_Data_Semiring_Generic_genericOne() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericOne.Do(func() {
		cache_Data_Semiring_Generic_genericOne = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericOne(dictGeneric_0_box, dictGenericSemiring_1_box)
})
	})
	return cache_Data_Semiring_Generic_genericOne
}

var cache_Data_Semiring_Generic_genericMul_prime gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul_prime sync.Once
func Get_Data_Semiring_Generic_genericMul_prime() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul_prime.Do(func() {
		cache_Data_Semiring_Generic_genericMul_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericMul_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring](dict_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericMul_prime
}

var cache_Data_Semiring_Generic_genericMul gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul sync.Once
func Get_Data_Semiring_Generic_genericMul() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul.Do(func() {
		cache_Data_Semiring_Generic_genericMul = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericMul(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring](dictGenericSemiring_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_Semiring_Generic_genericMul
}

var cache_Data_Semiring_Generic_genericAdd_prime gopurs_runtime.Value
var once_Data_Semiring_Generic_genericAdd_prime sync.Once
func Get_Data_Semiring_Generic_genericAdd_prime() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericAdd_prime.Do(func() {
		cache_Data_Semiring_Generic_genericAdd_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericAdd_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring](dict_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericAdd_prime
}

var cache_Data_Semiring_Generic_genericSemiringConstructor gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringConstructor sync.Once
func Get_Data_Semiring_Generic_genericSemiringConstructor() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringConstructor.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringConstructor = gopurs_runtime.Func(func(dictGenericSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericSemiringConstructor(dictGenericSemiring_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericSemiringConstructor
}

var cache_Data_Semiring_Generic_genericSemiringProduct gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringProduct sync.Once
func Get_Data_Semiring_Generic_genericSemiringProduct() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringProduct.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringProduct = gopurs_runtime.Func2(func(dictGenericSemiring_0_box gopurs_runtime.Value, dictGenericSemiring1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericSemiringProduct(dictGenericSemiring_0_box, dictGenericSemiring1_1_box)
})
	})
	return cache_Data_Semiring_Generic_genericSemiringProduct
}

var cache_Data_Semiring_Generic_genericAdd gopurs_runtime.Value
var once_Data_Semiring_Generic_genericAdd sync.Once
func Get_Data_Semiring_Generic_genericAdd() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericAdd.Do(func() {
		cache_Data_Semiring_Generic_genericAdd = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericAdd(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring](dictGenericSemiring_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_Semiring_Generic_genericAdd
}

var cache_Data_Semiring_Generic_genericAdd_prime__1201800327 gopurs_runtime.Value
var once_Data_Semiring_Generic_genericAdd_prime__1201800327 sync.Once
func Get_Data_Semiring_Generic_genericAdd_prime__1201800327() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericAdd_prime__1201800327.Do(func() {
		cache_Data_Semiring_Generic_genericAdd_prime__1201800327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericAdd_prime__1201800327(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring](dict_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericAdd_prime__1201800327
}

var cache_Data_Semiring_Generic_genericMul_prime__1201800327 gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul_prime__1201800327 sync.Once
func Get_Data_Semiring_Generic_genericMul_prime__1201800327() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul_prime__1201800327.Do(func() {
		cache_Data_Semiring_Generic_genericMul_prime__1201800327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericMul_prime__1201800327(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring](dict_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericMul_prime__1201800327
}

var cache_Data_Semiring_Generic_genericOne_prime__3878335718 gopurs_runtime.Value
var once_Data_Semiring_Generic_genericOne_prime__3878335718 sync.Once
func Get_Data_Semiring_Generic_genericOne_prime__3878335718() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericOne_prime__3878335718.Do(func() {
		cache_Data_Semiring_Generic_genericOne_prime__3878335718 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericOne_prime__3878335718(dict_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericOne_prime__3878335718
}

var cache_Data_Semiring_Generic_genericZero_prime__3878335718 gopurs_runtime.Value
var once_Data_Semiring_Generic_genericZero_prime__3878335718 sync.Once
func Get_Data_Semiring_Generic_genericZero_prime__3878335718() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericZero_prime__3878335718.Do(func() {
		cache_Data_Semiring_Generic_genericZero_prime__3878335718 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericZero_prime__3878335718(dict_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericZero_prime__3878335718
}

type Constructor_Data_Semiring_Generic_GenericSemiring struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3093875941] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semiring_Generic_GenericSemiring)(ptr)
		_ = c
		switch key {
		case "genericAdd'": return gopurs_runtime.Box(c.V0)
		case "genericMul'": return gopurs_runtime.Box(c.V1)
		case "genericOne'": return gopurs_runtime.Box(c.V2)
		case "genericZero'": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Semiring_Generic_GenericSemiring: " + key)
		}
	}
}


func Call_Data_Semiring_Generic_GenericSemiring_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semiring_Generic_genericZero_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericZero'")
}

func Call_Data_Semiring_Generic_genericZero(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericZero'"))
}

func Call_Data_Semiring_Generic_genericSemiringArgument(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semiring_Generic_GenericSemiring{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring_0, "zero")})}
}

func Call_Data_Semiring_Generic_genericOne_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericOne'")
}

func Call_Data_Semiring_Generic_genericOne(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericOne'"))
}

func Call_Data_Semiring_Generic_genericMul_prime(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_Generic_genericMul(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericSemiring_1_loop *Constructor_Data_Semiring_Generic_GenericSemiring, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 *Constructor_Data_Semiring_Generic_GenericSemiring = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_Semiring_Generic_genericMul_prime__1201800327(), gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(dictGenericSemiring_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_Semiring_Generic_genericAdd_prime(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_Generic_genericSemiringConstructor(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semiring_Generic_GenericSemiring{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'")})}
}

func Call_Data_Semiring_Generic_genericSemiringProduct(dictGenericSemiring_0_loop gopurs_runtime.Value, dictGenericSemiring1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
var dictGenericSemiring1_1 gopurs_runtime.Value = dictGenericSemiring1_1_loop
_ = dictGenericSemiring1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semiring_Generic_GenericSemiring{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericAdd'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericMul'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericOne'")})}, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'"), gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericZero'")})}})}
}

func Call_Data_Semiring_Generic_genericAdd(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericSemiring_1_loop *Constructor_Data_Semiring_Generic_GenericSemiring, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 *Constructor_Data_Semiring_Generic_GenericSemiring = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_Semiring_Generic_genericAdd_prime__1201800327(), gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(dictGenericSemiring_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_Semiring_Generic_genericAdd_prime__1201800327(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_Generic_genericMul_prime__1201800327(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_Generic_genericOne_prime__3878335718(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericOne'")
}

func Call_Data_Semiring_Generic_genericZero_prime__3878335718(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericZero'")
}


