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
		cache_showDurationComponent = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
		cache_showDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Duration "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_show(), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showDuration
}

var cache_newtypeDuration gopurs_runtime.Value
var once_newtypeDuration sync.Once
func Get_newtypeDuration() gopurs_runtime.Value {
	once_newtypeDuration.Do(func() {
		cache_newtypeDuration = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDuration
}

var cache_eqDurationComponent gopurs_runtime.Value
var once_eqDurationComponent sync.Once
func Get_eqDurationComponent() gopurs_runtime.Value {
	once_eqDurationComponent.Do(func() {
		cache_eqDurationComponent = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
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
		cache_ordDurationComponent = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDurationComponent()
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3908053364) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3908053364) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3908053364) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 217821258) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 217821258) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 217821258) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1292308612) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1292308612) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1292308612) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2311060696) {
var __t4 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 2311060696) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 2311060696) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 401302776) {
var __t5 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 401302776) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_5:
__t0 = __t5
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 401302776) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3327533908) {
var __t6 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3327533908) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_6:
__t0 = __t6
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3327533908) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 3631736139)) && ((y_1.Type == 9 && y_1.IntVal == 3631736139)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
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
		cache_semigroupDuration = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_ordDurationComponent(), "compare"), gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
	})
	return cache_semigroupDuration
}

var cache_monoidDuration gopurs_runtime.Value
var once_monoidDuration sync.Once
func Get_monoidDuration() gopurs_runtime.Value {
	once_monoidDuration.Do(func() {
		cache_monoidDuration = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDuration()
}), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil})
	})
	return cache_monoidDuration
}

var cache_eqDuration gopurs_runtime.Value
var once_eqDuration sync.Once
func Get_eqDuration() gopurs_runtime.Value {
	once_eqDuration.Do(func() {
		cache_eqDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_eq(), x_0, y_1)
}))
	})
	return cache_eqDuration
}

var cache_ordDuration gopurs_runtime.Value
var once_ordDuration sync.Once
func Get_ordDuration() gopurs_runtime.Value {
	once_ordDuration.Do(func() {
		cache_ordDuration = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDuration()
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_compare(), x_0, y_1)
}))
	})
	return cache_ordDuration
}

var cache_hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		cache_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hour(v_0_box.FloatVal())
})
	})
	return cache_hour
}

var cache_millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		cache_millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_millisecond(x_0_box.FloatVal())
})
	})
	return cache_millisecond
}

var cache_minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		cache_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minute(v_0_box.FloatVal())
})
	})
	return cache_minute
}

var cache_month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		cache_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_month(v_0_box.FloatVal())
})
	})
	return cache_month
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_second(v_0_box.FloatVal())
})
	})
	return cache_second
}

var cache_week gopurs_runtime.Value
var once_week sync.Once
func Get_week() gopurs_runtime.Value {
	once_week.Do(func() {
		cache_week = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_week(v_0_box.FloatVal())
})
	})
	return cache_week
}

var cache_year gopurs_runtime.Value
var once_year sync.Once
func Get_year() gopurs_runtime.Value {
	once_year.Do(func() {
		cache_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_year(v_0_box.FloatVal())
})
	})
	return cache_year
}

var cache_day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		cache_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_day(v_0_box.FloatVal())
})
	})
	return cache_day
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

type Record_alt_gopurs_runtime_Value struct {
	alt gopurs_runtime.Value
}

type Record_ struct {
	
}

type Record_pure_gopurs_runtime_Value struct {
	pure gopurs_runtime.Value
}

type Record_apply_gopurs_runtime_Value struct {
	apply gopurs_runtime.Value
}

type Record_bipure_gopurs_runtime_Value struct {
	bipure gopurs_runtime.Value
}

type Record_biapply_gopurs_runtime_Value struct {
	biapply gopurs_runtime.Value
}

type Record_bind_gopurs_runtime_Value struct {
	bind gopurs_runtime.Value
}

type Record_discard_gopurs_runtime_Value struct {
	discard gopurs_runtime.Value
}

