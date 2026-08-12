package Data_DateTime

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_DateTime gopurs_runtime.Value
var once_DateTime sync.Once
func Get_DateTime() gopurs_runtime.Value {
	once_DateTime.Do(func() {
		cache_DateTime = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](value0), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](value1)})}
})
})
	})
	return cache_DateTime
}

var cache_toRecord gopurs_runtime.Value
var once_toRecord sync.Once
func Get_toRecord() gopurs_runtime.Value {
	once_toRecord.Do(func() {
		cache_toRecord = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toRecord(gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_0_box))
})
	})
	return cache_toRecord
}

var cache_time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		cache_time = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_time(gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_0_box)))}
})
	})
	return cache_time
}

var cache_showDateTime gopurs_runtime.Value
var once_showDateTime sync.Once
func Get_showDateTime() gopurs_runtime.Value {
	once_showDateTime.Do(func() {
		cache_showDateTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(DateTime "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date.Get_showDate(), "show"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(v_0.UnsafePtr).V0)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time.Get_showTime(), "show"), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(v_0.UnsafePtr).V1)}), gopurs_runtime.Str(")"))))).StrVal())
}))
	})
	return cache_showDateTime
}

var cache_modifyTimeF gopurs_runtime.Value
var once_modifyTimeF sync.Once
func Get_modifyTimeF() gopurs_runtime.Value {
	once_modifyTimeF.Do(func() {
		cache_modifyTimeF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyTimeF(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_2_box))
})
	})
	return cache_modifyTimeF
}

var cache_modifyTime gopurs_runtime.Value
var once_modifyTime sync.Once
func Get_modifyTime() gopurs_runtime.Value {
	once_modifyTime.Do(func() {
		cache_modifyTime = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_modifyTime(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_1_box)))}
})
	})
	return cache_modifyTime
}

var cache_modifyDateF gopurs_runtime.Value
var once_modifyDateF sync.Once
func Get_modifyDateF() gopurs_runtime.Value {
	once_modifyDateF.Do(func() {
		cache_modifyDateF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyDateF(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_2_box))
})
	})
	return cache_modifyDateF
}

var cache_modifyDate gopurs_runtime.Value
var once_modifyDate sync.Once
func Get_modifyDate() gopurs_runtime.Value {
	once_modifyDate.Do(func() {
		cache_modifyDate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_modifyDate(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_1_box)))}
})
	})
	return cache_modifyDate
}

var cache_eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		cache_eqDateTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date.Get_eqDate(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V0)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time.Get_eqTime(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V1)})).IntVal) != (0))
})
}))
	})
	return cache_eqDateTime
}

var cache_ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		cache_ordDateTime = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDateTime()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date.Get_ordDate(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V0)})
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time.Get_ordTime(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V1)}).IntVal)), UnsafePtr: nil}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t1.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordDateTime
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_DateTime](dt1_1_box), gopurs_runtime.CoerceToStruct[Constructor_DateTime](dt2_2_box))
})
	})
	return cache_diff
}

var cache_date gopurs_runtime.Value
var once_date sync.Once
func Get_date() gopurs_runtime.Value {
	once_date.Do(func() {
		cache_date = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_date(gopurs_runtime.CoerceToStruct[Constructor_DateTime](v_0_box)))}
})
	})
	return cache_date
}

var cache_boundedDateTime gopurs_runtime.Value
var once_boundedDateTime sync.Once
func Get_boundedDateTime() gopurs_runtime.Value {
	once_boundedDateTime.Do(func() {
		cache_boundedDateTime = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDateTime()
}), gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.RecordGet(pkg_Data_Date.Get_boundedDate(), "bottom")), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.RecordGet(pkg_Data_Time.Get_boundedTime(), "bottom"))})}))}, gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.RecordGet(pkg_Data_Date.Get_boundedDate(), "top")), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.RecordGet(pkg_Data_Time.Get_boundedTime(), "top"))})}))})
	})
	return cache_boundedDateTime
}

var cache_adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		cache_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, dt_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_adjust(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), d_1_box, gopurs_runtime.CoerceToStruct[Constructor_DateTime](dt_2_box)))}
})
	})
	return cache_adjust
}

type Constructor_DateTime struct {
	Rc uint32
	V0 *pkg_Data_Date.Constructor_Date
	V1 *pkg_Data_Time.Constructor_Time
}


func Call_toRecord(v_0_loop *Constructor_DateTime) gopurs_runtime.Value {
var v_0 *Constructor_DateTime = v_0_loop
_ = v_0
return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V2)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumHour(), "fromEnum"), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V0)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMillisecond(), "fromEnum"), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V3)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMinute(), "fromEnum"), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V1), UnsafePtr: nil}).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumSecond(), "fromEnum"), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V2)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "fromEnum"), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V0)).IntVal)})
}

func Call_time(v_0_loop *Constructor_DateTime) *pkg_Data_Time.Constructor_Time {
var v_0 *Constructor_DateTime = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})
}

func Call_modifyTimeF(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_DateTime = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply(Get_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0)}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)}))
}

func Call_modifyTime(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_DateTime) *Constructor_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_DateTime = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)}), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}))})})
}

func Call_modifyDateF(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_DateTime = v_2_loop
_ = v_2
__local_var_3_0 := gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)}
_ = __local_var_3_0
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](a_4), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](__local_var_3_0)})}
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0)}))
}

func Call_modifyDate(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_DateTime) *Constructor_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_DateTime = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)})), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})})
}

func Call_diff(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], dt1_1_loop *Constructor_DateTime, dt2_2_loop *Constructor_DateTime) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 *Constructor_DateTime = dt1_1_loop
_ = dt1_1
var dt2_2 *Constructor_DateTime = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.UncurriedApp2(Get_calcDiff(), Call_toRecord(dt1_1), Call_toRecord(dt2_2)))
}

func Call_date(v_0_loop *Constructor_DateTime) *pkg_Data_Date.Constructor_Date {
var v_0 *Constructor_DateTime = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)})
}

func Call_adjust(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], d_1_loop gopurs_runtime.Value, dt_2_loop *Constructor_DateTime) *pkg_Data_Maybe.Constructor_Just[*Constructor_DateTime] {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var dt_2 *Constructor_DateTime = dt_2_loop
_ = dt_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_DateTime]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply4(Get_adjustImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Apply(dictDuration_0.V0, d_1), Call_toRecord(dt_2)), gopurs_runtime.Func(func(rec_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_DateTime]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_DateTime(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Date.Get_exactDate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "year"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumMonth(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "month"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "day"))), pkg_Control_Bind.Get_identity())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Time.Get_Time(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumHour(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "hour"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMinute(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "minute"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumSecond(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "second"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMillisecond(), "toEnum"), gopurs_runtime.RecordGet(rec_3, "millisecond"))))))}
})))
}

func Get_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_AdjustImpl
}

func Get_calcDiff() gopurs_runtime.Value {
	return _Gopurs_CalcDiff
}
