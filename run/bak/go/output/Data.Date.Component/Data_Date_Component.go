package Data_Date_Component

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var Monday gopurs_runtime.Value
var once_Monday sync.Once
func Get_Monday() gopurs_runtime.Value {
	once_Monday.Do(func() {
		Monday = gopurs_runtime.Constructor0("Monday")
	})
	return Monday
}

var Tuesday gopurs_runtime.Value
var once_Tuesday sync.Once
func Get_Tuesday() gopurs_runtime.Value {
	once_Tuesday.Do(func() {
		Tuesday = gopurs_runtime.Constructor0("Tuesday")
	})
	return Tuesday
}

var Wednesday gopurs_runtime.Value
var once_Wednesday sync.Once
func Get_Wednesday() gopurs_runtime.Value {
	once_Wednesday.Do(func() {
		Wednesday = gopurs_runtime.Constructor0("Wednesday")
	})
	return Wednesday
}

var Thursday gopurs_runtime.Value
var once_Thursday sync.Once
func Get_Thursday() gopurs_runtime.Value {
	once_Thursday.Do(func() {
		Thursday = gopurs_runtime.Constructor0("Thursday")
	})
	return Thursday
}

var Friday gopurs_runtime.Value
var once_Friday sync.Once
func Get_Friday() gopurs_runtime.Value {
	once_Friday.Do(func() {
		Friday = gopurs_runtime.Constructor0("Friday")
	})
	return Friday
}

var Saturday gopurs_runtime.Value
var once_Saturday sync.Once
func Get_Saturday() gopurs_runtime.Value {
	once_Saturday.Do(func() {
		Saturday = gopurs_runtime.Constructor0("Saturday")
	})
	return Saturday
}

var Sunday gopurs_runtime.Value
var once_Sunday sync.Once
func Get_Sunday() gopurs_runtime.Value {
	once_Sunday.Do(func() {
		Sunday = gopurs_runtime.Constructor0("Sunday")
	})
	return Sunday
}

var January gopurs_runtime.Value
var once_January sync.Once
func Get_January() gopurs_runtime.Value {
	once_January.Do(func() {
		January = gopurs_runtime.Constructor0("January")
	})
	return January
}

var February gopurs_runtime.Value
var once_February sync.Once
func Get_February() gopurs_runtime.Value {
	once_February.Do(func() {
		February = gopurs_runtime.Constructor0("February")
	})
	return February
}

var March gopurs_runtime.Value
var once_March sync.Once
func Get_March() gopurs_runtime.Value {
	once_March.Do(func() {
		March = gopurs_runtime.Constructor0("March")
	})
	return March
}

var April gopurs_runtime.Value
var once_April sync.Once
func Get_April() gopurs_runtime.Value {
	once_April.Do(func() {
		April = gopurs_runtime.Constructor0("April")
	})
	return April
}

var May gopurs_runtime.Value
var once_May sync.Once
func Get_May() gopurs_runtime.Value {
	once_May.Do(func() {
		May = gopurs_runtime.Constructor0("May")
	})
	return May
}

var June gopurs_runtime.Value
var once_June sync.Once
func Get_June() gopurs_runtime.Value {
	once_June.Do(func() {
		June = gopurs_runtime.Constructor0("June")
	})
	return June
}

var July gopurs_runtime.Value
var once_July sync.Once
func Get_July() gopurs_runtime.Value {
	once_July.Do(func() {
		July = gopurs_runtime.Constructor0("July")
	})
	return July
}

var August gopurs_runtime.Value
var once_August sync.Once
func Get_August() gopurs_runtime.Value {
	once_August.Do(func() {
		August = gopurs_runtime.Constructor0("August")
	})
	return August
}

var September gopurs_runtime.Value
var once_September sync.Once
func Get_September() gopurs_runtime.Value {
	once_September.Do(func() {
		September = gopurs_runtime.Constructor0("September")
	})
	return September
}

var October gopurs_runtime.Value
var once_October sync.Once
func Get_October() gopurs_runtime.Value {
	once_October.Do(func() {
		October = gopurs_runtime.Constructor0("October")
	})
	return October
}

var November gopurs_runtime.Value
var once_November sync.Once
func Get_November() gopurs_runtime.Value {
	once_November.Do(func() {
		November = gopurs_runtime.Constructor0("November")
	})
	return November
}

