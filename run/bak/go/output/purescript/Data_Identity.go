package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
		cache_Data_Identity_newtypeIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
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
		cache_Data_Identity_functorIdentity = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Identity_functorIdentity
}

var cache_Data_Identity_invariantIdentity gopurs_runtime.Value
var once_Data_Identity_invariantIdentity sync.Once
func Get_Data_Identity_invariantIdentity() gopurs_runtime.Value {
	once_Data_Identity_invariantIdentity.Do(func() {
		cache_Data_Identity_invariantIdentity = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Invariant_Invariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_2)
})
})
})})}
	})
	return cache_Data_Identity_invariantIdentity
}

var cache_Data_Identity_extendIdentity gopurs_runtime.Value
var once_Data_Identity_extendIdentity sync.Once
func Get_Data_Identity_extendIdentity() gopurs_runtime.Value {
	once_Data_Identity_extendIdentity.Do(func() {
		cache_Data_Identity_extendIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
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
		cache_Data_Identity_eq1Identity = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
})})}
	})
	return cache_Data_Identity_eq1Identity
}

var cache_Data_Identity_ord1Identity gopurs_runtime.Value
var once_Data_Identity_ord1Identity sync.Once
func Get_Data_Identity_ord1Identity() gopurs_runtime.Value {
	once_Data_Identity_ord1Identity.Do(func() {
		cache_Data_Identity_ord1Identity = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Identity_eq1Identity()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
})})}
	})
	return cache_Data_Identity_ord1Identity
}

var cache_Data_Identity_comonadIdentity gopurs_runtime.Value
var once_Data_Identity_comonadIdentity sync.Once
func Get_Data_Identity_comonadIdentity() gopurs_runtime.Value {
	once_Data_Identity_comonadIdentity.Do(func() {
		cache_Data_Identity_comonadIdentity = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](Get_Data_Identity_extendIdentity()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})})}
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
		cache_Data_Identity_applyIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
})})}
	})
	return cache_Data_Identity_applyIdentity
}

var cache_Data_Identity_bindIdentity gopurs_runtime.Value
var once_Data_Identity_bindIdentity sync.Once
func Get_Data_Identity_bindIdentity() gopurs_runtime.Value {
	once_Data_Identity_bindIdentity.Do(func() {
		cache_Data_Identity_bindIdentity = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Identity_applyIdentity()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
})})}
	})
	return cache_Data_Identity_bindIdentity
}

var cache_Data_Identity_applicativeIdentity gopurs_runtime.Value
var once_Data_Identity_applicativeIdentity sync.Once
func Get_Data_Identity_applicativeIdentity() gopurs_runtime.Value {
	once_Data_Identity_applicativeIdentity.Do(func() {
		cache_Data_Identity_applicativeIdentity = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Identity_applyIdentity()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Identity_applicativeIdentity
}

var cache_Data_Identity_monadIdentity gopurs_runtime.Value
var once_Data_Identity_monadIdentity sync.Once
func Get_Data_Identity_monadIdentity() gopurs_runtime.Value {
	once_Data_Identity_monadIdentity.Do(func() {
		cache_Data_Identity_monadIdentity = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Identity_applicativeIdentity()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Identity_bindIdentity()))}
})})}
	})
	return cache_Data_Identity_monadIdentity
}

var cache_Data_Identity_altIdentity gopurs_runtime.Value
var once_Data_Identity_altIdentity sync.Once
func Get_Data_Identity_altIdentity() gopurs_runtime.Value {
	once_Data_Identity_altIdentity.Do(func() {
		cache_Data_Identity_altIdentity = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
})})}
	})
	return cache_Data_Identity_altIdentity
}

