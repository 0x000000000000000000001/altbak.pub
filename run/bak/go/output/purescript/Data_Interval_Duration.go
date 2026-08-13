package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Interval_Duration_add gopurs_runtime.Value
var once_Data_Interval_Duration_add sync.Once
func Get_Data_Interval_Duration_add() gopurs_runtime.Value {
	once_Data_Interval_Duration_add.Do(func() {
		cache_Data_Interval_Duration_add = Get_Data_Semiring_numAdd()
	})
	return cache_Data_Interval_Duration_add
}

var cache_Data_Interval_Duration_Second gopurs_runtime.Value
var once_Data_Interval_Duration_Second sync.Once
func Get_Data_Interval_Duration_Second() gopurs_runtime.Value {
	once_Data_Interval_Duration_Second.Do(func() {
		cache_Data_Interval_Duration_Second = gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Second
}

var cache_Data_Interval_Duration_Minute gopurs_runtime.Value
var once_Data_Interval_Duration_Minute sync.Once
func Get_Data_Interval_Duration_Minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_Minute.Do(func() {
		cache_Data_Interval_Duration_Minute = gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Minute
}

var cache_Data_Interval_Duration_Hour gopurs_runtime.Value
var once_Data_Interval_Duration_Hour sync.Once
func Get_Data_Interval_Duration_Hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_Hour.Do(func() {
		cache_Data_Interval_Duration_Hour = gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Hour
}

var cache_Data_Interval_Duration_Day gopurs_runtime.Value
var once_Data_Interval_Duration_Day sync.Once
func Get_Data_Interval_Duration_Day() gopurs_runtime.Value {
	once_Data_Interval_Duration_Day.Do(func() {
		cache_Data_Interval_Duration_Day = gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Day
}

var cache_Data_Interval_Duration_Week gopurs_runtime.Value
var once_Data_Interval_Duration_Week sync.Once
func Get_Data_Interval_Duration_Week() gopurs_runtime.Value {
	once_Data_Interval_Duration_Week.Do(func() {
		cache_Data_Interval_Duration_Week = gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Week
}

var cache_Data_Interval_Duration_Month gopurs_runtime.Value
var once_Data_Interval_Duration_Month sync.Once
func Get_Data_Interval_Duration_Month() gopurs_runtime.Value {
	once_Data_Interval_Duration_Month.Do(func() {
		cache_Data_Interval_Duration_Month = gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Month
}

var cache_Data_Interval_Duration_Year gopurs_runtime.Value
var once_Data_Interval_Duration_Year sync.Once
func Get_Data_Interval_Duration_Year() gopurs_runtime.Value {
	once_Data_Interval_Duration_Year.Do(func() {
		cache_Data_Interval_Duration_Year = gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Year
}

var cache_Data_Interval_Duration_Duration gopurs_runtime.Value
var once_Data_Interval_Duration_Duration sync.Once
func Get_Data_Interval_Duration_Duration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Duration.Do(func() {
		cache_Data_Interval_Duration_Duration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_Duration_Duration(x_0_box)
})
	})
	return cache_Data_Interval_Duration_Duration
}

var cache_Data_Interval_Duration_showDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_showDurationComponent sync.Once
func Get_Data_Interval_Duration_showDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDurationComponent.Do(func() {
		cache_Data_Interval_Duration_showDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (uint32(v_0.IntVal) == 217821258) {
__t0 = "Minute"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 3908053364) {
__t0 = "Second"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 1292308612) {
__t0 = "Hour"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 2311060696) {
__t0 = "Day"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 401302776) {
__t0 = "Week"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 3327533908) {
__t0 = "Month"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 3631736139) {
__t0 = "Year"
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
})})}
	})
	return cache_Data_Interval_Duration_showDurationComponent
}

var cache_Data_Interval_Duration_showMap gopurs_runtime.Value
var once_Data_Interval_Duration_showMap sync.Once
func Get_Data_Interval_Duration_showMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_showMap.Do(func() {
		cache_Data_Interval_Duration_showMap = func() gopurs_runtime.Value {
// TAST (Let): showArray_0_0 -> *Constructor_Data_Show_Show
showArray_0_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 string
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_1.IntVal) == 217821258) {
__t8 = "(Tuple Minute "
goto end_branch_8
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_2.IntVal) == 3908053364) {
__t8 = "(Tuple Second "
goto end_branch_8
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_3.IntVal) == 1292308612) {
__t8 = "(Tuple Hour "
goto end_branch_8
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_4.IntVal) == 2311060696) {
__t8 = "(Tuple Day "
goto end_branch_8
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_5.IntVal) == 401302776) {
__t8 = "(Tuple Week "
goto end_branch_8
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_6.IntVal) == 3327533908) {
__t8 = "(Tuple Month "
goto end_branch_8
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
if (uint32(__t_tag_7.IntVal) == 3631736139) {
__t8 = "(Tuple Year "
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_8:
return gopurs_runtime.Str(((__t8) + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1).StrVal())) + (")"))
}))}
_ = showArray_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(as_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_0_0.V0), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil) {
__t9 = (*Constructor_Data_Maybe_Just)(v_3.UnsafePtr).V0
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](as_1), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")"))
})})}
}()
	})
	return cache_Data_Interval_Duration_showMap
}

var cache_Data_Interval_Duration_showDuration gopurs_runtime.Value
var once_Data_Interval_Duration_showDuration sync.Once
func Get_Data_Interval_Duration_showDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDuration.Do(func() {
		cache_Data_Interval_Duration_showDuration = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Duration (fromFoldable ") + (gopurs_runtime.Apply2(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 string
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_0.IntVal) == 217821258) {
__t7 = "(Tuple Minute "
goto end_branch_7
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_1.IntVal) == 3908053364) {
__t7 = "(Tuple Second "
goto end_branch_7
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_2.IntVal) == 1292308612) {
__t7 = "(Tuple Hour "
goto end_branch_7
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_3.IntVal) == 2311060696) {
__t7 = "(Tuple Day "
goto end_branch_7
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_4.IntVal) == 401302776) {
__t7 = "(Tuple Week "
goto end_branch_7
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_5.IntVal) == 3327533908) {
__t7 = "(Tuple Month "
goto end_branch_7
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_6.IntVal) == 3631736139) {
__t7 = "(Tuple Year "
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_7:
return gopurs_runtime.Str(((__t7) + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1).StrVal())) + (")"))
}), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t8 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + ("))"))
})})}
	})
	return cache_Data_Interval_Duration_showDuration
}

