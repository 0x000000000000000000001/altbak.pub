package Data_Interval_Duration

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	unsafe "unsafe"
)

var cache_Second gopurs_runtime.Value
var once_Second sync.Once
func Get_Second() gopurs_runtime.Value {
	once_Second.Do(func() {
		cache_Second = gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: nil}
	})
	return cache_Second
}

var cache_Minute gopurs_runtime.Value
var once_Minute sync.Once
func Get_Minute() gopurs_runtime.Value {
	once_Minute.Do(func() {
		cache_Minute = gopurs_runtime.Value{Type: 9, IntVal: 217821258, UnsafePtr: nil}
	})
	return cache_Minute
}

var cache_Hour gopurs_runtime.Value
var once_Hour sync.Once
func Get_Hour() gopurs_runtime.Value {
	once_Hour.Do(func() {
		cache_Hour = gopurs_runtime.Value{Type: 9, IntVal: 1292308612, UnsafePtr: nil}
	})
	return cache_Hour
}

var cache_Day gopurs_runtime.Value
var once_Day sync.Once
func Get_Day() gopurs_runtime.Value {
	once_Day.Do(func() {
		cache_Day = gopurs_runtime.Value{Type: 9, IntVal: 2311060696, UnsafePtr: nil}
	})
	return cache_Day
}

var cache_Week gopurs_runtime.Value
var once_Week sync.Once
func Get_Week() gopurs_runtime.Value {
	once_Week.Do(func() {
		cache_Week = gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: nil}
	})
	return cache_Week
}

var cache_Month gopurs_runtime.Value
var once_Month sync.Once
func Get_Month() gopurs_runtime.Value {
	once_Month.Do(func() {
		cache_Month = gopurs_runtime.Value{Type: 9, IntVal: 3327533908, UnsafePtr: nil}
	})
	return cache_Month
}

var cache_Year gopurs_runtime.Value
var once_Year sync.Once
func Get_Year() gopurs_runtime.Value {
	once_Year.Do(func() {
		cache_Year = gopurs_runtime.Value{Type: 9, IntVal: 3631736139, UnsafePtr: nil}
	})
	return cache_Year
}

var cache_Duration gopurs_runtime.Value
var once_Duration sync.Once
func Get_Duration() gopurs_runtime.Value {
	once_Duration.Do(func() {
		cache_Duration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Duration(x_0_box)
})
	})
	return cache_Duration
}

var cache_showDurationComponent gopurs_runtime.Value
var once_showDurationComponent sync.Once
func Get_showDurationComponent() gopurs_runtime.Value {
	once_showDurationComponent.Do(func() {
		cache_showDurationComponent = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))))
	})
	return cache_showDurationComponent
}

var cache_show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		cache_show = gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_showMap(), Get_showDurationComponent(), pkg_Data_Show.Get_showNumber()), "show")
	})
	return cache_show
}

var cache_showDuration gopurs_runtime.Value
var once_showDuration sync.Once
func Get_showDuration() gopurs_runtime.Value {
	once_showDuration.Do(func() {
		cache_showDuration = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Duration "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_show(), v_0), gopurs_runtime.Str(")")))
}))))
	})
	return cache_showDuration
}

var cache_newtypeDuration gopurs_runtime.Value
var once_newtypeDuration sync.Once
func Get_newtypeDuration() gopurs_runtime.Value {
	once_newtypeDuration.Do(func() {
		cache_newtypeDuration = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))))
	})
	return cache_newtypeDuration
}

var cache_eqDurationComponent gopurs_runtime.Value
var once_eqDurationComponent sync.Once
func Get_eqDurationComponent() gopurs_runtime.Value {
	once_eqDurationComponent.Do(func() {
		cache_eqDurationComponent = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 3631736139)) && ((y_1.Type == 9 && y_1.IntVal == 3631736139)))
}
end_branch_0:
return __t0
}))))
	})
	return cache_eqDurationComponent
}

var cache_eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		cache_eq = gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), Get_eqDurationComponent(), pkg_Data_Eq.Get_eqNumber()), "eq")
	})
	return cache_eq
}

