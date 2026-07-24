package Control_Monad_Writer_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var WriterT gopurs_runtime.Value
var once_WriterT sync.Once
func Get_WriterT() gopurs_runtime.Value {
	once_WriterT.Do(func() {
		WriterT = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return WriterT
}

var runWriterT gopurs_runtime.Value
var once_runWriterT sync.Once
func Get_runWriterT() gopurs_runtime.Value {
	once_runWriterT.Do(func() {
		runWriterT = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}()
})
	})
	return runWriterT
}

var newtypeWriterT gopurs_runtime.Value
var once_newtypeWriterT sync.Once
func Get_newtypeWriterT() gopurs_runtime.Value {
	once_newtypeWriterT.Do(func() {
		newtypeWriterT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeWriterT
}

var monadTransWriterT gopurs_runtime.Value
var once_monadTransWriterT sync.Once
func Get_monadTransWriterT() gopurs_runtime.Value {
	once_monadTransWriterT.Do(func() {
		monadTransWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func2(func(dictMonad_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor2("Tuple", a_4, mempty_1_0))
}))
}))
}()
})
	})
	return monadTransWriterT
}

var mapWriterT gopurs_runtime.Value
var once_mapWriterT sync.Once
func Get_mapWriterT() gopurs_runtime.Value {
	once_mapWriterT.Do(func() {
		mapWriterT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWriterT(f_0_box, v_1_box)
})
	})
	return mapWriterT
}

var functorWriterT gopurs_runtime.Value
var once_functorWriterT sync.Once
func Get_functorWriterT() gopurs_runtime.Value {
	once_functorWriterT.Do(func() {
		functorWriterT = gopurs_runtime.Func(func(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
}))
}))
}()
})
	})
	return functorWriterT
}

var execWriterT gopurs_runtime.Value
var once_execWriterT sync.Once
func Get_execWriterT() gopurs_runtime.Value {
	once_execWriterT.Do(func() {
		execWriterT = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execWriterT(dictFunctor_0_box, v_1_box)
})
	})
	return execWriterT
}

var applyWriterT gopurs_runtime.Value
var once_applyWriterT sync.Once
func Get_applyWriterT() gopurs_runtime.Value {
	once_applyWriterT.Do(func() {
		applyWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyWriterT(dictSemigroup_0_box, dictApply_1_box)
})
	})
	return applyWriterT
}

var bindWriterT gopurs_runtime.Value
var once_bindWriterT sync.Once
func Get_bindWriterT() gopurs_runtime.Value {
	once_bindWriterT.Do(func() {
		bindWriterT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictBind_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindWriterT(dictSemigroup_0_box, dictBind_1_box)
})
	})
	return bindWriterT
}

var semigroupWriterT gopurs_runtime.Value
var once_semigroupWriterT sync.Once
func Get_semigroupWriterT() gopurs_runtime.Value {
	once_semigroupWriterT.Do(func() {
		semigroupWriterT = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupWriterT(dictApply_0_box, dictSemigroup_1_box)
})
	})
	return semigroupWriterT
}

var applicativeWriterT gopurs_runtime.Value
var once_applicativeWriterT sync.Once
func Get_applicativeWriterT() gopurs_runtime.Value {
	once_applicativeWriterT.Do(func() {
		applicativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
Functor0_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_5_3
functorWriterT1_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_5_3, "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_6_5
applyWriterT2_6_4 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_5_3, "map"), gopurs_runtime.Func2(func(v3_9 gopurs_runtime.Value, v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v3_9.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v4_10.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*[1024]gopurs_runtime.Value)(v3_9.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v4_10.UnsafePtr)[1]))
}), v_7), v1_8)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_5
}))
_ = applyWriterT2_6_4
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Constructor2("Tuple", a_7, mempty_1_0))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_4
}))
})
}()
})
	})
	return applicativeWriterT
}

var monadWriterT gopurs_runtime.Value
var once_monadWriterT sync.Once
func Get_monadWriterT() gopurs_runtime.Value {
	once_monadWriterT.Do(func() {
		monadWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applicativeWriterT1_1_0 := gopurs_runtime.Apply(Get_applicativeWriterT(), dictMonoid_0)
_ = applicativeWriterT1_1_0
bindWriterT1_2_1 := gopurs_runtime.Apply(Get_bindWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindWriterT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT2_4_2 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeWriterT2_4_2
bindWriterT2_5_3 := gopurs_runtime.Apply(bindWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = bindWriterT2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_5_3
}))
})
}()
})
	})
	return monadWriterT
}

