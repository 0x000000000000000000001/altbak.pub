package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_Trans_WriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_WriterT sync.Once
func Get_Control_Monad_Writer_Trans_WriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_WriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_WriterT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_WriterT(x_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_WriterT
}

var cache_Control_Monad_Writer_Trans_runWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_runWriterT sync.Once
func Get_Control_Monad_Writer_Trans_runWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_runWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_runWriterT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_runWriterT(v_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_runWriterT
}

var cache_Control_Monad_Writer_Trans_newtypeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_newtypeWriterT sync.Once
func Get_Control_Monad_Writer_Trans_newtypeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_newtypeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_newtypeWriterT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Monad_Writer_Trans_newtypeWriterT
}

var cache_Control_Monad_Writer_Trans_monadTransWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadTransWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadTransWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadTransWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadTransWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadTransWriterT
}

var cache_Control_Monad_Writer_Trans_mapWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_mapWriterT sync.Once
func Get_Control_Monad_Writer_Trans_mapWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_mapWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_mapWriterT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_mapWriterT(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_mapWriterT
}

var cache_Control_Monad_Writer_Trans_functorWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_functorWriterT sync.Once
func Get_Control_Monad_Writer_Trans_functorWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_functorWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_functorWriterT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_functorWriterT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_functorWriterT
}

var cache_Control_Monad_Writer_Trans_execWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_execWriterT sync.Once
func Get_Control_Monad_Writer_Trans_execWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_execWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_execWriterT = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_execWriterT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_execWriterT
}

var cache_Control_Monad_Writer_Trans_applyWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_applyWriterT sync.Once
func Get_Control_Monad_Writer_Trans_applyWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_applyWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_applyWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0_box, dictApply_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_applyWriterT
}

var cache_Control_Monad_Writer_Trans_bindWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_bindWriterT sync.Once
func Get_Control_Monad_Writer_Trans_bindWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_bindWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_bindWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictBind_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_bindWriterT(dictSemigroup_0_box, dictBind_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_bindWriterT
}

var cache_Control_Monad_Writer_Trans_semigroupWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_semigroupWriterT sync.Once
func Get_Control_Monad_Writer_Trans_semigroupWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_semigroupWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_semigroupWriterT = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_box, dictSemigroup_1_box)
})
	})
	return cache_Control_Monad_Writer_Trans_semigroupWriterT
}

var cache_Control_Monad_Writer_Trans_applicativeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_applicativeWriterT sync.Once
func Get_Control_Monad_Writer_Trans_applicativeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_applicativeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_applicativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_applicativeWriterT
}

var cache_Control_Monad_Writer_Trans_monadWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadWriterT
}

var cache_Control_Monad_Writer_Trans_monadAskWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadAskWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadAskWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadAskWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadAskWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadAskWriterT
}

var cache_Control_Monad_Writer_Trans_monadReaderWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadReaderWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadReaderWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadReaderWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadReaderWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadReaderWriterT
}

var cache_Control_Monad_Writer_Trans_monadContWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadContWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadContWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadContWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadContWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadContWriterT
}

var cache_Control_Monad_Writer_Trans_monadEffectWriter gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadEffectWriter sync.Once
func Get_Control_Monad_Writer_Trans_monadEffectWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadEffectWriter.Do(func() {
		cache_Control_Monad_Writer_Trans_monadEffectWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadEffectWriter
}

var cache_Control_Monad_Writer_Trans_monadRecWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadRecWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadRecWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadRecWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadRecWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadRecWriterT
}

var cache_Control_Monad_Writer_Trans_monadStateWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadStateWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadStateWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadStateWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadStateWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadStateWriterT
}

var cache_Control_Monad_Writer_Trans_monadTellWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadTellWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadTellWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadTellWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadTellWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadTellWriterT
}

var cache_Control_Monad_Writer_Trans_monadWriterWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadWriterWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadWriterWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadWriterWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadWriterWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadWriterWriterT
}

var cache_Control_Monad_Writer_Trans_monadThrowWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadThrowWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadThrowWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadThrowWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadThrowWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadThrowWriterT
}

var cache_Control_Monad_Writer_Trans_monadErrorWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadErrorWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadErrorWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadErrorWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadErrorWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadErrorWriterT
}

var cache_Control_Monad_Writer_Trans_monadSTWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadSTWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadSTWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadSTWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadSTWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadSTWriterT
}

var cache_Control_Monad_Writer_Trans_monoidWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monoidWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monoidWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monoidWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monoidWriterT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monoidWriterT
}

var cache_Control_Monad_Writer_Trans_altWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_altWriterT sync.Once
func Get_Control_Monad_Writer_Trans_altWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_altWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_altWriterT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_altWriterT(dictAlt_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_altWriterT
}

var cache_Control_Monad_Writer_Trans_plusWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_plusWriterT sync.Once
func Get_Control_Monad_Writer_Trans_plusWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_plusWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_plusWriterT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_plusWriterT(dictPlus_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_plusWriterT
}

var cache_Control_Monad_Writer_Trans_alternativeWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_alternativeWriterT sync.Once
func Get_Control_Monad_Writer_Trans_alternativeWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_alternativeWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_alternativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_alternativeWriterT
}

