package Data_Enum_Generic

import (
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genericToEnum_prime gopurs_runtime.Value
var once_genericToEnum_prime sync.Once
func Get_genericToEnum_prime() gopurs_runtime.Value {
	once_genericToEnum_prime.Do(func() {
		cache_genericToEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericToEnum_prime
}

var cache_genericToEnum_prime__gopurs_runtime_Value_244040409 gopurs_runtime.Value
var once_genericToEnum_prime__gopurs_runtime_Value_244040409 sync.Once
func Get_genericToEnum_prime__gopurs_runtime_Value_244040409() gopurs_runtime.Value {
	once_genericToEnum_prime__gopurs_runtime_Value_244040409.Do(func() {
		cache_genericToEnum_prime__gopurs_runtime_Value_244040409 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum_prime__gopurs_runtime_Value_244040409(gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericToEnum_prime__gopurs_runtime_Value_244040409
}

var cache_genericToEnum gopurs_runtime.Value
var once_genericToEnum sync.Once
func Get_genericToEnum() gopurs_runtime.Value {
	once_genericToEnum.Do(func() {
		cache_genericToEnum = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
})
	})
	return cache_genericToEnum
}

var cache_genericSucc_prime gopurs_runtime.Value
var once_genericSucc_prime sync.Once
func Get_genericSucc_prime() gopurs_runtime.Value {
	once_genericSucc_prime.Do(func() {
		cache_genericSucc_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericSucc_prime
}

var cache_genericSucc_prime__gopurs_runtime_Value_3707548623 gopurs_runtime.Value
var once_genericSucc_prime__gopurs_runtime_Value_3707548623 sync.Once
func Get_genericSucc_prime__gopurs_runtime_Value_3707548623() gopurs_runtime.Value {
	once_genericSucc_prime__gopurs_runtime_Value_3707548623.Do(func() {
		cache_genericSucc_prime__gopurs_runtime_Value_3707548623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc_prime__gopurs_runtime_Value_3707548623(gopurs_runtime.CoerceToStruct[Constructor_GenericEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericSucc_prime__gopurs_runtime_Value_3707548623
}

var cache_genericSucc gopurs_runtime.Value
var once_genericSucc sync.Once
func Get_genericSucc() gopurs_runtime.Value {
	once_genericSucc.Do(func() {
		cache_genericSucc = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
})
	})
	return cache_genericSucc
}

var cache_genericPred_prime gopurs_runtime.Value
var once_genericPred_prime sync.Once
func Get_genericPred_prime() gopurs_runtime.Value {
	once_genericPred_prime.Do(func() {
		cache_genericPred_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericPred_prime
}

var cache_genericPred_prime__gopurs_runtime_Value_3707548623 gopurs_runtime.Value
var once_genericPred_prime__gopurs_runtime_Value_3707548623 sync.Once
func Get_genericPred_prime__gopurs_runtime_Value_3707548623() gopurs_runtime.Value {
	once_genericPred_prime__gopurs_runtime_Value_3707548623.Do(func() {
		cache_genericPred_prime__gopurs_runtime_Value_3707548623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred_prime__gopurs_runtime_Value_3707548623(gopurs_runtime.CoerceToStruct[Constructor_GenericEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericPred_prime__gopurs_runtime_Value_3707548623
}

var cache_genericPred gopurs_runtime.Value
var once_genericPred sync.Once
func Get_genericPred() gopurs_runtime.Value {
	once_genericPred.Do(func() {
		cache_genericPred = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
})
	})
	return cache_genericPred
}

var cache_genericFromEnum_prime gopurs_runtime.Value
var once_genericFromEnum_prime sync.Once
func Get_genericFromEnum_prime() gopurs_runtime.Value {
	once_genericFromEnum_prime.Do(func() {
		cache_genericFromEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericFromEnum_prime
}

var cache_genericFromEnum_prime__gopurs_runtime_Value_4119640152 gopurs_runtime.Value
var once_genericFromEnum_prime__gopurs_runtime_Value_4119640152 sync.Once
func Get_genericFromEnum_prime__gopurs_runtime_Value_4119640152() gopurs_runtime.Value {
	once_genericFromEnum_prime__gopurs_runtime_Value_4119640152.Do(func() {
		cache_genericFromEnum_prime__gopurs_runtime_Value_4119640152 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum_prime__gopurs_runtime_Value_4119640152(gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericFromEnum_prime__gopurs_runtime_Value_4119640152
}

var cache_genericFromEnum gopurs_runtime.Value
var once_genericFromEnum sync.Once
func Get_genericFromEnum() gopurs_runtime.Value {
	once_genericFromEnum.Do(func() {
		cache_genericFromEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_genericFromEnum(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dictGenericBoundedEnum_1_box), x_2_box))
})
	})
	return cache_genericFromEnum
}

var cache_genericEnumSum gopurs_runtime.Value
var once_genericEnumSum sync.Once
func Get_genericEnumSum() gopurs_runtime.Value {
	once_genericEnumSum.Do(func() {
		cache_genericEnumSum = gopurs_runtime.Func4(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value, dictGenericEnum1_2_box gopurs_runtime.Value, dictGenericBottom_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumSum(dictGenericEnum_0_box, dictGenericTop_1_box, dictGenericEnum1_2_box, dictGenericBottom_3_box)
})
	})
	return cache_genericEnumSum
}

var cache_genericEnumProduct gopurs_runtime.Value
var once_genericEnumProduct sync.Once
func Get_genericEnumProduct() gopurs_runtime.Value {
	once_genericEnumProduct.Do(func() {
		cache_genericEnumProduct = gopurs_runtime.Func6(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value, dictGenericBottom_2_box gopurs_runtime.Value, dictGenericEnum1_3_box gopurs_runtime.Value, dictGenericTop1_4_box gopurs_runtime.Value, dictGenericBottom1_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumProduct(dictGenericEnum_0_box, dictGenericTop_1_box, dictGenericBottom_2_box, dictGenericEnum1_3_box, dictGenericTop1_4_box, dictGenericBottom1_5_box)
})
	})
	return cache_genericEnumProduct
}

var cache_genericEnumNoArguments gopurs_runtime.Value
var once_genericEnumNoArguments sync.Once
func Get_genericEnumNoArguments() gopurs_runtime.Value {
	once_genericEnumNoArguments.Do(func() {
		cache_genericEnumNoArguments = gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}))
	})
	return cache_genericEnumNoArguments
}

var cache_genericEnumConstructor gopurs_runtime.Value
var once_genericEnumConstructor sync.Once
func Get_genericEnumConstructor() gopurs_runtime.Value {
	once_genericEnumConstructor.Do(func() {
		cache_genericEnumConstructor = gopurs_runtime.Func(func(dictGenericEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumConstructor(dictGenericEnum_0_box)
})
	})
	return cache_genericEnumConstructor
}

var cache_genericEnumArgument gopurs_runtime.Value
var once_genericEnumArgument sync.Once
func Get_genericEnumArgument() gopurs_runtime.Value {
	once_genericEnumArgument.Do(func() {
		cache_genericEnumArgument = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumArgument(dictEnum_0_box)
})
	})
	return cache_genericEnumArgument
}

var cache_genericCardinality_prime gopurs_runtime.Value
var once_genericCardinality_prime sync.Once
func Get_genericCardinality_prime() gopurs_runtime.Value {
	once_genericCardinality_prime.Do(func() {
		cache_genericCardinality_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCardinality_prime(dict_0_box)
})
	})
	return cache_genericCardinality_prime
}

var cache_genericCardinality gopurs_runtime.Value
var once_genericCardinality sync.Once
func Get_genericCardinality() gopurs_runtime.Value {
	once_genericCardinality.Do(func() {
		cache_genericCardinality = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericCardinality(dictGeneric_0_box, dictGenericBoundedEnum_1_box)
})
	})
	return cache_genericCardinality
}

var cache_genericBoundedEnumSum gopurs_runtime.Value
var once_genericBoundedEnumSum sync.Once
func Get_genericBoundedEnumSum() gopurs_runtime.Value {
	once_genericBoundedEnumSum.Do(func() {
		cache_genericBoundedEnumSum = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumSum(dictGenericBoundedEnum_0_box)
})
	})
	return cache_genericBoundedEnumSum
}

