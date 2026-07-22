package Data_Date

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Int "gopurs/output/Data.Int"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
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
	return fromJust
}

var greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		greaterThan = gopurs_runtime.Func(func(a1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(a1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(a2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(a1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(a2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], a1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), a2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")
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
	})
	return greaterThan
}

var Date gopurs_runtime.Value
var once_Date sync.Once
func Get_Date() gopurs_runtime.Value {
	once_Date.Do(func() {
		Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": value0, "value1": value1, "value2": value2})
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
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return year
}

var weekday gopurs_runtime.Value
var once_weekday sync.Once
func Get_weekday() gopurs_runtime.Value {
	once_weekday.Do(func() {
		weekday = gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
n_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), Get_calcWeekday()), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t1), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), n_2_0), gopurs_runtime.Int(0))).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Sunday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Monday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuesday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Wednesday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Thursday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Friday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Saturday")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(n_2_0.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Sunday")})
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
}))
	})
	return weekday
}

var showDate gopurs_runtime.Value
var once_showDate sync.Once
func Get_showDate() gopurs_runtime.Value {
	once_showDate.Do(func() {
		showDate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Str("January")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Str("February")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Str("March")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Str("April")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Str("May")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Str("June")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Str("July")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Str("August")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Str("September")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Str("October")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Str("November")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t0 = gopurs_runtime.Str("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Date ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Year ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), __t0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Day ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"])), gopurs_runtime.Str(")")))), gopurs_runtime.Str(")")))))))
})})
	})
	return showDate
}

var month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		month = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return month
}

var isLeapYear gopurs_runtime.Value
var once_isLeapYear sync.Once
func Get_isLeapYear() gopurs_runtime.Value {
	once_isLeapYear.Do(func() {
		isLeapYear = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), y_0), gopurs_runtime.Int(4))), gopurs_runtime.Int(0))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), y_0), gopurs_runtime.Int(400))), gopurs_runtime.Int(0))), gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), y_0), gopurs_runtime.Int(100))), gopurs_runtime.Int(0)))))
})
	})
	return isLeapYear
}

var lastDayOfMonth gopurs_runtime.Value
var once_lastDayOfMonth sync.Once
func Get_lastDayOfMonth() gopurs_runtime.Value {
	once_lastDayOfMonth.Do(func() {
		lastDayOfMonth = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
}))
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(Get_isLeapYear(), y_0)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(29)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(28)))
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(30)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(30)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(30)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(30)))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(31)))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
})
	})
	return lastDayOfMonth
}

var eqDate gopurs_runtime.Value
var once_eqDate sync.Once
func Get_eqDate() gopurs_runtime.Value {
	once_eqDate.Do(func() {
		eqDate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December").IntVal != 0)
}
end_branch_0:
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), __t0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]))
})
})})
	})
	return eqDate
}

var ordDate gopurs_runtime.Value
var once_ordDate sync.Once
func Get_ordDate() gopurs_runtime.Value {
	once_ordDate.Do(func() {
		ordDate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_3
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Date_Component.Get_ordMonth().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v1_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDate()
})})
	})
	return ordDate
}

