package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_bottom gopurs_runtime.Value
var once_Data_Enum_bottom sync.Once
func Get_Data_Enum_bottom() gopurs_runtime.Value {
	once_Data_Enum_bottom.Do(func() {
		cache_Data_Enum_bottom = gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "bottom").IntVal)
	})
	return cache_Data_Enum_bottom
}

var cache_Data_Enum_fromJust gopurs_runtime.Value
var once_Data_Enum_fromJust sync.Once
func Get_Data_Enum_fromJust() gopurs_runtime.Value {
	once_Data_Enum_fromJust.Do(func() {
		cache_Data_Enum_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box))
})
	})
	return cache_Data_Enum_fromJust
}

var cache_Data_Enum_bottom1 gopurs_runtime.Value
var once_Data_Enum_bottom1 sync.Once
func Get_Data_Enum_bottom1() gopurs_runtime.Value {
	once_Data_Enum_bottom1.Do(func() {
		cache_Data_Enum_bottom1 = gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal())
	})
	return cache_Data_Enum_bottom1
}

var cache_Data_Enum_top gopurs_runtime.Value
var once_Data_Enum_top sync.Once
func Get_Data_Enum_top() gopurs_runtime.Value {
	once_Data_Enum_top.Do(func() {
		cache_Data_Enum_top = gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal())
	})
	return cache_Data_Enum_top
}

var cache_Data_Enum_Enum_dollarDict gopurs_runtime.Value
var once_Data_Enum_Enum_dollarDict sync.Once
func Get_Data_Enum_Enum_dollarDict() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollarDict.Do(func() {
		cache_Data_Enum_Enum_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Enum_dollarDict(x_0_box)
})
	})
	return cache_Data_Enum_Enum_dollarDict
}

var cache_Data_Enum_Cardinality gopurs_runtime.Value
var once_Data_Enum_Cardinality sync.Once
func Get_Data_Enum_Cardinality() gopurs_runtime.Value {
	once_Data_Enum_Cardinality.Do(func() {
		cache_Data_Enum_Cardinality = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality(x_0_box)
})
	})
	return cache_Data_Enum_Cardinality
}

var cache_Data_Enum_BoundedEnum_dollarDict gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollarDict sync.Once
func Get_Data_Enum_BoundedEnum_dollarDict() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollarDict.Do(func() {
		cache_Data_Enum_BoundedEnum_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_BoundedEnum_dollarDict(x_0_box)
})
	})
	return cache_Data_Enum_BoundedEnum_dollarDict
}

var cache_Data_Enum_toEnum gopurs_runtime.Value
var once_Data_Enum_toEnum sync.Once
func Get_Data_Enum_toEnum() gopurs_runtime.Value {
	once_Data_Enum_toEnum.Do(func() {
		cache_Data_Enum_toEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_toEnum
}

var cache_Data_Enum_succ gopurs_runtime.Value
var once_Data_Enum_succ sync.Once
func Get_Data_Enum_succ() gopurs_runtime.Value {
	once_Data_Enum_succ.Do(func() {
		cache_Data_Enum_succ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dict_0_box))
})
	})
	return cache_Data_Enum_succ
}

var cache_Data_Enum_upFromIncluding gopurs_runtime.Value
var once_Data_Enum_upFromIncluding sync.Once
func Get_Data_Enum_upFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_upFromIncluding.Do(func() {
		cache_Data_Enum_upFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_upFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_1_box))
})
	})
	return cache_Data_Enum_upFromIncluding
}

var cache_Data_Enum_showCardinality gopurs_runtime.Value
var once_Data_Enum_showCardinality sync.Once
func Get_Data_Enum_showCardinality() gopurs_runtime.Value {
	once_Data_Enum_showCardinality.Do(func() {
		cache_Data_Enum_showCardinality = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Cardinality ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
}))
	})
	return cache_Data_Enum_showCardinality
}

var cache_Data_Enum_pred gopurs_runtime.Value
var once_Data_Enum_pred sync.Once
func Get_Data_Enum_pred() gopurs_runtime.Value {
	once_Data_Enum_pred.Do(func() {
		cache_Data_Enum_pred = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dict_0_box))
})
	})
	return cache_Data_Enum_pred
}

var cache_Data_Enum_ordCardinality gopurs_runtime.Value
var once_Data_Enum_ordCardinality sync.Once
func Get_Data_Enum_ordCardinality() gopurs_runtime.Value {
	once_Data_Enum_ordCardinality.Do(func() {
		cache_Data_Enum_ordCardinality = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_Data_Enum_ordCardinality
}

var cache_Data_Enum_newtypeCardinality gopurs_runtime.Value
var once_Data_Enum_newtypeCardinality sync.Once
func Get_Data_Enum_newtypeCardinality() gopurs_runtime.Value {
	once_Data_Enum_newtypeCardinality.Do(func() {
		cache_Data_Enum_newtypeCardinality = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Enum_newtypeCardinality
}

var cache_Data_Enum_fromEnum gopurs_runtime.Value
var once_Data_Enum_fromEnum sync.Once
func Get_Data_Enum_fromEnum() gopurs_runtime.Value {
	once_Data_Enum_fromEnum.Do(func() {
		cache_Data_Enum_fromEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_fromEnum
}

var cache_Data_Enum_toEnumWithDefaults gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults sync.Once
func Get_Data_Enum_toEnumWithDefaults() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults.Do(func() {
		cache_Data_Enum_toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnumWithDefaults(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dictBoundedEnum_0_box))
})
	})
	return cache_Data_Enum_toEnumWithDefaults
}

var cache_Data_Enum_eqCardinality gopurs_runtime.Value
var once_Data_Enum_eqCardinality sync.Once
func Get_Data_Enum_eqCardinality() gopurs_runtime.Value {
	once_Data_Enum_eqCardinality.Do(func() {
		cache_Data_Enum_eqCardinality = gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
	})
	return cache_Data_Enum_eqCardinality
}

var cache_Data_Enum_enumUnit gopurs_runtime.Value
var once_Data_Enum_enumUnit sync.Once
func Get_Data_Enum_enumUnit() gopurs_runtime.Value {
	once_Data_Enum_enumUnit.Do(func() {
		cache_Data_Enum_enumUnit = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordUnit()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
	})
	return cache_Data_Enum_enumUnit
}

var cache_Data_Enum_enumTuple gopurs_runtime.Value
var once_Data_Enum_enumTuple sync.Once
func Get_Data_Enum_enumTuple() gopurs_runtime.Value {
	once_Data_Enum_enumTuple.Do(func() {
		cache_Data_Enum_enumTuple = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumTuple(dictEnum_0_box)
})
	})
	return cache_Data_Enum_enumTuple
}

var cache_Data_Enum_enumOrdering gopurs_runtime.Value
var once_Data_Enum_enumOrdering sync.Once
func Get_Data_Enum_enumOrdering() gopurs_runtime.Value {
	once_Data_Enum_enumOrdering.Do(func() {
		cache_Data_Enum_enumOrdering = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordOrdering()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (uint32(v_0.IntVal) == 1527465420) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if (uint32(v_0.IntVal) == 1527465420) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}
goto end_branch_1
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}}
goto end_branch_1
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
	})
	return cache_Data_Enum_enumOrdering
}