var cache_Data_Interval_Duration_newtypeDuration gopurs_runtime.Value
var once_Data_Interval_Duration_newtypeDuration sync.Once
func Get_Data_Interval_Duration_newtypeDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_newtypeDuration.Do(func() {
		cache_Data_Interval_Duration_newtypeDuration = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Interval_Duration_newtypeDuration
}

var cache_Data_Interval_Duration_eqDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_eqDurationComponent sync.Once
func Get_Data_Interval_Duration_eqDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDurationComponent.Do(func() {
		cache_Data_Interval_Duration_eqDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if (uint32(x_0.IntVal) == 3908053364) {
var __t0 bool
{
if (uint32(y_1.IntVal) == 3908053364) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 217821258) {
var __t1 bool
{
if (uint32(y_1.IntVal) == 217821258) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 1292308612) {
var __t2 bool
{
if (uint32(y_1.IntVal) == 1292308612) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 2311060696) {
var __t3 bool
{
if (uint32(y_1.IntVal) == 2311060696) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 401302776) {
var __t4 bool
{
if (uint32(y_1.IntVal) == 401302776) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 3327533908) {
var __t5 bool
{
if (uint32(y_1.IntVal) == 3327533908) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if ((uint32(x_0.IntVal) == 3631736139)) && ((uint32(y_1.IntVal) == 3631736139)) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
})})}
	})
	return cache_Data_Interval_Duration_eqDurationComponent
}

var cache_Data_Interval_Duration_eqMap gopurs_runtime.Value
var once_Data_Interval_Duration_eqMap sync.Once
func Get_Data_Interval_Duration_eqMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqMap.Do(func() {
		cache_Data_Interval_Duration_eqMap = func() gopurs_runtime.Value {
var go__go_0_1_0 gopurs_runtime.Value
go__go_0_1_0 = gopurs_runtime.Func(func(a_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_1_loop gopurs_runtime.Value = a_1_loop_val
var b_2_loop gopurs_runtime.Value = b_2_loop_val
go__go_0_1_0:
for {
if false { continue go__go_0_1_0 }
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
// TAST (Let): v_3_2 -> *Constructor_Data_Map_Internal_IterNext
v_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_1))
_ = v_3_2
var __t28 bool
{
if (v_3_2 != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Map_Internal_IterNext
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_2))
_ = v2_4_3
var __t27 bool
{
var __t_and_26 bool = false
if (v2_4_3 != nil) {

var __t25 bool
{
var __t_tag_4 gopurs_runtime.Value = (v_3_2).V0
if (uint32(__t_tag_4.IntVal) == 3908053364) {
var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = (v2_4_3).V0
if (uint32(__t_tag_5.IntVal) == 3908053364) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t25 = __t6
goto end_branch_25
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (v_3_2).V0
if (uint32(__t_tag_7.IntVal) == 217821258) {
var __t9 bool
{
var __t_tag_8 gopurs_runtime.Value = (v2_4_3).V0
if (uint32(__t_tag_8.IntVal) == 217821258) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
__t25 = __t9
goto end_branch_25
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = (v_3_2).V0
if (uint32(__t_tag_10.IntVal) == 1292308612) {
var __t12 bool
{
var __t_tag_11 gopurs_runtime.Value = (v2_4_3).V0
if (uint32(__t_tag_11.IntVal) == 1292308612) {
__t12 = true
goto end_branch_12
} else {

}
}
{
__t12 = false
}
end_branch_12:
__t25 = __t12
goto end_branch_25
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (v_3_2).V0
if (uint32(__t_tag_13.IntVal) == 2311060696) {
var __t15 bool
{
var __t_tag_14 gopurs_runtime.Value = (v2_4_3).V0
if (uint32(__t_tag_14.IntVal) == 2311060696) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t25 = __t15
goto end_branch_25
} else {

}
}
{
var __t_tag_16 gopurs_runtime.Value = (v_3_2).V0
if (uint32(__t_tag_16.IntVal) == 401302776) {
var __t18 bool
{
var __t_tag_17 gopurs_runtime.Value = (v2_4_3).V0
if (uint32(__t_tag_17.IntVal) == 401302776) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
__t25 = __t18
goto end_branch_25
} else {

}
}
{
var __t_tag_19 gopurs_runtime.Value = (v_3_2).V0
if (uint32(__t_tag_19.IntVal) == 3327533908) {
var __t21 bool
{
var __t_tag_20 gopurs_runtime.Value = (v2_4_3).V0
if (uint32(__t_tag_20.IntVal) == 3327533908) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
__t25 = __t21
goto end_branch_25
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = (v_3_2).V0
var __t_and_24 bool = false
if (uint32(__t_tag_22.IntVal) == 3631736139) {

var __t_tag_23 gopurs_runtime.Value = (v2_4_3).V0
__t_and_24 = (uint32(__t_tag_23.IntVal) == 3631736139)
}
if __t_and_24 {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = (__t25) && (((v_3_2).V1.FloatVal()) == ((v2_4_3).V1.FloatVal()))
}
if __t_and_26 {
a_1_loop = (v_3_2).V2
b_2_loop = (v2_4_3).V2
continue go__go_0_1_0
__t27 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v_3_2 == nil) {
__t28 = true
goto end_branch_28
} else {

}
}
{
__t28 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_28:
return gopurs_runtime.Bool(__t28)
}
}()
})
})
// TAST (Let): eqMapIter2_0_0 -> *Constructor_Data_Eq_Eq
eqMapIter2_0_0 := &Constructor_Data_Eq_Eq{1, go__go_0_1_0}
_ = eqMapIter2_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 bool
{
if (xs_1.Type == 9 && xs_1.IntVal == 324739070 && xs_1.UnsafePtr == nil) {
var __t29 bool
{
if (ys_2.Type == 9 && ys_2.IntVal == 324739070 && ys_2.UnsafePtr == nil) {
__t29 = true
goto end_branch_29
} else {

}
}
{
__t29 = false
}
end_branch_29:
__t31 = __t29
goto end_branch_31
} else {

}
}
{
if (xs_1.Type == 9 && xs_1.IntVal == 324739070 && xs_1.UnsafePtr != nil) {
var __t30 bool
{
if ((ys_2.Type == 9 && ys_2.IntVal == 324739070 && ys_2.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_1.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_2.UnsafePtr).V1)) {
__t30 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_0_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_1), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_30
} else {

}
}
{
__t30 = false
}
end_branch_30:
__t31 = __t30
goto end_branch_31
} else {

}
}
{
__t31 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_31:
return gopurs_runtime.Bool(__t31)
})
})})}
}()
	})
	return cache_Data_Interval_Duration_eqMap
}

