package Data_Tuple

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_Tuple gopurs_runtime.Value
var once_Tuple sync.Once
func Get_Tuple() gopurs_runtime.Value {
	once_Tuple.Do(func() {
		cache_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{value0, value1})}
})
})
	})
	return cache_Tuple
}

var cache_uncurry gopurs_runtime.Value
var once_uncurry sync.Once
func Get_uncurry() gopurs_runtime.Value {
	once_uncurry.Do(func() {
		cache_uncurry = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncurry(f_0_box, v_1_box)
})
	})
	return cache_uncurry
}

var cache_swap gopurs_runtime.Value
var once_swap sync.Once
func Get_swap() gopurs_runtime.Value {
	once_swap.Do(func() {
		cache_swap = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{(*Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V0})}
}()
})
	})
	return cache_swap
}

var cache_snd gopurs_runtime.Value
var once_snd sync.Once
func Get_snd() gopurs_runtime.Value {
	once_snd.Do(func() {
		cache_snd = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1
}()
})
	})
	return cache_snd
}

var cache_showTuple gopurs_runtime.Value
var once_showTuple sync.Once
func Get_showTuple() gopurs_runtime.Value {
	once_showTuple.Do(func() {
		cache_showTuple = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showTuple(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showTuple
}

var cache_semiringTuple gopurs_runtime.Value
var once_semiringTuple sync.Once
func Get_semiringTuple() gopurs_runtime.Value {
	once_semiringTuple.Do(func() {
		cache_semiringTuple = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
one_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "one")
_ = one_1_0
zero_2_1 := gopurs_runtime.RecordGet(dictSemiring_0, "zero")
_ = zero_2_1
return gopurs_runtime.Func(func(dictSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "one", "mul", "zero", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_3, "add"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{one_1_0, gopurs_runtime.RecordGet(dictSemiring1_3, "one")})}, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_3, "mul"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{zero_2_1, gopurs_runtime.RecordGet(dictSemiring1_3, "zero")})})
})
}()
})
	})
	return cache_semiringTuple
}

var cache_semigroupoidTuple gopurs_runtime.Value
var once_semigroupoidTuple sync.Once
func Get_semigroupoidTuple() gopurs_runtime.Value {
	once_semigroupoidTuple.Do(func() {
		cache_semigroupoidTuple = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{(*Data_Data_Tuple_Tuple)(v1_1.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1})}
}))
	})
	return cache_semigroupoidTuple
}

var cache_semigroupTuple gopurs_runtime.Value
var once_semigroupTuple sync.Once
func Get_semigroupTuple() gopurs_runtime.Value {
	once_semigroupTuple.Do(func() {
		cache_semigroupTuple = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictSemigroup1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupTuple(dictSemigroup_0_box, dictSemigroup1_1_box)
})
	})
	return cache_semigroupTuple
}

var cache_ringTuple gopurs_runtime.Value
var once_ringTuple sync.Once
func Get_ringTuple() gopurs_runtime.Value {
	once_ringTuple.Do(func() {
		cache_ringTuple = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_0
one_2_1 := gopurs_runtime.RecordGet(__local_var_1_0, "one")
_ = one_2_1
zero_3_3 := gopurs_runtime.RecordGet(__local_var_1_0, "zero")
_ = zero_3_3
semiringTuple1_3_2 := gopurs_runtime.Func(func(dictSemiring1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "one", "mul", "zero", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "add"), (*Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_4, "add"), (*Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{one_2_1, gopurs_runtime.RecordGet(dictSemiring1_4, "one")})}, gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mul"), (*Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_4, "mul"), (*Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{zero_3_3, gopurs_runtime.RecordGet(dictSemiring1_4, "zero")})})
})
_ = semiringTuple1_3_2
return gopurs_runtime.Func(func(dictRing1_4 gopurs_runtime.Value) gopurs_runtime.Value {
semiringTuple2_5_4 := gopurs_runtime.Apply(semiringTuple1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing1_4, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringTuple2_5_4
return gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), (*Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing1_4, "sub"), (*Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringTuple2_5_4
}))
})
}()
})
	})
	return cache_ringTuple
}