var December gopurs_runtime.Value
var once_December sync.Once
func Get_December() gopurs_runtime.Value {
	once_December.Do(func() {
		December = gopurs_runtime.Constructor0("December")
	})
	return December
}

var showYear gopurs_runtime.Value
var once_showYear sync.Once
func Get_showYear() gopurs_runtime.Value {
	once_showYear.Do(func() {
		showYear = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Year ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showYear
}

var showWeekday gopurs_runtime.Value
var once_showWeekday sync.Once
func Get_showWeekday() gopurs_runtime.Value {
	once_showWeekday.Do(func() {
		showWeekday = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "Monday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Monday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Tuesday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Tuesday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Wednesday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Wednesday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Thursday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Thursday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Friday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Friday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Saturday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Saturday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Sunday")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Sunday")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return showWeekday
}

var showMonth gopurs_runtime.Value
var once_showMonth sync.Once
func Get_showMonth() gopurs_runtime.Value {
	once_showMonth.Do(func() {
		showMonth = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Str("January")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Str("February")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Str("March")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Str("April")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Str("May")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Str("June")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Str("July")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Str("August")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Str("September")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Str("October")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Str("November")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "December")).IntVal != 0 {
__t0 = gopurs_runtime.Str("December")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return showMonth
}

var showDay gopurs_runtime.Value
var once_showDay sync.Once
func Get_showDay() gopurs_runtime.Value {
	once_showDay.Do(func() {
		showDay = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Day ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showDay
}

var ordYear gopurs_runtime.Value
var once_ordYear sync.Once
func Get_ordYear() gopurs_runtime.Value {
	once_ordYear.Do(func() {
		ordYear = pkg_Data_Ord.Get_ordInt()
	})
	return ordYear
}

var ordDay gopurs_runtime.Value
var once_ordDay sync.Once
func Get_ordDay() gopurs_runtime.Value {
	once_ordDay.Do(func() {
		ordDay = pkg_Data_Ord.Get_ordInt()
	})
	return ordDay
}

var eqYear gopurs_runtime.Value
var once_eqYear sync.Once
func Get_eqYear() gopurs_runtime.Value {
	once_eqYear.Do(func() {
		eqYear = pkg_Data_Eq.Get_eqInt()
	})
	return eqYear
}

var eqWeekday gopurs_runtime.Value
var once_eqWeekday sync.Once
func Get_eqWeekday() gopurs_runtime.Value {
	once_eqWeekday.Do(func() {
		eqWeekday = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "Monday")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Monday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Tuesday")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Tuesday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Wednesday")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Wednesday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Thursday")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Thursday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Friday")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Friday")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Saturday")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "Saturday")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "Sunday").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "Sunday").IntVal != 0)
}
end_branch_0:
return __t0
}))
	})
	return eqWeekday
}

var ordWeekday gopurs_runtime.Value
var once_ordWeekday sync.Once
func Get_ordWeekday() gopurs_runtime.Value {
	once_ordWeekday.Do(func() {
		ordWeekday = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "Monday")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "Monday")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("EQ")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("LT")
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "Monday")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Tuesday")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "Tuesday")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("EQ")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("LT")
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "Tuesday")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Wednesday")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "Wednesday")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("EQ")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("LT")
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "Wednesday")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Thursday")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "Thursday")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("EQ")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("LT")
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "Thursday")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Friday")).IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "Friday")).IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("EQ")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("LT")
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "Friday")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Saturday")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "Saturday")).IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("EQ")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Constructor0("LT")
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "Saturday")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "Sunday").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "Sunday").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("EQ")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqWeekday()
}))
	})
	return ordWeekday
}

var eqMonth gopurs_runtime.Value
var once_eqMonth sync.Once
func Get_eqMonth() gopurs_runtime.Value {
	once_eqMonth.Do(func() {
		eqMonth = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "January")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "February")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "March")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "April")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "May")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "June")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "July")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "August")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "September")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "October")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.StrVal == "November")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "December").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "December").IntVal != 0)
}
end_branch_0:
return __t0
}))
	})
	return eqMonth
}

