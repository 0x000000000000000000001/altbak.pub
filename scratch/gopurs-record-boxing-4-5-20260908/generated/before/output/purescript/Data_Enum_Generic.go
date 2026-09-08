package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_Generic_GenericEnum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_Generic_GenericEnum_dollar_Dict sync.Once
func Get_Data_Enum_Generic_GenericEnum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_Generic_GenericEnum_dollar_Dict.Do(func() {
		cache_Data_Enum_Generic_GenericEnum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3087587621, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Generic_GenericEnum_dollar_Dict(func() struct{
	genericPred_prime_ gopurs_runtime.Value
	genericSucc_prime_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericPred_prime_ gopurs_runtime.Value
	genericSucc_prime_ gopurs_runtime.Value
}{}
					clone.genericPred_prime_ = gopurs_runtime.RecordGet(orig, "genericPred'")
					clone.genericSucc_prime_ = gopurs_runtime.RecordGet(orig, "genericSucc'")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_Generic_GenericEnum_dollar_Dict
}

var cache_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict sync.Once
func Get_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict.Do(func() {
		cache_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4011582198, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict(func() struct{
	genericCardinality_prime_ gopurs_runtime.Value
	genericFromEnum_prime_ gopurs_runtime.Value
	genericToEnum_prime_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	genericCardinality_prime_ gopurs_runtime.Value
	genericFromEnum_prime_ gopurs_runtime.Value
	genericToEnum_prime_ gopurs_runtime.Value
}{}
					clone.genericCardinality_prime_ = gopurs_runtime.RecordGet(orig, "genericCardinality'")
					clone.genericFromEnum_prime_ = gopurs_runtime.RecordGet(orig, "genericFromEnum'")
					clone.genericToEnum_prime_ = gopurs_runtime.RecordGet(orig, "genericToEnum'")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict
}

var cache_Data_Enum_Generic_genericToEnum_prime_ gopurs_runtime.Value
var once_Data_Enum_Generic_genericToEnum_prime_ sync.Once
func Get_Data_Enum_Generic_genericToEnum_prime_() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericToEnum_prime_.Do(func() {
		cache_Data_Enum_Generic_genericToEnum_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_Generic_genericToEnum_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_Generic_genericToEnum_prime_
}

var cache_Data_Enum_Generic_genericToEnum gopurs_runtime.Value
var once_Data_Enum_Generic_genericToEnum sync.Once
func Get_Data_Enum_Generic_genericToEnum() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericToEnum.Do(func() {
		cache_Data_Enum_Generic_genericToEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_Generic_genericToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]](dictGenericBoundedEnum_1_box), x_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_Generic_genericToEnum
}

var cache_Data_Enum_Generic_genericSucc_prime_ gopurs_runtime.Value
var once_Data_Enum_Generic_genericSucc_prime_ sync.Once
func Get_Data_Enum_Generic_genericSucc_prime_() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericSucc_prime_.Do(func() {
		cache_Data_Enum_Generic_genericSucc_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_Generic_genericSucc_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_Generic_genericSucc_prime_
}

var cache_Data_Enum_Generic_genericSucc gopurs_runtime.Value
var once_Data_Enum_Generic_genericSucc sync.Once
func Get_Data_Enum_Generic_genericSucc() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericSucc.Do(func() {
		cache_Data_Enum_Generic_genericSucc = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_Generic_genericSucc(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]](dictGenericEnum_1_box), x_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_Generic_genericSucc
}

var cache_Data_Enum_Generic_genericPred_prime_ gopurs_runtime.Value
var once_Data_Enum_Generic_genericPred_prime_ sync.Once
func Get_Data_Enum_Generic_genericPred_prime_() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericPred_prime_.Do(func() {
		cache_Data_Enum_Generic_genericPred_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_Generic_genericPred_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_Generic_genericPred_prime_
}

var cache_Data_Enum_Generic_genericPred gopurs_runtime.Value
var once_Data_Enum_Generic_genericPred sync.Once
func Get_Data_Enum_Generic_genericPred() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericPred.Do(func() {
		cache_Data_Enum_Generic_genericPred = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_Generic_genericPred(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]](dictGenericEnum_1_box), x_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_Generic_genericPred
}

var cache_Data_Enum_Generic_genericFromEnum_prime_ gopurs_runtime.Value
var once_Data_Enum_Generic_genericFromEnum_prime_ sync.Once
func Get_Data_Enum_Generic_genericFromEnum_prime_() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericFromEnum_prime_.Do(func() {
		cache_Data_Enum_Generic_genericFromEnum_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericFromEnum_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Enum_Generic_genericFromEnum_prime_
}

var cache_Data_Enum_Generic_genericFromEnum gopurs_runtime.Value
var once_Data_Enum_Generic_genericFromEnum sync.Once
func Get_Data_Enum_Generic_genericFromEnum() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericFromEnum.Do(func() {
		cache_Data_Enum_Generic_genericFromEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_Generic_genericFromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]](dictGenericBoundedEnum_1_box), x_2_box))
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
		cache_Data_Enum_Generic_genericEnumNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 3087587621, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_4197715708_4181322513((&Constructor_Data_Enum_Generic_GenericEnum[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
})})))}
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

var cache_Data_Enum_Generic_genericCardinality_prime_ gopurs_runtime.Value
var once_Data_Enum_Generic_genericCardinality_prime_ sync.Once
func Get_Data_Enum_Generic_genericCardinality_prime_() gopurs_runtime.Value {
	once_Data_Enum_Generic_genericCardinality_prime_.Do(func() {
		cache_Data_Enum_Generic_genericCardinality_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Generic_genericCardinality_prime_(dict_0_box)
})
	})
	return cache_Data_Enum_Generic_genericCardinality_prime_
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
		cache_Data_Enum_Generic_genericBoundedEnumNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 4011582198, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_3675079727_3471223010((&Constructor_Data_Enum_Generic_GenericBoundedEnum[uint32]{1, gopurs_runtime.Int(int64(1)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(int64(0))
}), gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) == (int64(0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_622082505_3094389156(Rebox_Data_Enum_Generic_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
})})))}
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

