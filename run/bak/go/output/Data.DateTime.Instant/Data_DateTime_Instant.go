package Data_DateTime_Instant

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unInstant gopurs_runtime.Value
var once_unInstant sync.Once
func Get_unInstant() gopurs_runtime.Value {
	once_unInstant.Do(func() {
		cache_unInstant = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_unInstant(v_0_box.FloatVal()))
})
	})
	return cache_unInstant
}

var cache_toDateTime gopurs_runtime.Value
var once_toDateTime sync.Once
func Get_toDateTime() gopurs_runtime.Value {
	once_toDateTime.Do(func() {
		cache_toDateTime = gopurs_runtime.Apply(Get_toDateTimeImpl(), gopurs_runtime.Apply(Get_unsafePartial__3962928989(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mo_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mi_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ms_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&pkg_Data_DateTime.Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Apply3(pkg_Data_Date.Get_canonicalDate(), gopurs_runtime.Int(y_1.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_fromJust__4142563260(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toEnum__2309750950(mo_2.IntVal))}))})).IntVal)), UnsafePtr: nil}, gopurs_runtime.Int(d_3.IntVal))), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&pkg_Data_Time.Constructor_Time{1, h_4.IntVal, mi_5.IntVal, s_6.IntVal, ms_7.IntVal})})})}
})
})
})
})
})
})
})
})))
	})
	return cache_toDateTime
}

var cache_showInstant gopurs_runtime.Value
var once_showInstant sync.Once
func Get_showInstant() gopurs_runtime.Value {
	once_showInstant.Do(func() {
		cache_showInstant = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Instant "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__3380206610(gopurs_runtime.Float(v_0.FloatVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showInstant
}

var cache_ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		cache_ordDateTime = pkg_Data_Ord.Get_ordNumber()
	})
	return cache_ordDateTime
}

var cache_instant gopurs_runtime.Value
var once_instant sync.Once
func Get_instant() gopurs_runtime.Value {
	once_instant.Do(func() {
		cache_instant = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_instant(v_0_box.FloatVal()))}
})
	})
	return cache_instant
}

var cache_fromDateTime gopurs_runtime.Value
var once_fromDateTime sync.Once
func Get_fromDateTime() gopurs_runtime.Value {
	once_fromDateTime.Do(func() {
		cache_fromDateTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_fromDateTime(gopurs_runtime.CoerceToStruct[pkg_Data_DateTime.Constructor_DateTime](v_0_box)))
})
	})
	return cache_fromDateTime
}

var cache_fromDate gopurs_runtime.Value
var once_fromDate sync.Once
func Get_fromDate() gopurs_runtime.Value {
	once_fromDate.Do(func() {
		cache_fromDate = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_fromDate(gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](d_0_box)))
})
	})
	return cache_fromDate
}

var cache_eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		cache_eqDateTime = pkg_Data_Eq.Get_eqNumber()
	})
	return cache_eqDateTime
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), dt1_1_box.FloatVal(), dt2_2_box.FloatVal())
})
	})
	return cache_diff
}

var cache_boundedInstant gopurs_runtime.Value
var once_boundedInstant sync.Once
func Get_boundedInstant() gopurs_runtime.Value {
	once_boundedInstant.Do(func() {
		cache_boundedInstant = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
}), gopurs_runtime.Float(Call_negate__2151342916(gopurs_runtime.Float(8639977881600000.0)).FloatVal()), gopurs_runtime.Float(8639977881599999.0))
	})
	return cache_boundedInstant
}

var cache_compose__1987728071 gopurs_runtime.Value
var once_compose__1987728071 sync.Once
func Get_compose__1987728071() gopurs_runtime.Value {
	once_compose__1987728071.Do(func() {
		cache_compose__1987728071 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1987728071(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__1987728071
}

var cache_compose__1555187646 gopurs_runtime.Value
var once_compose__1555187646 sync.Once
func Get_compose__1555187646() gopurs_runtime.Value {
	once_compose__1555187646.Do(func() {
		cache_compose__1555187646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1555187646(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1555187646
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

var cache_fromEnum__1196942535 gopurs_runtime.Value
var once_fromEnum__1196942535 sync.Once
func Get_fromEnum__1196942535() gopurs_runtime.Value {
	once_fromEnum__1196942535.Do(func() {
		cache_fromEnum__1196942535 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_fromEnum__1196942535(uint32(v_0_box.IntVal)))
})
	})
	return cache_fromEnum__1196942535
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

var cache_toEnum__2309750950 gopurs_runtime.Value
var once_toEnum__2309750950 sync.Once
func Get_toEnum__2309750950() gopurs_runtime.Value {
	once_toEnum__2309750950.Do(func() {
		cache_toEnum__2309750950 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toEnum__2309750950(v_0_box.IntVal))}
})
	})
	return cache_toEnum__2309750950
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
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
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
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
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
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

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_fromJust__4142563260 gopurs_runtime.Value
var once_fromJust__4142563260 sync.Once
func Get_fromJust__4142563260() gopurs_runtime.Value {
	once_fromJust__4142563260.Do(func() {
		cache_fromJust__4142563260 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__4142563260(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__4142563260
}

var cache_over__3209389325 gopurs_runtime.Value
var once_over__3209389325 sync.Once
func Get_over__3209389325() gopurs_runtime.Value {
	once_over__3209389325.Do(func() {
		cache_over__3209389325 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__3209389325(v_0_box)
})
	})
	return cache_over__3209389325
}

var cache_over__3306352994 gopurs_runtime.Value
var once_over__3306352994 sync.Once
func Get_over__3306352994() gopurs_runtime.Value {
	once_over__3306352994.Do(func() {
		cache_over__3306352994 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__3306352994(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_1_box), v_2_box)
})
	})
	return cache_over__3306352994
}

