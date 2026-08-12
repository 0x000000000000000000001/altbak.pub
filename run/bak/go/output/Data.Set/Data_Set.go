package Data_Set

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_identity(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_identity
}

var cache_Set gopurs_runtime.Value
var once_Set sync.Once
func Get_Set() gopurs_runtime.Value {
	once_Set.Do(func() {
		cache_Set = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Set(x_0_box)
})
	})
	return cache_Set
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_union
}

var cache_toggle gopurs_runtime.Value
var once_toggle sync.Once
func Get_toggle() gopurs_runtime.Value {
	once_toggle.Do(func() {
		cache_toggle = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_toggle(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_toggle
}

var cache_toMap gopurs_runtime.Value
var once_toMap sync.Once
func Get_toMap() gopurs_runtime.Value {
	once_toMap.Do(func() {
		cache_toMap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_toMap(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toMap
}

var cache_toMap__gopurs_runtime_Value_556171355 gopurs_runtime.Value
var once_toMap__gopurs_runtime_Value_556171355 sync.Once
func Get_toMap__gopurs_runtime_Value_556171355() gopurs_runtime.Value {
	once_toMap__gopurs_runtime_Value_556171355.Do(func() {
		cache_toMap__gopurs_runtime_Value_556171355 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_toMap__gopurs_runtime_Value_556171355(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toMap__gopurs_runtime_Value_556171355
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

var cache_toUnfoldable__gopurs_runtime_Value_871784487 gopurs_runtime.Value
var once_toUnfoldable__gopurs_runtime_Value_871784487 sync.Once
func Get_toUnfoldable__gopurs_runtime_Value_871784487() gopurs_runtime.Value {
	once_toUnfoldable__gopurs_runtime_Value_871784487.Do(func() {
		cache_toUnfoldable__gopurs_runtime_Value_871784487 = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable__gopurs_runtime_Value_871784487(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldable__gopurs_runtime_Value_871784487
}

var cache_size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		cache_size = pkg_Data_Map_Internal.Get_size()
	})
	return cache_size
}

var cache_size__gopurs_runtime_Value_909390430 gopurs_runtime.Value
var once_size__gopurs_runtime_Value_909390430 sync.Once
func Get_size__gopurs_runtime_Value_909390430() gopurs_runtime.Value {
	once_size__gopurs_runtime_Value_909390430.Do(func() {
		cache_size__gopurs_runtime_Value_909390430 = pkg_Data_Map_Internal.Get_size()
	})
	return cache_size__gopurs_runtime_Value_909390430
}

var cache_size__gopurs_runtime_Value_667039429 gopurs_runtime.Value
var once_size__gopurs_runtime_Value_667039429 sync.Once
func Get_size__gopurs_runtime_Value_667039429() gopurs_runtime.Value {
	once_size__gopurs_runtime_Value_667039429.Do(func() {
		cache_size__gopurs_runtime_Value_667039429 = pkg_Data_Map_Internal.Get_size()
	})
	return cache_size__gopurs_runtime_Value_667039429
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton(a_0_box))}
})
	})
	return cache_singleton
}

var cache_singleton__gopurs_runtime_Value_3724491835 gopurs_runtime.Value
var once_singleton__gopurs_runtime_Value_3724491835 sync.Once
func Get_singleton__gopurs_runtime_Value_3724491835() gopurs_runtime.Value {
	once_singleton__gopurs_runtime_Value_3724491835.Do(func() {
		cache_singleton__gopurs_runtime_Value_3724491835 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__gopurs_runtime_Value_3724491835(a_0_box))}
})
	})
	return cache_singleton__gopurs_runtime_Value_3724491835
}

var cache_showSet gopurs_runtime.Value
var once_showSet sync.Once
func Get_showSet() gopurs_runtime.Value {
	once_showSet.Do(func() {
		cache_showSet = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showSet(dictShow_0_box)
})
	})
	return cache_showSet
}

var cache_semigroupSet gopurs_runtime.Value
var once_semigroupSet sync.Once
func Get_semigroupSet() gopurs_runtime.Value {
	once_semigroupSet.Do(func() {
		cache_semigroupSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupSet(dictOrd_0_box)
})
	})
	return cache_semigroupSet
}

var cache_member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		cache_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_member(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_member
}

