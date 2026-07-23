package Data_Ord_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericOrdNoConstructors gopurs_runtime.Value
var once_genericOrdNoConstructors sync.Once
func Get_genericOrdNoConstructors() gopurs_runtime.Value {
	once_genericOrdNoConstructors.Do(func() {
		genericOrdNoConstructors = gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
}))
	})
	return genericOrdNoConstructors
}

var genericOrdNoArguments gopurs_runtime.Value
var once_genericOrdNoArguments sync.Once
func Get_genericOrdNoArguments() gopurs_runtime.Value {
	once_genericOrdNoArguments.Do(func() {
		genericOrdNoArguments = gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("EQ"))
}))
	})
	return genericOrdNoArguments
}

var genericOrdArgument gopurs_runtime.Value
var once_genericOrdArgument sync.Once
func Get_genericOrdArgument() gopurs_runtime.Value {
	once_genericOrdArgument.Do(func() {
		genericOrdArgument = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
}))
})
	})
	return genericOrdArgument
}

var genericCompare_prime gopurs_runtime.Value
var once_genericCompare_prime sync.Once
func Get_genericCompare_prime() gopurs_runtime.Value {
	once_genericCompare_prime.Do(func() {
		genericCompare_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericCompare'")
})
	})
	return genericCompare_prime
}

var genericOrdConstructor gopurs_runtime.Value
var once_genericOrdConstructor sync.Once
func Get_genericOrdConstructor() gopurs_runtime.Value {
	once_genericOrdConstructor.Do(func() {
		genericOrdConstructor = gopurs_runtime.Func(func(dictGenericOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), v_1, v1_2)
}))
})
	})
	return genericOrdConstructor
}

var genericOrdProduct gopurs_runtime.Value
var once_genericOrdProduct sync.Once
func Get_genericOrdProduct() gopurs_runtime.Value {
	once_genericOrdProduct.Do(func() {
		genericOrdProduct = gopurs_runtime.Func2(func(dictGenericOrd_0 gopurs_runtime.Value, dictGenericOrd1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), gopurs_runtime.RecordGet(v_2, "value0"), gopurs_runtime.RecordGet(v1_3, "value0"))
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4_0, "_tag").StrVal == "EQ")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd1_1, "genericCompare'"), gopurs_runtime.RecordGet(v_2, "value1"), gopurs_runtime.RecordGet(v1_3, "value1"))
goto end_branch_1
} else {

}
}
{
__t1 = v2_4_0
}
end_branch_1:
return __t1
}))
})
	})
	return genericOrdProduct
}

var genericOrdSum gopurs_runtime.Value
var once_genericOrdSum sync.Once
func Get_genericOrdSum() gopurs_runtime.Value {
	once_genericOrdSum.Do(func() {
		genericOrdSum = gopurs_runtime.Func2(func(dictGenericOrd_0 gopurs_runtime.Value, dictGenericOrd1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericCompare'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Inl")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Inl")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_0, "genericCompare'"), gopurs_runtime.RecordGet(v_2, "value0"), gopurs_runtime.RecordGet(v1_3, "value0"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Inr")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Inr")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Inr")).IntVal != 0 {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd1_1, "genericCompare'"), gopurs_runtime.RecordGet(v_2, "value0"), gopurs_runtime.RecordGet(v1_3, "value0"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Inl")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = __t2
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
	return genericOrdSum
}

var genericCompare gopurs_runtime.Value
var once_genericCompare sync.Once
func Get_genericCompare() gopurs_runtime.Value {
	once_genericCompare.Do(func() {
		genericCompare = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericOrd_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericOrd_1, "genericCompare'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3))
})
	})
	return genericCompare
}


