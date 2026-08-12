package Data_List_ZipList

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Lazy "gopurs/output/Control.Lazy"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Partial "gopurs/output/Partial"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_List_Lazy.Get_zipWith(), pkg_Data_Function.Get_apply(), v_0, v1_1)
})
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
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
go__go_1_0_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_0})}
}))
}))
return go__go_1_0_0
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), v_0, gopurs_runtime.Apply2(pkg_Data_List_Lazy.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), v_0).IntVal), v1_1))
})
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

var cache_pure__189931222 gopurs_runtime.Value
var once_pure__189931222 sync.Once
func Get_pure__189931222() gopurs_runtime.Value {
	once_pure__189931222.Do(func() {
		cache_pure__189931222 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__189931222(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__189931222
}

var cache_pure__3236307030 gopurs_runtime.Value
var once_pure__3236307030 sync.Once
func Get_pure__3236307030() gopurs_runtime.Value {
	once_pure__3236307030.Do(func() {
		cache_pure__3236307030 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3236307030(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3236307030
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__355615152 gopurs_runtime.Value
var once_pure__355615152 sync.Once
func Get_pure__355615152() gopurs_runtime.Value {
	once_pure__355615152.Do(func() {
		cache_pure__355615152 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__355615152(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__355615152
}

var cache_apply__4203183626 gopurs_runtime.Value
var once_apply__4203183626 sync.Once
func Get_apply__4203183626() gopurs_runtime.Value {
	once_apply__4203183626.Do(func() {
		cache_apply__4203183626 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__4203183626(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__4203183626
}

var cache_apply__2962221386 gopurs_runtime.Value
var once_apply__2962221386 sync.Once
func Get_apply__2962221386() gopurs_runtime.Value {
	once_apply__2962221386.Do(func() {
		cache_apply__2962221386 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2962221386(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2962221386
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__1851858028 gopurs_runtime.Value
var once_apply__1851858028 sync.Once
func Get_apply__1851858028() gopurs_runtime.Value {
	once_apply__1851858028.Do(func() {
		cache_apply__1851858028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1851858028(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__1851858028
}

var cache_apply__3620326986 gopurs_runtime.Value
var once_apply__3620326986 sync.Once
func Get_apply__3620326986() gopurs_runtime.Value {
	once_apply__3620326986.Do(func() {
		cache_apply__3620326986 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__3620326986(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_apply__3620326986
}

var cache_defer__3258767445 gopurs_runtime.Value
var once_defer__3258767445 sync.Once
func Get_defer__3258767445() gopurs_runtime.Value {
	once_defer__3258767445.Do(func() {
		cache_defer__3258767445 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__3258767445(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__3258767445
}

var cache_defer__3967925939 gopurs_runtime.Value
var once_defer__3967925939 sync.Once
func Get_defer__3967925939() gopurs_runtime.Value {
	once_defer__3967925939.Do(func() {
		cache_defer__3967925939 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__3967925939(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__3967925939
}

var cache_fix__1475205859 gopurs_runtime.Value
var once_fix__1475205859 sync.Once
func Get_fix__1475205859() gopurs_runtime.Value {
	once_fix__1475205859.Do(func() {
		cache_fix__1475205859 = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fix__1475205859(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_0_box), f_1_box)
})
	})
	return cache_fix__1475205859
}

var cache_fix__3570066147 gopurs_runtime.Value
var once_fix__3570066147 sync.Once
func Get_fix__3570066147() gopurs_runtime.Value {
	once_fix__3570066147.Do(func() {
		cache_fix__3570066147 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fix__3570066147(f_0_box)
})
	})
	return cache_fix__3570066147
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

var cache_foldl__1422885860 gopurs_runtime.Value
var once_foldl__1422885860 sync.Once
func Get_foldl__1422885860() gopurs_runtime.Value {
	once_foldl__1422885860.Do(func() {
		cache_foldl__1422885860 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1422885860(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__1422885860
}

var cache_foldl__267332164 gopurs_runtime.Value
var once_foldl__267332164 sync.Once
func Get_foldl__267332164() gopurs_runtime.Value {
	once_foldl__267332164.Do(func() {
		cache_foldl__267332164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__267332164(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__267332164
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__524683195 gopurs_runtime.Value
var once_foldl__524683195 sync.Once
func Get_foldl__524683195() gopurs_runtime.Value {
	once_foldl__524683195.Do(func() {
		cache_foldl__524683195 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__524683195(op_0_box)
})
	})
	return cache_foldl__524683195
}

var cache_foldl__3379885725 gopurs_runtime.Value
var once_foldl__3379885725 sync.Once
func Get_foldl__3379885725() gopurs_runtime.Value {
	once_foldl__3379885725.Do(func() {
		cache_foldl__3379885725 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3379885725(op_0_box)
})
	})
	return cache_foldl__3379885725
}

var cache_foldl__1985071933 gopurs_runtime.Value
var once_foldl__1985071933 sync.Once
func Get_foldl__1985071933() gopurs_runtime.Value {
	once_foldl__1985071933.Do(func() {
		cache_foldl__1985071933 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1985071933(op_0_box)
})
	})
	return cache_foldl__1985071933
}

var cache_foldl__536153533 gopurs_runtime.Value
var once_foldl__536153533 sync.Once
func Get_foldl__536153533() gopurs_runtime.Value {
	once_foldl__536153533.Do(func() {
		cache_foldl__536153533 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__536153533(op_0_box)
})
	})
	return cache_foldl__536153533
}

var cache_foldr__2111289130 gopurs_runtime.Value
var once_foldr__2111289130 sync.Once
func Get_foldr__2111289130() gopurs_runtime.Value {
	once_foldr__2111289130.Do(func() {
		cache_foldr__2111289130 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2111289130(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2111289130
}

var cache_foldr__926146538 gopurs_runtime.Value
var once_foldr__926146538 sync.Once
func Get_foldr__926146538() gopurs_runtime.Value {
	once_foldr__926146538.Do(func() {
		cache_foldr__926146538 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__926146538(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__926146538
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__1985071933 gopurs_runtime.Value
var once_foldr__1985071933 sync.Once
func Get_foldr__1985071933() gopurs_runtime.Value {
	once_foldr__1985071933.Do(func() {
		cache_foldr__1985071933 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__1985071933(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_foldr__1985071933
}

var cache_foldr__2389967549 gopurs_runtime.Value
var once_foldr__2389967549 sync.Once
func Get_foldr__2389967549() gopurs_runtime.Value {
	once_foldr__2389967549.Do(func() {
		cache_foldr__2389967549 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2389967549(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_foldr__2389967549
}

var cache_apply__458711162 gopurs_runtime.Value
var once_apply__458711162 sync.Once
func Get_apply__458711162() gopurs_runtime.Value {
	once_apply__458711162.Do(func() {
		cache_apply__458711162 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__458711162(f_0_box, x_1_box)
})
	})
	return cache_apply__458711162
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__3658931456 gopurs_runtime.Value
var once_flip__3658931456 sync.Once
func Get_flip__3658931456() gopurs_runtime.Value {
	once_flip__3658931456.Do(func() {
		cache_flip__3658931456 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3658931456(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3658931456
}

var cache_flip__3019832928 gopurs_runtime.Value
var once_flip__3019832928 sync.Once
func Get_flip__3019832928() gopurs_runtime.Value {
	once_flip__3019832928.Do(func() {
		cache_flip__3019832928 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3019832928(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3019832928
}

var cache_flip__2175652032 gopurs_runtime.Value
var once_flip__2175652032 sync.Once
func Get_flip__2175652032() gopurs_runtime.Value {
	once_flip__2175652032.Do(func() {
		cache_flip__2175652032 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2175652032(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2175652032
}

var cache_flip__848188896 gopurs_runtime.Value
var once_flip__848188896 sync.Once
func Get_flip__848188896() gopurs_runtime.Value {
	once_flip__848188896.Do(func() {
		cache_flip__848188896 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__848188896(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__848188896
}

var cache_flip__913470112 gopurs_runtime.Value
var once_flip__913470112 sync.Once
func Get_flip__913470112() gopurs_runtime.Value {
	once_flip__913470112.Do(func() {
		cache_flip__913470112 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__913470112(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__913470112
}

var cache_map__2665381605 gopurs_runtime.Value
var once_map__2665381605 sync.Once
func Get_map__2665381605() gopurs_runtime.Value {
	once_map__2665381605.Do(func() {
		cache_map__2665381605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2665381605(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2665381605
}

var cache_map__1542634789 gopurs_runtime.Value
var once_map__1542634789 sync.Once
func Get_map__1542634789() gopurs_runtime.Value {
	once_map__1542634789.Do(func() {
		cache_map__1542634789 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1542634789(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1542634789
}

var cache_map__2675323109 gopurs_runtime.Value
var once_map__2675323109 sync.Once
func Get_map__2675323109() gopurs_runtime.Value {
	once_map__2675323109.Do(func() {
		cache_map__2675323109 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2675323109(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2675323109
}

var cache_map__3871729957 gopurs_runtime.Value
var once_map__3871729957 sync.Once
func Get_map__3871729957() gopurs_runtime.Value {
	once_map__3871729957.Do(func() {
		cache_map__3871729957 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3871729957(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3871729957
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

var cache_map__831829748 gopurs_runtime.Value
var once_map__831829748 sync.Once
func Get_map__831829748() gopurs_runtime.Value {
	once_map__831829748.Do(func() {
		cache_map__831829748 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__831829748(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__831829748
}

var cache_map__1510739772 gopurs_runtime.Value
var once_map__1510739772 sync.Once
func Get_map__1510739772() gopurs_runtime.Value {
	once_map__1510739772.Do(func() {
		cache_map__1510739772 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1510739772(f_0_box, l_1_box)
})
	})
	return cache_map__1510739772
}

var cache_map__3565923196 gopurs_runtime.Value
var once_map__3565923196 sync.Once
func Get_map__3565923196() gopurs_runtime.Value {
	once_map__3565923196.Do(func() {
		cache_map__3565923196 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3565923196(f_0_box, l_1_box)
})
	})
	return cache_map__3565923196
}

var cache_map__109003388 gopurs_runtime.Value
var once_map__109003388 sync.Once
func Get_map__109003388() gopurs_runtime.Value {
	once_map__109003388.Do(func() {
		cache_map__109003388 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__109003388(f_0_box, l_1_box)
})
	})
	return cache_map__109003388
}

var cache_map__2156385148 gopurs_runtime.Value
var once_map__2156385148 sync.Once
func Get_map__2156385148() gopurs_runtime.Value {
	once_map__2156385148.Do(func() {
		cache_map__2156385148 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2156385148(f_0_box, l_1_box)
})
	})
	return cache_map__2156385148
}

var cache_applyLazy__879424557 gopurs_runtime.Value
var once_applyLazy__879424557 sync.Once
func Get_applyLazy__879424557() gopurs_runtime.Value {
	once_applyLazy__879424557.Do(func() {
		cache_applyLazy__879424557 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Lazy.Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Lazy.Get_force(), f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
}))
})
}))
	})
	return cache_applyLazy__879424557
}

var cache_functorLazy__491347738 gopurs_runtime.Value
var once_functorLazy__491347738 sync.Once
func Get_functorLazy__491347738() gopurs_runtime.Value {
	once_functorLazy__491347738.Do(func() {
		cache_functorLazy__491347738 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__491347738
}

var cache_functorLazy__3988504945 gopurs_runtime.Value
var once_functorLazy__3988504945 sync.Once
func Get_functorLazy__3988504945() gopurs_runtime.Value {
	once_functorLazy__3988504945.Do(func() {
		cache_functorLazy__3988504945 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__3988504945
}

var cache_cons__716923058 gopurs_runtime.Value
var once_cons__716923058 sync.Once
func Get_cons__716923058() gopurs_runtime.Value {
	once_cons__716923058.Do(func() {
		cache_cons__716923058 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__716923058(x_0_box, xs_1_box)
})
	})
	return cache_cons__716923058
}

var cache_cons__720046150 gopurs_runtime.Value
var once_cons__720046150 sync.Once
func Get_cons__720046150() gopurs_runtime.Value {
	once_cons__720046150.Do(func() {
		cache_cons__720046150 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__720046150(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](x_0_box), xs_1_box)
})
	})
	return cache_cons__720046150
}

var cache_cons__2305074921 gopurs_runtime.Value
var once_cons__2305074921 sync.Once
func Get_cons__2305074921() gopurs_runtime.Value {
	once_cons__2305074921.Do(func() {
		cache_cons__2305074921 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons__2305074921(x_0_box, xs_1_box)
})
	})
	return cache_cons__2305074921
}

var cache_foldableList__4097915271 gopurs_runtime.Value
var once_foldableList__4097915271 sync.Once
func Get_foldableList__4097915271() gopurs_runtime.Value {
	once_foldableList__4097915271.Do(func() {
		cache_foldableList__4097915271 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_7 gopurs_runtime.Value
go__go_1_2_7 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_7:
for {
if false { continue go__go_1_2_7 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_7
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return go__go_1_2_7
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__4097915271
}

var cache_foldableList__331628915 gopurs_runtime.Value
var once_foldableList__331628915 sync.Once
func Get_foldableList__331628915() gopurs_runtime.Value {
	once_foldableList__331628915.Do(func() {
		cache_foldableList__331628915 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_8 gopurs_runtime.Value
go__go_1_2_8 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_8:
for {
if false { continue go__go_1_2_8 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_8
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return go__go_1_2_8
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__331628915
}

var cache_foldableList__3094856796 gopurs_runtime.Value
var once_foldableList__3094856796 sync.Once
func Get_foldableList__3094856796() gopurs_runtime.Value {
	once_foldableList__3094856796.Do(func() {
		cache_foldableList__3094856796 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_9 gopurs_runtime.Value
go__go_1_2_9 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_9:
for {
if false { continue go__go_1_2_9 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_9
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return go__go_1_2_9
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__3094856796
}

var cache_foldableList__1218280485 gopurs_runtime.Value
var once_foldableList__1218280485 sync.Once
func Get_foldableList__1218280485() gopurs_runtime.Value {
	once_foldableList__1218280485.Do(func() {
		cache_foldableList__1218280485 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, b_4, gopurs_runtime.Apply(f_3, a_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_10 gopurs_runtime.Value
go__go_1_2_10 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_10:
for {
if false { continue go__go_1_2_10 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_3)}.UnsafePtr).V1
continue go__go_1_2_10
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return go__go_1_2_10
}), gopurs_runtime.Func(func(op_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
})
})
}))
	})
	return cache_foldableList__1218280485
}

var cache_functorList__699353223 gopurs_runtime.Value
var once_functorList__699353223 sync.Once
func Get_functorList__699353223() gopurs_runtime.Value {
	once_functorList__699353223.Do(func() {
		cache_functorList__699353223 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_functorList__699353223
}

var cache_functorList__3996674161 gopurs_runtime.Value
var once_functorList__3996674161 sync.Once
func Get_functorList__3996674161() gopurs_runtime.Value {
	once_functorList__3996674161.Do(func() {
		cache_functorList__3996674161 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_1)
})
}))
	})
	return cache_functorList__3996674161
}

var cache_lazyList__601034736 gopurs_runtime.Value
var once_lazyList__601034736 sync.Once
func Get_lazyList__601034736() gopurs_runtime.Value {
	once_lazyList__601034736.Do(func() {
		cache_lazyList__601034736 = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(f_0, x_1))
}))
}))
	})
	return cache_lazyList__601034736
}

var cache_monoidList__245587391 gopurs_runtime.Value
var once_monoidList__245587391 sync.Once
func Get_monoidList__245587391() gopurs_runtime.Value {
	once_monoidList__245587391.Do(func() {
		cache_monoidList__245587391 = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_semigroupList()
}), pkg_Data_List_Lazy_Types.Get_nil())
	})
	return cache_monoidList__245587391
}

var cache_nil__1478684294 gopurs_runtime.Value
var once_nil__1478684294 sync.Once
func Get_nil__1478684294() gopurs_runtime.Value {
	once_nil__1478684294.Do(func() {
		cache_nil__1478684294 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__1478684294
}

var cache_nil__3988504114 gopurs_runtime.Value
var once_nil__3988504114 sync.Once
func Get_nil__3988504114() gopurs_runtime.Value {
	once_nil__3988504114.Do(func() {
		cache_nil__3988504114 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__3988504114
}

var cache_nil__2012296605 gopurs_runtime.Value
var once_nil__2012296605 sync.Once
func Get_nil__2012296605() gopurs_runtime.Value {
	once_nil__2012296605.Do(func() {
		cache_nil__2012296605 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__2012296605
}

var cache_nil__4122162182 gopurs_runtime.Value
var once_nil__4122162182 sync.Once
func Get_nil__4122162182() gopurs_runtime.Value {
	once_nil__4122162182.Do(func() {
		cache_nil__4122162182 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_nil__4122162182
}

var cache_semigroupList__1199693447 gopurs_runtime.Value
var once_semigroupList__1199693447 sync.Once
func Get_semigroupList__1199693447() gopurs_runtime.Value {
	once_semigroupList__1199693447.Do(func() {
		cache_semigroupList__1199693447 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__1199693447
}

var cache_semigroupList__3612943602 gopurs_runtime.Value
var once_semigroupList__3612943602 sync.Once
func Get_semigroupList__3612943602() gopurs_runtime.Value {
	once_semigroupList__3612943602.Do(func() {
		cache_semigroupList__3612943602 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__3612943602
}

var cache_semigroupList__4136327256 gopurs_runtime.Value
var once_semigroupList__4136327256 sync.Once
func Get_semigroupList__4136327256() gopurs_runtime.Value {
	once_semigroupList__4136327256.Do(func() {
		cache_semigroupList__4136327256 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
})
}))
	})
	return cache_semigroupList__4136327256
}

var cache_step__3545407802 gopurs_runtime.Value
var once_step__3545407802 sync.Once
func Get_step__3545407802() gopurs_runtime.Value {
	once_step__3545407802.Do(func() {
		cache_step__3545407802 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__3545407802(x_0_box))}
})
	})
	return cache_step__3545407802
}

var cache_step__4184651873 gopurs_runtime.Value
var once_step__4184651873 sync.Once
func Get_step__4184651873() gopurs_runtime.Value {
	once_step__4184651873.Do(func() {
		cache_step__4184651873 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__4184651873(x_0_box))}
})
	})
	return cache_step__4184651873
}

var cache_step__4057057377 gopurs_runtime.Value
var once_step__4057057377 sync.Once
func Get_step__4057057377() gopurs_runtime.Value {
	once_step__4057057377.Do(func() {
		cache_step__4057057377 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__4057057377(x_0_box))}
})
	})
	return cache_step__4057057377
}

var cache_step__2999566881 gopurs_runtime.Value
var once_step__2999566881 sync.Once
func Get_step__2999566881() gopurs_runtime.Value {
	once_step__2999566881.Do(func() {
		cache_step__2999566881 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_step__2999566881(x_0_box))}
})
	})
	return cache_step__2999566881
}

var cache_traversableList__3068288903 gopurs_runtime.Value
var once_traversableList__3068288903 sync.Once
func Get_traversableList__3068288903() gopurs_runtime.Value {
	once_traversableList__3068288903.Do(func() {
		cache_traversableList__3068288903 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_List_Lazy_Types.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}))
	})
	return cache_traversableList__3068288903
}

var cache_traversableList__589375054 gopurs_runtime.Value
var once_traversableList__589375054 sync.Once
func Get_traversableList__589375054() gopurs_runtime.Value {
	once_traversableList__589375054.Do(func() {
		cache_traversableList__589375054 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_List_Lazy_Types.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}))
	})
	return cache_traversableList__589375054
}

var cache_drop__4024047148 gopurs_runtime.Value
var once_drop__4024047148 sync.Once
func Get_drop__4024047148() gopurs_runtime.Value {
	once_drop__4024047148.Do(func() {
		cache_drop__4024047148 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop__4024047148(n_0_box.IntVal)
})
	})
	return cache_drop__4024047148
}

var cache_length__162861552 gopurs_runtime.Value
var once_length__162861552 sync.Once
func Get_length__162861552() gopurs_runtime.Value {
	once_length__162861552.Do(func() {
		cache_length__162861552 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(l_0.IntVal), gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length__162861552
}

var cache_repeat__2462085934 gopurs_runtime.Value
var once_repeat__2462085934 sync.Once
func Get_repeat__2462085934() gopurs_runtime.Value {
	once_repeat__2462085934.Do(func() {
		cache_repeat__2462085934 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat__2462085934(x_0_box)
})
	})
	return cache_repeat__2462085934
}

var cache_zipWith__3539178005 gopurs_runtime.Value
var once_zipWith__3539178005 sync.Once
func Get_zipWith__3539178005() gopurs_runtime.Value {
	once_zipWith__3539178005.Do(func() {
		cache_zipWith__3539178005 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__3539178005(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__3539178005
}

var cache_zipWith__2064722709 gopurs_runtime.Value
var once_zipWith__2064722709 sync.Once
func Get_zipWith__2064722709() gopurs_runtime.Value {
	once_zipWith__2064722709.Do(func() {
		cache_zipWith__2064722709 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith__2064722709(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith__2064722709
}

var cache_altZipList__3296309911 gopurs_runtime.Value
var once_altZipList__3296309911 sync.Once
func Get_altZipList__3296309911() gopurs_runtime.Value {
	once_altZipList__3296309911.Do(func() {
		cache_altZipList__3296309911 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), v_0, gopurs_runtime.Apply2(pkg_Data_List_Lazy.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), v_0).IntVal), v1_1))
})
}))
	})
	return cache_altZipList__3296309911
}

var cache_applicativeZipList__37190504 gopurs_runtime.Value
var once_applicativeZipList__37190504 sync.Once
func Get_applicativeZipList__37190504() gopurs_runtime.Value {
	once_applicativeZipList__37190504.Do(func() {
		cache_applicativeZipList__37190504 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyZipList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_13 gopurs_runtime.Value
_ = go__go_1_0_13
go__go_1_0_13 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_13})}
}))
}))
return go__go_1_0_13
}))
	})
	return cache_applicativeZipList__37190504
}

var cache_applyZipList__1358886895 gopurs_runtime.Value
var once_applyZipList__1358886895 sync.Once
func Get_applyZipList__1358886895() gopurs_runtime.Value {
	once_applyZipList__1358886895.Do(func() {
		cache_applyZipList__1358886895 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_List_Lazy.Get_zipWith(), pkg_Data_Function.Get_apply(), v_0, v1_1)
})
}))
	})
	return cache_applyZipList__1358886895
}

