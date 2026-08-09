package Data_Enum

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThan
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_greaterThan
}

var cache_guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		cache_guard = gopurs_runtime.Apply(pkg_Control_Alternative.Get_guard(), pkg_Data_Maybe.Get_alternativeMaybe())
	})
	return cache_guard
}

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThanOrEq
}

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_Cardinality gopurs_runtime.Value
var once_Cardinality sync.Once
func Get_Cardinality() gopurs_runtime.Value {
	once_Cardinality.Do(func() {
		cache_Cardinality = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Cardinality(x_0_box)
})
	})
	return cache_Cardinality
}

var cache_toEnum gopurs_runtime.Value
var once_toEnum sync.Once
func Get_toEnum() gopurs_runtime.Value {
	once_toEnum.Do(func() {
		cache_toEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum(dict_0_box)
})
	})
	return cache_toEnum
}

var cache_toEnum__gopurs_runtime_Value_180026696 gopurs_runtime.Value
var once_toEnum__gopurs_runtime_Value_180026696 sync.Once
func Get_toEnum__gopurs_runtime_Value_180026696() gopurs_runtime.Value {
	once_toEnum__gopurs_runtime_Value_180026696.Do(func() {
		cache_toEnum__gopurs_runtime_Value_180026696 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__gopurs_runtime_Value_180026696(dict_0_box)
})
	})
	return cache_toEnum__gopurs_runtime_Value_180026696
}

var cache_succ gopurs_runtime.Value
var once_succ sync.Once
func Get_succ() gopurs_runtime.Value {
	once_succ.Do(func() {
		cache_succ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ(dict_0_box)
})
	})
	return cache_succ
}

var cache_succ__gopurs_runtime_Value_4187127373 gopurs_runtime.Value
var once_succ__gopurs_runtime_Value_4187127373 sync.Once
func Get_succ__gopurs_runtime_Value_4187127373() gopurs_runtime.Value {
	once_succ__gopurs_runtime_Value_4187127373.Do(func() {
		cache_succ__gopurs_runtime_Value_4187127373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__gopurs_runtime_Value_4187127373(dict_0_box)
})
	})
	return cache_succ__gopurs_runtime_Value_4187127373
}

var cache_upFromIncluding gopurs_runtime.Value
var once_upFromIncluding sync.Once
func Get_upFromIncluding() gopurs_runtime.Value {
	once_upFromIncluding.Do(func() {
		cache_upFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_upFromIncluding(dictEnum_0_box, dictUnfoldable1_1_box)
})
	})
	return cache_upFromIncluding
}

var cache_showCardinality gopurs_runtime.Value
var once_showCardinality sync.Once
func Get_showCardinality() gopurs_runtime.Value {
	once_showCardinality.Do(func() {
		cache_showCardinality = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Cardinality "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showCardinality
}

var cache_pred gopurs_runtime.Value
var once_pred sync.Once
func Get_pred() gopurs_runtime.Value {
	once_pred.Do(func() {
		cache_pred = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred(dict_0_box)
})
	})
	return cache_pred
}

var cache_pred__gopurs_runtime_Value_4187127373 gopurs_runtime.Value
var once_pred__gopurs_runtime_Value_4187127373 sync.Once
func Get_pred__gopurs_runtime_Value_4187127373() gopurs_runtime.Value {
	once_pred__gopurs_runtime_Value_4187127373.Do(func() {
		cache_pred__gopurs_runtime_Value_4187127373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred__gopurs_runtime_Value_4187127373(dict_0_box)
})
	})
	return cache_pred__gopurs_runtime_Value_4187127373
}

var cache_ordCardinality gopurs_runtime.Value
var once_ordCardinality sync.Once
func Get_ordCardinality() gopurs_runtime.Value {
	once_ordCardinality.Do(func() {
		cache_ordCardinality = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
	})
	return cache_ordCardinality
}

var cache_newtypeCardinality gopurs_runtime.Value
var once_newtypeCardinality sync.Once
func Get_newtypeCardinality() gopurs_runtime.Value {
	once_newtypeCardinality.Do(func() {
		cache_newtypeCardinality = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCardinality
}

var cache_fromEnum gopurs_runtime.Value
var once_fromEnum sync.Once
func Get_fromEnum() gopurs_runtime.Value {
	once_fromEnum.Do(func() {
		cache_fromEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum(dict_0_box)
})
	})
	return cache_fromEnum
}

