package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Rec_Class_Loop gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Loop sync.Once
func Get_Control_Monad_Rec_Class_Loop() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Loop.Do(func() {
		cache_Control_Monad_Rec_Class_Loop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, value0})}
})
	})
	return cache_Control_Monad_Rec_Class_Loop
}

var cache_Control_Monad_Rec_Class_Done gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Done sync.Once
func Get_Control_Monad_Rec_Class_Done() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Done.Do(func() {
		cache_Control_Monad_Rec_Class_Done = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, value0})}
})
	})
	return cache_Control_Monad_Rec_Class_Done
}

var cache_Control_Monad_Rec_Class_MonadRec_dollarDict gopurs_runtime.Value
var once_Control_Monad_Rec_Class_MonadRec_dollarDict sync.Once
func Get_Control_Monad_Rec_Class_MonadRec_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_MonadRec_dollarDict.Do(func() {
		cache_Control_Monad_Rec_Class_MonadRec_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_MonadRec_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_MonadRec_dollarDict
}

var cache_Control_Monad_Rec_Class_tailRecM gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM sync.Once
func Get_Control_Monad_Rec_Class_tailRecM() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM
}

var cache_Control_Monad_Rec_Class_tailRecM2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2
}

var cache_Control_Monad_Rec_Class_tailRecM3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM3 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM3.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM3 = gopurs_runtime.Func5(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM3(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box, c_4_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM3
}

var cache_Control_Monad_Rec_Class_untilJust gopurs_runtime.Value
var once_Control_Monad_Rec_Class_untilJust sync.Once
func Get_Control_Monad_Rec_Class_untilJust() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_untilJust.Do(func() {
		cache_Control_Monad_Rec_Class_untilJust = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_untilJust(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_untilJust
}

var cache_Control_Monad_Rec_Class_whileJust gopurs_runtime.Value
var once_Control_Monad_Rec_Class_whileJust sync.Once
func Get_Control_Monad_Rec_Class_whileJust() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_whileJust.Do(func() {
		cache_Control_Monad_Rec_Class_whileJust = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_whileJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_whileJust
}

var cache_Control_Monad_Rec_Class_tailRec gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec sync.Once
func Get_Control_Monad_Rec_Class_tailRec() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec(f_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec
}

var cache_Control_Monad_Rec_Class_tailRec2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec2 sync.Once
func Get_Control_Monad_Rec_Class_tailRec2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec2.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec2
}

var cache_Control_Monad_Rec_Class_tailRec3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec3 sync.Once
func Get_Control_Monad_Rec_Class_tailRec3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec3.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec3 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, c_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec3(f_0_box, a_1_box, b_2_box, c_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec3
}

var cache_Control_Monad_Rec_Class_monadRecMaybe gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecMaybe sync.Once
func Get_Control_Monad_Rec_Class_monadRecMaybe() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecMaybe.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Data_Maybe_monadMaybe()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Control_Monad_Rec_Class_Done
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t4 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0.UnsafePtr).V0)))}})}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t3)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t4)}
})
_ = __local_var_2_0
var go__go_3_5_3 gopurs_runtime.Value
go__go_3_5_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_3:
for {
if false { continue go__go_3_5_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0)
continue go__go_3_5_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*Constructor_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}()
})
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_3_5_3, gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, a0_1)))}))))}
})
})})}
	})
	return cache_Control_Monad_Rec_Class_monadRecMaybe
}

var cache_Control_Monad_Rec_Class_monadRecIdentity gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecIdentity sync.Once
func Get_Control_Monad_Rec_Class_monadRecIdentity() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecIdentity.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Data_Identity_monadIdentity()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_1_4 gopurs_runtime.Value
go__go_1_1_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_1_4:
for {
if false { continue go__go_1_1_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__go_1_1_4
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t2 = (*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_1_4, gopurs_runtime.Apply(f_0, x_2))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
})})}
	})
	return cache_Control_Monad_Rec_Class_monadRecIdentity
}

var cache_Control_Monad_Rec_Class_monadRecFunction gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecFunction sync.Once
func Get_Control_Monad_Rec_Class_monadRecFunction() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecFunction.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecFunction = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Control_Monad_monadFn()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_5 gopurs_runtime.Value
go__go_3_0_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_5:
for {
if false { continue go__go_3_0_5 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0, e_2)
continue go__go_3_0_5
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Apply(go__go_3_0_5, gopurs_runtime.Apply2(f_0, a0_1, e_2))
})
})
})})}
	})
	return cache_Control_Monad_Rec_Class_monadRecFunction
}