var cache_functorZipList__3996674161 gopurs_runtime.Value
var once_functorZipList__3996674161 sync.Once
func Get_functorZipList__3996674161() gopurs_runtime.Value {
	once_functorZipList__3996674161.Do(func() {
		cache_functorZipList__3996674161 = pkg_Data_List_Lazy_Types.Get_functorList()
	})
	return cache_functorZipList__3996674161
}

var cache_plusZipList__3460472018 gopurs_runtime.Value
var once_plusZipList__3460472018 sync.Once
func Get_plusZipList__3460472018() gopurs_runtime.Value {
	once_plusZipList__3460472018.Do(func() {
		cache_plusZipList__3460472018 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altZipList()
}), gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monoidList(), "mempty"))
	})
	return cache_plusZipList__3460472018
}

var cache_unwrap__1997596318 gopurs_runtime.Value
var once_unwrap__1997596318 sync.Once
func Get_unwrap__1997596318() gopurs_runtime.Value {
	once_unwrap__1997596318.Do(func() {
		cache_unwrap__1997596318 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__1997596318(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__1997596318
}

var cache_unwrap__3159071614 gopurs_runtime.Value
var once_unwrap__3159071614 sync.Once
func Get_unwrap__3159071614() gopurs_runtime.Value {
	once_unwrap__3159071614.Do(func() {
		cache_unwrap__3159071614 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3159071614(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3159071614
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_unwrap__831442259 gopurs_runtime.Value
var once_unwrap__831442259 sync.Once
func Get_unwrap__831442259() gopurs_runtime.Value {
	once_unwrap__831442259.Do(func() {
		cache_unwrap__831442259 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__831442259(__eta0_0_box)
})
	})
	return cache_unwrap__831442259
}

var cache_unwrap__4291124211 gopurs_runtime.Value
var once_unwrap__4291124211 sync.Once
func Get_unwrap__4291124211() gopurs_runtime.Value {
	once_unwrap__4291124211.Do(func() {
		cache_unwrap__4291124211 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__4291124211(__eta0_0_box)
})
	})
	return cache_unwrap__4291124211
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__2462288412 gopurs_runtime.Value
var once_append__2462288412 sync.Once
func Get_append__2462288412() gopurs_runtime.Value {
	once_append__2462288412.Do(func() {
		cache_append__2462288412 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__2462288412(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__2462288412
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

var cache_append__2734706680 gopurs_runtime.Value
var once_append__2734706680 sync.Once
func Get_append__2734706680() gopurs_runtime.Value {
	once_append__2734706680.Do(func() {
		cache_append__2734706680 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__2734706680(xs_0_box, ys_1_box)
})
	})
	return cache_append__2734706680
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
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

var cache_show__3316320786 gopurs_runtime.Value
var once_show__3316320786 sync.Once
func Get_show__3316320786() gopurs_runtime.Value {
	once_show__3316320786.Do(func() {
		cache_show__3316320786 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3316320786(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__3316320786
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__894989549 gopurs_runtime.Value
var once_traverse__894989549 sync.Once
func Get_traverse__894989549() gopurs_runtime.Value {
	once_traverse__894989549.Do(func() {
		cache_traverse__894989549 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__894989549(dictApplicative_0_box)
})
	})
	return cache_traverse__894989549
}

var cache_traverse__1157172365 gopurs_runtime.Value
var once_traverse__1157172365 sync.Once
func Get_traverse__1157172365() gopurs_runtime.Value {
	once_traverse__1157172365.Do(func() {
		cache_traverse__1157172365 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__1157172365(dictApplicative_0_box)
})
	})
	return cache_traverse__1157172365
}

var cache_unsafeCrashWith__69763299 gopurs_runtime.Value
var once_unsafeCrashWith__69763299 sync.Once
func Get_unsafeCrashWith__69763299() gopurs_runtime.Value {
	once_unsafeCrashWith__69763299.Do(func() {
		cache_unsafeCrashWith__69763299 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__69763299(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__69763299
}

var cache_unsafeCrashWith__551270687 gopurs_runtime.Value
var once_unsafeCrashWith__551270687 sync.Once
func Get_unsafeCrashWith__551270687() gopurs_runtime.Value {
	once_unsafeCrashWith__551270687.Do(func() {
		cache_unsafeCrashWith__551270687 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__551270687(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__551270687
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_crashWith__1894115486 gopurs_runtime.Value
var once_crashWith__1894115486 sync.Once
func Get_crashWith__1894115486() gopurs_runtime.Value {
	once_crashWith__1894115486.Do(func() {
		cache_crashWith__1894115486 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith__1894115486(_dollar__unused_0_box)
})
	})
	return cache_crashWith__1894115486
}

func Call_ZipList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showZipList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showList_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_showList(), dictShow_0))
_ = showList_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(ZipList "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showList_1_0.V0, v_2).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
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
}), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("bind: unreachable"))
})))
}

func Call_pure__189931222(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3236307030(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__355615152(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__4203183626(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2962221386(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__1851858028(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__3620326986(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
Bind1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monadList(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_monadList(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Apply2(Bind1_2_0.V1, __eta0_0, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_0.V1, __eta1_1, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_3_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
}

func Call_defer__3258767445(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_defer__3967925939(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fix__1475205859(dictLazy_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_1 gopurs_runtime.Value
_ = go__go_2_0_1
go__go_2_0_1 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__go_2_0_1)
}))
return go__go_2_0_1
}

func Call_fix__3570066147(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_2 gopurs_runtime.Value
_ = go__go_1_0_2
go__go_1_0_2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, go__go_1_0_2)
}))
return go__go_1_0_2
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__1422885860(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__267332164(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__524683195(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_3 gopurs_runtime.Value
go__go_1_0_3 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_3:
for {
if false { continue go__go_1_0_3 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_3
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return go__go_1_0_3
}

func Call_foldl__3379885725(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_4 gopurs_runtime.Value
go__go_1_0_4 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop int64 = b_2_loop_val.IntVal
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_4:
for {
if false { continue go__go_1_0_4 }
var b_2 int64 = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(b_2)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, gopurs_runtime.Int(b_2), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0))}).IntVal
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_4
__t2 = gopurs_runtime.Int(gopurs_runtime.Value{}.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Int(__t2.IntVal)
}
}()
})
})
return go__go_1_0_4
}

func Call_foldl__1985071933(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_5 gopurs_runtime.Value
go__go_1_0_5 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_5:
for {
if false { continue go__go_1_0_5 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0)
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_5
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return go__go_1_0_5
}

func Call_foldl__536153533(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_6 gopurs_runtime.Value
go__go_1_0_6 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_6:
for {
if false { continue go__go_1_0_6 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
v_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0))})
xs_3_loop = (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V1
continue go__go_1_0_6
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return go__go_1_0_6
}

func Call_foldr__2111289130(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__926146538(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__1985071933(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
}

func Call_foldr__2389967549(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_4, b_3})}
}))
})
}), pkg_Data_List_Lazy_Types.Get_nil(), xs_2))
}

