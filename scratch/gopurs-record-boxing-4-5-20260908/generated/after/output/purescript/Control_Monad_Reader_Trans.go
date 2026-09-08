package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Reader_Trans_ReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_ReaderT sync.Once
func Get_Control_Monad_Reader_Trans_ReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_ReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_ReaderT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_ReaderT(x_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_ReaderT
}

var cache_Control_Monad_Reader_Trans_withReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_withReaderT sync.Once
func Get_Control_Monad_Reader_Trans_withReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_withReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_withReaderT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_withReaderT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_withReaderT
}

var cache_Control_Monad_Reader_Trans_runReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_runReaderT sync.Once
func Get_Control_Monad_Reader_Trans_runReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_runReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_runReaderT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_runReaderT(v_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_runReaderT
}

var cache_Control_Monad_Reader_Trans_newtypeReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_newtypeReaderT sync.Once
func Get_Control_Monad_Reader_Trans_newtypeReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_newtypeReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_newtypeReaderT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Monad_Reader_Trans_newtypeReaderT
}

var cache_Control_Monad_Reader_Trans_monadTransReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadTransReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadTransReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadTransReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadTransReaderT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})}))}
	})
	return cache_Control_Monad_Reader_Trans_monadTransReaderT
}

var cache_Control_Monad_Reader_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_lift sync.Once
func Get_Control_Monad_Reader_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_lift.Do(func() {
		cache_Control_Monad_Reader_Trans_lift = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), x_1_box, v_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_lift
}

var cache_Control_Monad_Reader_Trans_mapReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_mapReaderT sync.Once
func Get_Control_Monad_Reader_Trans_mapReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_mapReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_mapReaderT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_mapReaderT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_mapReaderT
}

var cache_Control_Monad_Reader_Trans_functorReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_functorReaderT sync.Once
func Get_Control_Monad_Reader_Trans_functorReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_functorReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_functorReaderT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_functorReaderT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_functorReaderT
}

var cache_Control_Monad_Reader_Trans_distributiveReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_distributiveReaderT sync.Once
func Get_Control_Monad_Reader_Trans_distributiveReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_distributiveReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_distributiveReaderT = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_distributiveReaderT(dictDistributive_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_distributiveReaderT
}

var cache_Control_Monad_Reader_Trans_applyReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_applyReaderT sync.Once
func Get_Control_Monad_Reader_Trans_applyReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_applyReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_applyReaderT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_applyReaderT(dictApply_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_applyReaderT
}

var cache_Control_Monad_Reader_Trans_bindReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_bindReaderT sync.Once
func Get_Control_Monad_Reader_Trans_bindReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_bindReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_bindReaderT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_bindReaderT(dictBind_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_bindReaderT
}

var cache_Control_Monad_Reader_Trans_semigroupReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_semigroupReaderT sync.Once
func Get_Control_Monad_Reader_Trans_semigroupReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_semigroupReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_semigroupReaderT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_semigroupReaderT(dictApply_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_semigroupReaderT
}

var cache_Control_Monad_Reader_Trans_applicativeReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_applicativeReaderT sync.Once
func Get_Control_Monad_Reader_Trans_applicativeReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_applicativeReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_applicativeReaderT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_applicativeReaderT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_applicativeReaderT
}

var cache_Control_Monad_Reader_Trans_monadReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadReaderT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadReaderT
}

var cache_Control_Monad_Reader_Trans_monadAskReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadAskReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadAskReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadAskReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadAskReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadAskReaderT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadAskReaderT
}

var cache_Control_Monad_Reader_Trans_monadReaderReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadReaderReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadReaderReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadReaderReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadReaderReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadReaderReaderT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadReaderReaderT
}

var cache_Control_Monad_Reader_Trans_monadContReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadContReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadContReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadContReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadContReaderT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadContReaderT(dictMonadCont_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadContReaderT
}

