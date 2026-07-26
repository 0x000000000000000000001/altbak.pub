package Data_Profunctor_Star

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Either "gopurs/output/Data.Either"
	unsafe "unsafe"
)

var cache_Star gopurs_runtime.Value
var once_Star sync.Once
func Get_Star() gopurs_runtime.Value {
	once_Star.Do(func() {
		cache_Star = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Star(x_0_box)
})
	})
	return cache_Star
}

var cache_semigroupoidStar gopurs_runtime.Value
var once_semigroupoidStar sync.Once
func Get_semigroupoidStar() gopurs_runtime.Value {
	once_semigroupoidStar.Do(func() {
		cache_semigroupoidStar = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupoidStar(dictBind_0_box)
})
	})
	return cache_semigroupoidStar
}

var cache_profunctorStar gopurs_runtime.Value
var once_profunctorStar sync.Once
func Get_profunctorStar() gopurs_runtime.Value {
	once_profunctorStar.Do(func() {
		cache_profunctorStar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_profunctorStar(dictFunctor_0_box)
})
	})
	return cache_profunctorStar
}

var cache_strongStar gopurs_runtime.Value
var once_strongStar sync.Once
func Get_strongStar() gopurs_runtime.Value {
	once_strongStar.Do(func() {
		cache_strongStar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_strongStar(dictFunctor_0_box)
})
	})
	return cache_strongStar
}

var cache_newtypeStar gopurs_runtime.Value
var once_newtypeStar sync.Once
func Get_newtypeStar() gopurs_runtime.Value {
	once_newtypeStar.Do(func() {
		cache_newtypeStar = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeStar
}

var cache_invariantStar gopurs_runtime.Value
var once_invariantStar sync.Once
func Get_invariantStar() gopurs_runtime.Value {
	once_invariantStar.Do(func() {
		cache_invariantStar = gopurs_runtime.Func(func(dictInvariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_invariantStar(dictInvariant_0_box)
})
	})
	return cache_invariantStar
}

var cache_hoistStar gopurs_runtime.Value
var once_hoistStar sync.Once
func Get_hoistStar() gopurs_runtime.Value {
	once_hoistStar.Do(func() {
		cache_hoistStar = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistStar(f_0_box, v_1_box)
})
	})
	return cache_hoistStar
}

var cache_functorStar gopurs_runtime.Value
var once_functorStar sync.Once
func Get_functorStar() gopurs_runtime.Value {
	once_functorStar.Do(func() {
		cache_functorStar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorStar(dictFunctor_0_box)
})
	})
	return cache_functorStar
}

var cache_distributiveStar gopurs_runtime.Value
var once_distributiveStar sync.Once
func Get_distributiveStar() gopurs_runtime.Value {
	once_distributiveStar.Do(func() {
		cache_distributiveStar = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distributiveStar(dictDistributive_0_box)
})
	})
	return cache_distributiveStar
}

var cache_closedStar gopurs_runtime.Value
var once_closedStar sync.Once
func Get_closedStar() gopurs_runtime.Value {
	once_closedStar.Do(func() {
		cache_closedStar = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_closedStar(dictDistributive_0_box)
})
	})
	return cache_closedStar
}

var cache_choiceStar gopurs_runtime.Value
var once_choiceStar sync.Once
func Get_choiceStar() gopurs_runtime.Value {
	once_choiceStar.Do(func() {
		cache_choiceStar = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choiceStar(dictApplicative_0_box)
})
	})
	return cache_choiceStar
}

var cache_categoryStar gopurs_runtime.Value
var once_categoryStar sync.Once
func Get_categoryStar() gopurs_runtime.Value {
	once_categoryStar.Do(func() {
		cache_categoryStar = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_categoryStar(dictMonad_0_box)
})
	})
	return cache_categoryStar
}

var cache_applyStar gopurs_runtime.Value
var once_applyStar sync.Once
func Get_applyStar() gopurs_runtime.Value {
	once_applyStar.Do(func() {
		cache_applyStar = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyStar(dictApply_0_box)
})
	})
	return cache_applyStar
}

var cache_bindStar gopurs_runtime.Value
var once_bindStar sync.Once
func Get_bindStar() gopurs_runtime.Value {
	once_bindStar.Do(func() {
		cache_bindStar = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindStar(dictBind_0_box)
})
	})
	return cache_bindStar
}

var cache_applicativeStar gopurs_runtime.Value
var once_applicativeStar sync.Once
func Get_applicativeStar() gopurs_runtime.Value {
	once_applicativeStar.Do(func() {
		cache_applicativeStar = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeStar(dictApplicative_0_box)
})
	})
	return cache_applicativeStar
}

var cache_monadStar gopurs_runtime.Value
var once_monadStar sync.Once
func Get_monadStar() gopurs_runtime.Value {
	once_monadStar.Do(func() {
		cache_monadStar = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStar(dictMonad_0_box)
})
	})
	return cache_monadStar
}

var cache_altStar gopurs_runtime.Value
var once_altStar sync.Once
func Get_altStar() gopurs_runtime.Value {
	once_altStar.Do(func() {
		cache_altStar = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altStar(dictAlt_0_box)
})
	})
	return cache_altStar
}