var cache_Data_Identity_applicativeIdentity__995286821 gopurs_runtime.Value
var once_Data_Identity_applicativeIdentity__995286821 sync.Once
func Get_Data_Identity_applicativeIdentity__995286821() gopurs_runtime.Value {
	once_Data_Identity_applicativeIdentity__995286821.Do(func() {
		cache_Data_Identity_applicativeIdentity__995286821 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Identity_applyIdentity()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Identity_applicativeIdentity__995286821
}

var cache_Data_Identity_applyIdentity__111100453 gopurs_runtime.Value
var once_Data_Identity_applyIdentity__111100453 sync.Once
func Get_Data_Identity_applyIdentity__111100453() gopurs_runtime.Value {
	once_Data_Identity_applyIdentity__111100453.Do(func() {
		cache_Data_Identity_applyIdentity__111100453 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
})})}
	})
	return cache_Data_Identity_applyIdentity__111100453
}

var cache_Data_Identity_bindIdentity__2224694053 gopurs_runtime.Value
var once_Data_Identity_bindIdentity__2224694053 sync.Once
func Get_Data_Identity_bindIdentity__2224694053() gopurs_runtime.Value {
	once_Data_Identity_bindIdentity__2224694053.Do(func() {
		cache_Data_Identity_bindIdentity__2224694053 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Identity_applyIdentity()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
})})}
	})
	return cache_Data_Identity_bindIdentity__2224694053
}

var cache_Data_Identity_eq1Identity__294625475 gopurs_runtime.Value
var once_Data_Identity_eq1Identity__294625475 sync.Once
func Get_Data_Identity_eq1Identity__294625475() gopurs_runtime.Value {
	once_Data_Identity_eq1Identity__294625475.Do(func() {
		cache_Data_Identity_eq1Identity__294625475 = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
})})}
	})
	return cache_Data_Identity_eq1Identity__294625475
}

var cache_Data_Identity_extendIdentity__982429765 gopurs_runtime.Value
var once_Data_Identity_extendIdentity__982429765 sync.Once
func Get_Data_Identity_extendIdentity__982429765() gopurs_runtime.Value {
	once_Data_Identity_extendIdentity__982429765.Do(func() {
		cache_Data_Identity_extendIdentity__982429765 = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Identity_extendIdentity__982429765
}

var cache_Data_Identity_functorIdentity__850816530 gopurs_runtime.Value
var once_Data_Identity_functorIdentity__850816530 sync.Once
func Get_Data_Identity_functorIdentity__850816530() gopurs_runtime.Value {
	once_Data_Identity_functorIdentity__850816530.Do(func() {
		cache_Data_Identity_functorIdentity__850816530 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Identity_functorIdentity__850816530
}

var cache_Data_Identity_monadIdentity__2437051429 gopurs_runtime.Value
var once_Data_Identity_monadIdentity__2437051429 sync.Once
func Get_Data_Identity_monadIdentity__2437051429() gopurs_runtime.Value {
	once_Data_Identity_monadIdentity__2437051429.Do(func() {
		cache_Data_Identity_monadIdentity__2437051429 = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Identity_applicativeIdentity()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Identity_bindIdentity()))}
})})}
	})
	return cache_Data_Identity_monadIdentity__2437051429
}

func Call_Data_Identity_Identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Identity_showIdentity(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Identity ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Identity_semiringIdentity(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_0))}
}

func Call_Data_Identity_semigroupIdentity(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0))}
}

func Call_Data_Identity_ringIdentity(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_0))}
}

func Call_Data_Identity_ordIdentity(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0))}
}

func Call_Data_Identity_monoidIdentity(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}
}

func Call_Data_Identity_lazyIdentity(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
return gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_0))}
}

func Call_Data_Identity_heytingAlgebraIdentity(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_0))}
}

func Call_Data_Identity_euclideanRingIdentity(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
return gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dictEuclideanRing_0))}
}

func Call_Data_Identity_eqIdentity(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}

func Call_Data_Identity_commutativeRingIdentity(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing](dictCommutativeRing_0))}
}

func Call_Data_Identity_boundedIdentity(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](dictBounded_0))}
}

func Call_Data_Identity_booleanAlgebraIdentity(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
return gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_BooleanAlgebra_BooleanAlgebra](dictBooleanAlgebra_0))}
}