var cache_Data_Enum_enumMaybe gopurs_runtime.Value
var once_Data_Enum_enumMaybe sync.Once
func Get_Data_Enum_enumMaybe() gopurs_runtime.Value {
	once_Data_Enum_enumMaybe.Do(func() {
		cache_Data_Enum_enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumMaybe(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_enumMaybe
}

var cache_Data_Enum_enumInt gopurs_runtime.Value
var once_Data_Enum_enumInt sync.Once
func Get_Data_Enum_enumInt() gopurs_runtime.Value {
	once_Data_Enum_enumInt.Do(func() {
		cache_Data_Enum_enumInt = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (n_0.IntVal) > (gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "bottom").IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((n_0.IntVal) - (1))}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t2 bool
{
if (n_0.IntVal) < (gopurs_runtime.RecordGet(Get_Data_Bounded_boundedInt(), "top").IntVal) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((n_0.IntVal) + (1))}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}))
	})
	return cache_Data_Enum_enumInt
}

var cache_Data_Enum_enumFromTo gopurs_runtime.Value
var once_Data_Enum_enumFromTo sync.Once
func Get_Data_Enum_enumFromTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromTo.Do(func() {
		cache_Data_Enum_enumFromTo = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box))
})
	})
	return cache_Data_Enum_enumFromTo
}

var cache_Data_Enum_enumFromThenTo gopurs_runtime.Value
var once_Data_Enum_enumFromThenTo sync.Once
func Get_Data_Enum_enumFromThenTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromThenTo.Do(func() {
		cache_Data_Enum_enumFromThenTo = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictBoundedEnum_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromThenTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dictBoundedEnum_2_box))
})
	})
	return cache_Data_Enum_enumFromThenTo
}

var cache_Data_Enum_enumEither gopurs_runtime.Value
var once_Data_Enum_enumEither sync.Once
func Get_Data_Enum_enumEither() gopurs_runtime.Value {
	once_Data_Enum_enumEither.Do(func() {
		cache_Data_Enum_enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumEither(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_enumEither
}

var cache_Data_Enum_enumBoolean gopurs_runtime.Value
var once_Data_Enum_enumBoolean sync.Once
func Get_Data_Enum_enumBoolean() gopurs_runtime.Value {
	once_Data_Enum_enumBoolean.Do(func() {
		cache_Data_Enum_enumBoolean = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordBoolean()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.IntVal) != (0) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Bool(false)}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if ((v_0.IntVal) != (0)) != (true) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Bool(true)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
	})
	return cache_Data_Enum_enumBoolean
}

var cache_Data_Enum_downFromIncluding gopurs_runtime.Value
var once_Data_Enum_downFromIncluding sync.Once
func Get_Data_Enum_downFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_downFromIncluding.Do(func() {
		cache_Data_Enum_downFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_downFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_1_box))
})
	})
	return cache_Data_Enum_downFromIncluding
}

var cache_Data_Enum_diag gopurs_runtime.Value
var once_Data_Enum_diag sync.Once
func Get_Data_Enum_diag() gopurs_runtime.Value {
	once_Data_Enum_diag.Do(func() {
		cache_Data_Enum_diag = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Enum_diag(a_0_box))}
})
	})
	return cache_Data_Enum_diag
}

var cache_Data_Enum_downFrom gopurs_runtime.Value
var once_Data_Enum_downFrom sync.Once
func Get_Data_Enum_downFrom() gopurs_runtime.Value {
	once_Data_Enum_downFrom.Do(func() {
		cache_Data_Enum_downFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_downFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_1_box))
})
	})
	return cache_Data_Enum_downFrom
}

var cache_Data_Enum_upFrom gopurs_runtime.Value
var once_Data_Enum_upFrom sync.Once
func Get_Data_Enum_upFrom() gopurs_runtime.Value {
	once_Data_Enum_upFrom.Do(func() {
		cache_Data_Enum_upFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_upFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_1_box))
})
	})
	return cache_Data_Enum_upFrom
}

var cache_Data_Enum_defaultToEnum gopurs_runtime.Value
var once_Data_Enum_defaultToEnum sync.Once
func Get_Data_Enum_defaultToEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultToEnum.Do(func() {
		cache_Data_Enum_defaultToEnum = gopurs_runtime.Func3(func(dictBounded_0_box gopurs_runtime.Value, dictEnum_1_box gopurs_runtime.Value, i_prime_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](dictBounded_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_1_box), i_prime_2_box.IntVal))}
})
	})
	return cache_Data_Enum_defaultToEnum
}

var cache_Data_Enum_defaultSucc gopurs_runtime.Value
var once_Data_Enum_defaultSucc sync.Once
func Get_Data_Enum_defaultSucc() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc.Do(func() {
		cache_Data_Enum_defaultSucc = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultSucc(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_Data_Enum_defaultSucc
}

var cache_Data_Enum_defaultPred gopurs_runtime.Value
var once_Data_Enum_defaultPred sync.Once
func Get_Data_Enum_defaultPred() gopurs_runtime.Value {
	once_Data_Enum_defaultPred.Do(func() {
		cache_Data_Enum_defaultPred = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultPred(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_Data_Enum_defaultPred
}

var cache_Data_Enum_defaultFromEnum gopurs_runtime.Value
var once_Data_Enum_defaultFromEnum sync.Once
func Get_Data_Enum_defaultFromEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultFromEnum.Do(func() {
		cache_Data_Enum_defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_defaultFromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box))
})
	})
	return cache_Data_Enum_defaultFromEnum
}

var cache_Data_Enum_defaultCardinality gopurs_runtime.Value
var once_Data_Enum_defaultCardinality sync.Once
func Get_Data_Enum_defaultCardinality() gopurs_runtime.Value {
	once_Data_Enum_defaultCardinality.Do(func() {
		cache_Data_Enum_defaultCardinality = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_defaultCardinality(dictBounded_0_box)
})
	})
	return cache_Data_Enum_defaultCardinality
}

var cache_Data_Enum_charToEnum gopurs_runtime.Value
var once_Data_Enum_charToEnum sync.Once
func Get_Data_Enum_charToEnum() gopurs_runtime.Value {
	once_Data_Enum_charToEnum.Do(func() {
		cache_Data_Enum_charToEnum = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_charToEnum(v_0_box.IntVal))}
})
	})
	return cache_Data_Enum_charToEnum
}

var cache_Data_Enum_enumChar gopurs_runtime.Value
var once_Data_Enum_enumChar sync.Once
func Get_Data_Enum_enumChar() gopurs_runtime.Value {
	once_Data_Enum_enumChar.Do(func() {
		cache_Data_Enum_enumChar = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordChar()
}), gopurs_runtime.Apply2(Get_Data_Enum_defaultPred__1581620096(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()), gopurs_runtime.Apply2(Get_Data_Enum_defaultSucc__1581620096(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()))
	})
	return cache_Data_Enum_enumChar
}

var cache_Data_Enum_cardinality gopurs_runtime.Value
var once_Data_Enum_cardinality sync.Once
func Get_Data_Enum_cardinality() gopurs_runtime.Value {
	once_Data_Enum_cardinality.Do(func() {
		cache_Data_Enum_cardinality = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_cardinality(dict_0_box)
})
	})
	return cache_Data_Enum_cardinality
}

