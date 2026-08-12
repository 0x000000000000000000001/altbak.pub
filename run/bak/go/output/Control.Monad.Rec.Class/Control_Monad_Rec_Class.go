package Control_Monad_Rec_Class

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Ref "gopurs/output/Effect.Ref"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
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
var __t4 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}})}
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
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
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_3
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_3_5_3, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))))}
})
}))
	})
	return cache_monadRecMaybe
}

var cache_monadRecIdentity gopurs_runtime.Value
var once_monadRecIdentity sync.Once
func Get_monadRecIdentity() gopurs_runtime.Value {
	once_monadRecIdentity.Do(func() {
		cache_monadRecIdentity = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_monadIdentity()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_4 gopurs_runtime.Value
go__go_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_4:
for {
if false { continue go__go_1_0_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_4
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
return gopurs_runtime.Apply(go__go_1_0_4, gopurs_runtime.Apply(f_0, x_2))
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
v_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, e_2)
continue go__go_3_0_5
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
return gopurs_runtime.Apply(go__go_3_0_5, gopurs_runtime.Apply2(f_0, a0_1, e_2))
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
var __t4 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0})}})}
goto end_branch_4
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
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
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
continue go__go_3_5_6
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
return gopurs_runtime.Apply(go__go_3_5_6, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))
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
return gopurs_runtime.Apply2(Get_bind__2951621345(), gopurs_runtime.Apply2(pkg_Effect_Ref.Get_write(), e_4, r_2), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_pure__4209427318(), gopurs_runtime.Bool(false))
}))
}))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t0 = gopurs_runtime.Apply(Get_pure__4209427318(), gopurs_runtime.Bool(true))
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

