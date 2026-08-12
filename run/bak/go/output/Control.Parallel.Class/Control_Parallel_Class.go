package Control_Parallel_Class

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Control_Monad_Writer_Trans "gopurs/output/Control.Monad.Writer.Trans"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_Costar "gopurs/output/Data.Functor.Costar"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect_Class "gopurs/output/Effect.Class"
	pkg_Effect_Ref "gopurs/output/Effect.Ref"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_ParCont gopurs_runtime.Value
var once_ParCont sync.Once
func Get_ParCont() gopurs_runtime.Value {
	once_ParCont.Do(func() {
		cache_ParCont = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ParCont(x_0_box)
})
	})
	return cache_ParCont
}

var cache_sequential gopurs_runtime.Value
var once_sequential sync.Once
func Get_sequential() gopurs_runtime.Value {
	once_sequential.Do(func() {
		cache_sequential = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequential
}

var cache_parallel gopurs_runtime.Value
var once_parallel sync.Once
func Get_parallel() gopurs_runtime.Value {
	once_parallel.Do(func() {
		cache_parallel = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_parallel
}

var cache_newtypeParCont gopurs_runtime.Value
var once_newtypeParCont sync.Once
func Get_newtypeParCont() gopurs_runtime.Value {
	once_newtypeParCont.Do(func() {
		cache_newtypeParCont = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeParCont
}

var cache_monadParWriterT gopurs_runtime.Value
var once_monadParWriterT sync.Once
func Get_monadParWriterT() gopurs_runtime.Value {
	once_monadParWriterT.Do(func() {
		cache_monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParWriterT(dictMonoid_0_box)
})
	})
	return cache_monadParWriterT
}

var cache_monadParStar gopurs_runtime.Value
var once_monadParStar sync.Once
func Get_monadParStar() gopurs_runtime.Value {
	once_monadParStar.Do(func() {
		cache_monadParStar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParStar(dictParallel_0_box)
})
	})
	return cache_monadParStar
}

var cache_monadParReaderT gopurs_runtime.Value
var once_monadParReaderT sync.Once
func Get_monadParReaderT() gopurs_runtime.Value {
	once_monadParReaderT.Do(func() {
		cache_monadParReaderT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParReaderT(dictParallel_0_box)
})
	})
	return cache_monadParReaderT
}

var cache_monadParMaybeT gopurs_runtime.Value
var once_monadParMaybeT sync.Once
func Get_monadParMaybeT() gopurs_runtime.Value {
	once_monadParMaybeT.Do(func() {
		cache_monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParMaybeT(dictParallel_0_box)
})
	})
	return cache_monadParMaybeT
}

var cache_monadParExceptT gopurs_runtime.Value
var once_monadParExceptT sync.Once
func Get_monadParExceptT() gopurs_runtime.Value {
	once_monadParExceptT.Do(func() {
		cache_monadParExceptT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParExceptT(dictParallel_0_box)
})
	})
	return cache_monadParExceptT
}

var cache_monadParCostar gopurs_runtime.Value
var once_monadParCostar sync.Once
func Get_monadParCostar() gopurs_runtime.Value {
	once_monadParCostar.Do(func() {
		cache_monadParCostar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParCostar(dictParallel_0_box)
})
	})
	return cache_monadParCostar
}

var cache_monadParParCont gopurs_runtime.Value
var once_monadParParCont sync.Once
func Get_monadParParCont() gopurs_runtime.Value {
	once_monadParParCont.Do(func() {
		cache_monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParParCont(dictMonadEffect_0_box)
})
	})
	return cache_monadParParCont
}

var cache_functorParCont gopurs_runtime.Value
var once_functorParCont sync.Once
func Get_functorParCont() gopurs_runtime.Value {
	once_functorParCont.Do(func() {
		cache_functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorParCont(dictMonadEffect_0_box)
})
	})
	return cache_functorParCont
}

var cache_applyParCont gopurs_runtime.Value
var once_applyParCont sync.Once
func Get_applyParCont() gopurs_runtime.Value {
	once_applyParCont.Do(func() {
		cache_applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyParCont(dictMonadEffect_0_box)
})
	})
	return cache_applyParCont
}

