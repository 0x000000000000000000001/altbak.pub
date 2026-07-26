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
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
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
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
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
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415)) != (true))
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
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420)) != (true))
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
return Call_toEnum((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_toEnum
}

var cache_succ gopurs_runtime.Value
var once_succ sync.Once
func Get_succ() gopurs_runtime.Value {
	once_succ.Do(func() {
		cache_succ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_succ
}

var cache_upFromIncluding gopurs_runtime.Value
var once_upFromIncluding sync.Once
func Get_upFromIncluding() gopurs_runtime.Value {
	once_upFromIncluding.Do(func() {
		cache_upFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_upFromIncluding((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr), (*Record_unfoldr1_gopurs_runtime_Value)(dictUnfoldable1_1_box.UnsafePtr))
})
	})
	return cache_upFromIncluding
}

var cache_showCardinality gopurs_runtime.Value
var once_showCardinality sync.Once
func Get_showCardinality() gopurs_runtime.Value {
	once_showCardinality.Do(func() {
		cache_showCardinality = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Cardinality "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showCardinality
}

var cache_pred gopurs_runtime.Value
var once_pred sync.Once
func Get_pred() gopurs_runtime.Value {
	once_pred.Do(func() {
		cache_pred = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pred((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_pred
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
return Call_fromEnum((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
})
	})
	return cache_fromEnum
}

var cache_toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		cache_toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dictBoundedEnum_0_box.UnsafePtr))
})
	})
	return cache_toEnumWithDefaults
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
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}))
	})
	return cache_enumUnit
}

var cache_enumTuple gopurs_runtime.Value
var once_enumTuple sync.Once
func Get_enumTuple() gopurs_runtime.Value {
	once_enumTuple.Do(func() {
		cache_enumTuple = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumTuple((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr))
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
	})
	return cache_enumOrdering
}

var cache_enumMaybe gopurs_runtime.Value
var once_enumMaybe sync.Once
func Get_enumMaybe() gopurs_runtime.Value {
	once_enumMaybe.Do(func() {
		cache_enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumMaybe((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dictBoundedEnum_0_box.UnsafePtr))
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int((n_0.IntVal) - (1))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), n_0, gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedInt(), "top")).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int((n_0.IntVal) + (1))})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
return __t1
}))
	})
	return cache_enumInt
}

var cache_enumFromTo gopurs_runtime.Value
var once_enumFromTo sync.Once
func Get_enumFromTo() gopurs_runtime.Value {
	once_enumFromTo.Do(func() {
		cache_enumFromTo = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromTo((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr))
})
	})
	return cache_enumFromTo
}

var cache_enumFromThenTo gopurs_runtime.Value
var once_enumFromThenTo sync.Once
func Get_enumFromThenTo() gopurs_runtime.Value {
	once_enumFromThenTo.Do(func() {
		cache_enumFromThenTo = gopurs_runtime.Func6(func(dictUnfoldable_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictBoundedEnum_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value, b_4_box gopurs_runtime.Value, c_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromThenTo((*Record_unfoldr_gopurs_runtime_Value)(dictUnfoldable_0_box.UnsafePtr), (*Record_map__gopurs_runtime_Value)(dictFunctor_1_box.UnsafePtr), (*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dictBoundedEnum_2_box.UnsafePtr), a_3_box, b_4_box, c_5_box)
})
	})
	return cache_enumFromThenTo
}

var cache_enumEither gopurs_runtime.Value
var once_enumEither sync.Once
func Get_enumEither() gopurs_runtime.Value {
	once_enumEither.Do(func() {
		cache_enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumEither((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dictBoundedEnum_0_box.UnsafePtr))
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Bool(false)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_0.IntVal) != (0)) != (true) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Bool(true)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
return __t1
}))
	})
	return cache_enumBoolean
}

