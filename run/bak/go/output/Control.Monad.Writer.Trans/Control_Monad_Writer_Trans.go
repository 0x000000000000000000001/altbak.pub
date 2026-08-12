package Control_Monad_Writer_Trans

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_Cont_Class "gopurs/output/Control.Monad.Cont.Class"
	pkg_Control_Monad_Error_Class "gopurs/output/Control.Monad.Error.Class"
	pkg_Control_Monad_Reader_Class "gopurs/output/Control.Monad.Reader.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Monad_State_Class "gopurs/output/Control.Monad.State.Class"
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

var cache_WriterT gopurs_runtime.Value
var once_WriterT sync.Once
func Get_WriterT() gopurs_runtime.Value {
	once_WriterT.Do(func() {
		cache_WriterT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_WriterT(x_0_box)
})
	})
	return cache_WriterT
}

var cache_runWriterT gopurs_runtime.Value
var once_runWriterT sync.Once
func Get_runWriterT() gopurs_runtime.Value {
	once_runWriterT.Do(func() {
		cache_runWriterT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runWriterT(v_0_box)
})
	})
	return cache_runWriterT
}

var cache_newtypeWriterT gopurs_runtime.Value
var once_newtypeWriterT sync.Once
func Get_newtypeWriterT() gopurs_runtime.Value {
	once_newtypeWriterT.Do(func() {
		cache_newtypeWriterT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeWriterT
}

var cache_monadTransWriterT gopurs_runtime.Value
var once_monadTransWriterT sync.Once
func Get_monadTransWriterT() gopurs_runtime.Value {
	once_monadTransWriterT.Do(func() {
		cache_monadTransWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTransWriterT(dictMonoid_0_box)
})
	})
	return cache_monadTransWriterT
}

var cache_mapWriterT gopurs_runtime.Value
var once_mapWriterT sync.Once
func Get_mapWriterT() gopurs_runtime.Value {
	once_mapWriterT.Do(func() {
		cache_mapWriterT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWriterT(f_0_box, v_1_box)
})
	})
	return cache_mapWriterT
}

var cache_functorWriterT gopurs_runtime.Value
var once_functorWriterT sync.Once
func Get_functorWriterT() gopurs_runtime.Value {
	once_functorWriterT.Do(func() {
		cache_functorWriterT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWriterT(dictFunctor_0_box)
})
	})
	return cache_functorWriterT
}

var cache_execWriterT gopurs_runtime.Value
var once_execWriterT sync.Once
func Get_execWriterT() gopurs_runtime.Value {
	once_execWriterT.Do(func() {
		cache_execWriterT = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execWriterT(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
})
	})
	return cache_execWriterT
}

var cache_applyWriterT gopurs_runtime.Value
var once_applyWriterT sync.Once
func Get_applyWriterT() gopurs_runtime.Value {
	once_applyWriterT.Do(func() {
		cache_applyWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyWriterT(dictSemigroup_0_box, dictApply_1_box)
})
	})
	return cache_applyWriterT
}

var cache_bindWriterT gopurs_runtime.Value
var once_bindWriterT sync.Once
func Get_bindWriterT() gopurs_runtime.Value {
	once_bindWriterT.Do(func() {
		cache_bindWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictBind_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindWriterT(dictSemigroup_0_box, dictBind_1_box)
})
	})
	return cache_bindWriterT
}

var cache_semigroupWriterT gopurs_runtime.Value
var once_semigroupWriterT sync.Once
func Get_semigroupWriterT() gopurs_runtime.Value {
	once_semigroupWriterT.Do(func() {
		cache_semigroupWriterT = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupWriterT(dictApply_0_box, dictSemigroup_1_box)
})
	})
	return cache_semigroupWriterT
}

var cache_applicativeWriterT gopurs_runtime.Value
var once_applicativeWriterT sync.Once
func Get_applicativeWriterT() gopurs_runtime.Value {
	once_applicativeWriterT.Do(func() {
		cache_applicativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeWriterT(dictMonoid_0_box)
})
	})
	return cache_applicativeWriterT
}

var cache_monadWriterT gopurs_runtime.Value
var once_monadWriterT sync.Once
func Get_monadWriterT() gopurs_runtime.Value {
	once_monadWriterT.Do(func() {
		cache_monadWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterT(dictMonoid_0_box)
})
	})
	return cache_monadWriterT
}

