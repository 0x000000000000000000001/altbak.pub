package Control_Monad_Except_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var ExceptT gopurs_runtime.Value
var once_ExceptT sync.Once
func Get_ExceptT() gopurs_runtime.Value {
	once_ExceptT.Do(func() {
		ExceptT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return ExceptT
}

var withExceptT gopurs_runtime.Value
var once_withExceptT sync.Once
func Get_withExceptT() gopurs_runtime.Value {
	once_withExceptT.Do(func() {
		withExceptT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(f_1, v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})), v_2)
})
})
})
	})
	return withExceptT
}

var runExceptT gopurs_runtime.Value
var once_runExceptT sync.Once
func Get_runExceptT() gopurs_runtime.Value {
	once_runExceptT.Do(func() {
		runExceptT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runExceptT
}

var newtypeExceptT gopurs_runtime.Value
var once_newtypeExceptT sync.Once
func Get_newtypeExceptT() gopurs_runtime.Value {
	once_newtypeExceptT.Do(func() {
		newtypeExceptT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeExceptT
}

var monadTransExceptT gopurs_runtime.Value
var once_monadTransExceptT sync.Once
func Get_monadTransExceptT() gopurs_runtime.Value {
	once_monadTransExceptT.Do(func() {
		monadTransExceptT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lift": gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], m_1), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_2}))
}))
})
})})
	})
	return monadTransExceptT
}

var mapExceptT gopurs_runtime.Value
var once_mapExceptT sync.Once
func Get_mapExceptT() gopurs_runtime.Value {
	once_mapExceptT.Do(func() {
		mapExceptT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
})
	})
	return mapExceptT
}

var functorExceptT gopurs_runtime.Value
var once_functorExceptT sync.Once
func Get_functorExceptT() gopurs_runtime.Value {
	once_functorExceptT.Do(func() {
		functorExceptT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(f_1, m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})})
})
	})
	return functorExceptT
}

var except gopurs_runtime.Value
var once_except sync.Once
func Get_except() gopurs_runtime.Value {
	once_except.Do(func() {
		except = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_1)
})
})
	})
	return except
}

var monadExceptT gopurs_runtime.Value
var once_monadExceptT sync.Once
func Get_monadExceptT() gopurs_runtime.Value {
	once_monadExceptT.Do(func() {
		monadExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_0)
})})
})
	})
	return monadExceptT
}

var bindExceptT gopurs_runtime.Value
var once_bindExceptT sync.Once
func Get_bindExceptT() gopurs_runtime.Value {
	once_bindExceptT.Do(func() {
		bindExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_1), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(k_2, v2_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyExceptT(), dictMonad_0)
})})
})
	})
	return bindExceptT
}

var applyExceptT gopurs_runtime.Value
var once_applyExceptT sync.Once
func Get_applyExceptT() gopurs_runtime.Value {
	once_applyExceptT.Do(func() {
		applyExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(m_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(f_2, m_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
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
})})
__local_var_3_3 := gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["bind"], f_4), gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_5), gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_2_1
})})
})
	})
	return applyExceptT
}

var applicativeExceptT gopurs_runtime.Value
var once_applicativeExceptT sync.Once
func Get_applicativeExceptT() gopurs_runtime.Value {
	once_applicativeExceptT.Do(func() {
		applicativeExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_1}))
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyExceptT(), dictMonad_0)
})})
})
	})
	return applicativeExceptT
}

var semigroupExceptT gopurs_runtime.Value
var once_semigroupExceptT sync.Once
func Get_semigroupExceptT() gopurs_runtime.Value {
	once_semigroupExceptT.Do(func() {
		semigroupExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_applyExceptT(), dictMonad_0)
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], __local_var_3_1), a_4)), b_5)
})
})})
})
})
	})
	return semigroupExceptT
}

var monadAskExceptT gopurs_runtime.Value
var once_monadAskExceptT sync.Once
func Get_monadAskExceptT() gopurs_runtime.Value {
	once_monadAskExceptT.Do(func() {
		monadAskExceptT = gopurs_runtime.Func(func(dictMonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], dictMonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"]), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_3}))
})), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
})})
})
	})
	return monadAskExceptT
}