var cache_downFromIncluding gopurs_runtime.Value
var once_downFromIncluding sync.Once
func Get_downFromIncluding() gopurs_runtime.Value {
	once_downFromIncluding.Do(func() {
		cache_downFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_downFromIncluding((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr), (*Record_unfoldr1_gopurs_runtime_Value)(dictUnfoldable1_1_box.UnsafePtr))
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

var cache_downFrom gopurs_runtime.Value
var once_downFrom sync.Once
func Get_downFrom() gopurs_runtime.Value {
	once_downFrom.Do(func() {
		cache_downFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_downFrom((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr), (*Record_unfoldr_gopurs_runtime_Value)(dictUnfoldable_1_box.UnsafePtr))
})
	})
	return cache_downFrom
}

var cache_upFrom gopurs_runtime.Value
var once_upFrom sync.Once
func Get_upFrom() gopurs_runtime.Value {
	once_upFrom.Do(func() {
		cache_upFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_upFrom((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr), (*Record_unfoldr_gopurs_runtime_Value)(dictUnfoldable_1_box.UnsafePtr))
})
	})
	return cache_upFrom
}

var cache_defaultToEnum gopurs_runtime.Value
var once_defaultToEnum sync.Once
func Get_defaultToEnum() gopurs_runtime.Value {
	once_defaultToEnum.Do(func() {
		cache_defaultToEnum = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultToEnum((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_0_box.UnsafePtr))
})
	})
	return cache_defaultToEnum
}

var cache_defaultSucc gopurs_runtime.Value
var once_defaultSucc sync.Once
func Get_defaultSucc() gopurs_runtime.Value {
	once_defaultSucc.Do(func() {
		cache_defaultSucc = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultSucc(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box)
})
	})
	return cache_defaultSucc
}

var cache_defaultPred gopurs_runtime.Value
var once_defaultPred sync.Once
func Get_defaultPred() gopurs_runtime.Value {
	once_defaultPred.Do(func() {
		cache_defaultPred = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultPred(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box)
})
	})
	return cache_defaultPred
}

var cache_defaultFromEnum gopurs_runtime.Value
var once_defaultFromEnum sync.Once
func Get_defaultFromEnum() gopurs_runtime.Value {
	once_defaultFromEnum.Do(func() {
		cache_defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultFromEnum((*Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value)(dictEnum_0_box.UnsafePtr))
})
	})
	return cache_defaultFromEnum
}

var cache_defaultCardinality gopurs_runtime.Value
var once_defaultCardinality sync.Once
func Get_defaultCardinality() gopurs_runtime.Value {
	once_defaultCardinality.Do(func() {
		cache_defaultCardinality = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultCardinality((*Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value)(dictBounded_0_box.UnsafePtr))
})
	})
	return cache_defaultCardinality
}

var cache_charToEnum gopurs_runtime.Value
var once_charToEnum sync.Once
func Get_charToEnum() gopurs_runtime.Value {
	once_charToEnum.Do(func() {
		cache_charToEnum = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_charToEnum(v_0_box.IntVal)
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
return gopurs_runtime.Apply(Get_charToEnum(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_toCharCode(), a_0).IntVal) - (1)))
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_charToEnum(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_toCharCode(), a_0).IntVal) + (1)))
}))
	})
	return cache_enumChar
}

var cache_cardinality gopurs_runtime.Value
var once_cardinality sync.Once
func Get_cardinality() gopurs_runtime.Value {
	once_cardinality.Do(func() {
		cache_cardinality = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cardinality((*Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value)(dict_0_box.UnsafePtr))
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{pkg_Data_Unit.Get_unit()})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
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
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (1) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (2) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
return __t1
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
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Bool(false)})}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (1) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Bool(true)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
return __t1
}))
	})
	return cache_boundedEnumBoolean
}

type Record_alt_gopurs_runtime_Value struct {
	alt gopurs_runtime.Value
}

type Record_ struct {
	
}

type Record_pure_gopurs_runtime_Value struct {
	pure gopurs_runtime.Value
}

type Record_apply_gopurs_runtime_Value struct {
	apply gopurs_runtime.Value
}

type Record_bipure_gopurs_runtime_Value struct {
	bipure gopurs_runtime.Value
}

type Record_biapply_gopurs_runtime_Value struct {
	biapply gopurs_runtime.Value
}

type Record_bind_gopurs_runtime_Value struct {
	bind gopurs_runtime.Value
}

type Record_discard_gopurs_runtime_Value struct {
	discard gopurs_runtime.Value
}

type Record_identity_gopurs_runtime_Value struct {
	identity gopurs_runtime.Value
}

type Record_ask_gopurs_runtime_Value struct {
	ask gopurs_runtime.Value
}