var cache_Data_Enum_boundedEnumUnit gopurs_runtime.Value
var once_Data_Enum_boundedEnumUnit sync.Once
func Get_Data_Enum_boundedEnumUnit() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumUnit.Do(func() {
		cache_Data_Enum_boundedEnumUnit = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bounded_boundedUnit()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordUnit()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
}), gopurs_runtime.Int(1), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0.IntVal) == (0) {
__t0 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}))
	})
	return cache_Data_Enum_boundedEnumUnit
}

var cache_Data_Enum_boundedEnumOrdering gopurs_runtime.Value
var once_Data_Enum_boundedEnumOrdering sync.Once
func Get_Data_Enum_boundedEnumOrdering() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumOrdering.Do(func() {
		cache_Data_Enum_boundedEnumOrdering = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bounded_boundedOrdering()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordOrdering()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (uint32(v_1.IntVal) == 1527465420) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (uint32(v_1.IntVal) == 902936544) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (uint32(v_1.IntVal) == 380165415) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if (uint32(v_1.IntVal) == 1527465420) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}
goto end_branch_1
} else {

}
}
{
if (uint32(v_1.IntVal) == 902936544) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}}
goto end_branch_1
} else {

}
}
{
if (uint32(v_1.IntVal) == 380165415) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}), gopurs_runtime.Int(3), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 int64
{
if (uint32(v_0.IntVal) == 1527465420) {
__t2 = 0
goto end_branch_2
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t2 = 1
goto end_branch_2
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t2 = 2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_0.IntVal) == (0) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}}
goto end_branch_3
} else {

}
}
{
if (v_0.IntVal) == (1) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}
goto end_branch_3
} else {

}
}
{
if (v_0.IntVal) == (2) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}))
	})
	return cache_Data_Enum_boundedEnumOrdering
}

var cache_Data_Enum_boundedEnumChar gopurs_runtime.Value
var once_Data_Enum_boundedEnumChar sync.Once
func Get_Data_Enum_boundedEnumChar() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumChar.Do(func() {
		cache_Data_Enum_boundedEnumChar = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bounded_boundedChar()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordChar()
}), gopurs_runtime.Apply2(Get_Data_Enum_defaultPred__1581620096(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()), gopurs_runtime.Apply2(Get_Data_Enum_defaultSucc__1581620096(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()))
}), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal())).IntVal) - (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal())).IntVal)), Get_Data_Enum_toCharCode(), Get_Data_Enum_charToEnum())
	})
	return cache_Data_Enum_boundedEnumChar
}

var cache_Data_Enum_boundedEnumBoolean gopurs_runtime.Value
var once_Data_Enum_boundedEnumBoolean sync.Once
func Get_Data_Enum_boundedEnumBoolean() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumBoolean.Do(func() {
		cache_Data_Enum_boundedEnumBoolean = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bounded_boundedBoolean()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordBoolean()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_1.IntVal) != (0) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Bool(false)}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if ((v_1.IntVal) != (0)) != (true) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Bool(true)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}), gopurs_runtime.Int(2), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 int64
{
if ((v_0.IntVal) != (0)) != (true) {
__t2 = 0
goto end_branch_2
} else {

}
}
{
if (v_0.IntVal) != (0) {
__t2 = 1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_0.IntVal) == (0) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Bool(false)}
goto end_branch_3
} else {

}
}
{
if (v_0.IntVal) == (1) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Bool(true)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}))
	})
	return cache_Data_Enum_boundedEnumBoolean
}

var cache_Data_Enum_defaultPred__2391565248 gopurs_runtime.Value
var once_Data_Enum_defaultPred__2391565248 sync.Once
func Get_Data_Enum_defaultPred__2391565248() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__2391565248.Do(func() {
		cache_Data_Enum_defaultPred__2391565248 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultPred__2391565248(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.IntVal))}
})
	})
	return cache_Data_Enum_defaultPred__2391565248
}

var cache_Data_Enum_defaultPred__1581620096 gopurs_runtime.Value
var once_Data_Enum_defaultPred__1581620096 sync.Once
func Get_Data_Enum_defaultPred__1581620096() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__1581620096.Do(func() {
		cache_Data_Enum_defaultPred__1581620096 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultPred__1581620096(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.StrVal()))}
})
	})
	return cache_Data_Enum_defaultPred__1581620096
}

var cache_Data_Enum_defaultPred__2204581824 gopurs_runtime.Value
var once_Data_Enum_defaultPred__2204581824 sync.Once
func Get_Data_Enum_defaultPred__2204581824() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__2204581824.Do(func() {
		cache_Data_Enum_defaultPred__2204581824 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultPred__2204581824(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_Data_Enum_defaultPred__2204581824
}

var cache_Data_Enum_defaultSucc__2391565248 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__2391565248 sync.Once
func Get_Data_Enum_defaultSucc__2391565248() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__2391565248.Do(func() {
		cache_Data_Enum_defaultSucc__2391565248 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultSucc__2391565248(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.IntVal))}
})
	})
	return cache_Data_Enum_defaultSucc__2391565248
}

var cache_Data_Enum_defaultSucc__1581620096 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__1581620096 sync.Once
func Get_Data_Enum_defaultSucc__1581620096() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__1581620096.Do(func() {
		cache_Data_Enum_defaultSucc__1581620096 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultSucc__1581620096(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box.StrVal()))}
})
	})
	return cache_Data_Enum_defaultSucc__1581620096
}

var cache_Data_Enum_defaultSucc__2204581824 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__2204581824 sync.Once
func Get_Data_Enum_defaultSucc__2204581824() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__2204581824.Do(func() {
		cache_Data_Enum_defaultSucc__2204581824 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_defaultSucc__2204581824(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_Data_Enum_defaultSucc__2204581824
}

var cache_Data_Enum_diag__2627963987 gopurs_runtime.Value
var once_Data_Enum_diag__2627963987 sync.Once
func Get_Data_Enum_diag__2627963987() gopurs_runtime.Value {
	once_Data_Enum_diag__2627963987.Do(func() {
		cache_Data_Enum_diag__2627963987 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Enum_diag__2627963987(a_0_box))}
})
	})
	return cache_Data_Enum_diag__2627963987
}

var cache_Data_Enum_enumFromTo__1480115131 gopurs_runtime.Value
var once_Data_Enum_enumFromTo__1480115131 sync.Once
func Get_Data_Enum_enumFromTo__1480115131() gopurs_runtime.Value {
	once_Data_Enum_enumFromTo__1480115131.Do(func() {
		cache_Data_Enum_enumFromTo__1480115131 = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromTo__1480115131(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dictEnum_0_box))
})
	})
	return cache_Data_Enum_enumFromTo__1480115131
}