var ordMonth gopurs_runtime.Value
var once_ordMonth sync.Once
func Get_ordMonth() gopurs_runtime.Value {
	once_ordMonth.Do(func() {
		ordMonth = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "January")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("EQ")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("LT")
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "February")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "February")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("EQ")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("LT")
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "March")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "March")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("EQ")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("LT")
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "April")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "April")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("EQ")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("LT")
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "May")).IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "May")).IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("EQ")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("LT")
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "June")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "June")).IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("EQ")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Constructor0("LT")
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "July")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "July")).IntVal != 0 {
__t7 = gopurs_runtime.Constructor0("EQ")
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Constructor0("LT")
}
end_branch_7:
__t0 = __t7
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "August")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "August")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor0("EQ")
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Constructor0("LT")
}
end_branch_8:
__t0 = __t8
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "September")).IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "September")).IntVal != 0 {
__t9 = gopurs_runtime.Constructor0("EQ")
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Constructor0("LT")
}
end_branch_9:
__t0 = __t9
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "October")).IntVal != 0 {
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "October")).IntVal != 0 {
__t10 = gopurs_runtime.Constructor0("EQ")
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Constructor0("LT")
}
end_branch_10:
__t0 = __t10
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "November")).IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.StrVal == "November")).IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("EQ")
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Constructor0("LT")
}
end_branch_11:
__t0 = __t11
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.StrVal == "December").IntVal != 0 && gopurs_runtime.Bool(y_1.StrVal == "December").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("EQ")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqMonth()
}))
	})
	return ordMonth
}

var eqDay gopurs_runtime.Value
var once_eqDay sync.Once
func Get_eqDay() gopurs_runtime.Value {
	once_eqDay.Do(func() {
		eqDay = pkg_Data_Eq.Get_eqInt()
	})
	return eqDay
}

var boundedYear gopurs_runtime.Value
var once_boundedYear sync.Once
func Get_boundedYear() gopurs_runtime.Value {
	once_boundedYear.Do(func() {
		boundedYear = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(-271820), gopurs_runtime.Int(275759), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return boundedYear
}

var boundedWeekday gopurs_runtime.Value
var once_boundedWeekday sync.Once
func Get_boundedWeekday() gopurs_runtime.Value {
	once_boundedWeekday.Do(func() {
		boundedWeekday = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Constructor0("Monday"), gopurs_runtime.Constructor0("Sunday"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordWeekday()
}))
	})
	return boundedWeekday
}

var boundedMonth gopurs_runtime.Value
var once_boundedMonth sync.Once
func Get_boundedMonth() gopurs_runtime.Value {
	once_boundedMonth.Do(func() {
		boundedMonth = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Constructor0("January"), gopurs_runtime.Constructor0("December"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMonth()
}))
	})
	return boundedMonth
}

var boundedEnumYear gopurs_runtime.Value
var once_boundedEnumYear sync.Once
func Get_boundedEnumYear() gopurs_runtime.Value {
	once_boundedEnumYear.Do(func() {
		boundedEnumYear = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(547580), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(-271820).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(275759).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedYear()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumYear()
}))
	})
	return boundedEnumYear
}

var enumYear gopurs_runtime.Value
var once_enumYear sync.Once
func Get_enumYear() gopurs_runtime.Value {
	once_enumYear.Do(func() {
		enumYear = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(-271820).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(275759).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Int(x_0.IntVal - gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(-271820).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(275759).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumYear
}

var boundedEnumWeekday gopurs_runtime.Value
var once_boundedEnumWeekday sync.Once
func Get_boundedEnumWeekday() gopurs_runtime.Value {
	once_boundedEnumWeekday.Do(func() {
		boundedEnumWeekday = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(7), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Monday"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Tuesday"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Wednesday"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Thursday"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Friday"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Saturday"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Sunday"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "Monday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Tuesday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Wednesday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Thursday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Friday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Saturday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "Sunday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedWeekday()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumWeekday()
}))
	})
	return boundedEnumWeekday
}