var cache_applicativeFn__3751223912 gopurs_runtime.Value
var once_applicativeFn__3751223912 sync.Once
func Get_applicativeFn__3751223912() gopurs_runtime.Value {
	once_applicativeFn__3751223912.Do(func() {
		cache_applicativeFn__3751223912 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_applicativeFn__3751223912
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__4209427318 gopurs_runtime.Value
var once_pure__4209427318 sync.Once
func Get_pure__4209427318() gopurs_runtime.Value {
	once_pure__4209427318.Do(func() {
		cache_pure__4209427318 = gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure")
	})
	return cache_pure__4209427318
}

var cache_applyFn__4042184691 gopurs_runtime.Value
var once_applyFn__4042184691 sync.Once
func Get_applyFn__4042184691() gopurs_runtime.Value {
	once_applyFn__4042184691.Do(func() {
		cache_applyFn__4042184691 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_applyFn__4042184691
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__2951621345 gopurs_runtime.Value
var once_bind__2951621345 sync.Once
func Get_bind__2951621345() gopurs_runtime.Value {
	once_bind__2951621345.Do(func() {
		cache_bind__2951621345 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__2951621345
}

var cache_bind__3119797153 gopurs_runtime.Value
var once_bind__3119797153 sync.Once
func Get_bind__3119797153() gopurs_runtime.Value {
	once_bind__3119797153.Do(func() {
		cache_bind__3119797153 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__3119797153
}

var cache_bind__1325495585 gopurs_runtime.Value
var once_bind__1325495585 sync.Once
func Get_bind__1325495585() gopurs_runtime.Value {
	once_bind__1325495585.Do(func() {
		cache_bind__1325495585 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__1325495585
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_bindFlipped__1317599105 gopurs_runtime.Value
var once_bindFlipped__1317599105 sync.Once
func Get_bindFlipped__1317599105() gopurs_runtime.Value {
	once_bindFlipped__1317599105.Do(func() {
		cache_bindFlipped__1317599105 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1317599105(b_0_box, a_1_box)
})
	})
	return cache_bindFlipped__1317599105
}

var cache_bindFn__1648334822 gopurs_runtime.Value
var once_bindFn__1648334822 sync.Once
func Get_bindFn__1648334822() gopurs_runtime.Value {
	once_bindFn__1648334822.Do(func() {
		cache_bindFn__1648334822 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
})
})
}))
	})
	return cache_bindFn__1648334822
}

var cache_discard__1979268384 gopurs_runtime.Value
var once_discard__1979268384 sync.Once
func Get_discard__1979268384() gopurs_runtime.Value {
	once_discard__1979268384.Do(func() {
		cache_discard__1979268384 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard__1979268384
}

var cache_discard__317162198 gopurs_runtime.Value
var once_discard__317162198 sync.Once
func Get_discard__317162198() gopurs_runtime.Value {
	once_discard__317162198.Do(func() {
		cache_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__317162198(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard__317162198
}

var cache_discardUnit__2687062302 gopurs_runtime.Value
var once_discardUnit__2687062302 sync.Once
func Get_discardUnit__2687062302() gopurs_runtime.Value {
	once_discardUnit__2687062302.Do(func() {
		cache_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit__2687062302
}

var cache_tailRec__2110844386 gopurs_runtime.Value
var once_tailRec__2110844386 sync.Once
func Get_tailRec__2110844386() gopurs_runtime.Value {
	once_tailRec__2110844386.Do(func() {
		cache_tailRec__2110844386 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2110844386(f_0_box)
})
	})
	return cache_tailRec__2110844386
}

var cache_tailRec__2334182452 gopurs_runtime.Value
var once_tailRec__2334182452 sync.Once
func Get_tailRec__2334182452() gopurs_runtime.Value {
	once_tailRec__2334182452.Do(func() {
		cache_tailRec__2334182452 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2334182452(f_0_box)
})
	})
	return cache_tailRec__2334182452
}

var cache_tailRec__2666749533 gopurs_runtime.Value
var once_tailRec__2666749533 sync.Once
func Get_tailRec__2666749533() gopurs_runtime.Value {
	once_tailRec__2666749533.Do(func() {
		cache_tailRec__2666749533 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2666749533(f_0_box)
})
	})
	return cache_tailRec__2666749533
}

var cache_tailRec__2045907654 gopurs_runtime.Value
var once_tailRec__2045907654 sync.Once
func Get_tailRec__2045907654() gopurs_runtime.Value {
	once_tailRec__2045907654.Do(func() {
		cache_tailRec__2045907654 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2045907654(f_0_box)
})
	})
	return cache_tailRec__2045907654
}

var cache_tailRec__2929877587 gopurs_runtime.Value
var once_tailRec__2929877587 sync.Once
func Get_tailRec__2929877587() gopurs_runtime.Value {
	once_tailRec__2929877587.Do(func() {
		cache_tailRec__2929877587 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec__2929877587(f_0_box)
})
	})
	return cache_tailRec__2929877587
}

var cache_tailRecM__2220253896 gopurs_runtime.Value
var once_tailRecM__2220253896 sync.Once
func Get_tailRecM__2220253896() gopurs_runtime.Value {
	once_tailRecM__2220253896.Do(func() {
		cache_tailRecM__2220253896 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__2220253896(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__2220253896
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__1444729948 gopurs_runtime.Value
var once_tailRecM__1444729948 sync.Once
func Get_tailRecM__1444729948() gopurs_runtime.Value {
	once_tailRecM__1444729948.Do(func() {
		cache_tailRecM__1444729948 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__1444729948(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__1444729948
}

var cache_tailRecM__2222286441 gopurs_runtime.Value
var once_tailRecM__2222286441 sync.Once
func Get_tailRecM__2222286441() gopurs_runtime.Value {
	once_tailRecM__2222286441.Do(func() {
		cache_tailRecM__2222286441 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__2222286441(gopurs_runtime.CoerceToStruct[Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__2222286441
}

var cache_monadFn__1938941618 gopurs_runtime.Value
var once_monadFn__1938941618 sync.Once
func Get_monadFn__1938941618() gopurs_runtime.Value {
	once_monadFn__1938941618.Do(func() {
		cache_monadFn__1938941618 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeFn()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindFn()
}))
	})
	return cache_monadFn__1938941618
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_applicativeEither__2440223464 gopurs_runtime.Value
var once_applicativeEither__2440223464 sync.Once
func Get_applicativeEither__2440223464() gopurs_runtime.Value {
	once_applicativeEither__2440223464.Do(func() {
		cache_applicativeEither__2440223464 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_applyEither()
}), pkg_Data_Either.Get_Right())
	})
	return cache_applicativeEither__2440223464
}

var cache_applyEither__3806012498 gopurs_runtime.Value
var once_applyEither__3806012498 sync.Once
func Get_applyEither__3806012498() gopurs_runtime.Value {
	once_applyEither__3806012498.Do(func() {
		cache_applyEither__3806012498 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map"), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
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
	return cache_applyEither__3806012498
}

var cache_bindEither__3337174823 gopurs_runtime.Value
var once_bindEither__3337174823 sync.Once
func Get_bindEither__3337174823() gopurs_runtime.Value {
	once_bindEither__3337174823.Do(func() {
		cache_bindEither__3337174823 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_applyEither()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__local_var_1_0 := (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_1_0})}
})
goto end_branch_2
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__local_var_1_1 := (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_1
__t2 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, __local_var_1_1)
})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
	})
	return cache_bindEither__3337174823
}

var cache_either__2158544585 gopurs_runtime.Value
var once_either__2158544585 sync.Once
func Get_either__2158544585() gopurs_runtime.Value {
	once_either__2158544585.Do(func() {
		cache_either__2158544585 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__2158544585(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__2158544585
}

var cache_either__271265665 gopurs_runtime.Value
var once_either__271265665 sync.Once
func Get_either__271265665() gopurs_runtime.Value {
	once_either__271265665.Do(func() {
		cache_either__271265665 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__271265665(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__271265665
}

var cache_functorEither__13820179 gopurs_runtime.Value
var once_functorEither__13820179 sync.Once
func Get_functorEither__13820179() gopurs_runtime.Value {
	once_functorEither__13820179.Do(func() {
		cache_functorEither__13820179 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
	return cache_functorEither__13820179
}

var cache_functorEither__1771778897 gopurs_runtime.Value
var once_functorEither__1771778897 sync.Once
func Get_functorEither__1771778897() gopurs_runtime.Value {
	once_functorEither__1771778897.Do(func() {
		cache_functorEither__1771778897 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
	return cache_functorEither__1771778897
}

var cache_monadEither__2975460307 gopurs_runtime.Value
var once_monadEither__2975460307 sync.Once
func Get_monadEither__2975460307() gopurs_runtime.Value {
	once_monadEither__2975460307.Do(func() {
		cache_monadEither__2975460307 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_bindEither()
}))
	})
	return cache_monadEither__2975460307
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
}

var cache_functorFn__20325936 gopurs_runtime.Value
var once_functorFn__20325936 sync.Once
func Get_functorFn__20325936() gopurs_runtime.Value {
	once_functorFn__20325936.Do(func() {
		cache_functorFn__20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__20325936
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__3699108444 gopurs_runtime.Value
var once_map__3699108444 sync.Once
func Get_map__3699108444() gopurs_runtime.Value {
	once_map__3699108444.Do(func() {
		cache_map__3699108444 = gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map")
	})
	return cache_map__3699108444
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_map__4102685939 gopurs_runtime.Value
var once_map__4102685939 sync.Once
func Get_map__4102685939() gopurs_runtime.Value {
	once_map__4102685939.Do(func() {
		cache_map__4102685939 = gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map")
	})
	return cache_map__4102685939
}

var cache_mapFlipped__4215217780 gopurs_runtime.Value
var once_mapFlipped__4215217780 sync.Once
func Get_mapFlipped__4215217780() gopurs_runtime.Value {
	once_mapFlipped__4215217780.Do(func() {
		cache_mapFlipped__4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__4215217780(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__4215217780
}

var cache_mapFlipped__3249733428 gopurs_runtime.Value
var once_mapFlipped__3249733428 sync.Once
func Get_mapFlipped__3249733428() gopurs_runtime.Value {
	once_mapFlipped__3249733428.Do(func() {
		cache_mapFlipped__3249733428 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__3249733428(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__3249733428
}

var cache_mapFlipped__1087756276 gopurs_runtime.Value
var once_mapFlipped__1087756276 sync.Once
func Get_mapFlipped__1087756276() gopurs_runtime.Value {
	once_mapFlipped__1087756276.Do(func() {
		cache_mapFlipped__1087756276 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__1087756276(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__1087756276
}

var cache_voidRight__1142845180 gopurs_runtime.Value
var once_voidRight__1142845180 sync.Once
func Get_voidRight__1142845180() gopurs_runtime.Value {
	once_voidRight__1142845180.Do(func() {
		cache_voidRight__1142845180 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight__1142845180(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
})
	})
	return cache_voidRight__1142845180
}

var cache_voidRight__698766972 gopurs_runtime.Value
var once_voidRight__698766972 sync.Once
func Get_voidRight__698766972() gopurs_runtime.Value {
	once_voidRight__698766972.Do(func() {
		cache_voidRight__698766972 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight__698766972(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
})
	})
	return cache_voidRight__698766972
}

var cache_applicativeIdentity__4045440648 gopurs_runtime.Value
var once_applicativeIdentity__4045440648 sync.Once
func Get_applicativeIdentity__4045440648() gopurs_runtime.Value {
	once_applicativeIdentity__4045440648.Do(func() {
		cache_applicativeIdentity__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_applyIdentity()
}), pkg_Data_Identity.Get_Identity())
	})
	return cache_applicativeIdentity__4045440648
}

var cache_applyIdentity__3199351098 gopurs_runtime.Value
var once_applyIdentity__3199351098 sync.Once
func Get_applyIdentity__3199351098() gopurs_runtime.Value {
	once_applyIdentity__3199351098.Do(func() {
		cache_applyIdentity__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyIdentity__3199351098
}

var cache_bindIdentity__329376103 gopurs_runtime.Value
var once_bindIdentity__329376103 sync.Once
func Get_bindIdentity__329376103() gopurs_runtime.Value {
	once_bindIdentity__329376103.Do(func() {
		cache_bindIdentity__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_applyIdentity()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindIdentity__329376103
}

var cache_functorIdentity__943655089 gopurs_runtime.Value
var once_functorIdentity__943655089 sync.Once
func Get_functorIdentity__943655089() gopurs_runtime.Value {
	once_functorIdentity__943655089.Do(func() {
		cache_functorIdentity__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorIdentity__943655089
}

var cache_monadIdentity__1104192371 gopurs_runtime.Value
var once_monadIdentity__1104192371 sync.Once
func Get_monadIdentity__1104192371() gopurs_runtime.Value {
	once_monadIdentity__1104192371.Do(func() {
		cache_monadIdentity__1104192371 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_applicativeIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_bindIdentity()
}))
	})
	return cache_monadIdentity__1104192371
}

var cache_applicativeMaybe__500933224 gopurs_runtime.Value
var once_applicativeMaybe__500933224 sync.Once
func Get_applicativeMaybe__500933224() gopurs_runtime.Value {
	once_applicativeMaybe__500933224.Do(func() {
		cache_applicativeMaybe__500933224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__500933224
}

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__3591110311 gopurs_runtime.Value
var once_bindMaybe__3591110311 sync.Once
func Get_bindMaybe__3591110311() gopurs_runtime.Value {
	once_bindMaybe__3591110311.Do(func() {
		cache_bindMaybe__3591110311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__3591110311
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_monadMaybe__3072900051 gopurs_runtime.Value
var once_monadMaybe__3072900051 sync.Once
func Get_monadMaybe__3072900051() gopurs_runtime.Value {
	once_monadMaybe__3072900051.Do(func() {
		cache_monadMaybe__3072900051 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_bindMaybe()
}))
	})
	return cache_monadMaybe__3072900051
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_new__2045356851 gopurs_runtime.Value
var once_new__2045356851 sync.Once
func Get_new__2045356851() gopurs_runtime.Value {
	once_new__2045356851.Do(func() {
		cache_new__2045356851 = pkg_Effect_Ref.Get__new()
	})
	return cache_new__2045356851
}

var cache_applicativeEffect__284161122 gopurs_runtime.Value
var once_applicativeEffect__284161122 sync.Once
func Get_applicativeEffect__284161122() gopurs_runtime.Value {
	once_applicativeEffect__284161122.Do(func() {
		cache_applicativeEffect__284161122 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_pureE())
	})
	return cache_applicativeEffect__284161122
}

var cache_applicativeEffect__1969567048 gopurs_runtime.Value
var once_applicativeEffect__1969567048 sync.Once
func Get_applicativeEffect__1969567048() gopurs_runtime.Value {
	once_applicativeEffect__1969567048.Do(func() {
		cache_applicativeEffect__1969567048 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_pureE())
	})
	return cache_applicativeEffect__1969567048
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__2113658466 gopurs_runtime.Value
var once_bindEffect__2113658466 sync.Once
func Get_bindEffect__2113658466() gopurs_runtime.Value {
	once_bindEffect__2113658466.Do(func() {
		cache_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__2113658466
}

var cache_bindEffect__3856311079 gopurs_runtime.Value
var once_bindEffect__3856311079 sync.Once
func Get_bindEffect__3856311079() gopurs_runtime.Value {
	once_bindEffect__3856311079.Do(func() {
		cache_bindEffect__3856311079 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__3856311079
}

var cache_functorEffect__347161653 gopurs_runtime.Value
var once_functorEffect__347161653 sync.Once
func Get_functorEffect__347161653() gopurs_runtime.Value {
	once_functorEffect__347161653.Do(func() {
		cache_functorEffect__347161653 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__347161653
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

var cache_monadEffect__3527935219 gopurs_runtime.Value
var once_monadEffect__3527935219 sync.Once
func Get_monadEffect__3527935219() gopurs_runtime.Value {
	once_monadEffect__3527935219.Do(func() {
		cache_monadEffect__3527935219 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_bindEffect()
}))
	})
	return cache_monadEffect__3527935219
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__254560477 gopurs_runtime.Value
var once_unsafePartial__254560477 sync.Once
func Get_unsafePartial__254560477() gopurs_runtime.Value {
	once_unsafePartial__254560477.Do(func() {
		cache_unsafePartial__254560477 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__254560477
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

func Call_tailRec2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_4_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, "b"))
continue go__go_3_0_1
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
return gopurs_runtime.Apply(go__go_3_0_1, gopurs_runtime.Apply2(f_0, a_1, b_2))
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
v_5_loop = gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "b"), gopurs_runtime.RecordGet((*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, "c"))
continue go__go_4_0_2
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
return gopurs_runtime.Apply(go__go_4_0_2, gopurs_runtime.Apply3(f_0, a_1, b_2, c_3))
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

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bindFlipped__1485397639(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_bindFlipped__1317599105(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), a_1, b_0)
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_tailRec__2110844386(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_7
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
return gopurs_runtime.Apply(go__go_1_0_7, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec__2334182452(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_8
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
return gopurs_runtime.Apply(go__go_1_0_8, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec__2666749533(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_9
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
return gopurs_runtime.Apply(go__go_1_0_9, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec__2045907654(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_10
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
return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRec__2929877587(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
continue go__go_1_0_11
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
return gopurs_runtime.Apply(go__go_1_0_11, gopurs_runtime.Apply(f_0, x_2))
})
}

func Call_tailRecM__2220253896(dict_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3865988408(dict_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__1444729948(dict_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__2222286441(dict_0_loop *Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_either__2158544585(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_either__271265665(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapFlipped__4215217780(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__3249733428(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__1087756276(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_voidRight__1142845180(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_voidRight__698766972(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


