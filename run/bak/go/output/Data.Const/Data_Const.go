package Data_Const

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Const gopurs_runtime.Value
var once_Const sync.Once
func Get_Const() gopurs_runtime.Value {
	once_Const.Do(func() {
		Const = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Const
}

var showConst gopurs_runtime.Value
var once_showConst sync.Once
func Get_showConst() gopurs_runtime.Value {
	once_showConst.Do(func() {
		showConst = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Const ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showConst
}

var semiringConst gopurs_runtime.Value
var once_semiringConst sync.Once
func Get_semiringConst() gopurs_runtime.Value {
	once_semiringConst.Do(func() {
		semiringConst = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictSemiring_0
})
	})
	return semiringConst
}

var semigroupoidConst gopurs_runtime.Value
var once_semigroupoidConst sync.Once
func Get_semigroupoidConst() gopurs_runtime.Value {
	once_semigroupoidConst.Do(func() {
		semigroupoidConst = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return semigroupoidConst
}

var semigroupConst gopurs_runtime.Value
var once_semigroupConst sync.Once
func Get_semigroupConst() gopurs_runtime.Value {
	once_semigroupConst.Do(func() {
		semigroupConst = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictSemigroup_0
})
	})
	return semigroupConst
}

var ringConst gopurs_runtime.Value
var once_ringConst sync.Once
func Get_ringConst() gopurs_runtime.Value {
	once_ringConst.Do(func() {
		ringConst = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictRing_0
})
	})
	return ringConst
}

var ordConst gopurs_runtime.Value
var once_ordConst sync.Once
func Get_ordConst() gopurs_runtime.Value {
	once_ordConst.Do(func() {
		ordConst = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordConst
}

var newtypeConst gopurs_runtime.Value
var once_newtypeConst sync.Once
func Get_newtypeConst() gopurs_runtime.Value {
	once_newtypeConst.Do(func() {
		newtypeConst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeConst
}

var monoidConst gopurs_runtime.Value
var once_monoidConst sync.Once
func Get_monoidConst() gopurs_runtime.Value {
	once_monoidConst.Do(func() {
		monoidConst = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
})
	})
	return monoidConst
}

var heytingAlgebraConst gopurs_runtime.Value
var once_heytingAlgebraConst sync.Once
func Get_heytingAlgebraConst() gopurs_runtime.Value {
	once_heytingAlgebraConst.Do(func() {
		heytingAlgebraConst = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictHeytingAlgebra_0
})
	})
	return heytingAlgebraConst
}

var functorConst gopurs_runtime.Value
var once_functorConst sync.Once
func Get_functorConst() gopurs_runtime.Value {
	once_functorConst.Do(func() {
		functorConst = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return m_1
}))
	})
	return functorConst
}

var invariantConst gopurs_runtime.Value
var once_invariantConst sync.Once
func Get_invariantConst() gopurs_runtime.Value {
	once_invariantConst.Do(func() {
		invariantConst = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return m_2
}))
	})
	return invariantConst
}

var euclideanRingConst gopurs_runtime.Value
var once_euclideanRingConst sync.Once
func Get_euclideanRingConst() gopurs_runtime.Value {
	once_euclideanRingConst.Do(func() {
		euclideanRingConst = gopurs_runtime.Func(func(dictEuclideanRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEuclideanRing_0
})
	})
	return euclideanRingConst
}

var eqConst gopurs_runtime.Value
var once_eqConst sync.Once
func Get_eqConst() gopurs_runtime.Value {
	once_eqConst.Do(func() {
		eqConst = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqConst
}

var eq1Const gopurs_runtime.Value
var once_eq1Const sync.Once
func Get_eq1Const() gopurs_runtime.Value {
	once_eq1Const.Do(func() {
		eq1Const = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
eq_1_0 := gopurs_runtime.RecordGet(dictEq_0, "eq")
_ = eq_1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eq_1_0
}))
})
	})
	return eq1Const
}

var ord1Const gopurs_runtime.Value
var once_ord1Const sync.Once
func Get_ord1Const() gopurs_runtime.Value {
	once_ord1Const.Do(func() {
		ord1Const = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
	})
	return ord1Const
}

var commutativeRingConst gopurs_runtime.Value
var once_commutativeRingConst sync.Once
func Get_commutativeRingConst() gopurs_runtime.Value {
	once_commutativeRingConst.Do(func() {
		commutativeRingConst = gopurs_runtime.Func(func(dictCommutativeRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictCommutativeRing_0
})
	})
	return commutativeRingConst
}

var boundedConst gopurs_runtime.Value
var once_boundedConst sync.Once
func Get_boundedConst() gopurs_runtime.Value {
	once_boundedConst.Do(func() {
		boundedConst = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedConst
}

var booleanAlgebraConst gopurs_runtime.Value
var once_booleanAlgebraConst sync.Once
func Get_booleanAlgebraConst() gopurs_runtime.Value {
	once_booleanAlgebraConst.Do(func() {
		booleanAlgebraConst = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBooleanAlgebra_0
})
	})
	return booleanAlgebraConst
}

var applyConst gopurs_runtime.Value
var once_applyConst sync.Once
func Get_applyConst() gopurs_runtime.Value {
	once_applyConst.Do(func() {
		applyConst = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConst()
}))
})
	})
	return applyConst
}

var applicativeConst gopurs_runtime.Value
var once_applicativeConst sync.Once
func Get_applicativeConst() gopurs_runtime.Value {
	once_applicativeConst.Do(func() {
		applicativeConst = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
	})
	return applicativeConst
}


