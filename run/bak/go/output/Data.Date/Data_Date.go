package Data_Date

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Int "gopurs/output/Data.Int"
	unsafe "unsafe"
)

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_greaterThan(a1_0_box, a2_1_box)
})
	})
	return cache_greaterThan
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

var cache_greaterThan1 gopurs_runtime.Value
var once_greaterThan1 sync.Once
func Get_greaterThan1() gopurs_runtime.Value {
	once_greaterThan1.Do(func() {
		cache_greaterThan1 = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
})
}()
	})
	return cache_greaterThan1
}

var cache_Date gopurs_runtime.Value
var once_Date sync.Once
func Get_Date() gopurs_runtime.Value {
	once_Date.Do(func() {
		cache_Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{value0, value1, value2})}
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
return Call_year(v_0_box)
})
	})
	return cache_year
}

var cache_weekday gopurs_runtime.Value
var once_weekday sync.Once
func Get_weekday() gopurs_runtime.Value {
	once_weekday.Do(func() {
		cache_weekday = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_weekday(v_0_box)
})
	})
	return cache_weekday
}

var cache_showDate gopurs_runtime.Value
var once_showDate sync.Once
func Get_showDate() gopurs_runtime.Value {
	once_showDate.Do(func() {
		cache_showDate = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Date "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showYear(), "show"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showMonth(), "show"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showDay(), "show"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V2), gopurs_runtime.Str(")")))))))
}))
	})
	return cache_showDate
}

var cache_month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		cache_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_month(v_0_box)
})
	})
	return cache_month
}

var cache_isLeapYear gopurs_runtime.Value
var once_isLeapYear sync.Once
func Get_isLeapYear() gopurs_runtime.Value {
	once_isLeapYear.Do(func() {
		cache_isLeapYear = gopurs_runtime.Func(func(y_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_isLeapYear(y_0_box)
})
	})
	return cache_isLeapYear
}

var cache_lastDayOfMonth gopurs_runtime.Value
var once_lastDayOfMonth sync.Once
func Get_lastDayOfMonth() gopurs_runtime.Value {
	once_lastDayOfMonth.Do(func() {
		cache_lastDayOfMonth = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lastDayOfMonth(y_0_box, m_1_box)
})
	})
	return cache_lastDayOfMonth
}