var cache_Control_Monad_Reader_Trans_monadEffectReader gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadEffectReader sync.Once
func Get_Control_Monad_Reader_Trans_monadEffectReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadEffectReader.Do(func() {
		cache_Control_Monad_Reader_Trans_monadEffectReader = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadEffectReader(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadEffectReader
}

var cache_Control_Monad_Reader_Trans_monadRecReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadRecReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadRecReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadRecReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadRecReaderT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadRecReaderT(dictMonadRec_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadRecReaderT
}

var cache_Control_Monad_Reader_Trans_monadStateReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadStateReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadStateReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadStateReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadStateReaderT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadStateReaderT(dictMonadState_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadStateReaderT
}

var cache_Control_Monad_Reader_Trans_monadTellReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadTellReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadTellReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadTellReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadTellReaderT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadTellReaderT(dictMonadTell_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadTellReaderT
}

var cache_Control_Monad_Reader_Trans_monadWriterReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadWriterReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadWriterReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadWriterReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadWriterReaderT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadWriterReaderT(dictMonadWriter_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadWriterReaderT
}

var cache_Control_Monad_Reader_Trans_monadThrowReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadThrowReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadThrowReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadThrowReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadThrowReaderT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadThrowReaderT(dictMonadThrow_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadThrowReaderT
}

var cache_Control_Monad_Reader_Trans_monadErrorReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadErrorReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadErrorReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadErrorReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadErrorReaderT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadErrorReaderT(dictMonadError_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadErrorReaderT
}

var cache_Control_Monad_Reader_Trans_monadSTReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadSTReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadSTReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadSTReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadSTReaderT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadSTReaderT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadSTReaderT
}

var cache_Control_Monad_Reader_Trans_monoidReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monoidReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monoidReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monoidReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monoidReaderT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monoidReaderT(dictApplicative_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monoidReaderT
}

var cache_Control_Monad_Reader_Trans_altReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_altReaderT sync.Once
func Get_Control_Monad_Reader_Trans_altReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_altReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_altReaderT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_altReaderT(dictAlt_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_altReaderT
}

var cache_Control_Monad_Reader_Trans_plusReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_plusReaderT sync.Once
func Get_Control_Monad_Reader_Trans_plusReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_plusReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_plusReaderT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_plusReaderT(dictPlus_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_plusReaderT
}

var cache_Control_Monad_Reader_Trans_alternativeReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_alternativeReaderT sync.Once
func Get_Control_Monad_Reader_Trans_alternativeReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_alternativeReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_alternativeReaderT = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_alternativeReaderT(dictAlternative_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_alternativeReaderT
}

var cache_Control_Monad_Reader_Trans_monadPlusReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadPlusReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadPlusReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadPlusReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadPlusReaderT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_monadPlusReaderT(dictMonadPlus_0_box)
})
	})
	return cache_Control_Monad_Reader_Trans_monadPlusReaderT
}

func Call_Control_Monad_Reader_Trans_ReaderT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Reader_Trans_withReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_Control_Monad_Reader_Trans_runReaderT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_Reader_Trans_lift(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return x_1
}

func Call_Control_Monad_Reader_Trans_mapReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_Reader_Trans_functorReaderT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), x_1)
_ = __local_var_2_0
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(v_3, x_4))
})
})}))}
}

func Call_Control_Monad_Reader_Trans_distributiveReaderT(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveReaderT:
for {
if false { continue distributiveReaderT }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorReaderT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e), (TypeVar g)])])
functorReaderT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})})
_ = functorReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer((&Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_1_0)}
}), gopurs_runtime.Func2(func(dictFunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e), (TypeVar g), (TypeApp (TypeVar g) [(TypeVar b)])])
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Reader_Trans_distributiveReaderT(dictDistributive_0), "distribute"), dictFunctor_2)
_ = __local_var_4_3
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e), (TypeVar g), (TypeVar b)])]))
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(__local_var_5_4, x_6))
})
}), gopurs_runtime.Func3(func(dictFunctor_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_2))}, gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(r_5, e_4)
}), a_3)
})}))}
}
}

func Call_Control_Monad_Reader_Trans_applyReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorReaderT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})})
_ = functorReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
})}))}
}

func Call_Control_Monad_Reader_Trans_bindReaderT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})})
_ = functorReaderT1_2_2
// TAST (Let): applyReaderT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})})
_ = applyReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_3, a_5, r_4)
}))
})}))}
}