var cache_isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		cache_isEmpty = pkg_Data_Map_Internal.Get_isEmpty()
	})
	return cache_isEmpty
}

var cache_isEmpty__gopurs_runtime_Value_1620059593 gopurs_runtime.Value
var once_isEmpty__gopurs_runtime_Value_1620059593 sync.Once
func Get_isEmpty__gopurs_runtime_Value_1620059593() gopurs_runtime.Value {
	once_isEmpty__gopurs_runtime_Value_1620059593.Do(func() {
		cache_isEmpty__gopurs_runtime_Value_1620059593 = pkg_Data_Map_Internal.Get_isEmpty()
	})
	return cache_isEmpty__gopurs_runtime_Value_1620059593
}

var cache_isEmpty__gopurs_runtime_Value_1390394866 gopurs_runtime.Value
var once_isEmpty__gopurs_runtime_Value_1390394866 sync.Once
func Get_isEmpty__gopurs_runtime_Value_1390394866() gopurs_runtime.Value {
	once_isEmpty__gopurs_runtime_Value_1390394866.Do(func() {
		cache_isEmpty__gopurs_runtime_Value_1390394866 = pkg_Data_Map_Internal.Get_isEmpty()
	})
	return cache_isEmpty__gopurs_runtime_Value_1390394866
}

var cache_intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		cache_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersection(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_intersection
}

var cache_intersection__gopurs_runtime_Value_833514008 gopurs_runtime.Value
var once_intersection__gopurs_runtime_Value_833514008 sync.Once
func Get_intersection__gopurs_runtime_Value_833514008() gopurs_runtime.Value {
	once_intersection__gopurs_runtime_Value_833514008.Do(func() {
		cache_intersection__gopurs_runtime_Value_833514008 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersection__gopurs_runtime_Value_833514008(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_intersection__gopurs_runtime_Value_833514008
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_insert(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_insert
}

var cache_insert__gopurs_runtime_Value_4217907800 gopurs_runtime.Value
var once_insert__gopurs_runtime_Value_4217907800 sync.Once
func Get_insert__gopurs_runtime_Value_4217907800() gopurs_runtime.Value {
	once_insert__gopurs_runtime_Value_4217907800.Do(func() {
		cache_insert__gopurs_runtime_Value_4217907800 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_insert__gopurs_runtime_Value_4217907800(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_insert__gopurs_runtime_Value_4217907800
}

var cache_fromMap gopurs_runtime.Value
var once_fromMap sync.Once
func Get_fromMap() gopurs_runtime.Value {
	once_fromMap.Do(func() {
		cache_fromMap = Get_Set()
	})
	return cache_fromMap
}

var cache_fromMap__gopurs_runtime_Value_556171355 gopurs_runtime.Value
var once_fromMap__gopurs_runtime_Value_556171355 sync.Once
func Get_fromMap__gopurs_runtime_Value_556171355() gopurs_runtime.Value {
	once_fromMap__gopurs_runtime_Value_556171355.Do(func() {
		cache_fromMap__gopurs_runtime_Value_556171355 = Get_Set()
	})
	return cache_fromMap__gopurs_runtime_Value_556171355
}

var cache_foldableSet gopurs_runtime.Value
var once_foldableSet sync.Once
func Get_foldableSet() gopurs_runtime.Value {
	once_foldableSet.Do(func() {
		cache_foldableSet = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), dictMonoid_0, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}))
	})
	return cache_foldableSet
}

var cache_foldableSet__ptrData_Foldable_Constructor_Foldable_ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value___2661034898 gopurs_runtime.Value
var once_foldableSet__ptrData_Foldable_Constructor_Foldable_ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value___2661034898 sync.Once
func Get_foldableSet__ptrData_Foldable_Constructor_Foldable_ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value___2661034898() gopurs_runtime.Value {
	once_foldableSet__ptrData_Foldable_Constructor_Foldable_ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value___2661034898.Do(func() {
		cache_foldableSet__ptrData_Foldable_Constructor_Foldable_ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value___2661034898 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), dictMonoid_0, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
})})}
	})
	return cache_foldableSet__ptrData_Foldable_Constructor_Foldable_ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value___2661034898
}

