package Data_Date_Component

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_Monday gopurs_runtime.Value
var once_Monday sync.Once
func Get_Monday() gopurs_runtime.Value {
	once_Monday.Do(func() {
		cache_Monday = gopurs_runtime.Value{Type: 9, IntVal: 2900196686, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Monday{})}
	})
	return cache_Monday
}

var cache_Tuesday gopurs_runtime.Value
var once_Tuesday sync.Once
func Get_Tuesday() gopurs_runtime.Value {
	once_Tuesday.Do(func() {
		cache_Tuesday = gopurs_runtime.Value{Type: 9, IntVal: 20457557, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Tuesday{})}
	})
	return cache_Tuesday
}

var cache_Wednesday gopurs_runtime.Value
var once_Wednesday sync.Once
func Get_Wednesday() gopurs_runtime.Value {
	once_Wednesday.Do(func() {
		cache_Wednesday = gopurs_runtime.Value{Type: 9, IntVal: 4227105004, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Wednesday{})}
	})
	return cache_Wednesday
}

var cache_Thursday gopurs_runtime.Value
var once_Thursday sync.Once
func Get_Thursday() gopurs_runtime.Value {
	once_Thursday.Do(func() {
		cache_Thursday = gopurs_runtime.Value{Type: 9, IntVal: 3818857258, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Thursday{})}
	})
	return cache_Thursday
}

var cache_Friday gopurs_runtime.Value
var once_Friday sync.Once
func Get_Friday() gopurs_runtime.Value {
	once_Friday.Do(func() {
		cache_Friday = gopurs_runtime.Value{Type: 9, IntVal: 2946274527, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Friday{})}
	})
	return cache_Friday
}

var cache_Saturday gopurs_runtime.Value
var once_Saturday sync.Once
func Get_Saturday() gopurs_runtime.Value {
	once_Saturday.Do(func() {
		cache_Saturday = gopurs_runtime.Value{Type: 9, IntVal: 1070786179, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Saturday{})}
	})
	return cache_Saturday
}

var cache_Sunday gopurs_runtime.Value
var once_Sunday sync.Once
func Get_Sunday() gopurs_runtime.Value {
	once_Sunday.Do(func() {
		cache_Sunday = gopurs_runtime.Value{Type: 9, IntVal: 1326716170, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Sunday{})}
	})
	return cache_Sunday
}

var cache_January gopurs_runtime.Value
var once_January sync.Once
func Get_January() gopurs_runtime.Value {
	once_January.Do(func() {
		cache_January = gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_January{})}
	})
	return cache_January
}

var cache_February gopurs_runtime.Value
var once_February sync.Once
func Get_February() gopurs_runtime.Value {
	once_February.Do(func() {
		cache_February = gopurs_runtime.Value{Type: 9, IntVal: 2455627378, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_February{})}
	})
	return cache_February
}

var cache_March gopurs_runtime.Value
var once_March sync.Once
func Get_March() gopurs_runtime.Value {
	once_March.Do(func() {
		cache_March = gopurs_runtime.Value{Type: 9, IntVal: 4162469099, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_March{})}
	})
	return cache_March
}

var cache_April gopurs_runtime.Value
var once_April sync.Once
func Get_April() gopurs_runtime.Value {
	once_April.Do(func() {
		cache_April = gopurs_runtime.Value{Type: 9, IntVal: 1692989816, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_April{})}
	})
	return cache_April
}

var cache_May gopurs_runtime.Value
var once_May sync.Once
func Get_May() gopurs_runtime.Value {
	once_May.Do(func() {
		cache_May = gopurs_runtime.Value{Type: 9, IntVal: 330658827, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_May{})}
	})
	return cache_May
}

var cache_June gopurs_runtime.Value
var once_June sync.Once
func Get_June() gopurs_runtime.Value {
	once_June.Do(func() {
		cache_June = gopurs_runtime.Value{Type: 9, IntVal: 4067355978, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_June{})}
	})
	return cache_June
}

