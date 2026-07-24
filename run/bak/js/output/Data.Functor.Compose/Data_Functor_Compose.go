package Data_Functor_Compose

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Compose gopurs_runtime.Value
var once_Compose sync.Once
func Get_Compose() gopurs_runtime.Value {
	once_Compose.Do(func() {
		Compose = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Compose
}

var showCompose gopurs_runtime.Value
var once_showCompose sync.Once
func Get_showCompose() gopurs_runtime.Value {
	once_showCompose.Do(func() {
		showCompose = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Compose " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal + ")")
}))
})
	})
	return showCompose
}

var newtypeCompose gopurs_runtime.Value
var once_newtypeCompose sync.Once
func Get_newtypeCompose() gopurs_runtime.Value {
	once_newtypeCompose.Do(func() {
		newtypeCompose = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeCompose
}

var functorCompose gopurs_runtime.Value
var once_functorCompose sync.Once
func Get_functorCompose() gopurs_runtime.Value {
	once_functorCompose.Do(func() {
		functorCompose = gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
}))
})
	})
	return functorCompose
}

var eqCompose gopurs_runtime.Value
var once_eqCompose sync.Once
func Get_eqCompose() gopurs_runtime.Value {
	once_eqCompose.Do(func() {
		eqCompose = gopurs_runtime.Func3(func(dictEq1_0 gopurs_runtime.Value, dictEq11_1 gopurs_runtime.Value, dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq11_3_1
eq11_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_1, x_4, y_5)
})))
_ = eq11_3_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_0, v_4, v1_5)
}))
})
	})
	return eqCompose
}

var ordCompose gopurs_runtime.Value
var once_ordCompose sync.Once
func Get_ordCompose() gopurs_runtime.Value {
	once_ordCompose.Do(func() {
		ordCompose = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
compare11_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_9 gopurs_runtime.Value, y_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_6_4, x_9, y_10)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_8_6
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
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_6_3, v_9, v1_10)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCompose3_8_9
}))
})
})
})
	})
	return ordCompose
}

var eq1Compose gopurs_runtime.Value
var once_eq1Compose sync.Once
func Get_eq1Compose() gopurs_runtime.Value {
	once_eq1Compose.Do(func() {
		eq1Compose = gopurs_runtime.Func2(func(dictEq1_0 gopurs_runtime.Value, dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), dictEq_2)
_ = eq11_3_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_3_0, x_4, y_5)
})))
}))
})
	})
	return eq1Compose
}

var ord1Compose gopurs_runtime.Value
var once_ord1Compose sync.Once
func Get_ord1Compose() gopurs_runtime.Value {
	once_ord1Compose.Do(func() {
		ord1Compose = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordCompose1_1_0 := gopurs_runtime.Apply(Get_ordCompose(), dictOrd1_0)
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
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordCompose2_4_2, dictOrd_7), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Compose2_6_4
}))
})
})
	})
	return ord1Compose
}

var bihoistCompose gopurs_runtime.Value
var once_bihoistCompose sync.Once
func Get_bihoistCompose() gopurs_runtime.Value {
	once_bihoistCompose.Do(func() {
		bihoistCompose = gopurs_runtime.Func4(func(dictFunctor_0 gopurs_runtime.Value, natF_1 gopurs_runtime.Value, natG_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(natF_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), natG_2, v_3))
})
	})
	return bihoistCompose
}

var applyCompose gopurs_runtime.Value
var once_applyCompose sync.Once
func Get_applyCompose() gopurs_runtime.Value {
	once_applyCompose.Do(func() {
		applyCompose = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), apply1_3_1, v_6), v1_7)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}))
})
})
	})
	return applyCompose
}

var applicativeCompose gopurs_runtime.Value
var once_applicativeCompose sync.Once
func Get_applicativeCompose() gopurs_runtime.Value {
	once_applicativeCompose.Do(func() {
		applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
applyCompose2_6_4 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), apply1_5_3, v_8), v1_9)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}))
_ = applyCompose2_6_4
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "pure"), x_7))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose2_6_4
}))
})
})
	})
	return applicativeCompose
}

var altCompose gopurs_runtime.Value
var once_altCompose sync.Once
func Get_altCompose() gopurs_runtime.Value {
	once_altCompose.Do(func() {
		altCompose = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3), v_4)
}))
_ = functorCompose2_3_1
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_4, v1_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_1
}))
})
})
	})
	return altCompose
}

var plusCompose gopurs_runtime.Value
var once_plusCompose sync.Once
func Get_plusCompose() gopurs_runtime.Value {
	once_plusCompose.Do(func() {
		plusCompose = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
altCompose2_6_4 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "alt"), v_6, v1_7)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
}))
_ = altCompose2_6_4
return gopurs_runtime.RecordDict2("empty", "Alt0", empty_1_0, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_6_4
}))
})
})
	})
	return plusCompose
}

var alternativeCompose gopurs_runtime.Value
var once_alternativeCompose sync.Once
func Get_alternativeCompose() gopurs_runtime.Value {
	once_alternativeCompose.Do(func() {
		alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeCompose1_1_0 := gopurs_runtime.Apply(Get_applicativeCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}))
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
altCompose2_8_7 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "alt"), v_8, v1_9)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
}))
_ = altCompose2_8_7
return gopurs_runtime.RecordDict2("empty", "Alt0", empty_3_2, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_8_7
}))
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
})
	})
	return alternativeCompose
}




