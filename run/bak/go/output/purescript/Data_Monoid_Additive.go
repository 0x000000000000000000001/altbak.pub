package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Additive_Additive gopurs_runtime.Value
var once_Data_Monoid_Additive_Additive sync.Once
func Get_Data_Monoid_Additive_Additive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_Additive.Do(func() {
		cache_Data_Monoid_Additive_Additive = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_Additive(x_0_box)
})
	})
	return cache_Data_Monoid_Additive_Additive
}

var cache_Data_Monoid_Additive_showAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_showAdditive sync.Once
func Get_Data_Monoid_Additive_showAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_showAdditive.Do(func() {
		cache_Data_Monoid_Additive_showAdditive = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_showAdditive(dictShow_0_box)
})
	})
	return cache_Data_Monoid_Additive_showAdditive
}

var cache_Data_Monoid_Additive_semigroupAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_semigroupAdditive sync.Once
func Get_Data_Monoid_Additive_semigroupAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_semigroupAdditive.Do(func() {
		cache_Data_Monoid_Additive_semigroupAdditive = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_semigroupAdditive(dictSemiring_0_box)
})
	})
	return cache_Data_Monoid_Additive_semigroupAdditive
}

var cache_Data_Monoid_Additive_ordAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_ordAdditive sync.Once
func Get_Data_Monoid_Additive_ordAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_ordAdditive.Do(func() {
		cache_Data_Monoid_Additive_ordAdditive = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_ordAdditive(dictOrd_0_box)
})
	})
	return cache_Data_Monoid_Additive_ordAdditive
}

var cache_Data_Monoid_Additive_monoidAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_monoidAdditive sync.Once
func Get_Data_Monoid_Additive_monoidAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_monoidAdditive.Do(func() {
		cache_Data_Monoid_Additive_monoidAdditive = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_monoidAdditive(dictSemiring_0_box)
})
	})
	return cache_Data_Monoid_Additive_monoidAdditive
}

var cache_Data_Monoid_Additive_functorAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_functorAdditive sync.Once
func Get_Data_Monoid_Additive_functorAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_functorAdditive.Do(func() {
		cache_Data_Monoid_Additive_functorAdditive = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Monoid_Additive_functorAdditive
}

var cache_Data_Monoid_Additive_eqAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_eqAdditive sync.Once
func Get_Data_Monoid_Additive_eqAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_eqAdditive.Do(func() {
		cache_Data_Monoid_Additive_eqAdditive = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_eqAdditive(dictEq_0_box)
})
	})
	return cache_Data_Monoid_Additive_eqAdditive
}

var cache_Data_Monoid_Additive_eq1Additive gopurs_runtime.Value
var once_Data_Monoid_Additive_eq1Additive sync.Once
func Get_Data_Monoid_Additive_eq1Additive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_eq1Additive.Do(func() {
		cache_Data_Monoid_Additive_eq1Additive = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
})})}
	})
	return cache_Data_Monoid_Additive_eq1Additive
}

var cache_Data_Monoid_Additive_ord1Additive gopurs_runtime.Value
var once_Data_Monoid_Additive_ord1Additive sync.Once
func Get_Data_Monoid_Additive_ord1Additive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_ord1Additive.Do(func() {
		cache_Data_Monoid_Additive_ord1Additive = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Monoid_Additive_eq1Additive()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
})})}
	})
	return cache_Data_Monoid_Additive_ord1Additive
}

var cache_Data_Monoid_Additive_boundedAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_boundedAdditive sync.Once
func Get_Data_Monoid_Additive_boundedAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_boundedAdditive.Do(func() {
		cache_Data_Monoid_Additive_boundedAdditive = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_boundedAdditive(dictBounded_0_box)
})
	})
	return cache_Data_Monoid_Additive_boundedAdditive
}

var cache_Data_Monoid_Additive_applyAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_applyAdditive sync.Once
func Get_Data_Monoid_Additive_applyAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applyAdditive.Do(func() {
		cache_Data_Monoid_Additive_applyAdditive = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Monoid_Additive_functorAdditive()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
})})}
	})
	return cache_Data_Monoid_Additive_applyAdditive
}

