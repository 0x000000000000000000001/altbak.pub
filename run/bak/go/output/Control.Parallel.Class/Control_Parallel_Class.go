package Control_Parallel_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Data_Functor_Costar "gopurs/output/Data.Functor.Costar"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var ParCont gopurs_runtime.Value
var once_ParCont sync.Once
func Get_ParCont() gopurs_runtime.Value {
	once_ParCont.Do(func() {
		ParCont = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return ParCont
}

var sequential gopurs_runtime.Value
var once_sequential sync.Once
func Get_sequential() gopurs_runtime.Value {
	once_sequential.Do(func() {
		sequential = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "sequential")
}()
})
	})
	return sequential
}

var parallel gopurs_runtime.Value
var once_parallel sync.Once
func Get_parallel() gopurs_runtime.Value {
	once_parallel.Do(func() {
		parallel = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "parallel")
}()
})
	})
	return parallel
}

var newtypeParCont gopurs_runtime.Value
var once_newtypeParCont sync.Once
func Get_newtypeParCont() gopurs_runtime.Value {
	once_newtypeParCont.Do(func() {
		newtypeParCont = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeParCont
}

var monadParWriterT gopurs_runtime.Value
var once_monadParWriterT sync.Once
func Get_monadParWriterT() gopurs_runtime.Value {
	once_monadParWriterT.Do(func() {
		monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
applyWriterT_2_1 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_3_2
functorWriterT1_4_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1])
}))
}))
_ = functorWriterT1_4_3
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_3_2, "map"), gopurs_runtime.Func2(func(v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v3_7.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v4_8.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*[1024]gopurs_runtime.Value)(v3_7.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v4_8.UnsafePtr)[1]))
}), v_5), v1_6)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_4_3
}))
})
_ = applyWriterT_2_1
return gopurs_runtime.Func(func(dictParallel_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyWriterT1_4_4 := gopurs_runtime.UncurriedApp(applyWriterT_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT1_4_4
applyWriterT2_5_5 := gopurs_runtime.UncurriedApp(applyWriterT_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "Apply1"), gopurs_runtime.Value{}))
_ = applyWriterT2_5_5
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "parallel"), v_6)
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "sequential"), v_6)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT1_4_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_5
}))
})
}()
})
	})
	return monadParWriterT
}

var monadParStar gopurs_runtime.Value
var once_monadParStar sync.Once
func Get_monadParStar() gopurs_runtime.Value {
	once_monadParStar.Do(func() {
		monadParStar = gopurs_runtime.Func(func(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorStar1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
}))
_ = functorStar1_3_3
applyStar_3_2 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_3
}))
_ = applyStar_3_2
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
functorStar1_6_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "map"), f_6)
_ = __local_var_8_9
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_9, gopurs_runtime.Apply(v_7, x_9))
})
}))
_ = functorStar1_6_8
applyStar1_6_7 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "apply"), gopurs_runtime.Apply(v_7, a_9), gopurs_runtime.Apply(v1_8, a_9))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_6_8
}))
_ = applyStar1_6_7
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "parallel"), gopurs_runtime.Apply(v_7, x_8))
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "sequential"), gopurs_runtime.Apply(v_7, x_8))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_6_7
}))
}()
})
	})
	return monadParStar
}

var monadParReaderT gopurs_runtime.Value
var once_monadParReaderT sync.Once
func Get_monadParReaderT() gopurs_runtime.Value {
	once_monadParReaderT.Do(func() {
		monadParReaderT = gopurs_runtime.Func(func(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_3
applyReaderT_3_2 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
}))
_ = applyReaderT_3_2
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
functorReaderT1_6_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "map"), x_6)
_ = __local_var_7_9
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_9, gopurs_runtime.Apply(v_8, x_9))
})
}))
_ = functorReaderT1_6_8
applyReaderT1_6_7 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "apply"), gopurs_runtime.Apply(v_7, r_9), gopurs_runtime.Apply(v1_8, r_9))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_6_8
}))
_ = applyReaderT1_6_7
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "parallel"), gopurs_runtime.Apply(v_7, x_8))
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "sequential"), gopurs_runtime.Apply(v_7, x_8))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_6_7
}))
}()
})
	})
	return monadParReaderT
}

var monadParMaybeT gopurs_runtime.Value
var once_monadParMaybeT sync.Once
func Get_monadParMaybeT() gopurs_runtime.Value {
	once_monadParMaybeT.Do(func() {
		monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
functorCompose2_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_5.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]))
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
_ = functorCompose2_3_2
applyCompose_4_4 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), v_4), v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_2
}))
_ = applyCompose_4_4
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applyMaybeT_6_5 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_applyMaybeT(), dictMonad_5)
_ = applyMaybeT_6_5
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "parallel"), v_7)
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "sequential"), v_7)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMaybeT_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_4_4
}))
})
}()
})
	})
	return monadParMaybeT
}

var monadParExceptT gopurs_runtime.Value
var once_monadParExceptT sync.Once
func Get_monadParExceptT() gopurs_runtime.Value {
	once_monadParExceptT.Do(func() {
		monadParExceptT = gopurs_runtime.Func(func(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
functorCompose2_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_5.StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_5.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(m_5.StrVal == "Right").IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(m_5.UnsafePtr)[0]))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), v_4)
}))
_ = functorCompose2_3_2
applyCompose_4_4 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "apply"), v_4), v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_2
}))
_ = applyCompose_4_4
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applyExceptT_6_5 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_applyExceptT(), dictMonad_5)
_ = applyExceptT_6_5
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "parallel"), v_7)
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "sequential"), v_7)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyExceptT_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_4_4
}))
})
}()
})
	})
	return monadParExceptT
}