var cache_monadAskWriterT gopurs_runtime.Value
var once_monadAskWriterT sync.Once
func Get_monadAskWriterT() gopurs_runtime.Value {
	once_monadAskWriterT.Do(func() {
		cache_monadAskWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskWriterT(dictMonoid_0_box)
})
	})
	return cache_monadAskWriterT
}

var cache_monadReaderWriterT gopurs_runtime.Value
var once_monadReaderWriterT sync.Once
func Get_monadReaderWriterT() gopurs_runtime.Value {
	once_monadReaderWriterT.Do(func() {
		cache_monadReaderWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderWriterT(dictMonoid_0_box)
})
	})
	return cache_monadReaderWriterT
}

var cache_monadContWriterT gopurs_runtime.Value
var once_monadContWriterT sync.Once
func Get_monadContWriterT() gopurs_runtime.Value {
	once_monadContWriterT.Do(func() {
		cache_monadContWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContWriterT(dictMonoid_0_box)
})
	})
	return cache_monadContWriterT
}

var cache_monadEffectWriter gopurs_runtime.Value
var once_monadEffectWriter sync.Once
func Get_monadEffectWriter() gopurs_runtime.Value {
	once_monadEffectWriter.Do(func() {
		cache_monadEffectWriter = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectWriter(dictMonoid_0_box)
})
	})
	return cache_monadEffectWriter
}

var cache_monadRecWriterT gopurs_runtime.Value
var once_monadRecWriterT sync.Once
func Get_monadRecWriterT() gopurs_runtime.Value {
	once_monadRecWriterT.Do(func() {
		cache_monadRecWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecWriterT(dictMonoid_0_box)
})
	})
	return cache_monadRecWriterT
}

var cache_monadStateWriterT gopurs_runtime.Value
var once_monadStateWriterT sync.Once
func Get_monadStateWriterT() gopurs_runtime.Value {
	once_monadStateWriterT.Do(func() {
		cache_monadStateWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateWriterT(dictMonoid_0_box)
})
	})
	return cache_monadStateWriterT
}

var cache_monadTellWriterT gopurs_runtime.Value
var once_monadTellWriterT sync.Once
func Get_monadTellWriterT() gopurs_runtime.Value {
	once_monadTellWriterT.Do(func() {
		cache_monadTellWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellWriterT(dictMonoid_0_box)
})
	})
	return cache_monadTellWriterT
}

var cache_monadWriterWriterT gopurs_runtime.Value
var once_monadWriterWriterT sync.Once
func Get_monadWriterWriterT() gopurs_runtime.Value {
	once_monadWriterWriterT.Do(func() {
		cache_monadWriterWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterWriterT(dictMonoid_0_box)
})
	})
	return cache_monadWriterWriterT
}

var cache_monadThrowWriterT gopurs_runtime.Value
var once_monadThrowWriterT sync.Once
func Get_monadThrowWriterT() gopurs_runtime.Value {
	once_monadThrowWriterT.Do(func() {
		cache_monadThrowWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowWriterT(dictMonoid_0_box)
})
	})
	return cache_monadThrowWriterT
}

var cache_monadErrorWriterT gopurs_runtime.Value
var once_monadErrorWriterT sync.Once
func Get_monadErrorWriterT() gopurs_runtime.Value {
	once_monadErrorWriterT.Do(func() {
		cache_monadErrorWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorWriterT(dictMonoid_0_box)
})
	})
	return cache_monadErrorWriterT
}

var cache_monadSTWriterT gopurs_runtime.Value
var once_monadSTWriterT sync.Once
func Get_monadSTWriterT() gopurs_runtime.Value {
	once_monadSTWriterT.Do(func() {
		cache_monadSTWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTWriterT(dictMonoid_0_box)
})
	})
	return cache_monadSTWriterT
}

var cache_monoidWriterT gopurs_runtime.Value
var once_monoidWriterT sync.Once
func Get_monoidWriterT() gopurs_runtime.Value {
	once_monoidWriterT.Do(func() {
		cache_monoidWriterT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidWriterT(dictApplicative_0_box)
})
	})
	return cache_monoidWriterT
}

var cache_altWriterT gopurs_runtime.Value
var once_altWriterT sync.Once
func Get_altWriterT() gopurs_runtime.Value {
	once_altWriterT.Do(func() {
		cache_altWriterT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altWriterT(dictAlt_0_box)
})
	})
	return cache_altWriterT
}

