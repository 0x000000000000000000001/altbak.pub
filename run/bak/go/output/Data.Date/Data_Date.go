package Data_Date

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_ordMaybe gopurs_runtime.Value
var once_ordMaybe sync.Once
func Get_ordMaybe() gopurs_runtime.Value {
	once_ordMaybe.Do(func() {
		cache_ordMaybe = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordDay(), "Eq0"), gopurs_runtime.Value{})
_ = __local_var_0_1
eqMaybe1_0_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t2 bool
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
}))
_ = eqMaybe1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t4 uint32
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordDay(), "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t5.IntVal)), UnsafePtr: nil}
})
}))))}
}()
	})
	return cache_ordMaybe
}

var cache_Date gopurs_runtime.Value
var once_Date sync.Once
func Get_Date() gopurs_runtime.Value {
	once_Date.Do(func() {
		cache_Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, value0.IntVal, uint32(value1.IntVal), value2.IntVal})}
})
})
})
	})
	return cache_Date
}

var cache_year gopurs_runtime.Value
var once_year sync.Once
func Get_year() gopurs_runtime.Value {
	once_year.Do(func() {
		cache_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_year(gopurs_runtime.CoerceToStruct[Constructor_Date](v_0_box)))
})
	})
	return cache_year
}

var cache_weekday gopurs_runtime.Value
var once_weekday sync.Once
func Get_weekday() gopurs_runtime.Value {
	once_weekday.Do(func() {
		cache_weekday = gopurs_runtime.Apply(Get_unsafePartial__2741364381(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
n_2_0 := gopurs_runtime.UncurriedApp3(Get_calcWeekday(), gopurs_runtime.Int((*Constructor_Date)(v_1.UnsafePtr).V0), gopurs_runtime.Apply(Get_fromEnum__1196942535(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_1.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Int((*Constructor_Date)(v_1.UnsafePtr).V2))
_ = n_2_0
var __t1 uint32
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), n_2_0, gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = uint32(Call_fromJust__3809843644(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_toEnum__2793813158(), gopurs_runtime.Int(7)))).IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = uint32(Call_fromJust__3809843644(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_toEnum__2793813158(), n_2_0))).IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
})
}))
	})
	return cache_weekday
}

var cache_showDate gopurs_runtime.Value
var once_showDate sync.Once
func Get_showDate() gopurs_runtime.Value {
	once_showDate.Do(func() {
		cache_showDate = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Date "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(Get_show__1488465650(), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0)), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(Get_show__1626410898(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(Get_show__1488465650(), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V2)), gopurs_runtime.Str(")"))))))).StrVal())
}))
	})
	return cache_showDate
}

var cache_month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		cache_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_month(gopurs_runtime.CoerceToStruct[Constructor_Date](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_month
}

var cache_isLeapYear gopurs_runtime.Value
var once_isLeapYear sync.Once
func Get_isLeapYear() gopurs_runtime.Value {
	once_isLeapYear.Do(func() {
		cache_isLeapYear = gopurs_runtime.Func(func(y_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isLeapYear(y_0_box.IntVal))
})
	})
	return cache_isLeapYear
}

var cache_lastDayOfMonth gopurs_runtime.Value
var once_lastDayOfMonth sync.Once
func Get_lastDayOfMonth() gopurs_runtime.Value {
	once_lastDayOfMonth.Do(func() {
		cache_lastDayOfMonth = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_lastDayOfMonth(y_0_box.IntVal, uint32(m_1_box.IntVal)))
})
	})
	return cache_lastDayOfMonth
}

var cache_eqDate gopurs_runtime.Value
var once_eqDate sync.Once
func Get_eqDate() gopurs_runtime.Value {
	once_eqDate.Do(func() {
		cache_eqDate = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V0)), gopurs_runtime.Apply2(Get_eq__3887832182(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil})), gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V2))).IntVal) != (0))
})
}))
	})
	return cache_eqDate
}

