package Data_DateTime

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
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
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(DateTime "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(Get_show__1723386194(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(v_0.UnsafePtr).V0)}), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(Get_show__1073032466(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(v_0.UnsafePtr).V1)}), gopurs_runtime.Str(")"))))).StrVal())
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply2(Get_eq__1204755874(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V0)}), gopurs_runtime.Apply2(Get_eq__1287514754(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V1)})).IntVal) != (0))
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
v_2_0 := gopurs_runtime.Apply2(Get_compare__146529112(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V0)})
_ = v_2_0
var __t1 uint32
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t1 = 1527465420
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t1 = 380165415
goto end_branch_1
} else {

}
}
{
__t1 = uint32(gopurs_runtime.Apply2(Get_compare__463614392(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(y_1.UnsafePtr).V1)}).IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
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
}), gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.RecordGet(pkg_Data_Date.Get_boundedDate(), "bottom")), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.RecordGet(pkg_Data_Time.Get_boundedTime(), "bottom"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.RecordGet(pkg_Data_Date.Get_boundedDate(), "top")), gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.RecordGet(pkg_Data_Time.Get_boundedTime(), "top"))})})
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

var cache_apply__420223018 gopurs_runtime.Value
var once_apply__420223018 sync.Once
func Get_apply__420223018() gopurs_runtime.Value {
	once_apply__420223018.Do(func() {
		cache_apply__420223018 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__420223018
}

var cache_apply__3882563466 gopurs_runtime.Value
var once_apply__3882563466 sync.Once
func Get_apply__3882563466() gopurs_runtime.Value {
	once_apply__3882563466.Do(func() {
		cache_apply__3882563466 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3882563466
}

var cache_apply__3867059818 gopurs_runtime.Value
var once_apply__3867059818 sync.Once
func Get_apply__3867059818() gopurs_runtime.Value {
	once_apply__3867059818.Do(func() {
		cache_apply__3867059818 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3867059818
}

var cache_apply__3489442218 gopurs_runtime.Value
var once_apply__3489442218 sync.Once
func Get_apply__3489442218() gopurs_runtime.Value {
	once_apply__3489442218.Do(func() {
		cache_apply__3489442218 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3489442218
}

var cache_apply__1183285642 gopurs_runtime.Value
var once_apply__1183285642 sync.Once
func Get_apply__1183285642() gopurs_runtime.Value {
	once_apply__1183285642.Do(func() {
		cache_apply__1183285642 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__1183285642
}

var cache_apply__3534390890 gopurs_runtime.Value
var once_apply__3534390890 sync.Once
func Get_apply__3534390890() gopurs_runtime.Value {
	once_apply__3534390890.Do(func() {
		cache_apply__3534390890 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3534390890
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

var cache_bind__737692327 gopurs_runtime.Value
var once_bind__737692327 sync.Once
func Get_bind__737692327() gopurs_runtime.Value {
	once_bind__737692327.Do(func() {
		cache_bind__737692327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__737692327(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__737692327
}

var cache_bind__1702199617 gopurs_runtime.Value
var once_bind__1702199617 sync.Once
func Get_bind__1702199617() gopurs_runtime.Value {
	once_bind__1702199617.Do(func() {
		cache_bind__1702199617 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind")
	})
	return cache_bind__1702199617
}

var cache_join__1635241211 gopurs_runtime.Value
var once_join__1635241211 sync.Once
func Get_join__1635241211() gopurs_runtime.Value {
	once_join__1635241211.Do(func() {
		cache_join__1635241211 = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join__1635241211(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), m_1_box)
})
	})
	return cache_join__1635241211
}

var cache_join__880516349 gopurs_runtime.Value
var once_join__880516349 sync.Once
func Get_join__880516349() gopurs_runtime.Value {
	once_join__880516349.Do(func() {
		cache_join__880516349 = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join__880516349(m_0_box)
})
	})
	return cache_join__880516349
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
		cache_eq__1204755874 = gopurs_runtime.RecordGet(pkg_Data_Date.Get_eqDate(), "eq")
	})
	return cache_eq__1204755874
}

var cache_eq__1287514754 gopurs_runtime.Value
var once_eq__1287514754 sync.Once
func Get_eq__1287514754() gopurs_runtime.Value {
	once_eq__1287514754.Do(func() {
		cache_eq__1287514754 = gopurs_runtime.RecordGet(pkg_Data_Time.Get_eqTime(), "eq")
	})
	return cache_eq__1287514754
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

var cache_flip__682087456 gopurs_runtime.Value
var once_flip__682087456 sync.Once
func Get_flip__682087456() gopurs_runtime.Value {
	once_flip__682087456.Do(func() {
		cache_flip__682087456 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__682087456(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__682087456
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

var cache_map__366271444 gopurs_runtime.Value
var once_map__366271444 sync.Once
func Get_map__366271444() gopurs_runtime.Value {
	once_map__366271444.Do(func() {
		cache_map__366271444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__366271444(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__366271444
}

var cache_map__3592908820 gopurs_runtime.Value
var once_map__3592908820 sync.Once
func Get_map__3592908820() gopurs_runtime.Value {
	once_map__3592908820.Do(func() {
		cache_map__3592908820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3592908820(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3592908820
}

var cache_map__2165919164 gopurs_runtime.Value
var once_map__2165919164 sync.Once
func Get_map__2165919164() gopurs_runtime.Value {
	once_map__2165919164.Do(func() {
		cache_map__2165919164 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2165919164
}

var cache_map__2389909756 gopurs_runtime.Value
var once_map__2389909756 sync.Once
func Get_map__2389909756() gopurs_runtime.Value {
	once_map__2389909756.Do(func() {
		cache_map__2389909756 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2389909756
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_map__63598588 gopurs_runtime.Value
var once_map__63598588 sync.Once
func Get_map__63598588() gopurs_runtime.Value {
	once_map__63598588.Do(func() {
		cache_map__63598588 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__63598588
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

var cache_compare__146529112 gopurs_runtime.Value
var once_compare__146529112 sync.Once
func Get_compare__146529112() gopurs_runtime.Value {
	once_compare__146529112.Do(func() {
		cache_compare__146529112 = gopurs_runtime.RecordGet(pkg_Data_Date.Get_ordDate(), "compare")
	})
	return cache_compare__146529112
}

var cache_compare__463614392 gopurs_runtime.Value
var once_compare__463614392 sync.Once
func Get_compare__463614392() gopurs_runtime.Value {
	once_compare__463614392.Do(func() {
		cache_compare__463614392 = gopurs_runtime.RecordGet(pkg_Data_Time.Get_ordTime(), "compare")
	})
	return cache_compare__463614392
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

var cache_show__1723386194 gopurs_runtime.Value
var once_show__1723386194 sync.Once
func Get_show__1723386194() gopurs_runtime.Value {
	once_show__1723386194.Do(func() {
		cache_show__1723386194 = gopurs_runtime.RecordGet(pkg_Data_Date.Get_showDate(), "show")
	})
	return cache_show__1723386194
}

var cache_show__1073032466 gopurs_runtime.Value
var once_show__1073032466 sync.Once
func Get_show__1073032466() gopurs_runtime.Value {
	once_show__1073032466.Do(func() {
		cache_show__1073032466 = gopurs_runtime.RecordGet(pkg_Data_Time.Get_showTime(), "show")
	})
	return cache_show__1073032466
}

var cache_fromDuration__1721614606 gopurs_runtime.Value
var once_fromDuration__1721614606 sync.Once
func Get_fromDuration__1721614606() gopurs_runtime.Value {
	once_fromDuration__1721614606.Do(func() {
		cache_fromDuration__1721614606 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromDuration__1721614606(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromDuration__1721614606
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

type Constructor_DateTime struct {
	Rc uint32
	V0 *pkg_Data_Date.Constructor_Date
	V1 *pkg_Data_Time.Constructor_Time
}


func Call_toRecord(v_0_loop *Constructor_DateTime) gopurs_runtime.Value {
var v_0 *Constructor_DateTime = v_0_loop
_ = v_0
return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V2)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V0)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V3)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__1196942535(), gopurs_runtime.Value{Type: 9, IntVal: int64((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V1), UnsafePtr: nil}).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*pkg_Data_Time.Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V2)).IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*pkg_Data_Date.Constructor_Date)(gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)}.UnsafePtr).V0)).IntVal)})
}

func Call_time(v_0_loop *Constructor_DateTime) *pkg_Data_Time.Constructor_Time {
var v_0 *Constructor_DateTime = v_0_loop
_ = v_0
return (*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
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
return gopurs_runtime.CoerceToStruct[Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, (*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_Time.Constructor_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}))})})
}

func Call_modifyDateF(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_DateTime = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply2(Get_flip__682087456(), Get_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0)}))
}

func Call_modifyDate(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_DateTime) *Constructor_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_DateTime = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_DateTime{1, gopurs_runtime.CoerceToStruct[pkg_Data_Date.Constructor_Date](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)})), (*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1})})
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
return (*Constructor_DateTime)(gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_adjust(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], d_1_loop gopurs_runtime.Value, dt_2_loop *Constructor_DateTime) *pkg_Data_Maybe.Constructor_Just[*Constructor_DateTime] {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var dt_2 *Constructor_DateTime = dt_2_loop
_ = dt_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_DateTime]](gopurs_runtime.Apply2(Get_bind__1702199617(), gopurs_runtime.Apply4(Get_adjustImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(dictDuration_0.V0, d_1), Call_toRecord(dt_2)), gopurs_runtime.Func(func(rec_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_DateTime]](gopurs_runtime.Apply2(Get_apply__3534390890(), gopurs_runtime.Apply2(Get_map__63598588(), Get_DateTime(), Call_join__880516349(gopurs_runtime.Apply2(Get_apply__420223018(), gopurs_runtime.Apply2(Get_apply__1183285642(), gopurs_runtime.Apply2(Get_map__2389909756(), pkg_Data_Date.Get_exactDate(), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.RecordGet(rec_3, "year"))), gopurs_runtime.Apply(Get_toEnum__2309750950(), gopurs_runtime.RecordGet(rec_3, "month"))), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.RecordGet(rec_3, "day"))))), gopurs_runtime.Apply2(Get_apply__3882563466(), gopurs_runtime.Apply2(Get_apply__3867059818(), gopurs_runtime.Apply2(Get_apply__3489442218(), gopurs_runtime.Apply2(Get_map__2165919164(), pkg_Data_Time.Get_Time(), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.RecordGet(rec_3, "hour"))), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.RecordGet(rec_3, "minute"))), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.RecordGet(rec_3, "second"))), gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.RecordGet(rec_3, "millisecond"))))))}
})))
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

func Call_bind__737692327(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_join__1635241211(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(dictBind_0.V1, m_1, pkg_Control_Bind.Get_identity())
}

func Call_join__880516349(m_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), m_0, pkg_Control_Bind.Get_identity())
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

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__682087456(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_map__366271444(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3592908820(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_fromDuration__1721614606(dict_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_AdjustImpl
}

func Get_calcDiff() gopurs_runtime.Value {
	return _Gopurs_CalcDiff
}