var cache_Data_Interval_Duration_ordDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_ordDurationComponent sync.Once
func Get_Data_Interval_Duration_ordDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDurationComponent.Do(func() {
		cache_Data_Interval_Duration_ordDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_eqDurationComponent()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
{
if (uint32(x_0.IntVal) == 3908053364) {
var __t0 uint32
{
if (uint32(y_1.IntVal) == 3908053364) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 3908053364) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 217821258) {
var __t1 uint32
{
if (uint32(y_1.IntVal) == 217821258) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 217821258) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 1292308612) {
var __t2 uint32
{
if (uint32(y_1.IntVal) == 1292308612) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 1292308612) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 2311060696) {
var __t3 uint32
{
if (uint32(y_1.IntVal) == 2311060696) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 2311060696) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 401302776) {
var __t4 uint32
{
if (uint32(y_1.IntVal) == 401302776) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 401302776) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 3327533908) {
var __t5 uint32
{
if (uint32(y_1.IntVal) == 3327533908) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 3327533908) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if ((uint32(x_0.IntVal) == 3631736139)) && ((uint32(y_1.IntVal) == 3631736139)) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Interval_Duration_ordDurationComponent
}

var cache_Data_Interval_Duration_ordMap gopurs_runtime.Value
var once_Data_Interval_Duration_ordMap sync.Once
func Get_Data_Interval_Duration_ordMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordMap.Do(func() {
		cache_Data_Interval_Duration_ordMap = func() gopurs_runtime.Value {
var go__go_0_2_1 gopurs_runtime.Value
go__go_0_2_1 = gopurs_runtime.Func(func(a_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_1_loop gopurs_runtime.Value = a_1_loop_val
var b_2_loop gopurs_runtime.Value = b_2_loop_val
go__go_0_2_1:
for {
if false { continue go__go_0_2_1 }
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
// TAST (Let): v_3_3 -> *Constructor_Data_Map_Internal_IterNext
v_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_1))
_ = v_3_3
var __t29 bool
{
if (v_3_3 != nil) {
// TAST (Let): v2_4_4 -> *Constructor_Data_Map_Internal_IterNext
v2_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_2))
_ = v2_4_4
var __t28 bool
{
var __t_and_27 bool = false
if (v2_4_4 != nil) {

var __t26 bool
{
var __t_tag_5 gopurs_runtime.Value = (v_3_3).V0
if (uint32(__t_tag_5.IntVal) == 3908053364) {
var __t7 bool
{
var __t_tag_6 gopurs_runtime.Value = (v2_4_4).V0
if (uint32(__t_tag_6.IntVal) == 3908053364) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t26 = __t7
goto end_branch_26
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (v_3_3).V0
if (uint32(__t_tag_8.IntVal) == 217821258) {
var __t10 bool
{
var __t_tag_9 gopurs_runtime.Value = (v2_4_4).V0
if (uint32(__t_tag_9.IntVal) == 217821258) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t26 = __t10
goto end_branch_26
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (v_3_3).V0
if (uint32(__t_tag_11.IntVal) == 1292308612) {
var __t13 bool
{
var __t_tag_12 gopurs_runtime.Value = (v2_4_4).V0
if (uint32(__t_tag_12.IntVal) == 1292308612) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
__t26 = __t13
goto end_branch_26
} else {

}
}
{
var __t_tag_14 gopurs_runtime.Value = (v_3_3).V0
if (uint32(__t_tag_14.IntVal) == 2311060696) {
var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = (v2_4_4).V0
if (uint32(__t_tag_15.IntVal) == 2311060696) {
__t16 = true
goto end_branch_16
} else {

}
}
{
__t16 = false
}
end_branch_16:
__t26 = __t16
goto end_branch_26
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = (v_3_3).V0
if (uint32(__t_tag_17.IntVal) == 401302776) {
var __t19 bool
{
var __t_tag_18 gopurs_runtime.Value = (v2_4_4).V0
if (uint32(__t_tag_18.IntVal) == 401302776) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = false
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
var __t_tag_20 gopurs_runtime.Value = (v_3_3).V0
if (uint32(__t_tag_20.IntVal) == 3327533908) {
var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = (v2_4_4).V0
if (uint32(__t_tag_21.IntVal) == 3327533908) {
__t22 = true
goto end_branch_22
} else {

}
}
{
__t22 = false
}
end_branch_22:
__t26 = __t22
goto end_branch_26
} else {

}
}
{
var __t_tag_23 gopurs_runtime.Value = (v_3_3).V0
var __t_and_25 bool = false
if (uint32(__t_tag_23.IntVal) == 3631736139) {

var __t_tag_24 gopurs_runtime.Value = (v2_4_4).V0
__t_and_25 = (uint32(__t_tag_24.IntVal) == 3631736139)
}
if __t_and_25 {
__t26 = true
goto end_branch_26
} else {

}
}
{
__t26 = false
}
end_branch_26:
__t_and_27 = (__t26) && (((v_3_3).V1.FloatVal()) == ((v2_4_4).V1.FloatVal()))
}
if __t_and_27 {
a_1_loop = (v_3_3).V2
b_2_loop = (v2_4_4).V2
continue go__go_0_2_1
__t28 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_28
} else {

}
}
{
__t28 = false
}
end_branch_28:
__t29 = __t28
goto end_branch_29
} else {

}
}
{
if (v_3_3 == nil) {
__t29 = true
goto end_branch_29
} else {

}
}
{
__t29 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_29:
return gopurs_runtime.Bool(__t29)
}
}()
})
})
// TAST (Let): eqMapIter2_0_1 -> *Constructor_Data_Eq_Eq
eqMapIter2_0_1 := &Constructor_Data_Eq_Eq{1, go__go_0_2_1}
_ = eqMapIter2_0_1
var go__go_1_30_2 gopurs_runtime.Value
go__go_1_30_2 = gopurs_runtime.Func(func(a_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_2_loop gopurs_runtime.Value = a_2_loop_val
var b_3_loop gopurs_runtime.Value = b_3_loop_val
go__go_1_30_2:
for {
if false { continue go__go_1_30_2 }
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
// TAST (Let): v_4_31 -> *Constructor_Data_Map_Internal_IterNext
v_4_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_3))
_ = v_4_31
// TAST (Let): v1_5_32 -> *Constructor_Data_Map_Internal_IterNext
v1_5_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_2))
_ = v1_5_32
var __t67 uint32
{
if (v1_5_32 != nil) {
var __t65 uint32
{
if (v_4_31 != nil) {
var __t61 uint32
{
var __t_tag_34 gopurs_runtime.Value = (v1_5_32).V0
if (uint32(__t_tag_34.IntVal) == 3908053364) {
var __t36 uint32
{
var __t_tag_35 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_35.IntVal) == 3908053364) {
__t36 = 902936544
goto end_branch_36
} else {

}
}
{
__t36 = 1527465420
}
end_branch_36:
__t61 = __t36
goto end_branch_61
} else {

}
}
{
var __t_tag_37 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_37.IntVal) == 3908053364) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_38 gopurs_runtime.Value = (v1_5_32).V0
if (uint32(__t_tag_38.IntVal) == 217821258) {
var __t40 uint32
{
var __t_tag_39 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_39.IntVal) == 217821258) {
__t40 = 902936544
goto end_branch_40
} else {

}
}
{
__t40 = 1527465420
}
end_branch_40:
__t61 = __t40
goto end_branch_61
} else {

}
}
{
var __t_tag_41 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_41.IntVal) == 217821258) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_42 gopurs_runtime.Value = (v1_5_32).V0
if (uint32(__t_tag_42.IntVal) == 1292308612) {
var __t44 uint32
{
var __t_tag_43 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_43.IntVal) == 1292308612) {
__t44 = 902936544
goto end_branch_44
} else {

}
}
{
__t44 = 1527465420
}
end_branch_44:
__t61 = __t44
goto end_branch_61
} else {

}
}
{
var __t_tag_45 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_45.IntVal) == 1292308612) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_46 gopurs_runtime.Value = (v1_5_32).V0
if (uint32(__t_tag_46.IntVal) == 2311060696) {
var __t48 uint32
{
var __t_tag_47 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_47.IntVal) == 2311060696) {
__t48 = 902936544
goto end_branch_48
} else {

}
}
{
__t48 = 1527465420
}
end_branch_48:
__t61 = __t48
goto end_branch_61
} else {

}
}
{
var __t_tag_49 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_49.IntVal) == 2311060696) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_50 gopurs_runtime.Value = (v1_5_32).V0
if (uint32(__t_tag_50.IntVal) == 401302776) {
var __t52 uint32
{
var __t_tag_51 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_51.IntVal) == 401302776) {
__t52 = 902936544
goto end_branch_52
} else {

}
}
{
__t52 = 1527465420
}
end_branch_52:
__t61 = __t52
goto end_branch_61
} else {

}
}
{
var __t_tag_53 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_53.IntVal) == 401302776) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_54 gopurs_runtime.Value = (v1_5_32).V0
if (uint32(__t_tag_54.IntVal) == 3327533908) {
var __t56 uint32
{
var __t_tag_55 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_55.IntVal) == 3327533908) {
__t56 = 902936544
goto end_branch_56
} else {

}
}
{
__t56 = 1527465420
}
end_branch_56:
__t61 = __t56
goto end_branch_61
} else {

}
}
{
var __t_tag_57 gopurs_runtime.Value = (v_4_31).V0
if (uint32(__t_tag_57.IntVal) == 3327533908) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_58 gopurs_runtime.Value = (v1_5_32).V0
var __t_and_60 bool = false
if (uint32(__t_tag_58.IntVal) == 3631736139) {

var __t_tag_59 gopurs_runtime.Value = (v_4_31).V0
__t_and_60 = (uint32(__t_tag_59.IntVal) == 3631736139)
}
if __t_and_60 {
__t61 = 902936544
goto end_branch_61
} else {

}
}
{
__t61 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_61:
// TAST (Let): v3_6_33 -> uint32
v3_6_33 := __t61
_ = v3_6_33
var __t64 uint32
{
if (v3_6_33 == 902936544) {
// TAST (Let): v4_7_62 -> gopurs_runtime.Value
v4_7_62 := gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (v1_5_32).V1, (v_4_31).V1)
_ = v4_7_62
var __t63 uint32
{
if (uint32(v4_7_62.IntVal) == 902936544) {
a_2_loop = (v1_5_32).V2
b_3_loop = (v_4_31).V2
continue go__go_1_30_2
__t63 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_63
} else {

}
}
{
__t63 = uint32(v4_7_62.IntVal)
}
end_branch_63:
__t64 = __t63
goto end_branch_64
} else {

}
}
{
__t64 = v3_6_33
}
end_branch_64:
__t65 = __t64
goto end_branch_65
} else {

}
}
{
if (v_4_31 == nil) {
__t65 = 380165415
goto end_branch_65
} else {

}
}
{
__t65 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_65:
__t67 = __t65
goto end_branch_67
} else {

}
}
{
if (v1_5_32 == nil) {
var __t66 uint32
{
if (v_4_31 == nil) {
__t66 = 902936544
goto end_branch_66
} else {

}
}
{
__t66 = 1527465420
}
end_branch_66:
__t67 = __t66
goto end_branch_67
} else {

}
}
{
if (v_4_31 == nil) {
__t67 = 380165415
goto end_branch_67
} else {

}
}
{
__t67 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_67:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t67), UnsafePtr: nil}
}
}()
})
})
// TAST (Let): ordMapIter2_0_0 -> *Constructor_Data_Ord_Ord
ordMapIter2_0_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_0_1)}
}), go__go_1_30_2}
_ = ordMapIter2_0_0
var go__go_1_70_3 gopurs_runtime.Value
go__go_1_70_3 = gopurs_runtime.Func(func(a_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_2_loop gopurs_runtime.Value = a_2_loop_val
var b_3_loop gopurs_runtime.Value = b_3_loop_val
go__go_1_70_3:
for {
if false { continue go__go_1_70_3 }
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
// TAST (Let): v_4_71 -> *Constructor_Data_Map_Internal_IterNext
v_4_71 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_2))
_ = v_4_71
var __t97 bool
{
if (v_4_71 != nil) {
// TAST (Let): v2_5_72 -> *Constructor_Data_Map_Internal_IterNext
v2_5_72 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_3))
_ = v2_5_72
var __t96 bool
{
var __t_and_95 bool = false
if (v2_5_72 != nil) {

var __t94 bool
{
var __t_tag_73 gopurs_runtime.Value = (v_4_71).V0
if (uint32(__t_tag_73.IntVal) == 3908053364) {
var __t75 bool
{
var __t_tag_74 gopurs_runtime.Value = (v2_5_72).V0
if (uint32(__t_tag_74.IntVal) == 3908053364) {
__t75 = true
goto end_branch_75
} else {

}
}
{
__t75 = false
}
end_branch_75:
__t94 = __t75
goto end_branch_94
} else {

}
}
{
var __t_tag_76 gopurs_runtime.Value = (v_4_71).V0
if (uint32(__t_tag_76.IntVal) == 217821258) {
var __t78 bool
{
var __t_tag_77 gopurs_runtime.Value = (v2_5_72).V0
if (uint32(__t_tag_77.IntVal) == 217821258) {
__t78 = true
goto end_branch_78
} else {

}
}
{
__t78 = false
}
end_branch_78:
__t94 = __t78
goto end_branch_94
} else {

}
}
{
var __t_tag_79 gopurs_runtime.Value = (v_4_71).V0
if (uint32(__t_tag_79.IntVal) == 1292308612) {
var __t81 bool
{
var __t_tag_80 gopurs_runtime.Value = (v2_5_72).V0
if (uint32(__t_tag_80.IntVal) == 1292308612) {
__t81 = true
goto end_branch_81
} else {

}
}
{
__t81 = false
}
end_branch_81:
__t94 = __t81
goto end_branch_94
} else {

}
}
{
var __t_tag_82 gopurs_runtime.Value = (v_4_71).V0
if (uint32(__t_tag_82.IntVal) == 2311060696) {
var __t84 bool
{
var __t_tag_83 gopurs_runtime.Value = (v2_5_72).V0
if (uint32(__t_tag_83.IntVal) == 2311060696) {
__t84 = true
goto end_branch_84
} else {

}
}
{
__t84 = false
}
end_branch_84:
__t94 = __t84
goto end_branch_94
} else {

}
}
{
var __t_tag_85 gopurs_runtime.Value = (v_4_71).V0
if (uint32(__t_tag_85.IntVal) == 401302776) {
var __t87 bool
{
var __t_tag_86 gopurs_runtime.Value = (v2_5_72).V0
if (uint32(__t_tag_86.IntVal) == 401302776) {
__t87 = true
goto end_branch_87
} else {

}
}
{
__t87 = false
}
end_branch_87:
__t94 = __t87
goto end_branch_94
} else {

}
}
{
var __t_tag_88 gopurs_runtime.Value = (v_4_71).V0
if (uint32(__t_tag_88.IntVal) == 3327533908) {
var __t90 bool
{
var __t_tag_89 gopurs_runtime.Value = (v2_5_72).V0
if (uint32(__t_tag_89.IntVal) == 3327533908) {
__t90 = true
goto end_branch_90
} else {

}
}
{
__t90 = false
}
end_branch_90:
__t94 = __t90
goto end_branch_94
} else {

}
}
{
var __t_tag_91 gopurs_runtime.Value = (v_4_71).V0
var __t_and_93 bool = false
if (uint32(__t_tag_91.IntVal) == 3631736139) {

var __t_tag_92 gopurs_runtime.Value = (v2_5_72).V0
__t_and_93 = (uint32(__t_tag_92.IntVal) == 3631736139)
}
if __t_and_93 {
__t94 = true
goto end_branch_94
} else {

}
}
{
__t94 = false
}
end_branch_94:
__t_and_95 = (__t94) && (((v_4_71).V1.FloatVal()) == ((v2_5_72).V1.FloatVal()))
}
if __t_and_95 {
a_2_loop = (v_4_71).V2
b_3_loop = (v2_5_72).V2
continue go__go_1_70_3
__t96 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_96
} else {

}
}
{
__t96 = false
}
end_branch_96:
__t97 = __t96
goto end_branch_97
} else {

}
}
{
if (v_4_71 == nil) {
__t97 = true
goto end_branch_97
} else {

}
}
{
__t97 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_97:
return gopurs_runtime.Bool(__t97)
}
}()
})
})
// TAST (Let): eqMapIter2_1_69 -> *Constructor_Data_Eq_Eq
eqMapIter2_1_69 := &Constructor_Data_Eq_Eq{1, go__go_1_70_3}
_ = eqMapIter2_1_69
// TAST (Let): eqMap2_1_68 -> *Constructor_Data_Eq_Eq
eqMap2_1_68 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t100 bool
{
if (xs_2.Type == 9 && xs_2.IntVal == 324739070 && xs_2.UnsafePtr == nil) {
var __t98 bool
{
if (ys_3.Type == 9 && ys_3.IntVal == 324739070 && ys_3.UnsafePtr == nil) {
__t98 = true
goto end_branch_98
} else {

}
}
{
__t98 = false
}
end_branch_98:
__t100 = __t98
goto end_branch_100
} else {

}
}
{
if (xs_2.Type == 9 && xs_2.IntVal == 324739070 && xs_2.UnsafePtr != nil) {
var __t99 bool
{
if ((ys_3.Type == 9 && ys_3.IntVal == 324739070 && ys_3.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_2.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_3.UnsafePtr).V1)) {
__t99 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_1_69.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_99
} else {

}
}
{
__t99 = false
}
end_branch_99:
__t100 = __t99
goto end_branch_100
} else {

}
}
{
__t100 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_100:
return gopurs_runtime.Bool(__t100)
})
})}
_ = eqMap2_1_68
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMap2_1_68)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t103 uint32
{
if (xs_2.Type == 9 && xs_2.IntVal == 324739070 && xs_2.UnsafePtr == nil) {
var __t102 uint32
{
if (ys_3.Type == 9 && ys_3.IntVal == 324739070 && ys_3.UnsafePtr == nil) {
__t102 = 902936544
goto end_branch_102
} else {

}
}
{
__t102 = 1527465420
}
end_branch_102:
__t103 = __t102
goto end_branch_103
} else {

}
}
{
var __t101 uint32
{
if (ys_3.Type == 9 && ys_3.IntVal == 324739070 && ys_3.UnsafePtr == nil) {
__t101 = 380165415
goto end_branch_101
} else {

}
}
{
__t101 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_0_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
}
end_branch_101:
__t103 = __t101
}
end_branch_103:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t103), UnsafePtr: nil}
})
})})}
}()
	})
	return cache_Data_Interval_Duration_ordMap
}

