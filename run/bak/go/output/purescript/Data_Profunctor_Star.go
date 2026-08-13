package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Star_Star gopurs_runtime.Value
var once_Data_Profunctor_Star_Star sync.Once
func Get_Data_Profunctor_Star_Star() gopurs_runtime.Value {
	once_Data_Profunctor_Star_Star.Do(func() {
		cache_Data_Profunctor_Star_Star = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_Star(x_0_box)
})
	})
	return cache_Data_Profunctor_Star_Star
}

var cache_Data_Profunctor_Star_semigroupoidStar gopurs_runtime.Value
var once_Data_Profunctor_Star_semigroupoidStar sync.Once
func Get_Data_Profunctor_Star_semigroupoidStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_semigroupoidStar.Do(func() {
		cache_Data_Profunctor_Star_semigroupoidStar = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_semigroupoidStar(dictBind_0_box)
})
	})
	return cache_Data_Profunctor_Star_semigroupoidStar
}

var cache_Data_Profunctor_Star_profunctorStar gopurs_runtime.Value
var once_Data_Profunctor_Star_profunctorStar sync.Once
func Get_Data_Profunctor_Star_profunctorStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_profunctorStar.Do(func() {
		cache_Data_Profunctor_Star_profunctorStar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_profunctorStar(dictFunctor_0_box)
})
	})
	return cache_Data_Profunctor_Star_profunctorStar
}

var cache_Data_Profunctor_Star_strongStar gopurs_runtime.Value
var once_Data_Profunctor_Star_strongStar sync.Once
func Get_Data_Profunctor_Star_strongStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_strongStar.Do(func() {
		cache_Data_Profunctor_Star_strongStar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_strongStar(dictFunctor_0_box)
})
	})
	return cache_Data_Profunctor_Star_strongStar
}

var cache_Data_Profunctor_Star_newtypeStar gopurs_runtime.Value
var once_Data_Profunctor_Star_newtypeStar sync.Once
func Get_Data_Profunctor_Star_newtypeStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_newtypeStar.Do(func() {
		cache_Data_Profunctor_Star_newtypeStar = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Profunctor_Star_newtypeStar
}

var cache_Data_Profunctor_Star_invariantStar gopurs_runtime.Value
var once_Data_Profunctor_Star_invariantStar sync.Once
func Get_Data_Profunctor_Star_invariantStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_invariantStar.Do(func() {
		cache_Data_Profunctor_Star_invariantStar = gopurs_runtime.Func(func(dictInvariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_invariantStar(dictInvariant_0_box)
})
	})
	return cache_Data_Profunctor_Star_invariantStar
}

var cache_Data_Profunctor_Star_hoistStar gopurs_runtime.Value
var once_Data_Profunctor_Star_hoistStar sync.Once
func Get_Data_Profunctor_Star_hoistStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_hoistStar.Do(func() {
		cache_Data_Profunctor_Star_hoistStar = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_hoistStar(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Data_Profunctor_Star_hoistStar
}

var cache_Data_Profunctor_Star_functorStar gopurs_runtime.Value
var once_Data_Profunctor_Star_functorStar sync.Once
func Get_Data_Profunctor_Star_functorStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_functorStar.Do(func() {
		cache_Data_Profunctor_Star_functorStar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_functorStar(dictFunctor_0_box)
})
	})
	return cache_Data_Profunctor_Star_functorStar
}

var cache_Data_Profunctor_Star_distributiveStar gopurs_runtime.Value
var once_Data_Profunctor_Star_distributiveStar sync.Once
func Get_Data_Profunctor_Star_distributiveStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_distributiveStar.Do(func() {
		cache_Data_Profunctor_Star_distributiveStar = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_distributiveStar(dictDistributive_0_box)
})
	})
	return cache_Data_Profunctor_Star_distributiveStar
}

var cache_Data_Profunctor_Star_closedStar gopurs_runtime.Value
var once_Data_Profunctor_Star_closedStar sync.Once
func Get_Data_Profunctor_Star_closedStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_closedStar.Do(func() {
		cache_Data_Profunctor_Star_closedStar = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_closedStar(dictDistributive_0_box)
})
	})
	return cache_Data_Profunctor_Star_closedStar
}

var cache_Data_Profunctor_Star_choiceStar gopurs_runtime.Value
var once_Data_Profunctor_Star_choiceStar sync.Once
func Get_Data_Profunctor_Star_choiceStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_choiceStar.Do(func() {
		cache_Data_Profunctor_Star_choiceStar = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_choiceStar(dictApplicative_0_box)
})
	})
	return cache_Data_Profunctor_Star_choiceStar
}