var enumWeekday gopurs_runtime.Value
var once_enumWeekday sync.Once
func Get_enumWeekday() gopurs_runtime.Value {
	once_enumWeekday.Do(func() {
		enumWeekday = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "Monday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Tuesday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Wednesday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Thursday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Friday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Saturday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Sunday")).IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__local_var_1_0 := __t1
_ = __local_var_1_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Monday"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Tuesday"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Wednesday"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Thursday"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Friday"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Saturday"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Sunday"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "Monday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(0)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Tuesday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Wednesday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(2)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Thursday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Friday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(4)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Saturday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(5)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "Sunday")).IntVal != 0 {
__t4 = gopurs_runtime.Int(6)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__local_var_1_3 := __t4
_ = __local_var_1_3
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Monday"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Tuesday"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Wednesday"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Thursday"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Friday"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Saturday"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("Sunday"))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordWeekday()
}))
	})
	return enumWeekday
}

var boundedEnumMonth gopurs_runtime.Value
var once_boundedEnumMonth sync.Once
func Get_boundedEnumMonth() gopurs_runtime.Value {
	once_boundedEnumMonth.Do(func() {
		boundedEnumMonth = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(12), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("January"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("February"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("March"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("April"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("May"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("June"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("July"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("August"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("September"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("October"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("November"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("December"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "February")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "March")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "April")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "May")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "June")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "July")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "August")).IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "September")).IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "October")).IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "November")).IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "December")).IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMonth()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMonth()
}))
	})
	return boundedEnumMonth
}

var enumMonth gopurs_runtime.Value
var once_enumMonth sync.Once
func Get_enumMonth() gopurs_runtime.Value {
	once_enumMonth.Do(func() {
		enumMonth = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "January")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "February")).IntVal != 0 {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "March")).IntVal != 0 {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "April")).IntVal != 0 {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "May")).IntVal != 0 {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "June")).IntVal != 0 {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "July")).IntVal != 0 {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "August")).IntVal != 0 {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "September")).IntVal != 0 {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "October")).IntVal != 0 {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "November")).IntVal != 0 {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "December")).IntVal != 0 {
__t1 = gopurs_runtime.Int(13)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__local_var_1_0 := __t1
_ = __local_var_1_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("January"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("February"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("March"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("April"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("May"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("June"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("July"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("August"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("September"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("October"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("November"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_0.IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("December"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.StrVal == "January")).IntVal != 0 {
__t4 = gopurs_runtime.Int(0)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "February")).IntVal != 0 {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "March")).IntVal != 0 {
__t4 = gopurs_runtime.Int(2)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "April")).IntVal != 0 {
__t4 = gopurs_runtime.Int(3)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "May")).IntVal != 0 {
__t4 = gopurs_runtime.Int(4)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "June")).IntVal != 0 {
__t4 = gopurs_runtime.Int(5)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "July")).IntVal != 0 {
__t4 = gopurs_runtime.Int(6)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "August")).IntVal != 0 {
__t4 = gopurs_runtime.Int(7)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "September")).IntVal != 0 {
__t4 = gopurs_runtime.Int(8)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "October")).IntVal != 0 {
__t4 = gopurs_runtime.Int(9)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "November")).IntVal != 0 {
__t4 = gopurs_runtime.Int(10)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(x_0.StrVal == "December")).IntVal != 0 {
__t4 = gopurs_runtime.Int(11)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__local_var_1_3 := __t4
_ = __local_var_1_3
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("January"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("February"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("March"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("April"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("May"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("June"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("July"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("August"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("September"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("October"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("November"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_1_3.IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("December"))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMonth()
}))
	})
	return enumMonth
}

var boundedDay gopurs_runtime.Value
var once_boundedDay sync.Once
func Get_boundedDay() gopurs_runtime.Value {
	once_boundedDay.Do(func() {
		boundedDay = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(1), gopurs_runtime.Int(31), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return boundedDay
}

var boundedEnumDay gopurs_runtime.Value
var once_boundedEnumDay sync.Once
func Get_boundedEnumDay() gopurs_runtime.Value {
	once_boundedEnumDay.Do(func() {
		boundedEnumDay = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(31), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedDay()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumDay()
}))
	})
	return boundedEnumDay
}

var enumDay gopurs_runtime.Value
var once_enumDay sync.Once
func Get_enumDay() gopurs_runtime.Value {
	once_enumDay.Do(func() {
		enumDay = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Int(x_0.IntVal - gopurs_runtime.Int(1).IntVal)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumDay
}