var cache_monoidTuple gopurs_runtime.Value
var once_monoidTuple sync.Once
func Get_monoidTuple() gopurs_runtime.Value {
	once_monoidTuple.Do(func() {
		cache_monoidTuple = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonoid1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_3, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_4_2
semigroupTuple2_5_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), (*Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "append"), (*Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1)})}
}))
_ = semigroupTuple2_5_3
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{mempty_1_0, gopurs_runtime.RecordGet(dictMonoid1_3, "mempty")})}, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupTuple2_5_3
}))
})
}()
})
	})
	return cache_monoidTuple
}

var cache_heytingAlgebraTuple gopurs_runtime.Value
var once_heytingAlgebraTuple sync.Once
func Get_heytingAlgebraTuple() gopurs_runtime.Value {
	once_heytingAlgebraTuple.Do(func() {
		cache_heytingAlgebraTuple = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
tt_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
_ = tt_1_0
ff_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
_ = ff_2_1
return gopurs_runtime.Func(func(dictHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"tt", "ff", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{tt_1_0, gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "tt")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{ff_2_1, gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "ff")})}, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "implies"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "conj"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "disj"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "not"), (*Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)})}
})})
})
}()
})
	})
	return cache_heytingAlgebraTuple
}

var cache_genericTuple gopurs_runtime.Value
var once_genericTuple sync.Once
func Get_genericTuple() gopurs_runtime.Value {
	once_genericTuple.Do(func() {
		cache_genericTuple = gopurs_runtime.RecordDict2("to", "from", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{(*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(x_0.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product)(x_0.UnsafePtr).V1})}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Data_Data_Generic_Rep_Product{(*Data_Data_Tuple_Tuple)(x_0.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(x_0.UnsafePtr).V1})}
}))
	})
	return cache_genericTuple
}

var cache_functorTuple gopurs_runtime.Value
var once_functorTuple sync.Once
func Get_functorTuple() gopurs_runtime.Value {
	once_functorTuple.Do(func() {
		cache_functorTuple = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{(*Data_Data_Tuple_Tuple)(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Data_Data_Tuple_Tuple)(m_1.UnsafePtr).V1)})}
}))
	})
	return cache_functorTuple
}

var cache_invariantTuple gopurs_runtime.Value
var once_invariantTuple sync.Once
func Get_invariantTuple() gopurs_runtime.Value {
	once_invariantTuple.Do(func() {
		cache_invariantTuple = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{(*Data_Data_Tuple_Tuple)(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Data_Data_Tuple_Tuple)(m_2.UnsafePtr).V1)})}
}))
	})
	return cache_invariantTuple
}

var cache_fst gopurs_runtime.Value
var once_fst sync.Once
func Get_fst() gopurs_runtime.Value {
	once_fst.Do(func() {
		cache_fst = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
}()
})
	})
	return cache_fst
}

var cache_lazyTuple gopurs_runtime.Value
var once_lazyTuple sync.Once
func Get_lazyTuple() gopurs_runtime.Value {
	once_lazyTuple.Do(func() {
		cache_lazyTuple = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, dictLazy1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lazyTuple(dictLazy_0_box, dictLazy1_1_box)
})
	})
	return cache_lazyTuple
}

var cache_extendTuple gopurs_runtime.Value
var once_extendTuple sync.Once
func Get_extendTuple() gopurs_runtime.Value {
	once_extendTuple.Do(func() {
		cache_extendTuple = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{(*Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, v_1)})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}))
	})
	return cache_extendTuple
}

var cache_eqTuple gopurs_runtime.Value
var once_eqTuple sync.Once
func Get_eqTuple() gopurs_runtime.Value {
	once_eqTuple.Do(func() {
		cache_eqTuple = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqTuple(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqTuple
}

var cache_ordTuple gopurs_runtime.Value
var once_ordTuple sync.Once
func Get_ordTuple() gopurs_runtime.Value {
	once_ordTuple.Do(func() {
		cache_ordTuple = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_1
eqTuple2_4_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Data_Data_Tuple_Tuple)(x_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), (*Data_Data_Tuple_Tuple)(x_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_5.UnsafePtr).V1).IntVal) != (0)))
}))
_ = eqTuple2_4_2
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Tuple_Tuple)(x_5.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_6.UnsafePtr).V0)
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Data_Data_Tuple_Tuple)(x_5.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_6.UnsafePtr).V1)
}
end_branch_4:
return __t4
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_4_2
}))
})
}()
})
	})
	return cache_ordTuple
}

