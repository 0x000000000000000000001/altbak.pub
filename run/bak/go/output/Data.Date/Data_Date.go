package Data_Date

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Int "gopurs/output/Data.Int"
)

var Date gopurs_runtime.Value
var once_Date sync.Once
func Get_Date() gopurs_runtime.Value {
	once_Date.Do(func() {
		Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("Date", value0, value1, value2)
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
		year = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
})
	})
	return year
}

var weekday gopurs_runtime.Value
var once_weekday sync.Once
func Get_weekday() gopurs_runtime.Value {
	once_weekday.Do(func() {
		weekday = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "December").IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
n_1_0 := gopurs_runtime.UncurriedApp3(Get_calcWeekday(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], __t1, (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2])
_ = n_1_0
var __t2 gopurs_runtime.Value
{
if n_1_0.IntVal == 0 {
__t2 = gopurs_runtime.Constructor0("Sunday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 1 {
__t2 = gopurs_runtime.Constructor0("Monday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 2 {
__t2 = gopurs_runtime.Constructor0("Tuesday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 3 {
__t2 = gopurs_runtime.Constructor0("Wednesday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 4 {
__t2 = gopurs_runtime.Constructor0("Thursday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 5 {
__t2 = gopurs_runtime.Constructor0("Friday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 6 {
__t2 = gopurs_runtime.Constructor0("Saturday")
goto end_branch_2
} else {

}
}
{
if n_1_0.IntVal == 7 {
__t2 = gopurs_runtime.Constructor0("Sunday")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Str("January")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t0 = gopurs_runtime.Str("February")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Str("March")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Str("April")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Str("May")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Str("June")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Str("July")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Str("August")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Str("September")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Str("October")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Str("November")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "December").IntVal != 0 {
__t0 = gopurs_runtime.Str("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str("(Date (Year " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]).StrVal + ") " + __t0.StrVal + " (Day " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2]).StrVal + "))")
}))
	})
	return showDate
}

var month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		month = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
})
	})
	return month
}

var isLeapYear gopurs_runtime.Value
var once_isLeapYear sync.Once
func Get_isLeapYear() gopurs_runtime.Value {
	once_isLeapYear.Do(func() {
		isLeapYear = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), y_0, gopurs_runtime.Int(4)).IntVal == 0 && gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), y_0, gopurs_runtime.Int(400)).IntVal == 0 || gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), y_0, gopurs_runtime.Int(100)).IntVal != 0)
})
	})
	return isLeapYear
}

var lastDayOfMonth gopurs_runtime.Value
var once_lastDayOfMonth sync.Once
func Get_lastDayOfMonth() gopurs_runtime.Value {
	once_lastDayOfMonth.Do(func() {
		lastDayOfMonth = gopurs_runtime.Func2(func(y_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_1.StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "February").IntVal != 0 {
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
if gopurs_runtime.Bool(m_1.StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Int(31)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Int(30)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "December").IntVal != 0 {
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "January")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "February")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "March")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "April")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "May")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "June")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "July")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "August")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "September")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "October")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "November")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].StrVal == "December").IntVal != 0 && gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].StrVal == "December").IntVal != 0)
}
end_branch_0:
return gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].IntVal == (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].IntVal && __t0.IntVal != 0 && (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[2].IntVal == (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[2].IntVal)
}))
	})
	return eqDate
}

