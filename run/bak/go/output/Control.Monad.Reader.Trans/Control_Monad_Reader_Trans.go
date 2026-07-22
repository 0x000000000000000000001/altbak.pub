package Control_Monad_Reader_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var ReaderT gopurs_runtime.Value
var once_ReaderT sync.Once
func Get_ReaderT() gopurs_runtime.Value {
	once_ReaderT.Do(func() {
		ReaderT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return ReaderT
}

var withReaderT gopurs_runtime.Value
var once_withReaderT sync.Once
func Get_withReaderT() gopurs_runtime.Value {
	once_withReaderT.Do(func() {
		withReaderT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})
	})
	return withReaderT
}

var runReaderT gopurs_runtime.Value
var once_runReaderT sync.Once
func Get_runReaderT() gopurs_runtime.Value {
	once_runReaderT.Do(func() {
		runReaderT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runReaderT
}

var newtypeReaderT gopurs_runtime.Value
var once_newtypeReaderT sync.Once
func Get_newtypeReaderT() gopurs_runtime.Value {
	once_newtypeReaderT.Do(func() {
		newtypeReaderT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeReaderT
}

var monadTransReaderT gopurs_runtime.Value
var once_monadTransReaderT sync.Once
func Get_monadTransReaderT() gopurs_runtime.Value {
	once_monadTransReaderT.Do(func() {
		monadTransReaderT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lift": gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
})
})})
	})
	return monadTransReaderT
}

var mapReaderT gopurs_runtime.Value
var once_mapReaderT sync.Once
func Get_mapReaderT() gopurs_runtime.Value {
	once_mapReaderT.Do(func() {
		mapReaderT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
})
	})
	return mapReaderT
}

var functorReaderT gopurs_runtime.Value
var once_functorReaderT sync.Once
func Get_functorReaderT() gopurs_runtime.Value {
	once_functorReaderT.Do(func() {
		functorReaderT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_1)
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(v_3, x_4))
})
})
})})
})
	})
	return functorReaderT
}

var distributiveReaderT gopurs_runtime.Value
var once_distributiveReaderT sync.Once
func Get_distributiveReaderT() gopurs_runtime.Value {
	once_distributiveReaderT.Do(func() {
		distributiveReaderT = gopurs_runtime.Func(func(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
distributiveReaderT:
for {
if false { continue distributiveReaderT }
var dictDistributive_0 = dictDistributive_0_loop
_ = dictDistributive_0
__local_var_1_0 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"distribute": gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_4_3 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["collect"], dictFunctor_3)
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(collect1_4_3, gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(r_7, e_6)
})), a_5)
})
})
}), "collect": gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_distributiveReaderT(), dictDistributive_0).PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_3)
__local_var_6_5 := gopurs_runtime.Apply(dictFunctor_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(__local_var_6_5, x_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
})})
}
}()
})
	})
	return distributiveReaderT
}

var applyReaderT gopurs_runtime.Value
var once_applyReaderT sync.Once
func Get_applyReaderT() gopurs_runtime.Value {
	once_applyReaderT.Do(func() {
		applyReaderT = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_3, r_5)), gopurs_runtime.Apply(v1_4, r_5))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
})})
})
	})
	return applyReaderT
}

var bindReaderT gopurs_runtime.Value
var once_bindReaderT sync.Once
func Get_bindReaderT() gopurs_runtime.Value {
	once_bindReaderT.Do(func() {
		bindReaderT = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_3)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
})})
applyReaderT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_4, r_6)), gopurs_runtime.Apply(v1_5, r_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v_4, r_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(k_5, a_7), r_6)
}))
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_3_2
})})
})
	})
	return bindReaderT
}

var semigroupReaderT gopurs_runtime.Value
var once_semigroupReaderT sync.Once
func Get_semigroupReaderT() gopurs_runtime.Value {
	once_semigroupReaderT.Do(func() {
		semigroupReaderT = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], __local_var_3_1)
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(a_4, r_7))), gopurs_runtime.Apply(b_5, r_7))
})
})
})})
})
})
	})
	return semigroupReaderT
}

var applicativeReaderT gopurs_runtime.Value
var once_applicativeReaderT sync.Once
func Get_applicativeReaderT() gopurs_runtime.Value {
	once_applicativeReaderT.Do(func() {
		applicativeReaderT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_3)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(v_5, x_6))
})
})
})})
applyReaderT1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_4, r_6)), gopurs_runtime.Apply(v1_5, r_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_2
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_5
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_4_4
})})
})
	})
	return applicativeReaderT
}