var cache_fromEnum__gopurs_runtime_Value_2576611849 gopurs_runtime.Value
var once_fromEnum__gopurs_runtime_Value_2576611849 sync.Once
func Get_fromEnum__gopurs_runtime_Value_2576611849() gopurs_runtime.Value {
	once_fromEnum__gopurs_runtime_Value_2576611849.Do(func() {
		cache_fromEnum__gopurs_runtime_Value_2576611849 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__gopurs_runtime_Value_2576611849(dict_0_box)
})
	})
	return cache_fromEnum__gopurs_runtime_Value_2576611849
}

var cache_toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		cache_toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults(dictBoundedEnum_0_box)
})
	})
	return cache_toEnumWithDefaults
}

var cache_toEnumWithDefaults__gopurs_runtime_Value_3074683145 gopurs_runtime.Value
var once_toEnumWithDefaults__gopurs_runtime_Value_3074683145 sync.Once
func Get_toEnumWithDefaults__gopurs_runtime_Value_3074683145() gopurs_runtime.Value {
	once_toEnumWithDefaults__gopurs_runtime_Value_3074683145.Do(func() {
		cache_toEnumWithDefaults__gopurs_runtime_Value_3074683145 = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults__gopurs_runtime_Value_3074683145(dictBoundedEnum_0_box)
})
	})
	return cache_toEnumWithDefaults__gopurs_runtime_Value_3074683145
}

var cache_toEnumWithDefaults__gopurs_runtime_Value_1689825737 gopurs_runtime.Value
var once_toEnumWithDefaults__gopurs_runtime_Value_1689825737 sync.Once
func Get_toEnumWithDefaults__gopurs_runtime_Value_1689825737() gopurs_runtime.Value {
	once_toEnumWithDefaults__gopurs_runtime_Value_1689825737.Do(func() {
		cache_toEnumWithDefaults__gopurs_runtime_Value_1689825737 = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults__gopurs_runtime_Value_1689825737(dictBoundedEnum_0_box)
})
	})
	return cache_toEnumWithDefaults__gopurs_runtime_Value_1689825737
}

var cache_eqCardinality gopurs_runtime.Value
var once_eqCardinality sync.Once
func Get_eqCardinality() gopurs_runtime.Value {
	once_eqCardinality.Do(func() {
		cache_eqCardinality = gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
	})
	return cache_eqCardinality
}

var cache_enumUnit gopurs_runtime.Value
var once_enumUnit sync.Once
func Get_enumUnit() gopurs_runtime.Value {
	once_enumUnit.Do(func() {
		cache_enumUnit = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordUnit()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}))
	})
	return cache_enumUnit
}

var cache_enumTuple gopurs_runtime.Value
var once_enumTuple sync.Once
func Get_enumTuple() gopurs_runtime.Value {
	once_enumTuple.Do(func() {
		cache_enumTuple = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumTuple(dictEnum_0_box)
})
	})
	return cache_enumTuple
}

var cache_enumOrdering gopurs_runtime.Value
var once_enumOrdering sync.Once
func Get_enumOrdering() gopurs_runtime.Value {
	once_enumOrdering.Do(func() {
		cache_enumOrdering = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordOrdering()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t0.UnsafePtr))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
	})
	return cache_enumOrdering
}

var cache_enumMaybe gopurs_runtime.Value
var once_enumMaybe sync.Once
func Get_enumMaybe() gopurs_runtime.Value {
	once_enumMaybe.Do(func() {
		cache_enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumMaybe(dictBoundedEnum_0_box)
})
	})
	return cache_enumMaybe
}

var cache_enumInt gopurs_runtime.Value
var once_enumInt sync.Once
func Get_enumInt() gopurs_runtime.Value {
	once_enumInt.Do(func() {
		cache_enumInt = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), n_0, gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "bottom")).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((n_0.IntVal) - (1))})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(__t0.UnsafePtr))}
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), n_0, gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "top")).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((n_0.IntVal) + (1))})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(__t1.UnsafePtr))}
}))
	})
	return cache_enumInt
}

var cache_enumFromTo gopurs_runtime.Value
var once_enumFromTo sync.Once
func Get_enumFromTo() gopurs_runtime.Value {
	once_enumFromTo.Do(func() {
		cache_enumFromTo = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromTo(dictEnum_0_box)
})
	})
	return cache_enumFromTo
}

var cache_enumFromTo__gopurs_runtime_Value_631221758 gopurs_runtime.Value
var once_enumFromTo__gopurs_runtime_Value_631221758 sync.Once
func Get_enumFromTo__gopurs_runtime_Value_631221758() gopurs_runtime.Value {
	once_enumFromTo__gopurs_runtime_Value_631221758.Do(func() {
		cache_enumFromTo__gopurs_runtime_Value_631221758 = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromTo__gopurs_runtime_Value_631221758(dictEnum_0_box)
})
	})
	return cache_enumFromTo__gopurs_runtime_Value_631221758
}