var enumDate gopurs_runtime.Value
var once_enumDate sync.Once
func Get_enumDate() gopurs_runtime.Value {
	once_enumDate.Do(func() {
		enumDate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
sm_1_0 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_enumMonth().PtrVal.(map[string]gopurs_runtime.Value)["succ"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
v1_2_1 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Int(1)))
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(Get_greaterThan(), v1_2_1), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_lastDayOfMonth(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])}))).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_3
} else {

}
}
{
__t3 = v1_2_1
}
end_branch_3:
sd_3_2 := __t3
var __t5 gopurs_runtime.Value
{
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(sd_3_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(sd_3_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(sm_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(sm_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t7 = gopurs_runtime.Bool(false)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), __t6), __t7)).IntVal != 0 {
__t5 = gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumYear().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
}
end_branch_5:
__local_var_4_4 := __t5
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t10 gopurs_runtime.Value
{
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(sd_3_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t11 = gopurs_runtime.Bool(true)
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(sd_3_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t11 = gopurs_runtime.Bool(false)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
if (__t11).IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(sm_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")})
goto end_branch_12
} else {

}
}
{
if (gopurs_runtime.Bool(sm_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t12 = sm_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t10 = __t12
goto end_branch_10
} else {

}
}
{
__t10 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}
end_branch_10:
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Date(), __local_var_4_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t10)
var __t14 gopurs_runtime.Value
{
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(sd_3_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t15 = gopurs_runtime.Bool(true)
goto end_branch_15
} else {

}
}
{
if (gopurs_runtime.Bool(sd_3_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t15 = gopurs_runtime.Bool(false)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
if (__t15).IntVal != 0 {
__t14 = gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(1))
goto end_branch_14
} else {

}
}
{
__t14 = sd_3_2
}
end_branch_14:
__local_var_6_13 := __t14
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_13.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_5_9, __local_var_6_13.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_16:
__t8 = __t16
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_8:
return __t8
}), "pred": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
pm_1_17 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_enumMonth().PtrVal.(map[string]gopurs_runtime.Value)["pred"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
pd_2_18 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Int(1)))
var __t20 gopurs_runtime.Value
{
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t21 = gopurs_runtime.Bool(true)
goto end_branch_21
} else {

}
}
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t21 = gopurs_runtime.Bool(false)
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
if (__t21).IntVal != 0 {
var __t22 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(pm_1_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t22 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")})
goto end_branch_22
} else {

}
}
{
if (gopurs_runtime.Bool(pm_1_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t22 = pm_1_17.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t20 = __t22
goto end_branch_20
} else {

}
}
{
__t20 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}
end_branch_20:
m_prime_3_19 := __t20
l_4_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_lastDayOfMonth(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), m_prime_3_19)
var __t25 gopurs_runtime.Value
{
var __t26 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t26 = gopurs_runtime.Bool(true)
goto end_branch_26
} else {

}
}
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t26 = gopurs_runtime.Bool(false)
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
var __t27 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(pm_1_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t27 = gopurs_runtime.Bool(true)
goto end_branch_27
} else {

}
}
{
if (gopurs_runtime.Bool(pm_1_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t27 = gopurs_runtime.Bool(false)
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), __t26), __t27)).IntVal != 0 {
__t25 = gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumYear().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Int(1)))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
}
end_branch_25:
__local_var_5_24 := __t25
var __t28 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_24.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_29 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Date(), __local_var_5_24.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), m_prime_3_19)
var __t30 gopurs_runtime.Value
{
var __t31 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t31 = gopurs_runtime.Bool(true)
goto end_branch_31
} else {

}
}
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t31 = gopurs_runtime.Bool(false)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
if (__t31).IntVal != 0 {
__t30 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_6_29, l_4_23)})
goto end_branch_30
} else {

}
}
{
if (gopurs_runtime.Bool(pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t30 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_6_29, pd_2_18.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_30:
__t28 = __t30
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_28:
return __t28
}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
})})
	})
	return enumDate
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func(func(dictDuration_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
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
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(dictDuration_0.PtrVal.(map[string]gopurs_runtime.Value)["toDuration"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn6(), Get_calcDiff()), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t0), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t1), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"]))
})
})
})
	})
	return diff
}

var day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		day = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
})
	})
	return day
}

var canonicalDate gopurs_runtime.Value
var once_canonicalDate sync.Once
func Get_canonicalDate() gopurs_runtime.Value {
	once_canonicalDate.Do(func() {
		canonicalDate = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), Get_canonicalDateImpl()), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("February")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("March")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("April")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("May")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("June")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("July")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("August")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("September")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("October")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("November")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": y_prime_4, "value1": __t0, "value2": d_prime_6})
})
})
})
}))), y_0), __t1), d_2)
})
})
})
	})
	return canonicalDate
}

