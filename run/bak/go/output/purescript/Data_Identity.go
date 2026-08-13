package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Identity_Identity gopurs_runtime.Value
var once_Data_Identity_Identity sync.Once
func Get_Data_Identity_Identity() gopurs_runtime.Value {
	once_Data_Identity_Identity.Do(func() {
		cache_Data_Identity_Identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_Identity(x_0_box)
})
	})
	return cache_Data_Identity_Identity
}

var cache_Data_Identity_showIdentity gopurs_runtime.Value
var once_Data_Identity_showIdentity sync.Once
func Get_Data_Identity_showIdentity() gopurs_runtime.Value {
	once_Data_Identity_showIdentity.Do(func() {
		cache_Data_Identity_showIdentity = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_showIdentity(dictShow_0_box)
})
	})
	return cache_Data_Identity_showIdentity
}

var cache_Data_Identity_semiringIdentity gopurs_runtime.Value
var once_Data_Identity_semiringIdentity sync.Once
func Get_Data_Identity_semiringIdentity() gopurs_runtime.Value {
	once_Data_Identity_semiringIdentity.Do(func() {
		cache_Data_Identity_semiringIdentity = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_semiringIdentity(dictSemiring_0_box)
})
	})
	return cache_Data_Identity_semiringIdentity
}

var cache_Data_Identity_semigroupIdentity gopurs_runtime.Value
var once_Data_Identity_semigroupIdentity sync.Once
func Get_Data_Identity_semigroupIdentity() gopurs_runtime.Value {
	once_Data_Identity_semigroupIdentity.Do(func() {
		cache_Data_Identity_semigroupIdentity = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_semigroupIdentity(dictSemigroup_0_box)
})
	})
	return cache_Data_Identity_semigroupIdentity
}

var cache_Data_Identity_ringIdentity gopurs_runtime.Value
var once_Data_Identity_ringIdentity sync.Once
func Get_Data_Identity_ringIdentity() gopurs_runtime.Value {
	once_Data_Identity_ringIdentity.Do(func() {
		cache_Data_Identity_ringIdentity = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_ringIdentity(dictRing_0_box)
})
	})
	return cache_Data_Identity_ringIdentity
}

var cache_Data_Identity_ordIdentity gopurs_runtime.Value
var once_Data_Identity_ordIdentity sync.Once
func Get_Data_Identity_ordIdentity() gopurs_runtime.Value {
	once_Data_Identity_ordIdentity.Do(func() {
		cache_Data_Identity_ordIdentity = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_ordIdentity(dictOrd_0_box)
})
	})
	return cache_Data_Identity_ordIdentity
}

var cache_Data_Identity_newtypeIdentity gopurs_runtime.Value
var once_Data_Identity_newtypeIdentity sync.Once
func Get_Data_Identity_newtypeIdentity() gopurs_runtime.Value {
	once_Data_Identity_newtypeIdentity.Do(func() {
		cache_Data_Identity_newtypeIdentity = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Identity_newtypeIdentity
}

var cache_Data_Identity_monoidIdentity gopurs_runtime.Value
var once_Data_Identity_monoidIdentity sync.Once
func Get_Data_Identity_monoidIdentity() gopurs_runtime.Value {
	once_Data_Identity_monoidIdentity.Do(func() {
		cache_Data_Identity_monoidIdentity = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_monoidIdentity(dictMonoid_0_box)
})
	})
	return cache_Data_Identity_monoidIdentity
}

var cache_Data_Identity_lazyIdentity gopurs_runtime.Value
var once_Data_Identity_lazyIdentity sync.Once
func Get_Data_Identity_lazyIdentity() gopurs_runtime.Value {
	once_Data_Identity_lazyIdentity.Do(func() {
		cache_Data_Identity_lazyIdentity = gopurs_runtime.Func(func(dictLazy_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_lazyIdentity(dictLazy_0_box)
})
	})
	return cache_Data_Identity_lazyIdentity
}