var cache_enumFromTo__gopurs_runtime_Value_1236426744 gopurs_runtime.Value
var once_enumFromTo__gopurs_runtime_Value_1236426744 sync.Once
func Get_enumFromTo__gopurs_runtime_Value_1236426744() gopurs_runtime.Value {
	once_enumFromTo__gopurs_runtime_Value_1236426744.Do(func() {
		cache_enumFromTo__gopurs_runtime_Value_1236426744 = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromTo__gopurs_runtime_Value_1236426744(dictEnum_0_box)
})
	})
	return cache_enumFromTo__gopurs_runtime_Value_1236426744
}

var cache_enumFromThenTo gopurs_runtime.Value
var once_enumFromThenTo sync.Once
func Get_enumFromThenTo() gopurs_runtime.Value {
	once_enumFromThenTo.Do(func() {
		cache_enumFromThenTo = gopurs_runtime.Func6(func(dictUnfoldable_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictBoundedEnum_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value, b_4_box gopurs_runtime.Value, c_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromThenTo(dictUnfoldable_0_box, dictFunctor_1_box, dictBoundedEnum_2_box, a_3_box, b_4_box, c_5_box)
})
	})
	return cache_enumFromThenTo
}

var cache_enumEither gopurs_runtime.Value
var once_enumEither sync.Once
func Get_enumEither() gopurs_runtime.Value {
	once_enumEither.Do(func() {
		cache_enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumEither(dictBoundedEnum_0_box)
})
	})
	return cache_enumEither
}

var cache_enumBoolean gopurs_runtime.Value
var once_enumBoolean sync.Once
func Get_enumBoolean() gopurs_runtime.Value {
	once_enumBoolean.Do(func() {
		cache_enumBoolean = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordBoolean()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(false)})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(__t0.UnsafePtr))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_0.IntVal) != (0)) != (true) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(true)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(__t1.UnsafePtr))}
}))
	})
	return cache_enumBoolean
}

var cache_downFromIncluding gopurs_runtime.Value
var once_downFromIncluding sync.Once
func Get_downFromIncluding() gopurs_runtime.Value {
	once_downFromIncluding.Do(func() {
		cache_downFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_downFromIncluding(dictEnum_0_box, dictUnfoldable1_1_box)
})
	})
	return cache_downFromIncluding
}

var cache_diag gopurs_runtime.Value
var once_diag sync.Once
func Get_diag() gopurs_runtime.Value {
	once_diag.Do(func() {
		cache_diag = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diag(a_0_box)
})
	})
	return cache_diag
}

var cache_diag__gopurs_runtime_Value_2627963987 gopurs_runtime.Value
var once_diag__gopurs_runtime_Value_2627963987 sync.Once
func Get_diag__gopurs_runtime_Value_2627963987() gopurs_runtime.Value {
	once_diag__gopurs_runtime_Value_2627963987.Do(func() {
		cache_diag__gopurs_runtime_Value_2627963987 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diag__gopurs_runtime_Value_2627963987(a_0_box)
})
	})
	return cache_diag__gopurs_runtime_Value_2627963987
}

var cache_downFrom gopurs_runtime.Value
var once_downFrom sync.Once
func Get_downFrom() gopurs_runtime.Value {
	once_downFrom.Do(func() {
		cache_downFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_downFrom(dictEnum_0_box, dictUnfoldable_1_box)
})
	})
	return cache_downFrom
}

var cache_upFrom gopurs_runtime.Value
var once_upFrom sync.Once
func Get_upFrom() gopurs_runtime.Value {
	once_upFrom.Do(func() {
		cache_upFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_upFrom(dictEnum_0_box, dictUnfoldable_1_box)
})
	})
	return cache_upFrom
}

var cache_defaultToEnum gopurs_runtime.Value
var once_defaultToEnum sync.Once
func Get_defaultToEnum() gopurs_runtime.Value {
	once_defaultToEnum.Do(func() {
		cache_defaultToEnum = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultToEnum(dictBounded_0_box)
})
	})
	return cache_defaultToEnum
}

