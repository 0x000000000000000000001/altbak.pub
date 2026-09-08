package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Except_Trans_ExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_ExceptT sync.Once
func Get_Control_Monad_Except_Trans_ExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_ExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_ExceptT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_ExceptT(x_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_ExceptT
}

var cache_Control_Monad_Except_Trans_withExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_withExceptT sync.Once
func Get_Control_Monad_Except_Trans_withExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_withExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_withExceptT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_withExceptT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_Except_Trans_withExceptT
}

var cache_Control_Monad_Except_Trans_runExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_runExceptT sync.Once
func Get_Control_Monad_Except_Trans_runExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_runExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_runExceptT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_runExceptT(v_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_runExceptT
}

var cache_Control_Monad_Except_Trans_newtypeExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_newtypeExceptT sync.Once
func Get_Control_Monad_Except_Trans_newtypeExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_newtypeExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_newtypeExceptT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Monad_Except_Trans_newtypeExceptT
}

var cache_Control_Monad_Except_Trans_monadTransExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadTransExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadTransExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadTransExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadTransExceptT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})
})}))}
	})
	return cache_Control_Monad_Except_Trans_monadTransExceptT
}

var cache_Control_Monad_Except_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Except_Trans_lift sync.Once
func Get_Control_Monad_Except_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_lift.Do(func() {
		cache_Control_Monad_Except_Trans_lift = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_Except_Trans_lift
}

var cache_Control_Monad_Except_Trans_mapExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_mapExceptT sync.Once
func Get_Control_Monad_Except_Trans_mapExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_mapExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_mapExceptT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_mapExceptT(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_mapExceptT
}

var cache_Control_Monad_Except_Trans_functorExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_functorExceptT sync.Once
func Get_Control_Monad_Except_Trans_functorExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_functorExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_functorExceptT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_functorExceptT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_functorExceptT
}

var cache_Control_Monad_Except_Trans_except gopurs_runtime.Value
var once_Control_Monad_Except_Trans_except sync.Once
func Get_Control_Monad_Except_Trans_except() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_except.Do(func() {
		cache_Control_Monad_Except_Trans_except = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_except(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), x_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_except
}

var cache_Control_Monad_Except_Trans_monadExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadExceptT
}

var cache_Control_Monad_Except_Trans_bindExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_bindExceptT sync.Once
func Get_Control_Monad_Except_Trans_bindExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_bindExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_bindExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_bindExceptT
}

var cache_Control_Monad_Except_Trans_applyExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_applyExceptT sync.Once
func Get_Control_Monad_Except_Trans_applyExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_applyExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_applyExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_applyExceptT
}

var cache_Control_Monad_Except_Trans_applicativeExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_applicativeExceptT sync.Once
func Get_Control_Monad_Except_Trans_applicativeExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_applicativeExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_applicativeExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_applicativeExceptT
}

var cache_Control_Monad_Except_Trans_semigroupExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_semigroupExceptT sync.Once
func Get_Control_Monad_Except_Trans_semigroupExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_semigroupExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_semigroupExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_semigroupExceptT
}

var cache_Control_Monad_Except_Trans_monadAskExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadAskExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadAskExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadAskExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadAskExceptT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadAskExceptT(dictMonadAsk_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadAskExceptT
}

var cache_Control_Monad_Except_Trans_monadReaderExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadReaderExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadReaderExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadReaderExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadReaderExceptT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadReaderExceptT(dictMonadReader_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadReaderExceptT
}

var cache_Control_Monad_Except_Trans_monadContExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadContExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadContExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadContExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadContExceptT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadContExceptT(dictMonadCont_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadContExceptT
}

var cache_Control_Monad_Except_Trans_monadEffectExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadEffectExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadEffectExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadEffectExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadEffectExceptT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadEffectExceptT(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadEffectExceptT
}

var cache_Control_Monad_Except_Trans_monadRecExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadRecExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadRecExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadRecExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadRecExceptT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadRecExceptT(dictMonadRec_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadRecExceptT
}

var cache_Control_Monad_Except_Trans_monadStateExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadStateExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadStateExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadStateExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadStateExceptT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadStateExceptT(dictMonadState_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadStateExceptT
}

var cache_Control_Monad_Except_Trans_monadTellExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadTellExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadTellExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadTellExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadTellExceptT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadTellExceptT(dictMonadTell_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadTellExceptT
}

var cache_Control_Monad_Except_Trans_monadWriterExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadWriterExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadWriterExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadWriterExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadWriterExceptT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadWriterExceptT(dictMonadWriter_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadWriterExceptT
}

var cache_Control_Monad_Except_Trans_monadThrowExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadThrowExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadThrowExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadThrowExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadThrowExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadThrowExceptT
}

var cache_Control_Monad_Except_Trans_monadErrorExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadErrorExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadErrorExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadErrorExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadErrorExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadErrorExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadErrorExceptT
}

var cache_Control_Monad_Except_Trans_monadSTExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadSTExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadSTExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadSTExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadSTExceptT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadSTExceptT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadSTExceptT
}

var cache_Control_Monad_Except_Trans_monoidExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monoidExceptT sync.Once
func Get_Control_Monad_Except_Trans_monoidExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monoidExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monoidExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monoidExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monoidExceptT
}

var cache_Control_Monad_Except_Trans_altExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_altExceptT sync.Once
func Get_Control_Monad_Except_Trans_altExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_altExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_altExceptT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_altExceptT(dictSemigroup_0_box, dictMonad_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_altExceptT
}

var cache_Control_Monad_Except_Trans_plusExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_plusExceptT sync.Once
func Get_Control_Monad_Except_Trans_plusExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_plusExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_plusExceptT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_plusExceptT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_plusExceptT
}

var cache_Control_Monad_Except_Trans_alternativeExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_alternativeExceptT sync.Once
func Get_Control_Monad_Except_Trans_alternativeExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_alternativeExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_alternativeExceptT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_alternativeExceptT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_alternativeExceptT
}

var cache_Control_Monad_Except_Trans_monadPlusExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadPlusExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadPlusExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadPlusExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadPlusExceptT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadPlusExceptT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadPlusExceptT
}

func Call_Control_Monad_Except_Trans_ExceptT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Except_Trans_withExceptT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(f_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2)
}

func Call_Control_Monad_Except_Trans_runExceptT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_Except_Trans_lift(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})
}

func Call_Control_Monad_Except_Trans_mapExceptT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Except_Trans_functorExceptT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_2.Type == 9 && m_2.IntVal == 3711209382) {
__t1 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_1
} else {

}
}
{
if (m_2.Type == 9 && m_2.IntVal == 2465973597) {
__t1 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
})}))}
}

func Call_Control_Monad_Except_Trans_except(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), x_1)
}

func Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0)))}
})}))}
}

func Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_3, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(pure_2_1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(k_4, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)
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
})}))}
}

func Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorExceptT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (m_3.Type == 9 && m_3.IntVal == 3711209382) {
__t3 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_3
} else {

}
}
{
if (m_3.Type == 9 && m_3.IntVal == 2465973597) {
__t3 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, v_4)
})
})})
_ = functorExceptT1_1_0
// TAST (Let): __local_var_2_4 shape=App(Var) bindingType=Any
__local_var_2_4 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_2_4
// TAST (Let): Bind1_3_5 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_5
// TAST (Let): Applicative0_4_6 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_6
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_1_0)}
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), f_5, gopurs_runtime.Func(func(f_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), a_6, gopurs_runtime.Func(func(a_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_6.V1), gopurs_runtime.Apply(f_prime__7, a_prime__8))
}))
}))
})}))}
}

func Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_1, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}

func Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorExceptT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_1_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "map"), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_3.Type == 9 && m_3.IntVal == 3711209382) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
if (m_3.Type == 9 && m_3.IntVal == 2465973597) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
_ = __local_var_3_3
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, v_4)
})
})})
_ = functorExceptT1_1_1
// TAST (Let): __local_var_2_5 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_2_5 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_6 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_6
// TAST (Let): pure_4_7 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_4_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_6.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(pure_4_7, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_8
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
})}))}
})})
_ = __local_var_2_5
// TAST (Let): Bind1_3_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_5.V1), gopurs_runtime.Value{}))
_ = Bind1_3_9
// TAST (Let): Applicative0_4_10 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_5.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_10
// TAST (Let): applyExceptT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
applyExceptT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_1_1)}
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_9.V1), f_5, gopurs_runtime.Func(func(f_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_9.V1), a_6, gopurs_runtime.Func(func(a_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_10.V1), gopurs_runtime.Apply(f_prime__7, a_prime__8))
}))
}))
})})
_ = applyExceptT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_11 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_11
// TAST (Let): __local_var_4_12 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_4_12 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_12
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_11.V0), __local_var_4_12, a_5), b_6)
})}))}
})
}

func Call_Control_Monad_Except_Trans_monadAskExceptT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadExceptT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})})
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 shape=App(Var) bindingType=Any
__local_var_5_6 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_1_1)
_ = __local_var_5_6
// TAST (Let): Bind1_6_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_9
// TAST (Let): pure_4_10 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_4_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_12 shape=App(Other) bindingType=Any
__local_var_6_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_12
// TAST (Let): functorExceptT1_6_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 shape=App(Other) bindingType=Any
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_12, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
_ = __local_var_8_13
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, v_9)
})
})})
_ = functorExceptT1_6_11
// TAST (Let): __local_var_7_15 shape=App(Var) bindingType=Any
__local_var_7_15 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_1_1)
_ = __local_var_7_15
// TAST (Let): Bind1_8_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_16
// TAST (Let): Applicative0_9_17 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_11)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_17.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_9.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_4_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_18
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})}))}
})})
_ = monadExceptT1_1_0
// TAST (Let): __local_var_2_19 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
__local_var_2_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_2_19
// TAST (Let): pure_3_20 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_3_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_19.V0), gopurs_runtime.Value{}), "pure")
_ = pure_3_20
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_1_0)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_19.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_20, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))}))}
}

func Call_Control_Monad_Except_Trans_monadReaderExceptT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): monadExceptT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_2_2 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 shape=App(Other) bindingType=Any
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorExceptT1_5_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 shape=App(Other) bindingType=Any
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t7 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_7
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t7 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
})})
_ = functorExceptT1_5_4
// TAST (Let): __local_var_6_8 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_6_8 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 shape=App(Other) bindingType=Any
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorExceptT1_8_9 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_8_9 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 shape=App(Other) bindingType=Any
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t12 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_12
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t12 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_10.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
})})
_ = functorExceptT1_8_9
// TAST (Let): __local_var_9_13 shape=App(Var) bindingType=Any
__local_var_9_13 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_2_3)
_ = __local_var_9_13
// TAST (Let): Bind1_10_14 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_14
// TAST (Let): Applicative0_11_15 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_15
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_9)}
}), gopurs_runtime.Func2(func(f_12 gopurs_runtime.Value, a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), f_12, gopurs_runtime.Func(func(f_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), a_13, gopurs_runtime.Func(func(a_prime__15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_15.V1), gopurs_runtime.Apply(f_prime__14, a_prime__15))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_16
// TAST (Let): pure_8_17 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_8_17 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_17
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_19 shape=App(Other) bindingType=Any
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_19
// TAST (Let): functorExceptT1_10_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_10_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 shape=App(Other) bindingType=Any
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_19, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t21 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_21
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t21 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}))
_ = __local_var_12_20
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_20, v_13)
})
})})
_ = functorExceptT1_10_18
// TAST (Let): __local_var_11_22 shape=App(Var) bindingType=Any
__local_var_11_22 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_2_3)
_ = __local_var_11_22
// TAST (Let): Bind1_12_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_12_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_23
// TAST (Let): Applicative0_13_24 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_13_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_24
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_18)}
}), gopurs_runtime.Func2(func(f_14 gopurs_runtime.Value, a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_23.V1), f_14, gopurs_runtime.Func(func(f_prime__16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_23.V1), a_15, gopurs_runtime.Func(func(a_prime__17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_24.V1), gopurs_runtime.Apply(f_prime__16, a_prime__17))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_16.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t25 = gopurs_runtime.Apply(pure_8_17, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_25
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t25 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}))
})}))}
})})
_ = __local_var_6_8
// TAST (Let): Bind1_7_26 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_8.V1), gopurs_runtime.Value{}))
_ = Bind1_7_26
// TAST (Let): Applicative0_8_27 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_8.V0), gopurs_runtime.Value{}))
_ = Applicative0_8_27
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_4)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_27.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_28 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_28
// TAST (Let): pure_5_29 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_29 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_29
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_31 shape=App(Other) bindingType=Any
__local_var_7_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_31
// TAST (Let): functorExceptT1_7_30 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_30 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 shape=App(Other) bindingType=Any
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_31, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t33 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_33
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t33 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
}))
_ = __local_var_9_32
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_32, v_10)
})
})})
_ = functorExceptT1_7_30
// TAST (Let): __local_var_8_34 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_8_34 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_36 shape=App(Other) bindingType=Any
__local_var_10_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_36
// TAST (Let): functorExceptT1_10_35 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_10_35 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_37 shape=App(Other) bindingType=Any
__local_var_12_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_36, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t38 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_38
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t38 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}))
_ = __local_var_12_37
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_37, v_13)
})
})})
_ = functorExceptT1_10_35
// TAST (Let): __local_var_11_39 shape=App(Var) bindingType=Any
__local_var_11_39 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_2_3)
_ = __local_var_11_39
// TAST (Let): Bind1_12_40 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_40
// TAST (Let): Applicative0_13_41 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_13_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_41
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_35)}
}), gopurs_runtime.Func2(func(f_14 gopurs_runtime.Value, a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), f_14, gopurs_runtime.Func(func(f_prime__16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), a_15, gopurs_runtime.Func(func(a_prime__17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_41.V1), gopurs_runtime.Apply(f_prime__16, a_prime__17))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_9, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_42 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_42
// TAST (Let): pure_10_43 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_10_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_43
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_45 shape=App(Other) bindingType=Any
__local_var_12_45 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_45
// TAST (Let): functorExceptT1_12_44 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_12_44 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_46 shape=App(Other) bindingType=Any
__local_var_14_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_45, "map"), gopurs_runtime.Func(func(m_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 gopurs_runtime.Value
{
if (m_14.Type == 9 && m_14.IntVal == 3711209382) {
__t47 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_14.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_47
} else {

}
}
{
if (m_14.Type == 9 && m_14.IntVal == 2465973597) {
__t47 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_13, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_14.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}))
_ = __local_var_14_46
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_46, v_15)
})
})})
_ = functorExceptT1_12_44
// TAST (Let): __local_var_13_48 shape=App(Var) bindingType=Any
__local_var_13_48 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_2_3)
_ = __local_var_13_48
// TAST (Let): Bind1_14_49 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_14_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_48, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_49
// TAST (Let): Applicative0_15_50 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_15_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_48, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_50
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_12_44)}
}), gopurs_runtime.Func2(func(f_16 gopurs_runtime.Value, a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_49.V1), f_16, gopurs_runtime.Func(func(f_prime__18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_49.V1), a_17, gopurs_runtime.Func(func(a_prime__19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_50.V1), gopurs_runtime.Apply(f_prime__18, a_prime__19))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_11 gopurs_runtime.Value, k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_42.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t51 = gopurs_runtime.Apply(pure_10_43, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_51
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t51 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0)
goto end_branch_51
} else {

}
}
{
__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_51:
return __t51
}))
})}))}
})})
_ = __local_var_8_34
// TAST (Let): Bind1_9_52 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_34.V1), gopurs_runtime.Value{}))
_ = Bind1_9_52
// TAST (Let): Applicative0_10_53 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_34.V0), gopurs_runtime.Value{}))
_ = Applicative0_10_53
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_30)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_52.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_52.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_53.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_28.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t54 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t54 = gopurs_runtime.Apply(pure_5_29, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_54
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t54 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
return __t54
}))
})}))}
})})
_ = monadExceptT1_2_2
// TAST (Let): __local_var_3_55 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
__local_var_3_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_55
// TAST (Let): pure_4_56 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_4_56 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_55.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_56
// TAST (Let): monadAskExceptT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r), (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadAskExceptT1_1_0 := (&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_2)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_55.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_56, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_5, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))})
_ = monadAskExceptT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskExceptT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_57 shape=App(Other) bindingType=Any
__local_var_3_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_57
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_57, v_4)
})
})}))}
}