var cache_eqDate gopurs_runtime.Value
var once_eqDate sync.Once
func Get_eqDate() gopurs_runtime.Value {
	once_eqDate.Do(func() {
		cache_eqDate = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqYear(), "eq"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V0, (*Data_Data_Date_Date)(y_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqMonth(), "eq"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V1, (*Data_Data_Date_Date)(y_1.UnsafePtr).V1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqDay(), "eq"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V2, (*Data_Data_Date_Date)(y_1.UnsafePtr).V2))
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
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordYear(), "compare"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V0, (*Data_Data_Date_Date)(y_1.UnsafePtr).V0)
_ = v_2_0
var __t3 gopurs_runtime.Value
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordMonth(), "compare"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V1, (*Data_Data_Date_Date)(y_1.UnsafePtr).V1)
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordDay(), "compare"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V2, (*Data_Data_Date_Date)(y_1.UnsafePtr).V2)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
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
pm_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V1)
_ = pm_1_0
pd_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "pred"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V2)
_ = pd_2_1
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 3589588149) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136) {
__t4 = gopurs_runtime.Bool(false)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
if (__t4.IntVal) != (0) {
var __t5 gopurs_runtime.Value
{
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 3589588149) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 930809136) {
__t5 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(pm_1_0.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__t3 = (*Data_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_3:
m_prime_3_2 := __t3
_ = m_prime_3_2
l_4_6 := Call_lastDayOfMonth((*Data_Data_Date_Date)(v_0.UnsafePtr).V0, m_prime_3_2)
_ = l_4_6
var __t7 gopurs_runtime.Value
{
var __t8 gopurs_runtime.Value
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 3589588149) {
__t8 = gopurs_runtime.Bool(true)
goto end_branch_8
} else {

}
}
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136) {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
var __t9 gopurs_runtime.Value
{
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 3589588149) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 930809136) {
__t9 = gopurs_runtime.Bool(false)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t8, __t9).IntVal) != (0) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumYear(), "pred"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Date_Date)(v_0.UnsafePtr).V0})}
}
end_branch_7:
var __t10 gopurs_runtime.Value
{
var __t11 gopurs_runtime.Value
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 3589588149) {
__t11 = gopurs_runtime.Bool(true)
goto end_branch_11
} else {

}
}
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136) {
__t11 = gopurs_runtime.Bool(false)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
if (__t11.IntVal) != (0) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{l_4_6})}
goto end_branch_10
} else {

}
}
{
__t10 = pd_2_1
}
end_branch_10:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Date(), __t7), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), m_prime_3_2)), __t10)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
sm_1_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "succ"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V1)
_ = sm_1_12
v1_2_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "succ"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V2)
_ = v1_2_13
var __t15 gopurs_runtime.Value
{
if (Call_greaterThan(v1_2_13, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{Call_lastDayOfMonth((*Data_Data_Date_Date)(v_0.UnsafePtr).V0, (*Data_Data_Date_Date)(v_0.UnsafePtr).V1)})}).IntVal) != (0) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_15
} else {

}
}
{
__t15 = v1_2_13
}
end_branch_15:
sd_3_14 := __t15
_ = sd_3_14
var __t16 gopurs_runtime.Value
{
var __t17 gopurs_runtime.Value
{
if (sd_3_14.Type == 9 && sd_3_14.IntVal == 3589588149) {
__t17 = gopurs_runtime.Bool(true)
goto end_branch_17
} else {

}
}
{
if (sd_3_14.Type == 9 && sd_3_14.IntVal == 930809136) {
__t17 = gopurs_runtime.Bool(false)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
var __t18 gopurs_runtime.Value
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 3589588149) {
__t18 = gopurs_runtime.Bool(true)
goto end_branch_18
} else {

}
}
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 930809136) {
__t18 = gopurs_runtime.Bool(false)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t17, __t18).IntVal) != (0) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumYear(), "succ"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Data_Date_Date)(v_0.UnsafePtr).V0})}
}
end_branch_16:
var __t19 gopurs_runtime.Value
{
var __t20 gopurs_runtime.Value
{
if (sd_3_14.Type == 9 && sd_3_14.IntVal == 3589588149) {
__t20 = gopurs_runtime.Bool(true)
goto end_branch_20
} else {

}
}
{
if (sd_3_14.Type == 9 && sd_3_14.IntVal == 930809136) {
__t20 = gopurs_runtime.Bool(false)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
if (__t20.IntVal) != (0) {
var __t21 gopurs_runtime.Value
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 3589588149) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: nil}
goto end_branch_21
} else {

}
}
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 930809136) {
__t21 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(sm_1_12.UnsafePtr).V0
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t19 = __t21
goto end_branch_19
} else {

}
}
{
__t19 = (*Data_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_19:
var __t22 gopurs_runtime.Value
{
var __t23 gopurs_runtime.Value
{
if (sd_3_14.Type == 9 && sd_3_14.IntVal == 3589588149) {
__t23 = gopurs_runtime.Bool(true)
goto end_branch_23
} else {

}
}
{
if (sd_3_14.Type == 9 && sd_3_14.IntVal == 930809136) {
__t23 = gopurs_runtime.Bool(false)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
if (__t23.IntVal) != (0) {
__t22 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(1))
goto end_branch_22
} else {

}
}
{
__t22 = sd_3_14
}
end_branch_22:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Date(), __t16), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), __t19)), __t22)
}))
	})
	return cache_enumDate
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff((*Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value)(dictDuration_0_box.UnsafePtr), v_1_box, v1_2_box)
})
	})
	return cache_diff
}

var cache_day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		cache_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_day(v_0_box)
})
	})
	return cache_day
}

var cache_canonicalDate gopurs_runtime.Value
var once_canonicalDate sync.Once
func Get_canonicalDate() gopurs_runtime.Value {
	once_canonicalDate.Do(func() {
		cache_canonicalDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_canonicalDate(y_0_box, m_1_box, d_2_box)
})
	})
	return cache_canonicalDate
}

var cache_exactDate gopurs_runtime.Value
var once_exactDate sync.Once
func Get_exactDate() gopurs_runtime.Value {
	once_exactDate.Do(func() {
		cache_exactDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_exactDate(y_0_box, m_1_box, d_2_box)
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
}), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedYear(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedYear(), "top"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "top"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "top")})})
	})
	return cache_boundedDate
}

var cache_adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		cache_adjust = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, date_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_adjust(v_0_box, date_1_box)
})
	})
	return cache_adjust
}