var cache_foldableSet__gopurs_runtime_Value_1081688304 gopurs_runtime.Value
var once_foldableSet__gopurs_runtime_Value_1081688304 sync.Once
func Get_foldableSet__gopurs_runtime_Value_1081688304() gopurs_runtime.Value {
	once_foldableSet__gopurs_runtime_Value_1081688304.Do(func() {
		cache_foldableSet__gopurs_runtime_Value_1081688304 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), dictMonoid_0, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}))
	})
	return cache_foldableSet__gopurs_runtime_Value_1081688304
}

var cache_findMin gopurs_runtime.Value
var once_findMin sync.Once
func Get_findMin() gopurs_runtime.Value {
	once_findMin.Do(func() {
		cache_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin
}

var cache_findMin__gopurs_runtime_Value_611847841 gopurs_runtime.Value
var once_findMin__gopurs_runtime_Value_611847841 sync.Once
func Get_findMin__gopurs_runtime_Value_611847841() gopurs_runtime.Value {
	once_findMin__gopurs_runtime_Value_611847841.Do(func() {
		cache_findMin__gopurs_runtime_Value_611847841 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin__gopurs_runtime_Value_611847841(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin__gopurs_runtime_Value_611847841
}

var cache_findMax gopurs_runtime.Value
var once_findMax sync.Once
func Get_findMax() gopurs_runtime.Value {
	once_findMax.Do(func() {
		cache_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax
}

var cache_findMax__gopurs_runtime_Value_611847841 gopurs_runtime.Value
var once_findMax__gopurs_runtime_Value_611847841 sync.Once
func Get_findMax__gopurs_runtime_Value_611847841() gopurs_runtime.Value {
	once_findMax__gopurs_runtime_Value_611847841.Do(func() {
		cache_findMax__gopurs_runtime_Value_611847841 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax__gopurs_runtime_Value_611847841(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax__gopurs_runtime_Value_611847841
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_filter
}

var cache_eqSet gopurs_runtime.Value
var once_eqSet sync.Once
func Get_eqSet() gopurs_runtime.Value {
	once_eqSet.Do(func() {
		cache_eqSet = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqSet(dictEq_0_box)
})
	})
	return cache_eqSet
}

var cache_ordSet gopurs_runtime.Value
var once_ordSet sync.Once
func Get_ordSet() gopurs_runtime.Value {
	once_ordSet.Do(func() {
		cache_ordSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordSet(dictOrd_0_box)
})
	})
	return cache_ordSet
}

var cache_eq1Set gopurs_runtime.Value
var once_eq1Set sync.Once
func Get_eq1Set() gopurs_runtime.Value {
	once_eq1Set.Do(func() {
		cache_eq1Set = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqSet(dictEq_0), "eq")
}))
	})
	return cache_eq1Set
}

var cache_eq1Set__gopurs_runtime_Value_656709182 gopurs_runtime.Value
var once_eq1Set__gopurs_runtime_Value_656709182 sync.Once
func Get_eq1Set__gopurs_runtime_Value_656709182() gopurs_runtime.Value {
	once_eq1Set__gopurs_runtime_Value_656709182.Do(func() {
		cache_eq1Set__gopurs_runtime_Value_656709182 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqSet(dictEq_0), "eq")
}))
	})
	return cache_eq1Set__gopurs_runtime_Value_656709182
}

var cache_ord1Set gopurs_runtime.Value
var once_ord1Set sync.Once
func Get_ord1Set() gopurs_runtime.Value {
	once_ord1Set.Do(func() {
		cache_ord1Set = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Set()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_ordSet(dictOrd_0), "compare")
}))
	})
	return cache_ord1Set
}

var cache_ord1Set__gopurs_runtime_Value_1315498729 gopurs_runtime.Value
var once_ord1Set__gopurs_runtime_Value_1315498729 sync.Once
func Get_ord1Set__gopurs_runtime_Value_1315498729() gopurs_runtime.Value {
	once_ord1Set__gopurs_runtime_Value_1315498729.Do(func() {
		cache_ord1Set__gopurs_runtime_Value_1315498729 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Set()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_ordSet(dictOrd_0), "compare")
}))
	})
	return cache_ord1Set__gopurs_runtime_Value_1315498729
}

var cache_empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		cache_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}))}
	})
	return cache_empty
}

