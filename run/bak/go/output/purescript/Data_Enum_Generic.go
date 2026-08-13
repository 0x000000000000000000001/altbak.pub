package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_Generic_GenericEnum_dollarDict gopurs_runtime.Value
var once_Data_Enum_Generic_GenericEnum_dollarDict sync.Once
func Get_Data_Enum_Generic_GenericEnum_dollarDict() gopurs_runtime.Value {
	once_Data_Enum_Generic_GenericEnum_dollarDict.Do(func() {
		cache_Data_Enum_Generic_GenericEnum_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_GenericEnum_dollarDict(x_0_box)
})
	})
	return cache_Data_Enum_Generic_GenericEnum_dollarDict
}

var cache_Data_Enum_Generic_GenericBoundedEnum_dollarDict gopurs_runtime.Value
var once_Data_Enum_Generic_GenericBoundedEnum_dollarDict sync.Once
func Get_Data_Enum_Generic_GenericBoundedEnum_dollarDict() gopurs_runtime.Value {
	once_Data_Enum_Generic_GenericBoundedEnum_dollarDict.Do(func() {
		cache_Data_Enum_Generic_GenericBoundedEnum_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_GenericBoundedEnum_dollarDict(x_0_box)
})
	})
	return cache_Data_Enum_Generic_GenericBoundedEnum_dollarDict
}

var cache_Data_Enum_Generic_genericToEnum_prime gopurs_runtime.Value
var once_Data_Enum_Generic_genericToEnum_prime sync.Once
func Get_Data_Enum_Generic_genericToEnum_prime() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericToEnum_prime.Do(func() {
		cache_Data_Enum_Generic_genericToEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericToEnum_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericToEnum_prime
}

var cache_Data_Enum_Generic_genericToEnum gopurs_runtime.Value
var once_Data_Enum_Generic_genericToEnum sync.Once
func Get_Data_Enum_Generic_genericToEnum() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericToEnum.Do(func() {
		cache_Data_Enum_Generic_genericToEnum = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box))
})
	})
	return cache_Data_Enum_Generic_genericToEnum
}

var cache_Data_Enum_Generic_genericSucc_prime gopurs_runtime.Value
var once_Data_Enum_Generic_genericSucc_prime sync.Once
func Get_Data_Enum_Generic_genericSucc_prime() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericSucc_prime.Do(func() {
		cache_Data_Enum_Generic_genericSucc_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericSucc_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericSucc_prime
}

var cache_Data_Enum_Generic_genericSucc gopurs_runtime.Value
var once_Data_Enum_Generic_genericSucc sync.Once
func Get_Data_Enum_Generic_genericSucc() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericSucc.Do(func() {
		cache_Data_Enum_Generic_genericSucc = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericSucc(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box))
})
	})
	return cache_Data_Enum_Generic_genericSucc
}

var cache_Data_Enum_Generic_genericPred_prime gopurs_runtime.Value
var once_Data_Enum_Generic_genericPred_prime sync.Once
func Get_Data_Enum_Generic_genericPred_prime() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericPred_prime.Do(func() {
		cache_Data_Enum_Generic_genericPred_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericPred_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericPred_prime
}

var cache_Data_Enum_Generic_genericPred gopurs_runtime.Value
var once_Data_Enum_Generic_genericPred sync.Once
func Get_Data_Enum_Generic_genericPred() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericPred.Do(func() {
		cache_Data_Enum_Generic_genericPred = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericPred(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box))
})
	})
	return cache_Data_Enum_Generic_genericPred
}

var cache_Data_Enum_Generic_genericFromEnum_prime gopurs_runtime.Value
var once_Data_Enum_Generic_genericFromEnum_prime sync.Once
func Get_Data_Enum_Generic_genericFromEnum_prime() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericFromEnum_prime.Do(func() {
		cache_Data_Enum_Generic_genericFromEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericFromEnum_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericFromEnum_prime
}

var cache_Data_Enum_Generic_genericFromEnum gopurs_runtime.Value
var once_Data_Enum_Generic_genericFromEnum sync.Once
func Get_Data_Enum_Generic_genericFromEnum() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericFromEnum.Do(func() {
		cache_Data_Enum_Generic_genericFromEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_Generic_genericFromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum](dictGenericBoundedEnum_1_box), x_2_box))
})
	})
	return cache_Data_Enum_Generic_genericFromEnum
}

