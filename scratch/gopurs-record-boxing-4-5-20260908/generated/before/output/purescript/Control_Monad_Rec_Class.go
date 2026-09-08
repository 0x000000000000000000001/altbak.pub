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
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Control_Monad_Rec_Class_Loop
}

var cache_Control_Monad_Rec_Class_Done gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Done sync.Once
func Get_Control_Monad_Rec_Class_Done() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Done.Do(func() {
		cache_Control_Monad_Rec_Class_Done = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Control_Monad_Rec_Class_Done
}

var cache_Control_Monad_Rec_Class_MonadRec_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Rec_Class_MonadRec_dollar_Dict sync.Once
func Get_Control_Monad_Rec_Class_MonadRec_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_MonadRec_dollar_Dict.Do(func() {
		cache_Control_Monad_Rec_Class_MonadRec_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Rec_Class_MonadRec_dollar_Dict(func() struct{
	Monad0 gopurs_runtime.Value
	tailRecM gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad0 gopurs_runtime.Value
	tailRecM gopurs_runtime.Value
}{}
					clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
					clone.tailRecM = gopurs_runtime.RecordGet(orig, "tailRecM")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Rec_Class_MonadRec_dollar_Dict
}

var cache_Control_Monad_Rec_Class_tailRecM gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM sync.Once
func Get_Control_Monad_Rec_Class_tailRecM() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM
}

var cache_Control_Monad_Rec_Class_tailRecM2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM2(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2
}

var cache_Control_Monad_Rec_Class_tailRecM3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM3 sync.Once
func Get_Control_Monad_Rec_Class_tailRecM3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM3.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM3 = gopurs_runtime.Func5(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_tailRecM3(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box, c_4_box)
})
	})
	return cache_Control_Monad_Rec_Class_tailRecM3
}

var cache_Control_Monad_Rec_Class_untilJust gopurs_runtime.Value
var once_Control_Monad_Rec_Class_untilJust sync.Once
func Get_Control_Monad_Rec_Class_untilJust() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_untilJust.Do(func() {
		cache_Control_Monad_Rec_Class_untilJust = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_untilJust(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_untilJust
}

var cache_Control_Monad_Rec_Class_whileJust gopurs_runtime.Value
var once_Control_Monad_Rec_Class_whileJust sync.Once
func Get_Control_Monad_Rec_Class_whileJust() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_whileJust.Do(func() {
		cache_Control_Monad_Rec_Class_whileJust = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_whileJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
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
		cache_Control_Monad_Rec_Class_monadRecMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_1542299734_4130553207((&Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_1642601656_2568689657(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_monadMaybe())))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Control_Monad_Rec_Class___local_var_2_0 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class___local_var_2_0
var __local_var_2_0 gopurs_runtime.Value
_ = __local_var_2_0
Call_local_Control_Monad_Rec_Class___local_var_2_0 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t6 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_802708012_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
goto end_branch_6
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 525585346) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_3406595152_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 60402430) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_802708012_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
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
__local_var_2_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class___local_var_2_0(v_2_loop_val)
})
var Call_local_Control_Monad_Rec_Class_go__go_3_7_3 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_3_7_3
var go__go_3_7_3 gopurs_runtime.Value
_ = go__go_3_7_3
Call_local_Control_Monad_Rec_Class_go__go_3_7_3 = func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_7_3:
for {
if false { continue go__go_3_7_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t8 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = Call_local_Control_Monad_Rec_Class___local_var_2_0((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_7_3
__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_8
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t8 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}
go__go_3_7_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_3_7_3(v_4_loop_val)
})
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Control_Monad_Rec_Class_go__go_3_7_3(Call_local_Control_Monad_Rec_Class___local_var_2_0(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a0_1)))}))))}
})})))}
	})
	return cache_Control_Monad_Rec_Class_monadRecMaybe
}

var cache_Control_Monad_Rec_Class_monadRecIdentity gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecIdentity sync.Once
func Get_Control_Monad_Rec_Class_monadRecIdentity() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecIdentity.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Control_Monad_Rec_Class_go__go_1_1_4 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_1_1_4
var go__go_1_1_4 gopurs_runtime.Value
_ = go__go_1_1_4
Call_local_Control_Monad_Rec_Class_go__go_1_1_4 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_1_4:
for {
if false { continue go__go_1_1_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_1_4
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t2 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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
}
go__go_1_1_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_1_1_4(v_2_loop_val)
})
// TAST (Let): __local_var_1_0 shape=LetRec(Abs(App(Other))) bindingType=(Func [(TypeVar a)] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_1_1_4(gopurs_runtime.Apply(f_0, x_2))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecIdentity
}