var ordDate gopurs_runtime.Value
var once_ordDate sync.Once
func Get_ordDate() gopurs_runtime.Value {
	once_ordDate.Do(func() {
		ordDate = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0])
_ = v_2_0
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2_0.StrVal == "LT").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("LT")
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_2_0.StrVal == "GT").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("GT")
goto end_branch_3
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_ordMonth(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1])
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_1.StrVal == "LT").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("LT")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_1.StrVal == "GT").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("GT")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[2])
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
sm_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "succ"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
_ = sm_1_0
__local_var_2_1 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2].IntVal + 1
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
if __local_var_2_1 >= 1 && __local_var_2_1 <= 31 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(__local_var_2_1))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
v1_3_2 := __t3
_ = v1_3_2
var __t10 gopurs_runtime.Value
{
__local_var_4_11 := gopurs_runtime.Apply2(Get_lastDayOfMonth(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
_ = __local_var_4_11
var __t12 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Nothing").IntVal != 0 {
__t12 = gopurs_runtime.Bool(false)
goto end_branch_12
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
__t12 = gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_3_2.UnsafePtr)[0].IntVal > __local_var_4_11.IntVal)
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
if gopurs_runtime.Bool(sm_1_0.StrVal == "Nothing").IntVal != 0 {
__t15 = gopurs_runtime.Bool(true)
goto end_branch_15
} else {

}
}
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Just").IntVal != 0 {
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
__local_var_4_16 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].IntVal + 1
_ = __local_var_4_16
var __t17 gopurs_runtime.Value
{
if __local_var_4_16 >= -271820 && __local_var_4_16 <= 275759 {
var __t18 gopurs_runtime.Value
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Nothing").IntVal != 0 {
__t18 = gopurs_runtime.Constructor0("January")
goto end_branch_18
} else {

}
}
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Just").IntVal != 0 {
__t18 = (*[1024]gopurs_runtime.Value)(sm_1_0.UnsafePtr)[0]
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t17 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor3("Date", gopurs_runtime.Int(__local_var_4_16), __t18, gopurs_runtime.Int(1)))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Constructor0("Nothing")
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
if gopurs_runtime.Bool(sm_1_0.StrVal == "Nothing").IntVal != 0 {
__t13 = gopurs_runtime.Constructor0("January")
goto end_branch_13
} else {

}
}
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Just").IntVal != 0 {
__t13 = (*[1024]gopurs_runtime.Value)(sm_1_0.UnsafePtr)[0]
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t14 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor3("Date", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], __t13, gopurs_runtime.Int(1)))
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
if gopurs_runtime.Bool(v1_3_2.StrVal == "Nothing").IntVal != 0 {
__t19 = gopurs_runtime.Bool(true)
goto end_branch_19
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
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
if gopurs_runtime.Bool(sm_1_0.StrVal == "Nothing").IntVal != 0 {
__t20 = gopurs_runtime.Bool(true)
goto end_branch_20
} else {

}
}
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Just").IntVal != 0 {
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
__local_var_4_21 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].IntVal + 1
_ = __local_var_4_21
var __t22 gopurs_runtime.Value
{
if __local_var_4_21 >= -271820 && __local_var_4_21 <= 275759 {
var __t24 gopurs_runtime.Value
{
var __t25 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Nothing").IntVal != 0 {
__t25 = gopurs_runtime.Bool(true)
goto end_branch_25
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
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
if gopurs_runtime.Bool(sm_1_0.StrVal == "Nothing").IntVal != 0 {
__t26 = gopurs_runtime.Constructor0("January")
goto end_branch_26
} else {

}
}
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Just").IntVal != 0 {
__t26 = (*[1024]gopurs_runtime.Value)(sm_1_0.UnsafePtr)[0]
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
__t24 = (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
}
end_branch_24:
__local_var_5_23 := gopurs_runtime.Apply2(Get_Date(), gopurs_runtime.Int(__local_var_4_21), __t24)
_ = __local_var_5_23
var __t27 gopurs_runtime.Value
{
var __t28 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Nothing").IntVal != 0 {
__t28 = gopurs_runtime.Bool(true)
goto end_branch_28
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
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
__t27 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_5_23, gopurs_runtime.Int(1)))
goto end_branch_27
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
__t27 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_5_23, (*[1024]gopurs_runtime.Value)(v1_3_2.UnsafePtr)[0]))
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_27:
__t22 = __t27
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.Constructor0("Nothing")
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
if gopurs_runtime.Bool(v1_3_2.StrVal == "Nothing").IntVal != 0 {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
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
if gopurs_runtime.Bool(sm_1_0.StrVal == "Nothing").IntVal != 0 {
__t7 = gopurs_runtime.Constructor0("January")
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(sm_1_0.StrVal == "Just").IntVal != 0 {
__t7 = (*[1024]gopurs_runtime.Value)(sm_1_0.UnsafePtr)[0]
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
__t5 = (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
}
end_branch_5:
__local_var_4_4 := gopurs_runtime.Apply2(Get_Date(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], __t5)
_ = __local_var_4_4
var __t8 gopurs_runtime.Value
{
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Nothing").IntVal != 0 {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
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
__t8 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Int(1)))
goto end_branch_8
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_2.StrVal == "Just").IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_4_4, (*[1024]gopurs_runtime.Value)(v1_3_2.UnsafePtr)[0]))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_8:
__t10 = __t8
}
end_branch_10:
return __t10
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
pm_1_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
_ = pm_1_29
__local_var_2_30 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2].IntVal - 1
_ = __local_var_2_30
var __t38 gopurs_runtime.Value
{
if __local_var_2_30 >= 1 && __local_var_2_30 <= 31 {
__t38 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor3("Date", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], gopurs_runtime.Int(__local_var_2_30)))
goto end_branch_38
} else {

}
}
{
var __t32 gopurs_runtime.Value
{
if gopurs_runtime.Bool(pm_1_29.StrVal == "Nothing").IntVal != 0 {
__t32 = gopurs_runtime.Constructor0("December")
goto end_branch_32
} else {

}
}
{
if gopurs_runtime.Bool(pm_1_29.StrVal == "Just").IntVal != 0 {
__t32 = (*[1024]gopurs_runtime.Value)(pm_1_29.UnsafePtr)[0]
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
l_4_33 := gopurs_runtime.Apply2(Get_lastDayOfMonth(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], m_prime_3_31)
_ = l_4_33
var __t34 gopurs_runtime.Value
{
var __t35 gopurs_runtime.Value
{
if gopurs_runtime.Bool(pm_1_29.StrVal == "Nothing").IntVal != 0 {
__t35 = gopurs_runtime.Bool(true)
goto end_branch_35
} else {

}
}
{
if gopurs_runtime.Bool(pm_1_29.StrVal == "Just").IntVal != 0 {
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
__local_var_5_36 := (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].IntVal - 1
_ = __local_var_5_36
var __t37 gopurs_runtime.Value
{
if __local_var_5_36 >= -271820 && __local_var_5_36 <= 275759 {
__t37 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor3("Date", gopurs_runtime.Int(__local_var_5_36), m_prime_3_31, l_4_33))
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_37:
__t34 = __t37
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor3("Date", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], m_prime_3_31, l_4_33))
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
		diff = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1].StrVal == "December").IntVal != 0 {
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
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1].StrVal == "December").IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.UncurriedApp6(Get_calcDiff(), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], __t0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0], __t1, (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[2]))
})
	})
	return diff
}