var cache_Data_Monoid_Additive_bindAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_bindAdditive sync.Once
func Get_Data_Monoid_Additive_bindAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_bindAdditive.Do(func() {
		cache_Data_Monoid_Additive_bindAdditive = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Additive_applyAdditive()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
})})}
	})
	return cache_Data_Monoid_Additive_bindAdditive
}

var cache_Data_Monoid_Additive_applicativeAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_applicativeAdditive sync.Once
func Get_Data_Monoid_Additive_applicativeAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applicativeAdditive.Do(func() {
		cache_Data_Monoid_Additive_applicativeAdditive = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Additive_applyAdditive()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Monoid_Additive_applicativeAdditive
}

var cache_Data_Monoid_Additive_monadAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_monadAdditive sync.Once
func Get_Data_Monoid_Additive_monadAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_monadAdditive.Do(func() {
		cache_Data_Monoid_Additive_monadAdditive = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Monoid_Additive_applicativeAdditive()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Monoid_Additive_bindAdditive()))}
})})}
	})
	return cache_Data_Monoid_Additive_monadAdditive
}

var cache_Data_Monoid_Additive_applicativeAdditive__995286821 gopurs_runtime.Value
var once_Data_Monoid_Additive_applicativeAdditive__995286821 sync.Once
func Get_Data_Monoid_Additive_applicativeAdditive__995286821() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applicativeAdditive__995286821.Do(func() {
		cache_Data_Monoid_Additive_applicativeAdditive__995286821 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Additive_applyAdditive()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})})}
	})
	return cache_Data_Monoid_Additive_applicativeAdditive__995286821
}

var cache_Data_Monoid_Additive_applyAdditive__111100453 gopurs_runtime.Value
var once_Data_Monoid_Additive_applyAdditive__111100453 sync.Once
func Get_Data_Monoid_Additive_applyAdditive__111100453() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applyAdditive__111100453.Do(func() {
		cache_Data_Monoid_Additive_applyAdditive__111100453 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Monoid_Additive_functorAdditive()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
})})}
	})
	return cache_Data_Monoid_Additive_applyAdditive__111100453
}

var cache_Data_Monoid_Additive_bindAdditive__2224694053 gopurs_runtime.Value
var once_Data_Monoid_Additive_bindAdditive__2224694053 sync.Once
func Get_Data_Monoid_Additive_bindAdditive__2224694053() gopurs_runtime.Value {
	once_Data_Monoid_Additive_bindAdditive__2224694053.Do(func() {
		cache_Data_Monoid_Additive_bindAdditive__2224694053 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Monoid_Additive_applyAdditive()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
})})}
	})
	return cache_Data_Monoid_Additive_bindAdditive__2224694053
}

var cache_Data_Monoid_Additive_eq1Additive__294625475 gopurs_runtime.Value
var once_Data_Monoid_Additive_eq1Additive__294625475 sync.Once
func Get_Data_Monoid_Additive_eq1Additive__294625475() gopurs_runtime.Value {
	once_Data_Monoid_Additive_eq1Additive__294625475.Do(func() {
		cache_Data_Monoid_Additive_eq1Additive__294625475 = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
})})}
	})
	return cache_Data_Monoid_Additive_eq1Additive__294625475
}

var cache_Data_Monoid_Additive_functorAdditive__850816530 gopurs_runtime.Value
var once_Data_Monoid_Additive_functorAdditive__850816530 sync.Once
func Get_Data_Monoid_Additive_functorAdditive__850816530() gopurs_runtime.Value {
	once_Data_Monoid_Additive_functorAdditive__850816530.Do(func() {
		cache_Data_Monoid_Additive_functorAdditive__850816530 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})}
	})
	return cache_Data_Monoid_Additive_functorAdditive__850816530
}

func Call_Data_Monoid_Additive_Additive(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Additive_showAdditive(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Additive ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Monoid_Additive_semigroupAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
})})}
}

func Call_Data_Monoid_Additive_ordAdditive(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0))}
}

func Call_Data_Monoid_Additive_monoidAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
// TAST (Let): semigroupAdditive1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupAdditive1_1_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
})}
_ = semigroupAdditive1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_1_0)}
}), gopurs_runtime.RecordGet(dictSemiring_0, "zero")})}
}

func Call_Data_Monoid_Additive_eqAdditive(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}

func Call_Data_Monoid_Additive_boundedAdditive(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](dictBounded_0))}
}


