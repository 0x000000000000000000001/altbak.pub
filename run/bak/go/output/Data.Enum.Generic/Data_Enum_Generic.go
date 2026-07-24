package Data_Enum_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var genericToEnum_prime gopurs_runtime.Value
var once_genericToEnum_prime sync.Once
func Get_genericToEnum_prime() gopurs_runtime.Value {
	once_genericToEnum_prime.Do(func() {
		genericToEnum_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericToEnum'")
}()
})
	})
	return genericToEnum_prime
}

var genericToEnum gopurs_runtime.Value
var once_genericToEnum sync.Once
func Get_genericToEnum() gopurs_runtime.Value {
	once_genericToEnum.Do(func() {
		genericToEnum = gopurs_runtime.Func3(Call_genericToEnum)
	})
	return genericToEnum
}

var genericSucc_prime gopurs_runtime.Value
var once_genericSucc_prime sync.Once
func Get_genericSucc_prime() gopurs_runtime.Value {
	once_genericSucc_prime.Do(func() {
		genericSucc_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericSucc'")
}()
})
	})
	return genericSucc_prime
}

var genericSucc gopurs_runtime.Value
var once_genericSucc sync.Once
func Get_genericSucc() gopurs_runtime.Value {
	once_genericSucc.Do(func() {
		genericSucc = gopurs_runtime.Func3(Call_genericSucc)
	})
	return genericSucc
}

var genericPred_prime gopurs_runtime.Value
var once_genericPred_prime sync.Once
func Get_genericPred_prime() gopurs_runtime.Value {
	once_genericPred_prime.Do(func() {
		genericPred_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericPred'")
}()
})
	})
	return genericPred_prime
}

var genericPred gopurs_runtime.Value
var once_genericPred sync.Once
func Get_genericPred() gopurs_runtime.Value {
	once_genericPred.Do(func() {
		genericPred = gopurs_runtime.Func3(Call_genericPred)
	})
	return genericPred
}

var genericFromEnum_prime gopurs_runtime.Value
var once_genericFromEnum_prime sync.Once
func Get_genericFromEnum_prime() gopurs_runtime.Value {
	once_genericFromEnum_prime.Do(func() {
		genericFromEnum_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericFromEnum'")
}()
})
	})
	return genericFromEnum_prime
}

var genericFromEnum gopurs_runtime.Value
var once_genericFromEnum sync.Once
func Get_genericFromEnum() gopurs_runtime.Value {
	once_genericFromEnum.Do(func() {
		genericFromEnum = gopurs_runtime.Func3(Call_genericFromEnum)
	})
	return genericFromEnum
}

var genericEnumSum gopurs_runtime.Value
var once_genericEnumSum sync.Once
func Get_genericEnumSum() gopurs_runtime.Value {
	once_genericEnumSum.Do(func() {
		genericEnumSum = gopurs_runtime.Func2(Call_genericEnumSum)
	})
	return genericEnumSum
}

var genericEnumProduct gopurs_runtime.Value
var once_genericEnumProduct sync.Once
func Get_genericEnumProduct() gopurs_runtime.Value {
	once_genericEnumProduct.Do(func() {
		genericEnumProduct = gopurs_runtime.Func5(Call_genericEnumProduct)
	})
	return genericEnumProduct
}

var genericEnumNoArguments gopurs_runtime.Value
var once_genericEnumNoArguments sync.Once
func Get_genericEnumNoArguments() gopurs_runtime.Value {
	once_genericEnumNoArguments.Do(func() {
		genericEnumNoArguments = gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}))
	})
	return genericEnumNoArguments
}

var genericEnumConstructor gopurs_runtime.Value
var once_genericEnumConstructor sync.Once
func Get_genericEnumConstructor() gopurs_runtime.Value {
	once_genericEnumConstructor.Do(func() {
		genericEnumConstructor = gopurs_runtime.Func(func(dictGenericEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0_loop, "genericPred'"), v_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0_loop, "genericSucc'"), v_1)
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_2.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_2_2.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}))
}()
})
	})
	return genericEnumConstructor
}

var genericEnumArgument gopurs_runtime.Value
var once_genericEnumArgument sync.Once
func Get_genericEnumArgument() gopurs_runtime.Value {
	once_genericEnumArgument.Do(func() {
		genericEnumArgument = gopurs_runtime.Func(func(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0_loop, "pred"), v_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0_loop, "succ"), v_1)
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_2.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_2_2.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}))
}()
})
	})
	return genericEnumArgument
}

