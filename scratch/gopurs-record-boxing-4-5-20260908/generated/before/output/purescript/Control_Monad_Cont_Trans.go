package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Cont_Trans_ContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_ContT sync.Once
func Get_Control_Monad_Cont_Trans_ContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_ContT.Do(func() {
		cache_Control_Monad_Cont_Trans_ContT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_ContT(x_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_ContT
}

var cache_Control_Monad_Cont_Trans_withContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_withContT sync.Once
func Get_Control_Monad_Cont_Trans_withContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_withContT.Do(func() {
		cache_Control_Monad_Cont_Trans_withContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_withContT(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_Control_Monad_Cont_Trans_withContT
}

var cache_Control_Monad_Cont_Trans_runContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_runContT sync.Once
func Get_Control_Monad_Cont_Trans_runContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_runContT.Do(func() {
		cache_Control_Monad_Cont_Trans_runContT = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_runContT(v_0_box, k_1_box)
})
	})
	return cache_Control_Monad_Cont_Trans_runContT
}

var cache_Control_Monad_Cont_Trans_newtypeContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_newtypeContT sync.Once
func Get_Control_Monad_Cont_Trans_newtypeContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_newtypeContT.Do(func() {
		cache_Control_Monad_Cont_Trans_newtypeContT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Monad_Cont_Trans_newtypeContT
}

var cache_Control_Monad_Cont_Trans_monadTransContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadTransContT sync.Once
func Get_Control_Monad_Cont_Trans_monadTransContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadTransContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadTransContT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_2, k_3)
})
})}))}
	})
	return cache_Control_Monad_Cont_Trans_monadTransContT
}

var cache_Control_Monad_Cont_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_lift sync.Once
func Get_Control_Monad_Cont_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_lift.Do(func() {
		cache_Control_Monad_Cont_Trans_lift = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_Cont_Trans_lift
}

var cache_Control_Monad_Cont_Trans_mapContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_mapContT sync.Once
func Get_Control_Monad_Cont_Trans_mapContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_mapContT.Do(func() {
		cache_Control_Monad_Cont_Trans_mapContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_mapContT(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_Control_Monad_Cont_Trans_mapContT
}

var cache_Control_Monad_Cont_Trans_functorContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_functorContT sync.Once
func Get_Control_Monad_Cont_Trans_functorContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_functorContT.Do(func() {
		cache_Control_Monad_Cont_Trans_functorContT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_functorContT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_functorContT
}

var cache_Control_Monad_Cont_Trans_applyContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_applyContT sync.Once
func Get_Control_Monad_Cont_Trans_applyContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_applyContT.Do(func() {
		cache_Control_Monad_Cont_Trans_applyContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_applyContT(dictApply_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_applyContT
}

var cache_Control_Monad_Cont_Trans_bindContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_bindContT sync.Once
func Get_Control_Monad_Cont_Trans_bindContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_bindContT.Do(func() {
		cache_Control_Monad_Cont_Trans_bindContT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_bindContT(dictBind_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_bindContT
}

var cache_Control_Monad_Cont_Trans_semigroupContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_semigroupContT sync.Once
func Get_Control_Monad_Cont_Trans_semigroupContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_semigroupContT.Do(func() {
		cache_Control_Monad_Cont_Trans_semigroupContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_semigroupContT(dictApply_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_semigroupContT
}

var cache_Control_Monad_Cont_Trans_applicativeContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_applicativeContT sync.Once
func Get_Control_Monad_Cont_Trans_applicativeContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_applicativeContT.Do(func() {
		cache_Control_Monad_Cont_Trans_applicativeContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_applicativeContT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_applicativeContT
}

var cache_Control_Monad_Cont_Trans_monadContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadContT sync.Once
func Get_Control_Monad_Cont_Trans_monadContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadContT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadContT
}

var cache_Control_Monad_Cont_Trans_monadAskContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadAskContT sync.Once
func Get_Control_Monad_Cont_Trans_monadAskContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadAskContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadAskContT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadAskContT(dictMonadAsk_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadAskContT
}

var cache_Control_Monad_Cont_Trans_monadReaderContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadReaderContT sync.Once
func Get_Control_Monad_Cont_Trans_monadReaderContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadReaderContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadReaderContT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadReaderContT(dictMonadReader_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadReaderContT
}

var cache_Control_Monad_Cont_Trans_monadContContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadContContT sync.Once
func Get_Control_Monad_Cont_Trans_monadContContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadContContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadContContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadContContT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadContContT
}

var cache_Control_Monad_Cont_Trans_monadEffectContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadEffectContT sync.Once
func Get_Control_Monad_Cont_Trans_monadEffectContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadEffectContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadEffectContT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadEffectContT(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadEffectContT
}

var cache_Control_Monad_Cont_Trans_monadStateContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadStateContT sync.Once
func Get_Control_Monad_Cont_Trans_monadStateContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadStateContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadStateContT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadStateContT(dictMonadState_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadStateContT
}

var cache_Control_Monad_Cont_Trans_monadSTContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monadSTContT sync.Once
func Get_Control_Monad_Cont_Trans_monadSTContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monadSTContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monadSTContT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monadSTContT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monadSTContT
}

var cache_Control_Monad_Cont_Trans_monoidContT gopurs_runtime.Value
var once_Control_Monad_Cont_Trans_monoidContT sync.Once
func Get_Control_Monad_Cont_Trans_monoidContT() gopurs_runtime.Value {
	once_Control_Monad_Cont_Trans_monoidContT.Do(func() {
		cache_Control_Monad_Cont_Trans_monoidContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Trans_monoidContT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_Cont_Trans_monoidContT
}

func Call_Control_Monad_Cont_Trans_ContT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Cont_Trans_withContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
}

func Call_Control_Monad_Cont_Trans_runContT(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_Control_Monad_Cont_Trans_lift(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_2, k_3)
})
}

func Call_Control_Monad_Cont_Trans_mapContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_Control_Monad_Cont_Trans_functorContT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})}))}
}

