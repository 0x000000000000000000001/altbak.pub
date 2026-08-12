package Data_Enum_Generic

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
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
if (gopurs_runtime.Apply2(Get_eq__2843686287(), i_0, gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](__t0))}
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

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__724144906 gopurs_runtime.Value
var once_apply__724144906 sync.Once
func Get_apply__724144906() gopurs_runtime.Value {
	once_apply__724144906.Do(func() {
		cache_apply__724144906 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__724144906
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_genericBottom_prime__2482077850 gopurs_runtime.Value
var once_genericBottom_prime__2482077850 sync.Once
func Get_genericBottom_prime__2482077850() gopurs_runtime.Value {
	once_genericBottom_prime__2482077850.Do(func() {
		cache_genericBottom_prime__2482077850 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom_prime__2482077850(dict_0_box)
})
	})
	return cache_genericBottom_prime__2482077850
}

var cache_genericTop_prime__1114003774 gopurs_runtime.Value
var once_genericTop_prime__1114003774 sync.Once
func Get_genericTop_prime__1114003774() gopurs_runtime.Value {
	once_genericTop_prime__1114003774.Do(func() {
		cache_genericTop_prime__1114003774 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop_prime__1114003774(dict_0_box)
})
	})
	return cache_genericTop_prime__1114003774
}

