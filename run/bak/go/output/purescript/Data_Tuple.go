package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Tuple_Tuple gopurs_runtime.Value
var once_Data_Tuple_Tuple sync.Once
func Get_Data_Tuple_Tuple() gopurs_runtime.Value {
	once_Data_Tuple_Tuple.Do(func() {
		cache_Data_Tuple_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, value0, value1})}
})
})
	})
	return cache_Data_Tuple_Tuple
}

var cache_Data_Tuple_uncurry gopurs_runtime.Value
var once_Data_Tuple_uncurry sync.Once
func Get_Data_Tuple_uncurry() gopurs_runtime.Value {
	once_Data_Tuple_uncurry.Do(func() {
		cache_Data_Tuple_uncurry = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_uncurry(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box))
})
	})
	return cache_Data_Tuple_uncurry
}

var cache_Data_Tuple_swap gopurs_runtime.Value
var once_Data_Tuple_swap sync.Once
func Get_Data_Tuple_swap() gopurs_runtime.Value {
	once_Data_Tuple_swap.Do(func() {
		cache_Data_Tuple_swap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_swap(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Data_Tuple_swap
}

var cache_Data_Tuple_snd gopurs_runtime.Value
var once_Data_Tuple_snd sync.Once
func Get_Data_Tuple_snd() gopurs_runtime.Value {
	once_Data_Tuple_snd.Do(func() {
		cache_Data_Tuple_snd = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd
}

var cache_Data_Tuple_showTuple gopurs_runtime.Value
var once_Data_Tuple_showTuple sync.Once
func Get_Data_Tuple_showTuple() gopurs_runtime.Value {
	once_Data_Tuple_showTuple.Do(func() {
		cache_Data_Tuple_showTuple = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_showTuple(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Tuple_showTuple
}

var cache_Data_Tuple_semiringTuple gopurs_runtime.Value
var once_Data_Tuple_semiringTuple sync.Once
func Get_Data_Tuple_semiringTuple() gopurs_runtime.Value {
	once_Data_Tuple_semiringTuple.Do(func() {
		cache_Data_Tuple_semiringTuple = gopurs_runtime.Func2(func(dictSemiring_0_box gopurs_runtime.Value, dictSemiring1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_semiringTuple(dictSemiring_0_box, dictSemiring1_1_box)
})
	})
	return cache_Data_Tuple_semiringTuple
}

var cache_Data_Tuple_semigroupoidTuple gopurs_runtime.Value
var once_Data_Tuple_semigroupoidTuple sync.Once
func Get_Data_Tuple_semigroupoidTuple() gopurs_runtime.Value {
	once_Data_Tuple_semigroupoidTuple.Do(func() {
		cache_Data_Tuple_semigroupoidTuple = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1})}
})
}))
	})
	return cache_Data_Tuple_semigroupoidTuple
}

var cache_Data_Tuple_semigroupTuple gopurs_runtime.Value
var once_Data_Tuple_semigroupTuple sync.Once
func Get_Data_Tuple_semigroupTuple() gopurs_runtime.Value {
	once_Data_Tuple_semigroupTuple.Do(func() {
		cache_Data_Tuple_semigroupTuple = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictSemigroup1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_semigroupTuple(dictSemigroup_0_box, dictSemigroup1_1_box)
})
	})
	return cache_Data_Tuple_semigroupTuple
}

var cache_Data_Tuple_ringTuple gopurs_runtime.Value
var once_Data_Tuple_ringTuple sync.Once
func Get_Data_Tuple_ringTuple() gopurs_runtime.Value {
	once_Data_Tuple_ringTuple.Do(func() {
		cache_Data_Tuple_ringTuple = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_ringTuple(dictRing_0_box)
})
	})
	return cache_Data_Tuple_ringTuple
}

var cache_Data_Tuple_monoidTuple gopurs_runtime.Value
var once_Data_Tuple_monoidTuple sync.Once
func Get_Data_Tuple_monoidTuple() gopurs_runtime.Value {
	once_Data_Tuple_monoidTuple.Do(func() {
		cache_Data_Tuple_monoidTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_monoidTuple(dictMonoid_0_box)
})
	})
	return cache_Data_Tuple_monoidTuple
}

var cache_Data_Tuple_heytingAlgebraTuple gopurs_runtime.Value
var once_Data_Tuple_heytingAlgebraTuple sync.Once
func Get_Data_Tuple_heytingAlgebraTuple() gopurs_runtime.Value {
	once_Data_Tuple_heytingAlgebraTuple.Do(func() {
		cache_Data_Tuple_heytingAlgebraTuple = gopurs_runtime.Func2(func(dictHeytingAlgebra_0_box gopurs_runtime.Value, dictHeytingAlgebra1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_heytingAlgebraTuple(dictHeytingAlgebra_0_box, dictHeytingAlgebra1_1_box)
})
	})
	return cache_Data_Tuple_heytingAlgebraTuple
}