var monadAskWriterT gopurs_runtime.Value
var once_monadAskWriterT sync.Once
func Get_monadAskWriterT() gopurs_runtime.Value {
	once_monadAskWriterT.Do(func() {
		monadAskWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadAsk_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_5_3
return gopurs_runtime.RecordDict2("ask", "Monad0", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_3, "ask"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor2("Tuple", a_6, mempty_1_0))
})), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}))
})
}()
})
	})
	return monadAskWriterT
}

var monadReaderWriterT gopurs_runtime.Value
var once_monadReaderWriterT sync.Once
func Get_monadReaderWriterT() gopurs_runtime.Value {
	once_monadReaderWriterT.Do(func() {
		monadReaderWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadAskWriterT1_1_0 := gopurs_runtime.Apply(Get_monadAskWriterT(), dictMonoid_0)
_ = monadAskWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadReader_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskWriterT2_3_1 := gopurs_runtime.Apply(monadAskWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskWriterT2_3_1
return gopurs_runtime.RecordDict2("local", "MonadAsk0", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_2, "local"), f_4)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskWriterT2_3_1
}))
})
}()
})
	})
	return monadReaderWriterT
}

var monadContWriterT gopurs_runtime.Value
var once_monadContWriterT sync.Once
func Get_monadContWriterT() gopurs_runtime.Value {
	once_monadContWriterT.Do(func() {
		monadContWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadCont_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_4_2
return gopurs_runtime.RecordDict2("callCC", "Monad0", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_3, "callCC"), gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_6, gopurs_runtime.Constructor2("Tuple", a_7, mempty_1_0))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}))
})
}()
})
	})
	return monadContWriterT
}

var monadEffectWriter gopurs_runtime.Value
var once_monadEffectWriter sync.Once
func Get_monadEffectWriter() gopurs_runtime.Value {
	once_monadEffectWriter.Do(func() {
		monadEffectWriter = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadEffect_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_5_3
return gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_3, "liftEffect"), x_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor2("Tuple", a_7, mempty_1_0))
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}))
})
}()
})
	})
	return monadEffectWriter
}

var monadRecWriterT gopurs_runtime.Value
var once_monadRecWriterT sync.Once
func Get_monadRecWriterT() gopurs_runtime.Value {
	once_monadRecWriterT.Do(func() {
		monadRecWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
monadWriterT1_3_2 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_3_2
return gopurs_runtime.Func(func(dictMonadRec_4 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_4, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_5_3
monadWriterT2_6_4 := gopurs_runtime.Apply(monadWriterT1_3_2, Monad0_5_3)
_ = monadWriterT2_6_4
return gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_4, "tailRecM"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_5 := (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[1]
_ = __local_var_10_5
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(f_7, (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[0]), gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v2_11.UnsafePtr)[0].StrVal == "Loop").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Loop", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_11.UnsafePtr)[0].UnsafePtr)[0], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), __local_var_10_5, (*[1024]gopurs_runtime.Value)(v2_11.UnsafePtr)[1])))
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v2_11.UnsafePtr)[0].StrVal == "Done").IntVal != 0 {
__t6 = gopurs_runtime.Constructor1("Done", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v2_11.UnsafePtr)[0].UnsafePtr)[0], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), __local_var_10_5, (*[1024]gopurs_runtime.Value)(v2_11.UnsafePtr)[1])))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_5_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t6)
}))
}), gopurs_runtime.Constructor2("Tuple", a_8, mempty_2_1))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_6_4
}))
})
}()
})
	})
	return monadRecWriterT
}

var monadStateWriterT gopurs_runtime.Value
var once_monadStateWriterT sync.Once
func Get_monadStateWriterT() gopurs_runtime.Value {
	once_monadStateWriterT.Do(func() {
		monadStateWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadState_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_5_3
return gopurs_runtime.RecordDict2("state", "Monad0", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_3, "state"), f_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor2("Tuple", a_7, mempty_1_0))
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}))
})
}()
})
	})
	return monadStateWriterT
}

