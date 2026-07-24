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
		withExceptT = gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_3.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v2_3.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_3.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v2_3.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2)
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
		newtypeExceptT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeExceptT
}

var monadTransExceptT gopurs_runtime.Value
var once_monadTransExceptT sync.Once
func Get_monadTransExceptT() gopurs_runtime.Value {
	once_monadTransExceptT.Do(func() {
		monadTransExceptT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func2(func(dictMonad_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_1, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", a_2))
}))
}))
	})
	return monadTransExceptT
}

var mapExceptT gopurs_runtime.Value
var once_mapExceptT sync.Once
func Get_mapExceptT() gopurs_runtime.Value {
	once_mapExceptT.Do(func() {
		mapExceptT = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
	})
	return mapExceptT
}

var functorExceptT gopurs_runtime.Value
var once_functorExceptT sync.Once
func Get_functorExceptT() gopurs_runtime.Value {
	once_functorExceptT.Do(func() {
		functorExceptT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_2.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0]))
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
}))
})
	})
	return functorExceptT
}

var except gopurs_runtime.Value
var once_except sync.Once
func Get_except() gopurs_runtime.Value {
	once_except.Do(func() {
		except = gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_1)
})
	})
	return except
}

var monadExceptT gopurs_runtime.Value
var once_monadExceptT sync.Once
func Get_monadExceptT() gopurs_runtime.Value {
	once_monadExceptT.Do(func() {
		monadExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0_loop)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_0_loop)
}))
}()
})
	})
	return monadExceptT
}

var bindExceptT gopurs_runtime.Value
var once_bindExceptT sync.Once
func Get_bindExceptT() gopurs_runtime.Value {
	once_bindExceptT.Do(func() {
		bindExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}), "bind"), v_1, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_3.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(v2_3.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_3.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Apply(k_2, (*[1024]gopurs_runtime.Value)(v2_3.UnsafePtr)[0])
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
return gopurs_runtime.Apply(Get_applyExceptT(), dictMonad_0_loop)
}))
}()
})
	})
	return bindExceptT
}

var applyExceptT gopurs_runtime.Value
var once_applyExceptT sync.Once
func Get_applyExceptT() gopurs_runtime.Value {
	once_applyExceptT.Do(func() {
		applyExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorExceptT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_3.StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_3.UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(m_3.StrVal == "Right").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(m_3.UnsafePtr)[0]))
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
_ = functorExceptT1_2_1
__local_var_3_3 := gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_0_loop)
_ = __local_var_3_3
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "bind"), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "bind"), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0_loop), "pure"), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_2_1
}))
}()
})
	})
	return applyExceptT
}

var applicativeExceptT gopurs_runtime.Value
var once_applicativeExceptT sync.Once
func Get_applicativeExceptT() gopurs_runtime.Value {
	once_applicativeExceptT.Do(func() {
		applicativeExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", x_1))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyExceptT(), dictMonad_0_loop)
}))
}()
})
	})
	return applicativeExceptT
}

var semigroupExceptT gopurs_runtime.Value
var once_semigroupExceptT sync.Once
func Get_semigroupExceptT() gopurs_runtime.Value {
	once_semigroupExceptT.Do(func() {
		semigroupExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(Get_applyExceptT(), dictMonad_0_loop)
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
}))
})
}()
})
	})
	return semigroupExceptT
}

var monadAskExceptT gopurs_runtime.Value
var once_monadAskExceptT sync.Once
func Get_monadAskExceptT() gopurs_runtime.Value {
	once_monadAskExceptT.Do(func() {
		monadAskExceptT = gopurs_runtime.Func(func(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("ask", "Monad0", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0_loop, "ask"), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", a_3))
})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}))
}()
})
	})
	return monadAskExceptT
}

var monadReaderExceptT gopurs_runtime.Value
var once_monadReaderExceptT sync.Once
func Get_monadReaderExceptT() gopurs_runtime.Value {
	once_monadReaderExceptT.Do(func() {
		monadReaderExceptT = gopurs_runtime.Func(func(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
monadAskExceptT1_1_0 := gopurs_runtime.Apply(Get_monadAskExceptT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0_loop, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskExceptT1_1_0
return gopurs_runtime.RecordDict2("local", "MonadAsk0", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0_loop, "local"), f_2)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskExceptT1_1_0
}))
}()
})
	})
	return monadReaderExceptT
}

var monadContExceptT gopurs_runtime.Value
var once_monadContExceptT sync.Once
func Get_monadContExceptT() gopurs_runtime.Value {
	once_monadContExceptT.Do(func() {
		monadContExceptT = gopurs_runtime.Func(func(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_0
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), __local_var_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), __local_var_1_0)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("callCC", "Monad0", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0_loop, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_4, gopurs_runtime.Constructor1("Right", a_5))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}))
}()
})
	})
	return monadContExceptT
}