var cache_over__1462660013 gopurs_runtime.Value
var once_over__1462660013 sync.Once
func Get_over__1462660013() gopurs_runtime.Value {
	once_over__1462660013.Do(func() {
		cache_over__1462660013 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__1462660013(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_1_box), v_2_box)
})
	})
	return cache_over__1462660013
}

var cache_over__2794260290 gopurs_runtime.Value
var once_over__2794260290 sync.Once
func Get_over__2794260290() gopurs_runtime.Value {
	once_over__2794260290.Do(func() {
		cache_over__2794260290 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__2794260290(_dollar__unused_0_box, v_1_box)
})
	})
	return cache_over__2794260290
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
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

var cache_greaterThanOrEq__1061005983 gopurs_runtime.Value
var once_greaterThanOrEq__1061005983 sync.Once
func Get_greaterThanOrEq__1061005983() gopurs_runtime.Value {
	once_greaterThanOrEq__1061005983.Do(func() {
		cache_greaterThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__1061005983
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

var cache_lessThanOrEq__1061005983 gopurs_runtime.Value
var once_lessThanOrEq__1061005983 sync.Once
func Get_lessThanOrEq__1061005983() gopurs_runtime.Value {
	once_lessThanOrEq__1061005983.Do(func() {
		cache_lessThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__1061005983
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_negate__2151342916 gopurs_runtime.Value
var once_negate__2151342916 sync.Once
func Get_negate__2151342916() gopurs_runtime.Value {
	once_negate__2151342916.Do(func() {
		cache_negate__2151342916 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2151342916(__eta0_0_box)
})
	})
	return cache_negate__2151342916
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
}

var cache_sub__1124926121 gopurs_runtime.Value
var once_sub__1124926121 sync.Once
func Get_sub__1124926121() gopurs_runtime.Value {
	once_sub__1124926121.Do(func() {
		cache_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__1124926121
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
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

var cache_append__3678571768 gopurs_runtime.Value
var once_append__3678571768 sync.Once
func Get_append__3678571768() gopurs_runtime.Value {
	once_append__3678571768.Do(func() {
		cache_append__3678571768 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_append__3678571768(v_0_box.FloatVal(), v1_1_box.FloatVal()))
})
	})
	return cache_append__3678571768
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_add__101133084 gopurs_runtime.Value
var once_add__101133084 sync.Once
func Get_add__101133084() gopurs_runtime.Value {
	once_add__101133084.Do(func() {
		cache_add__101133084 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__101133084(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[float64]](dict_0_box))
})
	})
	return cache_add__101133084
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__137136408 gopurs_runtime.Value
var once_add__137136408 sync.Once
func Get_add__137136408() gopurs_runtime.Value {
	once_add__137136408.Do(func() {
		cache_add__137136408 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__137136408(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__137136408
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

var cache_zero__1556010056 gopurs_runtime.Value
var once_zero__1556010056 sync.Once
func Get_zero__1556010056() gopurs_runtime.Value {
	once_zero__1556010056.Do(func() {
		cache_zero__1556010056 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1556010056(dict_0_box)
})
	})
	return cache_zero__1556010056
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_show__3380206610 gopurs_runtime.Value
var once_show__3380206610 sync.Once
func Get_show__3380206610() gopurs_runtime.Value {
	once_show__3380206610.Do(func() {
		cache_show__3380206610 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3380206610(__eta0_0_box)
})
	})
	return cache_show__3380206610
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_negateDuration__4195558286 gopurs_runtime.Value
var once_negateDuration__4195558286 sync.Once
func Get_negateDuration__4195558286() gopurs_runtime.Value {
	once_negateDuration__4195558286.Do(func() {
		cache_negateDuration__4195558286 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negateDuration__4195558286(__eta0_0_box)
})
	})
	return cache_negateDuration__4195558286
}

var cache_negateDuration__3870190523 gopurs_runtime.Value
var once_negateDuration__3870190523 sync.Once
func Get_negateDuration__3870190523() gopurs_runtime.Value {
	once_negateDuration__3870190523.Do(func() {
		cache_negateDuration__3870190523 = gopurs_runtime.Func2(func(dictDuration_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negateDuration__3870190523(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), x_1_box)
})
	})
	return cache_negateDuration__3870190523
}

var cache_toDuration__2440169646 gopurs_runtime.Value
var once_toDuration__2440169646 sync.Once
func Get_toDuration__2440169646() gopurs_runtime.Value {
	once_toDuration__2440169646.Do(func() {
		cache_toDuration__2440169646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toDuration__2440169646(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toDuration__2440169646
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1130268957 gopurs_runtime.Value
var once_unsafePartial__1130268957 sync.Once
func Get_unsafePartial__1130268957() gopurs_runtime.Value {
	once_unsafePartial__1130268957.Do(func() {
		cache_unsafePartial__1130268957 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1130268957
}

var cache_unsafePartial__3962928989 gopurs_runtime.Value
var once_unsafePartial__3962928989 sync.Once
func Get_unsafePartial__3962928989() gopurs_runtime.Value {
	once_unsafePartial__3962928989.Do(func() {
		cache_unsafePartial__3962928989 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__3962928989
}

func Call_unInstant(v_0_loop float64) float64 {
var v_0 float64 = v_0_loop
_ = v_0
return v_0
}

func Call_instant(v_0_loop float64) *pkg_Data_Maybe.Constructor_Just[float64] {
var v_0 float64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(gopurs_runtime.Float(v_0), gopurs_runtime.Float(Call_negate__2151342916(gopurs_runtime.Float(8639977881600000.0)).FloatVal()))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(gopurs_runtime.Float(v_0), gopurs_runtime.Float(8639977881599999.0))).IntVal) != (0))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(v_0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[float64]](__t0)
}

func Call_fromDateTime(v_0_loop *pkg_Data_DateTime.Constructor_DateTime) float64 {
var v_0 *pkg_Data_DateTime.Constructor_DateTime = v_0_loop
_ = v_0
return gopurs_runtime.UncurriedApp7(Get_fromDateTimeImpl(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V0), gopurs_runtime.Int(gopurs_runtime.Int(Call_fromEnum__1196942535((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V1)).IntVal), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V2), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V2), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V3)).FloatVal()
}

func Call_fromDate(d_0_loop *pkg_Data_Date.Constructor_Date) float64 {
var d_0 *pkg_Data_Date.Constructor_Date = d_0_loop
_ = d_0
return gopurs_runtime.UncurriedApp7(Get_fromDateTimeImpl(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(d_0)}.UnsafePtr).V0), gopurs_runtime.Int(gopurs_runtime.Int(Call_fromEnum__1196942535((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(d_0)}.UnsafePtr).V1)).IntVal), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(d_0)}.UnsafePtr).V2), gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "bottom").IntVal), gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "bottom").IntVal), gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "bottom").IntVal), gopurs_runtime.Int(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "bottom").IntVal)).FloatVal()
}