func Call_Control_Monad_Cont_Trans_applyContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): functorContT1_1_0 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})}))}
}

func Call_Control_Monad_Cont_Trans_bindContT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): functorContT1_1_1 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_1
// TAST (Let): applyContT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_1)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, k_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_3, a_5, k_prime__4)
}))
})}))}
}

func Call_Control_Monad_Cont_Trans_semigroupContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): functorContT1_1_1 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_1
// TAST (Let): applyContT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_1)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyContT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): __local_var_4_3 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_4_3 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyContT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), __local_var_4_3, a_5), b_6)
})}))}
})
}

func Call_Control_Monad_Cont_Trans_applicativeContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): functorContT1_1_1 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_1
// TAST (Let): applyContT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_1)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_0)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})}))}
}

func Call_Control_Monad_Cont_Trans_monadContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): functorContT1_1_2 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_2
// TAST (Let): applyContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_1 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_2)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_1
// TAST (Let): applicativeContT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_1)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_0
// TAST (Let): functorContT1_2_5 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_5
// TAST (Let): applyContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_4
// TAST (Let): bindContT1_2_3 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_2_3 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_2_4)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime__5)
}))
})})
_ = bindContT1_2_3
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_3)}
})}))}
}

func Call_Control_Monad_Cont_Trans_monadAskContT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): functorContT1_1_3 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_3
// TAST (Let): applyContT1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_3)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_2
// TAST (Let): applicativeContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_1 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_2)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_1
// TAST (Let): functorContT1_2_6 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_6
// TAST (Let): applyContT1_2_5 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_6)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_5
// TAST (Let): bindContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_2_4 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime__5)
}))
})})
_ = bindContT1_2_4
// TAST (Let): monadContT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadContT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_4)}
})})
_ = monadContT1_1_0
// TAST (Let): Bind1_2_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_7
// TAST (Let): __local_var_3_8 shape=Other bindingType=(TypeApp (TypeVar m) [(TypeVar r1)])
__local_var_3_8 := gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")
_ = __local_var_3_8
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
}), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_7.V1), __local_var_3_8, k_4)
})}))}
}

func Call_Control_Monad_Cont_Trans_monadReaderContT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): MonadAsk0_1_0 shape=App(Other) bindingType=Any
MonadAsk0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = MonadAsk0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): ask_3_2 shape=Other bindingType=(TypeApp (TypeVar m) [(TypeVar r1)])
ask_3_2 := gopurs_runtime.RecordGet(MonadAsk0_1_0, "ask")
_ = ask_3_2
// TAST (Let): functorContT1_4_7 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_4_7 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_6, gopurs_runtime.Apply(f_4, a_7))
}))
})})
_ = functorContT1_4_7
// TAST (Let): applyContT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_4_6 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_4_7)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(g_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, gopurs_runtime.Apply(g_8, a_9))
}))
}))
})})
_ = applyContT1_4_6
// TAST (Let): applicativeContT1_4_5 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_4_5 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_4_6)}
}), gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_6, a_5)
})})
_ = applicativeContT1_4_5
// TAST (Let): functorContT1_5_10 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_5_10 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, gopurs_runtime.Apply(f_5, a_8))
}))
})})
_ = functorContT1_5_10
// TAST (Let): applyContT1_5_9 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_5_9 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_5_10)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(g_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, gopurs_runtime.Apply(g_9, a_10))
}))
}))
})})
_ = applyContT1_5_9
// TAST (Let): bindContT1_5_8 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_5_8 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_5_9)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value, k_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_7, a_9, k_prime__8)
}))
})})
_ = bindContT1_5_8
// TAST (Let): monadContT1_4_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadContT1_4_4 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_4_5)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_5_8)}
})})
_ = monadContT1_4_4
// TAST (Let): Bind1_5_11 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_11
// TAST (Let): __local_var_6_12 shape=Other bindingType=(TypeApp (TypeVar m) [(TypeVar r1)])
__local_var_6_12 := gopurs_runtime.RecordGet(MonadAsk0_1_0, "ask")
_ = __local_var_6_12
// TAST (Let): monadAskContT1_4_3 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r1), (TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadAskContT1_4_3 := (&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_4_4)}
}), gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_11.V1), __local_var_6_12, k_7)
})})
_ = monadAskContT1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskContT1_4_3)}
}), gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), ask_3_2, gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_13 shape=App(Other) bindingType=Any
__local_var_9_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return r_8
}))
_ = __local_var_9_13
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_5, gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_13, gopurs_runtime.Apply(k_7, x_10))
})))
}))
})}))}
}

