package Data_Map

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Set "gopurs/output/Data.Set"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_SemigroupMap gopurs_runtime.Value
var once_SemigroupMap sync.Once
func Get_SemigroupMap() gopurs_runtime.Value {
	once_SemigroupMap.Do(func() {
		cache_SemigroupMap = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_SemigroupMap(x_0_box)
})
	})
	return cache_SemigroupMap
}

var cache_traversableWithIndexSemigroupMap gopurs_runtime.Value
var once_traversableWithIndexSemigroupMap sync.Once
func Get_traversableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_traversableWithIndexSemigroupMap.Do(func() {
		cache_traversableWithIndexSemigroupMap = pkg_Data_Map_Internal.Get_traversableWithIndexMap()
	})
	return cache_traversableWithIndexSemigroupMap
}

var cache_traversableSemigroupMap gopurs_runtime.Value
var once_traversableSemigroupMap sync.Once
func Get_traversableSemigroupMap() gopurs_runtime.Value {
	once_traversableSemigroupMap.Do(func() {
		cache_traversableSemigroupMap = pkg_Data_Map_Internal.Get_traversableMap()
	})
	return cache_traversableSemigroupMap
}

var cache_showSemigroupMap gopurs_runtime.Value
var once_showSemigroupMap sync.Once
func Get_showSemigroupMap() gopurs_runtime.Value {
	once_showSemigroupMap.Do(func() {
		cache_showSemigroupMap = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showSemigroupMap(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showSemigroupMap
}

var cache_semigroupSemigroupMap gopurs_runtime.Value
var once_semigroupSemigroupMap sync.Once
func Get_semigroupSemigroupMap() gopurs_runtime.Value {
	once_semigroupSemigroupMap.Do(func() {
		cache_semigroupSemigroupMap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupSemigroupMap(dictOrd_0_box, dictSemigroup_1_box)
})
	})
	return cache_semigroupSemigroupMap
}

var cache_plusSemigroupMap gopurs_runtime.Value
var once_plusSemigroupMap sync.Once
func Get_plusSemigroupMap() gopurs_runtime.Value {
	once_plusSemigroupMap.Do(func() {
		cache_plusSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusSemigroupMap(dictOrd_0_box)
})
	})
	return cache_plusSemigroupMap
}

var cache_ordSemigroupMap gopurs_runtime.Value
var once_ordSemigroupMap sync.Once
func Get_ordSemigroupMap() gopurs_runtime.Value {
	once_ordSemigroupMap.Do(func() {
		cache_ordSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordSemigroupMap(dictOrd_0_box)
})
	})
	return cache_ordSemigroupMap
}

var cache_ord1SemigroupMap gopurs_runtime.Value
var once_ord1SemigroupMap sync.Once
func Get_ord1SemigroupMap() gopurs_runtime.Value {
	once_ord1SemigroupMap.Do(func() {
		cache_ord1SemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1SemigroupMap(dictOrd_0_box)
})
	})
	return cache_ord1SemigroupMap
}

var cache_newtypeSemigroupMap gopurs_runtime.Value
var once_newtypeSemigroupMap sync.Once
func Get_newtypeSemigroupMap() gopurs_runtime.Value {
	once_newtypeSemigroupMap.Do(func() {
		cache_newtypeSemigroupMap = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeSemigroupMap
}

var cache_monoidSemigroupMap gopurs_runtime.Value
var once_monoidSemigroupMap sync.Once
func Get_monoidSemigroupMap() gopurs_runtime.Value {
	once_monoidSemigroupMap.Do(func() {
		cache_monoidSemigroupMap = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidSemigroupMap(dictOrd_0_box, dictSemigroup_1_box)
})
	})
	return cache_monoidSemigroupMap
}

var cache_keys gopurs_runtime.Value
var once_keys sync.Once
func Get_keys() gopurs_runtime.Value {
	once_keys.Do(func() {
		cache_keys = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Map_Internal.Get_functorMap(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, x_1)
})
}()
	})
	return cache_keys
}