var genericCardinality_prime gopurs_runtime.Value
var once_genericCardinality_prime sync.Once
func Get_genericCardinality_prime() gopurs_runtime.Value {
	once_genericCardinality_prime.Do(func() {
		genericCardinality_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericCardinality'")
}()
})
	})
	return genericCardinality_prime
}

var genericCardinality gopurs_runtime.Value
var once_genericCardinality sync.Once
func Get_genericCardinality() gopurs_runtime.Value {
	once_genericCardinality.Do(func() {
		genericCardinality = gopurs_runtime.Func2(Call_genericCardinality)
	})
	return genericCardinality
}

var genericBoundedEnumSum gopurs_runtime.Value
var once_genericBoundedEnumSum sync.Once
func Get_genericBoundedEnumSum() gopurs_runtime.Value {
	once_genericBoundedEnumSum.Do(func() {
		genericBoundedEnumSum = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
genericCardinality_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericCardinality'")
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("genericCardinality'", "genericToEnum'", "genericFromEnum'", gopurs_runtime.Int(genericCardinality_prime1_1_0.IntVal + gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'").IntVal), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if n_3.IntVal >= 0 && n_3.IntVal < genericCardinality_prime1_1_0.IntVal {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericToEnum'"), n_3)
_ = __local_var_4_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_4.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inl", (*[1024]gopurs_runtime.Value)(__local_var_4_4.UnsafePtr)[0]))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int(n_3.IntVal - genericCardinality_prime1_1_0.IntVal))
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_1.StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inr", (*[1024]gopurs_runtime.Value)(__local_var_4_1.UnsafePtr)[0]))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Inl").IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericFromEnum'"), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Inr").IntVal != 0 {
__t6 = gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]).IntVal + genericCardinality_prime1_1_0.IntVal)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
})
}()
})
	})
	return genericBoundedEnumSum
}

var genericBoundedEnumProduct gopurs_runtime.Value
var once_genericBoundedEnumProduct sync.Once
func Get_genericBoundedEnumProduct() gopurs_runtime.Value {
	once_genericBoundedEnumProduct.Do(func() {
		genericBoundedEnumProduct = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
genericCardinality_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericCardinality'")
_ = genericCardinality_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum1_2 gopurs_runtime.Value) gopurs_runtime.Value {
genericCardinality_prime2_3_1 := gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'")
_ = genericCardinality_prime2_3_1
return gopurs_runtime.RecordDict3("genericCardinality'", "genericToEnum'", "genericFromEnum'", gopurs_runtime.Int(genericCardinality_prime1_1_0.IntVal * genericCardinality_prime2_3_1.IntVal), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericToEnum'"), gopurs_runtime.Int(n_4.IntVal / genericCardinality_prime2_3_1.IntVal))
_ = __local_var_5_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_2.StrVal == "Just").IntVal != 0 {
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), n_4, genericCardinality_prime2_3_1))
_ = __local_var_6_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_4.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Product", (*[1024]gopurs_runtime.Value)(__local_var_5_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_6_4.UnsafePtr)[0]))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericFromEnum'"), (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[0]).IntVal * genericCardinality_prime2_3_1.IntVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[1]).IntVal)
}))
})
}()
})
	})
	return genericBoundedEnumProduct
}

var genericBoundedEnumNoArguments gopurs_runtime.Value
var once_genericBoundedEnumNoArguments sync.Once
func Get_genericBoundedEnumNoArguments() gopurs_runtime.Value {
	once_genericBoundedEnumNoArguments.Do(func() {
		genericBoundedEnumNoArguments = gopurs_runtime.RecordDict3("genericCardinality'", "genericToEnum'", "genericFromEnum'", gopurs_runtime.Int(1), gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if i_0.IntVal == 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("NoArguments"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}))
	})
	return genericBoundedEnumNoArguments
}

var genericBoundedEnumConstructor gopurs_runtime.Value
var once_genericBoundedEnumConstructor sync.Once
func Get_genericBoundedEnumConstructor() gopurs_runtime.Value {
	once_genericBoundedEnumConstructor.Do(func() {
		genericBoundedEnumConstructor = gopurs_runtime.Func(func(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericToEnum'", "genericFromEnum'", gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericCardinality'"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericToEnum'"), i_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0_loop, "genericFromEnum'"), v_1)
}))
}()
})
	})
	return genericBoundedEnumConstructor
}

var genericBoundedEnumArgument gopurs_runtime.Value
var once_genericBoundedEnumArgument sync.Once
func Get_genericBoundedEnumArgument() gopurs_runtime.Value {
	once_genericBoundedEnumArgument.Do(func() {
		genericBoundedEnumArgument = gopurs_runtime.Func(func(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericToEnum'", "genericFromEnum'", gopurs_runtime.RecordGet(dictBoundedEnum_0_loop, "cardinality"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0_loop, "toEnum"), i_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0_loop, "fromEnum"), v_1)
}))
}()
})
	})
	return genericBoundedEnumArgument
}