var cache_Data_Tuple_genericTuple gopurs_runtime.Value
var once_Data_Tuple_genericTuple sync.Once
func Get_Data_Tuple_genericTuple() gopurs_runtime.Value {
	once_Data_Tuple_genericTuple.Do(func() {
		cache_Data_Tuple_genericTuple = gopurs_runtime.RecordDict2("from", "to", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, (*Constructor_Data_Tuple_Tuple)(x_0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(x_0.UnsafePtr).V1})}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Generic_Rep_Product)(x_0.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(x_0.UnsafePtr).V1})}
}))
	})
	return cache_Data_Tuple_genericTuple
}

var cache_Data_Tuple_functorTuple gopurs_runtime.Value
var once_Data_Tuple_functorTuple sync.Once
func Get_Data_Tuple_functorTuple() gopurs_runtime.Value {
	once_Data_Tuple_functorTuple.Do(func() {
		cache_Data_Tuple_functorTuple = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_Tuple_functorTuple
}

var cache_Data_Tuple_invariantTuple gopurs_runtime.Value
var once_Data_Tuple_invariantTuple sync.Once
func Get_Data_Tuple_invariantTuple() gopurs_runtime.Value {
	once_Data_Tuple_invariantTuple.Do(func() {
		cache_Data_Tuple_invariantTuple = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), f_0)
})
}))
	})
	return cache_Data_Tuple_invariantTuple
}

var cache_Data_Tuple_fst gopurs_runtime.Value
var once_Data_Tuple_fst sync.Once
func Get_Data_Tuple_fst() gopurs_runtime.Value {
	once_Data_Tuple_fst.Do(func() {
		cache_Data_Tuple_fst = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst
}

var cache_Data_Tuple_lazyTuple gopurs_runtime.Value
var once_Data_Tuple_lazyTuple sync.Once
func Get_Data_Tuple_lazyTuple() gopurs_runtime.Value {
	once_Data_Tuple_lazyTuple.Do(func() {
		cache_Data_Tuple_lazyTuple = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, dictLazy1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_lazyTuple(dictLazy_0_box, dictLazy1_1_box)
})
	})
	return cache_Data_Tuple_lazyTuple
}

var cache_Data_Tuple_extendTuple gopurs_runtime.Value
var once_Data_Tuple_extendTuple sync.Once
func Get_Data_Tuple_extendTuple() gopurs_runtime.Value {
	once_Data_Tuple_extendTuple.Do(func() {
		cache_Data_Tuple_extendTuple = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1))})})}
})
}))
	})
	return cache_Data_Tuple_extendTuple
}

var cache_Data_Tuple_eqTuple gopurs_runtime.Value
var once_Data_Tuple_eqTuple sync.Once
func Get_Data_Tuple_eqTuple() gopurs_runtime.Value {
	once_Data_Tuple_eqTuple.Do(func() {
		cache_Data_Tuple_eqTuple = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_eqTuple(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Tuple_eqTuple
}

var cache_Data_Tuple_ordTuple gopurs_runtime.Value
var once_Data_Tuple_ordTuple sync.Once
func Get_Data_Tuple_ordTuple() gopurs_runtime.Value {
	once_Data_Tuple_ordTuple.Do(func() {
		cache_Data_Tuple_ordTuple = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_ordTuple(dictOrd_0_box)
})
	})
	return cache_Data_Tuple_ordTuple
}

var cache_Data_Tuple_eq1Tuple gopurs_runtime.Value
var once_Data_Tuple_eq1Tuple sync.Once
func Get_Data_Tuple_eq1Tuple() gopurs_runtime.Value {
	once_Data_Tuple_eq1Tuple.Do(func() {
		cache_Data_Tuple_eq1Tuple = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_eq1Tuple(dictEq_0_box)
})
	})
	return cache_Data_Tuple_eq1Tuple
}

var cache_Data_Tuple_ord1Tuple gopurs_runtime.Value
var once_Data_Tuple_ord1Tuple sync.Once
func Get_Data_Tuple_ord1Tuple() gopurs_runtime.Value {
	once_Data_Tuple_ord1Tuple.Do(func() {
		cache_Data_Tuple_ord1Tuple = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_ord1Tuple(dictOrd_0_box)
})
	})
	return cache_Data_Tuple_ord1Tuple
}