var cache_functorWithIndexSemigroupMap gopurs_runtime.Value
var once_functorWithIndexSemigroupMap sync.Once
func Get_functorWithIndexSemigroupMap() gopurs_runtime.Value {
	once_functorWithIndexSemigroupMap.Do(func() {
		cache_functorWithIndexSemigroupMap = pkg_Data_Map_Internal.Get_functorWithIndexMap()
	})
	return cache_functorWithIndexSemigroupMap
}

var cache_functorSemigroupMap gopurs_runtime.Value
var once_functorSemigroupMap sync.Once
func Get_functorSemigroupMap() gopurs_runtime.Value {
	once_functorSemigroupMap.Do(func() {
		cache_functorSemigroupMap = pkg_Data_Map_Internal.Get_functorMap()
	})
	return cache_functorSemigroupMap
}

var cache_foldableWithIndexSemigroupMap gopurs_runtime.Value
var once_foldableWithIndexSemigroupMap sync.Once
func Get_foldableWithIndexSemigroupMap() gopurs_runtime.Value {
	once_foldableWithIndexSemigroupMap.Do(func() {
		cache_foldableWithIndexSemigroupMap = pkg_Data_Map_Internal.Get_foldableWithIndexMap()
	})
	return cache_foldableWithIndexSemigroupMap
}

var cache_foldableSemigroupMap gopurs_runtime.Value
var once_foldableSemigroupMap sync.Once
func Get_foldableSemigroupMap() gopurs_runtime.Value {
	once_foldableSemigroupMap.Do(func() {
		cache_foldableSemigroupMap = pkg_Data_Map_Internal.Get_foldableMap()
	})
	return cache_foldableSemigroupMap
}

var cache_eqSemigroupMap gopurs_runtime.Value
var once_eqSemigroupMap sync.Once
func Get_eqSemigroupMap() gopurs_runtime.Value {
	once_eqSemigroupMap.Do(func() {
		cache_eqSemigroupMap = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqSemigroupMap(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqSemigroupMap
}

var cache_eq1SemigroupMap gopurs_runtime.Value
var once_eq1SemigroupMap sync.Once
func Get_eq1SemigroupMap() gopurs_runtime.Value {
	once_eq1SemigroupMap.Do(func() {
		cache_eq1SemigroupMap = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1SemigroupMap(dictEq_0_box)
})
	})
	return cache_eq1SemigroupMap
}

var cache_bindSemigroupMap gopurs_runtime.Value
var once_bindSemigroupMap sync.Once
func Get_bindSemigroupMap() gopurs_runtime.Value {
	once_bindSemigroupMap.Do(func() {
		cache_bindSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindSemigroupMap(dictOrd_0_box)
})
	})
	return cache_bindSemigroupMap
}

var cache_applySemigroupMap gopurs_runtime.Value
var once_applySemigroupMap sync.Once
func Get_applySemigroupMap() gopurs_runtime.Value {
	once_applySemigroupMap.Do(func() {
		cache_applySemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applySemigroupMap(dictOrd_0_box)
})
	})
	return cache_applySemigroupMap
}

var cache_altSemigroupMap gopurs_runtime.Value
var once_altSemigroupMap sync.Once
func Get_altSemigroupMap() gopurs_runtime.Value {
	once_altSemigroupMap.Do(func() {
		cache_altSemigroupMap = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altSemigroupMap(dictOrd_0_box)
})
	})
	return cache_altSemigroupMap
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

