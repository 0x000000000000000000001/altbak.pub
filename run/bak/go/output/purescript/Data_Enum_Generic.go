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
		cache_Data_Enum_Generic_genericToEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Generic_genericToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum](dictGenericBoundedEnum_1_box), x_2_box.IntVal))}
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
		cache_Data_Enum_Generic_genericSucc = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Generic_genericSucc(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum](dictGenericEnum_1_box), x_2_box))}
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
		cache_Data_Enum_Generic_genericPred = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Generic_genericPred(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum](dictGenericEnum_1_box), x_2_box))}
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

func Call_Data_Enum_Generic_genericToEnum(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericBoundedEnum_1_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum, x_2_loop int64) *Constructor_Data_Maybe_Just {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Constructor_Data_Enum_Generic_GenericBoundedEnum = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 int64 = x_2_loop
_ = x_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericBoundedEnum_1.V2), gopurs_runtime.Int(x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Enum_Generic_genericSucc_prime(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_Generic_genericSucc(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericEnum_1_loop *Constructor_Data_Enum_Generic_GenericEnum, x_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEnum_1 *Constructor_Data_Enum_Generic_GenericEnum = dictGenericEnum_1_loop
_ = dictGenericEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericEnum_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Enum_Generic_genericPred_prime(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Enum_Generic_genericPred(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericEnum_1_loop *Constructor_Data_Enum_Generic_GenericEnum, x_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEnum_1 *Constructor_Data_Enum_Generic_GenericEnum = dictGenericEnum_1_loop
_ = dictGenericEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericEnum_1.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
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
var __t4 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
// TAST (Let): __local_var_5_0 -> *Constructor_Data_Maybe_Just
__local_var_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*Constructor_Data_Generic_Rep_Inl)(v_4.UnsafePtr).V0))
_ = __local_var_5_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_5_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, (__local_var_5_0).V0})}}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
// TAST (Let): v1_5_2 -> gopurs_runtime.Value
v1_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericPred'"), (*Constructor_Data_Generic_Rep_Inr)(v_4.UnsafePtr).V0)
_ = v1_5_2
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 930809136 && v1_5_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'")})}}
goto end_branch_3
} else {

}
}
{
if (v1_5_2.Type == 9 && v1_5_2.IntVal == 930809136 && v1_5_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, (*Constructor_Data_Maybe_Just)(v1_5_2.UnsafePtr).V0})}}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
// TAST (Let): v1_5_5 -> gopurs_runtime.Value
v1_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*Constructor_Data_Generic_Rep_Inl)(v_4.UnsafePtr).V0)
_ = v1_5_5
var __t6 *Constructor_Data_Maybe_Just
{
if (v1_5_5.Type == 9 && v1_5_5.IntVal == 930809136 && v1_5_5.UnsafePtr == nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, gopurs_runtime.RecordGet(dictGenericBottom_3, "genericBottom'")})}}
goto end_branch_6
} else {

}
}
{
if (v1_5_5.Type == 9 && v1_5_5.IntVal == 930809136 && v1_5_5.UnsafePtr != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, (*Constructor_Data_Maybe_Just)(v1_5_5.UnsafePtr).V0})}}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
__t9 = __t6
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
// TAST (Let): __local_var_5_7 -> *Constructor_Data_Maybe_Just
__local_var_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericSucc'"), (*Constructor_Data_Generic_Rep_Inr)(v_4.UnsafePtr).V0))
_ = __local_var_5_7
var __t8 *Constructor_Data_Maybe_Just
{
if (__local_var_5_7 != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, (__local_var_5_7).V0})}}
goto end_branch_8
} else {

}
}
{
__t8 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
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
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_7_0.UnsafePtr).V0})}}
goto end_branch_5
} else {

}
}
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr == nil) {
// TAST (Let): __local_var_8_2 -> gopurs_runtime.Value
__local_var_8_2 := gopurs_runtime.RecordGet(dictGenericTop1_4, "genericTop'")
_ = __local_var_8_2
// TAST (Let): __local_var_8_1 -> gopurs_runtime.Value
__local_var_8_1 := gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, a_9, __local_var_8_2})}
})
_ = __local_var_8_1
// TAST (Let): __local_var_9_3 -> *Constructor_Data_Maybe_Just
__local_var_9_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0))
_ = __local_var_9_3
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_9_3 != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_8_1, (__local_var_9_3).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_7_6 -> gopurs_runtime.Value
v1_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V1)
_ = v1_7_6
var __t11 *Constructor_Data_Maybe_Just
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr != nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(v1_7_6.UnsafePtr).V0})}}
goto end_branch_11
} else {

}
}
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr == nil) {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.RecordGet(dictGenericBottom1_5, "genericBottom'")
_ = __local_var_8_8
// TAST (Let): __local_var_8_7 -> gopurs_runtime.Value
__local_var_8_7 := gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, a_9, __local_var_8_8})}
})
_ = __local_var_8_7
// TAST (Let): __local_var_9_9 -> *Constructor_Data_Maybe_Just
__local_var_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*Constructor_Data_Generic_Rep_Product)(v_6.UnsafePtr).V0))
_ = __local_var_9_9
var __t10 *Constructor_Data_Maybe_Just
{
if (__local_var_9_9 != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_8_7, (__local_var_9_9).V0)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}
}))
}