var cache_Data_Enum_Generic_genericEnumSum gopurs_runtime.Value
var once_Data_Enum_Generic_genericEnumSum sync.Once
func Get_Data_Enum_Generic_genericEnumSum() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericEnumSum.Do(func() {
		cache_Data_Enum_Generic_genericEnumSum = gopurs_runtime.Func4(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value, dictGenericEnum1_2_box gopurs_runtime.Value, dictGenericBottom_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericEnumSum(dictGenericEnum_0_box, dictGenericTop_1_box, dictGenericEnum1_2_box, dictGenericBottom_3_box)
})
	})
	return cache_Data_Enum_Generic_genericEnumSum
}

var cache_Data_Enum_Generic_genericEnumProduct gopurs_runtime.Value
var once_Data_Enum_Generic_genericEnumProduct sync.Once
func Get_Data_Enum_Generic_genericEnumProduct() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericEnumProduct.Do(func() {
		cache_Data_Enum_Generic_genericEnumProduct = gopurs_runtime.Func6(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value, dictGenericBottom_2_box gopurs_runtime.Value, dictGenericEnum1_3_box gopurs_runtime.Value, dictGenericTop1_4_box gopurs_runtime.Value, dictGenericBottom1_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericEnumProduct(dictGenericEnum_0_box, dictGenericTop_1_box, dictGenericBottom_2_box, dictGenericEnum1_3_box, dictGenericTop1_4_box, dictGenericBottom1_5_box)
})
	})
	return cache_Data_Enum_Generic_genericEnumProduct
}

var cache_Data_Enum_Generic_genericEnumNoArguments gopurs_runtime.Value
var once_Data_Enum_Generic_genericEnumNoArguments sync.Once
func Get_Data_Enum_Generic_genericEnumNoArguments() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericEnumNoArguments.Do(func() {
		cache_Data_Enum_Generic_genericEnumNoArguments = gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Enum_Generic_genericEnumNoArguments
}

var cache_Data_Enum_Generic_genericEnumConstructor gopurs_runtime.Value
var once_Data_Enum_Generic_genericEnumConstructor sync.Once
func Get_Data_Enum_Generic_genericEnumConstructor() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericEnumConstructor.Do(func() {
		cache_Data_Enum_Generic_genericEnumConstructor = gopurs_runtime.Func(func(dictGenericEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericEnumConstructor(dictGenericEnum_0_box)
})
	})
	return cache_Data_Enum_Generic_genericEnumConstructor
}

var cache_Data_Enum_Generic_genericEnumArgument gopurs_runtime.Value
var once_Data_Enum_Generic_genericEnumArgument sync.Once
func Get_Data_Enum_Generic_genericEnumArgument() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericEnumArgument.Do(func() {
		cache_Data_Enum_Generic_genericEnumArgument = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericEnumArgument(dictEnum_0_box)
})
	})
	return cache_Data_Enum_Generic_genericEnumArgument
}

var cache_Data_Enum_Generic_genericCardinality_prime gopurs_runtime.Value
var once_Data_Enum_Generic_genericCardinality_prime sync.Once
func Get_Data_Enum_Generic_genericCardinality_prime() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericCardinality_prime.Do(func() {
		cache_Data_Enum_Generic_genericCardinality_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericCardinality_prime(dict_0_box)
})
	})
	return cache_Data_Enum_Generic_genericCardinality_prime
}

var cache_Data_Enum_Generic_genericCardinality gopurs_runtime.Value
var once_Data_Enum_Generic_genericCardinality sync.Once
func Get_Data_Enum_Generic_genericCardinality() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericCardinality.Do(func() {
		cache_Data_Enum_Generic_genericCardinality = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericCardinality(dictGeneric_0_box, dictGenericBoundedEnum_1_box)
})
	})
	return cache_Data_Enum_Generic_genericCardinality
}

