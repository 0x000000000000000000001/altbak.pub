package Data_Show_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Type_Proxy "gopurs/output/Type.Proxy"
	unsafe "unsafe"
)

var genericShowArgsNoArguments gopurs_runtime.Value
var once_genericShowArgsNoArguments sync.Once
func Get_genericShowArgsNoArguments() gopurs_runtime.Value {
	once_genericShowArgsNoArguments.Do(func() {
		genericShowArgsNoArguments = gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{})
}))
	})
	return genericShowArgsNoArguments
}

var genericShowArgsArgument gopurs_runtime.Value
var once_genericShowArgsArgument sync.Once
func Get_genericShowArgsArgument() gopurs_runtime.Value {
	once_genericShowArgsArgument.Do(func() {
		genericShowArgsArgument = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1)})
}))
}()
})
	})
	return genericShowArgsArgument
}

var genericShowArgs gopurs_runtime.Value
var once_genericShowArgs sync.Once
func Get_genericShowArgs() gopurs_runtime.Value {
	once_genericShowArgs.Do(func() {
		genericShowArgs = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericShowArgs")
}()
})
	})
	return genericShowArgs
}

var genericShowArgsProduct gopurs_runtime.Value
var once_genericShowArgsProduct sync.Once
func Get_genericShowArgsProduct() gopurs_runtime.Value {
	once_genericShowArgsProduct.Do(func() {
		genericShowArgsProduct = gopurs_runtime.Func2(func(dictGenericShowArgs_0_box gopurs_runtime.Value, dictGenericShowArgs1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShowArgsProduct(dictGenericShowArgs_0_box, dictGenericShowArgs1_1_box)
})
	})
	return genericShowArgsProduct
}

var genericShowConstructor gopurs_runtime.Value
var once_genericShowConstructor sync.Once
func Get_genericShowConstructor() gopurs_runtime.Value {
	once_genericShowConstructor.Do(func() {
		genericShowConstructor = gopurs_runtime.Func2(func(dictGenericShowArgs_0_box gopurs_runtime.Value, dictIsSymbol_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShowConstructor(dictGenericShowArgs_0_box, dictIsSymbol_1_box)
})
	})
	return genericShowConstructor
}

var genericShow_prime gopurs_runtime.Value
var once_genericShow_prime sync.Once
func Get_genericShow_prime() gopurs_runtime.Value {
	once_genericShow_prime.Do(func() {
		genericShow_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericShow'")
}()
})
	})
	return genericShow_prime
}

var genericShowNoConstructors gopurs_runtime.Value
var once_genericShowNoConstructors sync.Once
func Get_genericShowNoConstructors() gopurs_runtime.Value {
	once_genericShowNoConstructors.Do(func() {
		genericShowNoConstructors = gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_genericShowNoConstructors(), "genericShow'"), a_0)
}))
	})
	return genericShowNoConstructors
}

var genericShowSum gopurs_runtime.Value
var once_genericShowSum sync.Once
func Get_genericShowSum() gopurs_runtime.Value {
	once_genericShowSum.Do(func() {
		genericShowSum = gopurs_runtime.Func2(func(dictGenericShow_0_box gopurs_runtime.Value, dictGenericShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShowSum(dictGenericShow_0_box, dictGenericShow1_1_box)
})
	})
	return genericShowSum
}

var genericShow gopurs_runtime.Value
var once_genericShow sync.Once
func Get_genericShow() gopurs_runtime.Value {
	once_genericShow.Do(func() {
		genericShow = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericShow_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericShow(dictGeneric_0_box, dictGenericShow_1_box, x_2_box)
})
	})
	return genericShow
}

func Call_genericShowArgsProduct(dictGenericShowArgs_0_loop gopurs_runtime.Value, dictGenericShowArgs1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericShowArgs_0 gopurs_runtime.Value = dictGenericShowArgs_0_loop
_ = dictGenericShowArgs_0
var dictGenericShowArgs1_1 gopurs_runtime.Value = dictGenericShowArgs1_1_loop
_ = dictGenericShowArgs1_1
return gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs1_1, "genericShowArgs"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1))
}))
}

func Call_genericShowConstructor(dictGenericShowArgs_0_loop gopurs_runtime.Value, dictIsSymbol_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericShowArgs_0 gopurs_runtime.Value = dictGenericShowArgs_0_loop
_ = dictGenericShowArgs_0
var dictIsSymbol_1 gopurs_runtime.Value = dictIsSymbol_1_loop
_ = dictIsSymbol_1
return gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
ctor_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_1, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 3178699476, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})})
_ = ctor_3_0
v1_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), v_2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Int(int64(len(v1_4_1.PtrVal().([]gopurs_runtime.Value)))).IntVal == 0 {
__t2 = ctor_3_0
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Str("(" + gopurs_runtime.Apply2(Get_intercalate(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{ctor_3_0}), v1_4_1)).StrVal() + ")")
}
end_branch_2:
return __t2
}))
}

func Call_genericShowSum(dictGenericShow_0_loop gopurs_runtime.Value, dictGenericShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericShow_0 gopurs_runtime.Value = dictGenericShow_0_loop
_ = dictGenericShow_0
var dictGenericShow1_1 gopurs_runtime.Value = dictGenericShow1_1_loop
_ = dictGenericShow1_1
return gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 164387955) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_0, "genericShow'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inl)(v_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 4051932077) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow1_1, "genericShow'"), (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Inr)(v_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_genericShow(dictGeneric_0_loop gopurs_runtime.Value, dictGenericShow_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericShow_1 gopurs_runtime.Value = dictGenericShow_1_loop
_ = dictGenericShow_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_1, "genericShow'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2))
}

func Get_intercalate() gopurs_runtime.Value {
	return _Gopurs_Intercalate
}
