package Data_List_ZipList

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Partial "gopurs/output/Partial"
	unsafe "unsafe"
)

var cache_ZipList gopurs_runtime.Value
var once_ZipList sync.Once
func Get_ZipList() gopurs_runtime.Value {
	once_ZipList.Do(func() {
		cache_ZipList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ZipList(x_0_box)
})
	})
	return cache_ZipList
}

var cache_traversableZipList gopurs_runtime.Value
var once_traversableZipList sync.Once
func Get_traversableZipList() gopurs_runtime.Value {
	once_traversableZipList.Do(func() {
		cache_traversableZipList = pkg_Data_List_Lazy_Types.Get_traversableList()
	})
	return cache_traversableZipList
}

var cache_showZipList gopurs_runtime.Value
var once_showZipList sync.Once
func Get_showZipList() gopurs_runtime.Value {
	once_showZipList.Do(func() {
		cache_showZipList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showZipList(dictShow_0_box)
})
	})
	return cache_showZipList
}

var cache_semigroupZipList gopurs_runtime.Value
var once_semigroupZipList sync.Once
func Get_semigroupZipList() gopurs_runtime.Value {
	once_semigroupZipList.Do(func() {
		cache_semigroupZipList = pkg_Data_List_Lazy_Types.Get_semigroupList()
	})
	return cache_semigroupZipList
}

var cache_ordZipList gopurs_runtime.Value
var once_ordZipList sync.Once
func Get_ordZipList() gopurs_runtime.Value {
	once_ordZipList.Do(func() {
		cache_ordZipList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordZipList(dictOrd_0_box)
})
	})
	return cache_ordZipList
}

var cache_newtypeZipList gopurs_runtime.Value
var once_newtypeZipList sync.Once
func Get_newtypeZipList() gopurs_runtime.Value {
	once_newtypeZipList.Do(func() {
		cache_newtypeZipList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeZipList
}

var cache_monoidZipList gopurs_runtime.Value
var once_monoidZipList sync.Once
func Get_monoidZipList() gopurs_runtime.Value {
	once_monoidZipList.Do(func() {
		cache_monoidZipList = pkg_Data_List_Lazy_Types.Get_monoidList()
	})
	return cache_monoidZipList
}

var cache_functorZipList gopurs_runtime.Value
var once_functorZipList sync.Once
func Get_functorZipList() gopurs_runtime.Value {
	once_functorZipList.Do(func() {
		cache_functorZipList = pkg_Data_List_Lazy_Types.Get_functorList()
	})
	return cache_functorZipList
}

var cache_foldableZipList gopurs_runtime.Value
var once_foldableZipList sync.Once
func Get_foldableZipList() gopurs_runtime.Value {
	once_foldableZipList.Do(func() {
		cache_foldableZipList = pkg_Data_List_Lazy_Types.Get_foldableList()
	})
	return cache_foldableZipList
}

var cache_eqZipList gopurs_runtime.Value
var once_eqZipList sync.Once
func Get_eqZipList() gopurs_runtime.Value {
	once_eqZipList.Do(func() {
		cache_eqZipList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqZipList(dictEq_0_box)
})
	})
	return cache_eqZipList
}

var cache_applyZipList gopurs_runtime.Value
var once_applyZipList sync.Once
func Get_applyZipList() gopurs_runtime.Value {
	once_applyZipList.Do(func() {
		cache_applyZipList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_List_Lazy.Get_zipWith(), pkg_Data_Function.Get_apply(), v_0, v1_1)
}))
	})
	return cache_applyZipList
}

var cache_zipListIsNotBind gopurs_runtime.Value
var once_zipListIsNotBind sync.Once
func Get_zipListIsNotBind() gopurs_runtime.Value {
	once_zipListIsNotBind.Do(func() {
		cache_zipListIsNotBind = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipListIsNotBind(_dollar__unused_0_box)
})
	})
	return cache_zipListIsNotBind
}

var cache_applicativeZipList gopurs_runtime.Value
var once_applicativeZipList sync.Once
func Get_applicativeZipList() gopurs_runtime.Value {
	once_applicativeZipList.Do(func() {
		cache_applicativeZipList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyZipList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
_ = go__1_0
go__1_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons{x_0, go__1_0})}
}))
return go__1_0
}))
	})
	return cache_applicativeZipList
}

var cache_altZipList gopurs_runtime.Value
var once_altZipList sync.Once
func Get_altZipList() gopurs_runtime.Value {
	once_altZipList.Do(func() {
		cache_altZipList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), v_0, gopurs_runtime.Apply2(pkg_Data_List_Lazy.Get_drop(), gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), v_0), v1_1))
}))
	})
	return cache_altZipList
}

var cache_plusZipList gopurs_runtime.Value
var once_plusZipList sync.Once
func Get_plusZipList() gopurs_runtime.Value {
	once_plusZipList.Do(func() {
		cache_plusZipList = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altZipList()
}), gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monoidList(), "mempty"))
	})
	return cache_plusZipList
}

var cache_alternativeZipList gopurs_runtime.Value
var once_alternativeZipList sync.Once
func Get_alternativeZipList() gopurs_runtime.Value {
	once_alternativeZipList.Do(func() {
		cache_alternativeZipList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeZipList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusZipList()
}))
	})
	return cache_alternativeZipList
}

func Call_ZipList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showZipList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(ZipList "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_showList(), dictShow_0), "show"), v_1), gopurs_runtime.Str(")")))
}))
}

func Call_ordZipList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_ordList(), dictOrd_0)
}

func Call_eqZipList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_eq1List(), "eq1"), dictEq_0))
}

func Call_zipListIsNotBind(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyZipList()
}), gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("bind: unreachable")))
}