var cache_ordDate gopurs_runtime.Value
var once_ordDate sync.Once
func Get_ordDate() gopurs_runtime.Value {
	once_ordDate.Do(func() {
		cache_ordDate = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDate()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V0))
_ = v_2_0
var __t3 uint32
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(Get_compare__696857420(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil})
_ = v1_3_1
var __t2 uint32
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
__t2 = 380165415
goto end_branch_2
} else {

}
}
{
__t2 = uint32(gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V2)).IntVal)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
}))
	})
	return cache_ordDate
}

var cache_enumDate gopurs_runtime.Value
var once_enumDate sync.Once
func Get_enumDate() gopurs_runtime.Value {
	once_enumDate.Do(func() {
		cache_enumDate = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
pm_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_pred__2010692236(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}))
_ = pm_1_0
pd_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_pred__2914940949(), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V2)))
_ = pd_2_1
var __t3 uint32
{
if (gopurs_runtime.Bool(Call_isNothing__1358705270(pd_2_1)).IntVal) != (0) {
__t3 = uint32(Call_fromMaybe__18840980(gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}, pm_1_0).IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Date)(v_0.UnsafePtr).V1
}
end_branch_3:
m_prime_3_2 := __t3
_ = m_prime_3_2
l_4_4 := gopurs_runtime.Int(Call_lastDayOfMonth((*Constructor_Date)(v_0.UnsafePtr).V0, m_prime_3_2))
_ = l_4_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool(Call_isNothing__1358705270(pd_2_1)), gopurs_runtime.Bool(Call_isNothing__2787066607(pm_1_0))).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_pred__2914940949(), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0)})}
}
end_branch_5:
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_isNothing__1358705270(pd_2_1)).IntVal) != (0) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, l_4_4})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(pd_2_1)}
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(Get_apply__1471729482(), gopurs_runtime.Apply2(Get_apply__1572009162(), gopurs_runtime.Apply2(Get_map__4155962236(), Get_Date(), __t5), gopurs_runtime.Apply(Get_pure__3181299446(), gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime_3_2), UnsafePtr: nil})), __t6)))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
sm_1_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_succ__2010692236(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}))
_ = sm_1_7
v1_2_8 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_succ__2914940949(), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V2)))
_ = v1_2_8
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThan__2400628110(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_8)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_lastDayOfMonth((*Constructor_Date)(v_0.UnsafePtr).V0, (*Constructor_Date)(v_0.UnsafePtr).V1))})})).IntVal) != (0) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_2_8)}
}
end_branch_10:
sd_3_9 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t10)
_ = sd_3_9
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool(Call_isNothing__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_9)}))), gopurs_runtime.Bool(Call_isNothing__2787066607(sm_1_7))).IntVal) != (0) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_succ__2914940949(), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0)})}
}
end_branch_11:
var __t12 uint32
{
if (gopurs_runtime.Bool(Call_isNothing__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_9)}))).IntVal) != (0) {
__t12 = uint32(Call_fromMaybe__18840980(gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}, sm_1_7).IntVal)
goto end_branch_12
} else {

}
}
{
__t12 = (*Constructor_Date)(v_0.UnsafePtr).V1
}
end_branch_12:
var __t13 *pkg_Data_Maybe.Constructor_Just[int64]
{
if (gopurs_runtime.Bool(Call_isNothing__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_9)}))).IntVal) != (0) {
__t13 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.Int(1)))
goto end_branch_13
} else {

}
}
{
__t13 = sd_3_9
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(Get_apply__1471729482(), gopurs_runtime.Apply2(Get_apply__1572009162(), gopurs_runtime.Apply2(Get_map__4155962236(), Get_Date(), __t11), gopurs_runtime.Apply(Get_pure__3181299446(), gopurs_runtime.Value{Type: 9, IntVal: int64(__t12), UnsafePtr: nil})), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t13)})))}
}))
	})
	return cache_enumDate
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Date](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Date](v1_2_box))
})
	})
	return cache_diff
}

var cache_day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		cache_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_day(gopurs_runtime.CoerceToStruct[Constructor_Date](v_0_box)))
})
	})
	return cache_day
}