var cache_defaultSucc gopurs_runtime.Value
var once_defaultSucc sync.Once
func Get_defaultSucc() gopurs_runtime.Value {
	once_defaultSucc.Do(func() {
		cache_defaultSucc = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc
}

var cache_defaultSucc__gopurs_runtime_Value_1807708987 gopurs_runtime.Value
var once_defaultSucc__gopurs_runtime_Value_1807708987 sync.Once
func Get_defaultSucc__gopurs_runtime_Value_1807708987() gopurs_runtime.Value {
	once_defaultSucc__gopurs_runtime_Value_1807708987.Do(func() {
		cache_defaultSucc__gopurs_runtime_Value_1807708987 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__gopurs_runtime_Value_1807708987(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.IntVal))}
})
	})
	return cache_defaultSucc__gopurs_runtime_Value_1807708987
}

var cache_defaultSucc__gopurs_runtime_Value_166819931 gopurs_runtime.Value
var once_defaultSucc__gopurs_runtime_Value_166819931 sync.Once
func Get_defaultSucc__gopurs_runtime_Value_166819931() gopurs_runtime.Value {
	once_defaultSucc__gopurs_runtime_Value_166819931.Do(func() {
		cache_defaultSucc__gopurs_runtime_Value_166819931 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__gopurs_runtime_Value_166819931(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.StrVal()))}
})
	})
	return cache_defaultSucc__gopurs_runtime_Value_166819931
}

var cache_defaultSucc__gopurs_runtime_Value_346157499 gopurs_runtime.Value
var once_defaultSucc__gopurs_runtime_Value_346157499 sync.Once
func Get_defaultSucc__gopurs_runtime_Value_346157499() gopurs_runtime.Value {
	once_defaultSucc__gopurs_runtime_Value_346157499.Do(func() {
		cache_defaultSucc__gopurs_runtime_Value_346157499 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__gopurs_runtime_Value_346157499(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc__gopurs_runtime_Value_346157499
}

var cache_defaultPred gopurs_runtime.Value
var once_defaultPred sync.Once
func Get_defaultPred() gopurs_runtime.Value {
	once_defaultPred.Do(func() {
		cache_defaultPred = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred
}

var cache_defaultPred__gopurs_runtime_Value_1807708987 gopurs_runtime.Value
var once_defaultPred__gopurs_runtime_Value_1807708987 sync.Once
func Get_defaultPred__gopurs_runtime_Value_1807708987() gopurs_runtime.Value {
	once_defaultPred__gopurs_runtime_Value_1807708987.Do(func() {
		cache_defaultPred__gopurs_runtime_Value_1807708987 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__gopurs_runtime_Value_1807708987(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.IntVal))}
})
	})
	return cache_defaultPred__gopurs_runtime_Value_1807708987
}

var cache_defaultPred__gopurs_runtime_Value_166819931 gopurs_runtime.Value
var once_defaultPred__gopurs_runtime_Value_166819931 sync.Once
func Get_defaultPred__gopurs_runtime_Value_166819931() gopurs_runtime.Value {
	once_defaultPred__gopurs_runtime_Value_166819931.Do(func() {
		cache_defaultPred__gopurs_runtime_Value_166819931 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__gopurs_runtime_Value_166819931(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.StrVal()))}
})
	})
	return cache_defaultPred__gopurs_runtime_Value_166819931
}

var cache_defaultPred__gopurs_runtime_Value_346157499 gopurs_runtime.Value
var once_defaultPred__gopurs_runtime_Value_346157499 sync.Once
func Get_defaultPred__gopurs_runtime_Value_346157499() gopurs_runtime.Value {
	once_defaultPred__gopurs_runtime_Value_346157499.Do(func() {
		cache_defaultPred__gopurs_runtime_Value_346157499 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__gopurs_runtime_Value_346157499(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred__gopurs_runtime_Value_346157499
}

var cache_defaultFromEnum gopurs_runtime.Value
var once_defaultFromEnum sync.Once
func Get_defaultFromEnum() gopurs_runtime.Value {
	once_defaultFromEnum.Do(func() {
		cache_defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultFromEnum(dictEnum_0_box)
})
	})
	return cache_defaultFromEnum
}

var cache_defaultCardinality gopurs_runtime.Value
var once_defaultCardinality sync.Once
func Get_defaultCardinality() gopurs_runtime.Value {
	once_defaultCardinality.Do(func() {
		cache_defaultCardinality = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultCardinality(dictBounded_0_box)
})
	})
	return cache_defaultCardinality
}

var cache_charToEnum gopurs_runtime.Value
var once_charToEnum sync.Once
func Get_charToEnum() gopurs_runtime.Value {
	once_charToEnum.Do(func() {
		cache_charToEnum = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_charToEnum(v_0_box.IntVal))}
})
	})
	return cache_charToEnum
}

