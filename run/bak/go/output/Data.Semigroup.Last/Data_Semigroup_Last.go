package Data_Semigroup_Last

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Last gopurs_runtime.Value
var once_Last sync.Once
func Get_Last() gopurs_runtime.Value {
	once_Last.Do(func() {
		cache_Last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Last(x_0_box)
})
	})
	return cache_Last
}

var cache_showLast gopurs_runtime.Value
var once_showLast sync.Once
func Get_showLast() gopurs_runtime.Value {
	once_showLast.Do(func() {
		cache_showLast = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showLast(dictShow_0_box)
})
	})
	return cache_showLast
}

var cache_semigroupLast gopurs_runtime.Value
var once_semigroupLast sync.Once
func Get_semigroupLast() gopurs_runtime.Value {
	once_semigroupLast.Do(func() {
		cache_semigroupLast = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_semigroupLast
}

var cache_ordLast gopurs_runtime.Value
var once_ordLast sync.Once
func Get_ordLast() gopurs_runtime.Value {
	once_ordLast.Do(func() {
		cache_ordLast = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordLast(dictOrd_0_box)
})
	})
	return cache_ordLast
}

var cache_functorLast gopurs_runtime.Value
var once_functorLast sync.Once
func Get_functorLast() gopurs_runtime.Value {
	once_functorLast.Do(func() {
		cache_functorLast = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorLast
}

var cache_eqLast gopurs_runtime.Value
var once_eqLast sync.Once
func Get_eqLast() gopurs_runtime.Value {
	once_eqLast.Do(func() {
		cache_eqLast = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqLast(dictEq_0_box)
})
	})
	return cache_eqLast
}

var cache_eq1Last gopurs_runtime.Value
var once_eq1Last sync.Once
func Get_eq1Last() gopurs_runtime.Value {
	once_eq1Last.Do(func() {
		cache_eq1Last = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Last
}

var cache_ord1Last gopurs_runtime.Value
var once_ord1Last sync.Once
func Get_ord1Last() gopurs_runtime.Value {
	once_ord1Last.Do(func() {
		cache_ord1Last = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Last()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1Last
}

var cache_boundedLast gopurs_runtime.Value
var once_boundedLast sync.Once
func Get_boundedLast() gopurs_runtime.Value {
	once_boundedLast.Do(func() {
		cache_boundedLast = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedLast(dictBounded_0_box)
})
	})
	return cache_boundedLast
}

var cache_applyLast gopurs_runtime.Value
var once_applyLast sync.Once
func Get_applyLast() gopurs_runtime.Value {
	once_applyLast.Do(func() {
		cache_applyLast = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyLast
}

var cache_bindLast gopurs_runtime.Value
var once_bindLast sync.Once
func Get_bindLast() gopurs_runtime.Value {
	once_bindLast.Do(func() {
		cache_bindLast = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindLast
}

var cache_applicativeLast gopurs_runtime.Value
var once_applicativeLast sync.Once
func Get_applicativeLast() gopurs_runtime.Value {
	once_applicativeLast.Do(func() {
		cache_applicativeLast = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLast()
}), Get_Last())
	})
	return cache_applicativeLast
}

var cache_monadLast gopurs_runtime.Value
var once_monadLast sync.Once
func Get_monadLast() gopurs_runtime.Value {
	once_monadLast.Do(func() {
		cache_monadLast = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindLast()
}))
	})
	return cache_monadLast
}

var cache_applicativeLast__4045440648 gopurs_runtime.Value
var once_applicativeLast__4045440648 sync.Once
func Get_applicativeLast__4045440648() gopurs_runtime.Value {
	once_applicativeLast__4045440648.Do(func() {
		cache_applicativeLast__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLast()
}), Get_Last())
	})
	return cache_applicativeLast__4045440648
}

var cache_applyLast__3199351098 gopurs_runtime.Value
var once_applyLast__3199351098 sync.Once
func Get_applyLast__3199351098() gopurs_runtime.Value {
	once_applyLast__3199351098.Do(func() {
		cache_applyLast__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyLast__3199351098
}

var cache_bindLast__329376103 gopurs_runtime.Value
var once_bindLast__329376103 sync.Once
func Get_bindLast__329376103() gopurs_runtime.Value {
	once_bindLast__329376103.Do(func() {
		cache_bindLast__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindLast__329376103
}

var cache_eq1Last__1905950174 gopurs_runtime.Value
var once_eq1Last__1905950174 sync.Once
func Get_eq1Last__1905950174() gopurs_runtime.Value {
	once_eq1Last__1905950174.Do(func() {
		cache_eq1Last__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Last__1905950174
}

var cache_functorLast__943655089 gopurs_runtime.Value
var once_functorLast__943655089 sync.Once
func Get_functorLast__943655089() gopurs_runtime.Value {
	once_functorLast__943655089.Do(func() {
		cache_functorLast__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorLast__943655089
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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

func Call_Last(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Last "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


