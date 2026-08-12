package Data_Monoid_Conj

import (
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Conj gopurs_runtime.Value
var once_Conj sync.Once
func Get_Conj() gopurs_runtime.Value {
	once_Conj.Do(func() {
		cache_Conj = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Conj(x_0_box)
})
	})
	return cache_Conj
}

var cache_showConj gopurs_runtime.Value
var once_showConj sync.Once
func Get_showConj() gopurs_runtime.Value {
	once_showConj.Do(func() {
		cache_showConj = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showConj(dictShow_0_box)
})
	})
	return cache_showConj
}

var cache_semiringConj gopurs_runtime.Value
var once_semiringConj sync.Once
func Get_semiringConj() gopurs_runtime.Value {
	once_semiringConj.Do(func() {
		cache_semiringConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringConj(dictHeytingAlgebra_0_box)
})
	})
	return cache_semiringConj
}

var cache_semigroupConj gopurs_runtime.Value
var once_semigroupConj sync.Once
func Get_semigroupConj() gopurs_runtime.Value {
	once_semigroupConj.Do(func() {
		cache_semigroupConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupConj(dictHeytingAlgebra_0_box)
})
	})
	return cache_semigroupConj
}

var cache_ordConj gopurs_runtime.Value
var once_ordConj sync.Once
func Get_ordConj() gopurs_runtime.Value {
	once_ordConj.Do(func() {
		cache_ordConj = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordConj(dictOrd_0_box)
})
	})
	return cache_ordConj
}

var cache_monoidConj gopurs_runtime.Value
var once_monoidConj sync.Once
func Get_monoidConj() gopurs_runtime.Value {
	once_monoidConj.Do(func() {
		cache_monoidConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidConj(dictHeytingAlgebra_0_box)
})
	})
	return cache_monoidConj
}

var cache_functorConj gopurs_runtime.Value
var once_functorConj sync.Once
func Get_functorConj() gopurs_runtime.Value {
	once_functorConj.Do(func() {
		cache_functorConj = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorConj
}

var cache_eqConj gopurs_runtime.Value
var once_eqConj sync.Once
func Get_eqConj() gopurs_runtime.Value {
	once_eqConj.Do(func() {
		cache_eqConj = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqConj(dictEq_0_box)
})
	})
	return cache_eqConj
}

var cache_eq1Conj gopurs_runtime.Value
var once_eq1Conj sync.Once
func Get_eq1Conj() gopurs_runtime.Value {
	once_eq1Conj.Do(func() {
		cache_eq1Conj = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Conj
}

var cache_ord1Conj gopurs_runtime.Value
var once_ord1Conj sync.Once
func Get_ord1Conj() gopurs_runtime.Value {
	once_ord1Conj.Do(func() {
		cache_ord1Conj = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Conj()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1Conj
}

var cache_boundedConj gopurs_runtime.Value
var once_boundedConj sync.Once
func Get_boundedConj() gopurs_runtime.Value {
	once_boundedConj.Do(func() {
		cache_boundedConj = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedConj(dictBounded_0_box)
})
	})
	return cache_boundedConj
}

var cache_applyConj gopurs_runtime.Value
var once_applyConj sync.Once
func Get_applyConj() gopurs_runtime.Value {
	once_applyConj.Do(func() {
		cache_applyConj = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyConj
}

var cache_bindConj gopurs_runtime.Value
var once_bindConj sync.Once
func Get_bindConj() gopurs_runtime.Value {
	once_bindConj.Do(func() {
		cache_bindConj = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindConj
}

var cache_applicativeConj gopurs_runtime.Value
var once_applicativeConj sync.Once
func Get_applicativeConj() gopurs_runtime.Value {
	once_applicativeConj.Do(func() {
		cache_applicativeConj = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}), Get_Conj())
	})
	return cache_applicativeConj
}

var cache_monadConj gopurs_runtime.Value
var once_monadConj sync.Once
func Get_monadConj() gopurs_runtime.Value {
	once_monadConj.Do(func() {
		cache_monadConj = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindConj()
}))
	})
	return cache_monadConj
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_ff__2527024921 gopurs_runtime.Value
var once_ff__2527024921 sync.Once
func Get_ff__2527024921() gopurs_runtime.Value {
	once_ff__2527024921.Do(func() {
		cache_ff__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ff__2527024921(dict_0_box)
})
	})
	return cache_ff__2527024921
}

var cache_tt__2527024921 gopurs_runtime.Value
var once_tt__2527024921 sync.Once
func Get_tt__2527024921() gopurs_runtime.Value {
	once_tt__2527024921.Do(func() {
		cache_tt__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tt__2527024921(dict_0_box)
})
	})
	return cache_tt__2527024921
}

var cache_applicativeConj__4045440648 gopurs_runtime.Value
var once_applicativeConj__4045440648 sync.Once
func Get_applicativeConj__4045440648() gopurs_runtime.Value {
	once_applicativeConj__4045440648.Do(func() {
		cache_applicativeConj__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}), Get_Conj())
	})
	return cache_applicativeConj__4045440648
}

var cache_applyConj__3199351098 gopurs_runtime.Value
var once_applyConj__3199351098 sync.Once
func Get_applyConj__3199351098() gopurs_runtime.Value {
	once_applyConj__3199351098.Do(func() {
		cache_applyConj__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyConj__3199351098
}

var cache_bindConj__329376103 gopurs_runtime.Value
var once_bindConj__329376103 sync.Once
func Get_bindConj__329376103() gopurs_runtime.Value {
	once_bindConj__329376103.Do(func() {
		cache_bindConj__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindConj__329376103
}

var cache_eq1Conj__1905950174 gopurs_runtime.Value
var once_eq1Conj__1905950174 sync.Once
func Get_eq1Conj__1905950174() gopurs_runtime.Value {
	once_eq1Conj__1905950174.Do(func() {
		cache_eq1Conj__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Conj__1905950174
}

var cache_functorConj__943655089 gopurs_runtime.Value
var once_functorConj__943655089 sync.Once
func Get_functorConj__943655089() gopurs_runtime.Value {
	once_functorConj__943655089.Do(func() {
		cache_functorConj__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorConj__943655089
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

func Call_Conj(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showConj(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Conj "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semiringConj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"))
}

func Call_semigroupConj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
})
}))
}

func Call_ordConj(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidConj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
semigroupConj1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
})
}))
_ = semigroupConj1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_1_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"))
}

func Call_eqConj(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedConj(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_ff__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_tt__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
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