func Call_Control_Monad_Except_Trans_monadContExceptT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadExceptT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})})
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 shape=App(Var) bindingType=Any
__local_var_5_6 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_1_1)
_ = __local_var_5_6
// TAST (Let): Bind1_6_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_9
// TAST (Let): pure_4_10 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_4_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_12 shape=App(Other) bindingType=Any
__local_var_6_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_12
// TAST (Let): functorExceptT1_6_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 shape=App(Other) bindingType=Any
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_12, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
_ = __local_var_8_13
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, v_9)
})
})})
_ = functorExceptT1_6_11
// TAST (Let): __local_var_7_15 shape=App(Var) bindingType=Any
__local_var_7_15 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_1_1)
_ = __local_var_7_15
// TAST (Let): Bind1_8_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_16
// TAST (Let): Applicative0_9_17 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_11)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_17.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_9.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_4_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_18
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})}))}
})})
_ = monadExceptT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_3, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
}))
})}))}
}

func Call_Control_Monad_Except_Trans_monadEffectExceptT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadExceptT1_2_1 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})})
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 shape=App(Var) bindingType=Any
__local_var_5_6 := Call_Control_Monad_Except_Trans_monadExceptT(Monad0_1_0)
_ = __local_var_5_6
// TAST (Let): Bind1_6_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_9
// TAST (Let): pure_4_10 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_4_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_12 shape=App(Other) bindingType=Any
__local_var_6_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_12
// TAST (Let): functorExceptT1_6_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 shape=App(Other) bindingType=Any
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_12, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
_ = __local_var_8_13
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, v_9)
})
})})
_ = functorExceptT1_6_11
// TAST (Let): __local_var_7_15 shape=App(Var) bindingType=Any
__local_var_7_15 := Call_Control_Monad_Except_Trans_monadExceptT(Monad0_1_0)
_ = __local_var_7_15
// TAST (Let): Bind1_8_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_16
// TAST (Let): Applicative0_9_17 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_11)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_17.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_9.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_4_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_18
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})}))}
})})
_ = monadExceptT1_2_1
// TAST (Let): Bind1_3_19 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_19
// TAST (Let): pure_4_20 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_4_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_20
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_1)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_19.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_5), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_20, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_6, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})}))}
}

func Call_Control_Monad_Except_Trans_monadRecExceptT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
// TAST (Let): monadExceptT1_4_3 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_4_3 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 shape=App(Other) bindingType=Any
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorExceptT1_6_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 shape=App(Other) bindingType=Any
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t7 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_7
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t7 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
_ = __local_var_8_6
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_6, v_9)
})
})})
_ = functorExceptT1_6_4
// TAST (Let): __local_var_7_8 shape=App(Var) bindingType=Any
__local_var_7_8 := Call_Control_Monad_Except_Trans_monadExceptT(Monad0_1_0)
_ = __local_var_7_8
// TAST (Let): Bind1_8_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_9
// TAST (Let): Applicative0_9_10 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_10
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_4)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_9.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_9.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_10.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_5, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_11 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_11
// TAST (Let): pure_6_12 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_6_12 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_12
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_14 shape=App(Other) bindingType=Any
__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_14
// TAST (Let): functorExceptT1_8_13 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_8_13 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_15 shape=App(Other) bindingType=Any
__local_var_10_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_14, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t16 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_16
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t16 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_10.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}))
_ = __local_var_10_15
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_15, v_11)
})
})})
_ = functorExceptT1_8_13
// TAST (Let): __local_var_9_17 shape=App(Var) bindingType=Any
__local_var_9_17 := Call_Control_Monad_Except_Trans_monadExceptT(Monad0_1_0)
_ = __local_var_9_17
// TAST (Let): Bind1_10_18 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_18
// TAST (Let): Applicative0_11_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_11_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_13)}
}), gopurs_runtime.Func2(func(f_12 gopurs_runtime.Value, a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_18.V1), f_12, gopurs_runtime.Func(func(f_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_18.V1), a_13, gopurs_runtime.Func(func(a_prime__15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_19.V1), gopurs_runtime.Apply(f_prime__14, a_prime__15))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_11.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_6_12, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_20
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}))
})}))}
})})
_ = monadExceptT1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_4_3)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_21 shape=App(Other) bindingType=(Func [(TypeVar a)] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
__local_var_6_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 gopurs_runtime.Value
{
if (m_prime__7.Type == 9 && m_prime__7.IntVal == 3711209382) {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__7.UnsafePtr).V0}))}}))}
goto end_branch_25
} else {

}
}
{
if (m_prime__7.Type == 9 && m_prime__7.IntVal == 2465973597) {
var __t24 gopurs_runtime.Value
{
var __t_tag_22 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__7.UnsafePtr).V0
if (__t_tag_22.Type == 9 && __t_tag_22.IntVal == 525585346) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__7.UnsafePtr).V0.UnsafePtr).V0}))}
goto end_branch_24
} else {

}
}
{
var __t_tag_23 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__7.UnsafePtr).V0
if (__t_tag_23.Type == 9 && __t_tag_23.IntVal == 60402430) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__7.UnsafePtr).V0.UnsafePtr).V0}))}}))}
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t25)
}))
}))
_ = __local_var_6_21
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_21, x_7)
})
})}))}
}

func Call_Control_Monad_Except_Trans_monadStateExceptT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadExceptT1_2_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 shape=App(Other) bindingType=Any
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t6 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_6
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t6 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
})})
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_7 shape=App(Var) bindingType=Any
__local_var_6_7 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_2_2)
_ = __local_var_6_7
// TAST (Let): Bind1_7_8 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_8
// TAST (Let): Applicative0_8_9 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_3)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_9.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_10 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_10
// TAST (Let): pure_5_11 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_11 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_11
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_13 shape=App(Other) bindingType=Any
__local_var_7_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_13
// TAST (Let): functorExceptT1_7_12 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_12 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_14 shape=App(Other) bindingType=Any
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_13, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t15 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_15
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t15 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}))
_ = __local_var_9_14
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_14, v_10)
})
})})
_ = functorExceptT1_7_12
// TAST (Let): __local_var_8_16 shape=App(Var) bindingType=Any
__local_var_8_16 := Call_Control_Monad_Except_Trans_monadExceptT(__local_var_2_2)
_ = __local_var_8_16
// TAST (Let): Bind1_9_17 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_17
// TAST (Let): Applicative0_10_18 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_12)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_18.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_10.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_5_11, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_19
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}))
})}))}
})})
_ = monadExceptT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_1)}
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_20 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_4_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_20
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_20, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_5, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})}))}
}

