package Control_Monad_Maybe_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var MaybeT gopurs_runtime.Value
var once_MaybeT sync.Once
func Get_MaybeT() gopurs_runtime.Value {
	once_MaybeT.Do(func() {
		MaybeT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return MaybeT
}

var runMaybeT gopurs_runtime.Value
var once_runMaybeT sync.Once
func Get_runMaybeT() gopurs_runtime.Value {
	once_runMaybeT.Do(func() {
		runMaybeT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runMaybeT
}

var newtypeMaybeT gopurs_runtime.Value
var once_newtypeMaybeT sync.Once
func Get_newtypeMaybeT() gopurs_runtime.Value {
	once_newtypeMaybeT.Do(func() {
		newtypeMaybeT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeMaybeT
}

var monadTransMaybeT gopurs_runtime.Value
var once_monadTransMaybeT sync.Once
func Get_monadTransMaybeT() gopurs_runtime.Value {
	once_monadTransMaybeT.Do(func() {
		monadTransMaybeT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func2(func(dictMonad_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), x_1, gopurs_runtime.Func(func(a_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_2))
}))
}))
	})
	return monadTransMaybeT
}

var mapMaybeT gopurs_runtime.Value
var once_mapMaybeT sync.Once
func Get_mapMaybeT() gopurs_runtime.Value {
	once_mapMaybeT.Do(func() {
		mapMaybeT = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
	})
	return mapMaybeT
}

var functorMaybeT gopurs_runtime.Value
var once_functorMaybeT sync.Once
func Get_functorMaybeT() gopurs_runtime.Value {
	once_functorMaybeT.Do(func() {
		functorMaybeT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3.StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_1, gopurs_runtime.ConstructorGet(v1_3, 0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), v_2)
}))
})
	})
	return functorMaybeT
}

var monadMaybeT gopurs_runtime.Value
var once_monadMaybeT sync.Once
func Get_monadMaybeT() gopurs_runtime.Value {
	once_monadMaybeT.Do(func() {
		monadMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), dictMonad_0)
}))
})
	})
	return monadMaybeT
}

var bindMaybeT gopurs_runtime.Value
var once_bindMaybeT sync.Once
func Get_bindMaybeT() gopurs_runtime.Value {
	once_bindMaybeT.Do(func() {
		bindMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), v_1, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3.StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor0("Nothing"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v1_3.StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(f_2, gopurs_runtime.ConstructorGet(v1_3, 0))
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
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyMaybeT(), dictMonad_0)
}))
})
	})
	return bindMaybeT
}

var applyMaybeT gopurs_runtime.Value
var once_applyMaybeT sync.Once
func Get_applyMaybeT() gopurs_runtime.Value {
	once_applyMaybeT.Do(func() {
		applyMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorMaybeT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4.StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_2, gopurs_runtime.ConstructorGet(v1_4, 0)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
}), v_3)
}))
_ = functorMaybeT1_2_1
__local_var_3_3 := gopurs_runtime.Apply(Get_bindMaybeT(), dictMonad_0)
_ = __local_var_3_3
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "bind"), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "bind"), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0), "pure"), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_2_1
}))
})
	})
	return applyMaybeT
}

var applicativeMaybeT gopurs_runtime.Value
var once_applicativeMaybeT sync.Once
func Get_applicativeMaybeT() gopurs_runtime.Value {
	once_applicativeMaybeT.Do(func() {
		applicativeMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", x_1))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyMaybeT(), dictMonad_0)
}))
})
	})
	return applicativeMaybeT
}

var semigroupMaybeT gopurs_runtime.Value
var once_semigroupMaybeT sync.Once
func Get_semigroupMaybeT() gopurs_runtime.Value {
	once_semigroupMaybeT.Do(func() {
		semigroupMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_applyMaybeT(), dictMonad_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
}))
})
})
	})
	return semigroupMaybeT
}