func Call_genericToEnum(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_1_loop, "genericToEnum'"), x_2_loop)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}

func Call_genericSucc(dictGeneric_0_loop gopurs_runtime.Value, dictGenericEnum_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEnum_1 gopurs_runtime.Value = dictGenericEnum_1_loop
_ = dictGenericEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_1_loop, "genericSucc'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), x_2_loop))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}

func Call_genericPred(dictGeneric_0_loop gopurs_runtime.Value, dictGenericEnum_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEnum_1 gopurs_runtime.Value = dictGenericEnum_1_loop
_ = dictGenericEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_1_loop, "genericPred'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), x_2_loop))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}

func Call_genericFromEnum(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_1_loop, "genericFromEnum'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), x_2_loop))
}

func Call_genericEnumSum(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
genericTop_prime_2_0 := gopurs_runtime.RecordGet(dictGenericTop_1_loop, "genericTop'")
_ = genericTop_prime_2_0
return gopurs_runtime.Func2(func(dictGenericEnum1_3 gopurs_runtime.Value, dictGenericBottom_4 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_5_1 := gopurs_runtime.RecordGet(dictGenericBottom_4, "genericBottom'")
_ = genericBottom_prime_5_1
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Inl").IntVal != 0 {
__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0_loop, "genericPred'"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])
_ = __local_var_7_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_7_3.StrVal == "Just").IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inl", (*[1024]gopurs_runtime.Value)(__local_var_7_3.UnsafePtr)[0]))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Inr").IntVal != 0 {
v1_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])
_ = v1_7_5
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7_5.StrVal == "Nothing").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inl", genericTop_prime_2_0))
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v1_7_5.StrVal == "Just").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inr", (*[1024]gopurs_runtime.Value)(v1_7_5.UnsafePtr)[0]))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t2 = __t6
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Inl").IntVal != 0 {
v1_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0_loop, "genericSucc'"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])
_ = v1_7_8
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7_8.StrVal == "Nothing").IntVal != 0 {
__t9 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inr", genericBottom_prime_5_1))
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool(v1_7_8.StrVal == "Just").IntVal != 0 {
__t9 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inl", (*[1024]gopurs_runtime.Value)(v1_7_8.UnsafePtr)[0]))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t7 = __t9
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Inr").IntVal != 0 {
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])
_ = __local_var_7_10
var __t11 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_7_10.StrVal == "Just").IntVal != 0 {
__t11 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Inr", (*[1024]gopurs_runtime.Value)(__local_var_7_10.UnsafePtr)[0]))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_11:
__t7 = __t11
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
})
}

func Call_genericEnumProduct(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value, dictGenericBottom_2_loop gopurs_runtime.Value, dictGenericEnum1_3_loop gopurs_runtime.Value, dictGenericTop1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
genericTop_prime_5_0 := gopurs_runtime.RecordGet(dictGenericTop1_4_loop, "genericTop'")
_ = genericTop_prime_5_0
return gopurs_runtime.Func(func(dictGenericBottom1_6 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_7_1 := gopurs_runtime.RecordGet(dictGenericBottom1_6, "genericBottom'")
_ = genericBottom_prime_7_1
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3_loop, "genericPred'"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1])
_ = v1_9_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_9_2.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Product", (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_9_2.UnsafePtr)[0]))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_9_2.StrVal == "Nothing").IntVal != 0 {
__local_var_10_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0_loop, "genericPred'"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0])
_ = __local_var_10_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_10_4.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Product", (*[1024]gopurs_runtime.Value)(__local_var_10_4.UnsafePtr)[0], genericTop_prime_5_0))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3_loop, "genericSucc'"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1])
_ = v1_9_6
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_9_6.StrVal == "Just").IntVal != 0 {
__t7 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Product", (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_9_6.UnsafePtr)[0]))
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(v1_9_6.StrVal == "Nothing").IntVal != 0 {
__local_var_10_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0_loop, "genericSucc'"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0])
_ = __local_var_10_8
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_10_8.StrVal == "Just").IntVal != 0 {
__t9 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Product", (*[1024]gopurs_runtime.Value)(__local_var_10_8.UnsafePtr)[0], genericBottom_prime_7_1))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_9:
__t7 = __t9
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
})
}

func Call_genericCardinality(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
return gopurs_runtime.RecordGet(dictGenericBoundedEnum_1_loop, "genericCardinality'")
}