var cache_Control_Monad_Rec_Class_monadRecFunction gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecFunction sync.Once
func Get_Control_Monad_Rec_Class_monadRecFunction() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecFunction.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecFunction = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_monadFn()))}
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value, e_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Control_Monad_Rec_Class_go__go_3_0_5 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_3_0_5
var go__go_3_0_5 gopurs_runtime.Value
_ = go__go_3_0_5
Call_local_Control_Monad_Rec_Class_go__go_3_0_5 = func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_0_5:
for {
if false { continue go__go_3_0_5 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, e_2)
continue go__go_3_0_5
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
}
go__go_3_0_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_3_0_5(v_4_loop_val)
})
return Call_local_Control_Monad_Rec_Class_go__go_3_0_5(gopurs_runtime.Apply2(f_0, a0_1, e_2))
})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecFunction
}

var cache_Control_Monad_Rec_Class_monadRecEither gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecEither sync.Once
func Get_Control_Monad_Rec_Class_monadRecEither() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecEither.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecEither = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Either_monadEither()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Control_Monad_Rec_Class___local_var_2_0 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class___local_var_2_0
var __local_var_2_0 gopurs_runtime.Value
_ = __local_var_2_0
Call_local_Control_Monad_Rec_Class___local_var_2_0 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t4 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0}))}}))}
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)}))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0}))}}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
__local_var_2_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class___local_var_2_0(v_2_loop_val)
})
var Call_local_Control_Monad_Rec_Class_go__go_3_5_6 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_3_5_6
var go__go_3_5_6 gopurs_runtime.Value
_ = go__go_3_5_6
Call_local_Control_Monad_Rec_Class_go__go_3_5_6 = func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_5_6:
for {
if false { continue go__go_3_5_6 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = Call_local_Control_Monad_Rec_Class___local_var_2_0((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_6
__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
}
go__go_3_5_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_3_5_6(v_4_loop_val)
})
return Call_local_Control_Monad_Rec_Class_go__go_3_5_6(Call_local_Control_Monad_Rec_Class___local_var_2_0(gopurs_runtime.Apply(f_0, a0_1)))
})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecEither
}

var cache_Control_Monad_Rec_Class_monadRecEffect gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecEffect sync.Once
func Get_Control_Monad_Rec_Class_monadRecEffect() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecEffect.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecEffect = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Effect","Effect"] [(ADT ["Control","Monad","Rec","Class","Step"] [(TypeVar a), (TypeVar b)])])
__local_var_2_0 := gopurs_runtime.Apply(f_0, a_1)
_ = __local_var_2_0
__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_2
r_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), __local_var_3_2), gopurs_runtime.Value{})
_ = r_3_1
_dollar___unused_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 shape=App(Var) bindingType=Any
__local_var_4_4 := gopurs_runtime.Apply(Get_Effect_Ref_read(), r_3_1)
_ = __local_var_4_4
v_5_5 := gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Value{})
_ = v_5_5
var __t8 gopurs_runtime.Value
{
if (v_5_5.Type == 9 && v_5_5.IntVal == 525585346) {
__t8 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
e_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_5.UnsafePtr).V0), gopurs_runtime.Value{})
_ = e_6_6
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), e_6_6, r_3_1), gopurs_runtime.Value{})
_ = __local_var_7_7
return gopurs_runtime.Bool(false)
})
goto end_branch_8
} else {

}
}
{
if (v_5_5.Type == 9 && v_5_5.IntVal == 60402430) {
__t8 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Apply(__t8, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
_ = _dollar___unused_4_3
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 60402430) {
__t9 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Apply(Get_Effect_Ref_read(), r_3_1)), gopurs_runtime.Value{})
})
})}))}
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
		cache_Control_Monad_Rec_Class_functorStep = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0}))}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Control_Monad_Rec_Class_functorStep
}

var cache_Control_Monad_Rec_Class_forever gopurs_runtime.Value
var once_Control_Monad_Rec_Class_forever sync.Once
func Get_Control_Monad_Rec_Class_forever() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_forever.Do(func() {
		cache_Control_Monad_Rec_Class_forever = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Rec_Class_forever(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Control_Monad_Rec_Class_forever
}

var cache_Control_Monad_Rec_Class_bifunctorStep gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bifunctorStep sync.Once
func Get_Control_Monad_Rec_Class_bifunctorStep() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bifunctorStep.Do(func() {
		cache_Control_Monad_Rec_Class_bifunctorStep = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)}))}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Control_Monad_Rec_Class_bifunctorStep
}

type Constructor_Control_Monad_Rec_Class_Loop[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Control_Monad_Rec_Class_Done[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}


type Constructor_Control_Monad_Rec_Class_MonadRec[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3709389635] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Rec_Class_MonadRec[any])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "tailRecM": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Rec_Class_MonadRec: " + key)
		}
	}
}