var cache_plusWriterT gopurs_runtime.Value
var once_plusWriterT sync.Once
func Get_plusWriterT() gopurs_runtime.Value {
	once_plusWriterT.Do(func() {
		cache_plusWriterT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusWriterT(dictPlus_0_box)
})
	})
	return cache_plusWriterT
}

var cache_alternativeWriterT gopurs_runtime.Value
var once_alternativeWriterT sync.Once
func Get_alternativeWriterT() gopurs_runtime.Value {
	once_alternativeWriterT.Do(func() {
		cache_alternativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeWriterT(dictMonoid_0_box)
})
	})
	return cache_alternativeWriterT
}

var cache_monadPlusWriterT gopurs_runtime.Value
var once_monadPlusWriterT sync.Once
func Get_monadPlusWriterT() gopurs_runtime.Value {
	once_monadPlusWriterT.Do(func() {
		cache_monadPlusWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusWriterT(dictMonoid_0_box)
})
	})
	return cache_monadPlusWriterT
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

var cache_alt__3232097233 gopurs_runtime.Value
var once_alt__3232097233 sync.Once
func Get_alt__3232097233() gopurs_runtime.Value {
	once_alt__3232097233.Do(func() {
		cache_alt__3232097233 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__3232097233(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__3232097233
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

var cache_pure__3012389648 gopurs_runtime.Value
var once_pure__3012389648 sync.Once
func Get_pure__3012389648() gopurs_runtime.Value {
	once_pure__3012389648.Do(func() {
		cache_pure__3012389648 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3012389648(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3012389648
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

var cache_apply__2087590060 gopurs_runtime.Value
var once_apply__2087590060 sync.Once
func Get_apply__2087590060() gopurs_runtime.Value {
	once_apply__2087590060.Do(func() {
		cache_apply__2087590060 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2087590060(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2087590060
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

var cache_lift2__1517638032 gopurs_runtime.Value
var once_lift2__1517638032 sync.Once
func Get_lift2__1517638032() gopurs_runtime.Value {
	once_lift2__1517638032.Do(func() {
		cache_lift2__1517638032 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__1517638032(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__1517638032
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

var cache_bind__889812231 gopurs_runtime.Value
var once_bind__889812231 sync.Once
func Get_bind__889812231() gopurs_runtime.Value {
	once_bind__889812231.Do(func() {
		cache_bind__889812231 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__889812231(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__889812231
}

var cache_bind__2669704711 gopurs_runtime.Value
var once_bind__2669704711 sync.Once
func Get_bind__2669704711() gopurs_runtime.Value {
	once_bind__2669704711.Do(func() {
		cache_bind__2669704711 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2669704711(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2669704711
}

var cache_bind__2140072327 gopurs_runtime.Value
var once_bind__2140072327 sync.Once
func Get_bind__2140072327() gopurs_runtime.Value {
	once_bind__2140072327.Do(func() {
		cache_bind__2140072327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2140072327(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2140072327
}

var cache_bind__882000455 gopurs_runtime.Value
var once_bind__882000455 sync.Once
func Get_bind__882000455() gopurs_runtime.Value {
	once_bind__882000455.Do(func() {
		cache_bind__882000455 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__882000455(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__882000455
}

var cache_bind__1756310855 gopurs_runtime.Value
var once_bind__1756310855 sync.Once
func Get_bind__1756310855() gopurs_runtime.Value {
	once_bind__1756310855.Do(func() {
		cache_bind__1756310855 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1756310855(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1756310855
}

var cache_callCC__1888484333 gopurs_runtime.Value
var once_callCC__1888484333 sync.Once
func Get_callCC__1888484333() gopurs_runtime.Value {
	once_callCC__1888484333.Do(func() {
		cache_callCC__1888484333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_callCC__1888484333(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Cont_Class.Constructor_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_callCC__1888484333
}

var cache_callCC__2318135621 gopurs_runtime.Value
var once_callCC__2318135621 sync.Once
func Get_callCC__2318135621() gopurs_runtime.Value {
	once_callCC__2318135621.Do(func() {
		cache_callCC__2318135621 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_callCC__2318135621(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Cont_Class.Constructor_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_callCC__2318135621
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

var cache_catchError__3649261295 gopurs_runtime.Value
var once_catchError__3649261295 sync.Once
func Get_catchError__3649261295() gopurs_runtime.Value {
	once_catchError__3649261295.Do(func() {
		cache_catchError__3649261295 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catchError__3649261295(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_catchError__3649261295
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

var cache_local__1299460031 gopurs_runtime.Value
var once_local__1299460031 sync.Once
func Get_local__1299460031() gopurs_runtime.Value {
	once_local__1299460031.Do(func() {
		cache_local__1299460031 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local__1299460031(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_local__1299460031
}

var cache_local__909940799 gopurs_runtime.Value
var once_local__909940799 sync.Once
func Get_local__909940799() gopurs_runtime.Value {
	once_local__909940799.Do(func() {
		cache_local__909940799 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local__909940799(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_local__909940799
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

var cache_tailRecM__3478800400 gopurs_runtime.Value
var once_tailRecM__3478800400 sync.Once
func Get_tailRecM__3478800400() gopurs_runtime.Value {
	once_tailRecM__3478800400.Do(func() {
		cache_tailRecM__3478800400 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3478800400(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3478800400
}

var cache_state__3572857840 gopurs_runtime.Value
var once_state__3572857840 sync.Once
func Get_state__3572857840() gopurs_runtime.Value {
	once_state__3572857840.Do(func() {
		cache_state__3572857840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_state__3572857840(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_State_Class.Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_state__3572857840
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

var cache_lift__115114023 gopurs_runtime.Value
var once_lift__115114023 sync.Once
func Get_lift__115114023() gopurs_runtime.Value {
	once_lift__115114023.Do(func() {
		cache_lift__115114023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__115114023(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lift__115114023
}

var cache_mapWriterT__2842489082 gopurs_runtime.Value
var once_mapWriterT__2842489082 sync.Once
func Get_mapWriterT__2842489082() gopurs_runtime.Value {
	once_mapWriterT__2842489082.Do(func() {
		cache_mapWriterT__2842489082 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWriterT__2842489082(f_0_box, v_1_box)
})
	})
	return cache_mapWriterT__2842489082
}

var cache_mapWriterT__77717660 gopurs_runtime.Value
var once_mapWriterT__77717660 sync.Once
func Get_mapWriterT__77717660() gopurs_runtime.Value {
	once_mapWriterT__77717660.Do(func() {
		cache_mapWriterT__77717660 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWriterT__77717660(f_0_box, v_1_box)
})
	})
	return cache_mapWriterT__77717660
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

var cache_map__2345808404 gopurs_runtime.Value
var once_map__2345808404 sync.Once
func Get_map__2345808404() gopurs_runtime.Value {
	once_map__2345808404.Do(func() {
		cache_map__2345808404 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2345808404(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2345808404
}

var cache_map__1352087572 gopurs_runtime.Value
var once_map__1352087572 sync.Once
func Get_map__1352087572() gopurs_runtime.Value {
	once_map__1352087572.Do(func() {
		cache_map__1352087572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1352087572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1352087572
}

var cache_map__381330420 gopurs_runtime.Value
var once_map__381330420 sync.Once
func Get_map__381330420() gopurs_runtime.Value {
	once_map__381330420.Do(func() {
		cache_map__381330420 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__381330420(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__381330420
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

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

func Call_WriterT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runWriterT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_monadTransWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_0.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
}))
}

func Call_mapWriterT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorWriterT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1})}
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
}))
}

func Call_execWriterT(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(dictFunctor_0.V0, pkg_Data_Tuple.Get_snd(), v_1)
}

func Call_applyWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
functorWriterT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1})}
}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, v_6)
})
}))
_ = functorWriterT1_3_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_1, "apply"), gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_6.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_7.UnsafePtr).V1)})}
})
}), v_4), v1_5)
})
}))
}

func Call_bindWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictBind_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictBind_1 gopurs_runtime.Value = dictBind_1_loop
_ = dictBind_1
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_0
Functor0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
applyWriterT2_4_2 := Call_applyWriterT(dictSemigroup_0, Apply0_2_0)
_ = applyWriterT2_4_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_1, "bind"), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1
_ = __local_var_8_3
return gopurs_runtime.Apply2(Functor0_3_1.V0, gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_9.UnsafePtr).V1)})}
}), gopurs_runtime.Apply(k_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0))
}))
})
}))
}

func Call_semigroupWriterT(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
applyWriterT1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](Call_applyWriterT(dictSemigroup_1, dictApply_0))
_ = applyWriterT1_2_0
return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT1_2_0.V0, gopurs_runtime.Value{}))
_ = Functor0_4_1
__local_var_5_2 := gopurs_runtime.RecordGet(dictSemigroup1_3, "append")
_ = __local_var_5_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(applyWriterT1_2_0.V1, gopurs_runtime.Apply2(Functor0_4_1.V0, __local_var_5_2, a_6), b_7)
})
}))
})
}

