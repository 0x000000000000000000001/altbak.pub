package Control_Monad_Rec_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Effect "gopurs/output/Effect"
	unsafe "unsafe"
)

var Loop gopurs_runtime.Value
var once_Loop sync.Once
func Get_Loop() gopurs_runtime.Value {
	once_Loop.Do(func() {
		Loop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{value0})}
})
	})
	return Loop
}

var Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		Done = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{value0})}
})
	})
	return Done
}

var tailRecM gopurs_runtime.Value
var once_tailRecM sync.Once
func Get_tailRecM() gopurs_runtime.Value {
	once_tailRecM.Do(func() {
		tailRecM = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tailRecM")
}()
})
	})
	return tailRecM
}

var tailRecM2 gopurs_runtime.Value
var once_tailRecM2 sync.Once
func Get_tailRecM2() gopurs_runtime.Value {
	once_tailRecM2.Do(func() {
		tailRecM2 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2(dictMonadRec_0_box, f_1_box, a_2_box, b_3_box)
})
	})
	return tailRecM2
}

var tailRecM3 gopurs_runtime.Value
var once_tailRecM3 sync.Once
func Get_tailRecM3() gopurs_runtime.Value {
	once_tailRecM3.Do(func() {
		tailRecM3 = gopurs_runtime.Func5(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM3(dictMonadRec_0_box, f_1_box, a_2_box, b_3_box, c_4_box)
})
	})
	return tailRecM3
}

var untilJust gopurs_runtime.Value
var once_untilJust sync.Once
func Get_untilJust() gopurs_runtime.Value {
	once_untilJust.Do(func() {
		untilJust = gopurs_runtime.Func(func(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 42808261) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{pkg_Data_Unit.Get_unit()})}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1354639136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_4.UnsafePtr).V0})}
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
}()
})
	})
	return untilJust
}

var whileJust gopurs_runtime.Value
var once_whileJust sync.Once
func Get_whileJust() gopurs_runtime.Value {
	once_whileJust.Do(func() {
		whileJust = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(dictMonadRec_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_2, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_2, "tailRecM"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 42808261) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{v_5})}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1354639136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), v_5, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_6.UnsafePtr).V0)})}
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
}), mempty_1_0)
})
})
}()
})
	})
	return whileJust
}

var tailRec gopurs_runtime.Value
var once_tailRec sync.Once
func Get_tailRec() gopurs_runtime.Value {
	once_tailRec.Do(func() {
		tailRec = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Data_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__1_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Data_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(f_0, x_2))
})
}()
})
	})
	return tailRec
}

var tailRec2 gopurs_runtime.Value
var once_tailRec2 sync.Once
func Get_tailRec2() gopurs_runtime.Value {
	once_tailRec2.Do(func() {
		tailRec2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec2(f_0_box, a_1_box, b_2_box)
})
	})
	return tailRec2
}

var tailRec3 gopurs_runtime.Value
var once_tailRec3 sync.Once
func Get_tailRec3() gopurs_runtime.Value {
	once_tailRec3.Do(func() {
		tailRec3 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, c_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRec3(f_0_box, a_1_box, b_2_box, c_3_box)
})
	})
	return tailRec3
}

var monadRecMaybe gopurs_runtime.Value
var once_monadRecMaybe sync.Once
func Get_monadRecMaybe() gopurs_runtime.Value {
	once_monadRecMaybe.Do(func() {
		monadRecMaybe = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 42808261) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1354639136) {
var __t2 gopurs_runtime.Value
{
if ((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_2.UnsafePtr).V0.Type == 9 && (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_2.UnsafePtr).V0.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.Apply(f_0, (*Data_Control_Monad_Rec_Class_Loop)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if ((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_2.UnsafePtr).V0.Type == 9 && (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_2.UnsafePtr).V0.IntVal == 60402430) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*Data_Control_Monad_Rec_Class_Done)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
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
var go__3_3 gopurs_runtime.Value
go__3_3 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_3:
for {
if false { continue go__3_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t4 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Data_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0)
continue go__3_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t4 = (*Data_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
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
}()
})
return gopurs_runtime.Apply(go__3_3, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
}))
	})
	return monadRecMaybe
}