type Constructor_Data_Enum_Generic_GenericEnum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3087587621] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Generic_GenericEnum[any])(ptr)
		_ = c
		switch key {
		case "genericPred'": return gopurs_runtime.Box(c.V0)
		case "genericSucc'": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Generic_GenericEnum: " + key)
		}
	}
}


type Constructor_Data_Enum_Generic_GenericBoundedEnum[T_a any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4011582198] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Generic_GenericBoundedEnum[any])(ptr)
		_ = c
		switch key {
		case "genericCardinality'": return gopurs_runtime.Box(c.V0)
		case "genericFromEnum'": return gopurs_runtime.Box(c.V1)
		case "genericToEnum'": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Generic_GenericBoundedEnum: " + key)
		}
	}
}


func Call_Data_Enum_Generic_GenericEnum_dollar_Dict(x_0_loop struct{
	genericPred_prime_ gopurs_runtime.Value
	genericSucc_prime_ gopurs_runtime.Value
}) *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] {
var x_0 struct{
	genericPred_prime_ gopurs_runtime.Value
	genericSucc_prime_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"genericPred'", "genericSucc'"}, []gopurs_runtime.Value{orig.genericPred_prime_, orig.genericSucc_prime_})
				}())
}

func Call_Data_Enum_Generic_GenericBoundedEnum_dollar_Dict(x_0_loop struct{
	genericCardinality_prime_ gopurs_runtime.Value
	genericFromEnum_prime_ gopurs_runtime.Value
	genericToEnum_prime_ gopurs_runtime.Value
}) *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] {
var x_0 struct{
	genericCardinality_prime_ gopurs_runtime.Value
	genericFromEnum_prime_ gopurs_runtime.Value
	genericToEnum_prime_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"genericCardinality'", "genericFromEnum'", "genericToEnum'"}, []gopurs_runtime.Value{orig.genericCardinality_prime_, orig.genericFromEnum_prime_, orig.genericToEnum_prime_})
				}())
}

func Call_Data_Enum_Generic_genericToEnum_prime_(dict_0_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V2)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_Generic_genericToEnum(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericBoundedEnum_1_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value], x_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 int64 = x_2_loop
_ = x_2
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericBoundedEnum_1.V2), gopurs_runtime.Int(x_2)))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), (__local_var_3_0).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_Generic_genericSucc_prime_(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V1)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_Generic_genericSucc(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericEnum_1_loop *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEnum_1 *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] = dictGenericEnum_1_loop
_ = dictGenericEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericEnum_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_Generic_genericPred_prime_(dict_0_loop *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V0)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_Generic_genericPred(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericEnum_1_loop *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEnum_1 *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] = dictGenericEnum_1_loop
_ = dictGenericEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictGenericEnum_1.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_Generic_genericFromEnum_prime_(dict_0_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_Generic_genericFromEnum(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericBoundedEnum_1_loop *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) int64 {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] = dictGenericBoundedEnum_1_loop
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
return gopurs_runtime.Value{Type: 9, IntVal: 3087587621, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
// TAST (Let): __local_var_5_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))
_ = __local_var_5_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_5_0).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
// TAST (Let): v1_5_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v1_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericPred'"), (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))
_ = v1_5_2
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_5_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'")}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
if (v1_5_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_5_2).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
// TAST (Let): v1_5_5 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))
_ = v1_5_5
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_5_5 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericBottom_3, "genericBottom'")}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
if (v1_5_5 != nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_5_5).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
__t9 = __t6
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericSucc'"), (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))
_ = __local_var_5_7
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_7 != nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_5_7).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
})}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3087587621, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_1304835134_4181322513((&Constructor_Data_Enum_Generic_GenericEnum[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_7_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v1_7_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
_ = v1_7_0
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_7_0 != nil) {
__t3 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (v1_7_0).V0}))}})
goto end_branch_3
} else {

}
}
{
if (v1_7_0 == nil) {
// TAST (Let): __local_var_8_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_8_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_8_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_8_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_8_1).V0, gopurs_runtime.RecordGet(dictGenericTop1_4, "genericTop'")}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = Rebox_Data_Enum_Generic_2230344971_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]] { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_7_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v1_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
_ = v1_7_4
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_7_4 != nil) {
__t7 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (v1_7_4).V0}))}})
goto end_branch_7
} else {

}
}
{
if (v1_7_4 == nil) {
// TAST (Let): __local_var_8_5 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_8_5
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_8_5 != nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_8_5).V0, gopurs_runtime.RecordGet(dictGenericBottom1_5, "genericBottom'")}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = Rebox_Data_Enum_Generic_2230344971_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]] { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
})})))}
}