var monadTellWriterT gopurs_runtime.Value
var once_monadTellWriterT sync.Once
func Get_monadTellWriterT() gopurs_runtime.Value {
	once_monadTellWriterT.Do(func() {
		monadTellWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
Semigroup0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_2_1, dictMonad_3)
_ = monadWriterT2_4_2
__local_var_5_3 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), pkg_Data_Unit.Get_unit())
_ = __local_var_5_3
return gopurs_runtime.RecordDict3("tell", "Semigroup0", "Monad1", gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(__local_var_5_3, x_6))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}))
})
}()
})
	})
	return monadTellWriterT
}

var monadWriterWriterT gopurs_runtime.Value
var once_monadWriterWriterT sync.Once
func Get_monadWriterWriterT() gopurs_runtime.Value {
	once_monadWriterWriterT.Do(func() {
		monadWriterWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadTellWriterT1_1_0 := gopurs_runtime.Apply(Get_monadTellWriterT(), dictMonoid_0)
_ = monadTellWriterT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_2
monadTellWriterT2_5_3 := gopurs_runtime.Apply(monadTellWriterT1_1_0, dictMonad_2)
_ = monadTellWriterT2_5_3
return gopurs_runtime.RecordDict4("listen", "pass", "Monoid0", "MonadTell1", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "bind"), v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "pure"), gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[1]), (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[1]))
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "bind"), v_6, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "pure"), gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0].UnsafePtr)[0], gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0].UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[1])))
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellWriterT2_5_3
}))
})
}()
})
	})
	return monadWriterWriterT
}

var monadThrowWriterT gopurs_runtime.Value
var once_monadThrowWriterT sync.Once
func Get_monadThrowWriterT() gopurs_runtime.Value {
	once_monadThrowWriterT.Do(func() {
		monadThrowWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadThrow_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_5_3
return gopurs_runtime.RecordDict2("throwError", "Monad0", gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_3, "throwError"), e_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor2("Tuple", a_7, mempty_1_0))
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}))
})
}()
})
	})
	return monadThrowWriterT
}

var monadErrorWriterT gopurs_runtime.Value
var once_monadErrorWriterT sync.Once
func Get_monadErrorWriterT() gopurs_runtime.Value {
	once_monadErrorWriterT.Do(func() {
		monadErrorWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadThrowWriterT1_1_0 := gopurs_runtime.Apply(Get_monadThrowWriterT(), dictMonoid_0)
_ = monadThrowWriterT1_1_0
return gopurs_runtime.Func(func(dictMonadError_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowWriterT2_3_1 := gopurs_runtime.Apply(monadThrowWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_2, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowWriterT2_3_1
return gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_2, "catchError"), v_4, gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_5, e_6)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowWriterT2_3_1
}))
})
}()
})
	})
	return monadErrorWriterT
}

var monadSTWriterT gopurs_runtime.Value
var once_monadSTWriterT sync.Once
func Get_monadSTWriterT() gopurs_runtime.Value {
	once_monadSTWriterT.Do(func() {
		monadSTWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadST_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_3, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_2
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
_ = monadWriterT2_5_3
return gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_3, "liftST"), x_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor2("Tuple", a_7, mempty_1_0))
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
}))
})
}()
})
	})
	return monadSTWriterT
}

var monoidWriterT gopurs_runtime.Value
var once_monoidWriterT sync.Once
func Get_monoidWriterT() gopurs_runtime.Value {
	once_monoidWriterT.Do(func() {
		monoidWriterT = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_1
Functor0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_4_2
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_5, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_6_3
semigroupWriterT3_7_4 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_2, "map"), gopurs_runtime.Func2(func(v3_9 gopurs_runtime.Value, v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v3_9.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v4_10.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "append"), (*[1024]gopurs_runtime.Value)(v3_9.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v4_10.UnsafePtr)[1]))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_2, "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_3, "append"), (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[1])
}), a_7)), b_8)
}))
_ = semigroupWriterT3_7_4
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_applicativeWriterT(), dictMonoid_2, dictApplicative_0), "pure"), gopurs_runtime.RecordGet(dictMonoid1_5, "mempty")), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupWriterT3_7_4
}))
})
})
}()
})
	})
	return monoidWriterT
}