var cache_ordDurationComponent gopurs_runtime.Value
var once_ordDurationComponent sync.Once
func Get_ordDurationComponent() gopurs_runtime.Value {
	once_ordDurationComponent.Do(func() {
		cache_ordDurationComponent = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDurationComponent()
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3908053364) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3908053364) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil})
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3908053364) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 217821258) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 217821258) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil})
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 217821258) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1292308612) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1292308612) {
__t3 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil})
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1292308612) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2311060696) {
var __t4 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2311060696) {
__t4 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil})
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2311060696) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 401302776) {
var __t5 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 401302776) {
__t5 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil})
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 401302776) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3327533908) {
var __t6 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3327533908) {
__t6 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil})
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3327533908) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 3631736139)) && ((y_1.Type == 9 && y_1.IntVal == 3631736139)) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))))
	})
	return cache_ordDurationComponent
}

var cache_compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		cache_compare = gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_ordMap(), Get_ordDurationComponent(), pkg_Data_Ord.Get_ordNumber()), "compare")
	})
	return cache_compare
}

var cache_semigroupDuration gopurs_runtime.Value
var once_semigroupDuration sync.Once
func Get_semigroupDuration() gopurs_runtime.Value {
	once_semigroupDuration.Do(func() {
		cache_semigroupDuration = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_ordDurationComponent(), "compare"), gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))))
	})
	return cache_semigroupDuration
}

var cache_monoidDuration gopurs_runtime.Value
var once_monoidDuration sync.Once
func Get_monoidDuration() gopurs_runtime.Value {
	once_monoidDuration.Do(func() {
		cache_monoidDuration = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDuration()
}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}))))
	})
	return cache_monoidDuration
}

var cache_eqDuration gopurs_runtime.Value
var once_eqDuration sync.Once
func Get_eqDuration() gopurs_runtime.Value {
	once_eqDuration.Do(func() {
		cache_eqDuration = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_eq(), x_0, y_1)
}))))
	})
	return cache_eqDuration
}

var cache_ordDuration gopurs_runtime.Value
var once_ordDuration sync.Once
func Get_ordDuration() gopurs_runtime.Value {
	once_ordDuration.Do(func() {
		cache_ordDuration = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDuration()
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_compare(), x_0, y_1)
}))))
	})
	return cache_ordDuration
}

var cache_hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		cache_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_hour(v_0_box.FloatVal()))}
})
	})
	return cache_hour
}

var cache_millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		cache_millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_millisecond(x_0_box.FloatVal()))}
})
	})
	return cache_millisecond
}

var cache_minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		cache_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_minute(v_0_box.FloatVal()))}
})
	})
	return cache_minute
}

var cache_month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		cache_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_month(v_0_box.FloatVal()))}
})
	})
	return cache_month
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_second(v_0_box.FloatVal()))}
})
	})
	return cache_second
}

var cache_week gopurs_runtime.Value
var once_week sync.Once
func Get_week() gopurs_runtime.Value {
	once_week.Do(func() {
		cache_week = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_week(v_0_box.FloatVal()))}
})
	})
	return cache_week
}

var cache_year gopurs_runtime.Value
var once_year sync.Once
func Get_year() gopurs_runtime.Value {
	once_year.Do(func() {
		cache_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_year(v_0_box.FloatVal()))}
})
	})
	return cache_year
}

var cache_day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		cache_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_day(v_0_box.FloatVal()))}
})
	})
	return cache_day
}

type Constructor_Second struct {
	
}


type Constructor_Minute struct {
	
}


type Constructor_Hour struct {
	
}


type Constructor_Day struct {
	
}


type Constructor_Week struct {
	
}


type Constructor_Month struct {
	
}


type Constructor_Year struct {
	
}


func Call_Duration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_hour(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 1292308612, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_millisecond(x_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var x_0 float64 = x_0_loop
_ = x_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Float(x_0), gopurs_runtime.Float(1000.0))), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_minute(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 217821258, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_month(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3327533908, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_second(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_week(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_year(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3631736139, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}

func Call_day(v_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, float64])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 2311060696, UnsafePtr: nil}, gopurs_runtime.UnboxAny(gopurs_runtime.Float(v_0)), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr), (*pkg_Data_Map_Internal.Constructor_Node[interface{}, interface{}])(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}).UnsafePtr)})}.UnsafePtr)
}