var cache_Control_Monad_Rec_Class_monadRecEither gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecEither sync.Once
func Get_Control_Monad_Rec_Class_monadRecEither() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecEither.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecEither = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Data_Either_monadEither()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Control_Monad_Rec_Class_Done
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t4 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v_2.UnsafePtr).V0})}}
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t3)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t4)}
})
_ = __local_var_2_0
var go__go_3_5_6 gopurs_runtime.Value
go__go_3_5_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_6:
for {
if false { continue go__go_3_5_6 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0)
continue go__go_3_5_6
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*Constructor_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}()
})
return gopurs_runtime.Apply(go__go_3_5_6, gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))
})
})})}
	})
	return cache_Control_Monad_Rec_Class_monadRecEither
}

var cache_Control_Monad_Rec_Class_monadRecEffect gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecEffect sync.Once
func Get_Control_Monad_Rec_Class_monadRecEffect() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecEffect.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecEffect = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Effect_monadEffect()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): fromDone_2_0 -> gopurs_runtime.Value
fromDone_2_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_3.UnsafePtr).V0
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
}))
_ = fromDone_2_0
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(f_0, a_1)
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{})
_ = __local_var_4_4
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), __local_var_4_4), gopurs_runtime.Value{})
})
_ = __local_var_3_2
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_5 := gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): __local_var_5_7 -> gopurs_runtime.Value
__local_var_5_7 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_5)
_ = __local_var_5_7
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_8 := gopurs_runtime.Apply(__local_var_5_7, gopurs_runtime.Value{})
_ = __local_var_6_8
var __t11 gopurs_runtime.Value
{
if (__local_var_6_8.Type == 9 && __local_var_6_8.IntVal == 525585346) {
__t11 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(__local_var_6_8.UnsafePtr).V0), gopurs_runtime.Value{})
_ = __local_var_7_9
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), __local_var_7_9, __local_var_4_5), gopurs_runtime.Value{})
_ = __local_var_8_10
return gopurs_runtime.Bool(false)
})
goto end_branch_11
} else {

}
}
{
if (__local_var_6_8.Type == 9 && __local_var_6_8.IntVal == 60402430) {
__t11 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Apply(__t11, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
_ = __local_var_5_6
__local_var_6_12 := fromDone_2_0
_ = __local_var_6_12
__local_var_7_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_5), gopurs_runtime.Value{})
_ = __local_var_7_13
return gopurs_runtime.Apply(__local_var_6_12, __local_var_7_13)
})
})
})})}
	})
	return cache_Control_Monad_Rec_Class_monadRecEffect
}

var cache_Control_Monad_Rec_Class_loop3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_loop3 sync.Once
func Get_Control_Monad_Rec_Class_loop3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_loop3.Do(func() {
		cache_Control_Monad_Rec_Class_loop3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_loop3(a_0_box, b_1_box, c_2_box)
})
	})
	return cache_Control_Monad_Rec_Class_loop3
}

var cache_Control_Monad_Rec_Class_loop2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_loop2 sync.Once
func Get_Control_Monad_Rec_Class_loop2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_loop2.Do(func() {
		cache_Control_Monad_Rec_Class_loop2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_loop2(a_0_box, b_1_box)
})
	})
	return cache_Control_Monad_Rec_Class_loop2
}

var cache_Control_Monad_Rec_Class_functorStep gopurs_runtime.Value
var once_Control_Monad_Rec_Class_functorStep sync.Once
func Get_Control_Monad_Rec_Class_functorStep() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_functorStep.Do(func() {
		cache_Control_Monad_Rec_Class_functorStep = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, (*Constructor_Control_Monad_Rec_Class_Loop)(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Done)(m_1.UnsafePtr).V0)})}
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
})})}
	})
	return cache_Control_Monad_Rec_Class_functorStep
}

var cache_Control_Monad_Rec_Class_forever gopurs_runtime.Value
var once_Control_Monad_Rec_Class_forever sync.Once
func Get_Control_Monad_Rec_Class_forever() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_forever.Do(func() {
		cache_Control_Monad_Rec_Class_forever = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_forever(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_forever
}

var cache_Control_Monad_Rec_Class_bifunctorStep gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bifunctorStep sync.Once
func Get_Control_Monad_Rec_Class_bifunctorStep() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bifunctorStep.Do(func() {
		cache_Control_Monad_Rec_Class_bifunctorStep = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Apply(v_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Apply(v1_1, (*Constructor_Control_Monad_Rec_Class_Done)(v2_2.UnsafePtr).V0)})}
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
})})}
	})
	return cache_Control_Monad_Rec_Class_bifunctorStep
}

