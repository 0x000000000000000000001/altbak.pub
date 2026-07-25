package Data_Identity

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Identity gopurs_runtime.Value
var once_Identity sync.Once
func Get_Identity() gopurs_runtime.Value {
	once_Identity.Do(func() {
		cache_Identity = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Identity
}

var cache_showIdentity gopurs_runtime.Value
var once_showIdentity sync.Once
func Get_showIdentity() gopurs_runtime.Value {
	once_showIdentity.Do(func() {
		cache_showIdentity = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Identity ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showIdentity
}

var cache_semiringIdentity gopurs_runtime.Value
var once_semiringIdentity sync.Once
func Get_semiringIdentity() gopurs_runtime.Value {
	once_semiringIdentity.Do(func() {
		cache_semiringIdentity = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return dictSemiring_0
}()
})
	})
	return cache_semiringIdentity
}

var cache_semigroupIdentity gopurs_runtime.Value
var once_semigroupIdentity sync.Once
func Get_semigroupIdentity() gopurs_runtime.Value {
	once_semigroupIdentity.Do(func() {
		cache_semigroupIdentity = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return dictSemigroup_0
}()
})
	})
	return cache_semigroupIdentity
}

var cache_ringIdentity gopurs_runtime.Value
var once_ringIdentity sync.Once
func Get_ringIdentity() gopurs_runtime.Value {
	once_ringIdentity.Do(func() {
		cache_ringIdentity = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return dictRing_0
}()
})
	})
	return cache_ringIdentity
}

var cache_ordIdentity gopurs_runtime.Value
var once_ordIdentity sync.Once
func Get_ordIdentity() gopurs_runtime.Value {
	once_ordIdentity.Do(func() {
		cache_ordIdentity = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
})
	})
	return cache_ordIdentity
}

var cache_newtypeIdentity gopurs_runtime.Value
var once_newtypeIdentity sync.Once
func Get_newtypeIdentity() gopurs_runtime.Value {
	once_newtypeIdentity.Do(func() {
		cache_newtypeIdentity = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeIdentity
}

var cache_monoidIdentity gopurs_runtime.Value
var once_monoidIdentity sync.Once
func Get_monoidIdentity() gopurs_runtime.Value {
	once_monoidIdentity.Do(func() {
		cache_monoidIdentity = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return dictMonoid_0
}()
})
	})
	return cache_monoidIdentity
}

var cache_lazyIdentity gopurs_runtime.Value
var once_lazyIdentity sync.Once
func Get_lazyIdentity() gopurs_runtime.Value {
	once_lazyIdentity.Do(func() {
		cache_lazyIdentity = gopurs_runtime.Func(func(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
return dictLazy_0
}()
})
	})
	return cache_lazyIdentity
}

var cache_heytingAlgebraIdentity gopurs_runtime.Value
var once_heytingAlgebraIdentity sync.Once
func Get_heytingAlgebraIdentity() gopurs_runtime.Value {
	once_heytingAlgebraIdentity.Do(func() {
		cache_heytingAlgebraIdentity = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return dictHeytingAlgebra_0
}()
})
	})
	return cache_heytingAlgebraIdentity
}

var cache_functorIdentity gopurs_runtime.Value
var once_functorIdentity sync.Once
func Get_functorIdentity() gopurs_runtime.Value {
	once_functorIdentity.Do(func() {
		cache_functorIdentity = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return cache_functorIdentity
}

var cache_invariantIdentity gopurs_runtime.Value
var once_invariantIdentity sync.Once
func Get_invariantIdentity() gopurs_runtime.Value {
	once_invariantIdentity.Do(func() {
		cache_invariantIdentity = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_2)
}))
	})
	return cache_invariantIdentity
}

var cache_extendIdentity gopurs_runtime.Value
var once_extendIdentity sync.Once
func Get_extendIdentity() gopurs_runtime.Value {
	once_extendIdentity.Do(func() {
		cache_extendIdentity = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
}))
	})
	return cache_extendIdentity
}

