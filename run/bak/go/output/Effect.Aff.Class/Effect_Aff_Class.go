package Effect_Aff_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Control_Monad_Cont_Trans "gopurs/output/Control.Monad.Cont.Trans"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Control_Monad_List_Trans "gopurs/output/Control.Monad.List.Trans"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
	pkg_Control_Monad_Writer_Trans "gopurs/output/Control.Monad.Writer.Trans"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_monadAffAff gopurs_runtime.Value
var once_monadAffAff sync.Once
func Get_monadAffAff() gopurs_runtime.Value {
	once_monadAffAff.Do(func() {
		cache_monadAffAff = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_monadEffectAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))))
	})
	return cache_monadAffAff
}

var cache_liftAff gopurs_runtime.Value
var once_liftAff sync.Once
func Get_liftAff() gopurs_runtime.Value {
	once_liftAff.Do(func() {
		cache_liftAff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftAff(dict_0_box)
})
	})
	return cache_liftAff
}

var cache_monadAffContT gopurs_runtime.Value
var once_monadAffContT sync.Once
func Get_monadAffContT() gopurs_runtime.Value {
	once_monadAffContT.Do(func() {
		cache_monadAffContT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadAffContT(dictMonadAff_0_box))
})
	})
	return cache_monadAffContT
}

var cache_monadAffExceptT gopurs_runtime.Value
var once_monadAffExceptT sync.Once
func Get_monadAffExceptT() gopurs_runtime.Value {
	once_monadAffExceptT.Do(func() {
		cache_monadAffExceptT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadAffExceptT(dictMonadAff_0_box))
})
	})
	return cache_monadAffExceptT
}

var cache_monadAffListT gopurs_runtime.Value
var once_monadAffListT sync.Once
func Get_monadAffListT() gopurs_runtime.Value {
	once_monadAffListT.Do(func() {
		cache_monadAffListT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadAffListT(dictMonadAff_0_box))
})
	})
	return cache_monadAffListT
}

var cache_monadAffMaybe gopurs_runtime.Value
var once_monadAffMaybe sync.Once
func Get_monadAffMaybe() gopurs_runtime.Value {
	once_monadAffMaybe.Do(func() {
		cache_monadAffMaybe = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadAffMaybe(dictMonadAff_0_box))
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
return gopurs_runtime.Any(Call_monadAffReader(dictMonadAff_0_box))
})
	})
	return cache_monadAffReader
}

var cache_monadAffState gopurs_runtime.Value
var once_monadAffState sync.Once
func Get_monadAffState() gopurs_runtime.Value {
	once_monadAffState.Do(func() {
		cache_monadAffState = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadAffState(dictMonadAff_0_box))
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

func Call_liftAff(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "liftAff")
}

func Call_monadAffContT(dictMonadAff_0_loop gopurs_runtime.Value) interface{} {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectContT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Cont_Trans.Get_monadEffectContT(), MonadEffect0_1_0)
_ = monadEffectContT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Cont_Trans.Get_monadTransContT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectContT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
})))
}

func Call_monadAffExceptT(dictMonadAff_0_loop gopurs_runtime.Value) interface{} {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectExceptT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_monadEffectExceptT(), MonadEffect0_1_0)
_ = monadEffectExceptT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Except_Trans.Get_monadTransExceptT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectExceptT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
})))
}

func Call_monadAffListT(dictMonadAff_0_loop gopurs_runtime.Value) interface{} {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectListT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_List_Trans.Get_monadEffectListT(), MonadEffect0_1_0)
_ = monadEffectListT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_List_Trans.Get_monadTransListT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectListT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
})))
}

func Call_monadAffMaybe(dictMonadAff_0_loop gopurs_runtime.Value) interface{} {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectMaybe_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_monadEffectMaybe(), MonadEffect0_1_0)
_ = monadEffectMaybe_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Maybe_Trans.Get_monadTransMaybeT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectMaybe_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
})))
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
mempty_5_3 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = mempty_5_3
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectRWS_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_6)
_ = __local_var_7_4
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}), "bind"), __local_var_7_4, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(s_9), gopurs_runtime.UnboxAny(a_10), gopurs_runtime.UnboxAny(mempty_5_3)})}))
}))
})
}))
})
}

func Call_monadAffReader(dictMonadAff_0_loop gopurs_runtime.Value) interface{} {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectReader_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Reader_Trans.Get_monadEffectReader(), MonadEffect0_1_0)
_ = monadEffectReader_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_Reader_Trans.Get_monadTransReaderT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectReader_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
})))
}

func Call_monadAffState(dictMonadAff_0_loop gopurs_runtime.Value) interface{} {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
monadEffectState_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_State_Trans.Get_monadEffectState(), MonadEffect0_1_0)
_ = monadEffectState_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_State_Trans.Get_monadTransStateT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectState_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
})))
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
mempty_5_3 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = mempty_5_3
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectWriter_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_7), gopurs_runtime.UnboxAny(mempty_5_3)})}))
}))
}))
})
}