var cache_Data_Interval_Duration_semigroupDuration gopurs_runtime.Value
var once_Data_Interval_Duration_semigroupDuration sync.Once
func Get_Data_Interval_Duration_semigroupDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_semigroupDuration.Do(func() {
		cache_Data_Interval_Duration_semigroupDuration = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
{
if (uint32(x_2.IntVal) == 3908053364) {
var __t0 uint32
{
if (uint32(y_3.IntVal) == 3908053364) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (uint32(y_3.IntVal) == 3908053364) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_2.IntVal) == 217821258) {
var __t1 uint32
{
if (uint32(y_3.IntVal) == 217821258) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (uint32(y_3.IntVal) == 217821258) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_2.IntVal) == 1292308612) {
var __t2 uint32
{
if (uint32(y_3.IntVal) == 1292308612) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (uint32(y_3.IntVal) == 1292308612) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_2.IntVal) == 2311060696) {
var __t3 uint32
{
if (uint32(y_3.IntVal) == 2311060696) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (uint32(y_3.IntVal) == 2311060696) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_2.IntVal) == 401302776) {
var __t4 uint32
{
if (uint32(y_3.IntVal) == 401302776) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (uint32(y_3.IntVal) == 401302776) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if (uint32(x_2.IntVal) == 3327533908) {
var __t5 uint32
{
if (uint32(y_3.IntVal) == 3327533908) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (uint32(y_3.IntVal) == 3327533908) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if ((uint32(x_2.IntVal) == 3631736139)) && ((uint32(y_3.IntVal) == 3631736139)) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
})
}), Get_Data_Semiring_numAdd(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_1))})))}
})
})})}
	})
	return cache_Data_Interval_Duration_semigroupDuration
}