func Call_diff(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], dt1_1_loop float64, dt2_2_loop float64) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 float64 = dt1_1_loop
_ = dt1_1
var dt2_2 float64 = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Float(gopurs_runtime.Float(Call_append__3678571768(dt1_1, Call_negateDuration__4195558286(gopurs_runtime.Float(dt2_2)).FloatVal())).FloatVal()))
}

func Call_compose__1987728071(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1555187646(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fromEnum__1196942535(v_0_loop uint32) int64 {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 1908470532) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_0 == 2455627378) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (v_0 == 4162469099) {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (v_0 == 1692989816) {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (v_0 == 330658827) {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (v_0 == 4067355978) {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (v_0 == 2276710548) {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (v_0 == 243771071) {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (v_0 == 215731793) {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (v_0 == 8639228) {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (v_0 == 49471444) {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (v_0 == 3889233761) {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnum__2309750950(v_0_loop int64) *pkg_Data_Maybe.Constructor_Just[uint32] {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (2) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (3) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (4) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (5) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (6) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (7) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (8) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (9) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (10) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (11) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v_0) == (12) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[uint32]](__t0)
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_fromJust__4142563260(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
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

func Call_over__3209389325(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over__3306352994(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_1_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over__1462660013(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_1_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over__2794260290(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) < (a2_1.FloatVal()) {
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

func Call_lessThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) > (a2_1.FloatVal()) {
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

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
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

func Call_negate__2151342916(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Float(-(__eta0_0.FloatVal()))
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__3678571768(v_0_loop float64, v1_1_loop float64) float64 {
var v_0 float64 = v_0_loop
_ = v_0
var v1_1 float64 = v1_1_loop
_ = v1_1
return (v_0) + (v1_1)
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__101133084(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[float64]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[float64] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__137136408(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) + (__eta1_1.FloatVal()))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_zero__1556010056(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__3380206610(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), __eta0_0).StrVal())
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_negateDuration__4195558286(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), __eta0_0)))
}

func Call_negateDuration__3870190523(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(dictDuration_0.V0, x_1)))
}

func Call_toDuration__2440169646(dict_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Get_fromDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_FromDateTimeImpl
}

func Get_toDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_ToDateTimeImpl
}
