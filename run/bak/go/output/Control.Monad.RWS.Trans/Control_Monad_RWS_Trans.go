package Control_Monad_RWS_Trans

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_Error_Class "gopurs/output/Control.Monad.Error.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Monad_Trans_Class "gopurs/output/Control.Monad.Trans.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_RWSResult gopurs_runtime.Value
var once_RWSResult sync.Once
func Get_RWSResult() gopurs_runtime.Value {
	once_RWSResult.Do(func() {
		cache_RWSResult = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_RWSResult
}

var cache_RWST gopurs_runtime.Value
var once_RWST sync.Once
func Get_RWST() gopurs_runtime.Value {
	once_RWST.Do(func() {
		cache_RWST = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_RWST(x_0_box)
})
	})
	return cache_RWST
}

var cache_withRWST gopurs_runtime.Value
var once_withRWST sync.Once
func Get_withRWST() gopurs_runtime.Value {
	once_withRWST.Do(func() {
		cache_withRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withRWST(f_0_box, m_1_box, r_2_box, s_3_box)
})
	})
	return cache_withRWST
}

var cache_runRWST gopurs_runtime.Value
var once_runRWST sync.Once
func Get_runRWST() gopurs_runtime.Value {
	once_runRWST.Do(func() {
		cache_runRWST = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runRWST(v_0_box)
})
	})
	return cache_runRWST
}

var cache_newtypeRWST gopurs_runtime.Value
var once_newtypeRWST sync.Once
func Get_newtypeRWST() gopurs_runtime.Value {
	once_newtypeRWST.Do(func() {
		cache_newtypeRWST = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeRWST
}

var cache_monadTransRWST gopurs_runtime.Value
var once_monadTransRWST sync.Once
func Get_monadTransRWST() gopurs_runtime.Value {
	once_monadTransRWST.Do(func() {
		cache_monadTransRWST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTransRWST(dictMonoid_0_box)
})
	})
	return cache_monadTransRWST
}

var cache_mapRWST gopurs_runtime.Value
var once_mapRWST sync.Once
func Get_mapRWST() gopurs_runtime.Value {
	once_mapRWST.Do(func() {
		cache_mapRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapRWST(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_mapRWST
}

var cache_lazyRWST gopurs_runtime.Value
var once_lazyRWST sync.Once
func Get_lazyRWST() gopurs_runtime.Value {
	once_lazyRWST.Do(func() {
		cache_lazyRWST = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), r_1, s_2)
})
})
}))
	})
	return cache_lazyRWST
}

var cache_functorRWST gopurs_runtime.Value
var once_functorRWST sync.Once
func Get_functorRWST() gopurs_runtime.Value {
	once_functorRWST.Do(func() {
		cache_functorRWST = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorRWST(dictFunctor_0_box)
})
	})
	return cache_functorRWST
}

var cache_execRWST gopurs_runtime.Value
var once_execRWST sync.Once
func Get_execRWST() gopurs_runtime.Value {
	once_execRWST.Do(func() {
		cache_execRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execRWST(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_execRWST
}

var cache_evalRWST gopurs_runtime.Value
var once_evalRWST sync.Once
func Get_evalRWST() gopurs_runtime.Value {
	once_evalRWST.Do(func() {
		cache_evalRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalRWST(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_evalRWST
}

var cache_applyRWST gopurs_runtime.Value
var once_applyRWST sync.Once
func Get_applyRWST() gopurs_runtime.Value {
	once_applyRWST.Do(func() {
		cache_applyRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyRWST(dictBind_0_box)
})
	})
	return cache_applyRWST
}

var cache_bindRWST gopurs_runtime.Value
var once_bindRWST sync.Once
func Get_bindRWST() gopurs_runtime.Value {
	once_bindRWST.Do(func() {
		cache_bindRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindRWST(dictBind_0_box)
})
	})
	return cache_bindRWST
}

var cache_semigroupRWST gopurs_runtime.Value
var once_semigroupRWST sync.Once
func Get_semigroupRWST() gopurs_runtime.Value {
	once_semigroupRWST.Do(func() {
		cache_semigroupRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupRWST(dictBind_0_box)
})
	})
	return cache_semigroupRWST
}