var monadReaderT gopurs_runtime.Value
var once_monadReaderT sync.Once
func Get_monadReaderT() gopurs_runtime.Value {
	once_monadReaderT.Do(func() {
		monadReaderT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_3_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_4)
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(v_6, x_7))
})
})
})})
applyReaderT1_5_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_5, r_7)), gopurs_runtime.Apply(v1_6, r_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_4
})})
applicativeReaderT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_6)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_7
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_5_6
})})
bindReaderT1_4_8 := gopurs_runtime.Apply(Get_bindReaderT(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_3_2
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bindReaderT1_4_8
})})
})
	})
	return monadReaderT
}

var monadAskReaderT gopurs_runtime.Value
var once_monadAskReaderT sync.Once
func Get_monadAskReaderT() gopurs_runtime.Value {
	once_monadAskReaderT.Do(func() {
		monadAskReaderT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), dictMonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
})
	})
	return monadAskReaderT
}

var monadReaderReaderT gopurs_runtime.Value
var once_monadReaderReaderT sync.Once
func Get_monadReaderReaderT() gopurs_runtime.Value {
	once_monadReaderReaderT.Do(func() {
		monadReaderReaderT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), dictMonad_0)
monadAskReaderT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": Get_withReaderT(), "MonadAsk0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskReaderT1_2_1
})})
})
	})
	return monadReaderReaderT
}

var monadContReaderT gopurs_runtime.Value
var once_monadContReaderT sync.Once
func Get_monadContReaderT() gopurs_runtime.Value {
	once_monadContReaderT.Do(func() {
		monadContReaderT = gopurs_runtime.Func(func(dictMonadCont_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadCont_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"callCC": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadCont_0.PtrVal.(map[string]gopurs_runtime.Value)["callCC"], gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(c_4, x_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_1
})
})), r_3)
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
})
	})
	return monadContReaderT
}

var monadEffectReader gopurs_runtime.Value
var once_monadEffectReader sync.Once
func Get_monadEffectReader() gopurs_runtime.Value {
	once_monadEffectReader.Do(func() {
		monadEffectReader = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftEffect": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
})
	})
	return monadEffectReader
}

var monadRecReaderT gopurs_runtime.Value
var once_monadRecReaderT sync.Once
func Get_monadRecReaderT() gopurs_runtime.Value {
	once_monadRecReaderT.Do(func() {
		monadRecReaderT = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
pure_3_2 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"]
monadReaderT1_4_3 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(k_5, a_prime_8), r_7)), pure_3_2)
})), a_6)
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_4_3
})})
})
	})
	return monadRecReaderT
}

var monadStateReaderT gopurs_runtime.Value
var once_monadStateReaderT sync.Once
func Get_monadStateReaderT() gopurs_runtime.Value {
	once_monadStateReaderT.Do(func() {
		monadStateReaderT = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
})
	})
	return monadStateReaderT
}

var monadTellReaderT gopurs_runtime.Value
var once_monadTellReaderT sync.Once
func Get_monadTellReaderT() gopurs_runtime.Value {
	once_monadTellReaderT.Do(func() {
		monadTellReaderT = gopurs_runtime.Func(func(dictMonadTell_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tell": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["tell"], x_3)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_2
})
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), "Monad1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
})})
})
	})
	return monadTellReaderT
}

var monadWriterReaderT gopurs_runtime.Value
var once_monadWriterReaderT sync.Once
func Get_monadWriterReaderT() gopurs_runtime.Value {
	once_monadWriterReaderT.Do(func() {
		monadWriterReaderT = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monoid0_1_0 := gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["Monoid0"], gopurs_runtime.Value{})
monadTellReaderT1_2_1 := gopurs_runtime.Apply(Get_monadTellReaderT(), gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadTell1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"listen": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["listen"], gopurs_runtime.Apply(v_3, x_4))
})
}), "pass": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["pass"], gopurs_runtime.Apply(v_3, x_4))
})
}), "Monoid0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_1_0
}), "MonadTell1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellReaderT1_2_1
})})
})
	})
	return monadWriterReaderT
}

