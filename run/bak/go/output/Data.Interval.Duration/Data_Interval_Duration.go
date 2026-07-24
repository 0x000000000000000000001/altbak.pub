package Data_Interval_Duration

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	unsafe "unsafe"
)

var Second gopurs_runtime.Value
var once_Second sync.Once
func Get_Second() gopurs_runtime.Value {
	once_Second.Do(func() {
		Second = gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Second{})}
	})
	return Second
}

var Minute gopurs_runtime.Value
var once_Minute sync.Once
func Get_Minute() gopurs_runtime.Value {
	once_Minute.Do(func() {
		Minute = gopurs_runtime.Value{Type: 9, IntVal: 217821258, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Minute{})}
	})
	return Minute
}

var Hour gopurs_runtime.Value
var once_Hour sync.Once
func Get_Hour() gopurs_runtime.Value {
	once_Hour.Do(func() {
		Hour = gopurs_runtime.Value{Type: 9, IntVal: 1292308612, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Hour{})}
	})
	return Hour
}

var Day gopurs_runtime.Value
var once_Day sync.Once
func Get_Day() gopurs_runtime.Value {
	once_Day.Do(func() {
		Day = gopurs_runtime.Value{Type: 9, IntVal: 2311060696, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Day{})}
	})
	return Day
}

var Week gopurs_runtime.Value
var once_Week sync.Once
func Get_Week() gopurs_runtime.Value {
	once_Week.Do(func() {
		Week = gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Week{})}
	})
	return Week
}

var Month gopurs_runtime.Value
var once_Month sync.Once
func Get_Month() gopurs_runtime.Value {
	once_Month.Do(func() {
		Month = gopurs_runtime.Value{Type: 9, IntVal: 3327533908, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Month{})}
	})
	return Month
}

var Year gopurs_runtime.Value
var once_Year sync.Once
func Get_Year() gopurs_runtime.Value {
	once_Year.Do(func() {
		Year = gopurs_runtime.Value{Type: 9, IntVal: 3631736139, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Year{})}
	})
	return Year
}

var Duration gopurs_runtime.Value
var once_Duration sync.Once
func Get_Duration() gopurs_runtime.Value {
	once_Duration.Do(func() {
		Duration = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return Duration
}

var showDurationComponent gopurs_runtime.Value
var once_showDurationComponent sync.Once
func Get_showDurationComponent() gopurs_runtime.Value {
	once_showDurationComponent.Do(func() {
		showDurationComponent = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 217821258) {
__t0 = gopurs_runtime.Str("Minute")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3908053364) {
__t0 = gopurs_runtime.Str("Second")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1292308612) {
__t0 = gopurs_runtime.Str("Hour")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2311060696) {
__t0 = gopurs_runtime.Str("Day")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 401302776) {
__t0 = gopurs_runtime.Str("Week")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3327533908) {
__t0 = gopurs_runtime.Str("Month")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3631736139) {
__t0 = gopurs_runtime.Str("Year")
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
	return showDurationComponent
}

var show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		show = gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_showMap(), Get_showDurationComponent(), pkg_Data_Show.Get_showNumber()), "show")
	})
	return show
}

var showDuration gopurs_runtime.Value
var once_showDuration sync.Once
func Get_showDuration() gopurs_runtime.Value {
	once_showDuration.Do(func() {
		showDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Duration " + gopurs_runtime.Apply(Get_show(), v_0).StrVal() + ")")
}))
	})
	return showDuration
}

var newtypeDuration gopurs_runtime.Value
var once_newtypeDuration sync.Once
func Get_newtypeDuration() gopurs_runtime.Value {
	once_newtypeDuration.Do(func() {
		newtypeDuration = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeDuration
}

var eqDurationComponent gopurs_runtime.Value
var once_eqDurationComponent sync.Once
func Get_eqDurationComponent() gopurs_runtime.Value {
	once_eqDurationComponent.Do(func() {
		eqDurationComponent = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3908053364) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 3908053364))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 217821258) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 217821258))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1292308612) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1292308612))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2311060696) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 2311060696))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 401302776) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 401302776))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3327533908) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 3327533908))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool((x_0.Type == 9 && x_0.IntVal == 3631736139) && (y_1.Type == 9 && y_1.IntVal == 3631736139))
}
end_branch_0:
return __t0
}))
	})
	return eqDurationComponent
}

var eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		eq = gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), Get_eqDurationComponent(), pkg_Data_Eq.Get_eqNumber()), "eq")
	})
	return eq
}

var ordDurationComponent gopurs_runtime.Value
var once_ordDurationComponent sync.Once
func Get_ordDurationComponent() gopurs_runtime.Value {
	once_ordDurationComponent.Do(func() {
		ordDurationComponent = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3908053364) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3908053364) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3908053364) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 217821258) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 217821258) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 217821258) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1292308612) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1292308612) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1292308612) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2311060696) {
var __t4 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2311060696) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2311060696) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 401302776) {
var __t5 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 401302776) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 401302776) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3327533908) {
var __t6 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3327533908) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3866105248, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3327533908) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2098047435, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3631736139) && (y_1.Type == 9 && y_1.IntVal == 3631736139) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
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
return Get_eqDurationComponent()
}))
	})
	return ordDurationComponent
}

var compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		compare = gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_ordMap(), Get_ordDurationComponent(), pkg_Data_Ord.Get_ordNumber()), "compare")
	})
	return compare
}

var semigroupDuration gopurs_runtime.Value
var once_semigroupDuration sync.Once
func Get_semigroupDuration() gopurs_runtime.Value {
	once_semigroupDuration.Do(func() {
		semigroupDuration = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_ordDurationComponent(), "compare"), pkg_Data_Semiring.Get_numAdd(), v_0, v1_1)
}))
	})
	return semigroupDuration
}

var monoidDuration gopurs_runtime.Value
var once_monoidDuration sync.Once
func Get_monoidDuration() gopurs_runtime.Value {
	once_monoidDuration.Do(func() {
		monoidDuration = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDuration()
}))
	})
	return monoidDuration
}

var eqDuration gopurs_runtime.Value
var once_eqDuration sync.Once
func Get_eqDuration() gopurs_runtime.Value {
	once_eqDuration.Do(func() {
		eqDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_eq(), x_0, y_1)
}))
	})
	return eqDuration
}

var ordDuration gopurs_runtime.Value
var once_ordDuration sync.Once
func Get_ordDuration() gopurs_runtime.Value {
	once_ordDuration.Do(func() {
		ordDuration = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_compare(), x_0, y_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDuration()
}))
	})
	return ordDuration
}

var hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		hour = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 1292308612, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Hour{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return hour
}

var millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		millisecond = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Second{})}, gopurs_runtime.Float(x_0.FloatVal() / 1000.0), gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return millisecond
}

var minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		minute = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 217821258, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Minute{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return minute
}

var month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		month = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 3327533908, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Month{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return month
}

var second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		second = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Second{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return second
}

var week gopurs_runtime.Value
var once_week sync.Once
func Get_week() gopurs_runtime.Value {
	once_week.Do(func() {
		week = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Week{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return week
}

var year gopurs_runtime.Value
var once_year sync.Once
func Get_year() gopurs_runtime.Value {
	once_year.Do(func() {
		year = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 3631736139, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Year{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return year
}

var day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		day = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1240286680, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{gopurs_runtime.Int(1), gopurs_runtime.Int(1), gopurs_runtime.Value{Type: 9, IntVal: 2311060696, UnsafePtr: unsafe.Pointer(&Data_Data_Interval_Duration_Day{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 1144344694, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Leaf{})}})}
}()
})
	})
	return day
}

type Data_Data_Interval_Duration_Second struct {
	
}
func Is_Data_Data_Interval_Duration_Second(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3908053364
}

type Data_Data_Interval_Duration_Minute struct {
	
}
func Is_Data_Data_Interval_Duration_Minute(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 217821258
}

type Data_Data_Interval_Duration_Hour struct {
	
}
func Is_Data_Data_Interval_Duration_Hour(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1292308612
}

type Data_Data_Interval_Duration_Day struct {
	
}
func Is_Data_Data_Interval_Duration_Day(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2311060696
}

type Data_Data_Interval_Duration_Week struct {
	
}
func Is_Data_Data_Interval_Duration_Week(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 401302776
}

type Data_Data_Interval_Duration_Month struct {
	
}
func Is_Data_Data_Interval_Duration_Month(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3327533908
}

type Data_Data_Interval_Duration_Year struct {
	
}
func Is_Data_Data_Interval_Duration_Year(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3631736139
}