func Call_Data_Enum_Generic_genericEnumConstructor(dictGenericEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.Value{Type: 9, IntVal: 3087587621, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), v_1))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_2_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), v_1))
_ = __local_var_2_2
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_2_2).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})}))}
}

func Call_Data_Enum_Generic_genericEnumArgument(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.Value{Type: 9, IntVal: 3087587621, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), v_1))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_2_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), v_1))
_ = __local_var_2_2
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_2_2).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})}))}
}

func Call_Data_Enum_Generic_genericCardinality_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.Int(gopurs_runtime.RecordGet(dict_0, "genericCardinality'").IntVal)
}

func Call_Data_Enum_Generic_genericCardinality(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
return gopurs_runtime.RecordGet(dictGenericBoundedEnum_1, "genericCardinality'")
}

func Call_Data_Enum_Generic_genericBoundedEnumSum(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
// TAST (Let): genericCardinality_prime_1_1_0 shape=Other bindingType=(TypeApp Int [(TypeVar a)])
genericCardinality_prime_1_1_0 := gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'").IntVal)
_ = genericCardinality_prime_1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4011582198, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]{1, gopurs_runtime.Int((genericCardinality_prime_1_1_0.IntVal) + (gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'").IntVal)).IntVal, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 int64
{
if (v_3.Type == 9 && v_3.IntVal == 3478632216) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 492034566) {
__t1 = (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) + (genericCardinality_prime_1_1_0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((n_3.IntVal) >= (int64(0))) && ((n_3.IntVal) < (genericCardinality_prime_1_1_0.IntVal)) {
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int(n_3.IntVal)))
_ = __local_var_4_4
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_4 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_4_4).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_3.IntVal) - (genericCardinality_prime_1_1_0.IntVal))))
_ = __local_var_4_2
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_4_2).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_3:
__t6 = __t3
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
})}))}
})
}

func Call_Data_Enum_Generic_genericBoundedEnumProduct(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
// TAST (Let): genericCardinality_prime_1_1_0 shape=Other bindingType=(TypeApp Int [(TypeVar a)])
genericCardinality_prime_1_1_0 := gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'").IntVal)
_ = genericCardinality_prime_1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): genericCardinality_prime_2_3_1 shape=Other bindingType=(TypeApp Int [(TypeVar b)])
genericCardinality_prime_2_3_1 := gopurs_runtime.Int(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'").IntVal)
_ = genericCardinality_prime_2_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 4011582198, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_Generic_61850029_3471223010((&Constructor_Data_Enum_Generic_GenericBoundedEnum[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Int((genericCardinality_prime_1_1_0.IntVal) * (genericCardinality_prime_2_3_1.IntVal)).IntVal, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) * (genericCardinality_prime_2_3_1.IntVal)) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1).IntVal))
}), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int((n_4.IntVal) / (genericCardinality_prime_2_3_1.IntVal))))
_ = __local_var_5_2
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_2 != nil) {
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_4.IntVal) % (genericCardinality_prime_2_3_1.IntVal))))
_ = __local_var_6_3
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_6_3 != nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_5_2).V0, (__local_var_6_3).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})})))}
})
}

func Call_Data_Enum_Generic_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.Value{Type: 9, IntVal: 4011582198, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'").IntVal, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Int(i_1.IntVal)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_2_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})}))}
}

func Call_Data_Enum_Generic_genericBoundedEnumArgument(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.Value{Type: 9, IntVal: 4011582198, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBoundedEnum_0, "cardinality").IntVal, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), v_1).IntVal)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), gopurs_runtime.Int(i_1.IntVal)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_2_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})}))}
}

func Rebox_Data_Enum_Generic_1304835134_4181322513(in *Constructor_Data_Enum_Generic_GenericEnum[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_Generic_2230344971_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_Generic_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Enum_Generic_3675079727_3471223010(in *Constructor_Data_Enum_Generic_GenericBoundedEnum[uint32]) *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_Generic_4197715708_4181322513(in *Constructor_Data_Enum_Generic_GenericEnum[uint32]) *Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Generic_GenericEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_Generic_61850029_3471223010(in *Constructor_Data_Enum_Generic_GenericBoundedEnum[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Generic_GenericBoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_Generic_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}


