package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ord_Generic_GenericOrd_dollar_Dict gopurs_runtime.Value
var once_Data_Ord_Generic_GenericOrd_dollar_Dict sync.Once
func Get_Data_Ord_Generic_GenericOrd_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ord_Generic_GenericOrd_dollar_Dict.Do(func() {
		cache_Data_Ord_Generic_GenericOrd_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer(Call_Data_Ord_Generic_GenericOrd_dollar_Dict(func() struct{
	genericCompare_prime_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericCompare_prime_ gopurs_runtime.Value
}{}
					clone.genericCompare_prime_ = gopurs_runtime.RecordGet(orig, "genericCompare'")
					return clone
				}()))}
})
	})
	return cache_Data_Ord_Generic_GenericOrd_dollar_Dict
}

var cache_Data_Ord_Generic_genericOrdNoConstructors gopurs_runtime.Value
var once_Data_Ord_Generic_genericOrdNoConstructors sync.Once
func Get_Data_Ord_Generic_genericOrdNoConstructors() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericOrdNoConstructors.Do(func() {
		cache_Data_Ord_Generic_genericOrdNoConstructors = gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Ord_Generic_genericOrdNoConstructors
}

var cache_Data_Ord_Generic_genericOrdNoArguments gopurs_runtime.Value
var once_Data_Ord_Generic_genericOrdNoArguments sync.Once
func Get_Data_Ord_Generic_genericOrdNoArguments() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericOrdNoArguments.Do(func() {
		cache_Data_Ord_Generic_genericOrdNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_Generic_2756679068_3050046641((&Constructor_Data_Ord_Generic_GenericOrd[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Ord_Generic_genericOrdNoArguments
}

var cache_Data_Ord_Generic_genericOrdArgument gopurs_runtime.Value
var once_Data_Ord_Generic_genericOrdArgument sync.Once
func Get_Data_Ord_Generic_genericOrdArgument() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericOrdArgument.Do(func() {
		cache_Data_Ord_Generic_genericOrdArgument = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Generic_genericOrdArgument(dictOrd_0_box)
})
	})
	return cache_Data_Ord_Generic_genericOrdArgument
}

var cache_Data_Ord_Generic_genericCompare_prime_ gopurs_runtime.Value
var once_Data_Ord_Generic_genericCompare_prime_ sync.Once
func Get_Data_Ord_Generic_genericCompare_prime_() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericCompare_prime_.Do(func() {
		cache_Data_Ord_Generic_genericCompare_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Generic_genericCompare_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ord_Generic_genericCompare_prime_
}

var cache_Data_Ord_Generic_genericOrdConstructor gopurs_runtime.Value
var once_Data_Ord_Generic_genericOrdConstructor sync.Once
func Get_Data_Ord_Generic_genericOrdConstructor() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericOrdConstructor.Do(func() {
		cache_Data_Ord_Generic_genericOrdConstructor = gopurs_runtime.Func(func(dictGenericOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Generic_genericOrdConstructor(dictGenericOrd_0_box)
})
	})
	return cache_Data_Ord_Generic_genericOrdConstructor
}

var cache_Data_Ord_Generic_genericOrdProduct gopurs_runtime.Value
var once_Data_Ord_Generic_genericOrdProduct sync.Once
func Get_Data_Ord_Generic_genericOrdProduct() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericOrdProduct.Do(func() {
		cache_Data_Ord_Generic_genericOrdProduct = gopurs_runtime.Func2(func(dictGenericOrd_0_box gopurs_runtime.Value, dictGenericOrd1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Generic_genericOrdProduct(dictGenericOrd_0_box, dictGenericOrd1_1_box)
})
	})
	return cache_Data_Ord_Generic_genericOrdProduct
}

var cache_Data_Ord_Generic_genericOrdSum gopurs_runtime.Value
var once_Data_Ord_Generic_genericOrdSum sync.Once
func Get_Data_Ord_Generic_genericOrdSum() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericOrdSum.Do(func() {
		cache_Data_Ord_Generic_genericOrdSum = gopurs_runtime.Func2(func(dictGenericOrd_0_box gopurs_runtime.Value, dictGenericOrd1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Generic_genericOrdSum(dictGenericOrd_0_box, dictGenericOrd1_1_box)
})
	})
	return cache_Data_Ord_Generic_genericOrdSum
}