var cache_Data_Profunctor_Star_categoryStar gopurs_runtime.Value
var once_Data_Profunctor_Star_categoryStar sync.Once
func Get_Data_Profunctor_Star_categoryStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_categoryStar.Do(func() {
		cache_Data_Profunctor_Star_categoryStar = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_categoryStar(dictMonad_0_box)
})
	})
	return cache_Data_Profunctor_Star_categoryStar
}

var cache_Data_Profunctor_Star_applyStar gopurs_runtime.Value
var once_Data_Profunctor_Star_applyStar sync.Once
func Get_Data_Profunctor_Star_applyStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_applyStar.Do(func() {
		cache_Data_Profunctor_Star_applyStar = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_applyStar(dictApply_0_box)
})
	})
	return cache_Data_Profunctor_Star_applyStar
}

var cache_Data_Profunctor_Star_bindStar gopurs_runtime.Value
var once_Data_Profunctor_Star_bindStar sync.Once
func Get_Data_Profunctor_Star_bindStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_bindStar.Do(func() {
		cache_Data_Profunctor_Star_bindStar = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_bindStar(dictBind_0_box)
})
	})
	return cache_Data_Profunctor_Star_bindStar
}

var cache_Data_Profunctor_Star_applicativeStar gopurs_runtime.Value
var once_Data_Profunctor_Star_applicativeStar sync.Once
func Get_Data_Profunctor_Star_applicativeStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_applicativeStar.Do(func() {
		cache_Data_Profunctor_Star_applicativeStar = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_applicativeStar(dictApplicative_0_box)
})
	})
	return cache_Data_Profunctor_Star_applicativeStar
}

var cache_Data_Profunctor_Star_monadStar gopurs_runtime.Value
var once_Data_Profunctor_Star_monadStar sync.Once
func Get_Data_Profunctor_Star_monadStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_monadStar.Do(func() {
		cache_Data_Profunctor_Star_monadStar = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_monadStar(dictMonad_0_box)
})
	})
	return cache_Data_Profunctor_Star_monadStar
}

var cache_Data_Profunctor_Star_altStar gopurs_runtime.Value
var once_Data_Profunctor_Star_altStar sync.Once
func Get_Data_Profunctor_Star_altStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_altStar.Do(func() {
		cache_Data_Profunctor_Star_altStar = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_altStar(dictAlt_0_box)
})
	})
	return cache_Data_Profunctor_Star_altStar
}

var cache_Data_Profunctor_Star_plusStar gopurs_runtime.Value
var once_Data_Profunctor_Star_plusStar sync.Once
func Get_Data_Profunctor_Star_plusStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_plusStar.Do(func() {
		cache_Data_Profunctor_Star_plusStar = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_plusStar(dictPlus_0_box)
})
	})
	return cache_Data_Profunctor_Star_plusStar
}

var cache_Data_Profunctor_Star_alternativeStar gopurs_runtime.Value
var once_Data_Profunctor_Star_alternativeStar sync.Once
func Get_Data_Profunctor_Star_alternativeStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_alternativeStar.Do(func() {
		cache_Data_Profunctor_Star_alternativeStar = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_alternativeStar(dictAlternative_0_box)
})
	})
	return cache_Data_Profunctor_Star_alternativeStar
}

var cache_Data_Profunctor_Star_monadPlusStar gopurs_runtime.Value
var once_Data_Profunctor_Star_monadPlusStar sync.Once
func Get_Data_Profunctor_Star_monadPlusStar() gopurs_runtime.Value {
	once_Data_Profunctor_Star_monadPlusStar.Do(func() {
		cache_Data_Profunctor_Star_monadPlusStar = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Star_monadPlusStar(dictMonadPlus_0_box)
})
	})
	return cache_Data_Profunctor_Star_monadPlusStar
}

func Call_Data_Profunctor_Star_Star(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Star_semigroupoidStar(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Profunctor_Star_profunctorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2)
_ = __local_var_4_1
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(v_3, x_5))
})
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(f_1, x_5))
})
})
})
}))
}

func Call_Data_Profunctor_Star_strongStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
// TAST (Let): profunctorStar1_1_0 -> gopurs_runtime.Value
profunctorStar1_1_0 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2)
_ = __local_var_4_2
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(f_1, x_5))
})
})
})
}))
_ = profunctorStar1_1_0
return gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1
_ = __local_var_4_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, v2_5, __local_var_4_3})}
}), gopurs_runtime.Apply(v_2, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0))
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply(v_2, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1))
})
}))
}

func Call_Data_Profunctor_Star_invariantStar(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), f_1, g_2)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(v_3, x_5))
})
})
})
}))
}

func Call_Data_Profunctor_Star_hoistStar(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Data_Profunctor_Star_functorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Apply(v_2, x_4))
})
})
}))
}

