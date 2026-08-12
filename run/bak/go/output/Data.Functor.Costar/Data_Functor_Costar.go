package Data_Functor_Costar

import (
	pkg_Control_Comonad "gopurs/output/Control.Comonad"
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_Contravariant "gopurs/output/Data.Functor.Contravariant"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_functorCostar = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
}))
	})
	return cache_functorCostar
}

var cache_invariantCostar gopurs_runtime.Value
var once_invariantCostar sync.Once
func Get_invariantCostar() gopurs_runtime.Value {
	once_invariantCostar.Do(func() {
		cache_invariantCostar = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorCostar(), "map"), f_0)
})
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
}), gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_distributiveCostar(), "distribute"), dictFunctor_0)
_ = __local_var_2_0
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
})
})
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(v_0, x_2), x_2)
})
})
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
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
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

var cache_extract__1031647521 gopurs_runtime.Value
var once_extract__1031647521 sync.Once
func Get_extract__1031647521() gopurs_runtime.Value {
	once_extract__1031647521.Do(func() {
		cache_extract__1031647521 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract__1031647521(gopurs_runtime.CoerceToStruct[pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extract__1031647521
}

var cache_extract__3319904577 gopurs_runtime.Value
var once_extract__3319904577 sync.Once
func Get_extract__3319904577() gopurs_runtime.Value {
	once_extract__3319904577.Do(func() {
		cache_extract__3319904577 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract__3319904577(gopurs_runtime.CoerceToStruct[pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extract__3319904577
}

var cache_composeCoKleisliFlipped__1582554720 gopurs_runtime.Value
var once_composeCoKleisliFlipped__1582554720 sync.Once
func Get_composeCoKleisliFlipped__1582554720() gopurs_runtime.Value {
	once_composeCoKleisliFlipped__1582554720.Do(func() {
		cache_composeCoKleisliFlipped__1582554720 = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeCoKleisliFlipped__1582554720(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dictExtend_0_box), f_1_box, g_2_box, w_3_box)
})
	})
	return cache_composeCoKleisliFlipped__1582554720
}

var cache_extend__1264481661 gopurs_runtime.Value
var once_extend__1264481661 sync.Once
func Get_extend__1264481661() gopurs_runtime.Value {
	once_extend__1264481661.Do(func() {
		cache_extend__1264481661 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__1264481661(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__1264481661
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_cmap__326373820 gopurs_runtime.Value
var once_cmap__326373820 sync.Once
func Get_cmap__326373820() gopurs_runtime.Value {
	once_cmap__326373820.Do(func() {
		cache_cmap__326373820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmap__326373820(gopurs_runtime.CoerceToStruct[pkg_Data_Functor_Contravariant.Constructor_Contravariant[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_cmap__326373820
}

var cache_applicativeCostar__2238599400 gopurs_runtime.Value
var once_applicativeCostar__2238599400 sync.Once
func Get_applicativeCostar__2238599400() gopurs_runtime.Value {
	once_applicativeCostar__2238599400.Do(func() {
		cache_applicativeCostar__2238599400 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyCostar()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
}))
	})
	return cache_applicativeCostar__2238599400
}

var cache_applyCostar__1509714460 gopurs_runtime.Value
var once_applyCostar__1509714460 sync.Once
func Get_applyCostar__1509714460() gopurs_runtime.Value {
	once_applyCostar__1509714460.Do(func() {
		cache_applyCostar__1509714460 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorCostar()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
})
})
}))
	})
	return cache_applyCostar__1509714460
}

var cache_bindCostar__1019009222 gopurs_runtime.Value
var once_bindCostar__1019009222 sync.Once
func Get_bindCostar__1019009222() gopurs_runtime.Value {
	once_bindCostar__1019009222.Do(func() {
		cache_bindCostar__1019009222 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyCostar()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(v_0, x_2), x_2)
})
})
}))
	})
	return cache_bindCostar__1019009222
}

