package Data_DateTime_Instant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var unInstant gopurs_runtime.Value
var once_unInstant sync.Once
func Get_unInstant() gopurs_runtime.Value {
	once_unInstant.Do(func() {
		unInstant = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return unInstant
}

var toDateTime gopurs_runtime.Value
var once_toDateTime sync.Once
func Get_toDateTime() gopurs_runtime.Value {
	once_toDateTime.Do(func() {
		toDateTime = gopurs_runtime.Apply(Get_toDateTimeImpl(), gopurs_runtime.Func5(func(y_0 gopurs_runtime.Value, mo_1 gopurs_runtime.Value, d_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, mi_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func2(func(s_5 gopurs_runtime.Value, ms_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if mo_1.IntVal == 1 {
__t0 = gopurs_runtime.Constructor0("January")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 2 {
__t0 = gopurs_runtime.Constructor0("February")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 3 {
__t0 = gopurs_runtime.Constructor0("March")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 4 {
__t0 = gopurs_runtime.Constructor0("April")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 5 {
__t0 = gopurs_runtime.Constructor0("May")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 6 {
__t0 = gopurs_runtime.Constructor0("June")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 7 {
__t0 = gopurs_runtime.Constructor0("July")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 8 {
__t0 = gopurs_runtime.Constructor0("August")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 9 {
__t0 = gopurs_runtime.Constructor0("September")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 10 {
__t0 = gopurs_runtime.Constructor0("October")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 11 {
__t0 = gopurs_runtime.Constructor0("November")
goto end_branch_0
} else {

}
}
{
if mo_1.IntVal == 12 {
__t0 = gopurs_runtime.Constructor0("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Constructor2("DateTime", gopurs_runtime.Apply3(pkg_Data_Date.Get_canonicalDate(), y_0, __t0, d_2), gopurs_runtime.Constructor4("Time", h_3, mi_4, s_5, ms_6))
})
}))
	})
	return toDateTime
}

var showInstant gopurs_runtime.Value
var once_showInstant sync.Once
func Get_showInstant() gopurs_runtime.Value {
	once_showInstant.Do(func() {
		showInstant = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Instant (Milliseconds " + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), v_0).StrVal + "))")
}))
	})
	return showInstant
}

var ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		ordDateTime = pkg_Data_Ord.Get_ordNumber()
	})
	return ordDateTime
}

var instant gopurs_runtime.Value
var once_instant sync.Once
func Get_instant() gopurs_runtime.Value {
	once_instant.Do(func() {
		instant = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if v_0.FloatVal() >= -8639977881600000.0 && v_0.FloatVal() <= 8639977881599999.0 {
__t0 = gopurs_runtime.Constructor1("Just", v_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
})
	})
	return instant
}

var fromDateTime gopurs_runtime.Value
var once_fromDateTime sync.Once
func Get_fromDateTime() gopurs_runtime.Value {
	once_fromDateTime.Do(func() {
		fromDateTime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "December").IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UncurriedApp(Get_fromDateTimeImpl(), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[0], __t0, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[0], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[1], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[3])
})
	})
	return fromDateTime
}

var fromDate gopurs_runtime.Value
var once_fromDate sync.Once
func Get_fromDate() gopurs_runtime.Value {
	once_fromDate.Do(func() {
		fromDate = gopurs_runtime.Func(func(d_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[1].StrVal == "December").IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UncurriedApp(Get_fromDateTimeImpl(), (*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[0], __t0, (*[1024]gopurs_runtime.Value)(d_0.UnsafePtr)[2], gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0))
})
	})
	return fromDate
}

var eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		eqDateTime = pkg_Data_Eq.Get_eqNumber()
	})
	return eqDateTime
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, dt1_1 gopurs_runtime.Value, dt2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.Float(dt1_1.FloatVal() + gopurs_runtime.FloatNeg(dt2_2).FloatVal()))
})
	})
	return diff
}

var boundedInstant gopurs_runtime.Value
var once_boundedInstant sync.Once
func Get_boundedInstant() gopurs_runtime.Value {
	once_boundedInstant.Do(func() {
		boundedInstant = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Float(-8639977881600000.0), gopurs_runtime.Float(8639977881599999.0), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
}))
	})
	return boundedInstant
}



func Get_fromDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_FromDateTimeImpl
}

func Get_toDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_ToDateTimeImpl
}
