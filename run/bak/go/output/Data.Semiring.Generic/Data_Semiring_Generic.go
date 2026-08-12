package Data_Semiring_Generic

import (
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_genericSemiringNoArguments = gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil})
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
return Call_genericMul_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericMul_prime
}

var cache_genericMul gopurs_runtime.Value
var once_genericMul sync.Once
func Get_genericMul() gopurs_runtime.Value {
	once_genericMul.Do(func() {
		cache_genericMul = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMul(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericSemiring[gopurs_runtime.Value]](dictGenericSemiring_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericMul
}

var cache_genericAdd_prime gopurs_runtime.Value
var once_genericAdd_prime sync.Once
func Get_genericAdd_prime() gopurs_runtime.Value {
	once_genericAdd_prime.Do(func() {
		cache_genericAdd_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAdd_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericAdd_prime
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
		cache_genericSemiringProduct = gopurs_runtime.Func2(func(dictGenericSemiring_0_box gopurs_runtime.Value, dictGenericSemiring1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSemiringProduct(dictGenericSemiring_0_box, dictGenericSemiring1_1_box)
})
	})
	return cache_genericSemiringProduct
}

var cache_genericAdd gopurs_runtime.Value
var once_genericAdd sync.Once
func Get_genericAdd() gopurs_runtime.Value {
	once_genericAdd.Do(func() {
		cache_genericAdd = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAdd(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericSemiring[gopurs_runtime.Value]](dictGenericSemiring_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericAdd
}

var cache_from__1498760952 gopurs_runtime.Value
var once_from__1498760952 sync.Once
func Get_from__1498760952() gopurs_runtime.Value {
	once_from__1498760952.Do(func() {
		cache_from__1498760952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from__1498760952(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_from__1498760952
}

var cache_to__1498760952 gopurs_runtime.Value
var once_to__1498760952 sync.Once
func Get_to__1498760952() gopurs_runtime.Value {
	once_to__1498760952.Do(func() {
		cache_to__1498760952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_to__1498760952(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_to__1498760952
}

var cache_genericAdd_prime__1201800327 gopurs_runtime.Value
var once_genericAdd_prime__1201800327 sync.Once
func Get_genericAdd_prime__1201800327() gopurs_runtime.Value {
	once_genericAdd_prime__1201800327.Do(func() {
		cache_genericAdd_prime__1201800327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericAdd_prime__1201800327(gopurs_runtime.CoerceToStruct[Constructor_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericAdd_prime__1201800327
}

var cache_genericMul_prime__1201800327 gopurs_runtime.Value
var once_genericMul_prime__1201800327 sync.Once
func Get_genericMul_prime__1201800327() gopurs_runtime.Value {
	once_genericMul_prime__1201800327.Do(func() {
		cache_genericMul_prime__1201800327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMul_prime__1201800327(gopurs_runtime.CoerceToStruct[Constructor_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericMul_prime__1201800327
}

var cache_genericOne_prime__3878335718 gopurs_runtime.Value
var once_genericOne_prime__3878335718 sync.Once
func Get_genericOne_prime__3878335718() gopurs_runtime.Value {
	once_genericOne_prime__3878335718.Do(func() {
		cache_genericOne_prime__3878335718 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericOne_prime__3878335718(dict_0_box)
})
	})
	return cache_genericOne_prime__3878335718
}

var cache_genericZero_prime__3878335718 gopurs_runtime.Value
var once_genericZero_prime__3878335718 sync.Once
func Get_genericZero_prime__3878335718() gopurs_runtime.Value {
	once_genericZero_prime__3878335718.Do(func() {
		cache_genericZero_prime__3878335718 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericZero_prime__3878335718(dict_0_box)
})
	})
	return cache_genericZero_prime__3878335718
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

type Constructor_GenericSemiring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 T_a
}


func init() {
	gopurs_runtime.StructGetters[3093875941] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericSemiring[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericAdd'": return c.V0
		case "genericMul'": return c.V1
		case "genericOne'": return c.V2
		case "genericZero'": return c.V3
		default: panic("Key not found in dictionary Constructor_GenericSemiring: " + key)
		}
	}
}


func Call_genericZero_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring_0, "zero"))
}

func Call_genericOne_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_genericMul_prime(dict_0_loop *Constructor_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericMul(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericSemiring_1_loop *Constructor_GenericSemiring[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 *Constructor_GenericSemiring[gopurs_runtime.Value] = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply3(Get_genericMul_prime__1201800327(), gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(dictGenericSemiring_1)}, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}

func Call_genericAdd_prime(dict_0_loop *Constructor_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericSemiringConstructor(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'"))
}

func Call_genericSemiringProduct(dictGenericSemiring_0_loop gopurs_runtime.Value, dictGenericSemiring1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
var dictGenericSemiring1_1 gopurs_runtime.Value = dictGenericSemiring1_1_loop
_ = dictGenericSemiring1_1
return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericAdd'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericMul'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericOne'")})}, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'"), gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericZero'")})})
}

func Call_genericAdd(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericSemiring_1_loop *Constructor_GenericSemiring[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 *Constructor_GenericSemiring[gopurs_runtime.Value] = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply3(Get_genericAdd_prime__1201800327(), gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(dictGenericSemiring_1)}, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}

func Call_from__1498760952(dict_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_to__1498760952(dict_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericAdd_prime__1201800327(dict_0_loop *Constructor_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericMul_prime__1201800327(dict_0_loop *Constructor_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericOne_prime__3878335718(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericOne'")
}

func Call_genericZero_prime__3878335718(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericZero'")
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}