var cache_Control_Monad_Rec_Class_tailRec__2110844386 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec__2110844386 sync.Once
func Get_Control_Monad_Rec_Class_tailRec__2110844386() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec__2110844386.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec__2110844386 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec__2110844386(f_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec__2110844386
}

var cache_Control_Monad_Rec_Class_tailRec__2334182452 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec__2334182452 sync.Once
func Get_Control_Monad_Rec_Class_tailRec__2334182452() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec__2334182452.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec__2334182452 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec__2334182452(f_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec__2334182452
}

var cache_Control_Monad_Rec_Class_tailRec__2666749533 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec__2666749533 sync.Once
func Get_Control_Monad_Rec_Class_tailRec__2666749533() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec__2666749533.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec__2666749533 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec__2666749533(f_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec__2666749533
}

var cache_Control_Monad_Rec_Class_tailRec__2045907654 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec__2045907654 sync.Once
func Get_Control_Monad_Rec_Class_tailRec__2045907654() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec__2045907654.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec__2045907654 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec__2045907654(f_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec__2045907654
}

var cache_Control_Monad_Rec_Class_tailRec__2929877587 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec__2929877587 sync.Once
func Get_Control_Monad_Rec_Class_tailRec__2929877587() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec__2929877587.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec__2929877587 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRec__2929877587(f_0_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRec__2929877587
}

var cache_Control_Monad_Rec_Class_tailRecM__2468783173 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2468783173 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2468783173() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2468783173.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2468783173 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2468783173(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2468783173
}

var cache_Control_Monad_Rec_Class_tailRecM__2333741125 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2333741125 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2333741125() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2333741125.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2333741125 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2333741125(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2333741125
}

var cache_Control_Monad_Rec_Class_tailRecM__3369162518 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__3369162518 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__3369162518() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__3369162518.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__3369162518 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__3369162518(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__3369162518
}

var cache_Control_Monad_Rec_Class_tailRecM__2016576534 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2016576534 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2016576534() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2016576534.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2016576534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2016576534(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2016576534
}

var cache_Control_Monad_Rec_Class_tailRecM__2220253896 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2220253896 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2220253896() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2220253896.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2220253896 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2220253896(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2220253896
}

var cache_Control_Monad_Rec_Class_tailRecM__2301013304 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2301013304 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2301013304() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2301013304.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2301013304 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2301013304(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2301013304
}

var cache_Control_Monad_Rec_Class_tailRecM__1121640472 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__1121640472 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__1121640472() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__1121640472.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__1121640472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__1121640472(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__1121640472
}

var cache_Control_Monad_Rec_Class_tailRecM__3865988408 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__3865988408 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__3865988408() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__3865988408.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__3865988408
}

var cache_Control_Monad_Rec_Class_tailRecM__1299680280 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__1299680280 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__1299680280() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__1299680280.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__1299680280 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__1299680280(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__1299680280
}

var cache_Control_Monad_Rec_Class_tailRecM__2119835928 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2119835928 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2119835928() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2119835928.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2119835928 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2119835928(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2119835928
}

var cache_Control_Monad_Rec_Class_tailRecM__1700625784 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__1700625784 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__1700625784() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__1700625784.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__1700625784 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__1700625784(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__1700625784
}

var cache_Control_Monad_Rec_Class_tailRecM__2172295544 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2172295544 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2172295544() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2172295544.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2172295544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2172295544(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2172295544
}

var cache_Control_Monad_Rec_Class_tailRecM__3782365624 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__3782365624 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__3782365624() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__3782365624.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__3782365624 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__3782365624(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__3782365624
}

var cache_Control_Monad_Rec_Class_tailRecM__1818438580 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__1818438580 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__1818438580() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__1818438580.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__1818438580 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__1818438580(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__1818438580
}

var cache_Control_Monad_Rec_Class_tailRecM__3515372237 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__3515372237 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__3515372237() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__3515372237.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__3515372237 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__3515372237(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__3515372237
}

var cache_Control_Monad_Rec_Class_tailRecM__3965932080 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__3965932080 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__3965932080() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__3965932080.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__3965932080 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__3965932080(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__3965932080
}

var cache_Control_Monad_Rec_Class_tailRecM__2810501104 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2810501104 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2810501104() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2810501104.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2810501104 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2810501104(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2810501104
}