var cache_euclideanRingIdentity gopurs_runtime.Value
var once_euclideanRingIdentity sync.Once
func Get_euclideanRingIdentity() gopurs_runtime.Value {
	once_euclideanRingIdentity.Do(func() {
		cache_euclideanRingIdentity = gopurs_runtime.Func(func(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
return dictEuclideanRing_0
}()
})
	})
	return cache_euclideanRingIdentity
}

var cache_eqIdentity gopurs_runtime.Value
var once_eqIdentity sync.Once
func Get_eqIdentity() gopurs_runtime.Value {
	once_eqIdentity.Do(func() {
		cache_eqIdentity = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqIdentity
}

var cache_eq1Identity gopurs_runtime.Value
var once_eq1Identity sync.Once
func Get_eq1Identity() gopurs_runtime.Value {
	once_eq1Identity.Do(func() {
		cache_eq1Identity = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Identity
}

var cache_ord1Identity gopurs_runtime.Value
var once_ord1Identity sync.Once
func Get_ord1Identity() gopurs_runtime.Value {
	once_ord1Identity.Do(func() {
		cache_ord1Identity = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Identity()
}))
	})
	return cache_ord1Identity
}

var cache_comonadIdentity gopurs_runtime.Value
var once_comonadIdentity sync.Once
func Get_comonadIdentity() gopurs_runtime.Value {
	once_comonadIdentity.Do(func() {
		cache_comonadIdentity = gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendIdentity()
}))
	})
	return cache_comonadIdentity
}

var cache_commutativeRingIdentity gopurs_runtime.Value
var once_commutativeRingIdentity sync.Once
func Get_commutativeRingIdentity() gopurs_runtime.Value {
	once_commutativeRingIdentity.Do(func() {
		cache_commutativeRingIdentity = gopurs_runtime.Func(func(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
return dictCommutativeRing_0
}()
})
	})
	return cache_commutativeRingIdentity
}

var cache_boundedIdentity gopurs_runtime.Value
var once_boundedIdentity sync.Once
func Get_boundedIdentity() gopurs_runtime.Value {
	once_boundedIdentity.Do(func() {
		cache_boundedIdentity = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}()
})
	})
	return cache_boundedIdentity
}

var cache_booleanAlgebraIdentity gopurs_runtime.Value
var once_booleanAlgebraIdentity sync.Once
func Get_booleanAlgebraIdentity() gopurs_runtime.Value {
	once_booleanAlgebraIdentity.Do(func() {
		cache_booleanAlgebraIdentity = gopurs_runtime.Func(func(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
return dictBooleanAlgebra_0
}()
})
	})
	return cache_booleanAlgebraIdentity
}

var cache_applyIdentity gopurs_runtime.Value
var once_applyIdentity sync.Once
func Get_applyIdentity() gopurs_runtime.Value {
	once_applyIdentity.Do(func() {
		cache_applyIdentity = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
}))
	})
	return cache_applyIdentity
}

var cache_bindIdentity gopurs_runtime.Value
var once_bindIdentity sync.Once
func Get_bindIdentity() gopurs_runtime.Value {
	once_bindIdentity.Do(func() {
		cache_bindIdentity = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyIdentity()
}))
	})
	return cache_bindIdentity
}

var cache_applicativeIdentity gopurs_runtime.Value
var once_applicativeIdentity sync.Once
func Get_applicativeIdentity() gopurs_runtime.Value {
	once_applicativeIdentity.Do(func() {
		cache_applicativeIdentity = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Identity(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyIdentity()
}))
	})
	return cache_applicativeIdentity
}

var cache_monadIdentity gopurs_runtime.Value
var once_monadIdentity sync.Once
func Get_monadIdentity() gopurs_runtime.Value {
	once_monadIdentity.Do(func() {
		cache_monadIdentity = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindIdentity()
}))
	})
	return cache_monadIdentity
}

var cache_altIdentity gopurs_runtime.Value
var once_altIdentity sync.Once
func Get_altIdentity() gopurs_runtime.Value {
	once_altIdentity.Do(func() {
		cache_altIdentity = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
}))
	})
	return cache_altIdentity
}