var cache_Data_Tuple_curry gopurs_runtime.Value
var once_Data_Tuple_curry sync.Once
func Get_Data_Tuple_curry() gopurs_runtime.Value {
	once_Data_Tuple_curry.Do(func() {
		cache_Data_Tuple_curry = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_curry(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_Data_Tuple_curry
}

var cache_Data_Tuple_comonadTuple gopurs_runtime.Value
var once_Data_Tuple_comonadTuple sync.Once
func Get_Data_Tuple_comonadTuple() gopurs_runtime.Value {
	once_Data_Tuple_comonadTuple.Do(func() {
		cache_Data_Tuple_comonadTuple = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_extendTuple()
}), Get_Data_Tuple_snd())
	})
	return cache_Data_Tuple_comonadTuple
}

var cache_Data_Tuple_commutativeRingTuple gopurs_runtime.Value
var once_Data_Tuple_commutativeRingTuple sync.Once
func Get_Data_Tuple_commutativeRingTuple() gopurs_runtime.Value {
	once_Data_Tuple_commutativeRingTuple.Do(func() {
		cache_Data_Tuple_commutativeRingTuple = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_commutativeRingTuple(dictCommutativeRing_0_box)
})
	})
	return cache_Data_Tuple_commutativeRingTuple
}

var cache_Data_Tuple_boundedTuple gopurs_runtime.Value
var once_Data_Tuple_boundedTuple sync.Once
func Get_Data_Tuple_boundedTuple() gopurs_runtime.Value {
	once_Data_Tuple_boundedTuple.Do(func() {
		cache_Data_Tuple_boundedTuple = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_boundedTuple(dictBounded_0_box)
})
	})
	return cache_Data_Tuple_boundedTuple
}

var cache_Data_Tuple_booleanAlgebraTuple gopurs_runtime.Value
var once_Data_Tuple_booleanAlgebraTuple sync.Once
func Get_Data_Tuple_booleanAlgebraTuple() gopurs_runtime.Value {
	once_Data_Tuple_booleanAlgebraTuple.Do(func() {
		cache_Data_Tuple_booleanAlgebraTuple = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_booleanAlgebraTuple(dictBooleanAlgebra_0_box)
})
	})
	return cache_Data_Tuple_booleanAlgebraTuple
}

var cache_Data_Tuple_applyTuple gopurs_runtime.Value
var once_Data_Tuple_applyTuple sync.Once
func Get_Data_Tuple_applyTuple() gopurs_runtime.Value {
	once_Data_Tuple_applyTuple.Do(func() {
		cache_Data_Tuple_applyTuple = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_applyTuple(dictSemigroup_0_box)
})
	})
	return cache_Data_Tuple_applyTuple
}

var cache_Data_Tuple_bindTuple gopurs_runtime.Value
var once_Data_Tuple_bindTuple sync.Once
func Get_Data_Tuple_bindTuple() gopurs_runtime.Value {
	once_Data_Tuple_bindTuple.Do(func() {
		cache_Data_Tuple_bindTuple = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_bindTuple(dictSemigroup_0_box)
})
	})
	return cache_Data_Tuple_bindTuple
}

var cache_Data_Tuple_applicativeTuple gopurs_runtime.Value
var once_Data_Tuple_applicativeTuple sync.Once
func Get_Data_Tuple_applicativeTuple() gopurs_runtime.Value {
	once_Data_Tuple_applicativeTuple.Do(func() {
		cache_Data_Tuple_applicativeTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_applicativeTuple(dictMonoid_0_box)
})
	})
	return cache_Data_Tuple_applicativeTuple
}

var cache_Data_Tuple_monadTuple gopurs_runtime.Value
var once_Data_Tuple_monadTuple sync.Once
func Get_Data_Tuple_monadTuple() gopurs_runtime.Value {
	once_Data_Tuple_monadTuple.Do(func() {
		cache_Data_Tuple_monadTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_monadTuple(dictMonoid_0_box)
})
	})
	return cache_Data_Tuple_monadTuple
}

var cache_Data_Tuple_comonadTuple__409085367 gopurs_runtime.Value
var once_Data_Tuple_comonadTuple__409085367 sync.Once
func Get_Data_Tuple_comonadTuple__409085367() gopurs_runtime.Value {
	once_Data_Tuple_comonadTuple__409085367.Do(func() {
		cache_Data_Tuple_comonadTuple__409085367 = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_extendTuple()
}), Get_Data_Tuple_snd())
	})
	return cache_Data_Tuple_comonadTuple__409085367
}