type Record_identity_gopurs_runtime_Value struct {
	identity gopurs_runtime.Value
}

type Record_ask_gopurs_runtime_Value struct {
	ask gopurs_runtime.Value
}

type Record_local_gopurs_runtime_Value struct {
	local gopurs_runtime.Value
}

type Record_peek_gopurs_runtime_Value_pos_gopurs_runtime_Value struct {
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}

type Record_track_gopurs_runtime_Value struct {
	track gopurs_runtime.Value
}

type Record_extract_gopurs_runtime_Value struct {
	extract gopurs_runtime.Value
}

type Record_extend_gopurs_runtime_Value struct {
	extend gopurs_runtime.Value
}

type Record_defer__gopurs_runtime_Value struct {
	defer_ gopurs_runtime.Value
}

type Record_callCC_gopurs_runtime_Value struct {
	callCC gopurs_runtime.Value
}

type Record_catchError_gopurs_runtime_Value struct {
	catchError gopurs_runtime.Value
}

type Record_throwError_gopurs_runtime_Value struct {
	throwError gopurs_runtime.Value
}

type Record_chooseBool_gopurs_runtime_Value_chooseFloat_gopurs_runtime_Value_chooseInt_gopurs_runtime_Value_resize_gopurs_runtime_Value_sized_gopurs_runtime_Value struct {
	chooseBool gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt gopurs_runtime.Value
	resize gopurs_runtime.Value
	sized gopurs_runtime.Value
}

type Record_foldMap1_gopurs_runtime_Value_foldl1_gopurs_runtime_Value_foldr1_gopurs_runtime_Value struct {
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}

type Record_append__gopurs_runtime_Value struct {
	append_ gopurs_runtime.Value
}

type Record_tailRecM_gopurs_runtime_Value struct {
	tailRecM gopurs_runtime.Value
}

type Record_unfoldr_gopurs_runtime_Value struct {
	unfoldr gopurs_runtime.Value
}

type Record_map__gopurs_runtime_Value struct {
	map_ gopurs_runtime.Value
}

type Record_state_gopurs_runtime_Value struct {
	state gopurs_runtime.Value
}

type Record_lift_gopurs_runtime_Value struct {
	lift gopurs_runtime.Value
}

type Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value struct {
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}

type Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value struct {
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}

type Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value struct {
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}

type Record_mempty_gopurs_runtime_Value struct {
	mempty gopurs_runtime.Value
}

type Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value struct {
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}

type Record_empty_gopurs_runtime_Value struct {
	empty gopurs_runtime.Value
}

type Record_compose_gopurs_runtime_Value struct {
	compose gopurs_runtime.Value
}

type Record_eq_gopurs_runtime_Value struct {
	eq gopurs_runtime.Value
}

type Record_compare_gopurs_runtime_Value struct {
	compare gopurs_runtime.Value
}

type Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value struct {
	bifoldMap gopurs_runtime.Value
	bifoldl gopurs_runtime.Value
	bifoldr gopurs_runtime.Value
}

type Record_bimap_gopurs_runtime_Value struct {
	bimap gopurs_runtime.Value
}

type Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value struct {
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}

type Record_genericBottom_prime_gopurs_runtime_Value struct {
	genericBottom_prime gopurs_runtime.Value
}

type Record_genericTop_prime_gopurs_runtime_Value struct {
	genericTop_prime gopurs_runtime.Value
}

type Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value struct {
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}

type Record_lose_gopurs_runtime_Value struct {
	lose gopurs_runtime.Value
}

type Record_choose_gopurs_runtime_Value struct {
	choose gopurs_runtime.Value
}

type Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value struct {
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}

type Record_divide_gopurs_runtime_Value struct {
	divide gopurs_runtime.Value
}

type Record_recip_gopurs_runtime_Value struct {
	recip gopurs_runtime.Value
}

type Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value struct {
	genericCardinality_prime gopurs_runtime.Value
	genericFromEnum_prime gopurs_runtime.Value
	genericToEnum_prime gopurs_runtime.Value
}

type Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value struct {
	genericPred_prime gopurs_runtime.Value
	genericSucc_prime gopurs_runtime.Value
}

type Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value struct {
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}

type Record_unfoldr1_gopurs_runtime_Value struct {
	unfoldr1 gopurs_runtime.Value
}

type Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value struct {
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}

type Record_genericEq_prime_gopurs_runtime_Value struct {
	genericEq_prime gopurs_runtime.Value
}

type Record_eq1_gopurs_runtime_Value struct {
	eq1 gopurs_runtime.Value
}

type Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value struct {
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff gopurs_runtime.Value
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt gopurs_runtime.Value
}

type Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value struct {
	add gopurs_runtime.Value
	mul gopurs_runtime.Value
	one gopurs_runtime.Value
	zero gopurs_runtime.Value
}

type Record_foldMapWithIndex_gopurs_runtime_Value_foldlWithIndex_gopurs_runtime_Value_foldrWithIndex_gopurs_runtime_Value struct {
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}

type Record_cmap_gopurs_runtime_Value struct {
	cmap gopurs_runtime.Value
}

type Record_imap_gopurs_runtime_Value struct {
	imap gopurs_runtime.Value
}

type Record_mapWithIndex_gopurs_runtime_Value struct {
	mapWithIndex gopurs_runtime.Value
}

type Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value struct {
	from gopurs_runtime.Value
	to gopurs_runtime.Value
}

type Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value struct {
	genericConj_prime gopurs_runtime.Value
	genericDisj_prime gopurs_runtime.Value
	genericFF_prime gopurs_runtime.Value
	genericImplies_prime gopurs_runtime.Value
	genericNot_prime gopurs_runtime.Value
	genericTT_prime gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_bool_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_bool struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff bool
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt bool
}

type Record_genericMempty_prime_gopurs_runtime_Value struct {
	genericMempty_prime gopurs_runtime.Value
}

type Record_genericCompare_prime_gopurs_runtime_Value struct {
	genericCompare_prime gopurs_runtime.Value
}

type Record_sub_gopurs_runtime_Value struct {
	sub gopurs_runtime.Value
}

type Record_compare1_gopurs_runtime_Value struct {
	compare1 gopurs_runtime.Value
}

type Record_left_gopurs_runtime_Value_right_gopurs_runtime_Value struct {
	left gopurs_runtime.Value
	right gopurs_runtime.Value
}

type Record_first_gopurs_runtime_Value_second_gopurs_runtime_Value struct {
	first gopurs_runtime.Value
	second gopurs_runtime.Value
}

type Record_dimap_gopurs_runtime_Value struct {
	dimap gopurs_runtime.Value
}

type Record_genericSub_prime_gopurs_runtime_Value struct {
	genericSub_prime gopurs_runtime.Value
}

type Record_genericAppend_prime_gopurs_runtime_Value struct {
	genericAppend_prime gopurs_runtime.Value
}

type Record_sequence1_gopurs_runtime_Value_traverse1_gopurs_runtime_Value struct {
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}

type Record_genericAdd_prime_gopurs_runtime_Value_genericMul_prime_gopurs_runtime_Value_genericOne_prime_gopurs_runtime_Value_genericZero_prime_gopurs_runtime_Value struct {
	genericAdd_prime gopurs_runtime.Value
	genericMul_prime gopurs_runtime.Value
	genericOne_prime gopurs_runtime.Value
	genericZero_prime gopurs_runtime.Value
}

type Record_genericShow_prime_gopurs_runtime_Value struct {
	genericShow_prime gopurs_runtime.Value
}

type Record_genericShowArgs_gopurs_runtime_Value struct {
	genericShowArgs gopurs_runtime.Value
}

type Record_show_gopurs_runtime_Value struct {
	show gopurs_runtime.Value
}

type Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value struct {
	fromDuration gopurs_runtime.Value
	toDuration gopurs_runtime.Value
}