var cache_genericFromEnum_prime__4119640152 gopurs_runtime.Value
var once_genericFromEnum_prime__4119640152 sync.Once
func Get_genericFromEnum_prime__4119640152() gopurs_runtime.Value {
	once_genericFromEnum_prime__4119640152.Do(func() {
		cache_genericFromEnum_prime__4119640152 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFromEnum_prime__4119640152(gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericFromEnum_prime__4119640152
}

var cache_genericPred_prime__3707548623 gopurs_runtime.Value
var once_genericPred_prime__3707548623 sync.Once
func Get_genericPred_prime__3707548623() gopurs_runtime.Value {
	once_genericPred_prime__3707548623.Do(func() {
		cache_genericPred_prime__3707548623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericPred_prime__3707548623(gopurs_runtime.CoerceToStruct[Constructor_GenericEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericPred_prime__3707548623
}

var cache_genericSucc_prime__3707548623 gopurs_runtime.Value
var once_genericSucc_prime__3707548623 sync.Once
func Get_genericSucc_prime__3707548623() gopurs_runtime.Value {
	once_genericSucc_prime__3707548623.Do(func() {
		cache_genericSucc_prime__3707548623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericSucc_prime__3707548623(gopurs_runtime.CoerceToStruct[Constructor_GenericEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericSucc_prime__3707548623
}

var cache_genericToEnum_prime__244040409 gopurs_runtime.Value
var once_genericToEnum_prime__244040409 sync.Once
func Get_genericToEnum_prime__244040409() gopurs_runtime.Value {
	once_genericToEnum_prime__244040409.Do(func() {
		cache_genericToEnum_prime__244040409 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericToEnum_prime__244040409(gopurs_runtime.CoerceToStruct[Constructor_GenericBoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericToEnum_prime__244040409
}

var cache_fromEnum__1637084359 gopurs_runtime.Value
var once_fromEnum__1637084359 sync.Once
func Get_fromEnum__1637084359() gopurs_runtime.Value {
	once_fromEnum__1637084359.Do(func() {
		cache_fromEnum__1637084359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1637084359(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromEnum__1637084359
}

var cache_pred__3199041328 gopurs_runtime.Value
var once_pred__3199041328 sync.Once
func Get_pred__3199041328() gopurs_runtime.Value {
	once_pred__3199041328.Do(func() {
		cache_pred__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred__3199041328(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pred__3199041328
}

var cache_succ__3199041328 gopurs_runtime.Value
var once_succ__3199041328 sync.Once
func Get_succ__3199041328() gopurs_runtime.Value {
	once_succ__3199041328.Do(func() {
		cache_succ__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__3199041328(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_succ__3199041328
}

var cache_toEnum__3317293286 gopurs_runtime.Value
var once_toEnum__3317293286 sync.Once
func Get_toEnum__3317293286() gopurs_runtime.Value {
	once_toEnum__3317293286.Do(func() {
		cache_toEnum__3317293286 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__3317293286(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toEnum__3317293286
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = pkg_Data_Eq.Get_eqIntImpl()
	})
	return cache_eq__2843686287
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_div__2185172824 gopurs_runtime.Value
var once_div__2185172824 sync.Once
func Get_div__2185172824() gopurs_runtime.Value {
	once_div__2185172824.Do(func() {
		cache_div__2185172824 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "div")
	})
	return cache_div__2185172824
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod")
	})
	return cache_mod__2185172824
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__1540101856 gopurs_runtime.Value
var once_flip__1540101856 sync.Once
func Get_flip__1540101856() gopurs_runtime.Value {
	once_flip__1540101856.Do(func() {
		cache_flip__1540101856 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1540101856(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1540101856
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_map__2486200924 gopurs_runtime.Value
var once_map__2486200924 sync.Once
func Get_map__2486200924() gopurs_runtime.Value {
	once_map__2486200924.Do(func() {
		cache_map__2486200924 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2486200924
}

var cache_map__2670646620 gopurs_runtime.Value
var once_map__2670646620 sync.Once
func Get_map__2670646620() gopurs_runtime.Value {
	once_map__2670646620.Do(func() {
		cache_map__2670646620 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2670646620
}

var cache_map__48293596 gopurs_runtime.Value
var once_map__48293596 sync.Once
func Get_map__48293596() gopurs_runtime.Value {
	once_map__48293596.Do(func() {
		cache_map__48293596 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__48293596
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj")
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_applyMaybe__3561700045 gopurs_runtime.Value
var once_applyMaybe__3561700045 sync.Once
func Get_applyMaybe__3561700045() gopurs_runtime.Value {
	once_applyMaybe__3561700045.Do(func() {
		cache_applyMaybe__3561700045 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3561700045
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_unwrap__1132695379 gopurs_runtime.Value
var once_unwrap__1132695379 sync.Once
func Get_unwrap__1132695379() gopurs_runtime.Value {
	once_unwrap__1132695379.Do(func() {
		cache_unwrap__1132695379 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_unwrap__1132695379
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_unwrap__4214189139 gopurs_runtime.Value
var once_unwrap__4214189139 sync.Once
func Get_unwrap__4214189139() gopurs_runtime.Value {
	once_unwrap__4214189139.Do(func() {
		cache_unwrap__4214189139 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_unwrap__4214189139
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__560788792 gopurs_runtime.Value
var once_mul__560788792 sync.Once
func Get_mul__560788792() gopurs_runtime.Value {
	once_mul__560788792.Do(func() {
		cache_mul__560788792 = pkg_Data_Semiring.Get_intMul()
	})
	return cache_mul__560788792
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
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
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))))}
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
v1_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = v1_5_0
var __t1 gopurs_runtime.Value
{
if (v1_5_0.Type == 9 && v1_5_0.IntVal == 930809136 && v1_5_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'")})}})}
goto end_branch_1
} else {

}
}
{
if (v1_5_0.Type == 9 && v1_5_0.IntVal == 930809136 && v1_5_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_5_0.UnsafePtr).V0})}})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3478632216) {
v1_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = v1_5_3
var __t4 gopurs_runtime.Value
{
if (v1_5_3.Type == 9 && v1_5_3.IntVal == 930809136 && v1_5_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericBottom_3, "genericBottom'")})}})}
goto end_branch_4
} else {

}
}
{
if (v1_5_3.Type == 9 && v1_5_3.IntVal == 930809136 && v1_5_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_5_3.UnsafePtr).V0})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 492034566) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_2, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0))))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t5))}
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
var __t2 gopurs_runtime.Value
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_0.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
if (v1_7_0.Type == 9 && v1_7_0.IntVal == 930809136 && v1_7_0.UnsafePtr == nil) {
__local_var_8_1 := gopurs_runtime.RecordGet(dictGenericTop1_4, "genericTop'")
_ = __local_var_8_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, __local_var_8_1})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericPred'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](__t2))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
v1_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum1_3, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
_ = v1_7_3
var __t5 gopurs_runtime.Value
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_7_3.UnsafePtr).V0})}})}
goto end_branch_5
} else {

}
}
{
if (v1_7_3.Type == 9 && v1_7_3.IntVal == 930809136 && v1_7_3.UnsafePtr == nil) {
__local_var_8_4 := gopurs_runtime.RecordGet(dictGenericBottom1_5, "genericBottom'")
_ = __local_var_8_4
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, __local_var_8_4})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericEnum_0, "genericSucc'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](__t5))}
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
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), genericCardinality_prime1_1_0, gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericCardinality'")).IntVal), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), genericCardinality_prime1_1_0).IntVal)
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
var __t2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(n_3, gopurs_runtime.Int(0))), gopurs_runtime.Bool(Call_lessThan__4087042607(n_3, genericCardinality_prime1_1_0))).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), n_3)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Inr(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Apply2(Get_sub__1043827704(), n_3, genericCardinality_prime1_1_0))))
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
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
return gopurs_runtime.RecordDict3("genericCardinality'", "genericFromEnum'", "genericToEnum'", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_mul__560788792(), genericCardinality_prime1_1_0, genericCardinality_prime2_3_1).IntVal), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Apply2(Get_mul__560788792(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), genericCardinality_prime2_3_1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericFromEnum'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)).IntVal)
}), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Generic_Rep.Get_Product(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum_0, "genericToEnum'"), gopurs_runtime.Apply2(Get_div__2185172824(), n_4, genericCardinality_prime2_3_1))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericBoundedEnum1_2, "genericToEnum'"), gopurs_runtime.Apply2(Get_mod__2185172824(), n_4, genericCardinality_prime2_3_1)))))}
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

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericBottom_prime__2482077850(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericBottom'")
}

func Call_genericTop_prime__1114003774(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTop'")
}

func Call_genericFromEnum_prime__4119640152(dict_0_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericPred_prime__3707548623(dict_0_loop *Constructor_GenericEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericSucc_prime__3707548623(dict_0_loop *Constructor_GenericEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericToEnum_prime__244040409(dict_0_loop *Constructor_GenericBoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericBoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_pred__3199041328(dict_0_loop *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_succ__3199041328(dict_0_loop *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1540101856(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


