package Data_DateTime_Instant

import (
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
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
		cache_toDateTime = gopurs_runtime.Apply(Get_toDateTimeImpl(), gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mo_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mi_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ms_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum"), mo_1)
_ = __local_var_7_0
var __t1 gopurs_runtime.Value
{
if (__local_var_7_0.Type == 9 && __local_var_7_0.IntVal == 930809136 && __local_var_7_0.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_7_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_DateTime.Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&pkg_Data_DateTime.Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Apply3(pkg_Data_Date.Get_canonicalDate(), y_0, __t1, d_2)), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&pkg_Data_Time.Constructor_Time{1, h_3.IntVal, mi_4.IntVal, s_5.IntVal, ms_6.IntVal})})})}))}
})
})
})
})
})
})
}))
	})
	return cache_toDateTime
}

var cache_showInstant gopurs_runtime.Value
var once_showInstant sync.Once
func Get_showInstant() gopurs_runtime.Value {
	once_showInstant.Do(func() {
		cache_showInstant = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Instant "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_showMilliseconds(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
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
}), gopurs_runtime.Float(-8639977881600000.0), gopurs_runtime.Float(8639977881599999.0))
	})
	return cache_boundedInstant
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
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordNumber(), "compare"), gopurs_runtime.Float(v_0), gopurs_runtime.Float(-8639977881600000.0))
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
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordNumber(), "compare"), gopurs_runtime.Float(v_0), gopurs_runtime.Float(8639977881599999.0))
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 380165415) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(true)
}
end_branch_3:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t1, __t3).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[float64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(v_0)})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[float64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[float64]](__t0)
}

func Call_fromDateTime(v_0_loop *pkg_Data_DateTime.Constructor_DateTime) float64 {
var v_0 *pkg_Data_DateTime.Constructor_DateTime = v_0_loop
_ = v_0
return gopurs_runtime.UncurriedApp7(Get_fromDateTimeImpl(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), (*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V1), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V2), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V2), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*pkg_Data_DateTime.Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V3)).FloatVal()
}

func Call_fromDate(d_0_loop *pkg_Data_Date.Constructor_Date) float64 {
var d_0 *pkg_Data_Date.Constructor_Date = d_0_loop
_ = d_0
return gopurs_runtime.UncurriedApp7(Get_fromDateTimeImpl(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(d_0)}.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), (*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(d_0)}.UnsafePtr).V1), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(d_0)}.UnsafePtr).V2), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "bottom")).FloatVal()
}

func Call_diff(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], dt1_1_loop float64, dt2_2_loop float64) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 float64 = dt1_1_loop
_ = dt1_1
var dt2_2 float64 = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), gopurs_runtime.Float(dt1_1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Float(-(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), gopurs_runtime.Float(dt2_2)).FloatVal())))))
}

func Get_fromDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_FromDateTimeImpl
}

func Get_toDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_ToDateTimeImpl
}