var cache_Control_Monad_Writer_Trans_monadPlusWriterT gopurs_runtime.Value
var once_Control_Monad_Writer_Trans_monadPlusWriterT sync.Once
func Get_Control_Monad_Writer_Trans_monadPlusWriterT() gopurs_runtime.Value {
	once_Control_Monad_Writer_Trans_monadPlusWriterT.Do(func() {
		cache_Control_Monad_Writer_Trans_monadPlusWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Writer_Trans_monadPlusWriterT
}

func Call_Control_Monad_Writer_Trans_WriterT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Writer_Trans_runWriterT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): pure_3_1 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
})}))}
}

func Call_Control_Monad_Writer_Trans_mapWriterT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Writer_Trans_functorWriterT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
})}))}
}

func Call_Control_Monad_Writer_Trans_execWriterT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Data_Tuple_snd(), v_1)
}

func Call_Control_Monad_Writer_Trans_applyWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorWriterT1_3_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_3_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, v_6)
})
})})
_ = functorWriterT1_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_3_1)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func2(func(v3_6 gopurs_runtime.Value, v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_4), v1_5)
})}))}
}

func Call_Control_Monad_Writer_Trans_bindWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictBind_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictBind_1 gopurs_runtime.Value = dictBind_1_loop
_ = dictBind_1
// TAST (Let): Apply0_2_0 shape=App(Other) bindingType=Any
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_0
// TAST (Let): Functor0_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): __local_var_5_5 shape=App(Other) bindingType=Any
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorWriterT1_5_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
})})
_ = functorWriterT1_5_4
// TAST (Let): applyWriterT2_4_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_4)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})})
_ = applyWriterT2_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_2)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_1, "bind"), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 shape=Other bindingType=Any
__local_var_8_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1
_ = __local_var_8_7
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0))
}))
})}))}
}

func Call_Control_Monad_Writer_Trans_semigroupWriterT(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorWriterT1_3_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_3_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_5_4
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, v_6)
})
})})
_ = functorWriterT1_3_2
// TAST (Let): applyWriterT1_2_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT1_2_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_3_2)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func2(func(v3_6 gopurs_runtime.Value, v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_4), v1_5)
})})
_ = applyWriterT1_2_0
return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyWriterT1_2_0.V0), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_5_6 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_5_6 := gopurs_runtime.RecordGet(dictSemigroup1_3, "append")
_ = __local_var_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyWriterT1_2_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), __local_var_5_6, a_6), b_7)
})}))}
})
}

func Call_Control_Monad_Writer_Trans_applicativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): applyWriterT1__193435443_1_0 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=Any
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): functorWriterT1_4_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, v_7)
})
})})
_ = functorWriterT1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_3)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_6 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_6
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_6)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): applyWriterT1__193435443_1_1 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_1 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): functorWriterT1_4_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_6
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, v_7)
})
})})
_ = functorWriterT1_4_4
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_4)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_3.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_1
// TAST (Let): applicativeWriterT1__193435443_1_0 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_7 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_7
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_7)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_1_0
// TAST (Let): __local_var_2_9 shape=App(Other) bindingType=Any
__local_var_2_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_9
// TAST (Let): bindWriterT1__193435443_2_8 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_2_8 := gopurs_runtime.Func(func(dictBind_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_10 shape=App(Other) bindingType=Any
Apply0_4_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_10
// TAST (Let): Functor0_5_11 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_11
// TAST (Let): Functor0_6_13 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_13
// TAST (Let): __local_var_7_15 shape=App(Other) bindingType=Any
__local_var_7_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_15
// TAST (Let): functorWriterT1_7_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_7_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_16 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_9_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_9_16
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_16, v_10)
})
})})
_ = functorWriterT1_7_14
// TAST (Let): applyWriterT2_6_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_6_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_7_14)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_4_10, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_13.V0), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_9, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_8), v1_9)
})})
_ = applyWriterT2_6_12
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_6_12)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_3, "bind"), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_17 shape=Other bindingType=Any
__local_var_10_17 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1
_ = __local_var_10_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_11.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_9, "append"), __local_var_10_17, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_2_8
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_18 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_4_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_4_18
// TAST (Let): bindWriterT2_5_19 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_5_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_2_8, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_5_19
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_18)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_5_19)}
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadAskWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadTransWriterT1_1_0 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
})})
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_6 shape=App(Other) bindingType=Any
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): applyWriterT1__193435443_2_5 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_5 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_7
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=Any
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_9
// TAST (Let): functorWriterT1_5_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_10 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_10
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_10, v_8)
})
})})
_ = functorWriterT1_5_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_8)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_7.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_5
// TAST (Let): applicativeWriterT1__193435443_2_4 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_4 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_11 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_11
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_11)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_4
// TAST (Let): __local_var_3_13 shape=App(Other) bindingType=Any
__local_var_3_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_13
// TAST (Let): bindWriterT1__193435443_3_12 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_12 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_14 shape=App(Other) bindingType=Any
Apply0_5_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_14
// TAST (Let): Functor0_6_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_15
// TAST (Let): Functor0_7_17 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_17
// TAST (Let): __local_var_8_19 shape=App(Other) bindingType=Any
__local_var_8_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_19
// TAST (Let): functorWriterT1_8_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_20 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_19, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_20
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_20, v_11)
})
})})
_ = functorWriterT1_8_18
// TAST (Let): applyWriterT2_7_16 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_16 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_18)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_17.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_16)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_21 shape=Other bindingType=Any
__local_var_11_21 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_15.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "append"), __local_var_11_21, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_12
// TAST (Let): monadWriterT1__193435443_2_3 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_3 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_22 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_22
// TAST (Let): bindWriterT2_6_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_12, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_23
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_22)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_23)}
})}))}
})
_ = monadWriterT1__193435443_2_3
return gopurs_runtime.Func(func(dictMonadAsk_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_24 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_4_24
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_24)}
}), gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_3, "ask"))}))}
})
}

