package Data_Tuple

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Tuple gopurs_runtime.Value
var once_Tuple sync.Once
func Get_Tuple() gopurs_runtime.Value {
	once_Tuple.Do(func() {
		Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", value0, value1)
})
})
	})
	return Tuple
}

var uncurry gopurs_runtime.Value
var once_uncurry sync.Once
func Get_uncurry() gopurs_runtime.Value {
	once_uncurry.Do(func() {
		uncurry = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
})
	})
	return uncurry
}

var swap gopurs_runtime.Value
var once_swap sync.Once
func Get_swap() gopurs_runtime.Value {
	once_swap.Do(func() {
		swap = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
})
	})
	return swap
}

var snd gopurs_runtime.Value
var once_snd sync.Once
func Get_snd() gopurs_runtime.Value {
	once_snd.Do(func() {
		snd = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
})
	})
	return snd
}

var showTuple gopurs_runtime.Value
var once_showTuple sync.Once
func Get_showTuple() gopurs_runtime.Value {
	once_showTuple.Do(func() {
		showTuple = gopurs_runtime.Func2(func(dictShow_0 gopurs_runtime.Value, dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Tuple " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + ")")
}))
})
	})
	return showTuple
}

var semiringTuple gopurs_runtime.Value
var once_semiringTuple sync.Once
func Get_semiringTuple() gopurs_runtime.Value {
	once_semiringTuple.Do(func() {
		semiringTuple = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
one_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "one")
_ = one_1_0
zero_2_1 := gopurs_runtime.RecordGet(dictSemiring_0, "zero")
_ = zero_2_1
return gopurs_runtime.Func(func(dictSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "one", "mul", "zero", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_3, "add"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Constructor2("Tuple", one_1_0, gopurs_runtime.RecordGet(dictSemiring1_3, "one")), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_3, "mul"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Constructor2("Tuple", zero_2_1, gopurs_runtime.RecordGet(dictSemiring1_3, "zero")))
})
})
	})
	return semiringTuple
}

var semigroupoidTuple gopurs_runtime.Value
var once_semigroupoidTuple sync.Once
func Get_semigroupoidTuple() gopurs_runtime.Value {
	once_semigroupoidTuple.Do(func() {
		semigroupoidTuple = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
}))
	})
	return semigroupoidTuple
}

var semigroupTuple gopurs_runtime.Value
var once_semigroupTuple sync.Once
func Get_semigroupTuple() gopurs_runtime.Value {
	once_semigroupTuple.Do(func() {
		semigroupTuple = gopurs_runtime.Func2(func(dictSemigroup_0 gopurs_runtime.Value, dictSemigroup1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup1_1, "append"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]))
}))
})
	})
	return semigroupTuple
}

var ringTuple gopurs_runtime.Value
var once_ringTuple sync.Once
func Get_ringTuple() gopurs_runtime.Value {
	once_ringTuple.Do(func() {
		ringTuple = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_0
one_2_1 := gopurs_runtime.RecordGet(__local_var_1_0, "one")
_ = one_2_1
zero_3_3 := gopurs_runtime.RecordGet(__local_var_1_0, "zero")
_ = zero_3_3
semiringTuple1_3_2 := gopurs_runtime.Func(func(dictSemiring1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "one", "mul", "zero", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "add"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_4, "add"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]))
}), gopurs_runtime.Constructor2("Tuple", one_2_1, gopurs_runtime.RecordGet(dictSemiring1_4, "one")), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mul"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_4, "mul"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]))
}), gopurs_runtime.Constructor2("Tuple", zero_3_3, gopurs_runtime.RecordGet(dictSemiring1_4, "zero")))
})
_ = semiringTuple1_3_2
return gopurs_runtime.Func(func(dictRing1_4 gopurs_runtime.Value) gopurs_runtime.Value {
semiringTuple2_5_4 := gopurs_runtime.Apply(semiringTuple1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing1_4, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringTuple2_5_4
return gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing1_4, "sub"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringTuple2_5_4
}))
})
})
	})
	return ringTuple
}

var monoidTuple gopurs_runtime.Value
var once_monoidTuple sync.Once
func Get_monoidTuple() gopurs_runtime.Value {
	once_monoidTuple.Do(func() {
		monoidTuple = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonoid1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_3, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_4_2
semigroupTuple2_5_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "append"), (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]))
}))
_ = semigroupTuple2_5_3
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor2("Tuple", mempty_1_0, gopurs_runtime.RecordGet(dictMonoid1_3, "mempty")), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupTuple2_5_3
}))
})
})
	})
	return monoidTuple
}