func Call_Control_Monad_Except_Trans_monadTellExceptT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
// TAST (Let): monadExceptT1_3_2 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_3_2 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 shape=App(Other) bindingType=Any
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t6 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_6
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t6 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
})})
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_7 shape=App(Var) bindingType=Any
__local_var_6_7 := Call_Control_Monad_Except_Trans_monadExceptT(Monad1_1_0)
_ = __local_var_6_7
// TAST (Let): Bind1_7_8 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_8
// TAST (Let): Applicative0_8_9 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_3)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_9.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_10 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_10
// TAST (Let): pure_5_11 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_11 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_11
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_13 shape=App(Other) bindingType=Any
__local_var_7_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_13
// TAST (Let): functorExceptT1_7_12 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_12 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_14 shape=App(Other) bindingType=Any
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_13, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t15 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_15
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t15 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}))
_ = __local_var_9_14
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_14, v_10)
})
})})
_ = functorExceptT1_7_12
// TAST (Let): __local_var_8_16 shape=App(Var) bindingType=Any
__local_var_8_16 := Call_Control_Monad_Except_Trans_monadExceptT(Monad1_1_0)
_ = __local_var_8_16
// TAST (Let): Bind1_9_17 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_17
// TAST (Let): Applicative0_10_18 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_12)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_18.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_10.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_5_11, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_19
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}))
})}))}
})})
_ = monadExceptT1_3_2
// TAST (Let): Bind1_4_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_20
// TAST (Let): pure_5_21 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_5_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_21
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_3_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_2_1)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_20.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_21, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})}))}
}

func Call_Control_Monad_Except_Trans_monadWriterExceptT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): MonadTell1_1_0 shape=App(Other) bindingType=Any
MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
_ = MonadTell1_1_0
// TAST (Let): Monad1_2_1 shape=App(Other) bindingType=Any
Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_2_1
// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): pure_4_3 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])]))
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
// TAST (Let): Applicative0_5_4 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
// TAST (Let): Monoid0_6_5 shape=App(Other) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar w)])
Monoid0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{}))
_ = Monoid0_6_5
// TAST (Let): Monad1_7_7 shape=App(Other) bindingType=Any
Monad1_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_7_7
// TAST (Let): Semigroup0_8_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_8
// TAST (Let): monadExceptT1_9_9 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_9_9 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 shape=App(Other) bindingType=Any
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorExceptT1_11_10 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_11_10 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_12 shape=App(Other) bindingType=Any
__local_var_13_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t13 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_13
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t13 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_13.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
_ = __local_var_13_12
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_12, v_14)
})
})})
_ = functorExceptT1_11_10
// TAST (Let): __local_var_12_14 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_12_14 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_16 shape=App(Other) bindingType=Any
__local_var_14_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_16
// TAST (Let): functorExceptT1_14_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_14_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_17 shape=App(Other) bindingType=Any
__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_16, "map"), gopurs_runtime.Func(func(m_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (m_16.Type == 9 && m_16.IntVal == 3711209382) {
__t18 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_16.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_18
} else {

}
}
{
if (m_16.Type == 9 && m_16.IntVal == 2465973597) {
__t18 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_15, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_16.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
_ = __local_var_16_17
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_17, v_17)
})
})})
_ = functorExceptT1_14_15
// TAST (Let): __local_var_15_19 shape=App(Var) bindingType=Any
__local_var_15_19 := Call_Control_Monad_Except_Trans_monadExceptT(Monad1_7_7)
_ = __local_var_15_19
// TAST (Let): Bind1_16_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_16_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_20
// TAST (Let): Applicative0_17_21 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_17_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_14_15)}
}), gopurs_runtime.Func2(func(f_18 gopurs_runtime.Value, a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_20.V1), f_18, gopurs_runtime.Func(func(f_prime__20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_20.V1), a_19, gopurs_runtime.Func(func(a_prime__21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_21.V1), gopurs_runtime.Apply(f_prime__20, a_prime__21))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_13, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_22 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_13_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_22
// TAST (Let): pure_14_23 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_14_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_23
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_25 shape=App(Other) bindingType=Any
__local_var_16_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_25
// TAST (Let): functorExceptT1_16_24 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_16_24 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_26 shape=App(Other) bindingType=Any
__local_var_18_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_25, "map"), gopurs_runtime.Func(func(m_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
if (m_18.Type == 9 && m_18.IntVal == 3711209382) {
__t27 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_18.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_27
} else {

}
}
{
if (m_18.Type == 9 && m_18.IntVal == 2465973597) {
__t27 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_17, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_18.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return __t27
}))
_ = __local_var_18_26
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_26, v_19)
})
})})
_ = functorExceptT1_16_24
// TAST (Let): __local_var_17_28 shape=App(Var) bindingType=Any
__local_var_17_28 := Call_Control_Monad_Except_Trans_monadExceptT(Monad1_7_7)
_ = __local_var_17_28
// TAST (Let): Bind1_18_29 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_18_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_29
// TAST (Let): Applicative0_19_30 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_19_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_30
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_16_24)}
}), gopurs_runtime.Func2(func(f_20 gopurs_runtime.Value, a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_29.V1), f_20, gopurs_runtime.Func(func(f_prime__22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_29.V1), a_21, gopurs_runtime.Func(func(a_prime__23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_30.V1), gopurs_runtime.Apply(f_prime__22, a_prime__23))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_15 gopurs_runtime.Value, k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_22.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t31 = gopurs_runtime.Apply(pure_14_23, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_17.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_31
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t31 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_17.UnsafePtr).V0)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
}))
})}))}
})})
_ = __local_var_12_14
// TAST (Let): Bind1_13_32 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_13_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_12_14.V1), gopurs_runtime.Value{}))
_ = Bind1_13_32
// TAST (Let): Applicative0_14_33 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_14_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_12_14.V0), gopurs_runtime.Value{}))
_ = Applicative0_14_33
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_10)}
}), gopurs_runtime.Func2(func(f_15 gopurs_runtime.Value, a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_32.V1), f_15, gopurs_runtime.Func(func(f_prime__17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_32.V1), a_16, gopurs_runtime.Func(func(a_prime__18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_33.V1), gopurs_runtime.Apply(f_prime__17, a_prime__18))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_10, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_34 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_10_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_34
// TAST (Let): pure_11_35 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_11_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_35
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_37 shape=App(Other) bindingType=Any
__local_var_13_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_37
// TAST (Let): functorExceptT1_13_36 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_13_36 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_38 shape=App(Other) bindingType=Any
__local_var_15_38 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_37, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t39 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_15.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_39
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t39 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_15.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_39
} else {

}
}
{
__t39 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_39:
return __t39
}))
_ = __local_var_15_38
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_38, v_16)
})
})})
_ = functorExceptT1_13_36
// TAST (Let): __local_var_14_40 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_14_40 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_42 shape=App(Other) bindingType=Any
__local_var_16_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_42
// TAST (Let): functorExceptT1_16_41 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_16_41 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_43 shape=App(Other) bindingType=Any
__local_var_18_43 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_42, "map"), gopurs_runtime.Func(func(m_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (m_18.Type == 9 && m_18.IntVal == 3711209382) {
__t44 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_18.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_44
} else {

}
}
{
if (m_18.Type == 9 && m_18.IntVal == 2465973597) {
__t44 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_17, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_18.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_44
} else {

}
}
{
__t44 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_44:
return __t44
}))
_ = __local_var_18_43
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_43, v_19)
})
})})
_ = functorExceptT1_16_41
// TAST (Let): __local_var_17_45 shape=App(Var) bindingType=Any
__local_var_17_45 := Call_Control_Monad_Except_Trans_monadExceptT(Monad1_7_7)
_ = __local_var_17_45
// TAST (Let): Bind1_18_46 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_18_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_45, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_46
// TAST (Let): Applicative0_19_47 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_19_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_45, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_47
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_16_41)}
}), gopurs_runtime.Func2(func(f_20 gopurs_runtime.Value, a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_46.V1), f_20, gopurs_runtime.Func(func(f_prime__22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_46.V1), a_21, gopurs_runtime.Func(func(a_prime__23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_47.V1), gopurs_runtime.Apply(f_prime__22, a_prime__23))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_15, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_48 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_15_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_48
// TAST (Let): pure_16_49 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_16_49 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_49
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_51 shape=App(Other) bindingType=Any
__local_var_18_51 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_51
// TAST (Let): functorExceptT1_18_50 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_18_50 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_52 shape=App(Other) bindingType=Any
__local_var_20_52 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_51, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t53 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t53 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_20.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_53
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t53 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_20.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_53
} else {

}
}
{
__t53 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_53:
return __t53
}))
_ = __local_var_20_52
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_52, v_21)
})
})})
_ = functorExceptT1_18_50
// TAST (Let): __local_var_19_54 shape=App(Var) bindingType=Any
__local_var_19_54 := Call_Control_Monad_Except_Trans_monadExceptT(Monad1_7_7)
_ = __local_var_19_54
// TAST (Let): Bind1_20_55 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_20_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_54, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_55
// TAST (Let): Applicative0_21_56 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_21_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_54, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_56
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_50)}
}), gopurs_runtime.Func2(func(f_22 gopurs_runtime.Value, a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_55.V1), f_22, gopurs_runtime.Func(func(f_prime__24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_55.V1), a_23, gopurs_runtime.Func(func(a_prime__25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_56.V1), gopurs_runtime.Apply(f_prime__24, a_prime__25))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_17 gopurs_runtime.Value, k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_48.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t57 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t57 = gopurs_runtime.Apply(pure_16_49, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_19.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_57
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t57 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_19.UnsafePtr).V0)
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
return __t57
}))
})}))}
})})
_ = __local_var_14_40
// TAST (Let): Bind1_15_58 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_15_58 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_14_40.V1), gopurs_runtime.Value{}))
_ = Bind1_15_58
// TAST (Let): Applicative0_16_59 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_16_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_14_40.V0), gopurs_runtime.Value{}))
_ = Applicative0_16_59
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_36)}
}), gopurs_runtime.Func2(func(f_17 gopurs_runtime.Value, a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_58.V1), f_17, gopurs_runtime.Func(func(f_prime__19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_58.V1), a_18, gopurs_runtime.Func(func(a_prime__20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_59.V1), gopurs_runtime.Apply(f_prime__19, a_prime__20))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_12 gopurs_runtime.Value, k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t60 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t60 = gopurs_runtime.Apply(pure_11_35, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_60
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t60 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0)
goto end_branch_60
} else {

}
}
{
__t60 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_60:
return __t60
}))
})}))}
})})
_ = monadExceptT1_9_9
// TAST (Let): Bind1_10_61 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_10_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_61
// TAST (Let): pure_11_62 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_11_62 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_62
// TAST (Let): monadTellExceptT1_7_6 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w), (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadTellExceptT1_7_6 := (&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_9_9)}
}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_8_8)}
}), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_61.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "tell"), x_12), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_11_62, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_13, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})})
_ = monadTellExceptT1_7_6
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellExceptT1_7_6)}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Monoid0_6_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t65 gopurs_runtime.Value
{
var __t_tag_63 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
if (__t_tag_63.Type == 9 && __t_tag_63.IntVal == 3711209382) {
__t65 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_65
} else {

}
}
{
var __t_tag_64 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0
if (__t_tag_64.Type == 9 && __t_tag_64.IntVal == 2465973597) {
__t65 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_65
} else {

}
}
{
__t65 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_65:
return gopurs_runtime.Apply(pure_4_3, __t65)
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (a_9.Type == 9 && a_9.IntVal == 3711209382) {
__t66 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(a_9.UnsafePtr).V0}))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_66
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 2465973597) {
__t66 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(a_9.UnsafePtr).V0.UnsafePtr).V0}))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(a_9.UnsafePtr).V0.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_66
} else {

}
}
{
__t66 = func() *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_66:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t66)})
})))
})}))}
}

func Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadExceptT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorExceptT1_3_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_3_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=Any
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_5.Type == 9 && m_5.IntVal == 3711209382) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
if (m_5.Type == 9 && m_5.IntVal == 2465973597) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_4, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_5.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, v_6)
})
})})
_ = functorExceptT1_3_1
// TAST (Let): __local_var_4_5 shape=App(Var) bindingType=Any
__local_var_4_5 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_4_5
// TAST (Let): Bind1_5_6 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_6
// TAST (Let): Applicative0_6_7 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_3_1)}
}), gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), f_7, gopurs_runtime.Func(func(f_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), a_8, gopurs_runtime.Func(func(a_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_7.V1), gopurs_runtime.Apply(f_prime__9, a_prime__10))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_2, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_8 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_8
// TAST (Let): pure_3_9 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_3_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_11 shape=App(Other) bindingType=Any
__local_var_5_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_11
// TAST (Let): functorExceptT1_5_10 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_10 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_12 shape=App(Other) bindingType=Any
__local_var_7_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_11, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t13 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_13
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t13 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
_ = __local_var_7_12
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_12, v_8)
})
})})
_ = functorExceptT1_5_10
// TAST (Let): __local_var_6_14 shape=App(Var) bindingType=Any
__local_var_6_14 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_6_14
// TAST (Let): Bind1_7_15 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_15
// TAST (Let): Applicative0_8_16 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_16
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_10)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_15.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_15.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_16.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_8.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_3_9, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_17
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}))
})}))}
})})
_ = monadExceptT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_2, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}

func Call_Control_Monad_Except_Trans_monadErrorExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): monadExceptT1_3_3 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_3_3 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 shape=App(Other) bindingType=Any
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorExceptT1_5_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 shape=App(Other) bindingType=Any
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t7 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_7
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t7 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
})})
_ = functorExceptT1_5_4
// TAST (Let): __local_var_6_8 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_6_8 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 shape=App(Other) bindingType=Any
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorExceptT1_8_9 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_8_9 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 shape=App(Other) bindingType=Any
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t12 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_12
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t12 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_10.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
})})
_ = functorExceptT1_8_9
// TAST (Let): __local_var_9_13 shape=App(Var) bindingType=Any
__local_var_9_13 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_9_13
// TAST (Let): Bind1_10_14 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_14
// TAST (Let): Applicative0_11_15 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_15
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_9)}
}), gopurs_runtime.Func2(func(f_12 gopurs_runtime.Value, a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), f_12, gopurs_runtime.Func(func(f_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), a_13, gopurs_runtime.Func(func(a_prime__15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_15.V1), gopurs_runtime.Apply(f_prime__14, a_prime__15))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_16
// TAST (Let): pure_8_17 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_8_17 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_17
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_19 shape=App(Other) bindingType=Any
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_19
// TAST (Let): functorExceptT1_10_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_10_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 shape=App(Other) bindingType=Any
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_19, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t21 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_21
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t21 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}))
_ = __local_var_12_20
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_20, v_13)
})
})})
_ = functorExceptT1_10_18
// TAST (Let): __local_var_11_22 shape=App(Var) bindingType=Any
__local_var_11_22 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_11_22
// TAST (Let): Bind1_12_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_12_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_23
// TAST (Let): Applicative0_13_24 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_13_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_24
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_18)}
}), gopurs_runtime.Func2(func(f_14 gopurs_runtime.Value, a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_23.V1), f_14, gopurs_runtime.Func(func(f_prime__16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_23.V1), a_15, gopurs_runtime.Func(func(a_prime__17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_24.V1), gopurs_runtime.Apply(f_prime__16, a_prime__17))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_16.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t25 = gopurs_runtime.Apply(pure_8_17, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_25
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t25 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}))
})}))}
})})
_ = __local_var_6_8
// TAST (Let): Bind1_7_26 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_8.V1), gopurs_runtime.Value{}))
_ = Bind1_7_26
// TAST (Let): Applicative0_8_27 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_8.V0), gopurs_runtime.Value{}))
_ = Applicative0_8_27
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_4)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_27.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_28 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_28
// TAST (Let): pure_5_29 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_29 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_29
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_31 shape=App(Other) bindingType=Any
__local_var_7_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_31
// TAST (Let): functorExceptT1_7_30 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_30 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 shape=App(Other) bindingType=Any
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_31, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t33 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_33
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t33 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
}))
_ = __local_var_9_32
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_32, v_10)
})
})})
_ = functorExceptT1_7_30
// TAST (Let): __local_var_8_34 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_8_34 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_36 shape=App(Other) bindingType=Any
__local_var_10_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_36
// TAST (Let): functorExceptT1_10_35 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_10_35 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_37 shape=App(Other) bindingType=Any
__local_var_12_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_36, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t38 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_38
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t38 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_12.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}))
_ = __local_var_12_37
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_37, v_13)
})
})})
_ = functorExceptT1_10_35
// TAST (Let): __local_var_11_39 shape=App(Var) bindingType=Any
__local_var_11_39 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_11_39
// TAST (Let): Bind1_12_40 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_40
// TAST (Let): Applicative0_13_41 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_13_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_41
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_35)}
}), gopurs_runtime.Func2(func(f_14 gopurs_runtime.Value, a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), f_14, gopurs_runtime.Func(func(f_prime__16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), a_15, gopurs_runtime.Func(func(a_prime__17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_41.V1), gopurs_runtime.Apply(f_prime__16, a_prime__17))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_9, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_42 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_42
// TAST (Let): pure_10_43 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_10_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_43
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_45 shape=App(Other) bindingType=Any
__local_var_12_45 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_45
// TAST (Let): functorExceptT1_12_44 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_12_44 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_46 shape=App(Other) bindingType=Any
__local_var_14_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_45, "map"), gopurs_runtime.Func(func(m_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 gopurs_runtime.Value
{
if (m_14.Type == 9 && m_14.IntVal == 3711209382) {
__t47 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_14.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_47
} else {

}
}
{
if (m_14.Type == 9 && m_14.IntVal == 2465973597) {
__t47 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_13, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_14.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}))
_ = __local_var_14_46
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_46, v_15)
})
})})
_ = functorExceptT1_12_44
// TAST (Let): __local_var_13_48 shape=App(Var) bindingType=Any
__local_var_13_48 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0)
_ = __local_var_13_48
// TAST (Let): Bind1_14_49 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_14_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_48, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_49
// TAST (Let): Applicative0_15_50 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_15_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_48, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_50
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_12_44)}
}), gopurs_runtime.Func2(func(f_16 gopurs_runtime.Value, a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_49.V1), f_16, gopurs_runtime.Func(func(f_prime__18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_49.V1), a_17, gopurs_runtime.Func(func(a_prime__19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_50.V1), gopurs_runtime.Apply(f_prime__18, a_prime__19))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_11 gopurs_runtime.Value, k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_42.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t51 = gopurs_runtime.Apply(pure_10_43, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_51
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t51 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0)
goto end_branch_51
} else {

}
}
{
__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_51:
return __t51
}))
})}))}
})})
_ = __local_var_8_34
// TAST (Let): Bind1_9_52 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_34.V1), gopurs_runtime.Value{}))
_ = Bind1_9_52
// TAST (Let): Applicative0_10_53 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_34.V0), gopurs_runtime.Value{}))
_ = Applicative0_10_53
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_30)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_52.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_52.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_53.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_28.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t54 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t54 = gopurs_runtime.Apply(pure_5_29, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_54
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t54 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
return __t54
}))
})}))}
})})
_ = monadExceptT1_3_3
// TAST (Let): monadThrowExceptT1_3_2 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e), (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadThrowExceptT1_3_2 := (&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_3_3)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_4, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})})
_ = monadThrowExceptT1_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowExceptT1_3_2)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t55 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_55
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t55 = gopurs_runtime.Apply(pure_2_1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_55
} else {

}
}
{
__t55 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_55:
return __t55
}))
})}))}
}