var monadRecIdentity gopurs_runtime.Value
var once_monadRecIdentity sync.Once
func Get_monadRecIdentity() gopurs_runtime.Value {
	once_monadRecIdentity.Do(func() {
		monadRecIdentity = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 525585346) {
v_2_loop = gopurs_runtime.Apply(f_0, (*Data_Control_Monad_Rec_Class_Loop)(v_2.UnsafePtr).V0)
continue go__1_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 60402430) {
__t1 = (*Data_Control_Monad_Rec_Class_Done)(v_2.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(f_0, x_2))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_monadIdentity()
}))
	})
	return monadRecIdentity
}

var monadRecFunction gopurs_runtime.Value
var once_monadRecFunction sync.Once
func Get_monadRecFunction() gopurs_runtime.Value {
	once_monadRecFunction.Do(func() {
		monadRecFunction = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value, e_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*Data_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0, e_2)
continue go__3_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Data_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply2(f_0, a0_1, e_2))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
}))
	})
	return monadRecFunction
}

var monadRecEither gopurs_runtime.Value
var once_monadRecEither sync.Once
func Get_monadRecEither() gopurs_runtime.Value {
	once_monadRecEither.Do(func() {
		monadRecEither = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 590902115) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 590902115, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{(*pkg_Data_Either.Data_Data_Either_Left)(v_2.UnsafePtr).V0})}})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 4096564120) {
var __t2 gopurs_runtime.Value
{
if ((*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0.Type == 9 && (*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.Apply(f_0, (*Data_Control_Monad_Rec_Class_Loop)((*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
if ((*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0.Type == 9 && (*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0.IntVal == 60402430) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 4096564120, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{(*Data_Control_Monad_Rec_Class_Done)((*pkg_Data_Either.Data_Data_Either_Right)(v_2.UnsafePtr).V0.UnsafePtr).V0})}})}
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
var go__3_3 gopurs_runtime.Value
go__3_3 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_3:
for {
if false { continue go__3_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t4 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.UncurriedApp(__local_var_2_0, (*Data_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0)
continue go__3_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t4 = (*Data_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
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
}()
})
return gopurs_runtime.Apply(go__3_3, gopurs_runtime.UncurriedApp(__local_var_2_0, gopurs_runtime.Apply(f_0, a0_1)))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_monadEither()
}))
	})
	return monadRecEither
}

var monadRecEffect gopurs_runtime.Value
var once_monadRecEffect sync.Once
func Get_monadRecEffect() gopurs_runtime.Value {
	once_monadRecEffect.Do(func() {
		monadRecEffect = gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, a_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
__local_ref_3 := __local_var_3_1
_ = __local_ref_3
r_4_2 := gopurs_runtime.Value{PtrVal: &__local_ref_3}
_ = r_4_2
_dollar__unused_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
v_5_5 := *(r_4_2.PtrVal.(*gopurs_runtime.Value))
_ = v_5_5
var __t6 gopurs_runtime.Value
{
if (v_5_5.Type == 9 && v_5_5.IntVal == 525585346) {
__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
e_6_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, (*Data_Control_Monad_Rec_Class_Loop)(v_5_5.UnsafePtr).V0), gopurs_runtime.Value{})
_ = e_6_7
*(r_4_2.PtrVal.(*gopurs_runtime.Value)) = e_6_7
_dollar__unused_7_8 := e_6_7
_ = _dollar__unused_7_8
return gopurs_runtime.Bool(false)
})
goto end_branch_6
} else {

}
}
{
if (v_5_5.Type == 9 && v_5_5.IntVal == 60402430) {
__t6 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(__t6, gopurs_runtime.Value{})
})), gopurs_runtime.Value{})
_ = _dollar__unused_5_4
a_prime_6_9 := *(r_4_2.PtrVal.(*gopurs_runtime.Value))
_ = a_prime_6_9
var __t10 gopurs_runtime.Value
{
if (a_prime_6_9.Type == 9 && a_prime_6_9.IntVal == 60402430) {
__t10 = (*Data_Control_Monad_Rec_Class_Done)(a_prime_6_9.UnsafePtr).V0
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}))
	})
	return monadRecEffect
}