func Call_Control_Monad_Reader_Trans_semigroupReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorReaderT1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_1_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "map"), x_2)
_ = __local_var_3_3
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Apply(v_4, x_5))
})
})})
_ = functorReaderT1_1_1
// TAST (Let): applyReaderT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar s), (TypeVar m)])])
applyReaderT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_1_1)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
})})
_ = applyReaderT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyReaderT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_5 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_4_5 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyReaderT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), __local_var_4_5, a_5), b_6)
})}))}
})
}

func Call_Control_Monad_Reader_Trans_applicativeReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})})
_ = functorReaderT1_2_2
// TAST (Let): applyReaderT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})})
_ = applyReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_2)
_ = __local_var_3_5
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_5
})
})}))}
}

func Call_Control_Monad_Reader_Trans_monadReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorReaderT1_3_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_3_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), x_4)
_ = __local_var_5_6
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(v_6, x_7))
})
})})
_ = functorReaderT1_3_4
// TAST (Let): applyReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_2_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_3_4)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})})
_ = applyReaderT1_2_2
// TAST (Let): applicativeReaderT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_2_2)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), x_3)
_ = __local_var_4_7
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_7
})
})})
_ = applicativeReaderT1_1_0
// TAST (Let): __local_var_2_9 shape=App(Other) bindingType=Any
__local_var_2_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_9
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_9, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): functorReaderT1_4_12 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_12 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_14 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "map"), x_5)
_ = __local_var_6_14
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_14, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_12
// TAST (Let): applyReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_10 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_10
// TAST (Let): bindReaderT1_2_8 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_2_8 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_10)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_9, "bind"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_5, a_7, r_6)
}))
})})
_ = bindReaderT1_2_8
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_2_8)}
})}))}
}

func Call_Control_Monad_Reader_Trans_monadAskReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=Any
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): __local_var_3_6 shape=App(Other) bindingType=Any
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_6
// TAST (Let): functorReaderT1_3_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_3_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_6, "map"), x_4)
_ = __local_var_5_7
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_7, gopurs_runtime.Apply(v_6, x_7))
})
})})
_ = functorReaderT1_3_5
// TAST (Let): applyReaderT1_2_3 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_2_3 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_3_5)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})})
_ = applyReaderT1_2_3
// TAST (Let): applicativeReaderT1_1_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_1_1 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_2_3)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "pure"), x_3)
_ = __local_var_4_8
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_8
})
})})
_ = applicativeReaderT1_1_1
// TAST (Let): __local_var_2_10 shape=App(Other) bindingType=Any
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_10
// TAST (Let): __local_var_3_12 shape=App(Other) bindingType=Any
__local_var_3_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_10, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_12
// TAST (Let): __local_var_4_14 shape=App(Other) bindingType=Any
__local_var_4_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_14
// TAST (Let): functorReaderT1_4_13 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_13 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_15 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_14, "map"), x_5)
_ = __local_var_6_15
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_15, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_13
// TAST (Let): applyReaderT1_3_11 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_11 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_13)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_12, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_11
// TAST (Let): bindReaderT1_2_9 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_2_9 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_11)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_10, "bind"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_5, a_7, r_6)
}))
})})
_ = bindReaderT1_2_9
// TAST (Let): monadReaderT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_1_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_2_9)}
})})
_ = monadReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_1_0)}
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")}))}
}

func Call_Control_Monad_Reader_Trans_monadReaderReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=Any
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): __local_var_2_5 shape=App(Other) bindingType=Any
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_5
// TAST (Let): __local_var_3_7 shape=App(Other) bindingType=Any
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_7
// TAST (Let): functorReaderT1_3_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_3_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_5_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_7, "map"), x_4)
_ = __local_var_5_8
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_8, gopurs_runtime.Apply(v_6, x_7))
})
})})
_ = functorReaderT1_3_6
// TAST (Let): applyReaderT1_2_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_2_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_3_6)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_5, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})})
_ = applyReaderT1_2_4
// TAST (Let): applicativeReaderT1_1_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_1_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_2_4)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_4_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_3, "pure"), x_3)
_ = __local_var_4_9
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_9
})
})})
_ = applicativeReaderT1_1_2
// TAST (Let): __local_var_2_11 shape=App(Other) bindingType=Any
__local_var_2_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_11
// TAST (Let): __local_var_3_13 shape=App(Other) bindingType=Any
__local_var_3_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_13
// TAST (Let): __local_var_4_15 shape=App(Other) bindingType=Any
__local_var_4_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_15
// TAST (Let): functorReaderT1_4_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "map"), x_5)
_ = __local_var_6_16
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_16, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_14
// TAST (Let): applyReaderT1_3_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_14)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_12
// TAST (Let): bindReaderT1_2_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_2_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_12)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_11, "bind"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_5, a_7, r_6)
}))
})})
_ = bindReaderT1_2_10
// TAST (Let): monadReaderT1_1_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_1_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_1_2)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_2_10)}
})})
_ = monadReaderT1_1_1
// TAST (Let): monadAskReaderT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r), (TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadAskReaderT1_1_0 := (&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_1_1)}
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")})
_ = monadAskReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskReaderT1_1_0)}
}), Get_Control_Monad_Reader_Trans_withReaderT()}))}
}