var cache_pure__4233214992 gopurs_runtime.Value
var once_pure__4233214992 sync.Once
func Get_pure__4233214992() gopurs_runtime.Value {
	once_pure__4233214992.Do(func() {
		cache_pure__4233214992 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__4233214992(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__4233214992
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

var cache_apply__3783667596 gopurs_runtime.Value
var once_apply__3783667596 sync.Once
func Get_apply__3783667596() gopurs_runtime.Value {
	once_apply__3783667596.Do(func() {
		cache_apply__3783667596 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__3783667596(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__3783667596
}

var cache_apply__986161100 gopurs_runtime.Value
var once_apply__986161100 sync.Once
func Get_apply__986161100() gopurs_runtime.Value {
	once_apply__986161100.Do(func() {
		cache_apply__986161100 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__986161100(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__986161100
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

var cache_map__3658136916 gopurs_runtime.Value
var once_map__3658136916 sync.Once
func Get_map__3658136916() gopurs_runtime.Value {
	once_map__3658136916.Do(func() {
		cache_map__3658136916 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3658136916(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3658136916
}

var cache_empty__1818220131 gopurs_runtime.Value
var once_empty__1818220131 sync.Once
func Get_empty__1818220131() gopurs_runtime.Value {
	once_empty__1818220131.Do(func() {
		cache_empty__1818220131 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_empty__1818220131
}

var cache_foldableMap__373570208 gopurs_runtime.Value
var once_foldableMap__373570208 sync.Once
func Get_foldableMap__373570208() gopurs_runtime.Value {
	once_foldableMap__373570208.Do(func() {
		cache_foldableMap__373570208 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_0 gopurs_runtime.Value
_ = go__go_3_1_0
go__go_3_1_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
return go__go_3_1_0
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_1 gopurs_runtime.Value
go__go_2_3_1 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_1:
for {
if false { continue go__go_2_3_1 }
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __t4 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t4 = __local_var_3
goto end_branch_4
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_1, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_1, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_3_1, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_2 gopurs_runtime.Value
_ = go__go_2_5_2
go__go_2_5_2 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = __local_var_4
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_5_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableMap__373570208
}

var cache_foldableWithIndexMap__1966365627 gopurs_runtime.Value
var once_foldableWithIndexMap__1966365627 sync.Once
func Get_foldableWithIndexMap__1966365627() gopurs_runtime.Value {
	once_foldableWithIndexMap__1966365627.Do(func() {
		cache_foldableWithIndexMap__1966365627 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_3 gopurs_runtime.Value
_ = go__go_3_1_3
go__go_3_1_3 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
return go__go_3_1_3
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_4 gopurs_runtime.Value
go__go_2_3_4 = gopurs_runtime.Func(func(__local_var_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var __local_var_3_loop gopurs_runtime.Value = __local_var_3_loop_val
var __local_var_4_loop gopurs_runtime.Value = __local_var_4_loop_val
go__go_2_3_4:
for {
if false { continue go__go_2_3_4 }
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __local_var_4 gopurs_runtime.Value = __local_var_4_loop
_ = __local_var_4
var __t4 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t4 = __local_var_3
goto end_branch_4
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_4, gopurs_runtime.Apply3(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_4, __local_var_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
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
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_3_4, z_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_5 gopurs_runtime.Value
_ = go__go_2_5_5
go__go_2_5_5 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = __local_var_4
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_5_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, z_1)
})
})
}))
	})
	return cache_foldableWithIndexMap__1966365627
}

var cache_functorMap__2501170929 gopurs_runtime.Value
var once_functorMap__2501170929 sync.Once
func Get_functorMap__2501170929() gopurs_runtime.Value {
	once_functorMap__2501170929.Do(func() {
		cache_functorMap__2501170929 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_6 gopurs_runtime.Value
_ = go__go_1_0_6
go__go_1_0_6 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
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
return go__go_1_0_6
}))
	})
	return cache_functorMap__2501170929
}

var cache_functorWithIndexMap__3138419015 gopurs_runtime.Value
var once_functorWithIndexMap__3138419015 sync.Once
func Get_functorWithIndexMap__3138419015() gopurs_runtime.Value {
	once_functorWithIndexMap__3138419015.Do(func() {
		cache_functorWithIndexMap__3138419015 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0_7 gopurs_runtime.Value
_ = go__go_1_0_7
go__go_1_0_7 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, gopurs_runtime.Apply2(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)})), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_1_0_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}))})}
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
return go__go_1_0_7
}))
	})
	return cache_functorWithIndexMap__3138419015
}

var cache_singleton__2450056090 gopurs_runtime.Value
var once_singleton__2450056090 sync.Once
func Get_singleton__2450056090() gopurs_runtime.Value {
	once_singleton__2450056090.Do(func() {
		cache_singleton__2450056090 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__2450056090(k_0_box, v_1_box))}
})
	})
	return cache_singleton__2450056090
}