var cache_Data_Enum_fromEnum__2110387339 gopurs_runtime.Value
var once_Data_Enum_fromEnum__2110387339 sync.Once
func Get_Data_Enum_fromEnum__2110387339() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__2110387339.Do(func() {
		cache_Data_Enum_fromEnum__2110387339 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__2110387339(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_fromEnum__2110387339
}

var cache_Data_Enum_fromEnum__3599151655 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3599151655 sync.Once
func Get_Data_Enum_fromEnum__3599151655() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3599151655.Do(func() {
		cache_Data_Enum_fromEnum__3599151655 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3599151655(v_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__3599151655
}

var cache_Data_Enum_fromEnum__1606852103 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1606852103 sync.Once
func Get_Data_Enum_fromEnum__1606852103() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1606852103.Do(func() {
		cache_Data_Enum_fromEnum__1606852103 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__1606852103(__eta0_0_box)
})
	})
	return cache_Data_Enum_fromEnum__1606852103
}

var cache_Data_Enum_fromEnum__1637084359 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1637084359 sync.Once
func Get_Data_Enum_fromEnum__1637084359() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1637084359.Do(func() {
		cache_Data_Enum_fromEnum__1637084359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__1637084359(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_fromEnum__1637084359
}

var cache_Data_Enum_fromEnum__1196942535 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1196942535 sync.Once
func Get_Data_Enum_fromEnum__1196942535() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1196942535.Do(func() {
		cache_Data_Enum_fromEnum__1196942535 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__1196942535(uint32(v_0_box.IntVal)))
})
	})
	return cache_Data_Enum_fromEnum__1196942535
}

var cache_Data_Enum_pred__2914940949 gopurs_runtime.Value
var once_Data_Enum_pred__2914940949 sync.Once
func Get_Data_Enum_pred__2914940949() gopurs_runtime.Value {
	once_Data_Enum_pred__2914940949.Do(func() {
		cache_Data_Enum_pred__2914940949 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred__2914940949(__eta0_0_box)
})
	})
	return cache_Data_Enum_pred__2914940949
}

var cache_Data_Enum_pred__3199041328 gopurs_runtime.Value
var once_Data_Enum_pred__3199041328 sync.Once
func Get_Data_Enum_pred__3199041328() gopurs_runtime.Value {
	once_Data_Enum_pred__3199041328.Do(func() {
		cache_Data_Enum_pred__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred__3199041328(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dict_0_box))
})
	})
	return cache_Data_Enum_pred__3199041328
}

var cache_Data_Enum_pred__2010692236 gopurs_runtime.Value
var once_Data_Enum_pred__2010692236 sync.Once
func Get_Data_Enum_pred__2010692236() gopurs_runtime.Value {
	once_Data_Enum_pred__2010692236.Do(func() {
		cache_Data_Enum_pred__2010692236 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred__2010692236(__eta0_0_box)
})
	})
	return cache_Data_Enum_pred__2010692236
}

var cache_Data_Enum_succ__412946465 gopurs_runtime.Value
var once_Data_Enum_succ__412946465 sync.Once
func Get_Data_Enum_succ__412946465() gopurs_runtime.Value {
	once_Data_Enum_succ__412946465.Do(func() {
		cache_Data_Enum_succ__412946465 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__412946465(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dict_0_box))
})
	})
	return cache_Data_Enum_succ__412946465
}

var cache_Data_Enum_succ__2914940949 gopurs_runtime.Value
var once_Data_Enum_succ__2914940949 sync.Once
func Get_Data_Enum_succ__2914940949() gopurs_runtime.Value {
	once_Data_Enum_succ__2914940949.Do(func() {
		cache_Data_Enum_succ__2914940949 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__2914940949(__eta0_0_box)
})
	})
	return cache_Data_Enum_succ__2914940949
}

var cache_Data_Enum_succ__3199041328 gopurs_runtime.Value
var once_Data_Enum_succ__3199041328 sync.Once
func Get_Data_Enum_succ__3199041328() gopurs_runtime.Value {
	once_Data_Enum_succ__3199041328.Do(func() {
		cache_Data_Enum_succ__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__3199041328(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](dict_0_box))
})
	})
	return cache_Data_Enum_succ__3199041328
}

var cache_Data_Enum_succ__2010692236 gopurs_runtime.Value
var once_Data_Enum_succ__2010692236 sync.Once
func Get_Data_Enum_succ__2010692236() gopurs_runtime.Value {
	once_Data_Enum_succ__2010692236.Do(func() {
		cache_Data_Enum_succ__2010692236 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__2010692236(__eta0_0_box)
})
	})
	return cache_Data_Enum_succ__2010692236
}

var cache_Data_Enum_succ__2858180024 gopurs_runtime.Value
var once_Data_Enum_succ__2858180024 sync.Once
func Get_Data_Enum_succ__2858180024() gopurs_runtime.Value {
	once_Data_Enum_succ__2858180024.Do(func() {
		cache_Data_Enum_succ__2858180024 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_succ__2858180024(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))}
})
	})
	return cache_Data_Enum_succ__2858180024
}

var cache_Data_Enum_toEnum__2738794986 gopurs_runtime.Value
var once_Data_Enum_toEnum__2738794986 sync.Once
func Get_Data_Enum_toEnum__2738794986() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2738794986.Do(func() {
		cache_Data_Enum_toEnum__2738794986 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__2738794986(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_toEnum__2738794986
}

var cache_Data_Enum_toEnum__2203070892 gopurs_runtime.Value
var once_Data_Enum_toEnum__2203070892 sync.Once
func Get_Data_Enum_toEnum__2203070892() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2203070892.Do(func() {
		cache_Data_Enum_toEnum__2203070892 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__2203070892(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_toEnum__2203070892
}

var cache_Data_Enum_toEnum__2099864294 gopurs_runtime.Value
var once_Data_Enum_toEnum__2099864294 sync.Once
func Get_Data_Enum_toEnum__2099864294() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2099864294.Do(func() {
		cache_Data_Enum_toEnum__2099864294 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_toEnum__2099864294(n_0_box.IntVal))}
})
	})
	return cache_Data_Enum_toEnum__2099864294
}

var cache_Data_Enum_toEnum__3317293286 gopurs_runtime.Value
var once_Data_Enum_toEnum__3317293286 sync.Once
func Get_Data_Enum_toEnum__3317293286() gopurs_runtime.Value {
	once_Data_Enum_toEnum__3317293286.Do(func() {
		cache_Data_Enum_toEnum__3317293286 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__3317293286(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dict_0_box))
})
	})
	return cache_Data_Enum_toEnum__3317293286
}

var cache_Data_Enum_toEnum__2309750950 gopurs_runtime.Value
var once_Data_Enum_toEnum__2309750950 sync.Once
func Get_Data_Enum_toEnum__2309750950() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2309750950.Do(func() {
		cache_Data_Enum_toEnum__2309750950 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_toEnum__2309750950(v_0_box.IntVal))}
})
	})
	return cache_Data_Enum_toEnum__2309750950
}

var cache_Data_Enum_toEnum__2793813158 gopurs_runtime.Value
var once_Data_Enum_toEnum__2793813158 sync.Once
func Get_Data_Enum_toEnum__2793813158() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2793813158.Do(func() {
		cache_Data_Enum_toEnum__2793813158 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_toEnum__2793813158(v_0_box.IntVal))}
})
	})
	return cache_Data_Enum_toEnum__2793813158
}

var cache_Data_Enum_toEnumWithDefaults__3941305703 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__3941305703 sync.Once
func Get_Data_Enum_toEnumWithDefaults__3941305703() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__3941305703.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__3941305703 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnumWithDefaults__3941305703(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Enum_toEnumWithDefaults__3941305703
}

