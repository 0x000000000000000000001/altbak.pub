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
		cache_Data_Profunctor_Star_newtypeStar = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer((&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v1_2, x_3), v_1)
})}))}
}

func Call_Data_Profunctor_Star_profunctorStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar c)])] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2)
_ = __local_var_4_1
// TAST (Let): __local_var_4_0 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_4_0 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(v_3, x_5))
})
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(f_1, x_5))
})
})}))}
}

func Call_Data_Profunctor_Star_strongStar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
// TAST (Let): profunctorStar1_1_0 shape=LitRecord bindingType=(ADT ["Data","Profunctor","Profunctor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f)])])
profunctorStar1_1_0 := (&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar c)])] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2)
_ = __local_var_4_2
// TAST (Let): __local_var_4_1 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_4_1 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(f_1, x_5))
})
})})
_ = profunctorStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1323482783, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorStar1_1_0)}
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=Other bindingType=Any
__local_var_4_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v2_5, __local_var_4_3}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Apply(v_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply(v_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))
})}))}
}

func Call_Data_Profunctor_Star_invariantStar(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
return gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=Any
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), f_1, g_2)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(v_3, x_5))
})
})}))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=Any
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Apply(v_2, x_4))
})
})}))}
}

func Call_Data_Profunctor_Star_distributiveStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveStar:
for {
if false { continue distributiveStar }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStar1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=Any
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})})
_ = functorStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer((&Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_1_0)}
}), gopurs_runtime.Func2(func(dictFunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a), (TypeApp (TypeVar g) [(TypeVar b)])])
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Profunctor_Star_distributiveStar(dictDistributive_0), "distribute"), dictFunctor_2)
_ = __local_var_4_3
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a), (TypeVar b)])]))
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(__local_var_5_4, x_6))
})
}), gopurs_runtime.Func3(func(dictFunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_2))}, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, a_4)
}), f_3)
})}))}
}
}

func Call_Data_Profunctor_Star_closedStar(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): profunctorStar1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Profunctor","Profunctor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f)])])
profunctorStar1_1_0 := (&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar c)])] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), g_3)
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_5_2 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(v_4, x_6))
})
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_2, x_6))
})
})})
_ = profunctorStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 768764671, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorStar1_1_0)}
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictDistributive_0, "distribute"), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(g_3, x_4))
}))
})}))}
}

func Call_Data_Profunctor_Star_choiceStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=Any
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): profunctorStar1_3_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Profunctor","Profunctor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f)])])
profunctorStar1_3_2 := (&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar c)])] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), g_5)
_ = __local_var_7_5
// TAST (Let): __local_var_7_4 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(TypeVar d)]))
__local_var_7_4 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(v_6, x_8))
})
_ = __local_var_7_4
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(f_4, x_8))
})
})})
_ = profunctorStar1_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 3666633887, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorStar1_3_2)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 shape=App(Other) bindingType=Any
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Either_Left())
_ = __local_var_5_6
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(v_4, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0))
goto end_branch_7
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=Any
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Either_Right())
_ = __local_var_5_9
// TAST (Let): __local_var_5_8 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar c)])]))
__local_var_5_8 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(v_4, x_6))
})
_ = __local_var_5_8
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_10
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(__local_var_5_8, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
})}))}
}

func Call_Data_Profunctor_Star_categoryStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupoidStar1_1_0 shape=Let(LitRecord) bindingType=(TypeApp (ADT ["Control","Semigroupoid","Semigroupoid"] []) [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f)])])
semigroupoidStar1_1_0 := (&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "bind"), gopurs_runtime.Apply(v1_3, x_4), v_2)
})})
_ = semigroupoidStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer((&Constructor_Control_Category_Category[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(semigroupoidStar1_1_0)}
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")}))}
}

func Call_Data_Profunctor_Star_applyStar(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStar1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=Any
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})})
_ = functorStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_2, a_4), gopurs_runtime.Apply(v1_3, a_4))
})}))}
}

func Call_Data_Profunctor_Star_bindStar(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStar1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})})
_ = functorStar1_2_2
// TAST (Let): applyStar1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})})
_ = applyStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v_2, x_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, a_5, x_4)
}))
})}))}
}

func Call_Data_Profunctor_Star_applicativeStar(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStar1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})})
_ = functorStar1_2_2
// TAST (Let): applyStar1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})})
_ = applyStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_1_0)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_2)
})}))}
}