func Call_Control_Monad_Writer_Trans_monadReaderWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_1 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadTransWriterT1_1_1 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_2
// TAST (Let): pure_3_3 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_3_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_3
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_2.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
})})
_ = monadTransWriterT1_1_1
// TAST (Let): __local_var_2_7 shape=App(Other) bindingType=Any
__local_var_2_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_7
// TAST (Let): applyWriterT1__193435443_2_6 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_6 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_8
// TAST (Let): __local_var_5_10 shape=App(Other) bindingType=Any
__local_var_5_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_10
// TAST (Let): functorWriterT1_5_9 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_9 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_11
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_11, v_8)
})
})})
_ = functorWriterT1_5_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_9)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_8.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_7, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_6
// TAST (Let): applicativeWriterT1__193435443_2_5 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_5 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_12 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_6, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_12
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_12)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_5
// TAST (Let): __local_var_3_14 shape=App(Other) bindingType=Any
__local_var_3_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_14
// TAST (Let): bindWriterT1__193435443_3_13 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_13 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_15 shape=App(Other) bindingType=Any
Apply0_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_15
// TAST (Let): Functor0_6_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_16
// TAST (Let): Functor0_7_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_18
// TAST (Let): __local_var_8_20 shape=App(Other) bindingType=Any
__local_var_8_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_20
// TAST (Let): functorWriterT1_8_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_21 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_20, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_21
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_21, v_11)
})
})})
_ = functorWriterT1_8_19
// TAST (Let): applyWriterT2_7_17 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_17 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_19)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_15, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_18.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_17
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_17)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_22 shape=Other bindingType=Any
__local_var_11_22 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_16.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "append"), __local_var_11_22, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_13
// TAST (Let): monadWriterT1__193435443_2_4 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_4 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_23 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_23
// TAST (Let): bindWriterT2_6_24 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_13, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_24
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_23)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_24)}
})}))}
})
_ = monadWriterT1__193435443_2_4
// TAST (Let): monadAskWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
monadAskWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonadAsk_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_25 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_4_25
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_25)}
}), gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_1.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_3, "ask"))}))}
})
_ = monadAskWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonadReader_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadAskWriterT2_3_26 shape=App(Other) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r), (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadAskWriterT2_3_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadAskWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "MonadAsk0"), gopurs_runtime.Value{})))
_ = monadAskWriterT2_3_26
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskWriterT2_3_26)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_27 shape=App(Other) bindingType=Any
__local_var_5_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "local"), f_4)
_ = __local_var_5_27
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_27, v_6)
})
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadContWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): applyWriterT1__193435443_1_2 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_2 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorWriterT1_4_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_7
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_7, v_7)
})
})})
_ = functorWriterT1_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_5)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_2
// TAST (Let): applicativeWriterT1__193435443_1_1 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_1_1 := gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_8 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_8
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_8)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_1_1
// TAST (Let): __local_var_2_10 shape=App(Other) bindingType=Any
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_10
// TAST (Let): bindWriterT1__193435443_2_9 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_2_9 := gopurs_runtime.Func(func(dictBind_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_11 shape=App(Other) bindingType=Any
Apply0_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_11
// TAST (Let): Functor0_5_12 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_12
// TAST (Let): Functor0_6_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=Any
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_16
// TAST (Let): functorWriterT1_7_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_7_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_17 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_16, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_9_17
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_17, v_10)
})
})})
_ = functorWriterT1_7_15
// TAST (Let): applyWriterT2_6_13 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_6_13 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_7_15)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_4_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_8), v1_9)
})})
_ = applyWriterT2_6_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_6_13)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_3, "bind"), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_18 shape=Other bindingType=Any
__local_var_10_18 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1
_ = __local_var_10_18
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_12.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), __local_var_10_18, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_2_9
// TAST (Let): monadWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_4_19
// TAST (Let): bindWriterT2_5_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_2_9, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_19)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_5_20)}
})}))}
})
_ = monadWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonadCont_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_3_21 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_3_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_2, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_3_21
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_3_21)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_2, "callCC"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_6, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
}))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): applyWriterT1__193435443_1_2 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_2 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorWriterT1_4_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_7
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_7, v_7)
})
})})
_ = functorWriterT1_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_5)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_2
// TAST (Let): applicativeWriterT1__193435443_1_1 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_1_1 := gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_8 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_8
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_8)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_1_1
// TAST (Let): __local_var_2_10 shape=App(Other) bindingType=Any
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_10
// TAST (Let): bindWriterT1__193435443_2_9 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_2_9 := gopurs_runtime.Func(func(dictBind_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_11 shape=App(Other) bindingType=Any
Apply0_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_11
// TAST (Let): Functor0_5_12 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_12
// TAST (Let): Functor0_6_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=Any
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_16
// TAST (Let): functorWriterT1_7_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_7_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_17 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_16, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_9_17
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_17, v_10)
})
})})
_ = functorWriterT1_7_15
// TAST (Let): applyWriterT2_6_13 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_6_13 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_7_15)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_4_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_8), v1_9)
})})
_ = applyWriterT2_6_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_6_13)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_3, "bind"), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_18 shape=Other bindingType=Any
__local_var_10_18 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1
_ = __local_var_10_18
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_12.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), __local_var_10_18, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_2_9
// TAST (Let): monadWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_4_19
// TAST (Let): bindWriterT2_5_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_2_9, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_19)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_5_20)}
})}))}
})
_ = monadWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonadEffect_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_3_21 shape=App(Other) bindingType=Any
Monad0_3_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_21
// TAST (Let): monadWriterT2_4_22 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_1_0, Monad0_3_21))
_ = monadWriterT2_4_22
// TAST (Let): Bind1_5_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_21, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_23
// TAST (Let): pure_6_24 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_6_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_21, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_24
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_22)}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_23.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "liftEffect"), x_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_24, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadRecWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=Any
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): applyWriterT1__193435443_2_3 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_3 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): functorWriterT1_5_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_8
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_8, v_8)
})
})})
_ = functorWriterT1_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_6)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_3
// TAST (Let): applicativeWriterT1__193435443_2_2 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_2 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_9 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_9)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): bindWriterT1__193435443_3_10 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_10 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_12 shape=App(Other) bindingType=Any
Apply0_5_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_12
// TAST (Let): Functor0_6_13 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_13
// TAST (Let): Functor0_7_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_15
// TAST (Let): __local_var_8_17 shape=App(Other) bindingType=Any
__local_var_8_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_17
// TAST (Let): functorWriterT1_8_16 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_16 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_18 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_17, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_18
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_18, v_11)
})
})})
_ = functorWriterT1_8_16
// TAST (Let): applyWriterT2_7_14 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_14 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_16)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_12, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_15.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_14)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_19 shape=Other bindingType=Any
__local_var_11_19 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_19
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_13.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "append"), __local_var_11_19, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_10
// TAST (Let): monadWriterT1__193435443_2_1 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_1 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_20 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_20
// TAST (Let): bindWriterT2_6_21 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_10, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_21
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_20)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_21)}
})}))}
})
_ = monadWriterT1__193435443_2_1
return gopurs_runtime.Func(func(dictMonadRec_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_22 shape=App(Other) bindingType=Any
Monad0_4_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_22
// TAST (Let): Bind1_5_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_23
// TAST (Let): Applicative0_6_24 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_24
// TAST (Let): monadWriterT2_7_25 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_7_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_1, Monad0_4_22))
_ = monadWriterT2_7_25
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_7_25)}
}), gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_3, "tailRecM"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_26 shape=Other bindingType=Any
__local_var_11_26 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_11_26
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_23.V1), gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
var __t_tag_27 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0
if (__t_tag_27.Type == 9 && __t_tag_27.IntVal == 525585346) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Writer_Trans_3131450224_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), __local_var_11_26, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1)})})))}
goto end_branch_29
} else {

}
}
{
var __t_tag_28 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0
if (__t_tag_28.Type == 9 && __t_tag_28.IntVal == 60402430) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Writer_Trans_2886445004_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), __local_var_11_26, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1)})})))}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_24.V1), __t29)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadStateWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadTransWriterT1_1_0 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
})})
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_6 shape=App(Other) bindingType=Any
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): applyWriterT1__193435443_2_5 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_5 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_7
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=Any
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_9
// TAST (Let): functorWriterT1_5_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_10 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_10
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_10, v_8)
})
})})
_ = functorWriterT1_5_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_8)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_7.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_5
// TAST (Let): applicativeWriterT1__193435443_2_4 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_4 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_11 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_11
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_11)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_4
// TAST (Let): __local_var_3_13 shape=App(Other) bindingType=Any
__local_var_3_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_13
// TAST (Let): bindWriterT1__193435443_3_12 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_12 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_14 shape=App(Other) bindingType=Any
Apply0_5_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_14
// TAST (Let): Functor0_6_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_15
// TAST (Let): Functor0_7_17 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_17
// TAST (Let): __local_var_8_19 shape=App(Other) bindingType=Any
__local_var_8_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_19
// TAST (Let): functorWriterT1_8_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_20 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_19, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_20
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_20, v_11)
})
})})
_ = functorWriterT1_8_18
// TAST (Let): applyWriterT2_7_16 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_16 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_18)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_17.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_16)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_21 shape=Other bindingType=Any
__local_var_11_21 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_15.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "append"), __local_var_11_21, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_12
// TAST (Let): monadWriterT1__193435443_2_3 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_3 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_22 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_22
// TAST (Let): bindWriterT2_6_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_12, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_23
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_22)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_23)}
})}))}
})
_ = monadWriterT1__193435443_2_3
return gopurs_runtime.Func(func(dictMonadState_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_24 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
Monad0_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_24
// TAST (Let): monadWriterT2_5_25 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_5_25
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_5_25)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_24)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "state"), f_6))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadTellWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=Any
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): applyWriterT1__193435443_2_3 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_3 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): functorWriterT1_5_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_8
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_8, v_8)
})
})})
_ = functorWriterT1_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_6)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_3
// TAST (Let): applicativeWriterT1__193435443_2_2 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_2 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_9 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_9)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): bindWriterT1__193435443_3_10 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_10 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_12 shape=App(Other) bindingType=Any
Apply0_5_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_12
// TAST (Let): Functor0_6_13 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_13
// TAST (Let): Functor0_7_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_15
// TAST (Let): __local_var_8_17 shape=App(Other) bindingType=Any
__local_var_8_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_17
// TAST (Let): functorWriterT1_8_16 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_16 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_18 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_17, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_18
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_18, v_11)
})
})})
_ = functorWriterT1_8_16
// TAST (Let): applyWriterT2_7_14 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_14 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_16)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_12, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_15.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_14)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_19 shape=Other bindingType=Any
__local_var_11_19 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_19
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_13.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "append"), __local_var_11_19, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_10
// TAST (Let): monadWriterT1__193435443_2_1 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_1 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_20 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_20
// TAST (Let): bindWriterT2_6_21 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_10, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_21
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_20)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_21)}
})}))}
})
_ = monadWriterT1__193435443_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_22 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_1, dictMonad_3))
_ = monadWriterT2_4_22
// TAST (Let): __local_var_5_24 shape=App(Var) bindingType=(Func [(TypeVar w)] (ADT ["Data","Tuple","Tuple"] [Unit, (TypeVar w)]))
__local_var_5_24 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), Get_Data_Unit_unit())
_ = __local_var_5_24
// TAST (Let): __local_var_5_23 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar w)] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [Unit, (TypeVar w)])]))
__local_var_5_23 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(__local_var_5_24, x_6))
})
_ = __local_var_5_23
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_22)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_1_0)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_23, x_6)
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadWriterWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
// TAST (Let): __local_var_2_5 shape=App(Other) bindingType=Any
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_5
// TAST (Let): applyWriterT1__193435443_2_4 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_4 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_6
// TAST (Let): __local_var_5_8 shape=App(Other) bindingType=Any
__local_var_5_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_8
// TAST (Let): functorWriterT1_5_7 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_7 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_9 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_9
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_9, v_8)
})
})})
_ = functorWriterT1_5_7
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_7)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_6.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_4
// TAST (Let): applicativeWriterT1__193435443_2_3 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_3 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_10 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_10)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_3
// TAST (Let): __local_var_3_12 shape=App(Other) bindingType=Any
__local_var_3_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_12
// TAST (Let): bindWriterT1__193435443_3_11 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_11 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_13 shape=App(Other) bindingType=Any
Apply0_5_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_13
// TAST (Let): Functor0_6_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_13, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
// TAST (Let): Functor0_7_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_13, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_16
// TAST (Let): __local_var_8_18 shape=App(Other) bindingType=Any
__local_var_8_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_18
// TAST (Let): functorWriterT1_8_17 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_17 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_19 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_18, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_19
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_19, v_11)
})
})})
_ = functorWriterT1_8_17
// TAST (Let): applyWriterT2_7_15 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_15 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_17)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_13, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_16.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_12, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_15)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_20 shape=Other bindingType=Any
__local_var_11_20 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_20
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_12, "append"), __local_var_11_20, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_11
// TAST (Let): monadWriterT1__193435443_2_2 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_2 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_21 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_21
// TAST (Let): bindWriterT2_6_22 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_11, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_22
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_21)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_22)}
})}))}
})
_ = monadWriterT1__193435443_2_2
// TAST (Let): monadTellWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
monadTellWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_23 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_2, dictMonad_3))
_ = monadWriterT2_4_23
// TAST (Let): __local_var_5_25 shape=App(Var) bindingType=(Func [(TypeVar w)] (ADT ["Data","Tuple","Tuple"] [Unit, (TypeVar w)]))
__local_var_5_25 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), Get_Data_Unit_unit())
_ = __local_var_5_25
// TAST (Let): __local_var_5_24 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar w)] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [Unit, (TypeVar w)])]))
__local_var_5_24 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(__local_var_5_25, x_6))
})
_ = __local_var_5_24
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_23)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_1_1)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_24, x_6)
})}))}
})
_ = monadTellWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_26 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_26
// TAST (Let): Applicative0_4_27 shape=App(Other) bindingType=Any
Applicative0_4_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_4_27
// TAST (Let): monadTellWriterT2_5_28 shape=App(Other) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w), (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadTellWriterT2_5_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadTellWriterT1__193435443_1_0, dictMonad_2))
_ = monadTellWriterT2_5_28
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellWriterT2_5_28)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_26.V1), v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_27, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Writer_Trans_1785332133_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1}))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_26.V1), v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_27, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadThrowWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_0 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadTransWriterT1_1_0 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
})})
_ = monadTransWriterT1_1_0
// TAST (Let): __local_var_2_6 shape=App(Other) bindingType=Any
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): applyWriterT1__193435443_2_5 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_5 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_7
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=Any
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_9
// TAST (Let): functorWriterT1_5_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_10 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_10
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_10, v_8)
})
})})
_ = functorWriterT1_5_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_8)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_7.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_5
// TAST (Let): applicativeWriterT1__193435443_2_4 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_4 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_11 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_11
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_11)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_4
// TAST (Let): __local_var_3_13 shape=App(Other) bindingType=Any
__local_var_3_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_13
// TAST (Let): bindWriterT1__193435443_3_12 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_12 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_14 shape=App(Other) bindingType=Any
Apply0_5_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_14
// TAST (Let): Functor0_6_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_15
// TAST (Let): Functor0_7_17 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_17
// TAST (Let): __local_var_8_19 shape=App(Other) bindingType=Any
__local_var_8_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_19
// TAST (Let): functorWriterT1_8_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_20 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_19, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_20
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_20, v_11)
})
})})
_ = functorWriterT1_8_18
// TAST (Let): applyWriterT2_7_16 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_16 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_18)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_14, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_17.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_16)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_21 shape=Other bindingType=Any
__local_var_11_21 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_15.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "append"), __local_var_11_21, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_12
// TAST (Let): monadWriterT1__193435443_2_3 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_3 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_22 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_22
// TAST (Let): bindWriterT2_6_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_12, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_23
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_22)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_23)}
})}))}
})
_ = monadWriterT1__193435443_2_3
return gopurs_runtime.Func(func(dictMonadThrow_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_24 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
Monad0_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_24
// TAST (Let): monadWriterT2_5_25 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_5_25
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_5_25)}
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_24)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "throwError"), e_6))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadErrorWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): monadTransWriterT1_1_1 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadTransWriterT1_1_1 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_2
// TAST (Let): pure_3_3 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_3_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_3
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_2.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
})})
_ = monadTransWriterT1_1_1
// TAST (Let): __local_var_2_7 shape=App(Other) bindingType=Any
__local_var_2_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_7
// TAST (Let): applyWriterT1__193435443_2_6 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_6 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_8
// TAST (Let): __local_var_5_10 shape=App(Other) bindingType=Any
__local_var_5_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_10
// TAST (Let): functorWriterT1_5_9 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_9 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_11
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_11, v_8)
})
})})
_ = functorWriterT1_5_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_9)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_8.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_7, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_6
// TAST (Let): applicativeWriterT1__193435443_2_5 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_5 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_12 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_6, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_12
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_12)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_5
// TAST (Let): __local_var_3_14 shape=App(Other) bindingType=Any
__local_var_3_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_14
// TAST (Let): bindWriterT1__193435443_3_13 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_3_13 := gopurs_runtime.Func(func(dictBind_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_15 shape=App(Other) bindingType=Any
Apply0_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_15
// TAST (Let): Functor0_6_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_16
// TAST (Let): Functor0_7_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_18
// TAST (Let): __local_var_8_20 shape=App(Other) bindingType=Any
__local_var_8_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_20
// TAST (Let): functorWriterT1_8_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_8_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_21 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_10_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_20, "map"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_10_21
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_21, v_11)
})
})})
_ = functorWriterT1_8_19
// TAST (Let): applyWriterT2_7_17 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_7_17 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_8_19)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_5_15, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_18.V0), gopurs_runtime.Func2(func(v3_11 gopurs_runtime.Value, v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_9), v1_10)
})})
_ = applyWriterT2_7_17
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_7_17)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_4, "bind"), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_22 shape=Other bindingType=Any
__local_var_11_22 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1
_ = __local_var_11_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_16.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "append"), __local_var_11_22, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_9, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_3_13
// TAST (Let): monadWriterT1__193435443_2_4 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_2_4 := gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_5_23 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_5_23
// TAST (Let): bindWriterT2_6_24 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_3_13, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_6_24
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_5_23)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_6_24)}
})}))}
})
_ = monadWriterT1__193435443_2_4
// TAST (Let): monadThrowWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadThrowWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonadThrow_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_4_25 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
Monad0_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_25
// TAST (Let): monadWriterT2_5_26 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_2_4, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_5_26
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_5_26)}
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransWriterT1_1_1.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_25)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "throwError"), e_6))
})}))}
})
_ = monadThrowWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonadError_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadThrowWriterT2_3_27 shape=App(Other) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e), (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadThrowWriterT2_3_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadThrowWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_2, "MonadThrow0"), gopurs_runtime.Value{})))
_ = monadThrowWriterT2_3_27
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowWriterT2_3_27)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_2, "catchError"), v_4, gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_5, e_6)
}))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadSTWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): applyWriterT1__193435443_1_2 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_2 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorWriterT1_4_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_7
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_7, v_7)
})
})})
_ = functorWriterT1_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_5)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_2
// TAST (Let): applicativeWriterT1__193435443_1_1 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_1_1 := gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_8 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_8
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_8)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_1_1
// TAST (Let): __local_var_2_10 shape=App(Other) bindingType=Any
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_10
// TAST (Let): bindWriterT1__193435443_2_9 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_2_9 := gopurs_runtime.Func(func(dictBind_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_11 shape=App(Other) bindingType=Any
Apply0_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_11
// TAST (Let): Functor0_5_12 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_12
// TAST (Let): Functor0_6_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=Any
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_16
// TAST (Let): functorWriterT1_7_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_7_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_17 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_16, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_9_17
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_17, v_10)
})
})})
_ = functorWriterT1_7_15
// TAST (Let): applyWriterT2_6_13 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_6_13 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_7_15)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_4_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_8), v1_9)
})})
_ = applyWriterT2_6_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_6_13)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_3, "bind"), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_18 shape=Other bindingType=Any
__local_var_10_18 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1
_ = __local_var_10_18
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_12.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), __local_var_10_18, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_2_9
// TAST (Let): monadWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_4_19
// TAST (Let): bindWriterT2_5_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_2_9, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_19)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_5_20)}
})}))}
})
_ = monadWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonadST_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_3_21 shape=App(Other) bindingType=Any
Monad0_3_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_21
// TAST (Let): monadWriterT2_4_22 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_1_0, Monad0_3_21))
_ = monadWriterT2_4_22
// TAST (Let): Bind1_5_23 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_21, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_23
// TAST (Let): pure_6_24 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])]))
pure_6_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_21, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_24
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_22)}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_23.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "liftST"), x_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_24, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monoidWriterT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupWriterT1__193435443_1_0 shape=Let(Abs(Let(Abs(LitRecord)))) bindingType=Any
semigroupWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): functorWriterT1_4_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_6
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, v_7)
})
})})
_ = functorWriterT1_4_4
// TAST (Let): applyWriterT1_3_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT1_3_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_4)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_3.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})})
_ = applyWriterT1_3_2
return gopurs_runtime.Func(func(dictSemigroup1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyWriterT1_3_2.V0), gopurs_runtime.Value{}))
_ = Functor0_5_7
// TAST (Let): __local_var_6_8 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_6_8 := gopurs_runtime.RecordGet(dictSemigroup1_4, "append")
_ = __local_var_6_8
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyWriterT1_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_7.V0), __local_var_6_8, a_7), b_8)
})}))}
})
})
_ = semigroupWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_12 shape=App(Other) bindingType=Any
__local_var_4_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_12
// TAST (Let): Functor0_5_13 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_13
// TAST (Let): __local_var_6_15 shape=App(Other) bindingType=Any
__local_var_6_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_15
// TAST (Let): functorWriterT1_6_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_6_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_16 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_8_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_8_16
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_16, v_9)
})
})})
_ = functorWriterT1_6_14
// TAST (Let): applyWriterT2_3_10 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_10 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_6_14)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_12, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_13.V0), gopurs_runtime.Func2(func(v3_9 gopurs_runtime.Value, v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_10.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_10.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_7), v1_8)
})})
_ = applyWriterT2_3_10
// TAST (Let): applicativeWriterT1_3_9 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT1_3_9 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_10)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_2, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})})
_ = applicativeWriterT1_3_9
// TAST (Let): semigroupWriterT2__193435443_4_17 shape=App(Other) bindingType=Any
semigroupWriterT2__193435443_4_17 := gopurs_runtime.Apply(semigroupWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupWriterT2__193435443_4_17
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupWriterT3_6_18 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
semigroupWriterT3_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupWriterT2__193435443_4_17, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_5, "Semigroup0"), gopurs_runtime.Value{})))
_ = semigroupWriterT3_6_18
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupWriterT3_6_18)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeWriterT1_3_9.V1), gopurs_runtime.RecordGet(dictMonoid1_5, "mempty"))}))}
})
})
}

