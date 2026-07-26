package Data_Functor_App

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_App gopurs_runtime.Value
var once_App sync.Once
func Get_App() gopurs_runtime.Value {
	once_App.Do(func() {
		cache_App = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_App(x_0_box)
})
	})
	return cache_App
}

var cache_showApp gopurs_runtime.Value
var once_showApp sync.Once
func Get_showApp() gopurs_runtime.Value {
	once_showApp.Do(func() {
		cache_showApp = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showApp(dictShow_0_box)
})
	})
	return cache_showApp
}

var cache_semigroupApp gopurs_runtime.Value
var once_semigroupApp sync.Once
func Get_semigroupApp() gopurs_runtime.Value {
	once_semigroupApp.Do(func() {
		cache_semigroupApp = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupApp(dictApply_0_box, dictSemigroup_1_box)
})
	})
	return cache_semigroupApp
}

var cache_plusApp gopurs_runtime.Value
var once_plusApp sync.Once
func Get_plusApp() gopurs_runtime.Value {
	once_plusApp.Do(func() {
		cache_plusApp = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusApp(dictPlus_0_box)
})
	})
	return cache_plusApp
}

var cache_newtypeApp gopurs_runtime.Value
var once_newtypeApp sync.Once
func Get_newtypeApp() gopurs_runtime.Value {
	once_newtypeApp.Do(func() {
		cache_newtypeApp = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeApp
}

var cache_monoidApp gopurs_runtime.Value
var once_monoidApp sync.Once
func Get_monoidApp() gopurs_runtime.Value {
	once_monoidApp.Do(func() {
		cache_monoidApp = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidApp(dictApplicative_0_box)
})
	})
	return cache_monoidApp
}

var cache_monadPlusApp gopurs_runtime.Value
var once_monadPlusApp sync.Once
func Get_monadPlusApp() gopurs_runtime.Value {
	once_monadPlusApp.Do(func() {
		cache_monadPlusApp = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusApp(dictMonadPlus_0_box)
})
	})
	return cache_monadPlusApp
}

var cache_monadApp gopurs_runtime.Value
var once_monadApp sync.Once
func Get_monadApp() gopurs_runtime.Value {
	once_monadApp.Do(func() {
		cache_monadApp = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadApp(dictMonad_0_box)
})
	})
	return cache_monadApp
}

var cache_lazyApp gopurs_runtime.Value
var once_lazyApp sync.Once
func Get_lazyApp() gopurs_runtime.Value {
	once_lazyApp.Do(func() {
		cache_lazyApp = gopurs_runtime.Func(func(dictLazy_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lazyApp(dictLazy_0_box)
})
	})
	return cache_lazyApp
}

var cache_hoistLowerApp gopurs_runtime.Value
var once_hoistLowerApp sync.Once
func Get_hoistLowerApp() gopurs_runtime.Value {
	once_hoistLowerApp.Do(func() {
		cache_hoistLowerApp = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_hoistLowerApp
}

var cache_hoistLiftApp gopurs_runtime.Value
var once_hoistLiftApp sync.Once
func Get_hoistLiftApp() gopurs_runtime.Value {
	once_hoistLiftApp.Do(func() {
		cache_hoistLiftApp = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_hoistLiftApp
}

var cache_hoistApp gopurs_runtime.Value
var once_hoistApp sync.Once
func Get_hoistApp() gopurs_runtime.Value {
	once_hoistApp.Do(func() {
		cache_hoistApp = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistApp(f_0_box, v_1_box)
})
	})
	return cache_hoistApp
}

var cache_functorApp gopurs_runtime.Value
var once_functorApp sync.Once
func Get_functorApp() gopurs_runtime.Value {
	once_functorApp.Do(func() {
		cache_functorApp = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorApp(dictFunctor_0_box)
})
	})
	return cache_functorApp
}

var cache_extendApp gopurs_runtime.Value
var once_extendApp sync.Once
func Get_extendApp() gopurs_runtime.Value {
	once_extendApp.Do(func() {
		cache_extendApp = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendApp(dictExtend_0_box)
})
	})
	return cache_extendApp
}

var cache_eqApp gopurs_runtime.Value
var once_eqApp sync.Once
func Get_eqApp() gopurs_runtime.Value {
	once_eqApp.Do(func() {
		cache_eqApp = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqApp(dictEq1_0_box, dictEq_1_box)
})
	})
	return cache_eqApp
}

