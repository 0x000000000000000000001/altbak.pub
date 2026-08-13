package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_DateTime_Instant_bottom gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom sync.Once
func Get_Data_DateTime_Instant_bottom() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom.Do(func() {
		cache_Data_DateTime_Instant_bottom = gopurs_runtime.Int(0)
	})
	return cache_Data_DateTime_Instant_bottom
}

var cache_Data_DateTime_Instant_bottom1 gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom1 sync.Once
func Get_Data_DateTime_Instant_bottom1() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom1.Do(func() {
		cache_Data_DateTime_Instant_bottom1 = gopurs_runtime.Int(0)
	})
	return cache_Data_DateTime_Instant_bottom1
}

var cache_Data_DateTime_Instant_bottom2 gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom2 sync.Once
func Get_Data_DateTime_Instant_bottom2() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom2.Do(func() {
		cache_Data_DateTime_Instant_bottom2 = gopurs_runtime.Int(0)
	})
	return cache_Data_DateTime_Instant_bottom2
}

var cache_Data_DateTime_Instant_bottom3 gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom3 sync.Once
func Get_Data_DateTime_Instant_bottom3() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom3.Do(func() {
		cache_Data_DateTime_Instant_bottom3 = gopurs_runtime.Int(0)
	})
	return cache_Data_DateTime_Instant_bottom3
}

var cache_Data_DateTime_Instant_Instant gopurs_runtime.Value
var once_Data_DateTime_Instant_Instant sync.Once
func Get_Data_DateTime_Instant_Instant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_Instant.Do(func() {
		cache_Data_DateTime_Instant_Instant = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_Instant_Instant(x_0_box)
})
	})
	return cache_Data_DateTime_Instant_Instant
}

var cache_Data_DateTime_Instant_unInstant gopurs_runtime.Value
var once_Data_DateTime_Instant_unInstant sync.Once
func Get_Data_DateTime_Instant_unInstant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_unInstant.Do(func() {
		cache_Data_DateTime_Instant_unInstant = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_DateTime_Instant_unInstant(v_0_box.FloatVal()))
})
	})
	return cache_Data_DateTime_Instant_unInstant
}

var cache_Data_DateTime_Instant_toDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_toDateTime sync.Once
func Get_Data_DateTime_Instant_toDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_toDateTime.Do(func() {
		cache_Data_DateTime_Instant_toDateTime = gopurs_runtime.Apply(Get_Data_DateTime_Instant_toDateTimeImpl(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mo_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(mi_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ms_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_0 -> *Constructor_Data_Maybe_Just
__local_var_8_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2309750950(), gopurs_runtime.Int(mo_2.IntVal)))
_ = __local_var_8_0
var __t1 uint32
{
if (__local_var_8_0 != nil) {
__t1 = uint32((__local_var_8_0).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply3(Get_Data_Date_canonicalDate(), gopurs_runtime.Int(y_1.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}, gopurs_runtime.Int(d_3.IntVal))), &Constructor_Data_Time_Time{1, h_4.IntVal, mi_5.IntVal, s_6.IntVal, ms_7.IntVal}})}
})
})
})
})
})
})
})
})))
	})
	return cache_Data_DateTime_Instant_toDateTime
}

var cache_Data_DateTime_Instant_showInstant gopurs_runtime.Value
var once_Data_DateTime_Instant_showInstant sync.Once
func Get_Data_DateTime_Instant_showInstant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_showInstant.Do(func() {
		cache_Data_DateTime_Instant_showInstant = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Instant ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(v_0.FloatVal())).StrVal())) + (")"))
})})}
	})
	return cache_Data_DateTime_Instant_showInstant
}

var cache_Data_DateTime_Instant_ordDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_ordDateTime sync.Once
func Get_Data_DateTime_Instant_ordDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_ordDateTime.Do(func() {
		cache_Data_DateTime_Instant_ordDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordNumber()))}
	})
	return cache_Data_DateTime_Instant_ordDateTime
}

var cache_Data_DateTime_Instant_instant gopurs_runtime.Value
var once_Data_DateTime_Instant_instant sync.Once
func Get_Data_DateTime_Instant_instant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_instant.Do(func() {
		cache_Data_DateTime_Instant_instant = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_DateTime_Instant_instant(v_0_box.FloatVal()))}
})
	})
	return cache_Data_DateTime_Instant_instant
}

var cache_Data_DateTime_Instant_fromDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_fromDateTime sync.Once
func Get_Data_DateTime_Instant_fromDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_fromDateTime.Do(func() {
		cache_Data_DateTime_Instant_fromDateTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_DateTime_Instant_fromDateTime(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_0_box)))
})
	})
	return cache_Data_DateTime_Instant_fromDateTime
}

