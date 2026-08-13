package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Aff_Class_lift gopurs_runtime.Value
var once_Effect_Aff_Class_lift sync.Once
func Get_Effect_Aff_Class_lift() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift.Do(func() {
		cache_Effect_Aff_Class_lift = gopurs_runtime.RecordGet(Get_Control_Monad_Cont_Trans_monadTransContT(), "lift")
	})
	return cache_Effect_Aff_Class_lift
}

var cache_Effect_Aff_Class_lift1 gopurs_runtime.Value
var once_Effect_Aff_Class_lift1 sync.Once
func Get_Effect_Aff_Class_lift1() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift1.Do(func() {
		cache_Effect_Aff_Class_lift1 = gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift")
	})
	return cache_Effect_Aff_Class_lift1
}

var cache_Effect_Aff_Class_lift2 gopurs_runtime.Value
var once_Effect_Aff_Class_lift2 sync.Once
func Get_Effect_Aff_Class_lift2() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift2.Do(func() {
		cache_Effect_Aff_Class_lift2 = gopurs_runtime.RecordGet(Get_Control_Monad_List_Trans_monadTransListT(), "lift")
	})
	return cache_Effect_Aff_Class_lift2
}

var cache_Effect_Aff_Class_lift3 gopurs_runtime.Value
var once_Effect_Aff_Class_lift3 sync.Once
func Get_Effect_Aff_Class_lift3() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift3.Do(func() {
		cache_Effect_Aff_Class_lift3 = gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift")
	})
	return cache_Effect_Aff_Class_lift3
}

var cache_Effect_Aff_Class_lift4 gopurs_runtime.Value
var once_Effect_Aff_Class_lift4 sync.Once
func Get_Effect_Aff_Class_lift4() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift4.Do(func() {
		cache_Effect_Aff_Class_lift4 = gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift")
	})
	return cache_Effect_Aff_Class_lift4
}

var cache_Effect_Aff_Class_lift5 gopurs_runtime.Value
var once_Effect_Aff_Class_lift5 sync.Once
func Get_Effect_Aff_Class_lift5() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift5.Do(func() {
		cache_Effect_Aff_Class_lift5 = gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift")
	})
	return cache_Effect_Aff_Class_lift5
}

var cache_Effect_Aff_Class_MonadAff_dollarDict gopurs_runtime.Value
var once_Effect_Aff_Class_MonadAff_dollarDict sync.Once
func Get_Effect_Aff_Class_MonadAff_dollarDict() gopurs_runtime.Value {
	once_Effect_Aff_Class_MonadAff_dollarDict.Do(func() {
		cache_Effect_Aff_Class_MonadAff_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_MonadAff_dollarDict(x_0_box)
})
	})
	return cache_Effect_Aff_Class_MonadAff_dollarDict
}

var cache_Effect_Aff_Class_monadAffAff gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffAff sync.Once
func Get_Effect_Aff_Class_monadAffAff() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffAff.Do(func() {
		cache_Effect_Aff_Class_monadAffAff = gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_Aff_monadEffectAff()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Effect_Aff_Class_monadAffAff
}

var cache_Effect_Aff_Class_liftAff gopurs_runtime.Value
var once_Effect_Aff_Class_liftAff sync.Once
func Get_Effect_Aff_Class_liftAff() gopurs_runtime.Value {
	once_Effect_Aff_Class_liftAff.Do(func() {
		cache_Effect_Aff_Class_liftAff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff](dict_0_box))
})
	})
	return cache_Effect_Aff_Class_liftAff
}