var cache_empty__ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value__2198260019 gopurs_runtime.Value
var once_empty__ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value__2198260019 sync.Once
func Get_empty__ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value__2198260019() gopurs_runtime.Value {
	once_empty__ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value__2198260019.Do(func() {
		cache_empty__ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value__2198260019 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}))}
	})
	return cache_empty__ptrData_Map_Internal_Constructor_Node_gopurs_runtime_Value__gopurs_runtime_Value__2198260019
}

var cache_empty__gopurs_runtime_Value_3279308360 gopurs_runtime.Value
var once_empty__gopurs_runtime_Value_3279308360 sync.Once
func Get_empty__gopurs_runtime_Value_3279308360() gopurs_runtime.Value {
	once_empty__gopurs_runtime_Value_3279308360.Do(func() {
		cache_empty__gopurs_runtime_Value_3279308360 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}))}
	})
	return cache_empty__gopurs_runtime_Value_3279308360
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_fromFoldable
}

var cache_go__map gopurs_runtime.Value
var once_go__map sync.Once
func Get_go__map() gopurs_runtime.Value {
	once_go__map.Do(func() {
		cache_go__map = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__map(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_go__map
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_mapMaybe
}

var cache_mapMaybe__gopurs_runtime_Value_2748927564 gopurs_runtime.Value
var once_mapMaybe__gopurs_runtime_Value_2748927564 sync.Once
func Get_mapMaybe__gopurs_runtime_Value_2748927564() gopurs_runtime.Value {
	once_mapMaybe__gopurs_runtime_Value_2748927564.Do(func() {
		cache_mapMaybe__gopurs_runtime_Value_2748927564 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__gopurs_runtime_Value_2748927564(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_mapMaybe__gopurs_runtime_Value_2748927564
}

var cache_monoidSet gopurs_runtime.Value
var once_monoidSet sync.Once
func Get_monoidSet() gopurs_runtime.Value {
	once_monoidSet.Do(func() {
		cache_monoidSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidSet(dictOrd_0_box)
})
	})
	return cache_monoidSet
}

var cache_unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		cache_unions = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unions(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_unions
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_difference
}

var cache_difference__gopurs_runtime_Value_833514008 gopurs_runtime.Value
var once_difference__gopurs_runtime_Value_833514008 sync.Once
func Get_difference__gopurs_runtime_Value_833514008() gopurs_runtime.Value {
	once_difference__gopurs_runtime_Value_833514008.Do(func() {
		cache_difference__gopurs_runtime_Value_833514008 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference__gopurs_runtime_Value_833514008(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_difference__gopurs_runtime_Value_833514008
}

var cache_subset gopurs_runtime.Value
var once_subset sync.Once
func Get_subset() gopurs_runtime.Value {
	once_subset.Do(func() {
		cache_subset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_subset(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_2_box)))
})
	})
	return cache_subset
}

var cache_subset__gopurs_runtime_Value_3164788426 gopurs_runtime.Value
var once_subset__gopurs_runtime_Value_3164788426 sync.Once
func Get_subset__gopurs_runtime_Value_3164788426() gopurs_runtime.Value {
	once_subset__gopurs_runtime_Value_3164788426.Do(func() {
		cache_subset__gopurs_runtime_Value_3164788426 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_subset__gopurs_runtime_Value_3164788426(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_2_box)))
})
	})
	return cache_subset__gopurs_runtime_Value_3164788426
}

var cache_properSubset gopurs_runtime.Value
var once_properSubset sync.Once
func Get_properSubset() gopurs_runtime.Value {
	once_properSubset.Do(func() {
		cache_properSubset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_properSubset(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_2_box)))
})
	})
	return cache_properSubset
}

var cache_delete gopurs_runtime.Value
var once_delete sync.Once
func Get_delete() gopurs_runtime.Value {
	once_delete.Do(func() {
		cache_delete = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_delete
}

var cache_delete__gopurs_runtime_Value_4217907800 gopurs_runtime.Value
var once_delete__gopurs_runtime_Value_4217907800 sync.Once
func Get_delete__gopurs_runtime_Value_4217907800() gopurs_runtime.Value {
	once_delete__gopurs_runtime_Value_4217907800.Do(func() {
		cache_delete__gopurs_runtime_Value_4217907800 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete__gopurs_runtime_Value_4217907800(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_delete__gopurs_runtime_Value_4217907800
}

var cache_checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		cache_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_checkValid(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_checkValid
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catMaybes(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_catMaybes
}

func Call_identity(x_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var x_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Set(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_union(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
})
}

func Call_toggle(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_alter(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, pkg_Data_Unit.Get_unit()})}
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_toMap(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_toMap__gopurs_runtime_Value_556171355(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_2))
})
}

func Call_toUnfoldable__gopurs_runtime_Value_871784487(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_2))
})
}