func Call_applicativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applyWriterT1_1_0 := gopurs_runtime.Apply(Get_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyWriterT1_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyWriterT2_3_1 := gopurs_runtime.Apply(applyWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT2_3_1
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_3_1
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
}

func Call_monadWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applicativeWriterT1_1_0 := Call_applicativeWriterT(dictMonoid_0)
_ = applicativeWriterT1_1_0
bindWriterT1_2_1 := gopurs_runtime.Apply(Get_bindWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindWriterT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT2_4_2 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeWriterT2_4_2
bindWriterT2_5_3 := gopurs_runtime.Apply(bindWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = bindWriterT2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_5_3
}))
})
}

func Call_monadAskWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadTransWriterT1_1_0 := &pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
monadWriterT1_2_3 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_3
return gopurs_runtime.Func(func(dictMonadAsk_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_4 := gopurs_runtime.Apply(monadWriterT1_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_4_4
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_4
}), gopurs_runtime.Apply2(monadTransWriterT1_1_0.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{}), gopurs_runtime.RecordGet(dictMonadAsk_3, "ask")))
})
}

func Call_monadReaderWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadAskWriterT1_1_0 := Call_monadAskWriterT(dictMonoid_0)
_ = monadAskWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadReader_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskWriterT2_3_1 := gopurs_runtime.Apply(monadAskWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskWriterT2_3_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskWriterT2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "local"), f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, v_6)
})
}))
})
}