func Call_Control_Monad_Writer_Trans_altWriterT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorWriterT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, v_4)
})
})})
_ = functorWriterT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_1_0)}
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_2, v1_3)
})}))}
}

func Call_Control_Monad_Writer_Trans_plusWriterT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorWriterT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, v_5)
})
})})
_ = functorWriterT1_2_2
// TAST (Let): altWriterT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
altWriterT1_1_0 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_2_2)}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "alt"), v_3, v1_4)
})})
_ = altWriterT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altWriterT1_1_0)}
}), gopurs_runtime.RecordGet(dictPlus_0, "empty")}))}
}

func Call_Control_Monad_Writer_Trans_alternativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): applyWriterT1__193435443_1_1 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_1 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): functorWriterT1_4_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_6
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, v_7)
})
})})
_ = functorWriterT1_4_4
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_4)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_3.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_1
// TAST (Let): applicativeWriterT1__193435443_1_0 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_7 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_7
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_7)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_3_8 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_3_8
// TAST (Let): __local_var_4_10 shape=App(Other) bindingType=Any
__local_var_4_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_10
// TAST (Let): __local_var_5_12 shape=App(Other) bindingType=Any
__local_var_5_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_10, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_12
// TAST (Let): __local_var_6_14 shape=App(Other) bindingType=Any
__local_var_6_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_14
// TAST (Let): functorWriterT1_6_13 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_6_13 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_15 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_8_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_14, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_8_15
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_15, v_9)
})
})})
_ = functorWriterT1_6_13
// TAST (Let): altWriterT1_5_11 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
altWriterT1_5_11 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_6_13)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_12, "alt"), v_7, v1_8)
})})
_ = altWriterT1_5_11
// TAST (Let): plusWriterT1_4_9 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
plusWriterT1_4_9 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altWriterT1_5_11)}
}), gopurs_runtime.RecordGet(__local_var_4_10, "empty")})
_ = plusWriterT1_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_3_8)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusWriterT1_4_9)}
})}))}
})
}