func Call_Control_Monad_Reader_Trans_monadContReaderT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorReaderT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), x_5)
_ = __local_var_6_8
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_6
// TAST (Let): applyReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_4
// TAST (Let): applicativeReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_4)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_9
})
})})
_ = applicativeReaderT1_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): functorReaderT1_5_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "map"), x_6)
_ = __local_var_7_16
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_16, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_14
// TAST (Let): applyReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_12
// TAST (Let): bindReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_3_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "bind"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, r_7)
}))
})})
_ = bindReaderT1_3_10
// TAST (Let): monadReaderT1_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_3_10)}
})})
_ = monadReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_1_0)}
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_17 shape=App(Other) bindingType=(TypeApp (TypeVar m) [(TypeVar b)])
__local_var_6_17 := gopurs_runtime.Apply(c_4, x_5)
_ = __local_var_6_17
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_17
})
}), r_3)
}))
})}))}
}

func Call_Control_Monad_Reader_Trans_monadEffectReader(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorReaderT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), x_5)
_ = __local_var_6_8
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_6
// TAST (Let): applyReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_4
// TAST (Let): applicativeReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_4)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_9
})
})})
_ = applicativeReaderT1_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): functorReaderT1_5_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "map"), x_6)
_ = __local_var_7_16
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_16, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_14
// TAST (Let): applyReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_12
// TAST (Let): bindReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_3_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "bind"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, r_7)
}))
})})
_ = bindReaderT1_3_10
// TAST (Let): monadReaderT1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_3_10)}
})})
_ = monadReaderT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_2_1)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_17 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_3)
_ = __local_var_4_17
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_17
})
})}))}
}

func Call_Control_Monad_Reader_Trans_monadRecReaderT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(ADT ["Control","Monad","Rec","Class","Step"] [(TypeVar a), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","Rec","Class","Step"] [(TypeVar a), (TypeVar b)])]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=Any
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_9
// TAST (Let): functorReaderT1_6_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_6_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "map"), x_7)
_ = __local_var_8_10
return gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_10, gopurs_runtime.Apply(v_9, x_10))
})
})})
_ = functorReaderT1_6_8
// TAST (Let): applyReaderT1_5_6 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_5_6 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_6_8)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_7, "apply"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Apply(v1_8, r_9))
})})
_ = applyReaderT1_5_6
// TAST (Let): applicativeReaderT1_4_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_4_4 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_5_6)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "pure"), x_6)
_ = __local_var_7_11
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_11
})
})})
_ = applicativeReaderT1_4_4
// TAST (Let): __local_var_5_13 shape=App(Other) bindingType=Any
__local_var_5_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_13
// TAST (Let): __local_var_6_15 shape=App(Other) bindingType=Any
__local_var_6_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_13, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_15
// TAST (Let): __local_var_7_17 shape=App(Other) bindingType=Any
__local_var_7_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_17
// TAST (Let): functorReaderT1_7_16 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_7_16 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_18 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_9_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "map"), x_8)
_ = __local_var_9_18
return gopurs_runtime.Func2(func(v_10 gopurs_runtime.Value, x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_18, gopurs_runtime.Apply(v_10, x_11))
})
})})
_ = functorReaderT1_7_16
// TAST (Let): applyReaderT1_6_14 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_6_14 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_7_16)}
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_15, "apply"), gopurs_runtime.Apply(v_8, r_10), gopurs_runtime.Apply(v1_9, r_10))
})})
_ = applyReaderT1_6_14
// TAST (Let): bindReaderT1_5_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_5_12 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_6_14)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_13, "bind"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_8, a_10, r_9)
}))
})})
_ = bindReaderT1_5_12
// TAST (Let): monadReaderT1_4_3 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_4_3 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_4_4)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_5_12)}
})})
_ = monadReaderT1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_4_3)}
}), gopurs_runtime.Func3(func(k_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(k_5, a_prime__8, r_7), pure_3_2)
}), a_6)
})}))}
}