var cache_applicativeRWST gopurs_runtime.Value
var once_applicativeRWST sync.Once
func Get_applicativeRWST() gopurs_runtime.Value {
	once_applicativeRWST.Do(func() {
		cache_applicativeRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeRWST(dictMonad_0_box)
})
	})
	return cache_applicativeRWST
}

var cache_monadRWST gopurs_runtime.Value
var once_monadRWST sync.Once
func Get_monadRWST() gopurs_runtime.Value {
	once_monadRWST.Do(func() {
		cache_monadRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRWST(dictMonad_0_box)
})
	})
	return cache_monadRWST
}

var cache_monadAskRWST gopurs_runtime.Value
var once_monadAskRWST sync.Once
func Get_monadAskRWST() gopurs_runtime.Value {
	once_monadAskRWST.Do(func() {
		cache_monadAskRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskRWST(dictMonad_0_box)
})
	})
	return cache_monadAskRWST
}

var cache_monadReaderRWST gopurs_runtime.Value
var once_monadReaderRWST sync.Once
func Get_monadReaderRWST() gopurs_runtime.Value {
	once_monadReaderRWST.Do(func() {
		cache_monadReaderRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderRWST(dictMonad_0_box)
})
	})
	return cache_monadReaderRWST
}

var cache_monadEffectRWS gopurs_runtime.Value
var once_monadEffectRWS sync.Once
func Get_monadEffectRWS() gopurs_runtime.Value {
	once_monadEffectRWS.Do(func() {
		cache_monadEffectRWS = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictMonadEffect_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectRWS(dictMonoid_0_box, dictMonadEffect_1_box)
})
	})
	return cache_monadEffectRWS
}

var cache_monadRecRWST gopurs_runtime.Value
var once_monadRecRWST sync.Once
func Get_monadRecRWST() gopurs_runtime.Value {
	once_monadRecRWST.Do(func() {
		cache_monadRecRWST = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecRWST(dictMonadRec_0_box)
})
	})
	return cache_monadRecRWST
}

var cache_monadStateRWST gopurs_runtime.Value
var once_monadStateRWST sync.Once
func Get_monadStateRWST() gopurs_runtime.Value {
	once_monadStateRWST.Do(func() {
		cache_monadStateRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateRWST(dictMonad_0_box)
})
	})
	return cache_monadStateRWST
}

var cache_monadTellRWST gopurs_runtime.Value
var once_monadTellRWST sync.Once
func Get_monadTellRWST() gopurs_runtime.Value {
	once_monadTellRWST.Do(func() {
		cache_monadTellRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellRWST(dictMonad_0_box)
})
	})
	return cache_monadTellRWST
}

var cache_monadWriterRWST gopurs_runtime.Value
var once_monadWriterRWST sync.Once
func Get_monadWriterRWST() gopurs_runtime.Value {
	once_monadWriterRWST.Do(func() {
		cache_monadWriterRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterRWST(dictMonad_0_box)
})
	})
	return cache_monadWriterRWST
}

var cache_monadThrowRWST gopurs_runtime.Value
var once_monadThrowRWST sync.Once
func Get_monadThrowRWST() gopurs_runtime.Value {
	once_monadThrowRWST.Do(func() {
		cache_monadThrowRWST = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowRWST(dictMonadThrow_0_box)
})
	})
	return cache_monadThrowRWST
}

var cache_monadErrorRWST gopurs_runtime.Value
var once_monadErrorRWST sync.Once
func Get_monadErrorRWST() gopurs_runtime.Value {
	once_monadErrorRWST.Do(func() {
		cache_monadErrorRWST = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorRWST(dictMonadError_0_box)
})
	})
	return cache_monadErrorRWST
}

var cache_monadSTRWST gopurs_runtime.Value
var once_monadSTRWST sync.Once
func Get_monadSTRWST() gopurs_runtime.Value {
	once_monadSTRWST.Do(func() {
		cache_monadSTRWST = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictMonadST_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTRWST(dictMonoid_0_box, dictMonadST_1_box)
})
	})
	return cache_monadSTRWST
}