func Call_monadContWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadWriterT1_1_0 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadCont_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_3_1 := gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_2, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_3_1
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_2, "callCC"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
}))
}))
})
}

func Call_monadEffectWriter(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadWriterT1_1_0 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadEffect_2 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_1
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, Monad0_3_1)
_ = monadWriterT2_4_2
Bind1_5_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
pure_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_4
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_3.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "liftEffect"), x_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
}))
})
}

func Call_monadRecWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
monadWriterT1_2_1 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadRec_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
Bind1_5_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
monadWriterT2_7_5 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_7_5
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_7_5
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_3, "tailRecM"), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1
_ = __local_var_11_6
return gopurs_runtime.Apply2(Bind1_5_3.V1, gopurs_runtime.Apply(f_8, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 525585346) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(Semigroup0_1_0.V0, __local_var_11_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1)})}})}
goto end_branch_9
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 60402430) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(Semigroup0_1_0.V0, __local_var_11_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1)})}})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply(Applicative0_6_4.V1, __t9)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
})
}))
})
}

func Call_monadStateWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadTransWriterT1_1_0 := &pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
monadWriterT1_2_3 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_3
return gopurs_runtime.Func(func(dictMonadState_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_4
monadWriterT2_5_5 := gopurs_runtime.Apply(monadWriterT1_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_5_5
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_5
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(monadTransWriterT1_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_4)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "state"), f_6))
}))
})
}

func Call_monadTellWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
Semigroup0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_1_0
monadWriterT1_2_1 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_2_1, dictMonad_3)
_ = monadWriterT2_4_2
__local_var_5_3 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), pkg_Data_Unit.Get_unit())
_ = __local_var_5_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(__local_var_5_3, x_6))
}))
})
}

func Call_monadWriterWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadTellWriterT1_1_0 := Call_monadTellWriterT(dictMonoid_0)
_ = monadTellWriterT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
Applicative0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_4_2
monadTellWriterT2_5_3 := gopurs_runtime.Apply(monadTellWriterT1_1_0, dictMonad_2)
_ = monadTellWriterT2_5_3
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellWriterT2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_1.V1, v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1})})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_1.V1, v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_7.UnsafePtr).V1)})})
}))
}))
})
}