var cache_traversableMap__1002539403 gopurs_runtime.Value
var once_traversableMap__1002539403 sync.Once
func Get_traversableMap__1002539403() gopurs_runtime.Value {
	once_traversableMap__1002539403.Do(func() {
		cache_traversableMap__1002539403 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Map_Internal.Get_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Map_Internal.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_8 gopurs_runtime.Value
_ = go__go_4_2_8
go__go_4_2_8 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return go__go_4_2_8
})
}))
	})
	return cache_traversableMap__1002539403
}

var cache_traversableMap__2256206635 gopurs_runtime.Value
var once_traversableMap__2256206635 sync.Once
func Get_traversableMap__2256206635() gopurs_runtime.Value {
	once_traversableMap__2256206635.Do(func() {
		cache_traversableMap__2256206635 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_foldableMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Map_Internal.Get_traversableMap(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Map_Internal.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_9 gopurs_runtime.Value
_ = go__go_4_2_9
go__go_4_2_9 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return go__go_4_2_9
})
}))
	})
	return cache_traversableMap__2256206635
}

var cache_traversableWithIndexMap__3269014446 gopurs_runtime.Value
var once_traversableWithIndexMap__3269014446 sync.Once
func Get_traversableWithIndexMap__3269014446() gopurs_runtime.Value {
	once_traversableWithIndexMap__3269014446.Do(func() {
		cache_traversableWithIndexMap__3269014446 = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_foldableWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorWithIndexMap()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_traversableMap()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_10 gopurs_runtime.Value
_ = go__go_4_2_10
go__go_4_2_10 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply2(f_3, __local_var_7_4, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return go__go_4_2_10
})
}))
	})
	return cache_traversableWithIndexMap__3269014446
}

var cache_unionWith__2507192643 gopurs_runtime.Value
var once_unionWith__2507192643 sync.Once
func Get_unionWith__2507192643() gopurs_runtime.Value {
	once_unionWith__2507192643.Do(func() {
		cache_unionWith__2507192643 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionWith__2507192643(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_unionWith__2507192643
}

var cache_unsafeBalancedNode__1305301638 gopurs_runtime.Value
var once_unsafeBalancedNode__1305301638 sync.Once
func Get_unsafeBalancedNode__1305301638() gopurs_runtime.Value {
	once_unsafeBalancedNode__1305301638.Do(func() {
		cache_unsafeBalancedNode__1305301638 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeBalancedNode__1305301638(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeBalancedNode__1305301638
}

var cache_unsafeNode__1305301638 gopurs_runtime.Value
var once_unsafeNode__1305301638 sync.Once
func Get_unsafeNode__1305301638() gopurs_runtime.Value {
	once_unsafeNode__1305301638.Do(func() {
		cache_unsafeNode__1305301638 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeNode__1305301638(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeNode__1305301638
}

var cache_unsafeSplit__4154869695 gopurs_runtime.Value
var once_unsafeSplit__4154869695 sync.Once
func Get_unsafeSplit__4154869695() gopurs_runtime.Value {
	once_unsafeSplit__4154869695.Do(func() {
		cache_unsafeSplit__4154869695 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplit__4154869695(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_unsafeSplit__4154869695
}

var cache_unsafeUnionWith__4109280494 gopurs_runtime.Value
var once_unsafeUnionWith__4109280494 sync.Once
func Get_unsafeUnionWith__4109280494() gopurs_runtime.Value {
	once_unsafeUnionWith__4109280494.Do(func() {
		cache_unsafeUnionWith__4109280494 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeUnionWith__4109280494(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeUnionWith__4109280494
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
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

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
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

var cache_fromMap__556171355 gopurs_runtime.Value
var once_fromMap__556171355 sync.Once
func Get_fromMap__556171355() gopurs_runtime.Value {
	once_fromMap__556171355.Do(func() {
		cache_fromMap__556171355 = pkg_Data_Set.Get_Set()
	})
	return cache_fromMap__556171355
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

var cache_traverse__3640625269 gopurs_runtime.Value
var once_traverse__3640625269 sync.Once
func Get_traverse__3640625269() gopurs_runtime.Value {
	once_traverse__3640625269.Do(func() {
		cache_traverse__3640625269 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__3640625269(dictApplicative_0_box)
})
	})
	return cache_traverse__3640625269
}

func Call_SemigroupMap(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showSemigroupMap(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_showMap(), dictShow_0, dictShow1_1)
}

func Call_semigroupSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
append_2_0 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_4))})))}
})
}))
}

func Call_plusSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_plusMap(), dictOrd_0)
}