var cache_Data_Enum_Generic_genericBoundedEnumSum gopurs_runtime.Value
var once_Data_Enum_Generic_genericBoundedEnumSum sync.Once
func Get_Data_Enum_Generic_genericBoundedEnumSum() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericBoundedEnumSum.Do(func() {
		cache_Data_Enum_Generic_genericBoundedEnumSum = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericBoundedEnumSum(dictGenericBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_Generic_genericBoundedEnumSum
}

var cache_Data_Enum_Generic_genericBoundedEnumProduct gopurs_runtime.Value
var once_Data_Enum_Generic_genericBoundedEnumProduct sync.Once
func Get_Data_Enum_Generic_genericBoundedEnumProduct() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericBoundedEnumProduct.Do(func() {
		cache_Data_Enum_Generic_genericBoundedEnumProduct = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericBoundedEnumProduct(dictGenericBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_Generic_genericBoundedEnumProduct
}

var cache_Data_Enum_Generic_genericBoundedEnumNoArguments gopurs_runtime.Value
var once_Data_Enum_Generic_genericBoundedEnumNoArguments sync.Once
func Get_Data_Enum_Generic_genericBoundedEnumNoArguments() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericBoundedEnumNoArguments.Do(func() {
		cache_Data_Enum_Generic_genericBoundedEnumNoArguments = gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(1), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (i_0.IntVal) == (0) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}))
	})
	return cache_Data_Enum_Generic_genericBoundedEnumNoArguments
}

var cache_Data_Enum_Generic_genericBoundedEnumConstructor gopurs_runtime.Value
var once_Data_Enum_Generic_genericBoundedEnumConstructor sync.Once
func Get_Data_Enum_Generic_genericBoundedEnumConstructor() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericBoundedEnumConstructor.Do(func() {
		cache_Data_Enum_Generic_genericBoundedEnumConstructor = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_Generic_genericBoundedEnumConstructor
}

var cache_Data_Enum_Generic_genericBoundedEnumArgument gopurs_runtime.Value
var once_Data_Enum_Generic_genericBoundedEnumArgument sync.Once
func Get_Data_Enum_Generic_genericBoundedEnumArgument() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericBoundedEnumArgument.Do(func() {
		cache_Data_Enum_Generic_genericBoundedEnumArgument = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericBoundedEnumArgument(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_Generic_genericBoundedEnumArgument
}

var cache_Data_Enum_Generic_genericFromEnum_prime__4119640152 gopurs_runtime.Value
var once_Data_Enum_Generic_genericFromEnum_prime__4119640152 sync.Once
func Get_Data_Enum_Generic_genericFromEnum_prime__4119640152() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericFromEnum_prime__4119640152.Do(func() {
		cache_Data_Enum_Generic_genericFromEnum_prime__4119640152 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericFromEnum_prime__4119640152(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericFromEnum_prime__4119640152
}

var cache_Data_Enum_Generic_genericPred_prime__3707548623 gopurs_runtime.Value
var once_Data_Enum_Generic_genericPred_prime__3707548623 sync.Once
func Get_Data_Enum_Generic_genericPred_prime__3707548623() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericPred_prime__3707548623.Do(func() {
		cache_Data_Enum_Generic_genericPred_prime__3707548623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericPred_prime__3707548623(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericPred_prime__3707548623
}

var cache_Data_Enum_Generic_genericSucc_prime__3707548623 gopurs_runtime.Value
var once_Data_Enum_Generic_genericSucc_prime__3707548623 sync.Once
func Get_Data_Enum_Generic_genericSucc_prime__3707548623() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericSucc_prime__3707548623.Do(func() {
		cache_Data_Enum_Generic_genericSucc_prime__3707548623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericSucc_prime__3707548623(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericSucc_prime__3707548623
}

var cache_Data_Enum_Generic_genericToEnum_prime__244040409 gopurs_runtime.Value
var once_Data_Enum_Generic_genericToEnum_prime__244040409 sync.Once
func Get_Data_Enum_Generic_genericToEnum_prime__244040409() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericToEnum_prime__244040409.Do(func() {
		cache_Data_Enum_Generic_genericToEnum_prime__244040409 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericToEnum_prime__244040409(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericToEnum_prime__244040409
}

type Constructor_Data_Enum_Generic_GenericEnum struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3087587621] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Generic_GenericEnum)(ptr)
		_ = c
		switch key {
		case "genericPred'": return gopurs_runtime.Box(c.V0)
		case "genericSucc'": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Generic_GenericEnum: " + key)
		}
	}
}


type Constructor_Data_Enum_Generic_GenericBoundedEnum struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4011582198] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Generic_GenericBoundedEnum)(ptr)
		_ = c
		switch key {
		case "genericCardinality'": return gopurs_runtime.Box(c.V0)
		case "genericFromEnum'": return gopurs_runtime.Box(c.V1)
		case "genericToEnum'": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Generic_GenericBoundedEnum: " + key)
		}
	}
}


func Call_Data_Enum_Generic_GenericEnum_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Enum_Generic_GenericBoundedEnum_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Enum_Generic_genericToEnum_prime(dict_0_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericBoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Enum_Generic_genericToEnum(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
// TAST (Let): to_1_0 -> gopurs_runtime.Value
to_1_0 := gopurs_runtime.Box(dictGeneric_0.V1)
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_2, "genericToEnum'"), x_4))
})
})
}

