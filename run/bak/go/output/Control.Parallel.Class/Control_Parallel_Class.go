package Control_Parallel_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Data_Functor_Costar "gopurs/output/Data.Functor.Costar"
	pkg_Effect_Ref "gopurs/output/Effect.Ref"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var ParCont gopurs_runtime.Value
var once_ParCont sync.Once
func Get_ParCont() gopurs_runtime.Value {
	once_ParCont.Do(func() {
		ParCont = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return ParCont
}

var sequential gopurs_runtime.Value
var once_sequential sync.Once
func Get_sequential() gopurs_runtime.Value {
	once_sequential.Do(func() {
		sequential = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"]
})
	})
	return sequential
}

var parallel gopurs_runtime.Value
var once_parallel sync.Once
func Get_parallel() gopurs_runtime.Value {
	once_parallel.Do(func() {
		parallel = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"]
})
	})
	return parallel
}

var newtypeParCont gopurs_runtime.Value
var once_newtypeParCont sync.Once
func Get_newtypeParCont() gopurs_runtime.Value {
	once_newtypeParCont.Do(func() {
		newtypeParCont = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeParCont
}

var monadParWriterT gopurs_runtime.Value
var once_monadParWriterT sync.Once
func Get_monadParWriterT() gopurs_runtime.Value {
	once_monadParWriterT.Do(func() {
		monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
applyWriterT_2_1 := gopurs_runtime.Value{PtrVal: func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_3_2 := gopurs_runtime.Apply(dictApply_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Functor0_3_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_4, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_3_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v3_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v4_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v3_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v4_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})), v_5)), v1_6)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_4_3
})})
}}
return gopurs_runtime.Func(func(dictParallel_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyWriterT1_4_4 := applyWriterT_2_1.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(gopurs_runtime.Apply(dictParallel_3.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}))
applyWriterT2_5_5 := applyWriterT_2_1.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(gopurs_runtime.Apply(dictParallel_3.PtrVal.(map[string]gopurs_runtime.Value)["Apply1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_3.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], v_6)
}), "sequential": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_3.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], v_6)
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT1_4_4
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_5
})})
})
})
	})
	return monadParWriterT
}

var monadParStar gopurs_runtime.Value
var once_monadParStar sync.Once
func Get_monadParStar() gopurs_runtime.Value {
	once_monadParStar.Do(func() {
		monadParStar = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
})})
applyStar_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_4, a_6)), gopurs_runtime.Apply(v1_5, a_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_3
})})
__local_var_4_5 := gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply1"], gopurs_runtime.Value{})
__local_var_5_6 := gopurs_runtime.Apply(__local_var_4_5.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_6_8 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_9 := gopurs_runtime.Apply(__local_var_5_6.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_6)
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_9, gopurs_runtime.Apply(v_7, x_9))
})
})
})})
applyStar1_6_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_5.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_7, a_9)), gopurs_runtime.Apply(v1_8, a_9))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_6_8
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Apply(v_7, x_8))
})
}), "sequential": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(v_7, x_8))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar_3_2
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_6_7
})})
})
	})
	return monadParStar
}

var monadParReaderT gopurs_runtime.Value
var once_monadParReaderT sync.Once
func Get_monadParReaderT() gopurs_runtime.Value {
	once_monadParReaderT.Do(func() {
		monadParReaderT = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_3)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
})})
applyReaderT_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_4, r_6)), gopurs_runtime.Apply(v1_5, r_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
})})
__local_var_4_5 := gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply1"], gopurs_runtime.Value{})
__local_var_5_6 := gopurs_runtime.Apply(__local_var_4_5.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorReaderT1_6_8 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_9 := gopurs_runtime.Apply(__local_var_5_6.PtrVal.(map[string]gopurs_runtime.Value)["map"], x_6)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_9, gopurs_runtime.Apply(v_8, x_9))
})
})
})})
applyReaderT1_6_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_5.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_7, r_9)), gopurs_runtime.Apply(v1_8, r_9))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_6_8
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Apply(v_7, x_8))
})
}), "sequential": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], gopurs_runtime.Apply(v_7, x_8))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT_3_2
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_6_7
})})
})
	})
	return monadParReaderT
}

var monadParMaybeT gopurs_runtime.Value
var once_monadParMaybeT sync.Once
func Get_monadParMaybeT() gopurs_runtime.Value {
	once_monadParMaybeT.Do(func() {
		monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply1"], gopurs_runtime.Value{})
Functor0_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCompose2_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(f_3, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_3:
return __t3
})), v_4)
})
})})
applyCompose_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Maybe.Get_applyMaybe().PtrVal.(map[string]gopurs_runtime.Value)["apply"]), v_4)), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_2
})})
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applyMaybeT_6_5 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_applyMaybeT(), dictMonad_5)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], v_7)
}), "sequential": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], v_7)
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMaybeT_6_5
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_4_4
})})
})
})
	})
	return monadParMaybeT
}

