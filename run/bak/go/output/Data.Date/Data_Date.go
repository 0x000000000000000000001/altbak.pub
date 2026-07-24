package Data_Date

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Int "gopurs/output/Data.Int"
	unsafe "unsafe"
)

var Date gopurs_runtime.Value
var once_Date sync.Once
func Get_Date() gopurs_runtime.Value {
	once_Date.Do(func() {
		Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{value0, value1, value2})}
})
})
})
	})
	return Date
}

var year gopurs_runtime.Value
var once_year sync.Once
func Get_year() gopurs_runtime.Value {
	once_year.Do(func() {
		year = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Date_Date)(v_0.UnsafePtr).V0
}()
})
	})
	return year
}

var weekday gopurs_runtime.Value
var once_weekday sync.Once
func Get_weekday() gopurs_runtime.Value {
	once_weekday.Do(func() {
		weekday = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 3320970370) {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 904613236) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 2235536813) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 116409214) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 1527394637) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 2202783052) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 1676632594) {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 4203147001) {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 783850007) {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 2522709242) {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 25181906) {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 3004478759) {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
n_1_0 := gopurs_runtime.UncurriedApp3(Get_calcWeekday(), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0, __t1, (*Data_Data_Date_Date)(v_0.UnsafePtr).V2)
_ = n_1_0
var __t2 gopurs_runtime.Value
{
if n_1_0.IntVal == 0 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1355039372, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Sunday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 1 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 4055383368, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Monday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 2 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3750810643, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Tuesday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 3 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 647375594, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Wednesday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 4 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3164445356, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Thursday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 5 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1018578585, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Friday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 6 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3182748101, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Saturday{})}
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 7 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1355039372, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_Sunday{})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}()
})
	})
	return weekday
}

var showDate gopurs_runtime.Value
var once_showDate sync.Once
func Get_showDate() gopurs_runtime.Value {
	once_showDate.Do(func() {
		showDate = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 3320970370) {
__t0 = gopurs_runtime.Str("January")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 904613236) {
__t0 = gopurs_runtime.Str("February")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 2235536813) {
__t0 = gopurs_runtime.Str("March")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 116409214) {
__t0 = gopurs_runtime.Str("April")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 1527394637) {
__t0 = gopurs_runtime.Str("May")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 2202783052) {
__t0 = gopurs_runtime.Str("June")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 1676632594) {
__t0 = gopurs_runtime.Str("July")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 4203147001) {
__t0 = gopurs_runtime.Str("August")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 783850007) {
__t0 = gopurs_runtime.Str("September")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 2522709242) {
__t0 = gopurs_runtime.Str("October")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 25181906) {
__t0 = gopurs_runtime.Str("November")
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_0.UnsafePtr).V1.IntVal == 3004478759) {
__t0 = gopurs_runtime.Str("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str("(Date (Year " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0).StrVal() + ") " + __t0.StrVal() + " (Day " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*Data_Data_Date_Date)(v_0.UnsafePtr).V2).StrVal() + "))")
}))
	})
	return showDate
}

var month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		month = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Date_Date)(v_0.UnsafePtr).V1
}()
})
	})
	return month
}

var isLeapYear gopurs_runtime.Value
var once_isLeapYear sync.Once
func Get_isLeapYear() gopurs_runtime.Value {
	once_isLeapYear.Do(func() {
		isLeapYear = gopurs_runtime.Func(func(y_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
return gopurs_runtime.Bool(gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), y_0, gopurs_runtime.Int(4)).IntVal == 0 && gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), y_0, gopurs_runtime.Int(400)).IntVal == 0 || gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), y_0, gopurs_runtime.Int(100)).IntVal != 0)
}()
})
	})
	return isLeapYear
}

var lastDayOfMonth gopurs_runtime.Value
var once_lastDayOfMonth sync.Once
func Get_lastDayOfMonth() gopurs_runtime.Value {
	once_lastDayOfMonth.Do(func() {
		lastDayOfMonth = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lastDayOfMonth(y_0_box, m_1_box)
})
	})
	return lastDayOfMonth
}