var monadThrowReaderT gopurs_runtime.Value
var once_monadThrowReaderT sync.Once
func Get_monadThrowReaderT() gopurs_runtime.Value {
	once_monadThrowReaderT.Do(func() {
		monadThrowReaderT = gopurs_runtime.Func(func(dictMonadThrow_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadThrow_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"throwError": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadThrow_0.PtrVal.(map[string]gopurs_runtime.Value)["throwError"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
})
	})
	return monadThrowReaderT
}

var monadErrorReaderT gopurs_runtime.Value
var once_monadErrorReaderT sync.Once
func Get_monadErrorReaderT() gopurs_runtime.Value {
	once_monadErrorReaderT.Do(func() {
		monadErrorReaderT = gopurs_runtime.Func(func(dictMonadError_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowReaderT1_1_0 := gopurs_runtime.Apply(Get_monadThrowReaderT(), gopurs_runtime.Apply(dictMonadError_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadThrow0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"catchError": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadError_0.PtrVal.(map[string]gopurs_runtime.Value)["catchError"], gopurs_runtime.Apply(v_2, r_4)), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(h_3, e_5), r_4)
}))
})
})
}), "MonadThrow0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowReaderT1_1_0
})})
})
	})
	return monadErrorReaderT
}

var monadSTReaderT gopurs_runtime.Value
var once_monadSTReaderT sync.Once
func Get_monadSTReaderT() gopurs_runtime.Value {
	once_monadSTReaderT.Do(func() {
		monadSTReaderT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadST_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadST_0.PtrVal.(map[string]gopurs_runtime.Value)["liftST"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_1
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
})})
})
	})
	return monadSTReaderT
}

var monoidReaderT gopurs_runtime.Value
var once_monoidReaderT sync.Once
func Get_monoidReaderT() gopurs_runtime.Value {
	once_monoidReaderT.Do(func() {
		monoidReaderT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"]
semigroupReaderT2_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], __local_var_4_3)
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(a_5, r_8))), gopurs_runtime.Apply(b_6, r_8))
})
})
})})
__local_var_5_5 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_5
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupReaderT2_4_2
})})
})
})
	})
	return monoidReaderT
}

var altReaderT gopurs_runtime.Value
var once_altReaderT sync.Once
func Get_altReaderT() gopurs_runtime.Value {
	once_altReaderT.Do(func() {
		altReaderT = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_3, r_5)), gopurs_runtime.Apply(v1_4, r_5))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
})})
})
	})
	return altReaderT
}

var plusReaderT gopurs_runtime.Value
var once_plusReaderT sync.Once
func Get_plusReaderT() gopurs_runtime.Value {
	once_plusReaderT.Do(func() {
		plusReaderT = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_3)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
})})
altReaderT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_4, r_6)), gopurs_runtime.Apply(v1_5, r_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
})})
__local_var_4_5 := dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_5
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_3_2
})})
})
	})
	return plusReaderT
}

var alternativeReaderT gopurs_runtime.Value
var once_alternativeReaderT sync.Once
func Get_alternativeReaderT() gopurs_runtime.Value {
	once_alternativeReaderT.Do(func() {
		alternativeReaderT = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_3_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_4)
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(v_6, x_7))
})
})
})})
applyReaderT1_5_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_5, r_7)), gopurs_runtime.Apply(v1_6, r_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_4
})})
applicativeReaderT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_6)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_7
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_5_6
})})
__local_var_4_8 := gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
__local_var_5_9 := gopurs_runtime.Apply(__local_var_4_8.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_6_11 := gopurs_runtime.Apply(__local_var_5_9.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_7_12 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_13 := gopurs_runtime.Apply(__local_var_6_11.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_7)
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, gopurs_runtime.Apply(v_9, x_10))
})
})
})})
altReaderT1_8_14 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_9.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_8, r_10)), gopurs_runtime.Apply(v1_9, r_10))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_7_12
})})
__local_var_9_15 := __local_var_4_8.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
plusReaderT1_6_10 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_9_15
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_8_14
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_3_2
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusReaderT1_6_10
})})
})
	})
	return alternativeReaderT
}

var monadPlusReaderT gopurs_runtime.Value
var once_monadPlusReaderT sync.Once
func Get_monadPlusReaderT() gopurs_runtime.Value {
	once_monadPlusReaderT.Do(func() {
		monadPlusReaderT = gopurs_runtime.Func(func(dictMonadPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(dictMonadPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
alternativeReaderT1_2_1 := gopurs_runtime.Apply(Get_alternativeReaderT(), gopurs_runtime.Apply(dictMonadPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alternative1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeReaderT1_2_1
})})
})
	})
	return monadPlusReaderT
}


