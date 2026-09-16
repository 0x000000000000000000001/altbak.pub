package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict gopurs_runtime.Value
var once_Data_Semiring_Generic_GenericSemiring_dollar_Dict sync.Once
func Get_Data_Semiring_Generic_GenericSemiring_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semiring_Generic_GenericSemiring_dollar_Dict.Do(func() {
		cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict(func() struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ gopurs_runtime.Value
	genericZero_prime_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ gopurs_runtime.Value
	genericZero_prime_ gopurs_runtime.Value
}{}
					clone.genericAdd_prime_ = gopurs_runtime.RecordGet(orig, "genericAdd'")
					clone.genericMul_prime_ = gopurs_runtime.RecordGet(orig, "genericMul'")
					clone.genericOne_prime_ = gopurs_runtime.RecordGet(orig, "genericOne'")
					clone.genericZero_prime_ = gopurs_runtime.RecordGet(orig, "genericZero'")
					return clone
				}()))}
})
	})
	return cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict
}

var cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835 gopurs_runtime.Value
var once_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835 sync.Once
func Get_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835() gopurs_runtime.Value {
	once_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835.Do(func() {
		cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Rebox_Data_Semiring_Generic_2984329148_1889332433(Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835(func() struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ uint32
	genericZero_prime_ uint32
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ uint32
	genericZero_prime_ uint32
}{}
					clone.genericAdd_prime_ = gopurs_runtime.RecordGet(orig, "genericAdd'")
					clone.genericMul_prime_ = gopurs_runtime.RecordGet(orig, "genericMul'")
					clone.genericOne_prime_ = uint32(gopurs_runtime.RecordGet(orig, "genericOne'").IntVal)
					clone.genericZero_prime_ = uint32(gopurs_runtime.RecordGet(orig, "genericZero'").IntVal)
					return clone
				}())))}
})
	})
	return cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835
}

var cache_Data_Semiring_Generic_genericZero_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericZero_prime_ sync.Once
func Get_Data_Semiring_Generic_genericZero_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericZero_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericZero_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericZero_prime_(dict_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericZero_prime_
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
		cache_Data_Semiring_Generic_genericSemiringNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Rebox_Data_Semiring_Generic_2984329148_1889332433((&Constructor_Data_Semiring_Generic_GenericSemiring[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), 1454898258, 1454898258})))}
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

var cache_Data_Semiring_Generic_genericOne_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericOne_prime_ sync.Once
func Get_Data_Semiring_Generic_genericOne_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericOne_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericOne_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericOne_prime_(dict_0_box)
})
	})
	return cache_Data_Semiring_Generic_genericOne_prime_
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

var cache_Data_Semiring_Generic_genericMul_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul_prime_ sync.Once
func Get_Data_Semiring_Generic_genericMul_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericMul_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericMul_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericMul_prime_
}

var cache_Data_Semiring_Generic_genericMul gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul sync.Once
func Get_Data_Semiring_Generic_genericMul() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul.Do(func() {
		cache_Data_Semiring_Generic_genericMul = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericMul(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericMul
}

var cache_Data_Semiring_Generic_genericAdd_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericAdd_prime_ sync.Once
func Get_Data_Semiring_Generic_genericAdd_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericAdd_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericAdd_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericAdd_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericAdd_prime_
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
		cache_Data_Semiring_Generic_genericAdd = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Generic_genericAdd(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
})
	})
	return cache_Data_Semiring_Generic_genericAdd
}

type Constructor_Data_Semiring_Generic_GenericSemiring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 T_a
}


func init() {
	gopurs_runtime.StructGetters[3093875941] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value])(ptr)
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


func Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict(x_0_loop struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ gopurs_runtime.Value
	genericZero_prime_ gopurs_runtime.Value
}) *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] {
var x_0 struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ gopurs_runtime.Value
	genericZero_prime_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", orig.genericAdd_prime_, orig.genericMul_prime_, orig.genericOne_prime_, orig.genericZero_prime_)
				}())
}

func Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835(x_0_loop struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ uint32
	genericZero_prime_ uint32
}) *Constructor_Data_Semiring_Generic_GenericSemiring[uint32] {
GenericSemiring_dollar_Dict__1911646835:
for {
if false { continue GenericSemiring_dollar_Dict__1911646835 }
var x_0 struct{
	genericAdd_prime_ gopurs_runtime.Value
	genericMul_prime_ gopurs_runtime.Value
	genericOne_prime_ uint32
	genericZero_prime_ uint32
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", orig.genericAdd_prime_, orig.genericMul_prime_, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.genericOne_prime_), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.genericZero_prime_), UnsafePtr: nil})
				}())
}
}

func Call_Data_Semiring_Generic_genericZero_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring_0, "zero")}))}
}

func Call_Data_Semiring_Generic_genericOne_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Semiring_Generic_genericMul_prime_(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Data_Semiring_Generic_genericMul(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
// TAST (Let): to_1_0 shape=App(Var) bindingType=(Func [Any] (TypeVar a$scope21))
to_1_0 := Call_Data_Generic_Rep_to(dictGeneric_0)
_ = to_1_0
return gopurs_runtime.Func3(func(dictGenericSemiring_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(to_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_2, "genericMul'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_3), gopurs_runtime.Apply(dictGeneric_0.V0, y_4)))
})
}

func Call_Data_Semiring_Generic_genericAdd_prime_(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_Data_Semiring_Generic_genericSemiringConstructor(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'")}))}
}

func Call_Data_Semiring_Generic_genericSemiringProduct(dictGenericSemiring_0_loop gopurs_runtime.Value, dictGenericSemiring1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
var dictGenericSemiring1_1 gopurs_runtime.Value = dictGenericSemiring1_1_loop
_ = dictGenericSemiring1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Rebox_Data_Semiring_Generic_180684798_1889332433((&Constructor_Data_Semiring_Generic_GenericSemiring[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericAdd'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericMul'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
}), (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericOne'")}), (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'"), gopurs_runtime.RecordGet(dictGenericSemiring1_1, "genericZero'")})})))}
}

func Call_Data_Semiring_Generic_genericAdd(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
// TAST (Let): to_1_0 shape=App(Var) bindingType=(Func [Any] (TypeVar a$scope52))
to_1_0 := Call_Data_Generic_Rep_to(dictGeneric_0)
_ = to_1_0
return gopurs_runtime.Func3(func(dictGenericSemiring_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(to_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_2, "genericAdd'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_3), gopurs_runtime.Apply(dictGeneric_0.V0, y_4)))
})
}

func Rebox_Data_Semiring_Generic_180684798_1889332433(in *Constructor_Data_Semiring_Generic_GenericSemiring[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V2)}
		out.V3 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V3)}
	return out
}

func Rebox_Data_Semiring_Generic_2984329148_1889332433(in *Constructor_Data_Semiring_Generic_GenericSemiring[uint32]) *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V3), UnsafePtr: nil}
	return out
}


