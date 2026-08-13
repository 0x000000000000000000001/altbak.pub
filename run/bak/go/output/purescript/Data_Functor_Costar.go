package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Costar_Costar gopurs_runtime.Value
var once_Data_Functor_Costar_Costar sync.Once
func Get_Data_Functor_Costar_Costar() gopurs_runtime.Value {
	once_Data_Functor_Costar_Costar.Do(func() {
		cache_Data_Functor_Costar_Costar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_Costar(x_0_box)
})
	})
	return cache_Data_Functor_Costar_Costar
}

var cache_Data_Functor_Costar_semigroupoidCostar gopurs_runtime.Value
var once_Data_Functor_Costar_semigroupoidCostar sync.Once
func Get_Data_Functor_Costar_semigroupoidCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_semigroupoidCostar.Do(func() {
		cache_Data_Functor_Costar_semigroupoidCostar = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_semigroupoidCostar(dictExtend_0_box)
})
	})
	return cache_Data_Functor_Costar_semigroupoidCostar
}

var cache_Data_Functor_Costar_profunctorCostar gopurs_runtime.Value
var once_Data_Functor_Costar_profunctorCostar sync.Once
func Get_Data_Functor_Costar_profunctorCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_profunctorCostar.Do(func() {
		cache_Data_Functor_Costar_profunctorCostar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_profunctorCostar(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Costar_profunctorCostar
}

var cache_Data_Functor_Costar_strongCostar gopurs_runtime.Value
var once_Data_Functor_Costar_strongCostar sync.Once
func Get_Data_Functor_Costar_strongCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_strongCostar.Do(func() {
		cache_Data_Functor_Costar_strongCostar = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_strongCostar(dictComonad_0_box)
})
	})
	return cache_Data_Functor_Costar_strongCostar
}

var cache_Data_Functor_Costar_newtypeCostar gopurs_runtime.Value
var once_Data_Functor_Costar_newtypeCostar sync.Once
func Get_Data_Functor_Costar_newtypeCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_newtypeCostar.Do(func() {
		cache_Data_Functor_Costar_newtypeCostar = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Functor_Costar_newtypeCostar
}

var cache_Data_Functor_Costar_hoistCostar gopurs_runtime.Value
var once_Data_Functor_Costar_hoistCostar sync.Once
func Get_Data_Functor_Costar_hoistCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_hoistCostar.Do(func() {
		cache_Data_Functor_Costar_hoistCostar = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_hoistCostar(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Data_Functor_Costar_hoistCostar
}

var cache_Data_Functor_Costar_functorCostar gopurs_runtime.Value
var once_Data_Functor_Costar_functorCostar sync.Once
func Get_Data_Functor_Costar_functorCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_functorCostar.Do(func() {
		cache_Data_Functor_Costar_functorCostar = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
})})}
	})
	return cache_Data_Functor_Costar_functorCostar
}

var cache_Data_Functor_Costar_invariantCostar gopurs_runtime.Value
var once_Data_Functor_Costar_invariantCostar sync.Once
func Get_Data_Functor_Costar_invariantCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_invariantCostar.Do(func() {
		cache_Data_Functor_Costar_invariantCostar = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Invariant_Invariant{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_2, x_3))
})
})
})
})})}
	})
	return cache_Data_Functor_Costar_invariantCostar
}

var cache_Data_Functor_Costar_distributiveCostar gopurs_runtime.Value
var once_Data_Functor_Costar_distributiveCostar sync.Once
func Get_Data_Functor_Costar_distributiveCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_distributiveCostar.Do(func() {
		cache_Data_Functor_Costar_distributiveCostar = gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer(&Constructor_Data_Distributive_Distributive{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_Costar_functorCostar()))}
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Distributive_Distributive](Get_Data_Functor_Costar_distributiveCostar()).V2), dictFunctor_0)
_ = __local_var_2_0
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, a_2)
}), f_1)
})
})
})})}
	})
	return cache_Data_Functor_Costar_distributiveCostar
}

var cache_Data_Functor_Costar_closedCostar gopurs_runtime.Value
var once_Data_Functor_Costar_closedCostar sync.Once
func Get_Data_Functor_Costar_closedCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_closedCostar.Do(func() {
		cache_Data_Functor_Costar_closedCostar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_closedCostar(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Costar_closedCostar
}

var cache_Data_Functor_Costar_categoryCostar gopurs_runtime.Value
var once_Data_Functor_Costar_categoryCostar sync.Once
func Get_Data_Functor_Costar_categoryCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_categoryCostar.Do(func() {
		cache_Data_Functor_Costar_categoryCostar = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_categoryCostar(dictComonad_0_box)
})
	})
	return cache_Data_Functor_Costar_categoryCostar
}

var cache_Data_Functor_Costar_bifunctorCostar gopurs_runtime.Value
var once_Data_Functor_Costar_bifunctorCostar sync.Once
func Get_Data_Functor_Costar_bifunctorCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_bifunctorCostar.Do(func() {
		cache_Data_Functor_Costar_bifunctorCostar = gopurs_runtime.Func(func(dictContravariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Costar_bifunctorCostar(dictContravariant_0_box)
})
	})
	return cache_Data_Functor_Costar_bifunctorCostar
}

var cache_Data_Functor_Costar_applyCostar gopurs_runtime.Value
var once_Data_Functor_Costar_applyCostar sync.Once
func Get_Data_Functor_Costar_applyCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_applyCostar.Do(func() {
		cache_Data_Functor_Costar_applyCostar = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_Costar_functorCostar()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
})
})
})})}
	})
	return cache_Data_Functor_Costar_applyCostar
}