type Record_local_gopurs_runtime_Value struct {
	local gopurs_runtime.Value
}

type Record_peek_gopurs_runtime_Value_pos_gopurs_runtime_Value struct {
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}

type Record_track_gopurs_runtime_Value struct {
	track gopurs_runtime.Value
}

type Record_extract_gopurs_runtime_Value struct {
	extract gopurs_runtime.Value
}

type Record_extend_gopurs_runtime_Value struct {
	extend gopurs_runtime.Value
}

type Record_defer__gopurs_runtime_Value struct {
	defer_ gopurs_runtime.Value
}

type Record_callCC_gopurs_runtime_Value struct {
	callCC gopurs_runtime.Value
}

type Record_catchError_gopurs_runtime_Value struct {
	catchError gopurs_runtime.Value
}

type Record_throwError_gopurs_runtime_Value struct {
	throwError gopurs_runtime.Value
}

type Record_chooseBool_gopurs_runtime_Value_chooseFloat_gopurs_runtime_Value_chooseInt_gopurs_runtime_Value_resize_gopurs_runtime_Value_sized_gopurs_runtime_Value struct {
	chooseBool gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt gopurs_runtime.Value
	resize gopurs_runtime.Value
	sized gopurs_runtime.Value
}

type Record_foldMap1_gopurs_runtime_Value_foldl1_gopurs_runtime_Value_foldr1_gopurs_runtime_Value struct {
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}

type Record_append__gopurs_runtime_Value struct {
	append_ gopurs_runtime.Value
}

type Record_tailRecM_gopurs_runtime_Value struct {
	tailRecM gopurs_runtime.Value
}

type Record_unfoldr_gopurs_runtime_Value struct {
	unfoldr gopurs_runtime.Value
}

type Record_map__gopurs_runtime_Value struct {
	map_ gopurs_runtime.Value
}

type Record_state_gopurs_runtime_Value struct {
	state gopurs_runtime.Value
}

type Record_lift_gopurs_runtime_Value struct {
	lift gopurs_runtime.Value
}

type Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value struct {
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}

type Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value struct {
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}

type Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value struct {
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}

type Record_mempty_gopurs_runtime_Value struct {
	mempty gopurs_runtime.Value
}

type Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value struct {
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}

type Record_empty_gopurs_runtime_Value struct {
	empty gopurs_runtime.Value
}

type Record_compose_gopurs_runtime_Value struct {
	compose gopurs_runtime.Value
}

type Record_eq_gopurs_runtime_Value struct {
	eq gopurs_runtime.Value
}

type Record_compare_gopurs_runtime_Value struct {
	compare gopurs_runtime.Value
}

type Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value struct {
	bifoldMap gopurs_runtime.Value
	bifoldl gopurs_runtime.Value
	bifoldr gopurs_runtime.Value
}

type Record_bimap_gopurs_runtime_Value struct {
	bimap gopurs_runtime.Value
}

type Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value struct {
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}

type Record_genericBottom_prime_gopurs_runtime_Value struct {
	genericBottom_prime gopurs_runtime.Value
}

type Record_genericTop_prime_gopurs_runtime_Value struct {
	genericTop_prime gopurs_runtime.Value
}

type Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value struct {
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}

type Record_lose_gopurs_runtime_Value struct {
	lose gopurs_runtime.Value
}

type Record_choose_gopurs_runtime_Value struct {
	choose gopurs_runtime.Value
}

type Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value struct {
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}

type Record_divide_gopurs_runtime_Value struct {
	divide gopurs_runtime.Value
}

type Record_recip_gopurs_runtime_Value struct {
	recip gopurs_runtime.Value
}

type Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value struct {
	genericCardinality_prime gopurs_runtime.Value
	genericFromEnum_prime gopurs_runtime.Value
	genericToEnum_prime gopurs_runtime.Value
}

type Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value struct {
	genericPred_prime gopurs_runtime.Value
	genericSucc_prime gopurs_runtime.Value
}

type Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value struct {
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}

type Record_unfoldr1_gopurs_runtime_Value struct {
	unfoldr1 gopurs_runtime.Value
}

type Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value struct {
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}

type Record_genericEq_prime_gopurs_runtime_Value struct {
	genericEq_prime gopurs_runtime.Value
}