func Call_monadThrowWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadTransWriterT1_1_0 := &pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})}
_ = monadTransWriterT1_1_0
monadWriterT1_2_3 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_2_3
return gopurs_runtime.Func(func(dictMonadThrow_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_4_4
monadWriterT2_5_5 := gopurs_runtime.Apply(monadWriterT1_2_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_5_5
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_5
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(monadTransWriterT1_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_4_4)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "throwError"), e_6))
}))
})
}

func Call_monadErrorWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadThrowWriterT1_1_0 := Call_monadThrowWriterT(dictMonoid_0)
_ = monadThrowWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadError_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowWriterT2_3_1 := gopurs_runtime.Apply(monadThrowWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_2, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowWriterT2_3_1
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowWriterT2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_2, "catchError"), v_4, gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_5, e_6)
}))
})
}))
})
}

func Call_monadSTWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadWriterT1_1_0 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadST_2 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_1
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, Monad0_3_1)
_ = monadWriterT2_4_2
Bind1_5_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
pure_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_4
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_3.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "liftST"), x_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
}))
})
}

func Call_monoidWriterT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
semigroupWriterT1_1_0 := gopurs_runtime.Apply(Get_semigroupWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = semigroupWriterT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_applicativeWriterT(dictMonoid_2), dictApplicative_0))
_ = applicativeWriterT1_3_1
semigroupWriterT2_4_2 := gopurs_runtime.Apply(semigroupWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupWriterT2_4_2
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupWriterT3_6_3 := gopurs_runtime.Apply(semigroupWriterT2_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupWriterT3_6_3
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupWriterT3_6_3
}), gopurs_runtime.Apply(applicativeWriterT1_3_1.V1, gopurs_runtime.RecordGet(dictMonoid1_5, "mempty")))
})
})
}

func Call_altWriterT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorWriterT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, v_4)
})
}))
_ = functorWriterT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_2, v1_3)
})
}))
}

func Call_plusWriterT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorWriterT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1})}
}))
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, v_5)
})
}))
_ = functorWriterT1_2_2
altWriterT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "alt"), v_3, v1_4)
})
}))
_ = altWriterT1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_1_0
}), gopurs_runtime.RecordGet(dictPlus_0, "empty"))
}

func Call_alternativeWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applicativeWriterT1_1_0 := Call_applicativeWriterT(dictMonoid_0)
_ = applicativeWriterT1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT2_3_1 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeWriterT2_3_1
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_3
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_5
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
functorWriterT1_6_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1})}
}))
_ = __local_var_8_8
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_8, v_9)
})
}))
_ = functorWriterT1_6_6
altWriterT1_5_4 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_6
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "alt"), v_7, v1_8)
})
}))
_ = altWriterT1_5_4
plusWriterT1_4_2 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_5_4
}), gopurs_runtime.RecordGet(__local_var_4_3, "empty"))
_ = plusWriterT1_4_2
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusWriterT1_4_2
}))
})
}

func Call_monadPlusWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadWriterT1_1_0 := Call_monadWriterT(dictMonoid_0)
_ = monadWriterT1_1_0
alternativeWriterT1_2_1 := Call_alternativeWriterT(dictMonoid_0)
_ = alternativeWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadPlus_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_4_2
alternativeWriterT2_5_3 := gopurs_runtime.Apply(alternativeWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeWriterT2_5_3
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeWriterT2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}))
})
}

func Call_alt__267341625(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__3232097233(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3012389648(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_apply__2087590060(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_lift2__1517638032(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_bind__889812231(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2669704711(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2140072327(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__882000455(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1756310855(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_callCC__1888484333(dict_0_loop *pkg_Control_Monad_Cont_Class.Constructor_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Cont_Class.Constructor_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_callCC__2318135621(dict_0_loop *pkg_Control_Monad_Cont_Class.Constructor_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Cont_Class.Constructor_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchError__2657403463(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_catchError__3649261295(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadError[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_throwError__237885032(dict_0_loop *pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Error_Class.Constructor_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_local__1299460031(dict_0_loop *pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_local__909940799(dict_0_loop *pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__3478800400(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_state__3572857840(dict_0_loop *pkg_Control_Monad_State_Class.Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_State_Class.Constructor_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift__3816229929(dict_0_loop *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_lift__115114023(dict_0_loop *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapWriterT__2842489082(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_mapWriterT__77717660(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
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

func Call_map__2345808404(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1352087572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__381330420(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}


