package Data_Show_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
)

var cache_genericShowArgsNoArguments gopurs_runtime.Value
var once_genericShowArgsNoArguments sync.Once
func Get_genericShowArgsNoArguments() gopurs_runtime.Value {
	once_genericShowArgsNoArguments.Do(func() {
		cache_genericShowArgsNoArguments = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{})
}))))
	})
	return cache_genericShowArgsNoArguments
}

var cache_genericShowArgsArgument gopurs_runtime.Value
var once_genericShowArgsArgument sync.Once
func Get_genericShowArgsArgument() gopurs_runtime.Value {
	once_genericShowArgsArgument.Do(func() {
		cache_genericShowArgsArgument = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericShowArgsArgument(dictShow_0_box))
})
	})
	return cache_genericShowArgsArgument
}

var cache_genericShowArgs gopurs_runtime.Value
var once_genericShowArgs sync.Once
func Get_genericShowArgs() gopurs_runtime.Value {
	once_genericShowArgs.Do(func() {
		cache_genericShowArgs = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShowArgs(dict_0_box)
})
	})
	return cache_genericShowArgs
}

var cache_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322 gopurs_runtime.Value
var once_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322 sync.Once
func Get_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322() gopurs_runtime.Value {
	once_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322.Do(func() {
		cache_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322(dict_0_box)
})
	})
	return cache_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322
}

var cache_genericShowArgsProduct gopurs_runtime.Value
var once_genericShowArgsProduct sync.Once
func Get_genericShowArgsProduct() gopurs_runtime.Value {
	once_genericShowArgsProduct.Do(func() {
		cache_genericShowArgsProduct = gopurs_runtime.Func2(func(dictGenericShowArgs_0_box gopurs_runtime.Value, dictGenericShowArgs1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericShowArgsProduct(dictGenericShowArgs_0_box, dictGenericShowArgs1_1_box))
})
	})
	return cache_genericShowArgsProduct
}

var cache_genericShowConstructor gopurs_runtime.Value
var once_genericShowConstructor sync.Once
func Get_genericShowConstructor() gopurs_runtime.Value {
	once_genericShowConstructor.Do(func() {
		cache_genericShowConstructor = gopurs_runtime.Func2(func(dictGenericShowArgs_0_box gopurs_runtime.Value, dictIsSymbol_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericShowConstructor(dictGenericShowArgs_0_box, dictIsSymbol_1_box))
})
	})
	return cache_genericShowConstructor
}

var cache_genericShow_prime gopurs_runtime.Value
var once_genericShow_prime sync.Once
func Get_genericShow_prime() gopurs_runtime.Value {
	once_genericShow_prime.Do(func() {
		cache_genericShow_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShow_prime(dict_0_box)
})
	})
	return cache_genericShow_prime
}

var cache_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986 gopurs_runtime.Value
var once_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986 sync.Once
func Get_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986() gopurs_runtime.Value {
	once_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986.Do(func() {
		cache_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986(dict_0_box)
})
	})
	return cache_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986
}

var cache_genericShowNoConstructors gopurs_runtime.Value
var once_genericShowNoConstructors sync.Once
func Get_genericShowNoConstructors() gopurs_runtime.Value {
	once_genericShowNoConstructors.Do(func() {
		cache_genericShowNoConstructors = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_genericShowNoConstructors(), "genericShow'"), a_0)
}))))
	})
	return cache_genericShowNoConstructors
}

var cache_genericShowSum gopurs_runtime.Value
var once_genericShowSum sync.Once
func Get_genericShowSum() gopurs_runtime.Value {
	once_genericShowSum.Do(func() {
		cache_genericShowSum = gopurs_runtime.Func2(func(dictGenericShow_0_box gopurs_runtime.Value, dictGenericShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genericShowSum(dictGenericShow_0_box, dictGenericShow1_1_box))
})
	})
	return cache_genericShowSum
}

var cache_genericShow gopurs_runtime.Value
var once_genericShow sync.Once
func Get_genericShow() gopurs_runtime.Value {
	once_genericShow.Do(func() {
		cache_genericShow = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_genericShow(dictGeneric_0_box, dictGenericShow_1_box, gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_genericShow
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Intercalate(arg0.StrVal(), func() []string {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))
})
	})
	return cache_intercalate
}

func Call_genericShowArgsArgument(dictShow_0_loop gopurs_runtime.Value) interface{} {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1)})
})))
}

func Call_genericShowArgs(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericShowArgs")
}

func Call_genericShowArgs__func_gopurs_runtime_Value__interface____arrstring_2688816322(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericShowArgs")
}

func Call_genericShowArgsProduct(dictGenericShowArgs_0_loop gopurs_runtime.Value, dictGenericShowArgs1_1_loop gopurs_runtime.Value) interface{} {
var dictGenericShowArgs_0 gopurs_runtime.Value = dictGenericShowArgs_0_loop
_ = dictGenericShowArgs_0
var dictGenericShowArgs1_1 gopurs_runtime.Value = dictGenericShowArgs1_1_loop
_ = dictGenericShowArgs1_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs1_1, "genericShowArgs"), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)))
})))
}

func Call_genericShowConstructor(dictGenericShowArgs_0_loop gopurs_runtime.Value, dictIsSymbol_1_loop gopurs_runtime.Value) interface{} {
var dictGenericShowArgs_0 gopurs_runtime.Value = dictGenericShowArgs_0_loop
_ = dictGenericShowArgs_0
var dictIsSymbol_1 gopurs_runtime.Value = dictIsSymbol_1_loop
_ = dictIsSymbol_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
ctor_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_1, "reflectSymbol"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = ctor_3_0
v1_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), v_2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v1_4_1))).IntVal) == (0) {
__t2 = ctor_3_0
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("("), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(Get_intercalate(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array([]gopurs_runtime.Value{ctor_3_0}), v1_4_1)), gopurs_runtime.Str(")")))
}
end_branch_2:
return __t2
})))
}

func Call_genericShow_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericShow'")
}

func Call_genericShow_prime__func_gopurs_runtime_Value__interface____string_3379317986(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericShow'")
}

func Call_genericShowSum(dictGenericShow_0_loop gopurs_runtime.Value, dictGenericShow1_1_loop gopurs_runtime.Value) interface{} {
var dictGenericShow_0 gopurs_runtime.Value = dictGenericShow_0_loop
_ = dictGenericShow_0
var dictGenericShow1_1 gopurs_runtime.Value = dictGenericShow1_1_loop
_ = dictGenericShow1_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_0, "genericShow'"), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 492034566) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow1_1, "genericShow'"), gopurs_runtime.Any((*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}

func Call_genericShow(dictGeneric_0_loop gopurs_runtime.Value, dictGenericShow_1_loop gopurs_runtime.Value, x_2_loop interface{}) string {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericShow_1 gopurs_runtime.Value = dictGenericShow_1_loop
_ = dictGenericShow_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_1, "genericShow'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), gopurs_runtime.Any(x_2))).StrVal()
}
