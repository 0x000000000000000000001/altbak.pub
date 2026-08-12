package Data_Date

import (
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
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
var __t2 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t2 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
return gopurs_runtime.Bool((__t2.IntVal) != (0))
})
}))
_ = eqMaybe1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t5 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordDay(), "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t4.IntVal)), UnsafePtr: nil}
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
		cache_weekday = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_weekday(gopurs_runtime.CoerceToStruct[Constructor_Date](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_weekday
}

var cache_showDate gopurs_runtime.Value
var once_showDate sync.Once
func Get_showDate() gopurs_runtime.Value {
	once_showDate.Do(func() {
		cache_showDate = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Date "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showYear(), "show"), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showMonth(), "show"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_showDay(), "show"), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V2)), gopurs_runtime.Str(")"))))))).StrVal())
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqYear(), "eq"), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqMonth(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_eqDay(), "eq"), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V2))).IntVal) != (0))
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
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordYear(), "compare"), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V0))
_ = v_2_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordMonth(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil})
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordDay(), "compare"), gopurs_runtime.Int((*Constructor_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Date)(y_1.UnsafePtr).V2)).IntVal)), UnsafePtr: nil}
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t2.IntVal)), UnsafePtr: nil}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t3.IntVal)), UnsafePtr: nil}
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
pm_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})
_ = pm_1_0
pd_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "pred"), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V2))
_ = pd_2_1
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136 && pd_2_1.UnsafePtr == nil) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136 && pd_2_1.UnsafePtr != nil) {
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
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 930809136 && pm_1_0.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 930809136 && pm_1_0.UnsafePtr != nil) {
__t5 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(pm_1_0.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t5.IntVal)), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_3:
m_prime_3_2 := __t3
_ = m_prime_3_2
l_4_6 := gopurs_runtime.Int(Call_lastDayOfMonth(gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0).IntVal, uint32(m_prime_3_2.IntVal)))
_ = l_4_6
var __t7 gopurs_runtime.Value
{
var __t8 gopurs_runtime.Value
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136 && pd_2_1.UnsafePtr == nil) {
__t8 = gopurs_runtime.Bool(true)
goto end_branch_8
} else {

}
}
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136 && pd_2_1.UnsafePtr != nil) {
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
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 930809136 && pm_1_0.UnsafePtr == nil) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if (pm_1_0.Type == 9 && pm_1_0.IntVal == 930809136 && pm_1_0.UnsafePtr != nil) {
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
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumYear(), "pred"), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0)})}))}
}
end_branch_7:
var __t10 gopurs_runtime.Value
{
var __t11 gopurs_runtime.Value
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136 && pd_2_1.UnsafePtr == nil) {
__t11 = gopurs_runtime.Bool(true)
goto end_branch_11
} else {

}
}
{
if (pd_2_1.Type == 9 && pd_2_1.IntVal == 930809136 && pd_2_1.UnsafePtr != nil) {
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
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, l_4_6})}))}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](pd_2_1))}
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Date(), __t7), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), m_prime_3_2)), __t10)))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
sm_1_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "succ"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})
_ = sm_1_12
v1_2_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumDay(), "succ"), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V2))
_ = v1_2_13
var __t15 gopurs_runtime.Value
{
var __t16 gopurs_runtime.Value
{
var __t_tag_17 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](Get_ordMaybe()).V1, v1_2_13, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_lastDayOfMonth(gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}.IntVal)))})})
if (uint32(__t_tag_17.IntVal) == 380165415) {
__t16 = gopurs_runtime.Bool(true)
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Bool(false)
}
end_branch_16:
if (__t16.IntVal) != (0) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v1_2_13))}
}
end_branch_15:
sd_3_14 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t15)
_ = sd_3_14
var __t18 gopurs_runtime.Value
{
var __t19 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.UnsafePtr == nil) {
__t19 = gopurs_runtime.Bool(true)
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.UnsafePtr != nil) {
__t19 = gopurs_runtime.Bool(false)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
var __t20 gopurs_runtime.Value
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 930809136 && sm_1_12.UnsafePtr == nil) {
__t20 = gopurs_runtime.Bool(true)
goto end_branch_20
} else {

}
}
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 930809136 && sm_1_12.UnsafePtr != nil) {
__t20 = gopurs_runtime.Bool(false)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t19, __t20).IntVal) != (0) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumYear(), "succ"), gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Date)(v_0.UnsafePtr).V0)})}))}
}
end_branch_18:
var __t21 gopurs_runtime.Value
{
var __t22 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.UnsafePtr == nil) {
__t22 = gopurs_runtime.Bool(true)
goto end_branch_22
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.UnsafePtr != nil) {
__t22 = gopurs_runtime.Bool(false)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
if (__t22.IntVal) != (0) {
var __t23 gopurs_runtime.Value
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 930809136 && sm_1_12.UnsafePtr == nil) {
__t23 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_23
} else {

}
}
{
if (sm_1_12.Type == 9 && sm_1_12.IntVal == 930809136 && sm_1_12.UnsafePtr != nil) {
__t23 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(sm_1_12.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t21 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t23.IntVal)), UnsafePtr: nil}
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_21:
var __t24 gopurs_runtime.Value
{
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.UnsafePtr == nil) {
__t25 = gopurs_runtime.Bool(true)
goto end_branch_25
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}.UnsafePtr != nil) {
__t25 = gopurs_runtime.Bool(false)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
if (__t25.IntVal) != (0) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(1))))}
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(sd_3_14)}
}
end_branch_24:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Date(), __t18), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applicativeMaybe(), "pure"), __t21)), __t24)))}
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
}), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Date](gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedYear(), "bottom").IntVal, uint32(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom").IntVal), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom").IntVal})}))}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Date](gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedYear(), "top").IntVal, uint32(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "top").IntVal), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "top").IntVal})}))})
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

