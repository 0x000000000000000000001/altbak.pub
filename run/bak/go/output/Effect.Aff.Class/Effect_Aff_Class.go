package Effect_Aff_Class

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Cont_Trans "gopurs/output/Control.Monad.Cont.Trans"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Control_Monad_List_Trans "gopurs/output/Control.Monad.List.Trans"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
	pkg_Control_Monad_Writer_Trans "gopurs/output/Control.Monad.Writer.Trans"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadAffAff gopurs_runtime.Value
var once_monadAffAff sync.Once
func Get_monadAffAff() gopurs_runtime.Value {
	once_monadAffAff.Do(func() {
		cache_monadAffAff = gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_monadEffectAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_monadAffAff
}

var cache_liftAff gopurs_runtime.Value
var once_liftAff sync.Once
func Get_liftAff() gopurs_runtime.Value {
	once_liftAff.Do(func() {
		cache_liftAff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftAff(gopurs_runtime.CoerceToStruct[Constructor_MonadAff[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftAff
}

var cache_monadAffContT gopurs_runtime.Value
var once_monadAffContT sync.Once
func Get_monadAffContT() gopurs_runtime.Value {
	once_monadAffContT.Do(func() {
		cache_monadAffContT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffContT(dictMonadAff_0_box)
})
	})
	return cache_monadAffContT
}

var cache_monadAffExceptT gopurs_runtime.Value
var once_monadAffExceptT sync.Once
func Get_monadAffExceptT() gopurs_runtime.Value {
	once_monadAffExceptT.Do(func() {
		cache_monadAffExceptT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffExceptT(dictMonadAff_0_box)
})
	})
	return cache_monadAffExceptT
}

var cache_monadAffListT gopurs_runtime.Value
var once_monadAffListT sync.Once
func Get_monadAffListT() gopurs_runtime.Value {
	once_monadAffListT.Do(func() {
		cache_monadAffListT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffListT(dictMonadAff_0_box)
})
	})
	return cache_monadAffListT
}

var cache_monadAffMaybe gopurs_runtime.Value
var once_monadAffMaybe sync.Once
func Get_monadAffMaybe() gopurs_runtime.Value {
	once_monadAffMaybe.Do(func() {
		cache_monadAffMaybe = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffMaybe(dictMonadAff_0_box)
})
	})
	return cache_monadAffMaybe
}

var cache_monadAffRWS gopurs_runtime.Value
var once_monadAffRWS sync.Once
func Get_monadAffRWS() gopurs_runtime.Value {
	once_monadAffRWS.Do(func() {
		cache_monadAffRWS = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffRWS(dictMonadAff_0_box)
})
	})
	return cache_monadAffRWS
}

var cache_monadAffReader gopurs_runtime.Value
var once_monadAffReader sync.Once
func Get_monadAffReader() gopurs_runtime.Value {
	once_monadAffReader.Do(func() {
		cache_monadAffReader = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffReader(dictMonadAff_0_box)
})
	})
	return cache_monadAffReader
}

var cache_monadAffState gopurs_runtime.Value
var once_monadAffState sync.Once
func Get_monadAffState() gopurs_runtime.Value {
	once_monadAffState.Do(func() {
		cache_monadAffState = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffState(dictMonadAff_0_box)
})
	})
	return cache_monadAffState
}

var cache_monadAffWriter gopurs_runtime.Value
var once_monadAffWriter sync.Once
func Get_monadAffWriter() gopurs_runtime.Value {
	once_monadAffWriter.Do(func() {
		cache_monadAffWriter = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAffWriter(dictMonadAff_0_box)
})
	})
	return cache_monadAffWriter
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

var cache_applicativeAff__156155496 gopurs_runtime.Value
var once_applicativeAff__156155496 sync.Once
func Get_applicativeAff__156155496() gopurs_runtime.Value {
	once_applicativeAff__156155496.Do(func() {
		cache_applicativeAff__156155496 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applyAff()
}), pkg_Effect_Aff.Get__pure())
	})
	return cache_applicativeAff__156155496
}