var cache_Control_Monad_Rec_Class_tailRecM__3478800400 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__3478800400 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__3478800400() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__3478800400.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__3478800400 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__3478800400(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__3478800400
}

var cache_Control_Monad_Rec_Class_tailRecM__2063592441 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2063592441 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2063592441() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2063592441.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2063592441 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2063592441(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2063592441
}

var cache_Control_Monad_Rec_Class_tailRecM__1444729948 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__1444729948 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__1444729948() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__1444729948.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__1444729948 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__1444729948(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__1444729948
}

var cache_Control_Monad_Rec_Class_tailRecM__2222286441 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM__2222286441 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM__2222286441() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM__2222286441.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM__2222286441 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM__2222286441(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM__2222286441
}

var cache_Control_Monad_Rec_Class_tailRecM2__3241310373 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__3241310373 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__3241310373() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__3241310373.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__3241310373 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2__3241310373(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box.IntVal)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__3241310373
}

var cache_Control_Monad_Rec_Class_tailRecM2__2010968624 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__2010968624 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__2010968624() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__2010968624.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__2010968624 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2__2010968624(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__2010968624
}

var cache_Control_Monad_Rec_Class_tailRecM2__1800958704 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__1800958704 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__1800958704() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__1800958704.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__1800958704 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2__1800958704(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__1800958704
}

var cache_Control_Monad_Rec_Class_tailRecM2__1943630176 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__1943630176 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__1943630176() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__1943630176.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__1943630176 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2__1943630176(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__1943630176
}

var cache_Control_Monad_Rec_Class_tailRecM2__2551820843 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__2551820843 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__2551820843() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__2551820843.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__2551820843 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2__2551820843(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__2551820843
}

var cache_Control_Monad_Rec_Class_tailRecM2__3864450856 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__3864450856 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__3864450856() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__3864450856.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__3864450856 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Rec_Class_tailRecM2__3864450856(f_0_box, a_1_box, b_2_box))}
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__3864450856
}

var cache_Control_Monad_Rec_Class_tailRecM2__1136195496 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2__1136195496 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2__1136195496() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2__1136195496.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2__1136195496 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Rec_Class_tailRecM2__1136195496(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_2_box)))}
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2__1136195496
}

type Constructor_Control_Monad_Rec_Class_Loop struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Control_Monad_Rec_Class_Done struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Control_Monad_Rec_Class_MonadRec struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3709389635] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Rec_Class_MonadRec)(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "tailRecM": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Rec_Class_MonadRec: " + key)
		}
	}
}


func Call_Control_Monad_Rec_Class_MonadRec_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Rec_Class_tailRecM(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM2(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_Control_Monad_Rec_Class_tailRecM3(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var c_4 gopurs_runtime.Value = c_4_loop
_ = c_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, gopurs_runtime.RecordGet(o_5, "a"), gopurs_runtime.RecordGet(o_5, "b"), gopurs_runtime.RecordGet(o_5, "c"))
}), gopurs_runtime.RecordDict3("a", "b", "c", a_2, b_3, c_4))
}

func Call_Control_Monad_Rec_Class_untilJust(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, Get_Data_Unit_unit()})}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, (*Constructor_Data_Maybe_Just)(v1_4.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), m_2)
}), Get_Data_Unit_unit())
})
}

func Call_Control_Monad_Rec_Class_whileJust(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(dictMonadRec_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_2, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_2, "tailRecM"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, v_5})}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), v_5, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), m_4)
}), gopurs_runtime.Box(dictMonoid_0.V1))
})
})
}

func Call_Control_Monad_Rec_Class_tailRec(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_0, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRec2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var go__go_3_0_1 gopurs_runtime.Value
go__go_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_1:
for {
if false { continue go__go_3_0_1 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0, "b"))
continue go__go_3_0_1
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("a", "b", a_1, b_2)
_ = __local_var_4_2
return gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(__local_var_4_2, "a"), gopurs_runtime.RecordGet(__local_var_4_2, "b")))
}

func Call_Control_Monad_Rec_Class_tailRec3(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, c_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var c_3 gopurs_runtime.Value = c_3_loop
_ = c_3
var go__go_4_0_2 gopurs_runtime.Value
go__go_4_0_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_4_0_2:
for {
if false { continue go__go_4_0_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 525585346) {
v_5_loop = gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)(v_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)(v_5.UnsafePtr).V0, "b"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)(v_5.UnsafePtr).V0, "c"))
continue go__go_4_0_2
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.RecordDict3("a", "b", "c", a_1, b_2, c_3)
_ = __local_var_5_2
return gopurs_runtime.Apply(go__go_4_0_2, gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet(__local_var_5_2, "a"), gopurs_runtime.RecordGet(__local_var_5_2, "b"), gopurs_runtime.RecordGet(__local_var_5_2, "c")))
}

