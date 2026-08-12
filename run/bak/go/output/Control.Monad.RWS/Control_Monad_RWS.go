package Control_Monad_RWS

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_withRWS gopurs_runtime.Value
var once_withRWS sync.Once
func Get_withRWS() gopurs_runtime.Value {
	once_withRWS.Do(func() {
		cache_withRWS = pkg_Control_Monad_RWS_Trans.Get_withRWST()
	})
	return cache_withRWS
}

var cache_rws gopurs_runtime.Value
var once_rws sync.Once
func Get_rws() gopurs_runtime.Value {
	once_rws.Do(func() {
		cache_rws = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_rws(f_0_box, r_1_box, s_2_box))}
})
	})
	return cache_rws
}

var cache_runRWS gopurs_runtime.Value
var once_runRWS sync.Once
func Get_runRWS() gopurs_runtime.Value {
	once_runRWS.Do(func() {
		cache_runRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_runRWS(m_0_box, r_1_box, s_2_box))}
})
	})
	return cache_runRWS
}

var cache_mapRWS gopurs_runtime.Value
var once_mapRWS sync.Once
func Get_mapRWS() gopurs_runtime.Value {
	once_mapRWS.Do(func() {
		cache_mapRWS = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_mapRWS(f_0_box, v_1_box, r_2_box, s_3_box))}
})
	})
	return cache_mapRWS
}

var cache_execRWS gopurs_runtime.Value
var once_execRWS sync.Once
func Get_execRWS() gopurs_runtime.Value {
	once_execRWS.Do(func() {
		cache_execRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_execRWS(m_0_box, r_1_box, s_2_box))}
})
	})
	return cache_execRWS
}

var cache_evalRWS gopurs_runtime.Value
var once_evalRWS sync.Once
func Get_evalRWS() gopurs_runtime.Value {
	once_evalRWS.Do(func() {
		cache_evalRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_evalRWS(m_0_box, r_1_box, s_2_box))}
})
	})
	return cache_evalRWS
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

var cache_evalRWST__2982438712 gopurs_runtime.Value
var once_evalRWST__2982438712 sync.Once
func Get_evalRWST__2982438712() gopurs_runtime.Value {
	once_evalRWST__2982438712.Do(func() {
		cache_evalRWST__2982438712 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalRWST__2982438712(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_evalRWST__2982438712
}

var cache_evalRWST__2802039376 gopurs_runtime.Value
var once_evalRWST__2802039376 sync.Once
func Get_evalRWST__2802039376() gopurs_runtime.Value {
	once_evalRWST__2802039376.Do(func() {
		cache_evalRWST__2802039376 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, gopurs_runtime.Apply2(v_2, r_3, s_4), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2})})
}))
})
})
})
}()
	})
	return cache_evalRWST__2802039376
}

var cache_execRWST__2982438712 gopurs_runtime.Value
var once_execRWST__2982438712 sync.Once
func Get_execRWST__2982438712() gopurs_runtime.Value {
	once_execRWST__2982438712.Do(func() {
		cache_execRWST__2982438712 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execRWST__2982438712(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_execRWST__2982438712
}

var cache_execRWST__2802039376 gopurs_runtime.Value
var once_execRWST__2802039376 sync.Once
func Get_execRWST__2802039376() gopurs_runtime.Value {
	once_execRWST__2802039376.Do(func() {
		cache_execRWST__2802039376 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, gopurs_runtime.Apply2(v_2, r_3, s_4), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2})})
}))
})
})
})
}()
	})
	return cache_execRWST__2802039376
}

var cache_mapRWST__3506688348 gopurs_runtime.Value
var once_mapRWST__3506688348 sync.Once
func Get_mapRWST__3506688348() gopurs_runtime.Value {
	once_mapRWST__3506688348.Do(func() {
		cache_mapRWST__3506688348 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapRWST__3506688348(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_mapRWST__3506688348
}

var cache_mapRWST__1363965404 gopurs_runtime.Value
var once_mapRWST__1363965404 sync.Once
func Get_mapRWST__1363965404() gopurs_runtime.Value {
	once_mapRWST__1363965404.Do(func() {
		cache_mapRWST__1363965404 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapRWST__1363965404(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_mapRWST__1363965404
}

var cache_withRWST__673207610 gopurs_runtime.Value
var once_withRWST__673207610 sync.Once
func Get_withRWST__673207610() gopurs_runtime.Value {
	once_withRWST__673207610.Do(func() {
		cache_withRWST__673207610 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_withRWST__673207610(f_0_box, m_1_box, r_2_box, s_3_box))}
})
	})
	return cache_withRWST__673207610
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

var cache_monadIdentity__2437051429 gopurs_runtime.Value
var once_monadIdentity__2437051429 sync.Once
func Get_monadIdentity__2437051429() gopurs_runtime.Value {
	once_monadIdentity__2437051429.Do(func() {
		cache_monadIdentity__2437051429 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_applicativeIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_bindIdentity()
}))
	})
	return cache_monadIdentity__2437051429
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

var cache_uncurry__2421405441 gopurs_runtime.Value
var once_uncurry__2421405441 sync.Once
func Get_uncurry__2421405441() gopurs_runtime.Value {
	once_uncurry__2421405441.Do(func() {
		cache_uncurry__2421405441 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncurry__2421405441(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_uncurry__2421405441
}

func Call_rws(f_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_applicativeIdentity(), "pure"), gopurs_runtime.Apply2(f_0, r_1, s_2)))
}

func Call_runRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(m_0, r_1, s_2))
}

func Call_mapRWS(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) *pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.CoerceToStruct[pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3)))
}

func Call_execRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
Applicative0_3_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(m_0, r_1, s_2), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_3_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2})})
})))
}

func Call_evalRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
Applicative0_3_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(m_0, r_1, s_2), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_3_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V2})})
})))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1953455120(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3465679815(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_evalRWST__2982438712(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_execRWST__2982438712(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_mapRWST__3506688348(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_mapRWST__1363965404(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_withRWST__673207610(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) *pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
__local_var_4_0 := gopurs_runtime.Apply2(f_0, r_2, s_3)
_ = __local_var_4_0
return gopurs_runtime.CoerceToStruct[pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(m_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1))
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

func Call_uncurry__3533477633(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_uncurry__2421405441(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}