var cache_canonicalDate gopurs_runtime.Value
var once_canonicalDate sync.Once
func Get_canonicalDate() gopurs_runtime.Value {
	once_canonicalDate.Do(func() {
		cache_canonicalDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_canonicalDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal))}
})
	})
	return cache_canonicalDate
}

var cache_exactDate gopurs_runtime.Value
var once_exactDate sync.Once
func Get_exactDate() gopurs_runtime.Value {
	once_exactDate.Do(func() {
		cache_exactDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_exactDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal))}
})
	})
	return cache_exactDate
}

var cache_boundedDate gopurs_runtime.Value
var once_boundedDate sync.Once
func Get_boundedDate() gopurs_runtime.Value {
	once_boundedDate.Do(func() {
		cache_boundedDate = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
}), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedYear(), "bottom").IntVal, uint32(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom").IntVal), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom").IntVal})}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedYear(), "top").IntVal, uint32(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "top").IntVal), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "top").IntVal})})
	})
	return cache_boundedDate
}

var cache_adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		cache_adjust = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, date_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_adjust(v_0_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Date](date_1_box)))}
})
	})
	return cache_adjust
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__3181299446 gopurs_runtime.Value
var once_pure__3181299446 sync.Once
func Get_pure__3181299446() gopurs_runtime.Value {
	once_pure__3181299446.Do(func() {
		cache_pure__3181299446 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure")
	})
	return cache_pure__3181299446
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

var cache_apply__1471729482 gopurs_runtime.Value
var once_apply__1471729482 sync.Once
func Get_apply__1471729482() gopurs_runtime.Value {
	once_apply__1471729482.Do(func() {
		cache_apply__1471729482 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__1471729482
}

var cache_apply__1572009162 gopurs_runtime.Value
var once_apply__1572009162 sync.Once
func Get_apply__1572009162() gopurs_runtime.Value {
	once_apply__1572009162.Do(func() {
		cache_apply__1572009162 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__1572009162
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__4062037089 gopurs_runtime.Value
var once_bind__4062037089 sync.Once
func Get_bind__4062037089() gopurs_runtime.Value {
	once_bind__4062037089.Do(func() {
		cache_bind__4062037089 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind")
	})
	return cache_bind__4062037089
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_bindFlipped__3917280577 gopurs_runtime.Value
var once_bindFlipped__3917280577 sync.Once
func Get_bindFlipped__3917280577() gopurs_runtime.Value {
	once_bindFlipped__3917280577.Do(func() {
		cache_bindFlipped__3917280577 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__3917280577(b_0_box, a_1_box)
})
	})
	return cache_bindFlipped__3917280577
}

var cache_bindFlipped__1454086721 gopurs_runtime.Value
var once_bindFlipped__1454086721 sync.Once
func Get_bindFlipped__1454086721() gopurs_runtime.Value {
	once_bindFlipped__1454086721.Do(func() {
		cache_bindFlipped__1454086721 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1454086721(b_0_box, a_1_box)
})
	})
	return cache_bindFlipped__1454086721
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

var cache_fromEnum__3599151655 gopurs_runtime.Value
var once_fromEnum__3599151655 sync.Once
func Get_fromEnum__3599151655() gopurs_runtime.Value {
	once_fromEnum__3599151655.Do(func() {
		cache_fromEnum__3599151655 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "fromEnum")
	})
	return cache_fromEnum__3599151655
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
		cache_fromEnum__1196942535 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum")
	})
	return cache_fromEnum__1196942535
}

var cache_pred__2914940949 gopurs_runtime.Value
var once_pred__2914940949 sync.Once
func Get_pred__2914940949() gopurs_runtime.Value {
	once_pred__2914940949.Do(func() {
		cache_pred__2914940949 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "pred")
	})
	return cache_pred__2914940949
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

var cache_pred__2010692236 gopurs_runtime.Value
var once_pred__2010692236 sync.Once
func Get_pred__2010692236() gopurs_runtime.Value {
	once_pred__2010692236.Do(func() {
		cache_pred__2010692236 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred")
	})
	return cache_pred__2010692236
}

var cache_succ__2914940949 gopurs_runtime.Value
var once_succ__2914940949 sync.Once
func Get_succ__2914940949() gopurs_runtime.Value {
	once_succ__2914940949.Do(func() {
		cache_succ__2914940949 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "succ")
	})
	return cache_succ__2914940949
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

var cache_succ__2010692236 gopurs_runtime.Value
var once_succ__2010692236 sync.Once
func Get_succ__2010692236() gopurs_runtime.Value {
	once_succ__2010692236.Do(func() {
		cache_succ__2010692236 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "succ")
	})
	return cache_succ__2010692236
}

