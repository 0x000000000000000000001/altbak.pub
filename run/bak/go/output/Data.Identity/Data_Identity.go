package Data_Identity

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Identity gopurs_runtime.Value
var once_Identity sync.Once
func Get_Identity() gopurs_runtime.Value {
	once_Identity.Do(func() {
		Identity = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Identity
}

var showIdentity gopurs_runtime.Value
var once_showIdentity sync.Once
func Get_showIdentity() gopurs_runtime.Value {
	once_showIdentity.Do(func() {
		showIdentity = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Identity " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), v_1).StrVal + ")")
}))
}()
})
	})
	return showIdentity
}

var semiringIdentity gopurs_runtime.Value
var once_semiringIdentity sync.Once
func Get_semiringIdentity() gopurs_runtime.Value {
	once_semiringIdentity.Do(func() {
		semiringIdentity = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return dictSemiring_0_loop
}()
})
	})
	return semiringIdentity
}

var semigroupIdentity gopurs_runtime.Value
var once_semigroupIdentity sync.Once
func Get_semigroupIdentity() gopurs_runtime.Value {
	once_semigroupIdentity.Do(func() {
		semigroupIdentity = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return dictSemigroup_0_loop
}()
})
	})
	return semigroupIdentity
}

var ringIdentity gopurs_runtime.Value
var once_ringIdentity sync.Once
func Get_ringIdentity() gopurs_runtime.Value {
	once_ringIdentity.Do(func() {
		ringIdentity = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return dictRing_0_loop
}()
})
	})
	return ringIdentity
}

var ordIdentity gopurs_runtime.Value
var once_ordIdentity sync.Once
func Get_ordIdentity() gopurs_runtime.Value {
	once_ordIdentity.Do(func() {
		ordIdentity = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0_loop
}()
})
	})
	return ordIdentity
}

var newtypeIdentity gopurs_runtime.Value
var once_newtypeIdentity sync.Once
func Get_newtypeIdentity() gopurs_runtime.Value {
	once_newtypeIdentity.Do(func() {
		newtypeIdentity = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeIdentity
}

var monoidIdentity gopurs_runtime.Value
var once_monoidIdentity sync.Once
func Get_monoidIdentity() gopurs_runtime.Value {
	once_monoidIdentity.Do(func() {
		monoidIdentity = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return dictMonoid_0_loop
}()
})
	})
	return monoidIdentity
}

var lazyIdentity gopurs_runtime.Value
var once_lazyIdentity sync.Once
func Get_lazyIdentity() gopurs_runtime.Value {
	once_lazyIdentity.Do(func() {
		lazyIdentity = gopurs_runtime.Func(func(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
return dictLazy_0_loop
}()
})
	})
	return lazyIdentity
}

var heytingAlgebraIdentity gopurs_runtime.Value
var once_heytingAlgebraIdentity sync.Once
func Get_heytingAlgebraIdentity() gopurs_runtime.Value {
	once_heytingAlgebraIdentity.Do(func() {
		heytingAlgebraIdentity = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return dictHeytingAlgebra_0_loop
}()
})
	})
	return heytingAlgebraIdentity
}

var functorIdentity gopurs_runtime.Value
var once_functorIdentity sync.Once
func Get_functorIdentity() gopurs_runtime.Value {
	once_functorIdentity.Do(func() {
		functorIdentity = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return functorIdentity
}

var invariantIdentity gopurs_runtime.Value
var once_invariantIdentity sync.Once
func Get_invariantIdentity() gopurs_runtime.Value {
	once_invariantIdentity.Do(func() {
		invariantIdentity = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_2)
}))
	})
	return invariantIdentity
}

var extendIdentity gopurs_runtime.Value
var once_extendIdentity sync.Once
func Get_extendIdentity() gopurs_runtime.Value {
	once_extendIdentity.Do(func() {
		extendIdentity = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
}))
	})
	return extendIdentity
}