var cache_monoidRWST gopurs_runtime.Value
var once_monoidRWST sync.Once
func Get_monoidRWST() gopurs_runtime.Value {
	once_monoidRWST.Do(func() {
		cache_monoidRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidRWST(dictMonad_0_box)
})
	})
	return cache_monoidRWST
}

var cache_altRWST gopurs_runtime.Value
var once_altRWST sync.Once
func Get_altRWST() gopurs_runtime.Value {
	once_altRWST.Do(func() {
		cache_altRWST = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altRWST(dictAlt_0_box)
})
	})
	return cache_altRWST
}

var cache_plusRWST gopurs_runtime.Value
var once_plusRWST sync.Once
func Get_plusRWST() gopurs_runtime.Value {
	once_plusRWST.Do(func() {
		cache_plusRWST = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusRWST(dictPlus_0_box)
})
	})
	return cache_plusRWST
}

var cache_alternativeRWST gopurs_runtime.Value
var once_alternativeRWST sync.Once
func Get_alternativeRWST() gopurs_runtime.Value {
	once_alternativeRWST.Do(func() {
		cache_alternativeRWST = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeRWST(dictMonoid_0_box, dictAlternative_1_box)
})
	})
	return cache_alternativeRWST
}

var cache_alt__267341625 gopurs_runtime.Value
var once_alt__267341625 sync.Once
func Get_alt__267341625() gopurs_runtime.Value {
	once_alt__267341625.Do(func() {
		cache_alt__267341625 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__267341625(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__267341625
}

var cache_alt__3311196213 gopurs_runtime.Value
var once_alt__3311196213 sync.Once
func Get_alt__3311196213() gopurs_runtime.Value {
	once_alt__3311196213.Do(func() {
		cache_alt__3311196213 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__3311196213(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__3311196213
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

var cache_pure__1748760400 gopurs_runtime.Value
var once_pure__1748760400 sync.Once
func Get_pure__1748760400() gopurs_runtime.Value {
	once_pure__1748760400.Do(func() {
		cache_pure__1748760400 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1748760400(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1748760400
}

var cache_pure__1953455120 gopurs_runtime.Value
var once_pure__1953455120 sync.Once
func Get_pure__1953455120() gopurs_runtime.Value {
	once_pure__1953455120.Do(func() {
		cache_pure__1953455120 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1953455120(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1953455120
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__619377840 gopurs_runtime.Value
var once_lift2__619377840 sync.Once
func Get_lift2__619377840() gopurs_runtime.Value {
	once_lift2__619377840.Do(func() {
		cache_lift2__619377840 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__619377840(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__619377840
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

var cache_bind__1881387271 gopurs_runtime.Value
var once_bind__1881387271 sync.Once
func Get_bind__1881387271() gopurs_runtime.Value {
	once_bind__1881387271.Do(func() {
		cache_bind__1881387271 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1881387271(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1881387271
}

var cache_bind__1136545735 gopurs_runtime.Value
var once_bind__1136545735 sync.Once
func Get_bind__1136545735() gopurs_runtime.Value {
	once_bind__1136545735.Do(func() {
		cache_bind__1136545735 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1136545735(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1136545735
}

var cache_bind__1217453831 gopurs_runtime.Value
var once_bind__1217453831 sync.Once
func Get_bind__1217453831() gopurs_runtime.Value {
	once_bind__1217453831.Do(func() {
		cache_bind__1217453831 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1217453831(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1217453831
}

var cache_bind__3465679815 gopurs_runtime.Value
var once_bind__3465679815 sync.Once
func Get_bind__3465679815() gopurs_runtime.Value {
	once_bind__3465679815.Do(func() {
		cache_bind__3465679815 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3465679815(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3465679815
}

var cache_bind__1902451271 gopurs_runtime.Value
var once_bind__1902451271 sync.Once
func Get_bind__1902451271() gopurs_runtime.Value {
	once_bind__1902451271.Do(func() {
		cache_bind__1902451271 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1902451271(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1902451271
}

var cache_bind__277824519 gopurs_runtime.Value
var once_bind__277824519 sync.Once
func Get_bind__277824519() gopurs_runtime.Value {
	once_bind__277824519.Do(func() {
		cache_bind__277824519 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__277824519(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__277824519
}

var cache_bind__925691911 gopurs_runtime.Value
var once_bind__925691911 sync.Once
func Get_bind__925691911() gopurs_runtime.Value {
	once_bind__925691911.Do(func() {
		cache_bind__925691911 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__925691911(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__925691911
}

var cache_catchError__2657403463 gopurs_runtime.Value
var once_catchError__2657403463 sync.Once
func Get_catchError__2657403463() gopurs_runtime.Value {
	once_catchError__2657403463.Do(func() {
		cache_catchError__2657403463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError__2657403463(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError__2657403463
}

var cache_catchError__1102377099 gopurs_runtime.Value
var once_catchError__1102377099 sync.Once
func Get_catchError__1102377099() gopurs_runtime.Value {
	once_catchError__1102377099.Do(func() {
		cache_catchError__1102377099 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError__1102377099(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError__1102377099
}

var cache_throwError__237885032 gopurs_runtime.Value
var once_throwError__237885032 sync.Once
func Get_throwError__237885032() gopurs_runtime.Value {
	once_throwError__237885032.Do(func() {
		cache_throwError__237885032 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_throwError__237885032(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_throwError__237885032
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__1818438580 gopurs_runtime.Value
var once_tailRecM__1818438580 sync.Once
func Get_tailRecM__1818438580() gopurs_runtime.Value {
	once_tailRecM__1818438580.Do(func() {
		cache_tailRecM__1818438580 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__1818438580(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__1818438580
}

var cache_lift__3816229929 gopurs_runtime.Value
var once_lift__3816229929 sync.Once
func Get_lift__3816229929() gopurs_runtime.Value {
	once_lift__3816229929.Do(func() {
		cache_lift__3816229929 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__3816229929(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lift__3816229929
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

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__2300857972 gopurs_runtime.Value
var once_map__2300857972 sync.Once
func Get_map__2300857972() gopurs_runtime.Value {
	once_map__2300857972.Do(func() {
		cache_map__2300857972 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2300857972(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2300857972
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

var cache_mapFlipped__2213324916 gopurs_runtime.Value
var once_mapFlipped__2213324916 sync.Once
func Get_mapFlipped__2213324916() gopurs_runtime.Value {
	once_mapFlipped__2213324916.Do(func() {
		cache_mapFlipped__2213324916 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__2213324916(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__2213324916
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

var cache_uncurry__3533477633 gopurs_runtime.Value
var once_uncurry__3533477633 sync.Once
func Get_uncurry__3533477633() gopurs_runtime.Value {
	once_uncurry__3533477633.Do(func() {
		cache_uncurry__3533477633 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncurry__3533477633(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_uncurry__3533477633
}

var cache_uncurry__1489344417 gopurs_runtime.Value
var once_uncurry__1489344417 sync.Once
func Get_uncurry__1489344417() gopurs_runtime.Value {
	once_uncurry__1489344417.Do(func() {
		cache_uncurry__1489344417 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncurry__1489344417(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_uncurry__1489344417
}

type Constructor_RWSResult[T_state any, T_result any, T_writer any] struct {
	Rc uint32
	V0 T_state
	V1 T_result
	V2 T_writer
}


func Call_RWST(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withRWST(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
__local_var_4_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, r_2, s_3))
_ = __local_var_4_0
return gopurs_runtime.Apply2(m_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__local_var_4_0)}.UnsafePtr).V1)
}

func Call_runRWST(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_monadTransRWST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_0.V1, m_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
}))
}

func Call_mapRWST(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3))
}

func Call_functorRWST(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1), (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_2, r_3, s_4))
})
})
})
}))
}

func Call_execRWST(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_evalRWST(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_applyRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
functorRWST1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1), (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
}))
_ = functorRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_5_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_3_2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_5 := (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V2
_ = __local_var_11_5
return gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1), gopurs_runtime.Apply2(Semigroup0_5_4.V0, __local_var_11_5, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_7, r_8, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0))
}))
})
})
})
}))
})
}

func Call_bindRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
applyRWST1_2_1 := Call_applyRWST(dictBind_0)
_ = applyRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
applyRWST2_5_3 := gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3)
_ = applyRWST2_5_3
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_5_3
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_4 := (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2
_ = __local_var_11_4
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1, gopurs_runtime.Apply2(Semigroup0_4_2.V0, __local_var_11_4, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_7, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1, r_8, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})
})
})
}))
})
}

func Call_semigroupRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
applyRWST1_1_0 := Call_applyRWST(dictBind_0)
_ = applyRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyRWST2_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2))
_ = applyRWST2_3_1
return gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST2_3_1.V0, gopurs_runtime.Value{}))
_ = Functor0_5_2
__local_var_6_3 := gopurs_runtime.RecordGet(dictSemigroup_4, "append")
_ = __local_var_6_3
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(applyRWST2_3_1.V1, gopurs_runtime.Apply2(Functor0_5_2.V0, __local_var_6_3, a_7), b_8)
})
}))
})
})
}

func Call_applicativeRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
applyRWST1_2_1 := Call_applyRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = applyRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyRWST2_4_2 := gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3)
_ = applyRWST2_4_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_4_2
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
})
})
}))
})
}