func Call_Data_Profunctor_Star_monadStar(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorStar1_3_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_3_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=Any
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4)
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Apply(v_5, x_7))
})
})})
_ = functorStar1_3_4
// TAST (Let): applyStar1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_2_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_3_4)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})})
_ = applyStar1_2_2
// TAST (Let): applicativeStar1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applicativeStar1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_2_2)}
}), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), a_3)
})})
_ = applicativeStar1_1_0
// TAST (Let): __local_var_2_8 shape=App(Other) bindingType=Any
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_8
// TAST (Let): __local_var_3_10 shape=App(Other) bindingType=Any
__local_var_3_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_8, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_10
// TAST (Let): __local_var_4_12 shape=App(Other) bindingType=Any
__local_var_4_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_12
// TAST (Let): functorStar1_4_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_4_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_13 shape=App(Other) bindingType=Any
__local_var_7_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_12, "map"), f_5)
_ = __local_var_7_13
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_13, gopurs_runtime.Apply(v_6, x_8))
})
})})
_ = functorStar1_4_11
// TAST (Let): applyStar1_3_9 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_3_9 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_4_11)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_10, "apply"), gopurs_runtime.Apply(v_5, a_7), gopurs_runtime.Apply(v1_6, a_7))
})})
_ = applyStar1_3_9
// TAST (Let): bindStar1_2_7 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
bindStar1_2_7 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_3_9)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_8, "bind"), gopurs_runtime.Apply(v_4, x_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, a_7, x_6)
}))
})})
_ = bindStar1_2_7
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeStar1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindStar1_2_7)}
})}))}
}

func Call_Data_Profunctor_Star_altStar(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStar1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=Any
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})})
_ = functorStar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_1_0)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply(v_2, a_4), gopurs_runtime.Apply(v1_3, a_4))
})}))}
}

func Call_Data_Profunctor_Star_plusStar(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeVar a)])
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorStar1_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 shape=App(Other) bindingType=Any
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "map"), f_4)
_ = __local_var_6_5
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(v_5, x_7))
})
})})
_ = functorStar1_3_3
// TAST (Let): altStar1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
altStar1_2_1 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_3_3)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})})
_ = altStar1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altStar1_2_1)}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
})}))}
}

func Call_Data_Profunctor_Star_alternativeStar(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorStar1_3_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_3_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=Any
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4)
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, gopurs_runtime.Apply(v_5, x_7))
})
})})
_ = functorStar1_3_4
// TAST (Let): applyStar1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_2_2 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_3_4)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})})
_ = applyStar1_2_2
// TAST (Let): applicativeStar1_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applicativeStar1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_2_2)}
}), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), a_3)
})})
_ = applicativeStar1_1_0
// TAST (Let): __local_var_2_8 shape=App(Other) bindingType=Any
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_8
// TAST (Let): empty_3_9 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeVar a)])
empty_3_9 := gopurs_runtime.RecordGet(__local_var_2_8, "empty")
_ = empty_3_9
// TAST (Let): __local_var_4_11 shape=App(Other) bindingType=Any
__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_8, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_11
// TAST (Let): __local_var_5_13 shape=App(Other) bindingType=Any
__local_var_5_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_13
// TAST (Let): functorStar1_5_12 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_5_12 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_14 shape=App(Other) bindingType=Any
__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_13, "map"), f_6)
_ = __local_var_8_14
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_14, gopurs_runtime.Apply(v_7, x_9))
})
})})
_ = functorStar1_5_12
// TAST (Let): altStar1_4_10 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
altStar1_4_10 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_5_12)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_11, "alt"), gopurs_runtime.Apply(v_6, a_8), gopurs_runtime.Apply(v1_7, a_8))
})})
_ = altStar1_4_10
// TAST (Let): plusStar1_2_7 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
plusStar1_2_7 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altStar1_4_10)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_3_9
})})
_ = plusStar1_2_7
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeStar1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusStar1_2_7)}
})}))}
}