var cache_succ__2858180024 gopurs_runtime.Value
var once_succ__2858180024 sync.Once
func Get_succ__2858180024() gopurs_runtime.Value {
	once_succ__2858180024.Do(func() {
		cache_succ__2858180024 = gopurs_runtime.RecordGet(Get_enumDate(), "succ")
	})
	return cache_succ__2858180024
}

var cache_toEnum__2099864294 gopurs_runtime.Value
var once_toEnum__2099864294 sync.Once
func Get_toEnum__2099864294() gopurs_runtime.Value {
	once_toEnum__2099864294.Do(func() {
		cache_toEnum__2099864294 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum")
	})
	return cache_toEnum__2099864294
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
		cache_toEnum__2309750950 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum")
	})
	return cache_toEnum__2309750950
}

var cache_toEnum__2793813158 gopurs_runtime.Value
var once_toEnum__2793813158 sync.Once
func Get_toEnum__2793813158() gopurs_runtime.Value {
	once_toEnum__2793813158.Do(func() {
		cache_toEnum__2793813158 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumWeekday(), "toEnum")
	})
	return cache_toEnum__2793813158
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

var cache_eq__3887832182 gopurs_runtime.Value
var once_eq__3887832182 sync.Once
func Get_eq__3887832182() gopurs_runtime.Value {
	once_eq__3887832182.Do(func() {
		cache_eq__3887832182 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqMonth(), "eq")
	})
	return cache_eq__3887832182
}