func Call_monadRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeRWST1_1_0 := Call_applicativeRWST(dictMonad_0)
_ = applicativeRWST1_1_0
bindRWST1_2_1 := Call_bindRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST2_4_2 := gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3)
_ = applicativeRWST2_4_2
bindRWST2_5_3 := gopurs_runtime.Apply(bindRWST1_2_1, dictMonoid_3)
_ = bindRWST2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindRWST2_5_3
}))
})
}

func Call_monadAskRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
monadRWST1_2_1 := Call_monadRWST(dictMonad_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}), gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, r_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
})
}))
})
}

func Call_monadReaderRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadAskRWST1_1_0 := Call_monadAskRWST(dictMonad_0)
_ = monadAskRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskRWST2_3_1 := gopurs_runtime.Apply(monadAskRWST1_1_0, dictMonoid_2)
_ = monadAskRWST2_3_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskRWST2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_5, gopurs_runtime.Apply(f_4, r_6), s_7)
})
})
})
}))
})
}

func Call_monadEffectRWS(dictMonoid_0_loop gopurs_runtime.Value, dictMonadEffect_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictMonadEffect_1 gopurs_runtime.Value = dictMonadEffect_1_loop
_ = dictMonadEffect_1
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
monadRWST1_3_1 := gopurs_runtime.Apply(Call_monadRWST(Monad0_2_0), dictMonoid_0)
_ = monadRWST1_3_1
Bind1_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
__local_var_4_2 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_4_3.V1, m_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
_ = __local_var_4_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_3_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "liftEffect"), x_5))
}))
}