var cache_applicativeParCont gopurs_runtime.Value
var once_applicativeParCont sync.Once
func Get_applicativeParCont() gopurs_runtime.Value {
	once_applicativeParCont.Do(func() {
		cache_applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeParCont(dictMonadEffect_0_box)
})
	})
	return cache_applicativeParCont
}

var cache_altParCont gopurs_runtime.Value
var once_altParCont sync.Once
func Get_altParCont() gopurs_runtime.Value {
	once_altParCont.Do(func() {
		cache_altParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altParCont(dictMonadEffect_0_box)
})
	})
	return cache_altParCont
}

var cache_plusParCont gopurs_runtime.Value
var once_plusParCont sync.Once
func Get_plusParCont() gopurs_runtime.Value {
	once_plusParCont.Do(func() {
		cache_plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusParCont(dictMonadEffect_0_box)
})
	})
	return cache_plusParCont
}

var cache_alternativeParCont gopurs_runtime.Value
var once_alternativeParCont sync.Once
func Get_alternativeParCont() gopurs_runtime.Value {
	once_alternativeParCont.Do(func() {
		cache_alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeParCont(dictMonadEffect_0_box)
})
	})
	return cache_alternativeParCont
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
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

var cache_bind__3225218311 gopurs_runtime.Value
var once_bind__3225218311 sync.Once
func Get_bind__3225218311() gopurs_runtime.Value {
	once_bind__3225218311.Do(func() {
		cache_bind__3225218311 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3225218311(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3225218311
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

var cache_bind__1459396103 gopurs_runtime.Value
var once_bind__1459396103 sync.Once
func Get_bind__1459396103() gopurs_runtime.Value {
	once_bind__1459396103.Do(func() {
		cache_bind__1459396103 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1459396103(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1459396103
}

var cache_bind__877565287 gopurs_runtime.Value
var once_bind__877565287 sync.Once
func Get_bind__877565287() gopurs_runtime.Value {
	once_bind__877565287.Do(func() {
		cache_bind__877565287 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__877565287(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__877565287
}

var cache_bind__2510859911 gopurs_runtime.Value
var once_bind__2510859911 sync.Once
func Get_bind__2510859911() gopurs_runtime.Value {
	once_bind__2510859911.Do(func() {
		cache_bind__2510859911 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2510859911(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2510859911
}

var cache_bind__1918470055 gopurs_runtime.Value
var once_bind__1918470055 sync.Once
func Get_bind__1918470055() gopurs_runtime.Value {
	once_bind__1918470055.Do(func() {
		cache_bind__1918470055 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1918470055(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1918470055
}

var cache_bind__2763824711 gopurs_runtime.Value
var once_bind__2763824711 sync.Once
func Get_bind__2763824711() gopurs_runtime.Value {
	once_bind__2763824711.Do(func() {
		cache_bind__2763824711 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2763824711(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2763824711
}

var cache_discard__439597126 gopurs_runtime.Value
var once_discard__439597126 sync.Once
func Get_discard__439597126() gopurs_runtime.Value {
	once_discard__439597126.Do(func() {
		cache_discard__439597126 = gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard")
	})
	return cache_discard__439597126
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

var cache_runContT__2157838812 gopurs_runtime.Value
var once_runContT__2157838812 sync.Once
func Get_runContT__2157838812() gopurs_runtime.Value {
	once_runContT__2157838812.Do(func() {
		cache_runContT__2157838812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT__2157838812(v_0_box, k_1_box)
})
	})
	return cache_runContT__2157838812
}

var cache_runContT__3370980748 gopurs_runtime.Value
var once_runContT__3370980748 sync.Once
func Get_runContT__3370980748() gopurs_runtime.Value {
	once_runContT__3370980748.Do(func() {
		cache_runContT__3370980748 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT__3370980748(v_0_box, k_1_box)
})
	})
	return cache_runContT__3370980748
}

var cache_runContT__1411852724 gopurs_runtime.Value
var once_runContT__1411852724 sync.Once
func Get_runContT__1411852724() gopurs_runtime.Value {
	once_runContT__1411852724.Do(func() {
		cache_runContT__1411852724 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT__1411852724(v_0_box, k_1_box)
})
	})
	return cache_runContT__1411852724
}

var cache_mapReaderT__3691274100 gopurs_runtime.Value
var once_mapReaderT__3691274100 sync.Once
func Get_mapReaderT__3691274100() gopurs_runtime.Value {
	once_mapReaderT__3691274100.Do(func() {
		cache_mapReaderT__3691274100 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReaderT__3691274100(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapReaderT__3691274100
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

var cache_parallel__2242335472 gopurs_runtime.Value
var once_parallel__2242335472 sync.Once
func Get_parallel__2242335472() gopurs_runtime.Value {
	once_parallel__2242335472.Do(func() {
		cache_parallel__2242335472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__2242335472(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_parallel__2242335472
}

var cache_parallel__412321264 gopurs_runtime.Value
var once_parallel__412321264 sync.Once
func Get_parallel__412321264() gopurs_runtime.Value {
	once_parallel__412321264.Do(func() {
		cache_parallel__412321264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__412321264(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_parallel__412321264
}

var cache_parallel__1250658000 gopurs_runtime.Value
var once_parallel__1250658000 sync.Once
func Get_parallel__1250658000() gopurs_runtime.Value {
	once_parallel__1250658000.Do(func() {
		cache_parallel__1250658000 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__1250658000(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_parallel__1250658000
}

var cache_sequential__2242335472 gopurs_runtime.Value
var once_sequential__2242335472 sync.Once
func Get_sequential__2242335472() gopurs_runtime.Value {
	once_sequential__2242335472.Do(func() {
		cache_sequential__2242335472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential__2242335472(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequential__2242335472
}

var cache_sequential__412321264 gopurs_runtime.Value
var once_sequential__412321264 sync.Once
func Get_sequential__412321264() gopurs_runtime.Value {
	once_sequential__412321264.Do(func() {
		cache_sequential__412321264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential__412321264(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequential__412321264
}

var cache_sequential__1250658000 gopurs_runtime.Value
var once_sequential__1250658000 sync.Once
func Get_sequential__1250658000() gopurs_runtime.Value {
	once_sequential__1250658000.Do(func() {
		cache_sequential__1250658000 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential__1250658000(gopurs_runtime.CoerceToStruct[Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequential__1250658000
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

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
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

var cache_applyCostar__1509714460 gopurs_runtime.Value
var once_applyCostar__1509714460 sync.Once
func Get_applyCostar__1509714460() gopurs_runtime.Value {
	once_applyCostar__1509714460.Do(func() {
		cache_applyCostar__1509714460 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_functorCostar()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
})
})
}))
	})
	return cache_applyCostar__1509714460
}

var cache_functorCostar__735509168 gopurs_runtime.Value
var once_functorCostar__735509168 sync.Once
func Get_functorCostar__735509168() gopurs_runtime.Value {
	once_functorCostar__735509168.Do(func() {
		cache_functorCostar__735509168 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
}))
	})
	return cache_functorCostar__735509168
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

var cache_liftEffect__2407462165 gopurs_runtime.Value
var once_liftEffect__2407462165 sync.Once
func Get_liftEffect__2407462165() gopurs_runtime.Value {
	once_liftEffect__2407462165.Do(func() {
		cache_liftEffect__2407462165 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__2407462165(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__2407462165
}

var cache_liftEffect__3456588885 gopurs_runtime.Value
var once_liftEffect__3456588885 sync.Once
func Get_liftEffect__3456588885() gopurs_runtime.Value {
	once_liftEffect__3456588885.Do(func() {
		cache_liftEffect__3456588885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__3456588885(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__3456588885
}

var cache_liftEffect__1892566677 gopurs_runtime.Value
var once_liftEffect__1892566677 sync.Once
func Get_liftEffect__1892566677() gopurs_runtime.Value {
	once_liftEffect__1892566677.Do(func() {
		cache_liftEffect__1892566677 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__1892566677(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__1892566677
}

var cache_liftEffect__2322711157 gopurs_runtime.Value
var once_liftEffect__2322711157 sync.Once
func Get_liftEffect__2322711157() gopurs_runtime.Value {
	once_liftEffect__2322711157.Do(func() {
		cache_liftEffect__2322711157 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__2322711157(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__2322711157
}

var cache_liftEffect__735761941 gopurs_runtime.Value
var once_liftEffect__735761941 sync.Once
func Get_liftEffect__735761941() gopurs_runtime.Value {
	once_liftEffect__735761941.Do(func() {
		cache_liftEffect__735761941 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__735761941(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__735761941
}

var cache_liftEffect__3357942741 gopurs_runtime.Value
var once_liftEffect__3357942741 sync.Once
func Get_liftEffect__3357942741() gopurs_runtime.Value {
	once_liftEffect__3357942741.Do(func() {
		cache_liftEffect__3357942741 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__3357942741(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__3357942741
}

var cache_liftEffect__226852501 gopurs_runtime.Value
var once_liftEffect__226852501 sync.Once
func Get_liftEffect__226852501() gopurs_runtime.Value {
	once_liftEffect__226852501.Do(func() {
		cache_liftEffect__226852501 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__226852501(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__226852501
}

var cache_liftEffect__2550292213 gopurs_runtime.Value
var once_liftEffect__2550292213 sync.Once
func Get_liftEffect__2550292213() gopurs_runtime.Value {
	once_liftEffect__2550292213.Do(func() {
		cache_liftEffect__2550292213 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__2550292213(gopurs_runtime.CoerceToStruct[pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__2550292213
}

var cache_new__3544820218 gopurs_runtime.Value
var once_new__3544820218 sync.Once
func Get_new__3544820218() gopurs_runtime.Value {
	once_new__3544820218.Do(func() {
		cache_new__3544820218 = pkg_Effect_Ref.Get__new()
	})
	return cache_new__3544820218
}

var cache_new__1693026106 gopurs_runtime.Value
var once_new__1693026106 sync.Once
func Get_new__1693026106() gopurs_runtime.Value {
	once_new__1693026106.Do(func() {
		cache_new__1693026106 = pkg_Effect_Ref.Get__new()
	})
	return cache_new__1693026106
}

var cache_new__1017896474 gopurs_runtime.Value
var once_new__1017896474 sync.Once
func Get_new__1017896474() gopurs_runtime.Value {
	once_new__1017896474.Do(func() {
		cache_new__1017896474 = pkg_Effect_Ref.Get__new()
	})
	return cache_new__1017896474
}

var cache_new__337624346 gopurs_runtime.Value
var once_new__337624346 sync.Once
func Get_new__337624346() gopurs_runtime.Value {
	once_new__337624346.Do(func() {
		cache_new__337624346 = pkg_Effect_Ref.Get__new()
	})
	return cache_new__337624346
}

type Constructor_Parallel[T_f any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[327692956] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Apply0": return c.V0
		case "Apply1": return c.V1
		case "parallel": return c.V2
		case "sequential": return c.V3
		default: panic("Key not found in dictionary Constructor_Parallel: " + key)
		}
	}
}


func Call_ParCont(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_sequential(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_parallel(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_monadParWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applyWriterT_1_0 := gopurs_runtime.Apply(pkg_Control_Monad_Writer_Trans.Get_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyWriterT_1_0
return gopurs_runtime.Func(func(dictParallel_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyWriterT1_3_1 := gopurs_runtime.Apply(applyWriterT_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT1_3_1
applyWriterT2_4_2 := gopurs_runtime.Apply(applyWriterT_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply1"), gopurs_runtime.Value{}))
_ = applyWriterT2_4_2
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT1_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "parallel"), v_5)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "sequential"), v_5)
}))
})
}

func Call_monadParStar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
applyStar_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})
})
}))
_ = applyStar_1_0
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
functorStar1_3_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), f_4)
_ = __local_var_6_9
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_9, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_7
applyStar1_2_5 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_7
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})
})
}))
_ = applyStar1_2_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_2_5
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(v_3, x_4))
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_monadParReaderT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorReaderT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_2
applyReaderT_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})
})
}))
_ = applyReaderT_1_0
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
functorReaderT1_3_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(v_6, x_7))
})
})
}))
_ = functorReaderT1_3_7
applyReaderT1_2_5 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_7
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})
})
}))
_ = applyReaderT1_2_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_2_5
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(v_3, x_4))
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_monadParMaybeT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
functorCompose2_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "map"), f_5), v_6)
})
}))
_ = functorCompose2_4_4
applyCompose_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_4_4
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(Functor0_2_2.V0, gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), v_5), v1_6)
})
}))
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyMaybeT_3_6 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_applyMaybeT(), dictMonad_2)
_ = applyMaybeT_3_6
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMaybeT_3_6
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_1_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
}))
})
}