var cache_Data_Enum_toEnumWithDefaults__3558602759 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__3558602759 sync.Once
func Get_Data_Enum_toEnumWithDefaults__3558602759() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__3558602759.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__3558602759 = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnumWithDefaults__3558602759(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum](dictBoundedEnum_0_box))
})
	})
	return cache_Data_Enum_toEnumWithDefaults__3558602759
}

type Constructor_Data_Enum_Enum struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4075786298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Enum)(ptr)
		_ = c
		switch key {
		case "Ord0": return gopurs_runtime.Box(c.V0)
		case "pred": return gopurs_runtime.Box(c.V1)
		case "succ": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Enum: " + key)
		}
	}
}


type Constructor_Data_Enum_BoundedEnum struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 int64
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[287434377] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_BoundedEnum)(ptr)
		_ = c
		switch key {
		case "Bounded0": return gopurs_runtime.Box(c.V0)
		case "Enum1": return gopurs_runtime.Box(c.V1)
		case "cardinality": return gopurs_runtime.Box(c.V2)
		case "fromEnum": return gopurs_runtime.Box(c.V3)
		case "toEnum": return gopurs_runtime.Box(c.V4)
		default: panic("Key not found in dictionary Constructor_Data_Enum_BoundedEnum: " + key)
		}
	}
}


func Call_Data_Enum_fromJust(v_0_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Data_Enum_Enum_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Enum_Cardinality(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Enum_BoundedEnum_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Enum_toEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_Enum_succ(dict_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Enum_upFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum, dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_1.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), x_2)})}
}))
}

func Call_Data_Enum_pred(dict_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_fromEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Enum_toEnumWithDefaults(dictBoundedEnum_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dictBoundedEnum_0 *Constructor_Data_Enum_BoundedEnum = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): bottom2_1_0 -> gopurs_runtime.Value
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V0), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V4), gopurs_runtime.Int(x_4.IntVal))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t4 = (*Constructor_Data_Maybe_Just)(v_5_1.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (x_4.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V3), bottom2_1_0).IntVal) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = low_2
goto end_branch_3
} else {

}
}
{
__t3 = high_3
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
})
})
}

func Call_Data_Enum_enumTuple(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBoundedEnum_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bounded0_4_2 -> *Constructor_Data_Bounded_Bounded
Bounded0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_3, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_4_2
// TAST (Let): Enum1_5_3 -> *Constructor_Data_Enum_Enum
Enum1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_3, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_3, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): eqTuple2_7_6 -> gopurs_runtime.Value
eqTuple2_7_6 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*Constructor_Data_Tuple_Tuple)(x_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_9.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "eq"), (*Constructor_Data_Tuple_Tuple)(x_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_9.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqTuple2_7_6
// TAST (Let): ordTuple1_6_4 -> gopurs_runtime.Value
ordTuple1_6_4 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_7_6
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_10_8 -> gopurs_runtime.Value
v_10_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), (*Constructor_Data_Tuple_Tuple)(x_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_9.UnsafePtr).V0)
_ = v_10_8
var __t9 uint32
{
if (uint32(v_10_8.IntVal) == 1527465420) {
__t9 = 1527465420
goto end_branch_9
} else {

}
}
{
if (uint32(v_10_8.IntVal) == 380165415) {
__t9 = 380165415
goto end_branch_9
} else {

}
}
{
__t9 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_5, "compare"), (*Constructor_Data_Tuple_Tuple)(x_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_9.UnsafePtr).V1).IntVal)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
})
}))
_ = ordTuple1_6_4
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple1_6_4
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_11 -> gopurs_runtime.Value
__local_var_8_11 := gopurs_runtime.Box(Bounded0_4_2.V2)
_ = __local_var_8_11
// TAST (Let): __local_var_8_10 -> *Constructor_Data_Maybe_Just
__local_var_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, __local_var_8_11})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0)))}))
_ = __local_var_8_10
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0)
_ = __local_var_9_13
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_9_13, x_10)})}
})
_ = __local_var_9_12
// TAST (Let): __local_var_10_14 -> *Constructor_Data_Maybe_Just
__local_var_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_5_3.V1), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1))
_ = __local_var_10_14
var __t15 *Constructor_Data_Maybe_Just
{
if (__local_var_10_14 == nil) {
__t15 = __local_var_8_10
goto end_branch_15
} else {

}
}
{
if (__local_var_10_14 != nil) {
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(__local_var_9_12, (__local_var_10_14).V0))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t15)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_17 -> gopurs_runtime.Value
__local_var_8_17 := gopurs_runtime.Box(Bounded0_4_2.V1)
_ = __local_var_8_17
// TAST (Let): __local_var_8_16 -> *Constructor_Data_Maybe_Just
__local_var_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, __local_var_8_17})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0)))}))
_ = __local_var_8_16
// TAST (Let): __local_var_9_19 -> gopurs_runtime.Value
__local_var_9_19 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0)
_ = __local_var_9_19
// TAST (Let): __local_var_9_18 -> gopurs_runtime.Value
__local_var_9_18 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_9_19, x_10)})}
})
_ = __local_var_9_18
// TAST (Let): __local_var_10_20 -> *Constructor_Data_Maybe_Just
__local_var_10_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_5_3.V2), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1))
_ = __local_var_10_20
var __t21 *Constructor_Data_Maybe_Just
{
if (__local_var_10_20 == nil) {
__t21 = __local_var_8_16
goto end_branch_21
} else {

}
}
{
if (__local_var_10_20 != nil) {
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(__local_var_9_18, (__local_var_10_20).V0))
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_21:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t21)}
}))
})
}

func Call_Data_Enum_enumMaybe(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): Bounded0_1_0 -> *Constructor_Data_Bounded_Bounded
Bounded0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_1_0
// TAST (Let): Enum1_2_1 -> *Constructor_Data_Enum_Enum
Enum1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): eqMaybe1_4_4 -> gopurs_runtime.Value
eqMaybe1_4_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 bool
{
if (x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr == nil) {
var __t6 bool
{
if (y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr == nil) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr != nil)) && ((y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr != nil)) {
__t7 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "eq"), (*Constructor_Data_Maybe_Just)(x_5.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_6.UnsafePtr).V0).IntVal) != (0)
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
})
}))
_ = eqMaybe1_4_4
// TAST (Let): ordMaybe_3_2 -> gopurs_runtime.Value
ordMaybe_3_2 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_4_4
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 uint32
{
if (x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr == nil) {
var __t8 uint32
{
if (y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr == nil) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr == nil) {
__t9 = 380165415
goto end_branch_9
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 930809136 && x_5.UnsafePtr != nil)) && ((y_6.Type == 9 && y_6.IntVal == 930809136 && y_6.UnsafePtr != nil)) {
__t9 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "compare"), (*Constructor_Data_Maybe_Just)(x_5.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_6.UnsafePtr).V0).IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
})
}))
_ = ordMaybe_3_2
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t10 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_10
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_2_1.V1), (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0)))}}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Data_Maybe_Just
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Box(Bounded0_1_0.V1)})}}
goto end_branch_11
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_2_1.V2), (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0)))})))})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}
}))
}

