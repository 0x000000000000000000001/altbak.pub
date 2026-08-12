package Data_List_Lazy_NonEmpty

import (
	pkg_Control_Lazy "gopurs/output/Control.Lazy"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(v_0_box)
})
	})
	return cache_uncons
}

var cache_toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		cache_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList(v_0_box)
})
	})
	return cache_toList
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(v_0_box)
})
	})
	return cache_tail
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_applicativeNonEmptyList(), "pure")
	})
	return cache_singleton
}

var cache_repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		cache_repeat = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat(x_0_box)
})
	})
	return cache_repeat
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_length(v_0_box))
})
	})
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last(v_0_box)
})
	})
	return cache_last
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(f_0_box, x_1_box)
})
	})
	return cache_iterate
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_init(v_0_box)
})
	})
	return cache_init
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(v_0_box)
})
	})
	return cache_head
}

var cache_fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		cache_fromList = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList(l_0_box))}
})
	})
	return cache_fromList
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_fromFoldable
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(y_0_box, v_1_box)
})
	})
	return cache_cons
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concatMap(b_0_box, a_1_box)
})
	})
	return cache_concatMap
}

var cache_appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		cache_appendFoldable = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, nel_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_appendFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), nel_1_box, ys_2_box)
})
	})
	return cache_appendFoldable
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
		cache_foldl__524683195 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__524683195
}

var cache_foldl__3379885725 gopurs_runtime.Value
var once_foldl__3379885725 sync.Once
func Get_foldl__3379885725() gopurs_runtime.Value {
	once_foldl__3379885725.Do(func() {
		cache_foldl__3379885725 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__3379885725
}

var cache_foldl__536153533 gopurs_runtime.Value
var once_foldl__536153533 sync.Once
func Get_foldl__536153533() gopurs_runtime.Value {
	once_foldl__536153533.Do(func() {
		cache_foldl__536153533 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__536153533
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

var cache_foldr__2403185435 gopurs_runtime.Value
var once_foldr__2403185435 sync.Once
func Get_foldr__2403185435() gopurs_runtime.Value {
	once_foldr__2403185435.Do(func() {
		cache_foldr__2403185435 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2403185435(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2403185435
}

var cache_const__754826900 gopurs_runtime.Value
var once_const__754826900 sync.Once
func Get_const__754826900() gopurs_runtime.Value {
	once_const__754826900.Do(func() {
		cache_const__754826900 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__754826900(a_0_box, v_1_box)
})
	})
	return cache_const__754826900
}

var cache_const__4115900574 gopurs_runtime.Value
var once_const__4115900574 sync.Once
func Get_const__4115900574() gopurs_runtime.Value {
	once_const__4115900574.Do(func() {
		cache_const__4115900574 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4115900574(a_0_box, v_1_box)
})
	})
	return cache_const__4115900574
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

var cache_flip__1818565536 gopurs_runtime.Value
var once_flip__1818565536 sync.Once
func Get_flip__1818565536() gopurs_runtime.Value {
	once_flip__1818565536.Do(func() {
		cache_flip__1818565536 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1818565536(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1818565536
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

var cache_map__109003388 gopurs_runtime.Value
var once_map__109003388 sync.Once
func Get_map__109003388() gopurs_runtime.Value {
	once_map__109003388.Do(func() {
		cache_map__109003388 = gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map")
	})
	return cache_map__109003388
}

var cache_map__1808515292 gopurs_runtime.Value
var once_map__1808515292 sync.Once
func Get_map__1808515292() gopurs_runtime.Value {
	once_map__1808515292.Do(func() {
		cache_map__1808515292 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1808515292
}

var cache_map__1980149980 gopurs_runtime.Value
var once_map__1980149980 sync.Once
func Get_map__1980149980() gopurs_runtime.Value {
	once_map__1980149980.Do(func() {
		cache_map__1980149980 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1980149980
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

var cache_fromList__1361791219 gopurs_runtime.Value
var once_fromList__1361791219 sync.Once
func Get_fromList__1361791219() gopurs_runtime.Value {
	once_fromList__1361791219.Do(func() {
		cache_fromList__1361791219 = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList__1361791219(l_0_box))}
})
	})
	return cache_fromList__1361791219
}

var cache_head__3204870332 gopurs_runtime.Value
var once_head__3204870332 sync.Once
func Get_head__3204870332() gopurs_runtime.Value {
	once_head__3204870332.Do(func() {
		cache_head__3204870332 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head__3204870332(v_0_box)
})
	})
	return cache_head__3204870332
}

var cache_tail__4101396777 gopurs_runtime.Value
var once_tail__4101396777 sync.Once
func Get_tail__4101396777() gopurs_runtime.Value {
	once_tail__4101396777.Do(func() {
		cache_tail__4101396777 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail__4101396777(v_0_box)
})
	})
	return cache_tail__4101396777
}

var cache_toList__1017592434 gopurs_runtime.Value
var once_toList__1017592434 sync.Once
func Get_toList__1017592434() gopurs_runtime.Value {
	once_toList__1017592434.Do(func() {
		cache_toList__1017592434 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList__1017592434(v_0_box)
})
	})
	return cache_toList__1017592434
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
var go__go_1_2_3 gopurs_runtime.Value
go__go_1_2_3 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_3:
for {
if false { continue go__go_1_2_3 }
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
continue go__go_1_2_3
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
return go__go_1_2_3
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
var go__go_1_2_4 gopurs_runtime.Value
go__go_1_2_4 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_2_4:
for {
if false { continue go__go_1_2_4 }
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
continue go__go_1_2_4
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
return go__go_1_2_4
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

var cache_fromFoldable__4212258679 gopurs_runtime.Value
var once_fromFoldable__4212258679 sync.Once
func Get_fromFoldable__4212258679() gopurs_runtime.Value {
	once_fromFoldable__4212258679.Do(func() {
		cache_fromFoldable__4212258679 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable__4212258679(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_fromFoldable__4212258679
}

var cache_init__109194273 gopurs_runtime.Value
var once_init__109194273 sync.Once
func Get_init__109194273() gopurs_runtime.Value {
	once_init__109194273.Do(func() {
		cache_init__109194273 = func() gopurs_runtime.Value {
var go__go_0_0_5 gopurs_runtime.Value
_ = go__go_0_0_5
go__go_0_0_5 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, pkg_Data_List_Lazy_Types.Get_nil()})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_cons(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0), gopurs_runtime.Apply(go__go_0_0_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)))))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_init__109194273
}

var cache_iterate__455058292 gopurs_runtime.Value
var once_iterate__455058292 sync.Once
func Get_iterate__455058292() gopurs_runtime.Value {
	once_iterate__455058292.Do(func() {
		cache_iterate__455058292 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__455058292(f_0_box, x_1_box)
})
	})
	return cache_iterate__455058292
}

var cache_last__1102843348 gopurs_runtime.Value
var once_last__1102843348 sync.Once
func Get_last__1102843348() gopurs_runtime.Value {
	once_last__1102843348.Do(func() {
		cache_last__1102843348 = func() gopurs_runtime.Value {
var go__go_0_0_7 gopurs_runtime.Value
go__go_0_0_7 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
go__go_0_0_7:
for {
if false { continue go__go_0_0_7 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t4 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
v_1_loop = gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
continue go__go_0_0_7
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_7, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), x_1))
})
}()
	})
	return cache_last__1102843348
}

var cache_length__162861552 gopurs_runtime.Value
var once_length__162861552 sync.Once
func Get_length__162861552() gopurs_runtime.Value {
	once_length__162861552.Do(func() {
		cache_length__162861552 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), l_0, gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length__162861552
}

var cache_null__1674339719 gopurs_runtime.Value
var once_null__1674339719 sync.Once
func Get_null__1674339719() gopurs_runtime.Value {
	once_null__1674339719.Do(func() {
		cache_null__1674339719 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null__1674339719(x_0_box))
})
	})
	return cache_null__1674339719
}