var cache_ordApp gopurs_runtime.Value
var once_ordApp sync.Once
func Get_ordApp() gopurs_runtime.Value {
	once_ordApp.Do(func() {
		cache_ordApp = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordApp(dictOrd1_0_box)
})
	})
	return cache_ordApp
}

var cache_eq1App gopurs_runtime.Value
var once_eq1App sync.Once
func Get_eq1App() gopurs_runtime.Value {
	once_eq1App.Do(func() {
		cache_eq1App = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1App(dictEq1_0_box)
})
	})
	return cache_eq1App
}

var cache_ord1App gopurs_runtime.Value
var once_ord1App sync.Once
func Get_ord1App() gopurs_runtime.Value {
	once_ord1App.Do(func() {
		cache_ord1App = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1App(dictOrd1_0_box)
})
	})
	return cache_ord1App
}

var cache_comonadApp gopurs_runtime.Value
var once_comonadApp sync.Once
func Get_comonadApp() gopurs_runtime.Value {
	once_comonadApp.Do(func() {
		cache_comonadApp = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadApp(dictComonad_0_box)
})
	})
	return cache_comonadApp
}

var cache_bindApp gopurs_runtime.Value
var once_bindApp sync.Once
func Get_bindApp() gopurs_runtime.Value {
	once_bindApp.Do(func() {
		cache_bindApp = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindApp(dictBind_0_box)
})
	})
	return cache_bindApp
}

var cache_applyApp gopurs_runtime.Value
var once_applyApp sync.Once
func Get_applyApp() gopurs_runtime.Value {
	once_applyApp.Do(func() {
		cache_applyApp = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyApp(dictApply_0_box)
})
	})
	return cache_applyApp
}

var cache_applicativeApp gopurs_runtime.Value
var once_applicativeApp sync.Once
func Get_applicativeApp() gopurs_runtime.Value {
	once_applicativeApp.Do(func() {
		cache_applicativeApp = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeApp(dictApplicative_0_box)
})
	})
	return cache_applicativeApp
}

var cache_alternativeApp gopurs_runtime.Value
var once_alternativeApp sync.Once
func Get_alternativeApp() gopurs_runtime.Value {
	once_alternativeApp.Do(func() {
		cache_alternativeApp = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeApp(dictAlternative_0_box)
})
	})
	return cache_alternativeApp
}

var cache_altApp gopurs_runtime.Value
var once_altApp sync.Once
func Get_altApp() gopurs_runtime.Value {
	once_altApp.Do(func() {
		cache_altApp = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altApp(dictAlt_0_box)
})
	})
	return cache_altApp
}

func Call_App(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showApp(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(App "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupApp(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
append1_2_0 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append1_2_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), append1_2_0, v_3), v1_4)
}))
}

func Call_plusApp(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
return dictPlus_0
}

func Call_monoidApp(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
append1_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = append1_3_1
semigroupApp2_4_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), append1_3_1, v_4), v1_5)
}))
_ = semigroupApp2_4_2
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupApp2_4_2
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")))
})
}

func Call_monadPlusApp(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
return dictMonadPlus_0
}

func Call_monadApp(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return dictMonad_0
}

func Call_lazyApp(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
return dictLazy_0
}

func Call_hoistApp(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorApp(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return dictFunctor_0
}

func Call_extendApp(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return dictExtend_0
}

func Call_eqApp(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
}))
}

func Call_ordApp(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_2)
_ = compare11_3_1
eq11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_4_2
eqApp2_5_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_4_2, x_5, y_6)
}))
_ = eqApp2_5_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_5_3
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_3_1, x_6, y_7)
}))
})
}

func Call_eq1App(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
}))
}

func Call_ord1App(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1App1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), dictEq_3)
}))
_ = eq1App1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1App1_3_2
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_4)
_ = compare11_5_3
eq11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_6_4
eqApp2_7_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_6_4, x_7, y_8)
}))
_ = eqApp2_7_5
return gopurs_runtime.RecordGet(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_7_5
}), gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_5_3, x_8, y_9)
})), "compare")
}))
}

func Call_comonadApp(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
return dictComonad_0
}

func Call_bindApp(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return dictBind_0
}

func Call_applyApp(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return dictApply_0
}

func Call_applicativeApp(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return dictApplicative_0
}

func Call_alternativeApp(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
return dictAlternative_0
}

func Call_altApp(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
return dictAlt_0
}