func Call_Control_Monad_Reader_Trans_monadStateReaderT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorReaderT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), x_5)
_ = __local_var_6_8
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_6
// TAST (Let): applyReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_4
// TAST (Let): applicativeReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_4)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_9
})
})})
_ = applicativeReaderT1_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): functorReaderT1_5_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "map"), x_6)
_ = __local_var_7_16
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_16, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_14
// TAST (Let): applyReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_12
// TAST (Let): bindReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_3_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "bind"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, r_7)
}))
})})
_ = bindReaderT1_3_10
// TAST (Let): monadReaderT1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_3_10)}
})})
_ = monadReaderT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_2_1)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_17 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_3)
_ = __local_var_4_17
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_17
})
})}))}
}

func Call_Control_Monad_Reader_Trans_monadTellReaderT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): __local_var_5_8 shape=App(Other) bindingType=Any
__local_var_5_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_8
// TAST (Let): functorReaderT1_5_7 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_7 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_9 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "map"), x_6)
_ = __local_var_7_9
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_9, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_7
// TAST (Let): applyReaderT1_4_5 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_7)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_6, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_5
// TAST (Let): applicativeReaderT1_3_3 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_3_3 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_5)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_10 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_6_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "pure"), x_5)
_ = __local_var_6_10
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_10
})
})})
_ = applicativeReaderT1_3_3
// TAST (Let): __local_var_4_12 shape=App(Other) bindingType=Any
__local_var_4_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_12
// TAST (Let): __local_var_5_14 shape=App(Other) bindingType=Any
__local_var_5_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_12, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_14
// TAST (Let): __local_var_6_16 shape=App(Other) bindingType=Any
__local_var_6_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_16
// TAST (Let): functorReaderT1_6_15 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_6_15 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_17 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_8_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_16, "map"), x_7)
_ = __local_var_8_17
return gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_17, gopurs_runtime.Apply(v_9, x_10))
})
})})
_ = functorReaderT1_6_15
// TAST (Let): applyReaderT1_5_13 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_5_13 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_6_15)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_14, "apply"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Apply(v1_8, r_9))
})})
_ = applyReaderT1_5_13
// TAST (Let): bindReaderT1_4_11 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_4_11 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_5_13)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_12, "bind"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_7, a_9, r_8)
}))
})})
_ = bindReaderT1_4_11
// TAST (Let): monadReaderT1_3_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_3_2 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_3_3)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_4_11)}
})})
_ = monadReaderT1_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_3_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_18 shape=App(Other) bindingType=(TypeVar c)
__local_var_5_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_4)
_ = __local_var_5_18
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_18
})
})}))}
}