var cache_Data_Tuple_curry__2567276513 gopurs_runtime.Value
var once_Data_Tuple_curry__2567276513 sync.Once
func Get_Data_Tuple_curry__2567276513() gopurs_runtime.Value {
	once_Data_Tuple_curry__2567276513.Do(func() {
		cache_Data_Tuple_curry__2567276513 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_curry__2567276513(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_Data_Tuple_curry__2567276513
}

var cache_Data_Tuple_curry__4081152289 gopurs_runtime.Value
var once_Data_Tuple_curry__4081152289 sync.Once
func Get_Data_Tuple_curry__4081152289() gopurs_runtime.Value {
	once_Data_Tuple_curry__4081152289.Do(func() {
		cache_Data_Tuple_curry__4081152289 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_curry__4081152289(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_Data_Tuple_curry__4081152289
}

var cache_Data_Tuple_extendTuple__754007552 gopurs_runtime.Value
var once_Data_Tuple_extendTuple__754007552 sync.Once
func Get_Data_Tuple_extendTuple__754007552() gopurs_runtime.Value {
	once_Data_Tuple_extendTuple__754007552.Do(func() {
		cache_Data_Tuple_extendTuple__754007552 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1))})})}
})
}))
	})
	return cache_Data_Tuple_extendTuple__754007552
}

var cache_Data_Tuple_fst__3478736243 gopurs_runtime.Value
var once_Data_Tuple_fst__3478736243 sync.Once
func Get_Data_Tuple_fst__3478736243() gopurs_runtime.Value {
	once_Data_Tuple_fst__3478736243.Do(func() {
		cache_Data_Tuple_fst__3478736243 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Tuple_fst__3478736243(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))
})
	})
	return cache_Data_Tuple_fst__3478736243
}

var cache_Data_Tuple_fst__549384412 gopurs_runtime.Value
var once_Data_Tuple_fst__549384412 sync.Once
func Get_Data_Tuple_fst__549384412() gopurs_runtime.Value {
	once_Data_Tuple_fst__549384412.Do(func() {
		cache_Data_Tuple_fst__549384412 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Tuple_fst__549384412(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))
})
	})
	return cache_Data_Tuple_fst__549384412
}