func Call_Data_Profunctor_Star_distributiveStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveStar:
for {
if false { continue distributiveStar }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStar1_1_0 -> gopurs_runtime.Value
functorStar1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
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
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Profunctor_Star_distributiveStar(dictDistributive_0), "distribute"), dictFunctor_2)
_ = __local_var_4_3
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(__local_var_5_4, x_6))
})
})
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_2))}, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, a_4)
}), f_3)
})
})
}))
}
}

func Call_Data_Profunctor_Star_closedStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): profunctorStar1_1_0 -> gopurs_runtime.Value
profunctorStar1_1_0 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), g_3)
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(v_4, x_6))
})
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_2, x_6))
})
})
})
}))
_ = profunctorStar1_1_0
return gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictDistributive_0, "distribute"), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_functorFn()))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(g_3, x_4))
}))
})
}))
}

func Call_Data_Profunctor_Star_choiceStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): profunctorStar1_3_2 -> gopurs_runtime.Value
profunctorStar1_3_2 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), g_5)
_ = __local_var_7_5
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(v_6, x_8))
})
_ = __local_var_7_4
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(f_4, x_8))
})
})
})
}))
_ = profunctorStar1_3_2
return gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_7 -> gopurs_runtime.Value
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Either_Left())
_ = __local_var_5_7
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_7, gopurs_runtime.Apply(v_4, x_6))
})
_ = __local_var_5_6
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(__local_var_5_6, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0})})
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
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Either_Right())
_ = __local_var_5_10
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_10, gopurs_runtime.Apply(v_4, x_6))
})
_ = __local_var_5_9
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0})})
goto end_branch_11
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(__local_var_5_9, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
}))
}

func Call_Data_Profunctor_Star_categoryStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupoidStar1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_applyStar(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStar1_1_0 -> gopurs_runtime.Value
functorStar1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_bindStar(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStar1_2_2 -> gopurs_runtime.Value
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
// TAST (Let): applyStar1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_applicativeStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStar1_2_2 -> gopurs_runtime.Value
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
// TAST (Let): applyStar1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_monadStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorStar1_3_4 -> gopurs_runtime.Value
functorStar1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4)
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_4
// TAST (Let): applyStar1_2_2 -> gopurs_runtime.Value
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
// TAST (Let): applicativeStar1_1_0 -> gopurs_runtime.Value
applicativeStar1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_2_2
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), a_3)
})
}))
_ = applicativeStar1_1_0
// TAST (Let): bindStar1_2_7 -> gopurs_runtime.Value
bindStar1_2_7 := Call_Data_Profunctor_Star_bindStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindStar1_2_7
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindStar1_2_7
}))
}

func Call_Data_Profunctor_Star_altStar(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStar1_1_0 -> gopurs_runtime.Value
functorStar1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_plusStar(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorStar1_3_3 -> gopurs_runtime.Value
functorStar1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "map"), f_4)
_ = __local_var_6_5
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_3
// TAST (Let): altStar1_2_1 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_alternativeStar(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorStar1_3_4 -> gopurs_runtime.Value
functorStar1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4)
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_4
// TAST (Let): applyStar1_2_2 -> gopurs_runtime.Value
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
// TAST (Let): applicativeStar1_1_0 -> gopurs_runtime.Value
applicativeStar1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_2_2
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), a_3)
})
}))
_ = applicativeStar1_1_0
// TAST (Let): __local_var_2_8 -> gopurs_runtime.Value
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_8
// TAST (Let): empty_3_9 -> gopurs_runtime.Value
empty_3_9 := gopurs_runtime.RecordGet(__local_var_2_8, "empty")
_ = empty_3_9
// TAST (Let): __local_var_4_11 -> gopurs_runtime.Value
__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_8, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_11
// TAST (Let): __local_var_5_13 -> gopurs_runtime.Value
__local_var_5_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_13
// TAST (Let): functorStar1_5_12 -> gopurs_runtime.Value
functorStar1_5_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_14 -> gopurs_runtime.Value
__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_13, "map"), f_6)
_ = __local_var_8_14
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_14, gopurs_runtime.Apply(v_7, x_9))
})
})
}))
_ = functorStar1_5_12
// TAST (Let): altStar1_4_10 -> gopurs_runtime.Value
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
// TAST (Let): plusStar1_2_7 -> gopurs_runtime.Value
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

func Call_Data_Profunctor_Star_monadPlusStar(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
// TAST (Let): monadStar1_1_0 -> gopurs_runtime.Value
monadStar1_1_0 := Call_Data_Profunctor_Star_monadStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadStar1_1_0
// TAST (Let): alternativeStar1_2_1 -> gopurs_runtime.Value
alternativeStar1_2_1 := Call_Data_Profunctor_Star_alternativeStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeStar1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStar1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStar1_1_0
}))
}