var cache_genericBoundedEnumProduct gopurs_runtime.Value
var once_genericBoundedEnumProduct sync.Once
func Get_genericBoundedEnumProduct() gopurs_runtime.Value {
	once_genericBoundedEnumProduct.Do(func() {
		cache_genericBoundedEnumProduct = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumProduct(dictGenericBoundedEnum_0_box)
})
	})
	return cache_genericBoundedEnumProduct
}

var cache_genericBoundedEnumNoArguments gopurs_runtime.Value
var once_genericBoundedEnumNoArguments sync.Once
func Get_genericBoundedEnumNoArguments() gopurs_runtime.Value {
	once_genericBoundedEnumNoArguments.Do(func() {
		cache_genericBoundedEnumNoArguments = gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(1), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: unsafe.Pointer(nil)}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
}))
	})
	return cache_genericBoundedEnumNoArguments
}

var cache_genericBoundedEnumConstructor gopurs_runtime.Value
var once_genericBoundedEnumConstructor sync.Once
func Get_genericBoundedEnumConstructor() gopurs_runtime.Value {
	once_genericBoundedEnumConstructor.Do(func() {
		cache_genericBoundedEnumConstructor = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_box)
})
	})
	return cache_genericBoundedEnumConstructor
}

var cache_genericBoundedEnumArgument gopurs_runtime.Value
var once_genericBoundedEnumArgument sync.Once
func Get_genericBoundedEnumArgument() gopurs_runtime.Value {
	once_genericBoundedEnumArgument.Do(func() {
		cache_genericBoundedEnumArgument = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBoundedEnumArgument(dictBoundedEnum_0_box)
})
	})
	return cache_genericBoundedEnumArgument
}

