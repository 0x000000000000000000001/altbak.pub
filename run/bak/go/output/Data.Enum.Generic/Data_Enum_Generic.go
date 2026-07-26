package Data_Enum_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	unsafe "unsafe"
)

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420)) != (true))
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_genericToEnum_prime gopurs_runtime.Value
var once_genericToEnum_prime sync.Once
func Get_genericToEnum_prime() gopurs_runtime.Value {
	once_genericToEnum_prime.Do(func() {
		cache_genericToEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum_prime(dict_0_box)
})
	})
	return cache_genericToEnum_prime
}

var cache_genericToEnum gopurs_runtime.Value
var once_genericToEnum sync.Once
func Get_genericToEnum() gopurs_runtime.Value {
	once_genericToEnum.Do(func() {
		cache_genericToEnum = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum(dictGeneric_0_box)
})
	})
	return cache_genericToEnum
}

var cache_genericSucc_prime gopurs_runtime.Value
var once_genericSucc_prime sync.Once
func Get_genericSucc_prime() gopurs_runtime.Value {
	once_genericSucc_prime.Do(func() {
		cache_genericSucc_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc_prime(dict_0_box)
})
	})
	return cache_genericSucc_prime
}

var cache_genericSucc gopurs_runtime.Value
var once_genericSucc sync.Once
func Get_genericSucc() gopurs_runtime.Value {
	once_genericSucc.Do(func() {
		cache_genericSucc = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc(dictGeneric_0_box)
})
	})
	return cache_genericSucc
}

var cache_genericPred_prime gopurs_runtime.Value
var once_genericPred_prime sync.Once
func Get_genericPred_prime() gopurs_runtime.Value {
	once_genericPred_prime.Do(func() {
		cache_genericPred_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred_prime(dict_0_box)
})
	})
	return cache_genericPred_prime
}

var cache_genericPred gopurs_runtime.Value
var once_genericPred sync.Once
func Get_genericPred() gopurs_runtime.Value {
	once_genericPred.Do(func() {
		cache_genericPred = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred(dictGeneric_0_box)
})
	})
	return cache_genericPred
}

var cache_genericFromEnum_prime gopurs_runtime.Value
var once_genericFromEnum_prime sync.Once
func Get_genericFromEnum_prime() gopurs_runtime.Value {
	once_genericFromEnum_prime.Do(func() {
		cache_genericFromEnum_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum_prime(dict_0_box)
})
	})
	return cache_genericFromEnum_prime
}

var cache_genericFromEnum gopurs_runtime.Value
var once_genericFromEnum sync.Once
func Get_genericFromEnum() gopurs_runtime.Value {
	once_genericFromEnum.Do(func() {
		cache_genericFromEnum = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBoundedEnum_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum(dictGeneric_0_box, dictGenericBoundedEnum_1_box, x_2_box)
})
	})
	return cache_genericFromEnum
}

var cache_genericEnumSum gopurs_runtime.Value
var once_genericEnumSum sync.Once
func Get_genericEnumSum() gopurs_runtime.Value {
	once_genericEnumSum.Do(func() {
		cache_genericEnumSum = gopurs_runtime.Func2(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumSum(dictGenericEnum_0_box, dictGenericTop_1_box)
})
	})
	return cache_genericEnumSum
}

var cache_genericEnumProduct gopurs_runtime.Value
var once_genericEnumProduct sync.Once
func Get_genericEnumProduct() gopurs_runtime.Value {
	once_genericEnumProduct.Do(func() {
		cache_genericEnumProduct = gopurs_runtime.Func5(func(dictGenericEnum_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value, dictGenericBottom_2_box gopurs_runtime.Value, dictGenericEnum1_3_box gopurs_runtime.Value, dictGenericTop1_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericEnumProduct(dictGenericEnum_0_box, dictGenericTop_1_box, dictGenericBottom_2_box, dictGenericEnum1_3_box, dictGenericTop1_4_box)
})
	})
	return cache_genericEnumProduct
}

var cache_genericEnumNoArguments gopurs_runtime.Value
var once_genericEnumNoArguments sync.Once
func Get_genericEnumNoArguments() gopurs_runtime.Value {
	once_genericEnumNoArguments.Do(func() {
		cache_genericEnumNoArguments = gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}
end_branch_0:
return __t0
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

func Call_genericToEnum_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericToEnum'")
}

func Call_genericToEnum(dictGeneric_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := gopurs_runtime.RecordGet(dictGeneric_0, "to")
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_2, "genericToEnum'"), x_4))
})
})
}

func Call_genericSucc_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericSucc'")
}