var cache_eq1Tuple gopurs_runtime.Value
var once_eq1Tuple sync.Once
func Get_eq1Tuple() gopurs_runtime.Value {
	once_eq1Tuple.Do(func() {
		cache_eq1Tuple = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Data_Data_Tuple_Tuple)(x_2.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Data_Data_Tuple_Tuple)(x_2.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_3.UnsafePtr).V1).IntVal) != (0)))
}))
}()
})
	})
	return cache_eq1Tuple
}

var cache_ord1Tuple gopurs_runtime.Value
var once_ord1Tuple sync.Once
func Get_ord1Tuple() gopurs_runtime.Value {
	once_ord1Tuple.Do(func() {
		cache_ord1Tuple = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eq1Tuple1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Data_Data_Tuple_Tuple)(x_3.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Data_Data_Tuple_Tuple)(x_3.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_4.UnsafePtr).V1).IntVal) != (0)))
}))
_ = eq1Tuple1_2_1
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func3(func(dictOrd1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Data_Data_Tuple_Tuple)(x_4.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_5.UnsafePtr).V0)
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_3
} else {

}
}
{
if (v_6_2.Type == 9 && v_6_2.IntVal == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*Data_Data_Tuple_Tuple)(x_4.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_5.UnsafePtr).V1)
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Tuple1_2_1
}))
}()
})
	})
	return cache_ord1Tuple
}

var cache_curry gopurs_runtime.Value
var once_curry sync.Once
func Get_curry() gopurs_runtime.Value {
	once_curry.Do(func() {
		cache_curry = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_curry(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_curry
}

var cache_comonadTuple gopurs_runtime.Value
var once_comonadTuple sync.Once
func Get_comonadTuple() gopurs_runtime.Value {
	once_comonadTuple.Do(func() {
		cache_comonadTuple = gopurs_runtime.RecordDict2("extract", "Extend0", Get_snd(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendTuple()
}))
	})
	return cache_comonadTuple
}

var cache_commutativeRingTuple gopurs_runtime.Value
var once_commutativeRingTuple sync.Once
func Get_commutativeRingTuple() gopurs_runtime.Value {
	once_commutativeRingTuple.Do(func() {
		cache_commutativeRingTuple = gopurs_runtime.Func(func(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
ringTuple1_1_0 := gopurs_runtime.Apply(Get_ringTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple1_1_0
return gopurs_runtime.Func(func(dictCommutativeRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
ringTuple2_3_1 := gopurs_runtime.Apply(ringTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing1_2, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple2_3_1
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringTuple2_3_1
}))
})
}()
})
	})
	return cache_commutativeRingTuple
}

var cache_boundedTuple gopurs_runtime.Value
var once_boundedTuple sync.Once
func Get_boundedTuple() gopurs_runtime.Value {
	once_boundedTuple.Do(func() {
		cache_boundedTuple = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
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
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "eq"), (*Data_Data_Tuple_Tuple)(x_8.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_9.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_5, "eq"), (*Data_Data_Tuple_Tuple)(x_8.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_9.UnsafePtr).V1).IntVal) != (0)))
}))
_ = eqTuple2_8_7
ordTuple2_8_6 := gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_9 gopurs_runtime.Value, y_10 gopurs_runtime.Value) gopurs_runtime.Value {
v_11_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "compare"), (*Data_Data_Tuple_Tuple)(x_9.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_10.UnsafePtr).V0)
_ = v_11_8
var __t9 gopurs_runtime.Value
{
if (v_11_8.Type == 9 && v_11_8.IntVal == 1527465420) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_9
} else {

}
}
{
if (v_11_8.Type == 9 && v_11_8.IntVal == 380165415) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_4, "compare"), (*Data_Data_Tuple_Tuple)(x_9.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_10.UnsafePtr).V1)
}
end_branch_9:
return __t9
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_8_7
}))
_ = ordTuple2_8_6
return gopurs_runtime.RecordDict3("top", "bottom", "Ord0", gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{top_1_0, gopurs_runtime.RecordGet(dictBounded1_5, "top")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{bottom_2_1, gopurs_runtime.RecordGet(dictBounded1_5, "bottom")})}, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple2_8_6
}))
})
}()
})
	})
	return cache_boundedTuple
}

