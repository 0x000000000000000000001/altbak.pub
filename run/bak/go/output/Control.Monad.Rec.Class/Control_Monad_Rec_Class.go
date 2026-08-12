package Control_Monad_Rec_Class

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Ref "gopurs/output/Effect.Ref"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Loop gopurs_runtime.Value
var once_Loop sync.Once
func Get_Loop() gopurs_runtime.Value {
	once_Loop.Do(func() {
		cache_Loop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Loop
}

var cache_Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		cache_Done = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Done
}

var cache_tailRecM gopurs_runtime.Value
var once_tailRecM sync.Once
func Get_tailRecM() gopurs_runtime.Value {
	once_tailRecM.Do(func() {
		cache_tailRecM = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM
}

var cache_tailRecM__gopurs_runtime_Value_3865988408 gopurs_runtime.Value
var once_tailRecM__gopurs_runtime_Value_3865988408 sync.Once
func Get_tailRecM__gopurs_runtime_Value_3865988408() gopurs_runtime.Value {
	once_tailRecM__gopurs_runtime_Value_3865988408.Do(func() {
		cache_tailRecM__gopurs_runtime_Value_3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__gopurs_runtime_Value_3865988408(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__gopurs_runtime_Value_3865988408
}

var cache_tailRecM2 gopurs_runtime.Value
var once_tailRecM2 sync.Once
func Get_tailRecM2() gopurs_runtime.Value {
	once_tailRecM2.Do(func() {
		cache_tailRecM2 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2
}

var cache_tailRecM2__gopurs_runtime_Value_1943630176 gopurs_runtime.Value
var once_tailRecM2__gopurs_runtime_Value_1943630176 sync.Once
func Get_tailRecM2__gopurs_runtime_Value_1943630176() gopurs_runtime.Value {
	once_tailRecM2__gopurs_runtime_Value_1943630176.Do(func() {
		cache_tailRecM2__gopurs_runtime_Value_1943630176 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__gopurs_runtime_Value_1943630176(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2__gopurs_runtime_Value_1943630176
}

var cache_tailRecM3 gopurs_runtime.Value
var once_tailRecM3 sync.Once
func Get_tailRecM3() gopurs_runtime.Value {
	once_tailRecM3.Do(func() {
		cache_tailRecM3 = gopurs_runtime.Func5(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM3(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box, c_4_box)
})
	})
	return cache_tailRecM3
}

var cache_untilJust gopurs_runtime.Value
var once_untilJust sync.Once
func Get_untilJust() gopurs_runtime.Value {
	once_untilJust.Do(func() {
		cache_untilJust = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_untilJust(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_untilJust
}

var cache_whileJust gopurs_runtime.Value
var once_whileJust sync.Once
func Get_whileJust() gopurs_runtime.Value {
	once_whileJust.Do(func() {
		cache_whileJust = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_whileJust(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_whileJust
}

var cache_tailRec gopurs_runtime.Value
var once_tailRec sync.Once
func Get_tailRec() gopurs_runtime.Value {
	once_tailRec.Do(func() {
		cache_tailRec = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec(f_0_box)
})
	})
	return cache_tailRec
}

var cache_tailRec__gopurs_runtime_Value_2110844386 gopurs_runtime.Value
var once_tailRec__gopurs_runtime_Value_2110844386 sync.Once
func Get_tailRec__gopurs_runtime_Value_2110844386() gopurs_runtime.Value {
	once_tailRec__gopurs_runtime_Value_2110844386.Do(func() {
		cache_tailRec__gopurs_runtime_Value_2110844386 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__gopurs_runtime_Value_2110844386(f_0_box)
})
	})
	return cache_tailRec__gopurs_runtime_Value_2110844386
}

var cache_tailRec2 gopurs_runtime.Value
var once_tailRec2 sync.Once
func Get_tailRec2() gopurs_runtime.Value {
	once_tailRec2.Do(func() {
		cache_tailRec2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_tailRec2
}

var cache_tailRec3 gopurs_runtime.Value
var once_tailRec3 sync.Once
func Get_tailRec3() gopurs_runtime.Value {
	once_tailRec3.Do(func() {
		cache_tailRec3 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, c_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec3(f_0_box, a_1_box, b_2_box, c_3_box)
})
	})
	return cache_tailRec3
}

var cache_monadRecMaybe gopurs_runtime.Value
var once_monadRecMaybe sync.Once
func Get_monadRecMaybe() gopurs_runtime.Value {
	once_monadRecMaybe.Do(func() {
		cache_monadRecMaybe = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 60402430) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
_ = __local_var_2_0
var go__go_3_5_4 gopurs_runtime.Value
go__go_3_5_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_4:
for {
if false { continue go__go_3_5_4 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_4
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_5_4, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))))}
})
}))
	})
	return cache_monadRecMaybe
}

var cache_monadRecMaybe__ptrConstructor_MonadRec_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___796215523 gopurs_runtime.Value
var once_monadRecMaybe__ptrConstructor_MonadRec_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___796215523 sync.Once
func Get_monadRecMaybe__ptrConstructor_MonadRec_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___796215523() gopurs_runtime.Value {
	once_monadRecMaybe__ptrConstructor_MonadRec_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___796215523.Do(func() {
		cache_monadRecMaybe__ptrConstructor_MonadRec_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___796215523 = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_MonadRec[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 60402430) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
_ = __local_var_2_0
var go__go_3_5_5 gopurs_runtime.Value
go__go_3_5_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_5:
for {
if false { continue go__go_3_5_5 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_5
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_5_5, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))))}
})
})})}
	})
	return cache_monadRecMaybe__ptrConstructor_MonadRec_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___796215523
}