var cache_Data_Identity_heytingAlgebraIdentity gopurs_runtime.Value
var once_Data_Identity_heytingAlgebraIdentity sync.Once
func Get_Data_Identity_heytingAlgebraIdentity() gopurs_runtime.Value {
	once_Data_Identity_heytingAlgebraIdentity.Do(func() {
		cache_Data_Identity_heytingAlgebraIdentity = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_heytingAlgebraIdentity(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_Identity_heytingAlgebraIdentity
}

var cache_Data_Identity_functorIdentity gopurs_runtime.Value
var once_Data_Identity_functorIdentity sync.Once
func Get_Data_Identity_functorIdentity() gopurs_runtime.Value {
	once_Data_Identity_functorIdentity.Do(func() {
		cache_Data_Identity_functorIdentity = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Identity_functorIdentity
}

var cache_Data_Identity_invariantIdentity gopurs_runtime.Value
var once_Data_Identity_invariantIdentity sync.Once
func Get_Data_Identity_invariantIdentity() gopurs_runtime.Value {
	once_Data_Identity_invariantIdentity.Do(func() {
		cache_Data_Identity_invariantIdentity = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
_ = __local_var_0_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "map"), f_1)
})
}))
}()
	})
	return cache_Data_Identity_invariantIdentity
}

var cache_Data_Identity_extendIdentity gopurs_runtime.Value
var once_Data_Identity_extendIdentity sync.Once
func Get_Data_Identity_extendIdentity() gopurs_runtime.Value {
	once_Data_Identity_extendIdentity.Do(func() {
		cache_Data_Identity_extendIdentity = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Identity_extendIdentity
}

var cache_Data_Identity_euclideanRingIdentity gopurs_runtime.Value
var once_Data_Identity_euclideanRingIdentity sync.Once
func Get_Data_Identity_euclideanRingIdentity() gopurs_runtime.Value {
	once_Data_Identity_euclideanRingIdentity.Do(func() {
		cache_Data_Identity_euclideanRingIdentity = gopurs_runtime.Func(func(dictEuclideanRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_euclideanRingIdentity(dictEuclideanRing_0_box)
})
	})
	return cache_Data_Identity_euclideanRingIdentity
}

var cache_Data_Identity_eqIdentity gopurs_runtime.Value
var once_Data_Identity_eqIdentity sync.Once
func Get_Data_Identity_eqIdentity() gopurs_runtime.Value {
	once_Data_Identity_eqIdentity.Do(func() {
		cache_Data_Identity_eqIdentity = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_eqIdentity(dictEq_0_box)
})
	})
	return cache_Data_Identity_eqIdentity
}

var cache_Data_Identity_eq1Identity gopurs_runtime.Value
var once_Data_Identity_eq1Identity sync.Once
func Get_Data_Identity_eq1Identity() gopurs_runtime.Value {
	once_Data_Identity_eq1Identity.Do(func() {
		cache_Data_Identity_eq1Identity = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Identity_eq1Identity
}

var cache_Data_Identity_ord1Identity gopurs_runtime.Value
var once_Data_Identity_ord1Identity sync.Once
func Get_Data_Identity_ord1Identity() gopurs_runtime.Value {
	once_Data_Identity_ord1Identity.Do(func() {
		cache_Data_Identity_ord1Identity = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_1, "eq")
}))
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_Data_Identity_ord1Identity
}

var cache_Data_Identity_comonadIdentity gopurs_runtime.Value
var once_Data_Identity_comonadIdentity sync.Once
func Get_Data_Identity_comonadIdentity() gopurs_runtime.Value {
	once_Data_Identity_comonadIdentity.Do(func() {
		cache_Data_Identity_comonadIdentity = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}))
	})
	return cache_Data_Identity_comonadIdentity
}