var loop3 gopurs_runtime.Value
var once_loop3 sync.Once
func Get_loop3() gopurs_runtime.Value {
	once_loop3.Do(func() {
		loop3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_loop3(a_0_box, b_1_box, c_2_box)
})
	})
	return loop3
}

var loop2 gopurs_runtime.Value
var once_loop2 sync.Once
func Get_loop2() gopurs_runtime.Value {
	once_loop2.Do(func() {
		loop2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_loop2(a_0_box, b_1_box)
})
	})
	return loop2
}

var functorStep gopurs_runtime.Value
var once_functorStep sync.Once
func Get_functorStep() gopurs_runtime.Value {
	once_functorStep.Do(func() {
		functorStep = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{(*Data_Control_Monad_Rec_Class_Loop)(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Apply(f_0, (*Data_Control_Monad_Rec_Class_Done)(m_1.UnsafePtr).V0)})}
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
	return functorStep
}

var forever gopurs_runtime.Value
var once_forever sync.Once
func Get_forever() gopurs_runtime.Value {
	once_forever.Do(func() {
		forever = gopurs_runtime.Func(func(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{u_3})}
}), ma_2)
}), pkg_Data_Unit.Get_unit())
})
}()
})
	})
	return forever
}

var bifunctorStep gopurs_runtime.Value
var once_bifunctorStep sync.Once
func Get_bifunctorStep() gopurs_runtime.Value {
	once_bifunctorStep.Do(func() {
		bifunctorStep = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.Apply(v_0, (*Data_Control_Monad_Rec_Class_Loop)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Apply(v1_1, (*Data_Control_Monad_Rec_Class_Done)(v2_2.UnsafePtr).V0)})}
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
	return bifunctorStep
}

type Data_Control_Monad_Rec_Class_Loop struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Control_Monad_Rec_Class_Loop(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 525585346
}

type Data_Control_Monad_Rec_Class_Done struct {
	V0 gopurs_runtime.Value
}
func Is_Data_Control_Monad_Rec_Class_Done(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 60402430
}

func Call_tailRecM2(dictMonadRec_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_tailRecM3(dictMonadRec_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var c_4 gopurs_runtime.Value = c_4_loop
_ = c_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, gopurs_runtime.RecordGet(o_5, "a"), gopurs_runtime.RecordGet(o_5, "b"), gopurs_runtime.RecordGet(o_5, "c"))
}), gopurs_runtime.RecordDict3("a", "b", "c", a_2, b_3, c_4))
}

func Call_tailRec2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 525585346) {
v_4_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet((*Data_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Data_Control_Monad_Rec_Class_Loop)(v_4.UnsafePtr).V0, "b"))
continue go__3_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t1 = (*Data_Control_Monad_Rec_Class_Done)(v_4.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply2(f_0, a_1, b_2))
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
var go__4_0 gopurs_runtime.Value
go__4_0 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_0:
for {
if false { continue go__4_0 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 525585346) {
v_5_loop = gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet((*Data_Control_Monad_Rec_Class_Loop)(v_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Data_Control_Monad_Rec_Class_Loop)(v_5.UnsafePtr).V0, "b"), gopurs_runtime.RecordGet((*Data_Control_Monad_Rec_Class_Loop)(v_5.UnsafePtr).V0, "c"))
continue go__4_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 60402430) {
__t1 = (*Data_Control_Monad_Rec_Class_Done)(v_5.UnsafePtr).V0
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
return gopurs_runtime.Apply(go__4_0, gopurs_runtime.Apply3(f_0, a_1, b_2, c_3))
}

func Call_loop3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.RecordDict3("a", "b", "c", a_0, b_1, c_2)})}
}

func Call_loop2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.RecordDict2("a", "b", a_0, b_1)})}
}