func Call_monadRecRWST(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
monadRWST1_4_3 := Call_monadRWST(Monad0_1_0)
_ = monadRWST1_4_3
return gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_6_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_4
monadRWST2_7_5 := gopurs_runtime.Apply(monadRWST1_4_3, dictMonoid_5)
_ = monadRWST2_7_5
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_7_5
}), gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_13_6 := (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V2
_ = __local_var_13_6
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply3(k_8, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1, r_10, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 525585346) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(Semigroup0_6_4.V0, __local_var_13_6, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2)})}})}
goto end_branch_9
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 60402430) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0, (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(Semigroup0_6_4.V0, __local_var_13_6, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2)})}})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply(Applicative0_3_2.V1, __t9)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_11, a_9, gopurs_runtime.RecordGet(dictMonoid_5, "mempty")})})
})
})
})
}))
})
}

func Call_monadStateRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
monadRWST1_2_1 := Call_monadRWST(dictMonad_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
v1_8_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, s_7))
_ = v1_8_3
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_8_3)}.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_8_3)}.UnsafePtr).V0, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
})
})
}))
})
}

func Call_monadTellRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
monadRWST1_2_1 := Call_monadRWST(dictMonad_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_4_2
monadRWST2_5_3 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_5_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_4_2
}), gopurs_runtime.Func(func(w_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, pkg_Data_Unit.Get_unit(), w_6})})
})
})
}))
})
}