var monadParExceptT gopurs_runtime.Value
var once_monadParExceptT sync.Once
func Get_monadParExceptT() gopurs_runtime.Value {
	once_monadParExceptT.Do(func() {
		monadParExceptT = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply1"], gopurs_runtime.Value{})
Functor0_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCompose2_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(m_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(f_3, m_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})), v_4)
})
})})
applyCompose_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_applyEither().PtrVal.(map[string]gopurs_runtime.Value)["apply"]), v_4)), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_2
})})
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applyExceptT_6_5 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_applyExceptT(), dictMonad_5)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], v_7)
}), "sequential": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], v_7)
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyExceptT_6_5
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_4_4
})})
})
})
	})
	return monadParExceptT
}

var monadParCostar gopurs_runtime.Value
var once_monadParCostar sync.Once
func Get_monadParCostar() gopurs_runtime.Value {
	once_monadParCostar.Do(func() {
		monadParCostar = gopurs_runtime.Func(func(dictParallel_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["sequential"], x_2))
})
}), "sequential": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(dictParallel_0.PtrVal.(map[string]gopurs_runtime.Value)["parallel"], x_2))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
})})
})
	})
	return monadParCostar
}

var monadParParCont gopurs_runtime.Value
var once_monadParParCont sync.Once
func Get_monadParParCont() gopurs_runtime.Value {
	once_monadParParCont.Do(func() {
		monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorContT1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
})})
applyContT_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"parallel": Get_ParCont(), "sequential": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return v_3
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT_2_1
}), "Apply1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyParCont(), dictMonadEffect_0)
})})
})
	})
	return monadParParCont
}

var functorParCont gopurs_runtime.Value
var once_functorParCont sync.Once
func Get_functorParCont() gopurs_runtime.Value {
	once_functorParCont.Do(func() {
		functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_monadParParCont(), dictMonadEffect_0).PtrVal.(map[string]gopurs_runtime.Value)["sequential"]
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_monadParParCont(), dictMonadEffect_0).PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_1, a_6))
}))
}))
})
})})
})
	})
	return functorParCont
}

var applyParCont gopurs_runtime.Value
var once_applyParCont sync.Once
func Get_applyParCont() gopurs_runtime.Value {
	once_applyParCont.Do(func() {
		applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get__new(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})))), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get__new(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})))), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), rb_6))), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(mb_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Ref.Get_write(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": a_7})), ra_5))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(mb_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, mb_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
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
}))), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), ra_5))), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(ma_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Ref.Get_write(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": b_8})), rb_6))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(ma_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(ma_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], b_8))
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
}))
}))
}))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_functorParCont(), dictMonadEffect_0)
})})
})
	})
	return applyParCont
}

var applicativeParCont gopurs_runtime.Value
var once_applicativeParCont sync.Once
func Get_applicativeParCont() gopurs_runtime.Value {
	once_applicativeParCont.Do(func() {
		applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyParCont1_1_0 := gopurs_runtime.Apply(Get_applyParCont(), dictMonadEffect_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_monadParParCont(), dictMonadEffect_0).PtrVal.(map[string]gopurs_runtime.Value)["parallel"], gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
}))
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_0
})})
})
	})
	return applicativeParCont
}

var altParCont gopurs_runtime.Value
var once_altParCont sync.Once
func Get_altParCont() gopurs_runtime.Value {
	once_altParCont.Do(func() {
		altParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
Bind1_2_1 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
functorParCont1_4_3 := gopurs_runtime.Apply(Get_functorParCont(), dictMonadEffect_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get__new(), gopurs_runtime.Bool(false)))), gopurs_runtime.Func(func(done_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), done_8))), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (b_10).IntVal != 0 {
__t4 = gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit())
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Ref.Get_write(), gopurs_runtime.Bool(true)), done_8))), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_9)
}))
}
end_branch_4:
return __t4
}))
}))), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), done_8))), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_11).IntVal != 0 {
__t5 = gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Ref.Get_write(), gopurs_runtime.Bool(true)), done_8))), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_10)
}))
}
end_branch_5:
return __t5
}))
}))
}))
}))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_4_3
})})
})
	})
	return altParCont
}

var plusParCont gopurs_runtime.Value
var once_plusParCont sync.Once
func Get_plusParCont() gopurs_runtime.Value {
	once_plusParCont.Do(func() {
		plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
altParCont1_1_0 := gopurs_runtime.Apply(Get_altParCont(), dictMonadEffect_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit())
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_1_0
})})
})
	})
	return plusParCont
}

var alternativeParCont gopurs_runtime.Value
var once_alternativeParCont sync.Once
func Get_alternativeParCont() gopurs_runtime.Value {
	once_alternativeParCont.Do(func() {
		alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeParCont1_1_0 := gopurs_runtime.Apply(Get_applicativeParCont(), dictMonadEffect_0)
plusParCont1_2_1 := gopurs_runtime.Apply(Get_plusParCont(), dictMonadEffect_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeParCont1_1_0
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusParCont1_2_1
})})
})
	})
	return alternativeParCont
}