var cache_Data_Tuple_fst__20422131 gopurs_runtime.Value
var once_Data_Tuple_fst__20422131 sync.Once
func Get_Data_Tuple_fst__20422131() gopurs_runtime.Value {
	once_Data_Tuple_fst__20422131.Do(func() {
		cache_Data_Tuple_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst__20422131(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst__20422131
}

var cache_Data_Tuple_fst__2554656696 gopurs_runtime.Value
var once_Data_Tuple_fst__2554656696 sync.Once
func Get_Data_Tuple_fst__2554656696() gopurs_runtime.Value {
	once_Data_Tuple_fst__2554656696.Do(func() {
		cache_Data_Tuple_fst__2554656696 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst__2554656696(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst__2554656696
}

var cache_Data_Tuple_fst__4285080947 gopurs_runtime.Value
var once_Data_Tuple_fst__4285080947 sync.Once
func Get_Data_Tuple_fst__4285080947() gopurs_runtime.Value {
	once_Data_Tuple_fst__4285080947.Do(func() {
		cache_Data_Tuple_fst__4285080947 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst__4285080947(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst__4285080947
}

var cache_Data_Tuple_fst__395594805 gopurs_runtime.Value
var once_Data_Tuple_fst__395594805 sync.Once
func Get_Data_Tuple_fst__395594805() gopurs_runtime.Value {
	once_Data_Tuple_fst__395594805.Do(func() {
		cache_Data_Tuple_fst__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst__395594805(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst__395594805
}

var cache_Data_Tuple_fst__82526164 gopurs_runtime.Value
var once_Data_Tuple_fst__82526164 sync.Once
func Get_Data_Tuple_fst__82526164() gopurs_runtime.Value {
	once_Data_Tuple_fst__82526164.Do(func() {
		cache_Data_Tuple_fst__82526164 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst__82526164(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst__82526164
}

var cache_Data_Tuple_fst__2754781306 gopurs_runtime.Value
var once_Data_Tuple_fst__2754781306 sync.Once
func Get_Data_Tuple_fst__2754781306() gopurs_runtime.Value {
	once_Data_Tuple_fst__2754781306.Do(func() {
		cache_Data_Tuple_fst__2754781306 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_fst__2754781306(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_fst__2754781306
}

var cache_Data_Tuple_functorTuple__2544689875 gopurs_runtime.Value
var once_Data_Tuple_functorTuple__2544689875 sync.Once
func Get_Data_Tuple_functorTuple__2544689875() gopurs_runtime.Value {
	once_Data_Tuple_functorTuple__2544689875.Do(func() {
		cache_Data_Tuple_functorTuple__2544689875 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_Tuple_functorTuple__2544689875
}

var cache_Data_Tuple_functorTuple__2249620049 gopurs_runtime.Value
var once_Data_Tuple_functorTuple__2249620049 sync.Once
func Get_Data_Tuple_functorTuple__2249620049() gopurs_runtime.Value {
	once_Data_Tuple_functorTuple__2249620049.Do(func() {
		cache_Data_Tuple_functorTuple__2249620049 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_Tuple_functorTuple__2249620049
}

var cache_Data_Tuple_snd__1234761462 gopurs_runtime.Value
var once_Data_Tuple_snd__1234761462 sync.Once
func Get_Data_Tuple_snd__1234761462() gopurs_runtime.Value {
	once_Data_Tuple_snd__1234761462.Do(func() {
		cache_Data_Tuple_snd__1234761462 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__1234761462(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__1234761462
}

var cache_Data_Tuple_snd__4069939766 gopurs_runtime.Value
var once_Data_Tuple_snd__4069939766 sync.Once
func Get_Data_Tuple_snd__4069939766() gopurs_runtime.Value {
	once_Data_Tuple_snd__4069939766.Do(func() {
		cache_Data_Tuple_snd__4069939766 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__4069939766(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__4069939766
}

var cache_Data_Tuple_snd__3058387254 gopurs_runtime.Value
var once_Data_Tuple_snd__3058387254 sync.Once
func Get_Data_Tuple_snd__3058387254() gopurs_runtime.Value {
	once_Data_Tuple_snd__3058387254.Do(func() {
		cache_Data_Tuple_snd__3058387254 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_snd__3058387254(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Data_Tuple_snd__3058387254
}

var cache_Data_Tuple_snd__2618926102 gopurs_runtime.Value
var once_Data_Tuple_snd__2618926102 sync.Once
func Get_Data_Tuple_snd__2618926102() gopurs_runtime.Value {
	once_Data_Tuple_snd__2618926102.Do(func() {
		cache_Data_Tuple_snd__2618926102 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__2618926102(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__2618926102
}

var cache_Data_Tuple_snd__21214742 gopurs_runtime.Value
var once_Data_Tuple_snd__21214742 sync.Once
func Get_Data_Tuple_snd__21214742() gopurs_runtime.Value {
	once_Data_Tuple_snd__21214742.Do(func() {
		cache_Data_Tuple_snd__21214742 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__21214742(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__21214742
}

var cache_Data_Tuple_snd__20422131 gopurs_runtime.Value
var once_Data_Tuple_snd__20422131 sync.Once
func Get_Data_Tuple_snd__20422131() gopurs_runtime.Value {
	once_Data_Tuple_snd__20422131.Do(func() {
		cache_Data_Tuple_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__20422131(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__20422131
}

var cache_Data_Tuple_snd__2198421043 gopurs_runtime.Value
var once_Data_Tuple_snd__2198421043 sync.Once
func Get_Data_Tuple_snd__2198421043() gopurs_runtime.Value {
	once_Data_Tuple_snd__2198421043.Do(func() {
		cache_Data_Tuple_snd__2198421043 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__2198421043(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__2198421043
}

var cache_Data_Tuple_snd__4038973427 gopurs_runtime.Value
var once_Data_Tuple_snd__4038973427 sync.Once
func Get_Data_Tuple_snd__4038973427() gopurs_runtime.Value {
	once_Data_Tuple_snd__4038973427.Do(func() {
		cache_Data_Tuple_snd__4038973427 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__4038973427(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__4038973427
}

var cache_Data_Tuple_snd__2019004820 gopurs_runtime.Value
var once_Data_Tuple_snd__2019004820 sync.Once
func Get_Data_Tuple_snd__2019004820() gopurs_runtime.Value {
	once_Data_Tuple_snd__2019004820.Do(func() {
		cache_Data_Tuple_snd__2019004820 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Tuple_snd__2019004820(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))
})
	})
	return cache_Data_Tuple_snd__2019004820
}

var cache_Data_Tuple_snd__4227940231 gopurs_runtime.Value
var once_Data_Tuple_snd__4227940231 sync.Once
func Get_Data_Tuple_snd__4227940231() gopurs_runtime.Value {
	once_Data_Tuple_snd__4227940231.Do(func() {
		cache_Data_Tuple_snd__4227940231 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__4227940231(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__4227940231
}

var cache_Data_Tuple_snd__3388842183 gopurs_runtime.Value
var once_Data_Tuple_snd__3388842183 sync.Once
func Get_Data_Tuple_snd__3388842183() gopurs_runtime.Value {
	once_Data_Tuple_snd__3388842183.Do(func() {
		cache_Data_Tuple_snd__3388842183 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__3388842183(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__3388842183
}

var cache_Data_Tuple_snd__1694562576 gopurs_runtime.Value
var once_Data_Tuple_snd__1694562576 sync.Once
func Get_Data_Tuple_snd__1694562576() gopurs_runtime.Value {
	once_Data_Tuple_snd__1694562576.Do(func() {
		cache_Data_Tuple_snd__1694562576 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__1694562576(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__1694562576
}

var cache_Data_Tuple_snd__395594805 gopurs_runtime.Value
var once_Data_Tuple_snd__395594805 sync.Once
func Get_Data_Tuple_snd__395594805() gopurs_runtime.Value {
	once_Data_Tuple_snd__395594805.Do(func() {
		cache_Data_Tuple_snd__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_snd__395594805(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box))
})
	})
	return cache_Data_Tuple_snd__395594805
}

var cache_Data_Tuple_swap__572690077 gopurs_runtime.Value
var once_Data_Tuple_swap__572690077 sync.Once
func Get_Data_Tuple_swap__572690077() gopurs_runtime.Value {
	once_Data_Tuple_swap__572690077.Do(func() {
		cache_Data_Tuple_swap__572690077 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_swap__572690077(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Data_Tuple_swap__572690077
}

var cache_Data_Tuple_swap__249254301 gopurs_runtime.Value
var once_Data_Tuple_swap__249254301 sync.Once
func Get_Data_Tuple_swap__249254301() gopurs_runtime.Value {
	once_Data_Tuple_swap__249254301.Do(func() {
		cache_Data_Tuple_swap__249254301 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_swap__249254301(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Data_Tuple_swap__249254301
}

var cache_Data_Tuple_swap__1502889949 gopurs_runtime.Value
var once_Data_Tuple_swap__1502889949 sync.Once
func Get_Data_Tuple_swap__1502889949() gopurs_runtime.Value {
	once_Data_Tuple_swap__1502889949.Do(func() {
		cache_Data_Tuple_swap__1502889949 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_swap__1502889949(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Data_Tuple_swap__1502889949
}

var cache_Data_Tuple_uncurry__3533477633 gopurs_runtime.Value
var once_Data_Tuple_uncurry__3533477633 sync.Once
func Get_Data_Tuple_uncurry__3533477633() gopurs_runtime.Value {
	once_Data_Tuple_uncurry__3533477633.Do(func() {
		cache_Data_Tuple_uncurry__3533477633 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_uncurry__3533477633(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box))
})
	})
	return cache_Data_Tuple_uncurry__3533477633
}

var cache_Data_Tuple_uncurry__2421405441 gopurs_runtime.Value
var once_Data_Tuple_uncurry__2421405441 sync.Once
func Get_Data_Tuple_uncurry__2421405441() gopurs_runtime.Value {
	once_Data_Tuple_uncurry__2421405441.Do(func() {
		cache_Data_Tuple_uncurry__2421405441 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_uncurry__2421405441(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Data_Tuple_uncurry__2421405441
}

var cache_Data_Tuple_uncurry__1489344417 gopurs_runtime.Value
var once_Data_Tuple_uncurry__1489344417 sync.Once
func Get_Data_Tuple_uncurry__1489344417() gopurs_runtime.Value {
	once_Data_Tuple_uncurry__1489344417.Do(func() {
		cache_Data_Tuple_uncurry__1489344417 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Tuple_uncurry__1489344417(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box))
})
	})
	return cache_Data_Tuple_uncurry__1489344417
}

var cache_Data_Tuple_uncurry__601611969 gopurs_runtime.Value
var once_Data_Tuple_uncurry__601611969 sync.Once
func Get_Data_Tuple_uncurry__601611969() gopurs_runtime.Value {
	once_Data_Tuple_uncurry__601611969.Do(func() {
		cache_Data_Tuple_uncurry__601611969 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_Tuple_uncurry__601611969(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Data_Tuple_uncurry__601611969
}

type Constructor_Data_Tuple_Tuple struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_Data_Tuple_uncurry(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1)
}

func Call_Data_Tuple_swap(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return &Constructor_Data_Tuple_Tuple{1, (v_0).V1, (v_0).V0}
}

func Call_Data_Tuple_snd(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_showTuple(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Tuple ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_Data_Tuple_semiringTuple(dictSemiring_0_loop gopurs_runtime.Value, dictSemiring1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
var dictSemiring1_1 gopurs_runtime.Value = dictSemiring1_1_loop
_ = dictSemiring1_1
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_1, "add"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_1, "mul"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring1_1, "one")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictSemiring_0, "zero"), gopurs_runtime.RecordGet(dictSemiring1_1, "zero")})})
}

func Call_Data_Tuple_semigroupTuple(dictSemigroup_0_loop gopurs_runtime.Value, dictSemigroup1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictSemigroup1_1 gopurs_runtime.Value = dictSemigroup1_1_loop
_ = dictSemigroup1_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup1_1, "append"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
})
}))
}

func Call_Data_Tuple_ringTuple(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing1_2, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): semiringTuple2_3_1 -> gopurs_runtime.Value
semiringTuple2_3_1 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "add"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "add"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mul"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "mul"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(__local_var_1_0, "one"), gopurs_runtime.RecordGet(__local_var_3_2, "one")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(__local_var_1_0, "zero"), gopurs_runtime.RecordGet(__local_var_3_2, "zero")})})
_ = semiringTuple2_3_1
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringTuple2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing1_2, "sub"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
})
}))
})
}

func Call_Data_Tuple_monoidTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictMonoid1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): semigroupTuple2_3_1 -> gopurs_runtime.Value
semigroupTuple2_3_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "append"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)})}
})
}))
_ = semigroupTuple2_3_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupTuple2_3_1
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.RecordGet(dictMonoid1_2, "mempty")})})
})
}

func Call_Data_Tuple_heytingAlgebraTuple(dictHeytingAlgebra_0_loop gopurs_runtime.Value, dictHeytingAlgebra1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
var dictHeytingAlgebra1_1 gopurs_runtime.Value = dictHeytingAlgebra1_1_loop
_ = dictHeytingAlgebra1_1
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "conj"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "disj"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "ff")})}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "implies"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "not"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.RecordGet(dictHeytingAlgebra1_1, "tt")})}})
}

func Call_Data_Tuple_fst(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_lazyTuple(dictLazy_0_loop gopurs_runtime.Value, dictLazy1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
var dictLazy1_1 gopurs_runtime.Value = dictLazy1_1_loop
_ = dictLazy1_1
return gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()).UnsafePtr).V0
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy1_1, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()).UnsafePtr).V1
}))})}
}))
}

func Call_Data_Tuple_eqTuple(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_3.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_Tuple_ordTuple(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eqTuple2_3_1 -> gopurs_runtime.Value
eqTuple2_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Tuple_Tuple)(x_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Tuple_Tuple)(x_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_5.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqTuple2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_3 -> gopurs_runtime.Value
v_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Tuple_Tuple)(x_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_5.UnsafePtr).V0)
_ = v_6_3
var __t4 uint32
{
if (uint32(v_6_3.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(v_6_3.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Tuple_Tuple)(x_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_5.UnsafePtr).V1).IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
}))
})
}