var cache_July gopurs_runtime.Value
var once_July sync.Once
func Get_July() gopurs_runtime.Value {
	once_July.Do(func() {
		cache_July = gopurs_runtime.Value{Type: 9, IntVal: 2276710548, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_July{})}
	})
	return cache_July
}

var cache_August gopurs_runtime.Value
var once_August sync.Once
func Get_August() gopurs_runtime.Value {
	once_August.Do(func() {
		cache_August = gopurs_runtime.Value{Type: 9, IntVal: 243771071, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_August{})}
	})
	return cache_August
}

var cache_September gopurs_runtime.Value
var once_September sync.Once
func Get_September() gopurs_runtime.Value {
	once_September.Do(func() {
		cache_September = gopurs_runtime.Value{Type: 9, IntVal: 215731793, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_September{})}
	})
	return cache_September
}

var cache_October gopurs_runtime.Value
var once_October sync.Once
func Get_October() gopurs_runtime.Value {
	once_October.Do(func() {
		cache_October = gopurs_runtime.Value{Type: 9, IntVal: 8639228, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_October{})}
	})
	return cache_October
}

var cache_November gopurs_runtime.Value
var once_November sync.Once
func Get_November() gopurs_runtime.Value {
	once_November.Do(func() {
		cache_November = gopurs_runtime.Value{Type: 9, IntVal: 49471444, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_November{})}
	})
	return cache_November
}

var cache_December gopurs_runtime.Value
var once_December sync.Once
func Get_December() gopurs_runtime.Value {
	once_December.Do(func() {
		cache_December = gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_December{})}
	})
	return cache_December
}

var cache_showYear gopurs_runtime.Value
var once_showYear sync.Once
func Get_showYear() gopurs_runtime.Value {
	once_showYear.Do(func() {
		cache_showYear = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Year ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showYear
}

var cache_showWeekday gopurs_runtime.Value
var once_showWeekday sync.Once
func Get_showWeekday() gopurs_runtime.Value {
	once_showWeekday.Do(func() {
		cache_showWeekday = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 2900196686) {
__t0 = gopurs_runtime.Str("Monday")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 20457557) {
__t0 = gopurs_runtime.Str("Tuesday")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 4227105004) {
__t0 = gopurs_runtime.Str("Wednesday")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3818857258) {
__t0 = gopurs_runtime.Str("Thursday")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2946274527) {
__t0 = gopurs_runtime.Str("Friday")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1070786179) {
__t0 = gopurs_runtime.Str("Saturday")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1326716170) {
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
	return cache_showWeekday
}

var cache_showMonth gopurs_runtime.Value
var once_showMonth sync.Once
func Get_showMonth() gopurs_runtime.Value {
	once_showMonth.Do(func() {
		cache_showMonth = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1908470532) {
__t0 = gopurs_runtime.Str("January")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2455627378) {
__t0 = gopurs_runtime.Str("February")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 4162469099) {
__t0 = gopurs_runtime.Str("March")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1692989816) {
__t0 = gopurs_runtime.Str("April")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 330658827) {
__t0 = gopurs_runtime.Str("May")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 4067355978) {
__t0 = gopurs_runtime.Str("June")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2276710548) {
__t0 = gopurs_runtime.Str("July")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 243771071) {
__t0 = gopurs_runtime.Str("August")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 215731793) {
__t0 = gopurs_runtime.Str("September")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 8639228) {
__t0 = gopurs_runtime.Str("October")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 49471444) {
__t0 = gopurs_runtime.Str("November")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3889233761) {
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
	return cache_showMonth
}

var cache_showDay gopurs_runtime.Value
var once_showDay sync.Once
func Get_showDay() gopurs_runtime.Value {
	once_showDay.Do(func() {
		cache_showDay = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Day ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showDay
}

var cache_ordYear gopurs_runtime.Value
var once_ordYear sync.Once
func Get_ordYear() gopurs_runtime.Value {
	once_ordYear.Do(func() {
		cache_ordYear = pkg_Data_Ord.Get_ordInt()
	})
	return cache_ordYear
}

var cache_ordDay gopurs_runtime.Value
var once_ordDay sync.Once
func Get_ordDay() gopurs_runtime.Value {
	once_ordDay.Do(func() {
		cache_ordDay = pkg_Data_Ord.Get_ordInt()
	})
	return cache_ordDay
}

var cache_eqYear gopurs_runtime.Value
var once_eqYear sync.Once
func Get_eqYear() gopurs_runtime.Value {
	once_eqYear.Do(func() {
		cache_eqYear = pkg_Data_Eq.Get_eqInt()
	})
	return cache_eqYear
}

var cache_eqWeekday gopurs_runtime.Value
var once_eqWeekday sync.Once
func Get_eqWeekday() gopurs_runtime.Value {
	once_eqWeekday.Do(func() {
		cache_eqWeekday = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2900196686) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2900196686))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 20457557) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 20457557))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4227105004) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 4227105004))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3818857258) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 3818857258))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2946274527) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2946274527))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1070786179) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1070786179))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 1326716170)) && ((y_1.Type == 9 && y_1.IntVal == 1326716170)))
}
end_branch_0:
return __t0
}))
	})
	return cache_eqWeekday
}