var cache_Data_Identity_commutativeRingIdentity gopurs_runtime.Value
var once_Data_Identity_commutativeRingIdentity sync.Once
func Get_Data_Identity_commutativeRingIdentity() gopurs_runtime.Value {
	once_Data_Identity_commutativeRingIdentity.Do(func() {
		cache_Data_Identity_commutativeRingIdentity = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_commutativeRingIdentity(dictCommutativeRing_0_box)
})
	})
	return cache_Data_Identity_commutativeRingIdentity
}

var cache_Data_Identity_boundedIdentity gopurs_runtime.Value
var once_Data_Identity_boundedIdentity sync.Once
func Get_Data_Identity_boundedIdentity() gopurs_runtime.Value {
	once_Data_Identity_boundedIdentity.Do(func() {
		cache_Data_Identity_boundedIdentity = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_boundedIdentity(dictBounded_0_box)
})
	})
	return cache_Data_Identity_boundedIdentity
}

var cache_Data_Identity_booleanAlgebraIdentity gopurs_runtime.Value
var once_Data_Identity_booleanAlgebraIdentity sync.Once
func Get_Data_Identity_booleanAlgebraIdentity() gopurs_runtime.Value {
	once_Data_Identity_booleanAlgebraIdentity.Do(func() {
		cache_Data_Identity_booleanAlgebraIdentity = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Identity_booleanAlgebraIdentity(dictBooleanAlgebra_0_box)
})
	})
	return cache_Data_Identity_booleanAlgebraIdentity
}

var cache_Data_Identity_applyIdentity gopurs_runtime.Value
var once_Data_Identity_applyIdentity sync.Once
func Get_Data_Identity_applyIdentity() gopurs_runtime.Value {
	once_Data_Identity_applyIdentity.Do(func() {
		cache_Data_Identity_applyIdentity = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Identity_applyIdentity
}

var cache_Data_Identity_bindIdentity gopurs_runtime.Value
var once_Data_Identity_bindIdentity sync.Once
func Get_Data_Identity_bindIdentity() gopurs_runtime.Value {
	once_Data_Identity_bindIdentity.Do(func() {
		cache_Data_Identity_bindIdentity = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Identity_bindIdentity
}

var cache_Data_Identity_applicativeIdentity gopurs_runtime.Value
var once_Data_Identity_applicativeIdentity sync.Once
func Get_Data_Identity_applicativeIdentity() gopurs_runtime.Value {
	once_Data_Identity_applicativeIdentity.Do(func() {
		cache_Data_Identity_applicativeIdentity = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Identity_applicativeIdentity
}

var cache_Data_Identity_monadIdentity gopurs_runtime.Value
var once_Data_Identity_monadIdentity sync.Once
func Get_Data_Identity_monadIdentity() gopurs_runtime.Value {
	once_Data_Identity_monadIdentity.Do(func() {
		cache_Data_Identity_monadIdentity = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, v_1)
})
}))
}))
	})
	return cache_Data_Identity_monadIdentity
}

