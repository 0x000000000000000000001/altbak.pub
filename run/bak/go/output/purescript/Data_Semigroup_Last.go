package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Semigroup_Last_Last gopurs_runtime.Value
var once_Data_Semigroup_Last_Last sync.Once
func Get_Data_Semigroup_Last_Last() gopurs_runtime.Value {
	once_Data_Semigroup_Last_Last.Do(func() {
		cache_Data_Semigroup_Last_Last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Last_Last(x_0_box)
})
	})
	return cache_Data_Semigroup_Last_Last
}

var cache_Data_Semigroup_Last_showLast gopurs_runtime.Value
var once_Data_Semigroup_Last_showLast sync.Once
func Get_Data_Semigroup_Last_showLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_showLast.Do(func() {
		cache_Data_Semigroup_Last_showLast = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Last_showLast(dictShow_0_box)
})
	})
	return cache_Data_Semigroup_Last_showLast
}

var cache_Data_Semigroup_Last_semigroupLast gopurs_runtime.Value
var once_Data_Semigroup_Last_semigroupLast sync.Once
func Get_Data_Semigroup_Last_semigroupLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_semigroupLast.Do(func() {
		cache_Data_Semigroup_Last_semigroupLast = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_Data_Semigroup_Last_semigroupLast
}

var cache_Data_Semigroup_Last_ordLast gopurs_runtime.Value
var once_Data_Semigroup_Last_ordLast sync.Once
func Get_Data_Semigroup_Last_ordLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_ordLast.Do(func() {
		cache_Data_Semigroup_Last_ordLast = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Last_ordLast(dictOrd_0_box)
})
	})
	return cache_Data_Semigroup_Last_ordLast
}

var cache_Data_Semigroup_Last_functorLast gopurs_runtime.Value
var once_Data_Semigroup_Last_functorLast sync.Once
func Get_Data_Semigroup_Last_functorLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_functorLast.Do(func() {
		cache_Data_Semigroup_Last_functorLast = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Semigroup_Last_functorLast
}

var cache_Data_Semigroup_Last_eqLast gopurs_runtime.Value
var once_Data_Semigroup_Last_eqLast sync.Once
func Get_Data_Semigroup_Last_eqLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_eqLast.Do(func() {
		cache_Data_Semigroup_Last_eqLast = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Last_eqLast(dictEq_0_box)
})
	})
	return cache_Data_Semigroup_Last_eqLast
}

var cache_Data_Semigroup_Last_eq1Last gopurs_runtime.Value
var once_Data_Semigroup_Last_eq1Last sync.Once
func Get_Data_Semigroup_Last_eq1Last() gopurs_runtime.Value {
	once_Data_Semigroup_Last_eq1Last.Do(func() {
		cache_Data_Semigroup_Last_eq1Last = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Semigroup_Last_eq1Last
}

var cache_Data_Semigroup_Last_ord1Last gopurs_runtime.Value
var once_Data_Semigroup_Last_ord1Last sync.Once
func Get_Data_Semigroup_Last_ord1Last() gopurs_runtime.Value {
	once_Data_Semigroup_Last_ord1Last.Do(func() {
		cache_Data_Semigroup_Last_ord1Last = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_eq1Last()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_Data_Semigroup_Last_ord1Last
}

var cache_Data_Semigroup_Last_boundedLast gopurs_runtime.Value
var once_Data_Semigroup_Last_boundedLast sync.Once
func Get_Data_Semigroup_Last_boundedLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_boundedLast.Do(func() {
		cache_Data_Semigroup_Last_boundedLast = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Last_boundedLast(dictBounded_0_box)
})
	})
	return cache_Data_Semigroup_Last_boundedLast
}

var cache_Data_Semigroup_Last_applyLast gopurs_runtime.Value
var once_Data_Semigroup_Last_applyLast sync.Once
func Get_Data_Semigroup_Last_applyLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_applyLast.Do(func() {
		cache_Data_Semigroup_Last_applyLast = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_functorLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Semigroup_Last_applyLast
}

var cache_Data_Semigroup_Last_bindLast gopurs_runtime.Value
var once_Data_Semigroup_Last_bindLast sync.Once
func Get_Data_Semigroup_Last_bindLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_bindLast.Do(func() {
		cache_Data_Semigroup_Last_bindLast = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_applyLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Semigroup_Last_bindLast
}

var cache_Data_Semigroup_Last_applicativeLast gopurs_runtime.Value
var once_Data_Semigroup_Last_applicativeLast sync.Once
func Get_Data_Semigroup_Last_applicativeLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_applicativeLast.Do(func() {
		cache_Data_Semigroup_Last_applicativeLast = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_applyLast()
}), Get_Data_Semigroup_Last_Last())
	})
	return cache_Data_Semigroup_Last_applicativeLast
}