var cache_enumChar gopurs_runtime.Value
var once_enumChar sync.Once
func Get_enumChar() gopurs_runtime.Value {
	once_enumChar.Do(func() {
		cache_enumChar = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_charToEnum((gopurs_runtime.Apply(Get_toCharCode(), a_0).IntVal) - (1)))}.UnsafePtr))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_charToEnum((gopurs_runtime.Apply(Get_toCharCode(), a_0).IntVal) + (1)))}.UnsafePtr))}
}))
	})
	return cache_enumChar
}

var cache_cardinality gopurs_runtime.Value
var once_cardinality sync.Once
func Get_cardinality() gopurs_runtime.Value {
	once_cardinality.Do(func() {
		cache_cardinality = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_cardinality(dict_0_box))
})
	})
	return cache_cardinality
}

var cache_boundedEnumUnit gopurs_runtime.Value
var once_boundedEnumUnit sync.Once
func Get_boundedEnumUnit() gopurs_runtime.Value {
	once_boundedEnumUnit.Do(func() {
		cache_boundedEnumUnit = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedUnit()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumUnit()
}), gopurs_runtime.Int(1), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t0.UnsafePtr))}
}))
	})
	return cache_boundedEnumUnit
}

var cache_boundedEnumOrdering gopurs_runtime.Value
var once_boundedEnumOrdering sync.Once
func Get_boundedEnumOrdering() gopurs_runtime.Value {
	once_boundedEnumOrdering.Do(func() {
		cache_boundedEnumOrdering = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedOrdering()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumOrdering()
}), gopurs_runtime.Int(3), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (1) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (2) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t1.UnsafePtr))}
}))
	})
	return cache_boundedEnumOrdering
}

var cache_boundedEnumChar gopurs_runtime.Value
var once_boundedEnumChar sync.Once
func Get_boundedEnumChar() gopurs_runtime.Value {
	once_boundedEnumChar.Do(func() {
		cache_boundedEnumChar = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedChar()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumChar()
}), gopurs_runtime.Int((gopurs_runtime.Apply(Get_toCharCode(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")).IntVal) - (gopurs_runtime.Apply(Get_toCharCode(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")).IntVal)), Get_toCharCode(), Get_charToEnum())
	})
	return cache_boundedEnumChar
}

var cache_boundedEnumBoolean gopurs_runtime.Value
var once_boundedEnumBoolean sync.Once
func Get_boundedEnumBoolean() gopurs_runtime.Value {
	once_boundedEnumBoolean.Do(func() {
		cache_boundedEnumBoolean = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedBoolean()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumBoolean()
}), gopurs_runtime.Int(2), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((v_0.IntVal) != (0)) != (true) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) != (0) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(false)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (1) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(true)})}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[bool])(__t1.UnsafePtr))}
}))
	})
	return cache_boundedEnumBoolean
}

func Call_Cardinality(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_toEnum(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "toEnum")
}

func Call_toEnum__gopurs_runtime_Value_180026696(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "toEnum")
}

func Call_succ(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "succ")
}

func Call_succ__gopurs_runtime_Value_4187127373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "succ")
}

func Call_upFromIncluding(dictEnum_0_loop gopurs_runtime.Value, dictUnfoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 gopurs_runtime.Value = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_1, "unfoldr1"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), x_2)})}
}))
}

func Call_pred(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "pred")
}

func Call_pred__gopurs_runtime_Value_4187127373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "pred")
}

func Call_fromEnum(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "fromEnum")
}

func Call_fromEnum__gopurs_runtime_Value_2576611849(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "fromEnum")
}

func Call_toEnumWithDefaults(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), x_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), x_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), bottom2_1_0)).IntVal) != (0) {
__t3 = low_2
goto end_branch_3
} else {

}
}
{
__t3 = high_3
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
})
})
}

func Call_toEnumWithDefaults__gopurs_runtime_Value_3074683145(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom").StrVal()
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), x_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Str((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0.StrVal())
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), x_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), gopurs_runtime.Str(bottom2_1_0))).IntVal) != (0) {
__t3 = gopurs_runtime.Str(low_2.StrVal())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Str(high_3.StrVal())
}
end_branch_3:
__t2 = gopurs_runtime.Str(__t3.StrVal())
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2.StrVal())
})
})
})
}

func Call_toEnumWithDefaults__gopurs_runtime_Value_1689825737(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), x_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), x_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), bottom2_1_0)).IntVal) != (0) {
__t3 = low_2
goto end_branch_3
} else {

}
}
{
__t3 = high_3
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
})
})
}