type Record_eq1_gopurs_runtime_Value struct {
	eq1 gopurs_runtime.Value
}

type Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value struct {
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff gopurs_runtime.Value
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt gopurs_runtime.Value
}

type Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value struct {
	add gopurs_runtime.Value
	mul gopurs_runtime.Value
	one gopurs_runtime.Value
	zero gopurs_runtime.Value
}

type Record_foldMapWithIndex_gopurs_runtime_Value_foldlWithIndex_gopurs_runtime_Value_foldrWithIndex_gopurs_runtime_Value struct {
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}

type Record_cmap_gopurs_runtime_Value struct {
	cmap gopurs_runtime.Value
}

type Record_imap_gopurs_runtime_Value struct {
	imap gopurs_runtime.Value
}

type Record_mapWithIndex_gopurs_runtime_Value struct {
	mapWithIndex gopurs_runtime.Value
}

type Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value struct {
	from gopurs_runtime.Value
	to gopurs_runtime.Value
}

type Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value struct {
	genericConj_prime gopurs_runtime.Value
	genericDisj_prime gopurs_runtime.Value
	genericFF_prime gopurs_runtime.Value
	genericImplies_prime gopurs_runtime.Value
	genericNot_prime gopurs_runtime.Value
	genericTT_prime gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_bool_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_bool struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff bool
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt bool
}

type Record_genericMempty_prime_gopurs_runtime_Value struct {
	genericMempty_prime gopurs_runtime.Value
}

type Record_genericCompare_prime_gopurs_runtime_Value struct {
	genericCompare_prime gopurs_runtime.Value
}

type Record_sub_gopurs_runtime_Value struct {
	sub gopurs_runtime.Value
}

type Record_compare1_gopurs_runtime_Value struct {
	compare1 gopurs_runtime.Value
}

type Record_left_gopurs_runtime_Value_right_gopurs_runtime_Value struct {
	left gopurs_runtime.Value
	right gopurs_runtime.Value
}

type Record_first_gopurs_runtime_Value_second_gopurs_runtime_Value struct {
	first gopurs_runtime.Value
	second gopurs_runtime.Value
}

type Record_dimap_gopurs_runtime_Value struct {
	dimap gopurs_runtime.Value
}

type Record_genericSub_prime_gopurs_runtime_Value struct {
	genericSub_prime gopurs_runtime.Value
}

type Record_genericAppend_prime_gopurs_runtime_Value struct {
	genericAppend_prime gopurs_runtime.Value
}

type Record_sequence1_gopurs_runtime_Value_traverse1_gopurs_runtime_Value struct {
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}

type Record_genericAdd_prime_gopurs_runtime_Value_genericMul_prime_gopurs_runtime_Value_genericOne_prime_gopurs_runtime_Value_genericZero_prime_gopurs_runtime_Value struct {
	genericAdd_prime gopurs_runtime.Value
	genericMul_prime gopurs_runtime.Value
	genericOne_prime gopurs_runtime.Value
	genericZero_prime gopurs_runtime.Value
}

type Record_genericShow_prime_gopurs_runtime_Value struct {
	genericShow_prime gopurs_runtime.Value
}

type Record_genericShowArgs_gopurs_runtime_Value struct {
	genericShowArgs gopurs_runtime.Value
}

type Record_show_gopurs_runtime_Value struct {
	show gopurs_runtime.Value
}

type Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value struct {
	fromDuration gopurs_runtime.Value
	toDuration gopurs_runtime.Value
}

type Record_traverseWithIndex_gopurs_runtime_Value struct {
	traverseWithIndex gopurs_runtime.Value
}

type Record_liftEffect_gopurs_runtime_Value struct {
	liftEffect gopurs_runtime.Value
}

type Record_mappend__gopurs_runtime_Value_mempty__gopurs_runtime_Value struct {
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}

type Record_proof_gopurs_runtime_Value struct {
	proof gopurs_runtime.Value
}

type Record_lower_gopurs_runtime_Value struct {
	lower gopurs_runtime.Value
}

type Record_liftST_gopurs_runtime_Value struct {
	liftST gopurs_runtime.Value
}

type Record_tell_gopurs_runtime_Value struct {
	tell gopurs_runtime.Value
}