var cache_eq__1204755874 gopurs_runtime.Value
var once_eq__1204755874 sync.Once
func Get_eq__1204755874() gopurs_runtime.Value {
	once_eq__1204755874.Do(func() {
		cache_eq__1204755874 = gopurs_runtime.RecordGet(Get_eqDate(), "eq")
	})
	return cache_eq__1204755874
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

var cache_const__220790420 gopurs_runtime.Value
var once_const__220790420 sync.Once
func Get_const__220790420() gopurs_runtime.Value {
	once_const__220790420.Do(func() {
		cache_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__220790420(a_0_box, v_1_box)
})
	})
	return cache_const__220790420
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_flip__1826582752 gopurs_runtime.Value
var once_flip__1826582752 sync.Once
func Get_flip__1826582752() gopurs_runtime.Value {
	once_flip__1826582752.Do(func() {
		cache_flip__1826582752 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1826582752(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1826582752
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

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
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

var cache_map__3447677596 gopurs_runtime.Value
var once_map__3447677596 sync.Once
func Get_map__3447677596() gopurs_runtime.Value {
	once_map__3447677596.Do(func() {
		cache_map__3447677596 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__3447677596
}

var cache_map__4155962236 gopurs_runtime.Value
var once_map__4155962236 sync.Once
func Get_map__4155962236() gopurs_runtime.Value {
	once_map__4155962236.Do(func() {
		cache_map__4155962236 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__4155962236
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
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

var cache_applicativeMaybe__3016118221 gopurs_runtime.Value
var once_applicativeMaybe__3016118221 sync.Once
func Get_applicativeMaybe__3016118221() gopurs_runtime.Value {
	once_applicativeMaybe__3016118221.Do(func() {
		cache_applicativeMaybe__3016118221 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__3016118221
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

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__1910292045 gopurs_runtime.Value
var once_bindMaybe__1910292045 sync.Once
func Get_bindMaybe__1910292045() gopurs_runtime.Value {
	once_bindMaybe__1910292045.Do(func() {
		cache_bindMaybe__1910292045 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
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
	return cache_bindMaybe__1910292045
}

var cache_fromJust__1577979644 gopurs_runtime.Value
var once_fromJust__1577979644 sync.Once
func Get_fromJust__1577979644() gopurs_runtime.Value {
	once_fromJust__1577979644.Do(func() {
		cache_fromJust__1577979644 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1577979644(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__1577979644
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

var cache_fromJust__3809843644 gopurs_runtime.Value
var once_fromJust__3809843644 sync.Once
func Get_fromJust__3809843644() gopurs_runtime.Value {
	once_fromJust__3809843644.Do(func() {
		cache_fromJust__3809843644 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__3809843644(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__3809843644
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
}

var cache_fromMaybe__18840980 gopurs_runtime.Value
var once_fromMaybe__18840980 sync.Once
func Get_fromMaybe__18840980() gopurs_runtime.Value {
	once_fromMaybe__18840980.Do(func() {
		cache_fromMaybe__18840980 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__18840980(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__18840980
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

var cache_isNothing__1358705270 gopurs_runtime.Value
var once_isNothing__1358705270 sync.Once
func Get_isNothing__1358705270() gopurs_runtime.Value {
	once_isNothing__1358705270.Do(func() {
		cache_isNothing__1358705270 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__1358705270
}

var cache_isNothing__4206805139 gopurs_runtime.Value
var once_isNothing__4206805139 sync.Once
func Get_isNothing__4206805139() gopurs_runtime.Value {
	once_isNothing__4206805139.Do(func() {
		cache_isNothing__4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__4206805139(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__4206805139
}

var cache_isNothing__2787066607 gopurs_runtime.Value
var once_isNothing__2787066607 sync.Once
func Get_isNothing__2787066607() gopurs_runtime.Value {
	once_isNothing__2787066607.Do(func() {
		cache_isNothing__2787066607 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2787066607(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2787066607
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_compare__372254389 gopurs_runtime.Value
var once_compare__372254389 sync.Once
func Get_compare__372254389() gopurs_runtime.Value {
	once_compare__372254389.Do(func() {
		cache_compare__372254389 = gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})
	})
	return cache_compare__372254389
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

var cache_compare__696857420 gopurs_runtime.Value
var once_compare__696857420 sync.Once
func Get_compare__696857420() gopurs_runtime.Value {
	once_compare__696857420.Do(func() {
		cache_compare__696857420 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordMonth(), "compare")
	})
	return cache_compare__696857420
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_greaterThan__2400628110 gopurs_runtime.Value
var once_greaterThan__2400628110 sync.Once
func Get_greaterThan__2400628110() gopurs_runtime.Value {
	once_greaterThan__2400628110.Do(func() {
		cache_greaterThan__2400628110 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__2400628110(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__2400628110
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

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
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

var cache_negate__2635823316 gopurs_runtime.Value
var once_negate__2635823316 sync.Once
func Get_negate__2635823316() gopurs_runtime.Value {
	once_negate__2635823316.Do(func() {
		cache_negate__2635823316 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2635823316(a_0_box)
})
	})
	return cache_negate__2635823316
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

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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

var cache_show__1488465650 gopurs_runtime.Value
var once_show__1488465650 sync.Once
func Get_show__1488465650() gopurs_runtime.Value {
	once_show__1488465650.Do(func() {
		cache_show__1488465650 = gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show")
	})
	return cache_show__1488465650
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

var cache_show__1626410898 gopurs_runtime.Value
var once_show__1626410898 sync.Once
func Get_show__1626410898() gopurs_runtime.Value {
	once_show__1626410898.Do(func() {
		cache_show__1626410898 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showMonth(), "show")
	})
	return cache_show__1626410898
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

var cache_unsafePartial__2741364381 gopurs_runtime.Value
var once_unsafePartial__2741364381 sync.Once
func Get_unsafePartial__2741364381() gopurs_runtime.Value {
	once_unsafePartial__2741364381.Do(func() {
		cache_unsafePartial__2741364381 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__2741364381
}

var cache_unsafePartial__1059034269 gopurs_runtime.Value
var once_unsafePartial__1059034269 sync.Once
func Get_unsafePartial__1059034269() gopurs_runtime.Value {
	once_unsafePartial__1059034269.Do(func() {
		cache_unsafePartial__1059034269 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1059034269
}

type Constructor_Date struct {
	Rc uint32
	V0 int64
	V1 uint32
	V2 int64
}


func Call_year(v_0_loop *Constructor_Date) int64 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
return (*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_month(v_0_loop *Constructor_Date) uint32 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
return (*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_isLeapYear(y_0_loop int64) bool {
var y_0 int64 = y_0_loop
_ = y_0
y_prime_1_0 := gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int(y_0)).IntVal
_ = y_prime_1_0
return (gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Apply2(Get_mod__2185172824(), gopurs_runtime.Int(y_prime_1_0), gopurs_runtime.Int(4)), gopurs_runtime.Int(0)), gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Apply2(Get_mod__2185172824(), gopurs_runtime.Int(y_prime_1_0), gopurs_runtime.Int(400)), gopurs_runtime.Int(0)), gopurs_runtime.Apply(Get_not__3201284355(), gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Apply2(Get_mod__2185172824(), gopurs_runtime.Int(y_prime_1_0), gopurs_runtime.Int(100)), gopurs_runtime.Int(0))))).IntVal) != (0)
}

func Call_lastDayOfMonth(y_0_loop int64, m_1_loop uint32) int64 {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
__local_var_2_0 := gopurs_runtime.Apply(Get_unsafePartial__1059034269(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__1577979644()
}))
_ = __local_var_2_0
var __t2 gopurs_runtime.Value
{
if (m_1 == 1908470532) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 2455627378) {
var __t1 int64
{
if (gopurs_runtime.Bool(Call_isLeapYear(y_0)).IntVal) != (0) {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(29))).IntVal
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(28))).IntVal
}
end_branch_1:
__t2 = gopurs_runtime.Int(__t1)
goto end_branch_2
} else {

}
}
{
if (m_1 == 4162469099) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 1692989816) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 330658827) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 4067355978) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 2276710548) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 243771071) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 215731793) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 8639228) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 49471444) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal)
goto end_branch_2
} else {

}
}
{
if (m_1 == 3889233761) {
__t2 = gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2.IntVal
}

func Call_diff(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], v_1_loop *Constructor_Date, v1_2_loop *Constructor_Date) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var v_1 *Constructor_Date = v_1_loop
_ = v_1
var v1_2 *Constructor_Date = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.UncurriedApp6(Get_calcDiff(), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0), gopurs_runtime.Apply(Get_fromEnum__1196942535(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V0), gopurs_runtime.Apply(Get_fromEnum__1196942535(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V2)))
}

func Call_day(v_0_loop *Constructor_Date) int64 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
return (*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2
}

func Call_canonicalDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) *Constructor_Date {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var d_2 int64 = d_2_loop
_ = d_2
return gopurs_runtime.CoerceToStruct[Constructor_Date](gopurs_runtime.UncurriedApp4(Get_canonicalDateImpl(), gopurs_runtime.Apply(Get_unsafePartial__1130268957(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, y_prime_4.IntVal, uint32(Call_fromJust__4142563260(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_toEnum__2309750950(), m_prime_5))).IntVal), d_prime_6.IntVal})}
})
})
})
})), gopurs_runtime.Int(y_0), gopurs_runtime.Apply(Get_fromEnum__1196942535(), gopurs_runtime.Value{Type: 9, IntVal: int64(m_1), UnsafePtr: nil}), gopurs_runtime.Int(d_2)))
}

func Call_exactDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) *pkg_Data_Maybe.Constructor_Just[*Constructor_Date] {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var d_2 int64 = d_2_loop
_ = d_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__1204755874(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_canonicalDate(y_0, m_1, d_2))}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, y_0, m_1, d_2})}).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, y_0, m_1, d_2})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](__t0)
}