var cache_Data_Ord_Generic_genericCompare gopurs_runtime.Value
var once_Data_Ord_Generic_genericCompare sync.Once
func Get_Data_Ord_Generic_genericCompare() gopurs_runtime.Value {
	once_Data_Ord_Generic_genericCompare.Do(func() {
		cache_Data_Ord_Generic_genericCompare = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericOrd_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ord_Generic_genericCompare(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]](dictGenericOrd_1_box), x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Ord_Generic_genericCompare
}

type Constructor_Data_Ord_Generic_GenericOrd[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[950481285] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ord_Generic_GenericOrd[any])(ptr)
		_ = c
		switch key {
		case "genericCompare'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Ord_Generic_GenericOrd: " + key)
		}
	}
}


func Call_Data_Ord_Generic_GenericOrd_dollar_Dict(x_0_loop struct{
	genericCompare_prime_ gopurs_runtime.Value
}) *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value] {
var x_0 struct{
	genericCompare_prime_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("genericCompare'", orig.genericCompare_prime_)
				}())
}

func Call_Data_Ord_Generic_genericOrdArgument(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2).IntVal)), UnsafePtr: nil}
})}))}
}

func Call_Data_Ord_Generic_genericCompare_prime_(dict_0_loop *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Ord_Generic_genericOrdConstructor(dictGenericOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), v_1, v1_2).IntVal)), UnsafePtr: nil}
})}))}
}

func Call_Data_Ord_Generic_genericOrdProduct(dictGenericOrd_0_loop gopurs_runtime.Value, dictGenericOrd1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
var dictGenericOrd1_1 gopurs_runtime.Value = dictGenericOrd1_1_loop
_ = dictGenericOrd1_1
return gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer(Rebox_Data_Ord_Generic_2274670046_3050046641((&Constructor_Data_Ord_Generic_GenericOrd[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v2_4_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_4_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)
_ = v2_4_0
var __t1 uint32
{
if (v2_4_0 == 902936544) {
__t1 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd1_1, "genericCompare'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1).IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = v2_4_0
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
})})))}
}

func Call_Data_Ord_Generic_genericOrdSum(dictGenericOrd_0_loop gopurs_runtime.Value, dictGenericOrd1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericOrd_0 gopurs_runtime.Value = dictGenericOrd_0_loop
_ = dictGenericOrd_0
var dictGenericOrd1_1 gopurs_runtime.Value = dictGenericOrd1_1_loop
_ = dictGenericOrd1_1
return gopurs_runtime.Value{Type: 9, IntVal: 950481285, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 uint32
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
var __t0 uint32
{
if (v1_3.Type == 9 && v1_3.IntVal == 3478632216) {
__t0 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 492034566) {
__t0 = 1527465420
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
__t2 = __t0
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 492034566) {
var __t1 uint32
{
if (v1_3.Type == 9 && v1_3.IntVal == 492034566) {
__t1 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd1_1, "genericCompare'"), (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 3478632216) {
__t1 = 380165415
goto end_branch_1
} else {

}
}
{
__t1 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t2), UnsafePtr: nil}
})}))}
}

func Call_Data_Ord_Generic_genericCompare(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericOrd_1_loop *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericOrd_1 *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value] = dictGenericOrd_1_loop
_ = dictGenericOrd_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictGenericOrd_1.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)).IntVal)
}

func Rebox_Data_Ord_Generic_2274670046_3050046641(in *Constructor_Data_Ord_Generic_GenericOrd[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Ord_Generic_2756679068_3050046641(in *Constructor_Data_Ord_Generic_GenericOrd[uint32]) *Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Generic_GenericOrd[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}


