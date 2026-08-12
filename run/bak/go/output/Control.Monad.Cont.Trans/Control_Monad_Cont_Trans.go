package Control_Monad_Cont_Trans

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_Reader_Class "gopurs/output/Control.Monad.Reader.Class"
	pkg_Control_Monad_Trans_Class "gopurs/output/Control.Monad.Trans.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_ContT gopurs_runtime.Value
var once_ContT sync.Once
func Get_ContT() gopurs_runtime.Value {
	once_ContT.Do(func() {
		cache_ContT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ContT(x_0_box)
})
	})
	return cache_ContT
}

var cache_withContT gopurs_runtime.Value
var once_withContT sync.Once
func Get_withContT() gopurs_runtime.Value {
	once_withContT.Do(func() {
		cache_withContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withContT(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_withContT
}

var cache_runContT gopurs_runtime.Value
var once_runContT sync.Once
func Get_runContT() gopurs_runtime.Value {
	once_runContT.Do(func() {
		cache_runContT = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT(v_0_box, k_1_box)
})
	})
	return cache_runContT
}

var cache_newtypeContT gopurs_runtime.Value
var once_newtypeContT sync.Once
func Get_newtypeContT() gopurs_runtime.Value {
	once_newtypeContT.Do(func() {
		cache_newtypeContT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeContT
}

var cache_monadTransContT gopurs_runtime.Value
var once_monadTransContT sync.Once
func Get_monadTransContT() gopurs_runtime.Value {
	once_monadTransContT.Do(func() {
		cache_monadTransContT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, m_2, k_3)
})
})
}))
	})
	return cache_monadTransContT
}

var cache_mapContT gopurs_runtime.Value
var once_mapContT sync.Once
func Get_mapContT() gopurs_runtime.Value {
	once_mapContT.Do(func() {
		cache_mapContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapContT(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_mapContT
}

var cache_functorContT gopurs_runtime.Value
var once_functorContT sync.Once
func Get_functorContT() gopurs_runtime.Value {
	once_functorContT.Do(func() {
		cache_functorContT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorContT(dictFunctor_0_box)
})
	})
	return cache_functorContT
}

var cache_applyContT gopurs_runtime.Value
var once_applyContT sync.Once
func Get_applyContT() gopurs_runtime.Value {
	once_applyContT.Do(func() {
		cache_applyContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyContT(dictApply_0_box)
})
	})
	return cache_applyContT
}

var cache_bindContT gopurs_runtime.Value
var once_bindContT sync.Once
func Get_bindContT() gopurs_runtime.Value {
	once_bindContT.Do(func() {
		cache_bindContT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindContT(dictBind_0_box)
})
	})
	return cache_bindContT
}

var cache_semigroupContT gopurs_runtime.Value
var once_semigroupContT sync.Once
func Get_semigroupContT() gopurs_runtime.Value {
	once_semigroupContT.Do(func() {
		cache_semigroupContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupContT(dictApply_0_box)
})
	})
	return cache_semigroupContT
}

var cache_applicativeContT gopurs_runtime.Value
var once_applicativeContT sync.Once
func Get_applicativeContT() gopurs_runtime.Value {
	once_applicativeContT.Do(func() {
		cache_applicativeContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeContT(dictApplicative_0_box)
})
	})
	return cache_applicativeContT
}

var cache_monadContT gopurs_runtime.Value
var once_monadContT sync.Once
func Get_monadContT() gopurs_runtime.Value {
	once_monadContT.Do(func() {
		cache_monadContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContT(dictMonad_0_box)
})
	})
	return cache_monadContT
}

var cache_monadAskContT gopurs_runtime.Value
var once_monadAskContT sync.Once
func Get_monadAskContT() gopurs_runtime.Value {
	once_monadAskContT.Do(func() {
		cache_monadAskContT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskContT(dictMonadAsk_0_box)
})
	})
	return cache_monadAskContT
}

var cache_monadReaderContT gopurs_runtime.Value
var once_monadReaderContT sync.Once
func Get_monadReaderContT() gopurs_runtime.Value {
	once_monadReaderContT.Do(func() {
		cache_monadReaderContT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderContT(dictMonadReader_0_box)
})
	})
	return cache_monadReaderContT
}

var cache_monadContContT gopurs_runtime.Value
var once_monadContContT sync.Once
func Get_monadContContT() gopurs_runtime.Value {
	once_monadContContT.Do(func() {
		cache_monadContContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContContT(dictMonad_0_box)
})
	})
	return cache_monadContContT
}

var cache_monadEffectContT gopurs_runtime.Value
var once_monadEffectContT sync.Once
func Get_monadEffectContT() gopurs_runtime.Value {
	once_monadEffectContT.Do(func() {
		cache_monadEffectContT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectContT(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectContT
}

var cache_monadStateContT gopurs_runtime.Value
var once_monadStateContT sync.Once
func Get_monadStateContT() gopurs_runtime.Value {
	once_monadStateContT.Do(func() {
		cache_monadStateContT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateContT(dictMonadState_0_box)
})
	})
	return cache_monadStateContT
}

var cache_monadSTContT gopurs_runtime.Value
var once_monadSTContT sync.Once
func Get_monadSTContT() gopurs_runtime.Value {
	once_monadSTContT.Do(func() {
		cache_monadSTContT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTContT(dictMonadST_0_box)
})
	})
	return cache_monadSTContT
}