type Constructor_GenericEnum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3087587621] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericEnum[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericPred'": return c.V0
		case "genericSucc'": return c.V1
		default: panic("Key not found in dictionary Constructor_GenericEnum: " + key)
		}
	}
}


type Constructor_GenericBoundedEnum[T_a any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4011582198] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericBoundedEnum[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericCardinality'": return c.V0
		case "genericFromEnum'": return c.V1
		case "genericToEnum'": return c.V2
		default: panic("Key not found in dictionary Constructor_GenericBoundedEnum: " + key)
		}
	}
}


func Call_genericToEnum_prime(dict_0_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_genericToEnum_prime__gopurs_runtime_Value_244040409(dict_0_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_genericToEnum(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := dictGeneric_0.V1
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_2, "genericToEnum'"), x_4))
})
})
}

func Call_genericSucc_prime(dict_0_loop *Constructor_GenericEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericSucc_prime__gopurs_runtime_Value_3707548623(dict_0_loop *Constructor_GenericEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericSucc(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := dictGeneric_0.V1
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericSucc'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_4)))
})
})
}

func Call_genericPred_prime(dict_0_loop *Constructor_GenericEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericPred_prime__gopurs_runtime_Value_3707548623(dict_0_loop *Constructor_GenericEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericPred(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := dictGeneric_0.V1
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericPred'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_4)))
})
})
}

func Call_genericFromEnum_prime(dict_0_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericFromEnum_prime__gopurs_runtime_Value_4119640152(dict_0_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericFromEnum(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericBoundedEnum_1_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) int64 {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictGenericBoundedEnum_1.V1, gopurs_runtime.Apply(dictGeneric_0.V0, x_2)).IntVal
}

func Call_genericEnumSum(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value, dictGenericEnum1_2_loop gopurs_runtime.Value, dictGenericBottom_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
var dictGenericEnum1_2 gopurs_runtime.Value = dictGenericEnum1_2_loop
_ = dictGenericEnum1_2
var dictGenericBottom_3 gopurs_runtime.Value = dictGenericBottom_3_loop
_ = dictGenericBottom_3
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))))}
goto end_branch_0
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
v1_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = v1_5_1
var __t2 gopurs_runtime.Value
{
if (v1_5_1.Type == 9 && v1_5_1.IntVal == 930809136 && v1_5_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'")})}})}))}
goto end_branch_2
} else {

}
}
{
if (v1_5_1.Type == 9 && v1_5_1.IntVal == 930809136 && v1_5_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_5_1.UnsafePtr).V0})}})}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
v1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = v1_5_4
var __t5 gopurs_runtime.Value
{
if (v1_5_4.Type == 9 && v1_5_4.IntVal == 930809136 && v1_5_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericBottom_3, "genericBottom'")})}})}))}
goto end_branch_5
} else {

}
}
{
if (v1_5_4.Type == 9 && v1_5_4.IntVal == 930809136 && v1_5_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_5_4.UnsafePtr).V0})}})}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
goto end_branch_3
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}))
}