func Call_Control_Monad_Cont_Trans_monadContContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): functorContT1_1_3 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_3
// TAST (Let): applyContT1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_3)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_2
// TAST (Let): applicativeContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_1 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_2)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_1
// TAST (Let): functorContT1_2_6 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_6
// TAST (Let): applyContT1_2_5 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_6)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_5
// TAST (Let): bindContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_2_4 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime__5)
}))
})})
_ = bindContT1_2_4
// TAST (Let): monadContT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadContT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_4)}
})})
_ = monadContT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_4)
}), k_3)
})}))}
}

func Call_Control_Monad_Cont_Trans_monadEffectContT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): functorContT1_1_3 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_3
// TAST (Let): applyContT1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_3)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_2
// TAST (Let): applicativeContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_1 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_2)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_1
// TAST (Let): functorContT1_2_6 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_6
// TAST (Let): applyContT1_2_5 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_6)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_5
// TAST (Let): bindContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_2_4 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime__5)
}))
})})
_ = bindContT1_2_4
// TAST (Let): monadContT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadContT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_4)}
})})
_ = monadContT1_1_0
// TAST (Let): Bind1_2_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_7
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_3)
_ = __local_var_4_8
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_7.V1), __local_var_4_8, k_5)
})
})}))}
}

func Call_Control_Monad_Cont_Trans_monadStateContT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): functorContT1_1_3 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_3
// TAST (Let): applyContT1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_3)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_2
// TAST (Let): applicativeContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_1 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_2)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_1
// TAST (Let): functorContT1_2_6 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_6
// TAST (Let): applyContT1_2_5 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_6)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_5
// TAST (Let): bindContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_2_4 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime__5)
}))
})})
_ = bindContT1_2_4
// TAST (Let): monadContT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadContT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_4)}
})})
_ = monadContT1_1_0
// TAST (Let): Bind1_2_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_7
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_3)
_ = __local_var_4_8
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_7.V1), __local_var_4_8, k_5)
})
})}))}
}

func Call_Control_Monad_Cont_Trans_monadSTContT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): functorContT1_1_3 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_3
// TAST (Let): applyContT1_1_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_3)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_2
// TAST (Let): applicativeContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_1 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_2)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_1
// TAST (Let): functorContT1_2_6 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_6
// TAST (Let): applyContT1_2_5 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_6)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_5
// TAST (Let): bindContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
bindContT1_2_4 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime__5)
}))
})})
_ = bindContT1_2_4
// TAST (Let): monadContT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
monadContT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeContT1_1_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindContT1_2_4)}
})})
_ = monadContT1_1_0
// TAST (Let): Bind1_2_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_7
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadContT1_1_0)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_3)
_ = __local_var_4_8
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_7.V1), __local_var_4_8, k_5)
})
})}))}
}

func Call_Control_Monad_Cont_Trans_monoidContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): functorContT1_1_2 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_2
// TAST (Let): applyContT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_1_1 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_2)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT1_1_1
// TAST (Let): applicativeContT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applicativeContT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT1_1_1)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})})
_ = applicativeContT1_1_0
// TAST (Let): functorContT1_2_5 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_2_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})})
_ = functorContT1_2_5
// TAST (Let): applyContT1_2_4 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
applyContT1_2_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_2_5)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})})
_ = applyContT1_2_4
// TAST (Let): semigroupContT1__193435443_2_3 shape=Let(Abs(LitRecord)) bindingType=Any
semigroupContT1__193435443_2_3 := gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyContT1_2_4.V0), gopurs_runtime.Value{}))
_ = Functor0_4_6
// TAST (Let): __local_var_5_7 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_5_7 := gopurs_runtime.RecordGet(dictSemigroup_3, "append")
_ = __local_var_5_7
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyContT1_2_4.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_6.V0), __local_var_5_7, a_6), b_7)
})}))}
})
_ = semigroupContT1__193435443_2_3
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupContT2_4_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m), (TypeVar a)])])
semigroupContT2_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupContT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})))
_ = semigroupContT2_4_8
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupContT2_4_8)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeContT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))}))}
})
}


