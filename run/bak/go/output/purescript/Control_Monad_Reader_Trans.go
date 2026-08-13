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
		cache_Control_Monad_Reader_Trans_newtypeReaderT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Monad_Reader_Trans_newtypeReaderT
}

var cache_Control_Monad_Reader_Trans_monadTransReaderT gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_monadTransReaderT sync.Once
func Get_Control_Monad_Reader_Trans_monadTransReaderT() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_monadTransReaderT.Do(func() {
		cache_Control_Monad_Reader_Trans_monadTransReaderT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
})
}))
	})
	return cache_Control_Monad_Reader_Trans_monadTransReaderT
}

var cache_Control_Monad_Reader_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_lift sync.Once
func Get_Control_Monad_Reader_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_lift.Do(func() {
		cache_Control_Monad_Reader_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift")
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

var cache_Control_Monad_Reader_Trans_mapReaderT__552640602 gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_mapReaderT__552640602 sync.Once
func Get_Control_Monad_Reader_Trans_mapReaderT__552640602() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_mapReaderT__552640602.Do(func() {
		cache_Control_Monad_Reader_Trans_mapReaderT__552640602 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_mapReaderT__552640602(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_mapReaderT__552640602
}

var cache_Control_Monad_Reader_Trans_mapReaderT__3395617690 gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_mapReaderT__3395617690 sync.Once
func Get_Control_Monad_Reader_Trans_mapReaderT__3395617690() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_mapReaderT__3395617690.Do(func() {
		cache_Control_Monad_Reader_Trans_mapReaderT__3395617690 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_mapReaderT__3395617690(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_mapReaderT__3395617690
}

var cache_Control_Monad_Reader_Trans_mapReaderT__3691274100 gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_mapReaderT__3691274100 sync.Once
func Get_Control_Monad_Reader_Trans_mapReaderT__3691274100() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_mapReaderT__3691274100.Do(func() {
		cache_Control_Monad_Reader_Trans_mapReaderT__3691274100 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_mapReaderT__3691274100(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_mapReaderT__3691274100
}

var cache_Control_Monad_Reader_Trans_mapReaderT__437052724 gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_mapReaderT__437052724 sync.Once
func Get_Control_Monad_Reader_Trans_mapReaderT__437052724() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_mapReaderT__437052724.Do(func() {
		cache_Control_Monad_Reader_Trans_mapReaderT__437052724 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_mapReaderT__437052724(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_mapReaderT__437052724
}

var cache_Control_Monad_Reader_Trans_withReaderT__552640602 gopurs_runtime.Value
var once_Control_Monad_Reader_Trans_withReaderT__552640602 sync.Once
func Get_Control_Monad_Reader_Trans_withReaderT__552640602() gopurs_runtime.Value {
	once_Control_Monad_Reader_Trans_withReaderT__552640602.Do(func() {
		cache_Control_Monad_Reader_Trans_withReaderT__552640602 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Trans_withReaderT__552640602(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_Trans_withReaderT__552640602
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
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), x_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(v_3, x_4))
})
})
}))
}

func Call_Control_Monad_Reader_Trans_distributiveReaderT(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveReaderT:
for {
if false { continue distributiveReaderT }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorReaderT1_1_0 -> gopurs_runtime.Value
functorReaderT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})
}))
_ = functorReaderT1_1_0
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_0
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Reader_Trans_distributiveReaderT(dictDistributive_0), "distribute"), dictFunctor_2)
_ = __local_var_4_3
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(__local_var_5_4, x_6))
})
})
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_2))}, gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(r_5, e_4)
}), a_3)
})
})
}))
}
}

func Call_Control_Monad_Reader_Trans_applyReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorReaderT1_1_0 -> gopurs_runtime.Value
functorReaderT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})
}))
_ = functorReaderT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
})
})
}))
}

func Call_Control_Monad_Reader_Trans_bindReaderT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 -> gopurs_runtime.Value
functorReaderT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_2
// TAST (Let): applyReaderT1_1_0 -> gopurs_runtime.Value
applyReaderT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})
})
}))
_ = applyReaderT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_3, a_5, r_4)
}))
})
})
}))
}

func Call_Control_Monad_Reader_Trans_semigroupReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorReaderT1_1_1 -> gopurs_runtime.Value
functorReaderT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "map"), x_2)
_ = __local_var_3_3
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Apply(v_4, x_5))
})
})
}))
_ = functorReaderT1_1_1
// TAST (Let): applyReaderT1_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
applyReaderT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
})
})
})))
_ = applyReaderT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_4 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyReaderT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_5
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyReaderT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), __local_var_4_5, a_5), b_6)
})
}))
})
}