func Call_enumTuple(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
ordTuple_1_0 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_ordTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordTuple_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
Bounded0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{})
_ = Bounded0_3_1
bottom2_4_2 := gopurs_runtime.RecordGet(Bounded0_3_1, "bottom")
_ = bottom2_4_2
Enum1_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_5_3
top2_6_4 := gopurs_runtime.RecordGet(Bounded0_3_1, "top")
_ = top2_6_4
ordTuple1_7_5 := gopurs_runtime.Apply(ordTuple_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_5_3, "Ord0"), gopurs_runtime.Value{}))
_ = ordTuple1_7_5
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple1_7_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, top2_6_4})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_9_6
__local_var_10_7 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
_ = __local_var_10_7
__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_5_3, "pred"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)
_ = __local_var_11_8
var __t9 gopurs_runtime.Value
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 930809136 && __local_var_11_8.UnsafePtr == nil) {
__t9 = __local_var_9_6
goto end_branch_9
} else {

}
}
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 930809136 && __local_var_11_8.UnsafePtr != nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_10_7, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_11_8.UnsafePtr).V0)})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t9.UnsafePtr))}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, bottom2_4_2})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_9_10
__local_var_10_11 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
_ = __local_var_10_11
__local_var_11_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_5_3, "succ"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1)
_ = __local_var_11_12
var __t13 gopurs_runtime.Value
{
if (__local_var_11_12.Type == 9 && __local_var_11_12.IntVal == 930809136 && __local_var_11_12.UnsafePtr == nil) {
__t13 = __local_var_9_10
goto end_branch_13
} else {

}
}
{
if (__local_var_11_12.Type == 9 && __local_var_11_12.IntVal == 930809136 && __local_var_11_12.UnsafePtr != nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_10_11, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_11_12.UnsafePtr).V0)})}
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t13.UnsafePtr))}
}))
})
}

func Call_enumMaybe(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
Enum1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_3_3
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_5
eqMaybe1_4_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr == nil) {
var __t7 gopurs_runtime.Value
{
if (y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Bool(false)
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr != nil)) && ((y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr != nil)) {
__t6 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(false)
}
end_branch_6:
return gopurs_runtime.Bool((__t6.IntVal) != (0))
})
}))
_ = eqMaybe1_4_4
ordMaybe_3_2 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_4_4
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr == nil) {
var __t9 gopurs_runtime.Value
{
if (y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr != nil)) && ((y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr != nil)) {
__t8 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_6.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
}))
_ = ordMaybe_3_2
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_10
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "pred"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0)})}.UnsafePtr))}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]])(__t10.UnsafePtr))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, bottom2_1_0})}})}.UnsafePtr))}
goto end_branch_11
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "succ"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0)).UnsafePtr))}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]])(__t11.UnsafePtr))}
}))
}

func Call_enumFromTo(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_1_0
return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}), "eq"), v_3, v1_4).IntVal) != (0) {
__t3 = gopurs_runtime.Apply3(pkg_Data_Unfoldable1.Get_replicate1(), dictUnfoldable1_2, gopurs_runtime.Int(1), v_3)
goto end_branch_3
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), v_3, v1_4)
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1527465420) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 380165415) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(true)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), __t6)).UnsafePtr))}
}))})}
}), v_3)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), __t1)).UnsafePtr))}
}))})}
}), v_3)
}
end_branch_3:
return __t3
})
})
})
}

func Call_enumFromTo__gopurs_runtime_Value_631221758(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_1_0
return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}), "eq"), v_3, v1_4).IntVal) != (0) {
__t3 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Unfoldable1.Get_replicate1(), dictUnfoldable1_2, gopurs_runtime.Int(1), v_3).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_3
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), v_3, v1_4)
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1527465420) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 380165415) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(true)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), __t6)).UnsafePtr))}
}))})}
}), v_3).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), __t1)).UnsafePtr))}
}))})}
}), v_3).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}
end_branch_3:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
})
})
}

func Call_enumFromTo__gopurs_runtime_Value_1236426744(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_1_0
return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}), "eq"), v_3, v1_4).IntVal) != (0) {
__t3 = gopurs_runtime.Apply3(pkg_Data_Unfoldable1.Get_replicate1(), dictUnfoldable1_2, gopurs_runtime.Int(1), v_3)
goto end_branch_3
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), v_3, v1_4)
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1527465420) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 380165415) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(true)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), __t6)).UnsafePtr))}
}))})}
}), v_3)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), __t1)).UnsafePtr))}
}))})}
}), v_3)
}
end_branch_3:
return __t3
})
})
})
}