var monadAskMaybeT gopurs_runtime.Value
var once_monadAskMaybeT sync.Once
func Get_monadAskMaybeT() gopurs_runtime.Value {
	once_monadAskMaybeT.Do(func() {
		monadAskMaybeT = gopurs_runtime.Func(func(dictMonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("ask", "Monad0", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_3))
})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadAskMaybeT
}

var monadReaderMaybeT gopurs_runtime.Value
var once_monadReaderMaybeT sync.Once
func Get_monadReaderMaybeT() gopurs_runtime.Value {
	once_monadReaderMaybeT.Do(func() {
		monadReaderMaybeT = gopurs_runtime.Func(func(dictMonadReader_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskMaybeT1_1_0 := gopurs_runtime.Apply(Get_monadAskMaybeT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskMaybeT1_1_0
return gopurs_runtime.RecordDict2("local", "MonadAsk0", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskMaybeT1_1_0
}))
})
	})
	return monadReaderMaybeT
}

var monadContMaybeT gopurs_runtime.Value
var once_monadContMaybeT sync.Once
func Get_monadContMaybeT() gopurs_runtime.Value {
	once_monadContMaybeT.Do(func() {
		monadContMaybeT = gopurs_runtime.Func(func(dictMonadCont_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), __local_var_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), __local_var_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("callCC", "Monad0", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_4, gopurs_runtime.Constructor1("Just", a_5))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadContMaybeT
}

var monadEffectMaybe gopurs_runtime.Value
var once_monadEffectMaybe sync.Once
func Get_monadEffectMaybe() gopurs_runtime.Value {
	once_monadEffectMaybe.Do(func() {
		monadEffectMaybe = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_3), gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadEffectMaybe
}

var monadRecMaybeT gopurs_runtime.Value
var once_monadRecMaybeT sync.Once
func Get_monadRecMaybeT() gopurs_runtime.Value {
	once_monadRecMaybeT.Do(func() {
		monadRecMaybeT = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(f_3, a_4), gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_5.StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Done", gopurs_runtime.Constructor0("Nothing"))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_5.StrVal == "Just")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(m_prime_5, 0).StrVal == "Loop")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Loop", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(m_prime_5, 0), 0))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(m_prime_5, 0).StrVal == "Done")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Done", gopurs_runtime.Constructor1("Just", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(m_prime_5, 0), 0)))
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
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t2)
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadRecMaybeT
}

var monadStateMaybeT gopurs_runtime.Value
var once_monadStateMaybeT sync.Once
func Get_monadStateMaybeT() gopurs_runtime.Value {
	once_monadStateMaybeT.Do(func() {
		monadStateMaybeT = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("state", "Monad0", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3), gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadStateMaybeT
}

var monadTellMaybeT gopurs_runtime.Value
var once_monadTellMaybeT sync.Once
func Get_monadTellMaybeT() gopurs_runtime.Value {
	once_monadTellMaybeT.Do(func() {
		monadTellMaybeT = gopurs_runtime.Func(func(dictMonadTell_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadMaybeT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad1_1_0)
}))
_ = monadMaybeT1_3_2
return gopurs_runtime.RecordDict3("tell", "Semigroup0", "Monad1", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_4), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_5))
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_2
}))
})
	})
	return monadTellMaybeT
}

var monadWriterMaybeT gopurs_runtime.Value
var once_monadWriterMaybeT sync.Once
func Get_monadWriterMaybeT() gopurs_runtime.Value {
	once_monadWriterMaybeT.Do(func() {
		monadWriterMaybeT = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
_ = MonadTell1_1_0
Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
Monoid0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_5_4
monadTellMaybeT1_6_5 := gopurs_runtime.Apply(Get_monadTellMaybeT(), MonadTell1_1_0)
_ = monadTellMaybeT1_6_5
return gopurs_runtime.RecordDict4("listen", "pass", "Monoid0", "MonadTell1", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_7), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(v_8, 0).StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(v_8, 0), 0), gopurs_runtime.ConstructorGet(v_8, 1)))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), __t6)
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), v_7, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(a_8.StrVal == "Nothing")).IntVal != 0 {
__t7 = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(a_8.StrVal == "Just")).IntVal != 0 {
__t7 = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor1("Just", gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(a_8, 0), 0)), gopurs_runtime.ConstructorGet(gopurs_runtime.ConstructorGet(a_8, 0), 1))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), __t7)
})))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellMaybeT1_6_5
}))
})
	})
	return monadWriterMaybeT
}