type Record_reflectSymbol_gopurs_runtime_Value struct {
	reflectSymbol gopurs_runtime.Value
}

type Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value struct {
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}

type Record_conquer_gopurs_runtime_Value struct {
	conquer gopurs_runtime.Value
}

type Record_inj_gopurs_runtime_Value_prj_gopurs_runtime_Value struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}

type Record_eqRecord_gopurs_runtime_Value struct {
	eqRecord gopurs_runtime.Value
}

type Record_conjRecord_gopurs_runtime_Value_disjRecord_gopurs_runtime_Value_ffRecord_gopurs_runtime_Value_impliesRecord_gopurs_runtime_Value_notRecord_gopurs_runtime_Value_ttRecord_gopurs_runtime_Value struct {
	conjRecord gopurs_runtime.Value
	disjRecord gopurs_runtime.Value
	ffRecord gopurs_runtime.Value
	impliesRecord gopurs_runtime.Value
	notRecord gopurs_runtime.Value
	ttRecord gopurs_runtime.Value
}

type Record_memptyRecord_gopurs_runtime_Value struct {
	memptyRecord gopurs_runtime.Value
}

type Record_compareRecord_gopurs_runtime_Value struct {
	compareRecord gopurs_runtime.Value
}

type Record_closed_gopurs_runtime_Value struct {
	closed gopurs_runtime.Value
}

type Record_unleft_gopurs_runtime_Value_unright_gopurs_runtime_Value struct {
	unleft gopurs_runtime.Value
	unright gopurs_runtime.Value
}

type Record_unfirst_gopurs_runtime_Value_unsecond_gopurs_runtime_Value struct {
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}

type Record_reflectType_gopurs_runtime_Value struct {
	reflectType gopurs_runtime.Value
}

type Record_subRecord_gopurs_runtime_Value struct {
	subRecord gopurs_runtime.Value
}

type Record_appendRecord_gopurs_runtime_Value struct {
	appendRecord gopurs_runtime.Value
}

type Record_addRecord_gopurs_runtime_Value_mulRecord_gopurs_runtime_Value_oneRecord_gopurs_runtime_Value_zeroRecord_gopurs_runtime_Value struct {
	addRecord gopurs_runtime.Value
	mulRecord gopurs_runtime.Value
	oneRecord gopurs_runtime.Value
	zeroRecord gopurs_runtime.Value
}

type Record_showRecordFields_gopurs_runtime_Value struct {
	showRecordFields gopurs_runtime.Value
}

type Record_nes_gopurs_runtime_Value struct {
	nes gopurs_runtime.Value
}

type Record_liftAff_gopurs_runtime_Value struct {
	liftAff gopurs_runtime.Value
}

func Call_Cardinality(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_toEnum(dict_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.toEnum
}

func Call_succ(dict_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.succ
}

func Call_upFromIncluding(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value, dictUnfoldable1_1_loop *Record_unfoldr1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 *Record_unfoldr1_gopurs_runtime_Value = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(dictUnfoldable1_1.unfoldr1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{x_2, gopurs_runtime.Apply(dictEnum_0.succ, x_2)})}
}))
}

func Call_pred(dict_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.pred
}

func Call_fromEnum(dict_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.fromEnum
}

func Call_toEnumWithDefaults(dictBoundedEnum_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBoundedEnum_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBoundedEnum_0)}, "Bounded0_NOT_FOUND"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func3(func(low_2 gopurs_runtime.Value, high_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(dictBoundedEnum_0.toEnum, x_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_5_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 3589588149) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), x_4, gopurs_runtime.Apply(dictBoundedEnum_0.fromEnum, bottom2_1_0)).IntVal) != (0) {
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
}