var cache_Data_Interval_Duration_monoidDuration gopurs_runtime.Value
var once_Data_Interval_Duration_monoidDuration sync.Once
func Get_Data_Interval_Duration_monoidDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_monoidDuration.Do(func() {
		cache_Data_Interval_Duration_monoidDuration = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Interval_Duration_semigroupDuration()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}})}
	})
	return cache_Data_Interval_Duration_monoidDuration
}

var cache_Data_Interval_Duration_eqDuration gopurs_runtime.Value
var once_Data_Interval_Duration_eqDuration sync.Once
func Get_Data_Interval_Duration_eqDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDuration.Do(func() {
		cache_Data_Interval_Duration_eqDuration = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_4 gopurs_runtime.Value
go__go_2_1_4 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_1_4:
for {
if false { continue go__go_2_1_4 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_2 -> *Constructor_Data_Map_Internal_IterNext
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_2
var __t28 bool
{
if (v_5_2 != nil) {
// TAST (Let): v2_6_3 -> *Constructor_Data_Map_Internal_IterNext
v2_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_3
var __t27 bool
{
var __t_and_26 bool = false
if (v2_6_3 != nil) {

var __t25 bool
{
var __t_tag_4 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_4.IntVal) == 3908053364) {
var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_5.IntVal) == 3908053364) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t25 = __t6
goto end_branch_25
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_7.IntVal) == 217821258) {
var __t9 bool
{
var __t_tag_8 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_8.IntVal) == 217821258) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
__t25 = __t9
goto end_branch_25
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_10.IntVal) == 1292308612) {
var __t12 bool
{
var __t_tag_11 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_11.IntVal) == 1292308612) {
__t12 = true
goto end_branch_12
} else {

}
}
{
__t12 = false
}
end_branch_12:
__t25 = __t12
goto end_branch_25
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_13.IntVal) == 2311060696) {
var __t15 bool
{
var __t_tag_14 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_14.IntVal) == 2311060696) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t25 = __t15
goto end_branch_25
} else {

}
}
{
var __t_tag_16 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_16.IntVal) == 401302776) {
var __t18 bool
{
var __t_tag_17 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_17.IntVal) == 401302776) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
__t25 = __t18
goto end_branch_25
} else {

}
}
{
var __t_tag_19 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_19.IntVal) == 3327533908) {
var __t21 bool
{
var __t_tag_20 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_20.IntVal) == 3327533908) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
__t25 = __t21
goto end_branch_25
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = (v_5_2).V0
var __t_and_24 bool = false
if (uint32(__t_tag_22.IntVal) == 3631736139) {

var __t_tag_23 gopurs_runtime.Value = (v2_6_3).V0
__t_and_24 = (uint32(__t_tag_23.IntVal) == 3631736139)
}
if __t_and_24 {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = (__t25) && (((v_5_2).V1.FloatVal()) == ((v2_6_3).V1.FloatVal()))
}
if __t_and_26 {
a_3_loop = (v_5_2).V2
b_4_loop = (v2_6_3).V2
continue go__go_2_1_4
__t27 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v_5_2 == nil) {
__t28 = true
goto end_branch_28
} else {

}
}
{
__t28 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_28:
return gopurs_runtime.Bool(__t28)
}
}()
})
})
// TAST (Let): eqMapIter2_2_0 -> *Constructor_Data_Eq_Eq
eqMapIter2_2_0 := &Constructor_Data_Eq_Eq{1, go__go_2_1_4}
_ = eqMapIter2_2_0
var __t35 bool
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0)
if (__t_tag_29 == nil) {
var __t31 bool
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if (__t_tag_30 == nil) {
__t31 = true
goto end_branch_31
} else {

}
}
{
__t31 = false
}
end_branch_31:
__t35 = __t31
goto end_branch_35
} else {

}
}
{
var __t_tag_32 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0)
if (__t_tag_32 != nil) {
var __t34 bool
{
var __t_tag_33 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if ((__t_tag_33 != nil)) && (((*Constructor_Data_Map_Internal_Node)(x_0.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(y_1.UnsafePtr).V1)) {
__t34 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_34
} else {

}
}
{
__t34 = false
}
end_branch_34:
__t35 = __t34
goto end_branch_35
} else {

}
}
{
__t35 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_35:
return gopurs_runtime.Bool(__t35)
})
})})}
	})
	return cache_Data_Interval_Duration_eqDuration
}

