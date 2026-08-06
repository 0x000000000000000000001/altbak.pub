package Data_Profunctor_Star

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_hoistStar = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistStar(f_0_box, v_1_box, x_2_box)
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
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v1_2, x_3), v_1)
})
})
}))
}

func Call_profunctorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(f_1, x_5)))
})
})
})
}))
}

func Call_strongStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
profunctorStar1_1_0 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(f_1, x_5)))
})
})
})
}))
_ = profunctorStar1_1_0
return gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_5, __local_var_4_2})}
}), gopurs_runtime.Apply(v_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply(v_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))
})
}))
}

func Call_invariantStar(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), f_1, g_2)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(v_3, x_5))
})
})
})
}))
}

func Call_hoistStar(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_functorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Apply(v_2, x_4))
})
})
}))
}

func Call_distributiveStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveStar:
for {
if false { continue distributiveStar }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorStar1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})
}))
_ = functorStar1_1_0
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_1_0
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_distributiveStar(dictDistributive_0), "distribute"), dictFunctor_2)
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(__local_var_5_4, x_6))
})
})
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), dictFunctor_2)
_ = collect1_3_5
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(collect1_3_5, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_6, a_5)
}), f_4)
})
})
}))
}
}

func Call_closedStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
distribute_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "distribute"), pkg_Data_Functor.Get_functorFn())
_ = distribute_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
profunctorStar1_2_1 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "map"), g_4)
_ = __local_var_6_3
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_3, gopurs_runtime.Apply(v_5, gopurs_runtime.Apply(f_3, x_7)))
})
})
})
}))
_ = profunctorStar1_2_1
return gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute_1_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(g_4, x_5))
}))
})
}))
}

func Call_choiceStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
profunctorStar1_2_1 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_1_0, "map"), g_3)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(f_2, x_6)))
})
})
})
}))
_ = profunctorStar1_2_1
return gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_1_0, "map"), pkg_Data_Either.Get_Left())
_ = __local_var_4_3
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(v_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
goto end_branch_4
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0})})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_1_0, "map"), pkg_Data_Either.Get_Right())
_ = __local_var_4_5
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0})})
goto end_branch_6
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t6 = gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(v_3, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
}))
}

func Call_categoryStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupoidStar1_1_0 := gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "bind"), gopurs_runtime.Apply(v1_3, x_4), v_2)
})
})
}))
_ = semigroupoidStar1_1_0
return gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidStar1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"))
}

func Call_applyStar(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorStar1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})
}))
_ = functorStar1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_2, a_4), gopurs_runtime.Apply(v1_3, a_4))
})
})
}))
}

func Call_bindStar(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
applyStar1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})
})
}))
_ = applyStar1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v_2, x_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, a_5, x_4)
}))
})
})
}))
}

func Call_applicativeStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
applyStar1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})
})
}))
_ = applyStar1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_2)
})
}))
}

func Call_monadStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
functorStar1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4)
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_4
applyStar1_2_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_4
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})
})
}))
_ = applyStar1_2_2
applicativeStar1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_2_2
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), a_3)
})
}))
_ = applicativeStar1_1_0
bindStar1_2_7 := Call_bindStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindStar1_2_7
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindStar1_2_7
}))
}

func Call_altStar(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorStar1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})
}))
_ = functorStar1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply(v_2, a_4), gopurs_runtime.Apply(v1_3, a_4))
})
})
}))
}

func Call_plusStar(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
functorStar1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "map"), f_4)
_ = __local_var_6_5
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_3
altStar1_2_1 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_3
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})
})
}))
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
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
functorStar1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4)
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_4
applyStar1_2_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_4
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})
})
}))
_ = applyStar1_2_2
applicativeStar1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_2_2
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), a_3)
})
}))
_ = applicativeStar1_1_0
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_8
empty_3_9 := gopurs_runtime.RecordGet(__local_var_2_8, "empty")
_ = empty_3_9
__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_8, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_11
__local_var_5_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_13
functorStar1_5_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_13, "map"), f_6)
_ = __local_var_8_14
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_14, gopurs_runtime.Apply(v_7, x_9))
})
})
}))
_ = functorStar1_5_12
altStar1_4_10 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_5_12
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_11, "alt"), gopurs_runtime.Apply(v_6, a_8), gopurs_runtime.Apply(v1_7, a_8))
})
})
}))
_ = altStar1_4_10
plusStar1_2_7 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altStar1_4_10
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_3_9
}))
_ = plusStar1_2_7
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStar1_2_7
}))
}

func Call_monadPlusStar(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
monadStar1_1_0 := Call_monadStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadStar1_1_0
alternativeStar1_2_1 := Call_alternativeStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeStar1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStar1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStar1_1_0
}))
}