type Constructor_Date struct {
	Rc uint32
	V0 int64
	V1 uint32
	V2 int64
}


func Call_year(v_0_loop *Constructor_Date) int64 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
return gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0).IntVal
}

func Call_weekday(v_0_loop *Constructor_Date) uint32 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
n_1_0 := gopurs_runtime.UncurriedApp3(Get_calcWeekday(), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2))
_ = n_1_0
var __t3 gopurs_runtime.Value
{
if (n_1_0.IntVal) == (0) {
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumWeekday(), "toEnum"), gopurs_runtime.Int(7))
_ = __local_var_2_4
var __t5 gopurs_runtime.Value
{
if (__local_var_2_4.Type == 9 && __local_var_2_4.IntVal == 930809136 && __local_var_2_4.UnsafePtr != nil) {
__t5 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_4.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t5.IntVal)), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumWeekday(), "toEnum"), n_1_0)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t2.IntVal)), UnsafePtr: nil}
}
end_branch_3:
return uint32(__t3.IntVal)
}

func Call_month(v_0_loop *Constructor_Date) uint32 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
return uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1), UnsafePtr: nil}.IntVal)
}

func Call_isLeapYear(y_0_loop int64) bool {
var y_0 int64 = y_0_loop
_ = y_0
y_prime_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "fromEnum"), gopurs_runtime.Int(y_0))
_ = y_prime_1_0
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), y_prime_1_0, gopurs_runtime.Int(4)).IntVal) == (0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), y_prime_1_0, gopurs_runtime.Int(400)).IntVal) == (0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), y_prime_1_0, gopurs_runtime.Int(100)).IntVal) == (0))))).IntVal) != (0)
}