var heytingAlgebraTuple gopurs_runtime.Value
var once_heytingAlgebraTuple sync.Once
func Get_heytingAlgebraTuple() gopurs_runtime.Value {
	once_heytingAlgebraTuple.Do(func() {
		heytingAlgebraTuple = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
tt_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
_ = tt_1_0
ff_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
_ = ff_2_1
return gopurs_runtime.Func(func(dictHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"tt", "ff", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Constructor2("Tuple", tt_1_0, gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "tt")), gopurs_runtime.Constructor2("Tuple", ff_2_1, gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "ff")), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "implies"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "conj"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "disj"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "not"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
})})
})
})
	})
	return heytingAlgebraTuple
}

var genericTuple gopurs_runtime.Value
var once_genericTuple sync.Once
func Get_genericTuple() gopurs_runtime.Value {
	once_genericTuple.Do(func() {
		genericTuple = gopurs_runtime.RecordDict2("to", "from", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1])
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1])
}))
	})
	return genericTuple
}

var functorTuple gopurs_runtime.Value
var once_functorTuple sync.Once
func Get_functorTuple() gopurs_runtime.Value {
	once_functorTuple.Do(func() {
		functorTuple = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(m_1.UnsafePtr)[0], gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(m_1.UnsafePtr)[1]))
}))
	})
	return functorTuple
}

var invariantTuple gopurs_runtime.Value
var once_invariantTuple sync.Once
func Get_invariantTuple() gopurs_runtime.Value {
	once_invariantTuple.Do(func() {
		invariantTuple = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0], gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[1]))
}))
	})
	return invariantTuple
}

var fst gopurs_runtime.Value
var once_fst sync.Once
func Get_fst() gopurs_runtime.Value {
	once_fst.Do(func() {
		fst = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
})
	})
	return fst
}

var lazyTuple gopurs_runtime.Value
var once_lazyTuple sync.Once
func Get_lazyTuple() gopurs_runtime.Value {
	once_lazyTuple.Do(func() {
		lazyTuple = gopurs_runtime.Func2(func(dictLazy_0 gopurs_runtime.Value, dictLazy1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).UnsafePtr)[0]
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy1_1, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).UnsafePtr)[1]
})))
}))
})
	})
	return lazyTuple
}

var extendTuple gopurs_runtime.Value
var once_extendTuple sync.Once
func Get_extendTuple() gopurs_runtime.Value {
	once_extendTuple.Do(func() {
		extendTuple = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], gopurs_runtime.Apply(f_0, v_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}))
	})
	return extendTuple
}

var eqTuple gopurs_runtime.Value
var once_eqTuple sync.Once
func Get_eqTuple() gopurs_runtime.Value {
	once_eqTuple.Do(func() {
		eqTuple = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[1]).IntVal != 0)
}))
})
	})
	return eqTuple
}

var ordTuple gopurs_runtime.Value
var once_ordTuple sync.Once
func Get_ordTuple() gopurs_runtime.Value {
	once_ordTuple.Do(func() {
		ordTuple = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_1
eqTuple2_4_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[1]).IntVal != 0)
}))
_ = eqTuple2_4_2
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0])
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7_3.StrVal == "LT").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("LT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_7_3.StrVal == "GT").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[1])
}
end_branch_4:
return __t4
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_4_2
}))
})
})
	})
	return ordTuple
}

var eq1Tuple gopurs_runtime.Value
var once_eq1Tuple sync.Once
func Get_eq1Tuple() gopurs_runtime.Value {
	once_eq1Tuple.Do(func() {
		eq1Tuple = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[1]).IntVal != 0)
}))
})
	})
	return eq1Tuple
}

var ord1Tuple gopurs_runtime.Value
var once_ord1Tuple sync.Once
func Get_ord1Tuple() gopurs_runtime.Value {
	once_ord1Tuple.Do(func() {
		ord1Tuple = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eq1Tuple1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[1]).IntVal != 0)
}))
_ = eq1Tuple1_2_1
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func3(func(dictOrd1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0])
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6_2.StrVal == "LT").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("LT")
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_6_2.StrVal == "GT").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("GT")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[1])
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Tuple1_2_1
}))
})
	})
	return ord1Tuple
}

var curry gopurs_runtime.Value
var once_curry sync.Once
func Get_curry() gopurs_runtime.Value {
	once_curry.Do(func() {
		curry = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Constructor2("Tuple", a_1, b_2))
})
	})
	return curry
}

var comonadTuple gopurs_runtime.Value
var once_comonadTuple sync.Once
func Get_comonadTuple() gopurs_runtime.Value {
	once_comonadTuple.Do(func() {
		comonadTuple = gopurs_runtime.RecordDict2("extract", "Extend0", Get_snd(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendTuple()
}))
	})
	return comonadTuple
}