func Call_Data_Enum_Generic_genericEnumConstructor(dictGenericEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), v_1))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, (__local_var_2_0).V0}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> *Constructor_Data_Maybe_Just
__local_var_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), v_1))
_ = __local_var_2_2
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_2_2 != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, (__local_var_2_2).V0}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}))
}

func Call_Data_Enum_Generic_genericEnumArgument(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), v_1))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, (__local_var_2_0).V0}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> *Constructor_Data_Maybe_Just
__local_var_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), v_1))
_ = __local_var_2_2
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_2_2 != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, (__local_var_2_2).V0}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
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
var __t9 *Constructor_Data_Maybe_Just
{
var __t4 bool
{
if (n_3.IntVal) < (0) {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
var __t_and_6 bool = false
if __t4 {

var __t5 bool
{
if (n_3.IntVal) < (genericCardinality_prime1_1_0.IntVal) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t_and_6 = __t5
}
if __t_and_6 {
// TAST (Let): __local_var_4_7 -> *Constructor_Data_Maybe_Just
__local_var_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int(n_3.IntVal)))
_ = __local_var_4_7
var __t8 *Constructor_Data_Maybe_Just
{
if (__local_var_4_7 != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, (__local_var_4_7).V0})}}
goto end_branch_8
} else {

}
}
{
__t8 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
// TAST (Let): __local_var_4_2 -> *Constructor_Data_Maybe_Just
__local_var_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_3.IntVal) - (genericCardinality_prime1_1_0.IntVal))))
_ = __local_var_4_2
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_4_2 != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, (__local_var_4_2).V0})}}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
__t9 = __t3
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
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
// TAST (Let): __local_var_5_3 -> *Constructor_Data_Maybe_Just
__local_var_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int((n_4.IntVal) / (genericCardinality_prime2_3_1.IntVal))))
_ = __local_var_5_3
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_5_3 != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Generic_Rep_Product(), (__local_var_5_3).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
// TAST (Let): __local_var_5_2 -> *Constructor_Data_Maybe_Just
var __local_var_5_2 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
// TAST (Let): __local_var_6_5 -> *Constructor_Data_Maybe_Just
__local_var_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(n_4.IntVal), gopurs_runtime.Int(genericCardinality_prime2_3_1.IntVal)).IntVal)))
_ = __local_var_6_5
var __t7 *Constructor_Data_Maybe_Just
{
if (__local_var_5_2 != nil) {
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_6_5 != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_2).V0, (__local_var_6_5).V0)}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (__local_var_5_2 == nil) {
__t7 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
}))
})
}

func Call_Data_Enum_Generic_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'").IntVal), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int(i_1.IntVal)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, (__local_var_2_0).V0}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Call_Data_Enum_Generic_genericBoundedEnumArgument(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.RecordGet(dictBoundedEnum_0, "cardinality").IntVal), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), gopurs_runtime.Int(i_1.IntVal)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, (__local_var_2_0).V0}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
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