func Call_enumFromThenTo(dictUnfoldable_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value, dictBoundedEnum_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value, b_4_loop gopurs_runtime.Value, c_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
var dictBoundedEnum_2 gopurs_runtime.Value = dictBoundedEnum_2_loop
_ = dictBoundedEnum_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var c_5 gopurs_runtime.Value = c_5_loop
_ = c_5
a_prime_6_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "fromEnum"), a_3)
_ = a_prime_6_0
__local_var_7_3 := (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "fromEnum"), b_4).IntVal) - (a_prime_6_0.IntVal)
_ = __local_var_7_3
__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "fromEnum"), c_5)
_ = __local_var_8_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_1, "map"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "toEnum"), x_7)
_ = __local_var_8_1
var __t2 gopurs_runtime.Value
{
if (__local_var_8_1.Type == 9 && __local_var_8_1.IntVal == 930809136 && __local_var_8_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_8_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), e_9, __local_var_8_4).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_9, gopurs_runtime.Int((e_9.IntVal) + (__local_var_7_3))})}})}.UnsafePtr))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t5.UnsafePtr))}
}), a_prime_6_0))
}

func Call_enumEither(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
Enum1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_1_0
top2_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "top")
_ = top2_2_1
ordEither_3_2 := gopurs_runtime.Apply(pkg_Data_Either.Get_ordEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither_3_2
return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_5_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_5_3
Enum11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{})
_ = Enum11_6_4
ordEither1_7_5 := gopurs_runtime.Apply(ordEither_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither1_7_5
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither1_7_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "pred"), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
_ = __local_var_9_7
var __t8 gopurs_runtime.Value
{
if (__local_var_9_7.Type == 9 && __local_var_9_7.IntVal == 930809136 && __local_var_9_7.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
if (__local_var_9_7.Type == 9 && __local_var_9_7.IntVal == 930809136 && __local_var_9_7.UnsafePtr != nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_9_7.UnsafePtr).V0})}})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t8.UnsafePtr))}
goto end_branch_6
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "pred"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
_ = __local_var_9_9
var __t10 gopurs_runtime.Value
{
if (__local_var_9_9.Type == 9 && __local_var_9_9.IntVal == 930809136 && __local_var_9_9.UnsafePtr == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, top2_2_1})}})}
goto end_branch_10
} else {

}
}
{
if (__local_var_9_9.Type == 9 && __local_var_9_9.IntVal == 930809136 && __local_var_9_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_9_9.UnsafePtr).V0})}})}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t10.UnsafePtr))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t6.UnsafePtr))}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
__local_var_9_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "succ"), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
_ = __local_var_9_12
var __t13 gopurs_runtime.Value
{
if (__local_var_9_12.Type == 9 && __local_var_9_12.IntVal == 930809136 && __local_var_9_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, bottom2_5_3})}})}
goto end_branch_13
} else {

}
}
{
if (__local_var_9_12.Type == 9 && __local_var_9_12.IntVal == 930809136 && __local_var_9_12.UnsafePtr != nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_9_12.UnsafePtr).V0})}})}
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t13.UnsafePtr))}
goto end_branch_11
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "succ"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)
_ = __local_var_9_14
var __t15 gopurs_runtime.Value
{
if (__local_var_9_14.Type == 9 && __local_var_9_14.IntVal == 930809136 && __local_var_9_14.UnsafePtr == nil) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_15
} else {

}
}
{
if (__local_var_9_14.Type == 9 && __local_var_9_14.IntVal == 930809136 && __local_var_9_14.UnsafePtr != nil) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_9_14.UnsafePtr).V0})}})}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t15.UnsafePtr))}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t11.UnsafePtr))}
}))
})
}

func Call_downFromIncluding(dictEnum_0_loop gopurs_runtime.Value, dictUnfoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 gopurs_runtime.Value = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_1, "unfoldr1"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), x_2)})}
}))
}

func Call_diag(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, a_0})}
}

func Call_diag__gopurs_runtime_Value_2627963987(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, a_0})}
}

func Call_downFrom(dictEnum_0_loop gopurs_runtime.Value, dictUnfoldable_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 gopurs_runtime.Value = dictUnfoldable_1_loop
_ = dictUnfoldable_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_diag())
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_1, "unfoldr"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), x_3))
}))
}

func Call_upFrom(dictEnum_0_loop gopurs_runtime.Value, dictUnfoldable_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 gopurs_runtime.Value = dictUnfoldable_1_loop
_ = dictUnfoldable_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_diag())
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_1, "unfoldr"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), x_3))
}))
}