func Call_Control_Monad_Except_Trans_monadSTExceptT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadExceptT1_2_1 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})})
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 shape=App(Var) bindingType=Any
__local_var_5_6 := Call_Control_Monad_Except_Trans_monadExceptT(Monad0_1_0)
_ = __local_var_5_6
// TAST (Let): Bind1_6_7 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_9
// TAST (Let): pure_4_10 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_4_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_12 shape=App(Other) bindingType=Any
__local_var_6_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_12
// TAST (Let): functorExceptT1_6_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 shape=App(Other) bindingType=Any
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_12, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t14 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
_ = __local_var_8_13
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, v_9)
})
})})
_ = functorExceptT1_6_11
// TAST (Let): __local_var_7_15 shape=App(Var) bindingType=Any
__local_var_7_15 := Call_Control_Monad_Except_Trans_monadExceptT(Monad0_1_0)
_ = __local_var_7_15
// TAST (Let): Bind1_8_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_16
// TAST (Let): Applicative0_9_17 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_11)}
}), gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), f_10, gopurs_runtime.Func(func(f_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), a_11, gopurs_runtime.Func(func(a_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_17.V1), gopurs_runtime.Apply(f_prime__12, a_prime__13))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_9.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_4_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_18
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})}))}
})})
_ = monadExceptT1_2_1
// TAST (Let): Bind1_3_19 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_19
// TAST (Let): pure_4_20 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])]))
pure_4_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_20
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_1)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_19.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_5), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_20, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, a_6, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
})}))}
}

func Call_Control_Monad_Except_Trans_monoidExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeExceptT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
applicativeExceptT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorExceptT1_2_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_2_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_4.Type == 9 && m_4.IntVal == 3711209382) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
if (m_4.Type == 9 && m_4.IntVal == 2465973597) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
_ = __local_var_4_3
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, v_5)
})
})})
_ = functorExceptT1_2_1
// TAST (Let): __local_var_3_5 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_3_5 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_6 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_6
// TAST (Let): pure_5_7 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_6.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(pure_5_7, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_8
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
})}))}
})})
_ = __local_var_3_5
// TAST (Let): Bind1_4_9 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_5.V1), gopurs_runtime.Value{}))
_ = Bind1_4_9
// TAST (Let): Applicative0_5_10 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_5.V0), gopurs_runtime.Value{}))
_ = Applicative0_5_10
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_2_1)}
}), gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_9.V1), f_6, gopurs_runtime.Func(func(f_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_9.V1), a_7, gopurs_runtime.Func(func(a_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_10.V1), gopurs_runtime.Apply(f_prime__8, a_prime__9))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_1, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})})
_ = applicativeExceptT1_1_0
// TAST (Let): __local_var_2_14 shape=App(Other) bindingType=Any
__local_var_2_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_14
// TAST (Let): functorExceptT1_2_13 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_2_13 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 shape=App(Other) bindingType=Any
__local_var_4_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_14, "map"), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (m_4.Type == 9 && m_4.IntVal == 3711209382) {
__t16 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_16
} else {

}
}
{
if (m_4.Type == 9 && m_4.IntVal == 2465973597) {
__t16 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}))
_ = __local_var_4_15
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_15, v_5)
})
})})
_ = functorExceptT1_2_13
// TAST (Let): __local_var_3_17 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_3_17 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_19 shape=App(Other) bindingType=Any
__local_var_5_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_19
// TAST (Let): functorExceptT1_5_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_20 shape=App(Other) bindingType=Any
__local_var_7_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_19, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t21 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_21
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t21 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}))
_ = __local_var_7_20
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_20, v_8)
})
})})
_ = functorExceptT1_5_18
// TAST (Let): __local_var_6_22 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_6_22 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_23
// TAST (Let): pure_8_24 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_8_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_23.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t25 = gopurs_runtime.Apply(pure_8_24, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_25
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t25 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}))
})}))}
})})
_ = __local_var_6_22
// TAST (Let): Bind1_7_26 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_22.V1), gopurs_runtime.Value{}))
_ = Bind1_7_26
// TAST (Let): Applicative0_8_27 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_22.V0), gopurs_runtime.Value{}))
_ = Applicative0_8_27
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_18)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_27.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_28 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_28
// TAST (Let): pure_5_29 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_29 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_29
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_31 shape=App(Other) bindingType=Any
__local_var_7_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_31
// TAST (Let): functorExceptT1_7_30 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_30 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 shape=App(Other) bindingType=Any
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_31, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t33 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_33
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t33 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
}))
_ = __local_var_9_32
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_32, v_10)
})
})})
_ = functorExceptT1_7_30
// TAST (Let): __local_var_8_34 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_8_34 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_9, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_35 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_35
// TAST (Let): pure_10_36 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_10_36 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_36
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func2(func(v_11 gopurs_runtime.Value, k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_35.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t37 = gopurs_runtime.Apply(pure_10_36, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_37
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t37 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0)
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return __t37
}))
})}))}
})})
_ = __local_var_8_34
// TAST (Let): Bind1_9_38 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_34.V1), gopurs_runtime.Value{}))
_ = Bind1_9_38
// TAST (Let): Applicative0_10_39 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_34.V0), gopurs_runtime.Value{}))
_ = Applicative0_10_39
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_30)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_38.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_38.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_39.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_28.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t40 = gopurs_runtime.Apply(pure_5_29, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_40
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t40 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}))
})}))}
})})
_ = __local_var_3_17
// TAST (Let): Bind1_4_41 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_17.V1), gopurs_runtime.Value{}))
_ = Bind1_4_41
// TAST (Let): Applicative0_5_42 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_17.V0), gopurs_runtime.Value{}))
_ = Applicative0_5_42
// TAST (Let): applyExceptT1_2_12 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
applyExceptT1_2_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_2_13)}
}), gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_41.V1), f_6, gopurs_runtime.Func(func(f_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_41.V1), a_7, gopurs_runtime.Func(func(a_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_42.V1), gopurs_runtime.Apply(f_prime__8, a_prime__9))
}))
}))
})})
_ = applyExceptT1_2_12
// TAST (Let): semigroupExceptT1__193435443_2_11 shape=Let(Abs(LitRecord)) bindingType=Any
semigroupExceptT1__193435443_2_11 := gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_43 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_43 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_2_12.V0), gopurs_runtime.Value{}))
_ = Functor0_4_43
// TAST (Let): __local_var_5_44 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_5_44 := gopurs_runtime.RecordGet(dictSemigroup_3, "append")
_ = __local_var_5_44
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_2_12.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_43.V0), __local_var_5_44, a_6), b_7)
})}))}
})
_ = semigroupExceptT1__193435443_2_11
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupExceptT2_4_45 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
semigroupExceptT2_4_45 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupExceptT1__193435443_2_11, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})))
_ = semigroupExceptT2_4_45
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupExceptT2_4_45)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeExceptT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))}))}
})
}