var day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		day = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2]
})
	})
	return day
}

var canonicalDate gopurs_runtime.Value
var once_canonicalDate sync.Once
func Get_canonicalDate() gopurs_runtime.Value {
	once_canonicalDate.Do(func() {
		canonicalDate = gopurs_runtime.Func3(func(y_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value, d_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_1.StrVal == "January").IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "February").IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "March").IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "April").IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "May").IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "June").IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "July").IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "August").IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "September").IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "October").IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "November").IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "December").IntVal != 0 {
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
__t0 = gopurs_runtime.Constructor0("January")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 2 {
__t0 = gopurs_runtime.Constructor0("February")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 3 {
__t0 = gopurs_runtime.Constructor0("March")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 4 {
__t0 = gopurs_runtime.Constructor0("April")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 5 {
__t0 = gopurs_runtime.Constructor0("May")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 6 {
__t0 = gopurs_runtime.Constructor0("June")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 7 {
__t0 = gopurs_runtime.Constructor0("July")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 8 {
__t0 = gopurs_runtime.Constructor0("August")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 9 {
__t0 = gopurs_runtime.Constructor0("September")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 10 {
__t0 = gopurs_runtime.Constructor0("October")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 11 {
__t0 = gopurs_runtime.Constructor0("November")
goto end_branch_0
} else {

}
}
{
if m_prime_4.IntVal == 12 {
__t0 = gopurs_runtime.Constructor0("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Constructor3("Date", y_prime_3, __t0, d_prime_5)
}), y_0, __t1, d_2)
})
	})
	return canonicalDate
}

var exactDate gopurs_runtime.Value
var once_exactDate sync.Once
func Get_exactDate() gopurs_runtime.Value {
	once_exactDate.Do(func() {
		exactDate = gopurs_runtime.Func3(func(y_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value, d_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqDate(), "eq"), gopurs_runtime.Apply3(Get_canonicalDate(), y_0, m_1, d_2), gopurs_runtime.Constructor3("Date", y_0, m_1, d_2)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor3("Date", y_0, m_1, d_2))
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
	return exactDate
}

var boundedDate gopurs_runtime.Value
var once_boundedDate sync.Once
func Get_boundedDate() gopurs_runtime.Value {
	once_boundedDate.Do(func() {
		boundedDate = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Constructor3("Date", gopurs_runtime.Int(-271820), gopurs_runtime.Constructor0("January"), gopurs_runtime.Int(1)), gopurs_runtime.Constructor3("Date", gopurs_runtime.Int(275759), gopurs_runtime.Constructor0("December"), gopurs_runtime.Int(31)), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
}))
	})
	return boundedDate
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, date_1 gopurs_runtime.Value) gopurs_runtime.Value {
var adj_2_0 gopurs_runtime.Value
_ = adj_2_0
adj_2_0 = gopurs_runtime.Func2(func(v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if v1_3.IntVal == 0 {
__t15 = gopurs_runtime.Constructor1("Just", v2_4)
goto end_branch_15
} else {

}
}
{
j_5_1 := v1_3.IntVal + (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[2].IntVal
_ = j_5_1
low_6_2 := j_5_1 < 1
_ = low_6_2
var __t4 gopurs_runtime.Value
{
if low_6_2 {
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_enumMonth(), "pred"), (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[1])
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_7_5.StrVal == "Nothing").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("December")
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_7_5.StrVal == "Just").IntVal != 0 {
__t6 = (*[1024]gopurs_runtime.Value)(__local_var_7_5.UnsafePtr)[0]
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
__t4 = (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[1]
}
end_branch_4:
l_7_3 := gopurs_runtime.Apply2(Get_lastDayOfMonth(), (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0], __t4)
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
__t13 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_enumDate(), "pred"), gopurs_runtime.Constructor3("Date", (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[1], gopurs_runtime.Int(1)))
goto end_branch_13
} else {

}
}
{
if hi_8_7 {
__t13 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_enumDate(), "succ"), gopurs_runtime.Constructor3("Date", (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[1], l_7_3))
goto end_branch_13
} else {

}
}
{
__local_var_10_11 := gopurs_runtime.Apply2(Get_Date(), (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[1])
_ = __local_var_10_11
var __t12 gopurs_runtime.Value
{
if j_5_1 >= 1 && j_5_1 <= 31 {
__t12 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_10_11, gopurs_runtime.Int(j_5_1)))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_12:
__t13 = __t12
}
end_branch_13:
__local_var_10_10 := __t13
_ = __local_var_10_10
var __t14 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_10_10.StrVal == "Just").IntVal != 0 {
__t14 = gopurs_runtime.Apply(__local_var_9_8, (*[1024]gopurs_runtime.Value)(__local_var_10_10.UnsafePtr)[0])
goto end_branch_14
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_10_10.StrVal == "Nothing").IntVal != 0 {
__t14 = gopurs_runtime.Constructor0("Nothing")
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
if gopurs_runtime.Bool(__local_var_3_16.StrVal == "Just").IntVal != 0 {
__t17 = gopurs_runtime.Apply2(adj_2_0, (*[1024]gopurs_runtime.Value)(__local_var_3_16.UnsafePtr)[0], date_1)
goto end_branch_17
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_16.StrVal == "Nothing").IntVal != 0 {
__t17 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
})
	})
	return adjust
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