func Call_enumTuple(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
ordTuple_1_0 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_ordTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEnum_0)}, "Ord0_NOT_FOUND"), gopurs_runtime.Value{}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_9, top2_6_4})}
}), gopurs_runtime.Apply(dictEnum_0.pred, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_8.UnsafePtr).V0))
_ = __local_var_9_6
__local_var_10_7 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_8.UnsafePtr).V0)
_ = __local_var_10_7
__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_5_3, "pred"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_8.UnsafePtr).V1)
_ = __local_var_11_8
var __t9 gopurs_runtime.Value
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 3589588149) {
__t9 = __local_var_9_6
goto end_branch_9
} else {

}
}
{
if (__local_var_11_8.Type == 9 && __local_var_11_8.IntVal == 930809136) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_10_7, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_11_8.UnsafePtr).V0)})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_9, bottom2_4_2})}
}), gopurs_runtime.Apply(dictEnum_0.succ, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_8.UnsafePtr).V0))
_ = __local_var_9_10
__local_var_10_11 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_8.UnsafePtr).V0)
_ = __local_var_10_11
__local_var_11_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_5_3, "succ"), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_8.UnsafePtr).V1)
_ = __local_var_11_12
var __t13 gopurs_runtime.Value
{
if (__local_var_11_12.Type == 9 && __local_var_11_12.IntVal == 3589588149) {
__t13 = __local_var_9_10
goto end_branch_13
} else {

}
}
{
if (__local_var_11_12.Type == 9 && __local_var_11_12.IntVal == 930809136) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_10_11, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_11_12.UnsafePtr).V0)})}
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
})
}

func Call_enumMaybe(dictBoundedEnum_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBoundedEnum_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBoundedEnum_0)}, "Bounded0_NOT_FOUND"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
Enum1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBoundedEnum_0)}, "Enum1_NOT_FOUND"), gopurs_runtime.Value{})
_ = Enum1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_3
eqMaybe1_5_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (x_5.Type == 9 && x_5.IntVal == 3589588149) {
__t6 = gopurs_runtime.Bool((y_6.Type == 9 && y_6.IntVal == 3589588149))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(((x_5.Type == 9 && x_5.IntVal == 930809136)) && (((y_6.Type == 9 && y_6.IntVal == 930809136)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "eq"), (*pkg_Data_Maybe.Data_Data_Maybe_Just)(x_5.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(y_6.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_6:
return __t6
}))
_ = eqMaybe1_5_5
ordMaybe_5_4 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_5_5
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (x_6.Type == 9 && x_6.IntVal == 3589588149) {
var __t8 gopurs_runtime.Value
{
if (y_7.Type == 9 && y_7.IntVal == 3589588149) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
if (y_7.Type == 9 && y_7.IntVal == 3589588149) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_7
} else {

}
}
{
if ((x_6.Type == 9 && x_6.IntVal == 930809136)) && ((y_7.Type == 9 && y_7.IntVal == 930809136)) {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "compare"), (*pkg_Data_Maybe.Data_Data_Maybe_Just)(x_6.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(y_7.UnsafePtr).V0)
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
_ = ordMaybe_5_4
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3589588149) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_9
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "pred"), (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_6.UnsafePtr).V0)})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3589588149) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{bottom2_1_0})}})}
goto end_branch_10
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136) {
__t10 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "succ"), (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_6.UnsafePtr).V0))
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}))
}

func Call_enumFromTo(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEnum_0)}, "Ord0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Ord0_1_0
return gopurs_runtime.Func3(func(dictUnfoldable1_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}), "eq"), v_3, v1_4).IntVal) != (0) {
__t2 = gopurs_runtime.Apply3(pkg_Data_Unfoldable1.Get_replicate1(), dictUnfoldable1_2, gopurs_runtime.Int(1), v_3)
goto end_branch_2
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), v_3, v1_4)
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1527465420) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(dictEnum_0.succ, a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.Bool(((__t_tag_4.Type == 9 && __t_tag_4.IntVal == 380165415)) != (true))))
}))})}
}), v_3)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(dictEnum_0.pred, a_5), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_6, v1_4)
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_6
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420)) != (true))))
}))})}
}), v_3)
}
end_branch_2:
return __t2
})
}

