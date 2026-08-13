package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_List_ZipList_ZipList gopurs_runtime.Value
var once_Data_List_ZipList_ZipList sync.Once
func Get_Data_List_ZipList_ZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_ZipList.Do(func() {
		cache_Data_List_ZipList_ZipList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_ZipList(x_0_box)
})
	})
	return cache_Data_List_ZipList_ZipList
}

var cache_Data_List_ZipList_traversableZipList gopurs_runtime.Value
var once_Data_List_ZipList_traversableZipList sync.Once
func Get_Data_List_ZipList_traversableZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_traversableZipList.Do(func() {
		cache_Data_List_ZipList_traversableZipList = Get_Data_List_Lazy_Types_traversableList()
	})
	return cache_Data_List_ZipList_traversableZipList
}

var cache_Data_List_ZipList_showZipList gopurs_runtime.Value
var once_Data_List_ZipList_showZipList sync.Once
func Get_Data_List_ZipList_showZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_showZipList.Do(func() {
		cache_Data_List_ZipList_showZipList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_showZipList(dictShow_0_box)
})
	})
	return cache_Data_List_ZipList_showZipList
}

var cache_Data_List_ZipList_semigroupZipList gopurs_runtime.Value
var once_Data_List_ZipList_semigroupZipList sync.Once
func Get_Data_List_ZipList_semigroupZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_semigroupZipList.Do(func() {
		cache_Data_List_ZipList_semigroupZipList = Get_Data_List_Lazy_Types_semigroupList()
	})
	return cache_Data_List_ZipList_semigroupZipList
}

var cache_Data_List_ZipList_ordZipList gopurs_runtime.Value
var once_Data_List_ZipList_ordZipList sync.Once
func Get_Data_List_ZipList_ordZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_ordZipList.Do(func() {
		cache_Data_List_ZipList_ordZipList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_ordZipList(dictOrd_0_box)
})
	})
	return cache_Data_List_ZipList_ordZipList
}

var cache_Data_List_ZipList_newtypeZipList gopurs_runtime.Value
var once_Data_List_ZipList_newtypeZipList sync.Once
func Get_Data_List_ZipList_newtypeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_newtypeZipList.Do(func() {
		cache_Data_List_ZipList_newtypeZipList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_List_ZipList_newtypeZipList
}

var cache_Data_List_ZipList_monoidZipList gopurs_runtime.Value
var once_Data_List_ZipList_monoidZipList sync.Once
func Get_Data_List_ZipList_monoidZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_monoidZipList.Do(func() {
		cache_Data_List_ZipList_monoidZipList = Get_Data_List_Lazy_Types_monoidList()
	})
	return cache_Data_List_ZipList_monoidZipList
}

var cache_Data_List_ZipList_functorZipList gopurs_runtime.Value
var once_Data_List_ZipList_functorZipList sync.Once
func Get_Data_List_ZipList_functorZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_functorZipList.Do(func() {
		cache_Data_List_ZipList_functorZipList = Get_Data_List_Lazy_Types_functorList()
	})
	return cache_Data_List_ZipList_functorZipList
}

var cache_Data_List_ZipList_foldableZipList gopurs_runtime.Value
var once_Data_List_ZipList_foldableZipList sync.Once
func Get_Data_List_ZipList_foldableZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_foldableZipList.Do(func() {
		cache_Data_List_ZipList_foldableZipList = Get_Data_List_Lazy_Types_foldableList()
	})
	return cache_Data_List_ZipList_foldableZipList
}

var cache_Data_List_ZipList_eqZipList gopurs_runtime.Value
var once_Data_List_ZipList_eqZipList sync.Once
func Get_Data_List_ZipList_eqZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_eqZipList.Do(func() {
		cache_Data_List_ZipList_eqZipList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_eqZipList(dictEq_0_box)
})
	})
	return cache_Data_List_ZipList_eqZipList
}

var cache_Data_List_ZipList_applyZipList gopurs_runtime.Value
var once_Data_List_ZipList_applyZipList sync.Once
func Get_Data_List_ZipList_applyZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_applyZipList.Do(func() {
		cache_Data_List_ZipList_applyZipList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_List_Lazy_zipWith(), Get_Data_Function_apply(), v_0, v1_1)
})
}))
	})
	return cache_Data_List_ZipList_applyZipList
}

var cache_Data_List_ZipList_zipListIsNotBind gopurs_runtime.Value
var once_Data_List_ZipList_zipListIsNotBind sync.Once
func Get_Data_List_ZipList_zipListIsNotBind() gopurs_runtime.Value {
	once_Data_List_ZipList_zipListIsNotBind.Do(func() {
		cache_Data_List_ZipList_zipListIsNotBind = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ZipList_zipListIsNotBind(_dollar__unused_0_box)
})
	})
	return cache_Data_List_ZipList_zipListIsNotBind
}