var cache_repeat__2149902581 gopurs_runtime.Value
var once_repeat__2149902581 sync.Once
func Get_repeat__2149902581() gopurs_runtime.Value {
	once_repeat__2149902581.Do(func() {
		cache_repeat__2149902581 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat__2149902581(x_0_box)
})
	})
	return cache_repeat__2149902581
}

var cache_uncons__3647012005 gopurs_runtime.Value
var once_uncons__3647012005 sync.Once
func Get_uncons__3647012005() gopurs_runtime.Value {
	once_uncons__3647012005.Do(func() {
		cache_uncons__3647012005 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__3647012005(xs_0_box))}
})
	})
	return cache_uncons__3647012005
}

var cache_uncons__1321764894 gopurs_runtime.Value
var once_uncons__1321764894 sync.Once
func Get_uncons__1321764894() gopurs_runtime.Value {
	once_uncons__1321764894.Do(func() {
		cache_uncons__1321764894 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__1321764894(xs_0_box))}
})
	})
	return cache_uncons__1321764894
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_isNothing__1401305026 gopurs_runtime.Value
var once_isNothing__1401305026 sync.Once
func Get_isNothing__1401305026() gopurs_runtime.Value {
	once_isNothing__1401305026.Do(func() {
		cache_isNothing__1401305026 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__1401305026(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__1401305026
}

var cache_maybe__2641488518 gopurs_runtime.Value
var once_maybe__2641488518 sync.Once
func Get_maybe__2641488518() gopurs_runtime.Value {
	once_maybe__2641488518.Do(func() {
		cache_maybe__2641488518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__2641488518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__2641488518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_maybe__3931678292 gopurs_runtime.Value
var once_maybe__3931678292 sync.Once
func Get_maybe__3931678292() gopurs_runtime.Value {
	once_maybe__3931678292.Do(func() {
		cache_maybe__3931678292 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3931678292(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3931678292
}

var cache_maybe__1845466785 gopurs_runtime.Value
var once_maybe__1845466785 sync.Once
func Get_maybe__1845466785() gopurs_runtime.Value {
	once_maybe__1845466785.Do(func() {
		cache_maybe__1845466785 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1845466785(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1845466785
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

var cache_unwrap__4291124211 gopurs_runtime.Value
var once_unwrap__4291124211 sync.Once
func Get_unwrap__4291124211() gopurs_runtime.Value {
	once_unwrap__4291124211.Do(func() {
		cache_unwrap__4291124211 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_unwrap__4291124211
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
		cache_append__2734706680 = gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append")
	})
	return cache_append__2734706680
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
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

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldr__1779806517 gopurs_runtime.Value
var once_unfoldr__1779806517 sync.Once
func Get_unfoldr__1779806517() gopurs_runtime.Value {
	once_unfoldr__1779806517.Do(func() {
		cache_unfoldr__1779806517 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1779806517(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1779806517
}

func Call_uncons(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
return gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V1)
}

func Call_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V0
_ = __local_var_2_1
__local_var_3_2 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V1
_ = __local_var_3_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, __local_var_3_2})}
}))
}

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.RecordGet(rec_2, "tail")})}
}), gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), xs_1))))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, Call_toList(x_2))
})
}

