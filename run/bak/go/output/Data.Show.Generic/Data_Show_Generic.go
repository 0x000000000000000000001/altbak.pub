package Data_Show_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
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
		genericShowArgsArgument = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1)})
}))
})
	})
	return genericShowArgsArgument
}

var genericShowArgs gopurs_runtime.Value
var once_genericShowArgs sync.Once
func Get_genericShowArgs() gopurs_runtime.Value {
	once_genericShowArgs.Do(func() {
		genericShowArgs = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericShowArgs")
})
	})
	return genericShowArgs
}

var genericShowArgsProduct gopurs_runtime.Value
var once_genericShowArgsProduct sync.Once
func Get_genericShowArgsProduct() gopurs_runtime.Value {
	once_genericShowArgsProduct.Do(func() {
		genericShowArgsProduct = gopurs_runtime.Func2(func(dictGenericShowArgs_0 gopurs_runtime.Value, dictGenericShowArgs1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericShowArgs", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs1_1, "genericShowArgs"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
}))
})
	})
	return genericShowArgsProduct
}

var genericShowConstructor gopurs_runtime.Value
var once_genericShowConstructor sync.Once
func Get_genericShowConstructor() gopurs_runtime.Value {
	once_genericShowConstructor.Do(func() {
		genericShowConstructor = gopurs_runtime.Func2(func(dictGenericShowArgs_0 gopurs_runtime.Value, dictIsSymbol_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
ctor_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_1, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = ctor_3_0
v1_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShowArgs_0, "genericShowArgs"), v_2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Int(int64(len(v1_4_1.PtrVal.([]gopurs_runtime.Value)))).IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = ctor_3_0
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(").StrVal + gopurs_runtime.Apply2(Get_intercalate(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{ctor_3_0}), v1_4_1)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}
end_branch_2:
return __t2
}))
})
	})
	return genericShowConstructor
}

var genericShow_prime gopurs_runtime.Value
var once_genericShow_prime sync.Once
func Get_genericShow_prime() gopurs_runtime.Value {
	once_genericShow_prime.Do(func() {
		genericShow_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericShow'")
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
		genericShowSum = gopurs_runtime.Func2(func(dictGenericShow_0 gopurs_runtime.Value, dictGenericShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericShow'", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.StrVal == "Inl")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_0, "genericShow'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.StrVal == "Inr")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow1_1, "genericShow'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0])
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
})
	})
	return genericShowSum
}

var genericShow gopurs_runtime.Value
var once_genericShow sync.Once
func Get_genericShow() gopurs_runtime.Value {
	once_genericShow.Do(func() {
		genericShow = gopurs_runtime.Func3(func(dictGeneric_0 gopurs_runtime.Value, dictGenericShow_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericShow_1, "genericShow'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2))
})
	})
	return genericShow
}

func Get_intercalate() gopurs_runtime.Value {
	return _Gopurs_Intercalate
}