func Call_Data_Tuple_eq1Tuple(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_3.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
}

func Call_Data_Tuple_ord1Tuple(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eq1Tuple1_1_0 -> gopurs_runtime.Value
eq1Tuple1_1_0 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Tuple_Tuple)(x_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Tuple_Tuple)(x_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_4.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
_ = eq1Tuple1_1_0
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Tuple1_1_0
}), gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 -> gopurs_runtime.Value
v_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Tuple_Tuple)(x_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_4.UnsafePtr).V0)
_ = v_5_2
var __t3 uint32
{
if (uint32(v_5_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_5_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Tuple_Tuple)(x_3.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_4.UnsafePtr).V1).IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
})
}))
}

func Call_Data_Tuple_curry(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_1, b_2})})
}

func Call_Data_Tuple_commutativeRingTuple(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
// TAST (Let): ringTuple1_1_0 -> gopurs_runtime.Value
ringTuple1_1_0 := Call_Data_Tuple_ringTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple1_1_0
return gopurs_runtime.Func(func(dictCommutativeRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ringTuple2_3_1 -> gopurs_runtime.Value
ringTuple2_3_1 := gopurs_runtime.Apply(ringTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing1_2, "Ring0"), gopurs_runtime.Value{}))
_ = ringTuple2_3_1
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringTuple2_3_1
}))
})
}