func Call_genericSucc(dictGeneric_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := gopurs_runtime.RecordGet(dictGeneric_0, "to")
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericSucc'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_4)))
})
})
}

func Call_genericPred_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericPred'")
}

func Call_genericPred(dictGeneric_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
to_1_0 := gopurs_runtime.RecordGet(dictGeneric_0, "to")
_ = to_1_0
return gopurs_runtime.Func(func(dictGenericEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), to_1_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_2, "genericPred'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_4)))
})
})
}

func Call_genericFromEnum_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericFromEnum'")
}

func Call_genericFromEnum(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBoundedEnum_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBoundedEnum_1 gopurs_runtime.Value = dictGenericBoundedEnum_1_loop
_ = dictGenericBoundedEnum_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_1, "genericFromEnum'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2))
}

func Call_genericEnumSum(dictGenericEnum_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
genericTop_prime_2_0 := gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'")
_ = genericTop_prime_2_0
return gopurs_runtime.Func2(func(dictGenericEnum1_3 gopurs_runtime.Value, dictGenericBottom_4 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_5_1 := gopurs_runtime.RecordGet(dictGenericBottom_4, "genericBottom'")
_ = genericBottom_prime_5_1
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3478632216) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 492034566) {
v1_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
_ = v1_7_3
var __t4 gopurs_runtime.Value
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{genericTop_prime_2_0})}})}
goto end_branch_4
} else {

}
}
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_3.UnsafePtr).V0})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t2 = __t4
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
var __t5 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3478632216) {
v1_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
_ = v1_7_6
var __t7 gopurs_runtime.Value
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{genericBottom_prime_5_1})}})}
goto end_branch_7
} else {

}
}
{
if (v1_7_6.Type == 9 && v1_7_6.IntVal == 930809136 && v1_7_6.UnsafePtr != nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_6.UnsafePtr).V0})}})}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t5 = __t7
goto end_branch_5
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 492034566) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
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
genericTop_prime_5_0 := gopurs_runtime.RecordGet(dictGenericTop1_4, "genericTop'")
_ = genericTop_prime_5_0
return gopurs_runtime.Func(func(dictGenericBottom1_6 gopurs_runtime.Value) gopurs_runtime.Value {
genericBottom_prime_7_1 := gopurs_runtime.RecordGet(dictGenericBottom1_6, "genericBottom'")
_ = genericBottom_prime_7_1
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
v1_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)
_ = v1_9_2
var __t3 gopurs_runtime.Value
{
if (v1_9_2.Type == 9 && v1_9_2.IntVal == 930809136 && v1_9_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_9_2.UnsafePtr).V0})}})}
goto end_branch_3
} else {

}
}
{
if (v1_9_2.Type == 9 && v1_9_2.IntVal == 930809136 && v1_9_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{a_10, genericTop_prime_5_0})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
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
v1_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)
_ = v1_9_4
var __t5 gopurs_runtime.Value
{
if (v1_9_4.Type == 9 && v1_9_4.IntVal == 930809136 && v1_9_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_9_4.UnsafePtr).V0})}})}
goto end_branch_5
} else {

}
}
{
if (v1_9_4.Type == 9 && v1_9_4.IntVal == 930809136 && v1_9_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{a_10, genericBottom_prime_7_1})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
}

func Call_genericEnumConstructor(dictGenericEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEnum_0 gopurs_runtime.Value = dictGenericEnum_0_loop
_ = dictGenericEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), v_1))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), v_1))
}))
}

func Call_genericEnumArgument(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
return gopurs_runtime.RecordDict2("genericPred'", "genericSucc'", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), v_1))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), v_1))
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
return gopurs_runtime.RecordGet(dictGenericBoundedEnum_1, "genericCardinality'")
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
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)
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
return __t1
}), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThanOrEq(), n_3, gopurs_runtime.Int(0)), gopurs_runtime.Apply2(Get_lessThan(), n_3, genericCardinality_prime1_1_0)).IntVal) != (0) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), n_3))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Int((n_3.IntVal) - (genericCardinality_prime1_1_0.IntVal))))
}
end_branch_2:
return __t2
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Product(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "div"), n_4, genericCardinality_prime2_3_1))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), n_4, genericCardinality_prime2_3_1)))
}))
})
}

func Call_genericBoundedEnumConstructor(dictGenericBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBoundedEnum_0 gopurs_runtime.Value = dictGenericBoundedEnum_0_loop
_ = dictGenericBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericCardinality'"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), v_1)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Constructor(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), i_1))
}))
}

func Call_genericBoundedEnumArgument(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.RecordGet(dictBoundedEnum_0, "cardinality"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), v_1)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Argument(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), i_1))
}))
}