func Call_Data_Enum_enumFromTo(dictEnum_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): Ord0_1_0 -> gopurs_runtime.Value
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V0), gopurs_runtime.Value{})
_ = Ord0_1_0
// TAST (Let): Eq0_2_1 -> *Constructor_Data_Eq_Eq
Eq0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}))
_ = Eq0_2_1
// TAST (Let): Ord01_3_2 -> *Constructor_Data_Ord_Ord
Ord01_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V0), gopurs_runtime.Value{}))
_ = Ord01_3_2
return gopurs_runtime.Func(func(dictUnfoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(Eq0_2_1.V0), v_5, v1_6).IntVal) != (0) {
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Tuple_Tuple
{
var __t6 bool
{
if (i_7.IntVal) > (0) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
if __t6 {
__t7 = &Constructor_Data_Tuple_Tuple{1, v_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_7
} else {

}
}
{
__t7 = &Constructor_Data_Tuple_Tuple{1, v_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_7.IntVal) - (1))})}}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t7)}
}), gopurs_runtime.Int(0))
goto end_branch_13
} else {

}
}
{
var __t9 bool
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(Ord01_3_2.V1), v_5, v1_6)
if (uint32(__t_tag_8.IntVal) == 1527465420) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
if __t9 {
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), a_7)))}, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Maybe_Just
{
var __t11 bool
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_8, v1_6)
if (uint32(__t_tag_10.IntVal) == 380165415) {
__t11 = false
goto end_branch_11
} else {

}
}
{
__t11 = true
}
end_branch_11:
if __t11 {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_applicativeMaybe(), "pure"), Get_Data_Unit_unit()))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(Get_Data_Maybe_plusMaybe(), "empty"))
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_8
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t12)})))}
}))))}})}
}), v_5)
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), a_7)))}, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_8, v1_6)
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_applicativeMaybe(), "pure"), Get_Data_Unit_unit()))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(Get_Data_Maybe_plusMaybe(), "empty"))
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_8
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})))}
}))))}})}
}), v_5)
}
end_branch_13:
return __t13
})
})
})
}

func Call_Data_Enum_enumFromThenTo(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable, dictFunctor_1_loop *Constructor_Data_Functor_Functor, dictBoundedEnum_2_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var dictFunctor_1 *Constructor_Data_Functor_Functor = dictFunctor_1_loop
_ = dictFunctor_1
var dictBoundedEnum_2 *Constructor_Data_Enum_BoundedEnum = dictBoundedEnum_2_loop
_ = dictBoundedEnum_2
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): a_prime_7_0 -> gopurs_runtime.Value
a_prime_7_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V3), a_4)
_ = a_prime_7_0
// TAST (Let): __local_var_8_3 -> int64
__local_var_8_3 := (gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V3), b_5).IntVal) - (a_prime_7_0.IntVal)
_ = __local_var_8_3
// TAST (Let): __local_var_9_4 -> int64
__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V3), c_6).IntVal
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_1.V0), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_1 -> gopurs_runtime.Value
__local_var_9_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V4), x_8)
_ = __local_var_9_1
var __t2 gopurs_runtime.Value
{
if (__local_var_9_1.Type == 9 && __local_var_9_1.IntVal == 930809136 && __local_var_9_1.UnsafePtr != nil) {
__t2 = (*Constructor_Data_Maybe_Just)(__local_var_9_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(e_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just
{
var __t5 bool
{
if (e_10.IntVal) > (__local_var_9_4) {
__t5 = false
goto end_branch_5
} else {

}
}
{
__t5 = true
}
end_branch_5:
if __t5 {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(e_10.IntVal), gopurs_runtime.Int((e_10.IntVal) + (__local_var_8_3))})}}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), gopurs_runtime.Int(a_prime_7_0.IntVal)))
})
})
})
}))
}

func Call_Data_Enum_enumEither(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): Enum1_1_0 -> *Constructor_Data_Enum_Enum
Enum1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_1_0
// TAST (Let): Bounded0_2_1 -> *Constructor_Data_Bounded_Bounded
Bounded0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_2_1
// TAST (Let): ordEither_3_2 -> gopurs_runtime.Value
ordEither_3_2 := gopurs_runtime.Apply(Get_Data_Either_ordEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))
_ = ordEither_3_2
return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bounded01_5_3 -> *Constructor_Data_Bounded_Bounded
Bounded01_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded01_5_3
// TAST (Let): Enum11_6_4 -> *Constructor_Data_Enum_Enum
Enum11_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}))
_ = Enum11_6_4
// TAST (Let): ordEither1_7_5 -> gopurs_runtime.Value
ordEither1_7_5 := gopurs_runtime.Apply(ordEither_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))
_ = ordEither1_7_5
return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither1_7_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
// TAST (Let): __local_var_9_6 -> *Constructor_Data_Maybe_Just
__local_var_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_1_0.V1), (*Constructor_Data_Either_Left)(v_8.UnsafePtr).V0))
_ = __local_var_9_6
var __t7 *Constructor_Data_Maybe_Just
{
if (__local_var_9_6 == nil) {
__t7 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_7
} else {

}
}
{
if (__local_var_9_6 != nil) {
__t7 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (__local_var_9_6).V0})}}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
__t10 = __t7
goto end_branch_10
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
// TAST (Let): __local_var_9_8 -> *Constructor_Data_Maybe_Just
__local_var_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum11_6_4.V1), (*Constructor_Data_Either_Right)(v_8.UnsafePtr).V0))
_ = __local_var_9_8
var __t9 *Constructor_Data_Maybe_Just
{
if (__local_var_9_8 == nil) {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Box(Bounded0_2_1.V2)})}}
goto end_branch_9
} else {

}
}
{
if (__local_var_9_8 != nil) {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (__local_var_9_8).V0})}}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 *Constructor_Data_Maybe_Just
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
// TAST (Let): __local_var_9_11 -> *Constructor_Data_Maybe_Just
__local_var_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_1_0.V2), (*Constructor_Data_Either_Left)(v_8.UnsafePtr).V0))
_ = __local_var_9_11
var __t12 *Constructor_Data_Maybe_Just
{
if (__local_var_9_11 == nil) {
__t12 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Box(Bounded01_5_3.V1)})}}
goto end_branch_12
} else {

}
}
{
if (__local_var_9_11 != nil) {
__t12 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (__local_var_9_11).V0})}}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
__t15 = __t12
goto end_branch_15
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
// TAST (Let): __local_var_9_13 -> *Constructor_Data_Maybe_Just
__local_var_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum11_6_4.V2), (*Constructor_Data_Either_Right)(v_8.UnsafePtr).V0))
_ = __local_var_9_13
var __t14 *Constructor_Data_Maybe_Just
{
if (__local_var_9_13 == nil) {
__t14 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_14
} else {

}
}
{
if (__local_var_9_13 != nil) {
__t14 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (__local_var_9_13).V0})}}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t15)}
}))
})
}

func Call_Data_Enum_downFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum, dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_1.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_2)})}
}))
}

func Call_Data_Enum_diag(a_0_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return &Constructor_Data_Tuple_Tuple{1, a_0, a_0}
}

func Call_Data_Enum_downFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum, dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_1_loop
_ = dictUnfoldable_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Enum_diag())
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_1.V1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_3))
}))
}

