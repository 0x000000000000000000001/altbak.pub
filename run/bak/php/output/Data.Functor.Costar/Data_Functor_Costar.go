package Data_Functor_Costar

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	unsafe "unsafe"
)

var cache_Costar gopurs_runtime.Value
var once_Costar sync.Once
func Get_Costar() gopurs_runtime.Value {
	once_Costar.Do(func() {
		cache_Costar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Costar(x_0_box)
})
	})
	return cache_Costar
}

var cache_semigroupoidCostar gopurs_runtime.Value
var once_semigroupoidCostar sync.Once
func Get_semigroupoidCostar() gopurs_runtime.Value {
	once_semigroupoidCostar.Do(func() {
		cache_semigroupoidCostar = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupoidCostar(dictExtend_0_box)
})
	})
	return cache_semigroupoidCostar
}

var cache_profunctorCostar gopurs_runtime.Value
var once_profunctorCostar sync.Once
func Get_profunctorCostar() gopurs_runtime.Value {
	once_profunctorCostar.Do(func() {
		cache_profunctorCostar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_profunctorCostar(dictFunctor_0_box)
})
	})
	return cache_profunctorCostar
}

var cache_strongCostar gopurs_runtime.Value
var once_strongCostar sync.Once
func Get_strongCostar() gopurs_runtime.Value {
	once_strongCostar.Do(func() {
		cache_strongCostar = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_strongCostar(dictComonad_0_box)
})
	})
	return cache_strongCostar
}

var cache_newtypeCostar gopurs_runtime.Value
var once_newtypeCostar sync.Once
func Get_newtypeCostar() gopurs_runtime.Value {
	once_newtypeCostar.Do(func() {
		cache_newtypeCostar = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCostar
}

var cache_hoistCostar gopurs_runtime.Value
var once_hoistCostar sync.Once
func Get_hoistCostar() gopurs_runtime.Value {
	once_hoistCostar.Do(func() {
		cache_hoistCostar = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistCostar(f_0_box, v_1_box)
})
	})
	return cache_hoistCostar
}

var cache_functorCostar gopurs_runtime.Value
var once_functorCostar sync.Once
func Get_functorCostar() gopurs_runtime.Value {
	once_functorCostar.Do(func() {
		cache_functorCostar = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, v_1)
}))
	})
	return cache_functorCostar
}

var cache_invariantCostar gopurs_runtime.Value
var once_invariantCostar sync.Once
func Get_invariantCostar() gopurs_runtime.Value {
	once_invariantCostar.Do(func() {
		cache_invariantCostar = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorCostar(), "map"), f_0)
}))
	})
	return cache_invariantCostar
}

var cache_distributiveCostar gopurs_runtime.Value
var once_distributiveCostar sync.Once
func Get_distributiveCostar() gopurs_runtime.Value {
	once_distributiveCostar.Do(func() {
		cache_distributiveCostar = gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorCostar()
}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_distributiveCostar(), "distribute"), dictFunctor_0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1))
}), gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, a_2)
}), f_1)
}))
	})
	return cache_distributiveCostar
}

var cache_closedCostar gopurs_runtime.Value
var once_closedCostar sync.Once
func Get_closedCostar() gopurs_runtime.Value {
	once_closedCostar.Do(func() {
		cache_closedCostar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_closedCostar(dictFunctor_0_box)
})
	})
	return cache_closedCostar
}

var cache_categoryCostar gopurs_runtime.Value
var once_categoryCostar sync.Once
func Get_categoryCostar() gopurs_runtime.Value {
	once_categoryCostar.Do(func() {
		cache_categoryCostar = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_categoryCostar(dictComonad_0_box)
})
	})
	return cache_categoryCostar
}

var cache_bifunctorCostar gopurs_runtime.Value
var once_bifunctorCostar sync.Once
func Get_bifunctorCostar() gopurs_runtime.Value {
	once_bifunctorCostar.Do(func() {
		cache_bifunctorCostar = gopurs_runtime.Func(func(dictContravariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifunctorCostar(dictContravariant_0_box)
})
	})
	return cache_bifunctorCostar
}

var cache_applyCostar gopurs_runtime.Value
var once_applyCostar sync.Once
func Get_applyCostar() gopurs_runtime.Value {
	once_applyCostar.Do(func() {
		cache_applyCostar = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorCostar()
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
}))
	})
	return cache_applyCostar
}

var cache_bindCostar gopurs_runtime.Value
var once_bindCostar sync.Once
func Get_bindCostar() gopurs_runtime.Value {
	once_bindCostar.Do(func() {
		cache_bindCostar = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyCostar()
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(v_0, x_2), x_2)
}))
	})
	return cache_bindCostar
}

var cache_applicativeCostar gopurs_runtime.Value
var once_applicativeCostar sync.Once
func Get_applicativeCostar() gopurs_runtime.Value {
	once_applicativeCostar.Do(func() {
		cache_applicativeCostar = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyCostar()
}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
	})
	return cache_applicativeCostar
}

var cache_monadCostar gopurs_runtime.Value
var once_monadCostar sync.Once
func Get_monadCostar() gopurs_runtime.Value {
	once_monadCostar.Do(func() {
		cache_monadCostar = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeCostar()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindCostar()
}))
	})
	return cache_monadCostar
}

func Call_Costar(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_semigroupoidCostar(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictExtend_0.UnsafePtr)).V0, v1_2, w_3))
}))
}

func Call_profunctorCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), g_2, v_3), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, f_1))
}))
}

func Call_strongCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0_NOT_FOUND"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
profunctorCostar1_2_1 := gopurs_runtime.Apply(Get_profunctorCostar(), Functor0_1_0)
_ = profunctorCostar1_2_1
return gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorCostar1_2_1
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(v_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), pkg_Data_Tuple.Get_fst(), x_4)), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, x_4).UnsafePtr).V1})}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, x_4).UnsafePtr).V0, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), pkg_Data_Tuple.Get_snd(), x_4))})}
}))
}

func Call_hoistCostar(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Profunctor.Get_profunctorFn(), "dimap"), f_0, pkg_Data_Profunctor.Get_identity(), v_1)
}

func Call_closedCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
profunctorCostar1_1_0 := gopurs_runtime.Apply(Get_profunctorCostar(), dictFunctor_0)
_ = profunctorCostar1_1_0
return gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorCostar1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, x_4)
}), g_3))
}))
}

func Call_categoryCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupoidCostar1_2_1 := gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, w_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), v1_3, w_4))
}))
_ = semigroupoidCostar1_2_1
return gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidCostar1_2_1
}), ((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0)
}

func Call_bifunctorCostar(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), g_2, v_3), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictContravariant_0.UnsafePtr)).V0, f_1))
}))
}


