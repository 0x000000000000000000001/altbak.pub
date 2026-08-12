package Data_Tuple

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Tuple gopurs_runtime.Value
var once_Tuple sync.Once
func Get_Tuple() gopurs_runtime.Value {
	once_Tuple.Do(func() {
		cache_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
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
return Call_uncurry(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_uncurry
}

var cache_swap gopurs_runtime.Value
var once_swap sync.Once
func Get_swap() gopurs_runtime.Value {
	once_swap.Do(func() {
		cache_swap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_swap(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_swap
}

var cache_snd gopurs_runtime.Value
var once_snd sync.Once
func Get_snd() gopurs_runtime.Value {
	once_snd.Do(func() {
		cache_snd = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
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
		cache_semiringTuple = gopurs_runtime.Func2(func(dictSemiring_0_box gopurs_runtime.Value, dictSemiring1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringTuple(dictSemiring_0_box, dictSemiring1_1_box)
})
	})
	return cache_semiringTuple
}

var cache_semigroupoidTuple gopurs_runtime.Value
var once_semigroupoidTuple sync.Once
func Get_semigroupoidTuple() gopurs_runtime.Value {
	once_semigroupoidTuple.Do(func() {
		cache_semigroupoidTuple = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1})}
})
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
		cache_ringTuple = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringTuple(dictRing_0_box)
})
	})
	return cache_ringTuple
}

var cache_monoidTuple gopurs_runtime.Value
var once_monoidTuple sync.Once
func Get_monoidTuple() gopurs_runtime.Value {
	once_monoidTuple.Do(func() {
		cache_monoidTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidTuple(dictMonoid_0_box)
})
	})
	return cache_monoidTuple
}

var cache_heytingAlgebraTuple gopurs_runtime.Value
var once_heytingAlgebraTuple sync.Once
func Get_heytingAlgebraTuple() gopurs_runtime.Value {
	once_heytingAlgebraTuple.Do(func() {
		cache_heytingAlgebraTuple = gopurs_runtime.Func2(func(dictHeytingAlgebra_0_box gopurs_runtime.Value, dictHeytingAlgebra1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraTuple(dictHeytingAlgebra_0_box, dictHeytingAlgebra1_1_box)
})
	})
	return cache_heytingAlgebraTuple
}

var cache_genericTuple gopurs_runtime.Value
var once_genericTuple sync.Once
func Get_genericTuple() gopurs_runtime.Value {
	once_genericTuple.Do(func() {
		cache_genericTuple = gopurs_runtime.RecordDict2("from", "to", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V1})}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V1})}
}))
	})
	return cache_genericTuple
}

var cache_functorTuple gopurs_runtime.Value
var once_functorTuple sync.Once
func Get_functorTuple() gopurs_runtime.Value {
	once_functorTuple.Do(func() {
		cache_functorTuple = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple
}

var cache_invariantTuple gopurs_runtime.Value
var once_invariantTuple sync.Once
func Get_invariantTuple() gopurs_runtime.Value {
	once_invariantTuple.Do(func() {
		cache_invariantTuple = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorTuple(), "map"), f_0)
})
}))
	})
	return cache_invariantTuple
}

var cache_fst gopurs_runtime.Value
var once_fst sync.Once
func Get_fst() gopurs_runtime.Value {
	once_fst.Do(func() {
		cache_fst = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
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
		cache_extendTuple = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))})})}
})
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
		cache_ordTuple = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordTuple(dictOrd_0_box)
})
	})
	return cache_ordTuple
}

var cache_eq1Tuple gopurs_runtime.Value
var once_eq1Tuple sync.Once
func Get_eq1Tuple() gopurs_runtime.Value {
	once_eq1Tuple.Do(func() {
		cache_eq1Tuple = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Tuple(dictEq_0_box)
})
	})
	return cache_eq1Tuple
}

var cache_ord1Tuple gopurs_runtime.Value
var once_ord1Tuple sync.Once
func Get_ord1Tuple() gopurs_runtime.Value {
	once_ord1Tuple.Do(func() {
		cache_ord1Tuple = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Tuple(dictOrd_0_box)
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
		cache_comonadTuple = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendTuple()
}), Get_snd())
	})
	return cache_comonadTuple
}

var cache_commutativeRingTuple gopurs_runtime.Value
var once_commutativeRingTuple sync.Once
func Get_commutativeRingTuple() gopurs_runtime.Value {
	once_commutativeRingTuple.Do(func() {
		cache_commutativeRingTuple = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_commutativeRingTuple(dictCommutativeRing_0_box)
})
	})
	return cache_commutativeRingTuple
}