func Call_enumFromThenTo(dictUnfoldable_0_loop *Record_unfoldr_gopurs_runtime_Value, dictFunctor_1_loop *Record_map__gopurs_runtime_Value, dictBoundedEnum_2_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value, a_3_loop gopurs_runtime.Value, b_4_loop gopurs_runtime.Value, c_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Record_unfoldr_gopurs_runtime_Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var dictFunctor_1 *Record_map__gopurs_runtime_Value = dictFunctor_1_loop
_ = dictFunctor_1
var dictBoundedEnum_2 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dictBoundedEnum_2_loop
_ = dictBoundedEnum_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var c_5 gopurs_runtime.Value = c_5_loop
_ = c_5
a_prime_6_0 := gopurs_runtime.Apply(dictBoundedEnum_2.fromEnum, a_3)
_ = a_prime_6_0
__local_var_7_3 := (gopurs_runtime.Apply(dictBoundedEnum_2.fromEnum, b_4).IntVal) - (a_prime_6_0.IntVal)
_ = __local_var_7_3
__local_var_8_4 := gopurs_runtime.Apply(dictBoundedEnum_2.fromEnum, c_5)
_ = __local_var_8_4
return gopurs_runtime.Apply2(dictFunctor_1.map_, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_1 := gopurs_runtime.Apply(dictBoundedEnum_2.toEnum, x_7)
_ = __local_var_8_1
var __t2 gopurs_runtime.Value
{
if (__local_var_8_1.Type == 9 && __local_var_8_1.IntVal == 930809136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_8_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(dictUnfoldable_0.unfoldr, gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), e_9, __local_var_8_4).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_9, gopurs_runtime.Int((e_9.IntVal) + (__local_var_7_3))})}})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_5:
return __t5
}), a_prime_6_0))
}

func Call_enumEither(dictBoundedEnum_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBoundedEnum_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
Enum1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBoundedEnum_0)}, "Enum1_NOT_FOUND"), gopurs_runtime.Value{})
_ = Enum1_1_0
top2_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictBoundedEnum_0)}, "Bounded0_NOT_FOUND"), gopurs_runtime.Value{}), "top")
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
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "pred"), (*pkg_Data_Either.Data_Data_Either_Left)(v_8.UnsafePtr).V0)
_ = __local_var_9_7
var __t8 gopurs_runtime.Value
{
if (__local_var_9_7.Type == 9 && __local_var_9_7.IntVal == 3589588149) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
if (__local_var_9_7.Type == 9 && __local_var_9_7.IntVal == 930809136) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_9_7.UnsafePtr).V0})}})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t6 = __t8
goto end_branch_6
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "pred"), (*pkg_Data_Either.Data_Data_Either_Right)(v_8.UnsafePtr).V0)
_ = __local_var_9_9
var __t10 gopurs_runtime.Value
{
if (__local_var_9_9.Type == 9 && __local_var_9_9.IntVal == 3589588149) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{top2_2_1})}})}
goto end_branch_10
} else {

}
}
{
if (__local_var_9_9.Type == 9 && __local_var_9_9.IntVal == 930809136) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_9_9.UnsafePtr).V0})}})}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t6 = __t10
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
__local_var_9_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "succ"), (*pkg_Data_Either.Data_Data_Either_Left)(v_8.UnsafePtr).V0)
_ = __local_var_9_12
var __t13 gopurs_runtime.Value
{
if (__local_var_9_12.Type == 9 && __local_var_9_12.IntVal == 3589588149) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{bottom2_5_3})}})}
goto end_branch_13
} else {

}
}
{
if (__local_var_9_12.Type == 9 && __local_var_9_12.IntVal == 930809136) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_9_12.UnsafePtr).V0})}})}
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "succ"), (*pkg_Data_Either.Data_Data_Either_Right)(v_8.UnsafePtr).V0)
_ = __local_var_9_14
var __t15 gopurs_runtime.Value
{
if (__local_var_9_14.Type == 9 && __local_var_9_14.IntVal == 3589588149) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_15
} else {

}
}
{
if (__local_var_9_14.Type == 9 && __local_var_9_14.IntVal == 930809136) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_9_14.UnsafePtr).V0})}})}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t11 = __t15
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}))
})
}

func Call_downFromIncluding(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value, dictUnfoldable1_1_loop *Record_unfoldr1_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 *Record_unfoldr1_gopurs_runtime_Value = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(dictUnfoldable1_1.unfoldr1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{x_2, gopurs_runtime.Apply(dictEnum_0.pred, x_2)})}
}))
}

func Call_diag(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_0, a_0})}
}