func Call_Data_Enum_upFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum, dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_1_loop
_ = dictUnfoldable_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Enum_diag())
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_1.V1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), x_3))
}))
}

func Call_Data_Enum_defaultToEnum(dictBounded_0_loop *Constructor_Data_Bounded_Bounded, dictEnum_1_loop *Constructor_Data_Enum_Enum, i_prime_2_loop int64) *Constructor_Data_Maybe_Just {
var dictBounded_0 *Constructor_Data_Bounded_Bounded = dictBounded_0_loop
_ = dictBounded_0
var dictEnum_1 *Constructor_Data_Enum_Enum = dictEnum_1_loop
_ = dictEnum_1
var i_prime_2 int64 = i_prime_2_loop
_ = i_prime_2
var go__go_3_0_0 gopurs_runtime.Value
go__go_3_0_0 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_4_loop int64 = i_4_loop_val.IntVal
var x_5_loop gopurs_runtime.Value = x_5_loop_val
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
var __t3 *Constructor_Data_Maybe_Just
{
if (i_4) == (0) {
__t3 = &Constructor_Data_Maybe_Just{1, x_5}
goto end_branch_3
} else {

}
}
{
// TAST (Let): v_6_1 -> gopurs_runtime.Value
v_6_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_1.V2), x_5)
_ = v_6_1
var __t2 *Constructor_Data_Maybe_Just
{
if (v_6_1.Type == 9 && v_6_1.IntVal == 930809136 && v_6_1.UnsafePtr != nil) {
i_4_loop = (i_4) - (1)
x_5_loop = (*Constructor_Data_Maybe_Just)(v_6_1.UnsafePtr).V0
continue go__go_3_0_0
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v_6_1.Type == 9 && v_6_1.IntVal == 930809136 && v_6_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
var __t5 *Constructor_Data_Maybe_Just
{
var __t4 bool
{
if (i_prime_2) < (0) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
if __t4 {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Int(i_prime_2), gopurs_runtime.Box(dictBounded_0.V1)))
}
end_branch_5:
return __t5
}

func Call_Data_Enum_defaultSucc(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) + (1))))
}

func Call_Data_Enum_defaultPred(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) - (1))))
}

func Call_Data_Enum_defaultFromEnum(dictEnum_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
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
// TAST (Let): v_4_1 -> gopurs_runtime.Value
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_3)
_ = v_4_1
var __t2 int64
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr != nil) {
i_2_loop = (i_2) + (1)
x_3_loop = (*Constructor_Data_Maybe_Just)(v_4_1.UnsafePtr).V0
continue go__go_1_0_1
__t2 = gopurs_runtime.Value{}.IntVal
goto end_branch_2
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr == nil) {
__t2 = i_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_1, gopurs_runtime.Int(0))
}

func Call_Data_Enum_defaultCardinality(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): bottom2_1_0 -> gopurs_runtime.Value
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
// TAST (Let): v_6_2 -> gopurs_runtime.Value
v_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5)
_ = v_6_2
var __t3 int64
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 930809136 && v_6_2.UnsafePtr != nil) {
i_4_loop = (i_4) + (1)
x_5_loop = (*Constructor_Data_Maybe_Just)(v_6_2.UnsafePtr).V0
continue go__go_3_1_2
__t3 = gopurs_runtime.Value{}.IntVal
goto end_branch_3
} else {

}
}
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 930809136 && v_6_2.UnsafePtr == nil) {
__t3 = i_4
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_3:
return gopurs_runtime.Int(__t3)
}
}()
})
})
return gopurs_runtime.Int(gopurs_runtime.Apply2(go__go_3_1_2, gopurs_runtime.Int(1), bottom2_1_0).IntVal)
})
}

func Call_Data_Enum_charToEnum(v_0_loop int64) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var __t3 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (v_0) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal())).IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t_and_2 bool = false
if __t0 {

var __t1 bool
{
if (v_0) > (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal())).IntVal) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
__t_and_2 = __t1
}
if __t_and_2 {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Enum_fromCharCode(), gopurs_runtime.Int(v_0)).StrVal())}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
}

func Call_Data_Enum_cardinality(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "cardinality")
}

func Call_Data_Enum_defaultPred__2391565248(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop int64) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 int64 = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Int(a_2)).IntVal) - (1))))
}

func Call_Data_Enum_defaultPred__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop string) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 string = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Str(a_2)).IntVal) - (1))))
}

func Call_Data_Enum_defaultPred__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) - (1))))
}

func Call_Data_Enum_defaultSucc__2391565248(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop int64) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 int64 = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Int(a_2)).IntVal) + (1))))
}

func Call_Data_Enum_defaultSucc__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop string) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 string = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, gopurs_runtime.Str(a_2)).IntVal) + (1))))
}

func Call_Data_Enum_defaultSucc__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal) + (1))))
}

func Call_Data_Enum_diag__2627963987(a_0_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return &Constructor_Data_Tuple_Tuple{1, a_0, a_0}
}

func Call_Data_Enum_enumFromTo__1480115131(dictEnum_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): Ord0_1_0 -> gopurs_runtime.Value
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V0), gopurs_runtime.Value{})
_ = Ord0_1_0
// TAST (Let): Eq0_2_1 -> *Constructor_Data_Eq_Eq
Eq0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}))
_ = Eq0_2_1
// TAST (Let): Ord01_3_2 -> *Constructor_Data_Ord_Ord
Ord01_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V0), gopurs_runtime.Value{}))
_ = Ord01_3_2
return gopurs_runtime.Func(func(dictUnfoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(Eq0_2_1.V0), v_5, v1_6).IntVal) != (0) {
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Tuple_Tuple
{
var __t6 bool
{
if (i_7.IntVal) > (0) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
if __t6 {
__t7 = &Constructor_Data_Tuple_Tuple{1, v_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_7
} else {

}
}
{
__t7 = &Constructor_Data_Tuple_Tuple{1, v_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_7.IntVal) - (1))})}}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t7)}
}), gopurs_runtime.Int(0))
goto end_branch_13
} else {

}
}
{
var __t9 bool
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(Ord01_3_2.V1), v_5, v1_6)
if (uint32(__t_tag_8.IntVal) == 1527465420) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
if __t9 {
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), a_7)))}, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Maybe_Just
{
var __t11 bool
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_8, v1_6)
if (uint32(__t_tag_10.IntVal) == 380165415) {
__t11 = false
goto end_branch_11
} else {

}
}
{
__t11 = true
}
end_branch_11:
if __t11 {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_applicativeMaybe(), "pure"), Get_Data_Unit_unit()))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(Get_Data_Maybe_plusMaybe(), "empty"))
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_8
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t12)})))}
}))))}})}
}), v_5)
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), a_7)))}, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_8, v1_6)
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_applicativeMaybe(), "pure"), Get_Data_Unit_unit()))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(Get_Data_Maybe_plusMaybe(), "empty"))
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_8
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})))}
}))))}})}
}), v_5)
}
end_branch_13:
return __t13
})
})
})
}

func Call_Data_Enum_fromEnum__2110387339(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Enum_fromEnum__3599151655(v_0_loop int64) int64 {
var v_0 int64 = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Enum_fromEnum__1606852103(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), __eta0_0).IntVal)
}