func Call_apply__458711162(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3658931456(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3019832928(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2175652032(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__848188896(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__913470112(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2665381605(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1542634789(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2675323109(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3871729957(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__831829748(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1510739772(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))
}))
}

func Call_map__3565923196(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1))))}
}))
}

func Call_map__109003388(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1)))})))}
}))
}

func Call_map__2156385148(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_1)))})
}))
}

func Call_cons__716923058(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_cons__720046150(x_0_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value], xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}, xs_1})}
}))
}

func Call_cons__2305074921(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, xs_1})}
}))
}

func Call_step__3545407802(x_0_loop gopurs_runtime.Value) *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__4184651873(x_0_loop gopurs_runtime.Value) *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__4057057377(x_0_loop gopurs_runtime.Value) *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_step__2999566881(x_0_loop gopurs_runtime.Value) *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_0))
}

func Call_drop__4024047148(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_11 gopurs_runtime.Value
go__go_1_0_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
v_2_loop = Call_sub__1043827704(gopurs_runtime.Int(v_2), gopurs_runtime.Int(1)).IntVal
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1))
continue go__go_1_0_11
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(go__go_1_0_11, gopurs_runtime.Int(n_0)))
_ = __local_var_2_3
__local_var_2_2 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, x_3)
})
}

func Call_repeat__2462085934(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_12 gopurs_runtime.Value
_ = go__go_1_0_12
go__go_1_0_12 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_12})}
}))
}))
return go__go_1_0_12
}

func Call_zipWith__3539178005(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply3(pkg_Data_List_Lazy.Get_zipWith(), f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}), xs_1), ys_2)
}

func Call_zipWith__2064722709(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 218341868 && v_3.UnsafePtr != nil)) && ((v1_4.Type == 9 && v1_4.IntVal == 218341868 && v1_4.UnsafePtr != nil)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply3(pkg_Data_List_Lazy.Get_zipWith(), f_0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}), xs_1), ys_2)
}

func Call_unwrap__1997596318(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3159071614(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__831442259(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_unwrap__4291124211(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__2462288412(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_append__2734706680(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1, ys_1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), xs_0)
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3316320786(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__894989549(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}

func Call_traverse__1157172365(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_List_Lazy_Types.Get_cons(), gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](a_4))})), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_List_Lazy_Types.Get_nil()))
})
}

func Call_unsafeCrashWith__69763299(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}))
}

func Call_unsafeCrashWith__551270687(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}))
}

func Call_crashWith__1894115486(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Partial.Get__crashWith()
}