var monadReaderExceptT gopurs_runtime.Value
var once_monadReaderExceptT sync.Once
func Get_monadReaderExceptT() gopurs_runtime.Value {
	once_monadReaderExceptT.Do(func() {
		monadReaderExceptT = gopurs_runtime.Func(func(dictMonadReader_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskExceptT1_1_0 := gopurs_runtime.Apply(Get_monadAskExceptT(), gopurs_runtime.Apply(dictMonadReader_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadAsk0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadReader_0.PtrVal.(map[string]gopurs_runtime.Value)["local"], f_2)
}), "MonadAsk0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskExceptT1_1_0
})})
})
	})
	return monadReaderExceptT
}

var monadContExceptT gopurs_runtime.Value
var once_monadContExceptT sync.Once
func Get_monadContExceptT() gopurs_runtime.Value {
	once_monadContExceptT.Do(func() {
		monadContExceptT = gopurs_runtime.Func(func(dictMonadCont_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonadCont_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), __local_var_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), __local_var_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"callCC": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadCont_0.PtrVal.(map[string]gopurs_runtime.Value)["callCC"], gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_4, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_5}))
}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
})})
})
	})
	return monadContExceptT
}

var monadEffectExceptT gopurs_runtime.Value
var once_monadEffectExceptT sync.Once
func Get_monadEffectExceptT() gopurs_runtime.Value {
	once_monadEffectExceptT.Do(func() {
		monadEffectExceptT = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftEffect": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], x_3)), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_4}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
})})
})
	})
	return monadEffectExceptT
}

var monadRecExceptT gopurs_runtime.Value
var once_monadRecExceptT sync.Once
func Get_monadRecExceptT() gopurs_runtime.Value {
	once_monadRecExceptT.Do(func() {
		monadRecExceptT = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(f_3, a_4)), gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": m_prime_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t2)
}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
})})
})
	})
	return monadRecExceptT
}

var monadStateExceptT gopurs_runtime.Value
var once_monadStateExceptT sync.Once
func Get_monadStateExceptT() gopurs_runtime.Value {
	once_monadStateExceptT.Do(func() {
		monadStateExceptT = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], f_3)), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_4}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
})})
})
	})
	return monadStateExceptT
}

var monadTellExceptT gopurs_runtime.Value
var once_monadTellExceptT sync.Once
func Get_monadTellExceptT() gopurs_runtime.Value {
	once_monadTellExceptT.Do(func() {
		monadTellExceptT = gopurs_runtime.Func(func(dictMonadTell_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad1_1_0 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{})
Semigroup0_2_1 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
monadExceptT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad1_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad1_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tell": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["tell"], x_4)), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_5}))
}))
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), "Monad1": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_2
})})
})
	})
	return monadTellExceptT
}

var monadWriterExceptT gopurs_runtime.Value
var once_monadWriterExceptT sync.Once
func Get_monadWriterExceptT() gopurs_runtime.Value {
	once_monadWriterExceptT.Do(func() {
		monadWriterExceptT = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadTell1_1_0 := gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadTell1"], gopurs_runtime.Value{})
Monad1_2_1 := gopurs_runtime.Apply(MonadTell1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(Monad1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_4_3 := gopurs_runtime.Apply(Monad1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
Monoid0_5_4 := gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["Monoid0"], gopurs_runtime.Value{})
monadTellExceptT1_6_5 := gopurs_runtime.Apply(Get_monadTellExceptT(), MonadTell1_1_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"listen": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["listen"], v_7)), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t6)
}))
}), "pass": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["pass"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(a_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": a_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]})
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(a_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": a_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t7)
})))
}), "Monoid0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), "MonadTell1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellExceptT1_6_5
})})
})
	})
	return monadWriterExceptT
}

var monadThrowExceptT gopurs_runtime.Value
var once_monadThrowExceptT sync.Once
func Get_monadThrowExceptT() gopurs_runtime.Value {
	once_monadThrowExceptT.Do(func() {
		monadThrowExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadExceptT1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"throwError": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_2}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
})})
})
	})
	return monadThrowExceptT
}

