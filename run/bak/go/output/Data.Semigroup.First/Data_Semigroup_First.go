package Data_Semigroup_First

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_First gopurs_runtime.Value
var once_First sync.Once
func Get_First() gopurs_runtime.Value {
	once_First.Do(func() {
		cache_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_First(x_0_box)
})
	})
	return cache_First
}

var cache_showFirst gopurs_runtime.Value
var once_showFirst sync.Once
func Get_showFirst() gopurs_runtime.Value {
	once_showFirst.Do(func() {
		cache_showFirst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showFirst(dictShow_0_box)
})
	})
	return cache_showFirst
}

var cache_semigroupFirst gopurs_runtime.Value
var once_semigroupFirst sync.Once
func Get_semigroupFirst() gopurs_runtime.Value {
	once_semigroupFirst.Do(func() {
		cache_semigroupFirst = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_semigroupFirst
}

var cache_ordFirst gopurs_runtime.Value
var once_ordFirst sync.Once
func Get_ordFirst() gopurs_runtime.Value {
	once_ordFirst.Do(func() {
		cache_ordFirst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordFirst(dictOrd_0_box)
})
	})
	return cache_ordFirst
}

var cache_functorFirst gopurs_runtime.Value
var once_functorFirst sync.Once
func Get_functorFirst() gopurs_runtime.Value {
	once_functorFirst.Do(func() {
		cache_functorFirst = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorFirst
}

var cache_eqFirst gopurs_runtime.Value
var once_eqFirst sync.Once
func Get_eqFirst() gopurs_runtime.Value {
	once_eqFirst.Do(func() {
		cache_eqFirst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqFirst(dictEq_0_box)
})
	})
	return cache_eqFirst
}

var cache_eq1First gopurs_runtime.Value
var once_eq1First sync.Once
func Get_eq1First() gopurs_runtime.Value {
	once_eq1First.Do(func() {
		cache_eq1First = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1First
}

var cache_ord1First gopurs_runtime.Value
var once_ord1First sync.Once
func Get_ord1First() gopurs_runtime.Value {
	once_ord1First.Do(func() {
		cache_ord1First = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1First()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1First
}

var cache_boundedFirst gopurs_runtime.Value
var once_boundedFirst sync.Once
func Get_boundedFirst() gopurs_runtime.Value {
	once_boundedFirst.Do(func() {
		cache_boundedFirst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedFirst(dictBounded_0_box)
})
	})
	return cache_boundedFirst
}

var cache_applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		cache_applyFirst = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorFirst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyFirst
}

var cache_bindFirst gopurs_runtime.Value
var once_bindFirst sync.Once
func Get_bindFirst() gopurs_runtime.Value {
	once_bindFirst.Do(func() {
		cache_bindFirst = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFirst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindFirst
}

var cache_applicativeFirst gopurs_runtime.Value
var once_applicativeFirst sync.Once
func Get_applicativeFirst() gopurs_runtime.Value {
	once_applicativeFirst.Do(func() {
		cache_applicativeFirst = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFirst()
}), Get_First())
	})
	return cache_applicativeFirst
}

var cache_monadFirst gopurs_runtime.Value
var once_monadFirst sync.Once
func Get_monadFirst() gopurs_runtime.Value {
	once_monadFirst.Do(func() {
		cache_monadFirst = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindFirst()
}))
	})
	return cache_monadFirst
}

var cache_applicativeFirst__4045440648 gopurs_runtime.Value
var once_applicativeFirst__4045440648 sync.Once
func Get_applicativeFirst__4045440648() gopurs_runtime.Value {
	once_applicativeFirst__4045440648.Do(func() {
		cache_applicativeFirst__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFirst()
}), Get_First())
	})
	return cache_applicativeFirst__4045440648
}

var cache_applyFirst__3199351098 gopurs_runtime.Value
var once_applyFirst__3199351098 sync.Once
func Get_applyFirst__3199351098() gopurs_runtime.Value {
	once_applyFirst__3199351098.Do(func() {
		cache_applyFirst__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorFirst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyFirst__3199351098
}

var cache_bindFirst__329376103 gopurs_runtime.Value
var once_bindFirst__329376103 sync.Once
func Get_bindFirst__329376103() gopurs_runtime.Value {
	once_bindFirst__329376103.Do(func() {
		cache_bindFirst__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyFirst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindFirst__329376103
}

var cache_eq1First__1905950174 gopurs_runtime.Value
var once_eq1First__1905950174 sync.Once
func Get_eq1First__1905950174() gopurs_runtime.Value {
	once_eq1First__1905950174.Do(func() {
		cache_eq1First__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1First__1905950174
}

var cache_functorFirst__943655089 gopurs_runtime.Value
var once_functorFirst__943655089 sync.Once
func Get_functorFirst__943655089() gopurs_runtime.Value {
	once_functorFirst__943655089.Do(func() {
		cache_functorFirst__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorFirst__943655089
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
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

func Call_First(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showFirst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(First "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_ordFirst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_eqFirst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedFirst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
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