var cache_functorCostar__735509168 gopurs_runtime.Value
var once_functorCostar__735509168 sync.Once
func Get_functorCostar__735509168() gopurs_runtime.Value {
	once_functorCostar__735509168.Do(func() {
		cache_functorCostar__735509168 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
}))
	})
	return cache_functorCostar__735509168
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__691697300 gopurs_runtime.Value
var once_map__691697300 sync.Once
func Get_map__691697300() gopurs_runtime.Value {
	once_map__691697300.Do(func() {
		cache_map__691697300 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__691697300(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__691697300
}

var cache_map__2345808404 gopurs_runtime.Value
var once_map__2345808404 sync.Once
func Get_map__2345808404() gopurs_runtime.Value {
	once_map__2345808404.Do(func() {
		cache_map__2345808404 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2345808404(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2345808404
}

var cache_map__1938733460 gopurs_runtime.Value
var once_map__1938733460 sync.Once
func Get_map__1938733460() gopurs_runtime.Value {
	once_map__1938733460.Do(func() {
		cache_map__1938733460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1938733460(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1938733460
}

var cache_map__2418274292 gopurs_runtime.Value
var once_map__2418274292 sync.Once
func Get_map__2418274292() gopurs_runtime.Value {
	once_map__2418274292.Do(func() {
		cache_map__2418274292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2418274292(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2418274292
}

var cache_dimap__650672149 gopurs_runtime.Value
var once_dimap__650672149 sync.Once
func Get_dimap__650672149() gopurs_runtime.Value {
	once_dimap__650672149.Do(func() {
		cache_dimap__650672149 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap__650672149(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_dimap__650672149
}

var cache_dimap__1466332548 gopurs_runtime.Value
var once_dimap__1466332548 sync.Once
func Get_dimap__1466332548() gopurs_runtime.Value {
	once_dimap__1466332548.Do(func() {
		cache_dimap__1466332548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap__1466332548(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_dimap__1466332548
}

var cache_lcmap__1762133278 gopurs_runtime.Value
var once_lcmap__1762133278 sync.Once
func Get_lcmap__1762133278() gopurs_runtime.Value {
	once_lcmap__1762133278.Do(func() {
		cache_lcmap__1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcmap__1762133278(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), a2b_1_box)
})
	})
	return cache_lcmap__1762133278
}

var cache_lcmap__351678174 gopurs_runtime.Value
var once_lcmap__351678174 sync.Once
func Get_lcmap__351678174() gopurs_runtime.Value {
	once_lcmap__351678174.Do(func() {
		cache_lcmap__351678174 = gopurs_runtime.Func(func(a2b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcmap__351678174(a2b_0_box)
})
	})
	return cache_lcmap__351678174
}

var cache_profunctorFn__542207281 gopurs_runtime.Value
var once_profunctorFn__542207281 sync.Once
func Get_profunctorFn__542207281() gopurs_runtime.Value {
	once_profunctorFn__542207281.Do(func() {
		cache_profunctorFn__542207281 = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_profunctorFn__542207281
}

var cache_fst__20422131 gopurs_runtime.Value
var once_fst__20422131 sync.Once
func Get_fst__20422131() gopurs_runtime.Value {
	once_fst__20422131.Do(func() {
		cache_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__20422131
}

var cache_fst__395594805 gopurs_runtime.Value
var once_fst__395594805 sync.Once
func Get_fst__395594805() gopurs_runtime.Value {
	once_fst__395594805.Do(func() {
		cache_fst__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__395594805(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__395594805
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_snd__395594805 gopurs_runtime.Value
var once_snd__395594805 sync.Once
func Get_snd__395594805() gopurs_runtime.Value {
	once_snd__395594805.Do(func() {
		cache_snd__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__395594805(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__395594805
}

func Call_Costar(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_semigroupoidCostar(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), v1_2, w_3))
})
})
}))
}

func Call_profunctorCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_0, x_5)))
})
})
})
}))
}

func Call_strongCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
Extend0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = Extend0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Extend0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Extend0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
profunctorCostar1_3_2 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), f_4)
_ = __local_var_7_4
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_5, gopurs_runtime.Apply(v_6, gopurs_runtime.Apply(__local_var_7_4, x_8)))
})
})
})
}))
_ = profunctorCostar1_3_2
return gopurs_runtime.RecordDict3("Profunctor0", "first", "second", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorCostar1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_Tuple.Get_fst(), x_5)), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_5).UnsafePtr).V1})}
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_5).UnsafePtr).V0, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_Tuple.Get_snd(), x_5))})}
})
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
profunctorCostar1_1_0 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_1, x_5)))
})
})
})
}))
_ = profunctorCostar1_1_0
return gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorCostar1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, x_4)
}), g_3))
})
})
}))
}

func Call_categoryCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupoidCostar1_1_0 := gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "extend"), v1_3, w_4))
})
})
}))
_ = semigroupoidCostar1_1_0
return gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidCostar1_1_0
}), gopurs_runtime.RecordGet(dictComonad_0, "extract"))
}

func Call_bifunctorCostar(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1)
_ = __local_var_4_0
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_0, x_5)))
})
})
})
}))
}

func Call_extract__1031647521(dict_0_loop *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extract__3319904577(dict_0_loop *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_composeCoKleisliFlipped__1582554720(dictExtend_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dictExtend_0_loop
_ = dictExtend_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var w_3 gopurs_runtime.Value = w_3_loop
_ = w_3
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(dictExtend_0.V1, g_2, w_3))
}

func Call_extend__1264481661(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_cmap__326373820(dict_0_loop *pkg_Data_Functor_Contravariant.Constructor_Contravariant[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor_Contravariant.Constructor_Contravariant[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__691697300(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2345808404(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1938733460(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2418274292(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_dimap__650672149(dict_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_dimap__1466332548(dict_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_lcmap__1762133278(dictProfunctor_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value], a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(dictProfunctor_0.V0, a2b_1, pkg_Data_Profunctor.Get_identity())
}

func Call_lcmap__351678174(a2b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a2b_0 gopurs_runtime.Value = a2b_0_loop
_ = a2b_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Profunctor.Get_profunctorFn(), "dimap"), a2b_0, pkg_Data_Profunctor.Get_identity())
}

func Call_fst__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_fst__395594805(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__395594805(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}