var monadThrowMaybeT gopurs_runtime.Value
var once_monadThrowMaybeT sync.Once
func Get_monadThrowMaybeT() gopurs_runtime.Value {
	once_monadThrowMaybeT.Do(func() {
		monadThrowMaybeT = gopurs_runtime.Func(func(dictMonadThrow_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("throwError", "Monad0", gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3), gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadThrowMaybeT
}

var monadErrorMaybeT gopurs_runtime.Value
var once_monadErrorMaybeT sync.Once
func Get_monadErrorMaybeT() gopurs_runtime.Value {
	once_monadErrorMaybeT.Do(func() {
		monadErrorMaybeT = gopurs_runtime.Func(func(dictMonadError_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowMaybeT1_1_0 := gopurs_runtime.Apply(Get_monadThrowMaybeT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowMaybeT1_1_0
return gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_3, a_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowMaybeT1_1_0
}))
})
	})
	return monadErrorMaybeT
}

var monadSTMaybeT gopurs_runtime.Value
var once_monadSTMaybeT sync.Once
func Get_monadSTMaybeT() gopurs_runtime.Value {
	once_monadSTMaybeT.Do(func() {
		monadSTMaybeT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_3), gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", a_prime_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}))
})
	})
	return monadSTMaybeT
}

var monoidMaybeT gopurs_runtime.Value
var once_monoidMaybeT sync.Once
func Get_monoidMaybeT() gopurs_runtime.Value {
	once_monoidMaybeT.Do(func() {
		monoidMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMaybeT1_1_0 := gopurs_runtime.Apply(Get_semigroupMaybeT(), dictMonad_0)
_ = semigroupMaybeT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMaybeT2_3_1 := gopurs_runtime.Apply(semigroupMaybeT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupMaybeT2_3_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybeT2_3_1
}))
})
})
	})
	return monoidMaybeT
}

var altMaybeT gopurs_runtime.Value
var once_altMaybeT sync.Once
func Get_altMaybeT() gopurs_runtime.Value {
	once_altMaybeT.Do(func() {
		altMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorMaybeT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_5.StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_3, gopurs_runtime.ConstructorGet(v1_5, 0)))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), v_4)
}))
_ = functorMaybeT1_3_2
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_6.StrVal == "Nothing")).IntVal != 0 {
__t4 = v1_5
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), m_6)
}
end_branch_4:
return __t4
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_2
}))
})
	})
	return altMaybeT
}

var plusMaybeT gopurs_runtime.Value
var once_plusMaybeT sync.Once
func Get_plusMaybeT() gopurs_runtime.Value {
	once_plusMaybeT.Do(func() {
		plusMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorMaybeT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_5.StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_3, gopurs_runtime.ConstructorGet(v1_5, 0)))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_4:
return __t4
}), v_4)
}))
_ = functorMaybeT1_3_3
altMaybeT1_3_2 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_6.StrVal == "Nothing")).IntVal != 0 {
__t5 = v1_5
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), m_6)
}
end_branch_5:
return __t5
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_3
}))
_ = altMaybeT1_3_2
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor0("Nothing")), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altMaybeT1_3_2
}))
})
	})
	return plusMaybeT
}

var alternativeMaybeT gopurs_runtime.Value
var once_alternativeMaybeT sync.Once
func Get_alternativeMaybeT() gopurs_runtime.Value {
	once_alternativeMaybeT.Do(func() {
		alternativeMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeMaybeT1_1_0 := gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0)
_ = applicativeMaybeT1_1_0
plusMaybeT1_2_1 := gopurs_runtime.Apply(Get_plusMaybeT(), dictMonad_0)
_ = plusMaybeT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeMaybeT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusMaybeT1_2_1
}))
})
	})
	return alternativeMaybeT
}

var monadPlusMaybeT gopurs_runtime.Value
var once_monadPlusMaybeT sync.Once
func Get_monadPlusMaybeT() gopurs_runtime.Value {
	once_monadPlusMaybeT.Do(func() {
		monadPlusMaybeT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadMaybeT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), dictMonad_0)
}))
_ = monadMaybeT1_1_0
alternativeMaybeT1_2_1 := gopurs_runtime.Apply(Get_alternativeMaybeT(), dictMonad_0)
_ = alternativeMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "Alternative1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeMaybeT1_2_1
}))
})
	})
	return monadPlusMaybeT
}