func Call_Control_Monad_Rec_Class_MonadRec_dollar_Dict(x_0_loop struct{
	Monad0 gopurs_runtime.Value
	tailRecM gopurs_runtime.Value
}) *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] {
var x_0 struct{
	Monad0 gopurs_runtime.Value
	tailRecM gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Monad0", "tailRecM"}, []gopurs_runtime.Value{orig.Monad0, orig.tailRecM})
				}())
}

func Call_Control_Monad_Rec_Class_tailRecM(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Rec_Class_tailRecM2(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), func() gopurs_runtime.Value {
				orig := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{a_2, b_3}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{orig.a, orig.b})
				}())
}

func Call_Control_Monad_Rec_Class_tailRecM3(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
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
}), func() gopurs_runtime.Value {
				orig := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
}{a_2, b_3, c_4}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b", "c"}, []gopurs_runtime.Value{orig.a, orig.b, orig.c})
				}())
}

func Call_Control_Monad_Rec_Class_untilJust(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_4)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_4)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_4.UnsafePtr).V0}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), m_2)
}), Get_Data_Unit_unit())
})
}

func Call_Control_Monad_Rec_Class_whileJust(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(dictMonadRec_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_2, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_2, "tailRecM"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_6)
if (__t_tag_2 == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_5}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_6)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), v_5, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_6.UnsafePtr).V0)}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), m_4)
}), gopurs_runtime.Box(dictMonoid_0.V1))
})
})
}

func Call_Control_Monad_Rec_Class_tailRec(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var Call_local_Control_Monad_Rec_Class_go__go_1_0_0 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_1_0_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
Call_local_Control_Monad_Rec_Class_go__go_1_0_0 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_0
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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
}
go__go_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_1_0_0(v_2_loop_val)
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_1_0_0(gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_Control_Monad_Rec_Class_tailRec2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var Call_local_Control_Monad_Rec_Class_go__go_3_0_1 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_3_0_1
var go__go_3_0_1 gopurs_runtime.Value
_ = go__go_3_0_1
Call_local_Control_Monad_Rec_Class_go__go_3_0_1 = func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_0_1:
for {
if false { continue go__go_3_0_1 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, "b"))
continue go__go_3_0_1
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
}
go__go_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_3_0_1(v_4_loop_val)
})
// TAST (Let): __local_var_4_2 shape=LitRecord bindingType=(Record (Row [a: (TypeVar a), b: (TypeVar b)] Any))
__local_var_4_2 := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{a_1, b_2}
_ = __local_var_4_2
return Call_local_Control_Monad_Rec_Class_go__go_3_0_1(gopurs_runtime.Apply2(f_0, __local_var_4_2.a, __local_var_4_2.b))
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
var Call_local_Control_Monad_Rec_Class_go__go_4_0_2 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Control_Monad_Rec_Class_go__go_4_0_2
var go__go_4_0_2 gopurs_runtime.Value
_ = go__go_4_0_2
Call_local_Control_Monad_Rec_Class_go__go_4_0_2 = func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_0_2:
for {
if false { continue go__go_4_0_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 525585346) {
v_5_loop = gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "b"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "c"))
continue go__go_4_0_2
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 60402430) {
__t1 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
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
}
go__go_4_0_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Control_Monad_Rec_Class_go__go_4_0_2(v_5_loop_val)
})
// TAST (Let): __local_var_5_2 shape=LitRecord bindingType=(Record (Row [a: (TypeVar a), b: (TypeVar b), c: (TypeVar c)] Any))
__local_var_5_2 := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
}{a_1, b_2, c_3}
_ = __local_var_5_2
return Call_local_Control_Monad_Rec_Class_go__go_4_0_2(gopurs_runtime.Apply3(f_0, __local_var_5_2.a, __local_var_5_2.b, __local_var_5_2.c))
}

func Call_Control_Monad_Rec_Class_loop3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_1030879592_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
} {
					orig := gopurs_runtime.RecordDict3("a", "b", "c", a_0, b_1, c_2)
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b")
					clone.c = gopurs_runtime.RecordGet(orig, "c")
					return clone
				}()})))}
}

func Call_Control_Monad_Rec_Class_loop2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
} {
					orig := gopurs_runtime.RecordDict2("a", "b", a_0, b_1)
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b")
					return clone
				}()})))}
}

func Call_Control_Monad_Rec_Class_forever(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, u_3}))}
}), ma_2)
}), Get_Data_Unit_unit())
})
}

func Rebox_Control_Monad_Rec_Class_1030879592_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b", "c"}, []gopurs_runtime.Value{orig.a, orig.b, orig.c})
				}()
	return out
}

func Rebox_Control_Monad_Rec_Class_1542299734_4130553207(in *Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Rec_Class_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Rec_Class_3406595152_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_Rec_Class_784458114_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{orig.a, orig.b})
				}()
	return out
}

func Rebox_Control_Monad_Rec_Class_802708012_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}


