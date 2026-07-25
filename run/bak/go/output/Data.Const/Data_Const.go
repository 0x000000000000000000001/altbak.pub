package Data_Const

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Const gopurs_runtime.Value
var once_Const sync.Once
func Get_Const() gopurs_runtime.Value {
	once_Const.Do(func() {
		cache_Const = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Const
}

var cache_showConst gopurs_runtime.Value
var once_showConst sync.Once
func Get_showConst() gopurs_runtime.Value {
	once_showConst.Do(func() {
		cache_showConst = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Const ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showConst
}

var cache_semiringConst gopurs_runtime.Value
var once_semiringConst sync.Once
func Get_semiringConst() gopurs_runtime.Value {
	once_semiringConst.Do(func() {
		cache_semiringConst = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return dictSemiring_0
}()
})
	})
	return cache_semiringConst
}

var cache_semigroupoidConst gopurs_runtime.Value
var once_semigroupoidConst sync.Once
func Get_semigroupoidConst() gopurs_runtime.Value {
	once_semigroupoidConst.Do(func() {
		cache_semigroupoidConst = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_semigroupoidConst
}

var cache_semigroupConst gopurs_runtime.Value
var once_semigroupConst sync.Once
func Get_semigroupConst() gopurs_runtime.Value {
	once_semigroupConst.Do(func() {
		cache_semigroupConst = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return dictSemigroup_0
}()
})
	})
	return cache_semigroupConst
}

var cache_ringConst gopurs_runtime.Value
var once_ringConst sync.Once
func Get_ringConst() gopurs_runtime.Value {
	once_ringConst.Do(func() {
		cache_ringConst = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return dictRing_0
}()
})
	})
	return cache_ringConst
}

var cache_ordConst gopurs_runtime.Value
var once_ordConst sync.Once
func Get_ordConst() gopurs_runtime.Value {
	once_ordConst.Do(func() {
		cache_ordConst = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
})
	})
	return cache_ordConst
}

var cache_newtypeConst gopurs_runtime.Value
var once_newtypeConst sync.Once
func Get_newtypeConst() gopurs_runtime.Value {
	once_newtypeConst.Do(func() {
		cache_newtypeConst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeConst
}

var cache_monoidConst gopurs_runtime.Value
var once_monoidConst sync.Once
func Get_monoidConst() gopurs_runtime.Value {
	once_monoidConst.Do(func() {
		cache_monoidConst = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return dictMonoid_0
}()
})
	})
	return cache_monoidConst
}

var cache_heytingAlgebraConst gopurs_runtime.Value
var once_heytingAlgebraConst sync.Once
func Get_heytingAlgebraConst() gopurs_runtime.Value {
	once_heytingAlgebraConst.Do(func() {
		cache_heytingAlgebraConst = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return dictHeytingAlgebra_0
}()
})
	})
	return cache_heytingAlgebraConst
}

var cache_functorConst gopurs_runtime.Value
var once_functorConst sync.Once
func Get_functorConst() gopurs_runtime.Value {
	once_functorConst.Do(func() {
		cache_functorConst = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return m_1
}))
	})
	return cache_functorConst
}

var cache_invariantConst gopurs_runtime.Value
var once_invariantConst sync.Once
func Get_invariantConst() gopurs_runtime.Value {
	once_invariantConst.Do(func() {
		cache_invariantConst = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return m_2
}))
	})
	return cache_invariantConst
}

var cache_euclideanRingConst gopurs_runtime.Value
var once_euclideanRingConst sync.Once
func Get_euclideanRingConst() gopurs_runtime.Value {
	once_euclideanRingConst.Do(func() {
		cache_euclideanRingConst = gopurs_runtime.Func(func(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
return dictEuclideanRing_0
}()
})
	})
	return cache_euclideanRingConst
}

var cache_eqConst gopurs_runtime.Value
var once_eqConst sync.Once
func Get_eqConst() gopurs_runtime.Value {
	once_eqConst.Do(func() {
		cache_eqConst = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqConst
}

var cache_eq1Const gopurs_runtime.Value
var once_eq1Const sync.Once
func Get_eq1Const() gopurs_runtime.Value {
	once_eq1Const.Do(func() {
		cache_eq1Const = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
eq_1_0 := gopurs_runtime.RecordGet(dictEq_0, "eq")
_ = eq_1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eq_1_0
}))
}()
})
	})
	return cache_eq1Const
}

var cache_ord1Const gopurs_runtime.Value
var once_ord1Const sync.Once
func Get_ord1Const() gopurs_runtime.Value {
	once_ord1Const.Do(func() {
		cache_ord1Const = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
eq_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq")
_ = eq_2_1
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return compare_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq_2_1
}))
}))
}()
})
	})
	return cache_ord1Const
}

var cache_commutativeRingConst gopurs_runtime.Value
var once_commutativeRingConst sync.Once
func Get_commutativeRingConst() gopurs_runtime.Value {
	once_commutativeRingConst.Do(func() {
		cache_commutativeRingConst = gopurs_runtime.Func(func(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
return dictCommutativeRing_0
}()
})
	})
	return cache_commutativeRingConst
}

var cache_boundedConst gopurs_runtime.Value
var once_boundedConst sync.Once
func Get_boundedConst() gopurs_runtime.Value {
	once_boundedConst.Do(func() {
		cache_boundedConst = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}()
})
	})
	return cache_boundedConst
}

var cache_booleanAlgebraConst gopurs_runtime.Value
var once_booleanAlgebraConst sync.Once
func Get_booleanAlgebraConst() gopurs_runtime.Value {
	once_booleanAlgebraConst.Do(func() {
		cache_booleanAlgebraConst = gopurs_runtime.Func(func(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
return dictBooleanAlgebra_0
}()
})
	})
	return cache_booleanAlgebraConst
}

var cache_applyConst gopurs_runtime.Value
var once_applyConst sync.Once
func Get_applyConst() gopurs_runtime.Value {
	once_applyConst.Do(func() {
		cache_applyConst = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConst()
}))
}()
})
	})
	return cache_applyConst
}

var cache_applicativeConst gopurs_runtime.Value
var once_applicativeConst sync.Once
func Get_applicativeConst() gopurs_runtime.Value {
	once_applicativeConst.Do(func() {
		cache_applicativeConst = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
applyConst1_3_2 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), v_3, v1_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConst()
}))
_ = applyConst1_3_2
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyConst1_3_2
}))
}()
})
	})
	return cache_applicativeConst
}




