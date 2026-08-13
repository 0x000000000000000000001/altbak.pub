package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Multiplicative_Multiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_Multiplicative sync.Once
func Get_Data_Monoid_Multiplicative_Multiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_Multiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_Multiplicative = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_Multiplicative(x_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_Multiplicative
}

var cache_Data_Monoid_Multiplicative_showMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_showMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_showMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_showMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_showMultiplicative = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_showMultiplicative(dictShow_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_showMultiplicative
}

var cache_Data_Monoid_Multiplicative_semigroupMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_semigroupMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_semigroupMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_semigroupMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_semigroupMultiplicative = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_semigroupMultiplicative(dictSemiring_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_semigroupMultiplicative
}

var cache_Data_Monoid_Multiplicative_ordMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_ordMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_ordMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_ordMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_ordMultiplicative = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_ordMultiplicative(dictOrd_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_ordMultiplicative
}

var cache_Data_Monoid_Multiplicative_monoidMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_monoidMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_monoidMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_monoidMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_monoidMultiplicative = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_monoidMultiplicative(dictSemiring_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_monoidMultiplicative
}

var cache_Data_Monoid_Multiplicative_functorMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_functorMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_functorMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_functorMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_functorMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Monoid_Multiplicative_functorMultiplicative
}

var cache_Data_Monoid_Multiplicative_eqMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_eqMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_eqMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_eqMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_eqMultiplicative = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_eqMultiplicative(dictEq_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_eqMultiplicative
}

var cache_Data_Monoid_Multiplicative_eq1Multiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_eq1Multiplicative sync.Once
func Get_Data_Monoid_Multiplicative_eq1Multiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_eq1Multiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_eq1Multiplicative = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
})})}
	})
	return cache_Data_Monoid_Multiplicative_eq1Multiplicative
}

var cache_Data_Monoid_Multiplicative_ord1Multiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_ord1Multiplicative sync.Once
func Get_Data_Monoid_Multiplicative_ord1Multiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_ord1Multiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_ord1Multiplicative = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Monoid_Multiplicative_eq1Multiplicative()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
})})}
	})
	return cache_Data_Monoid_Multiplicative_ord1Multiplicative
}

var cache_Data_Monoid_Multiplicative_boundedMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_boundedMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_boundedMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_boundedMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_boundedMultiplicative = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Multiplicative_boundedMultiplicative(dictBounded_0_box)
})
	})
	return cache_Data_Monoid_Multiplicative_boundedMultiplicative
}

var cache_Data_Monoid_Multiplicative_applyMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_applyMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_applyMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_applyMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_applyMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Monoid_Multiplicative_functorMultiplicative()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
})})}
	})
	return cache_Data_Monoid_Multiplicative_applyMultiplicative
}

var cache_Data_Monoid_Multiplicative_bindMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_bindMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_bindMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_bindMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_bindMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Multiplicative_applyMultiplicative()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
})})}
	})
	return cache_Data_Monoid_Multiplicative_bindMultiplicative
}

var cache_Data_Monoid_Multiplicative_applicativeMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_applicativeMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_applicativeMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_applicativeMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_applicativeMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Multiplicative_applyMultiplicative()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Monoid_Multiplicative_applicativeMultiplicative
}

var cache_Data_Monoid_Multiplicative_monadMultiplicative gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_monadMultiplicative sync.Once
func Get_Data_Monoid_Multiplicative_monadMultiplicative() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_monadMultiplicative.Do(func() {
		cache_Data_Monoid_Multiplicative_monadMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Monoid_Multiplicative_applicativeMultiplicative()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Monoid_Multiplicative_bindMultiplicative()))}
})})}
	})
	return cache_Data_Monoid_Multiplicative_monadMultiplicative
}

var cache_Data_Monoid_Multiplicative_applicativeMultiplicative__995286821 gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_applicativeMultiplicative__995286821 sync.Once
func Get_Data_Monoid_Multiplicative_applicativeMultiplicative__995286821() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_applicativeMultiplicative__995286821.Do(func() {
		cache_Data_Monoid_Multiplicative_applicativeMultiplicative__995286821 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Multiplicative_applyMultiplicative()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Monoid_Multiplicative_applicativeMultiplicative__995286821
}

var cache_Data_Monoid_Multiplicative_applyMultiplicative__111100453 gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_applyMultiplicative__111100453 sync.Once
func Get_Data_Monoid_Multiplicative_applyMultiplicative__111100453() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_applyMultiplicative__111100453.Do(func() {
		cache_Data_Monoid_Multiplicative_applyMultiplicative__111100453 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Monoid_Multiplicative_functorMultiplicative()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
})})}
	})
	return cache_Data_Monoid_Multiplicative_applyMultiplicative__111100453
}

var cache_Data_Monoid_Multiplicative_bindMultiplicative__2224694053 gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_bindMultiplicative__2224694053 sync.Once
func Get_Data_Monoid_Multiplicative_bindMultiplicative__2224694053() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_bindMultiplicative__2224694053.Do(func() {
		cache_Data_Monoid_Multiplicative_bindMultiplicative__2224694053 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Multiplicative_applyMultiplicative()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
})})}
	})
	return cache_Data_Monoid_Multiplicative_bindMultiplicative__2224694053
}

var cache_Data_Monoid_Multiplicative_eq1Multiplicative__294625475 gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_eq1Multiplicative__294625475 sync.Once
func Get_Data_Monoid_Multiplicative_eq1Multiplicative__294625475() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_eq1Multiplicative__294625475.Do(func() {
		cache_Data_Monoid_Multiplicative_eq1Multiplicative__294625475 = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
})})}
	})
	return cache_Data_Monoid_Multiplicative_eq1Multiplicative__294625475
}

var cache_Data_Monoid_Multiplicative_functorMultiplicative__850816530 gopurs_runtime.Value
var once_Data_Monoid_Multiplicative_functorMultiplicative__850816530 sync.Once
func Get_Data_Monoid_Multiplicative_functorMultiplicative__850816530() gopurs_runtime.Value {
	once_Data_Monoid_Multiplicative_functorMultiplicative__850816530.Do(func() {
		cache_Data_Monoid_Multiplicative_functorMultiplicative__850816530 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Monoid_Multiplicative_functorMultiplicative__850816530
}

func Call_Data_Monoid_Multiplicative_Multiplicative(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Multiplicative_showMultiplicative(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Multiplicative ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Monoid_Multiplicative_semigroupMultiplicative(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
})
})})}
}

func Call_Data_Monoid_Multiplicative_ordMultiplicative(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0))}
}

func Call_Data_Monoid_Multiplicative_monoidMultiplicative(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
// TAST (Let): semigroupMultiplicative1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupMultiplicative1_1_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
})
})}
_ = semigroupMultiplicative1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMultiplicative1_1_0)}
}), gopurs_runtime.RecordGet(dictSemiring_0, "one")})}
}

func Call_Data_Monoid_Multiplicative_eqMultiplicative(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}

func Call_Data_Monoid_Multiplicative_boundedMultiplicative(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](dictBounded_0))}
}