func Call_lastDayOfMonth(y_0_loop int64, m_1_loop uint32) int64 {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var __t0 gopurs_runtime.Value
{
if (m_1 == 1908470532) {
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Int(__t2.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 2455627378) {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_isLeapYear(y_0)).IntVal) != (0) {
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(29))
_ = __local_var_2_6
var __t7 gopurs_runtime.Value
{
if (__local_var_2_6.Type == 9 && __local_var_2_6.IntVal == 930809136 && __local_var_2_6.UnsafePtr != nil) {
__t7 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_6.UnsafePtr).V0
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t5 = gopurs_runtime.Int(__t7.IntVal)
goto end_branch_5
} else {

}
}
{
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(28))
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136 && __local_var_2_3.UnsafePtr != nil) {
__t4 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Int(__t4.IntVal)
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (m_1 == 4162469099) {
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_8
var __t9 gopurs_runtime.Value
{
if (__local_var_2_8.Type == 9 && __local_var_2_8.IntVal == 930809136 && __local_var_2_8.UnsafePtr != nil) {
__t9 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_8.UnsafePtr).V0
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t0 = gopurs_runtime.Int(__t9.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 1692989816) {
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_10
var __t11 gopurs_runtime.Value
{
if (__local_var_2_10.Type == 9 && __local_var_2_10.IntVal == 930809136 && __local_var_2_10.UnsafePtr != nil) {
__t11 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_10.UnsafePtr).V0
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t0 = gopurs_runtime.Int(__t11.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 330658827) {
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_12
var __t13 gopurs_runtime.Value
{
if (__local_var_2_12.Type == 9 && __local_var_2_12.IntVal == 930809136 && __local_var_2_12.UnsafePtr != nil) {
__t13 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_12.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t0 = gopurs_runtime.Int(__t13.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 4067355978) {
__local_var_2_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_14
var __t15 gopurs_runtime.Value
{
if (__local_var_2_14.Type == 9 && __local_var_2_14.IntVal == 930809136 && __local_var_2_14.UnsafePtr != nil) {
__t15 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_14.UnsafePtr).V0
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t0 = gopurs_runtime.Int(__t15.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 2276710548) {
__local_var_2_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_16
var __t17 gopurs_runtime.Value
{
if (__local_var_2_16.Type == 9 && __local_var_2_16.IntVal == 930809136 && __local_var_2_16.UnsafePtr != nil) {
__t17 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_16.UnsafePtr).V0
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t0 = gopurs_runtime.Int(__t17.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 243771071) {
__local_var_2_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_18
var __t19 gopurs_runtime.Value
{
if (__local_var_2_18.Type == 9 && __local_var_2_18.IntVal == 930809136 && __local_var_2_18.UnsafePtr != nil) {
__t19 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_18.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t0 = gopurs_runtime.Int(__t19.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 215731793) {
__local_var_2_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_20
var __t21 gopurs_runtime.Value
{
if (__local_var_2_20.Type == 9 && __local_var_2_20.IntVal == 930809136 && __local_var_2_20.UnsafePtr != nil) {
__t21 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_20.UnsafePtr).V0
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t0 = gopurs_runtime.Int(__t21.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 8639228) {
__local_var_2_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_22
var __t23 gopurs_runtime.Value
{
if (__local_var_2_22.Type == 9 && __local_var_2_22.IntVal == 930809136 && __local_var_2_22.UnsafePtr != nil) {
__t23 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_22.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t0 = gopurs_runtime.Int(__t23.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 49471444) {
__local_var_2_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))
_ = __local_var_2_24
var __t25 gopurs_runtime.Value
{
if (__local_var_2_24.Type == 9 && __local_var_2_24.IntVal == 930809136 && __local_var_2_24.UnsafePtr != nil) {
__t25 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_24.UnsafePtr).V0
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t0 = gopurs_runtime.Int(__t25.IntVal)
goto end_branch_0
} else {

}
}
{
if (m_1 == 3889233761) {
__local_var_2_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))
_ = __local_var_2_26
var __t27 gopurs_runtime.Value
{
if (__local_var_2_26.Type == 9 && __local_var_2_26.IntVal == 930809136 && __local_var_2_26.UnsafePtr != nil) {
__t27 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_26.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
__t0 = gopurs_runtime.Int(__t27.IntVal)
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

func Call_diff(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], v_1_loop *Constructor_Date, v1_2_loop *Constructor_Date) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var v_1 *Constructor_Date = v_1_loop
_ = v_1
var v1_2 *Constructor_Date = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.UncurriedApp6(Get_calcDiff(), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V2)))
}

func Call_day(v_0_loop *Constructor_Date) int64 {
var v_0 *Constructor_Date = v_0_loop
_ = v_0
return gopurs_runtime.Int((*Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2).IntVal
}

func Call_canonicalDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) *Constructor_Date {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var d_2 int64 = d_2_loop
_ = d_2
return gopurs_runtime.CoerceToStruct[Constructor_Date](gopurs_runtime.UncurriedApp4(Get_canonicalDateImpl(), gopurs_runtime.Func(func(y_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum"), m_prime_4)
_ = __local_var_6_0
var __t1 gopurs_runtime.Value
{
if (__local_var_6_0.Type == 9 && __local_var_6_0.IntVal == 930809136 && __local_var_6_0.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_6_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Date](gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, y_prime_3.IntVal, uint32(__t1.IntVal), d_prime_5.IntVal})}))}
})
})
}), gopurs_runtime.Int(y_0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64(m_1), UnsafePtr: nil}), gopurs_runtime.Int(d_2)))
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
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqDate(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_canonicalDate(y_0, m_1, d_2))}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, y_0, m_1, d_2})}).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, y_0, m_1, d_2})}})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
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
var __t12 gopurs_runtime.Value
{
if (v1_3.IntVal) == (0) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v2_4})}))}
goto end_branch_12
} else {

}
}
{
j_5_1 := (v1_3.IntVal) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V2)).IntVal)
_ = j_5_1
var __t3 gopurs_runtime.Value
{
if (j_5_1) < (1) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
low_6_2 := __t3
_ = low_6_2
var __t5 gopurs_runtime.Value
{
if (low_6_2.IntVal) != (0) {
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_7_6
var __t7 gopurs_runtime.Value
{
if (__local_var_7_6.Type == 9 && __local_var_7_6.IntVal == 930809136 && __local_var_7_6.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_7
} else {

}
}
{
if (__local_var_7_6.Type == 9 && __local_var_7_6.IntVal == 930809136 && __local_var_7_6.UnsafePtr != nil) {
__t7 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_7_6.UnsafePtr).V0
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t7.IntVal)), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_5:
l_7_4 := gopurs_runtime.Int(Call_lastDayOfMonth(gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V0).IntVal, uint32(__t5.IntVal)))
_ = l_7_4
var __t9 gopurs_runtime.Value
{
if (j_5_1) > (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), l_7_4).IntVal) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Bool(false)
}
end_branch_9:
hi_8_8 := __t9
_ = hi_8_8
var __t10 gopurs_runtime.Value
{
if (low_6_2.IntVal) != (0) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply2(Get_Date(), gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(1))), gopurs_runtime.RecordGet(Get_enumDate(), "pred"))))}
goto end_branch_10
} else {

}
}
{
if (hi_8_8.IntVal) != (0) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_enumDate(), "succ"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Date{1, gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}.IntVal), l_7_4.IntVal})})))}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply2(Get_Date(), gopurs_runtime.Int((*Constructor_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(j_5_1)))))}
}
end_branch_10:
var __t11 gopurs_runtime.Value
{
if (low_6_2.IntVal) != (0) {
__t11 = gopurs_runtime.Int(j_5_1)
goto end_branch_11
} else {

}
}
{
if (hi_8_8.IntVal) != (0) {
__t11 = gopurs_runtime.Int(((j_5_1) - (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), l_7_4).IntVal)) - (1))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Int(0)
}
end_branch_11:
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), __t10, gopurs_runtime.Apply(adj_2_0_0, __t11))))}
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](__t12))}
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Date]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(pkg_Data_Int.Get_fromNumber(), gopurs_runtime.Float(v_0)), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(adj_2_0_0, a_3, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(date_1)})
})))
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