var cache_booleanAlgebraTuple gopurs_runtime.Value
var once_booleanAlgebraTuple sync.Once
func Get_booleanAlgebraTuple() gopurs_runtime.Value {
	once_booleanAlgebraTuple.Do(func() {
		cache_booleanAlgebraTuple = gopurs_runtime.Func(func(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
heytingAlgebraTuple1_1_0 := gopurs_runtime.Apply(Get_heytingAlgebraTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraTuple1_1_0
return gopurs_runtime.Func(func(dictBooleanAlgebra1_2 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraTuple2_3_1 := gopurs_runtime.Apply(heytingAlgebraTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra1_2, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraTuple2_3_1
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraTuple2_3_1
}))
})
}()
})
	})
	return cache_booleanAlgebraTuple
}

var cache_applyTuple gopurs_runtime.Value
var once_applyTuple sync.Once
func Get_applyTuple() gopurs_runtime.Value {
	once_applyTuple.Do(func() {
		cache_applyTuple = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_2.UnsafePtr).V0), gopurs_runtime.Apply((*Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_2.UnsafePtr).V1)})}
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}))
}()
})
	})
	return cache_applyTuple
}

var cache_bindTuple gopurs_runtime.Value
var once_bindTuple sync.Once
func Get_bindTuple() gopurs_runtime.Value {
	once_bindTuple.Do(func() {
		cache_bindTuple = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
applyTuple1_1_0 := gopurs_runtime.Apply(Get_applyTuple(), dictSemigroup_0)
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.Apply(f_3, (*Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
_ = v1_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_4_1.UnsafePtr).V0), (*Data_Data_Tuple_Tuple)(v1_4_1.UnsafePtr).V1})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}))
}()
})
	})
	return cache_bindTuple
}

var cache_applicativeTuple gopurs_runtime.Value
var once_applicativeTuple sync.Once
func Get_applicativeTuple() gopurs_runtime.Value {
	once_applicativeTuple.Do(func() {
		cache_applicativeTuple = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applyTuple1_1_0 := gopurs_runtime.Apply(Get_applyTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}))
}()
})
	})
	return cache_applicativeTuple
}

var cache_monadTuple gopurs_runtime.Value
var once_monadTuple sync.Once
func Get_monadTuple() gopurs_runtime.Value {
	once_monadTuple.Do(func() {
		cache_monadTuple = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applicativeTuple1_1_0 := gopurs_runtime.Apply(Get_applicativeTuple(), dictMonoid_0)
_ = applicativeTuple1_1_0
bindTuple1_2_1 := gopurs_runtime.Apply(Get_bindTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindTuple1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeTuple1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindTuple1_2_1
}))
}()
})
	})
	return cache_monadTuple
}

type Data_Data_Tuple_Tuple struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Data_Tuple_Tuple(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2339352186
}

func Call_uncurry(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V1)
}

func Call_showTuple(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_semigroupTuple(dictSemigroup_0_loop gopurs_runtime.Value, dictSemigroup1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictSemigroup1_1 gopurs_runtime.Value = dictSemigroup1_1_loop
_ = dictSemigroup1_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup1_1, "append"), (*Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
}))
}

func Call_lazyTuple(dictLazy_0_loop gopurs_runtime.Value, dictLazy1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
var dictLazy1_1 gopurs_runtime.Value = dictLazy1_1_loop
_ = dictLazy1_1
return gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Data_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).UnsafePtr).V0
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy1_1, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Data_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).UnsafePtr).V1
}))})}
}))
}

func Call_eqTuple(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Data_Data_Tuple_Tuple)(x_2.UnsafePtr).V0, (*Data_Data_Tuple_Tuple)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Data_Data_Tuple_Tuple)(x_2.UnsafePtr).V1, (*Data_Data_Tuple_Tuple)(y_3.UnsafePtr).V1).IntVal) != (0)))
}))
}

func Call_curry(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Data_Data_Tuple_Tuple{a_1, b_2})})
}