var monadErrorExceptT gopurs_runtime.Value
var once_monadErrorExceptT sync.Once
func Get_monadErrorExceptT() gopurs_runtime.Value {
	once_monadErrorExceptT.Do(func() {
		monadErrorExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowExceptT1_1_0 := gopurs_runtime.Apply(Get_monadThrowExceptT(), dictMonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"catchError": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_2), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(k_3, v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
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
})
}), "MonadThrow0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowExceptT1_1_0
})})
})
	})
	return monadErrorExceptT
}

var monadSTExceptT gopurs_runtime.Value
var once_monadSTExceptT sync.Once
func Get_monadSTExceptT() gopurs_runtime.Value {
	once_monadSTExceptT.Do(func() {
		monadSTExceptT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadST_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadExceptT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadST_0.PtrVal.(map[string]gopurs_runtime.Value)["liftST"], x_3)), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": a_4}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
})})
})
	})
	return monadSTExceptT
}

var monoidExceptT gopurs_runtime.Value
var once_monoidExceptT sync.Once
func Get_monoidExceptT() gopurs_runtime.Value {
	once_monoidExceptT.Do(func() {
		monoidExceptT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupExceptT1_1_0 := gopurs_runtime.Apply(Get_semigroupExceptT(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupExceptT2_3_1 := gopurs_runtime.Apply(semigroupExceptT1_1_0, gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0).PtrVal.(map[string]gopurs_runtime.Value)["pure"], dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupExceptT2_3_1
})})
})
})
	})
	return monoidExceptT
}

var altExceptT gopurs_runtime.Value
var once_altExceptT sync.Once
func Get_altExceptT() gopurs_runtime.Value {
	once_altExceptT.Do(func() {
		altExceptT = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_0 := gopurs_runtime.Apply(dictMonad_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_3_1 := gopurs_runtime.Apply(dictMonad_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorExceptT1_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(m_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(f_5, m_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_6), gopurs_runtime.Func(func(rm_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(rm_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": rm_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(rm_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__local_var_9_6 := rm_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v1_7), gopurs_runtime.Func(func(rn_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(rn_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": rn_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(rn_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_9_6), rn_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"])}))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_3
})})
})
})
	})
	return altExceptT
}

var plusExceptT gopurs_runtime.Value
var once_plusExceptT sync.Once
func Get_plusExceptT() gopurs_runtime.Value {
	once_plusExceptT.Do(func() {
		plusExceptT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
altExceptT1_2_1 := gopurs_runtime.Apply(Get_altExceptT(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
altExceptT2_4_2 := gopurs_runtime.Apply(altExceptT1_2_1, dictMonad_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_monadThrowExceptT(), dictMonad_3).PtrVal.(map[string]gopurs_runtime.Value)["throwError"], mempty_1_0), "Alt0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altExceptT2_4_2
})})
})
})
	})
	return plusExceptT
}

var alternativeExceptT gopurs_runtime.Value
var once_alternativeExceptT sync.Once
func Get_alternativeExceptT() gopurs_runtime.Value {
	once_alternativeExceptT.Do(func() {
		alternativeExceptT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
plusExceptT1_1_0 := gopurs_runtime.Apply(Get_plusExceptT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeExceptT1_3_1 := gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_2)
plusExceptT2_4_2 := gopurs_runtime.Apply(plusExceptT1_1_0, dictMonad_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeExceptT1_3_1
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusExceptT2_4_2
})})
})
})
	})
	return alternativeExceptT
}

var monadPlusExceptT gopurs_runtime.Value
var once_monadPlusExceptT sync.Once
func Get_monadPlusExceptT() gopurs_runtime.Value {
	once_monadPlusExceptT.Do(func() {
		monadPlusExceptT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
alternativeExceptT1_1_0 := gopurs_runtime.Apply(Get_alternativeExceptT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadExceptT1_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_2)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_2)
})})
alternativeExceptT2_4_2 := gopurs_runtime.Apply(alternativeExceptT1_1_0, dictMonad_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_1
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeExceptT2_4_2
})})
})
})
	})
	return monadPlusExceptT
}