func Call_Control_Monad_Reader_Trans_monadWriterReaderT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): Monoid0_1_0 shape=App(Other) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar w)])
Monoid0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{}))
_ = Monoid0_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): Monad1_3_3 shape=App(Other) bindingType=Any
Monad1_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_3_3
// TAST (Let): Semigroup0_4_4 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_4
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_3_3, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=Any
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_6_9
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=Any
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_11
// TAST (Let): functorReaderT1_7_10 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_7_10 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_12 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_9_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "map"), x_8)
_ = __local_var_9_12
return gopurs_runtime.Func2(func(v_10 gopurs_runtime.Value, x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_12, gopurs_runtime.Apply(v_10, x_11))
})
})})
_ = functorReaderT1_7_10
// TAST (Let): applyReaderT1_6_8 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_6_8 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_7_10)}
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_9, "apply"), gopurs_runtime.Apply(v_8, r_10), gopurs_runtime.Apply(v1_9, r_10))
})})
_ = applyReaderT1_6_8
// TAST (Let): applicativeReaderT1_5_6 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_5_6 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_6_8)}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_13 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "pure"), x_7)
_ = __local_var_8_13
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_8_13
})
})})
_ = applicativeReaderT1_5_6
// TAST (Let): __local_var_6_15 shape=App(Other) bindingType=Any
__local_var_6_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_3_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_15
// TAST (Let): __local_var_7_17 shape=App(Other) bindingType=Any
__local_var_7_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_7_17
// TAST (Let): __local_var_8_19 shape=App(Other) bindingType=Any
__local_var_8_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_19
// TAST (Let): functorReaderT1_8_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_8_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_20 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_10_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_19, "map"), x_9)
_ = __local_var_10_20
return gopurs_runtime.Func2(func(v_11 gopurs_runtime.Value, x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_20, gopurs_runtime.Apply(v_11, x_12))
})
})})
_ = functorReaderT1_8_18
// TAST (Let): applyReaderT1_7_16 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_7_16 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_8_18)}
}), gopurs_runtime.Func3(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_17, "apply"), gopurs_runtime.Apply(v_9, r_11), gopurs_runtime.Apply(v1_10, r_11))
})})
_ = applyReaderT1_7_16
// TAST (Let): bindReaderT1_6_14 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_6_14 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_7_16)}
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_15, "bind"), gopurs_runtime.Apply(v_8, r_10), gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_9, a_11, r_10)
}))
})})
_ = bindReaderT1_6_14
// TAST (Let): monadReaderT1_5_5 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_5_5 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_5_6)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_6_14)}
})})
_ = monadReaderT1_5_5
// TAST (Let): monadTellReaderT1_2_1 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w), (TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadTellReaderT1_2_1 := (&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_5_5)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_4_4)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_21 shape=App(Other) bindingType=(TypeVar c)
__local_var_7_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "tell"), x_6)
_ = __local_var_7_21
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_21
})
})})
_ = monadTellReaderT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellReaderT1_2_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Monoid0_1_0)}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(v_3, x_4))
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply(v_3, x_4))
})}))}
}

func Call_Control_Monad_Reader_Trans_monadThrowReaderT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorReaderT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), x_5)
_ = __local_var_6_8
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_6
// TAST (Let): applyReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_4
// TAST (Let): applicativeReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_4)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_9
})
})})
_ = applicativeReaderT1_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): functorReaderT1_5_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "map"), x_6)
_ = __local_var_7_16
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_16, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_14
// TAST (Let): applyReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_12
// TAST (Let): bindReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_3_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "bind"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, r_7)
}))
})})
_ = bindReaderT1_3_10
// TAST (Let): monadReaderT1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_3_10)}
})})
_ = monadReaderT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_2_1)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_17 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), x_3)
_ = __local_var_4_17
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_17
})
})}))}
}

func Call_Control_Monad_Reader_Trans_monadErrorReaderT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Monad0_2_2 shape=App(Other) bindingType=Any
Monad0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_2
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_2, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=Any
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_9
// TAST (Let): functorReaderT1_5_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_10 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "map"), x_6)
_ = __local_var_7_10
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_10, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_8
// TAST (Let): applyReaderT1_4_6 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_6 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_8)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_7, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_6
// TAST (Let): applicativeReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_3_4 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_6)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_11 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_6_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "pure"), x_5)
_ = __local_var_6_11
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_11
})
})})
_ = applicativeReaderT1_3_4
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_2, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): __local_var_6_17 shape=App(Other) bindingType=Any
__local_var_6_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_17
// TAST (Let): functorReaderT1_6_16 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_6_16 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_18 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_8_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_17, "map"), x_7)
_ = __local_var_8_18
return gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_18, gopurs_runtime.Apply(v_9, x_10))
})
})})
_ = functorReaderT1_6_16
// TAST (Let): applyReaderT1_5_14 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_5_14 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_6_16)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_15, "apply"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Apply(v1_8, r_9))
})})
_ = applyReaderT1_5_14
// TAST (Let): bindReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_4_12 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "bind"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_7, a_9, r_8)
}))
})})
_ = bindReaderT1_4_12
// TAST (Let): monadReaderT1_3_3 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_3_3 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_3_4)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_4_12)}
})})
_ = monadReaderT1_3_3
// TAST (Let): monadThrowReaderT1_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e), (TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadThrowReaderT1_1_0 := (&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_3_3)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_19 shape=App(Other) bindingType=(TypeVar c)
__local_var_5_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "throwError"), x_4)
_ = __local_var_5_19
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_19
})
})})
_ = monadThrowReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowReaderT1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, r_4)
}))
})}))}
}