type Data_Data_Date_Date struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_Date_Date(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 745776346
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

func Call_greaterThan(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 gopurs_runtime.Value
{
if (a1_0.Type == 9 && a1_0.IntVal == 3589588149) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (a2_1.Type == 9 && a2_1.IntVal == 3589588149) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if ((a1_0.Type == 9 && a1_0.IntVal == 930809136)) && ((a2_1.Type == 9 && a2_1.IntVal == 930809136)) {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordDay(), "compare"), (*pkg_Data_Maybe.Data_Data_Maybe_Just)(a1_0.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(a2_1.UnsafePtr).V0)
__t0 = gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
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

func Call_year(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Date_Date)(v_0.UnsafePtr).V0
}

func Call_weekday(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
n_1_0 := gopurs_runtime.UncurriedApp3(Get_calcWeekday(), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V1), (*Data_Data_Date_Date)(v_0.UnsafePtr).V2)
_ = n_1_0
var __t3 gopurs_runtime.Value
{
if (n_1_0.IntVal) == (0) {
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumWeekday(), "toEnum"), gopurs_runtime.Int(7))
_ = __local_var_2_4
var __t5 gopurs_runtime.Value
{
if (__local_var_2_4.Type == 9 && __local_var_2_4.IntVal == 930809136) {
__t5 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_4.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumWeekday(), "toEnum"), n_1_0)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
}

func Call_month(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Date_Date)(v_0.UnsafePtr).V1
}

func Call_isLeapYear(y_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
y_prime_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "fromEnum"), y_0)
_ = y_prime_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), y_prime_1_0, gopurs_runtime.Int(4)).IntVal) == (0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), y_prime_1_0, gopurs_runtime.Int(400)).IntVal) == (0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), y_prime_1_0, gopurs_runtime.Int(100)).IntVal) == (0)))))
}