func Call_Control_Monad_Except_Trans_altExceptT(dictSemigroup_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
_ = dictMonad_1
// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_1
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 shape=App(Other) bindingType=Any
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})})
_ = functorExceptT1_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v_5, gopurs_runtime.Func(func(rm_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (rm_7.Type == 9 && rm_7.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rm_7.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_8
} else {

}
}
{
if (rm_7.Type == 9 && rm_7.IntVal == 3711209382) {
// TAST (Let): __local_var_8_6 shape=Other bindingType=Any
__local_var_8_6 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rm_7.UnsafePtr).V0
_ = __local_var_8_6
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v1_6, gopurs_runtime.Func(func(rn_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (rn_9.Type == 9 && rn_9.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rn_9.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_7
} else {

}
}
{
if (rn_9.Type == 9 && rn_9.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_6, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rn_9.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
})}))}
}

func Call_Control_Monad_Except_Trans_plusExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar e)
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): altExceptT1__193435443_2_1 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
altExceptT1__193435443_2_1 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_3 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): Applicative0_5_4 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=Any
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): functorExceptT1_6_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 shape=App(Other) bindingType=Any
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t8 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_8
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t8 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
_ = __local_var_8_7
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_7, v_9)
})
})})
_ = functorExceptT1_6_5
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_5)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v_7, gopurs_runtime.Func(func(rm_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (rm_9.Type == 9 && rm_9.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rm_9.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_11
} else {

}
}
{
if (rm_9.Type == 9 && rm_9.IntVal == 3711209382) {
// TAST (Let): __local_var_10_9 shape=Other bindingType=Any
__local_var_10_9 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rm_9.UnsafePtr).V0
_ = __local_var_10_9
__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v1_8, gopurs_runtime.Func(func(rn_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (rn_11.Type == 9 && rn_11.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rn_11.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_10
} else {

}
}
{
if (rn_11.Type == 9 && rn_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "append"), __local_var_10_9, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rn_11.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}))
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}))
})}))}
})
_ = altExceptT1__193435443_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): altExceptT2_4_12 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
altExceptT2_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(altExceptT1__193435443_2_1, dictMonad_3))
_ = altExceptT2_4_12
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altExceptT2_4_12)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{mempty_1_0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())}))}
})
}

func Call_Control_Monad_Except_Trans_alternativeExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_1 shape=Other bindingType=(TypeVar e)
mempty_1_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): altExceptT1__193435443_2_2 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
altExceptT1__193435443_2_2 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_4 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_4
// TAST (Let): Applicative0_5_5 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_5
// TAST (Let): __local_var_6_7 shape=App(Other) bindingType=Any
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorExceptT1_6_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 shape=App(Other) bindingType=Any
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t9 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_9
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t9 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}))
_ = __local_var_8_8
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_8, v_9)
})
})})
_ = functorExceptT1_6_6
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_6)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), v_7, gopurs_runtime.Func(func(rm_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (rm_9.Type == 9 && rm_9.IntVal == 2465973597) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rm_9.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_12
} else {

}
}
{
if (rm_9.Type == 9 && rm_9.IntVal == 3711209382) {
// TAST (Let): __local_var_10_10 shape=Other bindingType=Any
__local_var_10_10 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rm_9.UnsafePtr).V0
_ = __local_var_10_10
__t12 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), v1_8, gopurs_runtime.Func(func(rn_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (rn_11.Type == 9 && rn_11.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rn_11.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_11
} else {

}
}
{
if (rn_11.Type == 9 && rn_11.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), __local_var_10_10, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rn_11.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}))
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
})}))}
})
_ = altExceptT1__193435443_2_2
// TAST (Let): plusExceptT1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
plusExceptT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): altExceptT2_4_13 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
altExceptT2_4_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(altExceptT1__193435443_2_2, dictMonad_3))
_ = altExceptT2_4_13
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altExceptT2_4_13)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{mempty_1_1, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())}))}
})
_ = plusExceptT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeExceptT1_3_14 shape=LitRecord bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
applicativeExceptT1_3_14 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_16 shape=App(Other) bindingType=Any
__local_var_4_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_16
// TAST (Let): functorExceptT1_4_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_17 shape=App(Other) bindingType=Any
__local_var_6_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_16, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t18 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_18
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t18 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
_ = __local_var_6_17
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_17, v_7)
})
})})
_ = functorExceptT1_4_15
// TAST (Let): __local_var_5_19 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_5_19 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_2)))}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): pure_7_21 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_7_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_21
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_2)))}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t22 = gopurs_runtime.Apply(pure_7_21, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_22
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t22 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}))
})}))}
})})
_ = __local_var_5_19
// TAST (Let): Bind1_6_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_19.V1), gopurs_runtime.Value{}))
_ = Bind1_6_23
// TAST (Let): Applicative0_7_24 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_19.V0), gopurs_runtime.Value{}))
_ = Applicative0_7_24
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_15)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_23.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_23.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_24.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})})
_ = applicativeExceptT1_3_14
// TAST (Let): plusExceptT2_4_25 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
plusExceptT2_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(plusExceptT1__193435443_1_0, dictMonad_2))
_ = plusExceptT2_4_25
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeExceptT1_3_14)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusExceptT2_4_25)}
})}))}
})
}