var cache_Data_Interval_Duration_ordDuration gopurs_runtime.Value
var once_Data_Interval_Duration_ordDuration sync.Once
func Get_Data_Interval_Duration_ordDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDuration.Do(func() {
		cache_Data_Interval_Duration_ordDuration = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_eqDuration()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_2_5 gopurs_runtime.Value
go__go_2_2_5 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_2_5:
for {
if false { continue go__go_2_2_5 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_3 -> *Constructor_Data_Map_Internal_IterNext
v_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_3
var __t29 bool
{
if (v_5_3 != nil) {
// TAST (Let): v2_6_4 -> *Constructor_Data_Map_Internal_IterNext
v2_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_4
var __t28 bool
{
var __t_and_27 bool = false
if (v2_6_4 != nil) {

var __t26 bool
{
var __t_tag_5 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_5.IntVal) == 3908053364) {
var __t7 bool
{
var __t_tag_6 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_6.IntVal) == 3908053364) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t26 = __t7
goto end_branch_26
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_8.IntVal) == 217821258) {
var __t10 bool
{
var __t_tag_9 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_9.IntVal) == 217821258) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t26 = __t10
goto end_branch_26
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_11.IntVal) == 1292308612) {
var __t13 bool
{
var __t_tag_12 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_12.IntVal) == 1292308612) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
__t26 = __t13
goto end_branch_26
} else {

}
}
{
var __t_tag_14 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_14.IntVal) == 2311060696) {
var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_15.IntVal) == 2311060696) {
__t16 = true
goto end_branch_16
} else {

}
}
{
__t16 = false
}
end_branch_16:
__t26 = __t16
goto end_branch_26
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_17.IntVal) == 401302776) {
var __t19 bool
{
var __t_tag_18 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_18.IntVal) == 401302776) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = false
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
var __t_tag_20 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_20.IntVal) == 3327533908) {
var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_21.IntVal) == 3327533908) {
__t22 = true
goto end_branch_22
} else {

}
}
{
__t22 = false
}
end_branch_22:
__t26 = __t22
goto end_branch_26
} else {

}
}
{
var __t_tag_23 gopurs_runtime.Value = (v_5_3).V0
var __t_and_25 bool = false
if (uint32(__t_tag_23.IntVal) == 3631736139) {

var __t_tag_24 gopurs_runtime.Value = (v2_6_4).V0
__t_and_25 = (uint32(__t_tag_24.IntVal) == 3631736139)
}
if __t_and_25 {
__t26 = true
goto end_branch_26
} else {

}
}
{
__t26 = false
}
end_branch_26:
__t_and_27 = (__t26) && (((v_5_3).V1.FloatVal()) == ((v2_6_4).V1.FloatVal()))
}
if __t_and_27 {
a_3_loop = (v_5_3).V2
b_4_loop = (v2_6_4).V2
continue go__go_2_2_5
__t28 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_28
} else {

}
}
{
__t28 = false
}
end_branch_28:
__t29 = __t28
goto end_branch_29
} else {

}
}
{
if (v_5_3 == nil) {
__t29 = true
goto end_branch_29
} else {

}
}
{
__t29 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_29:
return gopurs_runtime.Bool(__t29)
}
}()
})
})
// TAST (Let): eqMapIter2_2_1 -> *Constructor_Data_Eq_Eq
eqMapIter2_2_1 := &Constructor_Data_Eq_Eq{1, go__go_2_2_5}
_ = eqMapIter2_2_1
var go__go_3_30_6 gopurs_runtime.Value
go__go_3_30_6 = gopurs_runtime.Func(func(a_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_4_loop gopurs_runtime.Value = a_4_loop_val
var b_5_loop gopurs_runtime.Value = b_5_loop_val
go__go_3_30_6:
for {
if false { continue go__go_3_30_6 }
var a_4 gopurs_runtime.Value = a_4_loop
_ = a_4
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
// TAST (Let): v_6_31 -> *Constructor_Data_Map_Internal_IterNext
v_6_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v_6_31
// TAST (Let): v1_7_32 -> *Constructor_Data_Map_Internal_IterNext
v1_7_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v1_7_32
var __t67 uint32
{
if (v1_7_32 != nil) {
var __t65 uint32
{
if (v_6_31 != nil) {
var __t61 uint32
{
var __t_tag_34 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_34.IntVal) == 3908053364) {
var __t36 uint32
{
var __t_tag_35 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_35.IntVal) == 3908053364) {
__t36 = 902936544
goto end_branch_36
} else {

}
}
{
__t36 = 1527465420
}
end_branch_36:
__t61 = __t36
goto end_branch_61
} else {

}
}
{
var __t_tag_37 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_37.IntVal) == 3908053364) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_38 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_38.IntVal) == 217821258) {
var __t40 uint32
{
var __t_tag_39 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_39.IntVal) == 217821258) {
__t40 = 902936544
goto end_branch_40
} else {

}
}
{
__t40 = 1527465420
}
end_branch_40:
__t61 = __t40
goto end_branch_61
} else {

}
}
{
var __t_tag_41 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_41.IntVal) == 217821258) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_42 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_42.IntVal) == 1292308612) {
var __t44 uint32
{
var __t_tag_43 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_43.IntVal) == 1292308612) {
__t44 = 902936544
goto end_branch_44
} else {

}
}
{
__t44 = 1527465420
}
end_branch_44:
__t61 = __t44
goto end_branch_61
} else {

}
}
{
var __t_tag_45 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_45.IntVal) == 1292308612) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_46 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_46.IntVal) == 2311060696) {
var __t48 uint32
{
var __t_tag_47 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_47.IntVal) == 2311060696) {
__t48 = 902936544
goto end_branch_48
} else {

}
}
{
__t48 = 1527465420
}
end_branch_48:
__t61 = __t48
goto end_branch_61
} else {

}
}
{
var __t_tag_49 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_49.IntVal) == 2311060696) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_50 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_50.IntVal) == 401302776) {
var __t52 uint32
{
var __t_tag_51 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_51.IntVal) == 401302776) {
__t52 = 902936544
goto end_branch_52
} else {

}
}
{
__t52 = 1527465420
}
end_branch_52:
__t61 = __t52
goto end_branch_61
} else {

}
}
{
var __t_tag_53 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_53.IntVal) == 401302776) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_54 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_54.IntVal) == 3327533908) {
var __t56 uint32
{
var __t_tag_55 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_55.IntVal) == 3327533908) {
__t56 = 902936544
goto end_branch_56
} else {

}
}
{
__t56 = 1527465420
}
end_branch_56:
__t61 = __t56
goto end_branch_61
} else {

}
}
{
var __t_tag_57 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_57.IntVal) == 3327533908) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_58 gopurs_runtime.Value = (v1_7_32).V0
var __t_and_60 bool = false
if (uint32(__t_tag_58.IntVal) == 3631736139) {

var __t_tag_59 gopurs_runtime.Value = (v_6_31).V0
__t_and_60 = (uint32(__t_tag_59.IntVal) == 3631736139)
}
if __t_and_60 {
__t61 = 902936544
goto end_branch_61
} else {

}
}
{
__t61 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_61:
// TAST (Let): v3_8_33 -> uint32
v3_8_33 := __t61
_ = v3_8_33
var __t64 uint32
{
if (v3_8_33 == 902936544) {
// TAST (Let): v4_9_62 -> gopurs_runtime.Value
v4_9_62 := gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (v1_7_32).V1, (v_6_31).V1)
_ = v4_9_62
var __t63 uint32
{
if (uint32(v4_9_62.IntVal) == 902936544) {
a_4_loop = (v1_7_32).V2
b_5_loop = (v_6_31).V2
continue go__go_3_30_6
__t63 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_63
} else {

}
}
{
__t63 = uint32(v4_9_62.IntVal)
}
end_branch_63:
__t64 = __t63
goto end_branch_64
} else {

}
}
{
__t64 = v3_8_33
}
end_branch_64:
__t65 = __t64
goto end_branch_65
} else {

}
}
{
if (v_6_31 == nil) {
__t65 = 380165415
goto end_branch_65
} else {

}
}
{
__t65 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_65:
__t67 = __t65
goto end_branch_67
} else {

}
}
{
if (v1_7_32 == nil) {
var __t66 uint32
{
if (v_6_31 == nil) {
__t66 = 902936544
goto end_branch_66
} else {

}
}
{
__t66 = 1527465420
}
end_branch_66:
__t67 = __t66
goto end_branch_67
} else {

}
}
{
if (v_6_31 == nil) {
__t67 = 380165415
goto end_branch_67
} else {

}
}
{
__t67 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_67:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t67), UnsafePtr: nil}
}
}()
})
})
// TAST (Let): ordMapIter2_2_0 -> *Constructor_Data_Ord_Ord
ordMapIter2_2_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_2_1)}
}), go__go_3_30_6}
_ = ordMapIter2_2_0
var __t73 uint32
{
var __t_tag_70 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0)
if (__t_tag_70 == nil) {
var __t72 uint32
{
var __t_tag_71 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if (__t_tag_71 == nil) {
__t72 = 902936544
goto end_branch_72
} else {

}
}
{
__t72 = 1527465420
}
end_branch_72:
__t73 = __t72
goto end_branch_73
} else {

}
}
{
var __t69 uint32
{
var __t_tag_68 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if (__t_tag_68 == nil) {
__t69 = 380165415
goto end_branch_69
} else {

}
}
{
__t69 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_2_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
}
end_branch_69:
__t73 = __t69
}
end_branch_73:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t73), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Interval_Duration_ordDuration
}