var eqDate gopurs_runtime.Value
var once_eqDate sync.Once
func Get_eqDate() gopurs_runtime.Value {
	once_eqDate.Do(func() {
		eqDate = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 3320970370) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 3320970370))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 904613236) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 904613236))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 2235536813) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 2235536813))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 116409214) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 116409214))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 1527394637) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 1527394637))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 2202783052) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 2202783052))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 1676632594) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 1676632594))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 4203147001) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 4203147001))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 783850007) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 783850007))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 2522709242) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 2522709242))
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 25181906) {
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 25181906))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((*Data_Data_Date_Date)(x_0.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V1.IntVal == 3004478759) && ((*Data_Data_Date_Date)(y_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(y_1.UnsafePtr).V1.IntVal == 3004478759))
}
end_branch_0:
return gopurs_runtime.Bool((*Data_Data_Date_Date)(x_0.UnsafePtr).V0.IntVal == (*Data_Data_Date_Date)(y_1.UnsafePtr).V0.IntVal && __t0.IntVal != 0 && (*Data_Data_Date_Date)(x_0.UnsafePtr).V2.IntVal == (*Data_Data_Date_Date)(y_1.UnsafePtr).V2.IntVal)
}))
	})
	return eqDate
}

var ordDate gopurs_runtime.Value
var once_ordDate sync.Once
func Get_ordDate() gopurs_runtime.Value {
	once_ordDate.Do(func() {
		ordDate = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V0, (*Data_Data_Date_Date)(y_1.UnsafePtr).V0)
_ = v_2_0
var __t3 gopurs_runtime.Value
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 3866105248) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_3
} else {

}
}
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 2098047435) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_3
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordMonth(), "compare"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V1, (*Data_Data_Date_Date)(y_1.UnsafePtr).V1)
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 3866105248) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_2
} else {

}
}
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 2098047435) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*Data_Data_Date_Date)(x_0.UnsafePtr).V2, (*Data_Data_Date_Date)(y_1.UnsafePtr).V2)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDate()
}))
	})
	return ordDate
}