var cache_boundedTuple gopurs_runtime.Value
var once_boundedTuple sync.Once
func Get_boundedTuple() gopurs_runtime.Value {
	once_boundedTuple.Do(func() {
		cache_boundedTuple = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedTuple(dictBounded_0_box)
})
	})
	return cache_boundedTuple
}

var cache_booleanAlgebraTuple gopurs_runtime.Value
var once_booleanAlgebraTuple sync.Once
func Get_booleanAlgebraTuple() gopurs_runtime.Value {
	once_booleanAlgebraTuple.Do(func() {
		cache_booleanAlgebraTuple = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_booleanAlgebraTuple(dictBooleanAlgebra_0_box)
})
	})
	return cache_booleanAlgebraTuple
}

var cache_applyTuple gopurs_runtime.Value
var once_applyTuple sync.Once
func Get_applyTuple() gopurs_runtime.Value {
	once_applyTuple.Do(func() {
		cache_applyTuple = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyTuple(dictSemigroup_0_box)
})
	})
	return cache_applyTuple
}

var cache_bindTuple gopurs_runtime.Value
var once_bindTuple sync.Once
func Get_bindTuple() gopurs_runtime.Value {
	once_bindTuple.Do(func() {
		cache_bindTuple = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindTuple(dictSemigroup_0_box)
})
	})
	return cache_bindTuple
}

var cache_applicativeTuple gopurs_runtime.Value
var once_applicativeTuple sync.Once
func Get_applicativeTuple() gopurs_runtime.Value {
	once_applicativeTuple.Do(func() {
		cache_applicativeTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeTuple(dictMonoid_0_box)
})
	})
	return cache_applicativeTuple
}

var cache_monadTuple gopurs_runtime.Value
var once_monadTuple sync.Once
func Get_monadTuple() gopurs_runtime.Value {
	once_monadTuple.Do(func() {
		cache_monadTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTuple(dictMonoid_0_box)
})
	})
	return cache_monadTuple
}

var cache_bottom__338427193 gopurs_runtime.Value
var once_bottom__338427193 sync.Once
func Get_bottom__338427193() gopurs_runtime.Value {
	once_bottom__338427193.Do(func() {
		cache_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__338427193(dict_0_box)
})
	})
	return cache_bottom__338427193
}

var cache_top__338427193 gopurs_runtime.Value
var once_top__338427193 sync.Once
func Get_top__338427193() gopurs_runtime.Value {
	once_top__338427193.Do(func() {
		cache_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__338427193(dict_0_box)
})
	})
	return cache_top__338427193
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
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

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
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

var cache_implies__3472268504 gopurs_runtime.Value
var once_implies__3472268504 sync.Once
func Get_implies__3472268504() gopurs_runtime.Value {
	once_implies__3472268504.Do(func() {
		cache_implies__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_implies__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_implies__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
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

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
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

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
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

var cache_extendTuple__754007552 gopurs_runtime.Value
var once_extendTuple__754007552 sync.Once
func Get_extendTuple__754007552() gopurs_runtime.Value {
	once_extendTuple__754007552.Do(func() {
		cache_extendTuple__754007552 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))})})}
})
}))
	})
	return cache_extendTuple__754007552
}

var cache_fst__395594805 gopurs_runtime.Value
var once_fst__395594805 sync.Once
func Get_fst__395594805() gopurs_runtime.Value {
	once_fst__395594805.Do(func() {
		cache_fst__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__395594805(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__395594805
}

var cache_functorTuple__2249620049 gopurs_runtime.Value
var once_functorTuple__2249620049 sync.Once
func Get_functorTuple__2249620049() gopurs_runtime.Value {
	once_functorTuple__2249620049.Do(func() {
		cache_functorTuple__2249620049 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple__2249620049
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_snd__395594805 gopurs_runtime.Value
var once_snd__395594805 sync.Once
func Get_snd__395594805() gopurs_runtime.Value {
	once_snd__395594805.Do(func() {
		cache_snd__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__395594805(gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__395594805
}

type Constructor_Tuple[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}


func Call_uncurry(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_swap(v_0_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})})
}

func Call_snd(v_0_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_showTuple(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Tuple "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
}))
}

func Call_semiringTuple(dictSemiring_0_loop gopurs_runtime.Value, dictSemiring1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
var dictSemiring1_1 gopurs_runtime.Value = dictSemiring1_1_loop
_ = dictSemiring1_1
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_1, "add"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_1, "mul"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring1_1, "one")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictSemiring_0, "zero"), gopurs_runtime.RecordGet(dictSemiring1_1, "zero")})})
}