var cache_Data_Identity_altIdentity gopurs_runtime.Value
var once_Data_Identity_altIdentity sync.Once
func Get_Data_Identity_altIdentity() gopurs_runtime.Value {
	once_Data_Identity_altIdentity.Do(func() {
		cache_Data_Identity_altIdentity = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_Data_Identity_altIdentity
}

var cache_Data_Identity_applicativeIdentity__4045440648 gopurs_runtime.Value
var once_Data_Identity_applicativeIdentity__4045440648 sync.Once
func Get_Data_Identity_applicativeIdentity__4045440648() gopurs_runtime.Value {
	once_Data_Identity_applicativeIdentity__4045440648.Do(func() {
		cache_Data_Identity_applicativeIdentity__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Identity_applicativeIdentity__4045440648
}

var cache_Data_Identity_applyIdentity__3199351098 gopurs_runtime.Value
var once_Data_Identity_applyIdentity__3199351098 sync.Once
func Get_Data_Identity_applyIdentity__3199351098() gopurs_runtime.Value {
	once_Data_Identity_applyIdentity__3199351098.Do(func() {
		cache_Data_Identity_applyIdentity__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Identity_applyIdentity__3199351098
}

var cache_Data_Identity_bindIdentity__329376103 gopurs_runtime.Value
var once_Data_Identity_bindIdentity__329376103 sync.Once
func Get_Data_Identity_bindIdentity__329376103() gopurs_runtime.Value {
	once_Data_Identity_bindIdentity__329376103.Do(func() {
		cache_Data_Identity_bindIdentity__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Identity_bindIdentity__329376103
}

var cache_Data_Identity_eq1Identity__1905950174 gopurs_runtime.Value
var once_Data_Identity_eq1Identity__1905950174 sync.Once
func Get_Data_Identity_eq1Identity__1905950174() gopurs_runtime.Value {
	once_Data_Identity_eq1Identity__1905950174.Do(func() {
		cache_Data_Identity_eq1Identity__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Identity_eq1Identity__1905950174
}

var cache_Data_Identity_extendIdentity__2317691272 gopurs_runtime.Value
var once_Data_Identity_extendIdentity__2317691272 sync.Once
func Get_Data_Identity_extendIdentity__2317691272() gopurs_runtime.Value {
	once_Data_Identity_extendIdentity__2317691272.Do(func() {
		cache_Data_Identity_extendIdentity__2317691272 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Identity_extendIdentity__2317691272
}

var cache_Data_Identity_functorIdentity__943655089 gopurs_runtime.Value
var once_Data_Identity_functorIdentity__943655089 sync.Once
func Get_Data_Identity_functorIdentity__943655089() gopurs_runtime.Value {
	once_Data_Identity_functorIdentity__943655089.Do(func() {
		cache_Data_Identity_functorIdentity__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Identity_functorIdentity__943655089
}

var cache_Data_Identity_monadIdentity__2437051429 gopurs_runtime.Value
var once_Data_Identity_monadIdentity__2437051429 sync.Once
func Get_Data_Identity_monadIdentity__2437051429() gopurs_runtime.Value {
	once_Data_Identity_monadIdentity__2437051429.Do(func() {
		cache_Data_Identity_monadIdentity__2437051429 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, v_1)
})
}))
}))
	})
	return cache_Data_Identity_monadIdentity__2437051429
}

var cache_Data_Identity_monadIdentity__1104192371 gopurs_runtime.Value
var once_Data_Identity_monadIdentity__1104192371 sync.Once
func Get_Data_Identity_monadIdentity__1104192371() gopurs_runtime.Value {
	once_Data_Identity_monadIdentity__1104192371.Do(func() {
		cache_Data_Identity_monadIdentity__1104192371 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, v_1)
})
}))
}))
	})
	return cache_Data_Identity_monadIdentity__1104192371
}

func Call_Data_Identity_Identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Identity_showIdentity(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Identity ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Identity_semiringIdentity(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return dictSemiring_0
}

func Call_Data_Identity_semigroupIdentity(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return dictSemigroup_0
}

func Call_Data_Identity_ringIdentity(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return dictRing_0
}

func Call_Data_Identity_ordIdentity(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Identity_monoidIdentity(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return dictMonoid_0
}

func Call_Data_Identity_lazyIdentity(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
return dictLazy_0
}

func Call_Data_Identity_heytingAlgebraIdentity(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return dictHeytingAlgebra_0
}

func Call_Data_Identity_euclideanRingIdentity(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
return dictEuclideanRing_0
}

func Call_Data_Identity_eqIdentity(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Identity_commutativeRingIdentity(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
return dictCommutativeRing_0
}

func Call_Data_Identity_boundedIdentity(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_Data_Identity_booleanAlgebraIdentity(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
return dictBooleanAlgebra_0
}