var cache_Data_List_ZipList_applicativeZipList gopurs_runtime.Value
var once_Data_List_ZipList_applicativeZipList sync.Once
func Get_Data_List_ZipList_applicativeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_applicativeZipList.Do(func() {
		cache_Data_List_ZipList_applicativeZipList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_applyZipList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Lazy_repeat(), x_0)
}))
	})
	return cache_Data_List_ZipList_applicativeZipList
}

var cache_Data_List_ZipList_altZipList gopurs_runtime.Value
var once_Data_List_ZipList_altZipList sync.Once
func Get_Data_List_ZipList_altZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_altZipList.Do(func() {
		cache_Data_List_ZipList_altZipList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), v_0, gopurs_runtime.Apply2(Get_Data_List_Lazy_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_List_Lazy_length(), v_0).IntVal), v1_1))
})
}))
	})
	return cache_Data_List_ZipList_altZipList
}

var cache_Data_List_ZipList_plusZipList gopurs_runtime.Value
var once_Data_List_ZipList_plusZipList sync.Once
func Get_Data_List_ZipList_plusZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_plusZipList.Do(func() {
		cache_Data_List_ZipList_plusZipList = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_altZipList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_ZipList_plusZipList
}

var cache_Data_List_ZipList_alternativeZipList gopurs_runtime.Value
var once_Data_List_ZipList_alternativeZipList sync.Once
func Get_Data_List_ZipList_alternativeZipList() gopurs_runtime.Value {
	once_Data_List_ZipList_alternativeZipList.Do(func() {
		cache_Data_List_ZipList_alternativeZipList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_applicativeZipList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_plusZipList()
}))
	})
	return cache_Data_List_ZipList_alternativeZipList
}

var cache_Data_List_ZipList_altZipList__3296309911 gopurs_runtime.Value
var once_Data_List_ZipList_altZipList__3296309911 sync.Once
func Get_Data_List_ZipList_altZipList__3296309911() gopurs_runtime.Value {
	once_Data_List_ZipList_altZipList__3296309911.Do(func() {
		cache_Data_List_ZipList_altZipList__3296309911 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), v_0, gopurs_runtime.Apply2(Get_Data_List_Lazy_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_List_Lazy_length(), v_0).IntVal), v1_1))
})
}))
	})
	return cache_Data_List_ZipList_altZipList__3296309911
}

var cache_Data_List_ZipList_applicativeZipList__37190504 gopurs_runtime.Value
var once_Data_List_ZipList_applicativeZipList__37190504 sync.Once
func Get_Data_List_ZipList_applicativeZipList__37190504() gopurs_runtime.Value {
	once_Data_List_ZipList_applicativeZipList__37190504.Do(func() {
		cache_Data_List_ZipList_applicativeZipList__37190504 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_applyZipList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Lazy_repeat(), x_0)
}))
	})
	return cache_Data_List_ZipList_applicativeZipList__37190504
}

var cache_Data_List_ZipList_applyZipList__1358886895 gopurs_runtime.Value
var once_Data_List_ZipList_applyZipList__1358886895 sync.Once
func Get_Data_List_ZipList_applyZipList__1358886895() gopurs_runtime.Value {
	once_Data_List_ZipList_applyZipList__1358886895.Do(func() {
		cache_Data_List_ZipList_applyZipList__1358886895 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_List_Lazy_zipWith(), Get_Data_Function_apply(), v_0, v1_1)
})
}))
	})
	return cache_Data_List_ZipList_applyZipList__1358886895
}

var cache_Data_List_ZipList_functorZipList__3996674161 gopurs_runtime.Value
var once_Data_List_ZipList_functorZipList__3996674161 sync.Once
func Get_Data_List_ZipList_functorZipList__3996674161() gopurs_runtime.Value {
	once_Data_List_ZipList_functorZipList__3996674161.Do(func() {
		cache_Data_List_ZipList_functorZipList__3996674161 = Get_Data_List_Lazy_Types_functorList()
	})
	return cache_Data_List_ZipList_functorZipList__3996674161
}

var cache_Data_List_ZipList_plusZipList__3460472018 gopurs_runtime.Value
var once_Data_List_ZipList_plusZipList__3460472018 sync.Once
func Get_Data_List_ZipList_plusZipList__3460472018() gopurs_runtime.Value {
	once_Data_List_ZipList_plusZipList__3460472018.Do(func() {
		cache_Data_List_ZipList_plusZipList__3460472018 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_altZipList()
}), Get_Data_List_Lazy_Types_nil())
	})
	return cache_Data_List_ZipList_plusZipList__3460472018
}

func Call_Data_List_ZipList_ZipList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_ZipList_showZipList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList_1_0 -> *Constructor_Data_Show_Show
showList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_showList(), dictShow_0))
_ = showList_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(ZipList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showList_1_0.V0), v_2).StrVal())) + (")"))
}))
}

func Call_Data_List_ZipList_ordZipList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_Types_ordList(), dictOrd_0)
}

func Call_Data_List_ZipList_eqZipList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_eq1List(), "eq1"), dictEq_0))
}

func Call_Data_List_ZipList_zipListIsNotBind(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_ZipList_applyZipList()
}), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("bind: unreachable"))
})))
}