func Call_semigroupTuple(dictSemigroup_0_loop gopurs_runtime.Value, dictSemigroup1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictSemigroup1_1 gopurs_runtime.Value = dictSemigroup1_1_loop
_ = dictSemigroup1_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup1_1, "append"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}))
}

func Call_ringTuple(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing1_2, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_3_2
semiringTuple2_3_1 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "add"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "add"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mul"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "mul"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(__local_var_1_0, "one"), gopurs_runtime.RecordGet(__local_var_3_2, "one")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(__local_var_1_0, "zero"), gopurs_runtime.RecordGet(__local_var_3_2, "zero")})})
_ = semiringTuple2_3_1
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringTuple2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing1_2, "sub"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}))
})
}

func Call_monoidTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictMonoid1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_2
semigroupTuple2_3_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}))
_ = semigroupTuple2_3_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupTuple2_3_1
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet(dictMonoid1_2, "mempty")})})
})
}

func Call_heytingAlgebraTuple(dictHeytingAlgebra_0_loop gopurs_runtime.Value, dictHeytingAlgebra1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
var dictHeytingAlgebra1_1 gopurs_runtime.Value = dictHeytingAlgebra1_1_loop
_ = dictHeytingAlgebra1_1
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "conj"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "disj"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "ff")})}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "implies"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "not"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "tt")})}})
}

func Call_fst(v_0_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_lazyTuple(dictLazy_0_loop gopurs_runtime.Value, dictLazy1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
var dictLazy1_1 gopurs_runtime.Value = dictLazy1_1_loop
_ = dictLazy1_1
return gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).UnsafePtr).V0
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy1_1, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).UnsafePtr).V1
}))})}
}))
}

func Call_eqTuple(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0))
})
}))
}

func Call_ordTuple(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqTuple1_1_0 := gopurs_runtime.Apply(Get_eqTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqTuple1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqTuple2_3_1 := gopurs_runtime.Apply(eqTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqTuple2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_2
var __t3 uint32
{
if (uint32(v_6_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_6_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
}))
})
}

func Call_eq1Tuple(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqTuple(dictEq_0, dictEq1_1), "eq")
}))
}

func Call_ord1Tuple(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordTuple1_1_0 := Call_ordTuple(dictOrd_0)
_ = ordTuple1_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
eq1Tuple1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqTuple(__local_var_2_2, dictEq1_3), "eq")
}))
_ = eq1Tuple1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Tuple1_2_1
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordTuple1_1_0, dictOrd1_3), "compare")
}))
}

func Call_curry(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, b_2})})
}

func Call_commutativeRingTuple(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
ringTuple1_1_0 := Call_ringTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple1_1_0
return gopurs_runtime.Func(func(dictCommutativeRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
ringTuple2_3_1 := gopurs_runtime.Apply(ringTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing1_2, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple2_3_1
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringTuple2_3_1
}))
})
}

func Call_boundedTuple(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
ordTuple1_1_0 := Call_ordTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordTuple1_1_0
return gopurs_runtime.Func(func(dictBounded1_2 gopurs_runtime.Value) gopurs_runtime.Value {
ordTuple2_3_1 := gopurs_runtime.Apply(ordTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_2, "Ord0"), gopurs_runtime.Value{}))
_ = ordTuple2_3_1
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple2_3_1
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded_0, "bottom"), gopurs_runtime.RecordGet(dictBounded1_2, "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded_0, "top"), gopurs_runtime.RecordGet(dictBounded1_2, "top")})})
})
}

func Call_booleanAlgebraTuple(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
}

func Call_applyTuple(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V1)})}
})
}))
}

func Call_bindTuple(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
applyTuple1_1_0 := Call_applyTuple(dictSemigroup_0)
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
_ = v1_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V0), (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v1_4_1)}.UnsafePtr).V1})}
})
}))
}

func Call_applicativeTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applyTuple1_1_0 := Call_applyTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}), gopurs_runtime.Apply(Get_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_monadTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
applicativeTuple1_1_0 := Call_applicativeTuple(dictMonoid_0)
_ = applicativeTuple1_1_0
bindTuple1_2_1 := Call_bindTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindTuple1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeTuple1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindTuple1_2_1
}))
}

func Call_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
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

func Call_implies__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_tt__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fst__395594805(v_0_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__395594805(v_0_loop *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}


