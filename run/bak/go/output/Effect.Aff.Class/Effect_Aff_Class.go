package Effect_Aff_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Control_Monad_Cont_Trans "gopurs/output/Control.Monad.Cont.Trans"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Control_Monad_List_Trans "gopurs/output/Control.Monad.List.Trans"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
	pkg_Control_Monad_Writer_Trans "gopurs/output/Control.Monad.Writer.Trans"
)

var monadAffAff gopurs_runtime.Value
var once_monadAffAff sync.Once
func Get_monadAffAff() gopurs_runtime.Value {
	once_monadAffAff.Do(func() {
		monadAffAff = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"], "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect_Aff.Get_monadEffectAff()
})})
	})
	return monadAffAff
}

var liftAff gopurs_runtime.Value
var once_liftAff sync.Once
func Get_liftAff() gopurs_runtime.Value {
	once_liftAff.Do(func() {
		liftAff = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"]
})
	})
	return liftAff
}

var monadAffContT gopurs_runtime.Value
var once_monadAffContT sync.Once
func Get_monadAffContT() gopurs_runtime.Value {
	once_monadAffContT.Do(func() {
		monadAffContT = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
monadEffectContT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Cont_Trans.Get_monadEffectContT(), MonadEffect0_1_0)
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_4))
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectContT_2_1
})})
})
	})
	return monadAffContT
}

var monadAffExceptT gopurs_runtime.Value
var once_monadAffExceptT sync.Once
func Get_monadAffExceptT() gopurs_runtime.Value {
	once_monadAffExceptT.Do(func() {
		monadAffExceptT = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
monadEffectExceptT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_monadEffectExceptT(), MonadEffect0_1_0)
__local_var_3_2 := gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_4)), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_5}))
}))
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectExceptT_2_1
})})
})
	})
	return monadAffExceptT
}

var monadAffListT gopurs_runtime.Value
var once_monadAffListT sync.Once
func Get_monadAffListT() gopurs_runtime.Value {
	once_monadAffListT.Do(func() {
		monadAffListT = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
monadEffectListT_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_List_Trans.Get_monadEffectListT(), MonadEffect0_1_0)
__local_var_3_2 := gopurs_runtime.Apply(pkg_Control_Monad_List_Trans.Get_fromEffect(), gopurs_runtime.Apply(gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_4))
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectListT_2_1
})})
})
	})
	return monadAffListT
}

var monadAffMaybe gopurs_runtime.Value
var once_monadAffMaybe sync.Once
func Get_monadAffMaybe() gopurs_runtime.Value {
	once_monadAffMaybe.Do(func() {
		monadAffMaybe = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
monadEffectMaybe_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_monadEffectMaybe(), MonadEffect0_1_0)
__local_var_3_2 := gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_4)), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": a_prime_5}))
}))
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectMaybe_2_1
})})
})
	})
	return monadAffMaybe
}

var monadAffRWS gopurs_runtime.Value
var once_monadAffRWS sync.Once
func Get_monadAffRWS() gopurs_runtime.Value {
	once_monadAffRWS.Do(func() {
		monadAffRWS = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
Monad0_2_1 := gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadEffectRWS_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_RWS_Trans.Get_monadEffectRWS(), dictMonoid_3), MonadEffect0_1_0)
mempty_5_3 := dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_6)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_7_4), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_9, "value1": a_10, "value2": mempty_5_3}))
}))
})
})
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectRWS_4_2
})})
})
})
	})
	return monadAffRWS
}

var monadAffReader gopurs_runtime.Value
var once_monadAffReader sync.Once
func Get_monadAffReader() gopurs_runtime.Value {
	once_monadAffReader.Do(func() {
		monadAffReader = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadEffectReader_1_0 := gopurs_runtime.Apply(pkg_Control_Monad_Reader_Trans.Get_monadEffectReader(), gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
})
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectReader_1_0
})})
})
	})
	return monadAffReader
}

var monadAffState gopurs_runtime.Value
var once_monadAffState sync.Once
func Get_monadAffState() gopurs_runtime.Value {
	once_monadAffState.Do(func() {
		monadAffState = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
monadEffectState_2_1 := gopurs_runtime.Apply(pkg_Control_Monad_State_Trans.Get_monadEffectState(), MonadEffect0_1_0)
__local_var_3_2 := gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_4)
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_5_3), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_7, "value1": s_6}))
}))
})
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectState_2_1
})})
})
	})
	return monadAffState
}

var monadAffWriter gopurs_runtime.Value
var once_monadAffWriter sync.Once
func Get_monadAffWriter() gopurs_runtime.Value {
	once_monadAffWriter.Do(func() {
		monadAffWriter = gopurs_runtime.Func(func(dictMonadAff_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadEffect0_1_0 := gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadEffect0"], gopurs_runtime.Value{})
Monad0_2_1 := gopurs_runtime.Apply(MonadEffect0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadEffectWriter_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_Writer_Trans.Get_monadEffectWriter(), dictMonoid_3), MonadEffect0_1_0)
mempty_5_3 := dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftAff": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadAff_0.PtrVal.(map[string]gopurs_runtime.Value)["liftAff"], x_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_5_3}))
}))
}), "MonadEffect0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadEffectWriter_4_2
})})
})
})
	})
	return monadAffWriter
}