func Call_Data_Enum_Generic_genericSucc_prime(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_Generic_genericSucc(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
// TAST (Let): to_1_0 -> gopurs_runtime.Value
to_1_0 := gopurs_runtime.Box(dictGeneric_0.V1)
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericSucc'"), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_4)))
})
})
}

func Call_Data_Enum_Generic_genericPred_prime(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Enum_Generic_genericPred(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
// TAST (Let): to_1_0 -> gopurs_runtime.Value
to_1_0 := gopurs_runtime.Box(dictGeneric_0.V1)
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericPred'"), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_4)))
})
})
}

func Call_Data_Enum_Generic_genericFromEnum_prime(dict_0_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericBoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_Generic_genericFromEnum(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericBoundedEnum_1_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum, x_2_loop gopurs_runtime.Value) int64 {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Constructor_Data_Enum_Generic_GenericBoundedEnum = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericBoundedEnum_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2)).IntVal
}

func Call_Data_Enum_Generic_genericEnumSum(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value, dictGenericEnum1_2_loop gopurs_runtime.Value, dictGenericBottom_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
var dictGenericEnum1_2 gopurs_runtime.Value = dictGenericEnum1_2_loop
_ = dictGenericEnum1_2
var dictGenericBottom_3 gopurs_runtime.Value = dictGenericBottom_3_loop
_ = dictGenericBottom_3
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Inl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*Constructor_Data_Generic_Rep_Inl)(v_4.UnsafePtr).V0)))}))
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
// TAST (Let): v1_5_0 -> gopurs_runtime.Value
v1_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericPred'"), (*Constructor_Data_Generic_Rep_Inr)(v_4.UnsafePtr).V0)
_ = v1_5_0
var __t1 *Constructor_Data_Maybe_Just
{
if (v1_5_0.Type == 9 && v1_5_0.IntVal == 930809136 && v1_5_0.UnsafePtr == nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'")})}}
goto end_branch_1
} else {

}
}
{
if (v1_5_0.Type == 9 && v1_5_0.IntVal == 930809136 && v1_5_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, (*Constructor_Data_Maybe_Just)(v1_5_0.UnsafePtr).V0})}}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
// TAST (Let): v1_5_3 -> gopurs_runtime.Value
v1_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*Constructor_Data_Generic_Rep_Inl)(v_4.UnsafePtr).V0)
_ = v1_5_3
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_5_3.Type == 9 && v1_5_3.IntVal == 930809136 && v1_5_3.UnsafePtr == nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, gopurs_runtime.RecordGet(dictGenericBottom_3, "genericBottom'")})}}
goto end_branch_4
} else {

}
}
{
if (v1_5_3.Type == 9 && v1_5_3.IntVal == 930809136 && v1_5_3.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, (*Constructor_Data_Maybe_Just)(v1_5_3.UnsafePtr).V0})}}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Inr(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericSucc'"), (*Constructor_Data_Generic_Rep_Inr)(v_4.UnsafePtr).V0)))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}))
}