func Call_Control_Monad_Reader_Trans_monadSTReaderT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorReaderT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), x_5)
_ = __local_var_6_8
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_6
// TAST (Let): applyReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_4
// TAST (Let): applicativeReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_4)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_9
})
})})
_ = applicativeReaderT1_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): functorReaderT1_5_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "map"), x_6)
_ = __local_var_7_16
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_16, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_14
// TAST (Let): applyReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_12
// TAST (Let): bindReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_3_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "bind"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, r_7)
}))
})})
_ = bindReaderT1_3_10
// TAST (Let): monadReaderT1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_2_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_3_10)}
})})
_ = monadReaderT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_2_1)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_17 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_3)
_ = __local_var_4_17
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_17
})
})}))}
}

func Call_Control_Monad_Reader_Trans_monoidReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=Any
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): functorReaderT1_2_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_2_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "map"), x_3)
_ = __local_var_4_5
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(v_5, x_6))
})
})})
_ = functorReaderT1_2_3
// TAST (Let): applyReaderT1_1_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_1_1 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_2_3)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})})
_ = applyReaderT1_1_1
// TAST (Let): applicativeReaderT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar s), (TypeVar m)])])
applicativeReaderT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_1_1)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_6 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_2)
_ = __local_var_3_6
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_6
})
})})
_ = applicativeReaderT1_1_0
// TAST (Let): __local_var_2_8 shape=App(Other) bindingType=Any
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_8
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_8, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): functorReaderT1_3_10 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_3_10 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_12 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_5_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "map"), x_4)
_ = __local_var_5_12
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_12, gopurs_runtime.Apply(v_6, x_7))
})
})})
_ = functorReaderT1_3_10
// TAST (Let): applyReaderT1_3_9 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar s), (TypeVar m)])])
applyReaderT1_3_9 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_3_10)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_8, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})})
_ = applyReaderT1_3_9
// TAST (Let): semigroupReaderT1__193435443_2_7 shape=Let(Let(Abs(LitRecord))) bindingType=Any
semigroupReaderT1__193435443_2_7 := gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_13 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyReaderT1_3_9.V0), gopurs_runtime.Value{}))
_ = Functor0_5_13
// TAST (Let): __local_var_6_14 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_6_14 := gopurs_runtime.RecordGet(dictSemigroup_4, "append")
_ = __local_var_6_14
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyReaderT1_3_9.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_13.V0), __local_var_6_14, a_7), b_8)
})}))}
})
_ = semigroupReaderT1__193435443_2_7
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupReaderT2_4_15 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar s), (TypeVar m), (TypeVar a)])])
semigroupReaderT2_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupReaderT1__193435443_2_7, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})))
_ = semigroupReaderT2_4_15
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupReaderT2_4_15)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeReaderT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))}))}
})
}

func Call_Control_Monad_Reader_Trans_altReaderT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorReaderT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})})
_ = functorReaderT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
})}))}
}

func Call_Control_Monad_Reader_Trans_plusReaderT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})})
_ = functorReaderT1_2_2
// TAST (Let): altReaderT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
altReaderT1_1_0 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "alt"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})})
_ = altReaderT1_1_0
// TAST (Let): __local_var_2_5 shape=Other bindingType=Any
__local_var_2_5 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = __local_var_2_5
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altReaderT1_1_0)}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_5
})}))}
}

func Call_Control_Monad_Reader_Trans_alternativeReaderT(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorReaderT1_3_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_3_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), x_4)
_ = __local_var_5_6
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(v_6, x_7))
})
})})
_ = functorReaderT1_3_4
// TAST (Let): applyReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_2_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_3_4)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})})
_ = applyReaderT1_2_2
// TAST (Let): applicativeReaderT1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_2_2)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), x_3)
_ = __local_var_4_7
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_7
})
})})
_ = applicativeReaderT1_1_0
// TAST (Let): __local_var_2_9 shape=App(Other) bindingType=Any
__local_var_2_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_9
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_9, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): functorReaderT1_4_12 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_12 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_14 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "map"), x_5)
_ = __local_var_6_14
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_14, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_12
// TAST (Let): altReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
altReaderT1_3_10 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "alt"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = altReaderT1_3_10
// TAST (Let): __local_var_4_15 shape=Other bindingType=Any
__local_var_4_15 := gopurs_runtime.RecordGet(__local_var_2_9, "empty")
_ = __local_var_4_15
// TAST (Let): plusReaderT1_2_8 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
plusReaderT1_2_8 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altReaderT1_3_10)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_15
})})
_ = plusReaderT1_2_8
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusReaderT1_2_8)}
})}))}
}