var exactDate gopurs_runtime.Value
var once_exactDate sync.Once
func Get_exactDate() gopurs_runtime.Value {
	once_exactDate.Do(func() {
		exactDate = gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqDate().PtrVal.(map[string]gopurs_runtime.Value)["eq"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_canonicalDate(), y_0), m_1), d_2)), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": y_0, "value1": m_1, "value2": d_2}))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": y_0, "value1": m_1, "value2": d_2})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_0:
return __t0
})
})
})
	})
	return exactDate
}

var boundedDate gopurs_runtime.Value
var once_boundedDate sync.Once
func Get_boundedDate() gopurs_runtime.Value {
	once_boundedDate.Do(func() {
		boundedDate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": pkg_Data_Date_Component.Get_boundedYear().PtrVal.(map[string]gopurs_runtime.Value)["bottom"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")}), "value2": gopurs_runtime.Int(1)}), "top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": gopurs_runtime.Int(275759), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")}), "value2": gopurs_runtime.Int(31)}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDate()
})})
	})
	return boundedDate
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(date_1 gopurs_runtime.Value) gopurs_runtime.Value {
var adj_2_0 gopurs_runtime.Value
adj_2_0 = gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t19 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v2_4})
goto end_branch_19
} else {

}
}
{
j_5_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_3), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
low_6_2 := gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], j_5_1), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")
var __t4 gopurs_runtime.Value
{
if (low_6_2).IntVal != 0 {
__local_var_7_5 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_enumMonth().PtrVal.(map[string]gopurs_runtime.Value)["pred"], v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")})
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = __local_var_7_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
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
__t4 = v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}
end_branch_4:
l_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_lastDayOfMonth(), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __t4)
hi_8_7 := gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], j_5_1), l_7_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")
var __t9 gopurs_runtime.Value
{
if (low_6_2).IntVal != 0 {
__t9 = j_5_1
goto end_branch_9
} else {

}
}
{
if (hi_8_7).IntVal != 0 {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), j_5_1), l_7_3)), gopurs_runtime.Int(1))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Int(0)
}
end_branch_9:
__local_var_9_8 := gopurs_runtime.Apply(adj_2_0, __t9)
var __t14 gopurs_runtime.Value
{
if (low_6_2).IntVal != 0 {
__local_var_10_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Date(), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
__local_var_11_16 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Int(1))
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_16.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t17 = gopurs_runtime.Apply(Get_enumDate().PtrVal.(map[string]gopurs_runtime.Value)["pred"], gopurs_runtime.Apply(__local_var_10_15, __local_var_11_16.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_17:
__t14 = __t17
goto end_branch_14
} else {

}
}
{
if (hi_8_7).IntVal != 0 {
__t14 = gopurs_runtime.Apply(Get_enumDate().PtrVal.(map[string]gopurs_runtime.Value)["succ"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": l_7_3}))
goto end_branch_14
} else {

}
}
{
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Date(), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
__local_var_11_12 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], j_5_1)
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_12.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_10_11, __local_var_11_12.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_13:
__t14 = __t13
}
end_branch_14:
__local_var_10_10 := __t14
var __t18 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t18 = gopurs_runtime.Apply(__local_var_9_8, __local_var_10_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_18
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_10_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t18 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t19 = __t18
}
end_branch_19:
return __t19
})
})
__local_var_3_20 := gopurs_runtime.Apply(pkg_Data_Int.Get_fromNumber(), v_0)
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_20.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t21 = gopurs_runtime.Apply(gopurs_runtime.Apply(adj_2_0, __local_var_3_20.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), date_1)
goto end_branch_21
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_20.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t21 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
})
})
	})
	return adjust
}

func Get_calcDiff() gopurs_runtime.Value {
	return CalcDiff
}

func Get_calcWeekday() gopurs_runtime.Value {
	return CalcWeekday
}

func Get_canonicalDateImpl() gopurs_runtime.Value {
	return CanonicalDateImpl
}