func Call_Data_Enum_fromEnum__1637084359(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Enum_fromEnum__1196942535(v_0_loop uint32) int64 {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 == 1908470532) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
if (v_0 == 2455627378) {
__t0 = 2
goto end_branch_0
} else {

}
}
{
if (v_0 == 4162469099) {
__t0 = 3
goto end_branch_0
} else {

}
}
{
if (v_0 == 1692989816) {
__t0 = 4
goto end_branch_0
} else {

}
}
{
if (v_0 == 330658827) {
__t0 = 5
goto end_branch_0
} else {

}
}
{
if (v_0 == 4067355978) {
__t0 = 6
goto end_branch_0
} else {

}
}
{
if (v_0 == 2276710548) {
__t0 = 7
goto end_branch_0
} else {

}
}
{
if (v_0 == 243771071) {
__t0 = 8
goto end_branch_0
} else {

}
}
{
if (v_0 == 215731793) {
__t0 = 9
goto end_branch_0
} else {

}
}
{
if (v_0 == 8639228) {
__t0 = 10
goto end_branch_0
} else {

}
}
{
if (v_0 == 49471444) {
__t0 = 11
goto end_branch_0
} else {

}
}
{
if (v_0 == 3889233761) {
__t0 = 12
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return __t0
}

func Call_Data_Enum_pred__2914940949(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), __eta0_0).IntVal) - (1)))))}
}

func Call_Data_Enum_pred__3199041328(dict_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Enum_pred__2010692236(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), __eta0_0).IntVal) - (1)))))}
}

func Call_Data_Enum_succ__412946465(dict_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Enum_succ__2914940949(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), __eta0_0).IntVal) + (1)))))}
}

func Call_Data_Enum_succ__3199041328(dict_0_loop *Constructor_Data_Enum_Enum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Enum_succ__2010692236(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), __eta0_0).IntVal) + (1)))))}
}

func Call_Data_Enum_succ__2858180024(v_0_loop *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
// TAST (Let): sm_1_0 -> *Constructor_Data_Maybe_Just
sm_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_enumMonth(), "succ"), gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}))
_ = sm_1_0
// TAST (Let): v1_2_1 -> *Constructor_Data_Maybe_Just
v1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_enumDay(), "succ"), gopurs_runtime.Int((v_0).V2)))
_ = v1_2_1
var __t5 *Constructor_Data_Maybe_Just
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_1)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Date_lastDayOfMonth(), gopurs_runtime.Int((v_0).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}).IntVal)})})
if (uint32(__t_tag_3.IntVal) == 380165415) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
if __t4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_5
} else {

}
}
{
__t5 = v1_2_1
}
end_branch_5:
// TAST (Let): sd_3_2 -> *Constructor_Data_Maybe_Just
sd_3_2 := __t5
_ = sd_3_2
var __t8 *Constructor_Data_Maybe_Just
{
var __t6 gopurs_runtime.Value
{
if (sd_3_2 == nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if (sd_3_2 != nil) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
var __t7 gopurs_runtime.Value
{
if (sm_1_0 == nil) {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
if (sm_1_0 != nil) {
__t7 = gopurs_runtime.Bool(false)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((__t6.IntVal) != (0)), gopurs_runtime.Bool((__t7.IntVal) != (0))).IntVal) != (0) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_enumYear(), "succ"), gopurs_runtime.Int((v_0).V0))))})
goto end_branch_8
} else {

}
}
{
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((v_0).V0)}
}
end_branch_8:
var __t11 uint32
{
var __t9 gopurs_runtime.Value
{
if (sd_3_2 == nil) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if (sd_3_2 != nil) {
__t9 = gopurs_runtime.Bool(false)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
if (__t9.IntVal) != (0) {
var __t10 gopurs_runtime.Value
{
if (sm_1_0 == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_January().IntVal)), UnsafePtr: nil}
goto end_branch_10
} else {

}
}
{
if (sm_1_0 != nil) {
__t10 = (sm_1_0).V0
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t11 = uint32(__t10.IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = (v_0).V1
}
end_branch_11:
var __t13 *Constructor_Data_Maybe_Just
{
var __t12 gopurs_runtime.Value
{
if (sd_3_2 == nil) {
__t12 = gopurs_runtime.Bool(true)
goto end_branch_12
} else {

}
}
{
if (sd_3_2 != nil) {
__t12 = gopurs_runtime.Bool(false)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
if (__t12.IntVal) != (0) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(1)))
goto end_branch_13
} else {

}
}
{
__t13 = sd_3_2
}
end_branch_13:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), Get_Data_Date_Date(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_applicativeMaybe(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil})))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t13)}))
}

func Call_Data_Enum_toEnum__2738794986(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_Enum_toEnum__2203070892(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_Enum_toEnum__2099864294(n_0_loop int64) *Constructor_Data_Maybe_Just {
var n_0 int64 = n_0_loop
_ = n_0
var __t2 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (gopurs_runtime.Int(n_0).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t1 bool
{
if (gopurs_runtime.Int(n_0).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool(__t0), gopurs_runtime.Bool(__t1)).IntVal) != (0) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0)}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_Enum_toEnum__3317293286(dict_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_Enum_toEnum__2309750950(v_0_loop int64) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0) == (1) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_January().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (2) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_February().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (3) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_March().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (4) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_April().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (5) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_May().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (6) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_June().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (7) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_July().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (8) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_August().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (9) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_September().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (10) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_October().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (11) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_November().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (12) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_December().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_Enum_toEnum__2793813158(v_0_loop int64) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0) == (1) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Monday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (2) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Tuesday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (3) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Wednesday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (4) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Thursday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (5) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Friday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (6) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Saturday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
if (v_0) == (7) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Sunday().IntVal)), UnsafePtr: nil}}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_Enum_toEnumWithDefaults__3941305703(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
// TAST (Let): v_3_0 -> gopurs_runtime.Value
var v_3_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Enum_charToEnum(__eta2_2.IntVal))}
var __t3 string
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 930809136 && v_3_0.UnsafePtr != nil) {
__t3 = (*Constructor_Data_Maybe_Just)(v_3_0.UnsafePtr).V0.StrVal()
goto end_branch_3
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 930809136 && v_3_0.UnsafePtr == nil) {
var __t2 string
{
var __t1 bool
{
if (__eta2_2.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal())).IntVal) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = __eta0_0.StrVal()
goto end_branch_2
} else {

}
}
{
__t2 = __eta1_1.StrVal()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_3:
return gopurs_runtime.Str(__t3)
}

func Call_Data_Enum_toEnumWithDefaults__3558602759(dictBoundedEnum_0_loop *Constructor_Data_Enum_BoundedEnum) gopurs_runtime.Value {
var dictBoundedEnum_0 *Constructor_Data_Enum_BoundedEnum = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): bottom2_1_0 -> gopurs_runtime.Value
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V0), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> gopurs_runtime.Value
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V4), gopurs_runtime.Int(x_4.IntVal))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t4 = (*Constructor_Data_Maybe_Just)(v_5_1.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (x_4.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V3), bottom2_1_0).IntVal) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = low_2
goto end_branch_3
} else {

}
}
{
__t3 = high_3
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
})
})
}

func Get_Data_Enum_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_FromCharCode
}

func Get_Data_Enum_toCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_ToCharCode
}