func Call_lastDayOfMonth(y_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 1908470532) {
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0
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
if (m_1.Type == 9 && m_1.IntVal == 2455627378) {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(Get_isLeapYear(), y_0).IntVal) != (0) {
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(29))
_ = __local_var_2_6
var __t7 gopurs_runtime.Value
{
if (__local_var_2_6.Type == 9 && __local_var_2_6.IntVal == 930809136) {
__t7 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_6.UnsafePtr).V0
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
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(28))
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136) {
__t4 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 4162469099) {
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_8
var __t9 gopurs_runtime.Value
{
if (__local_var_2_8.Type == 9 && __local_var_2_8.IntVal == 930809136) {
__t9 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_8.UnsafePtr).V0
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t0 = __t9
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 1692989816) {
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_10
var __t11 gopurs_runtime.Value
{
if (__local_var_2_10.Type == 9 && __local_var_2_10.IntVal == 930809136) {
__t11 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_10.UnsafePtr).V0
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t0 = __t11
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 330658827) {
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_12
var __t13 gopurs_runtime.Value
{
if (__local_var_2_12.Type == 9 && __local_var_2_12.IntVal == 930809136) {
__t13 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_12.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t0 = __t13
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 4067355978) {
__local_var_2_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_14
var __t15 gopurs_runtime.Value
{
if (__local_var_2_14.Type == 9 && __local_var_2_14.IntVal == 930809136) {
__t15 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_14.UnsafePtr).V0
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t0 = __t15
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2276710548) {
__local_var_2_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_16
var __t17 gopurs_runtime.Value
{
if (__local_var_2_16.Type == 9 && __local_var_2_16.IntVal == 930809136) {
__t17 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_16.UnsafePtr).V0
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t0 = __t17
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 243771071) {
__local_var_2_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_18
var __t19 gopurs_runtime.Value
{
if (__local_var_2_18.Type == 9 && __local_var_2_18.IntVal == 930809136) {
__t19 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_18.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t0 = __t19
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 215731793) {
__local_var_2_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_20
var __t21 gopurs_runtime.Value
{
if (__local_var_2_20.Type == 9 && __local_var_2_20.IntVal == 930809136) {
__t21 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_20.UnsafePtr).V0
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t0 = __t21
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 8639228) {
__local_var_2_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_22
var __t23 gopurs_runtime.Value
{
if (__local_var_2_22.Type == 9 && __local_var_2_22.IntVal == 930809136) {
__t23 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_22.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t0 = __t23
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 49471444) {
__local_var_2_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_24
var __t25 gopurs_runtime.Value
{
if (__local_var_2_24.Type == 9 && __local_var_2_24.IntVal == 930809136) {
__t25 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_24.UnsafePtr).V0
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t0 = __t25
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 3889233761) {
__local_var_2_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_26
var __t27 gopurs_runtime.Value
{
if (__local_var_2_26.Type == 9 && __local_var_2_26.IntVal == 930809136) {
__t27 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_26.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
__t0 = __t27
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

func Call_diff(dictDuration_0_loop *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value = dictDuration_0_loop
_ = dictDuration_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply(dictDuration_0.toDuration, gopurs_runtime.UncurriedApp6(Get_calcDiff(), (*Data_Data_Date_Date)(v_1.UnsafePtr).V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), (*Data_Data_Date_Date)(v_1.UnsafePtr).V1), (*Data_Data_Date_Date)(v_1.UnsafePtr).V2, (*Data_Data_Date_Date)(v1_2.UnsafePtr).V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1), (*Data_Data_Date_Date)(v1_2.UnsafePtr).V2))
}

func Call_day(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Date_Date)(v_0.UnsafePtr).V2
}

func Call_canonicalDate(y_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, d_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
return gopurs_runtime.UncurriedApp4(Get_canonicalDateImpl(), gopurs_runtime.Func3(func(y_prime_3 gopurs_runtime.Value, m_prime_4 gopurs_runtime.Value, d_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum"), m_prime_4)
_ = __local_var_6_0
var __t1 gopurs_runtime.Value
{
if (__local_var_6_0.Type == 9 && __local_var_6_0.IntVal == 930809136) {
__t1 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_6_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{y_prime_3, __t1, d_prime_5})}
}), y_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), m_1), d_2)
}

func Call_exactDate(y_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, d_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqDate(), "eq"), Call_canonicalDate(y_0, m_1, d_2), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{y_0, m_1, d_2})}).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{y_0, m_1, d_2})}})}
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

func Call_adjust(v_0_loop gopurs_runtime.Value, date_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var date_1 gopurs_runtime.Value = date_1_loop
_ = date_1
var adj_2_0 gopurs_runtime.Value
_ = adj_2_0
adj_2_0 = gopurs_runtime.Func2(func(v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v1_3.IntVal) == (0) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v2_4})}
goto end_branch_10
} else {

}
}
{
j_5_1 := (v1_3.IntVal) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), (*Data_Data_Date_Date)(v2_4.UnsafePtr).V2).IntVal)
_ = j_5_1
low_6_2 := gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(j_5_1), gopurs_runtime.Int(1))
_ = low_6_2
var __t4 gopurs_runtime.Value
{
if (low_6_2.IntVal) != (0) {
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1)
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if (__local_var_7_5.Type == 9 && __local_var_7_5.IntVal == 3589588149) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (__local_var_7_5.Type == 9 && __local_var_7_5.IntVal == 930809136) {
__t6 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_7_5.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
__t4 = (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_4:
l_7_3 := Call_lastDayOfMonth((*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, __t4)
_ = l_7_3
hi_8_7 := gopurs_runtime.Apply2(Get_greaterThan1(), gopurs_runtime.Int(j_5_1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), l_7_3))
_ = hi_8_7
var __t8 gopurs_runtime.Value
{
if (low_6_2.IntVal) != (0) {
__t8 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply2(Get_Date(), (*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(1))), gopurs_runtime.RecordGet(Get_enumDate(), "pred"))
goto end_branch_8
} else {

}
}
{
if (hi_8_7.IntVal) != (0) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_enumDate(), "succ"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{(*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1, l_7_3})})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply2(Get_Date(), (*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(j_5_1)))
}
end_branch_8:
var __t9 gopurs_runtime.Value
{
if (low_6_2.IntVal) != (0) {
__t9 = gopurs_runtime.Int(j_5_1)
goto end_branch_9
} else {

}
}
{
if (hi_8_7.IntVal) != (0) {
__t9 = gopurs_runtime.Int(((j_5_1) - (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), l_7_3).IntVal)) - (1))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Int(0)
}
end_branch_9:
__t10 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), __t8, gopurs_runtime.Apply(adj_2_0, __t9))
}
end_branch_10:
return __t10
})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(pkg_Data_Int.Get_fromNumber(), v_0), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(adj_2_0, a_3, date_1)
}))
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