var cache_plusStar gopurs_runtime.Value
var once_plusStar sync.Once
func Get_plusStar() gopurs_runtime.Value {
	once_plusStar.Do(func() {
		cache_plusStar = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusStar(dictPlus_0_box)
})
	})
	return cache_plusStar
}

var cache_alternativeStar gopurs_runtime.Value
var once_alternativeStar sync.Once
func Get_alternativeStar() gopurs_runtime.Value {
	once_alternativeStar.Do(func() {
		cache_alternativeStar = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeStar(dictAlternative_0_box)
})
	})
	return cache_alternativeStar
}

var cache_monadPlusStar gopurs_runtime.Value
var once_monadPlusStar sync.Once
func Get_monadPlusStar() gopurs_runtime.Value {
	once_monadPlusStar.Do(func() {
		cache_monadPlusStar = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusStar(dictMonadPlus_0_box)
})
	})
	return cache_monadPlusStar
}

func Call_Star(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_semigroupoidStar(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictBind_0.UnsafePtr)).V0, gopurs_runtime.Apply(v1_2, x_3), v_1)
}))
}

func Call_profunctorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, g_2), v_3), f_1)
}))
}

func Call_strongStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
profunctorStar1_1_0 := gopurs_runtime.Apply(Get_profunctorStar(), dictFunctor_0)
_ = profunctorStar1_1_0
return gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{v2_5, __local_var_4_1})}
}), gopurs_runtime.Apply(v_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply(v_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))
}))
}

func Call_invariantStar(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictInvariant_0.UnsafePtr)).V0, f_1, g_2), v_3)
}))
}

func Call_hoistStar(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, v_1)
}

func Call_functorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, f_1), v_2)
}))
}

func Call_distributiveStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveStar:
for {
if false { continue distributiveStar }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorStar1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2), v_3)
}))
_ = functorStar1_2_1
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_1
}), gopurs_runtime.Func2(func(dictFunctor_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_distributiveStar(), dictDistributive_0), "distribute"), dictFunctor_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_3, "map"), f_4))
}), gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_4_2 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictDistributive_0.UnsafePtr)).V0, dictFunctor_3)
_ = collect1_4_2
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(collect1_4_2, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_7, a_6)
}), f_5)
})
}))
}
}

func Call_closedStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
distribute_1_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictDistributive_0.UnsafePtr)).V1, pkg_Data_Functor.Get_functorFn())
_ = distribute_1_0
profunctorStar1_2_1 := gopurs_runtime.Apply(Get_profunctorStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = profunctorStar1_2_1
return gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_2_1
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute_1_0, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), v_3, g_4))
}))
}

func Call_choiceStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
pure_2_1 := ((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0
_ = pure_2_1
profunctorStar1_3_2 := gopurs_runtime.Apply(Get_profunctorStar(), Functor0_1_0)
_ = profunctorStar1_3_2
return gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_1_0, "map"), pkg_Data_Either.Get_Left()), v_4)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pure_2_1, pkg_Data_Either.Get_Right())
_ = __local_var_6_4
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t5 = gopurs_runtime.Apply(__local_var_5_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply(__local_var_6_4, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_6 := gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pure_2_1, pkg_Data_Either.Get_Left())
_ = __local_var_5_6
__local_var_6_7 := gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_1_0, "map"), pkg_Data_Either.Get_Right()), v_4)
_ = __local_var_6_7
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(__local_var_5_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(__local_var_6_7, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
}))
}

func Call_categoryStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupoidStar1_2_1 := gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply(v1_3, x_4), v_2)
}))
_ = semigroupoidStar1_2_1
return gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidStar1_2_1
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"))
}

func Call_applyStar(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorStar1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2), v_3)
}))
_ = functorStar1_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_1
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictApply_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
}))
}

func Call_bindStar(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
applyStar1_1_0 := gopurs_runtime.Apply(Get_applyStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applyStar1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictBind_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_2, x_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, a_5, x_4)
}))
}))
}

func Call_applicativeStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
applyStar1_1_0 := gopurs_runtime.Apply(Get_applyStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applyStar1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_1_0
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, a_2)
}))
}

func Call_monadStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeStar1_1_0 := gopurs_runtime.Apply(Get_applicativeStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applicativeStar1_1_0
bindStar1_2_1 := gopurs_runtime.Apply(Get_bindStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = bindStar1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindStar1_2_1
}))
}

func Call_altStar(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorStar1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2), v_3)
}))
_ = functorStar1_2_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_1
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictAlt_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
}))
}

func Call_plusStar(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := ((*gopurs_runtime.RecordData1)(dictPlus_0.UnsafePtr)).V0
_ = empty_1_0
altStar1_2_1 := gopurs_runtime.Apply(Get_altStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = altStar1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altStar1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
}))
}

func Call_alternativeStar(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
applicativeStar1_1_0 := gopurs_runtime.Apply(Get_applicativeStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applicativeStar1_1_0
plusStar1_2_1 := gopurs_runtime.Apply(Get_plusStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = plusStar1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStar1_2_1
}))
}

func Call_monadPlusStar(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
monadStar1_1_0 := gopurs_runtime.Apply(Get_monadStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadStar1_1_0
alternativeStar1_2_1 := gopurs_runtime.Apply(Get_alternativeStar(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = alternativeStar1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStar1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStar1_1_0
}))
}