var enumDate gopurs_runtime.Value
var once_enumDate sync.Once
func Get_enumDate() gopurs_runtime.Value {
	once_enumDate.Do(func() {
		enumDate = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
sm_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "succ"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V1)
_ = sm_1_0
__local_var_2_1 := (*Data_Data_Date_Date)(v_0.UnsafePtr).V2.IntVal + 1
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
if __local_var_2_1 >= 1 && __local_var_2_1 <= 31 {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_2_1)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
v1_3_2 := __t3
_ = v1_3_2
var __t10 gopurs_runtime.Value
{
__local_var_4_11 := Call_lastDayOfMonth((*Data_Data_Date_Date)(v_0.UnsafePtr).V0, (*Data_Data_Date_Date)(v_0.UnsafePtr).V1)
_ = __local_var_4_11
var __t12 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 42808261) {
__t12 = gopurs_runtime.Bool(false)
goto end_branch_12
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t12 = gopurs_runtime.Bool((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_3_2.UnsafePtr).V0.IntVal > __local_var_4_11.IntVal)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
if __t12.IntVal != 0 {
var __t14 gopurs_runtime.Value
{
var __t15 gopurs_runtime.Value
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 42808261) {
__t15 = gopurs_runtime.Bool(true)
goto end_branch_15
} else {

}
}
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 1354639136) {
__t15 = gopurs_runtime.Bool(false)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
if __t15.IntVal != 0 {
__local_var_4_16 := (*Data_Data_Date_Date)(v_0.UnsafePtr).V0.IntVal + 1
_ = __local_var_4_16
var __t17 gopurs_runtime.Value
{
if __local_var_4_16 >= -271820 && __local_var_4_16 <= 275759 {
var __t18 gopurs_runtime.Value
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 42808261) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 3320970370, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}
goto end_branch_18
} else {

}
}
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 1354639136) {
__t18 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(sm_1_0.UnsafePtr).V0
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{gopurs_runtime.Int(__local_var_4_16), __t18, gopurs_runtime.Int(1)})}})}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_17:
__t14 = __t17
goto end_branch_14
} else {

}
}
{
var __t13 gopurs_runtime.Value
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 42808261) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3320970370, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}
goto end_branch_13
} else {

}
}
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 1354639136) {
__t13 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(sm_1_0.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{(*Data_Data_Date_Date)(v_0.UnsafePtr).V0, __t13, gopurs_runtime.Int(1)})}})}
}
end_branch_14:
__t10 = __t14
goto end_branch_10
} else {

}
}
{
var __t19 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 42808261) {
__t19 = gopurs_runtime.Bool(true)
goto end_branch_19
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
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
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 42808261) {
__t20 = gopurs_runtime.Bool(true)
goto end_branch_20
} else {

}
}
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 1354639136) {
__t20 = gopurs_runtime.Bool(false)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
if __t19.IntVal != 0 && __t20.IntVal != 0 {
__local_var_4_21 := (*Data_Data_Date_Date)(v_0.UnsafePtr).V0.IntVal + 1
_ = __local_var_4_21
var __t22 gopurs_runtime.Value
{
if __local_var_4_21 >= -271820 && __local_var_4_21 <= 275759 {
var __t24 gopurs_runtime.Value
{
var __t25 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 42808261) {
__t25 = gopurs_runtime.Bool(true)
goto end_branch_25
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t25 = gopurs_runtime.Bool(false)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
if __t25.IntVal != 0 {
var __t26 gopurs_runtime.Value
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 42808261) {
__t26 = gopurs_runtime.Value{Type: 9, IntVal: 3320970370, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}
goto end_branch_26
} else {

}
}
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 1354639136) {
__t26 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(sm_1_0.UnsafePtr).V0
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
__t24 = __t26
goto end_branch_24
} else {

}
}
{
__t24 = (*Data_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_24:
__local_var_5_23 := gopurs_runtime.Apply2(Get_Date(), gopurs_runtime.Int(__local_var_4_21), __t24)
_ = __local_var_5_23
var __t27 gopurs_runtime.Value
{
var __t28 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 42808261) {
__t28 = gopurs_runtime.Bool(true)
goto end_branch_28
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t28 = gopurs_runtime.Bool(false)
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
if __t28.IntVal != 0 {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_5_23, gopurs_runtime.Int(1))})}
goto end_branch_27
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_5_23, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_3_2.UnsafePtr).V0)})}
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_27:
__t22 = __t27
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_22:
__t10 = __t22
goto end_branch_10
} else {

}
}
{
var __t5 gopurs_runtime.Value
{
var __t6 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 42808261) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
if __t6.IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 42808261) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3320970370, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}
goto end_branch_7
} else {

}
}
{
if (sm_1_0.Type == 9 && sm_1_0.IntVal == 1354639136) {
__t7 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(sm_1_0.UnsafePtr).V0
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
__t5 = (*Data_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_5:
__local_var_4_4 := gopurs_runtime.Apply2(Get_Date(), (*Data_Data_Date_Date)(v_0.UnsafePtr).V0, __t5)
_ = __local_var_4_4
var __t8 gopurs_runtime.Value
{
var __t9 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 42808261) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t9 = gopurs_runtime.Bool(false)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
if __t9.IntVal != 0 {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Int(1))})}
goto end_branch_8
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1354639136) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_4_4, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_3_2.UnsafePtr).V0)})}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_8:
__t10 = __t8
}
end_branch_10:
return __t10
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
pm_1_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), (*Data_Data_Date_Date)(v_0.UnsafePtr).V1)
_ = pm_1_29
__local_var_2_30 := (*Data_Data_Date_Date)(v_0.UnsafePtr).V2.IntVal - 1
_ = __local_var_2_30
var __t38 gopurs_runtime.Value
{
if __local_var_2_30 >= 1 && __local_var_2_30 <= 31 {
__t38 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{(*Data_Data_Date_Date)(v_0.UnsafePtr).V0, (*Data_Data_Date_Date)(v_0.UnsafePtr).V1, gopurs_runtime.Int(__local_var_2_30)})}})}
goto end_branch_38
} else {

}
}
{
var __t32 gopurs_runtime.Value
{
if (pm_1_29.Type == 9 && pm_1_29.IntVal == 42808261) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 3004478759, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_December{})}
goto end_branch_32
} else {

}
}
{
if (pm_1_29.Type == 9 && pm_1_29.IntVal == 1354639136) {
__t32 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(pm_1_29.UnsafePtr).V0
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
m_prime_3_31 := __t32
_ = m_prime_3_31
l_4_33 := Call_lastDayOfMonth((*Data_Data_Date_Date)(v_0.UnsafePtr).V0, m_prime_3_31)
_ = l_4_33
var __t34 gopurs_runtime.Value
{
var __t35 gopurs_runtime.Value
{
if (pm_1_29.Type == 9 && pm_1_29.IntVal == 42808261) {
__t35 = gopurs_runtime.Bool(true)
goto end_branch_35
} else {

}
}
{
if (pm_1_29.Type == 9 && pm_1_29.IntVal == 1354639136) {
__t35 = gopurs_runtime.Bool(false)
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
if __t35.IntVal != 0 {
__local_var_5_36 := (*Data_Data_Date_Date)(v_0.UnsafePtr).V0.IntVal - 1
_ = __local_var_5_36
var __t37 gopurs_runtime.Value
{
if __local_var_5_36 >= -271820 && __local_var_5_36 <= 275759 {
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{gopurs_runtime.Int(__local_var_5_36), m_prime_3_31, l_4_33})}})}
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_37:
__t34 = __t37
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{(*Data_Data_Date_Date)(v_0.UnsafePtr).V0, m_prime_3_31, l_4_33})}})}
}
end_branch_34:
__t38 = __t34
}
end_branch_38:
return __t38
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
}))
	})
	return enumDate
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(dictDuration_0_box, v_1_box, v1_2_box)
})
	})
	return diff
}