var cache_Effect_Aff_Class_monadAffContT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffContT sync.Once
func Get_Effect_Aff_Class_monadAffContT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffContT.Do(func() {
		cache_Effect_Aff_Class_monadAffContT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffContT(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffContT
}

var cache_Effect_Aff_Class_monadAffExceptT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffExceptT sync.Once
func Get_Effect_Aff_Class_monadAffExceptT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffExceptT.Do(func() {
		cache_Effect_Aff_Class_monadAffExceptT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffExceptT(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffExceptT
}

var cache_Effect_Aff_Class_monadAffListT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffListT sync.Once
func Get_Effect_Aff_Class_monadAffListT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffListT.Do(func() {
		cache_Effect_Aff_Class_monadAffListT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffListT(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffListT
}

var cache_Effect_Aff_Class_monadAffMaybe gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffMaybe sync.Once
func Get_Effect_Aff_Class_monadAffMaybe() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffMaybe.Do(func() {
		cache_Effect_Aff_Class_monadAffMaybe = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffMaybe(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffMaybe
}

var cache_Effect_Aff_Class_monadAffRWS gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffRWS sync.Once
func Get_Effect_Aff_Class_monadAffRWS() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffRWS.Do(func() {
		cache_Effect_Aff_Class_monadAffRWS = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffRWS(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffRWS
}

var cache_Effect_Aff_Class_monadAffReader gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffReader sync.Once
func Get_Effect_Aff_Class_monadAffReader() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffReader.Do(func() {
		cache_Effect_Aff_Class_monadAffReader = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffReader(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffReader
}

var cache_Effect_Aff_Class_monadAffState gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffState sync.Once
func Get_Effect_Aff_Class_monadAffState() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffState.Do(func() {
		cache_Effect_Aff_Class_monadAffState = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffState(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffState
}

var cache_Effect_Aff_Class_monadAffWriter gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffWriter sync.Once
func Get_Effect_Aff_Class_monadAffWriter() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffWriter.Do(func() {
		cache_Effect_Aff_Class_monadAffWriter = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Aff_Class_monadAffWriter(dictMonadAff_0_box)
})
	})
	return cache_Effect_Aff_Class_monadAffWriter
}

type Constructor_Effect_Aff_Class_MonadAff struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3183257445] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Effect_Aff_Class_MonadAff)(ptr)
		_ = c
		switch key {
		case "MonadEffect0": return gopurs_runtime.Box(c.V0)
		case "liftAff": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Effect_Aff_Class_MonadAff: " + key)
		}
	}
}


func Call_Effect_Aff_Class_MonadAff_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Effect_Aff_Class_liftAff(dict_0_loop *Constructor_Effect_Aff_Class_MonadAff) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Aff_Class_MonadAff = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Aff_Class_monadAffContT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectContT_2_1 -> gopurs_runtime.Value
monadEffectContT_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Cont_Trans_monadEffectContT(), MonadEffect0_1_0)
_ = monadEffectContT_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Cont_Trans_monadTransContT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectContT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_Effect_Aff_Class_monadAffExceptT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectExceptT_2_1 -> gopurs_runtime.Value
monadEffectExceptT_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_monadEffectExceptT(), MonadEffect0_1_0)
_ = monadEffectExceptT_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectExceptT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_Effect_Aff_Class_monadAffListT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectListT_2_1 -> gopurs_runtime.Value
monadEffectListT_2_1 := gopurs_runtime.Apply(Get_Control_Monad_List_Trans_monadEffectListT(), MonadEffect0_1_0)
_ = monadEffectListT_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_List_Trans_monadTransListT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectListT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_Effect_Aff_Class_monadAffMaybe(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectMaybe_2_1 -> gopurs_runtime.Value
monadEffectMaybe_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Maybe_Trans_monadEffectMaybe(), MonadEffect0_1_0)
_ = monadEffectMaybe_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectMaybe_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_Effect_Aff_Class_monadAffRWS(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): Monad0_2_1 -> gopurs_runtime.Value
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadEffectRWS_4_2 -> gopurs_runtime.Value
monadEffectRWS_4_2 := gopurs_runtime.Apply2(Get_Control_Monad_RWS_Trans_monadEffectRWS(), dictMonoid_3, MonadEffect0_1_0)
_ = monadEffectRWS_4_2
// TAST (Let): Bind1_5_4 -> *Constructor_Control_Bind_Bind
Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_4
// TAST (Let): pure_6_5 -> gopurs_runtime.Value
pure_6_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_5
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_4.V1), m_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_5, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_9, a_10, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
}))
})
})
})
_ = __local_var_5_3
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectRWS_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_6))
}))
})
}

func Call_Effect_Aff_Class_monadAffReader(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectReader_2_1 -> gopurs_runtime.Value
monadEffectReader_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Reader_Trans_monadEffectReader(), MonadEffect0_1_0)
_ = monadEffectReader_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectReader_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_Effect_Aff_Class_monadAffState(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): monadEffectState_2_1 -> gopurs_runtime.Value
monadEffectState_2_1 := gopurs_runtime.Apply(Get_Control_Monad_State_Trans_monadEffectState(), MonadEffect0_1_0)
_ = monadEffectState_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectState_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_4))
}))
}

func Call_Effect_Aff_Class_monadAffWriter(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
_ = dictMonadAff_0
// TAST (Let): MonadEffect0_1_0 -> gopurs_runtime.Value
MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
_ = MonadEffect0_1_0
// TAST (Let): Monad0_2_1 -> gopurs_runtime.Value
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadEffectWriter_4_2 -> gopurs_runtime.Value
monadEffectWriter_4_2 := gopurs_runtime.Apply2(Get_Control_Monad_Writer_Trans_monadEffectWriter(), dictMonoid_3, MonadEffect0_1_0)
_ = monadEffectWriter_4_2
// TAST (Let): Bind1_5_4 -> *Constructor_Control_Bind_Bind
Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_4
// TAST (Let): pure_6_5 -> gopurs_runtime.Value
pure_6_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_5
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_4.V1), m_7, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
}))
})
_ = __local_var_5_3
return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectWriter_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "liftAff"), x_6))
}))
})
}