func Call_Control_Monad_Reader_Trans_applicativeReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 -> gopurs_runtime.Value
functorReaderT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_2
// TAST (Let): applyReaderT1_1_0 -> gopurs_runtime.Value
applyReaderT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})
})
}))
_ = applyReaderT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_2)
_ = __local_var_3_5
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_5
})
}))
}

func Call_Control_Monad_Reader_Trans_monadReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorReaderT1_3_4 -> gopurs_runtime.Value
functorReaderT1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), x_4)
_ = __local_var_5_6
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(v_6, x_7))
})
})
}))
_ = functorReaderT1_3_4
// TAST (Let): applyReaderT1_2_2 -> gopurs_runtime.Value
applyReaderT1_2_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_4
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})
})
}))
_ = applyReaderT1_2_2
// TAST (Let): applicativeReaderT1_1_0 -> gopurs_runtime.Value
applicativeReaderT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_2_2
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 -> gopurs_runtime.Value
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), x_3)
_ = __local_var_4_7
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_7
})
}))
_ = applicativeReaderT1_1_0
// TAST (Let): bindReaderT1_2_8 -> gopurs_runtime.Value
bindReaderT1_2_8 := Call_Control_Monad_Reader_Trans_bindReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindReaderT1_2_8
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindReaderT1_2_8
}))
}

func Call_Control_Monad_Reader_Trans_monadAskReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadReaderT1_1_0 -> gopurs_runtime.Value
monadReaderT1_1_0 := Call_Control_Monad_Reader_Trans_monadReaderT(dictMonad_0)
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"))
}

func Call_Control_Monad_Reader_Trans_monadReaderReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadReaderT1_1_1 -> gopurs_runtime.Value
monadReaderT1_1_1 := Call_Control_Monad_Reader_Trans_monadReaderT(dictMonad_0)
_ = monadReaderT1_1_1
// TAST (Let): monadAskReaderT1_1_0 -> gopurs_runtime.Value
monadAskReaderT1_1_0 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_1
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"))
_ = monadAskReaderT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskReaderT1_1_0
}), Get_Control_Monad_Reader_Trans_withReaderT())
}

func Call_Control_Monad_Reader_Trans_monadContReaderT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): monadReaderT1_1_0 -> gopurs_runtime.Value
monadReaderT1_1_0 := Call_Control_Monad_Reader_Trans_monadReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_1 -> gopurs_runtime.Value
__local_var_6_1 := gopurs_runtime.Apply(c_4, x_5)
_ = __local_var_6_1
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_1
})
}), r_3)
}))
})
}))
}

func Call_Control_Monad_Reader_Trans_monadEffectReader(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadReaderT1_2_1 -> gopurs_runtime.Value
monadReaderT1_2_1 := Call_Control_Monad_Reader_Trans_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_Control_Monad_Reader_Trans_monadRecReaderT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): monadReaderT1_4_3 -> gopurs_runtime.Value
monadReaderT1_4_3 := Call_Control_Monad_Reader_Trans_monadReaderT(Monad0_1_0)
_ = monadReaderT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_4_3
}), gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(k_5, a_prime_8, r_7), pure_3_2)
}), a_6)
})
})
}))
}

func Call_Control_Monad_Reader_Trans_monadStateReaderT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadReaderT1_2_1 -> gopurs_runtime.Value
monadReaderT1_2_1 := Call_Control_Monad_Reader_Trans_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_4))
}))
}

func Call_Control_Monad_Reader_Trans_monadTellReaderT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 -> gopurs_runtime.Value
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
// TAST (Let): monadReaderT1_3_2 -> gopurs_runtime.Value
monadReaderT1_3_2 := Call_Control_Monad_Reader_Trans_monadReaderT(Monad1_1_0)
_ = monadReaderT1_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
}

func Call_Control_Monad_Reader_Trans_monadWriterReaderT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): Monoid0_1_0 -> gopurs_runtime.Value
Monoid0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_1_0
// TAST (Let): monadTellReaderT1_2_1 -> gopurs_runtime.Value
monadTellReaderT1_2_1 := Call_Control_Monad_Reader_Trans_monadTellReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{}))
_ = monadTellReaderT1_2_1
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(v_3, x_4))
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_Control_Monad_Reader_Trans_monadThrowReaderT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadReaderT1_2_1 -> gopurs_runtime.Value
monadReaderT1_2_1 := Call_Control_Monad_Reader_Trans_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), x_4))
}))
}

func Call_Control_Monad_Reader_Trans_monadErrorReaderT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): monadThrowReaderT1_1_0 -> gopurs_runtime.Value
monadThrowReaderT1_1_0 := Call_Control_Monad_Reader_Trans_monadThrowReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowReaderT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowReaderT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, r_4)
}))
})
})
}))
}