var cache_monadRecIdentity gopurs_runtime.Value
var once_monadRecIdentity sync.Once
func Get_monadRecIdentity() gopurs_runtime.Value {
	once_monadRecIdentity.Do(func() {
		cache_monadRecIdentity = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_monadIdentity()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_6 gopurs_runtime.Value
go__go_1_0_6 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_6:
for {
if false { continue go__go_1_0_6 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_6
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__go_1_0_6, gopurs_runtime.Apply(f_0, x_2))
})
}))
	})
	return cache_monadRecIdentity
}

var cache_monadRecFunction gopurs_runtime.Value
var once_monadRecFunction sync.Once
func Get_monadRecFunction() gopurs_runtime.Value {
	once_monadRecFunction.Do(func() {
		cache_monadRecFunction = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_7 gopurs_runtime.Value
go__go_3_0_7 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_7:
for {
if false { continue go__go_3_0_7 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, e_2)
continue go__go_3_0_7
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__go_3_0_7, gopurs_runtime.Apply2(f_0, a0_1, e_2))
})
})
}))
	})
	return cache_monadRecFunction
}

var cache_monadRecEither gopurs_runtime.Value
var once_monadRecEither sync.Once
func Get_monadRecEither() gopurs_runtime.Value {
	once_monadRecEither.Do(func() {
		cache_monadRecEither = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_monadEither()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0})}})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 60402430) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
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
_ = __local_var_2_0
var go__go_3_5_8 gopurs_runtime.Value
go__go_3_5_8 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_5_8:
for {
if false { continue go__go_3_5_8 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_8
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t6 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__go_3_5_8, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))
})
}))
	})
	return cache_monadRecEither
}

var cache_monadRecEffect gopurs_runtime.Value
var once_monadRecEffect sync.Once
func Get_monadRecEffect() gopurs_runtime.Value {
	once_monadRecEffect.Do(func() {
		cache_monadRecEffect = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(f_0, a_1), pkg_Effect_Ref.Get__new()), gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect(), gopurs_runtime.Apply(pkg_Effect.Get_untilE(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), r_2), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 525585346) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Effect_Ref.Get_write(), e_4, r_2), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.Bool(false))
}))
}))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.Bool(true))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), r_2))
}))
}))
})
}))
	})
	return cache_monadRecEffect
}

var cache_loop3 gopurs_runtime.Value
var once_loop3 sync.Once
func Get_loop3() gopurs_runtime.Value {
	once_loop3.Do(func() {
		cache_loop3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_loop3(a_0_box, b_1_box, c_2_box)
})
	})
	return cache_loop3
}

var cache_loop2 gopurs_runtime.Value
var once_loop2 sync.Once
func Get_loop2() gopurs_runtime.Value {
	once_loop2.Do(func() {
		cache_loop2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_loop2(a_0_box, b_1_box)
})
	})
	return cache_loop2
}