var commutativeRingTuple gopurs_runtime.Value
var once_commutativeRingTuple sync.Once
func Get_commutativeRingTuple() gopurs_runtime.Value {
	once_commutativeRingTuple.Do(func() {
		commutativeRingTuple = gopurs_runtime.Func(func(dictCommutativeRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
ringTuple1_1_0 := gopurs_runtime.Apply(Get_ringTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple1_1_0
return gopurs_runtime.Func(func(dictCommutativeRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
ringTuple2_3_1 := gopurs_runtime.Apply(ringTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing1_2, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple2_3_1
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringTuple2_3_1
}))
})
})
	})
	return commutativeRingTuple
}

var boundedTuple gopurs_runtime.Value
var once_boundedTuple sync.Once
func Get_boundedTuple() gopurs_runtime.Value {
	once_boundedTuple.Do(func() {
		boundedTuple = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
top_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "top")
_ = top_1_0
bottom_2_1 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_3
return gopurs_runtime.Func(func(dictBounded1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_5, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_7_5
eqTuple2_8_7 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "eq"), (*[1024]gopurs_runtime.Value)(x_8.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_5, "eq"), (*[1024]gopurs_runtime.Value)(x_8.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_9.UnsafePtr)[1]).IntVal != 0)
}))
_ = eqTuple2_8_7
ordTuple2_8_6 := gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_9 gopurs_runtime.Value, y_10 gopurs_runtime.Value) gopurs_runtime.Value {
v_11_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "compare"), (*[1024]gopurs_runtime.Value)(x_9.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[0])
_ = v_11_8
var __t9 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_11_8.StrVal == "LT").IntVal != 0 {
__t9 = gopurs_runtime.Constructor0("LT")
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool(v_11_8.StrVal == "GT").IntVal != 0 {
__t9 = gopurs_runtime.Constructor0("GT")
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_4, "compare"), (*[1024]gopurs_runtime.Value)(x_9.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_10.UnsafePtr)[1])
}
end_branch_9:
return __t9
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_8_7
}))
_ = ordTuple2_8_6
return gopurs_runtime.RecordDict3("top", "bottom", "Ord0", gopurs_runtime.Constructor2("Tuple", top_1_0, gopurs_runtime.RecordGet(dictBounded1_5, "top")), gopurs_runtime.Constructor2("Tuple", bottom_2_1, gopurs_runtime.RecordGet(dictBounded1_5, "bottom")), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple2_8_6
}))
})
})
	})
	return boundedTuple
}

var booleanAlgebraTuple gopurs_runtime.Value
var once_booleanAlgebraTuple sync.Once
func Get_booleanAlgebraTuple() gopurs_runtime.Value {
	once_booleanAlgebraTuple.Do(func() {
		booleanAlgebraTuple = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraTuple1_1_0 := gopurs_runtime.Apply(Get_heytingAlgebraTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraTuple1_1_0
return gopurs_runtime.Func(func(dictBooleanAlgebra1_2 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraTuple2_3_1 := gopurs_runtime.Apply(heytingAlgebraTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra1_2, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraTuple2_3_1
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraTuple2_3_1
}))
})
})
	})
	return booleanAlgebraTuple
}

var applyTuple gopurs_runtime.Value
var once_applyTuple sync.Once
func Get_applyTuple() gopurs_runtime.Value {
	once_applyTuple.Do(func() {
		applyTuple = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]), gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}))
})
	})
	return applyTuple
}

var bindTuple gopurs_runtime.Value
var once_bindTuple sync.Once
func Get_bindTuple() gopurs_runtime.Value {
	once_bindTuple.Do(func() {
		bindTuple = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyTuple1_1_0 := gopurs_runtime.Apply(Get_applyTuple(), dictSemigroup_0)
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.Apply(f_3, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
_ = v1_4_1
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_4_1.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v1_4_1.UnsafePtr)[1])
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}))
})
	})
	return bindTuple
}

var applicativeTuple gopurs_runtime.Value
var once_applicativeTuple sync.Once
func Get_applicativeTuple() gopurs_runtime.Value {
	once_applicativeTuple.Do(func() {
		applicativeTuple = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyTuple1_1_0 := gopurs_runtime.Apply(Get_applyTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}))
})
	})
	return applicativeTuple
}

var monadTuple gopurs_runtime.Value
var once_monadTuple sync.Once
func Get_monadTuple() gopurs_runtime.Value {
	once_monadTuple.Do(func() {
		monadTuple = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeTuple1_1_0 := gopurs_runtime.Apply(Get_applicativeTuple(), dictMonoid_0)
_ = applicativeTuple1_1_0
bindTuple1_2_1 := gopurs_runtime.Apply(Get_bindTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindTuple1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeTuple1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindTuple1_2_1
}))
})
	})
	return monadTuple
}