func Call_Control_Monad_Reader_Trans_monadPlusReaderT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorReaderT1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), x_5)
_ = __local_var_6_8
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, gopurs_runtime.Apply(v_7, x_8))
})
})})
_ = functorReaderT1_4_6
// TAST (Let): applyReaderT1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})})
_ = applyReaderT1_3_4
// TAST (Let): applicativeReaderT1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_3_4)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_9
})
})})
_ = applicativeReaderT1_2_2
// TAST (Let): __local_var_3_11 shape=App(Other) bindingType=Any
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): __local_var_5_15 shape=App(Other) bindingType=Any
__local_var_5_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_15
// TAST (Let): functorReaderT1_5_14 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_14 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_16 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_15, "map"), x_6)
_ = __local_var_7_16
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_16, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_14
// TAST (Let): applyReaderT1_4_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_12 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_14)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_12
// TAST (Let): bindReaderT1_3_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
bindReaderT1_3_10 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_12)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "bind"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, r_7)
}))
})})
_ = bindReaderT1_3_10
// TAST (Let): monadReaderT1_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
monadReaderT1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindReaderT1_3_10)}
})})
_ = monadReaderT1_1_0
// TAST (Let): __local_var_2_18 shape=App(Other) bindingType=Any
__local_var_2_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{})
_ = __local_var_2_18
// TAST (Let): __local_var_3_20 shape=App(Other) bindingType=Any
__local_var_3_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_18, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_20
// TAST (Let): __local_var_4_22 shape=App(Other) bindingType=Any
__local_var_4_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_20, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_22
// TAST (Let): __local_var_5_24 shape=App(Other) bindingType=Any
__local_var_5_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_22, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_24
// TAST (Let): functorReaderT1_5_23 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_5_23 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_25 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_7_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_24, "map"), x_6)
_ = __local_var_7_25
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_25, gopurs_runtime.Apply(v_8, x_9))
})
})})
_ = functorReaderT1_5_23
// TAST (Let): applyReaderT1_4_21 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applyReaderT1_4_21 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_5_23)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_22, "apply"), gopurs_runtime.Apply(v_6, r_8), gopurs_runtime.Apply(v1_7, r_8))
})})
_ = applyReaderT1_4_21
// TAST (Let): applicativeReaderT1_3_19 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
applicativeReaderT1_3_19 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_4_21)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_26 shape=App(Other) bindingType=(ForAll [a] (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)])))
__local_var_6_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_20, "pure"), x_5)
_ = __local_var_6_26
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_26
})
})})
_ = applicativeReaderT1_3_19
// TAST (Let): __local_var_4_28 shape=App(Other) bindingType=Any
__local_var_4_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_18, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_28
// TAST (Let): __local_var_5_30 shape=App(Other) bindingType=Any
__local_var_5_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_28, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_30
// TAST (Let): __local_var_6_32 shape=App(Other) bindingType=Any
__local_var_6_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_30, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_32
// TAST (Let): functorReaderT1_6_31 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_6_31 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_33 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_8_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_32, "map"), x_7)
_ = __local_var_8_33
return gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_33, gopurs_runtime.Apply(v_9, x_10))
})
})})
_ = functorReaderT1_6_31
// TAST (Let): altReaderT1_5_29 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
altReaderT1_5_29 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_6_31)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_30, "alt"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Apply(v1_8, r_9))
})})
_ = altReaderT1_5_29
// TAST (Let): __local_var_6_34 shape=Other bindingType=Any
__local_var_6_34 := gopurs_runtime.RecordGet(__local_var_4_28, "empty")
_ = __local_var_6_34
// TAST (Let): plusReaderT1_4_27 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
plusReaderT1_4_27 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altReaderT1_5_29)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_34
})})
_ = plusReaderT1_4_27
// TAST (Let): alternativeReaderT1_2_17 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
alternativeReaderT1_2_17 := (&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeReaderT1_3_19)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusReaderT1_4_27)}
})})
_ = alternativeReaderT1_2_17
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeReaderT1_2_17)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadReaderT1_1_0)}
})}))}
}