func Call_Control_Monad_Rec_Class_loop3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict3("a", "b", "c", a_0, b_1, c_2)})}
}

func Call_Control_Monad_Rec_Class_loop2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", a_0, b_1)})}
}

func Call_Control_Monad_Rec_Class_forever(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, u_3})}
}), ma_2)
}), Get_Data_Unit_unit())
})
}

func Call_Control_Monad_Rec_Class_tailRec__2110844386(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_7 gopurs_runtime.Value
go__go_1_0_7 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_7:
for {
if false { continue go__go_1_0_7 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__go_1_0_7
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_7, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRec__2334182452(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__go_1_0_8
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_8, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRec__2666749533(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_9 gopurs_runtime.Value
go__go_1_0_9 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0))})
continue go__go_1_0_9
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_9, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRec__2045907654(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__go_1_0_10
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRec__2929877587(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_11 gopurs_runtime.Value
go__go_1_0_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__go_1_0_11
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_11, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRecM__2468783173(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2333741125(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__3369162518(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2016576534(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2220253896(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2301013304(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__1121640472(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__3865988408(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__1299680280(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2119835928(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__1700625784(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2172295544(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__3782365624(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__1818438580(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__3515372237(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__3965932080(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2810501104(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__3478800400(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2063592441(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__1444729948(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM__2222286441(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM2__3241310373(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop int64) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 int64 = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.Int(gopurs_runtime.RecordGet(o_4, "b").IntVal))
}), gopurs_runtime.RecordDict2("a", "b", a_2, gopurs_runtime.Int(b_3)))
}

func Call_Control_Monad_Rec_Class_tailRecM2__2010968624(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_Control_Monad_Rec_Class_tailRecM2__1800958704(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_Control_Monad_Rec_Class_tailRecM2__1943630176(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_Control_Monad_Rec_Class_tailRecM2__2551820843(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_Control_Monad_Rec_Class_tailRecM2__3864450856(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.RecordDict2("a", "b", a_1, b_2)
_ = __local_var_3_0
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Control_Monad_Rec_Class_Done
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t5 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 525585346) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "b"))))}})}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 60402430) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t4)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t5)}
})
_ = __local_var_4_1
var go__go_5_6_12 gopurs_runtime.Value
go__go_5_6_12 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_5_6_12:
for {
if false { continue go__go_5_6_12 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t7 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 525585346) {
v_6_loop = gopurs_runtime.Apply(__local_var_4_1, (*Constructor_Control_Monad_Rec_Class_Loop)(v_6.UnsafePtr).V0)
continue go__go_5_6_12
__t7 = gopurs_runtime.Value{}
goto end_branch_7
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 60402430) {
__t7 = (*Constructor_Control_Monad_Rec_Class_Done)(v_6.UnsafePtr).V0
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}()
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_5_6_12, gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(__local_var_3_0, "a"), gopurs_runtime.RecordGet(__local_var_3_0, "b"))))})))
}

func Call_Control_Monad_Rec_Class_tailRecM2__1136195496(f_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_List_Types_Cons, b_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 *Constructor_Data_List_Types_Cons = a_1_loop
_ = a_1
var b_2 *Constructor_Data_List_Types_Cons = b_2_loop
_ = b_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)})
_ = __local_var_3_0
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Control_Monad_Rec_Class_Done
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t5 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 525585346) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "a")))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "b")))})))}})}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 60402430) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t4)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t5)}
})
_ = __local_var_4_1
var go__go_5_6_13 gopurs_runtime.Value
go__go_5_6_13 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_5_6_13:
for {
if false { continue go__go_5_6_13 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t7 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 525585346) {
v_6_loop = gopurs_runtime.Apply(__local_var_4_1, (*Constructor_Control_Monad_Rec_Class_Loop)(v_6.UnsafePtr).V0)
continue go__go_5_6_13
__t7 = gopurs_runtime.Value{}
goto end_branch_7
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 60402430) {
__t7 = (*Constructor_Control_Monad_Rec_Class_Done)(v_6.UnsafePtr).V0
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}()
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_5_6_13, gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(__local_var_3_0, "a")))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(__local_var_3_0, "b")))})))}))))})
}