var cache_Data_Functor_Costar_bindCostar gopurs_runtime.Value
var once_Data_Functor_Costar_bindCostar sync.Once
func Get_Data_Functor_Costar_bindCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_bindCostar.Do(func() {
		cache_Data_Functor_Costar_bindCostar = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Functor_Costar_applyCostar()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(v_0, x_2), x_2)
})
})
})})}
	})
	return cache_Data_Functor_Costar_bindCostar
}

var cache_Data_Functor_Costar_applicativeCostar gopurs_runtime.Value
var once_Data_Functor_Costar_applicativeCostar sync.Once
func Get_Data_Functor_Costar_applicativeCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_applicativeCostar.Do(func() {
		cache_Data_Functor_Costar_applicativeCostar = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Functor_Costar_applyCostar()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})})}
	})
	return cache_Data_Functor_Costar_applicativeCostar
}

var cache_Data_Functor_Costar_monadCostar gopurs_runtime.Value
var once_Data_Functor_Costar_monadCostar sync.Once
func Get_Data_Functor_Costar_monadCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_monadCostar.Do(func() {
		cache_Data_Functor_Costar_monadCostar = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Functor_Costar_applicativeCostar()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_Functor_Costar_bindCostar()))}
})})}
	})
	return cache_Data_Functor_Costar_monadCostar
}

var cache_Data_Functor_Costar_applicativeCostar__2806900578 gopurs_runtime.Value
var once_Data_Functor_Costar_applicativeCostar__2806900578 sync.Once
func Get_Data_Functor_Costar_applicativeCostar__2806900578() gopurs_runtime.Value {
	once_Data_Functor_Costar_applicativeCostar__2806900578.Do(func() {
		cache_Data_Functor_Costar_applicativeCostar__2806900578 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Functor_Costar_applyCostar()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})})}
	})
	return cache_Data_Functor_Costar_applicativeCostar__2806900578
}

var cache_Data_Functor_Costar_applyCostar__508451810 gopurs_runtime.Value
var once_Data_Functor_Costar_applyCostar__508451810 sync.Once
func Get_Data_Functor_Costar_applyCostar__508451810() gopurs_runtime.Value {
	once_Data_Functor_Costar_applyCostar__508451810.Do(func() {
		cache_Data_Functor_Costar_applyCostar__508451810 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_Costar_functorCostar()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
})
})
})})}
	})
	return cache_Data_Functor_Costar_applyCostar__508451810
}

var cache_Data_Functor_Costar_bindCostar__2325516642 gopurs_runtime.Value
var once_Data_Functor_Costar_bindCostar__2325516642 sync.Once
func Get_Data_Functor_Costar_bindCostar__2325516642() gopurs_runtime.Value {
	once_Data_Functor_Costar_bindCostar__2325516642.Do(func() {
		cache_Data_Functor_Costar_bindCostar__2325516642 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Functor_Costar_applyCostar()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(v_0, x_2), x_2)
})
})
})})}
	})
	return cache_Data_Functor_Costar_bindCostar__2325516642
}

var cache_Data_Functor_Costar_functorCostar__3212384053 gopurs_runtime.Value
var once_Data_Functor_Costar_functorCostar__3212384053 sync.Once
func Get_Data_Functor_Costar_functorCostar__3212384053() gopurs_runtime.Value {
	once_Data_Functor_Costar_functorCostar__3212384053.Do(func() {
		cache_Data_Functor_Costar_functorCostar__3212384053 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
})})}
	})
	return cache_Data_Functor_Costar_functorCostar__3212384053
}

func Call_Data_Functor_Costar_Costar(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Costar_semigroupoidCostar(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(&Constructor_Control_Semigroupoid_Semigroupoid{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), v1_2, w_3))
})
})
})})}
}

func Call_Data_Functor_Costar_profunctorCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Profunctor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_0, x_5)))
})
})
})
})})}
}

func Call_Data_Functor_Costar_strongCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): Extend0_1_0 -> gopurs_runtime.Value
Extend0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = Extend0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Extend0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Extend0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): profunctorCostar1_3_2 -> *Constructor_Data_Profunctor_Profunctor
profunctorCostar1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor](gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), f_4)
_ = __local_var_7_4
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_5, gopurs_runtime.Apply(v_6, gopurs_runtime.Apply(__local_var_7_4, x_8)))
})
})
})
})))
_ = profunctorCostar1_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 1323482783, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Strong_Strong{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorCostar1_3_2)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Tuple_fst(), x_5)), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_5).UnsafePtr).V1})}
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_5).UnsafePtr).V0, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Tuple_snd(), x_5))})}
})
})})}
}

func Call_Data_Functor_Costar_hoistCostar(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_Data_Functor_Costar_closedCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
// TAST (Let): profunctorCostar1_1_0 -> *Constructor_Data_Profunctor_Profunctor
profunctorCostar1_1_0 := &Constructor_Data_Profunctor_Profunctor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_1, x_5)))
})
})
})
})}
_ = profunctorCostar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 768764671, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Closed_Closed{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorCostar1_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, x_4)
}), g_3))
})
})
})})}
}

func Call_Data_Functor_Costar_categoryCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupoidCostar1_1_0 -> *Constructor_Control_Semigroupoid_Semigroupoid
semigroupoidCostar1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "extend"), v1_3, w_4))
})
})
})))
_ = semigroupoidCostar1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(&Constructor_Control_Category_Category{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(semigroupoidCostar1_1_0)}
}), gopurs_runtime.RecordGet(dictComonad_0, "extract")})}
}

func Call_Data_Functor_Costar_bifunctorCostar(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_0, x_5)))
})
})
})
})})}
}