func Call_Data_Tuple_boundedTuple(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBounded1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_3, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): eqTuple2_5_4 -> gopurs_runtime.Value
eqTuple2_5_4 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*Constructor_Data_Tuple_Tuple)(x_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "eq"), (*Constructor_Data_Tuple_Tuple)(x_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_7.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqTuple2_5_4
// TAST (Let): ordTuple2_4_2 -> gopurs_runtime.Value
ordTuple2_4_2 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_5_4
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_8_6 -> gopurs_runtime.Value
v_8_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), (*Constructor_Data_Tuple_Tuple)(x_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(y_7.UnsafePtr).V0)
_ = v_8_6
var __t7 uint32
{
if (uint32(v_8_6.IntVal) == 1527465420) {
__t7 = 1527465420
goto end_branch_7
} else {

}
}
{
if (uint32(v_8_6.IntVal) == 380165415) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
__t7 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "compare"), (*Constructor_Data_Tuple_Tuple)(x_6.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_7.UnsafePtr).V1).IntVal)
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})
}))
_ = ordTuple2_4_2
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple2_4_2
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictBounded_0, "bottom"), gopurs_runtime.RecordGet(dictBounded1_3, "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(dictBounded_0, "top"), gopurs_runtime.RecordGet(dictBounded1_3, "top")})})
})
}

