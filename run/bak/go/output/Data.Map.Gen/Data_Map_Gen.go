package Data_Map_Gen

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genMap gopurs_runtime.Value
var once_genMap sync.Once
func Get_genMap() gopurs_runtime.Value {
	once_genMap.Do(func() {
		cache_genMap = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMap(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_genMap
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

var cache_apply__2623550860 gopurs_runtime.Value
var once_apply__2623550860 sync.Once
func Get_apply__2623550860() gopurs_runtime.Value {
	once_apply__2623550860.Do(func() {
		cache_apply__2623550860 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2623550860(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2623550860
}

var cache_bind__2517150375 gopurs_runtime.Value
var once_bind__2517150375 sync.Once
func Get_bind__2517150375() gopurs_runtime.Value {
	once_bind__2517150375.Do(func() {
		cache_bind__2517150375 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2517150375(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2517150375
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__1801470023 gopurs_runtime.Value
var once_bind__1801470023 sync.Once
func Get_bind__1801470023() gopurs_runtime.Value {
	once_bind__1801470023.Do(func() {
		cache_bind__1801470023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1801470023(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1801470023
}

var cache_chooseInt__1063828903 gopurs_runtime.Value
var once_chooseInt__1063828903 sync.Once
func Get_chooseInt__1063828903() gopurs_runtime.Value {
	once_chooseInt__1063828903.Do(func() {
		cache_chooseInt__1063828903 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseInt__1063828903(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseInt__1063828903
}

var cache_resize__1050945947 gopurs_runtime.Value
var once_resize__1050945947 sync.Once
func Get_resize__1050945947() gopurs_runtime.Value {
	once_resize__1050945947.Do(func() {
		cache_resize__1050945947 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize__1050945947(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_resize__1050945947
}

var cache_resize__1313223195 gopurs_runtime.Value
var once_resize__1313223195 sync.Once
func Get_resize__1313223195() gopurs_runtime.Value {
	once_resize__1313223195.Do(func() {
		cache_resize__1313223195 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize__1313223195(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_resize__1313223195
}

var cache_sized__3147117623 gopurs_runtime.Value
var once_sized__3147117623 sync.Once
func Get_sized__3147117623() gopurs_runtime.Value {
	once_sized__3147117623.Do(func() {
		cache_sized__3147117623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__3147117623(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__3147117623
}

var cache_sized__4206899191 gopurs_runtime.Value
var once_sized__4206899191 sync.Once
func Get_sized__4206899191() gopurs_runtime.Value {
	once_sized__4206899191.Do(func() {
		cache_sized__4206899191 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__4206899191(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__4206899191
}

var cache_sized__2391211191 gopurs_runtime.Value
var once_sized__2391211191 sync.Once
func Get_sized__2391211191() gopurs_runtime.Value {
	once_sized__2391211191.Do(func() {
		cache_sized__2391211191 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__2391211191(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__2391211191
}

var cache_unfoldable__542018773 gopurs_runtime.Value
var once_unfoldable__542018773 sync.Once
func Get_unfoldable__542018773() gopurs_runtime.Value {
	once_unfoldable__542018773.Do(func() {
		cache_unfoldable__542018773 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable__542018773(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_unfoldable__542018773
}

var cache_unfoldable__4218841939 gopurs_runtime.Value
var once_unfoldable__4218841939 sync.Once
func Get_unfoldable__4218841939() gopurs_runtime.Value {
	once_unfoldable__4218841939.Do(func() {
		cache_unfoldable__4218841939 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable__4218841939(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_unfoldable__4218841939
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__2063592441 gopurs_runtime.Value
var once_tailRecM__2063592441 sync.Once
func Get_tailRecM__2063592441() gopurs_runtime.Value {
	once_tailRecM__2063592441.Do(func() {
		cache_tailRecM__2063592441 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__2063592441(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__2063592441
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

var cache_foldl__3943124669 gopurs_runtime.Value
var once_foldl__3943124669 sync.Once
func Get_foldl__3943124669() gopurs_runtime.Value {
	once_foldl__3943124669.Do(func() {
		cache_foldl__3943124669 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3943124669(f_0_box)
})
	})
	return cache_foldl__3943124669
}

var cache_foldl__396932925 gopurs_runtime.Value
var once_foldl__396932925 sync.Once
func Get_foldl__396932925() gopurs_runtime.Value {
	once_foldl__396932925.Do(func() {
		cache_foldl__396932925 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__396932925(f_0_box)
})
	})
	return cache_foldl__396932925
}

var cache_const__1243414737 gopurs_runtime.Value
var once_const__1243414737 sync.Once
func Get_const__1243414737() gopurs_runtime.Value {
	once_const__1243414737.Do(func() {
		cache_const__1243414737 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1243414737(a_0_box, v_1_box)
})
	})
	return cache_const__1243414737
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
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

var cache_flip__2974723072 gopurs_runtime.Value
var once_flip__2974723072 sync.Once
func Get_flip__2974723072() gopurs_runtime.Value {
	once_flip__2974723072.Do(func() {
		cache_flip__2974723072 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2974723072(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2974723072
}

var cache_flip__3563101792 gopurs_runtime.Value
var once_flip__3563101792 sync.Once
func Get_flip__3563101792() gopurs_runtime.Value {
	once_flip__3563101792.Do(func() {
		cache_flip__3563101792 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3563101792(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3563101792
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

var cache_map__3061653364 gopurs_runtime.Value
var once_map__3061653364 sync.Once
func Get_map__3061653364() gopurs_runtime.Value {
	once_map__3061653364.Do(func() {
		cache_map__3061653364 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3061653364(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3061653364
}

var cache_map__3061937460 gopurs_runtime.Value
var once_map__3061937460 sync.Once
func Get_map__3061937460() gopurs_runtime.Value {
	once_map__3061937460.Do(func() {
		cache_map__3061937460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3061937460(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3061937460
}

var cache_map__1504457012 gopurs_runtime.Value
var once_map__1504457012 sync.Once
func Get_map__1504457012() gopurs_runtime.Value {
	once_map__1504457012.Do(func() {
		cache_map__1504457012 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1504457012(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1504457012
}

var cache_foldableList__1753400174 gopurs_runtime.Value
var once_foldableList__1753400174 sync.Once
func Get_foldableList__1753400174() gopurs_runtime.Value {
	once_foldableList__1753400174.Do(func() {
		cache_foldableList__1753400174 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_2 gopurs_runtime.Value
go__go_1_3_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_2:
for {
if false { continue go__go_1_3_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_2
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
return go__go_1_3_2
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_3 gopurs_runtime.Value
go__go_3_7_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_3:
for {
if false { continue go__go_3_7_3 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_3
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__1753400174
}

var cache_unfoldable1List__3672302568 gopurs_runtime.Value
var once_unfoldable1List__3672302568 sync.Once
func Get_unfoldable1List__3672302568() gopurs_runtime.Value {
	once_unfoldable1List__3672302568.Do(func() {
		cache_unfoldable1List__3672302568 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_4 gopurs_runtime.Value
go__go_2_0_4 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})
continue go__go_2_0_4
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_4, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldable1List__3672302568
}

var cache_unfoldableList__2633941518 gopurs_runtime.Value
var once_unfoldableList__2633941518 sync.Once
func Get_unfoldableList__2633941518() gopurs_runtime.Value {
	once_unfoldableList__2633941518.Do(func() {
		cache_unfoldableList__2633941518 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_unfoldable1List()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_5 gopurs_runtime.Value
go__go_2_0_5 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_5:
for {
if false { continue go__go_2_0_5 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_4)})))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V0, memo_4})})
continue go__go_2_0_5
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_5, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldableList__2633941518
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
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

var cache_unfoldr__457386988 gopurs_runtime.Value
var once_unfoldr__457386988 sync.Once
func Get_unfoldr__457386988() gopurs_runtime.Value {
	once_unfoldr__457386988.Do(func() {
		cache_unfoldr__457386988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__457386988(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__457386988
}

func Call_genMap(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
Bind1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
Bind11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind11_4_2
Functor0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind11_4_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
Apply0_6_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind11_4_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_4
return gopurs_runtime.Func(func(dictOrd_7 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable_8_5 := gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_fromFoldable(), dictOrd_7, pkg_Data_List_Types.Get_foldableList())
_ = fromFoldable_8_5
return gopurs_runtime.Func(func(genKey_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genValue_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_1.V1, gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(0), gopurs_runtime.Int(size_11.IntVal)), gopurs_runtime.Func(func(newSize_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Apply(Get_const__1243414737(), gopurs_runtime.Int(newSize_12.IntVal)), gopurs_runtime.Apply2(Functor0_5_3.V0, fromFoldable_8_5, gopurs_runtime.Apply4(pkg_Control_Monad_Gen.Get_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](pkg_Data_List_Types.Get_unfoldableList()))}, gopurs_runtime.Apply2(Apply0_6_4.V1, gopurs_runtime.Apply2(Functor0_5_3.V0, pkg_Data_Tuple.Get_Tuple(), genKey_9), genValue_10))))
}))
}))
})
})
})
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2623550860(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2517150375(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1801470023(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_chooseInt__1063828903(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_resize__1050945947(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_resize__1313223195(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_sized__3147117623(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_sized__4206899191(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_sized__2391211191(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_unfoldable__542018773(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
Bind1_4_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
Functor0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(dictUnfoldable_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1.IntVal), gopurs_runtime.Int(0))).IntVal) != (0) {
__t8 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))}})})
goto end_branch_8
} else {

}
}
{
__local_var_9_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0
_ = __local_var_9_6
__local_var_10_7 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_10_7
__t8 = gopurs_runtime.Apply2(Bind1_4_2.V1, gen_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]{1, x_11, gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]](__local_var_9_6)})}, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_10_7.IntVal), gopurs_runtime.Int(1)).IntVal)})}})})
}))
}
end_branch_8:
return __t8
}))
_ = __local_var_8_5
__local_var_9_9 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_9_9
return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_6, "unfoldr"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 759514854 && v_8.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V1)}})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]]]](__t4))}
})), gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Apply(__local_var_9_9, x_10))
})))
})
})
}

func Call_unfoldable__4218841939(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
Bind1_4_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_2
Functor0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_5 := gopurs_runtime.Apply(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1.IntVal), gopurs_runtime.Int(0))).IntVal) != (0) {
__t8 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0))}})})
goto end_branch_8
} else {

}
}
{
__local_var_8_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
_ = __local_var_8_6
__local_var_9_7 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1
_ = __local_var_9_7
__t8 = gopurs_runtime.Apply2(Bind1_4_2.V1, gen_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]{1, x_10, gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]](__local_var_8_6)})}, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_9_7.IntVal), gopurs_runtime.Int(1)).IntVal)})}})})
}))
}
end_branch_8:
return __t8
}))
_ = __local_var_7_5
__local_var_8_9 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_8_9
return gopurs_runtime.Apply2(Functor0_5_3.V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_unfoldableList(), "unfoldr"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 759514854 && v_7.UnsafePtr != nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V1)}})}})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value]]]](__t4))}
})), gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(__local_var_8_9, x_9))
})))
})
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__2063592441(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_foldl__3943124669(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
return go__go_1_0_0
}

func Call_foldl__396932925(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var b_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_1
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_1
}

func Call_const__1243414737(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_flip__2974723072(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3563101792(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_map__3061653364(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3061937460(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1504457012(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
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

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__457386988(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