var cache_Data_Interval_Duration_durationFromComponent gopurs_runtime.Value
var once_Data_Interval_Duration_durationFromComponent sync.Once
func Get_Data_Interval_Duration_durationFromComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_durationFromComponent.Do(func() {
		cache_Data_Interval_Duration_durationFromComponent = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_durationFromComponent(uint32(k_0_box.IntVal), v_1_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_durationFromComponent
}

var cache_Data_Interval_Duration_hour gopurs_runtime.Value
var once_Data_Interval_Duration_hour sync.Once
func Get_Data_Interval_Duration_hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_hour.Do(func() {
		cache_Data_Interval_Duration_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_hour(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_hour
}

var cache_Data_Interval_Duration_millisecond gopurs_runtime.Value
var once_Data_Interval_Duration_millisecond sync.Once
func Get_Data_Interval_Duration_millisecond() gopurs_runtime.Value {
	once_Data_Interval_Duration_millisecond.Do(func() {
		cache_Data_Interval_Duration_millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_millisecond(x_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_millisecond
}

var cache_Data_Interval_Duration_minute gopurs_runtime.Value
var once_Data_Interval_Duration_minute sync.Once
func Get_Data_Interval_Duration_minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_minute.Do(func() {
		cache_Data_Interval_Duration_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_minute(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_minute
}

var cache_Data_Interval_Duration_month gopurs_runtime.Value
var once_Data_Interval_Duration_month sync.Once
func Get_Data_Interval_Duration_month() gopurs_runtime.Value {
	once_Data_Interval_Duration_month.Do(func() {
		cache_Data_Interval_Duration_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_month(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_month
}

var cache_Data_Interval_Duration_second gopurs_runtime.Value
var once_Data_Interval_Duration_second sync.Once
func Get_Data_Interval_Duration_second() gopurs_runtime.Value {
	once_Data_Interval_Duration_second.Do(func() {
		cache_Data_Interval_Duration_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_second(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_second
}

var cache_Data_Interval_Duration_week gopurs_runtime.Value
var once_Data_Interval_Duration_week sync.Once
func Get_Data_Interval_Duration_week() gopurs_runtime.Value {
	once_Data_Interval_Duration_week.Do(func() {
		cache_Data_Interval_Duration_week = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_week(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_week
}

var cache_Data_Interval_Duration_year gopurs_runtime.Value
var once_Data_Interval_Duration_year sync.Once
func Get_Data_Interval_Duration_year() gopurs_runtime.Value {
	once_Data_Interval_Duration_year.Do(func() {
		cache_Data_Interval_Duration_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_year(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_year
}

var cache_Data_Interval_Duration_day gopurs_runtime.Value
var once_Data_Interval_Duration_day sync.Once
func Get_Data_Interval_Duration_day() gopurs_runtime.Value {
	once_Data_Interval_Duration_day.Do(func() {
		cache_Data_Interval_Duration_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_day(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_day
}

type Constructor_Data_Interval_Duration_Second struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Minute struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Hour struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Day struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Week struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Month struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Year struct {
	Rc uint32
}


func Call_Data_Interval_Duration_Duration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Interval_Duration_durationFromComponent(k_0_loop uint32, v_1_loop float64) *Constructor_Data_Map_Internal_Node {
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Float(v_1), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_hour(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_millisecond(x_0_loop float64) *Constructor_Data_Map_Internal_Node {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Float((gopurs_runtime.Float(x_0).FloatVal()) / (1000.0)).FloatVal()), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_minute(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_month(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_second(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_week(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_year(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}

func Call_Data_Interval_Duration_day(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}, gopurs_runtime.Float(v_0), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})})
}