func Call_ordSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordMap_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_ordMap(), dictOrd_0)
_ = ordMap_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordMap_1_0, dictOrd1_2)
})
}

func Call_ord1SemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_ord1Map(), dictOrd_0)
}

func Call_monoidSemigroupMap(dictOrd_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
append_2_1 := gopurs_runtime.RecordGet(dictSemigroup_1, "append")
_ = append_2_1
semigroupSemigroupMap2_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), append_2_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_4))})))}
})
}))
_ = semigroupSemigroupMap2_2_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSemigroupMap2_2_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_eqSemigroupMap(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, dictEq1_1)
}

func Call_eq1SemigroupMap(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, dictEq1_1), "eq")
}))
}

func Call_bindSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_bindMap(), dictOrd_0)
}

func Call_applySemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Map_Internal.Get_identity2(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
}

func Call_altSemigroupMap(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_functorMap()
}), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}))
}

func Call_pure__189931222(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__4233214992(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__4203183626(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__3783667596(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__986161100(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2665381605(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3658136916(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_singleton__2450056090(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_unionWith__2507192643(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_unsafeBalancedNode__1305301638(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t27 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t6 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_6
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t5 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_4 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr != nil) {

var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t_and_4 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t3.IntVal))).IntVal) != (0)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_6:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
goto end_branch_27
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t26 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t19 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t12 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_11 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 324739070 && __t_tag_7.UnsafePtr != nil) {

var __t10 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 324739070 && __t_tag_8.UnsafePtr == nil) {
__t10 = gopurs_runtime.Int(0)
goto end_branch_10
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 324739070 && __t_tag_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t_and_11 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t10.IntVal))).IntVal) != (0)
}
if __t_and_11 {
__t12 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_12:
__t19 = __t12
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t18 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_17 bool = false
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {

var __t16 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 324739070 && __t_tag_14.UnsafePtr == nil) {
__t16 = gopurs_runtime.Int(0)
goto end_branch_16
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 324739070 && __t_tag_15.UnsafePtr != nil) {
__t16 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t_and_17 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t16.IntVal), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_17 {
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t25 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr != nil) {

var __t23 gopurs_runtime.Value
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr == nil) {
__t23 = gopurs_runtime.Int(0)
goto end_branch_23
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_22.Type == 9 && __t_tag_22.IntVal == 324739070 && __t_tag_22.UnsafePtr != nil) {
__t23 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t_and_24 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t23.IntVal), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_26:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t26)}
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t27))}
}

func Call_unsafeNode__1305301638(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t3 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)).IntVal, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_3
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)).IntVal, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t1 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0))).IntVal) != (0) {
__t1 = (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int(__t1)).IntVal, Call_add__560788792(gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
}

func Call_unsafeSplit__4154869695(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
v1_4_1 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
v1_4_2 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_unsafeUnionWith__4109280494(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v_4_0 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeSplit(), __local_var_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
l_prime_5_1 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
r_prime_6_2 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
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

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
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

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__3640625269(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_11 gopurs_runtime.Value
_ = go__go_4_2_11
go__go_4_2_11 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
__local_var_6_3 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_3
__local_var_7_4 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V2
_ = __local_var_7_4
__local_var_8_5 := (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_8_5
__t6 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5, __local_var_7_4, v_prime_10, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_9), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_11)})}
})
})
}), gopurs_runtime.Apply(go__go_4_2_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V5)}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return go__go_4_2_11
})
}


