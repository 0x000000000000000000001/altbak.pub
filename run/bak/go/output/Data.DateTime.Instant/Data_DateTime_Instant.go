package Data_DateTime_Instant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	unsafe "unsafe"
)

var cache_unInstant gopurs_runtime.Value
var once_unInstant sync.Once
func Get_unInstant() gopurs_runtime.Value {
	once_unInstant.Do(func() {
		cache_unInstant = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}()
})
	})
	return cache_unInstant
}

var cache_toDateTime gopurs_runtime.Value
var once_toDateTime sync.Once
func Get_toDateTime() gopurs_runtime.Value {
	once_toDateTime.Do(func() {
		cache_toDateTime = gopurs_runtime.Apply(Get_toDateTimeImpl(), gopurs_runtime.Func5(func(y_0 gopurs_runtime.Value, mo_1 gopurs_runtime.Value, d_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, mi_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func2(func(s_5 gopurs_runtime.Value, ms_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mo_1.IntVal) == (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (2) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2455627378, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_February{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (3) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 4162469099, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_March{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (4) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1692989816, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_April{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (5) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 330658827, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_May{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (6) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 4067355978, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_June{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (7) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2276710548, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_July{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (8) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 243771071, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_August{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (9) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 215731793, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_September{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (10) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 8639228, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_October{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (11) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 49471444, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_November{})}
goto end_branch_0
} else {

}
}
{
if (mo_1.IntVal) == (12) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_December{})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&pkg_Data_DateTime.Data_Data_DateTime_DateTime{gopurs_runtime.Apply3(pkg_Data_Date.Get_canonicalDate(), y_0, __t0, d_2), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&pkg_Data_Time.Data_Data_Time_Time{h_3, mi_4, s_5, ms_6})}})}
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
return gopurs_runtime.Str((("(Instant (Milliseconds ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal())) + ("))"))
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
		cache_instant = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if ((v_0.FloatVal()) >= (-8639977881600000.0)) && ((v_0.FloatVal()) <= (8639977881599999.0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}()
})
	})
	return cache_instant
}

var cache_fromDateTime gopurs_runtime.Value
var once_fromDateTime sync.Once
func Get_fromDateTime() gopurs_runtime.Value {
	once_fromDateTime.Do(func() {
		cache_fromDateTime = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 1908470532) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 2455627378) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 4162469099) {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 1692989816) {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 330658827) {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 4067355978) {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 2276710548) {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 243771071) {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 215731793) {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 8639228) {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 49471444) {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V1.IntVal == 3889233761) {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UncurriedApp7(Get_fromDateTimeImpl(), (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V0, __t0, (*pkg_Data_Date.Data_Data_Date_Date)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_Time.Data_Data_Time_Time)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V1.UnsafePtr).V0, (*pkg_Data_Time.Data_Data_Time_Time)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V1.UnsafePtr).V1, (*pkg_Data_Time.Data_Data_Time_Time)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V1.UnsafePtr).V2, (*pkg_Data_Time.Data_Data_Time_Time)((*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(v_0.UnsafePtr).V1.UnsafePtr).V3)
}()
})
	})
	return cache_fromDateTime
}

var cache_fromDate gopurs_runtime.Value
var once_fromDate sync.Once
func Get_fromDate() gopurs_runtime.Value {
	once_fromDate.Do(func() {
		cache_fromDate = gopurs_runtime.Func(func(d_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var d_0 gopurs_runtime.Value = d_0_loop
_ = d_0
var __t0 gopurs_runtime.Value
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 1908470532) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 2455627378) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 4162469099) {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 1692989816) {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 330658827) {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 4067355978) {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 2276710548) {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 243771071) {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 215731793) {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 8639228) {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 49471444) {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if ((*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V1.IntVal == 3889233761) {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UncurriedApp7(Get_fromDateTimeImpl(), (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V0, __t0, (*pkg_Data_Date.Data_Data_Date_Date)(d_0.UnsafePtr).V2, gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0))
}()
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
return Call_diff(dictDuration_0_box, dt1_1_box, dt2_2_box)
})
	})
	return cache_diff
}

var cache_boundedInstant gopurs_runtime.Value
var once_boundedInstant sync.Once
func Get_boundedInstant() gopurs_runtime.Value {
	once_boundedInstant.Do(func() {
		cache_boundedInstant = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Float(-8639977881600000.0), gopurs_runtime.Float(8639977881599999.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
}))
	})
	return cache_boundedInstant
}

func Call_diff(dictDuration_0_loop gopurs_runtime.Value, dt1_1_loop gopurs_runtime.Value, dt2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 gopurs_runtime.Value = dt1_1_loop
_ = dt1_1
var dt2_2 gopurs_runtime.Value = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.Float((dt1_1.FloatVal()) + (gopurs_runtime.FloatNeg(dt2_2).FloatVal())))
}

func Get_fromDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_FromDateTimeImpl
}

func Get_toDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_ToDateTimeImpl
}