var monadEffectExceptT gopurs_runtime.Value
var once_monadEffectExceptT sync.Once
func Get_monadEffectExceptT() gopurs_runtime.Value {
	once_monadEffectExceptT.Do(func() {
		monadEffectExceptT = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), x_3), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", a_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}))
}()
})
	})
	return monadEffectExceptT
}

var monadRecExceptT gopurs_runtime.Value
var once_monadRecExceptT sync.Once
func Get_monadRecExceptT() gopurs_runtime.Value {
	once_monadRecExceptT.Do(func() {
		monadRecExceptT = gopurs_runtime.Func(func(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0_loop, "tailRecM"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(f_3, a_4), gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_prime_5.StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Done", gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_prime_5.UnsafePtr)[0]))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(m_prime_5.StrVal == "Right").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(m_prime_5.UnsafePtr)[0].StrVal == "Loop").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Loop", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(m_prime_5.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(m_prime_5.UnsafePtr)[0].StrVal == "Done").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Done", gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(m_prime_5.UnsafePtr)[0].UnsafePtr)[0]))
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
return monadExceptT1_2_1
}))
}()
})
	})
	return monadRecExceptT
}

var monadStateExceptT gopurs_runtime.Value
var once_monadStateExceptT sync.Once
func Get_monadStateExceptT() gopurs_runtime.Value {
	once_monadStateExceptT.Do(func() {
		monadStateExceptT = gopurs_runtime.Func(func(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("state", "Monad0", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0_loop, "state"), f_3), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", a_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}))
}()
})
	})
	return monadStateExceptT
}

var monadTellExceptT gopurs_runtime.Value
var once_monadTellExceptT sync.Once
func Get_monadTellExceptT() gopurs_runtime.Value {
	once_monadTellExceptT.Do(func() {
		monadTellExceptT = gopurs_runtime.Func(func(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0_loop, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadExceptT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad1_1_0)
}))
_ = monadExceptT1_3_2
return gopurs_runtime.RecordDict3("tell", "Semigroup0", "Monad1", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0_loop, "tell"), x_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", a_5))
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_2
}))
}()
})
	})
	return monadTellExceptT
}

var monadWriterExceptT gopurs_runtime.Value
var once_monadWriterExceptT sync.Once
func Get_monadWriterExceptT() gopurs_runtime.Value {
	once_monadWriterExceptT.Do(func() {
		monadWriterExceptT = gopurs_runtime.Func(func(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0_loop, "MonadTell1"), gopurs_runtime.Value{})
_ = MonadTell1_1_0
Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
Monoid0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0_loop, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_5_4
monadTellExceptT1_6_5 := gopurs_runtime.Apply(Get_monadTellExceptT(), MonadTell1_1_0)
_ = monadTellExceptT1_6_5
return gopurs_runtime.RecordDict4("listen", "pass", "Monoid0", "MonadTell1", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0_loop, "listen"), v_7), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0].StrVal == "Left").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0].StrVal == "Right").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0].UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1]))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), __t6)
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0_loop, "pass"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), v_7, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(a_8.StrVal == "Left").IntVal != 0 {
__t7 = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(a_8.UnsafePtr)[0]), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(a_8.StrVal == "Right").IntVal != 0 {
__t7 = gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(a_8.UnsafePtr)[0].UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(a_8.UnsafePtr)[0].UnsafePtr)[1])
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
return monadTellExceptT1_6_5
}))
}()
})
	})
	return monadWriterExceptT
}

var monadThrowExceptT gopurs_runtime.Value
var once_monadThrowExceptT sync.Once
func Get_monadThrowExceptT() gopurs_runtime.Value {
	once_monadThrowExceptT.Do(func() {
		monadThrowExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadExceptT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0_loop)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_0_loop)
}))
_ = monadExceptT1_1_0
return gopurs_runtime.RecordDict2("throwError", "Monad0", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Left", x_2))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
}))
}()
})
	})
	return monadThrowExceptT
}

var monadErrorExceptT gopurs_runtime.Value
var once_monadErrorExceptT sync.Once
func Get_monadErrorExceptT() gopurs_runtime.Value {
	once_monadErrorExceptT.Do(func() {
		monadErrorExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadThrowExceptT1_1_0 := gopurs_runtime.Apply(Get_monadThrowExceptT(), dictMonad_0_loop)
_ = monadThrowExceptT1_1_0
return gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(k_3, (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v2_4.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0]))
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
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowExceptT1_1_0
}))
}()
})
	})
	return monadErrorExceptT
}

