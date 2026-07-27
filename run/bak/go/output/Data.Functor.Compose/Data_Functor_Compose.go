package Data_Functor_Compose

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Compose gopurs_runtime.Value
var once_Compose sync.Once
func Get_Compose() gopurs_runtime.Value {
	once_Compose.Do(func() {
		cache_Compose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Compose(x_0_box)
})
	})
	return cache_Compose
}

var cache_showCompose gopurs_runtime.Value
var once_showCompose sync.Once
func Get_showCompose() gopurs_runtime.Value {
	once_showCompose.Do(func() {
		cache_showCompose = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showCompose(dictShow_0_box))
})
	})
	return cache_showCompose
}

var cache_newtypeCompose gopurs_runtime.Value
var once_newtypeCompose sync.Once
func Get_newtypeCompose() gopurs_runtime.Value {
	once_newtypeCompose.Do(func() {
		cache_newtypeCompose = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))))
	})
	return cache_newtypeCompose
}

var cache_functorCompose gopurs_runtime.Value
var once_functorCompose sync.Once
func Get_functorCompose() gopurs_runtime.Value {
	once_functorCompose.Do(func() {
		cache_functorCompose = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_functorCompose(dictFunctor_0_box, dictFunctor1_1_box))
})
	})
	return cache_functorCompose
}

var cache_eqCompose gopurs_runtime.Value
var once_eqCompose sync.Once
func Get_eqCompose() gopurs_runtime.Value {
	once_eqCompose.Do(func() {
		cache_eqCompose = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eqCompose(dictEq1_0_box, dictEq11_1_box, dictEq_2_box))
})
	})
	return cache_eqCompose
}

var cache_ordCompose gopurs_runtime.Value
var once_ordCompose sync.Once
func Get_ordCompose() gopurs_runtime.Value {
	once_ordCompose.Do(func() {
		cache_ordCompose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordCompose(dictOrd1_0_box)
})
	})
	return cache_ordCompose
}

var cache_eq1Compose gopurs_runtime.Value
var once_eq1Compose sync.Once
func Get_eq1Compose() gopurs_runtime.Value {
	once_eq1Compose.Do(func() {
		cache_eq1Compose = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eq1Compose(dictEq1_0_box, dictEq11_1_box))
})
	})
	return cache_eq1Compose
}

var cache_ord1Compose gopurs_runtime.Value
var once_ord1Compose sync.Once
func Get_ord1Compose() gopurs_runtime.Value {
	once_ord1Compose.Do(func() {
		cache_ord1Compose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Compose(dictOrd1_0_box)
})
	})
	return cache_ord1Compose
}

var cache_bihoistCompose gopurs_runtime.Value
var once_bihoistCompose sync.Once
func Get_bihoistCompose() gopurs_runtime.Value {
	once_bihoistCompose.Do(func() {
		cache_bihoistCompose = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, natF_1_box gopurs_runtime.Value, natG_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bihoistCompose(dictFunctor_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(natF_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(natG_2_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(v_3_box)))
})
	})
	return cache_bihoistCompose
}

var cache_applyCompose gopurs_runtime.Value
var once_applyCompose sync.Once
func Get_applyCompose() gopurs_runtime.Value {
	once_applyCompose.Do(func() {
		cache_applyCompose = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyCompose(dictApply_0_box)
})
	})
	return cache_applyCompose
}

var cache_applicativeCompose gopurs_runtime.Value
var once_applicativeCompose sync.Once
func Get_applicativeCompose() gopurs_runtime.Value {
	once_applicativeCompose.Do(func() {
		cache_applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeCompose(dictApplicative_0_box)
})
	})
	return cache_applicativeCompose
}

var cache_altCompose gopurs_runtime.Value
var once_altCompose sync.Once
func Get_altCompose() gopurs_runtime.Value {
	once_altCompose.Do(func() {
		cache_altCompose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altCompose(dictAlt_0_box)
})
	})
	return cache_altCompose
}

var cache_plusCompose gopurs_runtime.Value
var once_plusCompose sync.Once
func Get_plusCompose() gopurs_runtime.Value {
	once_plusCompose.Do(func() {
		cache_plusCompose = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusCompose(dictPlus_0_box)
})
	})
	return cache_plusCompose
}

var cache_alternativeCompose gopurs_runtime.Value
var once_alternativeCompose sync.Once
func Get_alternativeCompose() gopurs_runtime.Value {
	once_alternativeCompose.Do(func() {
		cache_alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeCompose(dictAlternative_0_box)
})
	})
	return cache_alternativeCompose
}