var cache_ordWeekday gopurs_runtime.Value
var once_ordWeekday sync.Once
func Get_ordWeekday() gopurs_runtime.Value {
	once_ordWeekday.Do(func() {
		cache_ordWeekday = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2900196686) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2900196686) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2900196686) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 20457557) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 20457557) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 20457557) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4227105004) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 4227105004) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 4227105004) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3818857258) {
var __t4 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3818857258) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3818857258) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2946274527) {
var __t5 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2946274527) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2946274527) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1070786179) {
var __t6 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1070786179) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1070786179) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 1326716170)) && ((y_1.Type == 9 && y_1.IntVal == 1326716170)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
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
	return cache_ordWeekday
}

var cache_eqMonth gopurs_runtime.Value
var once_eqMonth sync.Once
func Get_eqMonth() gopurs_runtime.Value {
	once_eqMonth.Do(func() {
		cache_eqMonth = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1908470532) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1908470532))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2455627378) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2455627378))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4162469099) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 4162469099))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1692989816) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1692989816))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 330658827) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 330658827))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4067355978) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 4067355978))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2276710548) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2276710548))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 243771071) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 243771071))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 215731793) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 215731793))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 8639228) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 8639228))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 49471444) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 49471444))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 3889233761)) && ((y_1.Type == 9 && y_1.IntVal == 3889233761)))
}
end_branch_0:
return __t0
}))
	})
	return cache_eqMonth
}

var cache_ordMonth gopurs_runtime.Value
var once_ordMonth sync.Once
func Get_ordMonth() gopurs_runtime.Value {
	once_ordMonth.Do(func() {
		cache_ordMonth = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1908470532) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1908470532) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1908470532) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2455627378) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2455627378) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2455627378) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4162469099) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 4162469099) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 4162469099) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1692989816) {
var __t4 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1692989816) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1692989816) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 330658827) {
var __t5 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 330658827) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 330658827) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4067355978) {
var __t6 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 4067355978) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 4067355978) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2276710548) {
var __t7 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2276710548) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_7:
__t0 = __t7
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2276710548) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 243771071) {
var __t8 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 243771071) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_8:
__t0 = __t8
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 243771071) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 215731793) {
var __t9 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 215731793) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_9:
__t0 = __t9
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 215731793) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 8639228) {
var __t10 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 8639228) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_10:
__t0 = __t10
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 8639228) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 49471444) {
var __t11 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 49471444) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_11:
__t0 = __t11
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 49471444) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 3889233761)) && ((y_1.Type == 9 && y_1.IntVal == 3889233761)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
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
	return cache_ordMonth
}