func Call_defaultToEnum(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
bottom2_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_0 gopurs_runtime.Value
go__go_4_1_0 = gopurs_runtime.Func(func(i_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_5_loop gopurs_runtime.Value = i_5_loop_val
var x_6_loop gopurs_runtime.Value = x_6_loop_val
go__go_4_1_0:
for {
if false { continue go__go_4_1_0 }
var i_5 gopurs_runtime.Value = i_5_loop
_ = i_5
var x_6 gopurs_runtime.Value = x_6_loop
_ = x_6
var __t4 gopurs_runtime.Value
{
if (i_5.IntVal) == (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_6})}.UnsafePtr))}
goto end_branch_4
} else {

}
}
{
v_7_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_6)
_ = v_7_2
var __t3 gopurs_runtime.Value
{
if (v_7_2.Type == 9 && v_7_2.IntVal == 930809136 && v_7_2.UnsafePtr != nil) {
i_5_loop = gopurs_runtime.Int((i_5.IntVal) - (1))
x_6_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7_2.UnsafePtr).V0
continue go__go_4_1_0
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
if (v_7_2.Type == 9 && v_7_2.IntVal == 930809136 && v_7_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t3.UnsafePtr))}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t4.UnsafePtr))}
}
}()
})
})
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), i_prime_3, gopurs_runtime.Int(0)).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(go__go_4_1_0, i_prime_3, bottom2_1_0).UnsafePtr))}
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t5.UnsafePtr))}
})
})
}

func Call_defaultSucc(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) + (1))).UnsafePtr)
}

func Call_defaultSucc__gopurs_runtime_Value_1807708987(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 int64 = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Int(a_2)).IntVal) + (1))).UnsafePtr)
}

func Call_defaultSucc__gopurs_runtime_Value_166819931(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop string) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 string = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Str(a_2)).IntVal) + (1))).UnsafePtr)
}

func Call_defaultSucc__gopurs_runtime_Value_346157499(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) + (1))).UnsafePtr)
}

func Call_defaultPred(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) - (1))).UnsafePtr)
}

func Call_defaultPred__gopurs_runtime_Value_1807708987(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 int64 = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Int(a_2)).IntVal) - (1))).UnsafePtr)
}

func Call_defaultPred__gopurs_runtime_Value_166819931(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop string) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 string = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Str(a_2)).IntVal) - (1))).UnsafePtr)
}

func Call_defaultPred__gopurs_runtime_Value_346157499(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) - (1))).UnsafePtr)
}

func Call_defaultFromEnum(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_2_loop int64 = i_2_loop_val.IntVal
var x_3_loop gopurs_runtime.Value = x_3_loop_val
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var i_2 int64 = i_2_loop
_ = i_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), x_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr != nil) {
i_2_loop = (i_2) + (1)
x_3_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
continue go__go_1_0_1
__t2 = gopurs_runtime.Int(gopurs_runtime.Value{}.IntVal)
goto end_branch_2
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(i_2)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Int(__t2.IntVal)
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_1, gopurs_runtime.Int(0))
}

func Call_defaultCardinality(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
bottom2_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_2 gopurs_runtime.Value
go__go_3_1_2 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_4_loop int64 = i_4_loop_val.IntVal
var x_5_loop gopurs_runtime.Value = x_5_loop_val
go__go_3_1_2:
for {
if false { continue go__go_3_1_2 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
v_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5)
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 930809136 && v_6_2.UnsafePtr != nil) {
i_4_loop = (i_4) + (1)
x_5_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6_2.UnsafePtr).V0
continue go__go_3_1_2
__t3 = gopurs_runtime.Int(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 930809136 && v_6_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(i_4)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Int(__t3.IntVal)
}
}()
})
})
return gopurs_runtime.Int(gopurs_runtime.Apply2(go__go_3_1_2, gopurs_runtime.Int(1), bottom2_1_0).IntVal)
})
}

func Call_charToEnum(v_0_loop int64) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThanOrEq(), gopurs_runtime.Int(v_0), gopurs_runtime.Apply(Get_toCharCode(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"))), gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(v_0), gopurs_runtime.Apply(Get_toCharCode(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[string])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(v_0))})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[string])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t0.UnsafePtr)
}

func Call_cardinality(dict_0_loop gopurs_runtime.Value) int64 {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "cardinality").IntVal
}

func Get_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_FromCharCode
}

func Get_toCharCode() gopurs_runtime.Value {
	return _Gopurs_ToCharCode
}