func Call_Data_Tuple_booleanAlgebraTuple(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
// TAST (Let): heytingAlgebraTuple1_1_0 -> gopurs_runtime.Value
heytingAlgebraTuple1_1_0 := gopurs_runtime.Apply(Get_Data_Tuple_heytingAlgebraTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraTuple1_1_0
return gopurs_runtime.Func(func(dictBooleanAlgebra1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): heytingAlgebraTuple2_3_1 -> gopurs_runtime.Value
heytingAlgebraTuple2_3_1 := gopurs_runtime.Apply(heytingAlgebraTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra1_2, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraTuple2_3_1
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraTuple2_3_1
}))
})
}

func Call_Data_Tuple_applyTuple(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_2.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(v1_2.UnsafePtr).V1)})}
})
}))
}

func Call_Data_Tuple_bindTuple(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): applyTuple1_1_0 -> gopurs_runtime.Value
applyTuple1_1_0 := Call_Data_Tuple_applyTuple(dictSemigroup_0)
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_1 -> *Constructor_Data_Tuple_Tuple
v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1))
_ = v1_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, (v1_4_1).V0), (v1_4_1).V1})}
})
}))
}

func Call_Data_Tuple_applicativeTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): applyTuple1_1_0 -> gopurs_runtime.Value
applyTuple1_1_0 := Call_Data_Tuple_applyTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyTuple1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Call_Data_Tuple_monadTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): applicativeTuple1_1_0 -> gopurs_runtime.Value
applicativeTuple1_1_0 := Call_Data_Tuple_applicativeTuple(dictMonoid_0)
_ = applicativeTuple1_1_0
// TAST (Let): bindTuple1_2_1 -> gopurs_runtime.Value
bindTuple1_2_1 := Call_Data_Tuple_bindTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = bindTuple1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeTuple1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindTuple1_2_1
}))
}

func Call_Data_Tuple_curry__2567276513(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_1, b_2})})
}

func Call_Data_Tuple_curry__4081152289(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_1, b_2})})
}

func Call_Data_Tuple_fst__3478736243(v_0_loop *Constructor_Data_Tuple_Tuple) int64 {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0.IntVal
}

func Call_Data_Tuple_fst__549384412(v_0_loop *Constructor_Data_Tuple_Tuple) float64 {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0.FloatVal()
}

func Call_Data_Tuple_fst__20422131(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_fst__2554656696(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_fst__4285080947(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_fst__395594805(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_fst__82526164(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_fst__2754781306(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Tuple_snd__1234761462(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__4069939766(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__3058387254(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)
}

func Call_Data_Tuple_snd__2618926102(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__21214742(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__20422131(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__2198421043(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__4038973427(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__2019004820(v_0_loop *Constructor_Data_Tuple_Tuple) float64 {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1.FloatVal()
}

func Call_Data_Tuple_snd__4227940231(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__3388842183(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__1694562576(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_snd__395594805(v_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Tuple_swap__572690077(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return &Constructor_Data_Tuple_Tuple{1, (v_0).V1, (v_0).V0}
}

func Call_Data_Tuple_swap__249254301(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return &Constructor_Data_Tuple_Tuple{1, (v_0).V1, (v_0).V0}
}

func Call_Data_Tuple_swap__1502889949(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return &Constructor_Data_Tuple_Tuple{1, (v_0).V1, (v_0).V0}
}

func Call_Data_Tuple_uncurry__3533477633(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1)
}

func Call_Data_Tuple_uncurry__2421405441(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Control_Monad_RWS_Trans_RWSResult {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1))
}

func Call_Data_Tuple_uncurry__1489344417(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1)
}

func Call_Data_Tuple_uncurry__601611969(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1))
}