func Call_genericEnumProduct(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value, dictGenericBottom_2_loop gopurs_runtime.Value, dictGenericEnum1_3_loop gopurs_runtime.Value, dictGenericTop1_4_loop gopurs_runtime.Value, dictGenericBottom1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
var dictGenericBottom_2 gopurs_runtime.Value = dictGenericBottom_2_loop
_ = dictGenericBottom_2
var dictGenericEnum1_3 gopurs_runtime.Value = dictGenericEnum1_3_loop
_ = dictGenericEnum1_3
var dictGenericTop1_4 gopurs_runtime.Value = dictGenericTop1_4_loop
_ = dictGenericTop1_4
var dictGenericBottom1_5 gopurs_runtime.Value = dictGenericBottom1_5_loop
_ = dictGenericBottom1_5
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
v1_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
_ = v1_7_0
var __t1 gopurs_runtime.Value
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_0.UnsafePtr).V0})}})}))}
goto end_branch_1
} else {

}
}
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr == nil) {
__local_var_8_2 := gopurs_runtime.RecordGet(dictGenericTop1_4, "genericTop'")
_ = __local_var_8_2
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, __local_var_8_2})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](__t1))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
v1_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
_ = v1_7_3
var __t4 gopurs_runtime.Value
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_3.UnsafePtr).V0})}})}))}
goto end_branch_4
} else {

}
}
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr == nil) {
__local_var_8_5 := gopurs_runtime.RecordGet(dictGenericBottom1_5, "genericBottom'")
_ = __local_var_8_5
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, __local_var_8_5})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](__t4))}
}))
}

func Call_genericEnumConstructor(dictGenericEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), v_1))))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), v_1))))}
}))
}

func Call_genericEnumArgument(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), v_1))))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), v_1))))}
}))
}

func Call_genericCardinality_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericCardinality'")
}

func Call_genericCardinality(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
return gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_1, "genericCardinality'").IntVal)
}

func Call_genericBoundedEnumSum(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
genericCardinality_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'")
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int((genericCardinality_prime1_1_0.IntVal) + (gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'").IntVal)), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3478632216) {
__t1 = gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 492034566) {
__t1 = gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) + (genericCardinality_prime1_1_0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Int(__t1.IntVal)
}), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t3 gopurs_runtime.Value
{
if (n_3.IntVal) < (0) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(true)
}
end_branch_3:
var __t4 gopurs_runtime.Value
{
if (n_3.IntVal) < (genericCardinality_prime1_1_0.IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t3, __t4).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), n_3))))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_3.IntVal) - (genericCardinality_prime1_1_0.IntVal))))))}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
}))
})
}

func Call_genericBoundedEnumProduct(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
genericCardinality_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'")
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
genericCardinality_prime2_3_1 := gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'")
_ = genericCardinality_prime2_3_1
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int((genericCardinality_prime1_1_0.IntVal) * (genericCardinality_prime2_3_1.IntVal)), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) * (genericCardinality_prime2_3_1.IntVal)) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1).IntVal))
}), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Product(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "div"), n_4, genericCardinality_prime2_3_1))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), n_4, genericCardinality_prime2_3_1)))))}
}))
})
}

func Call_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'").IntVal), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), i_1))))}
}))
}

func Call_genericBoundedEnumArgument(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.RecordGet(dictBoundedEnum_0, "cardinality").IntVal), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), i_1))))}
}))
}