func Call_adjust(v_0_loop float64, date_1_loop *Constructor_Date) *pkg_Data_Maybe.Constructor_Just[*Constructor_Date] {
var v_0 float64 = v_0_loop
_ = v_0
var date_1 *Constructor_Date = date_1_loop
_ = date_1
var adj_2_0_0 gopurs_runtime.Value
_ = adj_2_0_0
adj_2_0_0 = gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v1_3.IntVal) == (0) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v2_4})}
goto end_branch_8
} else {

}
}
{
j_5_1 := gopurs_runtime.Apply2(Get_add__560788792(), v1_3, gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V2)))
_ = j_5_1
low_6_2 := gopurs_runtime.Bool(Call_lessThan__4087042607(j_5_1, gopurs_runtime.Int(1)))
_ = low_6_2
var __t4 uint32
{
if (low_6_2.IntVal) != (0) {
__t4 = uint32(Call_fromMaybe__18840980(gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_pred__2010692236(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}))).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Date)(v2_4.UnsafePtr).V1
}
end_branch_4:
l_7_3 := gopurs_runtime.Int(Call_lastDayOfMonth((*Constructor_Date)(v2_4.UnsafePtr).V0, __t4))
_ = l_7_3
hi_8_5 := gopurs_runtime.Bool(Call_greaterThan__4087042607(j_5_1, gopurs_runtime.Apply(Get_fromEnum__3599151655(), l_7_3)))
_ = hi_8_5
var __t6 int64
{
if (low_6_2.IntVal) != (0) {
__t6 = j_5_1.IntVal
goto end_branch_6
} else {

}
}
{
if (hi_8_5.IntVal) != (0) {
__t6 = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply2(Get_sub__1043827704(), j_5_1, gopurs_runtime.Apply(Get_fromEnum__3599151655(), l_7_3)), gopurs_runtime.Int(1)).IntVal
goto end_branch_6
} else {

}
}
{
__t6 = 0
}
end_branch_6:
var __t7 *pkg_Data_Maybe.Constructor_Just[*Constructor_Date]
{
if (low_6_2.IntVal) != (0) {
__t7 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply2(Get_map__3447677596(), gopurs_runtime.Apply2(Get_Date(), gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.Int(1))), gopurs_runtime.RecordGet(Get_enumDate(), "pred")))
goto end_branch_7
} else {

}
}
{
if (hi_8_5.IntVal) != (0) {
__t7 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply(Get_succ__2858180024(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, (*Constructor_Date)(v2_4.UnsafePtr).V0, (*Constructor_Date)(v2_4.UnsafePtr).V1, l_7_3.IntVal})}))
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(Get_map__3447677596(), gopurs_runtime.Apply2(Get_Date(), gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Apply(Get_toEnum__2099864294(), j_5_1)))
}
end_branch_7:
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](Call_bindFlipped__1454086721(gopurs_runtime.Apply(adj_2_0_0, gopurs_runtime.Int(__t6)), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)})))}
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](__t8))}
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(Get_bind__4062037089(), gopurs_runtime.Apply(pkg_Data_Int.Get_fromNumber(), gopurs_runtime.Float(v_0)), gopurs_runtime.Apply2(Get_flip__1826582752(), adj_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(date_1)})))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bindFlipped__1485397639(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_bindFlipped__3917280577(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), a_1, b_0)
}

func Call_bindFlipped__1454086721(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), a_1, b_0)
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

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_const__220790420(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__1826582752(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
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

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_fromJust__1577979644(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_fromJust__3809843644(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_fromMaybe__18840980(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_isNothing__1358705270(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__4206805139(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isNothing__2787066607(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
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

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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

func Call_greaterThan__2400628110(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](Get_ordMaybe()).V1, a1_0, a2_1)
if (uint32(__t_tag_0.IntVal) == 380165415) {
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

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
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

func Call_negate__2635823316(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Int(-(a_0.IntVal))
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

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_calcDiff() gopurs_runtime.Value {
	return _Gopurs_CalcDiff
}

func Get_calcWeekday() gopurs_runtime.Value {
	return _Gopurs_CalcWeekday
}

func Get_canonicalDateImpl() gopurs_runtime.Value {
	return _Gopurs_CanonicalDateImpl
}