var monadSTExceptT gopurs_runtime.Value
var once_monadSTExceptT sync.Once
func Get_monadSTExceptT() gopurs_runtime.Value {
	once_monadSTExceptT.Do(func() {
		monadSTExceptT = gopurs_runtime.Func(func(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), Monad0_1_0)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0_loop, "liftST"), x_3), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Right", a_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}))
}()
})
	})
	return monadSTExceptT
}

var monoidExceptT gopurs_runtime.Value
var once_monoidExceptT sync.Once
func Get_monoidExceptT() gopurs_runtime.Value {
	once_monoidExceptT.Do(func() {
		monoidExceptT = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
semigroupExceptT1_1_0 := gopurs_runtime.Apply(Get_semigroupExceptT(), dictMonad_0_loop)
_ = semigroupExceptT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupExceptT2_3_1 := gopurs_runtime.Apply(semigroupExceptT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupExceptT2_3_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_0_loop), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupExceptT2_3_1
}))
})
}()
})
	})
	return monoidExceptT
}

var altExceptT gopurs_runtime.Value
var once_altExceptT sync.Once
func Get_altExceptT() gopurs_runtime.Value {
	once_altExceptT.Do(func() {
		altExceptT = gopurs_runtime.Func2(Call_altExceptT)
	})
	return altExceptT
}

var plusExceptT gopurs_runtime.Value
var once_plusExceptT sync.Once
func Get_plusExceptT() gopurs_runtime.Value {
	once_plusExceptT.Do(func() {
		plusExceptT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = mempty_1_0
altExceptT1_2_1 := gopurs_runtime.Apply(Get_altExceptT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{}))
_ = altExceptT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
altExceptT2_4_2 := gopurs_runtime.Apply(altExceptT1_2_1, dictMonad_3)
_ = altExceptT2_4_2
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadThrowExceptT(), dictMonad_3), "throwError"), mempty_1_0), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altExceptT2_4_2
}))
})
}()
})
	})
	return plusExceptT
}

var alternativeExceptT gopurs_runtime.Value
var once_alternativeExceptT sync.Once
func Get_alternativeExceptT() gopurs_runtime.Value {
	once_alternativeExceptT.Do(func() {
		alternativeExceptT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
plusExceptT1_1_0 := gopurs_runtime.Apply(Get_plusExceptT(), dictMonoid_0_loop)
_ = plusExceptT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeExceptT1_3_1 := gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_2)
_ = applicativeExceptT1_3_1
plusExceptT2_4_2 := gopurs_runtime.Apply(plusExceptT1_1_0, dictMonad_2)
_ = plusExceptT2_4_2
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeExceptT1_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusExceptT2_4_2
}))
})
}()
})
	})
	return alternativeExceptT
}

var monadPlusExceptT gopurs_runtime.Value
var once_monadPlusExceptT sync.Once
func Get_monadPlusExceptT() gopurs_runtime.Value {
	once_monadPlusExceptT.Do(func() {
		monadPlusExceptT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
alternativeExceptT1_1_0 := gopurs_runtime.Apply(Get_alternativeExceptT(), dictMonoid_0_loop)
_ = alternativeExceptT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadExceptT1_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeExceptT(), dictMonad_2)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindExceptT(), dictMonad_2)
}))
_ = monadExceptT1_3_1
alternativeExceptT2_4_2 := gopurs_runtime.Apply(alternativeExceptT1_1_0, dictMonad_2)
_ = alternativeExceptT2_4_2
return gopurs_runtime.RecordDict2("Monad0", "Alternative1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeExceptT2_4_2
}))
})
}()
})
	})
	return monadPlusExceptT
}

func Call_altExceptT(dictSemigroup_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
_ = dictMonad_1
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1_loop, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1_loop, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorExceptT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_6.StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_6.UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(m_6.StrVal == "Right").IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(f_5, (*[1024]gopurs_runtime.Value)(m_6.UnsafePtr)[0]))
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
}))
_ = functorExceptT1_5_3
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_0, "bind"), v_6, gopurs_runtime.Func(func(rm_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(rm_8.StrVal == "Right").IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(rm_8.UnsafePtr)[0]))
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(rm_8.StrVal == "Left").IntVal != 0 {
__local_var_9_6 := (*[1024]gopurs_runtime.Value)(rm_8.UnsafePtr)[0]
_ = __local_var_9_6
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_0, "bind"), v1_7, gopurs_runtime.Func(func(rn_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(rn_10.StrVal == "Right").IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(rn_10.UnsafePtr)[0]))
goto end_branch_7
} else {

}
}
{
if gopurs_runtime.Bool(rn_10.StrVal == "Left").IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), __local_var_9_6, (*[1024]gopurs_runtime.Value)(rn_10.UnsafePtr)[0])))
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
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_3
}))
}