func Call_Control_Monad_Writer_Trans_monadPlusWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): applyWriterT1__193435443_1_2 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_1_2 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorWriterT1_4_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_4_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_6_7
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_7, v_7)
})
})})
_ = functorWriterT1_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_4_5)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_5), v1_6)
})}))}
})
_ = applyWriterT1__193435443_1_2
// TAST (Let): applicativeWriterT1__193435443_1_1 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_1_1 := gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_3_8 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_1_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_3_8
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_3_8)}
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_1_1
// TAST (Let): __local_var_2_10 shape=App(Other) bindingType=Any
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_10
// TAST (Let): bindWriterT1__193435443_2_9 shape=Let(Abs(Let(Let(Let(LitRecord))))) bindingType=Any
bindWriterT1__193435443_2_9 := gopurs_runtime.Func(func(dictBind_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_11 shape=App(Other) bindingType=Any
Apply0_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_11
// TAST (Let): Functor0_5_12 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_12
// TAST (Let): Functor0_6_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_14
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=Any
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_16
// TAST (Let): functorWriterT1_7_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_7_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_17 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_16, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_9_17
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_17, v_10)
})
})})
_ = functorWriterT1_7_15
// TAST (Let): applyWriterT2_6_13 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_6_13 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_7_15)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_4_11, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_14.V0), gopurs_runtime.Func2(func(v3_10 gopurs_runtime.Value, v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_10.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_8), v1_9)
})})
_ = applyWriterT2_6_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_6_13)}
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_3, "bind"), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_18 shape=Other bindingType=Any
__local_var_10_18 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1
_ = __local_var_10_18
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_12.V0), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "append"), __local_var_10_18, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_11.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(k_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0))
}))
})}))}
})
_ = bindWriterT1__193435443_2_9
// TAST (Let): monadWriterT1__193435443_1_0 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadWriterT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_19 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_4_19
// TAST (Let): bindWriterT2_5_20 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
bindWriterT2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindWriterT1__193435443_2_9, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})))
_ = bindWriterT2_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_19)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindWriterT2_5_20)}
})}))}
})
_ = monadWriterT1__193435443_1_0
// TAST (Let): __local_var_2_24 shape=App(Other) bindingType=Any
__local_var_2_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_24
// TAST (Let): applyWriterT1__193435443_2_23 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
applyWriterT1__193435443_2_23 := gopurs_runtime.Func(func(dictApply_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_25 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_25
// TAST (Let): __local_var_5_27 shape=App(Other) bindingType=Any
__local_var_5_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_27
// TAST (Let): functorWriterT1_5_26 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_5_26 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_28 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_7_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_27, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_7_28
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_28, v_8)
})
})})
_ = functorWriterT1_5_26
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_5_26)}
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_3, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_25.V0), gopurs_runtime.Func2(func(v3_8 gopurs_runtime.Value, v4_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_24, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_8.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_9.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), v_6), v1_7)
})}))}
})
_ = applyWriterT1__193435443_2_23
// TAST (Let): applicativeWriterT1__193435443_2_22 shape=Let(Abs(Let(LitRecord))) bindingType=Any
applicativeWriterT1__193435443_2_22 := gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT2_4_29 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1__193435443_2_23, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_29
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_29)}
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})}))}
})
_ = applicativeWriterT1__193435443_2_22
// TAST (Let): alternativeWriterT1__193435443_2_21 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
alternativeWriterT1__193435443_2_21 := gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeWriterT2_4_30 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applicativeWriterT2_4_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeWriterT1__193435443_2_22, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Applicative0"), gopurs_runtime.Value{})))
_ = applicativeWriterT2_4_30
// TAST (Let): __local_var_5_32 shape=App(Other) bindingType=Any
__local_var_5_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_5_32
// TAST (Let): __local_var_6_34 shape=App(Other) bindingType=Any
__local_var_6_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_32, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_6_34
// TAST (Let): __local_var_7_36 shape=App(Other) bindingType=Any
__local_var_7_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_34, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_36
// TAST (Let): functorWriterT1_7_35 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
functorWriterT1_7_35 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_37 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b), (TypeVar w)])]))
__local_var_9_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_36, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_9_37
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_37, v_10)
})
})})
_ = functorWriterT1_7_35
// TAST (Let): altWriterT1_6_33 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
altWriterT1_6_33 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorWriterT1_7_35)}
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_34, "alt"), v_8, v1_9)
})})
_ = altWriterT1_6_33
// TAST (Let): plusWriterT1_5_31 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
plusWriterT1_5_31 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altWriterT1_6_33)}
}), gopurs_runtime.RecordGet(__local_var_5_32, "empty")})
_ = plusWriterT1_5_31
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeWriterT2_4_30)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusWriterT1_5_31)}
})}))}
})
_ = alternativeWriterT1__193435443_2_21
return gopurs_runtime.Func(func(dictMonadPlus_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadWriterT2_4_38 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
monadWriterT2_4_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadWriterT1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Monad0"), gopurs_runtime.Value{})))
_ = monadWriterT2_4_38
// TAST (Let): alternativeWriterT2_5_39 shape=App(Other) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
alternativeWriterT2_5_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](gopurs_runtime.Apply(alternativeWriterT1__193435443_2_21, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Alternative1"), gopurs_runtime.Value{})))
_ = alternativeWriterT2_5_39
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeWriterT2_5_39)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadWriterT2_4_38)}
})}))}
})
}

func Rebox_Control_Monad_Writer_Trans_1785332133_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Writer_Trans_2886445004_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_Writer_Trans_3131450224_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}