func Call_tail(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V1
}

func Call_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_0 gopurs_runtime.Value
_ = go__go_2_0_0
go__go_2_0_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_2_0_0})}
}))
}))
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, go__go_2_0_0})}
}))
}

func Call_length(v_0_loop gopurs_runtime.Value) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V1)).IntVal
}

func Call_last(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_last(), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0
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

func Call_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply2(pkg_Data_List_Lazy.Get_iterate(), f_0, gopurs_runtime.Apply(f_0, x_1))})}
}))
}

func Call_init(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_init(), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V1)
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t3 = pkg_Data_List_Lazy_Types.Get_nil()
goto end_branch_3
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__local_var_3_2 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0
_ = __local_var_3_2
__t3 = gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V0, __local_var_3_2})}
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}

func Call_head(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V0
}

func Call_fromList(l_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var l_0 gopurs_runtime.Value = l_0_loop
_ = l_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_0))
_ = v_1_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__local_var_2_1 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0
_ = __local_var_2_1
__local_var_3_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1
_ = __local_var_3_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_1, __local_var_3_2})}
}))})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3)
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.Apply2(dictFoldable_0.V2, pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList(gopurs_runtime.Apply(__local_var_1_0, x_2)))}
})
}

func Call_cons(y_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v2_3_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_1)
_ = v2_3_0
__local_var_4_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3_0.UnsafePtr).V0
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3_0.UnsafePtr).V1
_ = __local_var_5_2
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, y_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_4_1, __local_var_5_2})}
}))})}
}))
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_bindNonEmptyList(), "bind"), a_1, b_0)
}

func Call_appendFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], nel_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var nel_1 gopurs_runtime.Value = nel_1_loop
_ = nel_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_1).UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_semigroupList(), "append"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), nel_1).UnsafePtr).V1, gopurs_runtime.Apply3(dictFoldable_0.V2, pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil(), ys_2))})}
}))
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

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2403185435(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_const__754826900(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4115900574(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
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

func Call_flip__3019832928(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1818565536(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromList__1361791219(l_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var l_0 gopurs_runtime.Value = l_0_loop
_ = l_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), l_0))
_ = v_1_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__local_var_2_1 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0
_ = __local_var_2_1
__local_var_3_2 := (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1
_ = __local_var_3_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_1, __local_var_3_2})}
}))})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3)
}

func Call_head__3204870332(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V0
}

func Call_tail__4101396777(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0).UnsafePtr).V1
}

func Call_toList__1017592434(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), v_0)
_ = v1_1_0
__local_var_2_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V0
_ = __local_var_2_1
__local_var_3_2 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1_0.UnsafePtr).V1
_ = __local_var_3_2
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_1, __local_var_3_2})}
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

func Call_fromFoldable__4212258679(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, pkg_Data_List_Lazy_Types.Get_cons(), pkg_Data_List_Lazy_Types.Get_nil())
}

func Call_iterate__455058292(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
go__go_2_0_6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_functorList(), "map"), f_0, go__go_2_0_6)
_ = __local_var_4_1
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1})}
}))
}))
return go__go_2_0_6
}

func Call_null__1674339719(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_uncons(), x_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_repeat__2149902581(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_8 gopurs_runtime.Value
_ = go__go_1_0_8
go__go_1_0_8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Lazy_Types.Get_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_8})}
}))
}))
return go__go_1_0_8
}

func Call_uncons__3647012005(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1)
}

func Call_uncons__1321764894(xs_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), xs_0))
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 218341868 && gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0, (*pkg_Data_List_Lazy_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1)
}

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_isNothing__1401305026(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__2641488518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3931678292(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__1845466785(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__1779806517(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