var cache_Data_Semigroup_Last_monadLast gopurs_runtime.Value
var once_Data_Semigroup_Last_monadLast sync.Once
func Get_Data_Semigroup_Last_monadLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_monadLast.Do(func() {
		cache_Data_Semigroup_Last_monadLast = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_applicativeLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_bindLast()
}))
	})
	return cache_Data_Semigroup_Last_monadLast
}

var cache_Data_Semigroup_Last_applicativeLast__4045440648 gopurs_runtime.Value
var once_Data_Semigroup_Last_applicativeLast__4045440648 sync.Once
func Get_Data_Semigroup_Last_applicativeLast__4045440648() gopurs_runtime.Value {
	once_Data_Semigroup_Last_applicativeLast__4045440648.Do(func() {
		cache_Data_Semigroup_Last_applicativeLast__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_applyLast()
}), Get_Data_Semigroup_Last_Last())
	})
	return cache_Data_Semigroup_Last_applicativeLast__4045440648
}

var cache_Data_Semigroup_Last_applyLast__3199351098 gopurs_runtime.Value
var once_Data_Semigroup_Last_applyLast__3199351098 sync.Once
func Get_Data_Semigroup_Last_applyLast__3199351098() gopurs_runtime.Value {
	once_Data_Semigroup_Last_applyLast__3199351098.Do(func() {
		cache_Data_Semigroup_Last_applyLast__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_functorLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Semigroup_Last_applyLast__3199351098
}

var cache_Data_Semigroup_Last_bindLast__329376103 gopurs_runtime.Value
var once_Data_Semigroup_Last_bindLast__329376103 sync.Once
func Get_Data_Semigroup_Last_bindLast__329376103() gopurs_runtime.Value {
	once_Data_Semigroup_Last_bindLast__329376103.Do(func() {
		cache_Data_Semigroup_Last_bindLast__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semigroup_Last_applyLast()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Semigroup_Last_bindLast__329376103
}

var cache_Data_Semigroup_Last_eq1Last__1905950174 gopurs_runtime.Value
var once_Data_Semigroup_Last_eq1Last__1905950174 sync.Once
func Get_Data_Semigroup_Last_eq1Last__1905950174() gopurs_runtime.Value {
	once_Data_Semigroup_Last_eq1Last__1905950174.Do(func() {
		cache_Data_Semigroup_Last_eq1Last__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Semigroup_Last_eq1Last__1905950174
}

var cache_Data_Semigroup_Last_functorLast__943655089 gopurs_runtime.Value
var once_Data_Semigroup_Last_functorLast__943655089 sync.Once
func Get_Data_Semigroup_Last_functorLast__943655089() gopurs_runtime.Value {
	once_Data_Semigroup_Last_functorLast__943655089.Do(func() {
		cache_Data_Semigroup_Last_functorLast__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Semigroup_Last_functorLast__943655089
}

var cache_Data_Semigroup_Last_semigroupLast__3224870556 gopurs_runtime.Value
var once_Data_Semigroup_Last_semigroupLast__3224870556 sync.Once
func Get_Data_Semigroup_Last_semigroupLast__3224870556() gopurs_runtime.Value {
	once_Data_Semigroup_Last_semigroupLast__3224870556.Do(func() {
		cache_Data_Semigroup_Last_semigroupLast__3224870556 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_Data_Semigroup_Last_semigroupLast__3224870556
}

var cache_Data_Semigroup_Last_semigroupLast__2108226578 gopurs_runtime.Value
var once_Data_Semigroup_Last_semigroupLast__2108226578 sync.Once
func Get_Data_Semigroup_Last_semigroupLast__2108226578() gopurs_runtime.Value {
	once_Data_Semigroup_Last_semigroupLast__2108226578.Do(func() {
		cache_Data_Semigroup_Last_semigroupLast__2108226578 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_Data_Semigroup_Last_semigroupLast__2108226578
}

var cache_Data_Semigroup_Last_semigroupLast__2246167645 gopurs_runtime.Value
var once_Data_Semigroup_Last_semigroupLast__2246167645 sync.Once
func Get_Data_Semigroup_Last_semigroupLast__2246167645() gopurs_runtime.Value {
	once_Data_Semigroup_Last_semigroupLast__2246167645.Do(func() {
		cache_Data_Semigroup_Last_semigroupLast__2246167645 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_Data_Semigroup_Last_semigroupLast__2246167645
}

func Call_Data_Semigroup_Last_Last(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Last_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Last ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Semigroup_Last_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Semigroup_Last_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Semigroup_Last_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}