var cache_monoidContT gopurs_runtime.Value
var once_monoidContT sync.Once
func Get_monoidContT() gopurs_runtime.Value {
	once_monoidContT.Do(func() {
		cache_monoidContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidContT(dictApplicative_0_box)
})
	})
	return cache_monoidContT
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

var cache_pure__1476406672 gopurs_runtime.Value
var once_pure__1476406672 sync.Once
func Get_pure__1476406672() gopurs_runtime.Value {
	once_pure__1476406672.Do(func() {
		cache_pure__1476406672 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1476406672(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1476406672
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

var cache_lift2__1699116848 gopurs_runtime.Value
var once_lift2__1699116848 sync.Once
func Get_lift2__1699116848() gopurs_runtime.Value {
	once_lift2__1699116848.Do(func() {
		cache_lift2__1699116848 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__1699116848(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__1699116848
}

var cache_bind__2879969985 gopurs_runtime.Value
var once_bind__2879969985 sync.Once
func Get_bind__2879969985() gopurs_runtime.Value {
	once_bind__2879969985.Do(func() {
		cache_bind__2879969985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2879969985(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2879969985
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

var cache_monadTransContT__1045393380 gopurs_runtime.Value
var once_monadTransContT__1045393380 sync.Once
func Get_monadTransContT__1045393380() gopurs_runtime.Value {
	once_monadTransContT__1045393380.Do(func() {
		cache_monadTransContT__1045393380 = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, m_2, k_3)
})
})
}))
	})
	return cache_monadTransContT__1045393380
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

var cache_lift__1331755881 gopurs_runtime.Value
var once_lift__1331755881 sync.Once
func Get_lift__1331755881() gopurs_runtime.Value {
	once_lift__1331755881.Do(func() {
		cache_lift__1331755881 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__1331755881(dictMonad_0_box)
})
	})
	return cache_lift__1331755881
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

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
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

func Call_ContT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
}

func Call_runContT(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_mapContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_functorContT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
}

func Call_applyContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
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
}

func Call_bindContT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
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
applyContT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
_ = applyContT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_3, a_5, k_prime_4)
}))
})
})
}))
}

func Call_semigroupContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
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
applyContT1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
})))
_ = applyContT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(applyContT1_1_0.V0, gopurs_runtime.Value{}))
_ = Functor0_3_2
__local_var_4_3 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_3
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(applyContT1_1_0.V1, gopurs_runtime.Apply2(Functor0_3_2.V0, __local_var_4_3, a_5), b_6)
})
}))
})
}

func Call_applicativeContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
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
applyContT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
_ = applyContT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})
}))
}

func Call_monadContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
functorContT1_1_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_2
applyContT1_1_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_2
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
_ = applyContT1_1_1
applicativeContT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_1
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})
}))
_ = applicativeContT1_1_0
functorContT1_2_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})
})
}))
_ = functorContT1_2_5
applyContT1_2_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_2_5
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})
})
}))
_ = applyContT1_2_4
bindContT1_2_3 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_2_4
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime_5)
}))
})
})
}))
_ = bindContT1_2_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeContT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindContT1_2_3
}))
}

func Call_monadAskContT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
monadContT1_1_0 := Call_monadContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadContT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_1_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_monadReaderContT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
MonadAsk0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = MonadAsk0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
ask_3_2 := gopurs_runtime.RecordGet(MonadAsk0_1_0, "ask")
_ = ask_3_2
monadAskContT1_4_3 := Call_monadAskContT(MonadAsk0_1_0)
_ = monadAskContT1_4_3
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskContT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, ask_3_2, gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return r_8
}))
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_5, gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_4, gopurs_runtime.Apply(k_7, x_10))
})))
}))
})
})
}))
}

func Call_monadContContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadContT1_1_0 := Call_monadContT(dictMonad_0)
_ = monadContT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_4)
})
}), k_3)
})
}))
}

func Call_monadEffectContT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_monadStateContT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_4))
}))
}

func Call_monadSTContT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_monoidContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
functorContT1_1_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_2
applyContT1_1_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_2
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
_ = applyContT1_1_1
applicativeContT1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_1
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})
})))
_ = applicativeContT1_1_0
functorContT1_2_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})
})
}))
_ = functorContT1_2_4
applyContT1_2_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_2_4
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})
})
})))
_ = applyContT1_2_3
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_6 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(applyContT1_2_3.V0, gopurs_runtime.Value{}))
_ = Functor0_4_6
__local_var_5_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_5_7
semigroupContT2_4_5 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(applyContT1_2_3.V1, gopurs_runtime.Apply2(Functor0_4_6.V0, __local_var_5_7, a_6), b_7)
})
}))
_ = semigroupContT2_4_5
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupContT2_4_5
}), gopurs_runtime.Apply(applicativeContT1_1_0.V1, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")))
})
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1476406672(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_lift2__1699116848(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_bind__2879969985(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_local__1299460031(dict_0_loop *pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Reader_Class.Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift__3816229929(dict_0_loop *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Trans_Class.Constructor_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_lift__1331755881(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, m_2, k_3)
})
})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
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

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}