var day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		day = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Date_Date)(v_0.UnsafePtr).V2
}()
})
	})
	return day
}

var canonicalDate gopurs_runtime.Value
var once_canonicalDate sync.Once
func Get_canonicalDate() gopurs_runtime.Value {
	once_canonicalDate.Do(func() {
		canonicalDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_canonicalDate(y_0_box, m_1_box, d_2_box)
})
	})
	return canonicalDate
}

var exactDate gopurs_runtime.Value
var once_exactDate sync.Once
func Get_exactDate() gopurs_runtime.Value {
	once_exactDate.Do(func() {
		exactDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_exactDate(y_0_box, m_1_box, d_2_box)
})
	})
	return exactDate
}

var boundedDate gopurs_runtime.Value
var once_boundedDate sync.Once
func Get_boundedDate() gopurs_runtime.Value {
	once_boundedDate.Do(func() {
		boundedDate = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{gopurs_runtime.Int(-271820), gopurs_runtime.Value{Type: 9, IntVal: 3320970370, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}, gopurs_runtime.Int(1)})}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{gopurs_runtime.Int(275759), gopurs_runtime.Value{Type: 9, IntVal: 3004478759, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_December{})}, gopurs_runtime.Int(31)})}, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
}))
	})
	return boundedDate
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, date_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_adjust(v_0_box, date_1_box)
})
	})
	return adjust
}

type Data_Data_Date_Date struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_Date_Date(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 745776346
}