var cache_eqDay gopurs_runtime.Value
var once_eqDay sync.Once
func Get_eqDay() gopurs_runtime.Value {
	once_eqDay.Do(func() {
		cache_eqDay = pkg_Data_Eq.Get_eqInt()
	})
	return cache_eqDay
}

var cache_boundedYear gopurs_runtime.Value
var once_boundedYear sync.Once
func Get_boundedYear() gopurs_runtime.Value {
	once_boundedYear.Do(func() {
		cache_boundedYear = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(-271820), gopurs_runtime.Int(275759), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_boundedYear
}

var cache_boundedWeekday gopurs_runtime.Value
var once_boundedWeekday sync.Once
func Get_boundedWeekday() gopurs_runtime.Value {
	once_boundedWeekday.Do(func() {
		cache_boundedWeekday = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Value{Type: 9, IntVal: 2900196686, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Monday{})}, gopurs_runtime.Value{Type: 9, IntVal: 1326716170, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Sunday{})}, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordWeekday()
}))
	})
	return cache_boundedWeekday
}

var cache_boundedMonth gopurs_runtime.Value
var once_boundedMonth sync.Once
func Get_boundedMonth() gopurs_runtime.Value {
	once_boundedMonth.Do(func() {
		cache_boundedMonth = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_January{})}, gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_December{})}, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMonth()
}))
	})
	return cache_boundedMonth
}