func Call_singleton(a_0_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, a_0, pkg_Data_Unit.Get_unit(), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})})})
}

func Call_singleton__gopurs_runtime_Value_3724491835(a_0_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, a_0, pkg_Data_Unit.Get_unit(), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})})})
}

func Call_showSet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showArray_1_0 := &pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))}
_ = showArray_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(fromFoldable "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(showArray_1_0.V0, gopurs_runtime.Apply2(pkg_Data_List.Get_toUnfoldable(), pkg_Data_Unfoldable.Get_unfoldableArray(), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), s_2))), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semigroupSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
}))
}

func Call_member(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_0 gopurs_runtime.Value
go__go_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_0:
for {
if false { continue go__go_2_0_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_2 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (uint32(v1_4_2.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_0
__t3 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_3
} else {

}
}
{
if (uint32(v1_4_2.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_0
__t3 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_3
} else {

}
}
{
if (uint32(v1_4_2.IntVal) == 902936544) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = gopurs_runtime.Bool((__t3.IntVal) != (0))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
}
}()
})
return go__go_2_0_0
}

func Call_intersection(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
})
}

func Call_intersection__gopurs_runtime_Value_833514008(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
})
}

func Call_insert(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, pkg_Data_Unit.Get_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_insert__gopurs_runtime_Value_4217907800(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, pkg_Data_Unit.Get_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_findMin(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_findMin__gopurs_runtime_Value_611847841(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_findMax(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_findMax__gopurs_runtime_Value_611847841(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_filter(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_filterKeys(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_eqSet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
eqMap_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, pkg_Data_Eq.Get_eqUnit()))
_ = eqMap_1_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eqMap_1_0.V0, v_2, v1_3).IntVal) != (0))
})
}))
}

func Call_ordSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordList_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), dictOrd_0))
_ = ordList_1_0
eqSet1_2_1 := Call_eqSet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqSet1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqSet1_2_1
}), gopurs_runtime.Func(func(s1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(ordList_1_0.V1, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), s1_3), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), s2_4)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictOrd_1_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)}, a_3, pkg_Data_Unit.Get_unit(), m_2)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})
}

func Call_go__map(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableSet(), "foldl"), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Apply(f_1, a_3), pkg_Data_Unit.Get_unit(), m_2)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})
}

func Call_mapMaybe(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableSet(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(f_1, a_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 930809136 && __local_var_4_0.UnsafePtr == nil) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 930809136 && __local_var_4_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, pkg_Data_Unit.Get_unit(), acc_3)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})
}

func Call_mapMaybe__gopurs_runtime_Value_2748927564(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableSet(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(f_1, a_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 930809136 && __local_var_4_0.UnsafePtr == nil) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 930809136 && __local_var_4_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, pkg_Data_Unit.Get_unit(), acc_3)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})
}

func Call_monoidSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
semigroupSet1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
}))
_ = semigroupSet1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSet1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil}))})
}

func Call_unions(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictOrd_1_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
compare_2_0 := dictOrd_1.V1
_ = compare_2_0
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_2_0, pkg_Data_Function.Get_go__const(), m1_3, m2_4)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})
}

func Call_difference(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, m1_2, m2_3)))}
})
})
}

func Call_difference__gopurs_runtime_Value_833514008(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, m1_2, m2_3)))}
})
})
}

func Call_subset(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], s1_1_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
var __t0 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), dictOrd_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_subset__gopurs_runtime_Value_3164788426(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], s1_1_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
var __t0 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), dictOrd_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_properSubset(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], s1_1_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}.UnsafePtr).V1).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}.UnsafePtr).V1).IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), dictOrd_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 324739070 && __t_tag_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((__t0.IntVal) == (__t1.IntVal)), gopurs_runtime.Bool(false)), __t2).IntVal) != (0)
}

func Call_delete(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_delete__gopurs_runtime_Value_4217907800(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_checkValid(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_checkValid(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_catMaybes(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_mapMaybe(dictOrd_0, Get_identity())
}