func Call_Data_Profunctor_Star_monadPlusStar(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorStar1_4_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_4_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 shape=App(Other) bindingType=Any
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), f_5)
_ = __local_var_7_8
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_8, gopurs_runtime.Apply(v_6, x_8))
})
})})
_ = functorStar1_4_6
// TAST (Let): applyStar1_3_4 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_3_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_4_6)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "apply"), gopurs_runtime.Apply(v_5, a_7), gopurs_runtime.Apply(v1_6, a_7))
})})
_ = applyStar1_3_4
// TAST (Let): applicativeStar1_2_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applicativeStar1_2_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_3_4)}
}), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "pure"), a_4)
})})
_ = applicativeStar1_2_2
// TAST (Let): __local_var_3_10 shape=App(Other) bindingType=Any
__local_var_3_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_10
// TAST (Let): __local_var_4_12 shape=App(Other) bindingType=Any
__local_var_4_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_10, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_12
// TAST (Let): __local_var_5_14 shape=App(Other) bindingType=Any
__local_var_5_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_14
// TAST (Let): functorStar1_5_13 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_5_13 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_15 shape=App(Other) bindingType=Any
__local_var_8_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_14, "map"), f_6)
_ = __local_var_8_15
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_15, gopurs_runtime.Apply(v_7, x_9))
})
})})
_ = functorStar1_5_13
// TAST (Let): applyStar1_4_11 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_4_11 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_5_13)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_12, "apply"), gopurs_runtime.Apply(v_6, a_8), gopurs_runtime.Apply(v1_7, a_8))
})})
_ = applyStar1_4_11
// TAST (Let): bindStar1_3_9 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
bindStar1_3_9 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_4_11)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, f_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_10, "bind"), gopurs_runtime.Apply(v_5, x_7), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, a_8, x_7)
}))
})})
_ = bindStar1_3_9
// TAST (Let): monadStar1_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Monad","Monad"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
monadStar1_1_0 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeStar1_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindStar1_3_9)}
})})
_ = monadStar1_1_0
// TAST (Let): __local_var_2_17 shape=App(Other) bindingType=Any
__local_var_2_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{})
_ = __local_var_2_17
// TAST (Let): __local_var_3_19 shape=App(Other) bindingType=Any
__local_var_3_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_17, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_19
// TAST (Let): __local_var_4_21 shape=App(Other) bindingType=Any
__local_var_4_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_19, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_21
// TAST (Let): __local_var_5_23 shape=App(Other) bindingType=Any
__local_var_5_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_21, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_23
// TAST (Let): functorStar1_5_22 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_5_22 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_24 shape=App(Other) bindingType=Any
__local_var_8_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_23, "map"), f_6)
_ = __local_var_8_24
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_24, gopurs_runtime.Apply(v_7, x_9))
})
})})
_ = functorStar1_5_22
// TAST (Let): applyStar1_4_20 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_4_20 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_5_22)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_21, "apply"), gopurs_runtime.Apply(v_6, a_8), gopurs_runtime.Apply(v1_7, a_8))
})})
_ = applyStar1_4_20
// TAST (Let): applicativeStar1_3_18 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applicativeStar1_3_18 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_4_20)}
}), gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_19, "pure"), a_5)
})})
_ = applicativeStar1_3_18
// TAST (Let): __local_var_4_26 shape=App(Other) bindingType=Any
__local_var_4_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_17, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_26
// TAST (Let): empty_5_27 shape=Other bindingType=(TypeApp (TypeVar f) [(TypeVar a)])
empty_5_27 := gopurs_runtime.RecordGet(__local_var_4_26, "empty")
_ = empty_5_27
// TAST (Let): __local_var_6_29 shape=App(Other) bindingType=Any
__local_var_6_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_26, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_6_29
// TAST (Let): __local_var_7_31 shape=App(Other) bindingType=Any
__local_var_7_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_31
// TAST (Let): functorStar1_7_30 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_7_30 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_32 shape=App(Other) bindingType=Any
__local_var_10_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_31, "map"), f_8)
_ = __local_var_10_32
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_32, gopurs_runtime.Apply(v_9, x_11))
})
})})
_ = functorStar1_7_30
// TAST (Let): altStar1_6_28 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
altStar1_6_28 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_7_30)}
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_29, "alt"), gopurs_runtime.Apply(v_8, a_10), gopurs_runtime.Apply(v1_9, a_10))
})})
_ = altStar1_6_28
// TAST (Let): plusStar1_4_25 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
plusStar1_4_25 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altStar1_6_28)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_5_27
})})
_ = plusStar1_4_25
// TAST (Let): alternativeStar1_2_16 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Alternative","Alternative"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
alternativeStar1_2_16 := (&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeStar1_3_18)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusStar1_4_25)}
})})
_ = alternativeStar1_2_16
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer((&Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeStar1_2_16)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadStar1_1_0)}
})}))}
}