var monadParCostar gopurs_runtime.Value
var once_monadParCostar sync.Once
func Get_monadParCostar() gopurs_runtime.Value {
	once_monadParCostar.Do(func() {
		monadParCostar = gopurs_runtime.Func(func(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "sequential"), x_2))
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0_loop, "parallel"), x_2))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}))
}()
})
	})
	return monadParCostar
}

var monadParParCont gopurs_runtime.Value
var once_monadParParCont sync.Once
func Get_monadParParCont() gopurs_runtime.Value {
	once_monadParParCont.Do(func() {
		monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
_ = functorContT1_1_0
applyContT_2_1 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}))
_ = applyContT_2_1
return gopurs_runtime.RecordDict4("parallel", "sequential", "Apply0", "Apply1", Get_ParCont(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return v_3
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyParCont(), dictMonadEffect_0_loop)
}))
}()
})
	})
	return monadParParCont
}

var functorParCont gopurs_runtime.Value
var once_functorParCont sync.Once
func Get_functorParCont() gopurs_runtime.Value {
	once_functorParCont.Do(func() {
		functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadParParCont(), dictMonadEffect_0_loop), "sequential")
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadParParCont(), dictMonadEffect_0_loop), "parallel"), gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_1, a_6))
}))
}))
})
}))
}()
})
	})
	return functorParCont
}

var applyParCont gopurs_runtime.Value
var once_applyParCont sync.Once
func Get_applyParCont() gopurs_runtime.Value {
	once_applyParCont.Do(func() {
		applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_1 := gopurs_runtime.Constructor0("Nothing")
_ = __local_ref_1
return gopurs_runtime.Value{PtrVal: &__local_ref_1}
})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_2 := gopurs_runtime.Constructor0("Nothing")
_ = __local_ref_2
return gopurs_runtime.Value{PtrVal: &__local_ref_2}
})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(rb_6.PtrVal.(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(mb_8.StrVal == "Nothing").IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(ra_5.PtrVal.(*gopurs_runtime.Value)) = gopurs_runtime.Constructor1("Just", a_7)
return gopurs_runtime.Constructor1("Just", a_7)
}))
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(mb_8.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*[1024]gopurs_runtime.Value)(mb_8.UnsafePtr)[0]))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(ra_5.PtrVal.(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(ma_9.StrVal == "Nothing").IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(rb_6.PtrVal.(*gopurs_runtime.Value)) = gopurs_runtime.Constructor1("Just", b_8)
return gopurs_runtime.Constructor1("Just", b_8)
}))
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(ma_9.StrVal == "Just").IntVal != 0 {
__t4 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(ma_9.UnsafePtr)[0], b_8))
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
}))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_functorParCont(), dictMonadEffect_0_loop)
}))
}()
})
	})
	return applyParCont
}

var applicativeParCont gopurs_runtime.Value
var once_applicativeParCont sync.Once
func Get_applicativeParCont() gopurs_runtime.Value {
	once_applicativeParCont.Do(func() {
		applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applyParCont1_1_0 := gopurs_runtime.Apply(Get_applyParCont(), dictMonadEffect_0_loop)
_ = applyParCont1_1_0
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_monadParParCont(), dictMonadEffect_0_loop), "parallel"), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_0
}))
}()
})
	})
	return applicativeParCont
}

var altParCont gopurs_runtime.Value
var once_altParCont sync.Once
func Get_altParCont() gopurs_runtime.Value {
	once_altParCont.Do(func() {
		altParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_2
functorParCont1_4_3 := gopurs_runtime.Apply(Get_functorParCont(), dictMonadEffect_0_loop)
_ = functorParCont1_4_3
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_4 := false
_ = __local_ref_4
return gopurs_runtime.Value{PtrVal: &__local_ref_4}
})), gopurs_runtime.Func(func(done_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(done_8.PtrVal.(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if b_10.IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(done_8.PtrVal.(*gopurs_runtime.Value)) = true
return true
})), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_9)
}))
}
end_branch_5:
return __t5
}))
})), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return *(done_8.PtrVal.(*gopurs_runtime.Value))
})), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if b_11.IntVal != 0 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(done_8.PtrVal.(*gopurs_runtime.Value)) = true
return true
})), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_10)
}))
}
end_branch_6:
return __t6
}))
}))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_4_3
}))
}()
})
	})
	return altParCont
}

var plusParCont gopurs_runtime.Value
var once_plusParCont sync.Once
func Get_plusParCont() gopurs_runtime.Value {
	once_plusParCont.Do(func() {
		plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
altParCont1_1_0 := gopurs_runtime.Apply(Get_altParCont(), dictMonadEffect_0_loop)
_ = altParCont1_1_0
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0_loop, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_1_0
}))
}()
})
	})
	return plusParCont
}

var alternativeParCont gopurs_runtime.Value
var once_alternativeParCont sync.Once
func Get_alternativeParCont() gopurs_runtime.Value {
	once_alternativeParCont.Do(func() {
		alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applicativeParCont1_1_0 := gopurs_runtime.Apply(Get_applicativeParCont(), dictMonadEffect_0_loop)
_ = applicativeParCont1_1_0
plusParCont1_2_1 := gopurs_runtime.Apply(Get_plusParCont(), dictMonadEffect_0_loop)
_ = plusParCont1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeParCont1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusParCont1_2_1
}))
}()
})
	})
	return alternativeParCont
}




