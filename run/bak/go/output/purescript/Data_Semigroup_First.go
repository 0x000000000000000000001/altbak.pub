package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Semigroup_First_First gopurs_runtime.Value
var once_Data_Semigroup_First_First sync.Once
func Get_Data_Semigroup_First_First() gopurs_runtime.Value {
	once_Data_Semigroup_First_First.Do(func() {
		cache_Data_Semigroup_First_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_First_First(x_0_box)
})
	})
	return cache_Data_Semigroup_First_First
}

var cache_Data_Semigroup_First_showFirst gopurs_runtime.Value
var once_Data_Semigroup_First_showFirst sync.Once
func Get_Data_Semigroup_First_showFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_showFirst.Do(func() {
		cache_Data_Semigroup_First_showFirst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_First_showFirst(dictShow_0_box)
})
	})
	return cache_Data_Semigroup_First_showFirst
}

var cache_Data_Semigroup_First_semigroupFirst gopurs_runtime.Value
var once_Data_Semigroup_First_semigroupFirst sync.Once
func Get_Data_Semigroup_First_semigroupFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_semigroupFirst.Do(func() {
		cache_Data_Semigroup_First_semigroupFirst = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_Data_Semigroup_First_semigroupFirst
}

var cache_Data_Semigroup_First_ordFirst gopurs_runtime.Value
var once_Data_Semigroup_First_ordFirst sync.Once
func Get_Data_Semigroup_First_ordFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_ordFirst.Do(func() {
		cache_Data_Semigroup_First_ordFirst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_First_ordFirst(dictOrd_0_box)
})
	})
	return cache_Data_Semigroup_First_ordFirst
}

var cache_Data_Semigroup_First_functorFirst gopurs_runtime.Value
var once_Data_Semigroup_First_functorFirst sync.Once
func Get_Data_Semigroup_First_functorFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_functorFirst.Do(func() {
		cache_Data_Semigroup_First_functorFirst = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Semigroup_First_functorFirst
}

var cache_Data_Semigroup_First_eqFirst gopurs_runtime.Value
var once_Data_Semigroup_First_eqFirst sync.Once
func Get_Data_Semigroup_First_eqFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_eqFirst.Do(func() {
		cache_Data_Semigroup_First_eqFirst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_First_eqFirst(dictEq_0_box)
})
	})
	return cache_Data_Semigroup_First_eqFirst
}

var cache_Data_Semigroup_First_eq1First gopurs_runtime.Value
var once_Data_Semigroup_First_eq1First sync.Once
func Get_Data_Semigroup_First_eq1First() gopurs_runtime.Value {
	once_Data_Semigroup_First_eq1First.Do(func() {
		cache_Data_Semigroup_First_eq1First = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Semigroup_First_eq1First
}

var cache_Data_Semigroup_First_ord1First gopurs_runtime.Value
var once_Data_Semigroup_First_ord1First sync.Once
func Get_Data_Semigroup_First_ord1First() gopurs_runtime.Value {
	once_Data_Semigroup_First_ord1First.Do(func() {
		cache_Data_Semigroup_First_ord1First = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_1, "eq")
}))
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_Data_Semigroup_First_ord1First
}

var cache_Data_Semigroup_First_boundedFirst gopurs_runtime.Value
var once_Data_Semigroup_First_boundedFirst sync.Once
func Get_Data_Semigroup_First_boundedFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_boundedFirst.Do(func() {
		cache_Data_Semigroup_First_boundedFirst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_First_boundedFirst(dictBounded_0_box)
})
	})
	return cache_Data_Semigroup_First_boundedFirst
}

var cache_Data_Semigroup_First_applyFirst gopurs_runtime.Value
var once_Data_Semigroup_First_applyFirst sync.Once
func Get_Data_Semigroup_First_applyFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_applyFirst.Do(func() {
		cache_Data_Semigroup_First_applyFirst = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Semigroup_First_applyFirst
}

var cache_Data_Semigroup_First_bindFirst gopurs_runtime.Value
var once_Data_Semigroup_First_bindFirst sync.Once
func Get_Data_Semigroup_First_bindFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_bindFirst.Do(func() {
		cache_Data_Semigroup_First_bindFirst = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Semigroup_First_bindFirst
}

var cache_Data_Semigroup_First_applicativeFirst gopurs_runtime.Value
var once_Data_Semigroup_First_applicativeFirst sync.Once
func Get_Data_Semigroup_First_applicativeFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_applicativeFirst.Do(func() {
		cache_Data_Semigroup_First_applicativeFirst = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Semigroup_First_applicativeFirst
}

var cache_Data_Semigroup_First_monadFirst gopurs_runtime.Value
var once_Data_Semigroup_First_monadFirst sync.Once
func Get_Data_Semigroup_First_monadFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_monadFirst.Do(func() {
		cache_Data_Semigroup_First_monadFirst = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, v_1)
})
}))
}))
	})
	return cache_Data_Semigroup_First_monadFirst
}

var cache_Data_Semigroup_First_applicativeFirst__4045440648 gopurs_runtime.Value
var once_Data_Semigroup_First_applicativeFirst__4045440648 sync.Once
func Get_Data_Semigroup_First_applicativeFirst__4045440648() gopurs_runtime.Value {
	once_Data_Semigroup_First_applicativeFirst__4045440648.Do(func() {
		cache_Data_Semigroup_First_applicativeFirst__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Semigroup_First_applicativeFirst__4045440648
}

var cache_Data_Semigroup_First_applyFirst__3199351098 gopurs_runtime.Value
var once_Data_Semigroup_First_applyFirst__3199351098 sync.Once
func Get_Data_Semigroup_First_applyFirst__3199351098() gopurs_runtime.Value {
	once_Data_Semigroup_First_applyFirst__3199351098.Do(func() {
		cache_Data_Semigroup_First_applyFirst__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Semigroup_First_applyFirst__3199351098
}

var cache_Data_Semigroup_First_bindFirst__329376103 gopurs_runtime.Value
var once_Data_Semigroup_First_bindFirst__329376103 sync.Once
func Get_Data_Semigroup_First_bindFirst__329376103() gopurs_runtime.Value {
	once_Data_Semigroup_First_bindFirst__329376103.Do(func() {
		cache_Data_Semigroup_First_bindFirst__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Semigroup_First_bindFirst__329376103
}

var cache_Data_Semigroup_First_eq1First__1905950174 gopurs_runtime.Value
var once_Data_Semigroup_First_eq1First__1905950174 sync.Once
func Get_Data_Semigroup_First_eq1First__1905950174() gopurs_runtime.Value {
	once_Data_Semigroup_First_eq1First__1905950174.Do(func() {
		cache_Data_Semigroup_First_eq1First__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Semigroup_First_eq1First__1905950174
}

var cache_Data_Semigroup_First_functorFirst__943655089 gopurs_runtime.Value
var once_Data_Semigroup_First_functorFirst__943655089 sync.Once
func Get_Data_Semigroup_First_functorFirst__943655089() gopurs_runtime.Value {
	once_Data_Semigroup_First_functorFirst__943655089.Do(func() {
		cache_Data_Semigroup_First_functorFirst__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Semigroup_First_functorFirst__943655089
}

func Call_Data_Semigroup_First_First(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_First_showFirst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(First ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Semigroup_First_ordFirst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Semigroup_First_eqFirst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Semigroup_First_boundedFirst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}