var altWriterT gopurs_runtime.Value
var once_altWriterT sync.Once
func Get_altWriterT() gopurs_runtime.Value {
	once_altWriterT.Do(func() {
		altWriterT = gopurs_runtime.Func(func(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorWriterT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_2_1
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_3, v1_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_2_1
}))
}()
})
	})
	return altWriterT
}

var plusWriterT gopurs_runtime.Value
var once_plusWriterT sync.Once
func Get_plusWriterT() gopurs_runtime.Value {
	once_plusWriterT.Do(func() {
		plusWriterT = gopurs_runtime.Func(func(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorWriterT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_3_3
altWriterT1_3_2 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "alt"), v_4, v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_3
}))
_ = altWriterT1_3_2
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.RecordGet(dictPlus_0, "empty"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_3_2
}))
}()
})
	})
	return plusWriterT
}

var alternativeWriterT gopurs_runtime.Value
var once_alternativeWriterT sync.Once
func Get_alternativeWriterT() gopurs_runtime.Value {
	once_alternativeWriterT.Do(func() {
		alternativeWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applicativeWriterT1_1_0 := gopurs_runtime.Apply(Get_applicativeWriterT(), dictMonoid_0)
_ = applicativeWriterT1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT2_3_1 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeWriterT2_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_3
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
functorWriterT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_7, (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_7_6
altWriterT1_8_7 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_3, "alt"), v_8, v1_9)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_7_6
}))
_ = altWriterT1_8_7
plusWriterT1_6_4 := gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.RecordGet(__local_var_4_2, "empty"), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_8_7
}))
_ = plusWriterT1_6_4
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_3_1
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusWriterT1_6_4
}))
})
}()
})
	})
	return alternativeWriterT
}

var monadPlusWriterT gopurs_runtime.Value
var once_monadPlusWriterT sync.Once
func Get_monadPlusWriterT() gopurs_runtime.Value {
	once_monadPlusWriterT.Do(func() {
		monadPlusWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
monadWriterT1_1_0 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
_ = monadWriterT1_1_0
alternativeWriterT1_2_1 := gopurs_runtime.Apply(Get_alternativeWriterT(), dictMonoid_0)
_ = alternativeWriterT1_2_1
return gopurs_runtime.Func(func(dictMonadPlus_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Monad0"), gopurs_runtime.Value{}))
_ = monadWriterT2_4_2
alternativeWriterT2_5_3 := gopurs_runtime.Apply(alternativeWriterT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_3, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeWriterT2_5_3
return gopurs_runtime.RecordDict2("Monad0", "Alternative1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeWriterT2_5_3
}))
})
}()
})
	})
	return monadPlusWriterT
}

func Call_mapWriterT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_execWriterT(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Data_Tuple.Get_snd(), v_1)
}

func Call_applyWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictApply_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictApply_1 gopurs_runtime.Value = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_1, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_0
functorWriterT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_2_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_3_1
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_0, "map"), gopurs_runtime.Func2(func(v3_6 gopurs_runtime.Value, v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v3_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v4_7.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v3_6.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v4_7.UnsafePtr)[1]))
}), v_4), v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_1
}))
}

func Call_bindWriterT(dictSemigroup_0_loop gopurs_runtime.Value, dictBind_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictBind_1 gopurs_runtime.Value = dictBind_1_loop
_ = dictBind_1
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_0
Functor0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_3_1
functorWriterT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_3_1, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_4_2
applyWriterT2_5_3 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_3_1, "map"), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v3_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v4_8.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v3_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v4_8.UnsafePtr)[1]))
}), v_5), v1_6)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_4_2
}))
_ = applyWriterT2_5_3
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_1, "bind"), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[1]
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v3_10.UnsafePtr)[0], gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_9_4, (*[1024]gopurs_runtime.Value)(v3_10.UnsafePtr)[1]))
}), gopurs_runtime.Apply(k_7, (*[1024]gopurs_runtime.Value)(v1_8.UnsafePtr)[0]))
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_3
}))
}

func Call_semigroupWriterT(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
Functor0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_0
return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_0, "map"), gopurs_runtime.Func2(func(v3_6 gopurs_runtime.Value, v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v3_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v4_7.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*[1024]gopurs_runtime.Value)(v3_6.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v4_7.UnsafePtr)[1]))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_0, "map"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroup1_3, "append"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1])
}), a_4)), b_5)
}))
})
}