func Call_Compose(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showCompose(dictShow_0_loop gopurs_runtime.Value) interface{} {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Compose "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")")))
})))
}

func Call_functorCompose(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
})))
}

func Call_eqCompose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) interface{} {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
eq11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq11_3_1
eq11_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_1, x_4, y_5)
})))
_ = eq11_3_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_0, v_4, v1_5)
})))
}

func Call_ordCompose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), dictOrd_5)
_ = compare11_6_4
eq11_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_7_5
eqApp2_8_6 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_5, x_8, y_9)
}))
_ = eqApp2_8_6
compare11_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_8_6
}), gopurs_runtime.Func2(func(x_9 gopurs_runtime.Value, y_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_6_4, x_9, y_10)
})))
_ = compare11_6_3
eq11_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_7_8
eq11_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_8, x_8, y_9)
})))
_ = eq11_7_7
eqCompose3_8_9 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_7, v_8, v1_9)
}))
_ = eqCompose3_8_9
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCompose3_8_9
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_6_3, v_9, v1_10)
}))
})
})
}

func Call_eq1Compose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) interface{} {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq11_3_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_0, x_4, y_5)
})))
})))
}

func Call_ord1Compose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
ordCompose1_1_0 := Call_ordCompose(dictOrd1_0)
_ = ordCompose1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordCompose2_4_2 := gopurs_runtime.Apply(ordCompose1_1_0, dictOrd11_3)
_ = ordCompose2_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_5_3
eq1Compose2_6_4 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_6 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "eq1"), dictEq_6)
_ = eq11_7_5
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_7_5, x_8, y_9)
})))
}))
_ = eq1Compose2_6_4
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Compose2_6_4
}), gopurs_runtime.Func(func(dictOrd_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordCompose2_4_2, dictOrd_7), "compare")
}))
})
}

func Call_bihoistCompose(dictFunctor_0_loop gopurs_runtime.Value, natF_1_loop func(interface{}) interface{}, natG_2_loop func(interface{}) interface{}, v_3_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var natF_1 func(interface{}) interface{} = natF_1_loop
_ = natF_1
var natG_2 func(interface{}) interface{} = natG_2_loop
_ = natG_2
var v_3 interface{} = v_3_loop
_ = v_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(natF_1(gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(natG_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(v_3))))))
}

func Call_applyCompose(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
apply1_3_1 := gopurs_runtime.RecordGet(dictApply1_2, "apply")
_ = apply1_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorCompose2_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "map"), f_5), v_6)
}))
_ = functorCompose2_5_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), apply1_3_1, v_6), v1_7)
}))
})
}

func Call_applicativeCompose(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_4_2
apply1_5_3 := gopurs_runtime.RecordGet(__local_var_4_2, "apply")
_ = apply1_5_3
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
functorCompose2_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "map"), f_7), v_8)
}))
_ = functorCompose2_7_6
applyCompose2_6_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), apply1_5_3, v_8), v1_9)
}))
_ = applyCompose2_6_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose2_6_4
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "pure"), x_7))
}))
})
}

func Call_altCompose(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3), v_4)
}))
_ = functorCompose2_3_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_1
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_4, v1_5)
}))
})
}

func Call_plusCompose(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_4, "map"), f_5), v_6)
}))
_ = functorCompose2_5_3
altCompose2_6_4 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "alt"), v_6, v1_7)
}))
_ = altCompose2_6_4
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_6_4
}), empty_1_0)
})
}

func Call_alternativeCompose(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
applicativeCompose1_1_0 := Call_applicativeCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}))
_ = applicativeCompose1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_1
empty_3_2 := gopurs_runtime.RecordGet(__local_var_2_1, "empty")
_ = empty_3_2
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_4
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
plusCompose1_4_3 := gopurs_runtime.Func(func(dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_6, "map"), f_7), v_8)
}))
_ = functorCompose2_7_6
altCompose2_8_7 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "alt"), v_8, v1_9)
}))
_ = altCompose2_8_7
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_8_7
}), empty_3_2)
})
_ = plusCompose1_4_3
return gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeCompose2_6_8 := gopurs_runtime.Apply(applicativeCompose1_1_0, dictApplicative_5)
_ = applicativeCompose2_6_8
plusCompose2_7_9 := gopurs_runtime.Apply(plusCompose1_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = plusCompose2_7_9
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeCompose2_6_8
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return plusCompose2_7_9
}))
})
}