var cache_applyAff__2964533948 gopurs_runtime.Value
var once_applyAff__2964533948 sync.Once
func Get_applyAff__2964533948() gopurs_runtime.Value {
	once_applyAff__2964533948.Do(func() {
		cache_applyAff__2964533948 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadAff(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadAff(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_functorAff()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyAff__2964533948
}

var cache_bindAff__1025486311 gopurs_runtime.Value
var once_bindAff__1025486311 sync.Once
func Get_bindAff__1025486311() gopurs_runtime.Value {
	once_bindAff__1025486311.Do(func() {
		cache_bindAff__1025486311 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applyAff()
}), pkg_Effect_Aff.Get__bind())
	})
	return cache_bindAff__1025486311
}

var cache_functorAff__2378915857 gopurs_runtime.Value
var once_functorAff__2378915857 sync.Once
func Get_functorAff__2378915857() gopurs_runtime.Value {
	once_functorAff__2378915857.Do(func() {
		cache_functorAff__2378915857 = gopurs_runtime.RecordDict1("map", pkg_Effect_Aff.Get__map())
	})
	return cache_functorAff__2378915857
}

var cache_monadAff__2914113427 gopurs_runtime.Value
var once_monadAff__2914113427 sync.Once
func Get_monadAff__2914113427() gopurs_runtime.Value {
	once_monadAff__2914113427.Do(func() {
		cache_monadAff__2914113427 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_applicativeAff()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_bindAff()
}))
	})
	return cache_monadAff__2914113427
}

var cache_monadEffectAff__1856968838 gopurs_runtime.Value
var once_monadEffectAff__1856968838 sync.Once
func Get_monadEffectAff__1856968838() gopurs_runtime.Value {
	once_monadEffectAff__1856968838.Do(func() {
		cache_monadEffectAff__1856968838 = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_monadAff()
}), pkg_Effect_Aff.Get__liftEffect())
	})
	return cache_monadEffectAff__1856968838
}

type Constructor_MonadAff[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3183257445] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadAff[gopurs_runtime.Value])(ptr)
		switch key {
		case "MonadEffect0": return c.V0
		case "liftAff": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadAff: " + key)
		}
	}
}


func Call_liftAff(dict_0_loop *Constructor_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadAff[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_monadAffContT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectContT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Cont_Trans.Get_monadEffectContT(), MonadEffect0_1_0)
_ = monadEffectContT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Cont_Trans.Get_monadTransContT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectContT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_monadAffExceptT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectExceptT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_monadEffectExceptT(), MonadEffect0_1_0)
_ = monadEffectExceptT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Except_Trans.Get_monadTransExceptT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectExceptT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_monadAffListT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectListT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_List_Trans.Get_monadEffectListT(), MonadEffect0_1_0)
_ = monadEffectListT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_List_Trans.Get_monadTransListT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectListT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_monadAffMaybe(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectMaybe_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_monadEffectMaybe(), MonadEffect0_1_0)
_ = monadEffectMaybe_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Maybe_Trans.Get_monadTransMaybeT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectMaybe_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_monadAffRWS(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadEffectRWS_4_2 := gopurs_runtime.Apply2(pkg_Control_Monad_RWS_Trans.Get_monadEffectRWS(), dictMonoid_3, MonadEffect0_1_0)
_ = monadEffectRWS_4_2
Bind1_5_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
pure_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_4
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectRWS_4_2
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_7)
_ = __local_var_8_5
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_3.V1, __local_var_8_5, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, a_11, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
}))
})
})
}))
})
}

func Call_monadAffReader(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectReader_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Reader_Trans.Get_monadEffectReader(), MonadEffect0_1_0)
_ = monadEffectReader_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Reader_Trans.Get_monadTransReaderT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectReader_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_monadAffState(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectState_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_State_Trans.Get_monadEffectState(), MonadEffect0_1_0)
_ = monadEffectState_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_State_Trans.Get_monadTransStateT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectState_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_monadAffWriter(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadEffectWriter_4_2 := gopurs_runtime.Apply2(pkg_Control_Monad_Writer_Trans.Get_monadEffectWriter(), dictMonoid_3, MonadEffect0_1_0)
_ = monadEffectWriter_4_2
Bind1_5_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
pure_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_4
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectWriter_4_2
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_3.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
}))
}))
})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