func Call_monadParExceptT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
functorCompose2_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "map"), f_5), v_6)
})
}))
_ = functorCompose2_4_4
applyCompose_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_4_4
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(Functor0_2_2.V0, gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "apply"), v_5), v1_6)
})
}))
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyExceptT_3_6 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_applyExceptT(), dictMonad_2)
_ = applyExceptT_3_6
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyExceptT_3_6
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_1_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
}))
})
}

func Call_monadParCostar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), x_2))
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), x_2))
})
}))
}

func Call_monadParParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
functorContT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_1
applyContT_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = applyContT_1_0
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT_1_0
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyParCont(dictMonadEffect_0)
}), Get_ParCont(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2
}))
}

func Call_functorParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
functorContT_1_0 := &pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
})}
_ = functorContT_1_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(functorContT_1_0.V0, f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_monadParParCont(dictMonadEffect_0), "parallel"), gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_monadParParCont(dictMonadEffect_0), "sequential"), x_4)))
})
}))
}

func Call_applyParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorParCont(dictMonadEffect_0)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Ref.Get__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Ref.Get__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_1_0)}, gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(pkg_Effect_Ref.Get_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_7})}, ra_5))
goto end_branch_1
} else {

}
}
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mb_8.UnsafePtr).V0))
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
})), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(pkg_Effect_Ref.Get_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, b_8})}, rb_6))
goto end_branch_2
} else {

}
}
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ma_9.UnsafePtr).V0, b_8))
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
}))
}))
}))
}))
})
})
}))
}