func Call_Control_Monad_Reader_Trans_monadSTReaderT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadReaderT1_2_1 -> gopurs_runtime.Value
monadReaderT1_2_1 := Call_Control_Monad_Reader_Trans_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_Reader_Trans_monoidReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): functorReaderT1_2_3 -> gopurs_runtime.Value
functorReaderT1_2_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "map"), x_3)
_ = __local_var_4_5
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_3
// TAST (Let): applyReaderT1_1_1 -> gopurs_runtime.Value
applyReaderT1_1_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_3
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})
})
}))
_ = applyReaderT1_1_1
// TAST (Let): applicativeReaderT1_1_0 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
applicativeReaderT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_1_1
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_2)
_ = __local_var_3_6
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_6
})
})))
_ = applicativeReaderT1_1_0
// TAST (Let): semigroupReaderT1_2_7 -> gopurs_runtime.Value
semigroupReaderT1_2_7 := Call_Control_Monad_Reader_Trans_semigroupReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = semigroupReaderT1_2_7
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupReaderT2_4_8 -> gopurs_runtime.Value
semigroupReaderT2_4_8 := gopurs_runtime.Apply(semigroupReaderT1_2_7, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupReaderT2_4_8
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupReaderT2_4_8
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeReaderT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty")))
})
}

func Call_Control_Monad_Reader_Trans_altReaderT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorReaderT1_1_0 -> gopurs_runtime.Value
functorReaderT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})
}))
_ = functorReaderT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
})
})
}))
}

func Call_Control_Monad_Reader_Trans_plusReaderT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 -> gopurs_runtime.Value
functorReaderT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_2
// TAST (Let): altReaderT1_1_0 -> gopurs_runtime.Value
altReaderT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "alt"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})
})
}))
_ = altReaderT1_1_0
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = __local_var_2_5
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_5
}))
}

func Call_Control_Monad_Reader_Trans_alternativeReaderT(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorReaderT1_3_4 -> gopurs_runtime.Value
functorReaderT1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), x_4)
_ = __local_var_5_6
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(v_6, x_7))
})
})
}))
_ = functorReaderT1_3_4
// TAST (Let): applyReaderT1_2_2 -> gopurs_runtime.Value
applyReaderT1_2_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_4
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})
})
}))
_ = applyReaderT1_2_2
// TAST (Let): applicativeReaderT1_1_0 -> gopurs_runtime.Value
applicativeReaderT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_2_2
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 -> gopurs_runtime.Value
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), x_3)
_ = __local_var_4_7
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_7
})
}))
_ = applicativeReaderT1_1_0
// TAST (Let): __local_var_2_9 -> gopurs_runtime.Value
__local_var_2_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_9
// TAST (Let): __local_var_3_11 -> gopurs_runtime.Value
__local_var_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_9, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_3_11
// TAST (Let): __local_var_4_13 -> gopurs_runtime.Value
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): functorReaderT1_4_12 -> gopurs_runtime.Value
functorReaderT1_4_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_14 -> gopurs_runtime.Value
__local_var_6_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "map"), x_5)
_ = __local_var_6_14
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_14, gopurs_runtime.Apply(v_7, x_8))
})
})
}))
_ = functorReaderT1_4_12
// TAST (Let): altReaderT1_3_10 -> gopurs_runtime.Value
altReaderT1_3_10 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_12
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_11, "alt"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
})
})
}))
_ = altReaderT1_3_10
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.RecordGet(__local_var_2_9, "empty")
_ = __local_var_4_15
// TAST (Let): plusReaderT1_2_8 -> gopurs_runtime.Value
plusReaderT1_2_8 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_3_10
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_15
}))
_ = plusReaderT1_2_8
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusReaderT1_2_8
}))
}

func Call_Control_Monad_Reader_Trans_monadPlusReaderT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
// TAST (Let): monadReaderT1_1_0 -> gopurs_runtime.Value
monadReaderT1_1_0 := Call_Control_Monad_Reader_Trans_monadReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
// TAST (Let): alternativeReaderT1_2_1 -> gopurs_runtime.Value
alternativeReaderT1_2_1 := Call_Control_Monad_Reader_Trans_alternativeReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeReaderT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}))
}

func Call_Control_Monad_Reader_Trans_mapReaderT__552640602(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_Reader_Trans_mapReaderT__3395617690(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_Reader_Trans_mapReaderT__3691274100(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_Reader_Trans_mapReaderT__437052724(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_Reader_Trans_withReaderT__552640602(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}