var euclideanRingIdentity gopurs_runtime.Value
var once_euclideanRingIdentity sync.Once
func Get_euclideanRingIdentity() gopurs_runtime.Value {
	once_euclideanRingIdentity.Do(func() {
		euclideanRingIdentity = gopurs_runtime.Func(func(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
return dictEuclideanRing_0_loop
}()
})
	})
	return euclideanRingIdentity
}

var eqIdentity gopurs_runtime.Value
var once_eqIdentity sync.Once
func Get_eqIdentity() gopurs_runtime.Value {
	once_eqIdentity.Do(func() {
		eqIdentity = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0_loop
}()
})
	})
	return eqIdentity
}

var eq1Identity gopurs_runtime.Value
var once_eq1Identity sync.Once
func Get_eq1Identity() gopurs_runtime.Value {
	once_eq1Identity.Do(func() {
		eq1Identity = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return eq1Identity
}

var ord1Identity gopurs_runtime.Value
var once_ord1Identity sync.Once
func Get_ord1Identity() gopurs_runtime.Value {
	once_ord1Identity.Do(func() {
		ord1Identity = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Identity()
}))
	})
	return ord1Identity
}

var comonadIdentity gopurs_runtime.Value
var once_comonadIdentity sync.Once
func Get_comonadIdentity() gopurs_runtime.Value {
	once_comonadIdentity.Do(func() {
		comonadIdentity = gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendIdentity()
}))
	})
	return comonadIdentity
}

var commutativeRingIdentity gopurs_runtime.Value
var once_commutativeRingIdentity sync.Once
func Get_commutativeRingIdentity() gopurs_runtime.Value {
	once_commutativeRingIdentity.Do(func() {
		commutativeRingIdentity = gopurs_runtime.Func(func(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
return dictCommutativeRing_0_loop
}()
})
	})
	return commutativeRingIdentity
}

var boundedIdentity gopurs_runtime.Value
var once_boundedIdentity sync.Once
func Get_boundedIdentity() gopurs_runtime.Value {
	once_boundedIdentity.Do(func() {
		boundedIdentity = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0_loop
}()
})
	})
	return boundedIdentity
}

var booleanAlgebraIdentity gopurs_runtime.Value
var once_booleanAlgebraIdentity sync.Once
func Get_booleanAlgebraIdentity() gopurs_runtime.Value {
	once_booleanAlgebraIdentity.Do(func() {
		booleanAlgebraIdentity = gopurs_runtime.Func(func(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
return dictBooleanAlgebra_0_loop
}()
})
	})
	return booleanAlgebraIdentity
}

var applyIdentity gopurs_runtime.Value
var once_applyIdentity sync.Once
func Get_applyIdentity() gopurs_runtime.Value {
	once_applyIdentity.Do(func() {
		applyIdentity = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
}))
	})
	return applyIdentity
}

var bindIdentity gopurs_runtime.Value
var once_bindIdentity sync.Once
func Get_bindIdentity() gopurs_runtime.Value {
	once_bindIdentity.Do(func() {
		bindIdentity = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyIdentity()
}))
	})
	return bindIdentity
}

var applicativeIdentity gopurs_runtime.Value
var once_applicativeIdentity sync.Once
func Get_applicativeIdentity() gopurs_runtime.Value {
	once_applicativeIdentity.Do(func() {
		applicativeIdentity = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Identity(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyIdentity()
}))
	})
	return applicativeIdentity
}

var monadIdentity gopurs_runtime.Value
var once_monadIdentity sync.Once
func Get_monadIdentity() gopurs_runtime.Value {
	once_monadIdentity.Do(func() {
		monadIdentity = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindIdentity()
}))
	})
	return monadIdentity
}

var altIdentity gopurs_runtime.Value
var once_altIdentity sync.Once
func Get_altIdentity() gopurs_runtime.Value {
	once_altIdentity.Do(func() {
		altIdentity = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
}))
	})
	return altIdentity
}