func Call_applicativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applyParCont1_1_0 := Call_applyParCont(dictMonadEffect_0)
_ = applyParCont1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_monadParParCont(dictMonadEffect_0), "parallel"), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
}))
}))
}

func Call_altParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
functorParCont1_4_3 := Call_functorParCont(dictMonadEffect_0)
_ = functorParCont1_4_3
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_4_3
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_new__3544820218(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(done_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_2_1)}, gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), done_8)), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (b_10.IntVal) != (0) {
__t4 = gopurs_runtime.Apply(Applicative0_3_2.V1, pkg_Data_Unit.Get_unit())
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_2_1)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(pkg_Effect_Ref.Get_write(), gopurs_runtime.Bool(true), done_8)), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_9)
}))
}
end_branch_4:
return __t4
}))
})), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), done_8)), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_11.IntVal) != (0) {
__t5 = gopurs_runtime.Apply(Applicative0_3_2.V1, pkg_Data_Unit.Get_unit())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Bind1_2_1)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(pkg_Effect_Ref.Get_write(), gopurs_runtime.Bool(true), done_8)), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_10)
}))
}
end_branch_5:
return __t5
}))
}))
}))
}))
})
})
}))
}

func Call_plusParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
altParCont1_2_1 := Call_altParCont(dictMonadEffect_0)
_ = altParCont1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
}))
}

func Call_alternativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applicativeParCont1_1_0 := Call_applicativeParCont(dictMonadEffect_0)
_ = applicativeParCont1_1_0
plusParCont1_2_1 := Call_plusParCont(dictMonadEffect_0)
_ = plusParCont1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeParCont1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusParCont1_2_1
}))
}

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3225218311(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1459396103(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__877565287(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2510859911(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1918470055(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2763824711(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_runContT__2157838812(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_runContT__3370980748(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_runContT__1411852724(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_mapReaderT__3691274100(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_mapWriterT__77717660(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_parallel__2242335472(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_parallel__412321264(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_parallel__1250658000(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_sequential__2242335472(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_sequential__412321264(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_sequential__1250658000(dict_0_loop *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_liftEffect__2407462165(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__3456588885(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__1892566677(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__2322711157(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__735761941(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__3357942741(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__226852501(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__2550292213(dict_0_loop *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Effect_Class.Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