func Call_Control_Monad_Except_Trans_monadPlusExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_2 shape=Other bindingType=(TypeVar e)
mempty_1_2 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_2
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=Any
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): altExceptT1__193435443_2_3 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
altExceptT1__193435443_2_3 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_5 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_5
// TAST (Let): Applicative0_5_6 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_6
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=Any
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_8
// TAST (Let): functorExceptT1_6_7 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_6_7 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=Any
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t10 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_10
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t10 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_8.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}))
_ = __local_var_8_9
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_9, v_9)
})
})})
_ = functorExceptT1_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_7)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_5.V1), v_7, gopurs_runtime.Func(func(rm_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (rm_9.Type == 9 && rm_9.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_6.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rm_9.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_13
} else {

}
}
{
if (rm_9.Type == 9 && rm_9.IntVal == 3711209382) {
// TAST (Let): __local_var_10_11 shape=Other bindingType=Any
__local_var_10_11 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rm_9.UnsafePtr).V0
_ = __local_var_10_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_5.V1), v1_8, gopurs_runtime.Func(func(rn_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (rn_11.Type == 9 && rn_11.IntVal == 2465973597) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_6.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(rn_11.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_12
} else {

}
}
{
if (rn_11.Type == 9 && rn_11.IntVal == 3711209382) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_6.V1), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), __local_var_10_11, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(rn_11.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
})}))}
})
_ = altExceptT1__193435443_2_3
// TAST (Let): plusExceptT1__193435443_1_1 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
plusExceptT1__193435443_1_1 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): altExceptT2_4_14 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
altExceptT2_4_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(altExceptT1__193435443_2_3, dictMonad_3))
_ = altExceptT2_4_14
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altExceptT2_4_14)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{mempty_1_2, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())}))}
})
_ = plusExceptT1__193435443_1_1
// TAST (Let): alternativeExceptT1__193435443_1_0 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
alternativeExceptT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeExceptT1_3_15 shape=LitRecord bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
applicativeExceptT1_3_15 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_17 shape=App(Other) bindingType=Any
__local_var_4_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_17
// TAST (Let): functorExceptT1_4_16 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_4_16 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_18 shape=App(Other) bindingType=Any
__local_var_6_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_17, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t19 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_19
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t19 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}))
_ = __local_var_6_18
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_18, v_7)
})
})})
_ = functorExceptT1_4_16
// TAST (Let): __local_var_5_20 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_5_20 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_22 shape=App(Other) bindingType=Any
__local_var_7_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_22
// TAST (Let): functorExceptT1_7_21 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_21 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_23 shape=App(Other) bindingType=Any
__local_var_9_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_22, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t24 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_24
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t24 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
return __t24
}))
_ = __local_var_9_23
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_23, v_10)
})
})})
_ = functorExceptT1_7_21
// TAST (Let): __local_var_8_25 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_8_25 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_2)))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_26 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_26
// TAST (Let): pure_10_27 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_10_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_27
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_2)))}
}), gopurs_runtime.Func2(func(v_11 gopurs_runtime.Value, k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_26.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t28 = gopurs_runtime.Apply(pure_10_27, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_28
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t28 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0)
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}))
})}))}
})})
_ = __local_var_8_25
// TAST (Let): Bind1_9_29 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_25.V1), gopurs_runtime.Value{}))
_ = Bind1_9_29
// TAST (Let): Applicative0_10_30 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_25.V0), gopurs_runtime.Value{}))
_ = Applicative0_10_30
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_21)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_29.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_29.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_30.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_6, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_31 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_31
// TAST (Let): pure_7_32 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_7_32 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_32
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_34 shape=App(Other) bindingType=Any
__local_var_9_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_34
// TAST (Let): functorExceptT1_9_33 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_9_33 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_35 shape=App(Other) bindingType=Any
__local_var_11_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_34, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t36 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_11.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_36
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t36 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_11.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_36
} else {

}
}
{
__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_36:
return __t36
}))
_ = __local_var_11_35
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_35, v_12)
})
})})
_ = functorExceptT1_9_33
// TAST (Let): __local_var_10_37 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
__local_var_10_37 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_2)))}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_11, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_38 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_11_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_38
// TAST (Let): pure_12_39 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_12_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_39
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_2)))}
}), gopurs_runtime.Func2(func(v_13 gopurs_runtime.Value, k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_38.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t40 = gopurs_runtime.Apply(pure_12_39, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_40
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t40 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0)
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}))
})}))}
})})
_ = __local_var_10_37
// TAST (Let): Bind1_11_41 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_11_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_10_37.V1), gopurs_runtime.Value{}))
_ = Bind1_11_41
// TAST (Let): Applicative0_12_42 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_12_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_10_37.V0), gopurs_runtime.Value{}))
_ = Applicative0_12_42
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_33)}
}), gopurs_runtime.Func2(func(f_13 gopurs_runtime.Value, a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_41.V1), f_13, gopurs_runtime.Func(func(f_prime__15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_41.V1), a_14, gopurs_runtime.Func(func(a_prime__16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_42.V1), gopurs_runtime.Apply(f_prime__15, a_prime__16))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_31.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t43 = gopurs_runtime.Apply(pure_7_32, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_43
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t43 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0)
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}))
})}))}
})})
_ = __local_var_5_20
// TAST (Let): Bind1_6_44 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_20.V1), gopurs_runtime.Value{}))
_ = Bind1_6_44
// TAST (Let): Applicative0_7_45 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_7_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_20.V0), gopurs_runtime.Value{}))
_ = Applicative0_7_45
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_16)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_44.V1), f_8, gopurs_runtime.Func(func(f_prime__10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_44.V1), a_9, gopurs_runtime.Func(func(a_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_45.V1), gopurs_runtime.Apply(f_prime__10, a_prime__11))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_3, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})})
_ = applicativeExceptT1_3_15
// TAST (Let): plusExceptT2_4_46 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
plusExceptT2_4_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(plusExceptT1__193435443_1_1, dictMonad_2))
_ = plusExceptT2_4_46
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeExceptT1_3_15)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusExceptT2_4_46)}
})}))}
})
_ = alternativeExceptT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadExceptT1_3_47 shape=LitRecord bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
monadExceptT1_3_47 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_49 shape=App(Other) bindingType=Any
__local_var_5_49 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_49
// TAST (Let): functorExceptT1_5_48 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_5_48 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_50 shape=App(Other) bindingType=Any
__local_var_7_50 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_49, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t51 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_51
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t51 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_7.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_51
} else {

}
}
{
__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_51:
return __t51
}))
_ = __local_var_7_50
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_50, v_8)
})
})})
_ = functorExceptT1_5_48
// TAST (Let): __local_var_6_52 shape=App(Var) bindingType=Any
__local_var_6_52 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_2)
_ = __local_var_6_52
// TAST (Let): Bind1_7_53 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_7_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_52, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_53
// TAST (Let): Applicative0_8_54 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_8_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_52, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_54
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_48)}
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_53.V1), f_9, gopurs_runtime.Func(func(f_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_53.V1), a_10, gopurs_runtime.Func(func(a_prime__12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_54.V1), gopurs_runtime.Apply(f_prime__11, a_prime__12))
}))
}))
})}))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_4, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
})}))}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_55 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_55
// TAST (Let): pure_5_56 shape=Other bindingType=(Func [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar b)])]))
pure_5_56 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_56
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_58 shape=App(Other) bindingType=Any
__local_var_7_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_58
// TAST (Let): functorExceptT1_7_57 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
functorExceptT1_7_57 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_59 shape=App(Other) bindingType=Any
__local_var_9_59 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_58, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t60 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t60 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_60
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t60 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_9.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_60
} else {

}
}
{
__t60 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_60:
return __t60
}))
_ = __local_var_9_59
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_59, v_10)
})
})})
_ = functorExceptT1_7_57
// TAST (Let): __local_var_8_61 shape=App(Var) bindingType=Any
__local_var_8_61 := Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_2)
_ = __local_var_8_61
// TAST (Let): Bind1_9_62 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_9_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_61, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_62
// TAST (Let): Applicative0_10_63 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_10_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_61, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_63
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_57)}
}), gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_62.V1), f_11, gopurs_runtime.Func(func(f_prime__13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_62.V1), a_12, gopurs_runtime.Func(func(a_prime__14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_63.V1), gopurs_runtime.Apply(f_prime__13, a_prime__14))
}))
}))
})}))}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_55.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t64 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t64 = gopurs_runtime.Apply(pure_5_56, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_64
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t64 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_64
} else {

}
}
{
__t64 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_64:
return __t64
}))
})}))}
})})
_ = monadExceptT1_3_47
// TAST (Let): alternativeExceptT2_4_65 shape=App(Other) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
alternativeExceptT2_4_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](gopurs_runtime.Apply(alternativeExceptT1__193435443_1_0, dictMonad_2))
_ = alternativeExceptT2_4_65
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeExceptT2_4_65)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_3_47)}
})}))}
})
}