type Record_traverseWithIndex_gopurs_runtime_Value struct {
	traverseWithIndex gopurs_runtime.Value
}

type Record_liftEffect_gopurs_runtime_Value struct {
	liftEffect gopurs_runtime.Value
}

type Record_mappend__gopurs_runtime_Value_mempty__gopurs_runtime_Value struct {
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}

type Record_proof_gopurs_runtime_Value struct {
	proof gopurs_runtime.Value
}

type Record_lower_gopurs_runtime_Value struct {
	lower gopurs_runtime.Value
}

type Record_liftST_gopurs_runtime_Value struct {
	liftST gopurs_runtime.Value
}

type Record_tell_gopurs_runtime_Value struct {
	tell gopurs_runtime.Value
}

type Record_reflectSymbol_gopurs_runtime_Value struct {
	reflectSymbol gopurs_runtime.Value
}

type Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value struct {
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}

type Record_conquer_gopurs_runtime_Value struct {
	conquer gopurs_runtime.Value
}

type Record_inj_gopurs_runtime_Value_prj_gopurs_runtime_Value struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}

type Record_eqRecord_gopurs_runtime_Value struct {
	eqRecord gopurs_runtime.Value
}

type Record_conjRecord_gopurs_runtime_Value_disjRecord_gopurs_runtime_Value_ffRecord_gopurs_runtime_Value_impliesRecord_gopurs_runtime_Value_notRecord_gopurs_runtime_Value_ttRecord_gopurs_runtime_Value struct {
	conjRecord gopurs_runtime.Value
	disjRecord gopurs_runtime.Value
	ffRecord gopurs_runtime.Value
	impliesRecord gopurs_runtime.Value
	notRecord gopurs_runtime.Value
	ttRecord gopurs_runtime.Value
}

type Record_memptyRecord_gopurs_runtime_Value struct {
	memptyRecord gopurs_runtime.Value
}

type Record_compareRecord_gopurs_runtime_Value struct {
	compareRecord gopurs_runtime.Value
}

type Record_closed_gopurs_runtime_Value struct {
	closed gopurs_runtime.Value
}

type Record_unleft_gopurs_runtime_Value_unright_gopurs_runtime_Value struct {
	unleft gopurs_runtime.Value
	unright gopurs_runtime.Value
}

type Record_unfirst_gopurs_runtime_Value_unsecond_gopurs_runtime_Value struct {
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}

type Record_reflectType_gopurs_runtime_Value struct {
	reflectType gopurs_runtime.Value
}

type Record_subRecord_gopurs_runtime_Value struct {
	subRecord gopurs_runtime.Value
}

type Record_appendRecord_gopurs_runtime_Value struct {
	appendRecord gopurs_runtime.Value
}

type Record_addRecord_gopurs_runtime_Value_mulRecord_gopurs_runtime_Value_oneRecord_gopurs_runtime_Value_zeroRecord_gopurs_runtime_Value struct {
	addRecord gopurs_runtime.Value
	mulRecord gopurs_runtime.Value
	oneRecord gopurs_runtime.Value
	zeroRecord gopurs_runtime.Value
}

type Record_showRecordFields_gopurs_runtime_Value struct {
	showRecordFields gopurs_runtime.Value
}

type Record_nes_gopurs_runtime_Value struct {
	nes gopurs_runtime.Value
}

type Record_liftAff_gopurs_runtime_Value struct {
	liftAff gopurs_runtime.Value
}

func Call_Duration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_hour(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 1292308612, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_millisecond(x_0_loop float64) gopurs_runtime.Value {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: nil}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Float(x_0), gopurs_runtime.Float(1000.0)), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_minute(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 217821258, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_month(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3327533908, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_second(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3908053364, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_week(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_year(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 3631736139, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}

func Call_day(v_0_loop float64) gopurs_runtime.Value {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Data_Data_Map_Internal_Node{1, 1, gopurs_runtime.Value{Type: 9, IntVal: 2311060696, UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 687041424, UnsafePtr: nil}})}
}