var cache_functorStep gopurs_runtime.Value
var once_functorStep sync.Once
func Get_functorStep() gopurs_runtime.Value {
	once_functorStep.Do(func() {
		cache_functorStep = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
}))
	})
	return cache_functorStep
}

var cache_forever gopurs_runtime.Value
var once_forever sync.Once
func Get_forever() gopurs_runtime.Value {
	once_forever.Do(func() {
		cache_forever = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_forever(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_forever
}

var cache_bifunctorStep gopurs_runtime.Value
var once_bifunctorStep sync.Once
func Get_bifunctorStep() gopurs_runtime.Value {
	once_bifunctorStep.Do(func() {
		cache_bifunctorStep = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
}))
	})
	return cache_bifunctorStep
}

var cache_bifunctorStep__ptrData_Bifunctor_Constructor_Bifunctor_gopurs_runtime_Value__1527141292 gopurs_runtime.Value
var once_bifunctorStep__ptrData_Bifunctor_Constructor_Bifunctor_gopurs_runtime_Value__1527141292 sync.Once
func Get_bifunctorStep__ptrData_Bifunctor_Constructor_Bifunctor_gopurs_runtime_Value__1527141292() gopurs_runtime.Value {
	once_bifunctorStep__ptrData_Bifunctor_Constructor_Bifunctor_gopurs_runtime_Value__1527141292.Do(func() {
		cache_bifunctorStep__ptrData_Bifunctor_Constructor_Bifunctor_gopurs_runtime_Value__1527141292 = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(&pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
	return cache_bifunctorStep__ptrData_Bifunctor_Constructor_Bifunctor_gopurs_runtime_Value__1527141292
}

type Constructor_Loop[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Done[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}


type Constructor_MonadRec[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3709389635] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadRec[gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "tailRecM": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadRec: " + key)
		}
	}
}


func Call_tailRecM(dict_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__gopurs_runtime_Value_3865988408(dict_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM2(dictMonadRec_0_loop *Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_tailRecM2__gopurs_runtime_Value_1943630176(dictMonadRec_0_loop *Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_tailRecM3(dictMonadRec_0_loop *Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var c_4 gopurs_runtime.Value = c_4_loop
_ = c_4
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, gopurs_runtime.RecordGet(o_5, "a"), gopurs_runtime.RecordGet(o_5, "b"), gopurs_runtime.RecordGet(o_5, "c"))
}), gopurs_runtime.RecordDict3("a", "b", "c", a_2, b_3, c_4))
}

func Call_untilJust(dictMonadRec_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_4.UnsafePtr).V0})}
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
}), pkg_Data_Unit.Get_unit())
})
}

func Call_whileJust(dictMonoid_0_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(dictMonadRec_2 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_2, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_2, "tailRecM"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_3_1.V0, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_5})}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Semigroup0_1_0.V0, v_5, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_6.UnsafePtr).V0)})}
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
}), dictMonoid_0.V1)
})
})
}

func Call_tailRec(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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

func Call_tailRec__gopurs_runtime_Value_2110844386(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_1
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__go_1_0_1, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var go__go_3_0_2 gopurs_runtime.Value
go__go_3_0_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_2:
for {
if false { continue go__go_3_0_2 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, "b"))
continue go__go_3_0_2
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__go_3_0_2, gopurs_runtime.Apply2(f_0, a_1, b_2))
}

func Call_tailRec3(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, c_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var c_3 gopurs_runtime.Value = c_3_loop
_ = c_3
var go__go_4_0_3 gopurs_runtime.Value
go__go_4_0_3 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_4_0_3:
for {
if false { continue go__go_4_0_3 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 525585346) {
v_5_loop = gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "b"), gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "c"))
continue go__go_4_0_3
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 60402430) {
__t1 = (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__go_4_0_3, gopurs_runtime.Apply3(f_0, a_1, b_2, c_3))
}

func Call_loop3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict3("a", "b", "c", a_0, b_1, c_2)})}
}

func Call_loop2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", a_0, b_1)})}
}

func Call_forever(dictMonadRec_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, u_3})}
}), ma_2)
}), pkg_Data_Unit.Get_unit())
})
}