func Call_Data_Enum_Generic_genericEnumProduct(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value, dictGenericBottom_2_loop gopurs_runtime.Value, dictGenericEnum1_3_loop gopurs_runtime.Value, dictGenericTop1_4_loop gopurs_runtime.Value, dictGenericBottom1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): v1_7_0 -> gopurs_runtime.Value
v1_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V1)
_ = v1_7_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_7_0.UnsafePtr).V0})}}
goto end_branch_2
} else {

}
}
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr == nil) {
// TAST (Let): __local_var_8_1 -> gopurs_runtime.Value
__local_var_8_1 := gopurs_runtime.RecordGet(dictGenericTop1_4, "genericTop'")
_ = __local_var_8_1
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, a_9, __local_var_8_1})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0)))})))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_7_3 -> gopurs_runtime.Value
v1_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V1)
_ = v1_7_3
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_7_3.UnsafePtr).V0})}}
goto end_branch_5
} else {

}
}
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr == nil) {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := gopurs_runtime.RecordGet(dictGenericBottom1_5, "genericBottom'")
_ = __local_var_8_4
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, a_9, __local_var_8_4})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0)))})))})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}))
}

func Call_Data_Enum_Generic_genericEnumConstructor(dictGenericEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Constructor(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), v_1)))})))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Constructor(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), v_1)))})))}
}))
}

func Call_Data_Enum_Generic_genericEnumArgument(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Argument(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), v_1)))})))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Argument(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), v_1)))})))}
}))
}

func Call_Data_Enum_Generic_genericCardinality_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericCardinality'")
}

func Call_Data_Enum_Generic_genericCardinality(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
return gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_1, "genericCardinality'").IntVal)
}

func Call_Data_Enum_Generic_genericBoundedEnumSum(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
// TAST (Let): genericCardinality_prime1_1_0 -> gopurs_runtime.Value
genericCardinality_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'")
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int((genericCardinality_prime1_1_0.IntVal) + (gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'").IntVal)), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 int64
{
if (v_3.Type == 9 && v_3.IntVal == 3478632216) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Inl)(v_3.UnsafePtr).V0).IntVal
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 492034566) {
__t1 = (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Inr)(v_3.UnsafePtr).V0).IntVal) + (genericCardinality_prime1_1_0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
var __t2 bool
{
if (n_3.IntVal) < (0) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
var __t_and_4 bool = false
if __t2 {

var __t3 bool
{
if (n_3.IntVal) < (genericCardinality_prime1_1_0.IntVal) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t_and_4 = __t3
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Inl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int(n_3.IntVal))))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Inr(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_3.IntVal) - (genericCardinality_prime1_1_0.IntVal)))))}))
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}))
})
}

func Call_Data_Enum_Generic_genericBoundedEnumProduct(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
// TAST (Let): genericCardinality_prime1_1_0 -> gopurs_runtime.Value
genericCardinality_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'")
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): genericCardinality_prime2_3_1 -> gopurs_runtime.Value
genericCardinality_prime2_3_1 := gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'")
_ = genericCardinality_prime2_3_1
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int((genericCardinality_prime1_1_0.IntVal) * (genericCardinality_prime2_3_1.IntVal)), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Product)(v1_4.UnsafePtr).V0).IntVal) * (genericCardinality_prime2_3_1.IntVal)) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Product)(v1_4.UnsafePtr).V1).IntVal))
}), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Product(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int((n_4.IntVal) / (genericCardinality_prime2_3_1.IntVal)))))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(n_4.IntVal), gopurs_runtime.Int(genericCardinality_prime2_3_1.IntVal)).IntVal))))})))}
}))
})
}

func Call_Data_Enum_Generic_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'").IntVal), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Constructor(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int(i_1.IntVal))))})))}
}))
}

func Call_Data_Enum_Generic_genericBoundedEnumArgument(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.RecordGet(dictBoundedEnum_0, "cardinality").IntVal), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Generic_Rep_Argument(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), gopurs_runtime.Int(i_1.IntVal))))})))}
}))
}

func Call_Data_Enum_Generic_genericFromEnum_prime__4119640152(dict_0_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericBoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_Generic_genericPred_prime__3707548623(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Enum_Generic_genericSucc_prime__3707548623(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_Generic_genericToEnum_prime__244040409(dict_0_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericBoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}