var cache_boundedEnumYear gopurs_runtime.Value
var once_boundedEnumYear sync.Once
func Get_boundedEnumYear() gopurs_runtime.Value {
	once_boundedEnumYear.Do(func() {
		cache_boundedEnumYear = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(547580), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (-271820)) && ((n_0.IntVal) <= (275759)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
	return cache_boundedEnumYear
}

var cache_enumYear gopurs_runtime.Value
var once_enumYear sync.Once
func Get_enumYear() gopurs_runtime.Value {
	once_enumYear.Do(func() {
		cache_enumYear = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (x_0.IntVal) + (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (-271820)) && ((__local_var_1_0) <= (275759)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := (x_0.IntVal) - (1)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_1_2) >= (-271820)) && ((__local_var_1_2) <= (275759)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_2)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_enumYear
}

var cache_boundedEnumWeekday gopurs_runtime.Value
var once_boundedEnumWeekday sync.Once
func Get_boundedEnumWeekday() gopurs_runtime.Value {
	once_boundedEnumWeekday.Do(func() {
		cache_boundedEnumWeekday = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(7), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) == (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2900196686, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Monday{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (2) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 20457557, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Tuesday{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (3) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4227105004, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Wednesday{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (4) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3818857258, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Thursday{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (5) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2946274527, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Friday{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (6) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1070786179, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Saturday{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (7) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1326716170, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Sunday{})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 2900196686) {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 20457557) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 4227105004) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3818857258) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2946274527) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1070786179) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1326716170) {
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
	return cache_boundedEnumWeekday
}

var cache_enumWeekday gopurs_runtime.Value
var once_enumWeekday sync.Once
func Get_enumWeekday() gopurs_runtime.Value {
	once_enumWeekday.Do(func() {
		cache_enumWeekday = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2900196686) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 20457557) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4227105004) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3818857258) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2946274527) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1070786179) {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1326716170) {
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
if (__local_var_1_0.IntVal) == (1) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2900196686, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Monday{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (2) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 20457557, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Tuesday{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (3) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4227105004, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Wednesday{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (4) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3818857258, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Thursday{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (5) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2946274527, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Friday{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (6) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1070786179, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Saturday{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (7) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1326716170, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Sunday{})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 2900196686) {
__t4 = gopurs_runtime.Int(0)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 20457557) {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4227105004) {
__t4 = gopurs_runtime.Int(2)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3818857258) {
__t4 = gopurs_runtime.Int(3)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2946274527) {
__t4 = gopurs_runtime.Int(4)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1070786179) {
__t4 = gopurs_runtime.Int(5)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1326716170) {
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
if (__local_var_1_3.IntVal) == (1) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2900196686, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Monday{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (2) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 20457557, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Tuesday{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (3) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4227105004, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Wednesday{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (4) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3818857258, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Thursday{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (5) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2946274527, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Friday{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (6) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1070786179, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Saturday{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (7) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1326716170, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_Sunday{})}})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordWeekday()
}))
	})
	return cache_enumWeekday
}

var cache_boundedEnumMonth gopurs_runtime.Value
var once_boundedEnumMonth sync.Once
func Get_boundedEnumMonth() gopurs_runtime.Value {
	once_boundedEnumMonth.Do(func() {
		cache_boundedEnumMonth = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(12), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) == (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_January{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (2) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2455627378, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_February{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (3) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4162469099, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_March{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (4) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1692989816, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_April{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (5) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 330658827, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_May{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (6) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4067355978, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_June{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (7) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2276710548, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_July{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (8) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 243771071, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_August{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (9) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 215731793, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_September{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (10) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 8639228, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_October{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (11) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 49471444, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_November{})}})}
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) == (12) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_December{})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1908470532) {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2455627378) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 4162469099) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1692989816) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 330658827) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 4067355978) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2276710548) {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 243771071) {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 215731793) {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 8639228) {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 49471444) {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3889233761) {
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
	return cache_boundedEnumMonth
}

var cache_enumMonth gopurs_runtime.Value
var once_enumMonth sync.Once
func Get_enumMonth() gopurs_runtime.Value {
	once_enumMonth.Do(func() {
		cache_enumMonth = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1908470532) {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2455627378) {
__t1 = gopurs_runtime.Int(3)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4162469099) {
__t1 = gopurs_runtime.Int(4)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1692989816) {
__t1 = gopurs_runtime.Int(5)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 330658827) {
__t1 = gopurs_runtime.Int(6)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4067355978) {
__t1 = gopurs_runtime.Int(7)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2276710548) {
__t1 = gopurs_runtime.Int(8)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 243771071) {
__t1 = gopurs_runtime.Int(9)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 215731793) {
__t1 = gopurs_runtime.Int(10)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 8639228) {
__t1 = gopurs_runtime.Int(11)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 49471444) {
__t1 = gopurs_runtime.Int(12)
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3889233761) {
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
if (__local_var_1_0.IntVal) == (1) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_January{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (2) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2455627378, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_February{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (3) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4162469099, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_March{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (4) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1692989816, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_April{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (5) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 330658827, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_May{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (6) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4067355978, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_June{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (7) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2276710548, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_July{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (8) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 243771071, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_August{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (9) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 215731793, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_September{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (10) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 8639228, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_October{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (11) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 49471444, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_November{})}})}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_0.IntVal) == (12) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_December{})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1908470532) {
__t4 = gopurs_runtime.Int(0)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2455627378) {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4162469099) {
__t4 = gopurs_runtime.Int(2)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1692989816) {
__t4 = gopurs_runtime.Int(3)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 330658827) {
__t4 = gopurs_runtime.Int(4)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 4067355978) {
__t4 = gopurs_runtime.Int(5)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2276710548) {
__t4 = gopurs_runtime.Int(6)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 243771071) {
__t4 = gopurs_runtime.Int(7)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 215731793) {
__t4 = gopurs_runtime.Int(8)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 8639228) {
__t4 = gopurs_runtime.Int(9)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 49471444) {
__t4 = gopurs_runtime.Int(10)
goto end_branch_4
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3889233761) {
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
if (__local_var_1_3.IntVal) == (1) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1908470532, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_January{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (2) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2455627378, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_February{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (3) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4162469099, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_March{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (4) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1692989816, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_April{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (5) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 330658827, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_May{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (6) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 4067355978, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_June{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (7) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2276710548, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_July{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (8) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 243771071, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_August{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (9) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 215731793, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_September{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (10) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 8639228, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_October{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (11) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 49471444, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_November{})}})}
goto end_branch_5
} else {

}
}
{
if (__local_var_1_3.IntVal) == (12) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3889233761, UnsafePtr: unsafe.Pointer(&Data_Data_Date_Component_December{})}})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMonth()
}))
	})
	return cache_enumMonth
}

var cache_boundedDay gopurs_runtime.Value
var once_boundedDay sync.Once
func Get_boundedDay() gopurs_runtime.Value {
	once_boundedDay.Do(func() {
		cache_boundedDay = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(1), gopurs_runtime.Int(31), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_boundedDay
}

var cache_boundedEnumDay gopurs_runtime.Value
var once_boundedEnumDay sync.Once
func Get_boundedEnumDay() gopurs_runtime.Value {
	once_boundedEnumDay.Do(func() {
		cache_boundedEnumDay = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(31), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (1)) && ((n_0.IntVal) <= (31)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{n_0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
	return cache_boundedEnumDay
}

var cache_enumDay gopurs_runtime.Value
var once_enumDay sync.Once
func Get_enumDay() gopurs_runtime.Value {
	once_enumDay.Do(func() {
		cache_enumDay = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (x_0.IntVal) + (1)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (1)) && ((__local_var_1_0) <= (31)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := (x_0.IntVal) - (1)
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_1_2) >= (1)) && ((__local_var_1_2) <= (31)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(__local_var_1_2)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return cache_enumDay
}

type Data_Data_Date_Component_Monday struct {
	
}
func Is_Data_Data_Date_Component_Monday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2900196686
}

type Data_Data_Date_Component_Tuesday struct {
	
}
func Is_Data_Data_Date_Component_Tuesday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 20457557
}

type Data_Data_Date_Component_Wednesday struct {
	
}
func Is_Data_Data_Date_Component_Wednesday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 4227105004
}

type Data_Data_Date_Component_Thursday struct {
	
}
func Is_Data_Data_Date_Component_Thursday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3818857258
}

type Data_Data_Date_Component_Friday struct {
	
}
func Is_Data_Data_Date_Component_Friday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2946274527
}

type Data_Data_Date_Component_Saturday struct {
	
}
func Is_Data_Data_Date_Component_Saturday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1070786179
}

type Data_Data_Date_Component_Sunday struct {
	
}
func Is_Data_Data_Date_Component_Sunday(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1326716170
}

type Data_Data_Date_Component_January struct {
	
}
func Is_Data_Data_Date_Component_January(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1908470532
}

type Data_Data_Date_Component_February struct {
	
}
func Is_Data_Data_Date_Component_February(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2455627378
}

type Data_Data_Date_Component_March struct {
	
}
func Is_Data_Data_Date_Component_March(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 4162469099
}

type Data_Data_Date_Component_April struct {
	
}
func Is_Data_Data_Date_Component_April(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1692989816
}

type Data_Data_Date_Component_May struct {
	
}
func Is_Data_Data_Date_Component_May(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 330658827
}

type Data_Data_Date_Component_June struct {
	
}
func Is_Data_Data_Date_Component_June(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 4067355978
}

type Data_Data_Date_Component_July struct {
	
}
func Is_Data_Data_Date_Component_July(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2276710548
}

type Data_Data_Date_Component_August struct {
	
}
func Is_Data_Data_Date_Component_August(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 243771071
}

type Data_Data_Date_Component_September struct {
	
}
func Is_Data_Data_Date_Component_September(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 215731793
}

type Data_Data_Date_Component_October struct {
	
}
func Is_Data_Data_Date_Component_October(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 8639228
}

type Data_Data_Date_Component_November struct {
	
}
func Is_Data_Data_Date_Component_November(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 49471444
}

type Data_Data_Date_Component_December struct {
	
}
func Is_Data_Data_Date_Component_December(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3889233761
}