var cache_Data_DateTime_Instant_fromDate gopurs_runtime.Value
var once_Data_DateTime_Instant_fromDate sync.Once
func Get_Data_DateTime_Instant_fromDate() gopurs_runtime.Value {
	once_Data_DateTime_Instant_fromDate.Do(func() {
		cache_Data_DateTime_Instant_fromDate = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_DateTime_Instant_fromDate(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](d_0_box)))
})
	})
	return cache_Data_DateTime_Instant_fromDate
}

var cache_Data_DateTime_Instant_eqDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_eqDateTime sync.Once
func Get_Data_DateTime_Instant_eqDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_eqDateTime.Do(func() {
		cache_Data_DateTime_Instant_eqDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqNumber()))}
	})
	return cache_Data_DateTime_Instant_eqDateTime
}

var cache_Data_DateTime_Instant_diff gopurs_runtime.Value
var once_Data_DateTime_Instant_diff sync.Once
func Get_Data_DateTime_Instant_diff() gopurs_runtime.Value {
	once_Data_DateTime_Instant_diff.Do(func() {
		cache_Data_DateTime_Instant_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_Instant_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), dt1_1_box.FloatVal(), dt2_2_box.FloatVal())
})
	})
	return cache_Data_DateTime_Instant_diff
}

var cache_Data_DateTime_Instant_boundedInstant gopurs_runtime.Value
var once_Data_DateTime_Instant_boundedInstant sync.Once
func Get_Data_DateTime_Instant_boundedInstant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_boundedInstant.Do(func() {
		cache_Data_DateTime_Instant_boundedInstant = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordNumber()))}
}), gopurs_runtime.Float(-8639977881600000.0), gopurs_runtime.Float(8639977881599999.0)})}
	})
	return cache_Data_DateTime_Instant_boundedInstant
}

func Call_Data_DateTime_Instant_Instant(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_DateTime_Instant_unInstant(v_0_loop float64) float64 {
var v_0 float64 = v_0_loop
_ = v_0
return v_0
}

func Call_Data_DateTime_Instant_instant(v_0_loop float64) *Constructor_Data_Maybe_Just {
var v_0 float64 = v_0_loop
_ = v_0
var __t5 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Float(-8639977881600000.0))
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
var __t_and_4 bool = false
if __t1 {

var __t3 bool
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Float(8639977881599999.0))
if (uint32(__t_tag_2.IntVal) == 380165415) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
__t_and_4 = __t3
}
if __t_and_4 {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Float(v_0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
}

func Call_Data_DateTime_Instant_fromDateTime(v_0_loop *Constructor_Data_DateTime_DateTime) float64 {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
var __t12 int64
{
var __t_tag_0 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = 1
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = 2
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = 3
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = 4
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = 5
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = 6
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = 7
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = 8
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = 9
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = 10
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = 11
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_11) == 3889233761) {
__t12 = 12
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_12:
return gopurs_runtime.UncurriedApp7(Get_Data_DateTime_Instant_fromDateTimeImpl(), gopurs_runtime.Int(((v_0).V0).V0), gopurs_runtime.Int(__t12), gopurs_runtime.Int(((v_0).V0).V2), gopurs_runtime.Int(((v_0).V1).V0), gopurs_runtime.Int(((v_0).V1).V1), gopurs_runtime.Int(((v_0).V1).V2), gopurs_runtime.Int(((v_0).V1).V3)).FloatVal()
}

func Call_Data_DateTime_Instant_fromDate(d_0_loop *Constructor_Data_Date_Date) float64 {
var d_0 *Constructor_Data_Date_Date = d_0_loop
_ = d_0
var __t12 int64
{
var __t_tag_0 uint32 = (d_0).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = 1
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = (d_0).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = 2
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = (d_0).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = 3
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = (d_0).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = 4
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = (d_0).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = 5
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = (d_0).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = 6
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = (d_0).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = 7
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = (d_0).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = 8
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = (d_0).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = 9
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = (d_0).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = 10
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = (d_0).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = 11
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = (d_0).V1
if (uint32(__t_tag_11) == 3889233761) {
__t12 = 12
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_12:
return gopurs_runtime.UncurriedApp7(Get_Data_DateTime_Instant_fromDateTimeImpl(), gopurs_runtime.Int((d_0).V0), gopurs_runtime.Int(__t12), gopurs_runtime.Int((d_0).V2), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0)).FloatVal()
}

func Call_Data_DateTime_Instant_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, dt1_1_loop float64, dt2_2_loop float64) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 float64 = dt1_1_loop
_ = dt1_1
var dt2_2 float64 = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Float((dt1_1) + (-(gopurs_runtime.Float(dt2_2).FloatVal()))))
}

func Get_Data_DateTime_Instant_fromDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_Instant_FromDateTimeImpl
}

func Get_Data_DateTime_Instant_toDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_Instant_ToDateTimeImpl
}