func Call_lastDayOfMonth(y_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3320970370) {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 904613236) {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Apply(Get_isLeapYear(), y_0).IntVal != 0 {
__t1 = gopurs_runtime.Int(29)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int(28)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2235536813) {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 116409214) {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 1527394637) {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2202783052) {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 1676632594) {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 4203147001) {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 783850007) {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2522709242) {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 25181906) {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 3004478759) {
__t0 = gopurs_runtime.Int(31)
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

func Call_diff(dictDuration_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 3320970370) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 904613236) {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 2235536813) {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 116409214) {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 1527394637) {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 2202783052) {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 1676632594) {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 4203147001) {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 783850007) {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 2522709242) {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 25181906) {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if ((*Data_Data_Date_Date)(v_1.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v_1.UnsafePtr).V1.IntVal == 3004478759) {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
var __t1 gopurs_runtime.Value
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 3320970370) {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 904613236) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 2235536813) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 116409214) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 1527394637) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 2202783052) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 1676632594) {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 4203147001) {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 783850007) {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 2522709242) {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 25181906) {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if ((*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.Type == 9 && (*Data_Data_Date_Date)(v1_2.UnsafePtr).V1.IntVal == 3004478759) {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.UncurriedApp6(Get_calcDiff(), (*Data_Data_Date_Date)(v_1.UnsafePtr).V0, __t0, (*Data_Data_Date_Date)(v_1.UnsafePtr).V2, (*Data_Data_Date_Date)(v1_2.UnsafePtr).V0, __t1, (*Data_Data_Date_Date)(v1_2.UnsafePtr).V2))
}

func Call_canonicalDate(y_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, d_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var __t1 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3320970370) {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 904613236) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2235536813) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 116409214) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 1527394637) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2202783052) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 1676632594) {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 4203147001) {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 783850007) {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2522709242) {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 25181906) {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 3004478759) {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.UncurriedApp4(Get_canonicalDateImpl(), gopurs_runtime.Func3(func(y_prime_3 gopurs_runtime.Value, m_prime_4 gopurs_runtime.Value, d_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if m_prime_4.IntVal == 1 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3320970370, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_January{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 2 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 904613236, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_February{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 3 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2235536813, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_March{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 4 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 116409214, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_April{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 5 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1527394637, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_May{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 6 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2202783052, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_June{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 7 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1676632594, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_July{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 8 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 4203147001, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_August{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 9 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 783850007, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_September{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 10 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2522709242, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_October{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 11 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 25181906, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_November{})}
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 12 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3004478759, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_December{})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{y_prime_3, __t0, d_prime_5})}
}), y_0, __t1, d_2)
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
if gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqDate(), "eq"), Call_canonicalDate(y_0, m_1, d_2), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{y_0, m_1, d_2})}).IntVal != 0 {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{y_0, m_1, d_2})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
var __t15 gopurs_runtime.Value
{
if v1_3.IntVal == 0 {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v2_4})}
goto end_branch_15
} else {

}
}
{
j_5_1 := v1_3.IntVal + (*Data_Data_Date_Date)(v2_4.UnsafePtr).V2.IntVal
_ = j_5_1
low_6_2 := j_5_1 < 1
_ = low_6_2
var __t4 gopurs_runtime.Value
{
if low_6_2 {
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1)
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if (__local_var_7_5.Type == 9 && __local_var_7_5.IntVal == 42808261) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3004478759, UnsafePtr: unsafe.Pointer(&pkg_Data_Date_Component.Data_Data_Date_Component_December{})}
goto end_branch_6
} else {

}
}
{
if (__local_var_7_5.Type == 9 && __local_var_7_5.IntVal == 1354639136) {
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
hi_8_7 := j_5_1 > l_7_3.IntVal
_ = hi_8_7
var __t9 gopurs_runtime.Value
{
if low_6_2 {
__t9 = gopurs_runtime.Int(j_5_1)
goto end_branch_9
} else {

}
}
{
if hi_8_7 {
__t9 = gopurs_runtime.Int(j_5_1 - l_7_3.IntVal - 1)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Int(0)
}
end_branch_9:
__local_var_9_8 := gopurs_runtime.Apply(adj_2_0, __t9)
_ = __local_var_9_8
var __t13 gopurs_runtime.Value
{
if low_6_2 {
__t13 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_enumDate(), "pred"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{(*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1, gopurs_runtime.Int(1)})})
goto end_branch_13
} else {

}
}
{
if hi_8_7 {
__t13 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_enumDate(), "succ"), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Date{(*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1, l_7_3})})
goto end_branch_13
} else {

}
}
{
__local_var_10_11 := gopurs_runtime.Apply2(Get_Date(), (*Data_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Data_Data_Date_Date)(v2_4.UnsafePtr).V1)
_ = __local_var_10_11
var __t12 gopurs_runtime.Value
{
if j_5_1 >= 1 && j_5_1 <= 31 {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_10_11, gopurs_runtime.Int(j_5_1))})}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_12:
__t13 = __t12
}
end_branch_13:
__local_var_10_10 := __t13
_ = __local_var_10_10
var __t14 gopurs_runtime.Value
{
if (__local_var_10_10.Type == 9 && __local_var_10_10.IntVal == 1354639136) {
__t14 = gopurs_runtime.Apply(__local_var_9_8, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_10_10.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
if (__local_var_10_10.Type == 9 && __local_var_10_10.IntVal == 42808261) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
__t15 = __t14
}
end_branch_15:
return __t15
})
__local_var_3_16 := gopurs_runtime.Apply(pkg_Data_Int.Get_fromNumber(), v_0)
_ = __local_var_3_16
var __t17 gopurs_runtime.Value
{
if (__local_var_3_16.Type == 9 && __local_var_3_16.IntVal == 1354639136) {
__t17 = gopurs_runtime.Apply2(adj_2_0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_3_16.UnsafePtr).V0, date_1)
goto end_branch_17
} else {

}
}
{
if (__local_var_3_16.Type == 9 && __local_var_3_16.IntVal == 42808261) {
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
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