func Call_monadWriterRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_2_1
monadTellRWST1_3_2 := Call_monadTellRWST(dictMonad_0)
_ = monadTellRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
monadTellRWST2_5_3 := gopurs_runtime.Apply(monadTellRWST1_3_2, dictMonoid_4)
_ = monadTellRWST2_5_3
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellRWST2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_4
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V2})}, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V2})})
}))
})
})
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1.UnsafePtr).V1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V2)})})
}))
})
})
}))
})
}

func Call_monadThrowRWST(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
monadRWST1_2_1 := Call_monadRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadTransRWST1_4_2 := &pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_5_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
pure_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_4
return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_3.V1, m_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_9, a_10, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
}))
})
})
})
})}
_ = monadTransRWST1_4_2
monadRWST2_5_5 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_5_5
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_5_5
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(monadTransRWST1_4_2.V0, gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_6))
}))
})
}

func Call_monadErrorRWST(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowRWST1_1_0 := Call_monadThrowRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowRWST2_3_1 := gopurs_runtime.Apply(monadThrowRWST1_1_0, dictMonoid_2)
_ = monadThrowRWST2_3_1
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowRWST2_3_1
}), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply2(m_4, r_6, s_7), gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(h_5, e_8, r_6, s_7)
}))
})
})
})
}))
})
}

func Call_monadSTRWST(dictMonoid_0_loop gopurs_runtime.Value, dictMonadST_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictMonadST_1 gopurs_runtime.Value = dictMonadST_1_loop
_ = dictMonadST_1
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
monadRWST1_3_1 := gopurs_runtime.Apply(Call_monadRWST(Monad0_2_0), dictMonoid_0)
_ = monadRWST1_3_1
Bind1_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
__local_var_4_2 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_4_3.V1, m_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
_ = __local_var_4_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_3_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "liftST"), x_5))
}))
}

func Call_monoidRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeRWST1_1_0 := Call_applicativeRWST(dictMonad_0)
_ = applicativeRWST1_1_0
semigroupRWST1_2_1 := Call_semigroupRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = semigroupRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST2_4_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3))
_ = applicativeRWST2_4_2
semigroupRWST2_5_3 := gopurs_runtime.Apply(semigroupRWST1_2_1, dictMonoid_3)
_ = semigroupRWST2_5_3
return gopurs_runtime.Func(func(dictMonoid1_6 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRWST3_7_4 := gopurs_runtime.Apply(semigroupRWST2_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupRWST3_7_4
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRWST3_7_4
}), gopurs_runtime.Apply(applicativeRWST2_4_2.V1, gopurs_runtime.RecordGet(dictMonoid1_6, "mempty")))
})
})
}

func Call_altRWST(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorRWST1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1), (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_3, r_4, s_5))
})
})
})
}))
_ = functorRWST1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(v_2, r_4, s_5), gopurs_runtime.Apply2(v1_3, r_4, s_5))
})
})
})
}))
}

func Call_plusRWST(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
functorRWST1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1), (*Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
}))
_ = functorRWST1_3_3
altRWST1_2_1 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_3_3
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), gopurs_runtime.Apply2(v_4, r_6, s_7), gopurs_runtime.Apply2(v1_5, r_6, s_7))
})
})
})
}))
_ = altRWST1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altRWST1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
})
}))
}

func Call_alternativeRWST(dictMonoid_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
plusRWST1_2_0 := Call_plusRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{}))
_ = plusRWST1_2_0
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST1_4_1 := gopurs_runtime.Apply(Call_applicativeRWST(dictMonad_3), dictMonoid_0)
_ = applicativeRWST1_4_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST1_4_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusRWST1_2_0
}))
})
}

func Call_alt__267341625(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__3311196213(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1748760400(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1953455120(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__619377840(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1881387271(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1136545735(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1217453831(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3465679815(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1902451271(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__277824519(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__925691911(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchError__2657403463(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchError__1102377099(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_throwError__237885032(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__1818438580(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift__3816229929(dict_0_loop *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2300857972(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_mapFlipped__2213324916(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
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

func Call_uncurry__3533477633(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_uncurry__1489344417(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}