func Call_downFrom(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value, dictUnfoldable_1_loop *Record_unfoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 *Record_unfoldr_gopurs_runtime_Value = dictUnfoldable_1_loop
_ = dictUnfoldable_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_diag())
_ = __local_var_2_0
return gopurs_runtime.Apply(dictUnfoldable_1.unfoldr, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(dictEnum_0.pred, x_3))
}))
}

func Call_upFrom(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value, dictUnfoldable_1_loop *Record_unfoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 *Record_unfoldr_gopurs_runtime_Value = dictUnfoldable_1_loop
_ = dictUnfoldable_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_diag())
_ = __local_var_2_0
return gopurs_runtime.Apply(dictUnfoldable_1.unfoldr, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(dictEnum_0.succ, x_3))
}))
}

func Call_defaultToEnum(dictBounded_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBounded_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_0_loop
_ = dictBounded_0
bottom2_1_0 := dictBounded_0.bottom
_ = bottom2_1_0
return gopurs_runtime.Func2(func(dictEnum_2 gopurs_runtime.Value, i_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(i_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_5_loop gopurs_runtime.Value = i_5_loop_val
var x_6_loop gopurs_runtime.Value = x_6_loop_val
go__4_1:
for {
if false { continue go__4_1 }
var i_5 gopurs_runtime.Value = i_5_loop
_ = i_5
var x_6 gopurs_runtime.Value = x_6_loop
_ = x_6
var __t4 gopurs_runtime.Value
{
if (i_5.IntVal) == (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{x_6})}
goto end_branch_4
} else {

}
}
{
v_7_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_6)
_ = v_7_2
var __t3 gopurs_runtime.Value
{
if (v_7_2.Type == 9 && v_7_2.IntVal == 930809136) {
i_5_loop = gopurs_runtime.Int((i_5.IntVal) - (1))
x_6_loop = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_7_2.UnsafePtr).V0
continue go__4_1
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v_7_2.Type == 9 && v_7_2.IntVal == 3589588149) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
return __t4
}
}()
})
})
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), i_prime_3, gopurs_runtime.Int(0)).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(go__4_1, i_prime_3, bottom2_1_0)
}
end_branch_5:
return __t5
})
}

func Call_defaultSucc(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) + (1)))
}

func Call_defaultPred(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) - (1)))
}

func Call_defaultFromEnum(dictEnum_0_loop *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEnum_0 *Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value = dictEnum_0_loop
_ = dictEnum_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_2_loop gopurs_runtime.Value = i_2_loop_val
var x_3_loop gopurs_runtime.Value = x_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var i_2 gopurs_runtime.Value = i_2_loop
_ = i_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
v_4_1 := gopurs_runtime.Apply(dictEnum_0.pred, x_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136) {
i_2_loop = gopurs_runtime.Int((i_2.IntVal) + (1))
x_3_loop = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4_1.UnsafePtr).V0
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 3589588149) {
__t2 = i_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0))
}

func Call_defaultCardinality(dictBounded_0_loop *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value) gopurs_runtime.Value {
var dictBounded_0 *Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value = dictBounded_0_loop
_ = dictBounded_0
bottom2_1_0 := dictBounded_0.bottom
_ = bottom2_1_0
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_4_loop gopurs_runtime.Value = i_4_loop_val
var x_5_loop gopurs_runtime.Value = x_5_loop_val
go__3_1:
for {
if false { continue go__3_1 }
var i_4 gopurs_runtime.Value = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
v_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5)
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 930809136) {
i_4_loop = gopurs_runtime.Int((i_4.IntVal) + (1))
x_5_loop = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_6_2.UnsafePtr).V0
continue go__3_1
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 3589588149) {
__t3 = i_4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
return gopurs_runtime.Apply2(go__3_1, gopurs_runtime.Int(1), bottom2_1_0)
})
}

func Call_charToEnum(v_0_loop int64) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThanOrEq(), gopurs_runtime.Int(v_0), gopurs_runtime.Apply(Get_toCharCode(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"))), gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(v_0), gopurs_runtime.Apply(Get_toCharCode(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(v_0))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}

func Call_cardinality(dict_0_loop *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value) gopurs_runtime.Value {
var dict_0 *Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value = dict_0_loop
_ = dict_0
return dict_0.cardinality
}

func Get_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_FromCharCode
}

func Get_toCharCode() gopurs_runtime.Value {
	return _Gopurs_ToCharCode
}
