package Data_Monoid_Dual

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Dual gopurs_runtime.Value
var once_Dual sync.Once
func Get_Dual() gopurs_runtime.Value {
	once_Dual.Do(func() {
		cache_Dual = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Dual(x_0_box)
})
	})
	return cache_Dual
}

var cache_showDual gopurs_runtime.Value
var once_showDual sync.Once
func Get_showDual() gopurs_runtime.Value {
	once_showDual.Do(func() {
		cache_showDual = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showDual(dictShow_0_box)
})
	})
	return cache_showDual
}

var cache_semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		cache_semigroupDual = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupDual(dictSemigroup_0_box)
})
	})
	return cache_semigroupDual
}

var cache_ordDual gopurs_runtime.Value
var once_ordDual sync.Once
func Get_ordDual() gopurs_runtime.Value {
	once_ordDual.Do(func() {
		cache_ordDual = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordDual(dictOrd_0_box)
})
	})
	return cache_ordDual
}

var cache_monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		cache_monoidDual = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidDual(dictMonoid_0_box)
})
	})
	return cache_monoidDual
}

var cache_functorDual gopurs_runtime.Value
var once_functorDual sync.Once
func Get_functorDual() gopurs_runtime.Value {
	once_functorDual.Do(func() {
		cache_functorDual = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorDual
}

var cache_eqDual gopurs_runtime.Value
var once_eqDual sync.Once
func Get_eqDual() gopurs_runtime.Value {
	once_eqDual.Do(func() {
		cache_eqDual = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqDual(dictEq_0_box)
})
	})
	return cache_eqDual
}

var cache_eq1Dual gopurs_runtime.Value
var once_eq1Dual sync.Once
func Get_eq1Dual() gopurs_runtime.Value {
	once_eq1Dual.Do(func() {
		cache_eq1Dual = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Dual
}

var cache_ord1Dual gopurs_runtime.Value
var once_ord1Dual sync.Once
func Get_ord1Dual() gopurs_runtime.Value {
	once_ord1Dual.Do(func() {
		cache_ord1Dual = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Dual()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1Dual
}

var cache_boundedDual gopurs_runtime.Value
var once_boundedDual sync.Once
func Get_boundedDual() gopurs_runtime.Value {
	once_boundedDual.Do(func() {
		cache_boundedDual = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedDual(dictBounded_0_box)
})
	})
	return cache_boundedDual
}

var cache_applyDual gopurs_runtime.Value
var once_applyDual sync.Once
func Get_applyDual() gopurs_runtime.Value {
	once_applyDual.Do(func() {
		cache_applyDual = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDual()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyDual
}

var cache_bindDual gopurs_runtime.Value
var once_bindDual sync.Once
func Get_bindDual() gopurs_runtime.Value {
	once_bindDual.Do(func() {
		cache_bindDual = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindDual
}

var cache_applicativeDual gopurs_runtime.Value
var once_applicativeDual sync.Once
func Get_applicativeDual() gopurs_runtime.Value {
	once_applicativeDual.Do(func() {
		cache_applicativeDual = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}), Get_Dual())
	})
	return cache_applicativeDual
}

var cache_monadDual gopurs_runtime.Value
var once_monadDual sync.Once
func Get_monadDual() gopurs_runtime.Value {
	once_monadDual.Do(func() {
		cache_monadDual = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDual()
}))
	})
	return cache_monadDual
}

var cache_applicativeDual__4045440648 gopurs_runtime.Value
var once_applicativeDual__4045440648 sync.Once
func Get_applicativeDual__4045440648() gopurs_runtime.Value {
	once_applicativeDual__4045440648.Do(func() {
		cache_applicativeDual__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}), Get_Dual())
	})
	return cache_applicativeDual__4045440648
}

var cache_applyDual__3199351098 gopurs_runtime.Value
var once_applyDual__3199351098 sync.Once
func Get_applyDual__3199351098() gopurs_runtime.Value {
	once_applyDual__3199351098.Do(func() {
		cache_applyDual__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDual()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyDual__3199351098
}

var cache_bindDual__329376103 gopurs_runtime.Value
var once_bindDual__329376103 sync.Once
func Get_bindDual__329376103() gopurs_runtime.Value {
	once_bindDual__329376103.Do(func() {
		cache_bindDual__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindDual__329376103
}

var cache_eq1Dual__1905950174 gopurs_runtime.Value
var once_eq1Dual__1905950174 sync.Once
func Get_eq1Dual__1905950174() gopurs_runtime.Value {
	once_eq1Dual__1905950174.Do(func() {
		cache_eq1Dual__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Dual__1905950174
}

var cache_functorDual__943655089 gopurs_runtime.Value
var once_functorDual__943655089 sync.Once
func Get_functorDual__943655089() gopurs_runtime.Value {
	once_functorDual__943655089.Do(func() {
		cache_functorDual__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorDual__943655089
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

func Call_Dual(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showDual(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Dual "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semigroupDual(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v1_2, v_1)
})
}))
}

func Call_ordDual(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidDual(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupDual1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), v1_3, v_2)
})
}))
_ = semigroupDual1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_1_0
}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}

func Call_eqDual(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedDual(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


